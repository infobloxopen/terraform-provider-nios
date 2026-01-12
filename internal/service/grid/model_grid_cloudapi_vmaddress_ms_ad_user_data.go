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

type GridCloudapiVmaddressMsAdUserDataModel struct {
	ActiveUsersCount types.Int64 `tfsdk:"active_users_count"`
}

var GridCloudapiVmaddressMsAdUserDataAttrTypes = map[string]attr.Type{
	"active_users_count": types.Int64Type,
}

var GridCloudapiVmaddressMsAdUserDataResourceSchemaAttributes = map[string]schema.Attribute{
	"active_users_count": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of active users.",
	},
}

func ExpandGridCloudapiVmaddressMsAdUserData(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridCloudapiVmaddressMsAdUserData {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridCloudapiVmaddressMsAdUserDataModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridCloudapiVmaddressMsAdUserDataModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridCloudapiVmaddressMsAdUserData {
	if m == nil {
		return nil
	}
	to := &grid.GridCloudapiVmaddressMsAdUserData{}
	return to
}

func FlattenGridCloudapiVmaddressMsAdUserData(ctx context.Context, from *grid.GridCloudapiVmaddressMsAdUserData, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridCloudapiVmaddressMsAdUserDataAttrTypes)
	}
	m := GridCloudapiVmaddressMsAdUserDataModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridCloudapiVmaddressMsAdUserDataAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridCloudapiVmaddressMsAdUserDataModel) Flatten(ctx context.Context, from *grid.GridCloudapiVmaddressMsAdUserData, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridCloudapiVmaddressMsAdUserDataModel{}
	}
	m.ActiveUsersCount = flex.FlattenInt64Pointer(from.ActiveUsersCount)
}
