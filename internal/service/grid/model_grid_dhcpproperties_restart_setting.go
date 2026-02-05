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

type GridDhcppropertiesRestartSettingModel struct {
	Delay          types.Int64 `tfsdk:"delay"`
	Timeout        types.Int64 `tfsdk:"timeout"`
	RestartOffline types.Bool  `tfsdk:"restart_offline"`
}

var GridDhcppropertiesRestartSettingAttrTypes = map[string]attr.Type{
	"delay":           types.Int64Type,
	"timeout":         types.Int64Type,
	"restart_offline": types.BoolType,
}

var GridDhcppropertiesRestartSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"delay": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The time duration to delay a restart for a restart group.",
	},
	"timeout": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The duration of timeout for a restart group. The value \"-1\" means infinite.",
	},
	"restart_offline": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the Grid should try to restart offline member.",
	},
}

func ExpandGridDhcppropertiesRestartSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridDhcppropertiesRestartSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridDhcppropertiesRestartSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridDhcppropertiesRestartSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridDhcppropertiesRestartSetting {
	if m == nil {
		return nil
	}
	to := &grid.GridDhcppropertiesRestartSetting{
		Delay:          flex.ExpandInt64Pointer(m.Delay),
		Timeout:        flex.ExpandInt64Pointer(m.Timeout),
		RestartOffline: flex.ExpandBoolPointer(m.RestartOffline),
	}
	return to
}

func FlattenGridDhcppropertiesRestartSetting(ctx context.Context, from *grid.GridDhcppropertiesRestartSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridDhcppropertiesRestartSettingAttrTypes)
	}
	m := GridDhcppropertiesRestartSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridDhcppropertiesRestartSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridDhcppropertiesRestartSettingModel) Flatten(ctx context.Context, from *grid.GridDhcppropertiesRestartSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridDhcppropertiesRestartSettingModel{}
	}
	m.Delay = flex.FlattenInt64Pointer(from.Delay)
	m.Timeout = flex.FlattenInt64Pointer(from.Timeout)
	m.RestartOffline = types.BoolPointerValue(from.RestartOffline)
}
