package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type FuncCallModel struct {
	AttributeName    types.String `tfsdk:"attribute_name"`
	ObjectFunction   types.String `tfsdk:"object_function"`
	Parameters       types.Map    `tfsdk:"parameters"`
	ResultField      types.String `tfsdk:"result_field"`
	Object           types.String `tfsdk:"object"`
	ObjectParameters types.Map    `tfsdk:"object_parameters"`
}

var FuncCallAttrTypes = map[string]attr.Type{
	"attribute_name":    types.StringType,
	"object_function":   types.StringType,
	"parameters":        types.MapType{ElemType: types.StringType},
	"result_field":      types.StringType,
	"object":            types.StringType,
	"object_parameters": types.MapType{ElemType: types.StringType},
}

var FuncCallResourceSchemaAttributes = map[string]schema.Attribute{
	"attribute_name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The attribute to be called.",
	},
	"object_function": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The function to be called.",
	},
	"parameters": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The parameters for the function.",
	},
	"result_field": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The result field of the function.",
	},
	"object": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The object to be called.",
	},
	"object_parameters": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The parameters for the object.",
	},
}

func ExpandFuncCall(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.FuncCall {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m FuncCallModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *FuncCallModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.FuncCall {
	if m == nil {
		return nil
	}
	to := &dns.FuncCall{
		AttributeName:    flex.ExpandString(m.AttributeName),
		ObjectFunction:   flex.ExpandStringPointer(m.ObjectFunction),
		Parameters:       flex.ExpandFrameworkMapString(ctx, m.Parameters, diags),
		ResultField:      flex.ExpandStringPointer(m.ResultField),
		Object:           flex.ExpandStringPointer(m.Object),
		ObjectParameters: flex.ExpandFrameworkMapString(ctx, m.ObjectParameters, diags),
	}
	return to
}

func FlattenFuncCall(ctx context.Context, from *dns.FuncCall, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(FuncCallAttrTypes)
	}
	m := FuncCallModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, FuncCallAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *FuncCallModel) Flatten(ctx context.Context, from *dns.FuncCall, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = FuncCallModel{}
	}
	m.AttributeName = flex.FlattenString(from.AttributeName)
	m.ObjectFunction = flex.FlattenStringPointer(from.ObjectFunction)
	m.Parameters = flex.FlattenFrameworkMapString(ctx, from.Parameters, diags)
	m.ResultField = flex.FlattenStringPointer(from.ResultField)
	m.Object = flex.FlattenStringPointer(from.Object)
	m.ObjectParameters = flex.FlattenFrameworkMapString(ctx, from.ObjectParameters, diags)
}
