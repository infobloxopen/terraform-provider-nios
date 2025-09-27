package notification

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/notification"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type NotificationRestEndpointTemplateInstanceModel struct {
	Template   types.String `tfsdk:"template"`
	Parameters types.List   `tfsdk:"parameters"`
}

var NotificationRestEndpointTemplateInstanceAttrTypes = map[string]attr.Type{
	"template":   types.StringType,
	"parameters": types.ListType{ElemType: types.ObjectType{AttrTypes: NotificationrestendpointtemplateinstanceParametersAttrTypes}},
}

var NotificationRestEndpointTemplateInstanceResourceSchemaAttributes = map[string]schema.Attribute{
	"template": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of the REST API template parameter.",
	},
	"parameters": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NotificationrestendpointtemplateinstanceParametersResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The notification REST template parameters.",
	},
}

func ExpandNotificationRestEndpointTemplateInstance(ctx context.Context, o types.Object, diags *diag.Diagnostics) *notification.NotificationRestEndpointTemplateInstance {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NotificationRestEndpointTemplateInstanceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NotificationRestEndpointTemplateInstanceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *notification.NotificationRestEndpointTemplateInstance {
	if m == nil {
		return nil
	}
	to := &notification.NotificationRestEndpointTemplateInstance{
		Template:   flex.ExpandStringPointer(m.Template),
		Parameters: flex.ExpandFrameworkListNestedBlock(ctx, m.Parameters, diags, ExpandNotificationrestendpointtemplateinstanceParameters),
	}
	return to
}

func FlattenNotificationRestEndpointTemplateInstance(ctx context.Context, from *notification.NotificationRestEndpointTemplateInstance, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NotificationRestEndpointTemplateInstanceAttrTypes)
	}
	m := NotificationRestEndpointTemplateInstanceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NotificationRestEndpointTemplateInstanceAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NotificationRestEndpointTemplateInstanceModel) Flatten(ctx context.Context, from *notification.NotificationRestEndpointTemplateInstance, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NotificationRestEndpointTemplateInstanceModel{}
	}
	m.Template = flex.FlattenStringPointer(from.Template)
	m.Parameters = flex.FlattenFrameworkListNestedBlock(ctx, from.Parameters, NotificationrestendpointtemplateinstanceParametersAttrTypes, diags, FlattenNotificationrestendpointtemplateinstanceParameters)
}
