package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
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
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Action to take on descendants of the object when the object is deleted.",
	},
	"flags": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "This field contains extensible attribute flags. Possible values: (A)udited, (C)loud API, Cloud (G)master, (I)nheritable, (L)isted, (M)andatory value, MGM (P)rivate, (R)ead Only, (S)ort enum values, Multiple (V)alues If there are two or more flags in the field, you must list them according to the order they are listed above. For example, 'CR' is a valid value for the 'flags' field because C = Cloud API is listed before R = Read only. However, the value 'RC' is invalid because the order for the 'flags' field is broken.",
	},
	"list_values": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ExtensibleattributedefListValuesResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
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
		Required:            true,
		MarkdownDescription: "Type for the Extensible Attribute Definition.",
	},
}

func (m *ExtensibleattributedefModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *grid.Extensibleattributedef {
	if m == nil {
		return nil
	}
	to := &grid.Extensibleattributedef{
		AllowedObjectTypes: flex.ExpandFrameworkListString(ctx, m.AllowedObjectTypes, diags),
		Comment:            flex.ExpandStringPointer(m.Comment),
		DefaultValue:       flex.ExpandStringPointer(m.DefaultValue),
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
	m.DefaultValue = flex.FlattenStringPointer(from.DefaultValue)
	m.DescendantsAction = FlattenExtensibleattributedefDescendantsAction(ctx, from.DescendantsAction, diags)
	m.Flags = flex.FlattenStringPointer(from.Flags)
	m.ListValues = flex.FlattenFrameworkListNestedBlock(ctx, from.ListValues, ExtensibleattributedefListValuesAttrTypes, diags, FlattenExtensibleattributedefListValues)
	m.Max = flex.FlattenInt64Pointer(from.Max)
	m.Min = flex.FlattenInt64Pointer(from.Min)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Namespace = flex.FlattenStringPointer(from.Namespace)
	m.Type = flex.FlattenStringPointer(from.Type)
}
