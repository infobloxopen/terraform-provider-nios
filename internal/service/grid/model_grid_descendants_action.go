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

type GridDescendantsActionModel struct {
	OptionWithEa    types.String `tfsdk:"option_with_ea"`
	OptionWithoutEa types.String `tfsdk:"option_without_ea"`
	OptionDeleteEa  types.String `tfsdk:"option_delete_ea"`
}

var GridDescendantsActionAttrTypes = map[string]attr.Type{
	"option_with_ea":    types.StringType,
	"option_without_ea": types.StringType,
	"option_delete_ea":  types.StringType,
}

var GridDescendantsActionResourceSchemaAttributes = map[string]schema.Attribute{
	"option_with_ea": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "This option describes which action must be taken if the extensible attribute exists for both the parent and descendant objects: * INHERIT: inherit the extensible attribute from the parent object. * RETAIN: retain the value of an extensible attribute that was set for the child object. * CONVERT: the value of the extensible attribute must be copied from the parent object.",
	},
	"option_without_ea": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "This option describes which action must be taken if the extensible attribute exists for the parent, but is absent from the descendant object: * INHERIT: inherit the extensible attribute from the parent object. * NOT_INHERIT: do nothing.",
	},
	"option_delete_ea": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "This option describes which action must be taken if the extensible attribute exists for the descendant, but is absent for the parent object: * RETAIN: retain the extensible attribute value for the descendant object. * REMOVE: remove this extensible attribute from the descendant object.",
	},
}

func ExpandGridDescendantsAction(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridDescendantsAction {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridDescendantsActionModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridDescendantsActionModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridDescendantsAction {
	if m == nil {
		return nil
	}
	to := &grid.GridDescendantsAction{
		OptionWithEa:    flex.ExpandStringPointer(m.OptionWithEa),
		OptionWithoutEa: flex.ExpandStringPointer(m.OptionWithoutEa),
		OptionDeleteEa:  flex.ExpandStringPointer(m.OptionDeleteEa),
	}
	return to
}

func FlattenGridDescendantsAction(ctx context.Context, from *grid.GridDescendantsAction, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridDescendantsActionAttrTypes)
	}
	m := GridDescendantsActionModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridDescendantsActionAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridDescendantsActionModel) Flatten(ctx context.Context, from *grid.GridDescendantsAction, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridDescendantsActionModel{}
	}
	m.OptionWithEa = flex.FlattenStringPointer(from.OptionWithEa)
	m.OptionWithoutEa = flex.FlattenStringPointer(from.OptionWithoutEa)
	m.OptionDeleteEa = flex.FlattenStringPointer(from.OptionDeleteEa)
}
