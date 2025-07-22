package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/ipam"

	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/flex"
)

type NetworkcontainercloudinfoDelegatedMemberModel struct {
	Ipv4addr types.String `tfsdk:"ipv4addr"`
	Ipv6addr types.String `tfsdk:"ipv6addr"`
	Name     types.String `tfsdk:"name"`
}

var NetworkcontainercloudinfoDelegatedMemberAttrTypes = map[string]attr.Type{
	"ipv4addr": types.StringType,
	"ipv6addr": types.StringType,
	"name":     types.StringType,
}

var NetworkcontainercloudinfoDelegatedMemberResourceSchemaAttributes = map[string]schema.Attribute{
	"ipv4addr": schema.StringAttribute{
		MarkdownDescription: "The IPv4 Address of the Grid Member.",
		Computed:            true,
	},
	"ipv6addr": schema.StringAttribute{
		MarkdownDescription: "The IPv6 Address of the Grid Member.",
		Computed:            true,
	},
	"name": schema.StringAttribute{
		MarkdownDescription: "The Grid member name",
		Computed:            true,
	},
}

func ExpandNetworkcontainercloudinfoDelegatedMember(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.NetworkcontainercloudinfoDelegatedMember {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkcontainercloudinfoDelegatedMemberModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NetworkcontainercloudinfoDelegatedMemberModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.NetworkcontainercloudinfoDelegatedMember {
	if m == nil {
		return nil
	}
	to := &ipam.NetworkcontainercloudinfoDelegatedMember{
		Ipv4addr: flex.ExpandStringPointer(m.Ipv4addr),
		Ipv6addr: flex.ExpandStringPointer(m.Ipv6addr),
		Name:     flex.ExpandStringPointer(m.Name),
	}
	return to
}

func FlattenNetworkcontainercloudinfoDelegatedMember(ctx context.Context, from *ipam.NetworkcontainercloudinfoDelegatedMember, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkcontainercloudinfoDelegatedMemberAttrTypes)
	}
	m := NetworkcontainercloudinfoDelegatedMemberModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkcontainercloudinfoDelegatedMemberAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NetworkcontainercloudinfoDelegatedMemberModel) Flatten(ctx context.Context, from *ipam.NetworkcontainercloudinfoDelegatedMember, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NetworkcontainercloudinfoDelegatedMemberModel{}
	}
	m.Ipv4addr = flex.FlattenStringPointer(from.Ipv4addr)
	m.Ipv6addr = flex.FlattenStringPointer(from.Ipv6addr)
	m.Name = flex.FlattenStringPointer(from.Name)
}
