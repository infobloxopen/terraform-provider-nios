package ipam

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

var readableAttributesForIpv6networkcontainer = "cloud_info,comment,ddns_domainname,ddns_enable_option_fqdn,ddns_generate_hostname,ddns_server_always_updates,ddns_ttl,discover_now_status,discovery_basic_poll_settings,discovery_blackout_setting,discovery_engine_type,discovery_member,domain_name_servers,enable_ddns,enable_discovery,endpoint_sources,extattrs,last_rir_registration_update_sent,last_rir_registration_update_status,logic_filter_rules,mgm_private,mgm_private_overridable,ms_ad_user_data,network,network_container,network_view,options,port_control_blackout_setting,preferred_lifetime,rir,rir_organization,rir_registration_status,same_port_control_discovery_blackout,subscribe_settings,unmanaged,update_dns_on_lease_renewal,use_blackout_setting,use_ddns_domainname,use_ddns_enable_option_fqdn,use_ddns_generate_hostname,use_ddns_ttl,use_discovery_basic_polling_settings,use_domain_name_servers,use_enable_ddns,use_enable_discovery,use_logic_filter_rules,use_mgm_private,use_options,use_preferred_lifetime,use_subscribe_settings,use_update_dns_on_lease_renewal,use_valid_lifetime,use_zone_associations,utilization,valid_lifetime,zone_associations"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &Ipv6networkcontainerResource{}
var _ resource.ResourceWithImportState = &Ipv6networkcontainerResource{}

func NewIpv6networkcontainerResource() resource.Resource {
	return &Ipv6networkcontainerResource{}
}

// Ipv6networkcontainerResource defines the resource implementation.
type Ipv6networkcontainerResource struct {
	client *niosclient.APIClient
}

func (r *Ipv6networkcontainerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "ipam_ipv6network_container"
}

func (r *Ipv6networkcontainerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages an IPv6 Network Container",
		Attributes:          Ipv6networkcontainerResourceSchemaAttributes,
	}
}

func (r *Ipv6networkcontainerResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *Ipv6networkcontainerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var diags diag.Diagnostics
	var data Ipv6networkcontainerModel

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

	// If the function call attributes are set, update the attribute name to match tfsdk tag
	origFunCallAttrs := data.FuncCall.Attributes()
	if len(origFunCallAttrs) > 0 {
		data.FuncCall = r.UpdateFuncCallAttributeName(ctx, data, &resp.Diagnostics)
	}

	apiRes, _, err := r.client.IPAMAPI.
		Ipv6networkcontainerAPI.
		Create(ctx).
		Ipv6networkcontainer(*data.Expand(ctx, &resp.Diagnostics, true)).
		ReturnFieldsPlus(readableAttributesForIpv6networkcontainer).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Ipv6networkcontainer, got error: %s", err))
		return
	}

	res := apiRes.CreateIpv6networkcontainerResponseAsObject.GetResult()
	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while create Ipv6networkcontainer due inherited Extensible attributes, got error: %s", err))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Retain the original function call attributes
	if len(origFunCallAttrs) > 0 {
		data.FuncCall = types.ObjectValueMust(FuncCallAttrTypes, origFunCallAttrs)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Ipv6networkcontainerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var diags diag.Diagnostics
	var data Ipv6networkcontainerModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.IPAMAPI.
		Ipv6networkcontainerAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForIpv6networkcontainer).
		ReturnAsObject(1).
		Execute()

	// If the resource is not found, try searching using Extensible Attributes
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound && r.ReadByExtAttrs(ctx, &data, resp) {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Ipv6networkcontainer, got error: %s", err))
		return
	}

	res := apiRes.GetIpv6networkcontainerResponseObjectAsResult.GetResult()

	apiTerraformId, ok := (*res.ExtAttrs)[terraformInternalIDEA]
	if !ok {
		apiTerraformId.Value = ""
	}

	stateExtAttrs := ExpandExtAttrs(ctx, data.ExtAttrsAll, &diags)
	if stateExtAttrs == nil {
		resp.Diagnostics.AddError(
			"Missing Internal ID",
			"Unable to read Ipv6networkcontainer because the internal ID (from extattrs_all) is missing or invalid.",
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while reading Ipv6networkcontainer due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Ipv6networkcontainerResource) ReadByExtAttrs(ctx context.Context, data *Ipv6networkcontainerModel, resp *resource.ReadResponse) bool {
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

	apiRes, _, err := r.client.IPAMAPI.
		Ipv6networkcontainerAPI.
		List(ctx).
		Extattrfilter(idMap).
		ReturnAsObject(1).
		ReturnFieldsPlus(readableAttributesForIpv6networkcontainer).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Ipv6networkcontainer by extattrs, got error: %s", err))
		return true
	}

	results := apiRes.ListIpv6networkcontainerResponseObject.GetResult()

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

func (r *Ipv6networkcontainerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data Ipv6networkcontainerModel

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

	apiRes, _, err := r.client.IPAMAPI.
		Ipv6networkcontainerAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Ipv6networkcontainer(*data.Expand(ctx, &resp.Diagnostics, false)).
		ReturnFieldsPlus(readableAttributesForIpv6networkcontainer).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Ipv6networkcontainer, got error: %s", err))
		return
	}

	res := apiRes.UpdateIpv6networkcontainerResponseAsObject.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, planExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while update Ipv6networkcontainer due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Ipv6networkcontainerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Ipv6networkcontainerModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.IPAMAPI.
		Ipv6networkcontainerAPI.
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Ipv6networkcontainer, got error: %s", err))
		return
	}
}

