package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/ipam"

	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/flex"
)

type Ipv6networkcontainerMsAdUserDataModel struct {
	ActiveUsersCount types.Int64 `tfsdk:"active_users_count"`
}

var Ipv6networkcontainerMsAdUserDataAttrTypes = map[string]attr.Type{
	"active_users_count": types.Int64Type,
}

var Ipv6networkcontainerMsAdUserDataResourceSchemaAttributes = map[string]schema.Attribute{
	"active_users_count": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of active users.",
	},
}

func ExpandIpv6networkcontainerMsAdUserData(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.Ipv6networkcontainerMsAdUserData {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6networkcontainerMsAdUserDataModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *Ipv6networkcontainerMsAdUserDataModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.Ipv6networkcontainerMsAdUserData {
	if m == nil {
		return nil
	}
	to := &ipam.Ipv6networkcontainerMsAdUserData{}
	return to
}

func FlattenIpv6networkcontainerMsAdUserData(ctx context.Context, from *ipam.Ipv6networkcontainerMsAdUserData, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6networkcontainerMsAdUserDataAttrTypes)
	}
	m := Ipv6networkcontainerMsAdUserDataModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6networkcontainerMsAdUserDataAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Ipv6networkcontainerMsAdUserDataModel) Flatten(ctx context.Context, from *ipam.Ipv6networkcontainerMsAdUserData, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Ipv6networkcontainerMsAdUserDataModel{}
	}
	m.ActiveUsersCount = flex.FlattenInt64Pointer(from.ActiveUsersCount)
}
