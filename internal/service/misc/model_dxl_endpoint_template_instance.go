package misc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/misc"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type DxlEndpointTemplateInstanceModel struct {
	Template   types.String `tfsdk:"template"`
	Parameters types.List   `tfsdk:"parameters"`
}

var DxlEndpointTemplateInstanceAttrTypes = map[string]attr.Type{
	"template":   types.StringType,
	"parameters": types.ListType{ElemType: types.ObjectType{AttrTypes: DxlendpointtemplateinstanceParametersAttrTypes}},
}

var DxlEndpointTemplateInstanceResourceSchemaAttributes = map[string]schema.Attribute{
	"template": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the REST API template parameter.",
	},
	"parameters": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: DxlendpointtemplateinstanceParametersResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The notification REST template parameters.",
	},
}

func ExpandDxlEndpointTemplateInstance(ctx context.Context, o types.Object, diags *diag.Diagnostics) *misc.DxlEndpointTemplateInstance {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DxlEndpointTemplateInstanceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DxlEndpointTemplateInstanceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *misc.DxlEndpointTemplateInstance {
	if m == nil {
		return nil
	}
	to := &misc.DxlEndpointTemplateInstance{
		Template:   flex.ExpandStringPointer(m.Template),
		Parameters: flex.ExpandFrameworkListNestedBlock(ctx, m.Parameters, diags, ExpandDxlendpointtemplateinstanceParameters),
	}
	return to
}

func FlattenDxlEndpointTemplateInstance(ctx context.Context, from *misc.DxlEndpointTemplateInstance, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DxlEndpointTemplateInstanceAttrTypes)
	}
	m := DxlEndpointTemplateInstanceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DxlEndpointTemplateInstanceAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DxlEndpointTemplateInstanceModel) Flatten(ctx context.Context, from *misc.DxlEndpointTemplateInstance, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DxlEndpointTemplateInstanceModel{}
	}
	m.Template = flex.FlattenStringPointer(from.Template)
	m.Parameters = flex.FlattenFrameworkListNestedBlock(ctx, from.Parameters, DxlendpointtemplateinstanceParametersAttrTypes, diags, FlattenDxlendpointtemplateinstanceParameters)
}
