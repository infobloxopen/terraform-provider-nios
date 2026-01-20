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

type GridTrafficCaptureRecQueriesSettingModel struct {
	RecursiveClientsCountTriggerEnable types.Bool  `tfsdk:"recursive_clients_count_trigger_enable"`
	RecursiveClientsCountThreshold     types.Int64 `tfsdk:"recursive_clients_count_threshold"`
	RecursiveClientsCountReset         types.Int64 `tfsdk:"recursive_clients_count_reset"`
}

var GridTrafficCaptureRecQueriesSettingAttrTypes = map[string]attr.Type{
	"recursive_clients_count_trigger_enable": types.BoolType,
	"recursive_clients_count_threshold":      types.Int64Type,
	"recursive_clients_count_reset":          types.Int64Type,
}

var GridTrafficCaptureRecQueriesSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"recursive_clients_count_trigger_enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enable triggering automated traffic capture based on outgoing recursive queries count.",
	},
	"recursive_clients_count_threshold": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Concurrent outgoing recursive queries count below which traffic capture will be triggered.",
	},
	"recursive_clients_count_reset": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Concurrent outgoing recursive queries count below which traffic capture will be stopped.",
	},
}

func ExpandGridTrafficCaptureRecQueriesSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridTrafficCaptureRecQueriesSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridTrafficCaptureRecQueriesSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridTrafficCaptureRecQueriesSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridTrafficCaptureRecQueriesSetting {
	if m == nil {
		return nil
	}
	to := &grid.GridTrafficCaptureRecQueriesSetting{
		RecursiveClientsCountTriggerEnable: flex.ExpandBoolPointer(m.RecursiveClientsCountTriggerEnable),
		RecursiveClientsCountThreshold:     flex.ExpandInt64Pointer(m.RecursiveClientsCountThreshold),
		RecursiveClientsCountReset:         flex.ExpandInt64Pointer(m.RecursiveClientsCountReset),
	}
	return to
}

func FlattenGridTrafficCaptureRecQueriesSetting(ctx context.Context, from *grid.GridTrafficCaptureRecQueriesSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridTrafficCaptureRecQueriesSettingAttrTypes)
	}
	m := GridTrafficCaptureRecQueriesSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridTrafficCaptureRecQueriesSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridTrafficCaptureRecQueriesSettingModel) Flatten(ctx context.Context, from *grid.GridTrafficCaptureRecQueriesSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridTrafficCaptureRecQueriesSettingModel{}
	}
	m.RecursiveClientsCountTriggerEnable = types.BoolPointerValue(from.RecursiveClientsCountTriggerEnable)
	m.RecursiveClientsCountThreshold = flex.FlattenInt64Pointer(from.RecursiveClientsCountThreshold)
	m.RecursiveClientsCountReset = flex.FlattenInt64Pointer(from.RecursiveClientsCountReset)
}
