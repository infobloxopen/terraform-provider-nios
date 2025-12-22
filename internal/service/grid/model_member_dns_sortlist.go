package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type MemberDnsSortlistModel struct {
	Address   types.String `tfsdk:"address"`
	MatchList types.List   `tfsdk:"match_list"`
}

var MemberDnsSortlistAttrTypes = map[string]attr.Type{
	"address":    types.StringType,
	"match_list": types.ListType{ElemType: types.StringType},
}

var MemberDnsSortlistResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The source address of a sortlist object.",
	},
	"match_list": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The match list of a sortlist.",
	},
}

func ExpandMemberDnsSortlist(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberDnsSortlist {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberDnsSortlistModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberDnsSortlistModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberDnsSortlist {
	if m == nil {
		return nil
	}
	to := &grid.MemberDnsSortlist{
		Address:   flex.ExpandStringPointer(m.Address),
		MatchList: flex.ExpandFrameworkListString(ctx, m.MatchList, diags),
	}
	return to
}

func FlattenMemberDnsSortlist(ctx context.Context, from *grid.MemberDnsSortlist, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberDnsSortlistAttrTypes)
	}
	m := MemberDnsSortlistModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberDnsSortlistAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberDnsSortlistModel) Flatten(ctx context.Context, from *grid.MemberDnsSortlist, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberDnsSortlistModel{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.MatchList = flex.FlattenFrameworkListString(ctx, from.MatchList, diags)
}
