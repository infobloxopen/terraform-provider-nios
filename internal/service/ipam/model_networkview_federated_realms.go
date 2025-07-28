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

type NetworkviewFederatedRealmsModel struct {
	Name types.String `tfsdk:"name"`
	Id   types.String `tfsdk:"id"`
}

var NetworkviewFederatedRealmsAttrTypes = map[string]attr.Type{
	"name": types.StringType,
	"id":   types.StringType,
}

var NetworkviewFederatedRealmsResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The federated realm name",
	},
	"id": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The federated realm id",
	},
}

func ExpandNetworkviewFederatedRealms(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.NetworkviewFederatedRealms {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkviewFederatedRealmsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NetworkviewFederatedRealmsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.NetworkviewFederatedRealms {
	if m == nil {
		return nil
	}
	to := &ipam.NetworkviewFederatedRealms{
		Name: flex.ExpandStringPointer(m.Name),
		Id:   flex.ExpandStringPointer(m.Id),
	}
	return to
}

func FlattenNetworkviewFederatedRealms(ctx context.Context, from *ipam.NetworkviewFederatedRealms, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkviewFederatedRealmsAttrTypes)
	}
	m := NetworkviewFederatedRealmsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkviewFederatedRealmsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NetworkviewFederatedRealmsModel) Flatten(ctx context.Context, from *ipam.NetworkviewFederatedRealms, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NetworkviewFederatedRealmsModel{}
	}
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Id = flex.FlattenStringPointer(from.Id)
}
