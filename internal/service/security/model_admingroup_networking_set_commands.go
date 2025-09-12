package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type AdmingroupNetworkingSetCommandsModel struct {
	SetConnectionLimit      types.Bool `tfsdk:"set_connection_limit"`
	SetDefaultRoute         types.Bool `tfsdk:"set_default_route"`
	SetInterface            types.Bool `tfsdk:"set_interface"`
	SetIpRateLimit          types.Bool `tfsdk:"set_ip_rate_limit"`
	SetIpv6DisableOnDad     types.Bool `tfsdk:"set_ipv6_disable_on_dad"`
	SetIpv6Neighbor         types.Bool `tfsdk:"set_ipv6_neighbor"`
	SetIpv6Ospf             types.Bool `tfsdk:"set_ipv6_ospf"`
	SetIpv6Status           types.Bool `tfsdk:"set_ipv6_status"`
	SetLom                  types.Bool `tfsdk:"set_lom"`
	SetMldVersion1          types.Bool `tfsdk:"set_mld_version_1"`
	SetNamedRecvSockBufSize types.Bool `tfsdk:"set_named_recv_sock_buf_size"`
	SetNamedTcpClientsLimit types.Bool `tfsdk:"set_named_tcp_clients_limit"`
	SetNetwork              types.Bool `tfsdk:"set_network"`
	SetOspf                 types.Bool `tfsdk:"set_ospf"`
	SetPrompt               types.Bool `tfsdk:"set_prompt"`
	SetRemoteConsole        types.Bool `tfsdk:"set_remote_console"`
	SetStaticRoute          types.Bool `tfsdk:"set_static_route"`
	SetTcpTimestamps        types.Bool `tfsdk:"set_tcp_timestamps"`
	SetTrafficCapture       types.Bool `tfsdk:"set_traffic_capture"`
	SetWinsForwarding       types.Bool `tfsdk:"set_wins_forwarding"`
	EnableAll               types.Bool `tfsdk:"enable_all"`
	DisableAll              types.Bool `tfsdk:"disable_all"`
}

var AdmingroupNetworkingSetCommandsAttrTypes = map[string]attr.Type{
	"set_connection_limit":         types.BoolType,
	"set_default_route":            types.BoolType,
	"set_interface":                types.BoolType,
	"set_ip_rate_limit":            types.BoolType,
	"set_ipv6_disable_on_dad":      types.BoolType,
	"set_ipv6_neighbor":            types.BoolType,
	"set_ipv6_ospf":                types.BoolType,
	"set_ipv6_status":              types.BoolType,
	"set_lom":                      types.BoolType,
	"set_mld_version_1":            types.BoolType,
	"set_named_recv_sock_buf_size": types.BoolType,
	"set_named_tcp_clients_limit":  types.BoolType,
	"set_network":                  types.BoolType,
	"set_ospf":                     types.BoolType,
	"set_prompt":                   types.BoolType,
	"set_remote_console":           types.BoolType,
	"set_static_route":             types.BoolType,
	"set_tcp_timestamps":           types.BoolType,
	"set_traffic_capture":          types.BoolType,
	"set_wins_forwarding":          types.BoolType,
	"enable_all":                   types.BoolType,
	"disable_all":                  types.BoolType,
}

var AdmingroupNetworkingSetCommandsResourceSchemaAttributes = map[string]schema.Attribute{
	"set_connection_limit": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_default_route": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_interface": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_ip_rate_limit": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_ipv6_disable_on_dad": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_ipv6_neighbor": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_ipv6_ospf": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_ipv6_status": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_lom": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_mld_version_1": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_named_recv_sock_buf_size": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_named_tcp_clients_limit": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_network": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_ospf": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_prompt": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_remote_console": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_static_route": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_tcp_timestamps": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_traffic_capture": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_wins_forwarding": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"enable_all": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then enable all fields",
	},
	"disable_all": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then disable all fields",
	},
}

