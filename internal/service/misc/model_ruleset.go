package misc

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
	"github.com/infobloxopen/infoblox-nios-go-client/misc"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type RulesetModel struct {
	Ref           types.String `tfsdk:"ref"`
	Comment       types.String `tfsdk:"comment"`
	Disabled      types.Bool   `tfsdk:"disabled"`
	Name          types.String `tfsdk:"name"`
	NxdomainRules types.List   `tfsdk:"nxdomain_rules"`
	Type          types.String `tfsdk:"type"`
}

var RulesetAttrTypes = map[string]attr.Type{
	"ref":            types.StringType,
	"comment":        types.StringType,
	"disabled":       types.BoolType,
	"name":           types.StringType,
	"nxdomain_rules": types.ListType{ElemType: types.ObjectType{AttrTypes: RulesetNxdomainRulesAttrTypes}},
	"type":           types.StringType,
}

var RulesetResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Should not have leading or trailing whitespace",
			),
		},
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "Descriptive comment about the Ruleset object.",
	},
	"disabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The flag that indicates if the Ruleset object is disabled.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of this Ruleset object.",
	},
	"nxdomain_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RulesetNxdomainRulesResourceSchemaAttributes,
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The list of Rules assigned to this Ruleset object. Rules can be set only when the Ruleset type is set to \"NXDOMAIN\".",
	},
	"type": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.OneOf("BLACKLIST", "NXDOMAIN"),
		},
		MarkdownDescription: "The type of this Ruleset object.",
	},
}

func ExpandRuleset(ctx context.Context, o types.Object, diags *diag.Diagnostics) *misc.Ruleset {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RulesetModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *RulesetModel) Expand(ctx context.Context, diags *diag.Diagnostics) *misc.Ruleset {
	if m == nil {
		return nil
	}
	to := &misc.Ruleset{
		Comment:       flex.ExpandStringPointer(m.Comment),
		Disabled:      flex.ExpandBoolPointer(m.Disabled),
		Name:          flex.ExpandStringPointer(m.Name),
		NxdomainRules: flex.ExpandFrameworkListNestedBlock(ctx, m.NxdomainRules, diags, ExpandRulesetNxdomainRules),
		Type:          flex.ExpandStringPointer(m.Type),
	}
	return to
}

func FlattenRuleset(ctx context.Context, from *misc.Ruleset, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RulesetAttrTypes)
	}
	m := RulesetModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RulesetAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RulesetModel) Flatten(ctx context.Context, from *misc.Ruleset, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RulesetModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disabled = types.BoolPointerValue(from.Disabled)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.NxdomainRules = flex.FlattenFrameworkListNestedBlock(ctx, from.NxdomainRules, RulesetNxdomainRulesAttrTypes, diags, FlattenRulesetNxdomainRules)
	m.Type = flex.FlattenStringPointer(from.Type)
}
