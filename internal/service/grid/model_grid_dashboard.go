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

type GridDashboardModel struct {
	Ref                                      types.String `tfsdk:"ref"`
	Uuid                                     types.String `tfsdk:"uuid"`
	AnalyticsTunnelingEventCriticalThreshold types.Int64  `tfsdk:"analytics_tunneling_event_critical_threshold"`
	AnalyticsTunnelingEventWarningThreshold  types.Int64  `tfsdk:"analytics_tunneling_event_warning_threshold"`
	AtpCriticalEventCriticalThreshold        types.Int64  `tfsdk:"atp_critical_event_critical_threshold"`
	AtpCriticalEventWarningThreshold         types.Int64  `tfsdk:"atp_critical_event_warning_threshold"`
	AtpMajorEventCriticalThreshold           types.Int64  `tfsdk:"atp_major_event_critical_threshold"`
	AtpMajorEventWarningThreshold            types.Int64  `tfsdk:"atp_major_event_warning_threshold"`
	AtpWarningEventCriticalThreshold         types.Int64  `tfsdk:"atp_warning_event_critical_threshold"`
	AtpWarningEventWarningThreshold          types.Int64  `tfsdk:"atp_warning_event_warning_threshold"`
	RpzBlockedHitCriticalThreshold           types.Int64  `tfsdk:"rpz_blocked_hit_critical_threshold"`
	RpzBlockedHitWarningThreshold            types.Int64  `tfsdk:"rpz_blocked_hit_warning_threshold"`
	RpzPassthruEventCriticalThreshold        types.Int64  `tfsdk:"rpz_passthru_event_critical_threshold"`
	RpzPassthruEventWarningThreshold         types.Int64  `tfsdk:"rpz_passthru_event_warning_threshold"`
	RpzSubstitutedHitCriticalThreshold       types.Int64  `tfsdk:"rpz_substituted_hit_critical_threshold"`
	RpzSubstitutedHitWarningThreshold        types.Int64  `tfsdk:"rpz_substituted_hit_warning_threshold"`
}

var GridDashboardAttrTypes = map[string]attr.Type{
	"ref": types.StringType,
	"uuid": types.StringType,
	"analytics_tunneling_event_critical_threshold": types.Int64Type,
	"analytics_tunneling_event_warning_threshold":  types.Int64Type,
	"atp_critical_event_critical_threshold":        types.Int64Type,
	"atp_critical_event_warning_threshold":         types.Int64Type,
	"atp_major_event_critical_threshold":           types.Int64Type,
	"atp_major_event_warning_threshold":            types.Int64Type,
	"atp_warning_event_critical_threshold":         types.Int64Type,
	"atp_warning_event_warning_threshold":          types.Int64Type,
	"rpz_blocked_hit_critical_threshold":           types.Int64Type,
	"rpz_blocked_hit_warning_threshold":            types.Int64Type,
	"rpz_passthru_event_critical_threshold":        types.Int64Type,
	"rpz_passthru_event_warning_threshold":         types.Int64Type,
	"rpz_substituted_hit_critical_threshold":       types.Int64Type,
	"rpz_substituted_hit_warning_threshold":        types.Int64Type,
}

var GridDashboardResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The universally unique identifier (UUID) for the dashboard.",
	},
	"analytics_tunneling_event_critical_threshold": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The Grid Dashboard critical threshold for Analytics tunneling events.",
	},
	"analytics_tunneling_event_warning_threshold": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The Grid Dashboard warning threshold for Analytics tunneling events.",
	},
	"atp_critical_event_critical_threshold": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The Grid Dashboard critical threshold for ATP critical events.",
	},
	"atp_critical_event_warning_threshold": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The Grid Dashboard warning threshold for ATP critical events.",
	},
	"atp_major_event_critical_threshold": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The Grid Dashboard critical threshold for ATP major events.",
	},
	"atp_major_event_warning_threshold": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The Grid Dashboard warning threshold for ATP major events.",
	},
	"atp_warning_event_critical_threshold": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The Grid Dashboard critical threshold for ATP warning events.",
	},
	"atp_warning_event_warning_threshold": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The Grid Dashboard warning threshold for ATP warning events.",
	},
	"rpz_blocked_hit_critical_threshold": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The critical threshold value for blocked RPZ hits in the Grid dashboard.",
	},
	"rpz_blocked_hit_warning_threshold": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The warning threshold value for blocked RPZ hits in the Grid dashboard.",
	},
	"rpz_passthru_event_critical_threshold": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The Grid Dashboard critical threshold for RPZ passthru events.",
	},
	"rpz_passthru_event_warning_threshold": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The Grid Dashboard warning threshold for RPZ passthru events.",
	},
	"rpz_substituted_hit_critical_threshold": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The critical threshold value for substituted RPZ hits in the Grid dashboard.",
	},
	"rpz_substituted_hit_warning_threshold": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The warning threshold value for substituted RPZ hits in the Grid dashboard.",
	},
}

