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
	"github.com/infobloxopen/infoblox-nios-go-client/grid"

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
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForGridServicerestartGroup).
		ReturnAsObject(1).
		Execute()

	// If the resource is not found, try searching using Extensible Attributes
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound && r.ReadByExtAttrs(ctx, &data, resp) {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read GridServicerestartGroup, got error: %s", err))
		return
	}

	res := apiRes.GetGridServicerestartGroupResponseObjectAsResult.GetResult()

	apiTerraformId, ok := (*res.ExtAttrs)[terraformInternalIDEA]
	if !ok {
		apiTerraformId.Value = ""
	}

	stateExtAttrs := ExpandExtAttrs(ctx, data.ExtAttrsAll, &diags)
	if stateExtAttrs == nil {
		resp.Diagnostics.AddError(
			"Missing Internal ID",
			"Unable to read GridServicerestartGroup because the internal ID (from extattrs_all) is missing or invalid.",
		)
		return
	}

	stateTerraformId := (*stateExtAttrs)[terraformInternalIDEA]
	if apiTerraformId.Value != stateTerraformId.Value {
		if r.ReadByExtAttrs(ctx, &data, resp) {
			return
		}
	}

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while reading GridServicerestartGroup due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GridServicerestartGroupResource) ReadByExtAttrs(ctx context.Context, data *GridServicerestartGroupModel, resp *resource.ReadResponse) bool {
	var diags diag.Diagnostics

	if data.ExtAttrsAll.IsNull() {
		return false
	}

	internalIdExtAttr := *ExpandExtAttrs(ctx, data.ExtAttrsAll, &diags)
	if diags.HasError() {
		return false
	}

	internalId := internalIdExtAttr[terraformInternalIDEA].Value
	if internalId == "" {
		return false
	}

	idMap := map[string]interface{}{
		terraformInternalIDEA: internalId,
	}

	apiRes, _, err := r.client.GridAPI.
		GridServicerestartGroupAPI.
		List(ctx).
		Extattrfilter(idMap).
		ReturnAsObject(1).
		ReturnFieldsPlus(readableAttributesForGridServicerestartGroup).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read GridServicerestartGroup by extattrs, got error: %s", err))
		return true
	}

	results := apiRes.ListGridServicerestartGroupResponseObject.GetResult()

	// If the list is empty, the resource no longer exists so remove it from state
	if len(results) == 0 {
		resp.State.RemoveResource(ctx)
		return true
	}

	res := results[0]

	// Remove inherited external attributes from extattrs
	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		return true
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)

	return true
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

	apiRes, _, err := r.client.GridAPI.
		GridServicerestartGroupAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
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
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
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

	scheduleAttr := data.RecurringSchedule.Attributes()["schedule"]
	if scheduleAttr.IsNull() || scheduleAttr.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("recurring_schedule"),
			"Invalid Configuration for Recurring Schedule",
			"Schedule must be set if recurring_schedule is set",
		)
		return
	}
	if !scheduleAttr.IsNull() && !scheduleAttr.IsUnknown() {
		scheduleObj, ok := scheduleAttr.(types.Object)
		if !ok {
			resp.Diagnostics.AddError(
				"Invalid Schedule Attribute",
				"Expected schedule to be an object but got different type",
			)
			return
		}
		schedule := scheduleObj.Attributes()
		recurringTime := schedule["recurring_time"]
		if !recurringTime.IsNull() && !recurringTime.IsUnknown() {
			if !schedule["hour_of_day"].IsNull() || !schedule["hour_of_day"].IsUnknown() || !schedule["year"].IsNull() || !schedule["year"].IsUnknown() || !schedule["month"].IsNull() || !schedule["month"].IsUnknown() || !schedule["day_of_month"].IsNull() || !schedule["day_of_month"].IsUnknown() {
				resp.Diagnostics.AddAttributeError(
					path.Root("recurring_schedule").AtName("schedule").AtName("recurring_time"),
					"Invalid Configuration for Schedule",
					"Cannot Set Recurring Time if any of hour_of_day, year, month, day_of_month is set",
				)
			}
		}

		repeat := schedule["repeat"]
		if !repeat.IsNull() && !repeat.IsUnknown() {
			repeatStr, ok := repeat.(types.String)
			if !ok {
				resp.Diagnostics.AddError(
					"Invalid Repeat Attribute",
					"Expected repeat to be a string but got different type",
				)
				return
			}
			if repeatStr.ValueString() == "ONCE" {
				if (!schedule["weekdays"].IsNull() && !schedule["weekdays"].IsUnknown()) || (!schedule["frequency"].IsNull() && !schedule["frequency"].IsUnknown()) || (!schedule["every"].IsNull() && !schedule["every"].IsUnknown()) {
					resp.Diagnostics.AddAttributeError(
						path.Root("recurring_schedule").AtName("schedule").AtName("repeat"),
						"Invalid Configuration for Repeat",
						"Cannot Set Frequency, Weekdays and Every if Repeat is set to ONCE",
					)
				}
				if schedule["month"].IsNull() || schedule["month"].IsUnknown() || schedule["day_of_month"].IsNull() || schedule["day_of_month"].IsUnknown() || schedule["hour_of_day"].IsNull() || schedule["hour_of_day"].IsUnknown() || schedule["minutes_past_hour"].IsNull() || schedule["minutes_past_hour"].IsUnknown() {
					resp.Diagnostics.AddAttributeError(
						path.Root("recurring_schedule").AtName("schedule").AtName("repeat"),
						"Invalid Configuration for Schedule",
						"If REPEAT is set to ONCE, then month, day_of_month, hour_of_day and minutes_past_hour must be set",
					)
				}
			} else {
				if (!schedule["month"].IsNull() && !schedule["month"].IsUnknown()) || (!schedule["day_of_month"].IsNull() && !schedule["day_of_month"].IsUnknown()) || (!schedule["year"].IsNull() && !schedule["year"].IsUnknown()) {
					resp.Diagnostics.AddAttributeError(
						path.Root("recurring_schedule").AtName("schedule").AtName("repeat"),
						"Invalid Configuration for Repeat",
						"Cannot Set Month, Day of Month and Year if Repeat is set to RECUR",
					)
				}

				if schedule["weekdays"].IsNull() || schedule["weekdays"].IsUnknown() || schedule["frequency"].IsNull() || schedule["frequency"].IsUnknown() || schedule["minutes_past_hour"].IsNull() || schedule["minutes_past_hour"].IsUnknown() || schedule["hour_of_day"].IsNull() || schedule["hour_of_day"].IsUnknown() {
					resp.Diagnostics.AddAttributeError(
						path.Root("recurring_schedule").AtName("schedule").AtName("repeat"),
						"Invalid Configuration for Schedule",
						"If REPEAT is set to RECUR, then weekdays, frequency, hour_of_day and minutes_past_hour must be set",
					)
				}
			}
		}
	}
}

func (r *GridServicerestartGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var diags diag.Diagnostics
	var data GridServicerestartGroupModel
	var goClientData grid.GridServicerestartGroup

	resourceRef := utils.ExtractResourceRef(req.ID)
	extattrs, diags := AddInternalIDToExtAttrs(ctx, data.ExtAttrs, diags)
	if diags.HasError() {
		return
	}
	goClientData.ExtAttrsPlus = ExpandExtAttrs(ctx, extattrs, &diags)
	data.ExtAttrsAll = extattrs

	updateRes, _, err := r.client.GridAPI.
		GridServicerestartGroupAPI.
		Update(ctx, resourceRef).
		GridServicerestartGroup(goClientData).
		ReturnFieldsPlus(readableAttributesForGridServicerestartGroup).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Import Failed", fmt.Sprintf("Unable to update GridServicerestartGroup for import, got error: %s", err))
		return
	}

	res := updateRes.UpdateGridServicerestartGroupResponseAsObject.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrsAll, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while update GridServicerestartGroup due inherited Extensible attributes for import, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
