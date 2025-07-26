package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type Ipv6networkPortControlBlackoutSettingModel struct {
	EnableBlackout   types.Bool   `tfsdk:"enable_blackout"`
	BlackoutDuration types.Int64  `tfsdk:"blackout_duration"`
	BlackoutSchedule types.Object `tfsdk:"blackout_schedule"`
}

var Ipv6networkPortControlBlackoutSettingAttrTypes = map[string]attr.Type{
	"enable_blackout":   types.BoolType,
	"blackout_duration": types.Int64Type,
	"blackout_schedule": types.ObjectType{AttrTypes: Ipv6networkportcontrolblackoutsettingBlackoutScheduleAttrTypes},
}

var Ipv6networkPortControlBlackoutSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"enable_blackout": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether a blackout is enabled or not.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"blackout_duration": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The blackout duration in seconds; minimum value is 1 minute.",
		Computed:            true,
	},
	"blackout_schedule": schema.SingleNestedAttribute{
		Attributes: Ipv6networkportcontrolblackoutsettingBlackoutScheduleResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
}

func ExpandIpv6networkPortControlBlackoutSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.Ipv6networkPortControlBlackoutSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6networkPortControlBlackoutSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *Ipv6networkPortControlBlackoutSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.Ipv6networkPortControlBlackoutSetting {
	if m == nil {
		return nil
	}
	to := &ipam.Ipv6networkPortControlBlackoutSetting{
		EnableBlackout:   flex.ExpandBoolPointer(m.EnableBlackout),
		BlackoutDuration: flex.ExpandInt64Pointer(m.BlackoutDuration),
		BlackoutSchedule: ExpandIpv6networkportcontrolblackoutsettingBlackoutSchedule(ctx, m.BlackoutSchedule, diags),
	}
	return to
}

func FlattenIpv6networkPortControlBlackoutSetting(ctx context.Context, from *ipam.Ipv6networkPortControlBlackoutSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6networkPortControlBlackoutSettingAttrTypes)
	}
	m := Ipv6networkPortControlBlackoutSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6networkPortControlBlackoutSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Ipv6networkPortControlBlackoutSettingModel) Flatten(ctx context.Context, from *ipam.Ipv6networkPortControlBlackoutSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Ipv6networkPortControlBlackoutSettingModel{}
	}
	m.EnableBlackout = types.BoolPointerValue(from.EnableBlackout)
	m.BlackoutDuration = flex.FlattenInt64Pointer(from.BlackoutDuration)
	m.BlackoutSchedule = FlattenIpv6networkportcontrolblackoutsettingBlackoutSchedule(ctx, from.BlackoutSchedule, diags)
}
