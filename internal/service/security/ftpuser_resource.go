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
	"github.com/infobloxopen/infoblox-nios-go-client/security"
	"github.com/infobloxopen/terraform-provider-nios/internal/config"
	"github.com/infobloxopen/terraform-provider-nios/internal/retry"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForFtpuser = "extattrs,home_dir,permission,username"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &FtpuserResource{}
var _ resource.ResourceWithImportState = &FtpuserResource{}
var _ resource.ResourceWithModifyPlan = &FtpuserResource{}

func NewFtpuserResource() resource.Resource {
	return &FtpuserResource{}
}

// FtpuserResource defines the resource implementation.
type FtpuserResource struct {
	client *niosclient.APIClient
}

func (r *FtpuserResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "security_ftpuser"
}

func (r *FtpuserResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages an FTP user.",
		Attributes:          FtpuserResourceSchemaAttributes,
	}
}

func (r *FtpuserResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *FtpuserResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return
	}

	var stateRev types.Int64
	var planSecret types.String

	curRev := int64(0)
	if !req.State.Raw.IsNull() && req.State.Raw.IsKnown() {
		resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("password_version"), &stateRev)...)
		if resp.Diagnostics.HasError() {
			return
		}
		if !stateRev.IsNull() && !stateRev.IsUnknown() {
			curRev = stateRev.ValueInt64()
		}
	}

	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("password"), &planSecret)...)
	if resp.Diagnostics.HasError() {
		return
	}

	computeNewHash := !planSecret.IsNull() && !planSecret.IsUnknown()
	prevHashes := secretsHashState{}
	plannedHashes := secretsHashState{}

	if computeNewHash {
		var prev struct {
			Algo string `json:"algo"`
			Hash string `json:"hash"`
		}

		if b, diags := req.Private.GetKey(ctx, "password_hash"); diags != nil {
			resp.Diagnostics.Append(diags...)
		} else if b != nil {
			if err := json.Unmarshal(b, &prev); err != nil {
				prev.Hash = ""
			}
		}

		var plannedHash string
		if prev.Hash != "" {
			_ = json.Unmarshal([]byte(prev.Hash), &prevHashes)
		}

		if !planSecret.IsUnknown() {
			if planSecret.IsNull() {
				plannedHashes.Password = ""
			} else {
				h := sha256.New()
				h.Write([]byte(planSecret.ValueString()))
				plannedHashes.Password = hex.EncodeToString(h.Sum(nil))
			}
		}

		if data, err := json.Marshal(plannedHashes); err == nil {
			plannedHash = string(data)
		}

		if plannedHashes.Password != "" && plannedHashes.Password != prevHashes.Password {
			newRev := types.Int64Value(curRev + 1)
			resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("password_version"), newRev)...)

			val := map[string]string{"algo": "sha256", "hash": plannedHash}
			b, err := json.Marshal(val)
			if err != nil {
				resp.Diagnostics.AddError("Private State Marshal Error", err.Error())
				return
			}
			resp.Diagnostics.Append(resp.Private.SetKey(ctx, "password_hash", b)...)
		} else {
			resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("password_version"), curRev)...)
		}
	}
}

