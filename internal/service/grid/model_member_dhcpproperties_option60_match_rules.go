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

type MemberDhcppropertiesOption60MatchRulesModel struct {
	MatchValue      types.String `tfsdk:"match_value"`
	OptionSpace     types.String `tfsdk:"option_space"`
	IsSubstring     types.Bool   `tfsdk:"is_substring"`
	SubstringOffset types.Int64  `tfsdk:"substring_offset"`
	SubstringLength types.Int64  `tfsdk:"substring_length"`
}

var MemberDhcppropertiesOption60MatchRulesAttrTypes = map[string]attr.Type{
	"match_value":      types.StringType,
	"option_space":     types.StringType,
	"is_substring":     types.BoolType,
	"substring_offset": types.Int64Type,
	"substring_length": types.Int64Type,
}

var MemberDhcppropertiesOption60MatchRulesResourceSchemaAttributes = map[string]schema.Attribute{
	"match_value": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The match value for this DHCP Option 60 match rule.",
	},
	"option_space": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The option space for this DHCP Option 60 match rule.",
	},
	"is_substring": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the match value is a substring.",
	},
	"substring_offset": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The offset of match value for this DHCP Option 60 match rule.",
	},
	"substring_length": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The length of match value for this DHCP Option 60 match rule.",
	},
}

func ExpandMemberDhcppropertiesOption60MatchRules(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberDhcppropertiesOption60MatchRules {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberDhcppropertiesOption60MatchRulesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberDhcppropertiesOption60MatchRulesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberDhcppropertiesOption60MatchRules {
	if m == nil {
		return nil
	}
	to := &grid.MemberDhcppropertiesOption60MatchRules{
		MatchValue:      flex.ExpandStringPointer(m.MatchValue),
		OptionSpace:     flex.ExpandStringPointer(m.OptionSpace),
		IsSubstring:     flex.ExpandBoolPointer(m.IsSubstring),
		SubstringOffset: flex.ExpandInt64Pointer(m.SubstringOffset),
		SubstringLength: flex.ExpandInt64Pointer(m.SubstringLength),
	}
	return to
}

func FlattenMemberDhcppropertiesOption60MatchRules(ctx context.Context, from *grid.MemberDhcppropertiesOption60MatchRules, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberDhcppropertiesOption60MatchRulesAttrTypes)
	}
	m := MemberDhcppropertiesOption60MatchRulesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberDhcppropertiesOption60MatchRulesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberDhcppropertiesOption60MatchRulesModel) Flatten(ctx context.Context, from *grid.MemberDhcppropertiesOption60MatchRules, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberDhcppropertiesOption60MatchRulesModel{}
	}
	m.MatchValue = flex.FlattenStringPointer(from.MatchValue)
	m.OptionSpace = flex.FlattenStringPointer(from.OptionSpace)
	m.IsSubstring = types.BoolPointerValue(from.IsSubstring)
	m.SubstringOffset = flex.FlattenInt64Pointer(from.SubstringOffset)
	m.SubstringLength = flex.FlattenInt64Pointer(from.SubstringLength)
}
