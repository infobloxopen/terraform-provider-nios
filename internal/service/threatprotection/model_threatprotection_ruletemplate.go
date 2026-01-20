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

type ThreatprotectionRuletemplateModel struct {
	Ref            types.String `tfsdk:"ref"`
	Uuid           types.String `tfsdk:"uuid"`
	AllowedActions types.List   `tfsdk:"allowed_actions"`
	Category       types.String `tfsdk:"category"`
	DefaultConfig  types.Object `tfsdk:"default_config"`
	Description    types.String `tfsdk:"description"`
	Name           types.String `tfsdk:"name"`
	Ruleset        types.String `tfsdk:"ruleset"`
	Sid            types.Int64  `tfsdk:"sid"`
}

var ThreatprotectionRuletemplateAttrTypes = map[string]attr.Type{
	"ref":             types.StringType,
	"uuid":            types.StringType,
	"allowed_actions": types.ListType{ElemType: types.StringType},
	"category":        types.StringType,
	"default_config":  types.ObjectType{AttrTypes: ThreatprotectionRuletemplateDefaultConfigAttrTypes},
	"description":     types.StringType,
	"name":            types.StringType,
	"ruleset":         types.StringType,
	"sid":             types.Int64Type,
}

var ThreatprotectionRuletemplateResourceSchemaAttributes = map[string]schema.Attribute{
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
		MarkdownDescription: "The list of allowed actions of rhe rule template.",
	},
	"category": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The rule category this template assigned to.",
	},
	"default_config": schema.SingleNestedAttribute{
		Attributes: ThreatprotectionRuletemplateDefaultConfigResourceSchemaAttributes,
		Optional:   true,
	},
	"description": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The description of the rule template.",
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the rule template.",
	},
	"ruleset": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The version of the ruleset the template assigned to.",
	},
	"sid": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The Rule ID.",
	},
}

func ExpandThreatprotectionRuletemplate(ctx context.Context, o types.Object, diags *diag.Diagnostics) *threatprotection.ThreatprotectionRuletemplate {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ThreatprotectionRuletemplateModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ThreatprotectionRuletemplateModel) Expand(ctx context.Context, diags *diag.Diagnostics) *threatprotection.ThreatprotectionRuletemplate {
	if m == nil {
		return nil
	}
	to := &threatprotection.ThreatprotectionRuletemplate{
		Ref:           flex.ExpandStringPointer(m.Ref),
		Uuid:          flex.ExpandStringPointer(m.Uuid),
		DefaultConfig: ExpandThreatprotectionRuletemplateDefaultConfig(ctx, m.DefaultConfig, diags),
	}
	return to
}

func FlattenThreatprotectionRuletemplate(ctx context.Context, from *threatprotection.ThreatprotectionRuletemplate, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ThreatprotectionRuletemplateAttrTypes)
	}
	m := ThreatprotectionRuletemplateModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ThreatprotectionRuletemplateAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ThreatprotectionRuletemplateModel) Flatten(ctx context.Context, from *threatprotection.ThreatprotectionRuletemplate, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ThreatprotectionRuletemplateModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.AllowedActions = flex.FlattenFrameworkListString(ctx, from.AllowedActions, diags)
	m.Category = flex.FlattenStringPointer(from.Category)
	m.DefaultConfig = FlattenThreatprotectionRuletemplateDefaultConfig(ctx, from.DefaultConfig, diags)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Ruleset = flex.FlattenStringPointer(from.Ruleset)
	m.Sid = flex.FlattenInt64Pointer(from.Sid)
}
