package dhcp

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"

	"github.com/infobloxopen/terraform-provider-nios/internal/config"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForDhcpoptiondefinition = "code,name,space,type"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &DhcpoptiondefinitionResource{}
var _ resource.ResourceWithImportState = &DhcpoptiondefinitionResource{}

func NewDhcpoptiondefinitionResource() resource.Resource {
	return &DhcpoptiondefinitionResource{}
}

// DhcpoptiondefinitionResource defines the resource implementation.
type DhcpoptiondefinitionResource struct {
	client *niosclient.APIClient
}

func (r *DhcpoptiondefinitionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "dhcp_optiondefinition"
}

func (r *DhcpoptiondefinitionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a DHCP Option Definition.",
		Attributes:          DhcpoptiondefinitionResourceSchemaAttributes,
	}
}

func (r *DhcpoptiondefinitionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *DhcpoptiondefinitionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DhcpoptiondefinitionModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.DHCPAPI.
		DhcpoptiondefinitionAPI.
		Create(ctx).
		Dhcpoptiondefinition(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForDhcpoptiondefinition).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Dhcpoptiondefinition, got error: %s", err))
		return
	}

	res := apiRes.CreateDhcpoptiondefinitionResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DhcpoptiondefinitionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DhcpoptiondefinitionModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.DHCPAPI.
		DhcpoptiondefinitionAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForDhcpoptiondefinition).
		ReturnAsObject(1).
		ProxySearch(config.GetProxySearch()).
		Execute()

		// Handle not found case
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			// Resource no longer exists, remove from state
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Dhcpoptiondefinition, got error: %s", err))
		return
	}

	res := apiRes.GetDhcpoptiondefinitionResponseObjectAsResult.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DhcpoptiondefinitionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data DhcpoptiondefinitionModel
	var stateData DhcpoptiondefinitionModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

	if resp.Diagnostics.HasError() {
		return
	}

	diags = req.State.GetAttribute(ctx, path.Root("ref"), &data.Ref)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	// Update ref if option space has changed
	if !data.Space.Equal(stateData.Space) {
		r.updateRefIfOptionSpaceChanged(ctx, resp, &data, &stateData)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	apiRes, _, err := r.client.DHCPAPI.
		DhcpoptiondefinitionAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Dhcpoptiondefinition(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForDhcpoptiondefinition).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Dhcpoptiondefinition, got error: %s", err))
		return
	}

	res := apiRes.UpdateDhcpoptiondefinitionResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DhcpoptiondefinitionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DhcpoptiondefinitionModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.DHCPAPI.
		DhcpoptiondefinitionAPI.
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Dhcpoptiondefinition, got error: %s", err))
		return
	}
}

func (r *DhcpoptiondefinitionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("ref"), req, resp)
}

// updateRefIfOptionSpaceChanged updates the ref if the option space name changes by
// finding the option definition with the new space name and updating the data model accordingly.
func (r *DhcpoptiondefinitionResource) updateRefIfOptionSpaceChanged(ctx context.Context, resp *resource.UpdateResponse, data *DhcpoptiondefinitionModel, stateData *DhcpoptiondefinitionModel) {
	if resp.Diagnostics.HasError() {
		return
	}

	// Search for the option definition with the new space
	listApiRes, _, err := r.client.DHCPAPI.
		DhcpoptiondefinitionAPI.
		List(ctx).
		Filters(map[string]interface{}{
			"name":  stateData.Name.ValueString(),
			"space": data.Space.ValueString(),
			"code":  stateData.Code.ValueInt64(),
			"type":  stateData.Type.ValueString(),
		}).
		ReturnFieldsPlus(readableAttributesForDhcpoptiondefinition).
		ReturnAsObject(1).
		Execute()

	if err != nil {
		resp.State.RemoveResource(ctx)
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Dhcpoptiondefinition list, got error: %s", err))
		return
	}

	results := listApiRes.ListDhcpoptiondefinitionResponseObject.GetResult()

	if len(results) == 0 {
		return
	}

	data.Ref = types.StringValue(*results[0].Ref)
}
