package misc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/misc"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type DxlendpointtemplateinstanceParametersModel struct {
	Name         types.String `tfsdk:"name"`
	Value        types.String `tfsdk:"value"`
	DefaultValue types.String `tfsdk:"default_value"`
	Syntax       types.String `tfsdk:"syntax"`
}

var DxlendpointtemplateinstanceParametersAttrTypes = map[string]attr.Type{
	"name":          types.StringType,
	"value":         types.StringType,
	"default_value": types.StringType,
	"syntax":        types.StringType,
}

var DxlendpointtemplateinstanceParametersResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the REST API template parameter.",
	},
	"value": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The value of the REST API template parameter.",
	},
	"default_value": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The default value of the REST API template parameter.",
	},
	"syntax": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The syntax of the REST API template parameter.",
	},
}

func ExpandDxlendpointtemplateinstanceParameters(ctx context.Context, o types.Object, diags *diag.Diagnostics) *misc.DxlendpointtemplateinstanceParameters {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DxlendpointtemplateinstanceParametersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DxlendpointtemplateinstanceParametersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *misc.DxlendpointtemplateinstanceParameters {
	if m == nil {
		return nil
	}
	to := &misc.DxlendpointtemplateinstanceParameters{
		Name:   flex.ExpandStringPointer(m.Name),
		Value:  flex.ExpandStringPointer(m.Value),
		Syntax: flex.ExpandStringPointer(m.Syntax),
	}
	return to
}

func FlattenDxlendpointtemplateinstanceParameters(ctx context.Context, from *misc.DxlendpointtemplateinstanceParameters, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DxlendpointtemplateinstanceParametersAttrTypes)
	}
	m := DxlendpointtemplateinstanceParametersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DxlendpointtemplateinstanceParametersAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DxlendpointtemplateinstanceParametersModel) Flatten(ctx context.Context, from *misc.DxlendpointtemplateinstanceParameters, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DxlendpointtemplateinstanceParametersModel{}
	}
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Value = flex.FlattenStringPointer(from.Value)
	m.DefaultValue = flex.FlattenStringPointer(from.DefaultValue)
	m.Syntax = flex.FlattenStringPointer(from.Syntax)
}
