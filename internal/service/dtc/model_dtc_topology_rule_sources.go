package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dtc"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type DtcTopologyRuleSourcesModel struct {
	SourceType  types.String `tfsdk:"source_type"`
	SourceOp    types.String `tfsdk:"source_op"`
	SourceValue types.String `tfsdk:"source_value"`
}

var DtcTopologyRuleSourcesAttrTypes = map[string]attr.Type{
	"source_type":  types.StringType,
	"source_op":    types.StringType,
	"source_value": types.StringType,
}

var DtcTopologyRuleSourcesResourceSchemaAttributes = map[string]schema.Attribute{
	"source_type": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The source type.",
	},
	"source_op": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The operation used to match the value.",
	},
	"source_value": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The source value.",
	},
}

func ExpandDtcTopologyRuleSources(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dtc.DtcTopologyRuleSources {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DtcTopologyRuleSourcesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DtcTopologyRuleSourcesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dtc.DtcTopologyRuleSources {
	if m == nil {
		return nil
	}
	to := &dtc.DtcTopologyRuleSources{
		SourceType:  flex.ExpandStringPointer(m.SourceType),
		SourceOp:    flex.ExpandStringPointer(m.SourceOp),
		SourceValue: flex.ExpandStringPointer(m.SourceValue),
	}
	return to
}

func FlattenDtcTopologyRuleSources(ctx context.Context, from *dtc.DtcTopologyRuleSources, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DtcTopologyRuleSourcesAttrTypes)
	}
	m := DtcTopologyRuleSourcesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DtcTopologyRuleSourcesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DtcTopologyRuleSourcesModel) Flatten(ctx context.Context, from *dtc.DtcTopologyRuleSources, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DtcTopologyRuleSourcesModel{}
	}
	m.SourceType = flex.FlattenStringPointer(from.SourceType)
	m.SourceOp = flex.FlattenStringPointer(from.SourceOp)
	m.SourceValue = flex.FlattenStringPointer(from.SourceValue)
}
