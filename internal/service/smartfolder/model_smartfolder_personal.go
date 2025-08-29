package smartfolder

import (
	"context"
	"regexp"

	listvalidator "github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/smartfolder"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type SmartfolderPersonalModel struct {
	Ref        types.String `tfsdk:"ref"`
	Comment    types.String `tfsdk:"comment"`
	GroupBys   types.List   `tfsdk:"group_bys"`
	IsShortcut types.Bool   `tfsdk:"is_shortcut"`
	Name       types.String `tfsdk:"name"`
	QueryItems types.List   `tfsdk:"query_items"`
}

var SmartfolderPersonalAttrTypes = map[string]attr.Type{
	"ref":         types.StringType,
	"comment":     types.StringType,
	"group_bys":   types.ListType{ElemType: types.ObjectType{AttrTypes: SmartfolderPersonalGroupBysAttrTypes}},
	"is_shortcut": types.BoolType,
	"name":        types.StringType,
	"query_items": types.ListType{ElemType: types.ObjectType{AttrTypes: SmartfolderPersonalQueryItemsAttrTypes}},
}

var SmartfolderPersonalResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The personal Smart Folder descriptive comment.",
	},
	"group_bys": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: SmartfolderPersonalGroupBysResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The personal Smart Folder groupping rules.",
	},
	"is_shortcut": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines whether the personal Smart Folder is a shortcut.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Name should not have leading or trailing white space",
			),
		},
		MarkdownDescription: "The personal Smart Folder name.",
	},
	"query_items": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: SmartfolderPersonalQueryItemsResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The personal Smart Folder filter queries.",
	},
}

func (m *SmartfolderPersonalModel) Expand(ctx context.Context, diags *diag.Diagnostics) *smartfolder.SmartfolderPersonal {
	if m == nil {
		return nil
	}
	to := &smartfolder.SmartfolderPersonal{
		Comment:    flex.ExpandStringPointer(m.Comment),
		GroupBys:   flex.ExpandFrameworkListNestedBlock(ctx, m.GroupBys, diags, ExpandSmartfolderPersonalGroupBys),
		Name:       flex.ExpandStringPointer(m.Name),
		QueryItems: flex.ExpandFrameworkListNestedBlock(ctx, m.QueryItems, diags, ExpandSmartfolderPersonalQueryItems),
	}
	return to
}

func FlattenSmartfolderPersonal(ctx context.Context, from *smartfolder.SmartfolderPersonal, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SmartfolderPersonalAttrTypes)
	}
	m := SmartfolderPersonalModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, SmartfolderPersonalAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *SmartfolderPersonalModel) Flatten(ctx context.Context, from *smartfolder.SmartfolderPersonal, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = SmartfolderPersonalModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.GroupBys = flex.FlattenFrameworkListNestedBlock(ctx, from.GroupBys, SmartfolderPersonalGroupBysAttrTypes, diags, FlattenSmartfolderPersonalGroupBys)
	m.IsShortcut = types.BoolPointerValue(from.IsShortcut)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.QueryItems = flex.FlattenFrameworkListNestedBlock(ctx, from.QueryItems, SmartfolderPersonalQueryItemsAttrTypes, diags, FlattenSmartfolderPersonalQueryItems)
}
