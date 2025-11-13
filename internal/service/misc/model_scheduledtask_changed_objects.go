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

type ScheduledtaskChangedObjectsModel struct {
	Action     types.String `tfsdk:"action"`
	Name       types.String `tfsdk:"name"`
	Type       types.String `tfsdk:"type"`
	ObjectType types.String `tfsdk:"object_type"`
	Properties types.List   `tfsdk:"properties"`
}

var ScheduledtaskChangedObjectsAttrTypes = map[string]attr.Type{
	"action":      types.StringType,
	"name":        types.StringType,
	"type":        types.StringType,
	"object_type": types.StringType,
	"properties":  types.ListType{ElemType: types.StringType},
}

var ScheduledtaskChangedObjectsResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "This is a description of the action that is applied to this object.",
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The object name.",
	},
	"type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "A value of the object type, this may contain objects that are not yet available in WAPI.",
	},
	"object_type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The object type. This is undefined if the object is not yet supported.",
	},
	"properties": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Computed:            true,
		MarkdownDescription: "A list of properties that are being changed.",
	},
}

func ExpandScheduledtaskChangedObjects(ctx context.Context, o types.Object, diags *diag.Diagnostics) *misc.ScheduledtaskChangedObjects {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ScheduledtaskChangedObjectsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ScheduledtaskChangedObjectsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *misc.ScheduledtaskChangedObjects {
	if m == nil {
		return nil
	}
	to := &misc.ScheduledtaskChangedObjects{}
	return to
}

func FlattenScheduledtaskChangedObjects(ctx context.Context, from *misc.ScheduledtaskChangedObjects, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ScheduledtaskChangedObjectsAttrTypes)
	}
	m := ScheduledtaskChangedObjectsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ScheduledtaskChangedObjectsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ScheduledtaskChangedObjectsModel) Flatten(ctx context.Context, from *misc.ScheduledtaskChangedObjects, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ScheduledtaskChangedObjectsModel{}
	}
	m.Action = flex.FlattenStringPointer(from.Action)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Type = flex.FlattenStringPointer(from.Type)
	m.ObjectType = flex.FlattenStringPointer(from.ObjectType)
	m.Properties = flex.FlattenFrameworkListString(ctx, from.Properties, diags)
}
