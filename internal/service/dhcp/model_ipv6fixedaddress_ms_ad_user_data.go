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

type Ipv6fixedaddressMsAdUserDataModel struct {
	ActiveUsersCount types.Int64 `tfsdk:"active_users_count"`
}

var Ipv6fixedaddressMsAdUserDataAttrTypes = map[string]attr.Type{
	"active_users_count": types.Int64Type,
}

var Ipv6fixedaddressMsAdUserDataResourceSchemaAttributes = map[string]schema.Attribute{
	"active_users_count": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of active users.",
	},
}

func ExpandIpv6fixedaddressMsAdUserData(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dhcp.Ipv6fixedaddressMsAdUserData {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6fixedaddressMsAdUserDataModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *Ipv6fixedaddressMsAdUserDataModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.Ipv6fixedaddressMsAdUserData {
	if m == nil {
		return nil
	}
	to := &dhcp.Ipv6fixedaddressMsAdUserData{}
	return to
}

func FlattenIpv6fixedaddressMsAdUserData(ctx context.Context, from *dhcp.Ipv6fixedaddressMsAdUserData, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6fixedaddressMsAdUserDataAttrTypes)
	}
	m := Ipv6fixedaddressMsAdUserDataModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6fixedaddressMsAdUserDataAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Ipv6fixedaddressMsAdUserDataModel) Flatten(ctx context.Context, from *dhcp.Ipv6fixedaddressMsAdUserData, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Ipv6fixedaddressMsAdUserDataModel{}
	}
	m.ActiveUsersCount = flex.FlattenInt64Pointer(from.ActiveUsersCount)
}
