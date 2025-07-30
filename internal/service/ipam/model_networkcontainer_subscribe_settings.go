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

type NetworkcontainerSubscribeSettingsModel struct {
	EnabledAttributes  types.List `tfsdk:"enabled_attributes"`
	MappedEaAttributes types.List `tfsdk:"mapped_ea_attributes"`
}

var NetworkcontainerSubscribeSettingsAttrTypes = map[string]attr.Type{
	"enabled_attributes":   types.ListType{ElemType: types.StringType},
	"mapped_ea_attributes": types.ListType{ElemType: types.ObjectType{AttrTypes: NetworkcontainersubscribesettingsMappedEaAttributesAttrTypes}},
}

var NetworkcontainerSubscribeSettingsResourceSchemaAttributes = map[string]schema.Attribute{
	"enabled_attributes": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The list of Cisco ISE attributes allowed for subscription.",
	},
	"mapped_ea_attributes": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NetworkcontainersubscribesettingsMappedEaAttributesResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "The list of NIOS extensible attributes to Cisco ISE attributes mappings.",
	},
}

func ExpandNetworkcontainerSubscribeSettings(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.NetworkcontainerSubscribeSettings {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkcontainerSubscribeSettingsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NetworkcontainerSubscribeSettingsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.NetworkcontainerSubscribeSettings {
	if m == nil {
		return nil
	}
	to := &ipam.NetworkcontainerSubscribeSettings{
		EnabledAttributes:  flex.ExpandFrameworkListString(ctx, m.EnabledAttributes, diags),
		MappedEaAttributes: flex.ExpandFrameworkListNestedBlock(ctx, m.MappedEaAttributes, diags, ExpandNetworkcontainersubscribesettingsMappedEaAttributes),
	}
	return to
}

func FlattenNetworkcontainerSubscribeSettings(ctx context.Context, from *ipam.NetworkcontainerSubscribeSettings, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkcontainerSubscribeSettingsAttrTypes)
	}
	m := NetworkcontainerSubscribeSettingsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkcontainerSubscribeSettingsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NetworkcontainerSubscribeSettingsModel) Flatten(ctx context.Context, from *ipam.NetworkcontainerSubscribeSettings, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NetworkcontainerSubscribeSettingsModel{}
	}
	m.EnabledAttributes = flex.FlattenFrameworkListString(ctx, from.EnabledAttributes, diags)
	m.MappedEaAttributes = flex.FlattenFrameworkListNestedBlock(ctx, from.MappedEaAttributes, NetworkcontainersubscribesettingsMappedEaAttributesAttrTypes, diags, FlattenNetworkcontainersubscribesettingsMappedEaAttributes)
}
