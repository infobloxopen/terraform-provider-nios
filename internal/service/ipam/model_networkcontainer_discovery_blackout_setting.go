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

type NetworkcontainerDiscoveryBlackoutSettingModel struct {
	EnableBlackout   types.Bool   `tfsdk:"enable_blackout"`
	BlackoutDuration types.Int64  `tfsdk:"blackout_duration"`
	BlackoutSchedule types.Object `tfsdk:"blackout_schedule"`
}

var NetworkcontainerDiscoveryBlackoutSettingAttrTypes = map[string]attr.Type{
	"enable_blackout":   types.BoolType,
	"blackout_duration": types.Int64Type,
	"blackout_schedule": types.ObjectType{AttrTypes: NetworkcontainerdiscoveryblackoutsettingBlackoutScheduleAttrTypes},
}

var NetworkcontainerDiscoveryBlackoutSettingResourceSchemaAttributes = map[string]schema.Attribute{
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
		Attributes: NetworkcontainerdiscoveryblackoutsettingBlackoutScheduleResourceSchemaAttributes,
		Optional:   true,
	},
}

func ExpandNetworkcontainerDiscoveryBlackoutSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.NetworkcontainerDiscoveryBlackoutSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkcontainerDiscoveryBlackoutSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NetworkcontainerDiscoveryBlackoutSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.NetworkcontainerDiscoveryBlackoutSetting {
	if m == nil {
		return nil
	}
	to := &ipam.NetworkcontainerDiscoveryBlackoutSetting{
		EnableBlackout:   flex.ExpandBoolPointer(m.EnableBlackout),
		BlackoutDuration: flex.ExpandInt64Pointer(m.BlackoutDuration),
		BlackoutSchedule: ExpandNetworkcontainerdiscoveryblackoutsettingBlackoutSchedule(ctx, m.BlackoutSchedule, diags),
	}
	return to
}

func FlattenNetworkcontainerDiscoveryBlackoutSetting(ctx context.Context, from *ipam.NetworkcontainerDiscoveryBlackoutSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkcontainerDiscoveryBlackoutSettingAttrTypes)
	}
	m := NetworkcontainerDiscoveryBlackoutSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkcontainerDiscoveryBlackoutSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NetworkcontainerDiscoveryBlackoutSettingModel) Flatten(ctx context.Context, from *ipam.NetworkcontainerDiscoveryBlackoutSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NetworkcontainerDiscoveryBlackoutSettingModel{}
	}
	m.EnableBlackout = types.BoolPointerValue(from.EnableBlackout)
	m.BlackoutDuration = flex.FlattenInt64Pointer(from.BlackoutDuration)
	m.BlackoutSchedule = FlattenNetworkcontainerdiscoveryblackoutsettingBlackoutSchedule(ctx, from.BlackoutSchedule, diags)
}
