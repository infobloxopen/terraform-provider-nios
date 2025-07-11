package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dhcp"

	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/flex"
)

type RangePortControlBlackoutSettingModel struct {
	EnableBlackout   types.Bool   `tfsdk:"enable_blackout"`
	BlackoutDuration types.Int64  `tfsdk:"blackout_duration"`
	BlackoutSchedule types.Object `tfsdk:"blackout_schedule"`
}

var RangePortControlBlackoutSettingAttrTypes = map[string]attr.Type{
	"enable_blackout":   types.BoolType,
	"blackout_duration": types.Int64Type,
	"blackout_schedule": types.ObjectType{AttrTypes: RangeportcontrolblackoutsettingBlackoutScheduleAttrTypes},
}

var RangePortControlBlackoutSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"enable_blackout": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether a blackout is enabled or not.",
	},
	"blackout_duration": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The blackout duration in seconds; minimum value is 1 minute.",
	},
	"blackout_schedule": schema.SingleNestedAttribute{
		Attributes: RangeportcontrolblackoutsettingBlackoutScheduleResourceSchemaAttributes,
		Optional:   true,
	},
}

func ExpandRangePortControlBlackoutSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dhcp.RangePortControlBlackoutSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RangePortControlBlackoutSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *RangePortControlBlackoutSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.RangePortControlBlackoutSetting {
	if m == nil {
		return nil
	}
	to := &dhcp.RangePortControlBlackoutSetting{
		EnableBlackout:   flex.ExpandBoolPointer(m.EnableBlackout),
		BlackoutDuration: flex.ExpandInt64Pointer(m.BlackoutDuration),
		BlackoutSchedule: ExpandRangeportcontrolblackoutsettingBlackoutSchedule(ctx, m.BlackoutSchedule, diags),
	}
	return to
}

func FlattenRangePortControlBlackoutSetting(ctx context.Context, from *dhcp.RangePortControlBlackoutSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RangePortControlBlackoutSettingAttrTypes)
	}
	m := RangePortControlBlackoutSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RangePortControlBlackoutSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RangePortControlBlackoutSettingModel) Flatten(ctx context.Context, from *dhcp.RangePortControlBlackoutSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RangePortControlBlackoutSettingModel{}
	}
	m.EnableBlackout = types.BoolPointerValue(from.EnableBlackout)
	m.BlackoutDuration = flex.FlattenInt64Pointer(from.BlackoutDuration)
	m.BlackoutSchedule = FlattenRangeportcontrolblackoutsettingBlackoutSchedule(ctx, from.BlackoutSchedule, diags)
}
