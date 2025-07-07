package dns

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

	niosclient "github.com/Infoblox-CTO/infoblox-nios-go-client/client"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/utils"
)

var readableAttributesForZoneForward = "address,comment,disable,disable_ns_generation,display_domain,dns_fqdn,extattrs,external_ns_group," +
	"forward_to,forwarders_only,forwarding_servers,fqdn,locked,locked_by,mask_prefix,ms_ad_integrated,ms_ddns_mode,ms_managed," +
	"ms_read_only,ms_sync_master_name,ns_group,parent,prefix,using_srg_associations,view,zone_format"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ZoneForwardResource{}
var _ resource.ResourceWithImportState = &ZoneForwardResource{}

func NewZoneForwardResource() resource.Resource {
	return &ZoneForwardResource{}
}

// ZoneForwardResource defines the resource implementation.
type ZoneForwardResource struct {
	client *niosclient.APIClient
}

func (r *ZoneForwardResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "dns_zone_forward"
}

func (r *ZoneForwardResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages DNS Forward Zones.",
		Attributes:          ZoneForwardResourceSchemaAttributes,
	}
}

func (r *ZoneForwardResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ZoneForwardResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var diags diag.Diagnostics
	var data ZoneForwardModel

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

	apiRes, _, err := r.client.DNSAPI.
		ZoneForwardAPI.
		Create(ctx).
		ZoneForward(*data.Expand(ctx, &resp.Diagnostics, true)).
		ReturnFieldsPlus(readableAttributesForZoneForward).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create ZoneForward, got error: %s", err))
		return
	}

	res := apiRes.CreateZoneForwardResponseAsObject.GetResult()
	res.ExtAttrs, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while create ZoneForward due inherited Extensible attributes, got error: %s", err))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZoneForwardResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var diags diag.Diagnostics
	var data ZoneForwardModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.DNSAPI.
		ZoneForwardAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForZoneForward).
		ReturnAsObject(1).
		Execute()

	// If the resource is not found, try searching using Extensible Attributes
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound && r.ReadByExtAttrs(ctx, &data, resp) {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read ZoneForward, got error: %s", err))
		return
	}

	res := apiRes.GetZoneForwardResponseObjectAsResult.GetResult()
	if res.ExtAttrs == nil {
		resp.Diagnostics.AddError(
			"Missing Extensible Attributes",
			"Unable to read ZoneForward because no extensible attributes were returned from the API.",
		)
		return
	}

	res.ExtAttrs, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while reading ZoneForward due inherited Extensible attributes, got error: %s", diags))
		return
	}

	apiTerraformId, ok := (*res.ExtAttrs)["Terraform Internal ID"]
	if !ok {
		resp.Diagnostics.AddError(
			"Missing Terraform internal id Attributes",
			"Unable to read ZoneForward because terraform internal id does not exist.",
		)
		return
	}

	stateExtAttrs := ExpandExtAttr(ctx, data.ExtAttrsAll, &diags)
	if stateExtAttrs == nil {
		resp.Diagnostics.AddError(
			"Missing Internal ID",
			"Unable to read ZoneForward because the internal ID (from extattrs_all) is missing or invalid.",
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

func (r *ZoneForwardResource) ReadByExtAttrs(ctx context.Context, data *ZoneForwardModel, resp *resource.ReadResponse) bool {
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

	apiRes, _, err := r.client.DNSAPI.
		ZoneForwardAPI.
		List(ctx).
		Extattrfilter(idMap).
		ReturnAsObject(1).
		ReturnFieldsPlus(readableAttributesForZoneForward).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read ZoneForward by extattrs, got error: %s", err))
		return true
	}

	results := apiRes.ListZoneForwardResponseObject.GetResult()

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

func (r *ZoneForwardResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data ZoneForwardModel

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

	apiRes, _, err := r.client.DNSAPI.
		ZoneForwardAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ZoneForward(*data.Expand(ctx, &resp.Diagnostics, false)).
		ReturnFieldsPlus(readableAttributesForZoneForward).
		ReturnAsObject(1).
		Execute()

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update ZoneForward, got error: %s", err))
		return
	}

	res := apiRes.UpdateZoneForwardResponseAsObject.GetResult()

	res.ExtAttrs, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while update ZoneForward due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZoneForwardResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ZoneForwardModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.DNSAPI.
		ZoneForwardAPI.
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete ZoneForward, got error: %s", err))
		return
	}
}

func (r *ZoneForwardResource) addInternalIDToExtAttrs(ctx context.Context, data *ZoneForwardModel) error {
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

	r.client.DNSAPI.APIClient.Cfg.DefaultExtAttrs = map[string]struct{ Value string }{
		"Terraform Internal ID": {Value: internalId},
	}

	return nil
}

func (r *ZoneForwardResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("ref"), req, resp)
}
