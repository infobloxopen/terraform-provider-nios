package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type Ipv6rangePortControlBlackoutSettingModel struct {
	EnableBlackout   types.Bool   `tfsdk:"enable_blackout"`
	BlackoutDuration types.Int64  `tfsdk:"blackout_duration"`
	BlackoutSchedule types.Object `tfsdk:"blackout_schedule"`
}

var Ipv6rangePortControlBlackoutSettingAttrTypes = map[string]attr.Type{
	"enable_blackout":   types.BoolType,
	"blackout_duration": types.Int64Type,
	"blackout_schedule": types.ObjectType{AttrTypes: Ipv6rangeportcontrolblackoutsettingBlackoutScheduleAttrTypes},
}

var Ipv6rangePortControlBlackoutSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"enable_blackout": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Determines whether a blackout is enabled or not.",
	},
	"blackout_duration": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The blackout duration in seconds; minimum value is 1 minute.",
	},
	"blackout_schedule": schema.SingleNestedAttribute{
		Attributes:          Ipv6rangeportcontrolblackoutsettingBlackoutScheduleResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "A Schedule Setting struct that determines blackout schedule.",
	},
}

func ExpandIpv6rangePortControlBlackoutSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dhcp.Ipv6rangePortControlBlackoutSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6rangePortControlBlackoutSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *Ipv6rangePortControlBlackoutSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.Ipv6rangePortControlBlackoutSetting {
	if m == nil {
		return nil
	}
	to := &dhcp.Ipv6rangePortControlBlackoutSetting{
		EnableBlackout:   flex.ExpandBoolPointer(m.EnableBlackout),
		BlackoutDuration: flex.ExpandInt64Pointer(m.BlackoutDuration),
		BlackoutSchedule: ExpandIpv6rangeportcontrolblackoutsettingBlackoutSchedule(ctx, m.BlackoutSchedule, diags),
	}
	return to
}

func FlattenIpv6rangePortControlBlackoutSetting(ctx context.Context, from *dhcp.Ipv6rangePortControlBlackoutSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6rangePortControlBlackoutSettingAttrTypes)
	}
	m := Ipv6rangePortControlBlackoutSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6rangePortControlBlackoutSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Ipv6rangePortControlBlackoutSettingModel) Flatten(ctx context.Context, from *dhcp.Ipv6rangePortControlBlackoutSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Ipv6rangePortControlBlackoutSettingModel{}
	}
	m.EnableBlackout = types.BoolPointerValue(from.EnableBlackout)
	m.BlackoutDuration = flex.FlattenInt64Pointer(from.BlackoutDuration)
	m.BlackoutSchedule = FlattenIpv6rangeportcontrolblackoutsettingBlackoutSchedule(ctx, from.BlackoutSchedule, diags)
}
