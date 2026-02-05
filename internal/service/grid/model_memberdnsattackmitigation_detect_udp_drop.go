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

type MemberdnsattackmitigationDetectUdpDropModel struct {
	Enable       types.Bool  `tfsdk:"enable"`
	High         types.Int64 `tfsdk:"high"`
	IntervalMax  types.Int64 `tfsdk:"interval_max"`
	IntervalMin  types.Int64 `tfsdk:"interval_min"`
	IntervalTime types.Int64 `tfsdk:"interval_time"`
	Low          types.Int64 `tfsdk:"low"`
}

var MemberdnsattackmitigationDetectUdpDropAttrTypes = map[string]attr.Type{
	"enable":        types.BoolType,
	"high":          types.Int64Type,
	"interval_max":  types.Int64Type,
	"interval_min":  types.Int64Type,
	"interval_time": types.Int64Type,
	"low":           types.Int64Type,
}

var MemberdnsattackmitigationDetectUdpDropResourceSchemaAttributes = map[string]schema.Attribute{
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

func ExpandMemberdnsattackmitigationDetectUdpDrop(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberdnsattackmitigationDetectUdpDrop {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberdnsattackmitigationDetectUdpDropModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberdnsattackmitigationDetectUdpDropModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberdnsattackmitigationDetectUdpDrop {
	if m == nil {
		return nil
	}
	to := &grid.MemberdnsattackmitigationDetectUdpDrop{
		Enable:       flex.ExpandBoolPointer(m.Enable),
		High:         flex.ExpandInt64Pointer(m.High),
		IntervalMax:  flex.ExpandInt64Pointer(m.IntervalMax),
		IntervalMin:  flex.ExpandInt64Pointer(m.IntervalMin),
		IntervalTime: flex.ExpandInt64Pointer(m.IntervalTime),
		Low:          flex.ExpandInt64Pointer(m.Low),
	}
	return to
}

func FlattenMemberdnsattackmitigationDetectUdpDrop(ctx context.Context, from *grid.MemberdnsattackmitigationDetectUdpDrop, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberdnsattackmitigationDetectUdpDropAttrTypes)
	}
	m := MemberdnsattackmitigationDetectUdpDropModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberdnsattackmitigationDetectUdpDropAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberdnsattackmitigationDetectUdpDropModel) Flatten(ctx context.Context, from *grid.MemberdnsattackmitigationDetectUdpDrop, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberdnsattackmitigationDetectUdpDropModel{}
	}
	m.Enable = types.BoolPointerValue(from.Enable)
	m.High = flex.FlattenInt64Pointer(from.High)
	m.IntervalMax = flex.FlattenInt64Pointer(from.IntervalMax)
	m.IntervalMin = flex.FlattenInt64Pointer(from.IntervalMin)
	m.IntervalTime = flex.FlattenInt64Pointer(from.IntervalTime)
	m.Low = flex.FlattenInt64Pointer(from.Low)
}
