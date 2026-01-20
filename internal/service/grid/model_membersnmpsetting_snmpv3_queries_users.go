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

type MembersnmpsettingSnmpv3QueriesUsersModel struct {
	User    types.String `tfsdk:"user"`
	Comment types.String `tfsdk:"comment"`
}

var MembersnmpsettingSnmpv3QueriesUsersAttrTypes = map[string]attr.Type{
	"user":    types.StringType,
	"comment": types.StringType,
}

var MembersnmpsettingSnmpv3QueriesUsersResourceSchemaAttributes = map[string]schema.Attribute{
	"user": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The SNMPv3 user.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "A descriptive comment for this queries user.",
	},
}

func ExpandMembersnmpsettingSnmpv3QueriesUsers(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MembersnmpsettingSnmpv3QueriesUsers {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MembersnmpsettingSnmpv3QueriesUsersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MembersnmpsettingSnmpv3QueriesUsersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MembersnmpsettingSnmpv3QueriesUsers {
	if m == nil {
		return nil
	}
	to := &grid.MembersnmpsettingSnmpv3QueriesUsers{
		User:    flex.ExpandStringPointer(m.User),
		Comment: flex.ExpandStringPointer(m.Comment),
	}
	return to
}

func FlattenMembersnmpsettingSnmpv3QueriesUsers(ctx context.Context, from *grid.MembersnmpsettingSnmpv3QueriesUsers, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MembersnmpsettingSnmpv3QueriesUsersAttrTypes)
	}
	m := MembersnmpsettingSnmpv3QueriesUsersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MembersnmpsettingSnmpv3QueriesUsersAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MembersnmpsettingSnmpv3QueriesUsersModel) Flatten(ctx context.Context, from *grid.MembersnmpsettingSnmpv3QueriesUsers, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MembersnmpsettingSnmpv3QueriesUsersModel{}
	}
	m.User = flex.FlattenStringPointer(from.User)
	m.Comment = flex.FlattenStringPointer(from.Comment)
}
