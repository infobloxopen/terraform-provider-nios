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
	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForAdmingroup = "access_method,admin_set_commands,admin_show_commands,admin_toplevel_commands,cloud_set_commands,cloud_show_commands,comment,database_set_commands,database_show_commands,dhcp_set_commands,dhcp_show_commands,disable,disable_concurrent_login,dns_set_commands,dns_show_commands,dns_toplevel_commands,docker_set_commands,docker_show_commands,email_addresses,enable_restricted_user_access,extattrs,grid_set_commands,grid_show_commands,inactivity_lockout_setting,licensing_set_commands,licensing_show_commands,lockout_setting,machine_control_toplevel_commands,name,networking_set_commands,networking_show_commands,password_setting,roles,saml_setting,security_set_commands,security_show_commands,superuser,trouble_shooting_toplevel_commands,use_account_inactivity_lockout_enable,use_disable_concurrent_login,use_lockout_setting,use_password_setting,user_access"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AdmingroupResource{}
var _ resource.ResourceWithImportState = &AdmingroupResource{}
var _ resource.ResourceWithValidateConfig = &AdmingroupResource{}

func NewAdmingroupResource() resource.Resource {
	return &AdmingroupResource{}
}

// AdmingroupResource defines the resource implementation.
type AdmingroupResource struct {
	client *niosclient.APIClient
}

func (r *AdmingroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "security_admin_group"
}

func (r *AdmingroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages an Admin Group.",
		Attributes:          AdmingroupResourceSchemaAttributes,
	}
}

func (r *AdmingroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *AdmingroupResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var config AdmingroupModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	// Skip validation if UserAccess is not provided
	if config.UserAccess.IsNull() || config.UserAccess.IsUnknown() {
		return
	}

	refCount := 0
	// Track if we have any entries with address and permission
	hasAnyAddressWithPermission := false

	// Validate each user access entry
	for i, elem := range config.UserAccess.Elements() {
		obj := elem.(types.Object)
		attrMap := obj.Attributes()

		// Check field presence
		hasAddress := !attrMap["address"].IsUnknown() && !attrMap["address"].IsNull() && !attrMap["address"].Equal(types.StringValue(""))
		hasPermission := !attrMap["permission"].IsUnknown() && !attrMap["permission"].IsNull() && !attrMap["permission"].Equal(types.StringValue(""))
		hasRef := !attrMap["ref"].IsNull() && !attrMap["ref"].Equal(types.StringValue(""))

		// Rule 1: Can't have both ref and (address or permission)
		if hasRef && (hasAddress || hasPermission) {
			resp.Diagnostics.AddAttributeError(
				path.Root("user_access").AtListIndex(i),
				"Invalid combination of fields",
				"An entry cannot contain both ACL and ACEs. Address and Permission cannot be Set if ref for Named ACL is provided.",
			)
			continue
		}

		// Rule 2: Must have address if ref is not provided
		if !hasRef && !hasAddress {
			resp.Diagnostics.AddAttributeError(
				path.Root("user_access").AtListIndex(i),
				"Invalid Configuration for User Access",
				"An element must contain either 'address' for an ACE or 'ref' of the Named ACL.",
			)
			continue
		}

		// Track if we have an entry with both address and permission
		if hasAddress && hasPermission {
			hasAnyAddressWithPermission = true
		}

		// Count refs for collection-level validation
		if hasRef {
			refCount++
		}
	}

	// Collection-level validations
	if refCount > 0 && hasAnyAddressWithPermission {
		resp.Diagnostics.AddAttributeError(
			path.Root("user_access"),
			"Invalid combination of fields",
			"An element must contain either 'address' for an ACE or 'ref' of the Named ACL.",
		)
	}

	if refCount > 1 {
		resp.Diagnostics.AddAttributeError(
			path.Root("user_access"),
			"Too many references provided for Named ACL",
			"Either only one ACL or a set of ACEs is allowed in user_access.",
		)
	}
}

