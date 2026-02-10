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
)

var readableAttributesForZoneRp = "address,comment,disable,display_domain,dns_soa_email,extattrs,external_primaries,external_secondaries,fireeye_rule_mapping,fqdn,grid_primary,grid_secondaries,locked,locked_by,log_rpz,mask_prefix,member_soa_mnames,member_soa_serials,network_view,ns_group,parent,prefix,primary_type,record_name_policy,rpz_drop_ip_rule_enabled,rpz_drop_ip_rule_min_prefix_length_ipv4,rpz_drop_ip_rule_min_prefix_length_ipv6,rpz_last_updated_time,rpz_policy,rpz_priority,rpz_priority_end,rpz_severity,rpz_type,soa_default_ttl,soa_email,soa_expire,soa_negative_ttl,soa_refresh,soa_retry,soa_serial_number,substitute_name,use_external_primary,use_grid_zone_timer,use_log_rpz,use_record_name_policy,use_rpz_drop_ip_rule,use_soa_email,view"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ZoneRpResource{}
var _ resource.ResourceWithImportState = &ZoneRpResource{}
var _ resource.ResourceWithValidateConfig = &ZoneRpResource{}

func NewZoneRpResource() resource.Resource {
	return &ZoneRpResource{}
}

// ZoneRpResource defines the resource implementation.
type ZoneRpResource struct {
	client *niosclient.APIClient
}

func (r *ZoneRpResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "dns_zone_rp"
}

func (r *ZoneRpResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a DNS RPZ (Response Policy Zone).",
		Attributes:          ZoneRpResourceSchemaAttributes,
	}
}

func (r *ZoneRpResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ZoneRpResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {

	var data ZoneRpModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !data.UseGridZoneTimer.IsNull() && !data.UseGridZoneTimer.ValueBool() {
		if !data.SoaDefaultTtl.IsNull() || !data.SoaExpire.IsNull() || !data.SoaNegativeTtl.IsNull() || !data.SoaRefresh.IsNull() || !data.SoaRetry.IsNull() {
			resp.Diagnostics.AddError(
				"SOA Values Not Allowed",
				"When `use_grid_zone_timer` is set to false, the SOA Values (soa_default_ttl, soa_expire, soa_negative_ttl, soa_refresh, soa_retry) will reset to their default values. And hence they should not be set in the configuration. Either remove these values or set use_grid_zone_timer = true.",
			)
		}
	}

	// Validation for mutually exclusive primary servers
	var specifiedPrimaries []string

	if !data.GridPrimary.IsNull() && !data.GridPrimary.IsUnknown() {
		specifiedPrimaries = append(specifiedPrimaries, "grid_primary")
	}

	if !data.ExternalPrimaries.IsNull() && !data.ExternalPrimaries.IsUnknown() {
		specifiedPrimaries = append(specifiedPrimaries, "external_primaries")
	}

	// If more than one primary server is specified, raise an error
	if len(specifiedPrimaries) > 1 {
		resp.Diagnostics.AddError(
			"Conflicting Primary Servers",
			fmt.Sprintf(
				"Only one of grid_primary or external_primaries can be specified. Found: %s.",
				strings.Join(specifiedPrimaries, ", "),
			),
		)
		return
	}

	if !data.GridSecondaries.IsNull() && !data.GridSecondaries.IsUnknown() ||
		!data.ExternalSecondaries.IsNull() && !data.ExternalSecondaries.IsUnknown() {
		if len(specifiedPrimaries) != 1 {
			resp.Diagnostics.AddError(
				"Secondary Server Requires Exactly One Primary Server",
				"When secondary servers (grid_secondaries or  external_secondaries) are specified, exactly one primary server (grid_primary or external_primaries) is required.",
			)
		}
	}
}

func (r *ZoneRpResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var diags diag.Diagnostics
	var data ZoneRpModel

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
		ZoneRpAPI.
		Create(ctx).
		ZoneRp(*data.Expand(ctx, &resp.Diagnostics, true)).
		ReturnFieldsPlus(readableAttributesForZoneRp).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create ZoneRp, got error: %s", err))
		return
	}

	res := apiRes.CreateZoneRpResponseAsObject.GetResult()
	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while create ZoneRp due inherited Extensible attributes, got error: %s", err))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZoneRpResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var diags diag.Diagnostics
	var data ZoneRpModel

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
		ZoneRpAPI.
		Read(ctx, data.Uuid.ValueString()).
		ReturnFieldsPlus(readableAttributesForZoneRp).
		ReturnAsObject(1).
		ProxySearch(config.GetProxySearch()).
		Execute()

	// If the resource is not found, try searching using Extensible Attributes
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound && r.ReadByExtAttrs(ctx, &data, resp) {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read ZoneRp, got error: %s", err))
		return
	}

	res := apiRes.GetZoneRpResponseObjectAsResult.GetResult()

	apiTerraformId, ok := (*res.ExtAttrs)[terraformInternalIDEA]
	if !ok {
		apiTerraformId.Value = ""
	}

	if associateInternalId == nil {
		stateExtAttrs := ExpandExtAttrs(ctx, data.ExtAttrsAll, &diags)
		if stateExtAttrs == nil {
			resp.Diagnostics.AddError(
				"Missing Internal ID",
				"Unable to read ZoneRp because the internal ID (from extattrs_all) is missing or invalid.",
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while reading ZoneRp due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZoneRpResource) ReadByExtAttrs(ctx context.Context, data *ZoneRpModel, resp *resource.ReadResponse) bool {
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
		ZoneRpAPI.
		List(ctx).
		Extattrfilter(idMap).
		ReturnAsObject(1).
		ReturnFieldsPlus(readableAttributesForZoneRp).
		ProxySearch(config.GetProxySearch()).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read ZoneRp by extattrs, got error: %s", err))
		return true
	}

	results := apiRes.ListZoneRpResponseObject.GetResult()

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

func (r *ZoneRpResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data ZoneRpModel

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
		ZoneRpAPI.
		Update(ctx, data.Uuid.ValueString()).
		ZoneRp(*data.Expand(ctx, &resp.Diagnostics, false)).
		ReturnFieldsPlus(readableAttributesForZoneRp).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update ZoneRp, got error: %s", err))
		return
	}

	res := apiRes.UpdateZoneRpResponseAsObject.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, planExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while update ZoneRp due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

	if associateInternalId != nil {
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", nil)...)
	}
}

func (r *ZoneRpResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ZoneRpModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.DNSAPI.
		ZoneRpAPI.
		Delete(ctx, data.Uuid.ValueString()).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete ZoneRp, got error: %s", err))
		return
	}
}

func (r *ZoneRpResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("uuid"), req.ID)...)
	resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", []byte("true"))...)
}
