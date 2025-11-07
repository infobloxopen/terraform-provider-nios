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

type MemberModel struct {
	Ref                             types.String `tfsdk:"ref"`
	ActivePosition                  types.String `tfsdk:"active_position"`
	AdditionalIpList                types.List   `tfsdk:"additional_ip_list"`
	AutomatedTrafficCaptureSetting  types.Object `tfsdk:"automated_traffic_capture_setting"`
	BgpAs                           types.List   `tfsdk:"bgp_as"`
	Comment                         types.String `tfsdk:"comment"`
	ConfigAddrType                  types.String `tfsdk:"config_addr_type"`
	CspAccessKey                    types.List   `tfsdk:"csp_access_key"`
	CspMemberSetting                types.Object `tfsdk:"csp_member_setting"`
	DnsResolverSetting              types.Object `tfsdk:"dns_resolver_setting"`
	Dscp                            types.Int64  `tfsdk:"dscp"`
	EmailSetting                    types.Object `tfsdk:"email_setting"`
	EnableHa                        types.Bool   `tfsdk:"enable_ha"`
	EnableLom                       types.Bool   `tfsdk:"enable_lom"`
	EnableMemberRedirect            types.Bool   `tfsdk:"enable_member_redirect"`
	EnableRoApiAccess               types.Bool   `tfsdk:"enable_ro_api_access"`
	ExtAttrs                        types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll                     types.Map    `tfsdk:"extattrs_all"`
	ExternalSyslogBackupServers     types.List   `tfsdk:"external_syslog_backup_servers"`
	ExternalSyslogServerEnable      types.Bool   `tfsdk:"external_syslog_server_enable"`
	HaCloudPlatform                 types.String `tfsdk:"ha_cloud_platform"`
	HaOnCloud                       types.Bool   `tfsdk:"ha_on_cloud"`
	HostName                        types.String `tfsdk:"host_name"`
	Ipv6Setting                     types.Object `tfsdk:"ipv6_setting"`
	Ipv6StaticRoutes                types.List   `tfsdk:"ipv6_static_routes"`
	IsDscpCapable                   types.Bool   `tfsdk:"is_dscp_capable"`
	Lan2Enabled                     types.Bool   `tfsdk:"lan2_enabled"`
	Lan2PortSetting                 types.Object `tfsdk:"lan2_port_setting"`
	LomNetworkConfig                types.List   `tfsdk:"lom_network_config"`
	LomUsers                        types.List   `tfsdk:"lom_users"`
	MasterCandidate                 types.Bool   `tfsdk:"master_candidate"`
	MemberServiceCommunication      types.List   `tfsdk:"member_service_communication"`
	MgmtPortSetting                 types.Object `tfsdk:"mgmt_port_setting"`
	MmdbEaBuildTime                 types.Int64  `tfsdk:"mmdb_ea_build_time"`
	MmdbGeoipBuildTime              types.Int64  `tfsdk:"mmdb_geoip_build_time"`
	NatSetting                      types.Object `tfsdk:"nat_setting"`
	NodeInfo                        types.List   `tfsdk:"node_info"`
	NtpSetting                      types.Object `tfsdk:"ntp_setting"`
	OspfList                        types.List   `tfsdk:"ospf_list"`
	PassiveHaArpEnabled             types.Bool   `tfsdk:"passive_ha_arp_enabled"`
	Platform                        types.String `tfsdk:"platform"`
	PreProvisioning                 types.Object `tfsdk:"pre_provisioning"`
	PreserveIfOwnsDelegation        types.Bool   `tfsdk:"preserve_if_owns_delegation"`
	RemoteConsoleAccessEnable       types.Bool   `tfsdk:"remote_console_access_enable"`
	RouterId                        types.Int64  `tfsdk:"router_id"`
	ServiceStatus                   types.List   `tfsdk:"service_status"`
	ServiceTypeConfiguration        types.String `tfsdk:"service_type_configuration"`
	SnmpSetting                     types.Object `tfsdk:"snmp_setting"`
	StaticRoutes                    types.List   `tfsdk:"static_routes"`
	SupportAccessEnable             types.Bool   `tfsdk:"support_access_enable"`
	SupportAccessInfo               types.String `tfsdk:"support_access_info"`
	SyslogProxySetting              types.Object `tfsdk:"syslog_proxy_setting"`
	SyslogServers                   types.List   `tfsdk:"syslog_servers"`
	SyslogSize                      types.Int64  `tfsdk:"syslog_size"`
	ThresholdTraps                  types.List   `tfsdk:"threshold_traps"`
	TimeZone                        types.String `tfsdk:"time_zone"`
	TrafficCaptureAuthDnsSetting    types.Object `tfsdk:"traffic_capture_auth_dns_setting"`
	TrafficCaptureChrSetting        types.Object `tfsdk:"traffic_capture_chr_setting"`
	TrafficCaptureQpsSetting        types.Object `tfsdk:"traffic_capture_qps_setting"`
	TrafficCaptureRecDnsSetting     types.Object `tfsdk:"traffic_capture_rec_dns_setting"`
	TrafficCaptureRecQueriesSetting types.Object `tfsdk:"traffic_capture_rec_queries_setting"`
	TrapNotifications               types.List   `tfsdk:"trap_notifications"`
	UpgradeGroup                    types.String `tfsdk:"upgrade_group"`
	UseAutomatedTrafficCapture      types.Bool   `tfsdk:"use_automated_traffic_capture"`
	UseDnsResolverSetting           types.Bool   `tfsdk:"use_dns_resolver_setting"`
	UseDscp                         types.Bool   `tfsdk:"use_dscp"`
	UseEmailSetting                 types.Bool   `tfsdk:"use_email_setting"`
	UseEnableLom                    types.Bool   `tfsdk:"use_enable_lom"`
	UseEnableMemberRedirect         types.Bool   `tfsdk:"use_enable_member_redirect"`
	UseExternalSyslogBackupServers  types.Bool   `tfsdk:"use_external_syslog_backup_servers"`
	UseRemoteConsoleAccessEnable    types.Bool   `tfsdk:"use_remote_console_access_enable"`
	UseSnmpSetting                  types.Bool   `tfsdk:"use_snmp_setting"`
	UseSupportAccessEnable          types.Bool   `tfsdk:"use_support_access_enable"`
	UseSyslogProxySetting           types.Bool   `tfsdk:"use_syslog_proxy_setting"`
	UseThresholdTraps               types.Bool   `tfsdk:"use_threshold_traps"`
	UseTimeZone                     types.Bool   `tfsdk:"use_time_zone"`
	UseTrafficCaptureAuthDns        types.Bool   `tfsdk:"use_traffic_capture_auth_dns"`
	UseTrafficCaptureChr            types.Bool   `tfsdk:"use_traffic_capture_chr"`
	UseTrafficCaptureQps            types.Bool   `tfsdk:"use_traffic_capture_qps"`
	UseTrafficCaptureRecDns         types.Bool   `tfsdk:"use_traffic_capture_rec_dns"`
	UseTrafficCaptureRecQueries     types.Bool   `tfsdk:"use_traffic_capture_rec_queries"`
	UseTrapNotifications            types.Bool   `tfsdk:"use_trap_notifications"`
	UseV4Vrrp                       types.Bool   `tfsdk:"use_v4_vrrp"`
	VipSetting                      types.Object `tfsdk:"vip_setting"`
	VpnMtu                          types.Int64  `tfsdk:"vpn_mtu"`
}

