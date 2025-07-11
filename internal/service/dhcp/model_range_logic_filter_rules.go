package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dhcp"

	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/flex"
)

type RangeLogicFilterRulesModel struct {
	Filter types.String `tfsdk:"filter"`
	Type   types.String `tfsdk:"type"`
}

var RangeLogicFilterRulesAttrTypes = map[string]attr.Type{
	"filter": types.StringType,
	"type":   types.StringType,
}

var RangeLogicFilterRulesResourceSchemaAttributes = map[string]schema.Attribute{
	"filter": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The filter name.",
	},
	"type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The filter type. Valid values are: * MAC * NAC * Option",
	},
}

func ExpandRangeLogicFilterRules(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dhcp.RangeLogicFilterRules {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RangeLogicFilterRulesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *RangeLogicFilterRulesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.RangeLogicFilterRules {
	if m == nil {
		return nil
	}
	to := &dhcp.RangeLogicFilterRules{
		Filter: flex.ExpandStringPointer(m.Filter),
		Type:   flex.ExpandStringPointer(m.Type),
	}
	return to
}

func FlattenRangeLogicFilterRules(ctx context.Context, from *dhcp.RangeLogicFilterRules, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RangeLogicFilterRulesAttrTypes)
	}
	m := RangeLogicFilterRulesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RangeLogicFilterRulesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RangeLogicFilterRulesModel) Flatten(ctx context.Context, from *dhcp.RangeLogicFilterRules, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RangeLogicFilterRulesModel{}
	}
	m.Filter = flex.FlattenStringPointer(from.Filter)
	m.Type = flex.FlattenStringPointer(from.Type)
}
