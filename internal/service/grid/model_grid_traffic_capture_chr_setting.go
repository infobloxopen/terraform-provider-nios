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

type GridTrafficCaptureChrSettingModel struct {
	ChrTriggerEnable       types.Bool  `tfsdk:"chr_trigger_enable"`
	ChrThreshold           types.Int64 `tfsdk:"chr_threshold"`
	ChrReset               types.Int64 `tfsdk:"chr_reset"`
	ChrMinCacheUtilization types.Int64 `tfsdk:"chr_min_cache_utilization"`
}

var GridTrafficCaptureChrSettingAttrTypes = map[string]attr.Type{
	"chr_trigger_enable":        types.BoolType,
	"chr_threshold":             types.Int64Type,
	"chr_reset":                 types.Int64Type,
	"chr_min_cache_utilization": types.Int64Type,
}

var GridTrafficCaptureChrSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"chr_trigger_enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enable triggering automated traffic capture based on cache hit ratio thresholds.",
	},
	"chr_threshold": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "DNS Cache hit ratio threshold(%) below which traffic capture will be triggered.",
	},
	"chr_reset": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "DNS Cache hit ratio threshold(%) above which traffic capture will be triggered.",
	},
	"chr_min_cache_utilization": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Minimum DNS cache utilization threshold(%) for triggering traffic capture based on DNS cache hit ratio.",
	},
}

func ExpandGridTrafficCaptureChrSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridTrafficCaptureChrSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridTrafficCaptureChrSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridTrafficCaptureChrSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridTrafficCaptureChrSetting {
	if m == nil {
		return nil
	}
	to := &grid.GridTrafficCaptureChrSetting{
		ChrTriggerEnable:       flex.ExpandBoolPointer(m.ChrTriggerEnable),
		ChrThreshold:           flex.ExpandInt64Pointer(m.ChrThreshold),
		ChrReset:               flex.ExpandInt64Pointer(m.ChrReset),
		ChrMinCacheUtilization: flex.ExpandInt64Pointer(m.ChrMinCacheUtilization),
	}
	return to
}

func FlattenGridTrafficCaptureChrSetting(ctx context.Context, from *grid.GridTrafficCaptureChrSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridTrafficCaptureChrSettingAttrTypes)
	}
	m := GridTrafficCaptureChrSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridTrafficCaptureChrSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridTrafficCaptureChrSettingModel) Flatten(ctx context.Context, from *grid.GridTrafficCaptureChrSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridTrafficCaptureChrSettingModel{}
	}
	m.ChrTriggerEnable = types.BoolPointerValue(from.ChrTriggerEnable)
	m.ChrThreshold = flex.FlattenInt64Pointer(from.ChrThreshold)
	m.ChrReset = flex.FlattenInt64Pointer(from.ChrReset)
	m.ChrMinCacheUtilization = flex.FlattenInt64Pointer(from.ChrMinCacheUtilization)
}
