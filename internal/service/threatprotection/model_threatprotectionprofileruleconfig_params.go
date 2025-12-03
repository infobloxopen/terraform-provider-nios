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

type ThreatprotectionprofileruleconfigParamsModel struct {
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Syntax      types.String `tfsdk:"syntax"`
	Value       types.String `tfsdk:"value"`
	Min         types.Int64  `tfsdk:"min"`
	Max         types.Int64  `tfsdk:"max"`
	ReadOnly    types.Bool   `tfsdk:"read_only"`
	EnumValues  types.List   `tfsdk:"enum_values"`
}

var ThreatprotectionprofileruleconfigParamsAttrTypes = map[string]attr.Type{
	"name":        types.StringType,
	"description": types.StringType,
	"syntax":      types.StringType,
	"value":       types.StringType,
	"min":         types.Int64Type,
	"max":         types.Int64Type,
	"read_only":   types.BoolType,
	"enum_values": types.ListType{ElemType: types.StringType},
}

var ThreatprotectionprofileruleconfigParamsResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The rule parameter name.",
	},
	"description": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The rule parameter description.",
	},
	"syntax": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The rule parameter syntax.",
	},
	"value": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The rule parameter value.",
	},
	"min": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The rule parameter minimum.",
	},
	"max": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The rule parameter maximum.",
	},
	"read_only": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines if parameter value is editable at member level.",
	},
	"enum_values": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Computed:            true,
		MarkdownDescription: "The rule parameter enum values.",
	},
}

func ExpandThreatprotectionprofileruleconfigParams(ctx context.Context, o types.Object, diags *diag.Diagnostics) *threatprotection.ThreatprotectionprofileruleconfigParams {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ThreatprotectionprofileruleconfigParamsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ThreatprotectionprofileruleconfigParamsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *threatprotection.ThreatprotectionprofileruleconfigParams {
	if m == nil {
		return nil
	}
	to := &threatprotection.ThreatprotectionprofileruleconfigParams{
		Name:  flex.ExpandStringPointer(m.Name),
		Value: flex.ExpandStringPointer(m.Value),
	}
	return to
}

func FlattenThreatprotectionprofileruleconfigParams(ctx context.Context, from *threatprotection.ThreatprotectionprofileruleconfigParams, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ThreatprotectionprofileruleconfigParamsAttrTypes)
	}
	m := ThreatprotectionprofileruleconfigParamsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ThreatprotectionprofileruleconfigParamsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ThreatprotectionprofileruleconfigParamsModel) Flatten(ctx context.Context, from *threatprotection.ThreatprotectionprofileruleconfigParams, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ThreatprotectionprofileruleconfigParamsModel{}
	}
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Syntax = flex.FlattenStringPointer(from.Syntax)
	m.Value = flex.FlattenStringPointer(from.Value)
	m.Min = flex.FlattenInt64Pointer(from.Min)
	m.Max = flex.FlattenInt64Pointer(from.Max)
	m.ReadOnly = types.BoolPointerValue(from.ReadOnly)
	m.EnumValues = flex.FlattenFrameworkListString(ctx, from.EnumValues, diags)
}
