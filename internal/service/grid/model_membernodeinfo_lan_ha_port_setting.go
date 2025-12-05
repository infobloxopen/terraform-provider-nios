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

type MembernodeinfoLanHaPortSettingModel struct {
	MgmtLan          types.String `tfsdk:"mgmt_lan"`
	MgmtIpv6addr     types.String `tfsdk:"mgmt_ipv6addr"`
	HaIpAddress      types.String `tfsdk:"ha_ip_address"`
	LanPortSetting   types.Object `tfsdk:"lan_port_setting"`
	HaPortSetting    types.Object `tfsdk:"ha_port_setting"`
	HaCloudAttribute types.String `tfsdk:"ha_cloud_attribute"`
}

var MembernodeinfoLanHaPortSettingAttrTypes = map[string]attr.Type{
	"mgmt_lan":           types.StringType,
	"mgmt_ipv6addr":      types.StringType,
	"ha_ip_address":      types.StringType,
	"lan_port_setting":   types.ObjectType{AttrTypes: MembernodeinfolanhaportsettingLanPortSettingAttrTypes},
	"ha_port_setting":    types.ObjectType{AttrTypes: MembernodeinfolanhaportsettingHaPortSettingAttrTypes},
	"ha_cloud_attribute": types.StringType,
}

var MembernodeinfoLanHaPortSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"mgmt_lan": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Public IPv4 address for the LAN1 interface.",
	},
	"mgmt_ipv6addr": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Public IPv6 address for the LAN1 interface.",
	},
	"ha_ip_address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "HA IP address.",
	},
	"lan_port_setting": schema.SingleNestedAttribute{
		Attributes: MembernodeinfolanhaportsettingLanPortSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"ha_port_setting": schema.SingleNestedAttribute{
		Attributes: MembernodeinfolanhaportsettingHaPortSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"ha_cloud_attribute": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "HA cloud interface from cloud platform side.",
	},
}

func ExpandMembernodeinfoLanHaPortSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MembernodeinfoLanHaPortSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MembernodeinfoLanHaPortSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MembernodeinfoLanHaPortSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MembernodeinfoLanHaPortSetting {
	if m == nil {
		return nil
	}
	to := &grid.MembernodeinfoLanHaPortSetting{
		MgmtLan:          flex.ExpandStringPointer(m.MgmtLan),
		MgmtIpv6addr:     flex.ExpandStringPointer(m.MgmtIpv6addr),
		HaIpAddress:      flex.ExpandStringPointer(m.HaIpAddress),
		LanPortSetting:   ExpandMembernodeinfolanhaportsettingLanPortSetting(ctx, m.LanPortSetting, diags),
		HaPortSetting:    ExpandMembernodeinfolanhaportsettingHaPortSetting(ctx, m.HaPortSetting, diags),
		HaCloudAttribute: flex.ExpandStringPointer(m.HaCloudAttribute),
	}
	return to
}

func FlattenMembernodeinfoLanHaPortSetting(ctx context.Context, from *grid.MembernodeinfoLanHaPortSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MembernodeinfoLanHaPortSettingAttrTypes)
	}
	m := MembernodeinfoLanHaPortSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MembernodeinfoLanHaPortSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MembernodeinfoLanHaPortSettingModel) Flatten(ctx context.Context, from *grid.MembernodeinfoLanHaPortSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MembernodeinfoLanHaPortSettingModel{}
	}
	m.MgmtLan = flex.FlattenStringPointer(from.MgmtLan)
	m.MgmtIpv6addr = flex.FlattenStringPointer(from.MgmtIpv6addr)
	m.HaIpAddress = flex.FlattenStringPointer(from.HaIpAddress)
	m.LanPortSetting = FlattenMembernodeinfolanhaportsettingLanPortSetting(ctx, from.LanPortSetting, diags)
	m.HaPortSetting = FlattenMembernodeinfolanhaportsettingHaPortSetting(ctx, from.HaPortSetting, diags)
	m.HaCloudAttribute = flex.FlattenStringPointer(from.HaCloudAttribute)
}
