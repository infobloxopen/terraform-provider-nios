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

type GridDhcppropertiesModel struct {
	Ref                                  types.String `tfsdk:"ref"`
	Uuid                                 types.String `tfsdk:"uuid"`
	Authority                            types.Bool   `tfsdk:"authority"`
	Bootfile                             types.String `tfsdk:"bootfile"`
	Bootserver                           types.String `tfsdk:"bootserver"`
	CaptureHostname                      types.Bool   `tfsdk:"capture_hostname"`
	DdnsDomainname                       types.String `tfsdk:"ddns_domainname"`
	DdnsGenerateHostname                 types.Bool   `tfsdk:"ddns_generate_hostname"`
	DdnsRetryInterval                    types.Int64  `tfsdk:"ddns_retry_interval"`
	DdnsServerAlwaysUpdates              types.Bool   `tfsdk:"ddns_server_always_updates"`
	DdnsTtl                              types.Int64  `tfsdk:"ddns_ttl"`
	DdnsUpdateFixedAddresses             types.Bool   `tfsdk:"ddns_update_fixed_addresses"`
	DdnsUseOption81                      types.Bool   `tfsdk:"ddns_use_option81"`
	DenyBootp                            types.Bool   `tfsdk:"deny_bootp"`
	DisableAllNacFilters                 types.Bool   `tfsdk:"disable_all_nac_filters"`
	DnsUpdateStyle                       types.String `tfsdk:"dns_update_style"`
	EmailList                            types.List   `tfsdk:"email_list"`
	EnableDdns                           types.Bool   `tfsdk:"enable_ddns"`
	EnableDhcpThresholds                 types.Bool   `tfsdk:"enable_dhcp_thresholds"`
	EnableEmailWarnings                  types.Bool   `tfsdk:"enable_email_warnings"`
	EnableFingerprint                    types.Bool   `tfsdk:"enable_fingerprint"`
	EnableGssTsig                        types.Bool   `tfsdk:"enable_gss_tsig"`
	EnableHostnameRewrite                types.Bool   `tfsdk:"enable_hostname_rewrite"`
	EnableLeasequery                     types.Bool   `tfsdk:"enable_leasequery"`
	EnableRoamingHosts                   types.Bool   `tfsdk:"enable_roaming_hosts"`
	EnableSnmpWarnings                   types.Bool   `tfsdk:"enable_snmp_warnings"`
	FormatLogOption82                    types.String `tfsdk:"format_log_option_82"`
	Grid                                 types.String `tfsdk:"grid"`
	GssTsigKeys                          types.List   `tfsdk:"gss_tsig_keys"`
	HighWaterMark                        types.Int64  `tfsdk:"high_water_mark"`
	HighWaterMarkReset                   types.Int64  `tfsdk:"high_water_mark_reset"`
	HostnameRewritePolicy                types.String `tfsdk:"hostname_rewrite_policy"`
	IgnoreDhcpOptionListRequest          types.Bool   `tfsdk:"ignore_dhcp_option_list_request"`
	IgnoreId                             types.String `tfsdk:"ignore_id"`
	IgnoreMacAddresses                   types.List   `tfsdk:"ignore_mac_addresses"`
	ImmediateFaConfiguration             types.Bool   `tfsdk:"immediate_fa_configuration"`
	Ipv6CaptureHostname                  types.Bool   `tfsdk:"ipv6_capture_hostname"`
	Ipv6DdnsDomainname                   types.String `tfsdk:"ipv6_ddns_domainname"`
	Ipv6DdnsEnableOptionFqdn             types.Bool   `tfsdk:"ipv6_ddns_enable_option_fqdn"`
	Ipv6DdnsServerAlwaysUpdates          types.Bool   `tfsdk:"ipv6_ddns_server_always_updates"`
	Ipv6DdnsTtl                          types.Int64  `tfsdk:"ipv6_ddns_ttl"`
	Ipv6DefaultPrefix                    types.String `tfsdk:"ipv6_default_prefix"`
	Ipv6DnsUpdateStyle                   types.String `tfsdk:"ipv6_dns_update_style"`
	Ipv6DomainName                       types.String `tfsdk:"ipv6_domain_name"`
	Ipv6DomainNameServers                types.List   `tfsdk:"ipv6_domain_name_servers"`
	Ipv6EnableDdns                       types.Bool   `tfsdk:"ipv6_enable_ddns"`
	Ipv6EnableGssTsig                    types.Bool   `tfsdk:"ipv6_enable_gss_tsig"`
	Ipv6EnableLeaseScavenging            types.Bool   `tfsdk:"ipv6_enable_lease_scavenging"`
	Ipv6EnableRetryUpdates               types.Bool   `tfsdk:"ipv6_enable_retry_updates"`
	Ipv6GenerateHostname                 types.Bool   `tfsdk:"ipv6_generate_hostname"`
	Ipv6GssTsigKeys                      types.List   `tfsdk:"ipv6_gss_tsig_keys"`
	Ipv6KdcServer                        types.String `tfsdk:"ipv6_kdc_server"`
	Ipv6LeaseScavengingTime              types.Int64  `tfsdk:"ipv6_lease_scavenging_time"`
	Ipv6MicrosoftCodePage                types.String `tfsdk:"ipv6_microsoft_code_page"`
	Ipv6Options                          types.List   `tfsdk:"ipv6_options"`
	Ipv6Prefixes                         types.List   `tfsdk:"ipv6_prefixes"`
	Ipv6RecycleLeases                    types.Bool   `tfsdk:"ipv6_recycle_leases"`
	Ipv6RememberExpiredClientAssociation types.Bool   `tfsdk:"ipv6_remember_expired_client_association"`
	Ipv6RetryUpdatesInterval             types.Int64  `tfsdk:"ipv6_retry_updates_interval"`
	Ipv6TxtRecordHandling                types.String `tfsdk:"ipv6_txt_record_handling"`
	Ipv6UpdateDnsOnLeaseRenewal          types.Bool   `tfsdk:"ipv6_update_dns_on_lease_renewal"`
	KdcServer                            types.String `tfsdk:"kdc_server"`
	LeaseLoggingMember                   types.String `tfsdk:"lease_logging_member"`
	LeasePerClientSettings               types.String `tfsdk:"lease_per_client_settings"`
	LeaseScavengeTime                    types.Int64  `tfsdk:"lease_scavenge_time"`
	LogLeaseEvents                       types.Bool   `tfsdk:"log_lease_events"`
	LogicFilterRules                     types.List   `tfsdk:"logic_filter_rules"`
	LowWaterMark                         types.Int64  `tfsdk:"low_water_mark"`
	LowWaterMarkReset                    types.Int64  `tfsdk:"low_water_mark_reset"`
	MicrosoftCodePage                    types.String `tfsdk:"microsoft_code_page"`
	Nextserver                           types.String `tfsdk:"nextserver"`
	Option60MatchRules                   types.List   `tfsdk:"option60_match_rules"`
	Options                              types.List   `tfsdk:"options"`
	PingCount                            types.Int64  `tfsdk:"ping_count"`
	PingTimeout                          types.Int64  `tfsdk:"ping_timeout"`
	PreferredLifetime                    types.Int64  `tfsdk:"preferred_lifetime"`
	PrefixLengthMode                     types.String `tfsdk:"prefix_length_mode"`
	ProtocolHostnameRewritePolicies      types.List   `tfsdk:"protocol_hostname_rewrite_policies"`
	PxeLeaseTime                         types.Int64  `tfsdk:"pxe_lease_time"`
	RecycleLeases                        types.Bool   `tfsdk:"recycle_leases"`
	RestartSetting                       types.Object `tfsdk:"restart_setting"`
	RetryDdnsUpdates                     types.Bool   `tfsdk:"retry_ddns_updates"`
	SyslogFacility                       types.String `tfsdk:"syslog_facility"`
	TxtRecordHandling                    types.String `tfsdk:"txt_record_handling"`
	UpdateDnsOnLeaseRenewal              types.Bool   `tfsdk:"update_dns_on_lease_renewal"`
	ValidLifetime                        types.Int64  `tfsdk:"valid_lifetime"`
}

