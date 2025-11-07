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

type MemberDnsRecursiveQueryListModel struct {
	Address    types.String `tfsdk:"address"`
	Permission types.String `tfsdk:"permission"`
}

var MemberDnsRecursiveQueryListAttrTypes = map[string]attr.Type{
	"address":    types.StringType,
	"permission": types.StringType,
}

var MemberDnsRecursiveQueryListResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The address this rule applies to or \"Any\".",
	},
	"permission": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The permission to use for this address.",
	},
}

func ExpandMemberDnsRecursiveQueryList(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberDnsRecursiveQueryList {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberDnsRecursiveQueryListModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberDnsRecursiveQueryListModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberDnsRecursiveQueryList {
	if m == nil {
		return nil
	}
	to := &grid.MemberDnsRecursiveQueryList{
		Address:    flex.ExpandStringPointer(m.Address),
		Permission: flex.ExpandStringPointer(m.Permission),
	}
	return to
}

func FlattenMemberDnsRecursiveQueryList(ctx context.Context, from *grid.MemberDnsRecursiveQueryList, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberDnsRecursiveQueryListAttrTypes)
	}
	m := MemberDnsRecursiveQueryListModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberDnsRecursiveQueryListAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberDnsRecursiveQueryListModel) Flatten(ctx context.Context, from *grid.MemberDnsRecursiveQueryList, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberDnsRecursiveQueryListModel{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Permission = flex.FlattenStringPointer(from.Permission)
}
