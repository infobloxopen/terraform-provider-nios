package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	internaltypes "github.com/infobloxopen/terraform-provider-nios/internal/types"
)

type NetworkdiscoveryblackoutsettingBlackoutScheduleModel struct {
	Weekdays        internaltypes.UnorderedListValue `tfsdk:"weekdays"`
	TimeZone        types.String                     `tfsdk:"time_zone"`
	RecurringTime   types.Int64                      `tfsdk:"recurring_time"`
	Frequency       types.String                     `tfsdk:"frequency"`
	Every           types.Int64                      `tfsdk:"every"`
	MinutesPastHour types.Int64                      `tfsdk:"minutes_past_hour"`
	HourOfDay       types.Int64                      `tfsdk:"hour_of_day"`
	Year            types.Int64                      `tfsdk:"year"`
	Month           types.Int64                      `tfsdk:"month"`
	DayOfMonth      types.Int64                      `tfsdk:"day_of_month"`
	Repeat          types.String                     `tfsdk:"repeat"`
	Disable         types.Bool                       `tfsdk:"disable"`
}

var NetworkdiscoveryblackoutsettingBlackoutScheduleAttrTypes = map[string]attr.Type{
	"weekdays":          types.ListType{ElemType: types.StringType},
	"time_zone":         types.StringType,
	"recurring_time":    types.Int64Type,
	"frequency":         types.StringType,
	"every":             types.Int64Type,
	"minutes_past_hour": types.Int64Type,
	"hour_of_day":       types.Int64Type,
	"year":              types.Int64Type,
	"month":             types.Int64Type,
	"day_of_month":      types.Int64Type,
	"repeat":            types.StringType,
	"disable":           types.BoolType,
}

var NetworkdiscoveryblackoutsettingBlackoutScheduleResourceSchemaAttributes = map[string]schema.Attribute{
	"weekdays": schema.ListAttribute{
		ElementType:         types.StringType,
		CustomType:          internaltypes.UnorderedListOfStringType,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Days of the week when scheduling is triggered.",
		Validators: []validator.List{
			listvalidator.ValueStringsAre(
				stringvalidator.OneOf(
					"SUNDAY",
					"MONDAY",
					"TUESDAY",
					"WEDNESDAY",
					"THURSDAY",
					"FRIDAY",
					"SATURDAY",
				),
			),
			listvalidator.SizeAtLeast(1),
		},
	},
	"time_zone": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The time zone for the schedule.",
		Computed:            true,
		Default:             stringdefault.StaticString("UTC"),
	},
	"recurring_time": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The recurring time for the schedule in Epoch seconds format. This field is obsolete and is preserved only for backward compatibility purposes. Please use other applicable fields to define the recurring schedule. DO NOT use recurring_time together with these fields. If you use recurring_time with other fields to define the recurring schedule, recurring_time has priority over year, hour_of_day, and minutes_past_hour and will override the values of these fields, although it does not override month and day_of_month. In this case, the recurring time value might be different than the intended value that you define.",
	},
	"frequency": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The frequency for the scheduled task.",
		Computed:            true,
		Validators: []validator.String{
			stringvalidator.OneOf(
				"DAILY",
				"HOURLY",
				"WEEKLY",
				"MONTHLY",
			),
		},
	},
	"every": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The number of frequency to wait before repeating the scheduled task.",
		Computed:            true,
	},
	"minutes_past_hour": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The minutes past the hour for the scheduled task.",
		Computed:            true,
	},
	"hour_of_day": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The hour of day for the scheduled task.",
		Computed:            true,
		Validators: []validator.Int64{
			int64validator.Between(0, 23),
		},
	},
	"year": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The year for the scheduled task.",
	},
	"month": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The month for the scheduled task.",
		Computed:            true,
		Validators: []validator.Int64{
			int64validator.Between(1, 12),
		},
	},
	"day_of_month": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The day of the month for the scheduled task.",
		Computed:            true,
		Validators: []validator.Int64{
			int64validator.Between(1, 31),
		},
	},
	"repeat": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Indicates if the scheduled task will be repeated or run only once.",
		Computed:            true,
		Default:             stringdefault.StaticString("ONCE"),
		Validators: []validator.String{
			stringvalidator.OneOf(
				"ONCE",
				"RECUR",
			),
		},
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If set to True, the scheduled task is disabled.",
		Default:             booldefault.StaticBool(false),
		Computed:            true,
	},
}

func ExpandNetworkdiscoveryblackoutsettingBlackoutSchedule(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.NetworkdiscoveryblackoutsettingBlackoutSchedule {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkdiscoveryblackoutsettingBlackoutScheduleModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NetworkdiscoveryblackoutsettingBlackoutScheduleModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.NetworkdiscoveryblackoutsettingBlackoutSchedule {
	if m == nil {
		return nil
	}
	to := &ipam.NetworkdiscoveryblackoutsettingBlackoutSchedule{
		Weekdays:        flex.ExpandFrameworkListString(ctx, m.Weekdays, diags),
		TimeZone:        flex.ExpandStringPointer(m.TimeZone),
		RecurringTime:   flex.ExpandInt64Pointer(m.RecurringTime),
		Frequency:       flex.ExpandStringPointer(m.Frequency),
		Every:           flex.ExpandInt64Pointer(m.Every),
		MinutesPastHour: flex.ExpandInt64Pointer(m.MinutesPastHour),
		HourOfDay:       flex.ExpandInt64Pointer(m.HourOfDay),
		Year:            flex.ExpandInt64Pointer(m.Year),
		Month:           flex.ExpandInt64Pointer(m.Month),
		DayOfMonth:      flex.ExpandInt64Pointer(m.DayOfMonth),
		Repeat:          flex.ExpandStringPointer(m.Repeat),
		Disable:         flex.ExpandBoolPointer(m.Disable),
	}
	return to
}

func FlattenNetworkdiscoveryblackoutsettingBlackoutSchedule(ctx context.Context, from *ipam.NetworkdiscoveryblackoutsettingBlackoutSchedule, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkdiscoveryblackoutsettingBlackoutScheduleAttrTypes)
	}
	m := NetworkdiscoveryblackoutsettingBlackoutScheduleModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkdiscoveryblackoutsettingBlackoutScheduleAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NetworkdiscoveryblackoutsettingBlackoutScheduleModel) Flatten(ctx context.Context, from *ipam.NetworkdiscoveryblackoutsettingBlackoutSchedule, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NetworkdiscoveryblackoutsettingBlackoutScheduleModel{}
	}
	m.Weekdays = flex.FlattenFrameworkUnorderedList(ctx, types.StringType, from.Weekdays, diags)
	m.TimeZone = flex.FlattenStringPointer(from.TimeZone)
	m.RecurringTime = flex.FlattenInt64Pointer(from.RecurringTime)
	m.Frequency = flex.FlattenStringPointer(from.Frequency)
	m.Every = flex.FlattenInt64Pointer(from.Every)
	m.MinutesPastHour = flex.FlattenInt64Pointer(from.MinutesPastHour)
	m.HourOfDay = flex.FlattenInt64Pointer(from.HourOfDay)
	m.Year = flex.FlattenInt64Pointer(from.Year)
	m.Month = flex.FlattenInt64Pointer(from.Month)
	m.DayOfMonth = flex.FlattenInt64Pointer(from.DayOfMonth)
	m.Repeat = flex.FlattenStringPointer(from.Repeat)
	m.Disable = types.BoolPointerValue(from.Disable)
}
