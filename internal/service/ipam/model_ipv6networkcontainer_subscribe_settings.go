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

type Ipv6networkcontainerSubscribeSettingsModel struct {
	EnabledAttributes  types.List `tfsdk:"enabled_attributes"`
	MappedEaAttributes types.List `tfsdk:"mapped_ea_attributes"`
}

var Ipv6networkcontainerSubscribeSettingsAttrTypes = map[string]attr.Type{
	"enabled_attributes":   types.ListType{ElemType: types.StringType},
	"mapped_ea_attributes": types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6networkcontainersubscribesettingsMappedEaAttributesAttrTypes}},
}

var Ipv6networkcontainerSubscribeSettingsResourceSchemaAttributes = map[string]schema.Attribute{
	"enabled_attributes": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The list of Cisco ISE attributes allowed for subscription.",
	},
	"mapped_ea_attributes": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6networkcontainersubscribesettingsMappedEaAttributesResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "The list of NIOS extensible attributes to Cisco ISE attributes mappings.",
	},
}

func ExpandIpv6networkcontainerSubscribeSettings(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.Ipv6networkcontainerSubscribeSettings {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6networkcontainerSubscribeSettingsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *Ipv6networkcontainerSubscribeSettingsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.Ipv6networkcontainerSubscribeSettings {
	if m == nil {
		return nil
	}
	to := &ipam.Ipv6networkcontainerSubscribeSettings{
		EnabledAttributes:  flex.ExpandFrameworkListString(ctx, m.EnabledAttributes, diags),
		MappedEaAttributes: flex.ExpandFrameworkListNestedBlock(ctx, m.MappedEaAttributes, diags, ExpandIpv6networkcontainersubscribesettingsMappedEaAttributes),
	}
	return to
}

func FlattenIpv6networkcontainerSubscribeSettings(ctx context.Context, from *ipam.Ipv6networkcontainerSubscribeSettings, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6networkcontainerSubscribeSettingsAttrTypes)
	}
	m := Ipv6networkcontainerSubscribeSettingsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6networkcontainerSubscribeSettingsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Ipv6networkcontainerSubscribeSettingsModel) Flatten(ctx context.Context, from *ipam.Ipv6networkcontainerSubscribeSettings, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Ipv6networkcontainerSubscribeSettingsModel{}
	}
	m.EnabledAttributes = flex.FlattenFrameworkListString(ctx, from.EnabledAttributes, diags)
	m.MappedEaAttributes = flex.FlattenFrameworkListNestedBlock(ctx, from.MappedEaAttributes, Ipv6networkcontainersubscribesettingsMappedEaAttributesAttrTypes, diags, FlattenIpv6networkcontainersubscribesettingsMappedEaAttributes)
}
