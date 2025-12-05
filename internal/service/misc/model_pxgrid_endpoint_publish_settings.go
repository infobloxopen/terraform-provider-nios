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

type PxgridEndpointPublishSettingsModel struct {
	EnabledAttributes types.List `tfsdk:"enabled_attributes"`
}

var PxgridEndpointPublishSettingsAttrTypes = map[string]attr.Type{
	"enabled_attributes": types.ListType{ElemType: types.StringType},
}

var PxgridEndpointPublishSettingsResourceSchemaAttributes = map[string]schema.Attribute{
	"enabled_attributes": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of NIOS extensible attributes enalbed for publishsing to Cisco ISE endpoint.",
	},
}

func ExpandPxgridEndpointPublishSettings(ctx context.Context, o types.Object, diags *diag.Diagnostics) *misc.PxgridEndpointPublishSettings {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m PxgridEndpointPublishSettingsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *PxgridEndpointPublishSettingsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *misc.PxgridEndpointPublishSettings {
	if m == nil {
		return nil
	}
	to := &misc.PxgridEndpointPublishSettings{
		EnabledAttributes: flex.ExpandFrameworkListString(ctx, m.EnabledAttributes, diags),
	}
	return to
}

func FlattenPxgridEndpointPublishSettings(ctx context.Context, from *misc.PxgridEndpointPublishSettings, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(PxgridEndpointPublishSettingsAttrTypes)
	}
	m := PxgridEndpointPublishSettingsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, PxgridEndpointPublishSettingsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *PxgridEndpointPublishSettingsModel) Flatten(ctx context.Context, from *misc.PxgridEndpointPublishSettings, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = PxgridEndpointPublishSettingsModel{}
	}
	m.EnabledAttributes = flex.FlattenFrameworkListString(ctx, from.EnabledAttributes, diags)
}
