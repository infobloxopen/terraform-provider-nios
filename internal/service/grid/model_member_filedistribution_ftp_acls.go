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

type MemberFiledistributionFtpAclsModel struct {
	Address    types.String `tfsdk:"address"`
	Permission types.String `tfsdk:"permission"`
}

var MemberFiledistributionFtpAclsAttrTypes = map[string]attr.Type{
	"address":    types.StringType,
	"permission": types.StringType,
}

var MemberFiledistributionFtpAclsResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The address this rule applies to or \"Any\".",
	},
	"permission": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The permission to use for this address.",
	},
}

func ExpandMemberFiledistributionFtpAcls(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberFiledistributionFtpAcls {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberFiledistributionFtpAclsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberFiledistributionFtpAclsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberFiledistributionFtpAcls {
	if m == nil {
		return nil
	}
	to := &grid.MemberFiledistributionFtpAcls{
		Address:    flex.ExpandStringPointer(m.Address),
		Permission: flex.ExpandStringPointer(m.Permission),
	}
	return to
}

func FlattenMemberFiledistributionFtpAcls(ctx context.Context, from *grid.MemberFiledistributionFtpAcls, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberFiledistributionFtpAclsAttrTypes)
	}
	m := MemberFiledistributionFtpAclsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberFiledistributionFtpAclsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberFiledistributionFtpAclsModel) Flatten(ctx context.Context, from *grid.MemberFiledistributionFtpAcls, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberFiledistributionFtpAclsModel{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Permission = flex.FlattenStringPointer(from.Permission)
}