func (r *Ipv6networkcontainerResource) UpdateFuncCallAttributeName(ctx context.Context, data Ipv6networkcontainerModel, diags *diag.Diagnostics) types.Object {

	updatedFuncCallAttrs := data.FuncCall.Attributes()
	attrVal := updatedFuncCallAttrs["attribute_name"].(types.String).ValueString()
	pathVar, err := utils.FindModelFieldByTFSdkTag(data, attrVal)
	if !err {
		diags.AddError("Client Error", fmt.Sprintf("Unable to find attribute '%s' in Ipv6networkcontainer model, got error", attrVal))
		return types.ObjectNull(FuncCallAttrTypes)
	}
	updatedFuncCallAttrs["attribute_name"] = types.StringValue(pathVar)

	return types.ObjectValueMust(FuncCallAttrTypes, updatedFuncCallAttrs)
}

func (r *Ipv6networkcontainerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var diags diag.Diagnostics
	var data Ipv6networkcontainerModel

	resourceRef := utils.ExtractResourceRef(req.ID)

	apiRes, _, err := r.client.IPAMAPI.
		Ipv6networkcontainerAPI.
		Read(ctx, resourceRef).
		ReturnFieldsPlus(readableAttributesForIpv6networkcontainer).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Import Failed", fmt.Sprintf("Cannot read Ipv6networkcontainer for import, got error: %s", err))
		return
	}

	res := apiRes.GetIpv6networkcontainerResponseObjectAsResult.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while reading Ipv6networkcontainer for import due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	planExtAttrs := data.ExtAttrs
	data.ExtAttrs, diags = AddInheritedExtAttrs(ctx, data.ExtAttrs, data.ExtAttrsAll)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	data.ExtAttrs, diags = AddInternalIDToExtAttrs(ctx, data.ExtAttrs, diags)
	if diags.HasError() {
		return
	}

	updateRes, _, err := r.client.IPAMAPI.
		Ipv6networkcontainerAPI.
		Update(ctx, resourceRef).
		Ipv6networkcontainer(*data.Expand(ctx, &resp.Diagnostics, false)).
		ReturnFieldsPlus(readableAttributesForIpv6networkcontainer).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Import Failed", fmt.Sprintf("Unable to update Ipv6networkcontainer for import, got error: %s", err))
		return
	}

	res = updateRes.UpdateIpv6networkcontainerResponseAsObject.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, planExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while update Ipv6networkcontainer due inherited Extensible attributes for import, got error: %s", diags))
		return
	}
	data.Flatten(ctx, &res, &resp.Diagnostics)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
