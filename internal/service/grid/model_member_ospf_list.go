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

type MemberOspfListModel struct {
	AreaId                 types.String `tfsdk:"area_id"`
	AreaType               types.String `tfsdk:"area_type"`
	AuthenticationKey      types.String `tfsdk:"authentication_key"`
	AuthenticationType     types.String `tfsdk:"authentication_type"`
	AutoCalcCostEnabled    types.Bool   `tfsdk:"auto_calc_cost_enabled"`
	Comment                types.String `tfsdk:"comment"`
	Cost                   types.Int64  `tfsdk:"cost"`
	DeadInterval           types.Int64  `tfsdk:"dead_interval"`
	HelloInterval          types.Int64  `tfsdk:"hello_interval"`
	Interface              types.String `tfsdk:"interface"`
	IsIpv4                 types.Bool   `tfsdk:"is_ipv4"`
	KeyId                  types.Int64  `tfsdk:"key_id"`
	RetransmitInterval     types.Int64  `tfsdk:"retransmit_interval"`
	TransmitDelay          types.Int64  `tfsdk:"transmit_delay"`
	AdvertiseInterfaceVlan types.String `tfsdk:"advertise_interface_vlan"`
	BfdTemplate            types.String `tfsdk:"bfd_template"`
	EnableBfd              types.Bool   `tfsdk:"enable_bfd"`
	EnableBfdDnscheck      types.Bool   `tfsdk:"enable_bfd_dnscheck"`
}

var MemberOspfListAttrTypes = map[string]attr.Type{
	"area_id":                  types.StringType,
	"area_type":                types.StringType,
	"authentication_key":       types.StringType,
	"authentication_type":      types.StringType,
	"auto_calc_cost_enabled":   types.BoolType,
	"comment":                  types.StringType,
	"cost":                     types.Int64Type,
	"dead_interval":            types.Int64Type,
	"hello_interval":           types.Int64Type,
	"interface":                types.StringType,
	"is_ipv4":                  types.BoolType,
	"key_id":                   types.Int64Type,
	"retransmit_interval":      types.Int64Type,
	"transmit_delay":           types.Int64Type,
	"advertise_interface_vlan": types.StringType,
	"bfd_template":             types.StringType,
	"enable_bfd":               types.BoolType,
	"enable_bfd_dnscheck":      types.BoolType,
}

var MemberOspfListResourceSchemaAttributes = map[string]schema.Attribute{
	"area_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The area ID value of the OSPF settings.",
	},
	"area_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The OSPF area type.",
	},
	"authentication_key": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The authentication password to use for OSPF. The authentication key is valid only when authentication type is \"SIMPLE\" or \"MESSAGE_DIGEST\".",
	},
	"authentication_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The authentication type used for the OSPF advertisement.",
	},
	"auto_calc_cost_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if auto calculate cost is enabled or not.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "A descriptive comment of the OSPF configuration.",
	},
	"cost": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The cost metric associated with the OSPF advertisement.",
	},
	"dead_interval": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The dead interval value of OSPF (in seconds). The dead interval describes the time to wait before declaring the device is unavailable and down.",
	},
	"hello_interval": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The hello interval value of OSPF. The hello interval specifies how often to send OSPF hello advertisement, in seconds.",
	},
	"interface": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The interface that sends out OSPF advertisement information.",
	},
	"is_ipv4": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "The OSPF protocol version. Specify \"true\" if the IPv4 version of OSPF is used, or \"false\" if the IPv6 version of OSPF is used.",
	},
	"key_id": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The hash key identifier to use for \"MESSAGE_DIGEST\" authentication. The hash key identifier is valid only when authentication type is \"MESSAGE_DIGEST\".",
	},
	"retransmit_interval": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The retransmit interval time of OSPF (in seconds). The retransmit interval describes the time to wait before retransmitting OSPF advertisement.",
	},
	"transmit_delay": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The transmit delay value of OSPF (in seconds). The transmit delay describes the time to wait before sending an advertisement.",
	},
	"advertise_interface_vlan": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The VLAN used as the advertising interface for sending OSPF announcements.",
	},
	"bfd_template": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Determines BFD template name.",
	},
	"enable_bfd": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the BFD is enabled or not.",
	},
	"enable_bfd_dnscheck": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if BFD internal DNS monitor is enabled or not.",
	},
}

