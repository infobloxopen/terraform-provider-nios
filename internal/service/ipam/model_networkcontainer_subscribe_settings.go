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

type NetworkcontainerSubscribeSettingsModel struct {
	EnabledAttributes  types.List   `tfsdk:"enabled_attributes"`
	MappedEaAttributes types.Object `tfsdk:"mapped_ea_attributes"`
}

var NetworkcontainerSubscribeSettingsAttrTypes = map[string]attr.Type{
	"enabled_attributes":   types.ListType{ElemType: types.StringType},
	"mapped_ea_attributes": types.ObjectType{AttrTypes: NetworkcontainersubscribesettingsMappedEaAttributesAttrTypes},
}

var NetworkcontainerSubscribeSettingsResourceSchemaAttributes = map[string]schema.Attribute{
	"enabled_attributes": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The list of Cisco ISE attributes allowed for subscription.",
	},
	"mapped_ea_attributes": schema.SingleNestedAttribute{
		Attributes: NetworkcontainersubscribesettingsMappedEaAttributesResourceSchemaAttributes,
		Optional:   true,
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
		MappedEaAttributes: ExpandNetworkcontainersubscribesettingsMappedEaAttributes(ctx, m.MappedEaAttributes, diags),
	}
	return to
}

func FlattenNetworkcontainerSubscribeSettings(ctx context.Context, from *ipam.NetworkcontainerSubscribeSettings, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkcontainerSubscribeSettingsAttrTypes)
	}
	m := NetworkcontainerSubscribeSettingsModel{}
	m.Flatten(ctx, from, diags)
	// m.ExtAttrs = m.ExtAttrsAll
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
	m.MappedEaAttributes = FlattenNetworkcontainersubscribesettingsMappedEaAttributes(ctx, from.MappedEaAttributes, diags)
}
