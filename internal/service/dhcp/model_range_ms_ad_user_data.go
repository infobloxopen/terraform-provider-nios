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

type RangeMsAdUserDataModel struct {
	ActiveUsersCount types.Int64 `tfsdk:"active_users_count"`
}

var RangeMsAdUserDataAttrTypes = map[string]attr.Type{
	"active_users_count": types.Int64Type,
}

var RangeMsAdUserDataResourceSchemaAttributes = map[string]schema.Attribute{
	"active_users_count": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of active users.",
	},
}

func ExpandRangeMsAdUserData(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dhcp.RangeMsAdUserData {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RangeMsAdUserDataModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *RangeMsAdUserDataModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.RangeMsAdUserData {
	if m == nil {
		return nil
	}
	to := &dhcp.RangeMsAdUserData{}
	return to
}

func FlattenRangeMsAdUserData(ctx context.Context, from *dhcp.RangeMsAdUserData, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RangeMsAdUserDataAttrTypes)
	}
	m := RangeMsAdUserDataModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RangeMsAdUserDataAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RangeMsAdUserDataModel) Flatten(ctx context.Context, from *dhcp.RangeMsAdUserData, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RangeMsAdUserDataModel{}
	}
	m.ActiveUsersCount = flex.FlattenInt64Pointer(from.ActiveUsersCount)
}
