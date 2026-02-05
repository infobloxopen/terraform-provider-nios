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

type MemberIpv6StaticRoutesModel struct {
	Address types.String `tfsdk:"address"`
	Cidr    types.Int64  `tfsdk:"cidr"`
	Gateway types.String `tfsdk:"gateway"`
}

var MemberIpv6StaticRoutesAttrTypes = map[string]attr.Type{
	"address": types.StringType,
	"cidr":    types.Int64Type,
	"gateway": types.StringType,
}

var MemberIpv6StaticRoutesResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "IPv6 address.",
	},
	"cidr": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "IPv6 CIDR",
	},
	"gateway": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Gateway address.",
	},
}

func ExpandMemberIpv6StaticRoutes(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberIpv6StaticRoutes {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberIpv6StaticRoutesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberIpv6StaticRoutesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberIpv6StaticRoutes {
	if m == nil {
		return nil
	}
	to := &grid.MemberIpv6StaticRoutes{
		Address: flex.ExpandStringPointer(m.Address),
		Cidr:    flex.ExpandInt64Pointer(m.Cidr),
		Gateway: flex.ExpandStringPointer(m.Gateway),
	}
	return to
}

func FlattenMemberIpv6StaticRoutes(ctx context.Context, from *grid.MemberIpv6StaticRoutes, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberIpv6StaticRoutesAttrTypes)
	}
	m := MemberIpv6StaticRoutesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberIpv6StaticRoutesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberIpv6StaticRoutesModel) Flatten(ctx context.Context, from *grid.MemberIpv6StaticRoutes, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberIpv6StaticRoutesModel{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Cidr = flex.FlattenInt64Pointer(from.Cidr)
	m.Gateway = flex.FlattenStringPointer(from.Gateway)
}
