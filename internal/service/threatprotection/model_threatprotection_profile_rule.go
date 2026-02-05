package threatprotection

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/threatprotection"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type ThreatprotectionProfileRuleModel struct {
	Ref        types.String `tfsdk:"ref"`
	Config     types.Object `tfsdk:"config"`
	Disable    types.Bool   `tfsdk:"disable"`
	Profile    types.String `tfsdk:"profile"`
	Rule       types.String `tfsdk:"rule"`
	Sid        types.Int64  `tfsdk:"sid"`
	UseConfig  types.Bool   `tfsdk:"use_config"`
	UseDisable types.Bool   `tfsdk:"use_disable"`
}

var ThreatprotectionProfileRuleAttrTypes = map[string]attr.Type{
	"ref":         types.StringType,
	"config":      types.ObjectType{AttrTypes: ThreatprotectionProfileRuleConfigAttrTypes},
	"disable":     types.BoolType,
	"profile":     types.StringType,
	"rule":        types.StringType,
	"sid":         types.Int64Type,
	"use_config":  types.BoolType,
	"use_disable": types.BoolType,
}

var ThreatprotectionProfileRuleResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"config": schema.SingleNestedAttribute{
		Attributes: ThreatprotectionProfileRuleConfigResourceSchemaAttributes,
		Optional:   true,
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the rule is enabled or not for the profile.",
	},
	"profile": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the Threat protection profile.",
	},
	"rule": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The rule object name.",
	},
	"sid": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The snort rule ID.",
	},
	"use_config": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: config",
	},
	"use_disable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: disable",
	},
}

func ExpandThreatprotectionProfileRule(ctx context.Context, o types.Object, diags *diag.Diagnostics) *threatprotection.ThreatprotectionProfileRule {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ThreatprotectionProfileRuleModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ThreatprotectionProfileRuleModel) Expand(ctx context.Context, diags *diag.Diagnostics) *threatprotection.ThreatprotectionProfileRule {
	if m == nil {
		return nil
	}
	to := &threatprotection.ThreatprotectionProfileRule{
		Ref:        flex.ExpandStringPointer(m.Ref),
		Config:     ExpandThreatprotectionProfileRuleConfig(ctx, m.Config, diags),
		Disable:    flex.ExpandBoolPointer(m.Disable),
		UseConfig:  flex.ExpandBoolPointer(m.UseConfig),
		UseDisable: flex.ExpandBoolPointer(m.UseDisable),
	}
	return to
}

func FlattenThreatprotectionProfileRule(ctx context.Context, from *threatprotection.ThreatprotectionProfileRule, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ThreatprotectionProfileRuleAttrTypes)
	}
	m := ThreatprotectionProfileRuleModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ThreatprotectionProfileRuleAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ThreatprotectionProfileRuleModel) Flatten(ctx context.Context, from *threatprotection.ThreatprotectionProfileRule, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ThreatprotectionProfileRuleModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Config = FlattenThreatprotectionProfileRuleConfig(ctx, from.Config, diags)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.Profile = flex.FlattenStringPointer(from.Profile)
	m.Rule = flex.FlattenStringPointer(from.Rule)
	m.Sid = flex.FlattenInt64Pointer(from.Sid)
	m.UseConfig = types.BoolPointerValue(from.UseConfig)
	m.UseDisable = types.BoolPointerValue(from.UseDisable)
}
