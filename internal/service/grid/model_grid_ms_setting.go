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

type GridMsSettingModel struct {
	LogDestination       types.String `tfsdk:"log_destination"`
	EnableInvalidMac     types.Bool   `tfsdk:"enable_invalid_mac"`
	MaxConnection        types.Int64  `tfsdk:"max_connection"`
	RpcTimeout           types.Int64  `tfsdk:"rpc_timeout"`
	EnableDhcpMonitoring types.Bool   `tfsdk:"enable_dhcp_monitoring"`
	EnableDnsMonitoring  types.Bool   `tfsdk:"enable_dns_monitoring"`
	LdapTimeout          types.Int64  `tfsdk:"ldap_timeout"`
	DefaultIpSiteLink    types.String `tfsdk:"default_ip_site_link"`
	EnableNetworkUsers   types.Bool   `tfsdk:"enable_network_users"`
	EnableAdUserSync     types.Bool   `tfsdk:"enable_ad_user_sync"`
	AdUserDefaultTimeout types.Int64  `tfsdk:"ad_user_default_timeout"`
	EnableDnsReportsSync types.Bool   `tfsdk:"enable_dns_reports_sync"`
}

var GridMsSettingAttrTypes = map[string]attr.Type{
	"log_destination":         types.StringType,
	"enable_invalid_mac":      types.BoolType,
	"max_connection":          types.Int64Type,
	"rpc_timeout":             types.Int64Type,
	"enable_dhcp_monitoring":  types.BoolType,
	"enable_dns_monitoring":   types.BoolType,
	"ldap_timeout":            types.Int64Type,
	"default_ip_site_link":    types.StringType,
	"enable_network_users":    types.BoolType,
	"enable_ad_user_sync":     types.BoolType,
	"ad_user_default_timeout": types.Int64Type,
	"enable_dns_reports_sync": types.BoolType,
}

var GridMsSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"log_destination": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The logging of synchronization messages to the syslog or mslog.",
	},
	"enable_invalid_mac": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the invalid MAC address synchronization for DHCP leases and fixed addresses is enabled or not.",
	},
	"max_connection": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Determines the maximum number of connections to Microsoft servers.",
	},
	"rpc_timeout": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Determines the timeout value (in seconds) for RPC connections to all Microsoft servers.",
	},
	"enable_dhcp_monitoring": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the monitoring and control of DHCP service on all Microsoft servers in the Grid is enabled or not.",
	},
	"enable_dns_monitoring": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the monitoring and control of DNS service on all Microsoft servers in the Grid is enabled or not.",
	},
	"ldap_timeout": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Determines an LDAP connection timeout interval (in seconds) for all Microsoft servers.",
	},
	"default_ip_site_link": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The default IP site link for sites created on NIOS for all Microsoft servers.",
	},
	"enable_network_users": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the Network Users creation is enabled or not.",
	},
	"enable_ad_user_sync": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if Active Directory user synchronization for all Microsoft servers in the Grid is enabled or not.",
	},
	"ad_user_default_timeout": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Determines the default timeout value (in seconds) for Active Directory user synchronization for all Microsoft servers.",
	},
	"enable_dns_reports_sync": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if synchronization of DNS reporting data from all Microsoft servers in the Grid is enabled or not.",
	},
}

func ExpandGridMsSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridMsSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridMsSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridMsSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridMsSetting {
	if m == nil {
		return nil
	}
	to := &grid.GridMsSetting{
		LogDestination:       flex.ExpandStringPointer(m.LogDestination),
		EnableInvalidMac:     flex.ExpandBoolPointer(m.EnableInvalidMac),
		MaxConnection:        flex.ExpandInt64Pointer(m.MaxConnection),
		RpcTimeout:           flex.ExpandInt64Pointer(m.RpcTimeout),
		EnableDhcpMonitoring: flex.ExpandBoolPointer(m.EnableDhcpMonitoring),
		EnableDnsMonitoring:  flex.ExpandBoolPointer(m.EnableDnsMonitoring),
		LdapTimeout:          flex.ExpandInt64Pointer(m.LdapTimeout),
		DefaultIpSiteLink:    flex.ExpandStringPointer(m.DefaultIpSiteLink),
		EnableNetworkUsers:   flex.ExpandBoolPointer(m.EnableNetworkUsers),
		EnableAdUserSync:     flex.ExpandBoolPointer(m.EnableAdUserSync),
		AdUserDefaultTimeout: flex.ExpandInt64Pointer(m.AdUserDefaultTimeout),
		EnableDnsReportsSync: flex.ExpandBoolPointer(m.EnableDnsReportsSync),
	}
	return to
}

func FlattenGridMsSetting(ctx context.Context, from *grid.GridMsSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridMsSettingAttrTypes)
	}
	m := GridMsSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridMsSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridMsSettingModel) Flatten(ctx context.Context, from *grid.GridMsSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridMsSettingModel{}
	}
	m.LogDestination = flex.FlattenStringPointer(from.LogDestination)
	m.EnableInvalidMac = types.BoolPointerValue(from.EnableInvalidMac)
	m.MaxConnection = flex.FlattenInt64Pointer(from.MaxConnection)
	m.RpcTimeout = flex.FlattenInt64Pointer(from.RpcTimeout)
	m.EnableDhcpMonitoring = types.BoolPointerValue(from.EnableDhcpMonitoring)
	m.EnableDnsMonitoring = types.BoolPointerValue(from.EnableDnsMonitoring)
	m.LdapTimeout = flex.FlattenInt64Pointer(from.LdapTimeout)
	m.DefaultIpSiteLink = flex.FlattenStringPointer(from.DefaultIpSiteLink)
	m.EnableNetworkUsers = types.BoolPointerValue(from.EnableNetworkUsers)
	m.EnableAdUserSync = types.BoolPointerValue(from.EnableAdUserSync)
	m.AdUserDefaultTimeout = flex.FlattenInt64Pointer(from.AdUserDefaultTimeout)
	m.EnableDnsReportsSync = types.BoolPointerValue(from.EnableDnsReportsSync)
}
