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

type RangetemplateOptionFilterRulesModel struct {
	Filter     types.String `tfsdk:"filter"`
	Permission types.String `tfsdk:"permission"`
}

var RangetemplateOptionFilterRulesAttrTypes = map[string]attr.Type{
	"filter":     types.StringType,
	"permission": types.StringType,
}

var RangetemplateOptionFilterRulesResourceSchemaAttributes = map[string]schema.Attribute{
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

func ExpandRangetemplateOptionFilterRules(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dhcp.RangetemplateOptionFilterRules {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RangetemplateOptionFilterRulesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *RangetemplateOptionFilterRulesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.RangetemplateOptionFilterRules {
	if m == nil {
		return nil
	}
	to := &dhcp.RangetemplateOptionFilterRules{
		Filter:     flex.ExpandStringPointer(m.Filter),
		Permission: flex.ExpandStringPointer(m.Permission),
	}
	return to
}

func FlattenRangetemplateOptionFilterRules(ctx context.Context, from *dhcp.RangetemplateOptionFilterRules, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RangetemplateOptionFilterRulesAttrTypes)
	}
	m := RangetemplateOptionFilterRulesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RangetemplateOptionFilterRulesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RangetemplateOptionFilterRulesModel) Flatten(ctx context.Context, from *dhcp.RangetemplateOptionFilterRules, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RangetemplateOptionFilterRulesModel{}
	}
	m.Filter = flex.FlattenStringPointer(from.Filter)
	m.Permission = flex.FlattenStringPointer(from.Permission)
}
