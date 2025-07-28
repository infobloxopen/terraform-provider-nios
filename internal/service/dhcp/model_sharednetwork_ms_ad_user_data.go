package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type SharednetworkMsAdUserDataModel struct {
	ActiveUsersCount types.Int64 `tfsdk:"active_users_count"`
}

var SharednetworkMsAdUserDataAttrTypes = map[string]attr.Type{
	"active_users_count": types.Int64Type,
}

var SharednetworkMsAdUserDataResourceSchemaAttributes = map[string]schema.Attribute{
	"active_users_count": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of active users.",
	},
}

func ExpandSharednetworkMsAdUserData(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dhcp.SharednetworkMsAdUserData {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m SharednetworkMsAdUserDataModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *SharednetworkMsAdUserDataModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.SharednetworkMsAdUserData {
	if m == nil {
		return nil
	}
	to := &dhcp.SharednetworkMsAdUserData{}
	return to
}

func FlattenSharednetworkMsAdUserData(ctx context.Context, from *dhcp.SharednetworkMsAdUserData, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SharednetworkMsAdUserDataAttrTypes)
	}
	m := SharednetworkMsAdUserDataModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, SharednetworkMsAdUserDataAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *SharednetworkMsAdUserDataModel) Flatten(ctx context.Context, from *dhcp.SharednetworkMsAdUserData, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = SharednetworkMsAdUserDataModel{}
	}
	m.ActiveUsersCount = flex.FlattenInt64Pointer(from.ActiveUsersCount)
}
