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

type NetworkcontainerMsAdUserDataModel struct {
	ActiveUsersCount types.Int64 `tfsdk:"active_users_count"`
}

var NetworkcontainerMsAdUserDataAttrTypes = map[string]attr.Type{
	"active_users_count": types.Int64Type,
}

var NetworkcontainerMsAdUserDataResourceSchemaAttributes = map[string]schema.Attribute{
	"active_users_count": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of active users.",
	},
}

func ExpandNetworkcontainerMsAdUserData(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.NetworkcontainerMsAdUserData {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkcontainerMsAdUserDataModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NetworkcontainerMsAdUserDataModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.NetworkcontainerMsAdUserData {
	if m == nil {
		return nil
	}
	to := &ipam.NetworkcontainerMsAdUserData{}
	return to
}

func FlattenNetworkcontainerMsAdUserData(ctx context.Context, from *ipam.NetworkcontainerMsAdUserData, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkcontainerMsAdUserDataAttrTypes)
	}
	m := NetworkcontainerMsAdUserDataModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkcontainerMsAdUserDataAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NetworkcontainerMsAdUserDataModel) Flatten(ctx context.Context, from *ipam.NetworkcontainerMsAdUserData, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NetworkcontainerMsAdUserDataModel{}
	}
	m.ActiveUsersCount = flex.FlattenInt64Pointer(from.ActiveUsersCount)
}
