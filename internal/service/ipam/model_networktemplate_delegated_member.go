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

type NetworktemplateDelegatedMemberModel struct {
	Ipv4addr iptypes.IPv4Address `tfsdk:"ipv4addr"`
	Ipv6addr iptypes.IPv6Address `tfsdk:"ipv6addr"`
	Name     types.String        `tfsdk:"name"`
}

var NetworktemplateDelegatedMemberAttrTypes = map[string]attr.Type{
	"ipv4addr": iptypes.IPv4AddressType{},
	"ipv6addr": iptypes.IPv6AddressType{},
	"name":     types.StringType,
}

var NetworktemplateDelegatedMemberResourceSchemaAttributes = map[string]schema.Attribute{
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

func ExpandNetworktemplateDelegatedMember(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.NetworktemplateDelegatedMember {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworktemplateDelegatedMemberModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NetworktemplateDelegatedMemberModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.NetworktemplateDelegatedMember {
	if m == nil {
		return nil
	}
	to := &ipam.NetworktemplateDelegatedMember{
		Ipv4addr: flex.ExpandIPv4Address(m.Ipv4addr),
		Ipv6addr: flex.ExpandIPv6Address(m.Ipv6addr),
		Name:     flex.ExpandStringPointer(m.Name),
	}
	return to
}

func FlattenNetworktemplateDelegatedMember(ctx context.Context, from *ipam.NetworktemplateDelegatedMember, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworktemplateDelegatedMemberAttrTypes)
	}
	m := NetworktemplateDelegatedMemberModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworktemplateDelegatedMemberAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NetworktemplateDelegatedMemberModel) Flatten(ctx context.Context, from *ipam.NetworktemplateDelegatedMember, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NetworktemplateDelegatedMemberModel{}
	}
	m.Ipv4addr = flex.FlattenIPv4Address(from.Ipv4addr)
	m.Ipv6addr = flex.FlattenIPv6Address(from.Ipv6addr)
	m.Name = flex.FlattenStringPointer(from.Name)
}
