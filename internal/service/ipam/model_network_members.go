package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type NetworkMembersModel struct {
	Struct   types.String `tfsdk:"struct"`
	Ipv4addr types.String `tfsdk:"ipv4addr"`
	Ipv6addr types.String `tfsdk:"ipv6addr"`
	Name     types.String `tfsdk:"name"`
}

var NetworkMembersAttrTypes = map[string]attr.Type{
	"struct":   types.StringType,
	"ipv4addr": types.StringType,
	"ipv6addr": types.StringType,
	"name":     types.StringType,
}

var NetworkMembersResourceSchemaAttributes = map[string]schema.Attribute{
	"struct": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.OneOf("dhcpmember", "msdhcpserver"),
		},
		MarkdownDescription: "The struct type of the object. The value must be one of 'dhcpmember' or 'msdhcpserver'.",
	},
	"ipv4addr": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The IPv4 Address or FQDN of the Microsoft server.",
		Computed:            true,
	},
	"ipv6addr": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The IPv6 Address of the Grid Member.",
		Computed:            true,
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The Grid member name",
		Computed:            true,
	},
}

func ExpandNetworkMembers(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.NetworkMembers {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkMembersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NetworkMembersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.NetworkMembers {
	if m == nil {
		return nil
	}
	to := &ipam.NetworkMembers{
		Struct:   flex.ExpandStringPointer(m.Struct),
		Ipv4addr: flex.ExpandStringPointer(m.Ipv4addr),
		Ipv6addr: flex.ExpandStringPointer(m.Ipv6addr),
		Name:     flex.ExpandStringPointer(m.Name),
	}
	return to
}

func FlattenNetworkMembers(ctx context.Context, from *ipam.NetworkMembers, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkMembersAttrTypes)
	}
	m := NetworkMembersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkMembersAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NetworkMembersModel) Flatten(ctx context.Context, from *ipam.NetworkMembers, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NetworkMembersModel{}
	}
	m.Struct = flex.FlattenStringPointer(from.Struct)
	m.Ipv4addr = flex.FlattenStringPointer(from.Ipv4addr)
	m.Ipv6addr = flex.FlattenStringPointer(from.Ipv6addr)
	m.Name = flex.FlattenStringPointer(from.Name)
}