func ExpandMemberOspfList(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberOspfList {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberOspfListModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberOspfListModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberOspfList {
	if m == nil {
		return nil
	}
	to := &grid.MemberOspfList{
		AreaId:                 flex.ExpandStringPointer(m.AreaId),
		AreaType:               flex.ExpandStringPointer(m.AreaType),
		AuthenticationKey:      flex.ExpandStringPointer(m.AuthenticationKey),
		AuthenticationType:     flex.ExpandStringPointer(m.AuthenticationType),
		AutoCalcCostEnabled:    flex.ExpandBoolPointer(m.AutoCalcCostEnabled),
		Comment:                flex.ExpandStringPointer(m.Comment),
		Cost:                   flex.ExpandInt64Pointer(m.Cost),
		DeadInterval:           flex.ExpandInt64Pointer(m.DeadInterval),
		HelloInterval:          flex.ExpandInt64Pointer(m.HelloInterval),
		Interface:              flex.ExpandStringPointer(m.Interface),
		IsIpv4:                 flex.ExpandBoolPointer(m.IsIpv4),
		KeyId:                  flex.ExpandInt64Pointer(m.KeyId),
		RetransmitInterval:     flex.ExpandInt64Pointer(m.RetransmitInterval),
		TransmitDelay:          flex.ExpandInt64Pointer(m.TransmitDelay),
		AdvertiseInterfaceVlan: flex.ExpandStringPointer(m.AdvertiseInterfaceVlan),
		BfdTemplate:            flex.ExpandStringPointer(m.BfdTemplate),
		EnableBfd:              flex.ExpandBoolPointer(m.EnableBfd),
		EnableBfdDnscheck:      flex.ExpandBoolPointer(m.EnableBfdDnscheck),
	}
	return to
}

func FlattenMemberOspfList(ctx context.Context, from *grid.MemberOspfList, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberOspfListAttrTypes)
	}
	m := MemberOspfListModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberOspfListAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberOspfListModel) Flatten(ctx context.Context, from *grid.MemberOspfList, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberOspfListModel{}
	}
	m.AreaId = flex.FlattenStringPointer(from.AreaId)
	m.AreaType = flex.FlattenStringPointer(from.AreaType)
	m.AuthenticationKey = flex.FlattenStringPointer(from.AuthenticationKey)
	m.AuthenticationType = flex.FlattenStringPointer(from.AuthenticationType)
	m.AutoCalcCostEnabled = types.BoolPointerValue(from.AutoCalcCostEnabled)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Cost = flex.FlattenInt64Pointer(from.Cost)
	m.DeadInterval = flex.FlattenInt64Pointer(from.DeadInterval)
	m.HelloInterval = flex.FlattenInt64Pointer(from.HelloInterval)
	m.Interface = flex.FlattenStringPointer(from.Interface)
	m.IsIpv4 = types.BoolPointerValue(from.IsIpv4)
	m.KeyId = flex.FlattenInt64Pointer(from.KeyId)
	m.RetransmitInterval = flex.FlattenInt64Pointer(from.RetransmitInterval)
	m.TransmitDelay = flex.FlattenInt64Pointer(from.TransmitDelay)
	m.AdvertiseInterfaceVlan = flex.FlattenStringPointer(from.AdvertiseInterfaceVlan)
	m.BfdTemplate = flex.FlattenStringPointer(from.BfdTemplate)
	m.EnableBfd = types.BoolPointerValue(from.EnableBfd)
	m.EnableBfdDnscheck = types.BoolPointerValue(from.EnableBfdDnscheck)
}
