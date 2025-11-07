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

type PxgridEndpointTemplateInstanceModel struct {
	Template   types.String `tfsdk:"template"`
	Parameters types.List   `tfsdk:"parameters"`
}

var PxgridEndpointTemplateInstanceAttrTypes = map[string]attr.Type{
	"template":   types.StringType,
	"parameters": types.ListType{ElemType: types.ObjectType{AttrTypes: PxgridendpointtemplateinstanceParametersAttrTypes}},
}

var PxgridEndpointTemplateInstanceResourceSchemaAttributes = map[string]schema.Attribute{
	"template": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the REST API template parameter.",
	},
	"parameters": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: PxgridendpointtemplateinstanceParametersResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The notification REST template parameters.",
	},
}

func ExpandPxgridEndpointTemplateInstance(ctx context.Context, o types.Object, diags *diag.Diagnostics) *misc.PxgridEndpointTemplateInstance {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m PxgridEndpointTemplateInstanceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *PxgridEndpointTemplateInstanceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *misc.PxgridEndpointTemplateInstance {
	if m == nil {
		return nil
	}
	to := &misc.PxgridEndpointTemplateInstance{
		Template:   flex.ExpandStringPointer(m.Template),
		Parameters: flex.ExpandFrameworkListNestedBlock(ctx, m.Parameters, diags, ExpandPxgridendpointtemplateinstanceParameters),
	}
	return to
}

func FlattenPxgridEndpointTemplateInstance(ctx context.Context, from *misc.PxgridEndpointTemplateInstance, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(PxgridEndpointTemplateInstanceAttrTypes)
	}
	m := PxgridEndpointTemplateInstanceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, PxgridEndpointTemplateInstanceAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *PxgridEndpointTemplateInstanceModel) Flatten(ctx context.Context, from *misc.PxgridEndpointTemplateInstance, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = PxgridEndpointTemplateInstanceModel{}
	}
	m.Template = flex.FlattenStringPointer(from.Template)
	m.Parameters = flex.FlattenFrameworkListNestedBlock(ctx, from.Parameters, PxgridendpointtemplateinstanceParametersAttrTypes, diags, FlattenPxgridendpointtemplateinstanceParameters)
}
