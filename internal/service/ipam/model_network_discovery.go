package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/ipam"

	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/flex"
)

type NetworkDiscoveryModel struct {
	Ref types.String `tfsdk:"ref"`
}

var NetworkDiscoveryAttrTypes = map[string]attr.Type{
	"_ref": types.StringType,
}

var NetworkDiscoveryResourceSchemaAttributes = map[string]schema.Attribute{
	"_ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
}

func ExpandNetworkDiscovery(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.NetworkDiscovery {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkDiscoveryModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NetworkDiscoveryModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.NetworkDiscovery {
	if m == nil {
		return nil
	}
	to := &ipam.NetworkDiscovery{
		Ref: flex.ExpandStringPointer(m.Ref),
	}
	return to
}

func FlattenNetworkDiscovery(ctx context.Context, from *ipam.NetworkDiscovery, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkDiscoveryAttrTypes)
	}
	m := NetworkDiscoveryModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkDiscoveryAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NetworkDiscoveryModel) Flatten(ctx context.Context, from *ipam.NetworkDiscovery, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NetworkDiscoveryModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
}
