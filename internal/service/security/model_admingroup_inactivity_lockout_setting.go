package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type AdmingroupInactivityLockoutSettingModel struct {
	AccountInactivityLockoutEnable   types.Bool  `tfsdk:"account_inactivity_lockout_enable"`
	InactiveDays                     types.Int64 `tfsdk:"inactive_days"`
	ReminderDays                     types.Int64 `tfsdk:"reminder_days"`
	ReactivateViaSerialConsoleEnable types.Bool  `tfsdk:"reactivate_via_serial_console_enable"`
	ReactivateViaRemoteConsoleEnable types.Bool  `tfsdk:"reactivate_via_remote_console_enable"`
}

var AdmingroupInactivityLockoutSettingAttrTypes = map[string]attr.Type{
	"account_inactivity_lockout_enable":    types.BoolType,
	"inactive_days":                        types.Int64Type,
	"reminder_days":                        types.Int64Type,
	"reactivate_via_serial_console_enable": types.BoolType,
	"reactivate_via_remote_console_enable": types.BoolType,
}

var AdmingroupInactivityLockoutSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"account_inactivity_lockout_enable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Enable/disable the account inactivity lockout.",
	},
	"inactive_days": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(30),
		Validators: []validator.Int64{
			int64validator.Between(2, 9999),
		},
		MarkdownDescription: "Number of days after which account gets locked out if user does not login.",
	},
	"reminder_days": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(15),
		Validators: []validator.Int64{
			int64validator.Between(1, 30),
		},
		MarkdownDescription: "The number of days before the account lockout date when the appliance sends a reminder.",
	},
	"reactivate_via_serial_console_enable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Enable/disable reactivating user account by logging in from serial console.",
	},
	"reactivate_via_remote_console_enable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Enable/disable reactivating user account by logging in from remote console.",
	},
}

func ExpandAdmingroupInactivityLockoutSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.AdmingroupInactivityLockoutSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AdmingroupInactivityLockoutSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AdmingroupInactivityLockoutSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.AdmingroupInactivityLockoutSetting {
	if m == nil {
		return nil
	}
	to := &security.AdmingroupInactivityLockoutSetting{
		AccountInactivityLockoutEnable:   flex.ExpandBoolPointer(m.AccountInactivityLockoutEnable),
		InactiveDays:                     flex.ExpandInt64Pointer(m.InactiveDays),
		ReminderDays:                     flex.ExpandInt64Pointer(m.ReminderDays),
		ReactivateViaSerialConsoleEnable: flex.ExpandBoolPointer(m.ReactivateViaSerialConsoleEnable),
		ReactivateViaRemoteConsoleEnable: flex.ExpandBoolPointer(m.ReactivateViaRemoteConsoleEnable),
	}
	return to
}

func FlattenAdmingroupInactivityLockoutSetting(ctx context.Context, from *security.AdmingroupInactivityLockoutSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AdmingroupInactivityLockoutSettingAttrTypes)
	}
	m := AdmingroupInactivityLockoutSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AdmingroupInactivityLockoutSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AdmingroupInactivityLockoutSettingModel) Flatten(ctx context.Context, from *security.AdmingroupInactivityLockoutSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AdmingroupInactivityLockoutSettingModel{}
	}
	m.AccountInactivityLockoutEnable = types.BoolPointerValue(from.AccountInactivityLockoutEnable)
	m.InactiveDays = flex.FlattenInt64Pointer(from.InactiveDays)
	m.ReminderDays = flex.FlattenInt64Pointer(from.ReminderDays)
	m.ReactivateViaSerialConsoleEnable = types.BoolPointerValue(from.ReactivateViaSerialConsoleEnable)
	m.ReactivateViaRemoteConsoleEnable = types.BoolPointerValue(from.ReactivateViaRemoteConsoleEnable)
}
