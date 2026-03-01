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

	"github.com/infobloxopen/terraform-provider-nios/internal/config"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForUpgradeschedule = "active,start_time,time_zone,upgrade_groups"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &UpgradescheduleResource{}
var _ resource.ResourceWithImportState = &UpgradescheduleResource{}
var _ resource.ResourceWithValidateConfig = &UpgradescheduleResource{}

func NewUpgradescheduleResource() resource.Resource {
	return &UpgradescheduleResource{}
}

// UpgradescheduleResource defines the resource implementation.
type UpgradescheduleResource struct {
	client *niosclient.APIClient
}

func (r *UpgradescheduleResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "grid_upgradeschedule"
}

func (r *UpgradescheduleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages an Upgrade Schedule",
		Attributes:          UpgradescheduleResourceSchemaAttributes,
	}
}

func (r *UpgradescheduleResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *UpgradescheduleResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data UpgradescheduleModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !data.UpgradeGroups.IsNull() && !data.UpgradeGroups.IsUnknown() {
		var groups []UpgradescheduleUpgradeGroupsModel
		diags := data.UpgradeGroups.ElementsAs(ctx, &groups, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		for idx, group := range groups {
			if group.Name.IsNull() || group.Name.IsUnknown() {
				resp.Diagnostics.AddAttributeError(
					path.Root("upgrade_groups"),
					"Invalid upgrade_groups.name",
					fmt.Sprintf("upgrade_groups[%d].name must be set", idx),
				)
			}
		}
	}
}

func (r *UpgradescheduleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data UpgradescheduleModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	listRes, _, err := r.client.GridAPI.
		UpgradescheduleAPI.
		List(ctx).
		ReturnAsObject(1).
		ReturnFieldsPlus(readableAttributesForUpgradeschedule).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to list Upgradeschedule, got error: %s", err))
		return
	}

	list := listRes.ListUpgradescheduleResponseObject.GetResult()

	if len(list) == 0 {
		resp.Diagnostics.AddError("Not Found", "No UpgradeSchedule object exists in this Grid")
		return
	}

	// Extract the singleton ref
	listObj := list[0]

	// Update it with desired plan
	apiRes, _, err := r.client.GridAPI.
		UpgradescheduleAPI.
		Update(ctx, utils.ExtractResourceRef(listObj.GetRef())).
		Upgradeschedule(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForUpgradeschedule).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update UpgradeSchedule: %s", err))
		return
	}

	res := apiRes.UpdateUpgradescheduleResponseAsObject.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *UpgradescheduleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data UpgradescheduleModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.GridAPI.
		UpgradescheduleAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForUpgradeschedule).
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Upgradeschedule, got error: %s", err))
		return
	}

	res := apiRes.GetUpgradescheduleResponseObjectAsResult.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *UpgradescheduleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data UpgradescheduleModel

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
		UpgradescheduleAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Upgradeschedule(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForUpgradeschedule).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Upgradeschedule, got error: %s", err))
		return
	}

	res := apiRes.UpdateUpgradescheduleResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *UpgradescheduleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// UpgradeSchedule cannot be deleted, so just clear state
	resp.State.RemoveResource(ctx)
}

func (r *UpgradescheduleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("ref"), req, resp)
}
