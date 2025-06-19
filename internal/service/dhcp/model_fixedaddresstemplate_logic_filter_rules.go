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

type FixedaddresstemplateLogicFilterRulesModel struct {
	Filter types.String `tfsdk:"filter"`
	Type   types.String `tfsdk:"type"`
}

var FixedaddresstemplateLogicFilterRulesAttrTypes = map[string]attr.Type{
	"filter": types.StringType,
	"type":   types.StringType,
}

var FixedaddresstemplateLogicFilterRulesResourceSchemaAttributes = map[string]schema.Attribute{
	"filter": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The filter name.",
	},
	"type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The filter type. Valid values are: * MAC * NAC * Option",
	},
}

func ExpandFixedaddresstemplateLogicFilterRules(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dhcp.FixedaddresstemplateLogicFilterRules {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m FixedaddresstemplateLogicFilterRulesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *FixedaddresstemplateLogicFilterRulesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.FixedaddresstemplateLogicFilterRules {
	if m == nil {
		return nil
	}
	to := &dhcp.FixedaddresstemplateLogicFilterRules{
		Filter: flex.ExpandStringPointer(m.Filter),
		Type:   flex.ExpandStringPointer(m.Type),
	}
	return to
}

func FlattenFixedaddresstemplateLogicFilterRules(ctx context.Context, from *dhcp.FixedaddresstemplateLogicFilterRules, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(FixedaddresstemplateLogicFilterRulesAttrTypes)
	}
	m := FixedaddresstemplateLogicFilterRulesModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrs = m.ExtAttrsAll
	t, d := types.ObjectValueFrom(ctx, FixedaddresstemplateLogicFilterRulesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *FixedaddresstemplateLogicFilterRulesModel) Flatten(ctx context.Context, from *dhcp.FixedaddresstemplateLogicFilterRules, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = FixedaddresstemplateLogicFilterRulesModel{}
	}
	m.Filter = flex.FlattenStringPointer(from.Filter)
	m.Type = flex.FlattenStringPointer(from.Type)
}
