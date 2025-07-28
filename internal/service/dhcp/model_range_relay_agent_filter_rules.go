package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type RangeRelayAgentFilterRulesModel struct {
	Filter     types.String `tfsdk:"filter"`
	Permission types.String `tfsdk:"permission"`
}

var RangeRelayAgentFilterRulesAttrTypes = map[string]attr.Type{
	"filter":     types.StringType,
	"permission": types.StringType,
}

var RangeRelayAgentFilterRulesResourceSchemaAttributes = map[string]schema.Attribute{
	"filter": schema.StringAttribute{
		Required: 		   true,
		MarkdownDescription: "The name of the DHCP filter.",
	},
	"permission": schema.StringAttribute{
		Required:            true,
		Validators: []validator.String{
			stringvalidator.OneOf("Allow", "Deny"),
		},
		MarkdownDescription: "The permission to be applied.",
	},
}

func ExpandRangeRelayAgentFilterRules(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dhcp.RangeRelayAgentFilterRules {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RangeRelayAgentFilterRulesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *RangeRelayAgentFilterRulesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.RangeRelayAgentFilterRules {
	if m == nil {
		return nil
	}
	to := &dhcp.RangeRelayAgentFilterRules{
		Filter:     flex.ExpandStringPointer(m.Filter),
		Permission: flex.ExpandStringPointer(m.Permission),
	}
	return to
}

func FlattenRangeRelayAgentFilterRules(ctx context.Context, from *dhcp.RangeRelayAgentFilterRules, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RangeRelayAgentFilterRulesAttrTypes)
	}
	m := RangeRelayAgentFilterRulesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RangeRelayAgentFilterRulesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RangeRelayAgentFilterRulesModel) Flatten(ctx context.Context, from *dhcp.RangeRelayAgentFilterRules, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RangeRelayAgentFilterRulesModel{}
	}
	m.Filter = flex.FlattenStringPointer(from.Filter)
	m.Permission = flex.FlattenStringPointer(from.Permission)
}
