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

type MemberDhcppropertiesDdnsZonePrimariesModel struct {
	ZoneMatch      types.String `tfsdk:"zone_match"`
	DnsGridZone    types.String `tfsdk:"dns_grid_zone"`
	DnsGridPrimary types.String `tfsdk:"dns_grid_primary"`
	DnsExtZone     types.String `tfsdk:"dns_ext_zone"`
	DnsExtPrimary  types.String `tfsdk:"dns_ext_primary"`
}

var MemberDhcppropertiesDdnsZonePrimariesAttrTypes = map[string]attr.Type{
	"zone_match":       types.StringType,
	"dns_grid_zone":    types.StringType,
	"dns_grid_primary": types.StringType,
	"dns_ext_zone":     types.StringType,
	"dns_ext_primary":  types.StringType,
}

var MemberDhcppropertiesDdnsZonePrimariesResourceSchemaAttributes = map[string]schema.Attribute{
	"zone_match": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Indicate matching type.",
	},
	"dns_grid_zone": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The ref of a DNS zone.",
	},
	"dns_grid_primary": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of a Grid member.",
	},
	"dns_ext_zone": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of external zone in FQDN format.",
	},
	"dns_ext_primary": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The IP address of the External server. Valid when zone_match is \"EXTERNAL\" or \"ANY_EXTERNAL\".",
	},
}

func ExpandMemberDhcppropertiesDdnsZonePrimaries(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberDhcppropertiesDdnsZonePrimaries {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberDhcppropertiesDdnsZonePrimariesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberDhcppropertiesDdnsZonePrimariesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberDhcppropertiesDdnsZonePrimaries {
	if m == nil {
		return nil
	}
	to := &grid.MemberDhcppropertiesDdnsZonePrimaries{
		ZoneMatch:      flex.ExpandStringPointer(m.ZoneMatch),
		DnsGridZone:    flex.ExpandStringPointer(m.DnsGridZone),
		DnsGridPrimary: flex.ExpandStringPointer(m.DnsGridPrimary),
		DnsExtZone:     flex.ExpandStringPointer(m.DnsExtZone),
		DnsExtPrimary:  flex.ExpandStringPointer(m.DnsExtPrimary),
	}
	return to
}

func FlattenMemberDhcppropertiesDdnsZonePrimaries(ctx context.Context, from *grid.MemberDhcppropertiesDdnsZonePrimaries, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberDhcppropertiesDdnsZonePrimariesAttrTypes)
	}
	m := MemberDhcppropertiesDdnsZonePrimariesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberDhcppropertiesDdnsZonePrimariesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberDhcppropertiesDdnsZonePrimariesModel) Flatten(ctx context.Context, from *grid.MemberDhcppropertiesDdnsZonePrimaries, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberDhcppropertiesDdnsZonePrimariesModel{}
	}
	m.ZoneMatch = flex.FlattenStringPointer(from.ZoneMatch)
	m.DnsGridZone = flex.FlattenStringPointer(from.DnsGridZone)
	m.DnsGridPrimary = flex.FlattenStringPointer(from.DnsGridPrimary)
	m.DnsExtZone = flex.FlattenStringPointer(from.DnsExtZone)
	m.DnsExtPrimary = flex.FlattenStringPointer(from.DnsExtPrimary)
}
