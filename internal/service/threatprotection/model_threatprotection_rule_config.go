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

type ThreatprotectionRuleConfigModel struct {
	Action      types.String `tfsdk:"action"`
	LogSeverity types.String `tfsdk:"log_severity"`
	Params      types.List   `tfsdk:"params"`
}

var ThreatprotectionRuleConfigAttrTypes = map[string]attr.Type{
	"action":       types.StringType,
	"log_severity": types.StringType,
	"params":       types.ListType{ElemType: types.ObjectType{AttrTypes: ThreatprotectionruleconfigParamsAttrTypes}},
}

var ThreatprotectionRuleConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The rule action.",
	},
	"log_severity": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The rule log severity.",
	},
	"params": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ThreatprotectionruleconfigParamsResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The threat protection rule parameters.",
	},
}

func ExpandThreatprotectionRuleConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *threatprotection.ThreatprotectionRuleConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ThreatprotectionRuleConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ThreatprotectionRuleConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *threatprotection.ThreatprotectionRuleConfig {
	if m == nil {
		return nil
	}
	to := &threatprotection.ThreatprotectionRuleConfig{
		Action:      flex.ExpandStringPointer(m.Action),
		LogSeverity: flex.ExpandStringPointer(m.LogSeverity),
		Params:      flex.ExpandFrameworkListNestedBlock(ctx, m.Params, diags, ExpandThreatprotectionruleconfigParams),
	}
	return to
}

func FlattenThreatprotectionRuleConfig(ctx context.Context, from *threatprotection.ThreatprotectionRuleConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ThreatprotectionRuleConfigAttrTypes)
	}
	m := ThreatprotectionRuleConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ThreatprotectionRuleConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ThreatprotectionRuleConfigModel) Flatten(ctx context.Context, from *threatprotection.ThreatprotectionRuleConfig, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ThreatprotectionRuleConfigModel{}
	}
	m.Action = flex.FlattenStringPointer(from.Action)
	m.LogSeverity = flex.FlattenStringPointer(from.LogSeverity)
	m.Params = flex.FlattenFrameworkListNestedBlock(ctx, from.Params, ThreatprotectionruleconfigParamsAttrTypes, diags, FlattenThreatprotectionruleconfigParams)
}