func (r *FtpuserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var diags diag.Diagnostics
	var data FtpuserModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Add internal ID exists in the Extensible Attributes if not already present
	data.ExtAttrs, diags = AddInternalIDToExtAttrs(ctx, data.ExtAttrs, diags)
	if diags.HasError() {
		return
	}

	payload := data.Expand(ctx, &resp.Diagnostics, true)
	if resp.Diagnostics.HasError() {
		return
	}

	passwordVersion := types.Int64Value(0)
	var password types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("password"), &password)...)
	secretData := secretsHashState{}
	if !password.IsNull() && !password.IsUnknown() {
		payload.Password = password.ValueStringPointer()
		passwordVersion = types.Int64Value(1)
		h := sha256.New()
		h.Write([]byte(password.ValueString()))
		secretData.Password = hex.EncodeToString(h.Sum(nil))
		secretDataJSON, _ := json.Marshal(secretData)
		val := map[string]string{"algo": "sha256", "hash": string(secretDataJSON)}
		hashedSecret, err := json.Marshal(val)
		if err != nil {
			resp.Diagnostics.AddError("Private State Marshal Error", err.Error())
			return
		}
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, "password_hash", hashedSecret)...)
	}

	var apiRes *security.CreateFtpuserResponse

	err := retry.Do(ctx, retry.TransientErrors, func(ctx context.Context) (int, error) {
		var (
			httpRes *http.Response
			callErr error
		)
		apiRes, httpRes, callErr = r.client.SecurityAPI.
			FtpuserAPI.
			Create(ctx).
			Ftpuser(*payload).
			ReturnFieldsPlus(readableAttributesForFtpuser).
			ReturnAsObject(1).
			Execute()

		if httpRes != nil {
			return httpRes.StatusCode, callErr
		}
		return 0, callErr
	})

	if err != nil {
		if retry.IsAlreadyExistsErr(err) {
			// Resource already exists, import required
			resp.Diagnostics.AddError(
				"Resource Already Exists",
				fmt.Sprintf("Resource already exists, error: %s.\nPlease import the existing resource into terraform state.", err.Error()),
			)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Ftpuser, got error: %s", err))
		return
	}

	res := apiRes.CreateFtpuserResponseAsObject.GetResult()
	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while create Ftpuser due inherited Extensible attributes, got error: %s", err))
		return
	}

	data.PasswordVersion = passwordVersion
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FtpuserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var diags diag.Diagnostics
	var data FtpuserModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	associateInternalId, diags := req.Private.GetKey(ctx, "associate_internal_id")
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resourceRef := utils.ExtractResourceRef(data.Ref.ValueString())

	var (
		httpRes *http.Response
		apiRes  *security.GetFtpuserResponse
	)

	err := retry.Do(ctx, nil, func(ctx context.Context) (int, error) {
		var callErr error
		apiRes, httpRes, callErr = r.client.SecurityAPI.
			FtpuserAPI.
			Read(ctx, resourceRef).
			ReturnFieldsPlus(readableAttributesForFtpuser).
			ReturnAsObject(1).
			ProxySearch(config.GetProxySearch()).
			Execute()

		if httpRes != nil {
			return httpRes.StatusCode, callErr
		}
		return 0, callErr
	})

	// If the resource is not found, try searching using Extensible Attributes
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound && r.ReadByExtAttrs(ctx, &data, resp) {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Ftpuser, got error: %s", err))
		return
	}

	res := apiRes.GetFtpuserResponseObjectAsResult.GetResult()

	apiTerraformId, ok := (*res.ExtAttrs)[terraformInternalIDEA]
	if !ok {
		apiTerraformId.Value = ""
	}

	if associateInternalId == nil {
		stateExtAttrs := ExpandExtAttrs(ctx, data.ExtAttrsAll, &diags)
		if stateExtAttrs == nil {
			resp.Diagnostics.AddError(
				"Missing Internal ID",
				"Unable to read Ftpuser because the internal ID (from extattrs_all) is missing or invalid.",
			)
			return
		}

		stateTerraformId := (*stateExtAttrs)[terraformInternalIDEA]
		if apiTerraformId.Value != stateTerraformId.Value {
			if r.ReadByExtAttrs(ctx, &data, resp) {
				return
			}
		}
	}

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while reading Ftpuser due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FtpuserResource) ReadByExtAttrs(ctx context.Context, data *FtpuserModel, resp *resource.ReadResponse) bool {
	var diags diag.Diagnostics

	if data.ExtAttrsAll.IsNull() {
		return false
	}

	internalIdExtAttr := *ExpandExtAttrs(ctx, data.ExtAttrsAll, &diags)
	if diags.HasError() {
		return false
	}

	internalId := internalIdExtAttr[terraformInternalIDEA].Value
	if internalId == "" {
		return false
	}

	idMap := map[string]interface{}{
		terraformInternalIDEA: internalId,
	}

	apiRes, _, err := r.client.SecurityAPI.
		FtpuserAPI.
		List(ctx).
		Extattrfilter(idMap).
		ReturnAsObject(1).
		ReturnFieldsPlus(readableAttributesForFtpuser).
		ProxySearch(config.GetProxySearch()).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Ftpuser by extattrs, got error: %s", err))
		return true
	}

	results := apiRes.ListFtpuserResponseObject.GetResult()

	// If the list is empty, the resource no longer exists so remove it from state
	if len(results) == 0 {
		resp.State.RemoveResource(ctx)
		return true
	}

	res := results[0]

	// Remove inherited external attributes from extattrs
	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		return true
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)

	return true
}

