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

type GmcgroupMembersModel struct {
	Member types.String `tfsdk:"member"`
}

var GmcgroupMembersAttrTypes = map[string]attr.Type{
	"member": types.StringType,
}

var GmcgroupMembersResourceSchemaAttributes = map[string]schema.Attribute{
	"member": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The gmcgroup member name.",
	},
}

func ExpandGmcgroupMembers(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GmcgroupMembers {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GmcgroupMembersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GmcgroupMembersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GmcgroupMembers {
	if m == nil {
		return nil
	}
	to := &grid.GmcgroupMembers{
		Member: flex.ExpandStringPointer(m.Member),
	}
	return to
}

func FlattenGmcgroupMembers(ctx context.Context, from *grid.GmcgroupMembers, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GmcgroupMembersAttrTypes)
	}
	m := GmcgroupMembersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GmcgroupMembersAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GmcgroupMembersModel) Flatten(ctx context.Context, from *grid.GmcgroupMembers, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GmcgroupMembersModel{}
	}
	m.Member = flex.FlattenStringPointer(from.Member)
}
