package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type GridLicensePoolContainerModel struct {
	Ref                   types.String `tfsdk:"ref"`
	Uuid                  types.String `tfsdk:"uuid"`
	LastEntitlementUpdate types.Int64  `tfsdk:"last_entitlement_update"`
	LpcUid                types.String `tfsdk:"lpc_uid"`
}

var GridLicensePoolContainerAttrTypes = map[string]attr.Type{
	"ref":                     types.StringType,
	"uuid":                    types.StringType,
	"last_entitlement_update": types.Int64Type,
	"lpc_uid":                 types.StringType,
}

var GridLicensePoolContainerResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The UUID of the object.",
	},
	"last_entitlement_update": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The timestamp when the last pool licenses were updated.",
	},
	"lpc_uid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The world-wide unique ID for the license pool container.",
	},
}

func ExpandGridLicensePoolContainer(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridLicensePoolContainer {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridLicensePoolContainerModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridLicensePoolContainerModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridLicensePoolContainer {
	if m == nil {
		return nil
	}
	to := &grid.GridLicensePoolContainer{
		Ref: flex.ExpandStringPointer(m.Ref),
	}
	return to
}

func FlattenGridLicensePoolContainer(ctx context.Context, from *grid.GridLicensePoolContainer, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridLicensePoolContainerAttrTypes)
	}
	m := GridLicensePoolContainerModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridLicensePoolContainerAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridLicensePoolContainerModel) Flatten(ctx context.Context, from *grid.GridLicensePoolContainer, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridLicensePoolContainerModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.LastEntitlementUpdate = flex.FlattenInt64Pointer(from.LastEntitlementUpdate)
	m.LpcUid = flex.FlattenStringPointer(from.LpcUid)
}
