package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type AdmingroupNetworkingShowCommandsModel struct {
	ShowConnectionLimit      types.Bool `tfsdk:"show_connection_limit"`
	ShowConnections          types.Bool `tfsdk:"show_connections"`
	ShowInterface            types.Bool `tfsdk:"show_interface"`
	ShowIpRateLimit          types.Bool `tfsdk:"show_ip_rate_limit"`
	ShowIpv6Bgp              types.Bool `tfsdk:"show_ipv6_bgp"`
	ShowIpv6DisableOnDad     types.Bool `tfsdk:"show_ipv6_disable_on_dad"`
	ShowIpv6Neighbor         types.Bool `tfsdk:"show_ipv6_neighbor"`
	ShowIpv6Ospf             types.Bool `tfsdk:"show_ipv6_ospf"`
	ShowLom                  types.Bool `tfsdk:"show_lom"`
	ShowMldVersion           types.Bool `tfsdk:"show_mld_version"`
	ShowNamedRecvSockBufSize types.Bool `tfsdk:"show_named_recv_sock_buf_size"`
	ShowNamedTcpClientsLimit types.Bool `tfsdk:"show_named_tcp_clients_limit"`
	ShowNetwork              types.Bool `tfsdk:"show_network"`
	ShowOspf                 types.Bool `tfsdk:"show_ospf"`
	ShowRemoteConsole        types.Bool `tfsdk:"show_remote_console"`
	ShowRoutes               types.Bool `tfsdk:"show_routes"`
	ShowStaticRoutes         types.Bool `tfsdk:"show_static_routes"`
	ShowTcpTimestamps        types.Bool `tfsdk:"show_tcp_timestamps"`
	ShowTrafficCaptureStatus types.Bool `tfsdk:"show_traffic_capture_status"`
	ShowWinsForwarding       types.Bool `tfsdk:"show_wins_forwarding"`
	ShowDefaultRoute         types.Bool `tfsdk:"show_default_route"`
	ShowIproute              types.Bool `tfsdk:"show_iproute"`
	ShowIprule               types.Bool `tfsdk:"show_iprule"`
	ShowIptables             types.Bool `tfsdk:"show_iptables"`
	ShowMtuSize              types.Bool `tfsdk:"show_mtu_size"`
	ShowNetworkConnectivity  types.Bool `tfsdk:"show_network_connectivity"`
	ShowTrafficfiles         types.Bool `tfsdk:"show_trafficfiles"`
	ShowInterfaceStats       types.Bool `tfsdk:"show_interface_stats"`
	EnableAll                types.Bool `tfsdk:"enable_all"`
	DisableAll               types.Bool `tfsdk:"disable_all"`
}

var AdmingroupNetworkingShowCommandsAttrTypes = map[string]attr.Type{
	"show_connection_limit":         types.BoolType,
	"show_connections":              types.BoolType,
	"show_interface":                types.BoolType,
	"show_ip_rate_limit":            types.BoolType,
	"show_ipv6_bgp":                 types.BoolType,
	"show_ipv6_disable_on_dad":      types.BoolType,
	"show_ipv6_neighbor":            types.BoolType,
	"show_ipv6_ospf":                types.BoolType,
	"show_lom":                      types.BoolType,
	"show_mld_version":              types.BoolType,
	"show_named_recv_sock_buf_size": types.BoolType,
	"show_named_tcp_clients_limit":  types.BoolType,
	"show_network":                  types.BoolType,
	"show_ospf":                     types.BoolType,
	"show_remote_console":           types.BoolType,
	"show_routes":                   types.BoolType,
	"show_static_routes":            types.BoolType,
	"show_tcp_timestamps":           types.BoolType,
	"show_traffic_capture_status":   types.BoolType,
	"show_wins_forwarding":          types.BoolType,
	"show_default_route":            types.BoolType,
	"show_iproute":                  types.BoolType,
	"show_iprule":                   types.BoolType,
	"show_iptables":                 types.BoolType,
	"show_mtu_size":                 types.BoolType,
	"show_network_connectivity":     types.BoolType,
	"show_trafficfiles":             types.BoolType,
	"show_interface_stats":          types.BoolType,
	"enable_all":                    types.BoolType,
	"disable_all":                   types.BoolType,
}

