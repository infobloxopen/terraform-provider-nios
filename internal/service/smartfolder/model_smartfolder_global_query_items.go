package smartfolder

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/smartfolder"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type SmartfolderGlobalQueryItemsModel struct {
	Name      types.String `tfsdk:"name"`
	FieldType types.String `tfsdk:"field_type"`
	Operator  types.String `tfsdk:"operator"`
	OpMatch   types.Bool   `tfsdk:"op_match"`
	ValueType types.String `tfsdk:"value_type"`
	Value     types.Object `tfsdk:"value"`
}

var SmartfolderGlobalQueryItemsAttrTypes = map[string]attr.Type{
	"name":       types.StringType,
	"field_type": types.StringType,
	"operator":   types.StringType,
	"op_match":   types.BoolType,
	"value_type": types.StringType,
	"value":      types.ObjectType{AttrTypes: SmartfolderglobalqueryitemsValueAttrTypes},
}

var SmartfolderGlobalQueryItemsResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("type"),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The Smart Folder query name.",
	},
	"field_type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("NORMAL"),
		Validators: []validator.String{
			stringvalidator.OneOf("EXTATTR", "NORMAL"),
		},
		MarkdownDescription: "The Smart Folder query field type.",
	},
	"operator": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("EQ"),
		Validators: []validator.String{
			stringvalidator.OneOf(
				"BEGINS_WITH",
				"CONTAINS",
				"DROPS_BY",
				"ENDS_WITH",
				"EQ",
				"GEQ",
				"GT",
				"HAS_VALUE",
				"INHERITANCE_STATE_EQUALS",
				"IP_ADDR_WITHIN",
				"LEQ",
				"LT",
				"MATCH_EXPR",
				"RELATIVE_DATE",
				"RISES_BY",
				"SUFFIX_MATCH",
			),
		},
		MarkdownDescription: "The Smart Folder operator used in query.",
	},
	"op_match": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Determines whether the query operator should match.",
	},
	"value_type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("ENUM"),
		Validators: []validator.String{
			stringvalidator.OneOf(
				"BOOLEAN",
				"DATE",
				"EMAIL",
				"ENUM",
				"INTEGER",
				"OBJTYPE",
				"STRING",
				"URL",
			),
		},
		MarkdownDescription: "The Smart Folder query value type.",
	},
	"value": schema.SingleNestedAttribute{
		Attributes: SmartfolderglobalqueryitemsValueResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
}

func ExpandSmartfolderGlobalQueryItems(ctx context.Context, o types.Object, diags *diag.Diagnostics) *smartfolder.SmartfolderGlobalQueryItems {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m SmartfolderGlobalQueryItemsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *SmartfolderGlobalQueryItemsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *smartfolder.SmartfolderGlobalQueryItems {
	if m == nil {
		return nil
	}
	to := &smartfolder.SmartfolderGlobalQueryItems{
		Name:      flex.ExpandStringPointer(m.Name),
		FieldType: flex.ExpandStringPointer(m.FieldType),
		Operator:  flex.ExpandStringPointer(m.Operator),
		OpMatch:   flex.ExpandBoolPointer(m.OpMatch),
		ValueType: flex.ExpandStringPointer(m.ValueType),
		Value:     ExpandSmartfolderglobalqueryitemsValue(ctx, m.Value, diags),
	}
	return to
}

func FlattenSmartfolderGlobalQueryItems(ctx context.Context, from *smartfolder.SmartfolderGlobalQueryItems, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SmartfolderGlobalQueryItemsAttrTypes)
	}
	m := SmartfolderGlobalQueryItemsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, SmartfolderGlobalQueryItemsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *SmartfolderGlobalQueryItemsModel) Flatten(ctx context.Context, from *smartfolder.SmartfolderGlobalQueryItems, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = SmartfolderGlobalQueryItemsModel{}
	}
	m.Name = flex.FlattenStringPointer(from.Name)
	m.FieldType = flex.FlattenStringPointer(from.FieldType)
	m.Operator = flex.FlattenStringPointer(from.Operator)
	m.OpMatch = types.BoolPointerValue(from.OpMatch)
	m.ValueType = flex.FlattenStringPointer(from.ValueType)
	m.Value = FlattenSmartfolderglobalqueryitemsValue(ctx, from.Value, diags)
}
