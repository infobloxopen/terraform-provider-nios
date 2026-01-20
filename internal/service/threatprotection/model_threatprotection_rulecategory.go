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

type ThreatprotectionRulecategoryModel struct {
	Ref                   types.String `tfsdk:"ref"`
	Uuid                  types.String `tfsdk:"uuid"`
	IsFactoryResetEnabled types.Bool   `tfsdk:"is_factory_reset_enabled"`
	Name                  types.String `tfsdk:"name"`
	Ruleset               types.String `tfsdk:"ruleset"`
}

var ThreatprotectionRulecategoryAttrTypes = map[string]attr.Type{
	"ref":                      types.StringType,
	"uuid":                     types.StringType,
	"is_factory_reset_enabled": types.BoolType,
	"name":                     types.StringType,
	"ruleset":                  types.StringType,
}

var ThreatprotectionRulecategoryResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The uuid to the object.",
	},
	"is_factory_reset_enabled": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines if factory reset is enabled for this rule category.",
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the rule category.",
	},
	"ruleset": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The version of the ruleset the category assigned to.",
	},
}

func ExpandThreatprotectionRulecategory(ctx context.Context, o types.Object, diags *diag.Diagnostics) *threatprotection.ThreatprotectionRulecategory {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ThreatprotectionRulecategoryModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ThreatprotectionRulecategoryModel) Expand(ctx context.Context, diags *diag.Diagnostics) *threatprotection.ThreatprotectionRulecategory {
	if m == nil {
		return nil
	}
	to := &threatprotection.ThreatprotectionRulecategory{
		Ref:  flex.ExpandStringPointer(m.Ref),
		Uuid: flex.ExpandStringPointer(m.Uuid),
	}
	return to
}

func FlattenThreatprotectionRulecategory(ctx context.Context, from *threatprotection.ThreatprotectionRulecategory, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ThreatprotectionRulecategoryAttrTypes)
	}
	m := ThreatprotectionRulecategoryModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ThreatprotectionRulecategoryAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ThreatprotectionRulecategoryModel) Flatten(ctx context.Context, from *threatprotection.ThreatprotectionRulecategory, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ThreatprotectionRulecategoryModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.IsFactoryResetEnabled = types.BoolPointerValue(from.IsFactoryResetEnabled)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Ruleset = flex.FlattenStringPointer(from.Ruleset)
}
