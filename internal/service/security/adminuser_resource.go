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

var readableAttributesForAdminuser = "admin_groups,auth_method,auth_type,ca_certificate_issuer,client_certificate_serial_number,comment,disable,email,enable_certificate_authentication,extattrs,name,ssh_keys,status,time_zone,use_ssh_keys,use_time_zone"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AdminuserResource{}
var _ resource.ResourceWithImportState = &AdminuserResource{}
var _ resource.ResourceWithModifyPlan = &AdminuserResource{}

func NewAdminuserResource() resource.Resource {
	return &AdminuserResource{}
}

// AdminuserResource defines the resource implementation.
type AdminuserResource struct {
	client *niosclient.APIClient
}

type secretsHashState struct {
	Password string `json:"password_hash"`
}

func (r *AdminuserResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "security_admin_user"
}

func (r *AdminuserResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages an Admin User.",
		Attributes:          AdminuserResourceSchemaAttributes,
	}
}

func (r *AdminuserResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *AdminuserResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return
	}

	var statePwdRevision types.Int64
	var planPassword types.String

	// Normalize stateRev if null (e.g., first apply)
	curRev := int64(0)

	if !req.State.Raw.IsNull() && req.State.Raw.IsKnown() {
		resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("password_revision"), &statePwdRevision)...)
		if resp.Diagnostics.HasError() {
			return
		}
		if !statePwdRevision.IsNull() && !statePwdRevision.IsUnknown() {
			curRev = statePwdRevision.ValueInt64()
		}
	}
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("password"), &planPassword)...)
	if resp.Diagnostics.HasError() {
		return
	}

	computeNewHash := !planPassword.IsNull() && !planPassword.IsUnknown()

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
			// Increment revision and store new hash if password modified
			newRev := types.Int64Value(curRev + 1)
			resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("password_revision"), newRev)...)

			val := map[string]string{"algo": "sha256", "hash": plannedHash}
			b, err := json.Marshal(val)
			if err != nil {
				resp.Diagnostics.AddError("Private State Marshal Error", err.Error())
				return
			}
			resp.Diagnostics.Append(resp.Private.SetKey(ctx, "password_hash", b)...)
		} else {
			resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("password_revision"), curRev)...)
		}
	}

}

func (r *AdminuserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var diags diag.Diagnostics
	var data AdminuserModel

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

	payload := data.Expand(ctx, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	passwordRevision := types.Int64Value(0)
	var password types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("password"), &password)...)

	secretData := secretsHashState{}

	if !password.IsNull() && !password.IsUnknown() {

		payload.Password = password.ValueStringPointer()
		passwordRevision = types.Int64Value(1)
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
		AdminuserAPI.
		Create(ctx).
		Adminuser(*payload).
		ReturnFieldsPlus(readableAttributesForAdminuser).
		ReturnAsObject(1).
		Execute()

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Adminuser, got error: %s", err))
		return
	}

	res := apiRes.CreateAdminuserResponseAsObject.GetResult()
	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while create Adminuser due inherited Extensible attributes, got error: %s", err))
		return
	}

	data.PasswordRevision = passwordRevision
	plannedSshKeys := data.SshKeys
	data.Flatten(ctx, &res, &resp.Diagnostics)
	// The API may return ssh_keys even when not configured (e.g. use_ssh_keys=false).
	// Restore the planned value to avoid a plan/state inconsistency.
	if plannedSshKeys.IsNull() {
		data.SshKeys = plannedSshKeys
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AdminuserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var diags diag.Diagnostics
	var data AdminuserModel

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

	apiRes, httpRes, err := r.client.SecurityAPI.
		AdminuserAPI.
		Read(ctx, utils.ResolveIdentifier(data.Uuid, data.Ref)).
		ReturnFieldsPlus(readableAttributesForAdminuser).
		ReturnAsObject(1).
		ProxySearch(config.GetProxySearch()).
		Execute()

	// If the resource is not found, try searching using Extensible Attributes
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound && r.ReadByExtAttrs(ctx, &data, resp) {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Adminuser, got error: %s", err))
		return
	}

	res := apiRes.GetAdminuserResponseObjectAsResult.GetResult()

	apiTerraformId, ok := (*res.ExtAttrs)[terraformInternalIDEA]
	if !ok {
		apiTerraformId.Value = ""
	}

	if associateInternalId == nil {
		stateExtAttrs := ExpandExtAttrs(ctx, data.ExtAttrsAll, &diags)
		if stateExtAttrs == nil {
			resp.Diagnostics.AddError(
				"Missing Internal ID",
				"Unable to read Adminuser because the internal ID (from extattrs_all) is missing or invalid.",
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while reading Adminuser due inherited Extensible attributes, got error: %s", diags))
		return
	}

	priorSshKeys := data.SshKeys
	data.Flatten(ctx, &res, &resp.Diagnostics)
	if priorSshKeys.IsNull() {
		data.SshKeys = priorSshKeys
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AdminuserResource) ReadByExtAttrs(ctx context.Context, data *AdminuserModel, resp *resource.ReadResponse) bool {
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
		AdminuserAPI.
		List(ctx).
		Extattrfilter(idMap).
		ReturnAsObject(1).
		ReturnFieldsPlus(readableAttributesForAdminuser).
		ProxySearch(config.GetProxySearch()).
		Execute()

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Adminuser by extattrs, got error: %s", err))
		return true
	}

	results := apiRes.ListAdminuserResponseObject.GetResult()

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

	priorSshKeys := data.SshKeys
	data.Flatten(ctx, &res, &resp.Diagnostics)
	if priorSshKeys.IsNull() {
		data.SshKeys = priorSshKeys
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)

	return true
}

func (r *AdminuserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data AdminuserModel

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

	diags = req.State.GetAttribute(ctx, path.Root("uuid"), &data.Uuid)
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

	payload := data.Expand(ctx, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !password.IsNull() && !password.IsUnknown() {
		payload.Password = password.ValueStringPointer()
	}

	resourceRef := utils.ExtractResourceRef(data.Ref.ValueString())

	apiRes, _, err := r.client.SecurityAPI.
		AdminuserAPI.
		Update(ctx, resourceRef).
		Adminuser(*payload).
		ReturnFieldsPlus(readableAttributesForAdminuser).
		ReturnAsObject(1).
		Execute()

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Adminuser, got error: %s", err))
		return
	}

	res := apiRes.UpdateAdminuserResponseAsObject.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, planExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while update Adminuser due inherited Extensible attributes, got error: %s", diags))
		return
	}

	plannedSshKeys := data.SshKeys
	data.Flatten(ctx, &res, &resp.Diagnostics)
	// The API may return ssh_keys even when not configured (e.g. use_ssh_keys=false).
	// Restore the planned value to avoid a plan/state inconsistency.
	if plannedSshKeys.IsNull() {
		data.SshKeys = plannedSshKeys
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

	if associateInternalId != nil {
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", nil)...)
	}
}

func (r *AdminuserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AdminuserModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.SecurityAPI.
		AdminuserAPI.
		Delete(ctx, utils.ResolveIdentifier(data.Uuid, data.Ref)).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Adminuser, got error: %s", err))
		return
	}
}

func (r *AdminuserResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("uuid"), req.ID)...)
	resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", []byte("true"))...)
}
