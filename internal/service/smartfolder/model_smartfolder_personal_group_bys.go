package smartfolder

import (
	"context"
	"regexp"

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
)

type SmartfolderPersonalGroupBysModel struct {
	Value          types.String `tfsdk:"value"`
	ValueType      types.String `tfsdk:"value_type"`
	EnableGrouping types.Bool   `tfsdk:"enable_grouping"`
}

var SmartfolderPersonalGroupBysAttrTypes = map[string]attr.Type{
	"value":           types.StringType,
	"value_type":      types.StringType,
	"enable_grouping": types.BoolType,
}

var SmartfolderPersonalGroupBysResourceSchemaAttributes = map[string]schema.Attribute{
	"value": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Should not have leading or trailing white space",
			),
		},
		MarkdownDescription: "The name of the Smart Folder grouping attribute.",
	},
	"value_type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("EXTATTR", "NORMAL"),
		},
		Default:             stringdefault.StaticString("NORMAL"),
		MarkdownDescription: "The type of the Smart Folder grouping attribute value.",
	},
	"enable_grouping": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether the grouping is enabled.",
	},
}

func ExpandSmartfolderPersonalGroupBys(ctx context.Context, o types.Object, diags *diag.Diagnostics) *smartfolder.SmartfolderPersonalGroupBys {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m SmartfolderPersonalGroupBysModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *SmartfolderPersonalGroupBysModel) Expand(ctx context.Context, diags *diag.Diagnostics) *smartfolder.SmartfolderPersonalGroupBys {
	if m == nil {
		return nil
	}
	to := &smartfolder.SmartfolderPersonalGroupBys{
		Value:          flex.ExpandStringPointer(m.Value),
		ValueType:      flex.ExpandStringPointer(m.ValueType),
		EnableGrouping: flex.ExpandBoolPointer(m.EnableGrouping),
	}
	return to
}

func FlattenSmartfolderPersonalGroupBys(ctx context.Context, from *smartfolder.SmartfolderPersonalGroupBys, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SmartfolderPersonalGroupBysAttrTypes)
	}
	m := SmartfolderPersonalGroupBysModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, SmartfolderPersonalGroupBysAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *SmartfolderPersonalGroupBysModel) Flatten(ctx context.Context, from *smartfolder.SmartfolderPersonalGroupBys, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = SmartfolderPersonalGroupBysModel{}
	}
	m.Value = flex.FlattenStringPointer(from.Value)
	m.ValueType = flex.FlattenStringPointer(from.ValueType)
	m.EnableGrouping = types.BoolPointerValue(from.EnableGrouping)
}