func (r *AdmingroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var diags diag.Diagnostics
	var data AdmingroupModel

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

	apiRes, _, err := r.client.SecurityAPI.
		AdmingroupAPI.
		Create(ctx).
		Admingroup(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForAdmingroup).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Admingroup, got error: %s", err))
		return
	}

	res := apiRes.CreateAdmingroupResponseAsObject.GetResult()
	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while create Admingroup due inherited Extensible attributes, got error: %s", err))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AdmingroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var diags diag.Diagnostics
	var data AdmingroupModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.SecurityAPI.
		AdmingroupAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForAdmingroup).
		ReturnAsObject(1).
		Execute()

	// If the resource is not found, try searching using Extensible Attributes
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound && r.ReadByExtAttrs(ctx, &data, resp) {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Admingroup, got error: %s", err))
		return
	}

	res := apiRes.GetAdmingroupResponseObjectAsResult.GetResult()

	apiTerraformId, ok := (*res.ExtAttrs)[terraformInternalIDEA]
	if !ok {
		apiTerraformId.Value = ""
	}

	stateExtAttrs := ExpandExtAttrs(ctx, data.ExtAttrsAll, &diags)
	if stateExtAttrs == nil {
		resp.Diagnostics.AddError(
			"Missing Internal ID",
			"Unable to read Admingroup because the internal ID (from extattrs_all) is missing or invalid.",
		)
		return
	}

	stateTerraformId := (*stateExtAttrs)[terraformInternalIDEA]
	if apiTerraformId.Value != stateTerraformId.Value {
		if r.ReadByExtAttrs(ctx, &data, resp) {
			return
		}
	}

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while reading Admingroup due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AdmingroupResource) ReadByExtAttrs(ctx context.Context, data *AdmingroupModel, resp *resource.ReadResponse) bool {
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
		AdmingroupAPI.
		List(ctx).
		Extattrfilter(idMap).
		ReturnAsObject(1).
		ReturnFieldsPlus(readableAttributesForAdmingroup).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Admingroup by extattrs, got error: %s", err))
		return true
	}

	results := apiRes.ListAdmingroupResponseObject.GetResult()

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

func (r *AdmingroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data AdmingroupModel

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

	// Add Inherited Extensible Attributes
	data.ExtAttrs, diags = AddInheritedExtAttrs(ctx, data.ExtAttrs, data.ExtAttrsAll)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	apiRes, _, err := r.client.SecurityAPI.
		AdmingroupAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Admingroup(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForAdmingroup).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Admingroup, got error: %s", err))
		return
	}

	res := apiRes.UpdateAdmingroupResponseAsObject.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, planExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while update Admingroup due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AdmingroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AdmingroupModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.SecurityAPI.
		AdmingroupAPI.
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Admingroup, got error: %s", err))
		return
	}
}

func (r *AdmingroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var diags diag.Diagnostics
	var data AdmingroupModel
	var goClientData security.Admingroup

	resourceRef := utils.ExtractResourceRef(req.ID)
	extattrs, diags := AddInternalIDToExtAttrs(ctx, data.ExtAttrs, diags)
	if diags.HasError() {
		return
	}
	goClientData.ExtAttrsPlus = ExpandExtAttrs(ctx, extattrs, &diags)
	data.ExtAttrsAll = extattrs

	updateRes, _, err := r.client.SecurityAPI.
		AdmingroupAPI.
		Update(ctx, resourceRef).
		Admingroup(goClientData).
		ReturnFieldsPlus(readableAttributesForAdmingroup).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Import Failed", fmt.Sprintf("Unable to update Admingroup for import, got error: %s", err))
		return
	}

	res := updateRes.UpdateAdmingroupResponseAsObject.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrsAll, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while update Admingroup due inherited Extensible attributes for import, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("ref"), req.ID)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("extattrs_all"), data.ExtAttrsAll)...)
}
