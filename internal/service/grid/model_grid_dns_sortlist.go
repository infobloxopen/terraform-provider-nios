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

type GridDnsSortlistModel struct {
	Address   types.String `tfsdk:"address"`
	MatchList types.List   `tfsdk:"match_list"`
}

var GridDnsSortlistAttrTypes = map[string]attr.Type{
	"address":    types.StringType,
	"match_list": types.ListType{ElemType: types.StringType},
}

var GridDnsSortlistResourceSchemaAttributes = map[string]schema.Attribute{
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

func ExpandGridDnsSortlist(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridDnsSortlist {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridDnsSortlistModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridDnsSortlistModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridDnsSortlist {
	if m == nil {
		return nil
	}
	to := &grid.GridDnsSortlist{
		Address:   flex.ExpandStringPointer(m.Address),
		MatchList: flex.ExpandFrameworkListString(ctx, m.MatchList, diags),
	}
	return to
}

func FlattenGridDnsSortlist(ctx context.Context, from *grid.GridDnsSortlist, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridDnsSortlistAttrTypes)
	}
	m := GridDnsSortlistModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridDnsSortlistAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridDnsSortlistModel) Flatten(ctx context.Context, from *grid.GridDnsSortlist, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridDnsSortlistModel{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.MatchList = flex.FlattenFrameworkListString(ctx, from.MatchList, diags)
}
