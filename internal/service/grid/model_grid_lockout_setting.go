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

type GridLockoutSettingModel struct {
	EnableSequentialFailedLoginAttemptsLockout types.Bool  `tfsdk:"enable_sequential_failed_login_attempts_lockout"`
	SequentialAttempts                         types.Int64 `tfsdk:"sequential_attempts"`
	FailedLockoutDuration                      types.Int64 `tfsdk:"failed_lockout_duration"`
	NeverUnlockUser                            types.Bool  `tfsdk:"never_unlock_user"`
}

var GridLockoutSettingAttrTypes = map[string]attr.Type{
	"enable_sequential_failed_login_attempts_lockout": types.BoolType,
	"sequential_attempts":                             types.Int64Type,
	"failed_lockout_duration":                         types.Int64Type,
	"never_unlock_user":                               types.BoolType,
}

var GridLockoutSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"enable_sequential_failed_login_attempts_lockout": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enable/disable sequential failed login attempts lockout for local users",
	},
	"sequential_attempts": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The number of failed login attempts",
	},
	"failed_lockout_duration": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Time period the account remains locked after sequential failed login attempt lockout.",
	},
	"never_unlock_user": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Never unlock option is also provided and if set then user account is locked forever and only super user can unlock this account",
	},
}

func ExpandGridLockoutSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridLockoutSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridLockoutSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridLockoutSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridLockoutSetting {
	if m == nil {
		return nil
	}
	to := &grid.GridLockoutSetting{
		EnableSequentialFailedLoginAttemptsLockout: flex.ExpandBoolPointer(m.EnableSequentialFailedLoginAttemptsLockout),
		SequentialAttempts:                         flex.ExpandInt64Pointer(m.SequentialAttempts),
		FailedLockoutDuration:                      flex.ExpandInt64Pointer(m.FailedLockoutDuration),
		NeverUnlockUser:                            flex.ExpandBoolPointer(m.NeverUnlockUser),
	}
	return to
}

func FlattenGridLockoutSetting(ctx context.Context, from *grid.GridLockoutSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridLockoutSettingAttrTypes)
	}
	m := GridLockoutSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridLockoutSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridLockoutSettingModel) Flatten(ctx context.Context, from *grid.GridLockoutSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridLockoutSettingModel{}
	}
	m.EnableSequentialFailedLoginAttemptsLockout = types.BoolPointerValue(from.EnableSequentialFailedLoginAttemptsLockout)
	m.SequentialAttempts = flex.FlattenInt64Pointer(from.SequentialAttempts)
	m.FailedLockoutDuration = flex.FlattenInt64Pointer(from.FailedLockoutDuration)
	m.NeverUnlockUser = types.BoolPointerValue(from.NeverUnlockUser)
}
