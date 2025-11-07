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

type GridLicensePoolSubpoolsModel struct {
	Key        types.String `tfsdk:"key"`
	Installed  types.Int64  `tfsdk:"installed"`
	ExpiryDate types.Int64  `tfsdk:"expiry_date"`
}

var GridLicensePoolSubpoolsAttrTypes = map[string]attr.Type{
	"key":         types.StringType,
	"installed":   types.Int64Type,
	"expiry_date": types.Int64Type,
}

var GridLicensePoolSubpoolsResourceSchemaAttributes = map[string]schema.Attribute{
	"key": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The license string for the license subpool.",
	},
	"installed": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The total number of dynamic licenses allowed for this license subpool.",
	},
	"expiry_date": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "License expiration date.",
	},
}

func ExpandGridLicensePoolSubpools(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridLicensePoolSubpools {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridLicensePoolSubpoolsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridLicensePoolSubpoolsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridLicensePoolSubpools {
	if m == nil {
		return nil
	}
	to := &grid.GridLicensePoolSubpools{}
	return to
}

func FlattenGridLicensePoolSubpools(ctx context.Context, from *grid.GridLicensePoolSubpools, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridLicensePoolSubpoolsAttrTypes)
	}
	m := GridLicensePoolSubpoolsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridLicensePoolSubpoolsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridLicensePoolSubpoolsModel) Flatten(ctx context.Context, from *grid.GridLicensePoolSubpools, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridLicensePoolSubpoolsModel{}
	}
	m.Key = flex.FlattenStringPointer(from.Key)
	m.Installed = flex.FlattenInt64Pointer(from.Installed)
	m.ExpiryDate = flex.FlattenInt64Pointer(from.ExpiryDate)
}
