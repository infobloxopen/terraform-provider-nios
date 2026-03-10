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

type SyslogEndpointTemplateInstanceModel struct {
	Template   types.String `tfsdk:"template"`
	Parameters types.List   `tfsdk:"parameters"`
}

var SyslogEndpointTemplateInstanceAttrTypes = map[string]attr.Type{
	"template":   types.StringType,
	"parameters": types.ListType{ElemType: types.ObjectType{AttrTypes: SyslogendpointtemplateinstanceParametersAttrTypes}},
}

var SyslogEndpointTemplateInstanceResourceSchemaAttributes = map[string]schema.Attribute{
	"template": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the REST API template parameter.",
	},
	"parameters": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: SyslogendpointtemplateinstanceParametersResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The notification REST template parameters.",
	},
}

func ExpandSyslogEndpointTemplateInstance(ctx context.Context, o types.Object, diags *diag.Diagnostics) *misc.SyslogEndpointTemplateInstance {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m SyslogEndpointTemplateInstanceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *SyslogEndpointTemplateInstanceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *misc.SyslogEndpointTemplateInstance {
	if m == nil {
		return nil
	}
	to := &misc.SyslogEndpointTemplateInstance{
		Template:   flex.ExpandStringPointer(m.Template),
		Parameters: flex.ExpandFrameworkListNestedBlock(ctx, m.Parameters, diags, ExpandSyslogendpointtemplateinstanceParameters),
	}
	return to
}

func FlattenSyslogEndpointTemplateInstance(ctx context.Context, from *misc.SyslogEndpointTemplateInstance, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SyslogEndpointTemplateInstanceAttrTypes)
	}
	m := SyslogEndpointTemplateInstanceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, SyslogEndpointTemplateInstanceAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *SyslogEndpointTemplateInstanceModel) Flatten(ctx context.Context, from *misc.SyslogEndpointTemplateInstance, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = SyslogEndpointTemplateInstanceModel{}
	}
	m.Template = flex.FlattenStringPointer(from.Template)
	m.Parameters = flex.FlattenFrameworkListNestedBlock(ctx, from.Parameters, SyslogendpointtemplateinstanceParametersAttrTypes, diags, FlattenSyslogendpointtemplateinstanceParameters)
}
