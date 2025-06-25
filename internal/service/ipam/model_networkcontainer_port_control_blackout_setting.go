package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/ipam"

	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/flex"
)

type NetworkcontainerPortControlBlackoutSettingModel struct {
	EnableBlackout   types.Bool   `tfsdk:"enable_blackout"`
	BlackoutDuration types.Int64  `tfsdk:"blackout_duration"`
	BlackoutSchedule types.Object `tfsdk:"blackout_schedule"`
}

var NetworkcontainerPortControlBlackoutSettingAttrTypes = map[string]attr.Type{
	"enable_blackout":   types.BoolType,
	"blackout_duration": types.Int64Type,
	"blackout_schedule": types.ObjectType{AttrTypes: NetworkcontainerportcontrolblackoutsettingBlackoutScheduleAttrTypes},
}

var NetworkcontainerPortControlBlackoutSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"enable_blackout": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether a blackout is enabled or not.",
	},
	"blackout_duration": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The blackout duration in seconds; minimum value is 1 minute.",
	},
	"blackout_schedule": schema.SingleNestedAttribute{
		Attributes: NetworkcontainerportcontrolblackoutsettingBlackoutScheduleResourceSchemaAttributes,
		Optional:   true,
	},
}

func ExpandNetworkcontainerPortControlBlackoutSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.NetworkcontainerPortControlBlackoutSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkcontainerPortControlBlackoutSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NetworkcontainerPortControlBlackoutSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.NetworkcontainerPortControlBlackoutSetting {
	if m == nil {
		return nil
	}
	to := &ipam.NetworkcontainerPortControlBlackoutSetting{
		EnableBlackout:   flex.ExpandBoolPointer(m.EnableBlackout),
		BlackoutDuration: flex.ExpandInt64Pointer(m.BlackoutDuration),
		BlackoutSchedule: ExpandNetworkcontainerportcontrolblackoutsettingBlackoutSchedule(ctx, m.BlackoutSchedule, diags),
	}
	return to
}

func FlattenNetworkcontainerPortControlBlackoutSetting(ctx context.Context, from *ipam.NetworkcontainerPortControlBlackoutSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkcontainerPortControlBlackoutSettingAttrTypes)
	}
	m := NetworkcontainerPortControlBlackoutSettingModel{}
	m.Flatten(ctx, from, diags)
	// m.ExtAttrs = m.ExtAttrsAll
	t, d := types.ObjectValueFrom(ctx, NetworkcontainerPortControlBlackoutSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NetworkcontainerPortControlBlackoutSettingModel) Flatten(ctx context.Context, from *ipam.NetworkcontainerPortControlBlackoutSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NetworkcontainerPortControlBlackoutSettingModel{}
	}
	m.EnableBlackout = types.BoolPointerValue(from.EnableBlackout)
	m.BlackoutDuration = flex.FlattenInt64Pointer(from.BlackoutDuration)
	m.BlackoutSchedule = FlattenNetworkcontainerportcontrolblackoutsettingBlackoutSchedule(ctx, from.BlackoutSchedule, diags)
}
