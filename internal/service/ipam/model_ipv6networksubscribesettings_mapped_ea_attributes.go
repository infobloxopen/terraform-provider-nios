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

type Ipv6networksubscribesettingsMappedEaAttributesModel struct {
	Name     types.String `tfsdk:"name"`
	MappedEa types.String `tfsdk:"mapped_ea"`
}

var Ipv6networksubscribesettingsMappedEaAttributesAttrTypes = map[string]attr.Type{
	"name":      types.StringType,
	"mapped_ea": types.StringType,
}

var Ipv6networksubscribesettingsMappedEaAttributesResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The Cisco ISE attribute name that is enabled for publishsing from a Cisco ISE endpoint.",
	},
	"mapped_ea": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the extensible attribute definition object the Cisco ISE attribute that is enabled for subscription is mapped on.",
	},
}

func ExpandIpv6networksubscribesettingsMappedEaAttributes(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.Ipv6networksubscribesettingsMappedEaAttributes {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6networksubscribesettingsMappedEaAttributesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *Ipv6networksubscribesettingsMappedEaAttributesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.Ipv6networksubscribesettingsMappedEaAttributes {
	if m == nil {
		return nil
	}
	to := &ipam.Ipv6networksubscribesettingsMappedEaAttributes{
		Name:     flex.ExpandStringPointer(m.Name),
		MappedEa: flex.ExpandStringPointer(m.MappedEa),
	}
	return to
}

func FlattenIpv6networksubscribesettingsMappedEaAttributes(ctx context.Context, from *ipam.Ipv6networksubscribesettingsMappedEaAttributes, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6networksubscribesettingsMappedEaAttributesAttrTypes)
	}
	m := Ipv6networksubscribesettingsMappedEaAttributesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6networksubscribesettingsMappedEaAttributesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Ipv6networksubscribesettingsMappedEaAttributesModel) Flatten(ctx context.Context, from *ipam.Ipv6networksubscribesettingsMappedEaAttributes, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Ipv6networksubscribesettingsMappedEaAttributesModel{}
	}
	m.Name = flex.FlattenStringPointer(from.Name)
	m.MappedEa = flex.FlattenStringPointer(from.MappedEa)
}
