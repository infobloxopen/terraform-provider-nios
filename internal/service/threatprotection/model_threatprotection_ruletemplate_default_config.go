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

type ThreatprotectionRuletemplateDefaultConfigModel struct {
	Action      types.String `tfsdk:"action"`
	LogSeverity types.String `tfsdk:"log_severity"`
	Params      types.List   `tfsdk:"params"`
}

var ThreatprotectionRuletemplateDefaultConfigAttrTypes = map[string]attr.Type{
	"action":       types.StringType,
	"log_severity": types.StringType,
	"params":       types.ListType{ElemType: types.ObjectType{AttrTypes: ThreatprotectionruletemplatedefaultconfigParamsAttrTypes}},
}

var ThreatprotectionRuletemplateDefaultConfigResourceSchemaAttributes = map[string]schema.Attribute{
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
			Attributes: ThreatprotectionruletemplatedefaultconfigParamsResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The threat protection rule parameters.",
	},
}

func ExpandThreatprotectionRuletemplateDefaultConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *threatprotection.ThreatprotectionRuletemplateDefaultConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ThreatprotectionRuletemplateDefaultConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ThreatprotectionRuletemplateDefaultConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *threatprotection.ThreatprotectionRuletemplateDefaultConfig {
	if m == nil {
		return nil
	}
	to := &threatprotection.ThreatprotectionRuletemplateDefaultConfig{
		Action:      flex.ExpandStringPointer(m.Action),
		LogSeverity: flex.ExpandStringPointer(m.LogSeverity),
		Params:      flex.ExpandFrameworkListNestedBlock(ctx, m.Params, diags, ExpandThreatprotectionruletemplatedefaultconfigParams),
	}
	return to
}

func FlattenThreatprotectionRuletemplateDefaultConfig(ctx context.Context, from *threatprotection.ThreatprotectionRuletemplateDefaultConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ThreatprotectionRuletemplateDefaultConfigAttrTypes)
	}
	m := ThreatprotectionRuletemplateDefaultConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ThreatprotectionRuletemplateDefaultConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ThreatprotectionRuletemplateDefaultConfigModel) Flatten(ctx context.Context, from *threatprotection.ThreatprotectionRuletemplateDefaultConfig, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ThreatprotectionRuletemplateDefaultConfigModel{}
	}
	m.Action = flex.FlattenStringPointer(from.Action)
	m.LogSeverity = flex.FlattenStringPointer(from.LogSeverity)
	m.Params = flex.FlattenFrameworkListNestedBlock(ctx, from.Params, ThreatprotectionruletemplatedefaultconfigParamsAttrTypes, diags, FlattenThreatprotectionruletemplatedefaultconfigParams)
}
