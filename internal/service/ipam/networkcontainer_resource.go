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
	"github.com/infobloxopen/infoblox-nios-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForNetworkcontainer = "authority,bootfile,bootserver,cloud_info,comment,ddns_domainname,ddns_generate_hostname,ddns_server_always_updates,ddns_ttl,ddns_update_fixed_addresses,ddns_use_option81,deny_bootp,discover_now_status,discovery_basic_poll_settings,discovery_blackout_setting,discovery_engine_type,discovery_member,email_list,enable_ddns,enable_dhcp_thresholds,enable_discovery,enable_email_warnings,enable_pxe_lease_time,enable_snmp_warnings,endpoint_sources,extattrs,federated_realms,high_water_mark,high_water_mark_reset,ignore_dhcp_option_list_request,ignore_id,ignore_mac_addresses,ipam_email_addresses,ipam_threshold_settings,ipam_trap_settings,last_rir_registration_update_sent,last_rir_registration_update_status,lease_scavenge_time,logic_filter_rules,low_water_mark,low_water_mark_reset,mgm_private,mgm_private_overridable,ms_ad_user_data,network,network_container,network_view,nextserver,options,port_control_blackout_setting,pxe_lease_time,recycle_leases,rir,rir_organization,rir_registration_status,same_port_control_discovery_blackout,subscribe_settings,unmanaged,update_dns_on_lease_renewal,use_authority,use_blackout_setting,use_bootfile,use_bootserver,use_ddns_domainname,use_ddns_generate_hostname,use_ddns_ttl,use_ddns_update_fixed_addresses,use_ddns_use_option81,use_deny_bootp,use_discovery_basic_polling_settings,use_email_list,use_enable_ddns,use_enable_dhcp_thresholds,use_enable_discovery,use_ignore_dhcp_option_list_request,use_ignore_id,use_ipam_email_addresses,use_ipam_threshold_settings,use_ipam_trap_settings,use_lease_scavenge_time,use_logic_filter_rules,use_mgm_private,use_nextserver,use_options,use_pxe_lease_time,use_recycle_leases,use_subscribe_settings,use_update_dns_on_lease_renewal,use_zone_associations,utilization,zone_associations"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NetworkcontainerResource{}
var _ resource.ResourceWithImportState = &NetworkcontainerResource{}

func NewNetworkcontainerResource() resource.Resource {
	return &NetworkcontainerResource{}
}

// NetworkcontainerResource defines the resource implementation.
type NetworkcontainerResource struct {
	client *niosclient.APIClient
}

func (r *NetworkcontainerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "ipam_network_container"
}

func (r *NetworkcontainerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a Network Container",
		Attributes:          NetworkcontainerResourceSchemaAttributes,
	}
}

func (r *NetworkcontainerResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *NetworkcontainerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var diags diag.Diagnostics
	var data NetworkcontainerModel

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
		NetworkcontainerAPI.
		Create(ctx).
		Networkcontainer(*data.Expand(ctx, &resp.Diagnostics, true)).
		ReturnFieldsPlus(readableAttributesForNetworkcontainer).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Networkcontainer, got error: %s", err))
		return
	}

	res := apiRes.CreateNetworkcontainerResponseAsObject.GetResult()
	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while create Networkcontainer due inherited Extensible attributes, got error: %s", err))
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

