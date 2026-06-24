package security

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/terraform-provider-nios/internal/config"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForCertificateAuthservice = "auto_populate_login,ca_certificates,comment,disabled,enable_password_request,enable_remote_lookup,max_retries,name,ocsp_check,ocsp_responders,recovery_interval,remote_lookup_service,remote_lookup_username,response_timeout,trust_model,user_match_type"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &CertificateAuthserviceResource{}
var _ resource.ResourceWithImportState = &CertificateAuthserviceResource{}

var _ resource.ResourceWithModifyPlan = &CertificateAuthserviceResource{}

func NewCertificateAuthserviceResource() resource.Resource {
	return &CertificateAuthserviceResource{}
}

// CertificateAuthserviceResource defines the resource implementation.
type CertificateAuthserviceResource struct {
	client *niosclient.APIClient
}

func (r *CertificateAuthserviceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "security_certificate_authservice"
}

func (r *CertificateAuthserviceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a Certificate Authentication Service.",
		Attributes:          CertificateAuthserviceResourceSchemaAttributes,
	}
}

func (r *CertificateAuthserviceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*niosclient.APIClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *niosclient.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

type certAuthHashState struct {
	Password string `json:"password_hash"`
}

func (r *CertificateAuthserviceResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return
	}

	var statePwdVersion types.Int64
	var planPassword types.String

	// Normalize stateRev if null (e.g., first apply)
	curRev := int64(0)

	if !req.State.Raw.IsNull() && req.State.Raw.IsKnown() {
		resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("password_version"), &statePwdVersion)...)
		if resp.Diagnostics.HasError() {
			return
		}
		if !statePwdVersion.IsNull() && !statePwdVersion.IsUnknown() {
			curRev = statePwdVersion.ValueInt64()
		}
	}
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("remote_lookup_password"), &planPassword)...)
	if resp.Diagnostics.HasError() {
		return
	}

	computeNewHash := !planPassword.IsNull() && !planPassword.IsUnknown()

	prevHashes := certAuthHashState{}
	plannedHashes := certAuthHashState{}

	if computeNewHash {

		var prev struct {
			Algo string `json:"algo"`
			Hash string `json:"hash"`
		}

		if b, diags := req.Private.GetKey(ctx, "password_hash"); diags != nil {
			resp.Diagnostics.Append(diags...)
		} else if b != nil {
			if err := json.Unmarshal(b, &prev); err != nil {
				// Older buggy format: ignore and treat as different
				prev.Hash = ""
			}
		}
		var plannedHash string

		if prev.Hash != "" {
			// Best-effort parse; if this fails, treat prev.Hash as a legacy value and
			// leave prevHashes at its zero value so that we will recompute as needed.
			_ = json.Unmarshal([]byte(prev.Hash), &prevHashes)
		}

		if !planPassword.IsUnknown() {
			if planPassword.IsNull() {
				plannedHashes.Password = ""
			} else {
				h := sha256.New()
				h.Write([]byte(planPassword.ValueString()))
				plannedHashes.Password = hex.EncodeToString(h.Sum(nil))
			}
		}
		if data, err := json.Marshal(plannedHashes); err == nil {
			plannedHash = string(data)
		}

		if plannedHashes.Password != "" && plannedHashes.Password != prevHashes.Password {
			// Increment version and store new hash if password modified
			newRev := types.Int64Value(curRev + 1)
			resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("password_version"), newRev)...)

			val := map[string]string{"algo": "sha256", "hash": plannedHash}
			b, err := json.Marshal(val)
			if err != nil {
				resp.Diagnostics.AddError("error marshalling password hash", err.Error())
				return
			}
			resp.Diagnostics.Append(resp.Private.SetKey(ctx, "password_hash", b)...)
		} else {
			resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("password_version"), curRev)...)
		}
	}

}

