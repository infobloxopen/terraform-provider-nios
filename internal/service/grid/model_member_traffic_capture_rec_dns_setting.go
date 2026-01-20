package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type MemberTrafficCaptureRecDnsSettingModel struct {
	RecDnsLatencyTriggerEnable  types.Bool   `tfsdk:"rec_dns_latency_trigger_enable"`
	RecDnsLatencyThreshold      types.Int64  `tfsdk:"rec_dns_latency_threshold"`
	RecDnsLatencyReset          types.Int64  `tfsdk:"rec_dns_latency_reset"`
	RecDnsLatencyListenOnSource types.String `tfsdk:"rec_dns_latency_listen_on_source"`
	RecDnsLatencyListenOnIp     types.String `tfsdk:"rec_dns_latency_listen_on_ip"`
	KpiMonitoredDomains         types.List   `tfsdk:"kpi_monitored_domains"`
}

var MemberTrafficCaptureRecDnsSettingAttrTypes = map[string]attr.Type{
	"rec_dns_latency_trigger_enable":   types.BoolType,
	"rec_dns_latency_threshold":        types.Int64Type,
	"rec_dns_latency_reset":            types.Int64Type,
	"rec_dns_latency_listen_on_source": types.StringType,
	"rec_dns_latency_listen_on_ip":     types.StringType,
	"kpi_monitored_domains":            types.ListType{ElemType: types.ObjectType{AttrTypes: MembertrafficcapturerecdnssettingKpiMonitoredDomainsAttrTypes}},
}

var MemberTrafficCaptureRecDnsSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"rec_dns_latency_trigger_enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enable triggering automated traffic capture based on recursive DNS latency.",
	},
	"rec_dns_latency_threshold": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Recursive DNS latency below which traffic capture will be triggered.",
	},
	"rec_dns_latency_reset": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Recursive DNS latency above which traffic capture will be stopped.",
	},
	"rec_dns_latency_listen_on_source": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The local IP DNS service is listen on ( for recursive DNS latency trigger).",
	},
	"rec_dns_latency_listen_on_ip": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The DNS listen-on IP address used if rec_dns_latency_listen_on_source is IP.",
	},
	"kpi_monitored_domains": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MembertrafficcapturerecdnssettingKpiMonitoredDomainsResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "List of domains monitored by 'Recursive DNS Latency Threshold' trigger.",
	},
}

func ExpandMemberTrafficCaptureRecDnsSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberTrafficCaptureRecDnsSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberTrafficCaptureRecDnsSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberTrafficCaptureRecDnsSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberTrafficCaptureRecDnsSetting {
	if m == nil {
		return nil
	}
	to := &grid.MemberTrafficCaptureRecDnsSetting{
		RecDnsLatencyTriggerEnable:  flex.ExpandBoolPointer(m.RecDnsLatencyTriggerEnable),
		RecDnsLatencyThreshold:      flex.ExpandInt64Pointer(m.RecDnsLatencyThreshold),
		RecDnsLatencyReset:          flex.ExpandInt64Pointer(m.RecDnsLatencyReset),
		RecDnsLatencyListenOnSource: flex.ExpandStringPointer(m.RecDnsLatencyListenOnSource),
		RecDnsLatencyListenOnIp:     flex.ExpandStringPointer(m.RecDnsLatencyListenOnIp),
		KpiMonitoredDomains:         flex.ExpandFrameworkListNestedBlock(ctx, m.KpiMonitoredDomains, diags, ExpandMembertrafficcapturerecdnssettingKpiMonitoredDomains),
	}
	return to
}

func FlattenMemberTrafficCaptureRecDnsSetting(ctx context.Context, from *grid.MemberTrafficCaptureRecDnsSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberTrafficCaptureRecDnsSettingAttrTypes)
	}
	m := MemberTrafficCaptureRecDnsSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberTrafficCaptureRecDnsSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberTrafficCaptureRecDnsSettingModel) Flatten(ctx context.Context, from *grid.MemberTrafficCaptureRecDnsSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberTrafficCaptureRecDnsSettingModel{}
	}
	m.RecDnsLatencyTriggerEnable = types.BoolPointerValue(from.RecDnsLatencyTriggerEnable)
	m.RecDnsLatencyThreshold = flex.FlattenInt64Pointer(from.RecDnsLatencyThreshold)
	m.RecDnsLatencyReset = flex.FlattenInt64Pointer(from.RecDnsLatencyReset)
	m.RecDnsLatencyListenOnSource = flex.FlattenStringPointer(from.RecDnsLatencyListenOnSource)
	m.RecDnsLatencyListenOnIp = flex.FlattenStringPointer(from.RecDnsLatencyListenOnIp)
	m.KpiMonitoredDomains = flex.FlattenFrameworkListNestedBlock(ctx, from.KpiMonitoredDomains, MembertrafficcapturerecdnssettingKpiMonitoredDomainsAttrTypes, diags, FlattenMembertrafficcapturerecdnssettingKpiMonitoredDomains)
}
