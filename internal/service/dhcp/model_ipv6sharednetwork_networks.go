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

type Ipv6sharednetworkNetworksModel struct {
	Ref types.String `tfsdk:"ref"`
}

var Ipv6sharednetworkNetworksAttrTypes = map[string]attr.Type{
	"ref": types.StringType,
}

var Ipv6sharednetworkNetworksResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Reference to the IPv6 Network.",
	},
}

func ExpandIpv6sharednetworkNetworks(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dhcp.Ipv6sharednetworkNetworks {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6sharednetworkNetworksModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *Ipv6sharednetworkNetworksModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.Ipv6sharednetworkNetworks {
	if m == nil {
		return nil
	}
	to := &dhcp.Ipv6sharednetworkNetworks{
		Ref: flex.ExpandStringPointer(m.Ref),
	}
	return to
}

func FlattenIpv6sharednetworkNetworks(ctx context.Context, from *dhcp.Ipv6sharednetworkNetworks, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6sharednetworkNetworksAttrTypes)
	}
	m := Ipv6sharednetworkNetworksModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6sharednetworkNetworksAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Ipv6sharednetworkNetworksModel) Flatten(ctx context.Context, from *dhcp.Ipv6sharednetworkNetworks, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Ipv6sharednetworkNetworksModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
}