func (r *CertificateAuthserviceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	//var diags diag.Diagnostics
	var data CertificateAuthserviceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Process OCSP responders
	if !r.processOcspResponders(ctx, &data, &resp.Diagnostics) {
		return
	}

	payload := data.Expand(ctx, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	passwordVersion := types.Int64Value(0)
	var password types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("remote_lookup_password"), &password)...)

	secretData := certAuthHashState{}

	if !password.IsNull() && !password.IsUnknown() {

		payload.RemoteLookupPassword = password.ValueStringPointer()
		passwordVersion = types.Int64Value(1)
		h := sha256.New()
		h.Write([]byte(password.ValueString()))
		secretData.Password = hex.EncodeToString(h.Sum(nil))

		secretDataJSON, _ := json.Marshal(secretData)
		val := map[string]string{"algo": "sha256", "hash": string(secretDataJSON)}
		hashedPassword, err := json.Marshal(val)
		if err != nil {
			resp.Diagnostics.AddError("Private State Marshal Error", err.Error())
			return
		}
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, "password_hash", hashedPassword)...)
	}

	apiRes, _, err := r.client.SecurityAPI.
		CertificateAuthserviceAPI.
		Create(ctx).
		CertificateAuthservice(*payload).
		ReturnFieldsPlus(readableAttributesForCertificateAuthservice).
		ReturnAsObject(1).
		Execute()

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create CertificateAuthservice, got error: %s", err))
		return
	}

	res := apiRes.CreateCertificateAuthserviceResponseAsObject.GetResult()
	data.PasswordVersion = passwordVersion

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CertificateAuthserviceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	//var diags diag.Diagnostics
	var data CertificateAuthserviceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.SecurityAPI.
		CertificateAuthserviceAPI.
		Read(ctx, utils.ResolveIdentifier(data.Uuid, data.Ref)).
		ReturnFieldsPlus(readableAttributesForCertificateAuthservice).
		ReturnAsObject(1).
		ProxySearch(config.GetProxySearch()).
		Execute()

	//remove from the state if not found
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			// Resource no longer exists, remove from state
			resp.State.RemoveResource(ctx)
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read CertificateAuthservice, got error: %s", err))
		return
	}

	res := apiRes.GetCertificateAuthserviceResponseObjectAsResult.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CertificateAuthserviceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data CertificateAuthserviceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	payload := data.Expand(ctx, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Process OCSP responders
	if !r.processOcspResponders(ctx, &data, &resp.Diagnostics) {
		return
	}
	diags = req.State.GetAttribute(ctx, path.Root("ref"), &data.Ref)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	var password types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("remote_lookup_password"), &password)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if !password.IsNull() && !password.IsUnknown() {
		payload.RemoteLookupPassword = password.ValueStringPointer()
	}

	diags = req.State.GetAttribute(ctx, path.Root("uuid"), &data.Uuid)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	apiRes, _, err := r.client.SecurityAPI.
		CertificateAuthserviceAPI.
		Update(ctx, utils.ResolveIdentifier(data.Uuid, data.Ref)).
		CertificateAuthservice(*payload).
		ReturnFieldsPlus(readableAttributesForCertificateAuthservice).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update CertificateAuthservice, got error: %s", err))
		return
	}

	res := apiRes.UpdateCertificateAuthserviceResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CertificateAuthserviceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CertificateAuthserviceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.SecurityAPI.
		CertificateAuthserviceAPI.
		Delete(ctx, utils.ResolveIdentifier(data.Uuid, data.Ref)).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete CertificateAuthservice, got error: %s", err))
		return
	}
}

func (r *CertificateAuthserviceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("uuid"), req, resp)
}

