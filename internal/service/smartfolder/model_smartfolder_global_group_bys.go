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

type SmartfolderGlobalGroupBysModel struct {
	Value          types.String `tfsdk:"value"`
	ValueType      types.String `tfsdk:"value_type"`
	EnableGrouping types.Bool   `tfsdk:"enable_grouping"`
}

var SmartfolderGlobalGroupBysAttrTypes = map[string]attr.Type{
	"value":           types.StringType,
	"value_type":      types.StringType,
	"enable_grouping": types.BoolType,
}

var SmartfolderGlobalGroupBysResourceSchemaAttributes = map[string]schema.Attribute{
	"value": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The name of the Smart Folder grouping attribute.",
	},
	"value_type": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The type of the Smart Folder grouping attribute value.",
	},
	"enable_grouping": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Determines whether the grouping is enabled.",
	},
}

func ExpandSmartfolderGlobalGroupBys(ctx context.Context, o types.Object, diags *diag.Diagnostics) *smartfolder.SmartfolderGlobalGroupBys {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m SmartfolderGlobalGroupBysModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *SmartfolderGlobalGroupBysModel) Expand(ctx context.Context, diags *diag.Diagnostics) *smartfolder.SmartfolderGlobalGroupBys {
	if m == nil {
		return nil
	}
	to := &smartfolder.SmartfolderGlobalGroupBys{
		Value:          flex.ExpandStringPointer(m.Value),
		ValueType:      flex.ExpandStringPointer(m.ValueType),
		EnableGrouping: flex.ExpandBoolPointer(m.EnableGrouping),
	}
	return to
}

func FlattenSmartfolderGlobalGroupBys(ctx context.Context, from *smartfolder.SmartfolderGlobalGroupBys, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SmartfolderGlobalGroupBysAttrTypes)
	}
	m := SmartfolderGlobalGroupBysModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, SmartfolderGlobalGroupBysAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *SmartfolderGlobalGroupBysModel) Flatten(ctx context.Context, from *smartfolder.SmartfolderGlobalGroupBys, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = SmartfolderGlobalGroupBysModel{}
	}
	m.Value = flex.FlattenStringPointer(from.Value)
	m.ValueType = flex.FlattenStringPointer(from.ValueType)
	m.EnableGrouping = types.BoolPointerValue(from.EnableGrouping)
}
