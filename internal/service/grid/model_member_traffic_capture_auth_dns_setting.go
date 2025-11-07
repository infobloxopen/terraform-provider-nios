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

type MemberTrafficCaptureAuthDnsSettingModel struct {
	AuthDnsLatencyTriggerEnable  types.Bool   `tfsdk:"auth_dns_latency_trigger_enable"`
	AuthDnsLatencyThreshold      types.Int64  `tfsdk:"auth_dns_latency_threshold"`
	AuthDnsLatencyReset          types.Int64  `tfsdk:"auth_dns_latency_reset"`
	AuthDnsLatencyListenOnSource types.String `tfsdk:"auth_dns_latency_listen_on_source"`
	AuthDnsLatencyListenOnIp     types.String `tfsdk:"auth_dns_latency_listen_on_ip"`
}

var MemberTrafficCaptureAuthDnsSettingAttrTypes = map[string]attr.Type{
	"auth_dns_latency_trigger_enable":   types.BoolType,
	"auth_dns_latency_threshold":        types.Int64Type,
	"auth_dns_latency_reset":            types.Int64Type,
	"auth_dns_latency_listen_on_source": types.StringType,
	"auth_dns_latency_listen_on_ip":     types.StringType,
}

var MemberTrafficCaptureAuthDnsSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"auth_dns_latency_trigger_enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enabling trigger automated traffic capture based on authoritative DNS latency.",
	},
	"auth_dns_latency_threshold": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Authoritative DNS latency below which traffic capture will be triggered.",
	},
	"auth_dns_latency_reset": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Authoritative DNS latency above which traffic capture will stopped.",
	},
	"auth_dns_latency_listen_on_source": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The local IP DNS service is listen on (for authoritative DNS latency trigger).",
	},
	"auth_dns_latency_listen_on_ip": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The DNS listen-on IP address used if auth_dns_latency_on_source is IP.",
	},
}

func ExpandMemberTrafficCaptureAuthDnsSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberTrafficCaptureAuthDnsSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberTrafficCaptureAuthDnsSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberTrafficCaptureAuthDnsSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberTrafficCaptureAuthDnsSetting {
	if m == nil {
		return nil
	}
	to := &grid.MemberTrafficCaptureAuthDnsSetting{
		AuthDnsLatencyTriggerEnable:  flex.ExpandBoolPointer(m.AuthDnsLatencyTriggerEnable),
		AuthDnsLatencyThreshold:      flex.ExpandInt64Pointer(m.AuthDnsLatencyThreshold),
		AuthDnsLatencyReset:          flex.ExpandInt64Pointer(m.AuthDnsLatencyReset),
		AuthDnsLatencyListenOnSource: flex.ExpandStringPointer(m.AuthDnsLatencyListenOnSource),
		AuthDnsLatencyListenOnIp:     flex.ExpandStringPointer(m.AuthDnsLatencyListenOnIp),
	}
	return to
}

func FlattenMemberTrafficCaptureAuthDnsSetting(ctx context.Context, from *grid.MemberTrafficCaptureAuthDnsSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberTrafficCaptureAuthDnsSettingAttrTypes)
	}
	m := MemberTrafficCaptureAuthDnsSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberTrafficCaptureAuthDnsSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberTrafficCaptureAuthDnsSettingModel) Flatten(ctx context.Context, from *grid.MemberTrafficCaptureAuthDnsSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberTrafficCaptureAuthDnsSettingModel{}
	}
	m.AuthDnsLatencyTriggerEnable = types.BoolPointerValue(from.AuthDnsLatencyTriggerEnable)
	m.AuthDnsLatencyThreshold = flex.FlattenInt64Pointer(from.AuthDnsLatencyThreshold)
	m.AuthDnsLatencyReset = flex.FlattenInt64Pointer(from.AuthDnsLatencyReset)
	m.AuthDnsLatencyListenOnSource = flex.FlattenStringPointer(from.AuthDnsLatencyListenOnSource)
	m.AuthDnsLatencyListenOnIp = flex.FlattenStringPointer(from.AuthDnsLatencyListenOnIp)
}
