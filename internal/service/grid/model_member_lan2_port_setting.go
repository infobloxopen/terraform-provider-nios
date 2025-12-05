package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type MemberLan2PortSettingModel struct {
	VirtualRouterId             types.Int64  `tfsdk:"virtual_router_id"`
	Enabled                     types.Bool   `tfsdk:"enabled"`
	NetworkSetting              types.Object `tfsdk:"network_setting"`
	V6NetworkSetting            types.Object `tfsdk:"v6_network_setting"`
	NicFailoverEnabled          types.Bool   `tfsdk:"nic_failover_enabled"`
	NicFailoverEnablePrimary    types.Bool   `tfsdk:"nic_failover_enable_primary"`
	DefaultRouteFailoverEnabled types.Bool   `tfsdk:"default_route_failover_enabled"`
}

var MemberLan2PortSettingAttrTypes = map[string]attr.Type{
	"virtual_router_id":              types.Int64Type,
	"enabled":                        types.BoolType,
	"network_setting":                types.ObjectType{AttrTypes: Memberlan2portsettingNetworkSettingAttrTypes},
	"v6_network_setting":             types.ObjectType{AttrTypes: Memberlan2portsettingV6NetworkSettingAttrTypes},
	"nic_failover_enabled":           types.BoolType,
	"nic_failover_enable_primary":    types.BoolType,
	"default_route_failover_enabled": types.BoolType,
}

var MemberLan2PortSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"virtual_router_id": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "If the 'enabled' field is set to True, this defines the virtual router ID for the LAN2 port.",
	},
	"enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If this field is set to True, then it has its own IP settings. Otherwise, port redundancy mechanism is used, in which the LAN1 and LAN2 ports share the same IP settings for failover purposes.",
	},
	"network_setting": schema.SingleNestedAttribute{
		Attributes: Memberlan2portsettingNetworkSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"v6_network_setting": schema.SingleNestedAttribute{
		Attributes: Memberlan2portsettingV6NetworkSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"nic_failover_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if NIC failover is enabled or not.",
	},
	"nic_failover_enable_primary": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Prefer LAN1 when available.",
	},
	"default_route_failover_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Default route failover for LAN1 and LAN2.",
	},
}

func ExpandMemberLan2PortSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberLan2PortSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberLan2PortSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberLan2PortSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberLan2PortSetting {
	if m == nil {
		return nil
	}
	to := &grid.MemberLan2PortSetting{
		VirtualRouterId:             flex.ExpandInt64Pointer(m.VirtualRouterId),
		Enabled:                     flex.ExpandBoolPointer(m.Enabled),
		NetworkSetting:              ExpandMemberlan2portsettingNetworkSetting(ctx, m.NetworkSetting, diags),
		V6NetworkSetting:            ExpandMemberlan2portsettingV6NetworkSetting(ctx, m.V6NetworkSetting, diags),
		NicFailoverEnabled:          flex.ExpandBoolPointer(m.NicFailoverEnabled),
		NicFailoverEnablePrimary:    flex.ExpandBoolPointer(m.NicFailoverEnablePrimary),
		DefaultRouteFailoverEnabled: flex.ExpandBoolPointer(m.DefaultRouteFailoverEnabled),
	}
	return to
}

func FlattenMemberLan2PortSetting(ctx context.Context, from *grid.MemberLan2PortSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberLan2PortSettingAttrTypes)
	}
	m := MemberLan2PortSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberLan2PortSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberLan2PortSettingModel) Flatten(ctx context.Context, from *grid.MemberLan2PortSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberLan2PortSettingModel{}
	}
	m.VirtualRouterId = flex.FlattenInt64Pointer(from.VirtualRouterId)
	m.Enabled = types.BoolPointerValue(from.Enabled)
	m.NetworkSetting = FlattenMemberlan2portsettingNetworkSetting(ctx, from.NetworkSetting, diags)
	m.V6NetworkSetting = FlattenMemberlan2portsettingV6NetworkSetting(ctx, from.V6NetworkSetting, diags)
	m.NicFailoverEnabled = types.BoolPointerValue(from.NicFailoverEnabled)
	m.NicFailoverEnablePrimary = types.BoolPointerValue(from.NicFailoverEnablePrimary)
	m.DefaultRouteFailoverEnabled = types.BoolPointerValue(from.DefaultRouteFailoverEnabled)
}
