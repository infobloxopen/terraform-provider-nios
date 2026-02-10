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

type NetworkviewDdnsZonePrimariesDnsGridZoneModel struct {
	Ref  types.String `tfsdk:"ref"`
	Uuid types.String `tfsdk:"uuid"`
}

var NetworkviewDdnsZonePrimariesDnsGridZoneAttrTypes = map[string]attr.Type{
	"ref":  types.StringType,
	"uuid": types.StringType,
}

var NetworkviewDdnsZonePrimariesDnsGridZoneResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The reference to the DNS zone object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The uuid to the object.",
	},
}

func ExpandNetworkviewDdnsZonePrimariesDnsGridZone(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.NetworkviewDdnsZonePrimariesDnsGridZone {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkviewDdnsZonePrimariesDnsGridZoneModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NetworkviewDdnsZonePrimariesDnsGridZoneModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.NetworkviewDdnsZonePrimariesDnsGridZone {
	if m == nil {
		return nil
	}
	to := &ipam.NetworkviewDdnsZonePrimariesDnsGridZone{
		Ref: flex.ExpandStringPointer(m.Ref),
	}
	return to
}

func FlattenNetworkviewDdnsZonePrimariesDnsGridZone(ctx context.Context, from *ipam.NetworkviewDdnsZonePrimariesDnsGridZone, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkviewDdnsZonePrimariesDnsGridZoneAttrTypes)
	}
	m := NetworkviewDdnsZonePrimariesDnsGridZoneModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkviewDdnsZonePrimariesDnsGridZoneAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NetworkviewDdnsZonePrimariesDnsGridZoneModel) Flatten(ctx context.Context, from *ipam.NetworkviewDdnsZonePrimariesDnsGridZone, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NetworkviewDdnsZonePrimariesDnsGridZoneModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
}
