package parentalcontrol

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"

	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForParentalcontrolSubscribersite = "abss,api_members,api_port,block_size,blocking_ipv4_vip1,blocking_ipv4_vip2,blocking_ipv6_vip1,blocking_ipv6_vip2,comment,dca_sub_bw_list,dca_sub_query_count,enable_global_allow_list_rpz,enable_rpz_filtering_bypass,extattrs,first_port,global_allow_list_rpz,maximum_subscribers,members,msps,name,nas_gateways,nas_port,proxy_rpz_passthru,spms,stop_anycast,strict_nat,subscriber_collection_type"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ParentalcontrolSubscribersiteResource{}
var _ resource.ResourceWithImportState = &ParentalcontrolSubscribersiteResource{}

func NewParentalcontrolSubscribersiteResource() resource.Resource {
	return &ParentalcontrolSubscribersiteResource{}
}

// ParentalcontrolSubscribersiteResource defines the resource implementation.
type ParentalcontrolSubscribersiteResource struct {
	client *niosclient.APIClient
}

func (r *ParentalcontrolSubscribersiteResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "parentalcontrol_subscribersite"
}

func (r *ParentalcontrolSubscribersiteResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "",
		Attributes:          ParentalcontrolSubscribersiteResourceSchemaAttributes,
	}
}

func (r *ParentalcontrolSubscribersiteResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ParentalcontrolSubscribersiteResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var diags diag.Diagnostics
	var data ParentalcontrolSubscribersiteModel

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

	apiRes, _, err := r.client.ParentalControlAPI.
		ParentalcontrolSubscribersiteAPI.
		Create(ctx).
		ParentalcontrolSubscribersite(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForParentalcontrolSubscribersite).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create ParentalcontrolSubscribersite, got error: %s", err))
		return
	}

	res := apiRes.CreateParentalcontrolSubscribersiteResponseAsObject.GetResult()
	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while create ParentalcontrolSubscribersite due inherited Extensible attributes, got error: %s", err))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ParentalcontrolSubscribersiteResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var diags diag.Diagnostics
	var data ParentalcontrolSubscribersiteModel

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

	apiRes, httpRes, err := r.client.ParentalControlAPI.
		ParentalcontrolSubscribersiteAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForParentalcontrolSubscribersite).
		ReturnAsObject(1).
		Execute()

	// If the resource is not found, try searching using Extensible Attributes
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound && r.ReadByExtAttrs(ctx, &data, resp) {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read ParentalcontrolSubscribersite, got error: %s", err))
		return
	}

	res := apiRes.GetParentalcontrolSubscribersiteResponseObjectAsResult.GetResult()

	apiTerraformId, ok := (*res.ExtAttrs)[terraformInternalIDEA]
	if !ok {
		apiTerraformId.Value = ""
	}

	if associateInternalId == nil {
		stateExtAttrs := ExpandExtAttrs(ctx, data.ExtAttrsAll, &diags)
		if stateExtAttrs == nil {
			resp.Diagnostics.AddError(
				"Missing Internal ID",
				"Unable to read ParentalcontrolSubscribersite because the internal ID (from extattrs_all) is missing or invalid.",
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while reading ParentalcontrolSubscribersite due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ParentalcontrolSubscribersiteResource) ReadByExtAttrs(ctx context.Context, data *ParentalcontrolSubscribersiteModel, resp *resource.ReadResponse) bool {
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

	apiRes, _, err := r.client.ParentalControlAPI.
		ParentalcontrolSubscribersiteAPI.
		List(ctx).
		Extattrfilter(idMap).
		ReturnAsObject(1).
		ReturnFieldsPlus(readableAttributesForParentalcontrolSubscribersite).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read ParentalcontrolSubscribersite by extattrs, got error: %s", err))
		return true
	}

	results := apiRes.ListParentalcontrolSubscribersiteResponseObject.GetResult()

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

func (r *ParentalcontrolSubscribersiteResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data ParentalcontrolSubscribersiteModel

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

	apiRes, _, err := r.client.ParentalControlAPI.
		ParentalcontrolSubscribersiteAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ParentalcontrolSubscribersite(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForParentalcontrolSubscribersite).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update ParentalcontrolSubscribersite, got error: %s", err))
		return
	}

	res := apiRes.UpdateParentalcontrolSubscribersiteResponseAsObject.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, planExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while update ParentalcontrolSubscribersite due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

	if associateInternalId != nil {
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", nil)...)
	}
}

func (r *ParentalcontrolSubscribersiteResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ParentalcontrolSubscribersiteModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.ParentalControlAPI.
		ParentalcontrolSubscribersiteAPI.
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete ParentalcontrolSubscribersite, got error: %s", err))
		return
	}
}

func (r *ParentalcontrolSubscribersiteResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("ref"), req.ID)...)
	resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", []byte("true"))...)
}
