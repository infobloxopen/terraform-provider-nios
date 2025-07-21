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

type RangetemplateRelayAgentFilterRulesModel struct {
	Filter     types.String `tfsdk:"filter"`
	Permission types.String `tfsdk:"permission"`
}

var RangetemplateRelayAgentFilterRulesAttrTypes = map[string]attr.Type{
	"filter":     types.StringType,
	"permission": types.StringType,
}

var RangetemplateRelayAgentFilterRulesResourceSchemaAttributes = map[string]schema.Attribute{
	"filter": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The name of the DHCP filter.",
	},
	"permission": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The permission to be applied.",
	},
}

func ExpandRangetemplateRelayAgentFilterRules(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dhcp.RangetemplateRelayAgentFilterRules {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RangetemplateRelayAgentFilterRulesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *RangetemplateRelayAgentFilterRulesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.RangetemplateRelayAgentFilterRules {
	if m == nil {
		return nil
	}
	to := &dhcp.RangetemplateRelayAgentFilterRules{
		Filter:     flex.ExpandStringPointer(m.Filter),
		Permission: flex.ExpandStringPointer(m.Permission),
	}
	return to
}

func FlattenRangetemplateRelayAgentFilterRules(ctx context.Context, from *dhcp.RangetemplateRelayAgentFilterRules, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RangetemplateRelayAgentFilterRulesAttrTypes)
	}
	m := RangetemplateRelayAgentFilterRulesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RangetemplateRelayAgentFilterRulesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RangetemplateRelayAgentFilterRulesModel) Flatten(ctx context.Context, from *dhcp.RangetemplateRelayAgentFilterRules, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RangetemplateRelayAgentFilterRulesModel{}
	}
	m.Filter = flex.FlattenStringPointer(from.Filter)
	m.Permission = flex.FlattenStringPointer(from.Permission)
}
