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

type GridRestartBannerSettingModel struct {
	Enabled                  types.Bool `tfsdk:"enabled"`
	EnableDoubleConfirmation types.Bool `tfsdk:"enable_double_confirmation"`
}

var GridRestartBannerSettingAttrTypes = map[string]attr.Type{
	"enabled":                    types.BoolType,
	"enable_double_confirmation": types.BoolType,
}

var GridRestartBannerSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If set to True, the restart banner is enabled.",
	},
	"enable_double_confirmation": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If set to True, the user is required to input name before restarting the services.",
	},
}

func ExpandGridRestartBannerSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridRestartBannerSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridRestartBannerSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridRestartBannerSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridRestartBannerSetting {
	if m == nil {
		return nil
	}
	to := &grid.GridRestartBannerSetting{
		Enabled:                  flex.ExpandBoolPointer(m.Enabled),
		EnableDoubleConfirmation: flex.ExpandBoolPointer(m.EnableDoubleConfirmation),
	}
	return to
}

func FlattenGridRestartBannerSetting(ctx context.Context, from *grid.GridRestartBannerSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridRestartBannerSettingAttrTypes)
	}
	m := GridRestartBannerSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridRestartBannerSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridRestartBannerSettingModel) Flatten(ctx context.Context, from *grid.GridRestartBannerSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridRestartBannerSettingModel{}
	}
	m.Enabled = types.BoolPointerValue(from.Enabled)
	m.EnableDoubleConfirmation = types.BoolPointerValue(from.EnableDoubleConfirmation)
}
