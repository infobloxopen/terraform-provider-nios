package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type Ipv6networkSubscribeSettingsModel struct {
	EnabledAttributes  types.List `tfsdk:"enabled_attributes"`
	MappedEaAttributes types.List `tfsdk:"mapped_ea_attributes"`
}

var Ipv6networkSubscribeSettingsAttrTypes = map[string]attr.Type{
	"enabled_attributes":   types.ListType{ElemType: types.StringType},
	"mapped_ea_attributes": types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6networksubscribesettingsMappedEaAttributesAttrTypes}},
}

var Ipv6networkSubscribeSettingsResourceSchemaAttributes = map[string]schema.Attribute{
	"enabled_attributes": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The list of Cisco ISE attributes allowed for subscription.",
		Computed:            true,
		Validators: []validator.List{
			customvalidator.StringsInSlice([]string{
				"DOMAINNAME",
				"ENDPOINT_PROFILE",
				"SECURITY_GROUP",
				"SESSION_STATE",
				"SSID",
				"USERNAME",
				"VLAN",
			}),
			listvalidator.SizeAtLeast(1),
		},
	},
	"mapped_ea_attributes": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6networksubscribesettingsMappedEaAttributesResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "The list of NIOS extensible attributes to Cisco ISE attributes mappings.",
		Computed:            true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
	},
}

func ExpandIpv6networkSubscribeSettings(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.Ipv6networkSubscribeSettings {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6networkSubscribeSettingsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *Ipv6networkSubscribeSettingsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.Ipv6networkSubscribeSettings {
	if m == nil {
		return nil
	}
	to := &ipam.Ipv6networkSubscribeSettings{
		EnabledAttributes:  flex.ExpandFrameworkListString(ctx, m.EnabledAttributes, diags),
		MappedEaAttributes: flex.ExpandFrameworkListNestedBlock(ctx, m.MappedEaAttributes, diags, ExpandIpv6networksubscribesettingsMappedEaAttributes),
	}
	return to
}

func FlattenIpv6networkSubscribeSettings(ctx context.Context, from *ipam.Ipv6networkSubscribeSettings, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6networkSubscribeSettingsAttrTypes)
	}
	m := Ipv6networkSubscribeSettingsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6networkSubscribeSettingsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Ipv6networkSubscribeSettingsModel) Flatten(ctx context.Context, from *ipam.Ipv6networkSubscribeSettings, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Ipv6networkSubscribeSettingsModel{}
	}
	m.EnabledAttributes = flex.FlattenFrameworkListString(ctx, from.EnabledAttributes, diags)
	m.MappedEaAttributes = flex.FlattenFrameworkListNestedBlock(ctx, from.MappedEaAttributes, Ipv6networksubscribesettingsMappedEaAttributesAttrTypes, diags, FlattenIpv6networksubscribesettingsMappedEaAttributes)
}
