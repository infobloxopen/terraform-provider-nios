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

type GriddnsattackmitigationDetectChrModel struct {
	Enable       types.Bool  `tfsdk:"enable"`
	High         types.Int64 `tfsdk:"high"`
	IntervalMax  types.Int64 `tfsdk:"interval_max"`
	IntervalMin  types.Int64 `tfsdk:"interval_min"`
	IntervalTime types.Int64 `tfsdk:"interval_time"`
	Low          types.Int64 `tfsdk:"low"`
}

var GriddnsattackmitigationDetectChrAttrTypes = map[string]attr.Type{
	"enable":        types.BoolType,
	"high":          types.Int64Type,
	"interval_max":  types.Int64Type,
	"interval_min":  types.Int64Type,
	"interval_time": types.Int64Type,
	"low":           types.Int64Type,
}

var GriddnsattackmitigationDetectChrResourceSchemaAttributes = map[string]schema.Attribute{
	"enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if DNS attack detection is enabled or not.",
	},
	"high": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The high threshold value (in percentage) for starting DNS attack detection.",
	},
	"interval_max": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The maximum number of events that have occurred before processing DNS attack detection.",
	},
	"interval_min": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The minimum number of events that have occurred before processing DNS attack detection.",
	},
	"interval_time": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The time interval between detection processing.",
	},
	"low": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The low threshold value (in percentage) for starting DNS attack detection.",
	},
}

func ExpandGriddnsattackmitigationDetectChr(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GriddnsattackmitigationDetectChr {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GriddnsattackmitigationDetectChrModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GriddnsattackmitigationDetectChrModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GriddnsattackmitigationDetectChr {
	if m == nil {
		return nil
	}
	to := &grid.GriddnsattackmitigationDetectChr{
		Enable:       flex.ExpandBoolPointer(m.Enable),
		High:         flex.ExpandInt64Pointer(m.High),
		IntervalMax:  flex.ExpandInt64Pointer(m.IntervalMax),
		IntervalMin:  flex.ExpandInt64Pointer(m.IntervalMin),
		IntervalTime: flex.ExpandInt64Pointer(m.IntervalTime),
		Low:          flex.ExpandInt64Pointer(m.Low),
	}
	return to
}

func FlattenGriddnsattackmitigationDetectChr(ctx context.Context, from *grid.GriddnsattackmitigationDetectChr, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GriddnsattackmitigationDetectChrAttrTypes)
	}
	m := GriddnsattackmitigationDetectChrModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GriddnsattackmitigationDetectChrAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GriddnsattackmitigationDetectChrModel) Flatten(ctx context.Context, from *grid.GriddnsattackmitigationDetectChr, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GriddnsattackmitigationDetectChrModel{}
	}
	m.Enable = types.BoolPointerValue(from.Enable)
	m.High = flex.FlattenInt64Pointer(from.High)
	m.IntervalMax = flex.FlattenInt64Pointer(from.IntervalMax)
	m.IntervalMin = flex.FlattenInt64Pointer(from.IntervalMin)
	m.IntervalTime = flex.FlattenInt64Pointer(from.IntervalTime)
	m.Low = flex.FlattenInt64Pointer(from.Low)
}
