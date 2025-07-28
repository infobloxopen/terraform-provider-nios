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

// TODO: function call support for VLANs
type NetworkVlansModel struct {
	Vlan types.Map    `tfsdk:"vlan"`
	Id   types.Int64  `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

var NetworkVlansAttrTypes = map[string]attr.Type{
	"vlan": types.MapType{ElemType: types.StringType},
	"id":   types.Int64Type,
	"name": types.StringType,
}

var NetworkVlansResourceSchemaAttributes = map[string]schema.Attribute{
	"vlan": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "Reference to the underlying StaticVlan object vlan.",
	},
	"id": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "VLAN ID value.",
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Name of the VLAN.",
	},
}

func ExpandNetworkVlans(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.NetworkVlans {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkVlansModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NetworkVlansModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.NetworkVlans {
	if m == nil {
		return nil
	}
	to := &ipam.NetworkVlans{
		Vlan: flex.ExpandFrameworkMapString(ctx, m.Vlan, diags),
	}
	return to
}

func FlattenNetworkVlans(ctx context.Context, from *ipam.NetworkVlans, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkVlansAttrTypes)
	}
	m := NetworkVlansModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkVlansAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NetworkVlansModel) Flatten(ctx context.Context, from *ipam.NetworkVlans, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NetworkVlansModel{}
	}
	m.Vlan = flex.FlattenFrameworkMapString(ctx, from.Vlan, diags)
	m.Id = flex.FlattenInt64Pointer(from.Id)
	m.Name = flex.FlattenStringPointer(from.Name)
}