var GridDhcppropertiesAttrTypes = map[string]attr.Type{
	"ref":                                      types.StringType,
	"uuid":                                     types.StringType,
	"authority":                                types.BoolType,
	"bootfile":                                 types.StringType,
	"bootserver":                               types.StringType,
	"capture_hostname":                         types.BoolType,
	"ddns_domainname":                          types.StringType,
	"ddns_generate_hostname":                   types.BoolType,
	"ddns_retry_interval":                      types.Int64Type,
	"ddns_server_always_updates":               types.BoolType,
	"ddns_ttl":                                 types.Int64Type,
	"ddns_update_fixed_addresses":              types.BoolType,
	"ddns_use_option81":                        types.BoolType,
	"deny_bootp":                               types.BoolType,
	"disable_all_nac_filters":                  types.BoolType,
	"dns_update_style":                         types.StringType,
	"email_list":                               types.ListType{ElemType: types.StringType},
	"enable_ddns":                              types.BoolType,
	"enable_dhcp_thresholds":                   types.BoolType,
	"enable_email_warnings":                    types.BoolType,
	"enable_fingerprint":                       types.BoolType,
	"enable_gss_tsig":                          types.BoolType,
	"enable_hostname_rewrite":                  types.BoolType,
	"enable_leasequery":                        types.BoolType,
	"enable_roaming_hosts":                     types.BoolType,
	"enable_snmp_warnings":                     types.BoolType,
	"format_log_option_82":                     types.StringType,
	"grid":                                     types.StringType,
	"gss_tsig_keys":                            types.ListType{ElemType: types.StringType},
	"high_water_mark":                          types.Int64Type,
	"high_water_mark_reset":                    types.Int64Type,
	"hostname_rewrite_policy":                  types.StringType,
	"ignore_dhcp_option_list_request":          types.BoolType,
	"ignore_id":                                types.StringType,
	"ignore_mac_addresses":                     types.ListType{ElemType: types.StringType},
	"immediate_fa_configuration":               types.BoolType,
	"ipv6_capture_hostname":                    types.BoolType,
	"ipv6_ddns_domainname":                     types.StringType,
	"ipv6_ddns_enable_option_fqdn":             types.BoolType,
	"ipv6_ddns_server_always_updates":          types.BoolType,
	"ipv6_ddns_ttl":                            types.Int64Type,
	"ipv6_default_prefix":                      types.StringType,
	"ipv6_dns_update_style":                    types.StringType,
	"ipv6_domain_name":                         types.StringType,
	"ipv6_domain_name_servers":                 types.ListType{ElemType: types.StringType},
	"ipv6_enable_ddns":                         types.BoolType,
	"ipv6_enable_gss_tsig":                     types.BoolType,
	"ipv6_enable_lease_scavenging":             types.BoolType,
	"ipv6_enable_retry_updates":                types.BoolType,
	"ipv6_generate_hostname":                   types.BoolType,
	"ipv6_gss_tsig_keys":                       types.ListType{ElemType: types.StringType},
	"ipv6_kdc_server":                          types.StringType,
	"ipv6_lease_scavenging_time":               types.Int64Type,
	"ipv6_microsoft_code_page":                 types.StringType,
	"ipv6_options":                             types.ListType{ElemType: types.ObjectType{AttrTypes: GridDhcppropertiesIpv6OptionsAttrTypes}},
	"ipv6_prefixes":                            types.ListType{ElemType: types.StringType},
	"ipv6_recycle_leases":                      types.BoolType,
	"ipv6_remember_expired_client_association": types.BoolType,
	"ipv6_retry_updates_interval":              types.Int64Type,
	"ipv6_txt_record_handling":                 types.StringType,
	"ipv6_update_dns_on_lease_renewal":         types.BoolType,
	"kdc_server":                               types.StringType,
	"lease_logging_member":                     types.StringType,
	"lease_per_client_settings":                types.StringType,
	"lease_scavenge_time":                      types.Int64Type,
	"log_lease_events":                         types.BoolType,
	"logic_filter_rules":                       types.ListType{ElemType: types.ObjectType{AttrTypes: GridDhcppropertiesLogicFilterRulesAttrTypes}},
	"low_water_mark":                           types.Int64Type,
	"low_water_mark_reset":                     types.Int64Type,
	"microsoft_code_page":                      types.StringType,
	"nextserver":                               types.StringType,
	"option60_match_rules":                     types.ListType{ElemType: types.ObjectType{AttrTypes: GridDhcppropertiesOption60MatchRulesAttrTypes}},
	"options":                                  types.ListType{ElemType: types.ObjectType{AttrTypes: GridDhcppropertiesOptionsAttrTypes}},
	"ping_count":                               types.Int64Type,
	"ping_timeout":                             types.Int64Type,
	"preferred_lifetime":                       types.Int64Type,
	"prefix_length_mode":                       types.StringType,
	"protocol_hostname_rewrite_policies":       types.ListType{ElemType: types.StringType},
	"pxe_lease_time":                           types.Int64Type,
	"recycle_leases":                           types.BoolType,
	"restart_setting":                          types.ObjectType{AttrTypes: GridDhcppropertiesRestartSettingAttrTypes},
	"retry_ddns_updates":                       types.BoolType,
	"syslog_facility":                          types.StringType,
	"txt_record_handling":                      types.StringType,
	"update_dns_on_lease_renewal":              types.BoolType,
	"valid_lifetime":                           types.Int64Type,
}

var GridDhcppropertiesResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The universally unique identifier (UUID) for the DHCP properties.",
	},
	"authority": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "The Grid-level authority flag. This flag specifies whether a DHCP server is authoritative for a domain.",
	},
	"bootfile": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of a file that DHCP clients need to boot. Some DHCP clients use BOOTP (bootstrap protocol) or include the boot file name option in their DHCPREQUEST messages.",
	},
	"bootserver": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the server on which a boot file is stored.",
	},
	"capture_hostname": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "The Grid-level capture hostname flag. Set this flag to capture the hostname and lease time when assigning a fixed address.",
	},
	"ddns_domainname": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The member DDNS domain name value.",
	},
	"ddns_generate_hostname": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the ability of a DHCP server to generate a host name and update DNS with this host name when it receives a DHCP REQUEST message that does not include a host name is enabled or not.",
	},
	"ddns_retry_interval": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Determines the retry interval when the DHCP server makes repeated attempts to send DDNS updates to a DNS server.",
	},
	"ddns_server_always_updates": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines that only the DHCP server is allowed to update DNS, regardless of the requests from the DHCP clients.",
	},
	"ddns_ttl": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The DDNS TTL (Dynamic DNS Time To Live) value specifies the number of seconds an IP address for the name is cached.",
	},
	"ddns_update_fixed_addresses": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the Grid DHCP server's ability to update the A and PTR records with a fixed address is enabled or not.",
	},
	"ddns_use_option81": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if support for option 81 is enabled or not.",
	},
	"deny_bootp": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if deny BOOTP is enabled or not.",
	},
	"disable_all_nac_filters": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If set to True, NAC filters will be disabled on the Infoblox Grid.",
	},
	"dns_update_style": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The update style for dynamic DNS updates.",
	},
	"email_list": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The Grid-level email_list value. Specify an e-mail address to which you want the Infoblox appliance to send e-mail notifications when the DHCP address usage for the grid crosses a threshold. You can create a list of several e-mail addresses.",
	},
	"enable_ddns": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the member DHCP server's ability to send DDNS updates is enabled or not.",
	},
	"enable_dhcp_thresholds": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Represents the watermarks above or below which address usage in a network is unexpected and might warrant your attention.",
	},
	"enable_email_warnings": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if e-mail warnings are enabled or disabled. When DHCP threshold is enabled and DHCP address usage crosses a watermark threshold, the appliance sends an e-mail notification to an administrator.",
	},
	"enable_fingerprint": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the fingerprint feature is enabled or not. If you enable this feature, the server will match a fingerprint for incoming lease requests.",
	},
	"enable_gss_tsig": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether all appliances are enabled to receive GSS-TSIG authenticated updates from DHCP clients.",
	},
	"enable_hostname_rewrite": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the Grid-level host name rewrite feature is enabled or not.",
	},
	"enable_leasequery": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if lease query is allowed or not.",
	},
	"enable_roaming_hosts": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if DHCP servers in a Grid support roaming hosts or not.",
	},
	"enable_snmp_warnings": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determined if the SNMP warnings on Grid-level are enabled or not. When DHCP threshold is enabled and DHCP address usage crosses a watermark threshold, the appliance sends an SNMP trap to the trap receiver that you defined you defined at the Grid member level.",
	},
	"format_log_option_82": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The format option for Option 82 logging.",
	},
	"grid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Determines the Grid that serves DHCP. This specifies a group of Infoblox appliances that are connected together to provide a single point of device administration and service configuration in a secure, highly available environment.",
	},
	"gss_tsig_keys": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of GSS-TSIG keys for a Grid DHCP object.",
	},
	"high_water_mark": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Determines the high watermark value of a Grid DHCP server. If the percentage of allocated addresses exceeds this watermark, the appliance makes a syslog entry and sends an e-mail notification (if enabled). Specifies the percentage of allocated addresses. The range is from 1 to 100.",
	},
	"high_water_mark_reset": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Determines the high watermark reset value of a member DHCP server. If the percentage of allocated addresses drops below this value, a corresponding SNMP trap is reset. Specifies the percentage of allocated addresses. The range is from 1 to 100. The high watermark reset value must be lower than the high watermark value.",
	},
	"hostname_rewrite_policy": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the default hostname rewrite policy, which is also in the protocol_hostname_rewrite_policies array.",
	},
	"ignore_dhcp_option_list_request": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the ignore DHCP option list request flag of a Grid DHCP is enabled or not. If this flag is set to true all available DHCP options will be returned to the client.",
	},
	"ignore_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Indicates whether the appliance will ignore DHCP client IDs or MAC addresses. Valid values are \"NONE\", \"CLIENT\", or \"MACADDR\". The default is \"NONE\".",
	},
	"ignore_mac_addresses": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "A list of MAC addresses the appliance will ignore.",
	},
	"immediate_fa_configuration": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the fixed address configuration takes effect immediately without DHCP service restart or not.",
	},
	"ipv6_capture_hostname": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the IPv6 host name and lease time is captured or not while assigning a fixed address.",
	},
	"ipv6_ddns_domainname": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The Grid-level DDNS domain name value.",
	},
	"ipv6_ddns_enable_option_fqdn": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Controls whether the FQDN option sent by the client is to be used, or if the server can automatically generate the FQDN.",
	},
	"ipv6_ddns_server_always_updates": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the server always updates DNS or updates only if requested by the client.",
	},
	"ipv6_ddns_ttl": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The Grid-level IPv6 DDNS TTL value.",
	},
	"ipv6_default_prefix": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The Grid-level IPv6 default prefix.",
	},
	"ipv6_dns_update_style": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The update style for dynamic DHCPv6 DNS updates.",
	},
	"ipv6_domain_name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The IPv6 domain name.",
	},
	"ipv6_domain_name_servers": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The comma separated list of domain name server addresses in IPv6 address format.",
	},
	"ipv6_enable_ddns": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if sending DDNS updates by the DHCPv6 server is enabled or not.",
	},
	"ipv6_enable_gss_tsig": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the all appliances are enabled to receive GSS-TSIG authenticated updates from DHCPv6 clients.",
	},
	"ipv6_enable_lease_scavenging": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Indicates whether DHCPv6 lease scavenging is enabled or disabled.",
	},
	"ipv6_enable_retry_updates": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the DHCPv6 server retries failed dynamic DNS updates or not.",
	},
	"ipv6_generate_hostname": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the server generates the hostname if it is not sent by the client.",
	},
	"ipv6_gss_tsig_keys": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of GSS-TSIG keys for a Grid DHCPv6 object.",
	},
	"ipv6_kdc_server": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The IPv6 address or FQDN of the Kerberos server for DHCPv6 GSS-TSIG authentication.",
	},
	"ipv6_lease_scavenging_time": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The Grid-level grace period (in seconds) to keep an expired lease before it is deleted by the scavenging process.",
	},
	"ipv6_microsoft_code_page": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The Grid-level Microsoft client DHCP IPv6 code page value. This value is the hostname translation code page for Microsoft DHCP IPv6 clients.",
	},
	"ipv6_options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: GridDhcppropertiesIpv6OptionsResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "An array of DHCP option dhcpoption structs that lists the DHCPv6 options associated with the object.",
	},
	"ipv6_prefixes": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The Grid-level list of IPv6 prefixes.",
	},
	"ipv6_recycle_leases": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the IPv6 recycle leases feature is enabled or not. If the feature is enabled, leases are kept in the Recycle Bin until one week after expiration. When the feature is disabled, the leases are irrecoverably deleted.",
	},
	"ipv6_remember_expired_client_association": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enable binding for expired DHCPv6 leases.",
	},
	"ipv6_retry_updates_interval": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Determines the retry interval when the member DHCPv6 server makes repeated attempts to send DDNS updates to a DNS server.",
	},
	"ipv6_txt_record_handling": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The Grid-level TXT record handling value. This value specifies how DHCPv6 should treat the TXT records when performing DNS updates.",
	},
	"ipv6_update_dns_on_lease_renewal": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Controls whether the DHCPv6 server updates DNS when an IPv6 DHCP lease is renewed.",
	},
	"kdc_server": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The IPv4 address or FQDN of the Kerberos server for DHCPv4 GSS-TSIG authentication.",
	},
	"lease_logging_member": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The Grid member on which you want to store the DHCP lease history log. Infoblox recommends that you dedicate a member other than the master as a logging member. If possible, use this member solely for storing the DHCP lease history log. If you do not select a member, no logging can occur.",
	},
	"lease_per_client_settings": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Defines how the appliance releases DHCP leases. Valid values are \"RELEASE_MACHING_ID\", \"NEVER_RELEASE\", or \"ONE_LEASE_PER_CLIENT\". The default is \"RELEASE_MATCHING_ID\".",
	},
	"lease_scavenge_time": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Determines the lease scavenging time value. When this field is set, the appliance permanently deletes the free and backup leases, that remain in the database beyond a specified period of time. To disable lease scavenging, set the parameter to -1. The minimum positive value must be greater than 86400 seconds (1 day).",
	},
	"log_lease_events": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "This value specifies whether the Grid DHCP members log lease events is enabled or not.",
	},
	"logic_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: GridDhcppropertiesLogicFilterRulesResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "This field contains the logic filters to be applied on the Infoblox Grid. This list corresponds to the match rules that are written to the dhcpd configuration file.",
	},
	"low_water_mark": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Determines the low watermark value. If the percent of allocated addresses drops below this watermark, the appliance makes a syslog entry and if enabled, sends an e-mail notification.",
	},
	"low_water_mark_reset": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Determines the low watermark reset value.If the percentage of allocated addresses exceeds this value, a corresponding SNMP trap is reset. A number that specifies the percentage of allocated addresses. The range is from 1 to 100. The low watermark reset value must be higher than the low watermark value.",
	},
	"microsoft_code_page": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The Microsoft client DHCP IPv4 code page value of a Grid. This value is the hostname translation code page for Microsoft DHCP IPv4 clients.",
	},
	"nextserver": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The next server value of a DHCP server. This value is the IP address or name of the boot file server on which the boot file is stored.",
	},
	"option60_match_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: GridDhcppropertiesOption60MatchRulesResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of option 60 match rules.",
	},
	"options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: GridDhcppropertiesOptionsResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object. Note that WAPI does not return special options 'routers', 'domain-name-servers', 'domain-name' and 'broadcast-address' with empty values for this object.",
	},
	"ping_count": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Specifies the number of pings that the Infoblox appliance sends to an IP address to verify that it is not in use. Values are range is from 0 to 10, where 0 disables pings.",
	},
	"ping_timeout": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Indicates the number of milliseconds the appliance waits for a response to its ping. Valid values are 100, 500, 1000, 2000, 3000, 4000 and 5000 milliseconds.",
	},
	"preferred_lifetime": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The preferred lifetime value.",
	},
	"prefix_length_mode": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The Prefix length mode for DHCPv6.",
	},
	"protocol_hostname_rewrite_policies": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of hostname rewrite policies.",
	},
	"pxe_lease_time": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Specifies the duration of time it takes a host to connect to a boot server, such as a TFTP server, and download the file it needs to boot. A 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached.",
	},
	"recycle_leases": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the recycle leases feature is enabled or not. If you enabled this feature, and then delete a DHCP range, the appliance stores active leases from this range up to one week after the leases expires.",
	},
	"restart_setting": schema.SingleNestedAttribute{
		Attributes: GridDhcppropertiesRestartSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"retry_ddns_updates": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Indicates whether the DHCP server makes repeated attempts to send DDNS updates to a DNS server.",
	},
	"syslog_facility": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The syslog facility is the location on the syslog server to which you want to sort the syslog messages.",
	},
	"txt_record_handling": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The Grid-level TXT record handling value. This value specifies how DHCP should treat the TXT records when performing DNS updates.",
	},
	"update_dns_on_lease_renewal": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Controls whether the DHCP server updates DNS when a DHCP lease is renewed.",
	},
	"valid_lifetime": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The valid lifetime for the Grid members.",
	},
}

