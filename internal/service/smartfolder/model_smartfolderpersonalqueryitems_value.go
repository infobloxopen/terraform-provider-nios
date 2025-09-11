package smartfolder

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/smartfolder"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type SmartfolderpersonalqueryitemsValueModel struct {
	ValueInteger types.Int64  `tfsdk:"value_integer"`
	ValueString  types.String `tfsdk:"value_string"`
	ValueDate    types.Int64  `tfsdk:"value_date"`
	ValueBoolean types.Bool   `tfsdk:"value_boolean"`
}

var SmartfolderpersonalqueryitemsValueAttrTypes = map[string]attr.Type{
	"value_integer": types.Int64Type,
	"value_string":  types.StringType,
	"value_date":    types.Int64Type,
	"value_boolean": types.BoolType,
}

var SmartfolderpersonalqueryitemsValueResourceSchemaAttributes = map[string]schema.Attribute{
	"value_integer": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The integer value of the Smart Folder query.",
	},
	"value_string": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The string value of the Smart Folder query.",
	},
	"value_date": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The timestamp value of the Smart Folder query.",
	},
	"value_boolean": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "The boolean value of the Smart Folder query.",
	},
}

func ExpandSmartfolderpersonalqueryitemsValue(ctx context.Context, o types.Object, diags *diag.Diagnostics) *smartfolder.SmartfolderpersonalqueryitemsValue {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m SmartfolderpersonalqueryitemsValueModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *SmartfolderpersonalqueryitemsValueModel) Expand(ctx context.Context, diags *diag.Diagnostics) *smartfolder.SmartfolderpersonalqueryitemsValue {
	if m == nil {
		return nil
	}
	to := &smartfolder.SmartfolderpersonalqueryitemsValue{
		ValueInteger: flex.ExpandInt64Pointer(m.ValueInteger),
		ValueString:  flex.ExpandStringPointer(m.ValueString),
		ValueDate:    flex.ExpandInt64Pointer(m.ValueDate),
		ValueBoolean: flex.ExpandBoolPointer(m.ValueBoolean),
	}
	return to
}

func FlattenSmartfolderpersonalqueryitemsValue(ctx context.Context, from *smartfolder.SmartfolderpersonalqueryitemsValue, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SmartfolderpersonalqueryitemsValueAttrTypes)
	}
	m := SmartfolderpersonalqueryitemsValueModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, SmartfolderpersonalqueryitemsValueAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *SmartfolderpersonalqueryitemsValueModel) Flatten(ctx context.Context, from *smartfolder.SmartfolderpersonalqueryitemsValue, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = SmartfolderpersonalqueryitemsValueModel{}
	}
	m.ValueInteger = flex.FlattenInt64Pointer(from.ValueInteger)
	m.ValueString = flex.FlattenStringPointer(from.ValueString)
	m.ValueDate = flex.FlattenInt64Pointer(from.ValueDate)
	m.ValueBoolean = types.BoolPointerValue(from.ValueBoolean)
}
