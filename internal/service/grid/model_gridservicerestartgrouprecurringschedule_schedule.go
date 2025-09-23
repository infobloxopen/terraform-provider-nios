package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	internaltypes "github.com/infobloxopen/terraform-provider-nios/internal/types"
)

type GridservicerestartgrouprecurringscheduleScheduleModel struct {
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

var GridservicerestartgrouprecurringscheduleScheduleAttrTypes = map[string]attr.Type{
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

var GridservicerestartgrouprecurringscheduleScheduleResourceSchemaAttributes = map[string]schema.Attribute{
	"weekdays": schema.ListAttribute{
		CustomType:  internaltypes.UnorderedListOfStringType,
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.ValueStringsAre(
				stringvalidator.OneOf("MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"),
			),
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "Days of the week when scheduling is triggered.",
	},
	"time_zone": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString("UTC"),
		MarkdownDescription: "The time zone for the schedule.",
	},
	"recurring_time": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.ConflictsWith(
				path.MatchRelative().AtParent().AtName("hour_of_day"),
				path.MatchRelative().AtParent().AtName("year"),
				path.MatchRelative().AtParent().AtName("minutes_past_hour"),
			),
		},
		MarkdownDescription: "The recurring time for the schedule in Epoch seconds format. This field is obsolete and is preserved only for backward compatibility purposes. Please use other applicable fields to define the recurring schedule. DO NOT use recurring_time together with these fields. If you use recurring_time with other fields to define the recurring schedule, recurring_time has priority over year, hour_of_day, and minutes_past_hour and will override the values of these fields, although it does not override month and day_of_month. In this case, the recurring time value might be different than the intended value that you define.",
	},
	"frequency": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("DAILY", "HOURLY", "MONTHLY", "WEEKLY"),
		},
		MarkdownDescription: "The frequency for the scheduled task.",
	},
	"every": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The number of frequency to wait before repeating the scheduled task.",
	},
	"minutes_past_hour": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.Between(0, 59),
		},
		MarkdownDescription: "The minutes past the hour for the scheduled task.",
	},
	"hour_of_day": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.Between(0, 23),
		},
		MarkdownDescription: "The hour of day for the scheduled task.",
	},
	"year": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The year for the scheduled task.",
	},
	"month": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.Between(1, 12),
		},
		MarkdownDescription: "The month for the scheduled task.",
	},
	"day_of_month": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.Between(1, 31),
		},
		MarkdownDescription: "The day of the month for the scheduled task.",
	},
	"repeat": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("ONCE"),
		Validators: []validator.String{
			stringvalidator.OneOf("ONCE", "RECUR"),
		},
		MarkdownDescription: "Indicates if the scheduled task will be repeated or run only once.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If set to True, the scheduled task is disabled.",
	},
}

func ExpandGridservicerestartgrouprecurringscheduleSchedule(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridservicerestartgrouprecurringscheduleSchedule {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridservicerestartgrouprecurringscheduleScheduleModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridservicerestartgrouprecurringscheduleScheduleModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridservicerestartgrouprecurringscheduleSchedule {
	if m == nil {
		return nil
	}
	to := &grid.GridservicerestartgrouprecurringscheduleSchedule{
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

func FlattenGridservicerestartgrouprecurringscheduleSchedule(ctx context.Context, from *grid.GridservicerestartgrouprecurringscheduleSchedule, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridservicerestartgrouprecurringscheduleScheduleAttrTypes)
	}
	m := GridservicerestartgrouprecurringscheduleScheduleModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridservicerestartgrouprecurringscheduleScheduleAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridservicerestartgrouprecurringscheduleScheduleModel) Flatten(ctx context.Context, from *grid.GridservicerestartgrouprecurringscheduleSchedule, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridservicerestartgrouprecurringscheduleScheduleModel{}
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
