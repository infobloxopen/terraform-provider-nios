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

type ThreatprotectionRulesetModel struct {
	Ref                   types.String `tfsdk:"ref"`
	Uuid                  types.String `tfsdk:"uuid"`
	AddType               types.String `tfsdk:"add_type"`
	AddedTime             types.Int64  `tfsdk:"added_time"`
	Comment               types.String `tfsdk:"comment"`
	DoNotDelete           types.Bool   `tfsdk:"do_not_delete"`
	IsFactoryResetEnabled types.Bool   `tfsdk:"is_factory_reset_enabled"`
	UsedBy                types.List   `tfsdk:"used_by"`
	Version               types.String `tfsdk:"version"`
}

var ThreatprotectionRulesetAttrTypes = map[string]attr.Type{
	"ref":                      types.StringType,
	"uuid":                     types.StringType,
	"add_type":                 types.StringType,
	"added_time":               types.Int64Type,
	"comment":                  types.StringType,
	"do_not_delete":            types.BoolType,
	"is_factory_reset_enabled": types.BoolType,
	"used_by":                  types.ListType{ElemType: types.StringType},
	"version":                  types.StringType,
}

var ThreatprotectionRulesetResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The uuid to the object.",
	},
	"add_type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Determines the way the ruleset was added.",
	},
	"added_time": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The time when the ruleset was added.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The human readable comment for the ruleset.",
	},
	"do_not_delete": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the ruleset will not be deleted during upgrade.",
	},
	"is_factory_reset_enabled": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines if factory reset is enabled for this ruleset.",
	},
	"used_by": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Computed:            true,
		MarkdownDescription: "The users of the ruleset.",
	},
	"version": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The ruleset version.",
	},
}

func ExpandThreatprotectionRuleset(ctx context.Context, o types.Object, diags *diag.Diagnostics) *threatprotection.ThreatprotectionRuleset {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ThreatprotectionRulesetModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ThreatprotectionRulesetModel) Expand(ctx context.Context, diags *diag.Diagnostics) *threatprotection.ThreatprotectionRuleset {
	if m == nil {
		return nil
	}
	to := &threatprotection.ThreatprotectionRuleset{
		Ref:         flex.ExpandStringPointer(m.Ref),
		Comment:     flex.ExpandStringPointer(m.Comment),
		DoNotDelete: flex.ExpandBoolPointer(m.DoNotDelete),
		Uuid:        flex.ExpandStringPointer(m.Uuid),
	}
	return to
}

func FlattenThreatprotectionRuleset(ctx context.Context, from *threatprotection.ThreatprotectionRuleset, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ThreatprotectionRulesetAttrTypes)
	}
	m := ThreatprotectionRulesetModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ThreatprotectionRulesetAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ThreatprotectionRulesetModel) Flatten(ctx context.Context, from *threatprotection.ThreatprotectionRuleset, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ThreatprotectionRulesetModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.AddType = flex.FlattenStringPointer(from.AddType)
	m.AddedTime = flex.FlattenInt64Pointer(from.AddedTime)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.DoNotDelete = types.BoolPointerValue(from.DoNotDelete)
	m.IsFactoryResetEnabled = types.BoolPointerValue(from.IsFactoryResetEnabled)
	m.UsedBy = flex.FlattenFrameworkListString(ctx, from.UsedBy, diags)
	m.Version = flex.FlattenStringPointer(from.Version)
}
