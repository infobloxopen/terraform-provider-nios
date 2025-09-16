package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type GridServicerestartRequestChangedobjectModel struct {
	Ref               types.String `tfsdk:"ref"`
	Action            types.String `tfsdk:"action"`
	ChangedProperties types.List   `tfsdk:"changed_properties"`
	ChangedTime       types.Int64  `tfsdk:"changed_time"`
	ObjectName        types.String `tfsdk:"object_name"`
	ObjectType        types.String `tfsdk:"object_type"`
	UserName          types.String `tfsdk:"user_name"`
}

var GridServicerestartRequestChangedobjectAttrTypes = map[string]attr.Type{
	"ref":                types.StringType,
	"action":             types.StringType,
	"changed_properties": types.ListType{ElemType: types.StringType},
	"changed_time":       types.Int64Type,
	"object_name":        types.StringType,
	"object_type":        types.StringType,
	"user_name":          types.StringType,
}

var GridServicerestartRequestChangedobjectResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"action": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The operation on the changed object.",
	},
	"changed_properties": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Computed:            true,
		MarkdownDescription: "The list of changed properties in the object.",
	},
	"changed_time": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The time when the object was changed.",
	},
	"object_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the changed object.",
	},
	"object_type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The type of the changed object. This is undefined if the object is not supported.",
	},
	"user_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the user who changed the object properties.",
	},
}

func ExpandGridServicerestartRequestChangedobject(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridServicerestartRequestChangedobject {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridServicerestartRequestChangedobjectModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridServicerestartRequestChangedobjectModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridServicerestartRequestChangedobject {
	if m == nil {
		return nil
	}
	to := &grid.GridServicerestartRequestChangedobject{
		Ref: flex.ExpandStringPointer(m.Ref),
	}
	return to
}

func FlattenGridServicerestartRequestChangedobject(ctx context.Context, from *grid.GridServicerestartRequestChangedobject, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridServicerestartRequestChangedobjectAttrTypes)
	}
	m := GridServicerestartRequestChangedobjectModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridServicerestartRequestChangedobjectAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridServicerestartRequestChangedobjectModel) Flatten(ctx context.Context, from *grid.GridServicerestartRequestChangedobject, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridServicerestartRequestChangedobjectModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Action = flex.FlattenStringPointer(from.Action)
	m.ChangedProperties = flex.FlattenFrameworkListString(ctx, from.ChangedProperties, diags)
	m.ChangedTime = flex.FlattenInt64Pointer(from.ChangedTime)
	m.ObjectName = flex.FlattenStringPointer(from.ObjectName)
	m.ObjectType = flex.FlattenStringPointer(from.ObjectType)
	m.UserName = flex.FlattenStringPointer(from.UserName)
}
