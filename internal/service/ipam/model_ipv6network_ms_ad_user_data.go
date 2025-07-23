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

type Ipv6networkMsAdUserDataModel struct {
	ActiveUsersCount types.Int64 `tfsdk:"active_users_count"`
}

var Ipv6networkMsAdUserDataAttrTypes = map[string]attr.Type{
	"active_users_count": types.Int64Type,
}

var Ipv6networkMsAdUserDataResourceSchemaAttributes = map[string]schema.Attribute{
	"active_users_count": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of active users.",
	},
}

func ExpandIpv6networkMsAdUserData(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.Ipv6networkMsAdUserData {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6networkMsAdUserDataModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *Ipv6networkMsAdUserDataModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.Ipv6networkMsAdUserData {
	if m == nil {
		return nil
	}
	to := &ipam.Ipv6networkMsAdUserData{}
	return to
}

func FlattenIpv6networkMsAdUserData(ctx context.Context, from *ipam.Ipv6networkMsAdUserData, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6networkMsAdUserDataAttrTypes)
	}
	m := Ipv6networkMsAdUserDataModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6networkMsAdUserDataAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Ipv6networkMsAdUserDataModel) Flatten(ctx context.Context, from *ipam.Ipv6networkMsAdUserData, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Ipv6networkMsAdUserDataModel{}
	}
	m.ActiveUsersCount = flex.FlattenInt64Pointer(from.ActiveUsersCount)
}