func ExpandGridDhcpproperties(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridDhcpproperties {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridDhcppropertiesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridDhcppropertiesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridDhcpproperties {
	if m == nil {
		return nil
	}
	to := &grid.GridDhcpproperties{
		Ref:                                  flex.ExpandStringPointer(m.Ref),
		Uuid:                                 flex.ExpandStringPointer(m.Uuid),
		Authority:                            flex.ExpandBoolPointer(m.Authority),
		Bootfile:                             flex.ExpandStringPointer(m.Bootfile),
		Bootserver:                           flex.ExpandStringPointer(m.Bootserver),
		CaptureHostname:                      flex.ExpandBoolPointer(m.CaptureHostname),
		DdnsDomainname:                       flex.ExpandStringPointer(m.DdnsDomainname),
		DdnsGenerateHostname:                 flex.ExpandBoolPointer(m.DdnsGenerateHostname),
		DdnsRetryInterval:                    flex.ExpandInt64Pointer(m.DdnsRetryInterval),
		DdnsServerAlwaysUpdates:              flex.ExpandBoolPointer(m.DdnsServerAlwaysUpdates),
		DdnsTtl:                              flex.ExpandInt64Pointer(m.DdnsTtl),
		DdnsUpdateFixedAddresses:             flex.ExpandBoolPointer(m.DdnsUpdateFixedAddresses),
		DdnsUseOption81:                      flex.ExpandBoolPointer(m.DdnsUseOption81),
		DenyBootp:                            flex.ExpandBoolPointer(m.DenyBootp),
		DisableAllNacFilters:                 flex.ExpandBoolPointer(m.DisableAllNacFilters),
		DnsUpdateStyle:                       flex.ExpandStringPointer(m.DnsUpdateStyle),
		EmailList:                            flex.ExpandFrameworkListString(ctx, m.EmailList, diags),
		EnableDdns:                           flex.ExpandBoolPointer(m.EnableDdns),
		EnableDhcpThresholds:                 flex.ExpandBoolPointer(m.EnableDhcpThresholds),
		EnableEmailWarnings:                  flex.ExpandBoolPointer(m.EnableEmailWarnings),
		EnableFingerprint:                    flex.ExpandBoolPointer(m.EnableFingerprint),
		EnableGssTsig:                        flex.ExpandBoolPointer(m.EnableGssTsig),
		EnableHostnameRewrite:                flex.ExpandBoolPointer(m.EnableHostnameRewrite),
		EnableLeasequery:                     flex.ExpandBoolPointer(m.EnableLeasequery),
		EnableRoamingHosts:                   flex.ExpandBoolPointer(m.EnableRoamingHosts),
		EnableSnmpWarnings:                   flex.ExpandBoolPointer(m.EnableSnmpWarnings),
		FormatLogOption82:                    flex.ExpandStringPointer(m.FormatLogOption82),
		GssTsigKeys:                          flex.ExpandFrameworkListString(ctx, m.GssTsigKeys, diags),
		HighWaterMark:                        flex.ExpandInt64Pointer(m.HighWaterMark),
		HighWaterMarkReset:                   flex.ExpandInt64Pointer(m.HighWaterMarkReset),
		HostnameRewritePolicy:                flex.ExpandStringPointer(m.HostnameRewritePolicy),
		IgnoreDhcpOptionListRequest:          flex.ExpandBoolPointer(m.IgnoreDhcpOptionListRequest),
		IgnoreId:                             flex.ExpandStringPointer(m.IgnoreId),
		IgnoreMacAddresses:                   flex.ExpandFrameworkListString(ctx, m.IgnoreMacAddresses, diags),
		ImmediateFaConfiguration:             flex.ExpandBoolPointer(m.ImmediateFaConfiguration),
		Ipv6CaptureHostname:                  flex.ExpandBoolPointer(m.Ipv6CaptureHostname),
		Ipv6DdnsDomainname:                   flex.ExpandStringPointer(m.Ipv6DdnsDomainname),
		Ipv6DdnsEnableOptionFqdn:             flex.ExpandBoolPointer(m.Ipv6DdnsEnableOptionFqdn),
		Ipv6DdnsServerAlwaysUpdates:          flex.ExpandBoolPointer(m.Ipv6DdnsServerAlwaysUpdates),
		Ipv6DdnsTtl:                          flex.ExpandInt64Pointer(m.Ipv6DdnsTtl),
		Ipv6DefaultPrefix:                    flex.ExpandStringPointer(m.Ipv6DefaultPrefix),
		Ipv6DnsUpdateStyle:                   flex.ExpandStringPointer(m.Ipv6DnsUpdateStyle),
		Ipv6DomainName:                       flex.ExpandStringPointer(m.Ipv6DomainName),
		Ipv6DomainNameServers:                flex.ExpandFrameworkListString(ctx, m.Ipv6DomainNameServers, diags),
		Ipv6EnableDdns:                       flex.ExpandBoolPointer(m.Ipv6EnableDdns),
		Ipv6EnableGssTsig:                    flex.ExpandBoolPointer(m.Ipv6EnableGssTsig),
		Ipv6EnableLeaseScavenging:            flex.ExpandBoolPointer(m.Ipv6EnableLeaseScavenging),
		Ipv6EnableRetryUpdates:               flex.ExpandBoolPointer(m.Ipv6EnableRetryUpdates),
		Ipv6GenerateHostname:                 flex.ExpandBoolPointer(m.Ipv6GenerateHostname),
		Ipv6GssTsigKeys:                      flex.ExpandFrameworkListString(ctx, m.Ipv6GssTsigKeys, diags),
		Ipv6KdcServer:                        flex.ExpandStringPointer(m.Ipv6KdcServer),
		Ipv6LeaseScavengingTime:              flex.ExpandInt64Pointer(m.Ipv6LeaseScavengingTime),
		Ipv6MicrosoftCodePage:                flex.ExpandStringPointer(m.Ipv6MicrosoftCodePage),
		Ipv6Options:                          flex.ExpandFrameworkListNestedBlock(ctx, m.Ipv6Options, diags, ExpandGridDhcppropertiesIpv6Options),
		Ipv6Prefixes:                         flex.ExpandFrameworkListString(ctx, m.Ipv6Prefixes, diags),
		Ipv6RecycleLeases:                    flex.ExpandBoolPointer(m.Ipv6RecycleLeases),
		Ipv6RememberExpiredClientAssociation: flex.ExpandBoolPointer(m.Ipv6RememberExpiredClientAssociation),
		Ipv6RetryUpdatesInterval:             flex.ExpandInt64Pointer(m.Ipv6RetryUpdatesInterval),
		Ipv6TxtRecordHandling:                flex.ExpandStringPointer(m.Ipv6TxtRecordHandling),
		Ipv6UpdateDnsOnLeaseRenewal:          flex.ExpandBoolPointer(m.Ipv6UpdateDnsOnLeaseRenewal),
		KdcServer:                            flex.ExpandStringPointer(m.KdcServer),
		LeaseLoggingMember:                   flex.ExpandStringPointer(m.LeaseLoggingMember),
		LeasePerClientSettings:               flex.ExpandStringPointer(m.LeasePerClientSettings),
		LeaseScavengeTime:                    flex.ExpandInt64Pointer(m.LeaseScavengeTime),
		LogLeaseEvents:                       flex.ExpandBoolPointer(m.LogLeaseEvents),
		LogicFilterRules:                     flex.ExpandFrameworkListNestedBlock(ctx, m.LogicFilterRules, diags, ExpandGridDhcppropertiesLogicFilterRules),
		LowWaterMark:                         flex.ExpandInt64Pointer(m.LowWaterMark),
		LowWaterMarkReset:                    flex.ExpandInt64Pointer(m.LowWaterMarkReset),
		MicrosoftCodePage:                    flex.ExpandStringPointer(m.MicrosoftCodePage),
		Nextserver:                           flex.ExpandStringPointer(m.Nextserver),
		Option60MatchRules:                   flex.ExpandFrameworkListNestedBlock(ctx, m.Option60MatchRules, diags, ExpandGridDhcppropertiesOption60MatchRules),
		Options:                              flex.ExpandFrameworkListNestedBlock(ctx, m.Options, diags, ExpandGridDhcppropertiesOptions),
		PingCount:                            flex.ExpandInt64Pointer(m.PingCount),
		PingTimeout:                          flex.ExpandInt64Pointer(m.PingTimeout),
		PreferredLifetime:                    flex.ExpandInt64Pointer(m.PreferredLifetime),
		PrefixLengthMode:                     flex.ExpandStringPointer(m.PrefixLengthMode),
		ProtocolHostnameRewritePolicies:      flex.ExpandFrameworkListString(ctx, m.ProtocolHostnameRewritePolicies, diags),
		PxeLeaseTime:                         flex.ExpandInt64Pointer(m.PxeLeaseTime),
		RecycleLeases:                        flex.ExpandBoolPointer(m.RecycleLeases),
		RestartSetting:                       ExpandGridDhcppropertiesRestartSetting(ctx, m.RestartSetting, diags),
		RetryDdnsUpdates:                     flex.ExpandBoolPointer(m.RetryDdnsUpdates),
		SyslogFacility:                       flex.ExpandStringPointer(m.SyslogFacility),
		TxtRecordHandling:                    flex.ExpandStringPointer(m.TxtRecordHandling),
		UpdateDnsOnLeaseRenewal:              flex.ExpandBoolPointer(m.UpdateDnsOnLeaseRenewal),
		ValidLifetime:                        flex.ExpandInt64Pointer(m.ValidLifetime),
	}
	return to
}

func FlattenGridDhcpproperties(ctx context.Context, from *grid.GridDhcpproperties, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridDhcppropertiesAttrTypes)
	}
	m := GridDhcppropertiesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridDhcppropertiesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridDhcppropertiesModel) Flatten(ctx context.Context, from *grid.GridDhcpproperties, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridDhcppropertiesModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Authority = types.BoolPointerValue(from.Authority)
	m.Bootfile = flex.FlattenStringPointer(from.Bootfile)
	m.Bootserver = flex.FlattenStringPointer(from.Bootserver)
	m.CaptureHostname = types.BoolPointerValue(from.CaptureHostname)
	m.DdnsDomainname = flex.FlattenStringPointer(from.DdnsDomainname)
	m.DdnsGenerateHostname = types.BoolPointerValue(from.DdnsGenerateHostname)
	m.DdnsRetryInterval = flex.FlattenInt64Pointer(from.DdnsRetryInterval)
	m.DdnsServerAlwaysUpdates = types.BoolPointerValue(from.DdnsServerAlwaysUpdates)
	m.DdnsTtl = flex.FlattenInt64Pointer(from.DdnsTtl)
	m.DdnsUpdateFixedAddresses = types.BoolPointerValue(from.DdnsUpdateFixedAddresses)
	m.DdnsUseOption81 = types.BoolPointerValue(from.DdnsUseOption81)
	m.DenyBootp = types.BoolPointerValue(from.DenyBootp)
	m.DisableAllNacFilters = types.BoolPointerValue(from.DisableAllNacFilters)
	m.DnsUpdateStyle = flex.FlattenStringPointer(from.DnsUpdateStyle)
	m.EmailList = flex.FlattenFrameworkListString(ctx, from.EmailList, diags)
	m.EnableDdns = types.BoolPointerValue(from.EnableDdns)
	m.EnableDhcpThresholds = types.BoolPointerValue(from.EnableDhcpThresholds)
	m.EnableEmailWarnings = types.BoolPointerValue(from.EnableEmailWarnings)
	m.EnableFingerprint = types.BoolPointerValue(from.EnableFingerprint)
	m.EnableGssTsig = types.BoolPointerValue(from.EnableGssTsig)
	m.EnableHostnameRewrite = types.BoolPointerValue(from.EnableHostnameRewrite)
	m.EnableLeasequery = types.BoolPointerValue(from.EnableLeasequery)
	m.EnableRoamingHosts = types.BoolPointerValue(from.EnableRoamingHosts)
	m.EnableSnmpWarnings = types.BoolPointerValue(from.EnableSnmpWarnings)
	m.FormatLogOption82 = flex.FlattenStringPointer(from.FormatLogOption82)
	m.Grid = flex.FlattenStringPointer(from.Grid)
	m.GssTsigKeys = flex.FlattenFrameworkListString(ctx, from.GssTsigKeys, diags)
	m.HighWaterMark = flex.FlattenInt64Pointer(from.HighWaterMark)
	m.HighWaterMarkReset = flex.FlattenInt64Pointer(from.HighWaterMarkReset)
	m.HostnameRewritePolicy = flex.FlattenStringPointer(from.HostnameRewritePolicy)
	m.IgnoreDhcpOptionListRequest = types.BoolPointerValue(from.IgnoreDhcpOptionListRequest)
	m.IgnoreId = flex.FlattenStringPointer(from.IgnoreId)
	m.IgnoreMacAddresses = flex.FlattenFrameworkListString(ctx, from.IgnoreMacAddresses, diags)
	m.ImmediateFaConfiguration = types.BoolPointerValue(from.ImmediateFaConfiguration)
	m.Ipv6CaptureHostname = types.BoolPointerValue(from.Ipv6CaptureHostname)
	m.Ipv6DdnsDomainname = flex.FlattenStringPointer(from.Ipv6DdnsDomainname)
	m.Ipv6DdnsEnableOptionFqdn = types.BoolPointerValue(from.Ipv6DdnsEnableOptionFqdn)
	m.Ipv6DdnsServerAlwaysUpdates = types.BoolPointerValue(from.Ipv6DdnsServerAlwaysUpdates)
	m.Ipv6DdnsTtl = flex.FlattenInt64Pointer(from.Ipv6DdnsTtl)
	m.Ipv6DefaultPrefix = flex.FlattenStringPointer(from.Ipv6DefaultPrefix)
	m.Ipv6DnsUpdateStyle = flex.FlattenStringPointer(from.Ipv6DnsUpdateStyle)
	m.Ipv6DomainName = flex.FlattenStringPointer(from.Ipv6DomainName)
	m.Ipv6DomainNameServers = flex.FlattenFrameworkListString(ctx, from.Ipv6DomainNameServers, diags)
	m.Ipv6EnableDdns = types.BoolPointerValue(from.Ipv6EnableDdns)
	m.Ipv6EnableGssTsig = types.BoolPointerValue(from.Ipv6EnableGssTsig)
	m.Ipv6EnableLeaseScavenging = types.BoolPointerValue(from.Ipv6EnableLeaseScavenging)
	m.Ipv6EnableRetryUpdates = types.BoolPointerValue(from.Ipv6EnableRetryUpdates)
	m.Ipv6GenerateHostname = types.BoolPointerValue(from.Ipv6GenerateHostname)
	m.Ipv6GssTsigKeys = flex.FlattenFrameworkListString(ctx, from.Ipv6GssTsigKeys, diags)
	m.Ipv6KdcServer = flex.FlattenStringPointer(from.Ipv6KdcServer)
	m.Ipv6LeaseScavengingTime = flex.FlattenInt64Pointer(from.Ipv6LeaseScavengingTime)
	m.Ipv6MicrosoftCodePage = flex.FlattenStringPointer(from.Ipv6MicrosoftCodePage)
	m.Ipv6Options = flex.FlattenFrameworkListNestedBlock(ctx, from.Ipv6Options, GridDhcppropertiesIpv6OptionsAttrTypes, diags, FlattenGridDhcppropertiesIpv6Options)
	m.Ipv6Prefixes = flex.FlattenFrameworkListString(ctx, from.Ipv6Prefixes, diags)
	m.Ipv6RecycleLeases = types.BoolPointerValue(from.Ipv6RecycleLeases)
	m.Ipv6RememberExpiredClientAssociation = types.BoolPointerValue(from.Ipv6RememberExpiredClientAssociation)
	m.Ipv6RetryUpdatesInterval = flex.FlattenInt64Pointer(from.Ipv6RetryUpdatesInterval)
	m.Ipv6TxtRecordHandling = flex.FlattenStringPointer(from.Ipv6TxtRecordHandling)
	m.Ipv6UpdateDnsOnLeaseRenewal = types.BoolPointerValue(from.Ipv6UpdateDnsOnLeaseRenewal)
	m.KdcServer = flex.FlattenStringPointer(from.KdcServer)
	m.LeaseLoggingMember = flex.FlattenStringPointer(from.LeaseLoggingMember)
	m.LeasePerClientSettings = flex.FlattenStringPointer(from.LeasePerClientSettings)
	m.LeaseScavengeTime = flex.FlattenInt64Pointer(from.LeaseScavengeTime)
	m.LogLeaseEvents = types.BoolPointerValue(from.LogLeaseEvents)
	m.LogicFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.LogicFilterRules, GridDhcppropertiesLogicFilterRulesAttrTypes, diags, FlattenGridDhcppropertiesLogicFilterRules)
	m.LowWaterMark = flex.FlattenInt64Pointer(from.LowWaterMark)
	m.LowWaterMarkReset = flex.FlattenInt64Pointer(from.LowWaterMarkReset)
	m.MicrosoftCodePage = flex.FlattenStringPointer(from.MicrosoftCodePage)
	m.Nextserver = flex.FlattenStringPointer(from.Nextserver)
	m.Option60MatchRules = flex.FlattenFrameworkListNestedBlock(ctx, from.Option60MatchRules, GridDhcppropertiesOption60MatchRulesAttrTypes, diags, FlattenGridDhcppropertiesOption60MatchRules)
	m.Options = flex.FlattenFrameworkListNestedBlock(ctx, from.Options, GridDhcppropertiesOptionsAttrTypes, diags, FlattenGridDhcppropertiesOptions)
	m.PingCount = flex.FlattenInt64Pointer(from.PingCount)
	m.PingTimeout = flex.FlattenInt64Pointer(from.PingTimeout)
	m.PreferredLifetime = flex.FlattenInt64Pointer(from.PreferredLifetime)
	m.PrefixLengthMode = flex.FlattenStringPointer(from.PrefixLengthMode)
	m.ProtocolHostnameRewritePolicies = flex.FlattenFrameworkListString(ctx, from.ProtocolHostnameRewritePolicies, diags)
	m.PxeLeaseTime = flex.FlattenInt64Pointer(from.PxeLeaseTime)
	m.RecycleLeases = types.BoolPointerValue(from.RecycleLeases)
	m.RestartSetting = FlattenGridDhcppropertiesRestartSetting(ctx, from.RestartSetting, diags)
	m.RetryDdnsUpdates = types.BoolPointerValue(from.RetryDdnsUpdates)
	m.SyslogFacility = flex.FlattenStringPointer(from.SyslogFacility)
	m.TxtRecordHandling = flex.FlattenStringPointer(from.TxtRecordHandling)
	m.UpdateDnsOnLeaseRenewal = types.BoolPointerValue(from.UpdateDnsOnLeaseRenewal)
	m.ValidLifetime = flex.FlattenInt64Pointer(from.ValidLifetime)
}