var AdmingroupNetworkingShowCommandsResourceSchemaAttributes = map[string]schema.Attribute{
	"show_connection_limit": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_connections": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_interface": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_ip_rate_limit": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_ipv6_bgp": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_ipv6_disable_on_dad": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_ipv6_neighbor": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_ipv6_ospf": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_lom": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_mld_version": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_named_recv_sock_buf_size": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_named_tcp_clients_limit": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_network": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_ospf": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_remote_console": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_routes": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_static_routes": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_tcp_timestamps": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_traffic_capture_status": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_wins_forwarding": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_default_route": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_iproute": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_iprule": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_iptables": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_mtu_size": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_network_connectivity": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_trafficfiles": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_interface_stats": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"enable_all": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then enable all fields",
	},
	"disable_all": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then disable all fields",
	},
}

func ExpandAdmingroupNetworkingShowCommands(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.AdmingroupNetworkingShowCommands {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AdmingroupNetworkingShowCommandsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AdmingroupNetworkingShowCommandsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.AdmingroupNetworkingShowCommands {
	if m == nil {
		return nil
	}
	to := &security.AdmingroupNetworkingShowCommands{
		ShowConnectionLimit:      flex.ExpandBoolPointer(m.ShowConnectionLimit),
		ShowConnections:          flex.ExpandBoolPointer(m.ShowConnections),
		ShowInterface:            flex.ExpandBoolPointer(m.ShowInterface),
		ShowIpRateLimit:          flex.ExpandBoolPointer(m.ShowIpRateLimit),
		ShowIpv6Bgp:              flex.ExpandBoolPointer(m.ShowIpv6Bgp),
		ShowIpv6DisableOnDad:     flex.ExpandBoolPointer(m.ShowIpv6DisableOnDad),
		ShowIpv6Neighbor:         flex.ExpandBoolPointer(m.ShowIpv6Neighbor),
		ShowIpv6Ospf:             flex.ExpandBoolPointer(m.ShowIpv6Ospf),
		ShowLom:                  flex.ExpandBoolPointer(m.ShowLom),
		ShowMldVersion:           flex.ExpandBoolPointer(m.ShowMldVersion),
		ShowNamedRecvSockBufSize: flex.ExpandBoolPointer(m.ShowNamedRecvSockBufSize),
		ShowNamedTcpClientsLimit: flex.ExpandBoolPointer(m.ShowNamedTcpClientsLimit),
		ShowNetwork:              flex.ExpandBoolPointer(m.ShowNetwork),
		ShowOspf:                 flex.ExpandBoolPointer(m.ShowOspf),
		ShowRemoteConsole:        flex.ExpandBoolPointer(m.ShowRemoteConsole),
		ShowRoutes:               flex.ExpandBoolPointer(m.ShowRoutes),
		ShowStaticRoutes:         flex.ExpandBoolPointer(m.ShowStaticRoutes),
		ShowTcpTimestamps:        flex.ExpandBoolPointer(m.ShowTcpTimestamps),
		ShowTrafficCaptureStatus: flex.ExpandBoolPointer(m.ShowTrafficCaptureStatus),
		ShowWinsForwarding:       flex.ExpandBoolPointer(m.ShowWinsForwarding),
		ShowDefaultRoute:         flex.ExpandBoolPointer(m.ShowDefaultRoute),
		ShowIproute:              flex.ExpandBoolPointer(m.ShowIproute),
		ShowIprule:               flex.ExpandBoolPointer(m.ShowIprule),
		ShowIptables:             flex.ExpandBoolPointer(m.ShowIptables),
		ShowMtuSize:              flex.ExpandBoolPointer(m.ShowMtuSize),
		ShowNetworkConnectivity:  flex.ExpandBoolPointer(m.ShowNetworkConnectivity),
		ShowTrafficfiles:         flex.ExpandBoolPointer(m.ShowTrafficfiles),
		ShowInterfaceStats:       flex.ExpandBoolPointer(m.ShowInterfaceStats),
		EnableAll:                flex.ExpandBoolPointer(m.EnableAll),
		DisableAll:               flex.ExpandBoolPointer(m.DisableAll),
	}
	return to
}

func FlattenAdmingroupNetworkingShowCommands(ctx context.Context, from *security.AdmingroupNetworkingShowCommands, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AdmingroupNetworkingShowCommandsAttrTypes)
	}
	m := AdmingroupNetworkingShowCommandsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AdmingroupNetworkingShowCommandsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AdmingroupNetworkingShowCommandsModel) Flatten(ctx context.Context, from *security.AdmingroupNetworkingShowCommands, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AdmingroupNetworkingShowCommandsModel{}
	}
	m.ShowConnectionLimit = types.BoolPointerValue(from.ShowConnectionLimit)
	m.ShowConnections = types.BoolPointerValue(from.ShowConnections)
	m.ShowInterface = types.BoolPointerValue(from.ShowInterface)
	m.ShowIpRateLimit = types.BoolPointerValue(from.ShowIpRateLimit)
	m.ShowIpv6Bgp = types.BoolPointerValue(from.ShowIpv6Bgp)
	m.ShowIpv6DisableOnDad = types.BoolPointerValue(from.ShowIpv6DisableOnDad)
	m.ShowIpv6Neighbor = types.BoolPointerValue(from.ShowIpv6Neighbor)
	m.ShowIpv6Ospf = types.BoolPointerValue(from.ShowIpv6Ospf)
	m.ShowLom = types.BoolPointerValue(from.ShowLom)
	m.ShowMldVersion = types.BoolPointerValue(from.ShowMldVersion)
	m.ShowNamedRecvSockBufSize = types.BoolPointerValue(from.ShowNamedRecvSockBufSize)
	m.ShowNamedTcpClientsLimit = types.BoolPointerValue(from.ShowNamedTcpClientsLimit)
	m.ShowNetwork = types.BoolPointerValue(from.ShowNetwork)
	m.ShowOspf = types.BoolPointerValue(from.ShowOspf)
	m.ShowRemoteConsole = types.BoolPointerValue(from.ShowRemoteConsole)
	m.ShowRoutes = types.BoolPointerValue(from.ShowRoutes)
	m.ShowStaticRoutes = types.BoolPointerValue(from.ShowStaticRoutes)
	m.ShowTcpTimestamps = types.BoolPointerValue(from.ShowTcpTimestamps)
	m.ShowTrafficCaptureStatus = types.BoolPointerValue(from.ShowTrafficCaptureStatus)
	m.ShowWinsForwarding = types.BoolPointerValue(from.ShowWinsForwarding)
	m.ShowDefaultRoute = types.BoolPointerValue(from.ShowDefaultRoute)
	m.ShowIproute = types.BoolPointerValue(from.ShowIproute)
	m.ShowIprule = types.BoolPointerValue(from.ShowIprule)
	m.ShowIptables = types.BoolPointerValue(from.ShowIptables)
	m.ShowMtuSize = types.BoolPointerValue(from.ShowMtuSize)
	m.ShowNetworkConnectivity = types.BoolPointerValue(from.ShowNetworkConnectivity)
	m.ShowTrafficfiles = types.BoolPointerValue(from.ShowTrafficfiles)
	m.ShowInterfaceStats = types.BoolPointerValue(from.ShowInterfaceStats)
	m.EnableAll = types.BoolPointerValue(from.EnableAll)
	m.DisableAll = types.BoolPointerValue(from.DisableAll)
}
