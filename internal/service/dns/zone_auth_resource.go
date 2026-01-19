package dns

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"

	"github.com/infobloxopen/terraform-provider-nios/internal/config"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForZoneAuth = "address,allow_active_dir,allow_fixed_rrset_order,allow_gss_tsig_for_underscore_zone,allow_gss_tsig_zone_updates,allow_query,allow_transfer,allow_update,allow_update_forwarding,aws_rte53_zone_info,cloud_info,comment,copy_xfer_to_notify,create_underscore_zones,ddns_force_creation_timestamp_update,ddns_principal_group,ddns_principal_tracking,ddns_restrict_patterns,ddns_restrict_patterns_list,ddns_restrict_protected,ddns_restrict_secure,ddns_restrict_static,disable,disable_forwarding,display_domain,dns_fqdn,dns_integrity_enable,dns_integrity_frequency,dns_integrity_member,dns_integrity_verbose_logging,dns_soa_email,dnssec_key_params,dnssec_keys,dnssec_ksk_rollover_date,dnssec_zsk_rollover_date,effective_check_names_policy,effective_record_name_policy,extattrs,external_primaries,external_secondaries,fqdn,grid_primary,grid_primary_shared_with_ms_parent_delegation,grid_secondaries,is_dnssec_enabled,is_dnssec_signed,is_multimaster,last_queried,last_queried_acl,locked,locked_by,mask_prefix,member_soa_mnames,member_soa_serials,ms_ad_integrated,ms_allow_transfer,ms_allow_transfer_mode,ms_dc_ns_record_creation,ms_ddns_mode,ms_managed,ms_primaries,ms_read_only,ms_secondaries,ms_sync_disabled,ms_sync_master_name,network_associations,network_view,notify_delay,ns_group,parent,prefix,primary_type,record_name_policy,records_monitored,rr_not_queried_enabled_time,scavenging_settings,soa_default_ttl,soa_email,soa_expire,soa_negative_ttl,soa_refresh,soa_retry,soa_serial_number,srgs,update_forwarding,use_allow_active_dir,use_allow_query,use_allow_transfer,use_allow_update,use_allow_update_forwarding,use_check_names_policy,use_copy_xfer_to_notify,use_ddns_force_creation_timestamp_update,use_ddns_patterns_restriction,use_ddns_principal_security,use_ddns_restrict_protected,use_ddns_restrict_static,use_dnssec_key_params,use_external_primary,use_grid_zone_timer,use_import_from,use_notify_delay,use_record_name_policy,use_scavenging_settings,use_soa_email,using_srg_associations,view,zone_format,zone_not_queried_enabled_time"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ZoneAuthResource{}
var _ resource.ResourceWithImportState = &ZoneAuthResource{}
var _ resource.ResourceWithValidateConfig = &ZoneAuthResource{}

func NewZoneAuthResource() resource.Resource {
	return &ZoneAuthResource{}
}

// ZoneAuthResource defines the resource implementation.
type ZoneAuthResource struct {
	client *niosclient.APIClient
}

func (r *ZoneAuthResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "dns_zone_auth"
}

func (r *ZoneAuthResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages Auth Zones.",
		Attributes:          ZoneAuthResourceSchemaAttributes,
	}
}

