package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type ZoneAuthNetworkAssociationsModel struct {
	Ref         types.String `tfsdk:"ref"`
	Network     types.String `tfsdk:"network"`
	NetworkView types.String `tfsdk:"network_view"`
}

var ZoneAuthNetworkAssociationsAttrTypes = map[string]attr.Type{
	"ref":          types.StringType,
	"network":      types.StringType,
	"network_view": types.StringType,
}

var ZoneAuthNetworkAssociationsResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the network or network container object.",
	},
	"network": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The network address in CIDR notation.",
	},
	"network_view": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the network view.",
	},
}

func ExpandZoneAuthNetworkAssociations(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.ZoneAuthNetworkAssociations {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneAuthNetworkAssociationsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ZoneAuthNetworkAssociationsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.ZoneAuthNetworkAssociations {
	if m == nil {
		return nil
	}
	to := &dns.ZoneAuthNetworkAssociations{}
	return to
}

func FlattenZoneAuthNetworkAssociations(ctx context.Context, from *dns.ZoneAuthNetworkAssociations, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneAuthNetworkAssociationsAttrTypes)
	}
	m := ZoneAuthNetworkAssociationsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZoneAuthNetworkAssociationsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ZoneAuthNetworkAssociationsModel) Flatten(ctx context.Context, from *dns.ZoneAuthNetworkAssociations, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ZoneAuthNetworkAssociationsModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Network = flex.FlattenStringPointer(from.Network)
	m.NetworkView = flex.FlattenStringPointer(from.NetworkView)
}