var MemberAttrTypes = map[string]attr.Type{
	"ref":                                 types.StringType,
	"active_position":                     types.StringType,
	"additional_ip_list":                  types.ListType{ElemType: types.ObjectType{AttrTypes: MemberAdditionalIpListAttrTypes}},
	"automated_traffic_capture_setting":   types.ObjectType{AttrTypes: MemberAutomatedTrafficCaptureSettingAttrTypes},
	"bgp_as":                              types.ListType{ElemType: types.ObjectType{AttrTypes: MemberBgpAsAttrTypes}},
	"comment":                             types.StringType,
	"config_addr_type":                    types.StringType,
	"csp_access_key":                      types.ListType{ElemType: types.StringType},
	"csp_member_setting":                  types.ObjectType{AttrTypes: MemberCspMemberSettingAttrTypes},
	"dns_resolver_setting":                types.ObjectType{AttrTypes: MemberDnsResolverSettingAttrTypes},
	"dscp":                                types.Int64Type,
	"email_setting":                       types.ObjectType{AttrTypes: MemberEmailSettingAttrTypes},
	"enable_ha":                           types.BoolType,
	"enable_lom":                          types.BoolType,
	"enable_member_redirect":              types.BoolType,
	"enable_ro_api_access":                types.BoolType,
	"extattrs":                            types.MapType{ElemType: types.StringType},
	"extattrs_all":                        types.MapType{ElemType: types.StringType},
	"external_syslog_backup_servers":      types.ListType{ElemType: types.ObjectType{AttrTypes: MemberExternalSyslogBackupServersAttrTypes}},
	"external_syslog_server_enable":       types.BoolType,
	"ha_cloud_platform":                   types.StringType,
	"ha_on_cloud":                         types.BoolType,
	"host_name":                           types.StringType,
	"ipv6_setting":                        types.ObjectType{AttrTypes: MemberIpv6SettingAttrTypes},
	"ipv6_static_routes":                  types.ListType{ElemType: types.ObjectType{AttrTypes: MemberIpv6StaticRoutesAttrTypes}},
	"is_dscp_capable":                     types.BoolType,
	"lan2_enabled":                        types.BoolType,
	"lan2_port_setting":                   types.ObjectType{AttrTypes: MemberLan2PortSettingAttrTypes},
	"lom_network_config":                  types.ListType{ElemType: types.ObjectType{AttrTypes: MemberLomNetworkConfigAttrTypes}},
	"lom_users":                           types.ListType{ElemType: types.ObjectType{AttrTypes: MemberLomUsersAttrTypes}},
	"master_candidate":                    types.BoolType,
	"member_service_communication":        types.ListType{ElemType: types.ObjectType{AttrTypes: MemberMemberServiceCommunicationAttrTypes}},
	"mgmt_port_setting":                   types.ObjectType{AttrTypes: MemberMgmtPortSettingAttrTypes},
	"mmdb_ea_build_time":                  types.Int64Type,
	"mmdb_geoip_build_time":               types.Int64Type,
	"nat_setting":                         types.ObjectType{AttrTypes: MemberNatSettingAttrTypes},
	"node_info":                           types.ListType{ElemType: types.ObjectType{AttrTypes: MemberNodeInfoAttrTypes}},
	"ntp_setting":                         types.ObjectType{AttrTypes: MemberNtpSettingAttrTypes},
	"ospf_list":                           types.ListType{ElemType: types.ObjectType{AttrTypes: MemberOspfListAttrTypes}},
	"passive_ha_arp_enabled":              types.BoolType,
	"platform":                            types.StringType,
	"pre_provisioning":                    types.ObjectType{AttrTypes: MemberPreProvisioningAttrTypes},
	"preserve_if_owns_delegation":         types.BoolType,
	"remote_console_access_enable":        types.BoolType,
	"router_id":                           types.Int64Type,
	"service_status":                      types.ListType{ElemType: types.ObjectType{AttrTypes: MemberServiceStatusAttrTypes}},
	"service_type_configuration":          types.StringType,
	"snmp_setting":                        types.ObjectType{AttrTypes: MemberSnmpSettingAttrTypes},
	"static_routes":                       types.ListType{ElemType: types.ObjectType{AttrTypes: MemberStaticRoutesAttrTypes}},
	"support_access_enable":               types.BoolType,
	"support_access_info":                 types.StringType,
	"syslog_proxy_setting":                types.ObjectType{AttrTypes: MemberSyslogProxySettingAttrTypes},
	"syslog_servers":                      types.ListType{ElemType: types.ObjectType{AttrTypes: MemberSyslogServersAttrTypes}},
	"syslog_size":                         types.Int64Type,
	"threshold_traps":                     types.ListType{ElemType: types.ObjectType{AttrTypes: MemberThresholdTrapsAttrTypes}},
	"time_zone":                           types.StringType,
	"traffic_capture_auth_dns_setting":    types.ObjectType{AttrTypes: MemberTrafficCaptureAuthDnsSettingAttrTypes},
	"traffic_capture_chr_setting":         types.ObjectType{AttrTypes: MemberTrafficCaptureChrSettingAttrTypes},
	"traffic_capture_qps_setting":         types.ObjectType{AttrTypes: MemberTrafficCaptureQpsSettingAttrTypes},
	"traffic_capture_rec_dns_setting":     types.ObjectType{AttrTypes: MemberTrafficCaptureRecDnsSettingAttrTypes},
	"traffic_capture_rec_queries_setting": types.ObjectType{AttrTypes: MemberTrafficCaptureRecQueriesSettingAttrTypes},
	"trap_notifications":                  types.ListType{ElemType: types.ObjectType{AttrTypes: MemberTrapNotificationsAttrTypes}},
	"upgrade_group":                       types.StringType,
	"use_automated_traffic_capture":       types.BoolType,
	"use_dns_resolver_setting":            types.BoolType,
	"use_dscp":                            types.BoolType,
	"use_email_setting":                   types.BoolType,
	"use_enable_lom":                      types.BoolType,
	"use_enable_member_redirect":          types.BoolType,
	"use_external_syslog_backup_servers":  types.BoolType,
	"use_remote_console_access_enable":    types.BoolType,
	"use_snmp_setting":                    types.BoolType,
	"use_support_access_enable":           types.BoolType,
	"use_syslog_proxy_setting":            types.BoolType,
	"use_threshold_traps":                 types.BoolType,
	"use_time_zone":                       types.BoolType,
	"use_traffic_capture_auth_dns":        types.BoolType,
	"use_traffic_capture_chr":             types.BoolType,
	"use_traffic_capture_qps":             types.BoolType,
	"use_traffic_capture_rec_dns":         types.BoolType,
	"use_traffic_capture_rec_queries":     types.BoolType,
	"use_trap_notifications":              types.BoolType,
	"use_v4_vrrp":                         types.BoolType,
	"vip_setting":                         types.ObjectType{AttrTypes: MemberVipSettingAttrTypes},
	"vpn_mtu":                             types.Int64Type,
}

var MemberResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"active_position": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The active server of a Grid member.",
	},
	"additional_ip_list": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberAdditionalIpListResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The additional IP list of a Grid member. This list contains additional interface information that can be used at the member level. Note that interface structure(s) with interface type set to 'MGMT' are not supported.",
	},
	"automated_traffic_capture_setting": schema.SingleNestedAttribute{
		Attributes: MemberAutomatedTrafficCaptureSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"bgp_as": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberBgpAsResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The BGP configuration for anycast for a Grid member.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "A descriptive comment of the Grid member.",
	},
	"config_addr_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Address configuration type.",
	},
	"csp_access_key": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "CSP portal on-prem host access key",
	},
	"csp_member_setting": schema.SingleNestedAttribute{
		Attributes: MemberCspMemberSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"dns_resolver_setting": schema.SingleNestedAttribute{
		Attributes: MemberDnsResolverSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"dscp": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The DSCP (Differentiated Services Code Point) value.",
	},
	"email_setting": schema.SingleNestedAttribute{
		Attributes: MemberEmailSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"enable_ha": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If set to True, the member has two physical nodes (HA pair).",
	},
	"enable_lom": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the LOM functionality is enabled or not.",
	},
	"enable_member_redirect": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the member will redirect GUI connections to the Grid Master or not.",
	},
	"enable_ro_api_access": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If set to True and the member object is a Grid Master Candidate, then read-only API access is enabled.",
	},
	"extattrs": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.",
	},
	"extattrs_all": schema.MapAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object, including default attributes.",
	},
	"external_syslog_backup_servers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberExternalSyslogBackupServersResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of external syslog backup servers.",
	},
	"external_syslog_server_enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if external syslog servers should be enabled.",
	},
	"ha_cloud_platform": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Cloud platform for HA.",
	},
	"ha_on_cloud": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "True: HA on cloud. False: HA not on cloud.",
	},
	"host_name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The host name of the Grid member.",
	},
	"ipv6_setting": schema.SingleNestedAttribute{
		Attributes: MemberIpv6SettingResourceSchemaAttributes,
		Optional:   true,
	},
	"ipv6_static_routes": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberIpv6StaticRoutesResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "List of IPv6 static routes.",
	},
	"is_dscp_capable": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines if a Grid member supports DSCP (Differentiated Services Code Point).",
	},
	"lan2_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If this is set to \"true\", the LAN2 port is enabled as an independent port or as a port for failover purposes.",
	},
	"lan2_port_setting": schema.SingleNestedAttribute{
		Attributes: MemberLan2PortSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"lom_network_config": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberLomNetworkConfigResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The Network configurations for LOM.",
	},
	"lom_users": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberLomUsersResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of LOM users.",
	},
	"master_candidate": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if a Grid member is a Grid Master Candidate or not. This flag enables the Grid member to assume the role of the Grid Master as a disaster recovery measure.",
	},
	"member_service_communication": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberMemberServiceCommunicationResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "Configure communication type for various services.",
	},
	"mgmt_port_setting": schema.SingleNestedAttribute{
		Attributes: MemberMgmtPortSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"mmdb_ea_build_time": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Extensible attributes Topology database build time.",
	},
	"mmdb_geoip_build_time": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "GeoIP Topology database build time.",
	},
	"nat_setting": schema.SingleNestedAttribute{
		Attributes: MemberNatSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"node_info": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberNodeInfoResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The node information list with detailed status report on the operations of the Grid Member, mgmt_port_setting must be enabled when configuring the MGMT Port using the node_info field.",
	},
	"ntp_setting": schema.SingleNestedAttribute{
		Attributes: MemberNtpSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"ospf_list": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberOspfListResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The OSPF area configuration (for anycast) list for a Grid member.",
	},
	"passive_ha_arp_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "The ARP protocol setting on the passive node of an HA pair. If you do not specify a value, the default value is \"false\". You can only set this value to \"true\" if the member is an HA pair.",
	},
	"platform": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Hardware Platform.",
	},
	"pre_provisioning": schema.SingleNestedAttribute{
		Attributes: MemberPreProvisioningResourceSchemaAttributes,
		Optional:   true,
	},
	"preserve_if_owns_delegation": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Set this flag to \"true\" to prevent the deletion of the member if any delegated object remains attached to it.",
	},
	"remote_console_access_enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If set to True, superuser admins can access the Infoblox CLI from a remote location using an SSH (Secure Shell) v2 client.",
	},
	"router_id": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Virutal router identifier. Provide this ID if \"ha_enabled\" is set to \"true\". This is a unique VRID number (from 1 to 255) for the local subnet.",
	},
	"service_status": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberServiceStatusResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Computed:            true,
		MarkdownDescription: "The service status list of a grid member.",
	},
	"service_type_configuration": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Configure all services to the given type.",
	},
	"snmp_setting": schema.SingleNestedAttribute{
		Attributes: MemberSnmpSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"static_routes": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberStaticRoutesResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "List of static routes.",
	},
	"support_access_enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if support access for the Grid member should be enabled.",
	},
	"support_access_info": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The information string for support access.",
	},
	"syslog_proxy_setting": schema.SingleNestedAttribute{
		Attributes: MemberSyslogProxySettingResourceSchemaAttributes,
		Optional:   true,
	},
	"syslog_servers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberSyslogServersResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of external syslog servers.",
	},
	"syslog_size": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The maximum size for the syslog file expressed in megabytes.",
	},
	"threshold_traps": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberThresholdTrapsResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "Determines the list of threshold traps. The user can only change the values for each trap or remove traps.",
	},
	"time_zone": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The time zone of the Grid member. The UTC string that represents the time zone, such as \"Asia/Kolkata\".",
	},
	"traffic_capture_auth_dns_setting": schema.SingleNestedAttribute{
		Attributes: MemberTrafficCaptureAuthDnsSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"traffic_capture_chr_setting": schema.SingleNestedAttribute{
		Attributes: MemberTrafficCaptureChrSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"traffic_capture_qps_setting": schema.SingleNestedAttribute{
		Attributes: MemberTrafficCaptureQpsSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"traffic_capture_rec_dns_setting": schema.SingleNestedAttribute{
		Attributes: MemberTrafficCaptureRecDnsSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"traffic_capture_rec_queries_setting": schema.SingleNestedAttribute{
		Attributes: MemberTrafficCaptureRecQueriesSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"trap_notifications": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberTrapNotificationsResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "Determines configuration of the trap notifications.",
	},
	"upgrade_group": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the upgrade group to which this Grid member belongs.",
	},
	"use_automated_traffic_capture": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "This flag is the use flag for enabling automated traffic capture based on DNS cache ratio thresholds.",
	},
	"use_dns_resolver_setting": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: dns_resolver_setting",
	},
	"use_dscp": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: dscp",
	},
	"use_email_setting": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: email_setting",
	},
	"use_enable_lom": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: enable_lom",
	},
	"use_enable_member_redirect": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: enable_member_redirect",
	},
	"use_external_syslog_backup_servers": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: external_syslog_backup_servers",
	},
	"use_remote_console_access_enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: remote_console_access_enable",
	},
	"use_snmp_setting": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: snmp_setting",
	},
	"use_support_access_enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: support_access_enable",
	},
	"use_syslog_proxy_setting": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: external_syslog_server_enable , syslog_servers, syslog_proxy_setting, syslog_size",
	},
	"use_threshold_traps": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: threshold_traps",
	},
	"use_time_zone": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: time_zone",
	},
	"use_traffic_capture_auth_dns": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "This flag is the use flag for enabling automated traffic capture based on authorative DNS latency.",
	},
	"use_traffic_capture_chr": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "This flag is the use flag for automated traffic capture settings at member level.",
	},
	"use_traffic_capture_qps": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "This flag is the use flag for enabling automated traffic capture based on DNS querie per second thresholds.",
	},
	"use_traffic_capture_rec_dns": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "This flag is the use flag for enabling automated traffic capture based on recursive DNS latency.",
	},
	"use_traffic_capture_rec_queries": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "This flag is the use flag for enabling automated traffic capture based on outgoing recursive queries.",
	},
	"use_trap_notifications": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: trap_notifications",
	},
	"use_v4_vrrp": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Specify \"true\" to use VRRPv4 or \"false\" to use VRRPv6.",
	},
	"vip_setting": schema.SingleNestedAttribute{
		Attributes: MemberVipSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"vpn_mtu": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The VPN maximum transmission unit (MTU).",
	},
}

