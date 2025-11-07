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

var readableAttributesForParentalcontrolBlockingpolicy = "name,value"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ParentalcontrolBlockingpolicyResource{}
var _ resource.ResourceWithImportState = &ParentalcontrolBlockingpolicyResource{}

func NewParentalcontrolBlockingpolicyResource() resource.Resource {
	return &ParentalcontrolBlockingpolicyResource{}
}

// ParentalcontrolBlockingpolicyResource defines the resource implementation.
type ParentalcontrolBlockingpolicyResource struct {
	client *niosclient.APIClient
}

func (r *ParentalcontrolBlockingpolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "parentalcontrol_blockingpolicy"
}

func (r *ParentalcontrolBlockingpolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "",
		Attributes:          ParentalcontrolBlockingpolicyResourceSchemaAttributes,
	}
}

func (r *ParentalcontrolBlockingpolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ParentalcontrolBlockingpolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ParentalcontrolBlockingpolicyModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.ParentalControlAPI.
		ParentalcontrolBlockingpolicyAPI.
		Create(ctx).
		ParentalcontrolBlockingpolicy(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForParentalcontrolBlockingpolicy).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create ParentalcontrolBlockingpolicy, got error: %s", err))
		return
	}

	res := apiRes.CreateParentalcontrolBlockingpolicyResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ParentalcontrolBlockingpolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ParentalcontrolBlockingpolicyModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.ParentalControlAPI.
		ParentalcontrolBlockingpolicyAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForParentalcontrolBlockingpolicy).
		ReturnAsObject(1).
		Execute()

		// Handle not found case
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			// Resource no longer exists, remove from state
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read ParentalcontrolBlockingpolicy, got error: %s", err))
		return
	}

	res := apiRes.GetParentalcontrolBlockingpolicyResponseObjectAsResult.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ParentalcontrolBlockingpolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data ParentalcontrolBlockingpolicyModel

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

	apiRes, _, err := r.client.ParentalControlAPI.
		ParentalcontrolBlockingpolicyAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ParentalcontrolBlockingpolicy(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForParentalcontrolBlockingpolicy).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update ParentalcontrolBlockingpolicy, got error: %s", err))
		return
	}

	res := apiRes.UpdateParentalcontrolBlockingpolicyResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ParentalcontrolBlockingpolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ParentalcontrolBlockingpolicyModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.ParentalControlAPI.
		ParentalcontrolBlockingpolicyAPI.
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete ParentalcontrolBlockingpolicy, got error: %s", err))
		return
	}
}

func (r *ParentalcontrolBlockingpolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("ref"), req, resp)
}
