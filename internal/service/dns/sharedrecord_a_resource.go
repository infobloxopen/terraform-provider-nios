package dns

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

var readableAttributesForSharedrecordA = "comment,disable,dns_name,extattrs,ipv4addr,name,shared_record_group,ttl,use_ttl"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SharedrecordAResource{}
var _ resource.ResourceWithImportState = &SharedrecordAResource{}

func NewSharedrecordAResource() resource.Resource {
	return &SharedrecordAResource{}
}

// SharedrecordAResource defines the resource implementation.
type SharedrecordAResource struct {
	client *niosclient.APIClient
}

func (r *SharedrecordAResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "dns_sharedrecord_a"
}

func (r *SharedrecordAResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a DNS Shared A Record.",
		Attributes:          SharedrecordAResourceSchemaAttributes,
	}
}

func (r *SharedrecordAResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *SharedrecordAResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var diags diag.Diagnostics
	var data SharedrecordAModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.DNSAPI.
		SharedrecordAAPI.
		Create(ctx).
		SharedrecordA(*data.Expand(ctx, &resp.Diagnostics, true)).
		ReturnFieldsPlus(readableAttributesForSharedrecordA).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create SharedrecordA, got error: %s", err))
		return
	}

	res := apiRes.CreateSharedrecordAResponseAsObject.GetResult()
	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while create SharedrecordMx due inherited Extensible attributes, got error: %s", err))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SharedrecordAResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var diags diag.Diagnostics
	var data SharedrecordAModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.DNSAPI.
		SharedrecordAAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForSharedrecordA).
		ReturnAsObject(1).
		Execute()

	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			// Resource no longer exists, remove from state
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read SharedrecordA, got error: %s", err))
		return
	}

	res := apiRes.GetSharedrecordAResponseObjectAsResult.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while reading SharedrecordMx due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SharedrecordAResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data SharedrecordAModel

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

	// Add Inherited Extensible Attributes
	data.ExtAttrs, diags = AddInheritedExtAttrs(ctx, data.ExtAttrs, data.ExtAttrsAll)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	apiRes, _, err := r.client.DNSAPI.
		SharedrecordAAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		SharedrecordA(*data.Expand(ctx, &resp.Diagnostics, false)).
		ReturnFieldsPlus(readableAttributesForSharedrecordA).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update SharedrecordA, got error: %s", err))
		return
	}

	res := apiRes.UpdateSharedrecordAResponseAsObject.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, planExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while update SharedrecordMx due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SharedrecordAResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SharedrecordAModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.DNSAPI.
		SharedrecordAAPI.
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete SharedrecordA, got error: %s", err))
		return
	}
}

func (r *SharedrecordAResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("ref"), req.ID)...)
}
