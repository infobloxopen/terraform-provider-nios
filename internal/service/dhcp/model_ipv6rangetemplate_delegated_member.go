package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type Ipv6rangetemplateDelegatedMemberModel struct {
	Ipv4addr types.String `tfsdk:"ipv4addr"`
	Ipv6addr types.String `tfsdk:"ipv6addr"`
	Name     types.String `tfsdk:"name"`
}

var Ipv6rangetemplateDelegatedMemberAttrTypes = map[string]attr.Type{
	"ipv4addr": types.StringType,
	"ipv6addr": types.StringType,
	"name":     types.StringType,
}

var Ipv6rangetemplateDelegatedMemberResourceSchemaAttributes = map[string]schema.Attribute{
	"ipv4addr": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The IPv4 Address of the Grid Member.",
	},
	"ipv6addr": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The IPv6 Address of the Grid Member.",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The Grid member name",
	},
}

func ExpandIpv6rangetemplateDelegatedMember(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dhcp.Ipv6rangetemplateDelegatedMember {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6rangetemplateDelegatedMemberModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *Ipv6rangetemplateDelegatedMemberModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.Ipv6rangetemplateDelegatedMember {
	if m == nil {
		return nil
	}
	to := &dhcp.Ipv6rangetemplateDelegatedMember{
		Ipv4addr: flex.ExpandStringPointer(m.Ipv4addr),
		Ipv6addr: flex.ExpandStringPointer(m.Ipv6addr),
		Name:     flex.ExpandStringPointer(m.Name),
	}
	return to
}

func FlattenIpv6rangetemplateDelegatedMember(ctx context.Context, from *dhcp.Ipv6rangetemplateDelegatedMember, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6rangetemplateDelegatedMemberAttrTypes)
	}
	m := Ipv6rangetemplateDelegatedMemberModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, Ipv6rangetemplateDelegatedMemberAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Ipv6rangetemplateDelegatedMemberModel) Flatten(ctx context.Context, from *dhcp.Ipv6rangetemplateDelegatedMember, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Ipv6rangetemplateDelegatedMemberModel{}
	}
	m.Ipv4addr = flex.FlattenStringPointer(from.Ipv4addr)
	m.Ipv6addr = flex.FlattenStringPointer(from.Ipv6addr)
	m.Name = flex.FlattenStringPointer(from.Name)
}
