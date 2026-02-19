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
	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/config"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForMemberdfp = "dfp_forward_first,host_name,is_dfp_override"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &MemberdfpResource{}
var _ resource.ResourceWithImportState = &MemberdfpResource{}

func NewMemberdfpResource() resource.Resource {
	return &MemberdfpResource{}
}

// MemberdfpResource defines the resource implementation.
type MemberdfpResource struct {
	client *niosclient.APIClient
}

func (r *MemberdfpResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "grid_memberdfp"
}

func (r *MemberdfpResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a Memberdfp resource object.",
		Attributes:          MemberdfpResourceSchemaAttributes,
	}
}

func (r *MemberdfpResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *MemberdfpResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data MemberdfpModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Validate that host_name is provided
	if data.HostName.IsNull() || data.HostName.IsUnknown() {
		resp.Diagnostics.AddError(
			"Missing Required Field",
			"host_name is required to identify the member",
		)
		return
	}

	// List all memberdfp objects (host_name is not searchable via API)
	listResp, _, err := r.client.GridAPI.
		MemberdfpAPI.
		List(ctx).
		ReturnAsObject(1).
		ReturnFieldsPlus(readableAttributesForMemberdfp).
		ProxySearch(config.GetProxySearch()).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to list Memberdfp: %s", err))
		return
	}

	list := listResp.ListMemberdfpResponseObject.GetResult()

	// Find the memberdfp object with matching host_name
	var listObj *grid.Memberdfp
	targetHostName := data.HostName.ValueString()
	for i := range list {
		if list[i].GetHostName() == targetHostName {
			listObj = &list[i]
			break
		}
	}

	if listObj == nil {
		resp.Diagnostics.AddError(
			"Not Found",
			fmt.Sprintf("No Memberdfp object exists for host_name: %s. The member must exist in the grid before configuring its DFP settings.", targetHostName),
		)
		return
	}

	// Update it with desired plan
	apiRes, _, err := r.client.GridAPI.
		MemberdfpAPI.
		Update(ctx, utils.ExtractResourceRef(listObj.GetRef())).
		Memberdfp(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForMemberdfp).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Memberdfp: %s", err))
		return
	}

	res := apiRes.UpdateMemberdfpResponseAsObject.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MemberdfpResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data MemberdfpModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.GridAPI.
		MemberdfpAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForMemberdfp).
		ReturnAsObject(1).
		ProxySearch(config.GetProxySearch()).
		Execute()

	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Memberdfp, got error: %s", err))
		return
	}

	res := apiRes.GetMemberdfpResponseObjectAsResult.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MemberdfpResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data MemberdfpModel

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
		MemberdfpAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Memberdfp(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForMemberdfp).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Memberdfp, got error: %s", err))
		return
	}

	res := apiRes.UpdateMemberdfpResponseAsObject.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MemberdfpResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Memberdfp cannot be deleted (tied to grid member lifecycle), just clear state
	resp.State.RemoveResource(ctx)
}

func (r *MemberdfpResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("ref"), req.ID)...)
	resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", []byte("true"))...)
}
