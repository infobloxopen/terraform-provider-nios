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

type GridDnsLastQueriedAclModel struct {
	Address    types.String `tfsdk:"address"`
	Permission types.String `tfsdk:"permission"`
}

var GridDnsLastQueriedAclAttrTypes = map[string]attr.Type{
	"address":    types.StringType,
	"permission": types.StringType,
}

var GridDnsLastQueriedAclResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The address this rule applies to or \"Any\".",
	},
	"permission": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The permission to use for this address.",
	},
}

func ExpandGridDnsLastQueriedAcl(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridDnsLastQueriedAcl {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridDnsLastQueriedAclModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridDnsLastQueriedAclModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridDnsLastQueriedAcl {
	if m == nil {
		return nil
	}
	to := &grid.GridDnsLastQueriedAcl{
		Address:    flex.ExpandStringPointer(m.Address),
		Permission: flex.ExpandStringPointer(m.Permission),
	}
	return to
}

func FlattenGridDnsLastQueriedAcl(ctx context.Context, from *grid.GridDnsLastQueriedAcl, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridDnsLastQueriedAclAttrTypes)
	}
	m := GridDnsLastQueriedAclModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridDnsLastQueriedAclAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridDnsLastQueriedAclModel) Flatten(ctx context.Context, from *grid.GridDnsLastQueriedAcl, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridDnsLastQueriedAclModel{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Permission = flex.FlattenStringPointer(from.Permission)
}
