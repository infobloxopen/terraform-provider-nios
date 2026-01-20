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

type GridDnsRestartSettingModel struct {
	Delay          types.Int64 `tfsdk:"delay"`
	Timeout        types.Int64 `tfsdk:"timeout"`
	RestartOffline types.Bool  `tfsdk:"restart_offline"`
}

var GridDnsRestartSettingAttrTypes = map[string]attr.Type{
	"delay":           types.Int64Type,
	"timeout":         types.Int64Type,
	"restart_offline": types.BoolType,
}

var GridDnsRestartSettingResourceSchemaAttributes = map[string]schema.Attribute{
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

func ExpandGridDnsRestartSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridDnsRestartSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridDnsRestartSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridDnsRestartSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridDnsRestartSetting {
	if m == nil {
		return nil
	}
	to := &grid.GridDnsRestartSetting{
		Delay:          flex.ExpandInt64Pointer(m.Delay),
		Timeout:        flex.ExpandInt64Pointer(m.Timeout),
		RestartOffline: flex.ExpandBoolPointer(m.RestartOffline),
	}
	return to
}

func FlattenGridDnsRestartSetting(ctx context.Context, from *grid.GridDnsRestartSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridDnsRestartSettingAttrTypes)
	}
	m := GridDnsRestartSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridDnsRestartSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridDnsRestartSettingModel) Flatten(ctx context.Context, from *grid.GridDnsRestartSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridDnsRestartSettingModel{}
	}
	m.Delay = flex.FlattenInt64Pointer(from.Delay)
	m.Timeout = flex.FlattenInt64Pointer(from.Timeout)
	m.RestartOffline = types.BoolPointerValue(from.RestartOffline)
}
