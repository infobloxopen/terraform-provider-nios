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

type MemberFiledistributionTftpAclsModel struct {
	Address    types.String `tfsdk:"address"`
	Permission types.String `tfsdk:"permission"`
}

var MemberFiledistributionTftpAclsAttrTypes = map[string]attr.Type{
	"address":    types.StringType,
	"permission": types.StringType,
}

var MemberFiledistributionTftpAclsResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The address this rule applies to or \"Any\".",
	},
	"permission": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The permission to use for this address.",
	},
}

func ExpandMemberFiledistributionTftpAcls(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberFiledistributionTftpAcls {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberFiledistributionTftpAclsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberFiledistributionTftpAclsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberFiledistributionTftpAcls {
	if m == nil {
		return nil
	}
	to := &grid.MemberFiledistributionTftpAcls{
		Address:    flex.ExpandStringPointer(m.Address),
		Permission: flex.ExpandStringPointer(m.Permission),
	}
	return to
}

func FlattenMemberFiledistributionTftpAcls(ctx context.Context, from *grid.MemberFiledistributionTftpAcls, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberFiledistributionTftpAclsAttrTypes)
	}
	m := MemberFiledistributionTftpAclsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberFiledistributionTftpAclsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberFiledistributionTftpAclsModel) Flatten(ctx context.Context, from *grid.MemberFiledistributionTftpAcls, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberFiledistributionTftpAclsModel{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Permission = flex.FlattenStringPointer(from.Permission)
}
