package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dtc"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type DtcTopologyRulesInnerOneOf1SourcesInnerModel struct {
	SourceOp    types.String `tfsdk:"source_op"`
	SourceType  types.String `tfsdk:"source_type"`
	SourceValue types.String `tfsdk:"source_value"`
}

var DtcTopologyRulesInnerOneOf1SourcesInnerAttrTypes = map[string]attr.Type{
	"source_op":    types.StringType,
	"source_type":  types.StringType,
	"source_value": types.StringType,
}

var DtcTopologyRulesInnerOneOf1SourcesInnerResourceSchemaAttributes = map[string]schema.Attribute{
	"source_op": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("IS", "IS_NOT"),
		},
		MarkdownDescription: "Operation for matching the source.",
	},
	"source_type": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.OneOf("CITY", "CONTINENT", "COUNTRY", "EA0", "EA1", "EA2", "EA3", "SUBDIVISION", "SUBNET"),
		},
		MarkdownDescription: "Type of the source.",
	},
	"source_value": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Value of the source.",
	},
}

func ExpandDtcTopologyRulesInnerOneOf1SourcesInner(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dtc.DtcTopologyRulesInnerOneOf1SourcesInner {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DtcTopologyRulesInnerOneOf1SourcesInnerModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DtcTopologyRulesInnerOneOf1SourcesInnerModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dtc.DtcTopologyRulesInnerOneOf1SourcesInner {
	if m == nil {
		return nil
	}
	to := &dtc.DtcTopologyRulesInnerOneOf1SourcesInner{
		SourceOp:    flex.ExpandStringPointer(m.SourceOp),
		SourceType:  flex.ExpandStringPointer(m.SourceType),
		SourceValue: flex.ExpandStringPointer(m.SourceValue),
	}
	return to
}

func FlattenDtcTopologyRulesInnerOneOf1SourcesInner(ctx context.Context, from *dtc.DtcTopologyRulesInnerOneOf1SourcesInner, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DtcTopologyRulesInnerOneOf1SourcesInnerAttrTypes)
	}
	m := DtcTopologyRulesInnerOneOf1SourcesInnerModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DtcTopologyRulesInnerOneOf1SourcesInnerAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DtcTopologyRulesInnerOneOf1SourcesInnerModel) Flatten(ctx context.Context, from *dtc.DtcTopologyRulesInnerOneOf1SourcesInner, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DtcTopologyRulesInnerOneOf1SourcesInnerModel{}
	}
	m.SourceOp = flex.FlattenStringPointer(from.SourceOp)
	m.SourceType = flex.FlattenStringPointer(from.SourceType)
	m.SourceValue = flex.FlattenStringPointer(from.SourceValue)
}
