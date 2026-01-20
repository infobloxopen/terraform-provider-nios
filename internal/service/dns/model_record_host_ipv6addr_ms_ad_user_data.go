package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	"github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/plancontrol"
)

type RecordHostIpv6addrMsAdUserDataModel struct {
	ActiveUsersCount types.Int64 `tfsdk:"active_users_count"`
}

var RecordHostIpv6addrMsAdUserDataAttrTypes = map[string]attr.Type{
	"active_users_count": types.Int64Type,
}

var RecordHostIpv6addrMsAdUserDataResourceSchemaAttributes = map[string]schema.Attribute{
	"active_users_count": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of active users.",
		PlanModifiers: []planmodifier.Int64{
			plancontrol.UseStateForUnknownInt64(),
		},
	},
}

func ExpandRecordHostIpv6addrMsAdUserData(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.RecordHostIpv6addrMsAdUserData {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RecordHostIpv6addrMsAdUserDataModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *RecordHostIpv6addrMsAdUserDataModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.RecordHostIpv6addrMsAdUserData {
	if m == nil {
		return nil
	}
	to := &dns.RecordHostIpv6addrMsAdUserData{}
	return to
}

func FlattenRecordHostIpv6addrMsAdUserData(ctx context.Context, from *dns.RecordHostIpv6addrMsAdUserData, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordHostIpv6addrMsAdUserDataAttrTypes)
	}
	m := RecordHostIpv6addrMsAdUserDataModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RecordHostIpv6addrMsAdUserDataAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RecordHostIpv6addrMsAdUserDataModel) Flatten(ctx context.Context, from *dns.RecordHostIpv6addrMsAdUserData, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RecordHostIpv6addrMsAdUserDataModel{}
	}
	m.ActiveUsersCount = flex.FlattenInt64Pointer(from.ActiveUsersCount)
}
