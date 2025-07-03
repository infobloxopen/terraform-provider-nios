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

type NetworkIpamThresholdSettingsModel struct {
	TriggerValue types.Int64 `tfsdk:"trigger_value"`
	ResetValue   types.Int64 `tfsdk:"reset_value"`
}

var NetworkIpamThresholdSettingsAttrTypes = map[string]attr.Type{
	"trigger_value": types.Int64Type,
	"reset_value":   types.Int64Type,
}

var NetworkIpamThresholdSettingsResourceSchemaAttributes = map[string]schema.Attribute{
	"trigger_value": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Indicates the percentage point which triggers the email/SNMP trap sending.",
	},
	"reset_value": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Indicates the percentage point which resets the email/SNMP trap sending.",
	},
}

func ExpandNetworkIpamThresholdSettings(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.NetworkIpamThresholdSettings {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkIpamThresholdSettingsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NetworkIpamThresholdSettingsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.NetworkIpamThresholdSettings {
	if m == nil {
		return nil
	}
	to := &ipam.NetworkIpamThresholdSettings{
		TriggerValue: flex.ExpandInt64Pointer(m.TriggerValue),
		ResetValue:   flex.ExpandInt64Pointer(m.ResetValue),
	}
	return to
}

func FlattenNetworkIpamThresholdSettings(ctx context.Context, from *ipam.NetworkIpamThresholdSettings, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkIpamThresholdSettingsAttrTypes)
	}
	m := NetworkIpamThresholdSettingsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkIpamThresholdSettingsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NetworkIpamThresholdSettingsModel) Flatten(ctx context.Context, from *ipam.NetworkIpamThresholdSettings, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NetworkIpamThresholdSettingsModel{}
	}
	m.TriggerValue = flex.FlattenInt64Pointer(from.TriggerValue)
	m.ResetValue = flex.FlattenInt64Pointer(from.ResetValue)
}
