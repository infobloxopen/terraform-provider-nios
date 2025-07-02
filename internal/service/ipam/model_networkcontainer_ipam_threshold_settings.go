package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/ipam"

	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/flex"
)

type NetworkcontainerIpamThresholdSettingsModel struct {
	TriggerValue types.Int64 `tfsdk:"trigger_value"`
	ResetValue   types.Int64 `tfsdk:"reset_value"`
}

var NetworkcontainerIpamThresholdSettingsAttrTypes = map[string]attr.Type{
	"trigger_value": types.Int64Type,
	"reset_value":   types.Int64Type,
}

var NetworkcontainerIpamThresholdSettingsResourceSchemaAttributes = map[string]schema.Attribute{
	"trigger_value": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Indicates the percentage point which triggers the email/SNMP trap sending.",
		Default:             int64default.StaticInt64(95),
	},
	"reset_value": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(85),
		MarkdownDescription: "Indicates the percentage point which resets the email/SNMP trap sending.",
	},
}

func ExpandNetworkcontainerIpamThresholdSettings(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.NetworkcontainerIpamThresholdSettings {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkcontainerIpamThresholdSettingsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NetworkcontainerIpamThresholdSettingsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.NetworkcontainerIpamThresholdSettings {
	if m == nil {
		return nil
	}
	to := &ipam.NetworkcontainerIpamThresholdSettings{
		TriggerValue: flex.ExpandInt64Pointer(m.TriggerValue),
		ResetValue:   flex.ExpandInt64Pointer(m.ResetValue),
	}
	return to
}

func FlattenNetworkcontainerIpamThresholdSettings(ctx context.Context, from *ipam.NetworkcontainerIpamThresholdSettings, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkcontainerIpamThresholdSettingsAttrTypes)
	}
	m := NetworkcontainerIpamThresholdSettingsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkcontainerIpamThresholdSettingsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NetworkcontainerIpamThresholdSettingsModel) Flatten(ctx context.Context, from *ipam.NetworkcontainerIpamThresholdSettings, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NetworkcontainerIpamThresholdSettingsModel{}
	}
	m.TriggerValue = flex.FlattenInt64Pointer(from.TriggerValue)
	m.ResetValue = flex.FlattenInt64Pointer(from.ResetValue)
}
