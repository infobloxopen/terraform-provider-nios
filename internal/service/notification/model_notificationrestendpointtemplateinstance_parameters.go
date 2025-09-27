package notification

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/notification"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type NotificationrestendpointtemplateinstanceParametersModel struct {
	Name         types.String `tfsdk:"name"`
	Value        types.String `tfsdk:"value"`
	DefaultValue types.String `tfsdk:"default_value"`
	Syntax       types.String `tfsdk:"syntax"`
}

var NotificationrestendpointtemplateinstanceParametersAttrTypes = map[string]attr.Type{
	"name":          types.StringType,
	"value":         types.StringType,
	"default_value": types.StringType,
	"syntax":        types.StringType,
}

var NotificationrestendpointtemplateinstanceParametersResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of the REST API template parameter.",
	},
	"value": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The value of the REST API template parameter.",
	},
	"default_value": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The default value of the REST API template parameter.",
	},
	"syntax": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.OneOf("BOOL", "INT", "STR"),
		},
		MarkdownDescription: "The syntax of the REST API template parameter.",
	},
}

func ExpandNotificationrestendpointtemplateinstanceParameters(ctx context.Context, o types.Object, diags *diag.Diagnostics) *notification.NotificationrestendpointtemplateinstanceParameters {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NotificationrestendpointtemplateinstanceParametersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NotificationrestendpointtemplateinstanceParametersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *notification.NotificationrestendpointtemplateinstanceParameters {
	if m == nil {
		return nil
	}
	to := &notification.NotificationrestendpointtemplateinstanceParameters{
		Name:   flex.ExpandStringPointer(m.Name),
		Value:  flex.ExpandStringPointer(m.Value),
		Syntax: flex.ExpandStringPointer(m.Syntax),
	}
	return to
}

func FlattenNotificationrestendpointtemplateinstanceParameters(ctx context.Context, from *notification.NotificationrestendpointtemplateinstanceParameters, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NotificationrestendpointtemplateinstanceParametersAttrTypes)
	}
	m := NotificationrestendpointtemplateinstanceParametersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NotificationrestendpointtemplateinstanceParametersAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NotificationrestendpointtemplateinstanceParametersModel) Flatten(ctx context.Context, from *notification.NotificationrestendpointtemplateinstanceParameters, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NotificationrestendpointtemplateinstanceParametersModel{}
	}
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Value = flex.FlattenStringPointer(from.Value)
	m.DefaultValue = flex.FlattenStringPointer(from.DefaultValue)
	m.Syntax = flex.FlattenStringPointer(from.Syntax)
}
