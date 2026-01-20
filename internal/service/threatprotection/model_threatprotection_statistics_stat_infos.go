package threatprotection

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/threatprotection"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type ThreatprotectionStatisticsStatInfosModel struct {
	Timestamp     types.Int64 `tfsdk:"timestamp"`
	Critical      types.Map   `tfsdk:"critical"`
	Major         types.Map   `tfsdk:"major"`
	Warning       types.Map   `tfsdk:"warning"`
	Informational types.Map   `tfsdk:"informational"`
	Total         types.Map   `tfsdk:"total"`
}

var ThreatprotectionStatisticsStatInfosAttrTypes = map[string]attr.Type{
	"timestamp":     types.Int64Type,
	"critical":      types.MapType{ElemType: types.StringType},
	"major":         types.MapType{ElemType: types.StringType},
	"warning":       types.MapType{ElemType: types.StringType},
	"informational": types.MapType{ElemType: types.StringType},
	"total":         types.MapType{ElemType: types.StringType},
}

var ThreatprotectionStatisticsStatInfosResourceSchemaAttributes = map[string]schema.Attribute{
	"timestamp": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The timestamp when data was collected.",
	},
	"critical": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The number of critical events.",
	},
	"major": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The number of major events.",
	},
	"warning": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The number of warning events.",
	},
	"informational": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The number of informational events.",
	},
	"total": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The total number of events.",
	},
}

func ExpandThreatprotectionStatisticsStatInfos(ctx context.Context, o types.Object, diags *diag.Diagnostics) *threatprotection.ThreatprotectionStatisticsStatInfos {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ThreatprotectionStatisticsStatInfosModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ThreatprotectionStatisticsStatInfosModel) Expand(ctx context.Context, diags *diag.Diagnostics) *threatprotection.ThreatprotectionStatisticsStatInfos {
	if m == nil {
		return nil
	}
	to := &threatprotection.ThreatprotectionStatisticsStatInfos{
		Critical:      flex.ExpandFrameworkMapString(ctx, m.Critical, diags),
		Major:         flex.ExpandFrameworkMapString(ctx, m.Major, diags),
		Warning:       flex.ExpandFrameworkMapString(ctx, m.Warning, diags),
		Informational: flex.ExpandFrameworkMapString(ctx, m.Informational, diags),
		Total:         flex.ExpandFrameworkMapString(ctx, m.Total, diags),
	}
	return to
}

func FlattenThreatprotectionStatisticsStatInfos(ctx context.Context, from *threatprotection.ThreatprotectionStatisticsStatInfos, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ThreatprotectionStatisticsStatInfosAttrTypes)
	}
	m := ThreatprotectionStatisticsStatInfosModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ThreatprotectionStatisticsStatInfosAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ThreatprotectionStatisticsStatInfosModel) Flatten(ctx context.Context, from *threatprotection.ThreatprotectionStatisticsStatInfos, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ThreatprotectionStatisticsStatInfosModel{}
	}
	m.Timestamp = flex.FlattenInt64Pointer(from.Timestamp)
	m.Critical = flex.FlattenFrameworkMapString(ctx, from.Critical, diags)
	m.Major = flex.FlattenFrameworkMapString(ctx, from.Major, diags)
	m.Warning = flex.FlattenFrameworkMapString(ctx, from.Warning, diags)
	m.Informational = flex.FlattenFrameworkMapString(ctx, from.Informational, diags)
	m.Total = flex.FlattenFrameworkMapString(ctx, from.Total, diags)
}
