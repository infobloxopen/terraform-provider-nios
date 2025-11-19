package notification

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/notification"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	"github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/plancontrol"
)

type NotificationruletemplateinstanceParametersModel struct {
	Name         types.String `tfsdk:"name"`
	Value        types.String `tfsdk:"value"`
	DefaultValue types.String `tfsdk:"default_value"`
	Syntax       types.String `tfsdk:"syntax"`
}

var NotificationruletemplateinstanceParametersAttrTypes = map[string]attr.Type{
	"name":          types.StringType,
	"value":         types.StringType,
	"default_value": types.StringType,
	"syntax":        types.StringType,
}

var NotificationruletemplateinstanceParametersResourceSchemaAttributes = map[string]schema.Attribute{
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
		PlanModifiers: []planmodifier.String{
			plancontrol.UseStateForUnknownString(),
		},
	},
	"syntax": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("BOOL", "INT", "STR"),
		},
		MarkdownDescription: "The syntax of the REST API template parameter.",
	},
}

func ExpandNotificationruletemplateinstanceParameters(ctx context.Context, o types.Object, diags *diag.Diagnostics) *notification.NotificationruletemplateinstanceParameters {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NotificationruletemplateinstanceParametersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NotificationruletemplateinstanceParametersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *notification.NotificationruletemplateinstanceParameters {
	if m == nil {
		return nil
	}
	to := &notification.NotificationruletemplateinstanceParameters{
		Name:   flex.ExpandStringPointer(m.Name),
		Value:  flex.ExpandStringPointer(m.Value),
		Syntax: flex.ExpandStringPointer(m.Syntax),
	}
	return to
}

func FlattenNotificationruletemplateinstanceParameters(ctx context.Context, from *notification.NotificationruletemplateinstanceParameters, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NotificationruletemplateinstanceParametersAttrTypes)
	}
	m := NotificationruletemplateinstanceParametersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NotificationruletemplateinstanceParametersAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NotificationruletemplateinstanceParametersModel) Flatten(ctx context.Context, from *notification.NotificationruletemplateinstanceParameters, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NotificationruletemplateinstanceParametersModel{}
	}
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Value = flex.FlattenStringPointer(from.Value)
	m.DefaultValue = flex.FlattenStringPointer(from.DefaultValue)
	m.Syntax = flex.FlattenStringPointer(from.Syntax)
}
