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

type IPv6SharednetworkNetworksModel struct {
	Ref types.String `tfsdk:"ref"`
}

var IPv6SharednetworkNetworksAttrTypes = map[string]attr.Type{
	"ref": types.StringType,
}

var IPv6SharednetworkNetworksResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Reference to the IPv6 Network.",
	},
}

func ExpandIPv6SharednetworkNetworks(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dhcp.IPv6SharednetworkNetworks {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IPv6SharednetworkNetworksModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IPv6SharednetworkNetworksModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.IPv6SharednetworkNetworks {
	if m == nil {
		return nil
	}
	to := &dhcp.IPv6SharednetworkNetworks{
		Ref: flex.ExpandStringPointer(m.Ref),
	}
	return to
}

func FlattenIPv6SharednetworkNetworks(ctx context.Context, from *dhcp.IPv6SharednetworkNetworks, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IPv6SharednetworkNetworksAttrTypes)
	}
	m := IPv6SharednetworkNetworksModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IPv6SharednetworkNetworksAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IPv6SharednetworkNetworksModel) Flatten(ctx context.Context, from *dhcp.IPv6SharednetworkNetworks, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IPv6SharednetworkNetworksModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
}
