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

	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForIpv6range = "address_type,cloud_info,comment,disable,discover_now_status,discovery_basic_poll_settings,discovery_blackout_setting,discovery_member,enable_discovery,end_addr,endpoint_sources,exclude,extattrs,ipv6_end_prefix,ipv6_prefix_bits,ipv6_start_prefix,logic_filter_rules,member,name,network,network_view,option_filter_rules,port_control_blackout_setting,recycle_leases,same_port_control_discovery_blackout,server_association_type,start_addr,subscribe_settings,use_blackout_setting,use_discovery_basic_polling_settings,use_enable_discovery,use_logic_filter_rules,use_recycle_leases,use_subscribe_settings"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &Ipv6rangeResource{}
var _ resource.ResourceWithImportState = &Ipv6rangeResource{}
var _ resource.ResourceWithValidateConfig = &Ipv6rangeResource{}

func NewIpv6rangeResource() resource.Resource {
	return &Ipv6rangeResource{}
}

// Ipv6rangeResource defines the resource implementation.
type Ipv6rangeResource struct {
	client *niosclient.APIClient
}

func (r *Ipv6rangeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "dhcp_ipv6range"
}

func (r *Ipv6rangeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages an IPv6 Range.",
		Attributes:          Ipv6rangeResourceSchemaAttributes,
	}
}

func (r *Ipv6rangeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *Ipv6rangeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var diags diag.Diagnostics
	var data Ipv6rangeModel

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

	apiRes, _, err := r.client.DHCPAPI.
		Ipv6rangeAPI.
		Create(ctx).
		Ipv6range(*data.Expand(ctx, &resp.Diagnostics, true)).
		ReturnFieldsPlus(readableAttributesForIpv6range).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Ipv6range, got error: %s", err))
		return
	}

	res := apiRes.CreateIpv6rangeResponseAsObject.GetResult()
	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while create Ipv6range due inherited Extensible attributes, got error: %s", err))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Ipv6rangeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var diags diag.Diagnostics
	var data Ipv6rangeModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	associateInternalId, diags := req.Private.GetKey(ctx, "associate_internal_id")
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.DHCPAPI.
		Ipv6rangeAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForIpv6range).
		ReturnAsObject(1).
		Execute()

	// If the resource is not found, try searching using Extensible Attributes
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound && r.ReadByExtAttrs(ctx, &data, resp) {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Ipv6range, got error: %s", err))
		return
	}

	res := apiRes.GetIpv6rangeResponseObjectAsResult.GetResult()

	apiTerraformId, ok := (*res.ExtAttrs)[terraformInternalIDEA]
	if !ok {
		apiTerraformId.Value = ""
	}

	if associateInternalId == nil {
		stateExtAttrs := ExpandExtAttrs(ctx, data.ExtAttrsAll, &diags)
		if stateExtAttrs == nil {
			resp.Diagnostics.AddError(
				"Missing Internal ID",
				"Unable to read Ipv6range because the internal ID (from extattrs_all) is missing or invalid.",
			)
			return
		}

		stateTerraformId := (*stateExtAttrs)[terraformInternalIDEA]
		if apiTerraformId.Value != stateTerraformId.Value {
			if r.ReadByExtAttrs(ctx, &data, resp) {
				return
			}
		}
	}

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while reading Ipv6range due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Ipv6rangeResource) ReadByExtAttrs(ctx context.Context, data *Ipv6rangeModel, resp *resource.ReadResponse) bool {
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

	apiRes, _, err := r.client.DHCPAPI.
		Ipv6rangeAPI.
		List(ctx).
		Extattrfilter(idMap).
		ReturnAsObject(1).
		ReturnFieldsPlus(readableAttributesForIpv6range).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Ipv6range by extattrs, got error: %s", err))
		return true
	}

	results := apiRes.ListIpv6rangeResponseObject.GetResult()

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

func (r *Ipv6rangeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data Ipv6rangeModel

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

	apiRes, _, err := r.client.DHCPAPI.
		Ipv6rangeAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Ipv6range(*data.Expand(ctx, &resp.Diagnostics, false)).
		ReturnFieldsPlus(readableAttributesForIpv6range).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Ipv6range, got error: %s", err))
		return
	}

	res := apiRes.UpdateIpv6rangeResponseAsObject.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, planExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while update Ipv6range due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if associateInternalId != nil {
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", nil)...)
	}
}

func (r *Ipv6rangeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Ipv6rangeModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.DHCPAPI.
		Ipv6rangeAPI.
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Ipv6range, got error: %s", err))
		return
	}
}

