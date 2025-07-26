package ipam

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForIpv6network = "cloud_info,comment,ddns_domainname,ddns_enable_option_fqdn,ddns_generate_hostname,ddns_server_always_updates,ddns_ttl,disable,discover_now_status,discovered_bgp_as,discovered_bridge_domain,discovered_tenant,discovered_vlan_id,discovered_vlan_name,discovered_vrf_description,discovered_vrf_name,discovered_vrf_rd,discovery_basic_poll_settings,discovery_blackout_setting,discovery_engine_type,discovery_member,domain_name,domain_name_servers,enable_ddns,enable_discovery,enable_ifmap_publishing,endpoint_sources,extattrs,last_rir_registration_update_sent,last_rir_registration_update_status,logic_filter_rules,members,mgm_private,mgm_private_overridable,ms_ad_user_data,network,network_container,network_view,options,port_control_blackout_setting,preferred_lifetime,recycle_leases,rir,rir_organization,rir_registration_status,same_port_control_discovery_blackout,subscribe_settings,unmanaged,unmanaged_count,update_dns_on_lease_renewal,use_blackout_setting,use_ddns_domainname,use_ddns_enable_option_fqdn,use_ddns_generate_hostname,use_ddns_ttl,use_discovery_basic_polling_settings,use_domain_name,use_domain_name_servers,use_enable_ddns,use_enable_discovery,use_enable_ifmap_publishing,use_logic_filter_rules,use_mgm_private,use_options,use_preferred_lifetime,use_recycle_leases,use_subscribe_settings,use_update_dns_on_lease_renewal,use_valid_lifetime,use_zone_associations,valid_lifetime,vlans,zone_associations"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &Ipv6networkResource{}
var _ resource.ResourceWithImportState = &Ipv6networkResource{}

func NewIpv6networkResource() resource.Resource {
	return &Ipv6networkResource{}
}

// Ipv6networkResource defines the resource implementation.
type Ipv6networkResource struct {
	client *niosclient.APIClient
}

func (r *Ipv6networkResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "ipam_ipv6network"
}

func (r *Ipv6networkResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "",
		Attributes:          Ipv6networkResourceSchemaAttributes,
	}
}

func (r *Ipv6networkResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *Ipv6networkResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var diags diag.Diagnostics
	var data Ipv6networkModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Add internal ID exists in the Extensible Attributes if not already present
	if err := r.addInternalIDToExtAttrs(ctx, &data); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to add internal ID to Extensible Attributes, got error: %s", err))
		return
	}

	// If the function call attributes are set, update the attribute name to match tfsdk tag
	origFunCallAttrs := data.FuncCall.Attributes()
	if len(origFunCallAttrs) > 0 {
		data.FuncCall = r.UpdateFuncCallAttributeName(ctx, data, &resp.Diagnostics)
	}

	apiRes, _, err := r.client.IPAMAPI.
		Ipv6networkAPI.
		Create(ctx).
		Ipv6network(*data.Expand(ctx, &resp.Diagnostics, true)).
		ReturnFieldsPlus(readableAttributesForIpv6network).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Ipv6network, got error: %s", err))
		return
	}

	res := apiRes.CreateIpv6networkResponseAsObject.GetResult()
	res.ExtAttrs, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while create Ipv6network due inherited Extensible attributes, got error: %s", err))
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

