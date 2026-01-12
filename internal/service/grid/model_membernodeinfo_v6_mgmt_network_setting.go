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

type MembernodeinfoV6MgmtNetworkSettingModel struct {
	Enabled                 types.Bool   `tfsdk:"enabled"`
	VirtualIp               types.String `tfsdk:"virtual_ip"`
	CidrPrefix              types.Int64  `tfsdk:"cidr_prefix"`
	Gateway                 types.String `tfsdk:"gateway"`
	AutoRouterConfigEnabled types.Bool   `tfsdk:"auto_router_config_enabled"`
	VlanId                  types.Int64  `tfsdk:"vlan_id"`
	Primary                 types.Bool   `tfsdk:"primary"`
	Dscp                    types.Int64  `tfsdk:"dscp"`
	UseDscp                 types.Bool   `tfsdk:"use_dscp"`
}

var MembernodeinfoV6MgmtNetworkSettingAttrTypes = map[string]attr.Type{
	"enabled":                    types.BoolType,
	"virtual_ip":                 types.StringType,
	"cidr_prefix":                types.Int64Type,
	"gateway":                    types.StringType,
	"auto_router_config_enabled": types.BoolType,
	"vlan_id":                    types.Int64Type,
	"primary":                    types.BoolType,
	"dscp":                       types.Int64Type,
	"use_dscp":                   types.BoolType,
}

var MembernodeinfoV6MgmtNetworkSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if IPv6 networking should be enabled.",
	},
	"virtual_ip": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "IPv6 address.",
	},
	"cidr_prefix": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "IPv6 cidr prefix",
	},
	"gateway": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Gateway address.",
	},
	"auto_router_config_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if automatic router configuration should be enabled.",
	},
	"vlan_id": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The identifier for the VLAN. Valid values are from 1 to 4096.",
	},
	"primary": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the current address is the primary VLAN address or not.",
	},
	"dscp": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The DSCP (Differentiated Services Code Point) value determines relative priorities for the type of services on your network. The appliance implements QoS (Quality of Service) rules based on this configuration. Valid values are from 0 to 63.",
	},
	"use_dscp": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: dscp",
	},
}

func ExpandMembernodeinfoV6MgmtNetworkSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MembernodeinfoV6MgmtNetworkSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MembernodeinfoV6MgmtNetworkSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MembernodeinfoV6MgmtNetworkSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MembernodeinfoV6MgmtNetworkSetting {
	if m == nil {
		return nil
	}
	to := &grid.MembernodeinfoV6MgmtNetworkSetting{
		Enabled:                 flex.ExpandBoolPointer(m.Enabled),
		VirtualIp:               flex.ExpandStringPointer(m.VirtualIp),
		CidrPrefix:              flex.ExpandInt64Pointer(m.CidrPrefix),
		Gateway:                 flex.ExpandStringPointer(m.Gateway),
		AutoRouterConfigEnabled: flex.ExpandBoolPointer(m.AutoRouterConfigEnabled),
		VlanId:                  flex.ExpandInt64Pointer(m.VlanId),
		Primary:                 flex.ExpandBoolPointer(m.Primary),
		Dscp:                    flex.ExpandInt64Pointer(m.Dscp),
		UseDscp:                 flex.ExpandBoolPointer(m.UseDscp),
	}
	return to
}

func FlattenMembernodeinfoV6MgmtNetworkSetting(ctx context.Context, from *grid.MembernodeinfoV6MgmtNetworkSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MembernodeinfoV6MgmtNetworkSettingAttrTypes)
	}
	m := MembernodeinfoV6MgmtNetworkSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MembernodeinfoV6MgmtNetworkSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MembernodeinfoV6MgmtNetworkSettingModel) Flatten(ctx context.Context, from *grid.MembernodeinfoV6MgmtNetworkSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MembernodeinfoV6MgmtNetworkSettingModel{}
	}
	m.Enabled = types.BoolPointerValue(from.Enabled)
	m.VirtualIp = flex.FlattenStringPointer(from.VirtualIp)
	m.CidrPrefix = flex.FlattenInt64Pointer(from.CidrPrefix)
	m.Gateway = flex.FlattenStringPointer(from.Gateway)
	m.AutoRouterConfigEnabled = types.BoolPointerValue(from.AutoRouterConfigEnabled)
	m.VlanId = flex.FlattenInt64Pointer(from.VlanId)
	m.Primary = types.BoolPointerValue(from.Primary)
	m.Dscp = flex.FlattenInt64Pointer(from.Dscp)
	m.UseDscp = types.BoolPointerValue(from.UseDscp)
}