func (r *Ipv6rangeResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data Ipv6rangeModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	addressType := "ADDRESS"
	if !data.AddressType.IsNull() && !data.AddressType.IsUnknown() {
		addressType = data.AddressType.ValueString()
	}

	switch addressType {
	case "ADDRESS":
		if !data.StartAddr.IsUnknown() && !data.EndAddr.IsUnknown() {
			if data.StartAddr.IsNull() || data.EndAddr.IsNull() {
				resp.Diagnostics.AddError(
					"Configuration Error",
					"When address_type is set to 'ADDRESS' (default), both start_addr and end_addr must be specified.",
				)
			}
		}
		if !data.Ipv6StartPrefix.IsUnknown() && !data.Ipv6EndPrefix.IsUnknown() && !data.Ipv6PrefixBits.IsUnknown() {
			if !data.Ipv6StartPrefix.IsNull() || !data.Ipv6EndPrefix.IsNull() || !data.Ipv6PrefixBits.IsNull() {
				resp.Diagnostics.AddError(
					"Configuration Error",
					"When address_type is 'ADDRESS' (default), ipv6_start_prefix, ipv6_end_prefix, and ipv6_prefix_bits cannot be specified.",
				)
			}
		}
	case "PREFIX":
		if !data.Ipv6StartPrefix.IsUnknown() && !data.Ipv6EndPrefix.IsUnknown() && !data.Ipv6PrefixBits.IsUnknown() {
			if data.Ipv6StartPrefix.IsNull() || data.Ipv6EndPrefix.IsNull() || data.Ipv6PrefixBits.IsNull() {
				resp.Diagnostics.AddError(
					"Configuration Error",
					"When address_type is set to 'PREFIX', ipv6_start_prefix, ipv6_end_prefix, and ipv6_prefix_bits must be specified.",
				)
			}
		}
		if !data.StartAddr.IsUnknown() && !data.EndAddr.IsUnknown() {
			if !data.StartAddr.IsNull() || !data.EndAddr.IsNull() {
				resp.Diagnostics.AddError(
					"Configuration Error",
					"When address_type is 'PREFIX', start_addr and end_addr cannot be specified.",
				)
			}
		}
	case "BOTH":
		if !data.StartAddr.IsUnknown() && !data.EndAddr.IsUnknown() && !data.Ipv6StartPrefix.IsUnknown() && !data.Ipv6EndPrefix.IsUnknown() && !data.Ipv6PrefixBits.IsUnknown() {
			if data.StartAddr.IsNull() || data.EndAddr.IsNull() || data.Ipv6StartPrefix.IsNull() || data.Ipv6EndPrefix.IsNull() || data.Ipv6PrefixBits.IsNull() {
				resp.Diagnostics.AddError(
					"Configuration Error",
					"When address_type is set to 'BOTH', start_addr, end_addr, ipv6_start_prefix, ipv6_end_prefix, and ipv6_prefix_bits must be specified.",
				)
			}
		}
	}

	// Validate discovery_blackout_setting blackout_schedule
	if !data.DiscoveryBlackoutSetting.IsNull() && !data.DiscoveryBlackoutSetting.IsUnknown() {
		validateBlackoutSchedule(
			data.DiscoveryBlackoutSetting,
			path.Root("discovery_blackout_setting"),
			&resp.Diagnostics,
		)
	}

	// Validate port_control_blackout_setting blackout_schedule
	if !data.PortControlBlackoutSetting.IsNull() && !data.PortControlBlackoutSetting.IsUnknown() {
		validateBlackoutSchedule(
			data.PortControlBlackoutSetting,
			path.Root("port_control_blackout_setting"),
			&resp.Diagnostics,
		)
	}
}

