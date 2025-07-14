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

type NetworkMsAdUserDataModel struct {
	ActiveUsersCount types.Int64 `tfsdk:"active_users_count"`
}

var NetworkMsAdUserDataAttrTypes = map[string]attr.Type{
	"active_users_count": types.Int64Type,
}

var NetworkMsAdUserDataResourceSchemaAttributes = map[string]schema.Attribute{
	"active_users_count": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of active users.",
	},
}

func ExpandNetworkMsAdUserData(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.NetworkMsAdUserData {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkMsAdUserDataModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NetworkMsAdUserDataModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.NetworkMsAdUserData {
	if m == nil {
		return nil
	}
	to := &ipam.NetworkMsAdUserData{}
	return to
}

func FlattenNetworkMsAdUserData(ctx context.Context, from *ipam.NetworkMsAdUserData, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkMsAdUserDataAttrTypes)
	}
	m := NetworkMsAdUserDataModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkMsAdUserDataAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NetworkMsAdUserDataModel) Flatten(ctx context.Context, from *ipam.NetworkMsAdUserData, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NetworkMsAdUserDataModel{}
	}
	m.ActiveUsersCount = flex.FlattenInt64Pointer(from.ActiveUsersCount)
}
