package grid

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

var readableAttributesForDistributionschedule = "active,start_time,time_zone,upgrade_groups"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &DistributionscheduleResource{}
var _ resource.ResourceWithImportState = &DistributionscheduleResource{}

func NewDistributionscheduleResource() resource.Resource {
	return &DistributionscheduleResource{}
}

// DistributionscheduleResource defines the resource implementation.
type DistributionscheduleResource struct {
	client *niosclient.APIClient
}

func (r *DistributionscheduleResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "grid_distributionschedule"
}

func (r *DistributionscheduleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a Distribution Schedule resource within the NIOS Grid.",
		Attributes:          DistributionscheduleResourceSchemaAttributes,
	}
}

func (r *DistributionscheduleResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// func (r *DistributionscheduleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
// 	var data DistributionscheduleModel

// 	// Read Terraform plan data into the model
// 	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

// 	if resp.Diagnostics.HasError() {
// 		return
// 	}

// 	apiRes, _, err := r.client.GridAPI.
// 		DistributionscheduleAPI.
// 		Create(ctx).
// 		Distributionschedule(*data.Expand(ctx, &resp.Diagnostics)).
// 		ReturnFieldsPlus(readableAttributesForDistributionschedule).
// 		ReturnAsObject(1).
// 		Execute()
// 	if err != nil {
// 		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Distributionschedule, got error: %s", err))
// 		return
// 	}

// 	res := apiRes.CreateDistributionscheduleResponseAsObject.GetResult()

// 	data.Flatten(ctx, &res, &resp.Diagnostics)

// 	// Save data into Terraform state
// 	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
// }

func (r *DistributionscheduleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DistributionscheduleModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	listResp, _, err := r.client.GridAPI.
		DistributionscheduleAPI.
		List(ctx).
		ReturnAsObject(1).
		// ReturnFieldsPlus(readableAttributesForDistributionschedule).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to list DistributionSchedule: %s", err))
		return
	}

	list := listResp.ListDistributionscheduleResponseObject.GetResult()

	if len(list) == 0 {
		resp.Diagnostics.AddError("Not Found", "No DistributionSchedule object exists in this Grid")
		return
	}

	// Extract the singleton ref
	ref := list[0].GetRef()

	// Update it with desired plan
	apiRes, _, err := r.client.GridAPI.
		DistributionscheduleAPI.
		Update(ctx, utils.ExtractResourceRef(ref)).
		Distributionschedule(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForDistributionschedule).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update DistributionSchedule: %s", err))
		return
	}

	res := apiRes.UpdateDistributionscheduleResponseAsObject.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DistributionscheduleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DistributionscheduleModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.GridAPI.
		DistributionscheduleAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForDistributionschedule).
		ReturnAsObject(1).
		Execute()

	// Handle not found case
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			// Resource no longer exists, remove from state
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Distributionschedule, got error: %s", err))
		return
	}

	res := apiRes.GetDistributionscheduleResponseObjectAsResult.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DistributionscheduleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data DistributionscheduleModel

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

	apiRes, _, err := r.client.GridAPI.
		DistributionscheduleAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Distributionschedule(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForDistributionschedule).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Distributionschedule, got error: %s", err))
		return
	}

	res := apiRes.UpdateDistributionscheduleResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DistributionscheduleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// var data DistributionscheduleModel

	// // Read Terraform prior state data into the model
	// resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	// if resp.Diagnostics.HasError() {
	// 	return
	// }

	// httpRes, err := r.client.GridAPI.
	// 	DistributionscheduleAPI.
	// 	Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
	// 	Execute()
	// if err != nil {
	// 	if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
	// 		return
	// 	}
	// 	resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Distributionschedule, got error: %s", err))
	// 	return
	// }

	// DistributionSchedule cannot be deleted, so just clear state
	resp.State.RemoveResource(ctx)
}

func (r *DistributionscheduleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("ref"), req, resp)
}
