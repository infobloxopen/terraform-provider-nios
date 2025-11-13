package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/import"
)

type MemberDhcppropertiesModel struct {
	Ref                                   types.String `tfsdk:"ref"`
	AuthServerGroup                       types.String `tfsdk:"auth_server_group"`
	AuthnCaptivePortal                    types.String `tfsdk:"authn_captive_portal"`
	AuthnCaptivePortalAuthenticatedFilter types.String `tfsdk:"authn_captive_portal_authenticated_filter"`
	AuthnCaptivePortalEnabled             types.Bool   `tfsdk:"authn_captive_portal_enabled"`
	AuthnCaptivePortalGuestFilter         types.String `tfsdk:"authn_captive_portal_guest_filter"`
	AuthnServerGroupEnabled               types.Bool   `tfsdk:"authn_server_group_enabled"`
	Authority                             types.Bool   `tfsdk:"authority"`
	Bootfile                              types.String `tfsdk:"bootfile"`
	Bootserver                            types.String `tfsdk:"bootserver"`
	DdnsDomainname                        types.String `tfsdk:"ddns_domainname"`
	DdnsGenerateHostname                  types.Bool   `tfsdk:"ddns_generate_hostname"`
	DdnsRetryInterval                     types.Int64  `tfsdk:"ddns_retry_interval"`
	DdnsServerAlwaysUpdates               types.Bool   `tfsdk:"ddns_server_always_updates"`
	DdnsTtl                               types.Int64  `tfsdk:"ddns_ttl"`
	DdnsUpdateFixedAddresses              types.Bool   `tfsdk:"ddns_update_fixed_addresses"`
	DdnsUseOption81                       types.Bool   `tfsdk:"ddns_use_option81"`
	DdnsZonePrimaries                     types.List   `tfsdk:"ddns_zone_primaries"`
	DenyBootp                             types.Bool   `tfsdk:"deny_bootp"`
	DhcpUtilization                       types.Int64  `tfsdk:"dhcp_utilization"`
	DhcpUtilizationStatus                 types.String `tfsdk:"dhcp_utilization_status"`
	DnsUpdateStyle                        types.String `tfsdk:"dns_update_style"`
	DynamicHosts                          types.Int64  `tfsdk:"dynamic_hosts"`
	EmailList                             types.List   `tfsdk:"email_list"`
	EnableDdns                            types.Bool   `tfsdk:"enable_ddns"`
	EnableDhcp                            types.Bool   `tfsdk:"enable_dhcp"`
	EnableDhcpOnIpv6Lan2                  types.Bool   `tfsdk:"enable_dhcp_on_ipv6_lan2"`
	EnableDhcpOnLan2                      types.Bool   `tfsdk:"enable_dhcp_on_lan2"`
	EnableDhcpThresholds                  types.Bool   `tfsdk:"enable_dhcp_thresholds"`
	EnableDhcpv6Service                   types.Bool   `tfsdk:"enable_dhcpv6_service"`
	EnableEmailWarnings                   types.Bool   `tfsdk:"enable_email_warnings"`
	EnableFingerprint                     types.Bool   `tfsdk:"enable_fingerprint"`
	EnableGssTsig                         types.Bool   `tfsdk:"enable_gss_tsig"`
	EnableHostnameRewrite                 types.Bool   `tfsdk:"enable_hostname_rewrite"`
	EnableLeasequery                      types.Bool   `tfsdk:"enable_leasequery"`
	EnableSnmpWarnings                    types.Bool   `tfsdk:"enable_snmp_warnings"`
	ExtAttrs                              types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll                           types.Map    `tfsdk:"extattrs_all"`
	GssTsigKeys                           types.List   `tfsdk:"gss_tsig_keys"`
	HighWaterMark                         types.Int64  `tfsdk:"high_water_mark"`
	HighWaterMarkReset                    types.Int64  `tfsdk:"high_water_mark_reset"`
	HostName                              types.String `tfsdk:"host_name"`
	HostnameRewritePolicy                 types.String `tfsdk:"hostname_rewrite_policy"`
	IgnoreDhcpOptionListRequest           types.Bool   `tfsdk:"ignore_dhcp_option_list_request"`
	IgnoreId                              types.String `tfsdk:"ignore_id"`
	IgnoreMacAddresses                    types.List   `tfsdk:"ignore_mac_addresses"`
	ImmediateFaConfiguration              types.Bool   `tfsdk:"immediate_fa_configuration"`
	Ipv4addr                              types.String `tfsdk:"ipv4addr"`
	Ipv6DdnsDomainname                    types.String `tfsdk:"ipv6_ddns_domainname"`
	Ipv6DdnsEnableOptionFqdn              types.Bool   `tfsdk:"ipv6_ddns_enable_option_fqdn"`
	Ipv6DdnsHostname                      types.String `tfsdk:"ipv6_ddns_hostname"`
	Ipv6DdnsServerAlwaysUpdates           types.Bool   `tfsdk:"ipv6_ddns_server_always_updates"`
	Ipv6DdnsTtl                           types.Int64  `tfsdk:"ipv6_ddns_ttl"`
	Ipv6DnsUpdateStyle                    types.String `tfsdk:"ipv6_dns_update_style"`
	Ipv6DomainName                        types.String `tfsdk:"ipv6_domain_name"`
	Ipv6DomainNameServers                 types.List   `tfsdk:"ipv6_domain_name_servers"`
	Ipv6EnableDdns                        types.Bool   `tfsdk:"ipv6_enable_ddns"`
	Ipv6EnableGssTsig                     types.Bool   `tfsdk:"ipv6_enable_gss_tsig"`
	Ipv6EnableLeaseScavenging             types.Bool   `tfsdk:"ipv6_enable_lease_scavenging"`
	Ipv6EnableRetryUpdates                types.Bool   `tfsdk:"ipv6_enable_retry_updates"`
	Ipv6GenerateHostname                  types.Bool   `tfsdk:"ipv6_generate_hostname"`
	Ipv6GssTsigKeys                       types.List   `tfsdk:"ipv6_gss_tsig_keys"`
	Ipv6KdcServer                         types.String `tfsdk:"ipv6_kdc_server"`
	Ipv6LeaseScavengingTime               types.Int64  `tfsdk:"ipv6_lease_scavenging_time"`
	Ipv6MicrosoftCodePage                 types.String `tfsdk:"ipv6_microsoft_code_page"`
	Ipv6Options                           types.List   `tfsdk:"ipv6_options"`
	Ipv6RecycleLeases                     types.Bool   `tfsdk:"ipv6_recycle_leases"`
	Ipv6RememberExpiredClientAssociation  types.Bool   `tfsdk:"ipv6_remember_expired_client_association"`
	Ipv6RetryUpdatesInterval              types.Int64  `tfsdk:"ipv6_retry_updates_interval"`
	Ipv6ServerDuid                        types.String `tfsdk:"ipv6_server_duid"`
	Ipv6UpdateDnsOnLeaseRenewal           types.Bool   `tfsdk:"ipv6_update_dns_on_lease_renewal"`
	Ipv6addr                              types.String `tfsdk:"ipv6addr"`
	KdcServer                             types.String `tfsdk:"kdc_server"`
	LeasePerClientSettings                types.String `tfsdk:"lease_per_client_settings"`
	LeaseScavengeTime                     types.Int64  `tfsdk:"lease_scavenge_time"`
	LogLeaseEvents                        types.Bool   `tfsdk:"log_lease_events"`
	LogicFilterRules                      types.List   `tfsdk:"logic_filter_rules"`
	LowWaterMark                          types.Int64  `tfsdk:"low_water_mark"`
	LowWaterMarkReset                     types.Int64  `tfsdk:"low_water_mark_reset"`
	MicrosoftCodePage                     types.String `tfsdk:"microsoft_code_page"`
	Nextserver                            types.String `tfsdk:"nextserver"`
	Option60MatchRules                    types.List   `tfsdk:"option60_match_rules"`
	Options                               types.List   `tfsdk:"options"`
	PingCount                             types.Int64  `tfsdk:"ping_count"`
	PingTimeout                           types.Int64  `tfsdk:"ping_timeout"`
	PreferredLifetime                     types.Int64  `tfsdk:"preferred_lifetime"`
	PrefixLengthMode                      types.String `tfsdk:"prefix_length_mode"`
	PxeLeaseTime                          types.Int64  `tfsdk:"pxe_lease_time"`
	RecycleLeases                         types.Bool   `tfsdk:"recycle_leases"`
	RetryDdnsUpdates                      types.Bool   `tfsdk:"retry_ddns_updates"`
	StaticHosts                           types.Int64  `tfsdk:"static_hosts"`
	SyslogFacility                        types.String `tfsdk:"syslog_facility"`
	TotalHosts                            types.Int64  `tfsdk:"total_hosts"`
	UpdateDnsOnLeaseRenewal               types.Bool   `tfsdk:"update_dns_on_lease_renewal"`
	UseAuthority                          types.Bool   `tfsdk:"use_authority"`
	UseBootfile                           types.Bool   `tfsdk:"use_bootfile"`
	UseBootserver                         types.Bool   `tfsdk:"use_bootserver"`
	UseDdnsDomainname                     types.Bool   `tfsdk:"use_ddns_domainname"`
	UseDdnsGenerateHostname               types.Bool   `tfsdk:"use_ddns_generate_hostname"`
	UseDdnsTtl                            types.Bool   `tfsdk:"use_ddns_ttl"`
	UseDdnsUpdateFixedAddresses           types.Bool   `tfsdk:"use_ddns_update_fixed_addresses"`
	UseDdnsUseOption81                    types.Bool   `tfsdk:"use_ddns_use_option81"`
	UseDenyBootp                          types.Bool   `tfsdk:"use_deny_bootp"`
	UseDnsUpdateStyle                     types.Bool   `tfsdk:"use_dns_update_style"`
	UseEmailList                          types.Bool   `tfsdk:"use_email_list"`
	UseEnableDdns                         types.Bool   `tfsdk:"use_enable_ddns"`
	UseEnableDhcpThresholds               types.Bool   `tfsdk:"use_enable_dhcp_thresholds"`
	UseEnableFingerprint                  types.Bool   `tfsdk:"use_enable_fingerprint"`
	UseEnableGssTsig                      types.Bool   `tfsdk:"use_enable_gss_tsig"`
	UseEnableHostnameRewrite              types.Bool   `tfsdk:"use_enable_hostname_rewrite"`
	UseEnableLeasequery                   types.Bool   `tfsdk:"use_enable_leasequery"`
	UseEnableOneLeasePerClient            types.Bool   `tfsdk:"use_enable_one_lease_per_client"`
	UseGssTsigKeys                        types.Bool   `tfsdk:"use_gss_tsig_keys"`
	UseIgnoreDhcpOptionListRequest        types.Bool   `tfsdk:"use_ignore_dhcp_option_list_request"`
	UseIgnoreId                           types.Bool   `tfsdk:"use_ignore_id"`
	UseImmediateFaConfiguration           types.Bool   `tfsdk:"use_immediate_fa_configuration"`
	UseIpv6DdnsDomainname                 types.Bool   `tfsdk:"use_ipv6_ddns_domainname"`
	UseIpv6DdnsEnableOptionFqdn           types.Bool   `tfsdk:"use_ipv6_ddns_enable_option_fqdn"`
	UseIpv6DdnsHostname                   types.Bool   `tfsdk:"use_ipv6_ddns_hostname"`
	UseIpv6DdnsTtl                        types.Bool   `tfsdk:"use_ipv6_ddns_ttl"`
	UseIpv6DnsUpdateStyle                 types.Bool   `tfsdk:"use_ipv6_dns_update_style"`
	UseIpv6DomainName                     types.Bool   `tfsdk:"use_ipv6_domain_name"`
	UseIpv6DomainNameServers              types.Bool   `tfsdk:"use_ipv6_domain_name_servers"`
	UseIpv6EnableDdns                     types.Bool   `tfsdk:"use_ipv6_enable_ddns"`
	UseIpv6EnableGssTsig                  types.Bool   `tfsdk:"use_ipv6_enable_gss_tsig"`
	UseIpv6EnableRetryUpdates             types.Bool   `tfsdk:"use_ipv6_enable_retry_updates"`
	UseIpv6GenerateHostname               types.Bool   `tfsdk:"use_ipv6_generate_hostname"`
	UseIpv6GssTsigKeys                    types.Bool   `tfsdk:"use_ipv6_gss_tsig_keys"`
	UseIpv6LeaseScavenging                types.Bool   `tfsdk:"use_ipv6_lease_scavenging"`
	UseIpv6MicrosoftCodePage              types.Bool   `tfsdk:"use_ipv6_microsoft_code_page"`
	UseIpv6Options                        types.Bool   `tfsdk:"use_ipv6_options"`
	UseIpv6RecycleLeases                  types.Bool   `tfsdk:"use_ipv6_recycle_leases"`
	UseIpv6UpdateDnsOnLeaseRenewal        types.Bool   `tfsdk:"use_ipv6_update_dns_on_lease_renewal"`
	UseLeasePerClientSettings             types.Bool   `tfsdk:"use_lease_per_client_settings"`
	UseLeaseScavengeTime                  types.Bool   `tfsdk:"use_lease_scavenge_time"`
	UseLogLeaseEvents                     types.Bool   `tfsdk:"use_log_lease_events"`
	UseLogicFilterRules                   types.Bool   `tfsdk:"use_logic_filter_rules"`
	UseMicrosoftCodePage                  types.Bool   `tfsdk:"use_microsoft_code_page"`
	UseNextserver                         types.Bool   `tfsdk:"use_nextserver"`
	UseOptions                            types.Bool   `tfsdk:"use_options"`
	UsePingCount                          types.Bool   `tfsdk:"use_ping_count"`
	UsePingTimeout                        types.Bool   `tfsdk:"use_ping_timeout"`
	UsePreferredLifetime                  types.Bool   `tfsdk:"use_preferred_lifetime"`
	UsePrefixLengthMode                   types.Bool   `tfsdk:"use_prefix_length_mode"`
	UsePxeLeaseTime                       types.Bool   `tfsdk:"use_pxe_lease_time"`
	UseRecycleLeases                      types.Bool   `tfsdk:"use_recycle_leases"`
	UseRetryDdnsUpdates                   types.Bool   `tfsdk:"use_retry_ddns_updates"`
	UseSyslogFacility                     types.Bool   `tfsdk:"use_syslog_facility"`
	UseUpdateDnsOnLeaseRenewal            types.Bool   `tfsdk:"use_update_dns_on_lease_renewal"`
	UseValidLifetime                      types.Bool   `tfsdk:"use_valid_lifetime"`
	ValidLifetime                         types.Int64  `tfsdk:"valid_lifetime"`
}

