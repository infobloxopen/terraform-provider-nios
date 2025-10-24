package grid

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"

	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForExtensibleattributedef = "allowed_object_types,comment,default_value,flags,list_values,max,min,name,namespace,type"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ExtensibleattributedefResource{}
var _ resource.ResourceWithImportState = &ExtensibleattributedefResource{}

func NewExtensibleattributedefResource() resource.Resource {
	return &ExtensibleattributedefResource{}
}

// ExtensibleattributedefResource defines the resource implementation.
type ExtensibleattributedefResource struct {
	client *niosclient.APIClient
}

func (r *ExtensibleattributedefResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "grid_extensibleattributedef"
}

func (r *ExtensibleattributedefResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages an Extensible Attribute definition.",
		Attributes:          ExtensibleattributedefResourceSchemaAttributes,
	}
}

func (r *ExtensibleattributedefResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ExtensibleattributedefResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ExtensibleattributedefModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.GridAPI.
		ExtensibleattributedefAPI.
		Create(ctx).
		Extensibleattributedef(*data.Expand(ctx, &resp.Diagnostics, true)).
		ReturnFieldsPlus(readableAttributesForExtensibleattributedef).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Extensibleattributedef, got error: %s", err))
		return
	}

	res := apiRes.CreateExtensibleattributedefResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ExtensibleattributedefResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ExtensibleattributedefModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.GridAPI.
		ExtensibleattributedefAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForExtensibleattributedef).
		ReturnAsObject(1).
		Execute()

	// Handle not found case
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			// Resource no longer exists, remove from state
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Extensibleattributedef, got error: %s", err))
		return
	}

	res := apiRes.GetExtensibleattributedefResponseObjectAsResult.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ExtensibleattributedefResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data ExtensibleattributedefModel

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
		ExtensibleattributedefAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Extensibleattributedef(*data.Expand(ctx, &resp.Diagnostics, false)).
		ReturnFieldsPlus(readableAttributesForExtensibleattributedef).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Extensibleattributedef, got error: %s", err))
		return
	}

	res := apiRes.UpdateExtensibleattributedefResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ExtensibleattributedefResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ExtensibleattributedefModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.GridAPI.
		ExtensibleattributedefAPI.
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Extensibleattributedef, got error: %s", err))
		return
	}
}

func (r *ExtensibleattributedefResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data ExtensibleattributedefModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	eaType := data.Type
	minValue := data.Min
	maxValue := data.Max

	// If type is unknown, we can't validate
	if eaType.IsUnknown() {
		return
	}

	typeValue := eaType.ValueString()

	// Check if min or max are set with incompatible types
	if (!minValue.IsNull() || !maxValue.IsNull()) && typeValue != "INTEGER" && typeValue != "STRING" {
		resp.Diagnostics.AddError(
			"Invalid Min/Max Configuration",
			fmt.Sprintf("The 'min' and 'max' attributes are only valid for INTEGER and STRING extensible attribute types, but type is %q. Remove min and max attributes or change the type to INTEGER or STRING.", typeValue),
		)
	}

	// Check if type is INTEGER, then default_value should be a valid integer
	if !data.DefaultValue.IsNull() && !data.DefaultValue.IsUnknown() {
		if typeValue == "INTEGER" {
			defaultValue := data.DefaultValue.ValueString()
			_, err := strconv.ParseInt(defaultValue, 10, 32)
			if err != nil {
				resp.Diagnostics.AddError(
					"Invalid Integer Default Value",
					fmt.Sprintf("The default_value '%s' is not a valid integer.", defaultValue),
				)
				return
			}
		}
	}
}

func (r *ExtensibleattributedefResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("ref"), req, resp)
}
