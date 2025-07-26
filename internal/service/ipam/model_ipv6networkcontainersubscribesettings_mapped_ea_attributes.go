package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type Ipv6networkcontainersubscribesettingsMappedEaAttributesModel struct {
	Name     types.String `tfsdk:"name"`
	MappedEa types.String `tfsdk:"mapped_ea"`
}

var Ipv6networkcontainersubscribesettingsMappedEaAttributesAttrTypes = map[string]attr.Type{
	"name":      types.StringType,
	"mapped_ea": types.StringType,
}

var Ipv6networkcontainersubscribesettingsMappedEaAttributesResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The Cisco ISE attribute name that is enabled for publishsing from a Cisco ISE endpoint.",
	},
	"mapped_ea": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of the extensible attribute definition object the Cisco ISE attribute that is enabled for subscription is mapped on.",
		Validators: []validator.String{
			stringvalidator.OneOf(
				"ACCOUNT_SESSION_ID",
				"AUDIT_SESSION_ID",
				"EPS_STATUS",
				"IP_ADDRESS",
				"MAC",
				"NAS_IP_ADDRESS",
				"NAS_PORT_ID",
				"POSTURE_STATUS",
				"POSTURE_TIMESTAMP",
			),
		},
	},
}

func ExpandIpv6networkcontainersubscribesettingsMappedEaAttributes(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.Ipv6networkcontainersubscribesettingsMappedEaAttributes {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6networkcontainersubscribesettingsMappedEaAttributesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *Ipv6networkcontainersubscribesettingsMappedEaAttributesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.Ipv6networkcontainersubscribesettingsMappedEaAttributes {
	if m == nil {
		return nil
	}
	to := &ipam.Ipv6networkcontainersubscribesettingsMappedEaAttributes{
		Name:     flex.ExpandStringPointer(m.Name),
		MappedEa: flex.ExpandStringPointer(m.MappedEa),
	}
	return to
}

func FlattenIpv6networkcontainersubscribesettingsMappedEaAttributes(ctx context.Context, from *ipam.Ipv6networkcontainersubscribesettingsMappedEaAttributes, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6networkcontainersubscribesettingsMappedEaAttributesAttrTypes)
	}
	m := Ipv6networkcontainersubscribesettingsMappedEaAttributesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6networkcontainersubscribesettingsMappedEaAttributesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Ipv6networkcontainersubscribesettingsMappedEaAttributesModel) Flatten(ctx context.Context, from *ipam.Ipv6networkcontainersubscribesettingsMappedEaAttributes, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Ipv6networkcontainersubscribesettingsMappedEaAttributesModel{}
	}
	m.Name = flex.FlattenStringPointer(from.Name)
	m.MappedEa = flex.FlattenStringPointer(from.MappedEa)
}