func ExpandGridDashboard(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridDashboard {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridDashboardModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridDashboardModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridDashboard {
	if m == nil {
		return nil
	}
	to := &grid.GridDashboard{
		Ref:                                      flex.ExpandStringPointer(m.Ref),
		Uuid:                                     flex.ExpandStringPointer(m.Uuid),
		AnalyticsTunnelingEventCriticalThreshold: flex.ExpandInt64Pointer(m.AnalyticsTunnelingEventCriticalThreshold),
		AnalyticsTunnelingEventWarningThreshold:  flex.ExpandInt64Pointer(m.AnalyticsTunnelingEventWarningThreshold),
		AtpCriticalEventCriticalThreshold:        flex.ExpandInt64Pointer(m.AtpCriticalEventCriticalThreshold),
		AtpCriticalEventWarningThreshold:         flex.ExpandInt64Pointer(m.AtpCriticalEventWarningThreshold),
		AtpMajorEventCriticalThreshold:           flex.ExpandInt64Pointer(m.AtpMajorEventCriticalThreshold),
		AtpMajorEventWarningThreshold:            flex.ExpandInt64Pointer(m.AtpMajorEventWarningThreshold),
		AtpWarningEventCriticalThreshold:         flex.ExpandInt64Pointer(m.AtpWarningEventCriticalThreshold),
		AtpWarningEventWarningThreshold:          flex.ExpandInt64Pointer(m.AtpWarningEventWarningThreshold),
		RpzBlockedHitCriticalThreshold:           flex.ExpandInt64Pointer(m.RpzBlockedHitCriticalThreshold),
		RpzBlockedHitWarningThreshold:            flex.ExpandInt64Pointer(m.RpzBlockedHitWarningThreshold),
		RpzPassthruEventCriticalThreshold:        flex.ExpandInt64Pointer(m.RpzPassthruEventCriticalThreshold),
		RpzPassthruEventWarningThreshold:         flex.ExpandInt64Pointer(m.RpzPassthruEventWarningThreshold),
		RpzSubstitutedHitCriticalThreshold:       flex.ExpandInt64Pointer(m.RpzSubstitutedHitCriticalThreshold),
		RpzSubstitutedHitWarningThreshold:        flex.ExpandInt64Pointer(m.RpzSubstitutedHitWarningThreshold),
	}
	return to
}

func FlattenGridDashboard(ctx context.Context, from *grid.GridDashboard, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridDashboardAttrTypes)
	}
	m := GridDashboardModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridDashboardAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridDashboardModel) Flatten(ctx context.Context, from *grid.GridDashboard, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridDashboardModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.AnalyticsTunnelingEventCriticalThreshold = flex.FlattenInt64Pointer(from.AnalyticsTunnelingEventCriticalThreshold)
	m.AnalyticsTunnelingEventWarningThreshold = flex.FlattenInt64Pointer(from.AnalyticsTunnelingEventWarningThreshold)
	m.AtpCriticalEventCriticalThreshold = flex.FlattenInt64Pointer(from.AtpCriticalEventCriticalThreshold)
	m.AtpCriticalEventWarningThreshold = flex.FlattenInt64Pointer(from.AtpCriticalEventWarningThreshold)
	m.AtpMajorEventCriticalThreshold = flex.FlattenInt64Pointer(from.AtpMajorEventCriticalThreshold)
	m.AtpMajorEventWarningThreshold = flex.FlattenInt64Pointer(from.AtpMajorEventWarningThreshold)
	m.AtpWarningEventCriticalThreshold = flex.FlattenInt64Pointer(from.AtpWarningEventCriticalThreshold)
	m.AtpWarningEventWarningThreshold = flex.FlattenInt64Pointer(from.AtpWarningEventWarningThreshold)
	m.RpzBlockedHitCriticalThreshold = flex.FlattenInt64Pointer(from.RpzBlockedHitCriticalThreshold)
	m.RpzBlockedHitWarningThreshold = flex.FlattenInt64Pointer(from.RpzBlockedHitWarningThreshold)
	m.RpzPassthruEventCriticalThreshold = flex.FlattenInt64Pointer(from.RpzPassthruEventCriticalThreshold)
	m.RpzPassthruEventWarningThreshold = flex.FlattenInt64Pointer(from.RpzPassthruEventWarningThreshold)
	m.RpzSubstitutedHitCriticalThreshold = flex.FlattenInt64Pointer(from.RpzSubstitutedHitCriticalThreshold)
	m.RpzSubstitutedHitWarningThreshold = flex.FlattenInt64Pointer(from.RpzSubstitutedHitWarningThreshold)
}
