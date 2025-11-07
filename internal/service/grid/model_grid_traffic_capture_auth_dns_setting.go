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

type GridTrafficCaptureAuthDnsSettingModel struct {
	AuthDnsLatencyTriggerEnable  types.Bool   `tfsdk:"auth_dns_latency_trigger_enable"`
	AuthDnsLatencyThreshold      types.Int64  `tfsdk:"auth_dns_latency_threshold"`
	AuthDnsLatencyReset          types.Int64  `tfsdk:"auth_dns_latency_reset"`
	AuthDnsLatencyListenOnSource types.String `tfsdk:"auth_dns_latency_listen_on_source"`
}

var GridTrafficCaptureAuthDnsSettingAttrTypes = map[string]attr.Type{
	"auth_dns_latency_trigger_enable":   types.BoolType,
	"auth_dns_latency_threshold":        types.Int64Type,
	"auth_dns_latency_reset":            types.Int64Type,
	"auth_dns_latency_listen_on_source": types.StringType,
}

var GridTrafficCaptureAuthDnsSettingResourceSchemaAttributes = map[string]schema.Attribute{
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
}

func ExpandGridTrafficCaptureAuthDnsSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridTrafficCaptureAuthDnsSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridTrafficCaptureAuthDnsSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridTrafficCaptureAuthDnsSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridTrafficCaptureAuthDnsSetting {
	if m == nil {
		return nil
	}
	to := &grid.GridTrafficCaptureAuthDnsSetting{
		AuthDnsLatencyTriggerEnable:  flex.ExpandBoolPointer(m.AuthDnsLatencyTriggerEnable),
		AuthDnsLatencyThreshold:      flex.ExpandInt64Pointer(m.AuthDnsLatencyThreshold),
		AuthDnsLatencyReset:          flex.ExpandInt64Pointer(m.AuthDnsLatencyReset),
		AuthDnsLatencyListenOnSource: flex.ExpandStringPointer(m.AuthDnsLatencyListenOnSource),
	}
	return to
}

func FlattenGridTrafficCaptureAuthDnsSetting(ctx context.Context, from *grid.GridTrafficCaptureAuthDnsSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridTrafficCaptureAuthDnsSettingAttrTypes)
	}
	m := GridTrafficCaptureAuthDnsSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridTrafficCaptureAuthDnsSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridTrafficCaptureAuthDnsSettingModel) Flatten(ctx context.Context, from *grid.GridTrafficCaptureAuthDnsSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridTrafficCaptureAuthDnsSettingModel{}
	}
	m.AuthDnsLatencyTriggerEnable = types.BoolPointerValue(from.AuthDnsLatencyTriggerEnable)
	m.AuthDnsLatencyThreshold = flex.FlattenInt64Pointer(from.AuthDnsLatencyThreshold)
	m.AuthDnsLatencyReset = flex.FlattenInt64Pointer(from.AuthDnsLatencyReset)
	m.AuthDnsLatencyListenOnSource = flex.FlattenStringPointer(from.AuthDnsLatencyListenOnSource)
}
