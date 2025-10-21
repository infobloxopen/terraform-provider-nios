package security

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForCertificateAuthservice = "auto_populate_login,ca_certificates,comment,disabled,enable_password_request,enable_remote_lookup,max_retries,name,ocsp_check,ocsp_responders,recovery_interval,remote_lookup_service,remote_lookup_username,response_timeout,trust_model,user_match_type"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &CertificateAuthserviceResource{}
var _ resource.ResourceWithImportState = &CertificateAuthserviceResource{}

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

	apiRes, _, err := r.client.SecurityAPI.
		CertificateAuthserviceAPI.
		Create(ctx).
		CertificateAuthservice(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForCertificateAuthservice).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create CertificateAuthservice, got error: %s", err))
		return
	}

	res := apiRes.CreateCertificateAuthserviceResponseAsObject.GetResult()

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
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForCertificateAuthservice).
		ReturnAsObject(1).
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

	// Process OCSP responders
	if !r.processOcspResponders(ctx, &data, &resp.Diagnostics) {
		return
	}
	diags = req.State.GetAttribute(ctx, path.Root("ref"), &data.Ref)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	apiRes, _, err := r.client.SecurityAPI.
		CertificateAuthserviceAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		CertificateAuthservice(*data.Expand(ctx, &resp.Diagnostics)).
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
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
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
	resource.ImportStatePassthroughID(ctx, path.Root("ref"), req, resp)
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
	}
}
