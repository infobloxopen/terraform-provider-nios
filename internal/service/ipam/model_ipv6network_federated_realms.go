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

type Ipv6networkFederatedRealmsModel struct {
	Name types.String `tfsdk:"name"`
	Id   types.String `tfsdk:"id"`
}

var Ipv6networkFederatedRealmsAttrTypes = map[string]attr.Type{
	"name": types.StringType,
	"id":   types.StringType,
}

var Ipv6networkFederatedRealmsResourceSchemaAttributes = map[string]schema.Attribute{
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

func ExpandIpv6networkFederatedRealms(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.Ipv6networkFederatedRealms {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6networkFederatedRealmsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *Ipv6networkFederatedRealmsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.Ipv6networkFederatedRealms {
	if m == nil {
		return nil
	}
	to := &ipam.Ipv6networkFederatedRealms{
		Name: flex.ExpandStringPointer(m.Name),
		Id:   flex.ExpandStringPointer(m.Id),
	}
	return to
}

func FlattenIpv6networkFederatedRealms(ctx context.Context, from *ipam.Ipv6networkFederatedRealms, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6networkFederatedRealmsAttrTypes)
	}
	m := Ipv6networkFederatedRealmsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6networkFederatedRealmsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Ipv6networkFederatedRealmsModel) Flatten(ctx context.Context, from *ipam.Ipv6networkFederatedRealms, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Ipv6networkFederatedRealmsModel{}
	}
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Id = flex.FlattenStringPointer(from.Id)
}