// validateBlackoutSchedule validates the blackout_schedule configuration
func validateBlackoutSchedule(settingObj types.Object, basePath path.Path, diagnostics *diag.Diagnostics) {
	scheduleAttr := settingObj.Attributes()["blackout_schedule"]
	if scheduleAttr.IsNull() || scheduleAttr.IsUnknown() {
		return
	}

	scheduleObj, ok := scheduleAttr.(types.Object)
	if !ok {
		diagnostics.AddAttributeError(
			basePath.AtName("blackout_schedule"),
			"Invalid Blackout Schedule Attribute",
			"Expected blackout_schedule to be an object but got different type",
		)
		return
	}

	schedule := scheduleObj.Attributes()
	recurringTime := schedule["recurring_time"]
	repeat := schedule["repeat"]
	weekdays := schedule["weekdays"]
	frequency := schedule["frequency"]
	every := schedule["every"]
	minutesPastHour := schedule["minutes_past_hour"]
	month := schedule["month"]
	dayOfMonth := schedule["day_of_month"]
	hourOfDay := schedule["hour_of_day"]
	year := schedule["year"]

	if !recurringTime.IsNull() && !recurringTime.IsUnknown() {
		if !schedule["hour_of_day"].IsNull() || !schedule["hour_of_day"].IsUnknown() || !schedule["year"].IsNull() || !schedule["year"].IsUnknown() || !schedule["month"].IsNull() || !schedule["month"].IsUnknown() || !schedule["day_of_month"].IsNull() || !schedule["day_of_month"].IsUnknown() {
			diagnostics.AddAttributeError(
				basePath.AtName("blackout_schedule").AtName("schedule").AtName("recurring_time"),
				"Invalid Configuration for Schedule",
				"Cannot Set Recurring Time if any of hour_of_day, year, month, day_of_month is set",
			)
		}
	}

	if !repeat.IsNull() && !repeat.IsUnknown() {
		repeatStr, ok := repeat.(types.String)
		if !ok {
			diagnostics.AddAttributeError(
				basePath.AtName("blackout_schedule").AtName("schedule").AtName("repeat"),
				"Invalid Repeat Attribute",
				"Expected repeat to be a string but got different type",
			)
			return
		}

		switch repeatStr.ValueString() {
		case "ONCE":
			// For ONCE: cannot set weekdays, frequency, every
			if (!weekdays.IsNull() && !weekdays.IsUnknown()) ||
				(!frequency.IsNull() && !frequency.IsUnknown()) ||
				(!every.IsNull() && !every.IsUnknown()) {
				diagnostics.AddAttributeError(
					basePath.AtName("blackout_schedule").AtName("schedule").AtName("repeat"),
					"Invalid Configuration for Repeat",
					"Cannot set frequency, weekdays and every if repeat is set to ONCE",
				)
			}
			// For ONCE: must set month, day_of_month, hour_of_day, minutes_past_hour
			if month.IsNull() || month.IsUnknown() ||
				dayOfMonth.IsNull() || dayOfMonth.IsUnknown() ||
				hourOfDay.IsNull() || hourOfDay.IsUnknown() ||
				minutesPastHour.IsNull() || minutesPastHour.IsUnknown() {
				diagnostics.AddAttributeError(
					basePath.AtName("blackout_schedule").AtName("schedule").AtName("repeat"),
					"Invalid Configuration for Schedule",
					"If repeat is set to ONCE, then month, day_of_month, hour_of_day and minutes_past_hour must be set",
				)
			}
		case "RECUR":
			// For RECUR: cannot set month, day_of_month, year
			if (!month.IsNull() && !month.IsUnknown()) ||
				(!dayOfMonth.IsNull() && !dayOfMonth.IsUnknown()) ||
				(!year.IsNull() && !year.IsUnknown()) {
				diagnostics.AddAttributeError(
					basePath.AtName("blackout_schedule").AtName("schedule").AtName("repeat"),
					"Invalid Configuration for Repeat",
					"Cannot set month, day_of_month and year if repeat is set to RECUR",
				)
			}
			// For RECUR: must set frequency, hour_of_day, minutes_past_hour
			if frequency.IsNull() || frequency.IsUnknown() ||
				hourOfDay.IsNull() || hourOfDay.IsUnknown() ||
				minutesPastHour.IsNull() || minutesPastHour.IsUnknown() {
				diagnostics.AddAttributeError(
					basePath.AtName("blackout_schedule").AtName("schedule").AtName("repeat"),
					"Invalid Configuration for Schedule",
					"If repeat is set to RECUR, then frequency, hour_of_day and minutes_past_hour must be set",
				)
			}
			// Handle weekdays validation based on frequency for RECUR only
			if !frequency.IsNull() && !frequency.IsUnknown() {
				freqStr, ok := frequency.(types.String)
				if ok && freqStr.ValueString() == "WEEKLY" {
					// WEEKLY requires weekdays
					if weekdays.IsNull() || weekdays.IsUnknown() {
						diagnostics.AddAttributeError(
							basePath.AtName("blackout_schedule").AtName("schedule").AtName("weekdays"),
							"Invalid Configuration for Weekdays",
							"Weekdays must be set if frequency is set to WEEKLY",
						)
					}
				} else {
					// Non-WEEKLY cannot have weekdays
					if !weekdays.IsNull() && !weekdays.IsUnknown() {
						diagnostics.AddAttributeError(
							basePath.AtName("blackout_schedule").AtName("schedule").AtName("weekdays"),
							"Invalid Configuration for Weekdays",
							"Weekdays can only be set if frequency is set to WEEKLY",
						)
					}
				}
			}
		}
	}
}

func (r *Ipv6rangeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("ref"), req.ID)...)
	resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", []byte("true"))...)
}
