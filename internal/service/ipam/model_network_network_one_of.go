package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/ipam"

	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/flex"
)

type NetworkNetworkOneOfModel struct {
	ObjectFunction   types.String `tfsdk:"_object_function"`
	Parameters       types.Map    `tfsdk:"_parameters"`
	ResultField      types.String `tfsdk:"_result_field"`
	Object           types.String `tfsdk:"_object"`
	ObjectParameters types.Map    `tfsdk:"_object_parameters"`
	ObjectRef        types.String `tfsdk:"_object_ref"`
}

var NetworkNetworkOneOfAttrTypes = map[string]attr.Type{
	"_object_function":   types.StringType,
	"_parameters":        types.MapType{ElemType: types.StringType},
	"_result_field":      types.StringType,
	"_object":            types.StringType,
	"_object_parameters": types.MapType{ElemType: types.StringType},
	"_object_ref":        types.StringType,
}

var NetworkNetworkOneOfResourceSchemaAttributes = map[string]schema.Attribute{
	"_object_function": schema.StringAttribute{
		Optional: true,
	},
	"_parameters": schema.MapAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"_result_field": schema.StringAttribute{
		Optional: true,
	},
	"_object": schema.StringAttribute{
		Optional: true,
	},
	"_object_parameters": schema.MapAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"_object_ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "A WAPI object reference on which the function calls. Either _object or _object_ref must be set.",
	},
}

func ExpandNetworkNetworkOneOf(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.NetworkNetworkOneOf {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkNetworkOneOfModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NetworkNetworkOneOfModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.NetworkNetworkOneOf {
	if m == nil {
		return nil
	}
	to := &ipam.NetworkNetworkOneOf{
		ObjectFunction:   flex.ExpandStringPointer(m.ObjectFunction),
		Parameters:       flex.ExpandFrameworkMapString(ctx, m.Parameters, diags),
		ResultField:      flex.ExpandStringPointer(m.ResultField),
		Object:           flex.ExpandStringPointer(m.Object),
		ObjectParameters: flex.ExpandFrameworkMapString(ctx, m.ObjectParameters, diags),
		ObjectRef:        flex.ExpandStringPointer(m.ObjectRef),
	}
	return to
}

func FlattenNetworkNetworkOneOf(ctx context.Context, from *ipam.NetworkNetworkOneOf, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkNetworkOneOfAttrTypes)
	}
	m := NetworkNetworkOneOfModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkNetworkOneOfAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NetworkNetworkOneOfModel) Flatten(ctx context.Context, from *ipam.NetworkNetworkOneOf, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NetworkNetworkOneOfModel{}
	}
	m.ObjectFunction = flex.FlattenStringPointer(from.ObjectFunction)
	m.Parameters = flex.FlattenFrameworkMapString(ctx, from.Parameters, diags)
	m.ResultField = flex.FlattenStringPointer(from.ResultField)
	m.Object = flex.FlattenStringPointer(from.Object)
	m.ObjectParameters = flex.FlattenFrameworkMapString(ctx, from.ObjectParameters, diags)
	m.ObjectRef = flex.FlattenStringPointer(from.ObjectRef)
}
