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

type RecordAaaaMsAdUserDataModel struct {
	ActiveUsersCount types.Int64 `tfsdk:"active_users_count"`
}

var RecordAaaaMsAdUserDataAttrTypes = map[string]attr.Type{
	"active_users_count": types.Int64Type,
}

var RecordAaaaMsAdUserDataResourceSchemaAttributes = map[string]schema.Attribute{
	"active_users_count": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of active users.",
		PlanModifiers: []planmodifier.Int64{
			plancontrol.UseStateForUnknownInt64(),
		},
	},
}

func ExpandRecordAaaaMsAdUserData(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.RecordAaaaMsAdUserData {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RecordAaaaMsAdUserDataModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *RecordAaaaMsAdUserDataModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.RecordAaaaMsAdUserData {
	if m == nil {
		return nil
	}
	to := &dns.RecordAaaaMsAdUserData{}
	return to
}

func FlattenRecordAaaaMsAdUserData(ctx context.Context, from *dns.RecordAaaaMsAdUserData, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordAaaaMsAdUserDataAttrTypes)
	}
	m := RecordAaaaMsAdUserDataModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RecordAaaaMsAdUserDataAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RecordAaaaMsAdUserDataModel) Flatten(ctx context.Context, from *dns.RecordAaaaMsAdUserData, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RecordAaaaMsAdUserDataModel{}
	}
	m.ActiveUsersCount = flex.FlattenInt64Pointer(from.ActiveUsersCount)
}
