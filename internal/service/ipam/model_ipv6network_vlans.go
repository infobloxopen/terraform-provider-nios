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

type Ipv6networkVlansModel struct {
	Vlan types.Map    `tfsdk:"vlan"`
	Id   types.Int64  `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

var Ipv6networkVlansAttrTypes = map[string]attr.Type{
	"vlan": types.MapType{ElemType: types.StringType},
	"id":   types.Int64Type,
	"name": types.StringType,
}

var Ipv6networkVlansResourceSchemaAttributes = map[string]schema.Attribute{
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
		Optional:            true,
	},
}

func ExpandIpv6networkVlans(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.Ipv6networkVlans {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6networkVlansModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *Ipv6networkVlansModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.Ipv6networkVlans {
	if m == nil {
		return nil
	}
	to := &ipam.Ipv6networkVlans{
		Vlan: flex.ExpandFrameworkMapString(ctx, m.Vlan, diags),
	}
	return to
}

func FlattenIpv6networkVlans(ctx context.Context, from *ipam.Ipv6networkVlans, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6networkVlansAttrTypes)
	}
	m := Ipv6networkVlansModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6networkVlansAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Ipv6networkVlansModel) Flatten(ctx context.Context, from *ipam.Ipv6networkVlans, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Ipv6networkVlansModel{}
	}
	m.Vlan = flex.FlattenFrameworkMapString(ctx, from.Vlan, diags)
	m.Id = flex.FlattenInt64Pointer(from.Id)
	m.Name = flex.FlattenStringPointer(from.Name)
}
