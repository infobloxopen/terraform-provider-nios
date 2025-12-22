package threatinsight

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

var readableAttributesForThreatinsightAllowlist = "comment,disable,fqdn,type"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ThreatinsightAllowlistResource{}
var _ resource.ResourceWithImportState = &ThreatinsightAllowlistResource{}

func NewThreatinsightAllowlistResource() resource.Resource {
	return &ThreatinsightAllowlistResource{}
}

// ThreatinsightAllowlistResource defines the resource implementation.
type ThreatinsightAllowlistResource struct {
	client *niosclient.APIClient
}

func (r *ThreatinsightAllowlistResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "threatinsight_allowlist"
}

func (r *ThreatinsightAllowlistResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "",
		Attributes:          ThreatinsightAllowlistResourceSchemaAttributes,
	}
}

func (r *ThreatinsightAllowlistResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ThreatinsightAllowlistResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ThreatinsightAllowlistModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.ThreatInsightAPI.
		ThreatinsightAllowlistAPI.
		Create(ctx).
		ThreatinsightAllowlist(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForThreatinsightAllowlist).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create ThreatinsightAllowlist, got error: %s", err))
		return
	}

	res := apiRes.CreateThreatinsightAllowlistResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ThreatinsightAllowlistResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ThreatinsightAllowlistModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.ThreatInsightAPI.
		ThreatinsightAllowlistAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForThreatinsightAllowlist).
		ReturnAsObject(1).
		Execute()

		// Handle not found case
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			// Resource no longer exists, remove from state
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read ThreatinsightAllowlist, got error: %s", err))
		return
	}

	res := apiRes.GetThreatinsightAllowlistResponseObjectAsResult.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ThreatinsightAllowlistResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data ThreatinsightAllowlistModel

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

	apiRes, _, err := r.client.ThreatInsightAPI.
		ThreatinsightAllowlistAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ThreatinsightAllowlist(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForThreatinsightAllowlist).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update ThreatinsightAllowlist, got error: %s", err))
		return
	}

	res := apiRes.UpdateThreatinsightAllowlistResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ThreatinsightAllowlistResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ThreatinsightAllowlistModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.ThreatInsightAPI.
		ThreatinsightAllowlistAPI.
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete ThreatinsightAllowlist, got error: %s", err))
		return
	}
}

func (r *ThreatinsightAllowlistResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("ref"), req, resp)
}
