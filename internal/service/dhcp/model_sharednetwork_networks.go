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

type SharednetworkNetworksModel struct {
	Ref types.String `tfsdk:"ref"`
}

var SharednetworkNetworksAttrTypes = map[string]attr.Type{
	"ref": types.StringType,
}

var SharednetworkNetworksResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Reference to the Network.",
	},
}

func ExpandSharednetworkNetworks(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dhcp.SharednetworkNetworks {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m SharednetworkNetworksModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *SharednetworkNetworksModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.SharednetworkNetworks {
	if m == nil {
		return nil
	}
	to := &dhcp.SharednetworkNetworks{
		Ref: flex.ExpandStringPointer(m.Ref),
	}
	return to
}

func FlattenSharednetworkNetworks(ctx context.Context, from *dhcp.SharednetworkNetworks, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SharednetworkNetworksAttrTypes)
	}
	m := SharednetworkNetworksModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, SharednetworkNetworksAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *SharednetworkNetworksModel) Flatten(ctx context.Context, from *dhcp.SharednetworkNetworks, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = SharednetworkNetworksModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
}
