package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dhcp"

	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/flex"
)

type FixedaddressMsAdUserDataModel struct {
	ActiveUsersCount types.Int64 `tfsdk:"active_users_count"`
}

var FixedaddressMsAdUserDataAttrTypes = map[string]attr.Type{
	"active_users_count": types.Int64Type,
}

var FixedaddressMsAdUserDataResourceSchemaAttributes = map[string]schema.Attribute{
	"active_users_count": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of active users.",
	},
}

func ExpandFixedaddressMsAdUserData(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dhcp.FixedaddressMsAdUserData {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m FixedaddressMsAdUserDataModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *FixedaddressMsAdUserDataModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.FixedaddressMsAdUserData {
	if m == nil {
		return nil
	}
	to := &dhcp.FixedaddressMsAdUserData{}
	return to
}

func FlattenFixedaddressMsAdUserData(ctx context.Context, from *dhcp.FixedaddressMsAdUserData, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(FixedaddressMsAdUserDataAttrTypes)
	}
	m := FixedaddressMsAdUserDataModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrs = m.ExtAttrsAll
	t, d := types.ObjectValueFrom(ctx, FixedaddressMsAdUserDataAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *FixedaddressMsAdUserDataModel) Flatten(ctx context.Context, from *dhcp.FixedaddressMsAdUserData, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = FixedaddressMsAdUserDataModel{}
	}
	m.ActiveUsersCount = flex.FlattenInt64Pointer(from.ActiveUsersCount)
}
