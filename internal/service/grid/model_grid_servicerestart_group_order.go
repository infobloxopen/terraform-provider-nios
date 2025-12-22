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

type GridServicerestartGroupOrderModel struct {
	Ref    types.String `tfsdk:"ref"`
	Groups types.List   `tfsdk:"groups"`
}

var GridServicerestartGroupOrderAttrTypes = map[string]attr.Type{
	"ref":    types.StringType,
	"groups": types.ListType{ElemType: types.StringType},
}

var GridServicerestartGroupOrderResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"groups": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The ordered list of the Service Restart Group.",
	},
}

func ExpandGridServicerestartGroupOrder(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridServicerestartGroupOrder {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridServicerestartGroupOrderModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridServicerestartGroupOrderModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridServicerestartGroupOrder {
	if m == nil {
		return nil
	}
	to := &grid.GridServicerestartGroupOrder{
		Ref:    flex.ExpandStringPointer(m.Ref),
		Groups: flex.ExpandFrameworkListString(ctx, m.Groups, diags),
	}
	return to
}

func FlattenGridServicerestartGroupOrder(ctx context.Context, from *grid.GridServicerestartGroupOrder, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridServicerestartGroupOrderAttrTypes)
	}
	m := GridServicerestartGroupOrderModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridServicerestartGroupOrderAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridServicerestartGroupOrderModel) Flatten(ctx context.Context, from *grid.GridServicerestartGroupOrder, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridServicerestartGroupOrderModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Groups = flex.FlattenFrameworkListString(ctx, from.Groups, diags)
}
