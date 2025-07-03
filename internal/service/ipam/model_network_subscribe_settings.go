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

type NetworkSubscribeSettingsModel struct {
	EnabledAttributes  types.List   `tfsdk:"enabled_attributes"`
	MappedEaAttributes types.Object `tfsdk:"mapped_ea_attributes"`
}

var NetworkSubscribeSettingsAttrTypes = map[string]attr.Type{
	"enabled_attributes":   types.ListType{ElemType: types.StringType},
	"mapped_ea_attributes": types.ObjectType{AttrTypes: NetworksubscribesettingsMappedEaAttributesAttrTypes},
}

var NetworkSubscribeSettingsResourceSchemaAttributes = map[string]schema.Attribute{
	"enabled_attributes": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The list of Cisco ISE attributes allowed for subscription.",
	},
	"mapped_ea_attributes": schema.SingleNestedAttribute{
		Attributes: NetworksubscribesettingsMappedEaAttributesResourceSchemaAttributes,
		Optional:   true,
	},
}

func ExpandNetworkSubscribeSettings(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.NetworkSubscribeSettings {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkSubscribeSettingsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NetworkSubscribeSettingsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.NetworkSubscribeSettings {
	if m == nil {
		return nil
	}
	to := &ipam.NetworkSubscribeSettings{
		EnabledAttributes:  flex.ExpandFrameworkListString(ctx, m.EnabledAttributes, diags),
		MappedEaAttributes: ExpandNetworksubscribesettingsMappedEaAttributes(ctx, m.MappedEaAttributes, diags),
	}
	return to
}

func FlattenNetworkSubscribeSettings(ctx context.Context, from *ipam.NetworkSubscribeSettings, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkSubscribeSettingsAttrTypes)
	}
	m := NetworkSubscribeSettingsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkSubscribeSettingsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NetworkSubscribeSettingsModel) Flatten(ctx context.Context, from *ipam.NetworkSubscribeSettings, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NetworkSubscribeSettingsModel{}
	}
	m.EnabledAttributes = flex.FlattenFrameworkListString(ctx, from.EnabledAttributes, diags)
	m.MappedEaAttributes = FlattenNetworksubscribesettingsMappedEaAttributes(ctx, from.MappedEaAttributes, diags)
}
