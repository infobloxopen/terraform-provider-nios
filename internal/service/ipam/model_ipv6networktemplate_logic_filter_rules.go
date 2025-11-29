package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type Ipv6networktemplateLogicFilterRulesModel struct {
	Filter types.String `tfsdk:"filter"`
	Type   types.String `tfsdk:"type"`
}

var Ipv6networktemplateLogicFilterRulesAttrTypes = map[string]attr.Type{
	"filter": types.StringType,
	"type":   types.StringType,
}

var Ipv6networktemplateLogicFilterRulesResourceSchemaAttributes = map[string]schema.Attribute{
	"filter": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The filter name.",
	},
	"type": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The filter type. Valid values are: * MAC * NAC * Option",
	},
}

func ExpandIpv6networktemplateLogicFilterRules(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.Ipv6networktemplateLogicFilterRules {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6networktemplateLogicFilterRulesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *Ipv6networktemplateLogicFilterRulesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.Ipv6networktemplateLogicFilterRules {
	if m == nil {
		return nil
	}
	to := &ipam.Ipv6networktemplateLogicFilterRules{
		Filter: flex.ExpandStringPointer(m.Filter),
		Type:   flex.ExpandStringPointer(m.Type),
	}
	return to
}

func FlattenIpv6networktemplateLogicFilterRules(ctx context.Context, from *ipam.Ipv6networktemplateLogicFilterRules, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6networktemplateLogicFilterRulesAttrTypes)
	}
	m := Ipv6networktemplateLogicFilterRulesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6networktemplateLogicFilterRulesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Ipv6networktemplateLogicFilterRulesModel) Flatten(ctx context.Context, from *ipam.Ipv6networktemplateLogicFilterRules, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Ipv6networktemplateLogicFilterRulesModel{}
	}
	m.Filter = flex.FlattenStringPointer(from.Filter)
	m.Type = flex.FlattenStringPointer(from.Type)
}
