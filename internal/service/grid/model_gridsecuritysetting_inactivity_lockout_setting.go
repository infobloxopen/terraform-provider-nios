package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type GridsecuritysettingInactivityLockoutSettingModel struct {
	AccountInactivityLockoutEnable   types.Bool  `tfsdk:"account_inactivity_lockout_enable"`
	InactiveDays                     types.Int64 `tfsdk:"inactive_days"`
	ReminderDays                     types.Int64 `tfsdk:"reminder_days"`
	ReactivateViaSerialConsoleEnable types.Bool  `tfsdk:"reactivate_via_serial_console_enable"`
	ReactivateViaRemoteConsoleEnable types.Bool  `tfsdk:"reactivate_via_remote_console_enable"`
}

var GridsecuritysettingInactivityLockoutSettingAttrTypes = map[string]attr.Type{
	"account_inactivity_lockout_enable":    types.BoolType,
	"inactive_days":                        types.Int64Type,
	"reminder_days":                        types.Int64Type,
	"reactivate_via_serial_console_enable": types.BoolType,
	"reactivate_via_remote_console_enable": types.BoolType,
}

var GridsecuritysettingInactivityLockoutSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"account_inactivity_lockout_enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enable/disable the account inactivity lockout.",
	},
	"inactive_days": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Number of days after which account gets locked out if user does not login.",
	},
	"reminder_days": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The number of days before the account lockout date when the appliance sends a reminder.",
	},
	"reactivate_via_serial_console_enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enable/disable reactivating user account by logging in from serial console.",
	},
	"reactivate_via_remote_console_enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enable/disable reactivating user account by logging in from remote console.",
	},
}

func ExpandGridsecuritysettingInactivityLockoutSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridsecuritysettingInactivityLockoutSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridsecuritysettingInactivityLockoutSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridsecuritysettingInactivityLockoutSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridsecuritysettingInactivityLockoutSetting {
	if m == nil {
		return nil
	}
	to := &grid.GridsecuritysettingInactivityLockoutSetting{
		AccountInactivityLockoutEnable:   flex.ExpandBoolPointer(m.AccountInactivityLockoutEnable),
		InactiveDays:                     flex.ExpandInt64Pointer(m.InactiveDays),
		ReminderDays:                     flex.ExpandInt64Pointer(m.ReminderDays),
		ReactivateViaSerialConsoleEnable: flex.ExpandBoolPointer(m.ReactivateViaSerialConsoleEnable),
		ReactivateViaRemoteConsoleEnable: flex.ExpandBoolPointer(m.ReactivateViaRemoteConsoleEnable),
	}
	return to
}

func FlattenGridsecuritysettingInactivityLockoutSetting(ctx context.Context, from *grid.GridsecuritysettingInactivityLockoutSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridsecuritysettingInactivityLockoutSettingAttrTypes)
	}
	m := GridsecuritysettingInactivityLockoutSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridsecuritysettingInactivityLockoutSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridsecuritysettingInactivityLockoutSettingModel) Flatten(ctx context.Context, from *grid.GridsecuritysettingInactivityLockoutSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridsecuritysettingInactivityLockoutSettingModel{}
	}
	m.AccountInactivityLockoutEnable = types.BoolPointerValue(from.AccountInactivityLockoutEnable)
	m.InactiveDays = flex.FlattenInt64Pointer(from.InactiveDays)
	m.ReminderDays = flex.FlattenInt64Pointer(from.ReminderDays)
	m.ReactivateViaSerialConsoleEnable = types.BoolPointerValue(from.ReactivateViaSerialConsoleEnable)
	m.ReactivateViaRemoteConsoleEnable = types.BoolPointerValue(from.ReactivateViaRemoteConsoleEnable)
}
