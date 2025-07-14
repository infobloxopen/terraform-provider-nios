package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type NetworkFederatedRealmsModel struct {
	Name types.String `tfsdk:"name"`
	Id   types.String `tfsdk:"id"`
}

var NetworkFederatedRealmsAttrTypes = map[string]attr.Type{
	"name": types.StringType,
	"id":   types.StringType,
}

var NetworkFederatedRealmsResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The federated realm name",
		Computed:            true,
	},
	"id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The federated realm id",
		Computed:            true,
	},
}

func ExpandNetworkFederatedRealms(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.NetworkFederatedRealms {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkFederatedRealmsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NetworkFederatedRealmsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.NetworkFederatedRealms {
	if m == nil {
		return nil
	}
	to := &ipam.NetworkFederatedRealms{
		Name: flex.ExpandStringPointer(m.Name),
		Id:   flex.ExpandStringPointer(m.Id),
	}
	return to
}

func FlattenNetworkFederatedRealms(ctx context.Context, from *ipam.NetworkFederatedRealms, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkFederatedRealmsAttrTypes)
	}
	m := NetworkFederatedRealmsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkFederatedRealmsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NetworkFederatedRealmsModel) Flatten(ctx context.Context, from *ipam.NetworkFederatedRealms, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NetworkFederatedRealmsModel{}
	}
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Id = flex.FlattenStringPointer(from.Id)
}
