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

type MemberDnsDnsViewAddressSettingsModel struct {
	ViewName                       types.String `tfsdk:"view_name"`
	DnsNotifyTransferSource        types.String `tfsdk:"dns_notify_transfer_source"`
	DnsNotifyTransferSourceAddress types.String `tfsdk:"dns_notify_transfer_source_address"`
	DnsQuerySourceInterface        types.String `tfsdk:"dns_query_source_interface"`
	DnsQuerySourceAddress          types.String `tfsdk:"dns_query_source_address"`
	EnableNotifySourcePort         types.Bool   `tfsdk:"enable_notify_source_port"`
	NotifySourcePort               types.Int64  `tfsdk:"notify_source_port"`
	EnableQuerySourcePort          types.Bool   `tfsdk:"enable_query_source_port"`
	QuerySourcePort                types.Int64  `tfsdk:"query_source_port"`
	NotifyDelay                    types.Int64  `tfsdk:"notify_delay"`
	UseSourcePorts                 types.Bool   `tfsdk:"use_source_ports"`
	UseNotifyDelay                 types.Bool   `tfsdk:"use_notify_delay"`
}

var MemberDnsDnsViewAddressSettingsAttrTypes = map[string]attr.Type{
	"view_name":                          types.StringType,
	"dns_notify_transfer_source":         types.StringType,
	"dns_notify_transfer_source_address": types.StringType,
	"dns_query_source_interface":         types.StringType,
	"dns_query_source_address":           types.StringType,
	"enable_notify_source_port":          types.BoolType,
	"notify_source_port":                 types.Int64Type,
	"enable_query_source_port":           types.BoolType,
	"query_source_port":                  types.Int64Type,
	"notify_delay":                       types.Int64Type,
	"use_source_ports":                   types.BoolType,
	"use_notify_delay":                   types.BoolType,
}

var MemberDnsDnsViewAddressSettingsResourceSchemaAttributes = map[string]schema.Attribute{
	"view_name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to DNS View",
	},
	"dns_notify_transfer_source": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Determines which IP address is used as the source for DDNS notify and transfer operations.",
	},
	"dns_notify_transfer_source_address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The source address used if dns_notify_transfer_source type is \"IP\".",
	},
	"dns_query_source_interface": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Determines which IP address is used as the source for DDNS query operations.",
	},
	"dns_query_source_address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The source address used if dns_query_source_interface type is \"IP\".",
	},
	"enable_notify_source_port": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the notify source port for a view is enabled or not.",
	},
	"notify_source_port": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The source port for notify messages. When requesting zone transfers from the primary server, some secondary DNS servers use the source port number (the primary server used to send the notify message) as the destination port number in the zone transfer request. This setting overrides Grid static source port settings. Valid values are between 1 and 63999. The default is selected by BIND.",
	},
	"enable_query_source_port": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the query source port for a view is enabled or not.",
	},
	"query_source_port": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The source port for queries. Specifying a source port number for recursive queries ensures that a firewall will allow the response. Valid values are between 1 and 63999. The default is selected by BIND.",
	},
	"notify_delay": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Specifies the number of seconds of delay the notify messages are sent to secondaries.",
	},
	"use_source_ports": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: enable_notify_source_port , notify_source_port, enable_query_source_port, query_source_port",
	},
	"use_notify_delay": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: notify_delay",
	},
}

func ExpandMemberDnsDnsViewAddressSettings(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberDnsDnsViewAddressSettings {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberDnsDnsViewAddressSettingsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberDnsDnsViewAddressSettingsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberDnsDnsViewAddressSettings {
	if m == nil {
		return nil
	}
	to := &grid.MemberDnsDnsViewAddressSettings{
		ViewName:                       flex.ExpandStringPointer(m.ViewName),
		DnsNotifyTransferSource:        flex.ExpandStringPointer(m.DnsNotifyTransferSource),
		DnsNotifyTransferSourceAddress: flex.ExpandStringPointer(m.DnsNotifyTransferSourceAddress),
		DnsQuerySourceInterface:        flex.ExpandStringPointer(m.DnsQuerySourceInterface),
		DnsQuerySourceAddress:          flex.ExpandStringPointer(m.DnsQuerySourceAddress),
		EnableNotifySourcePort:         flex.ExpandBoolPointer(m.EnableNotifySourcePort),
		NotifySourcePort:               flex.ExpandInt64Pointer(m.NotifySourcePort),
		EnableQuerySourcePort:          flex.ExpandBoolPointer(m.EnableQuerySourcePort),
		QuerySourcePort:                flex.ExpandInt64Pointer(m.QuerySourcePort),
		NotifyDelay:                    flex.ExpandInt64Pointer(m.NotifyDelay),
		UseSourcePorts:                 flex.ExpandBoolPointer(m.UseSourcePorts),
		UseNotifyDelay:                 flex.ExpandBoolPointer(m.UseNotifyDelay),
	}
	return to
}

func FlattenMemberDnsDnsViewAddressSettings(ctx context.Context, from *grid.MemberDnsDnsViewAddressSettings, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberDnsDnsViewAddressSettingsAttrTypes)
	}
	m := MemberDnsDnsViewAddressSettingsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberDnsDnsViewAddressSettingsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberDnsDnsViewAddressSettingsModel) Flatten(ctx context.Context, from *grid.MemberDnsDnsViewAddressSettings, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberDnsDnsViewAddressSettingsModel{}
	}
	m.ViewName = flex.FlattenStringPointer(from.ViewName)
	m.DnsNotifyTransferSource = flex.FlattenStringPointer(from.DnsNotifyTransferSource)
	m.DnsNotifyTransferSourceAddress = flex.FlattenStringPointer(from.DnsNotifyTransferSourceAddress)
	m.DnsQuerySourceInterface = flex.FlattenStringPointer(from.DnsQuerySourceInterface)
	m.DnsQuerySourceAddress = flex.FlattenStringPointer(from.DnsQuerySourceAddress)
	m.EnableNotifySourcePort = types.BoolPointerValue(from.EnableNotifySourcePort)
	m.NotifySourcePort = flex.FlattenInt64Pointer(from.NotifySourcePort)
	m.EnableQuerySourcePort = types.BoolPointerValue(from.EnableQuerySourcePort)
	m.QuerySourcePort = flex.FlattenInt64Pointer(from.QuerySourcePort)
	m.NotifyDelay = flex.FlattenInt64Pointer(from.NotifyDelay)
	m.UseSourcePorts = types.BoolPointerValue(from.UseSourcePorts)
	m.UseNotifyDelay = types.BoolPointerValue(from.UseNotifyDelay)
}