// processOcspResponders processes certificate files in OCSP responders list
func (r *CertificateAuthserviceResource) processOcspResponders(
	ctx context.Context,
	data *CertificateAuthserviceModel,
	diag *diag.Diagnostics,
) bool {
	if data.OcspResponders.IsNull() || data.OcspResponders.IsUnknown() {
		return true
	}

	baseUrl := r.client.SecurityAPI.Cfg.NIOSHostURL
	username := r.client.SecurityAPI.Cfg.NIOSUsername
	password := r.client.SecurityAPI.Cfg.NIOSPassword

	var ocspResponders []CertificateAuthserviceOcspRespondersModel
	diagResult := data.OcspResponders.ElementsAs(ctx, &ocspResponders, false)
	diag.Append(diagResult...)
	if diag.HasError() {
		return false
	}

	for i, ocspResponder := range ocspResponders {
		if !ocspResponder.CertificateFilePath.IsNull() && !ocspResponder.CertificateFilePath.IsUnknown() {
			filePath := ocspResponder.CertificateFilePath.ValueString()
			token, err := utils.UploadFileWithToken(ctx, baseUrl, filePath, username, password)
			if err != nil {
				diag.AddError(
					"Client Error",
					fmt.Sprintf("Unable to process certificate file %s, got error: %s", filePath, err),
				)
				return false
			}
			ocspResponders[i].CertificateToken = types.StringValue(token)
		}
	}

	listValue, diagResult := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: CertificateAuthserviceOcspRespondersAttrTypes}, ocspResponders)
	diag.Append(diagResult...)
	if diag.HasError() {
		return false
	}

	data.OcspResponders = listValue
	return true
}

func (r *CertificateAuthserviceResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data CertificateAuthserviceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	ocspCheck := data.OcspCheck
	ocspResponders := data.OcspResponders

	isOcspCheckValid := !ocspCheck.IsNull() && !ocspCheck.IsUnknown()
	isManualCheck := ocspCheck.ValueString() == "MANUAL" || ocspCheck.ValueString() == "AIA_AND_MANUAL"

	// Handle when ocsp_check is valid and set to MANUAL or AIA_AND_MANUAL
	if (isOcspCheckValid && isManualCheck) || !isOcspCheckValid {
		if ocspResponders.IsNull() || ocspResponders.IsUnknown() {
			resp.Diagnostics.AddError(
				"Invalid Configuration",
				"At least one `ocsp_responders` must be specified when `ocsp_check` is set to `MANUAL` or `AIA_AND_MANUAL`, else set the ocsp_check to 'DISABLED'.",
			)
		}
	}

	// Check if remote lookup is enabled and validate required fields
	isRemoteLookupEnabled := !data.EnableRemoteLookup.IsNull() && !data.EnableRemoteLookup.IsUnknown() && data.EnableRemoteLookup.ValueBool()
	missingService := data.RemoteLookupService.IsNull() || data.RemoteLookupService.IsUnknown()
	missingUsername := data.RemoteLookupUsername.IsNull() || data.RemoteLookupUsername.IsUnknown()
	missingPassword := data.RemoteLookupPassword.IsNull() || data.RemoteLookupPassword.IsUnknown()

	if isRemoteLookupEnabled {
		// Validate required fields for remote lookup
		if missingService || missingUsername || missingPassword {
			resp.Diagnostics.AddError(
				"Invalid Configuration",
				"When `enable_remote_lookup` is set to `true`, all fields `remote_lookup_service`, `remote_lookup_username`, and `remote_lookup_password` must be provided.",
			)
		}

		// Validate enable_password_request setting
		if data.EnablePasswordRequest.IsNull() || data.EnablePasswordRequest.IsUnknown() || data.EnablePasswordRequest.ValueBool() {
			resp.Diagnostics.AddError(
				"Invalid Configuration",
				"When `enable_remote_lookup` is set to `true`, `enable_password_request` must be set to `false`.",
			)
		}

		if !data.UserMatchType.IsNull() && !data.UserMatchType.IsUnknown() && data.UserMatchType.ValueString() != "AUTO_MATCH" {
			resp.Diagnostics.AddError(
				"Invalid Configuration",
				"`user_match_type` must be set to \"AUTO_MATCH\" to use remote lookup services.",
			)
		}
	}
}