func ExpandMember(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.Member {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.Member {
	if m == nil {
		return nil
	}
	to := &grid.Member{
		Ref:                             flex.ExpandStringPointer(m.Ref),
		AdditionalIpList:                flex.ExpandFrameworkListNestedBlock(ctx, m.AdditionalIpList, diags, ExpandMemberAdditionalIpList),
		AutomatedTrafficCaptureSetting:  ExpandMemberAutomatedTrafficCaptureSetting(ctx, m.AutomatedTrafficCaptureSetting, diags),
		BgpAs:                           flex.ExpandFrameworkListNestedBlock(ctx, m.BgpAs, diags, ExpandMemberBgpAs),
		Comment:                         flex.ExpandStringPointer(m.Comment),
		ConfigAddrType:                  flex.ExpandStringPointer(m.ConfigAddrType),
		CspAccessKey:                    flex.ExpandFrameworkListString(ctx, m.CspAccessKey, diags),
		CspMemberSetting:                ExpandMemberCspMemberSetting(ctx, m.CspMemberSetting, diags),
		DnsResolverSetting:              ExpandMemberDnsResolverSetting(ctx, m.DnsResolverSetting, diags),
		Dscp:                            flex.ExpandInt64Pointer(m.Dscp),
		EmailSetting:                    ExpandMemberEmailSetting(ctx, m.EmailSetting, diags),
		EnableHa:                        flex.ExpandBoolPointer(m.EnableHa),
		EnableLom:                       flex.ExpandBoolPointer(m.EnableLom),
		EnableMemberRedirect:            flex.ExpandBoolPointer(m.EnableMemberRedirect),
		EnableRoApiAccess:               flex.ExpandBoolPointer(m.EnableRoApiAccess),
		ExtAttrs:                        ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		ExternalSyslogBackupServers:     flex.ExpandFrameworkListNestedBlock(ctx, m.ExternalSyslogBackupServers, diags, ExpandMemberExternalSyslogBackupServers),
		ExternalSyslogServerEnable:      flex.ExpandBoolPointer(m.ExternalSyslogServerEnable),
		HaCloudPlatform:                 flex.ExpandStringPointer(m.HaCloudPlatform),
		HaOnCloud:                       flex.ExpandBoolPointer(m.HaOnCloud),
		HostName:                        flex.ExpandStringPointer(m.HostName),
		Ipv6Setting:                     ExpandMemberIpv6Setting(ctx, m.Ipv6Setting, diags),
		Ipv6StaticRoutes:                flex.ExpandFrameworkListNestedBlock(ctx, m.Ipv6StaticRoutes, diags, ExpandMemberIpv6StaticRoutes),
		Lan2Enabled:                     flex.ExpandBoolPointer(m.Lan2Enabled),
		Lan2PortSetting:                 ExpandMemberLan2PortSetting(ctx, m.Lan2PortSetting, diags),
		LomNetworkConfig:                flex.ExpandFrameworkListNestedBlock(ctx, m.LomNetworkConfig, diags, ExpandMemberLomNetworkConfig),
		LomUsers:                        flex.ExpandFrameworkListNestedBlock(ctx, m.LomUsers, diags, ExpandMemberLomUsers),
		MasterCandidate:                 flex.ExpandBoolPointer(m.MasterCandidate),
		MemberServiceCommunication:      flex.ExpandFrameworkListNestedBlock(ctx, m.MemberServiceCommunication, diags, ExpandMemberMemberServiceCommunication),
		MgmtPortSetting:                 ExpandMemberMgmtPortSetting(ctx, m.MgmtPortSetting, diags),
		NatSetting:                      ExpandMemberNatSetting(ctx, m.NatSetting, diags),
		NodeInfo:                        flex.ExpandFrameworkListNestedBlock(ctx, m.NodeInfo, diags, ExpandMemberNodeInfo),
		NtpSetting:                      ExpandMemberNtpSetting(ctx, m.NtpSetting, diags),
		OspfList:                        flex.ExpandFrameworkListNestedBlock(ctx, m.OspfList, diags, ExpandMemberOspfList),
		PassiveHaArpEnabled:             flex.ExpandBoolPointer(m.PassiveHaArpEnabled),
		Platform:                        flex.ExpandStringPointer(m.Platform),
		PreProvisioning:                 ExpandMemberPreProvisioning(ctx, m.PreProvisioning, diags),
		PreserveIfOwnsDelegation:        flex.ExpandBoolPointer(m.PreserveIfOwnsDelegation),
		RemoteConsoleAccessEnable:       flex.ExpandBoolPointer(m.RemoteConsoleAccessEnable),
		RouterId:                        flex.ExpandInt64Pointer(m.RouterId),
		ServiceTypeConfiguration:        flex.ExpandStringPointer(m.ServiceTypeConfiguration),
		SnmpSetting:                     ExpandMemberSnmpSetting(ctx, m.SnmpSetting, diags),
		StaticRoutes:                    flex.ExpandFrameworkListNestedBlock(ctx, m.StaticRoutes, diags, ExpandMemberStaticRoutes),
		SupportAccessEnable:             flex.ExpandBoolPointer(m.SupportAccessEnable),
		SyslogProxySetting:              ExpandMemberSyslogProxySetting(ctx, m.SyslogProxySetting, diags),
		SyslogServers:                   flex.ExpandFrameworkListNestedBlock(ctx, m.SyslogServers, diags, ExpandMemberSyslogServers),
		SyslogSize:                      flex.ExpandInt64Pointer(m.SyslogSize),
		ThresholdTraps:                  flex.ExpandFrameworkListNestedBlock(ctx, m.ThresholdTraps, diags, ExpandMemberThresholdTraps),
		TimeZone:                        flex.ExpandStringPointer(m.TimeZone),
		TrafficCaptureAuthDnsSetting:    ExpandMemberTrafficCaptureAuthDnsSetting(ctx, m.TrafficCaptureAuthDnsSetting, diags),
		TrafficCaptureChrSetting:        ExpandMemberTrafficCaptureChrSetting(ctx, m.TrafficCaptureChrSetting, diags),
		TrafficCaptureQpsSetting:        ExpandMemberTrafficCaptureQpsSetting(ctx, m.TrafficCaptureQpsSetting, diags),
		TrafficCaptureRecDnsSetting:     ExpandMemberTrafficCaptureRecDnsSetting(ctx, m.TrafficCaptureRecDnsSetting, diags),
		TrafficCaptureRecQueriesSetting: ExpandMemberTrafficCaptureRecQueriesSetting(ctx, m.TrafficCaptureRecQueriesSetting, diags),
		TrapNotifications:               flex.ExpandFrameworkListNestedBlock(ctx, m.TrapNotifications, diags, ExpandMemberTrapNotifications),
		UpgradeGroup:                    flex.ExpandStringPointer(m.UpgradeGroup),
		UseAutomatedTrafficCapture:      flex.ExpandBoolPointer(m.UseAutomatedTrafficCapture),
		UseDnsResolverSetting:           flex.ExpandBoolPointer(m.UseDnsResolverSetting),
		UseDscp:                         flex.ExpandBoolPointer(m.UseDscp),
		UseEmailSetting:                 flex.ExpandBoolPointer(m.UseEmailSetting),
		UseEnableLom:                    flex.ExpandBoolPointer(m.UseEnableLom),
		UseEnableMemberRedirect:         flex.ExpandBoolPointer(m.UseEnableMemberRedirect),
		UseExternalSyslogBackupServers:  flex.ExpandBoolPointer(m.UseExternalSyslogBackupServers),
		UseRemoteConsoleAccessEnable:    flex.ExpandBoolPointer(m.UseRemoteConsoleAccessEnable),
		UseSnmpSetting:                  flex.ExpandBoolPointer(m.UseSnmpSetting),
		UseSupportAccessEnable:          flex.ExpandBoolPointer(m.UseSupportAccessEnable),
		UseSyslogProxySetting:           flex.ExpandBoolPointer(m.UseSyslogProxySetting),
		UseThresholdTraps:               flex.ExpandBoolPointer(m.UseThresholdTraps),
		UseTimeZone:                     flex.ExpandBoolPointer(m.UseTimeZone),
		UseTrafficCaptureAuthDns:        flex.ExpandBoolPointer(m.UseTrafficCaptureAuthDns),
		UseTrafficCaptureChr:            flex.ExpandBoolPointer(m.UseTrafficCaptureChr),
		UseTrafficCaptureQps:            flex.ExpandBoolPointer(m.UseTrafficCaptureQps),
		UseTrafficCaptureRecDns:         flex.ExpandBoolPointer(m.UseTrafficCaptureRecDns),
		UseTrafficCaptureRecQueries:     flex.ExpandBoolPointer(m.UseTrafficCaptureRecQueries),
		UseTrapNotifications:            flex.ExpandBoolPointer(m.UseTrapNotifications),
		UseV4Vrrp:                       flex.ExpandBoolPointer(m.UseV4Vrrp),
		VipSetting:                      ExpandMemberVipSetting(ctx, m.VipSetting, diags),
		VpnMtu:                          flex.ExpandInt64Pointer(m.VpnMtu),
	}
	return to
}

func FlattenMember(ctx context.Context, from *grid.Member, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberAttrTypes)
	}
	m := MemberModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, MemberAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberModel) Flatten(ctx context.Context, from *grid.Member, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.ActivePosition = flex.FlattenStringPointer(from.ActivePosition)
	m.AdditionalIpList = flex.FlattenFrameworkListNestedBlock(ctx, from.AdditionalIpList, MemberAdditionalIpListAttrTypes, diags, FlattenMemberAdditionalIpList)
	m.AutomatedTrafficCaptureSetting = FlattenMemberAutomatedTrafficCaptureSetting(ctx, from.AutomatedTrafficCaptureSetting, diags)
	m.BgpAs = flex.FlattenFrameworkListNestedBlock(ctx, from.BgpAs, MemberBgpAsAttrTypes, diags, FlattenMemberBgpAs)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.ConfigAddrType = flex.FlattenStringPointer(from.ConfigAddrType)
	m.CspAccessKey = flex.FlattenFrameworkListString(ctx, from.CspAccessKey, diags)
	m.CspMemberSetting = FlattenMemberCspMemberSetting(ctx, from.CspMemberSetting, diags)
	m.DnsResolverSetting = FlattenMemberDnsResolverSetting(ctx, from.DnsResolverSetting, diags)
	m.Dscp = flex.FlattenInt64Pointer(from.Dscp)
	m.EmailSetting = FlattenMemberEmailSetting(ctx, from.EmailSetting, diags)
	m.EnableHa = types.BoolPointerValue(from.EnableHa)
	m.EnableLom = types.BoolPointerValue(from.EnableLom)
	m.EnableMemberRedirect = types.BoolPointerValue(from.EnableMemberRedirect)
	m.EnableRoApiAccess = types.BoolPointerValue(from.EnableRoApiAccess)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.ExternalSyslogBackupServers = flex.FlattenFrameworkListNestedBlock(ctx, from.ExternalSyslogBackupServers, MemberExternalSyslogBackupServersAttrTypes, diags, FlattenMemberExternalSyslogBackupServers)
	m.ExternalSyslogServerEnable = types.BoolPointerValue(from.ExternalSyslogServerEnable)
	m.HaCloudPlatform = flex.FlattenStringPointer(from.HaCloudPlatform)
	m.HaOnCloud = types.BoolPointerValue(from.HaOnCloud)
	m.HostName = flex.FlattenStringPointer(from.HostName)
	m.Ipv6Setting = FlattenMemberIpv6Setting(ctx, from.Ipv6Setting, diags)
	m.Ipv6StaticRoutes = flex.FlattenFrameworkListNestedBlock(ctx, from.Ipv6StaticRoutes, MemberIpv6StaticRoutesAttrTypes, diags, FlattenMemberIpv6StaticRoutes)
	m.IsDscpCapable = types.BoolPointerValue(from.IsDscpCapable)
	m.Lan2Enabled = types.BoolPointerValue(from.Lan2Enabled)
	m.Lan2PortSetting = FlattenMemberLan2PortSetting(ctx, from.Lan2PortSetting, diags)
	m.LomNetworkConfig = flex.FlattenFrameworkListNestedBlock(ctx, from.LomNetworkConfig, MemberLomNetworkConfigAttrTypes, diags, FlattenMemberLomNetworkConfig)
	m.LomUsers = flex.FlattenFrameworkListNestedBlock(ctx, from.LomUsers, MemberLomUsersAttrTypes, diags, FlattenMemberLomUsers)
	m.MasterCandidate = types.BoolPointerValue(from.MasterCandidate)
	m.MemberServiceCommunication = flex.FlattenFrameworkListNestedBlock(ctx, from.MemberServiceCommunication, MemberMemberServiceCommunicationAttrTypes, diags, FlattenMemberMemberServiceCommunication)
	m.MgmtPortSetting = FlattenMemberMgmtPortSetting(ctx, from.MgmtPortSetting, diags)
	m.MmdbEaBuildTime = flex.FlattenInt64Pointer(from.MmdbEaBuildTime)
	m.MmdbGeoipBuildTime = flex.FlattenInt64Pointer(from.MmdbGeoipBuildTime)
	m.NatSetting = FlattenMemberNatSetting(ctx, from.NatSetting, diags)
	m.NodeInfo = flex.FlattenFrameworkListNestedBlock(ctx, from.NodeInfo, MemberNodeInfoAttrTypes, diags, FlattenMemberNodeInfo)
	m.NtpSetting = FlattenMemberNtpSetting(ctx, from.NtpSetting, diags)
	m.OspfList = flex.FlattenFrameworkListNestedBlock(ctx, from.OspfList, MemberOspfListAttrTypes, diags, FlattenMemberOspfList)
	m.PassiveHaArpEnabled = types.BoolPointerValue(from.PassiveHaArpEnabled)
	m.Platform = flex.FlattenStringPointer(from.Platform)
	m.PreProvisioning = FlattenMemberPreProvisioning(ctx, from.PreProvisioning, diags)
	m.PreserveIfOwnsDelegation = types.BoolPointerValue(from.PreserveIfOwnsDelegation)
	m.RemoteConsoleAccessEnable = types.BoolPointerValue(from.RemoteConsoleAccessEnable)
	m.RouterId = flex.FlattenInt64Pointer(from.RouterId)
	m.ServiceStatus = flex.FlattenFrameworkListNestedBlock(ctx, from.ServiceStatus, MemberServiceStatusAttrTypes, diags, FlattenMemberServiceStatus)
	m.ServiceTypeConfiguration = flex.FlattenStringPointer(from.ServiceTypeConfiguration)
	m.SnmpSetting = FlattenMemberSnmpSetting(ctx, from.SnmpSetting, diags)
	m.StaticRoutes = flex.FlattenFrameworkListNestedBlock(ctx, from.StaticRoutes, MemberStaticRoutesAttrTypes, diags, FlattenMemberStaticRoutes)
	m.SupportAccessEnable = types.BoolPointerValue(from.SupportAccessEnable)
	m.SupportAccessInfo = flex.FlattenStringPointer(from.SupportAccessInfo)
	m.SyslogProxySetting = FlattenMemberSyslogProxySetting(ctx, from.SyslogProxySetting, diags)
	m.SyslogServers = flex.FlattenFrameworkListNestedBlock(ctx, from.SyslogServers, MemberSyslogServersAttrTypes, diags, FlattenMemberSyslogServers)
	m.SyslogSize = flex.FlattenInt64Pointer(from.SyslogSize)
	m.ThresholdTraps = flex.FlattenFrameworkListNestedBlock(ctx, from.ThresholdTraps, MemberThresholdTrapsAttrTypes, diags, FlattenMemberThresholdTraps)
	m.TimeZone = flex.FlattenStringPointer(from.TimeZone)
	m.TrafficCaptureAuthDnsSetting = FlattenMemberTrafficCaptureAuthDnsSetting(ctx, from.TrafficCaptureAuthDnsSetting, diags)
	m.TrafficCaptureChrSetting = FlattenMemberTrafficCaptureChrSetting(ctx, from.TrafficCaptureChrSetting, diags)
	m.TrafficCaptureQpsSetting = FlattenMemberTrafficCaptureQpsSetting(ctx, from.TrafficCaptureQpsSetting, diags)
	m.TrafficCaptureRecDnsSetting = FlattenMemberTrafficCaptureRecDnsSetting(ctx, from.TrafficCaptureRecDnsSetting, diags)
	m.TrafficCaptureRecQueriesSetting = FlattenMemberTrafficCaptureRecQueriesSetting(ctx, from.TrafficCaptureRecQueriesSetting, diags)
	m.TrapNotifications = flex.FlattenFrameworkListNestedBlock(ctx, from.TrapNotifications, MemberTrapNotificationsAttrTypes, diags, FlattenMemberTrapNotifications)
	m.UpgradeGroup = flex.FlattenStringPointer(from.UpgradeGroup)
	m.UseAutomatedTrafficCapture = types.BoolPointerValue(from.UseAutomatedTrafficCapture)
	m.UseDnsResolverSetting = types.BoolPointerValue(from.UseDnsResolverSetting)
	m.UseDscp = types.BoolPointerValue(from.UseDscp)
	m.UseEmailSetting = types.BoolPointerValue(from.UseEmailSetting)
	m.UseEnableLom = types.BoolPointerValue(from.UseEnableLom)
	m.UseEnableMemberRedirect = types.BoolPointerValue(from.UseEnableMemberRedirect)
	m.UseExternalSyslogBackupServers = types.BoolPointerValue(from.UseExternalSyslogBackupServers)
	m.UseRemoteConsoleAccessEnable = types.BoolPointerValue(from.UseRemoteConsoleAccessEnable)
	m.UseSnmpSetting = types.BoolPointerValue(from.UseSnmpSetting)
	m.UseSupportAccessEnable = types.BoolPointerValue(from.UseSupportAccessEnable)
	m.UseSyslogProxySetting = types.BoolPointerValue(from.UseSyslogProxySetting)
	m.UseThresholdTraps = types.BoolPointerValue(from.UseThresholdTraps)
	m.UseTimeZone = types.BoolPointerValue(from.UseTimeZone)
	m.UseTrafficCaptureAuthDns = types.BoolPointerValue(from.UseTrafficCaptureAuthDns)
	m.UseTrafficCaptureChr = types.BoolPointerValue(from.UseTrafficCaptureChr)
	m.UseTrafficCaptureQps = types.BoolPointerValue(from.UseTrafficCaptureQps)
	m.UseTrafficCaptureRecDns = types.BoolPointerValue(from.UseTrafficCaptureRecDns)
	m.UseTrafficCaptureRecQueries = types.BoolPointerValue(from.UseTrafficCaptureRecQueries)
	m.UseTrapNotifications = types.BoolPointerValue(from.UseTrapNotifications)
	m.UseV4Vrrp = types.BoolPointerValue(from.UseV4Vrrp)
	m.VipSetting = FlattenMemberVipSetting(ctx, from.VipSetting, diags)
	m.VpnMtu = flex.FlattenInt64Pointer(from.VpnMtu)
}
