package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type MemberadditionaliplistIpv4NetworkSettingModel struct {
	Address       iptypes.IPv4Address `tfsdk:"address"`
	Gateway       types.String        `tfsdk:"gateway"`
	SubnetMask    types.String        `tfsdk:"subnet_mask"`
	VlanId        types.Int64         `tfsdk:"vlan_id"`
	Primary       types.Bool          `tfsdk:"primary"`
	Dscp          types.Int64         `tfsdk:"dscp"`
	LanSubnetMask types.String        `tfsdk:"lan_subnet_mask"`
	LanGateway    types.String        `tfsdk:"lan_gateway"`
	UseDscp       types.Bool          `tfsdk:"use_dscp"`
}

var MemberadditionaliplistIpv4NetworkSettingAttrTypes = map[string]attr.Type{
	"address":         iptypes.IPv4AddressType{},
	"gateway":         types.StringType,
	"subnet_mask":     types.StringType,
	"vlan_id":         types.Int64Type,
	"primary":         types.BoolType,
	"dscp":            types.Int64Type,
	"lan_subnet_mask": types.StringType,
	"lan_gateway":     types.StringType,
	"use_dscp":        types.BoolType,
}

var MemberadditionaliplistIpv4NetworkSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		CustomType:          iptypes.IPv4AddressType{},
		Computed:            true,
		Optional:            true,
		MarkdownDescription: "The IPv4 Address of the Grid Member.",
	},
	"gateway": schema.StringAttribute{
		Computed:            true,
		Optional:            true,
		MarkdownDescription: "The default gateway for the Grid Member.",
	},
	"subnet_mask": schema.StringAttribute{
		Computed:            true,
		Optional:            true,
		MarkdownDescription: "The subnet mask for the Grid Member.",
	},
	"vlan_id": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The identifier for the VLAN. Valid values are from 1 to 4096.",
	},
	"primary": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Determines if the current address is the primary VLAN address or not.",
	},
	"dscp": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(0),
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_dscp")),
		},
		MarkdownDescription: "The DSCP (Differentiated Services Code Point) value determines relative priorities for the type of services on your network. The appliance implements QoS (Quality of Service) rules based on this configuration. Valid values are from 0 to 63.",
	},
	"lan_subnet_mask": schema.StringAttribute{
		Computed:            true,
		Optional:            true,
		MarkdownDescription: "LAN netmask only for GCP HA.",
	},
	"lan_gateway": schema.StringAttribute{
		Computed:            true,
		Optional:            true,
		MarkdownDescription: "LAN gateway only for GCP HA.",
	},
	"use_dscp": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: dscp",
	},
}

func ExpandMemberadditionaliplistIpv4NetworkSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberadditionaliplistIpv4NetworkSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberadditionaliplistIpv4NetworkSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberadditionaliplistIpv4NetworkSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberadditionaliplistIpv4NetworkSetting {
	if m == nil {
		return nil
	}
	to := &grid.MemberadditionaliplistIpv4NetworkSetting{
		Address:       flex.ExpandIPv4Address(m.Address),
		Gateway:       flex.ExpandStringPointer(m.Gateway),
		SubnetMask:    flex.ExpandStringPointer(m.SubnetMask),
		VlanId:        flex.ExpandInt64Pointer(m.VlanId),
		Primary:       flex.ExpandBoolPointer(m.Primary),
		Dscp:          flex.ExpandInt64Pointer(m.Dscp),
		LanSubnetMask: flex.ExpandStringPointer(m.LanSubnetMask),
		LanGateway:    flex.ExpandStringPointer(m.LanGateway),
		UseDscp:       flex.ExpandBoolPointer(m.UseDscp),
	}
	return to
}

func FlattenMemberadditionaliplistIpv4NetworkSetting(ctx context.Context, from *grid.MemberadditionaliplistIpv4NetworkSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberadditionaliplistIpv4NetworkSettingAttrTypes)
	}
	m := MemberadditionaliplistIpv4NetworkSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberadditionaliplistIpv4NetworkSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberadditionaliplistIpv4NetworkSettingModel) Flatten(ctx context.Context, from *grid.MemberadditionaliplistIpv4NetworkSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberadditionaliplistIpv4NetworkSettingModel{}
	}
	m.Address = flex.FlattenIPv4Address(from.Address)
	m.Gateway = flex.FlattenStringPointer(from.Gateway)
	m.SubnetMask = flex.FlattenStringPointer(from.SubnetMask)
	m.VlanId = flex.FlattenInt64Pointer(from.VlanId)
	m.Primary = types.BoolPointerValue(from.Primary)
	m.Dscp = flex.FlattenInt64Pointer(from.Dscp)
	m.LanSubnetMask = flex.FlattenStringPointer(from.LanSubnetMask)
	m.LanGateway = flex.FlattenStringPointer(from.LanGateway)
	m.UseDscp = types.BoolPointerValue(from.UseDscp)
}
