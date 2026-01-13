package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type Ipv6networktemplateMembersModel struct {
	Ipv4addr iptypes.IPv4Address `tfsdk:"ipv4addr"`
	Ipv6addr iptypes.IPv6Address `tfsdk:"ipv6addr"`
	Name     types.String        `tfsdk:"name"`
}

var Ipv6networktemplateMembersAttrTypes = map[string]attr.Type{
	"ipv4addr": iptypes.IPv4AddressType{},
	"ipv6addr": iptypes.IPv6AddressType{},
	"name":     types.StringType,
}

var Ipv6networktemplateMembersResourceSchemaAttributes = map[string]schema.Attribute{
	"ipv4addr": schema.StringAttribute{
		CustomType:          iptypes.IPv4AddressType{},
		Computed:            true,
		Optional:            true,
		MarkdownDescription: "The IPv4 Address of the Grid Member.",
	},
	"ipv6addr": schema.StringAttribute{
		CustomType:          iptypes.IPv6AddressType{},
		Computed:            true,
		Optional:            true,
		MarkdownDescription: "The IPv6 Address of the Grid Member.",
	},
	"name": schema.StringAttribute{
		Computed:            true,
		Optional:            true,
		MarkdownDescription: "The Grid member name",
	},
}

func ExpandIpv6networktemplateMembers(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.Ipv6networktemplateMembers {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6networktemplateMembersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *Ipv6networktemplateMembersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.Ipv6networktemplateMembers {
	if m == nil {
		return nil
	}
	to := &ipam.Ipv6networktemplateMembers{
		Ipv4addr: flex.ExpandIPv4Address(m.Ipv4addr),
		Ipv6addr: flex.ExpandIPv6Address(m.Ipv6addr),
		Name:     flex.ExpandStringPointer(m.Name),
	}
	return to
}

func FlattenIpv6networktemplateMembers(ctx context.Context, from *ipam.Ipv6networktemplateMembers, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6networktemplateMembersAttrTypes)
	}
	m := Ipv6networktemplateMembersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6networktemplateMembersAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Ipv6networktemplateMembersModel) Flatten(ctx context.Context, from *ipam.Ipv6networktemplateMembers, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Ipv6networktemplateMembersModel{}
	}
	m.Ipv4addr = flex.FlattenIPv4Address(from.Ipv4addr)
	m.Ipv6addr = flex.FlattenIPv6Address(from.Ipv6addr)
	m.Name = flex.FlattenStringPointer(from.Name)
}