func ExpandAdmingroupNetworkingSetCommands(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.AdmingroupNetworkingSetCommands {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AdmingroupNetworkingSetCommandsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AdmingroupNetworkingSetCommandsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.AdmingroupNetworkingSetCommands {
	if m == nil {
		return nil
	}
	to := &security.AdmingroupNetworkingSetCommands{
		SetConnectionLimit:      flex.ExpandBoolPointer(m.SetConnectionLimit),
		SetDefaultRoute:         flex.ExpandBoolPointer(m.SetDefaultRoute),
		SetInterface:            flex.ExpandBoolPointer(m.SetInterface),
		SetIpRateLimit:          flex.ExpandBoolPointer(m.SetIpRateLimit),
		SetIpv6DisableOnDad:     flex.ExpandBoolPointer(m.SetIpv6DisableOnDad),
		SetIpv6Neighbor:         flex.ExpandBoolPointer(m.SetIpv6Neighbor),
		SetIpv6Ospf:             flex.ExpandBoolPointer(m.SetIpv6Ospf),
		SetIpv6Status:           flex.ExpandBoolPointer(m.SetIpv6Status),
		SetLom:                  flex.ExpandBoolPointer(m.SetLom),
		SetMldVersion1:          flex.ExpandBoolPointer(m.SetMldVersion1),
		SetNamedRecvSockBufSize: flex.ExpandBoolPointer(m.SetNamedRecvSockBufSize),
		SetNamedTcpClientsLimit: flex.ExpandBoolPointer(m.SetNamedTcpClientsLimit),
		SetNetwork:              flex.ExpandBoolPointer(m.SetNetwork),
		SetOspf:                 flex.ExpandBoolPointer(m.SetOspf),
		SetPrompt:               flex.ExpandBoolPointer(m.SetPrompt),
		SetRemoteConsole:        flex.ExpandBoolPointer(m.SetRemoteConsole),
		SetStaticRoute:          flex.ExpandBoolPointer(m.SetStaticRoute),
		SetTcpTimestamps:        flex.ExpandBoolPointer(m.SetTcpTimestamps),
		SetTrafficCapture:       flex.ExpandBoolPointer(m.SetTrafficCapture),
		SetWinsForwarding:       flex.ExpandBoolPointer(m.SetWinsForwarding),
		EnableAll:               flex.ExpandBoolPointer(m.EnableAll),
		DisableAll:              flex.ExpandBoolPointer(m.DisableAll),
	}
	return to
}

func FlattenAdmingroupNetworkingSetCommands(ctx context.Context, from *security.AdmingroupNetworkingSetCommands, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AdmingroupNetworkingSetCommandsAttrTypes)
	}
	m := AdmingroupNetworkingSetCommandsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AdmingroupNetworkingSetCommandsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AdmingroupNetworkingSetCommandsModel) Flatten(ctx context.Context, from *security.AdmingroupNetworkingSetCommands, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AdmingroupNetworkingSetCommandsModel{}
	}
	m.SetConnectionLimit = types.BoolPointerValue(from.SetConnectionLimit)
	m.SetDefaultRoute = types.BoolPointerValue(from.SetDefaultRoute)
	m.SetInterface = types.BoolPointerValue(from.SetInterface)
	m.SetIpRateLimit = types.BoolPointerValue(from.SetIpRateLimit)
	m.SetIpv6DisableOnDad = types.BoolPointerValue(from.SetIpv6DisableOnDad)
	m.SetIpv6Neighbor = types.BoolPointerValue(from.SetIpv6Neighbor)
	m.SetIpv6Ospf = types.BoolPointerValue(from.SetIpv6Ospf)
	m.SetIpv6Status = types.BoolPointerValue(from.SetIpv6Status)
	m.SetLom = types.BoolPointerValue(from.SetLom)
	m.SetMldVersion1 = types.BoolPointerValue(from.SetMldVersion1)
	m.SetNamedRecvSockBufSize = types.BoolPointerValue(from.SetNamedRecvSockBufSize)
	m.SetNamedTcpClientsLimit = types.BoolPointerValue(from.SetNamedTcpClientsLimit)
	m.SetNetwork = types.BoolPointerValue(from.SetNetwork)
	m.SetOspf = types.BoolPointerValue(from.SetOspf)
	m.SetPrompt = types.BoolPointerValue(from.SetPrompt)
	m.SetRemoteConsole = types.BoolPointerValue(from.SetRemoteConsole)
	m.SetStaticRoute = types.BoolPointerValue(from.SetStaticRoute)
	m.SetTcpTimestamps = types.BoolPointerValue(from.SetTcpTimestamps)
	m.SetTrafficCapture = types.BoolPointerValue(from.SetTrafficCapture)
	m.SetWinsForwarding = types.BoolPointerValue(from.SetWinsForwarding)
	m.EnableAll = types.BoolPointerValue(from.EnableAll)
	m.DisableAll = types.BoolPointerValue(from.DisableAll)
}
