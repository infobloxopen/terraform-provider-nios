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

type NetworkcontainerFederatedRealmsModel struct {
	Name types.String `tfsdk:"name"`
	Id   types.String `tfsdk:"id"`
}

var NetworkcontainerFederatedRealmsAttrTypes = map[string]attr.Type{
	"name": types.StringType,
	"id":   types.StringType,
}

var NetworkcontainerFederatedRealmsResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The federated realm name",
	},
	"id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The federated realm id",
	},
}

func ExpandNetworkcontainerFederatedRealms(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.NetworkcontainerFederatedRealms {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkcontainerFederatedRealmsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NetworkcontainerFederatedRealmsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.NetworkcontainerFederatedRealms {
	if m == nil {
		return nil
	}
	to := &ipam.NetworkcontainerFederatedRealms{
		Name: flex.ExpandStringPointer(m.Name),
		Id:   flex.ExpandStringPointer(m.Id),
	}
	return to
}

func FlattenNetworkcontainerFederatedRealms(ctx context.Context, from *ipam.NetworkcontainerFederatedRealms, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkcontainerFederatedRealmsAttrTypes)
	}
	m := NetworkcontainerFederatedRealmsModel{}
	m.Flatten(ctx, from, diags)
	// m.ExtAttrs = m.ExtAttrsAll
	t, d := types.ObjectValueFrom(ctx, NetworkcontainerFederatedRealmsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NetworkcontainerFederatedRealmsModel) Flatten(ctx context.Context, from *ipam.NetworkcontainerFederatedRealms, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NetworkcontainerFederatedRealmsModel{}
	}
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Id = flex.FlattenStringPointer(from.Id)
}