func (r *FtpuserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data FtpuserModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	planExtAttrs := data.ExtAttrs
	diags = req.State.GetAttribute(ctx, path.Root("ref"), &data.Ref)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	diags = req.State.GetAttribute(ctx, path.Root("extattrs_all"), &data.ExtAttrsAll)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	associateInternalId, diags := req.Private.GetKey(ctx, "associate_internal_id")
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}
	if associateInternalId != nil {
		data.ExtAttrs, diags = AddInternalIDToExtAttrs(ctx, data.ExtAttrs, diags)
		if diags.HasError() {
			return
		}
	}

	// Add Inherited Extensible Attributes
	data.ExtAttrs, diags = AddInheritedExtAttrs(ctx, data.ExtAttrs, data.ExtAttrsAll)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	var password types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("password"), &password)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resourceRef := utils.ExtractResourceRef(data.Ref.ValueString())

	payload := data.Expand(ctx, &resp.Diagnostics, false)
	if resp.Diagnostics.HasError() {
		return
	}

	if !password.IsNull() && !password.IsUnknown() {
		payload.Password = password.ValueStringPointer()
	}

	var apiRes *security.UpdateFtpuserResponse

	err := retry.Do(ctx, retry.TransientErrors, func(ctx context.Context) (int, error) {
		var (
			httpRes *http.Response
			callErr error
		)
		apiRes, httpRes, callErr = r.client.SecurityAPI.
			FtpuserAPI.
			Update(ctx, resourceRef).
			Ftpuser(*payload).
			ReturnFieldsPlus(readableAttributesForFtpuser).
			ReturnAsObject(1).
			Execute()

		if httpRes != nil {
			return httpRes.StatusCode, callErr
		}
		return 0, callErr
	})

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Ftpuser, got error: %s", err))
		return
	}

	res := apiRes.UpdateFtpuserResponseAsObject.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, planExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while update Ftpuser due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

	if associateInternalId != nil {
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", nil)...)
	}
}

func (r *FtpuserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data FtpuserModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resourceRef := utils.ExtractResourceRef(data.Ref.ValueString())

	err := retry.Do(ctx, retry.TransientErrors, func(ctx context.Context) (int, error) {
		httpRes, callErr := r.client.SecurityAPI.
			FtpuserAPI.
			Delete(ctx, resourceRef).
			Execute()

		if httpRes != nil {
			if httpRes.StatusCode == http.StatusNotFound {
				return 0, nil
			}
			return httpRes.StatusCode, callErr
		}
		return 0, callErr
	})

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Ftpuser, got error: %s", err))
		return
	}
}

func (r *FtpuserResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("ref"), req.ID)...)
	resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", []byte("true"))...)
}

func (r *FtpuserResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data FtpuserModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	createHomeDir := data.CreateHomeDir
	homeDir := data.HomeDir

	// If create_home_dir is unknown, we can't validate (during planning)
	if createHomeDir.IsUnknown() {
		return
	}

	// If create_home_dir is true and home_dir is specified, that's invalid
	if !createHomeDir.IsNull() && createHomeDir.ValueBool() && !homeDir.IsNull() && !homeDir.IsUnknown() {
		resp.Diagnostics.AddError(
			"Invalid Configuration",
			"You cannot specify 'home_dir' when 'create_home_dir' is set to true. When create_home_dir is true, the system automatically creates the home directory using the username.",
		)
	}
	if !createHomeDir.IsNull() && !createHomeDir.ValueBool() && (homeDir.IsNull() || homeDir.IsUnknown()) {
		resp.Diagnostics.AddError(
			"Invalid Configuration",
			"You must specify home_dir when create_home_dir is set to False.",
		)
	}
}
