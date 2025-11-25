package dtc

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

var readableAttributesForDtcRecordA = "auto_created,comment,disable,dtc_server,ipv4addr,ttl,use_ttl"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &DtcRecordAResource{}
var _ resource.ResourceWithImportState = &DtcRecordAResource{}

func NewDtcRecordAResource() resource.Resource {
	return &DtcRecordAResource{}
}

// DtcRecordAResource defines the resource implementation.
type DtcRecordAResource struct {
	client *niosclient.APIClient
}

func (r *DtcRecordAResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "dtc_record_a"
}

func (r *DtcRecordAResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "",
		Attributes:          DtcRecordAResourceSchemaAttributes,
	}
}

func (r *DtcRecordAResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *DtcRecordAResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DtcRecordAModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.DTCAPI.
		DtcRecordAAPI.
		Create(ctx).
		DtcRecordA(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForDtcRecordA).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create DtcRecordA, got error: %s", err))
		return
	}

	res := apiRes.CreateDtcRecordAResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DtcRecordAResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DtcRecordAModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.DTCAPI.
		DtcRecordAAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForDtcRecordA).
		ReturnAsObject(1).
		Execute()

		// Handle not found case
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			// Resource no longer exists, remove from state
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read DtcRecordA, got error: %s", err))
		return
	}

	res := apiRes.GetDtcRecordAResponseObjectAsResult.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DtcRecordAResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data DtcRecordAModel

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

	apiRes, _, err := r.client.DTCAPI.
		DtcRecordAAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		DtcRecordA(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForDtcRecordA).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update DtcRecordA, got error: %s", err))
		return
	}

	res := apiRes.UpdateDtcRecordAResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DtcRecordAResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DtcRecordAModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.DTCAPI.
		DtcRecordAAPI.
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete DtcRecordA, got error: %s", err))
		return
	}
}

func (r *DtcRecordAResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("ref"), req, resp)
}
