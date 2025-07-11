package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dtc"

	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/flex"
)

type DtcPoolLbDynamicRatioAlternateModel struct {
	Method              types.String `tfsdk:"method"`
	Monitor             types.String `tfsdk:"monitor"`
	MonitorMetric       types.String `tfsdk:"monitor_metric"`
	MonitorWeighing     types.String `tfsdk:"monitor_weighing"`
	InvertMonitorMetric types.Bool   `tfsdk:"invert_monitor_metric"`
}

var DtcPoolLbDynamicRatioAlternateAttrTypes = map[string]attr.Type{
	"method":                types.StringType,
	"monitor":               types.StringType,
	"monitor_metric":        types.StringType,
	"monitor_weighing":      types.StringType,
	"invert_monitor_metric": types.BoolType,
}

var DtcPoolLbDynamicRatioAlternateResourceSchemaAttributes = map[string]schema.Attribute{
	"method": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("MONITOR"),
		Validators: []validator.String{
			stringvalidator.OneOf("MONITOR", "ROUND_TRIP_DELAY"),
		},
		MarkdownDescription: "The method of the DTC dynamic ratio load balancing.",
	},
	"monitor": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The DTC monitor output of which will be used for dynamic ratio load balancing.",
	},
	"monitor_metric": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The metric of the DTC SNMP monitor that will be used for dynamic weighing.",
	},
	"monitor_weighing": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("RATIO"),
		Validators: []validator.String{
			stringvalidator.OneOf("PRIORITY", "RATIO"),
		},
		MarkdownDescription: "The DTC monitor weight. 'PRIORITY' means that all clients will be forwarded to the least loaded server. 'RATIO' means that distribution will be calculated based on dynamic weights.",
	},
	"invert_monitor_metric": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether the inverted values of the DTC SNMP monitor metric will be used.",
	},
}

func ExpandDtcPoolLbDynamicRatioAlternate(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dtc.DtcPoolLbDynamicRatioAlternate {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DtcPoolLbDynamicRatioAlternateModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DtcPoolLbDynamicRatioAlternateModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dtc.DtcPoolLbDynamicRatioAlternate {
	if m == nil {
		return nil
	}
	to := &dtc.DtcPoolLbDynamicRatioAlternate{
		Method:              flex.ExpandStringPointer(m.Method),
		Monitor:             flex.ExpandStringPointer(m.Monitor),
		MonitorMetric:       flex.ExpandStringPointer(m.MonitorMetric),
		MonitorWeighing:     flex.ExpandStringPointer(m.MonitorWeighing),
		InvertMonitorMetric: flex.ExpandBoolPointer(m.InvertMonitorMetric),
	}
	return to
}

func FlattenDtcPoolLbDynamicRatioAlternate(ctx context.Context, from *dtc.DtcPoolLbDynamicRatioAlternate, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DtcPoolLbDynamicRatioAlternateAttrTypes)
	}
	m := DtcPoolLbDynamicRatioAlternateModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DtcPoolLbDynamicRatioAlternateAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DtcPoolLbDynamicRatioAlternateModel) Flatten(ctx context.Context, from *dtc.DtcPoolLbDynamicRatioAlternate, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DtcPoolLbDynamicRatioAlternateModel{}
	}
	m.Method = flex.FlattenStringPointer(from.Method)
	m.Monitor = flex.FlattenStringPointer(from.Monitor)
	m.MonitorMetric = flex.FlattenStringPointer(from.MonitorMetric)
	m.MonitorWeighing = flex.FlattenStringPointer(from.MonitorWeighing)
	m.InvertMonitorMetric = types.BoolPointerValue(from.InvertMonitorMetric)
}
