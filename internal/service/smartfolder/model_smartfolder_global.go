package smartfolder

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/smartfolder"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type SmartfolderGlobalModel struct {
	Ref        types.String `tfsdk:"ref"`
	Comment    types.String `tfsdk:"comment"`
	GroupBys   types.List   `tfsdk:"group_bys"`
	Name       types.String `tfsdk:"name"`
	QueryItems types.List   `tfsdk:"query_items"`
}

var SmartfolderGlobalAttrTypes = map[string]attr.Type{
	"ref":         types.StringType,
	"comment":     types.StringType,
	"group_bys":   types.ListType{ElemType: types.ObjectType{AttrTypes: SmartfolderGlobalGroupBysAttrTypes}},
	"name":        types.StringType,
	"query_items": types.ListType{ElemType: types.ObjectType{AttrTypes: SmartfolderGlobalQueryItemsAttrTypes}},
}

var SmartfolderGlobalResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The global Smart Folder descriptive comment.",
	},
	"group_bys": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: SmartfolderGlobalGroupBysResourceSchemaAttributes,
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Global Smart Folder grouping rules.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The global Smart Folder name.",
	},
	"query_items": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: SmartfolderGlobalQueryItemsResourceSchemaAttributes,
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The global Smart Folder filter queries.",
	},
}

// func ExpandSmartfolderGlobal(ctx context.Context, o types.Object, diags *diag.Diagnostics) *smartfolder.SmartfolderGlobal {
// 	if o.IsNull() || o.IsUnknown() {
// 		return nil
// 	}
// 	var m SmartfolderGlobalModel
// 	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
// 	if diags.HasError() {
// 		return nil
// 	}
// 	return m.Expand(ctx, diags)
// }

func (m *SmartfolderGlobalModel) Expand(ctx context.Context, diags *diag.Diagnostics) *smartfolder.SmartfolderGlobal {
	if m == nil {
		return nil
	}
	to := &smartfolder.SmartfolderGlobal{
		Comment:    flex.ExpandStringPointer(m.Comment),
		GroupBys:   flex.ExpandFrameworkListNestedBlock(ctx, m.GroupBys, diags, ExpandSmartfolderGlobalGroupBys),
		Name:       flex.ExpandStringPointer(m.Name),
		QueryItems: flex.ExpandFrameworkListNestedBlock(ctx, m.QueryItems, diags, ExpandSmartfolderGlobalQueryItems),
	}
	return to
}

func FlattenSmartfolderGlobal(ctx context.Context, from *smartfolder.SmartfolderGlobal, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SmartfolderGlobalAttrTypes)
	}
	m := SmartfolderGlobalModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, SmartfolderGlobalAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *SmartfolderGlobalModel) Flatten(ctx context.Context, from *smartfolder.SmartfolderGlobal, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = SmartfolderGlobalModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.GroupBys = flex.FlattenFrameworkListNestedBlock(ctx, from.GroupBys, SmartfolderGlobalGroupBysAttrTypes, diags, FlattenSmartfolderGlobalGroupBys)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.QueryItems = flex.FlattenFrameworkListNestedBlock(ctx, from.QueryItems, SmartfolderGlobalQueryItemsAttrTypes, diags, FlattenSmartfolderGlobalQueryItems)
}
