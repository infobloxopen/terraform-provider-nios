package grid

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"
	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	planmodifiers "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/immutable"
)

type ExtensibleattributedefModel struct {
	Ref                types.String `tfsdk:"ref"`
	AllowedObjectTypes types.List   `tfsdk:"allowed_object_types"`
	Comment            types.String `tfsdk:"comment"`
	DefaultValue       types.String `tfsdk:"default_value"`
	DescendantsAction  types.Object `tfsdk:"descendants_action"`
	Flags              types.String `tfsdk:"flags"`
	ListValues         types.List   `tfsdk:"list_values"`
	Max                types.Int64  `tfsdk:"max"`
	Min                types.Int64  `tfsdk:"min"`
	Name               types.String `tfsdk:"name"`
	Namespace          types.String `tfsdk:"namespace"`
	Type               types.String `tfsdk:"type"`
}

var ExtensibleattributedefAttrTypes = map[string]attr.Type{
	"ref":                  types.StringType,
	"allowed_object_types": types.ListType{ElemType: types.StringType},
	"comment":              types.StringType,
	"default_value":        types.StringType,
	"descendants_action":   types.ObjectType{AttrTypes: ExtensibleattributedefDescendantsActionAttrTypes},
	"flags":                types.StringType,
	"list_values":          types.ListType{ElemType: types.ObjectType{AttrTypes: ExtensibleattributedefListValuesAttrTypes}},
	"max":                  types.Int64Type,
	"min":                  types.Int64Type,
	"name":                 types.StringType,
	"namespace":            types.StringType,
	"type":                 types.StringType,
}

var ExtensibleattributedefResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"allowed_object_types": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
		Default:     listdefault.StaticValue(types.ListNull(types.StringType)),
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The object types this extensible attribute is allowed to associate with.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "Comment for the Extensible Attribute Definition; maximum 256 characters.",
	},
	"default_value": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Default value used to pre-populate the attribute value in the GUI. For email, URL, and string types, the value is a string with a maximum of 256 characters. For an integer, the value is an integer from -2147483648 through 2147483647. For a date, the value is the number of seconds that have elapsed since January 1st, 1970 UTC.",
	},
	"descendants_action": schema.SingleNestedAttribute{
		Attributes:          ExtensibleattributedefDescendantsActionResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "This option describes the action that must be taken on the extensible attribute by its descendant in case the ‘Inheritable’ flag is set.",
	},
	"flags": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "This field contains extensible attribute flags. Possible values: (A)udited, (C)loud API, Cloud (G)master, (I)nheritable, (L)isted, (M)andatory value, MGM (P)rivate, (R)ead Only, (S)ort enum values, Multiple (V)alues If there are two or more flags in the field, you must list them according to the order they are listed above. For example, 'CR' is a valid value for the 'flags' field because C = Cloud API is listed before R = Read only. However, the value 'RC' is invalid because the order for the 'flags' field is broken.",
	},
	"list_values": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ExtensibleattributedefListValuesResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Default:             listdefault.StaticValue(types.ListNull(types.ObjectType{AttrTypes: ExtensibleattributedefListValuesAttrTypes})),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "List of Values. Applicable if the extensible attribute type is ENUM.",
	},
	"max": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Maximum allowed value of extensible attribute. Applicable if the extensible attribute type is INTEGER.",
	},
	"min": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Minimum allowed value of extensible attribute. Applicable if the extensible attribute type is INTEGER.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of the Extensible Attribute Definition.",
	},
	"namespace": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Namespace for the Extensible Attribute Definition.",
	},
	"type": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.OneOf("DATE", "EMAIL", "ENUM", "INTEGER", "STRING", "URL"),
		},
		MarkdownDescription: "Type for the Extensible Attribute Definition.",
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
	},
}