var MemberDhcppropertiesAttrTypes = map[string]attr.Type{
	"ref":                  types.StringType,
	"auth_server_group":    types.StringType,
	"authn_captive_portal": types.StringType,
	"authn_captive_portal_authenticated_filter": types.StringType,
	"authn_captive_portal_enabled":              types.BoolType,
	"authn_captive_portal_guest_filter":         types.StringType,
	"authn_server_group_enabled":                types.BoolType,
	"authority":                                 types.BoolType,
	"bootfile":                                  types.StringType,
	"bootserver":                                types.StringType,
	"ddns_domainname":                           types.StringType,
	"ddns_generate_hostname":                    types.BoolType,
	"ddns_retry_interval":                       types.Int64Type,
	"ddns_server_always_updates":                types.BoolType,
	"ddns_ttl":                                  types.Int64Type,
	"ddns_update_fixed_addresses":               types.BoolType,
	"ddns_use_option81":                         types.BoolType,
	"ddns_zone_primaries":                       types.ListType{ElemType: types.ObjectType{AttrTypes: MemberDhcppropertiesDdnsZonePrimariesAttrTypes}},
	"deny_bootp":                                types.BoolType,
	"dhcp_utilization":                          types.Int64Type,
	"dhcp_utilization_status":                   types.StringType,
	"dns_update_style":                          types.StringType,
	"dynamic_hosts":                             types.Int64Type,
	"email_list":                                types.ListType{ElemType: types.StringType},
	"enable_ddns":                               types.BoolType,
	"enable_dhcp":                               types.BoolType,
	"enable_dhcp_on_ipv6_lan2":                  types.BoolType,
	"enable_dhcp_on_lan2":                       types.BoolType,
	"enable_dhcp_thresholds":                    types.BoolType,
	"enable_dhcpv6_service":                     types.BoolType,
	"enable_email_warnings":                     types.BoolType,
	"enable_fingerprint":                        types.BoolType,
	"enable_gss_tsig":                           types.BoolType,
	"enable_hostname_rewrite":                   types.BoolType,
	"enable_leasequery":                         types.BoolType,
	"enable_snmp_warnings":                      types.BoolType,
	"extattrs":                                  types.MapType{ElemType: types.StringType},
	"extattrs_all":                              types.MapType{ElemType: types.StringType},
	"gss_tsig_keys":                             types.ListType{ElemType: types.StringType},
	"high_water_mark":                           types.Int64Type,
	"high_water_mark_reset":                     types.Int64Type,
	"host_name":                                 types.StringType,
	"hostname_rewrite_policy":                   types.StringType,
	"ignore_dhcp_option_list_request":           types.BoolType,
	"ignore_id":                                 types.StringType,
	"ignore_mac_addresses":                      types.ListType{ElemType: types.StringType},
	"immediate_fa_configuration":                types.BoolType,
	"ipv4addr":                                  types.StringType,
	"ipv6_ddns_domainname":                      types.StringType,
	"ipv6_ddns_enable_option_fqdn":              types.BoolType,
	"ipv6_ddns_hostname":                        types.StringType,
	"ipv6_ddns_server_always_updates":           types.BoolType,
	"ipv6_ddns_ttl":                             types.Int64Type,
	"ipv6_dns_update_style":                     types.StringType,
	"ipv6_domain_name":                          types.StringType,
	"ipv6_domain_name_servers":                  types.ListType{ElemType: types.StringType},
	"ipv6_enable_ddns":                          types.BoolType,
	"ipv6_enable_gss_tsig":                      types.BoolType,
	"ipv6_enable_lease_scavenging":              types.BoolType,
	"ipv6_enable_retry_updates":                 types.BoolType,
	"ipv6_generate_hostname":                    types.BoolType,
	"ipv6_gss_tsig_keys":                        types.ListType{ElemType: types.StringType},
	"ipv6_kdc_server":                           types.StringType,
	"ipv6_lease_scavenging_time":                types.Int64Type,
	"ipv6_microsoft_code_page":                  types.StringType,
	"ipv6_options":                              types.ListType{ElemType: types.ObjectType{AttrTypes: MemberDhcppropertiesIpv6OptionsAttrTypes}},
	"ipv6_recycle_leases":                       types.BoolType,
	"ipv6_remember_expired_client_association":  types.BoolType,
	"ipv6_retry_updates_interval":               types.Int64Type,
	"ipv6_server_duid":                          types.StringType,
	"ipv6_update_dns_on_lease_renewal":          types.BoolType,
	"ipv6addr":                                  types.StringType,
	"kdc_server":                                types.StringType,
	"lease_per_client_settings":                 types.StringType,
	"lease_scavenge_time":                       types.Int64Type,
	"log_lease_events":                          types.BoolType,
	"logic_filter_rules":                        types.ListType{ElemType: types.ObjectType{AttrTypes: MemberDhcppropertiesLogicFilterRulesAttrTypes}},
	"low_water_mark":                            types.Int64Type,
	"low_water_mark_reset":                      types.Int64Type,
	"microsoft_code_page":                       types.StringType,
	"nextserver":                                types.StringType,
	"option60_match_rules":                      types.ListType{ElemType: types.ObjectType{AttrTypes: MemberDhcppropertiesOption60MatchRulesAttrTypes}},
	"options":                                   types.ListType{ElemType: types.ObjectType{AttrTypes: MemberDhcppropertiesOptionsAttrTypes}},
	"ping_count":                                types.Int64Type,
	"ping_timeout":                              types.Int64Type,
	"preferred_lifetime":                        types.Int64Type,
	"prefix_length_mode":                        types.StringType,
	"pxe_lease_time":                            types.Int64Type,
	"recycle_leases":                            types.BoolType,
	"retry_ddns_updates":                        types.BoolType,
	"static_hosts":                              types.Int64Type,
	"syslog_facility":                           types.StringType,
	"total_hosts":                               types.Int64Type,
	"update_dns_on_lease_renewal":               types.BoolType,
	"use_authority":                             types.BoolType,
	"use_bootfile":                              types.BoolType,
	"use_bootserver":                            types.BoolType,
	"use_ddns_domainname":                       types.BoolType,
	"use_ddns_generate_hostname":                types.BoolType,
	"use_ddns_ttl":                              types.BoolType,
	"use_ddns_update_fixed_addresses":           types.BoolType,
	"use_ddns_use_option81":                     types.BoolType,
	"use_deny_bootp":                            types.BoolType,
	"use_dns_update_style":                      types.BoolType,
	"use_email_list":                            types.BoolType,
	"use_enable_ddns":                           types.BoolType,
	"use_enable_dhcp_thresholds":                types.BoolType,
	"use_enable_fingerprint":                    types.BoolType,
	"use_enable_gss_tsig":                       types.BoolType,
	"use_enable_hostname_rewrite":               types.BoolType,
	"use_enable_leasequery":                     types.BoolType,
	"use_enable_one_lease_per_client":           types.BoolType,
	"use_gss_tsig_keys":                         types.BoolType,
	"use_ignore_dhcp_option_list_request":       types.BoolType,
	"use_ignore_id":                             types.BoolType,
	"use_immediate_fa_configuration":            types.BoolType,
	"use_ipv6_ddns_domainname":                  types.BoolType,
	"use_ipv6_ddns_enable_option_fqdn":          types.BoolType,
	"use_ipv6_ddns_hostname":                    types.BoolType,
	"use_ipv6_ddns_ttl":                         types.BoolType,
	"use_ipv6_dns_update_style":                 types.BoolType,
	"use_ipv6_domain_name":                      types.BoolType,
	"use_ipv6_domain_name_servers":              types.BoolType,
	"use_ipv6_enable_ddns":                      types.BoolType,
	"use_ipv6_enable_gss_tsig":                  types.BoolType,
	"use_ipv6_enable_retry_updates":             types.BoolType,
	"use_ipv6_generate_hostname":                types.BoolType,
	"use_ipv6_gss_tsig_keys":                    types.BoolType,
	"use_ipv6_lease_scavenging":                 types.BoolType,
	"use_ipv6_microsoft_code_page":              types.BoolType,
	"use_ipv6_options":                          types.BoolType,
	"use_ipv6_recycle_leases":                   types.BoolType,
	"use_ipv6_update_dns_on_lease_renewal":      types.BoolType,
	"use_lease_per_client_settings":             types.BoolType,
	"use_lease_scavenge_time":                   types.BoolType,
	"use_log_lease_events":                      types.BoolType,
	"use_logic_filter_rules":                    types.BoolType,
	"use_microsoft_code_page":                   types.BoolType,
	"use_nextserver":                            types.BoolType,
	"use_options":                               types.BoolType,
	"use_ping_count":                            types.BoolType,
	"use_ping_timeout":                          types.BoolType,
	"use_preferred_lifetime":                    types.BoolType,
	"use_prefix_length_mode":                    types.BoolType,
	"use_pxe_lease_time":                        types.BoolType,
	"use_recycle_leases":                        types.BoolType,
	"use_retry_ddns_updates":                    types.BoolType,
	"use_syslog_facility":                       types.BoolType,
	"use_update_dns_on_lease_renewal":           types.BoolType,
	"use_valid_lifetime":                        types.BoolType,
	"valid_lifetime":                            types.Int64Type,
}

var MemberDhcppropertiesResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"auth_server_group": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The Authentication Server Group object associated with this member.",
	},
	"authn_captive_portal": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The captive portal responsible for authenticating this DHCP member.",
	},
	"authn_captive_portal_authenticated_filter": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The MAC filter representing the authenticated range.",
	},
	"authn_captive_portal_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "The flag that controls if this DHCP member is enabled for captive portal authentication.",
	},
	"authn_captive_portal_guest_filter": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The MAC filter representing the guest range.",
	},
	"authn_server_group_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "The flag that controls if this DHCP member can send authentication requests to an authentication server group.",
	},
	"authority": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "The authority flag of a Grid member. This flag specifies if a DHCP server is authoritative for a domain.",
	},
	"bootfile": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of a file that DHCP clients need to boot. This setting overrides the Grid level setting.",
	},
	"bootserver": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the server on which a boot file is stored. This setting overrides the Grid level setting.",
	},
	"ddns_domainname": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The member DDNS domain name value.",
	},
	"ddns_generate_hostname": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines the ability of a member DHCP server to generate a host name and update DNS with this host name when it receives a DHCP REQUEST message that does not include a host name.",
	},
	"ddns_retry_interval": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Determines the retry interval when the member DHCP server makes repeated attempts to send DDNS updates to a DNS server.",
	},
	"ddns_server_always_updates": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines that only the DHCP server is allowed to update DNS, regardless of the requests from the DHCP clients. This setting overrides the Grid level setting.",
	},
	"ddns_ttl": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The DDNS TTL (Dynamic DNS Time To Live) value specifies the number of seconds an IP address for the name is cached.",
	},
	"ddns_update_fixed_addresses": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the member DHCP server's ability to update the A and PTR records with a fixed address is enabled or not.",
	},
	"ddns_use_option81": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if support for option 81 is enabled or not.",
	},
	"ddns_zone_primaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberDhcppropertiesDdnsZonePrimariesResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "An ordered list of zone primaries that will receive DDNS updates.",
	},
	"deny_bootp": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if a BOOTP server denies BOOTP request or not. This setting overrides the Grid level setting.",
	},
	"dhcp_utilization": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The percentage of the total DHCP utilization of DHCP objects belonging to the Grid Member multiplied by 1000. This is the percentage of the total number of available IP addresses from all the DHCP objects belonging to the Grid Member versus the total number of all IP addresses in all of the DHCP objects on the Grid Member.",
	},
	"dhcp_utilization_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "A string describing the utilization level of DHCP objects that belong to the Grid Member.",
	},
	"dns_update_style": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The update style for dynamic DNS updates.",
	},
	"dynamic_hosts": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The total number of DHCP leases issued for the DHCP objects on the Grid Member.",
	},
	"email_list": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The email_list value of a member DHCP server.",
	},
	"enable_ddns": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the member DHCP server's ability to send DDNS updates is enabled or not.",
	},
	"enable_dhcp": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the DHCP service of a member is enabled or not.",
	},
	"enable_dhcp_on_ipv6_lan2": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the DHCP service on the IPv6 LAN2 interface is enabled or not.",
	},
	"enable_dhcp_on_lan2": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the DHCP service on the LAN2 interface is enabled or not.",
	},
	"enable_dhcp_thresholds": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Represents the watermarks above or below which address usage in a network is unexpected and might warrant your attention. This setting overrides the Grid level setting.",
	},
	"enable_dhcpv6_service": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if DHCPv6 service for the member is enabled or not.",
	},
	"enable_email_warnings": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if e-mail warnings are enabled or disabled. When DHCP threshold is enabled and DHCP address usage crosses a watermark threshold, the appliance sends an e-mail notification to an administrator.",
	},
	"enable_fingerprint": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if fingerprint feature is enabled on this member. If you enable this feature, the server will match a fingerprint for incoming lease requests.",
	},
	"enable_gss_tsig": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the appliance is enabled to receive GSS-TSIG authenticated updates from DHCP clients.",
	},
	"enable_hostname_rewrite": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the Grid member's host name rewrite feature is enabled or not.",
	},
	"enable_leasequery": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if lease query is allowed or not. This setting overrides the Grid-level setting.",
	},
	"enable_snmp_warnings": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if SNMP warnings are enabled or disabled on this DHCP member. When DHCP threshold is enabled and DHCP address usage crosses a watermark threshold, the appliance sends an SNMP trap to the trap receiver that was defined for the Grid member level.",
	},
	"extattrs": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.",
	},
	"extattrs_all": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "All extensible attributes associated with the object.",
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
	},
	"gss_tsig_keys": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of GSS-TSIG keys for a member DHCP object.",
	},
	"high_water_mark": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Determines the high watermark value of a member DHCP server. If the percentage of allocated addresses exceeds this watermark, the appliance makes a syslog entry and sends an e-mail notification (if enabled). Specifies the percentage of allocated addresses. The range is from 1 to 100.",
	},
	"high_water_mark_reset": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Determines the high watermark reset value of a member DHCP server. If the percentage of allocated addresses drops below this value, a corresponding SNMP trap is reset. Specifies the percentage of allocated addresses. The range is from 1 to 100. The high watermark reset value must be lower than the high watermark value.",
	},
	"host_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Host name of the Grid member.",
	},
	"hostname_rewrite_policy": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The hostname rewrite policy that is in the protocol hostname rewrite policies array of the Grid DHCP object. This attribute is mandatory if enable_hostname_rewrite is \"true\".",
	},
	"ignore_dhcp_option_list_request": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the ignore DHCP option list request flag of a Grid member DHCP is enabled or not. If this flag is set to true all available DHCP options will be returned to the client.",
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
		MarkdownDescription: "Determines if the Immediate Fixed address configuration apply feature for the DHCP member is enabled or not.",
	},
	"ipv4addr": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The IPv4 Address of the Grid member.",
	},
	"ipv6_ddns_domainname": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The member DDNS IPv6 domain name value.",
	},
	"ipv6_ddns_enable_option_fqdn": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Controls whether the FQDN option sent by the DHCPv6 client is to be used, or if the server can automatically generate the FQDN.",
	},
	"ipv6_ddns_hostname": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The member IPv6 DDNS hostname value.",
	},
	"ipv6_ddns_server_always_updates": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the server always updates DNS or updates only if requested by the client.",
	},
	"ipv6_ddns_ttl": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The member IPv6 DDNS TTL value.",
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
		MarkdownDescription: "Determines if sending DDNS updates by the member DHCPv6 server is enabled or not.",
	},
	"ipv6_enable_gss_tsig": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the appliance is enabled to receive GSS-TSIG authenticated updates from DHCPv6 clients.",
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
		MarkdownDescription: "The list of GSS-TSIG keys for a member DHCPv6 object.",
	},
	"ipv6_kdc_server": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Determines the IPv6 address or FQDN of the Kerberos server for DHCPv6 GSS-TSIG authentication. This setting overrides the Grid level setting.",
	},
	"ipv6_lease_scavenging_time": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The member-level grace period (in seconds) to keep an expired lease before it is deleted by the scavenging process.",
	},
	"ipv6_microsoft_code_page": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The Microsoft client DHCP IPv6 code page value of a Grid member. This value is the hostname translation code page for Microsoft DHCP IPv6 clients and overrides the Grid level Microsoft DHCP IPv6 client code page.",
	},
	"ipv6_options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberDhcppropertiesIpv6OptionsResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "An array of DHCP option dhcpoption structs that lists the DHCPv6 options associated with the object.",
	},
	"ipv6_recycle_leases": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the IPv6 recycle leases feature is enabled or not. If the feature is enabled, leases are kept in the Recycle Bin until one week after lease expiration. When the feature is disabled, the leases are irrecoverably deleted.",
	},
	"ipv6_remember_expired_client_association": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enable binding for expired DHCPv6 leases.",
	},
	"ipv6_retry_updates_interval": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Determines the retry interval when the member DHCPv6 server makes repeated attempts to send DDNS updates to a DNS server.",
	},
	"ipv6_server_duid": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The server DHCPv6 unique identifier (DUID) for the Grid member.",
	},
	"ipv6_update_dns_on_lease_renewal": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Controls whether the DHCPv6 server updates DNS when an IPv6 DHCP lease is renewed.",
	},
	"ipv6addr": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The IPv6 Address of the Grid member.",
	},
	"kdc_server": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The IPv4 address or FQDN of the Kerberos server for DHCPv4 GSS-TSIG authentication. This setting overrides the Grid level setting.",
	},
	"lease_per_client_settings": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Defines how the appliance releases DHCP leases. Valid values are \"RELEASE_MACHING_ID\", \"NEVER_RELEASE\", or \"ONE_LEASE_PER_CLIENT\". The default is \"RELEASE_MATCHING_ID\".",
	},
	"lease_scavenge_time": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Determines the lease scavenging time value. When this field is set, the appliance permanently deletes the free and backup leases that remain in the database beyond a specified period of time. To disable lease scavenging, set the parameter to -1. The minimum positive value must be greater than 86400 seconds (1 day).",
	},
	"log_lease_events": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "This value specifies whether the grid member logs lease events. This setting overrides the Grid level setting.",
	},
	"logic_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberDhcppropertiesLogicFilterRulesResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "This field contains the logic filters to be applied on the Grid member. This list corresponds to the match rules that are written to the dhcpd configuration file.",
	},
	"low_water_mark": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Determines the low watermark value. If the percent of allocated addresses drops below this watermark, the appliance makes a syslog entry and sends an e-mail notification (if enabled).",
	},
	"low_water_mark_reset": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Determines the low watermark reset value. If the percentage of allocated addresses exceeds this value, a corresponding SNMP trap is reset. A number that specifies the percentage of allocated addresses. The range is from 1 to 100. The low watermark reset value must be higher than the low watermark value.",
	},
	"microsoft_code_page": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The Microsoft client DHCP IPv4 code page value of a grid member. This value is the hostname translation code page for Microsoft DHCP IPv4 clients and overrides the Grid level Microsoft DHCP IPv4 client code page.",
	},
	"nextserver": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The next server value of a member DHCP server. This value is the IP address or name of the boot file server on which the boot file is stored.",
	},
	"option60_match_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberDhcppropertiesOption60MatchRulesResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of option 60 match rules.",
	},
	"options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberDhcppropertiesOptionsResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object.",
	},
	"ping_count": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Specifies the number of pings that the Infoblox appliance sends to an IP address to verify that it is not in use. Values are from 0 to 10, where 0 disables pings.",
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
	"pxe_lease_time": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Specifies the duration of time it takes a host to connect to a boot server, such as a TFTP server, and download the file it needs to boot. A 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached.",
	},
	"recycle_leases": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the recycle leases feature is enabled or not. If you enabled this feature and then delete a DHCP range, the appliance stores active leases from this range up to one week after the leases expires.",
	},
	"retry_ddns_updates": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Indicates whether the DHCP server makes repeated attempts to send DDNS updates to a DNS server.",
	},
	"static_hosts": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of static DHCP addresses configured in DHCP objects that belong to the Grid Member.",
	},
	"syslog_facility": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The syslog facility is the location on the syslog server to which you want to sort the syslog messages.",
	},
	"total_hosts": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The total number of DHCP addresses configured in DHCP objects that belong to the Grid Member.",
	},
	"update_dns_on_lease_renewal": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Controls whether the DHCP server updates DNS when a DHCP lease is renewed.",
	},
	"use_authority": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: authority",
	},
	"use_bootfile": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: bootfile",
	},
	"use_bootserver": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: bootserver",
	},
	"use_ddns_domainname": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ddns_domainname",
	},
	"use_ddns_generate_hostname": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ddns_generate_hostname",
	},
	"use_ddns_ttl": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ddns_ttl",
	},
	"use_ddns_update_fixed_addresses": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ddns_update_fixed_addresses",
	},
	"use_ddns_use_option81": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ddns_use_option81",
	},
	"use_deny_bootp": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: deny_bootp",
	},
	"use_dns_update_style": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: dns_update_style",
	},
	"use_email_list": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: email_list",
	},
	"use_enable_ddns": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: enable_ddns",
	},
	"use_enable_dhcp_thresholds": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: enable_dhcp_thresholds , high_water_mark, high_water_mark_reset, low_water_mark, low_water_mark_reset",
	},
	"use_enable_fingerprint": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: enable_fingerprint",
	},
	"use_enable_gss_tsig": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: kdc_server , enable_gss_tsig",
	},
	"use_enable_hostname_rewrite": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: enable_hostname_rewrite , hostname_rewrite_policy",
	},
	"use_enable_leasequery": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: enable_leasequery",
	},
	"use_enable_one_lease_per_client": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: enable_one_lease_per_client",
	},
	"use_gss_tsig_keys": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: gss_tsig_keys",
	},
	"use_ignore_dhcp_option_list_request": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ignore_dhcp_option_list_request",
	},
	"use_ignore_id": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ignore_id",
	},
	"use_immediate_fa_configuration": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: immediate_fa_configuration",
	},
	"use_ipv6_ddns_domainname": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ipv6_ddns_domainname",
	},
	"use_ipv6_ddns_enable_option_fqdn": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ipv6_ddns_enable_option_fqdn",
	},
	"use_ipv6_ddns_hostname": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ipv6_ddns_hostname",
	},
	"use_ipv6_ddns_ttl": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ipv6_ddns_ttl",
	},
	"use_ipv6_dns_update_style": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ipv6_dns_update_style",
	},
	"use_ipv6_domain_name": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ipv6_domain_name",
	},
	"use_ipv6_domain_name_servers": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ipv6_domain_name_servers",
	},
	"use_ipv6_enable_ddns": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ipv6_enable_ddns",
	},
	"use_ipv6_enable_gss_tsig": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ipv6_kdc_server , ipv6_enable_gss_tsig",
	},
	"use_ipv6_enable_retry_updates": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ipv6_enable_retry_updates , ipv6_retry_updates_interval",
	},
	"use_ipv6_generate_hostname": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ipv6_generate_hostname",
	},
	"use_ipv6_gss_tsig_keys": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ipv6_gss_tsig_keys",
	},
	"use_ipv6_lease_scavenging": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ipv6_enable_lease_scavenging , ipv6_lease_scavenging_time, ipv6_remember_expired_client_association",
	},
	"use_ipv6_microsoft_code_page": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ipv6_microsoft_code_page",
	},
	"use_ipv6_options": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ipv6_options",
	},
	"use_ipv6_recycle_leases": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ipv6_recycle_leases",
	},
	"use_ipv6_update_dns_on_lease_renewal": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ipv6_update_dns_on_lease_renewal",
	},
	"use_lease_per_client_settings": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: lease_per_client_settings",
	},
	"use_lease_scavenge_time": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: lease_scavenge_time",
	},
	"use_log_lease_events": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: log_lease_events",
	},
	"use_logic_filter_rules": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: logic_filter_rules",
	},
	"use_microsoft_code_page": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: microsoft_code_page",
	},
	"use_nextserver": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: nextserver",
	},
	"use_options": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: options",
	},
	"use_ping_count": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ping_count",
	},
	"use_ping_timeout": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ping_timeout",
	},
	"use_preferred_lifetime": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: preferred_lifetime",
	},
	"use_prefix_length_mode": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: prefix_length_mode",
	},
	"use_pxe_lease_time": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: pxe_lease_time",
	},
	"use_recycle_leases": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: recycle_leases",
	},
	"use_retry_ddns_updates": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ddns_retry_interval , retry_ddns_updates",
	},
	"use_syslog_facility": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: syslog_facility",
	},
	"use_update_dns_on_lease_renewal": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: update_dns_on_lease_renewal",
	},
	"use_valid_lifetime": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: valid_lifetime",
	},
	"valid_lifetime": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The valid lifetime for Grid Member DHCP. Specifies the length of time addresses that are assigned to DHCPv6 clients remain in the valid state.",
	},
}

