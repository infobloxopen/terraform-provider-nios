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

type GriddnsscavengingsettingsExpressionListModel struct {
	Op      types.String `tfsdk:"op"`
	Op1     types.String `tfsdk:"op1"`
	Op1Type types.String `tfsdk:"op1_type"`
	Op2     types.String `tfsdk:"op2"`
	Op2Type types.String `tfsdk:"op2_type"`
}

var GriddnsscavengingsettingsExpressionListAttrTypes = map[string]attr.Type{
	"op":       types.StringType,
	"op1":      types.StringType,
	"op1_type": types.StringType,
	"op2":      types.StringType,
	"op2_type": types.StringType,
}

var GriddnsscavengingsettingsExpressionListResourceSchemaAttributes = map[string]schema.Attribute{
	"op": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The operation name.",
	},
	"op1": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The first operand value.",
	},
	"op1_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The first operand type.",
	},
	"op2": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The second operand value.",
	},
	"op2_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The second operand type.",
	},
}

func ExpandGriddnsscavengingsettingsExpressionList(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GriddnsscavengingsettingsExpressionList {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GriddnsscavengingsettingsExpressionListModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GriddnsscavengingsettingsExpressionListModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GriddnsscavengingsettingsExpressionList {
	if m == nil {
		return nil
	}
	to := &grid.GriddnsscavengingsettingsExpressionList{
		Op:      flex.ExpandStringPointer(m.Op),
		Op1:     flex.ExpandStringPointer(m.Op1),
		Op1Type: flex.ExpandStringPointer(m.Op1Type),
		Op2:     flex.ExpandStringPointer(m.Op2),
		Op2Type: flex.ExpandStringPointer(m.Op2Type),
	}
	return to
}

func FlattenGriddnsscavengingsettingsExpressionList(ctx context.Context, from *grid.GriddnsscavengingsettingsExpressionList, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GriddnsscavengingsettingsExpressionListAttrTypes)
	}
	m := GriddnsscavengingsettingsExpressionListModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GriddnsscavengingsettingsExpressionListAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GriddnsscavengingsettingsExpressionListModel) Flatten(ctx context.Context, from *grid.GriddnsscavengingsettingsExpressionList, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GriddnsscavengingsettingsExpressionListModel{}
	}
	m.Op = flex.FlattenStringPointer(from.Op)
	m.Op1 = flex.FlattenStringPointer(from.Op1)
	m.Op1Type = flex.FlattenStringPointer(from.Op1Type)
	m.Op2 = flex.FlattenStringPointer(from.Op2)
	m.Op2Type = flex.FlattenStringPointer(from.Op2Type)
}
