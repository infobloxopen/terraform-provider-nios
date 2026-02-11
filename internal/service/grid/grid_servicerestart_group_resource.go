package grid

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
	internaltypes "github.com/infobloxopen/terraform-provider-nios/internal/types"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForGridServicerestartGroup = "comment,extattrs,is_default,last_updated_time,members,mode,name,position,recurring_schedule,requests,service,status"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &GridServicerestartGroupResource{}
var _ resource.ResourceWithImportState = &GridServicerestartGroupResource{}
var _ resource.ResourceWithValidateConfig = &GridServicerestartGroupResource{}

func NewGridServicerestartGroupResource() resource.Resource {
	return &GridServicerestartGroupResource{}
}

// GridServicerestartGroupResource defines the resource implementation.
type GridServicerestartGroupResource struct {
	client *niosclient.APIClient
}

func (r *GridServicerestartGroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "grid_servicerestart_group"
}

func (r *GridServicerestartGroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a GRID Service Restart Group.",
		Attributes:          GridServicerestartGroupResourceSchemaAttributes,
	}
}

func (r *GridServicerestartGroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *GridServicerestartGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var diags diag.Diagnostics
	var data GridServicerestartGroupModel

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

	apiRes, _, err := r.client.GridAPI.
		GridServicerestartGroupAPI.
		Create(ctx).
		GridServicerestartGroup(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForGridServicerestartGroup).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create GridServicerestartGroup, got error: %s", err))
		return
	}

	res := apiRes.CreateGridServicerestartGroupResponseAsObject.GetResult()
	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while create GridServicerestartGroup due inherited Extensible attributes, got error: %s", err))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GridServicerestartGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var diags diag.Diagnostics
	var data GridServicerestartGroupModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.GridAPI.
		GridServicerestartGroupAPI.
		Read(ctx, data.Uuid.ValueString()).
		ReturnFieldsPlus(readableAttributesForGridServicerestartGroup).
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read GridServicerestartGroup, got error: %s", err))
		return
	}

	res := apiRes.GetGridServicerestartGroupResponseObjectAsResult.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while reading GridServicerestartGroup due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}


func (r *GridServicerestartGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data GridServicerestartGroupModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	planExtAttrs := data.ExtAttrs
	diags = req.State.GetAttribute(ctx, path.Root("uuid"), &data.Uuid)
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

	apiRes, _, err := r.client.GridAPI.
		GridServicerestartGroupAPI.
		Update(ctx, data.Uuid.ValueString()).
		GridServicerestartGroup(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForGridServicerestartGroup).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update GridServicerestartGroup, got error: %s", err))
		return
	}

	res := apiRes.UpdateGridServicerestartGroupResponseAsObject.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, planExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while update GridServicerestartGroup due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

	if associateInternalId != nil {
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", nil)...)
	}
}

func (r *GridServicerestartGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data GridServicerestartGroupModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.GridAPI.
		GridServicerestartGroupAPI.
		Delete(ctx, data.Uuid.ValueString()).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete GridServicerestartGroup, got error: %s", err))
		return
	}
}

func (r *GridServicerestartGroupResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data GridServicerestartGroupModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	recurringScheduleAttr := data.RecurringSchedule.Attributes()
	if len(recurringScheduleAttr) == 0 {
		return
	}

	servicesAttr := data.RecurringSchedule.Attributes()["services"]
	if !servicesAttr.IsNull() && !servicesAttr.IsUnknown() {
		servicesList, ok := servicesAttr.(internaltypes.UnorderedListValue)
		if !ok {
			resp.Diagnostics.AddError(
				"Invalid Services Attribute",
				"Expected services to be a list but got different type",
			)
			return
		}
		// Traverse the list and check both the values is "DHCP" and "DNS" , if yes then return error
		hasDHCP := false
		hasDNS := false
		for _, v := range servicesList.Elements() {
			service, ok := v.(types.String)
			if !ok {
				resp.Diagnostics.AddError(
					"Invalid Service Value",
					"Expected service value to be a string but got different type",
				)
				return
			}
			if service.ValueString() == "DHCP" {
				hasDHCP = true
			}
			if service.ValueString() == "DNS" {
				hasDNS = true
			}
			if hasDHCP && hasDNS {
				resp.Diagnostics.AddAttributeError(
					path.Root("recurring_schedule").AtName("services"),
					"Invalid Services Configuration",
					"If both DHCP and DNS are selected in services, then the services must be set to ALL",
				)
				return
			}
		}
	}

	// Validate recurring_schedule schedule
	if !data.RecurringSchedule.IsNull() && !data.RecurringSchedule.IsUnknown() {
		utils.ValidateScheduleConfig(
			data.RecurringSchedule,
			"schedule",
			path.Root("recurring_schedule"),
			&resp.Diagnostics,
		)
	}
}

func (r *GridServicerestartGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("uuid"), req.ID)...)
	resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", []byte("true"))...)
}
