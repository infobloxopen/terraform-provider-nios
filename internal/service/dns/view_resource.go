package dns

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

	"github.com/infobloxopen/terraform-provider-nios/internal/config"
)

var readableAttributesForView = "blacklist_action,blacklist_log_query,blacklist_redirect_addresses,blacklist_redirect_ttl,blacklist_rulesets,cloud_info,comment,custom_root_name_servers,ddns_force_creation_timestamp_update,ddns_principal_group,ddns_principal_tracking,ddns_restrict_patterns,ddns_restrict_patterns_list,ddns_restrict_protected,ddns_restrict_secure,ddns_restrict_static,disable,dns64_enabled,dns64_groups,dnssec_enabled,dnssec_expired_signatures_enabled,dnssec_negative_trust_anchors,dnssec_trusted_keys,dnssec_validation_enabled,edns_udp_size,enable_blacklist,enable_fixed_rrset_order_fqdns,enable_match_recursive_only,extattrs,filter_aaaa,filter_aaaa_list,fixed_rrset_order_fqdns,forward_only,forwarders,is_default,last_queried_acl,match_clients,match_destinations,max_cache_ttl,max_ncache_ttl,max_udp_size,name,network_view,notify_delay,nxdomain_log_query,nxdomain_redirect,nxdomain_redirect_addresses,nxdomain_redirect_addresses_v6,nxdomain_redirect_ttl,nxdomain_rulesets,recursion,response_rate_limiting,root_name_server_type,rpz_drop_ip_rule_enabled,rpz_drop_ip_rule_min_prefix_length_ipv4,rpz_drop_ip_rule_min_prefix_length_ipv6,rpz_qname_wait_recurse,scavenging_settings,sortlist,use_blacklist,use_ddns_force_creation_timestamp_update,use_ddns_patterns_restriction,use_ddns_principal_security,use_ddns_restrict_protected,use_ddns_restrict_static,use_dns64,use_dnssec,use_edns_udp_size,use_filter_aaaa,use_fixed_rrset_order_fqdns,use_forwarders,use_max_cache_ttl,use_max_ncache_ttl,use_max_udp_size,use_nxdomain_redirect,use_recursion,use_response_rate_limiting,use_root_name_server,use_rpz_drop_ip_rule,use_rpz_qname_wait_recurse,use_scavenging_settings,use_sortlist"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ViewResource{}
var _ resource.ResourceWithImportState = &ViewResource{}

func NewViewResource() resource.Resource {
	return &ViewResource{}
}

// ViewResource defines the resource implementation.
type ViewResource struct {
	client *niosclient.APIClient
}

func (r *ViewResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "dns_view"
}

func (r *ViewResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Managing DNS View",
		Attributes:          ViewResourceSchemaAttributes,
	}
}

func (r *ViewResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ViewResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var diags diag.Diagnostics
	var data ViewModel

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

	apiRes, _, err := r.client.DNSAPI.
		ViewAPI.
		Create(ctx).
		View(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForView).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create View, got error: %s", err))
		return
	}

	res := apiRes.CreateViewResponseAsObject.GetResult()
	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while create View due inherited Extensible attributes, got error: %s", err))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ViewResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var diags diag.Diagnostics
	var data ViewModel

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

	apiRes, httpRes, err := r.client.DNSAPI.
		ViewAPI.
		Read(ctx, data.Uuid.ValueString()).
		ReturnFieldsPlus(readableAttributesForView).
		ReturnAsObject(1).
		ProxySearch(config.GetProxySearch()).
		Execute()

	// If the resource is not found, try searching using Extensible Attributes
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound && r.ReadByExtAttrs(ctx, &data, resp) {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read View, got error: %s", err))
		return
	}

	res := apiRes.GetViewResponseObjectAsResult.GetResult()

	apiTerraformId, ok := (*res.ExtAttrs)[terraformInternalIDEA]
	if !ok {
		apiTerraformId.Value = ""
	}

	stateExtAttrs := ExpandExtAttrs(ctx, data.ExtAttrsAll, &diags)

	if associateInternalId == nil {
		if stateExtAttrs == nil {
			resp.Diagnostics.AddError(
				"Missing Internal ID",
				"Unable to read View because the internal ID (from extattrs_all) is missing or invalid.",
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while reading View due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ViewResource) ReadByExtAttrs(ctx context.Context, data *ViewModel, resp *resource.ReadResponse) bool {
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

	apiRes, _, err := r.client.DNSAPI.
		ViewAPI.
		List(ctx).
		Extattrfilter(idMap).
		ReturnAsObject(1).
		ReturnFieldsPlus(readableAttributesForView).
		ProxySearch(config.GetProxySearch()).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read View by extattrs, got error: %s", err))
		return true
	}

	results := apiRes.ListViewResponseObject.GetResult()

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

func (r *ViewResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data ViewModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	planExtAttrs := data.ExtAttrs
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

	apiRes, _, err := r.client.DNSAPI.
		ViewAPI.
		Update(ctx, data.Uuid.ValueString()).
		View(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForView).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update View, got error: %s", err))
		return
	}

	res := apiRes.UpdateViewResponseAsObject.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, planExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while update View due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

	if associateInternalId != nil {
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", nil)...)
	}
}

func (r *ViewResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ViewModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.DNSAPI.
		ViewAPI.
		Delete(ctx, data.Uuid.ValueString()).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete View, got error: %s", err))
		return
	}
}

func (r *ViewResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("uuid"), req.ID)...)
	resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", []byte("true"))...)
}

func (r *ViewResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data ViewModel

	// Retrieve the resource configuration
	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Check if filter_aaaa_list contains items with a ref field
	if !data.FilterAaaaList.IsNull() && !data.FilterAaaaList.IsUnknown() {
		var filterAaaaListItems []types.Object
		diags = data.FilterAaaaList.ElementsAs(ctx, &filterAaaaListItems, false)
		resp.Diagnostics.Append(diags...)

		hasRefInList := false
		for _, item := range filterAaaaListItems {
			itemMap := item.Attributes()

			// Check if ref field exists and is not empty
			if refAttr, ok := itemMap["ref"]; ok {
				refValue, _ := refAttr.(types.String)
				if !refValue.IsNull() && !refValue.IsUnknown() && refValue.ValueString() != "" {
					hasRefInList = true
					break
				}
			}
		}

		// If ref field is found, validate filter_aaaa value
		if hasRefInList {
			if !data.FilterAaaa.IsNull() && !data.FilterAaaa.IsUnknown() {
				filterAaaaValue := data.FilterAaaa.ValueString()
				if filterAaaaValue == "NO" {
					resp.Diagnostics.AddAttributeError(
						path.Root("filter_aaaa"),
						"Invalid Filter AAAA Configuration",
						"When 'ref' field is provided in filter_aaaa_list, filter_aaaa must be set to 'YES' or 'BREAK_DNSSEC', not 'NO'.",
					)
				}
			} else {
				resp.Diagnostics.AddAttributeError(
					path.Root("filter_aaaa"),
					"Missing Filter AAAA Configuration",
					"When 'ref' field is provided in filter_aaaa_list, filter_aaaa must be set to 'YES' or 'BREAK_DNSSEC', it cannot be null or empty.",
				)
			}
		}
	}
}
