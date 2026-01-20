package threatprotection

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/threatprotection"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type ThreatprotectionGridRuleModel struct {
	Ref                   types.String `tfsdk:"ref"`
	Uuid                  types.String `tfsdk:"uuid"`
	AllowedActions        types.List   `tfsdk:"allowed_actions"`
	Category              types.String `tfsdk:"category"`
	Comment               types.String `tfsdk:"comment"`
	Config                types.Object `tfsdk:"config"`
	Description           types.String `tfsdk:"description"`
	Disabled              types.Bool   `tfsdk:"disabled"`
	IsFactoryResetEnabled types.Bool   `tfsdk:"is_factory_reset_enabled"`
	Name                  types.String `tfsdk:"name"`
	Ruleset               types.String `tfsdk:"ruleset"`
	Sid                   types.Int64  `tfsdk:"sid"`
	Template              types.String `tfsdk:"template"`
	Type                  types.String `tfsdk:"type"`
}

var ThreatprotectionGridRuleAttrTypes = map[string]attr.Type{
	"ref":                      types.StringType,
	"uuid":                     types.StringType,
	"allowed_actions":          types.ListType{ElemType: types.StringType},
	"category":                 types.StringType,
	"comment":                  types.StringType,
	"config":                   types.ObjectType{AttrTypes: ThreatprotectionGridRuleConfigAttrTypes},
	"description":              types.StringType,
	"disabled":                 types.BoolType,
	"is_factory_reset_enabled": types.BoolType,
	"name":                     types.StringType,
	"ruleset":                  types.StringType,
	"sid":                      types.Int64Type,
	"template":                 types.StringType,
	"type":                     types.StringType,
}

var ThreatprotectionGridRuleResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The uuid to the object.",
	},
	"allowed_actions": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Computed:            true,
		MarkdownDescription: "The list of allowed actions of the custom rule.",
	},
	"category": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The rule category the custom rule assigned to.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The human readable comment for the custom rule.",
	},
	"config": schema.SingleNestedAttribute{
		Attributes: ThreatprotectionGridRuleConfigResourceSchemaAttributes,
		Optional:   true,
	},
	"description": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The description of the custom rule.",
	},
	"disabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the custom rule is disabled.",
	},
	"is_factory_reset_enabled": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines if factory reset is enabled for the custom rule.",
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the rule custom rule concatenated with its rule config parameters.",
	},
	"ruleset": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The version of the ruleset the custom rule assigned to.",
	},
	"sid": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The Rule ID.",
	},
	"template": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The threat protection rule template used to create this rule.",
	},
	"type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The type of the custom rule.",
	},
}

func ExpandThreatprotectionGridRule(ctx context.Context, o types.Object, diags *diag.Diagnostics) *threatprotection.ThreatprotectionGridRule {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ThreatprotectionGridRuleModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ThreatprotectionGridRuleModel) Expand(ctx context.Context, diags *diag.Diagnostics) *threatprotection.ThreatprotectionGridRule {
	if m == nil {
		return nil
	}
	to := &threatprotection.ThreatprotectionGridRule{
		Ref:      flex.ExpandStringPointer(m.Ref),
		Comment:  flex.ExpandStringPointer(m.Comment),
		Config:   ExpandThreatprotectionGridRuleConfig(ctx, m.Config, diags),
		Disabled: flex.ExpandBoolPointer(m.Disabled),
		Template: flex.ExpandStringPointer(m.Template),
	}
	return to
}

func FlattenThreatprotectionGridRule(ctx context.Context, from *threatprotection.ThreatprotectionGridRule, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ThreatprotectionGridRuleAttrTypes)
	}
	m := ThreatprotectionGridRuleModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ThreatprotectionGridRuleAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ThreatprotectionGridRuleModel) Flatten(ctx context.Context, from *threatprotection.ThreatprotectionGridRule, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ThreatprotectionGridRuleModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.AllowedActions = flex.FlattenFrameworkListString(ctx, from.AllowedActions, diags)
	m.Category = flex.FlattenStringPointer(from.Category)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Config = FlattenThreatprotectionGridRuleConfig(ctx, from.Config, diags)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Disabled = types.BoolPointerValue(from.Disabled)
	m.IsFactoryResetEnabled = types.BoolPointerValue(from.IsFactoryResetEnabled)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Ruleset = flex.FlattenStringPointer(from.Ruleset)
	m.Sid = flex.FlattenInt64Pointer(from.Sid)
	m.Template = flex.FlattenStringPointer(from.Template)
	m.Type = flex.FlattenStringPointer(from.Type)
}