func ExpandMemberDhcpproperties(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberDhcpproperties {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberDhcppropertiesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberDhcppropertiesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberDhcpproperties {
	if m == nil {
		return nil
	}
	to := &grid.MemberDhcpproperties{
		Ref:                                   flex.ExpandStringPointer(m.Ref),
		AuthServerGroup:                       flex.ExpandStringPointer(m.AuthServerGroup),
		AuthnCaptivePortal:                    flex.ExpandStringPointer(m.AuthnCaptivePortal),
		AuthnCaptivePortalAuthenticatedFilter: flex.ExpandStringPointer(m.AuthnCaptivePortalAuthenticatedFilter),
		AuthnCaptivePortalEnabled:             flex.ExpandBoolPointer(m.AuthnCaptivePortalEnabled),
		AuthnCaptivePortalGuestFilter:         flex.ExpandStringPointer(m.AuthnCaptivePortalGuestFilter),
		AuthnServerGroupEnabled:               flex.ExpandBoolPointer(m.AuthnServerGroupEnabled),
		Authority:                             flex.ExpandBoolPointer(m.Authority),
		Bootfile:                              flex.ExpandStringPointer(m.Bootfile),
		Bootserver:                            flex.ExpandStringPointer(m.Bootserver),
		DdnsDomainname:                        flex.ExpandStringPointer(m.DdnsDomainname),
		DdnsGenerateHostname:                  flex.ExpandBoolPointer(m.DdnsGenerateHostname),
		DdnsRetryInterval:                     flex.ExpandInt64Pointer(m.DdnsRetryInterval),
		DdnsServerAlwaysUpdates:               flex.ExpandBoolPointer(m.DdnsServerAlwaysUpdates),
		DdnsTtl:                               flex.ExpandInt64Pointer(m.DdnsTtl),
		DdnsUpdateFixedAddresses:              flex.ExpandBoolPointer(m.DdnsUpdateFixedAddresses),
		DdnsUseOption81:                       flex.ExpandBoolPointer(m.DdnsUseOption81),
		DdnsZonePrimaries:                     flex.ExpandFrameworkListNestedBlock(ctx, m.DdnsZonePrimaries, diags, ExpandMemberDhcppropertiesDdnsZonePrimaries),
		DenyBootp:                             flex.ExpandBoolPointer(m.DenyBootp),
		DnsUpdateStyle:                        flex.ExpandStringPointer(m.DnsUpdateStyle),
		EmailList:                             flex.ExpandFrameworkListString(ctx, m.EmailList, diags),
		EnableDdns:                            flex.ExpandBoolPointer(m.EnableDdns),
		EnableDhcp:                            flex.ExpandBoolPointer(m.EnableDhcp),
		EnableDhcpOnIpv6Lan2:                  flex.ExpandBoolPointer(m.EnableDhcpOnIpv6Lan2),
		EnableDhcpOnLan2:                      flex.ExpandBoolPointer(m.EnableDhcpOnLan2),
		EnableDhcpThresholds:                  flex.ExpandBoolPointer(m.EnableDhcpThresholds),
		EnableDhcpv6Service:                   flex.ExpandBoolPointer(m.EnableDhcpv6Service),
		EnableEmailWarnings:                   flex.ExpandBoolPointer(m.EnableEmailWarnings),
		EnableFingerprint:                     flex.ExpandBoolPointer(m.EnableFingerprint),
		EnableGssTsig:                         flex.ExpandBoolPointer(m.EnableGssTsig),
		EnableHostnameRewrite:                 flex.ExpandBoolPointer(m.EnableHostnameRewrite),
		EnableLeasequery:                      flex.ExpandBoolPointer(m.EnableLeasequery),
		EnableSnmpWarnings:                    flex.ExpandBoolPointer(m.EnableSnmpWarnings),
		ExtAttrs:                              ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		GssTsigKeys:                           flex.ExpandFrameworkListString(ctx, m.GssTsigKeys, diags),
		HighWaterMark:                         flex.ExpandInt64Pointer(m.HighWaterMark),
		HighWaterMarkReset:                    flex.ExpandInt64Pointer(m.HighWaterMarkReset),
		HostnameRewritePolicy:                 flex.ExpandStringPointer(m.HostnameRewritePolicy),
		IgnoreDhcpOptionListRequest:           flex.ExpandBoolPointer(m.IgnoreDhcpOptionListRequest),
		IgnoreId:                              flex.ExpandStringPointer(m.IgnoreId),
		IgnoreMacAddresses:                    flex.ExpandFrameworkListString(ctx, m.IgnoreMacAddresses, diags),
		ImmediateFaConfiguration:              flex.ExpandBoolPointer(m.ImmediateFaConfiguration),
		Ipv6DdnsDomainname:                    flex.ExpandStringPointer(m.Ipv6DdnsDomainname),
		Ipv6DdnsEnableOptionFqdn:              flex.ExpandBoolPointer(m.Ipv6DdnsEnableOptionFqdn),
		Ipv6DdnsHostname:                      flex.ExpandStringPointer(m.Ipv6DdnsHostname),
		Ipv6DdnsServerAlwaysUpdates:           flex.ExpandBoolPointer(m.Ipv6DdnsServerAlwaysUpdates),
		Ipv6DdnsTtl:                           flex.ExpandInt64Pointer(m.Ipv6DdnsTtl),
		Ipv6DnsUpdateStyle:                    flex.ExpandStringPointer(m.Ipv6DnsUpdateStyle),
		Ipv6DomainName:                        flex.ExpandStringPointer(m.Ipv6DomainName),
		Ipv6DomainNameServers:                 flex.ExpandFrameworkListString(ctx, m.Ipv6DomainNameServers, diags),
		Ipv6EnableDdns:                        flex.ExpandBoolPointer(m.Ipv6EnableDdns),
		Ipv6EnableGssTsig:                     flex.ExpandBoolPointer(m.Ipv6EnableGssTsig),
		Ipv6EnableLeaseScavenging:             flex.ExpandBoolPointer(m.Ipv6EnableLeaseScavenging),
		Ipv6EnableRetryUpdates:                flex.ExpandBoolPointer(m.Ipv6EnableRetryUpdates),
		Ipv6GenerateHostname:                  flex.ExpandBoolPointer(m.Ipv6GenerateHostname),
		Ipv6GssTsigKeys:                       flex.ExpandFrameworkListString(ctx, m.Ipv6GssTsigKeys, diags),
		Ipv6KdcServer:                         flex.ExpandStringPointer(m.Ipv6KdcServer),
		Ipv6LeaseScavengingTime:               flex.ExpandInt64Pointer(m.Ipv6LeaseScavengingTime),
		Ipv6MicrosoftCodePage:                 flex.ExpandStringPointer(m.Ipv6MicrosoftCodePage),
		Ipv6Options:                           flex.ExpandFrameworkListNestedBlock(ctx, m.Ipv6Options, diags, ExpandMemberDhcppropertiesIpv6Options),
		Ipv6RecycleLeases:                     flex.ExpandBoolPointer(m.Ipv6RecycleLeases),
		Ipv6RememberExpiredClientAssociation:  flex.ExpandBoolPointer(m.Ipv6RememberExpiredClientAssociation),
		Ipv6RetryUpdatesInterval:              flex.ExpandInt64Pointer(m.Ipv6RetryUpdatesInterval),
		Ipv6ServerDuid:                        flex.ExpandStringPointer(m.Ipv6ServerDuid),
		Ipv6UpdateDnsOnLeaseRenewal:           flex.ExpandBoolPointer(m.Ipv6UpdateDnsOnLeaseRenewal),
		KdcServer:                             flex.ExpandStringPointer(m.KdcServer),
		LeasePerClientSettings:                flex.ExpandStringPointer(m.LeasePerClientSettings),
		LeaseScavengeTime:                     flex.ExpandInt64Pointer(m.LeaseScavengeTime),
		LogLeaseEvents:                        flex.ExpandBoolPointer(m.LogLeaseEvents),
		LogicFilterRules:                      flex.ExpandFrameworkListNestedBlock(ctx, m.LogicFilterRules, diags, ExpandMemberDhcppropertiesLogicFilterRules),
		LowWaterMark:                          flex.ExpandInt64Pointer(m.LowWaterMark),
		LowWaterMarkReset:                     flex.ExpandInt64Pointer(m.LowWaterMarkReset),
		MicrosoftCodePage:                     flex.ExpandStringPointer(m.MicrosoftCodePage),
		Nextserver:                            flex.ExpandStringPointer(m.Nextserver),
		Option60MatchRules:                    flex.ExpandFrameworkListNestedBlock(ctx, m.Option60MatchRules, diags, ExpandMemberDhcppropertiesOption60MatchRules),
		Options:                               flex.ExpandFrameworkListNestedBlock(ctx, m.Options, diags, ExpandMemberDhcppropertiesOptions),
		PingCount:                             flex.ExpandInt64Pointer(m.PingCount),
		PingTimeout:                           flex.ExpandInt64Pointer(m.PingTimeout),
		PreferredLifetime:                     flex.ExpandInt64Pointer(m.PreferredLifetime),
		PrefixLengthMode:                      flex.ExpandStringPointer(m.PrefixLengthMode),
		PxeLeaseTime:                          flex.ExpandInt64Pointer(m.PxeLeaseTime),
		RecycleLeases:                         flex.ExpandBoolPointer(m.RecycleLeases),
		RetryDdnsUpdates:                      flex.ExpandBoolPointer(m.RetryDdnsUpdates),
		SyslogFacility:                        flex.ExpandStringPointer(m.SyslogFacility),
		UpdateDnsOnLeaseRenewal:               flex.ExpandBoolPointer(m.UpdateDnsOnLeaseRenewal),
		UseAuthority:                          flex.ExpandBoolPointer(m.UseAuthority),
		UseBootfile:                           flex.ExpandBoolPointer(m.UseBootfile),
		UseBootserver:                         flex.ExpandBoolPointer(m.UseBootserver),
		UseDdnsDomainname:                     flex.ExpandBoolPointer(m.UseDdnsDomainname),
		UseDdnsGenerateHostname:               flex.ExpandBoolPointer(m.UseDdnsGenerateHostname),
		UseDdnsTtl:                            flex.ExpandBoolPointer(m.UseDdnsTtl),
		UseDdnsUpdateFixedAddresses:           flex.ExpandBoolPointer(m.UseDdnsUpdateFixedAddresses),
		UseDdnsUseOption81:                    flex.ExpandBoolPointer(m.UseDdnsUseOption81),
		UseDenyBootp:                          flex.ExpandBoolPointer(m.UseDenyBootp),
		UseDnsUpdateStyle:                     flex.ExpandBoolPointer(m.UseDnsUpdateStyle),
		UseEmailList:                          flex.ExpandBoolPointer(m.UseEmailList),
		UseEnableDdns:                         flex.ExpandBoolPointer(m.UseEnableDdns),
		UseEnableDhcpThresholds:               flex.ExpandBoolPointer(m.UseEnableDhcpThresholds),
		UseEnableFingerprint:                  flex.ExpandBoolPointer(m.UseEnableFingerprint),
		UseEnableGssTsig:                      flex.ExpandBoolPointer(m.UseEnableGssTsig),
		UseEnableHostnameRewrite:              flex.ExpandBoolPointer(m.UseEnableHostnameRewrite),
		UseEnableLeasequery:                   flex.ExpandBoolPointer(m.UseEnableLeasequery),
		UseEnableOneLeasePerClient:            flex.ExpandBoolPointer(m.UseEnableOneLeasePerClient),
		UseGssTsigKeys:                        flex.ExpandBoolPointer(m.UseGssTsigKeys),
		UseIgnoreDhcpOptionListRequest:        flex.ExpandBoolPointer(m.UseIgnoreDhcpOptionListRequest),
		UseIgnoreId:                           flex.ExpandBoolPointer(m.UseIgnoreId),
		UseImmediateFaConfiguration:           flex.ExpandBoolPointer(m.UseImmediateFaConfiguration),
		UseIpv6DdnsDomainname:                 flex.ExpandBoolPointer(m.UseIpv6DdnsDomainname),
		UseIpv6DdnsEnableOptionFqdn:           flex.ExpandBoolPointer(m.UseIpv6DdnsEnableOptionFqdn),
		UseIpv6DdnsHostname:                   flex.ExpandBoolPointer(m.UseIpv6DdnsHostname),
		UseIpv6DdnsTtl:                        flex.ExpandBoolPointer(m.UseIpv6DdnsTtl),
		UseIpv6DnsUpdateStyle:                 flex.ExpandBoolPointer(m.UseIpv6DnsUpdateStyle),
		UseIpv6DomainName:                     flex.ExpandBoolPointer(m.UseIpv6DomainName),
		UseIpv6DomainNameServers:              flex.ExpandBoolPointer(m.UseIpv6DomainNameServers),
		UseIpv6EnableDdns:                     flex.ExpandBoolPointer(m.UseIpv6EnableDdns),
		UseIpv6EnableGssTsig:                  flex.ExpandBoolPointer(m.UseIpv6EnableGssTsig),
		UseIpv6EnableRetryUpdates:             flex.ExpandBoolPointer(m.UseIpv6EnableRetryUpdates),
		UseIpv6GenerateHostname:               flex.ExpandBoolPointer(m.UseIpv6GenerateHostname),
		UseIpv6GssTsigKeys:                    flex.ExpandBoolPointer(m.UseIpv6GssTsigKeys),
		UseIpv6LeaseScavenging:                flex.ExpandBoolPointer(m.UseIpv6LeaseScavenging),
		UseIpv6MicrosoftCodePage:              flex.ExpandBoolPointer(m.UseIpv6MicrosoftCodePage),
		UseIpv6Options:                        flex.ExpandBoolPointer(m.UseIpv6Options),
		UseIpv6RecycleLeases:                  flex.ExpandBoolPointer(m.UseIpv6RecycleLeases),
		UseIpv6UpdateDnsOnLeaseRenewal:        flex.ExpandBoolPointer(m.UseIpv6UpdateDnsOnLeaseRenewal),
		UseLeasePerClientSettings:             flex.ExpandBoolPointer(m.UseLeasePerClientSettings),
		UseLeaseScavengeTime:                  flex.ExpandBoolPointer(m.UseLeaseScavengeTime),
		UseLogLeaseEvents:                     flex.ExpandBoolPointer(m.UseLogLeaseEvents),
		UseLogicFilterRules:                   flex.ExpandBoolPointer(m.UseLogicFilterRules),
		UseMicrosoftCodePage:                  flex.ExpandBoolPointer(m.UseMicrosoftCodePage),
		UseNextserver:                         flex.ExpandBoolPointer(m.UseNextserver),
		UseOptions:                            flex.ExpandBoolPointer(m.UseOptions),
		UsePingCount:                          flex.ExpandBoolPointer(m.UsePingCount),
		UsePingTimeout:                        flex.ExpandBoolPointer(m.UsePingTimeout),
		UsePreferredLifetime:                  flex.ExpandBoolPointer(m.UsePreferredLifetime),
		UsePrefixLengthMode:                   flex.ExpandBoolPointer(m.UsePrefixLengthMode),
		UsePxeLeaseTime:                       flex.ExpandBoolPointer(m.UsePxeLeaseTime),
		UseRecycleLeases:                      flex.ExpandBoolPointer(m.UseRecycleLeases),
		UseRetryDdnsUpdates:                   flex.ExpandBoolPointer(m.UseRetryDdnsUpdates),
		UseSyslogFacility:                     flex.ExpandBoolPointer(m.UseSyslogFacility),
		UseUpdateDnsOnLeaseRenewal:            flex.ExpandBoolPointer(m.UseUpdateDnsOnLeaseRenewal),
		UseValidLifetime:                      flex.ExpandBoolPointer(m.UseValidLifetime),
		ValidLifetime:                         flex.ExpandInt64Pointer(m.ValidLifetime),
	}
	return to
}

func FlattenMemberDhcpproperties(ctx context.Context, from *grid.MemberDhcpproperties, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberDhcppropertiesAttrTypes)
	}
	m := MemberDhcppropertiesModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, MemberDhcppropertiesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberDhcppropertiesModel) Flatten(ctx context.Context, from *grid.MemberDhcpproperties, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberDhcppropertiesModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.AuthServerGroup = flex.FlattenStringPointer(from.AuthServerGroup)
	m.AuthnCaptivePortal = flex.FlattenStringPointer(from.AuthnCaptivePortal)
	m.AuthnCaptivePortalAuthenticatedFilter = flex.FlattenStringPointer(from.AuthnCaptivePortalAuthenticatedFilter)
	m.AuthnCaptivePortalEnabled = types.BoolPointerValue(from.AuthnCaptivePortalEnabled)
	m.AuthnCaptivePortalGuestFilter = flex.FlattenStringPointer(from.AuthnCaptivePortalGuestFilter)
	m.AuthnServerGroupEnabled = types.BoolPointerValue(from.AuthnServerGroupEnabled)
	m.Authority = types.BoolPointerValue(from.Authority)
	m.Bootfile = flex.FlattenStringPointer(from.Bootfile)
	m.Bootserver = flex.FlattenStringPointer(from.Bootserver)
	m.DdnsDomainname = flex.FlattenStringPointer(from.DdnsDomainname)
	m.DdnsGenerateHostname = types.BoolPointerValue(from.DdnsGenerateHostname)
	m.DdnsRetryInterval = flex.FlattenInt64Pointer(from.DdnsRetryInterval)
	m.DdnsServerAlwaysUpdates = types.BoolPointerValue(from.DdnsServerAlwaysUpdates)
	m.DdnsTtl = flex.FlattenInt64Pointer(from.DdnsTtl)
	m.DdnsUpdateFixedAddresses = types.BoolPointerValue(from.DdnsUpdateFixedAddresses)
	m.DdnsUseOption81 = types.BoolPointerValue(from.DdnsUseOption81)
	m.DdnsZonePrimaries = flex.FlattenFrameworkListNestedBlock(ctx, from.DdnsZonePrimaries, MemberDhcppropertiesDdnsZonePrimariesAttrTypes, diags, FlattenMemberDhcppropertiesDdnsZonePrimaries)
	m.DenyBootp = types.BoolPointerValue(from.DenyBootp)
	m.DhcpUtilization = flex.FlattenInt64Pointer(from.DhcpUtilization)
	m.DhcpUtilizationStatus = flex.FlattenStringPointer(from.DhcpUtilizationStatus)
	m.DnsUpdateStyle = flex.FlattenStringPointer(from.DnsUpdateStyle)
	m.DynamicHosts = flex.FlattenInt64Pointer(from.DynamicHosts)
	m.EmailList = flex.FlattenFrameworkListString(ctx, from.EmailList, diags)
	m.EnableDdns = types.BoolPointerValue(from.EnableDdns)
	m.EnableDhcp = types.BoolPointerValue(from.EnableDhcp)
	m.EnableDhcpOnIpv6Lan2 = types.BoolPointerValue(from.EnableDhcpOnIpv6Lan2)
	m.EnableDhcpOnLan2 = types.BoolPointerValue(from.EnableDhcpOnLan2)
	m.EnableDhcpThresholds = types.BoolPointerValue(from.EnableDhcpThresholds)
	m.EnableDhcpv6Service = types.BoolPointerValue(from.EnableDhcpv6Service)
	m.EnableEmailWarnings = types.BoolPointerValue(from.EnableEmailWarnings)
	m.EnableFingerprint = types.BoolPointerValue(from.EnableFingerprint)
	m.EnableGssTsig = types.BoolPointerValue(from.EnableGssTsig)
	m.EnableHostnameRewrite = types.BoolPointerValue(from.EnableHostnameRewrite)
	m.EnableLeasequery = types.BoolPointerValue(from.EnableLeasequery)
	m.EnableSnmpWarnings = types.BoolPointerValue(from.EnableSnmpWarnings)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.GssTsigKeys = flex.FlattenFrameworkListString(ctx, from.GssTsigKeys, diags)
	m.HighWaterMark = flex.FlattenInt64Pointer(from.HighWaterMark)
	m.HighWaterMarkReset = flex.FlattenInt64Pointer(from.HighWaterMarkReset)
	m.HostName = flex.FlattenStringPointer(from.HostName)
	m.HostnameRewritePolicy = flex.FlattenStringPointer(from.HostnameRewritePolicy)
	m.IgnoreDhcpOptionListRequest = types.BoolPointerValue(from.IgnoreDhcpOptionListRequest)
	m.IgnoreId = flex.FlattenStringPointer(from.IgnoreId)
	m.IgnoreMacAddresses = flex.FlattenFrameworkListString(ctx, from.IgnoreMacAddresses, diags)
	m.ImmediateFaConfiguration = types.BoolPointerValue(from.ImmediateFaConfiguration)
	m.Ipv4addr = flex.FlattenStringPointer(from.Ipv4addr)
	m.Ipv6DdnsDomainname = flex.FlattenStringPointer(from.Ipv6DdnsDomainname)
	m.Ipv6DdnsEnableOptionFqdn = types.BoolPointerValue(from.Ipv6DdnsEnableOptionFqdn)
	m.Ipv6DdnsHostname = flex.FlattenStringPointer(from.Ipv6DdnsHostname)
	m.Ipv6DdnsServerAlwaysUpdates = types.BoolPointerValue(from.Ipv6DdnsServerAlwaysUpdates)
	m.Ipv6DdnsTtl = flex.FlattenInt64Pointer(from.Ipv6DdnsTtl)
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
	m.Ipv6Options = flex.FlattenFrameworkListNestedBlock(ctx, from.Ipv6Options, MemberDhcppropertiesIpv6OptionsAttrTypes, diags, FlattenMemberDhcppropertiesIpv6Options)
	m.Ipv6RecycleLeases = types.BoolPointerValue(from.Ipv6RecycleLeases)
	m.Ipv6RememberExpiredClientAssociation = types.BoolPointerValue(from.Ipv6RememberExpiredClientAssociation)
	m.Ipv6RetryUpdatesInterval = flex.FlattenInt64Pointer(from.Ipv6RetryUpdatesInterval)
	m.Ipv6ServerDuid = flex.FlattenStringPointer(from.Ipv6ServerDuid)
	m.Ipv6UpdateDnsOnLeaseRenewal = types.BoolPointerValue(from.Ipv6UpdateDnsOnLeaseRenewal)
	m.Ipv6addr = flex.FlattenStringPointer(from.Ipv6addr)
	m.KdcServer = flex.FlattenStringPointer(from.KdcServer)
	m.LeasePerClientSettings = flex.FlattenStringPointer(from.LeasePerClientSettings)
	m.LeaseScavengeTime = flex.FlattenInt64Pointer(from.LeaseScavengeTime)
	m.LogLeaseEvents = types.BoolPointerValue(from.LogLeaseEvents)
	m.LogicFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.LogicFilterRules, MemberDhcppropertiesLogicFilterRulesAttrTypes, diags, FlattenMemberDhcppropertiesLogicFilterRules)
	m.LowWaterMark = flex.FlattenInt64Pointer(from.LowWaterMark)
	m.LowWaterMarkReset = flex.FlattenInt64Pointer(from.LowWaterMarkReset)
	m.MicrosoftCodePage = flex.FlattenStringPointer(from.MicrosoftCodePage)
	m.Nextserver = flex.FlattenStringPointer(from.Nextserver)
	m.Option60MatchRules = flex.FlattenFrameworkListNestedBlock(ctx, from.Option60MatchRules, MemberDhcppropertiesOption60MatchRulesAttrTypes, diags, FlattenMemberDhcppropertiesOption60MatchRules)
	m.Options = flex.FlattenFrameworkListNestedBlock(ctx, from.Options, MemberDhcppropertiesOptionsAttrTypes, diags, FlattenMemberDhcppropertiesOptions)
	m.PingCount = flex.FlattenInt64Pointer(from.PingCount)
	m.PingTimeout = flex.FlattenInt64Pointer(from.PingTimeout)
	m.PreferredLifetime = flex.FlattenInt64Pointer(from.PreferredLifetime)
	m.PrefixLengthMode = flex.FlattenStringPointer(from.PrefixLengthMode)
	m.PxeLeaseTime = flex.FlattenInt64Pointer(from.PxeLeaseTime)
	m.RecycleLeases = types.BoolPointerValue(from.RecycleLeases)
	m.RetryDdnsUpdates = types.BoolPointerValue(from.RetryDdnsUpdates)
	m.StaticHosts = flex.FlattenInt64Pointer(from.StaticHosts)
	m.SyslogFacility = flex.FlattenStringPointer(from.SyslogFacility)
	m.TotalHosts = flex.FlattenInt64Pointer(from.TotalHosts)
	m.UpdateDnsOnLeaseRenewal = types.BoolPointerValue(from.UpdateDnsOnLeaseRenewal)
	m.UseAuthority = types.BoolPointerValue(from.UseAuthority)
	m.UseBootfile = types.BoolPointerValue(from.UseBootfile)
	m.UseBootserver = types.BoolPointerValue(from.UseBootserver)
	m.UseDdnsDomainname = types.BoolPointerValue(from.UseDdnsDomainname)
	m.UseDdnsGenerateHostname = types.BoolPointerValue(from.UseDdnsGenerateHostname)
	m.UseDdnsTtl = types.BoolPointerValue(from.UseDdnsTtl)
	m.UseDdnsUpdateFixedAddresses = types.BoolPointerValue(from.UseDdnsUpdateFixedAddresses)
	m.UseDdnsUseOption81 = types.BoolPointerValue(from.UseDdnsUseOption81)
	m.UseDenyBootp = types.BoolPointerValue(from.UseDenyBootp)
	m.UseDnsUpdateStyle = types.BoolPointerValue(from.UseDnsUpdateStyle)
	m.UseEmailList = types.BoolPointerValue(from.UseEmailList)
	m.UseEnableDdns = types.BoolPointerValue(from.UseEnableDdns)
	m.UseEnableDhcpThresholds = types.BoolPointerValue(from.UseEnableDhcpThresholds)
	m.UseEnableFingerprint = types.BoolPointerValue(from.UseEnableFingerprint)
	m.UseEnableGssTsig = types.BoolPointerValue(from.UseEnableGssTsig)
	m.UseEnableHostnameRewrite = types.BoolPointerValue(from.UseEnableHostnameRewrite)
	m.UseEnableLeasequery = types.BoolPointerValue(from.UseEnableLeasequery)
	m.UseEnableOneLeasePerClient = types.BoolPointerValue(from.UseEnableOneLeasePerClient)
	m.UseGssTsigKeys = types.BoolPointerValue(from.UseGssTsigKeys)
	m.UseIgnoreDhcpOptionListRequest = types.BoolPointerValue(from.UseIgnoreDhcpOptionListRequest)
	m.UseIgnoreId = types.BoolPointerValue(from.UseIgnoreId)
	m.UseImmediateFaConfiguration = types.BoolPointerValue(from.UseImmediateFaConfiguration)
	m.UseIpv6DdnsDomainname = types.BoolPointerValue(from.UseIpv6DdnsDomainname)
	m.UseIpv6DdnsEnableOptionFqdn = types.BoolPointerValue(from.UseIpv6DdnsEnableOptionFqdn)
	m.UseIpv6DdnsHostname = types.BoolPointerValue(from.UseIpv6DdnsHostname)
	m.UseIpv6DdnsTtl = types.BoolPointerValue(from.UseIpv6DdnsTtl)
	m.UseIpv6DnsUpdateStyle = types.BoolPointerValue(from.UseIpv6DnsUpdateStyle)
	m.UseIpv6DomainName = types.BoolPointerValue(from.UseIpv6DomainName)
	m.UseIpv6DomainNameServers = types.BoolPointerValue(from.UseIpv6DomainNameServers)
	m.UseIpv6EnableDdns = types.BoolPointerValue(from.UseIpv6EnableDdns)
	m.UseIpv6EnableGssTsig = types.BoolPointerValue(from.UseIpv6EnableGssTsig)
	m.UseIpv6EnableRetryUpdates = types.BoolPointerValue(from.UseIpv6EnableRetryUpdates)
	m.UseIpv6GenerateHostname = types.BoolPointerValue(from.UseIpv6GenerateHostname)
	m.UseIpv6GssTsigKeys = types.BoolPointerValue(from.UseIpv6GssTsigKeys)
	m.UseIpv6LeaseScavenging = types.BoolPointerValue(from.UseIpv6LeaseScavenging)
	m.UseIpv6MicrosoftCodePage = types.BoolPointerValue(from.UseIpv6MicrosoftCodePage)
	m.UseIpv6Options = types.BoolPointerValue(from.UseIpv6Options)
	m.UseIpv6RecycleLeases = types.BoolPointerValue(from.UseIpv6RecycleLeases)
	m.UseIpv6UpdateDnsOnLeaseRenewal = types.BoolPointerValue(from.UseIpv6UpdateDnsOnLeaseRenewal)
	m.UseLeasePerClientSettings = types.BoolPointerValue(from.UseLeasePerClientSettings)
	m.UseLeaseScavengeTime = types.BoolPointerValue(from.UseLeaseScavengeTime)
	m.UseLogLeaseEvents = types.BoolPointerValue(from.UseLogLeaseEvents)
	m.UseLogicFilterRules = types.BoolPointerValue(from.UseLogicFilterRules)
	m.UseMicrosoftCodePage = types.BoolPointerValue(from.UseMicrosoftCodePage)
	m.UseNextserver = types.BoolPointerValue(from.UseNextserver)
	m.UseOptions = types.BoolPointerValue(from.UseOptions)
	m.UsePingCount = types.BoolPointerValue(from.UsePingCount)
	m.UsePingTimeout = types.BoolPointerValue(from.UsePingTimeout)
	m.UsePreferredLifetime = types.BoolPointerValue(from.UsePreferredLifetime)
	m.UsePrefixLengthMode = types.BoolPointerValue(from.UsePrefixLengthMode)
	m.UsePxeLeaseTime = types.BoolPointerValue(from.UsePxeLeaseTime)
	m.UseRecycleLeases = types.BoolPointerValue(from.UseRecycleLeases)
	m.UseRetryDdnsUpdates = types.BoolPointerValue(from.UseRetryDdnsUpdates)
	m.UseSyslogFacility = types.BoolPointerValue(from.UseSyslogFacility)
	m.UseUpdateDnsOnLeaseRenewal = types.BoolPointerValue(from.UseUpdateDnsOnLeaseRenewal)
	m.UseValidLifetime = types.BoolPointerValue(from.UseValidLifetime)
	m.ValidLifetime = flex.FlattenInt64Pointer(from.ValidLifetime)
}
