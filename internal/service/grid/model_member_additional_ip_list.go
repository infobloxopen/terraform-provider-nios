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

type MemberAdditionalIpListModel struct {
	Anycast            types.Bool   `tfsdk:"anycast"`
	Ipv4NetworkSetting types.Object `tfsdk:"ipv4_network_setting"`
	Ipv6NetworkSetting types.Object `tfsdk:"ipv6_network_setting"`
	Comment            types.String `tfsdk:"comment"`
	EnableBgp          types.Bool   `tfsdk:"enable_bgp"`
	EnableOspf         types.Bool   `tfsdk:"enable_ospf"`
	Interface          types.String `tfsdk:"interface"`
}

var MemberAdditionalIpListAttrTypes = map[string]attr.Type{
	"anycast":              types.BoolType,
	"ipv4_network_setting": types.ObjectType{AttrTypes: MemberadditionaliplistIpv4NetworkSettingAttrTypes},
	"ipv6_network_setting": types.ObjectType{AttrTypes: MemberadditionaliplistIpv6NetworkSettingAttrTypes},
	"comment":              types.StringType,
	"enable_bgp":           types.BoolType,
	"enable_ospf":          types.BoolType,
	"interface":            types.StringType,
}

var MemberAdditionalIpListResourceSchemaAttributes = map[string]schema.Attribute{
	"anycast": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if anycast for the Interface object is enabled or not.",
	},
	"ipv4_network_setting": schema.SingleNestedAttribute{
		Attributes: MemberadditionaliplistIpv4NetworkSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"ipv6_network_setting": schema.SingleNestedAttribute{
		Attributes: MemberadditionaliplistIpv6NetworkSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "A descriptive comment of this structure.",
	},
	"enable_bgp": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the BGP advertisement setting is enabled for this interface or not.",
	},
	"enable_ospf": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the OSPF advertisement setting is enabled for this interface or not.",
	},
	"interface": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The interface type for the Interface object.",
	},
}

func ExpandMemberAdditionalIpList(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberAdditionalIpList {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberAdditionalIpListModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberAdditionalIpListModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberAdditionalIpList {
	if m == nil {
		return nil
	}
	to := &grid.MemberAdditionalIpList{
		Anycast:            flex.ExpandBoolPointer(m.Anycast),
		Ipv4NetworkSetting: ExpandMemberadditionaliplistIpv4NetworkSetting(ctx, m.Ipv4NetworkSetting, diags),
		Ipv6NetworkSetting: ExpandMemberadditionaliplistIpv6NetworkSetting(ctx, m.Ipv6NetworkSetting, diags),
		Comment:            flex.ExpandStringPointer(m.Comment),
		EnableBgp:          flex.ExpandBoolPointer(m.EnableBgp),
		EnableOspf:         flex.ExpandBoolPointer(m.EnableOspf),
		Interface:          flex.ExpandStringPointer(m.Interface),
	}
	return to
}

func FlattenMemberAdditionalIpList(ctx context.Context, from *grid.MemberAdditionalIpList, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberAdditionalIpListAttrTypes)
	}
	m := MemberAdditionalIpListModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberAdditionalIpListAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberAdditionalIpListModel) Flatten(ctx context.Context, from *grid.MemberAdditionalIpList, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberAdditionalIpListModel{}
	}
	m.Anycast = types.BoolPointerValue(from.Anycast)
	m.Ipv4NetworkSetting = FlattenMemberadditionaliplistIpv4NetworkSetting(ctx, from.Ipv4NetworkSetting, diags)
	m.Ipv6NetworkSetting = FlattenMemberadditionaliplistIpv6NetworkSetting(ctx, from.Ipv6NetworkSetting, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.EnableBgp = types.BoolPointerValue(from.EnableBgp)
	m.EnableOspf = types.BoolPointerValue(from.EnableOspf)
	m.Interface = flex.FlattenStringPointer(from.Interface)
}
