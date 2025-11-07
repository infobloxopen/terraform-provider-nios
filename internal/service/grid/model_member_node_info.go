package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type MemberNodeInfoModel struct {
	ServiceStatus        types.List   `tfsdk:"service_status"`
	PhysicalOid          types.String `tfsdk:"physical_oid"`
	HaStatus             types.String `tfsdk:"ha_status"`
	Hwid                 types.String `tfsdk:"hwid"`
	Hwmodel              types.String `tfsdk:"hwmodel"`
	Hwtype               types.String `tfsdk:"hwtype"`
	HostPlatform         types.String `tfsdk:"host_platform"`
	Hypervisor           types.String `tfsdk:"hypervisor"`
	PaidNios             types.Bool   `tfsdk:"paid_nios"`
	MgmtNetworkSetting   types.Object `tfsdk:"mgmt_network_setting"`
	LanHaPortSetting     types.Object `tfsdk:"lan_ha_port_setting"`
	MgmtPhysicalSetting  types.Object `tfsdk:"mgmt_physical_setting"`
	Lan2PhysicalSetting  types.Object `tfsdk:"lan2_physical_setting"`
	NatExternalIp        types.String `tfsdk:"nat_external_ip"`
	V6MgmtNetworkSetting types.Object `tfsdk:"v6_mgmt_network_setting"`
}

var MemberNodeInfoAttrTypes = map[string]attr.Type{
	"service_status":          types.ListType{ElemType: types.ObjectType{AttrTypes: MembernodeinfoServiceStatusAttrTypes}},
	"physical_oid":            types.StringType,
	"ha_status":               types.StringType,
	"hwid":                    types.StringType,
	"hwmodel":                 types.StringType,
	"hwtype":                  types.StringType,
	"host_platform":           types.StringType,
	"hypervisor":              types.StringType,
	"paid_nios":               types.BoolType,
	"mgmt_network_setting":    types.ObjectType{AttrTypes: MembernodeinfoMgmtNetworkSettingAttrTypes},
	"lan_ha_port_setting":     types.ObjectType{AttrTypes: MembernodeinfoLanHaPortSettingAttrTypes},
	"mgmt_physical_setting":   types.ObjectType{AttrTypes: MembernodeinfoMgmtPhysicalSettingAttrTypes},
	"lan2_physical_setting":   types.ObjectType{AttrTypes: MembernodeinfoLan2PhysicalSettingAttrTypes},
	"nat_external_ip":         types.StringType,
	"v6_mgmt_network_setting": types.ObjectType{AttrTypes: MembernodeinfoV6MgmtNetworkSettingAttrTypes},
}

var MemberNodeInfoResourceSchemaAttributes = map[string]schema.Attribute{
	"service_status": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MembernodeinfoServiceStatusResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The service status list of the Grid Member.",
	},
	"physical_oid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The OID of the physical node.",
	},
	"ha_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Status about the node of an HA pair.",
	},
	"hwid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Hardware ID.",
	},
	"hwmodel": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Hardware model.",
	},
	"hwtype": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Hardware type.",
	},
	"host_platform": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Host Platform",
	},
	"hypervisor": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Hypervisor",
	},
	"paid_nios": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "True if node is Paid NIOS.",
	},
	"mgmt_network_setting": schema.SingleNestedAttribute{
		Attributes: MembernodeinfoMgmtNetworkSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"lan_ha_port_setting": schema.SingleNestedAttribute{
		Attributes: MembernodeinfoLanHaPortSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"mgmt_physical_setting": schema.SingleNestedAttribute{
		Attributes: MembernodeinfoMgmtPhysicalSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"lan2_physical_setting": schema.SingleNestedAttribute{
		Attributes: MembernodeinfoLan2PhysicalSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"nat_external_ip": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The NAT external IP address for the node.",
	},
	"v6_mgmt_network_setting": schema.SingleNestedAttribute{
		Attributes: MembernodeinfoV6MgmtNetworkSettingResourceSchemaAttributes,
		Optional:   true,
	},
}

func ExpandMemberNodeInfo(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberNodeInfo {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberNodeInfoModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberNodeInfoModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberNodeInfo {
	if m == nil {
		return nil
	}
	to := &grid.MemberNodeInfo{
		ServiceStatus:        flex.ExpandFrameworkListNestedBlock(ctx, m.ServiceStatus, diags, ExpandMembernodeinfoServiceStatus),
		MgmtNetworkSetting:   ExpandMembernodeinfoMgmtNetworkSetting(ctx, m.MgmtNetworkSetting, diags),
		LanHaPortSetting:     ExpandMembernodeinfoLanHaPortSetting(ctx, m.LanHaPortSetting, diags),
		MgmtPhysicalSetting:  ExpandMembernodeinfoMgmtPhysicalSetting(ctx, m.MgmtPhysicalSetting, diags),
		Lan2PhysicalSetting:  ExpandMembernodeinfoLan2PhysicalSetting(ctx, m.Lan2PhysicalSetting, diags),
		NatExternalIp:        flex.ExpandStringPointer(m.NatExternalIp),
		V6MgmtNetworkSetting: ExpandMembernodeinfoV6MgmtNetworkSetting(ctx, m.V6MgmtNetworkSetting, diags),
	}
	return to
}

func FlattenMemberNodeInfo(ctx context.Context, from *grid.MemberNodeInfo, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberNodeInfoAttrTypes)
	}
	m := MemberNodeInfoModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberNodeInfoAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberNodeInfoModel) Flatten(ctx context.Context, from *grid.MemberNodeInfo, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberNodeInfoModel{}
	}
	m.ServiceStatus = flex.FlattenFrameworkListNestedBlock(ctx, from.ServiceStatus, MembernodeinfoServiceStatusAttrTypes, diags, FlattenMembernodeinfoServiceStatus)
	m.PhysicalOid = flex.FlattenStringPointer(from.PhysicalOid)
	m.HaStatus = flex.FlattenStringPointer(from.HaStatus)
	m.Hwid = flex.FlattenStringPointer(from.Hwid)
	m.Hwmodel = flex.FlattenStringPointer(from.Hwmodel)
	m.Hwtype = flex.FlattenStringPointer(from.Hwtype)
	m.HostPlatform = flex.FlattenStringPointer(from.HostPlatform)
	m.Hypervisor = flex.FlattenStringPointer(from.Hypervisor)
	m.PaidNios = types.BoolPointerValue(from.PaidNios)
	m.MgmtNetworkSetting = FlattenMembernodeinfoMgmtNetworkSetting(ctx, from.MgmtNetworkSetting, diags)
	m.LanHaPortSetting = FlattenMembernodeinfoLanHaPortSetting(ctx, from.LanHaPortSetting, diags)
	m.MgmtPhysicalSetting = FlattenMembernodeinfoMgmtPhysicalSetting(ctx, from.MgmtPhysicalSetting, diags)
	m.Lan2PhysicalSetting = FlattenMembernodeinfoLan2PhysicalSetting(ctx, from.Lan2PhysicalSetting, diags)
	m.NatExternalIp = flex.FlattenStringPointer(from.NatExternalIp)
	m.V6MgmtNetworkSetting = FlattenMembernodeinfoV6MgmtNetworkSetting(ctx, from.V6MgmtNetworkSetting, diags)
}
