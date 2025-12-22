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

type GridTrafficCaptureQpsSettingModel struct {
	QpsTriggerEnable types.Bool  `tfsdk:"qps_trigger_enable"`
	QpsThreshold     types.Int64 `tfsdk:"qps_threshold"`
	QpsReset         types.Int64 `tfsdk:"qps_reset"`
}

var GridTrafficCaptureQpsSettingAttrTypes = map[string]attr.Type{
	"qps_trigger_enable": types.BoolType,
	"qps_threshold":      types.Int64Type,
	"qps_reset":          types.Int64Type,
}

var GridTrafficCaptureQpsSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"qps_trigger_enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enable triggering automated traffic capture based on DNS queries per second threshold.",
	},
	"qps_threshold": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "DNS queries per second threshold below which traffic capture will be triggered.",
	},
	"qps_reset": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "DNS queries per second threshold below which traffic capture will be stopped.",
	},
}

func ExpandGridTrafficCaptureQpsSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridTrafficCaptureQpsSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridTrafficCaptureQpsSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridTrafficCaptureQpsSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridTrafficCaptureQpsSetting {
	if m == nil {
		return nil
	}
	to := &grid.GridTrafficCaptureQpsSetting{
		QpsTriggerEnable: flex.ExpandBoolPointer(m.QpsTriggerEnable),
		QpsThreshold:     flex.ExpandInt64Pointer(m.QpsThreshold),
		QpsReset:         flex.ExpandInt64Pointer(m.QpsReset),
	}
	return to
}

func FlattenGridTrafficCaptureQpsSetting(ctx context.Context, from *grid.GridTrafficCaptureQpsSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridTrafficCaptureQpsSettingAttrTypes)
	}
	m := GridTrafficCaptureQpsSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridTrafficCaptureQpsSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridTrafficCaptureQpsSettingModel) Flatten(ctx context.Context, from *grid.GridTrafficCaptureQpsSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridTrafficCaptureQpsSettingModel{}
	}
	m.QpsTriggerEnable = types.BoolPointerValue(from.QpsTriggerEnable)
	m.QpsThreshold = flex.FlattenInt64Pointer(from.QpsThreshold)
	m.QpsReset = flex.FlattenInt64Pointer(from.QpsReset)
}