func (r *ZoneAuthResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ZoneAuthResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {

	var data ZoneAuthModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !data.UseGridZoneTimer.IsNull() && !data.UseGridZoneTimer.ValueBool() {
		if !data.SoaDefaultTtl.IsNull() || !data.SoaExpire.IsNull() || !data.SoaNegativeTtl.IsNull() || !data.SoaRefresh.IsNull() || !data.SoaRetry.IsNull() {
			resp.Diagnostics.AddError(
				"SOA Values Not Allowed",
				"When grid_zone_timer is set to false, the SOA Values (soa_default_ttl, soa_expire, soa_negative_ttl, soa_refresh, soa_retry) will reset to their default values. And hence they should not be set in the configuration. Either remove these values or set use_grid_zone_timer = true.",
			)
		}
	}

	// Validation for mutually exclusive primary servers
	specifiedPrimaries := []string{}

	if !data.GridPrimary.IsNull() && !data.GridPrimary.IsUnknown() {
		specifiedPrimaries = append(specifiedPrimaries, "grid_primary")
	}

	if !data.ExternalPrimaries.IsNull() && !data.ExternalPrimaries.IsUnknown() {
		specifiedPrimaries = append(specifiedPrimaries, "external_primaries")
	}

	if !data.MsPrimaries.IsNull() && !data.MsPrimaries.IsUnknown() {
		specifiedPrimaries = append(specifiedPrimaries, "ms_primaries")
	}

	// If more than one primary server is specified, raise an error
	if len(specifiedPrimaries) > 1 {
		resp.Diagnostics.AddError(
			"Conflicting Primary Servers",
			fmt.Sprintf(
				"Only one of grid_primary, external_primaries, or ms_primaries can be specified. Found: %s.",
				strings.Join(specifiedPrimaries, ", "),
			),
		)
		return
	}

	if !data.GridSecondaries.IsNull() && !data.GridSecondaries.IsUnknown() ||
		!data.ExternalSecondaries.IsNull() && !data.ExternalSecondaries.IsUnknown() ||
		!data.MsSecondaries.IsNull() && !data.MsSecondaries.IsUnknown() {
		if len(specifiedPrimaries) == 0 || len(specifiedPrimaries) > 1 {
			resp.Diagnostics.AddError(
				"Secondary Server Requires Exactly One Primary Server",
				"When secondary servers (grid_secondaries, external_secondaries, or ms_secondaries) are specified, exactly one primary server (grid_primary, external_primaries, or ms_primaries) is required.",
			)
		}
	}
}

func (r *ZoneAuthResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var diags diag.Diagnostics
	var data ZoneAuthModel

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
		ZoneAuthAPI.
		Create(ctx).
		ZoneAuth(*data.Expand(ctx, &resp.Diagnostics, true)).
		ReturnFieldsPlus(readableAttributesForZoneAuth).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create ZoneAuth, got error: %s", err))
		return
	}

	res := apiRes.CreateZoneAuthResponseAsObject.GetResult()
	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while create ZoneAuth due inherited Extensible attributes, got error: %s", err))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZoneAuthResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var diags diag.Diagnostics
	var data ZoneAuthModel

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
		ZoneAuthAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForZoneAuth).
		ReturnAsObject(1).
		ProxySearch(config.GetProxySearch()).
		Execute()

	// If the resource is not found, try searching using Extensible Attributes
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound && r.ReadByExtAttrs(ctx, &data, resp) {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read ZoneAuth, got error: %s", err))
		return
	}

	res := apiRes.GetZoneAuthResponseObjectAsResult.GetResult()

	apiTerraformId, ok := (*res.ExtAttrs)[terraformInternalIDEA]
	if !ok {
		apiTerraformId.Value = ""
	}

	if associateInternalId == nil {
		stateExtAttrs := ExpandExtAttrs(ctx, data.ExtAttrsAll, &diags)
		if stateExtAttrs == nil {
			resp.Diagnostics.AddError(
				"Missing Internal ID",
				"Unable to read ZoneAuth because the internal ID (from extattrs_all) is missing or invalid.",
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while reading ZoneAuth due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZoneAuthResource) ReadByExtAttrs(ctx context.Context, data *ZoneAuthModel, resp *resource.ReadResponse) bool {
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
		ZoneAuthAPI.
		List(ctx).
		Extattrfilter(idMap).
		ReturnAsObject(1).
		ReturnFieldsPlus(readableAttributesForZoneAuth).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read ZoneAuth by extattrs, got error: %s", err))
		return true
	}

	results := apiRes.ListZoneAuthResponseObject.GetResult()

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

func (r *ZoneAuthResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data ZoneAuthModel

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

	apiRes, _, err := r.client.DNSAPI.
		ZoneAuthAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ZoneAuth(*data.Expand(ctx, &resp.Diagnostics, false)).
		ReturnFieldsPlus(readableAttributesForZoneAuth).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update ZoneAuth, got error: %s", err))
		return
	}

	res := apiRes.UpdateZoneAuthResponseAsObject.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, planExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while update ZoneAuth due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

	if associateInternalId != nil {
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", nil)...)
	}
}

func (r *ZoneAuthResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ZoneAuthModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.DNSAPI.
		ZoneAuthAPI.
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete ZoneAuth, got error: %s", err))
		return
	}
}

func (r *ZoneAuthResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("ref"), req.ID)...)
	resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", []byte("true"))...)
}
