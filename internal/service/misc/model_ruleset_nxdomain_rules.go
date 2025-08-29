package misc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/misc"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type RulesetNxdomainRulesModel struct {
	Action  types.String `tfsdk:"action"`
	Pattern types.String `tfsdk:"pattern"`
}

var RulesetNxdomainRulesAttrTypes = map[string]attr.Type{
	"action":  types.StringType,
	"pattern": types.StringType,
}

var RulesetNxdomainRulesResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("MODIFY", "PASS", "REDIRECT"),
		},
		Default:             stringdefault.StaticString("PASS"),
		MarkdownDescription: "The action to perform when a domain name matches the pattern defined in this Ruleset.",
	},
	"pattern": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The pattern that is used to match the domain name.",
	},
}

func ExpandRulesetNxdomainRules(ctx context.Context, o types.Object, diags *diag.Diagnostics) *misc.RulesetNxdomainRules {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RulesetNxdomainRulesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *RulesetNxdomainRulesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *misc.RulesetNxdomainRules {
	if m == nil {
		return nil
	}
	to := &misc.RulesetNxdomainRules{
		Action:  flex.ExpandStringPointer(m.Action),
		Pattern: flex.ExpandStringPointer(m.Pattern),
	}
	return to
}

func FlattenRulesetNxdomainRules(ctx context.Context, from *misc.RulesetNxdomainRules, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RulesetNxdomainRulesAttrTypes)
	}
	m := RulesetNxdomainRulesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RulesetNxdomainRulesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RulesetNxdomainRulesModel) Flatten(ctx context.Context, from *misc.RulesetNxdomainRules, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RulesetNxdomainRulesModel{}
	}
	m.Action = flex.FlattenStringPointer(from.Action)
	m.Pattern = flex.FlattenStringPointer(from.Pattern)
}
