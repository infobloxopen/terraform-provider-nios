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

type NetworkviewMsAdUserDataModel struct {
	ActiveUsersCount types.Int64 `tfsdk:"active_users_count"`
}

var NetworkviewMsAdUserDataAttrTypes = map[string]attr.Type{
	"active_users_count": types.Int64Type,
}

var NetworkviewMsAdUserDataResourceSchemaAttributes = map[string]schema.Attribute{
	"active_users_count": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of active users.",
	},
}

func ExpandNetworkviewMsAdUserData(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.NetworkviewMsAdUserData {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkviewMsAdUserDataModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NetworkviewMsAdUserDataModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.NetworkviewMsAdUserData {
	if m == nil {
		return nil
	}
	to := &ipam.NetworkviewMsAdUserData{}
	return to
}

func FlattenNetworkviewMsAdUserData(ctx context.Context, from *ipam.NetworkviewMsAdUserData, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkviewMsAdUserDataAttrTypes)
	}
	m := NetworkviewMsAdUserDataModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkviewMsAdUserDataAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NetworkviewMsAdUserDataModel) Flatten(ctx context.Context, from *ipam.NetworkviewMsAdUserData, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NetworkviewMsAdUserDataModel{}
	}
	m.ActiveUsersCount = flex.FlattenInt64Pointer(from.ActiveUsersCount)
}
