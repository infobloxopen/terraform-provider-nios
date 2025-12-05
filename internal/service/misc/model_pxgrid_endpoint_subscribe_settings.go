package misc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/misc"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type PxgridEndpointSubscribeSettingsModel struct {
	EnabledAttributes  types.List `tfsdk:"enabled_attributes"`
	MappedEaAttributes types.List `tfsdk:"mapped_ea_attributes"`
}

var PxgridEndpointSubscribeSettingsAttrTypes = map[string]attr.Type{
	"enabled_attributes":   types.ListType{ElemType: types.StringType},
	"mapped_ea_attributes": types.ListType{ElemType: types.ObjectType{AttrTypes: PxgridendpointsubscribesettingsMappedEaAttributesAttrTypes}},
}

var PxgridEndpointSubscribeSettingsResourceSchemaAttributes = map[string]schema.Attribute{
	"enabled_attributes": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of Cisco ISE attributes allowed for subscription.",
	},
	"mapped_ea_attributes": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: PxgridendpointsubscribesettingsMappedEaAttributesResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of NIOS extensible attributes to Cisco ISE attributes mappings.",
	},
}

func ExpandPxgridEndpointSubscribeSettings(ctx context.Context, o types.Object, diags *diag.Diagnostics) *misc.PxgridEndpointSubscribeSettings {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m PxgridEndpointSubscribeSettingsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *PxgridEndpointSubscribeSettingsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *misc.PxgridEndpointSubscribeSettings {
	if m == nil {
		return nil
	}
	to := &misc.PxgridEndpointSubscribeSettings{
		EnabledAttributes:  flex.ExpandFrameworkListString(ctx, m.EnabledAttributes, diags),
		MappedEaAttributes: flex.ExpandFrameworkListNestedBlock(ctx, m.MappedEaAttributes, diags, ExpandPxgridendpointsubscribesettingsMappedEaAttributes),
	}
	return to
}

func FlattenPxgridEndpointSubscribeSettings(ctx context.Context, from *misc.PxgridEndpointSubscribeSettings, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(PxgridEndpointSubscribeSettingsAttrTypes)
	}
	m := PxgridEndpointSubscribeSettingsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, PxgridEndpointSubscribeSettingsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *PxgridEndpointSubscribeSettingsModel) Flatten(ctx context.Context, from *misc.PxgridEndpointSubscribeSettings, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = PxgridEndpointSubscribeSettingsModel{}
	}
	m.EnabledAttributes = flex.FlattenFrameworkListString(ctx, from.EnabledAttributes, diags)
	m.MappedEaAttributes = flex.FlattenFrameworkListNestedBlock(ctx, from.MappedEaAttributes, PxgridendpointsubscribesettingsMappedEaAttributesAttrTypes, diags, FlattenPxgridendpointsubscribesettingsMappedEaAttributes)
}