func (r *NetworkcontainerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var diags diag.Diagnostics
	var data NetworkcontainerModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.IPAMAPI.
		NetworkcontainerAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForNetworkcontainer).
		ReturnAsObject(1).
		Execute()

	// If the resource is not found, try searching using Extensible Attributes
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound && r.ReadByExtAttrs(ctx, &data, resp) {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Networkcontainer, got error: %s", err))
		return
	}

	res := apiRes.GetNetworkcontainerResponseObjectAsResult.GetResult()

	apiTerraformId, ok := (*res.ExtAttrs)[terraformInternalIDEA]
	if !ok {
		apiTerraformId.Value = ""
	}

	stateExtAttrs := ExpandExtAttrs(ctx, data.ExtAttrsAll, &diags)
	if stateExtAttrs == nil {
		resp.Diagnostics.AddError(
			"Missing Internal ID",
			"Unable to read Networkcontainer because the internal ID (from extattrs_all) is missing or invalid.",
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while reading Networkcontainer due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NetworkcontainerResource) ReadByExtAttrs(ctx context.Context, data *NetworkcontainerModel, resp *resource.ReadResponse) bool {
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
		NetworkcontainerAPI.
		List(ctx).
		Extattrfilter(idMap).
		ReturnAsObject(1).
		ReturnFieldsPlus(readableAttributesForNetworkcontainer).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Networkcontainer by extattrs, got error: %s", err))
		return true
	}

	results := apiRes.ListNetworkcontainerResponseObject.GetResult()

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

func (r *NetworkcontainerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data NetworkcontainerModel

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
		NetworkcontainerAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Networkcontainer(*data.Expand(ctx, &resp.Diagnostics, false)).
		ReturnFieldsPlus(readableAttributesForNetworkcontainer).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Networkcontainer, got error: %s", err))
		return
	}

	res := apiRes.UpdateNetworkcontainerResponseAsObject.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, planExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while update Networkcontainer due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NetworkcontainerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NetworkcontainerModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.IPAMAPI.
		NetworkcontainerAPI.
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Networkcontainer, got error: %s", err))
		return
	}
}

func (r *NetworkcontainerResource) UpdateFuncCallAttributeName(ctx context.Context, data NetworkcontainerModel, diags *diag.Diagnostics) types.Object {

	updatedFuncCallAttrs := data.FuncCall.Attributes()
	attrVal := updatedFuncCallAttrs["attribute_name"].(types.String).ValueString()
	pathVar, err := utils.FindModelFieldByTFSdkTag(data, attrVal)
	if !err {
		diags.AddError("Client Error", fmt.Sprintf("Unable to find attribute '%s' in Networkcontainer model, got error", attrVal))
		return types.ObjectNull(FuncCallAttrTypes)
	}
	updatedFuncCallAttrs["attribute_name"] = types.StringValue(pathVar)

	return types.ObjectValueMust(FuncCallAttrTypes, updatedFuncCallAttrs)
}

func (r *NetworkcontainerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var diags diag.Diagnostics
	var data NetworkcontainerModel
	var goClientData ipam.Networkcontainer

	resourceRef := utils.ExtractResourceRef(req.ID)
	extattrs, diags := AddInternalIDToExtAttrs(ctx, data.ExtAttrs, diags)
	if diags.HasError() {
		return
	}
	goClientData.ExtAttrsPlus = ExpandExtAttrs(ctx, extattrs, &diags)
	data.ExtAttrsAll = extattrs

	updateRes, _, err := r.client.IPAMAPI.
		NetworkcontainerAPI.
		Update(ctx, resourceRef).
		Networkcontainer(goClientData).
		ReturnFieldsPlus(readableAttributesForNetworkcontainer).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Import Failed", fmt.Sprintf("Unable to update Networkcontainer for import, got error: %s", err))
		return
	}

	res := updateRes.UpdateNetworkcontainerResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NetworkcontainerResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data NetworkcontainerModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Check if options are defined
	if !data.Options.IsNull() && !data.Options.IsUnknown() {
		// Special DHCP option names that require use_option to be set
		specialOptions := map[string]bool{
			"routers":                  true,
			"router-templates":         true,
			"domain-name-servers":      true,
			"domain-name":              true,
			"broadcast-address":        true,
			"broadcast-address-offset": true,
			"dhcp-lease-time":          true,
			"dhcp6.name-servers":       true,
		}

		var options []NetworkcontainerOptionsModel
		diags := data.Options.ElementsAs(ctx, &options, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		for i, option := range options {
			if option.Name.IsNull() || option.Name.IsUnknown() {
				continue
			}

			optionName := option.Name.ValueString()
			isSpecialOption := specialOptions[optionName]

			if !isSpecialOption && !option.UseOption.IsNull() && !option.UseOption.IsUnknown() {
				resp.Diagnostics.AddAttributeError(
					path.Root("options").AtListIndex(i).AtName("use_option"),
					"Invalid configuration",
					fmt.Sprintf("The 'use_option' attribute should not be set for Custom DHCP Options '%s'. "+
						"It is only applicable for special options: routers, router-templates, domain-name-servers, "+
						"domain-name, broadcast-address, broadcast-address-offset, dhcp-lease-time, dhcp6.name-servers.",
						optionName),
				)
			}
		}
	}
}
