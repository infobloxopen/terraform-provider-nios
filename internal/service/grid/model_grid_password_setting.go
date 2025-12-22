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

type GridPasswordSettingModel struct {
	PasswordMinLength types.Int64 `tfsdk:"password_min_length"`
	NumLowerChar      types.Int64 `tfsdk:"num_lower_char"`
	NumUpperChar      types.Int64 `tfsdk:"num_upper_char"`
	NumNumericChar    types.Int64 `tfsdk:"num_numeric_char"`
	NumSymbolChar     types.Int64 `tfsdk:"num_symbol_char"`
	CharsToChange     types.Int64 `tfsdk:"chars_to_change"`
	ExpireDays        types.Int64 `tfsdk:"expire_days"`
	ReminderDays      types.Int64 `tfsdk:"reminder_days"`
	ForceResetEnable  types.Bool  `tfsdk:"force_reset_enable"`
	ExpireEnable      types.Bool  `tfsdk:"expire_enable"`
	HistoryEnable     types.Bool  `tfsdk:"history_enable"`
	NumPasswordsSaved types.Int64 `tfsdk:"num_passwords_saved"`
	MinPasswordAge    types.Int64 `tfsdk:"min_password_age"`
}

var GridPasswordSettingAttrTypes = map[string]attr.Type{
	"password_min_length": types.Int64Type,
	"num_lower_char":      types.Int64Type,
	"num_upper_char":      types.Int64Type,
	"num_numeric_char":    types.Int64Type,
	"num_symbol_char":     types.Int64Type,
	"chars_to_change":     types.Int64Type,
	"expire_days":         types.Int64Type,
	"reminder_days":       types.Int64Type,
	"force_reset_enable":  types.BoolType,
	"expire_enable":       types.BoolType,
	"history_enable":      types.BoolType,
	"num_passwords_saved": types.Int64Type,
	"min_password_age":    types.Int64Type,
}

var GridPasswordSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"password_min_length": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The minimum length of the password.",
	},
	"num_lower_char": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The minimum number of lowercase characters.",
	},
	"num_upper_char": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The minimum number of uppercase characters.",
	},
	"num_numeric_char": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The minimum number of numeric characters.",
	},
	"num_symbol_char": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The minimum number of symbol characters. The allowed characters are ! @ # $ % ^ & * ( ).",
	},
	"chars_to_change": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The minimum number of characters that must be changed when revising an admin password.",
	},
	"expire_days": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The number of days of the password expiration period (if enabled).",
	},
	"reminder_days": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The number of days before the password expiration date when the appliance sends a reminder.",
	},
	"force_reset_enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If set to True, all new users must change their passwords when they first log in to the system, and existing users must change the passwords that were just reset.",
	},
	"expire_enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If set to True, password expiration is enabled.",
	},
	"history_enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enable/disable the password history.",
	},
	"num_passwords_saved": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Number of saved passwords if password history is enabled. Can be set between 1 to 20.",
	},
	"min_password_age": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Minimum password age in days before password can be updated. Can be set between 1 to 9998 days.",
	},
}

func ExpandGridPasswordSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridPasswordSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridPasswordSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridPasswordSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridPasswordSetting {
	if m == nil {
		return nil
	}
	to := &grid.GridPasswordSetting{
		PasswordMinLength: flex.ExpandInt64Pointer(m.PasswordMinLength),
		NumLowerChar:      flex.ExpandInt64Pointer(m.NumLowerChar),
		NumUpperChar:      flex.ExpandInt64Pointer(m.NumUpperChar),
		NumNumericChar:    flex.ExpandInt64Pointer(m.NumNumericChar),
		NumSymbolChar:     flex.ExpandInt64Pointer(m.NumSymbolChar),
		CharsToChange:     flex.ExpandInt64Pointer(m.CharsToChange),
		ExpireDays:        flex.ExpandInt64Pointer(m.ExpireDays),
		ReminderDays:      flex.ExpandInt64Pointer(m.ReminderDays),
		ForceResetEnable:  flex.ExpandBoolPointer(m.ForceResetEnable),
		ExpireEnable:      flex.ExpandBoolPointer(m.ExpireEnable),
		HistoryEnable:     flex.ExpandBoolPointer(m.HistoryEnable),
		NumPasswordsSaved: flex.ExpandInt64Pointer(m.NumPasswordsSaved),
		MinPasswordAge:    flex.ExpandInt64Pointer(m.MinPasswordAge),
	}
	return to
}

func FlattenGridPasswordSetting(ctx context.Context, from *grid.GridPasswordSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridPasswordSettingAttrTypes)
	}
	m := GridPasswordSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridPasswordSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridPasswordSettingModel) Flatten(ctx context.Context, from *grid.GridPasswordSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridPasswordSettingModel{}
	}
	m.PasswordMinLength = flex.FlattenInt64Pointer(from.PasswordMinLength)
	m.NumLowerChar = flex.FlattenInt64Pointer(from.NumLowerChar)
	m.NumUpperChar = flex.FlattenInt64Pointer(from.NumUpperChar)
	m.NumNumericChar = flex.FlattenInt64Pointer(from.NumNumericChar)
	m.NumSymbolChar = flex.FlattenInt64Pointer(from.NumSymbolChar)
	m.CharsToChange = flex.FlattenInt64Pointer(from.CharsToChange)
	m.ExpireDays = flex.FlattenInt64Pointer(from.ExpireDays)
	m.ReminderDays = flex.FlattenInt64Pointer(from.ReminderDays)
	m.ForceResetEnable = types.BoolPointerValue(from.ForceResetEnable)
	m.ExpireEnable = types.BoolPointerValue(from.ExpireEnable)
	m.HistoryEnable = types.BoolPointerValue(from.HistoryEnable)
	m.NumPasswordsSaved = flex.FlattenInt64Pointer(from.NumPasswordsSaved)
	m.MinPasswordAge = flex.FlattenInt64Pointer(from.MinPasswordAge)
}
