package dtc

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

var readableAttributesForDtcPool = "extattrs,lb_preferred_method,auto_consolidated_monitors,availability,comment,consolidated_monitors,disable,health,lb_alternate_method,lb_alternate_topology,lb_dynamic_ratio_alternate,lb_dynamic_ratio_preferred,lb_preferred_topology,name,quorum,servers,ttl,use_ttl,monitors"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &DtcPoolResource{}
var _ resource.ResourceWithImportState = &DtcPoolResource{}

func NewDtcPoolResource() resource.Resource {
	return &DtcPoolResource{}
}

// DtcPoolResource defines the resource implementation.
type DtcPoolResource struct {
	client *niosclient.APIClient
}

func (r *DtcPoolResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "dtc_pool"
}

func (r *DtcPoolResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a DTC Pool",
		Attributes:          DtcPoolResourceSchemaAttributes,
	}
}

func (r *DtcPoolResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *DtcPoolResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var diags diag.Diagnostics
	var data DtcPoolModel

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

	apiRes, _, err := r.client.DTCAPI.
		DtcPoolAPI.
		Create(ctx).
		DtcPool(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForDtcPool).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create DtcPool, got error: %s", err))
		return
	}

	res := apiRes.CreateDtcPoolResponseAsObject.GetResult()
	res.ExtAttrs, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while create DtcPool due inherited Extensible attributes, got error: %s", err))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DtcPoolResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var diags diag.Diagnostics
	var data DtcPoolModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.DTCAPI.
		DtcPoolAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForDtcPool).
		ReturnAsObject(1).
		Execute()

	// If the resource is not found, try searching using Extensible Attributes
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound && r.ReadByExtAttrs(ctx, &data, resp) {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read DtcPool, got error: %s", err))
		return
	}

	res := apiRes.GetDtcPoolResponseObjectAsResult.GetResult()
	if res.ExtAttrs == nil {
		resp.Diagnostics.AddError(
			"Missing Extensible Attributes",
			"Unable to read DtcPool because no extensible attributes were returned from the API.",
		)
		return
	}

	res.ExtAttrs, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while reading DtcPool due inherited Extensible attributes, got error: %s", diags))
		return
	}

	apiTerraformId, ok := (*res.ExtAttrs)["Terraform Internal ID"]
	if !ok {
		resp.Diagnostics.AddError(
			"Missing Terraform internal id Attributes",
			"Unable to read DtcPool because terraform internal id does not exist.",
		)
		return
	}

	stateExtAttrs := ExpandExtAttr(ctx, data.ExtAttrsAll, &diags)
	if stateExtAttrs == nil {
		resp.Diagnostics.AddError(
			"Missing Internal ID",
			"Unable to read DtcPool because the internal ID (from extattrs_all) is missing or invalid.",
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

func (r *DtcPoolResource) ReadByExtAttrs(ctx context.Context, data *DtcPoolModel, resp *resource.ReadResponse) bool {
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

	apiRes, _, err := r.client.DTCAPI.
		DtcPoolAPI.
		List(ctx).
		Extattrfilter(idMap).
		ReturnAsObject(1).
		ReturnFieldsPlus(readableAttributesForDtcPool).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read DtcPool by extattrs, got error: %s", err))
		return true
	}

	results := apiRes.ListDtcPoolResponseObject.GetResult()

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

func (r *DtcPoolResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data DtcPoolModel

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

	apiRes, _, err := r.client.DTCAPI.
		DtcPoolAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		DtcPool(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForDtcPool).
		ReturnAsObject(1).
		Execute()

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update DtcPool, got error: %s", err))
		return
	}

	res := apiRes.UpdateDtcPoolResponseAsObject.GetResult()

	res.ExtAttrs, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while update DtcPool due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DtcPoolResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DtcPoolModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.DTCAPI.
		DtcPoolAPI.
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete DtcPool, got error: %s", err))
		return
	}
}

func (r *DtcPoolResource) addInternalIDToExtAttrs(ctx context.Context, data *DtcPoolModel) error {
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

	r.client.DTCAPI.APIClient.Cfg.DefaultExtAttrs = map[string]struct{ Value string }{
		"Terraform Internal ID": {Value: internalId},
	}

	return nil
}

func (r *DtcPoolResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("ref"), req, resp)
}
