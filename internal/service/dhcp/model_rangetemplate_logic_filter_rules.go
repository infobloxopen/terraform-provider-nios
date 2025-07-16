package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type RangetemplateLogicFilterRulesModel struct {
	Filter types.String `tfsdk:"filter"`
	Type   types.String `tfsdk:"type"`
}

var RangetemplateLogicFilterRulesAttrTypes = map[string]attr.Type{
	"filter": types.StringType,
	"type":   types.StringType,
}

var RangetemplateLogicFilterRulesResourceSchemaAttributes = map[string]schema.Attribute{
	"filter": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The filter name.",
	},
	"type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The filter type. Valid values are: * MAC * NAC * Option",
	},
}

func ExpandRangetemplateLogicFilterRules(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dhcp.RangetemplateLogicFilterRules {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RangetemplateLogicFilterRulesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *RangetemplateLogicFilterRulesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.RangetemplateLogicFilterRules {
	if m == nil {
		return nil
	}
	to := &dhcp.RangetemplateLogicFilterRules{
		Filter: flex.ExpandStringPointer(m.Filter),
		Type:   flex.ExpandStringPointer(m.Type),
	}
	return to
}

func FlattenRangetemplateLogicFilterRules(ctx context.Context, from *dhcp.RangetemplateLogicFilterRules, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RangetemplateLogicFilterRulesAttrTypes)
	}
	m := RangetemplateLogicFilterRulesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RangetemplateLogicFilterRulesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RangetemplateLogicFilterRulesModel) Flatten(ctx context.Context, from *dhcp.RangetemplateLogicFilterRules, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RangetemplateLogicFilterRulesModel{}
	}
	m.Filter = flex.FlattenStringPointer(from.Filter)
	m.Type = flex.FlattenStringPointer(from.Type)
}
