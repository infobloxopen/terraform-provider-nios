package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type NetworkviewDdnsZonePrimariesModel struct {
	ZoneMatch      types.String `tfsdk:"zone_match"`
	DnsGridZone    types.Object `tfsdk:"dns_grid_zone"`
	DnsGridPrimary types.String `tfsdk:"dns_grid_primary"`
	DnsExtZone     types.String `tfsdk:"dns_ext_zone"`
	DnsExtPrimary  types.String `tfsdk:"dns_ext_primary"`
}

var NetworkviewDdnsZonePrimariesAttrTypes = map[string]attr.Type{
	"zone_match":       types.StringType,
	"dns_grid_zone":    types.ObjectType{AttrTypes: NetworkviewDdnsZonePrimariesDnsGridZoneAttrTypes},
	"dns_grid_primary": types.StringType,
	"dns_ext_zone":     types.StringType,
	"dns_ext_primary":  types.StringType,
}

var NetworkviewDdnsZonePrimariesResourceSchemaAttributes = map[string]schema.Attribute{
	"zone_match": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.OneOf("ANY_EXTERNAL", "ANY_GRID", "EXTERNAL", "GRID"),
		},
		MarkdownDescription: "Indicate matching type.",
	},
	"dns_grid_zone": schema.SingleNestedAttribute{
		Attributes:          NetworkviewDdnsZonePrimariesDnsGridZoneResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "The ref of a DNS zone.",
	},
	"dns_grid_primary": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The name of a Grid member.",
	},
	"dns_ext_zone": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.IsValidFQDN(),
		},
		MarkdownDescription: "The name of external zone in FQDN format.",
	},
	"dns_ext_primary": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The IP address of the External server. Valid when zone_match is \"EXTERNAL\" or \"ANY_EXTERNAL\".",
	},
}

func ExpandNetworkviewDdnsZonePrimaries(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.NetworkviewDdnsZonePrimaries {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkviewDdnsZonePrimariesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NetworkviewDdnsZonePrimariesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.NetworkviewDdnsZonePrimaries {
	if m == nil {
		return nil
	}
	to := &ipam.NetworkviewDdnsZonePrimaries{
		ZoneMatch:      flex.ExpandStringPointer(m.ZoneMatch),
		DnsGridZone:    ExpandNetworkviewDdnsZonePrimariesDnsGridZone(ctx, m.DnsGridZone, diags),
		DnsGridPrimary: flex.ExpandStringPointer(m.DnsGridPrimary),
		DnsExtZone:     flex.ExpandStringPointer(m.DnsExtZone),
		DnsExtPrimary:  flex.ExpandStringPointer(m.DnsExtPrimary),
	}
	return to
}

func FlattenNetworkviewDdnsZonePrimaries(ctx context.Context, from *ipam.NetworkviewDdnsZonePrimaries, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkviewDdnsZonePrimariesAttrTypes)
	}
	m := NetworkviewDdnsZonePrimariesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkviewDdnsZonePrimariesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NetworkviewDdnsZonePrimariesModel) Flatten(ctx context.Context, from *ipam.NetworkviewDdnsZonePrimaries, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NetworkviewDdnsZonePrimariesModel{}
	}
	m.ZoneMatch = flex.FlattenStringPointer(from.ZoneMatch)
	m.DnsGridZone = FlattenNetworkviewDdnsZonePrimariesDnsGridZone(ctx, from.DnsGridZone, diags)
	m.DnsGridPrimary = flex.FlattenStringPointer(from.DnsGridPrimary)
	m.DnsExtZone = flex.FlattenStringPointer(from.DnsExtZone)
	m.DnsExtPrimary = flex.FlattenStringPointer(from.DnsExtPrimary)
}
