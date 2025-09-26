package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type RecordHostIpv4addrMsAdUserDataModel struct {
	ActiveUsersCount types.Int64 `tfsdk:"active_users_count"`
}

var RecordHostIpv4addrMsAdUserDataAttrTypes = map[string]attr.Type{
	"active_users_count": types.Int64Type,
}

var RecordHostIpv4addrMsAdUserDataResourceSchemaAttributes = map[string]schema.Attribute{
	"active_users_count": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of active users.",
	},
}

func ExpandRecordHostIpv4addrMsAdUserData(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.RecordHostIpv4addrMsAdUserData {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RecordHostIpv4addrMsAdUserDataModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *RecordHostIpv4addrMsAdUserDataModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.RecordHostIpv4addrMsAdUserData {
	if m == nil {
		return nil
	}
	to := &dns.RecordHostIpv4addrMsAdUserData{}
	return to
}

func FlattenRecordHostIpv4addrMsAdUserData(ctx context.Context, from *dns.RecordHostIpv4addrMsAdUserData, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordHostIpv4addrMsAdUserDataAttrTypes)
	}
	m := RecordHostIpv4addrMsAdUserDataModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RecordHostIpv4addrMsAdUserDataAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RecordHostIpv4addrMsAdUserDataModel) Flatten(ctx context.Context, from *dns.RecordHostIpv4addrMsAdUserData, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RecordHostIpv4addrMsAdUserDataModel{}
	}
	m.ActiveUsersCount = flex.FlattenInt64Pointer(from.ActiveUsersCount)
}