func (m *ExtensibleattributedefModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *grid.Extensibleattributedef {
	if m == nil {
		return nil
	}
	to := &grid.Extensibleattributedef{
		AllowedObjectTypes: flex.ExpandFrameworkListString(ctx, m.AllowedObjectTypes, diags),
		Comment:            flex.ExpandStringPointer(m.Comment),
		DefaultValue:       ExpandExtensibleattributedefDefaultValue(ctx, m.DefaultValue, m.Type, diags),
		DescendantsAction:  ExpandExtensibleattributedefDescendantsAction(ctx, m.DescendantsAction, diags),
		Flags:              flex.ExpandStringPointer(m.Flags),
		ListValues:         flex.ExpandFrameworkListNestedBlock(ctx, m.ListValues, diags, ExpandExtensibleattributedefListValues),
		Max:                flex.ExpandInt64Pointer(m.Max),
		Min:                flex.ExpandInt64Pointer(m.Min),
		Name:               flex.ExpandStringPointer(m.Name),
	}
	if isCreate {
		to.Type = flex.ExpandStringPointer(m.Type)
	}
	return to
}

func FlattenExtensibleattributedef(ctx context.Context, from *grid.Extensibleattributedef, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ExtensibleattributedefAttrTypes)
	}
	m := ExtensibleattributedefModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ExtensibleattributedefAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ExtensibleattributedefModel) Flatten(ctx context.Context, from *grid.Extensibleattributedef, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ExtensibleattributedefModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.AllowedObjectTypes = flex.FlattenFrameworkListString(ctx, from.AllowedObjectTypes, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.DefaultValue = FlattenExtensibleattributedefDefaultValue(ctx, from.DefaultValue, diags)
	m.DescendantsAction = FlattenExtensibleattributedefDescendantsAction(ctx, from.DescendantsAction, diags)
	m.Flags = flex.FlattenStringPointer(from.Flags)
	m.ListValues = flex.FlattenFrameworkListNestedBlock(ctx, from.ListValues, ExtensibleattributedefListValuesAttrTypes, diags, FlattenExtensibleattributedefListValues)
	m.Max = flex.FlattenInt64Pointer(from.Max)
	m.Min = flex.FlattenInt64Pointer(from.Min)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Namespace = flex.FlattenStringPointer(from.Namespace)
	m.Type = flex.FlattenStringPointer(from.Type)
}

func ExpandExtensibleattributedefDefaultValue(ctx context.Context, defaultValue types.String, eaType types.String, diags *diag.Diagnostics) *grid.ExtensibleattributedefDefaultValue {
	if defaultValue.IsNull() || defaultValue.IsUnknown() {
		return &grid.ExtensibleattributedefDefaultValue{}
	}

	value := defaultValue.ValueString()
	if value == "" {
		return &grid.ExtensibleattributedefDefaultValue{}
	}

	// Check the type to determine if we should send as integer or string
	if !eaType.IsNull() && !eaType.IsUnknown() && eaType.ValueString() == "INTEGER" {
		// Convert string to integer for INTEGER type
		if intVal, err := strconv.ParseInt(value, 10, 32); err == nil {
			int32Val := int32(intVal)
			return &grid.ExtensibleattributedefDefaultValue{
				Int32: &int32Val,
			}
		} else {
			diags.AddError(
				"Invalid Integer Default Value",
				fmt.Sprintf("Cannot convert default_value '%s' to integer: %v", value, err),
			)
			return &grid.ExtensibleattributedefDefaultValue{}
		}
	}

	// For all other types (STRING, EMAIL, URL, DATE, ENUM), send as string
	return &grid.ExtensibleattributedefDefaultValue{
		String: &value,
	}
}

func FlattenExtensibleattributedefDefaultValue(ctx context.Context, from *grid.ExtensibleattributedefDefaultValue, diags *diag.Diagnostics) types.String {
	if from == nil {
		return types.StringNull()
	}

	if from.Int32 != nil {
		// Convert int32 to string for Terraform
		return types.StringValue(strconv.FormatInt(int64(*from.Int32), 10))
	}

	// Check if string value is set
	if from.String != nil {
		return types.StringValue(*from.String)
	}

	// No value set
	return types.StringNull()
}
