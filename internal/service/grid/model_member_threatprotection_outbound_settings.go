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

type MemberThreatprotectionOutboundSettingsModel struct {
	EnableQueryFqdn types.Bool  `tfsdk:"enable_query_fqdn"`
	QueryFqdnLimit  types.Int64 `tfsdk:"query_fqdn_limit"`
}

var MemberThreatprotectionOutboundSettingsAttrTypes = map[string]attr.Type{
	"enable_query_fqdn": types.BoolType,
	"query_fqdn_limit":  types.Int64Type,
}

var MemberThreatprotectionOutboundSettingsResourceSchemaAttributes = map[string]schema.Attribute{
	"enable_query_fqdn": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Flag to enable using DNS query FQDN for Outbound.",
	},
	"query_fqdn_limit": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Max domain level for DNS Query FQDN",
	},
}

func ExpandMemberThreatprotectionOutboundSettings(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberThreatprotectionOutboundSettings {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberThreatprotectionOutboundSettingsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberThreatprotectionOutboundSettingsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberThreatprotectionOutboundSettings {
	if m == nil {
		return nil
	}
	to := &grid.MemberThreatprotectionOutboundSettings{
		EnableQueryFqdn: flex.ExpandBoolPointer(m.EnableQueryFqdn),
		QueryFqdnLimit:  flex.ExpandInt64Pointer(m.QueryFqdnLimit),
	}
	return to
}

func FlattenMemberThreatprotectionOutboundSettings(ctx context.Context, from *grid.MemberThreatprotectionOutboundSettings, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberThreatprotectionOutboundSettingsAttrTypes)
	}
	m := MemberThreatprotectionOutboundSettingsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberThreatprotectionOutboundSettingsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberThreatprotectionOutboundSettingsModel) Flatten(ctx context.Context, from *grid.MemberThreatprotectionOutboundSettings, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberThreatprotectionOutboundSettingsModel{}
	}
	m.EnableQueryFqdn = types.BoolPointerValue(from.EnableQueryFqdn)
	m.QueryFqdnLimit = flex.FlattenInt64Pointer(from.QueryFqdnLimit)
}
