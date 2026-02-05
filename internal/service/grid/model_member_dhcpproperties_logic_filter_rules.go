package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type MemberDhcppropertiesLogicFilterRulesModel struct {
	Filter types.String `tfsdk:"filter"`
	Type   types.String `tfsdk:"type"`
}

var MemberDhcppropertiesLogicFilterRulesAttrTypes = map[string]attr.Type{
	"filter": types.StringType,
	"type":   types.StringType,
}

var MemberDhcppropertiesLogicFilterRulesResourceSchemaAttributes = map[string]schema.Attribute{
	"filter": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The filter name.",
	},
	"type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The filter type. Valid values are: * MAC * NAC * Option",
	},
}

func ExpandMemberDhcppropertiesLogicFilterRules(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberDhcppropertiesLogicFilterRules {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberDhcppropertiesLogicFilterRulesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberDhcppropertiesLogicFilterRulesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberDhcppropertiesLogicFilterRules {
	if m == nil {
		return nil
	}
	to := &grid.MemberDhcppropertiesLogicFilterRules{
		Filter: flex.ExpandStringPointer(m.Filter),
		Type:   flex.ExpandStringPointer(m.Type),
	}
	return to
}

func FlattenMemberDhcppropertiesLogicFilterRules(ctx context.Context, from *grid.MemberDhcppropertiesLogicFilterRules, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberDhcppropertiesLogicFilterRulesAttrTypes)
	}
	m := MemberDhcppropertiesLogicFilterRulesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberDhcppropertiesLogicFilterRulesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberDhcppropertiesLogicFilterRulesModel) Flatten(ctx context.Context, from *grid.MemberDhcppropertiesLogicFilterRules, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberDhcppropertiesLogicFilterRulesModel{}
	}
	m.Filter = flex.FlattenStringPointer(from.Filter)
	m.Type = flex.FlattenStringPointer(from.Type)
}
