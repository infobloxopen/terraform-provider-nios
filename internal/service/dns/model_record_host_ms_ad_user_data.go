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

type RecordHostMsAdUserDataModel struct {
	ActiveUsersCount types.Int64 `tfsdk:"active_users_count"`
}

var RecordHostMsAdUserDataAttrTypes = map[string]attr.Type{
	"active_users_count": types.Int64Type,
}

var RecordHostMsAdUserDataResourceSchemaAttributes = map[string]schema.Attribute{
	"active_users_count": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of active users.",
	},
}

func ExpandRecordHostMsAdUserData(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.RecordHostMsAdUserData {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RecordHostMsAdUserDataModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *RecordHostMsAdUserDataModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.RecordHostMsAdUserData {
	if m == nil {
		return nil
	}
	to := &dns.RecordHostMsAdUserData{}
	return to
}

func FlattenRecordHostMsAdUserData(ctx context.Context, from *dns.RecordHostMsAdUserData, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordHostMsAdUserDataAttrTypes)
	}
	m := RecordHostMsAdUserDataModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RecordHostMsAdUserDataAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RecordHostMsAdUserDataModel) Flatten(ctx context.Context, from *dns.RecordHostMsAdUserData, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RecordHostMsAdUserDataModel{}
	}
	m.ActiveUsersCount = flex.FlattenInt64Pointer(from.ActiveUsersCount)
}
