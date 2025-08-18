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

type SmartfolderglobalqueryitemsValueModel struct {
	ValueInteger types.Int64  `tfsdk:"value_integer"`
	ValueString  types.String `tfsdk:"value_string"`
	ValueDate    types.Int64  `tfsdk:"value_date"`
	ValueBoolean types.Bool   `tfsdk:"value_boolean"`
}

var SmartfolderglobalqueryitemsValueAttrTypes = map[string]attr.Type{
	"value_integer": types.Int64Type,
	"value_string":  types.StringType,
	"value_date":    types.Int64Type,
	"value_boolean": types.BoolType,
}

var SmartfolderglobalqueryitemsValueResourceSchemaAttributes = map[string]schema.Attribute{
	"value_integer": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The integer value of the Smart Folder query.",
	},
	"value_string": schema.StringAttribute{
		Optional: true,
		Computed: true,
		//Default:             stringdefault.StaticString("Network/Zone/Range/Member"),
		MarkdownDescription: "The string value of the Smart Folder query.",
	},
	"value_date": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The timestamp value of the Smart Folder query.",
	},
	"value_boolean": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The boolean value of the Smart Folder query.",
	},
}

func ExpandSmartfolderglobalqueryitemsValue(ctx context.Context, o types.Object, diags *diag.Diagnostics) *smartfolder.SmartfolderglobalqueryitemsValue {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m SmartfolderglobalqueryitemsValueModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *SmartfolderglobalqueryitemsValueModel) Expand(ctx context.Context, diags *diag.Diagnostics) *smartfolder.SmartfolderglobalqueryitemsValue {
	if m == nil {
		return nil
	}
	to := &smartfolder.SmartfolderglobalqueryitemsValue{
		ValueInteger: flex.ExpandInt64Pointer(m.ValueInteger),
		ValueString:  flex.ExpandStringPointer(m.ValueString),
		ValueDate:    flex.ExpandInt64Pointer(m.ValueDate),
		ValueBoolean: flex.ExpandBoolPointer(m.ValueBoolean),
	}
	return to
}

func FlattenSmartfolderglobalqueryitemsValue(ctx context.Context, from *smartfolder.SmartfolderglobalqueryitemsValue, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SmartfolderglobalqueryitemsValueAttrTypes)
	}
	m := SmartfolderglobalqueryitemsValueModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, SmartfolderglobalqueryitemsValueAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *SmartfolderglobalqueryitemsValueModel) Flatten(ctx context.Context, from *smartfolder.SmartfolderglobalqueryitemsValue, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = SmartfolderglobalqueryitemsValueModel{}
	}
	m.ValueInteger = flex.FlattenInt64Pointer(from.ValueInteger)
	m.ValueString = flex.FlattenStringPointer(from.ValueString)
	m.ValueDate = flex.FlattenInt64Pointer(from.ValueDate)
	m.ValueBoolean = types.BoolPointerValue(from.ValueBoolean)
}
