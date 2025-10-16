package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type Ipv6rangetemplateMemberModel struct {
	Ipv4addr iptypes.IPv4Address `tfsdk:"ipv4addr"`
	Ipv6addr iptypes.IPv6Address `tfsdk:"ipv6addr"`
	Name     types.String        `tfsdk:"name"`
}

var Ipv6rangetemplateMemberAttrTypes = map[string]attr.Type{
	"ipv4addr": iptypes.IPv4AddressType{},
	"ipv6addr": iptypes.IPv6AddressType{},
	"name":     types.StringType,
}

var Ipv6rangetemplateMemberResourceSchemaAttributes = map[string]schema.Attribute{
	"ipv4addr": schema.StringAttribute{
		CustomType:          iptypes.IPv4AddressType{},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The IPv4 Address of the Grid Member.",
	},
	"ipv6addr": schema.StringAttribute{
		CustomType:          iptypes.IPv6AddressType{},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The IPv6 Address of the Grid Member.",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The Grid member name",
	},
}

func ExpandIpv6rangetemplateMember(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dhcp.Ipv6rangetemplateMember {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6rangetemplateMemberModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *Ipv6rangetemplateMemberModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.Ipv6rangetemplateMember {
	if m == nil {
		return nil
	}
	to := &dhcp.Ipv6rangetemplateMember{
		Ipv4addr: flex.ExpandIPv4Address(m.Ipv4addr),
		Ipv6addr: flex.ExpandIPv6Address(m.Ipv6addr),
		Name:     flex.ExpandStringPointer(m.Name),
	}
	return to
}

func FlattenIpv6rangetemplateMember(ctx context.Context, from *dhcp.Ipv6rangetemplateMember, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6rangetemplateMemberAttrTypes)
	}
	m := Ipv6rangetemplateMemberModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6rangetemplateMemberAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Ipv6rangetemplateMemberModel) Flatten(ctx context.Context, from *dhcp.Ipv6rangetemplateMember, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Ipv6rangetemplateMemberModel{}
	}
	m.Ipv4addr = flex.FlattenIPv4Address(from.Ipv4addr)
	m.Ipv6addr = flex.FlattenIPv6Address(from.Ipv6addr)
	m.Name = flex.FlattenStringPointer(from.Name)
}