func (r *Ipv6networkResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var diags diag.Diagnostics
	var data Ipv6networkModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.IPAMAPI.
		Ipv6networkAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForIpv6network).
		ReturnAsObject(1).
		Execute()

	// If the resource is not found, try searching using Extensible Attributes
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound && r.ReadByExtAttrs(ctx, &data, resp) {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Ipv6network, got error: %s", err))
		return
	}

	res := apiRes.GetIpv6networkResponseObjectAsResult.GetResult()
	if res.ExtAttrs == nil {
		resp.Diagnostics.AddError(
			"Missing Extensible Attributes",
			"Unable to read Ipv6network because no extensible attributes were returned from the API.",
		)
		return
	}

	res.ExtAttrs, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while reading Ipv6network due inherited Extensible attributes, got error: %s", diags))
		return
	}

	apiTerraformId, ok := (*res.ExtAttrs)["Terraform Internal ID"]
	if !ok {
		resp.Diagnostics.AddError(
			"Missing Terraform internal id Attributes",
			"Unable to read Ipv6network because terraform internal id does not exist.",
		)
		return
	}

	stateExtAttrs := ExpandExtAttr(ctx, data.ExtAttrsAll, &diags)
	if stateExtAttrs == nil {
		resp.Diagnostics.AddError(
			"Missing Internal ID",
			"Unable to read Ipv6network because the internal ID (from extattrs_all) is missing or invalid.",
		)
		return
	}

	stateTerraformId := (*stateExtAttrs)["Terraform Internal ID"]

	if apiTerraformId.Value != stateTerraformId.Value {
		if r.ReadByExtAttrs(ctx, &data, resp) {
			return
		}
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Ipv6networkResource) ReadByExtAttrs(ctx context.Context, data *Ipv6networkModel, resp *resource.ReadResponse) bool {
	var diags diag.Diagnostics

	if data.ExtAttrsAll.IsNull() {
		return false
	}

	internalIdExtAttr := *ExpandExtAttr(ctx, data.ExtAttrsAll, &diags)
	if diags.HasError() {
		return false
	}

	internalId := internalIdExtAttr["Terraform Internal ID"].Value
	if internalId == "" {
		return false
	}

	idMap := map[string]interface{}{
		"Terraform Internal ID": internalId,
	}

	apiRes, _, err := r.client.IPAMAPI.
		Ipv6networkAPI.
		List(ctx).
		Extattrfilter(idMap).
		ReturnAsObject(1).
		ReturnFieldsPlus(readableAttributesForIpv6network).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Ipv6network by extattrs, got error: %s", err))
		return true
	}

	results := apiRes.ListIpv6networkResponseObject.GetResult()

	// If the list is empty, the resource no longer exists so remove it from state
	if len(results) == 0 {
		resp.State.RemoveResource(ctx)
		return true
	}

	res := results[0]

	// Remove inherited external attributes and check for errors
	res.ExtAttrs, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		return true
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)

	return true
}

func (r *Ipv6networkResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data Ipv6networkModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

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

	// Add internal ID exists in the Extensible Attributes if not already present
	if err := r.addInternalIDToExtAttrs(ctx, &data); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to add internal ID to Extensible Attributes, got error: %s", err))
		return
	}

	apiRes, _, err := r.client.IPAMAPI.
		Ipv6networkAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Ipv6network(*data.Expand(ctx, &resp.Diagnostics, false)).
		ReturnFieldsPlus(readableAttributesForIpv6network).
		ReturnAsObject(1).
		Execute()

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Ipv6network, got error: %s", err))
		return
	}

	res := apiRes.UpdateIpv6networkResponseAsObject.GetResult()

	res.ExtAttrs, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while update Ipv6network due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Ipv6networkResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Ipv6networkModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.IPAMAPI.
		Ipv6networkAPI.
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Ipv6network, got error: %s", err))
		return
	}
}

func (r *Ipv6networkResource) addInternalIDToExtAttrs(ctx context.Context, data *Ipv6networkModel) error {
	var internalId string

	if !data.ExtAttrsAll.IsNull() {
		elements := data.ExtAttrsAll.Elements()
		if tId, ok := elements["Terraform Internal ID"]; ok {
			if tIdStr, ok := tId.(types.String); ok {
				internalId = tIdStr.ValueString()
			}
		}
	}

	if internalId == "" {
		var err error
		internalId, err = uuid.GenerateUUID()
		if err != nil {
			return err
		}
	}

	r.client.IPAMAPI.APIClient.Cfg.DefaultExtAttrs = map[string]struct{ Value string }{
		"Terraform Internal ID": {Value: internalId},
	}

	return nil
}
func (r *Ipv6networkResource) UpdateFuncCallAttributeName(ctx context.Context, data Ipv6networkModel, diags *diag.Diagnostics) types.Object {

	updatedFuncCallAttrs := data.FuncCall.Attributes()
	attrVal := updatedFuncCallAttrs["attribute_name"].(types.String).ValueString()
	pathVar, err := utils.FindModelFieldByTFSdkTag(data, attrVal)
	if !err {
		diags.AddError("Client Error", fmt.Sprintf("Unable to find attribute '%s' in RecordA model, got error", attrVal))
		return types.ObjectNull(FuncCallAttrTypes)
	}
	updatedFuncCallAttrs["attribute_name"] = types.StringValue(pathVar)

	return types.ObjectValueMust(FuncCallAttrTypes, updatedFuncCallAttrs)
}

func (r *Ipv6networkResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("ref"), req, resp)
}
