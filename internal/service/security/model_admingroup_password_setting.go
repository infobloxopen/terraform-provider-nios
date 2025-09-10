package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type AdmingroupPasswordSettingModel struct {
	ExpireEnable types.Bool  `tfsdk:"expire_enable"`
	ExpireDays   types.Int64 `tfsdk:"expire_days"`
	ReminderDays types.Int64 `tfsdk:"reminder_days"`
}

var AdmingroupPasswordSettingAttrTypes = map[string]attr.Type{
	"expire_enable": types.BoolType,
	"expire_days":   types.Int64Type,
	"reminder_days": types.Int64Type,
}

var AdmingroupPasswordSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"expire_enable": schema.BoolAttribute{
		Optional: true,
		//Computed:            true,
		//Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Whether password expiry enabled or not.",
	},
	"expire_days": schema.Int64Attribute{
		Optional: true,
		//Computed:            true,
		//Default:             int64default.StaticInt64(30),
		MarkdownDescription: "The days that password must expire",
	},
	"reminder_days": schema.Int64Attribute{
		Optional: true,
		//Computed:            true,
		//Default:             int64default.StaticInt64(15),
		MarkdownDescription: "Days to show up reminder prior to expiration",
	},
}

func ExpandAdmingroupPasswordSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.AdmingroupPasswordSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AdmingroupPasswordSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AdmingroupPasswordSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.AdmingroupPasswordSetting {
	if m == nil {
		return nil
	}
	to := &security.AdmingroupPasswordSetting{
		ExpireEnable: flex.ExpandBoolPointer(m.ExpireEnable),
		ExpireDays:   flex.ExpandInt64Pointer(m.ExpireDays),
		ReminderDays: flex.ExpandInt64Pointer(m.ReminderDays),
	}
	return to
}

func FlattenAdmingroupPasswordSetting(ctx context.Context, from *security.AdmingroupPasswordSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AdmingroupPasswordSettingAttrTypes)
	}
	m := AdmingroupPasswordSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AdmingroupPasswordSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AdmingroupPasswordSettingModel) Flatten(ctx context.Context, from *security.AdmingroupPasswordSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AdmingroupPasswordSettingModel{}
	}
	m.ExpireEnable = types.BoolPointerValue(from.ExpireEnable)
	m.ExpireDays = flex.FlattenInt64Pointer(from.ExpireDays)
	m.ReminderDays = flex.FlattenInt64Pointer(from.ReminderDays)
}
