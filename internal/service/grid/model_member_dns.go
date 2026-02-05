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

type MemberDnsModel struct {
	Ref                              types.String `tfsdk:"ref"`
	AddClientIpMacOptions            types.Bool   `tfsdk:"add_client_ip_mac_options"`
	AdditionalIpList                 types.List   `tfsdk:"additional_ip_list"`
	AdditionalIpListStruct           types.List   `tfsdk:"additional_ip_list_struct"`
	AllowGssTsigZoneUpdates          types.Bool   `tfsdk:"allow_gss_tsig_zone_updates"`
	AllowQuery                       types.List   `tfsdk:"allow_query"`
	AllowRecursiveQuery              types.Bool   `tfsdk:"allow_recursive_query"`
	AllowTransfer                    types.List   `tfsdk:"allow_transfer"`
	AllowUpdate                      types.List   `tfsdk:"allow_update"`
	AnonymizeResponseLogging         types.Bool   `tfsdk:"anonymize_response_logging"`
	AtcFwdEnable                     types.Bool   `tfsdk:"atc_fwd_enable"`
	AttackMitigation                 types.Object `tfsdk:"attack_mitigation"`
	AutoBlackhole                    types.Object `tfsdk:"auto_blackhole"`
	AutoCreateAAndPtrForLan2         types.Bool   `tfsdk:"auto_create_a_and_ptr_for_lan2"`
	AutoCreateAaaaAndIpv6ptrForLan2  types.Bool   `tfsdk:"auto_create_aaaa_and_ipv6ptr_for_lan2"`
	AutoSortViews                    types.Bool   `tfsdk:"auto_sort_views"`
	BindCheckNamesPolicy             types.String `tfsdk:"bind_check_names_policy"`
	BindHostnameDirective            types.String `tfsdk:"bind_hostname_directive"`
	BindHostnameDirectiveFqdn        types.String `tfsdk:"bind_hostname_directive_fqdn"`
	BlackholeList                    types.List   `tfsdk:"blackhole_list"`
	BlacklistAction                  types.String `tfsdk:"blacklist_action"`
	BlacklistLogQuery                types.Bool   `tfsdk:"blacklist_log_query"`
	BlacklistRedirectAddresses       types.List   `tfsdk:"blacklist_redirect_addresses"`
	BlacklistRedirectTtl             types.Int64  `tfsdk:"blacklist_redirect_ttl"`
	BlacklistRulesets                types.List   `tfsdk:"blacklist_rulesets"`
	CaptureDnsQueriesOnAllDomains    types.Bool   `tfsdk:"capture_dns_queries_on_all_domains"`
	CheckNamesForDdnsAndZoneTransfer types.Bool   `tfsdk:"check_names_for_ddns_and_zone_transfer"`
	CopyClientIpMacOptions           types.Bool   `tfsdk:"copy_client_ip_mac_options"`
	CopyXferToNotify                 types.Bool   `tfsdk:"copy_xfer_to_notify"`
	CustomRootNameServers            types.List   `tfsdk:"custom_root_name_servers"`
	DisableEdns                      types.Bool   `tfsdk:"disable_edns"`
	Dns64Groups                      types.List   `tfsdk:"dns64_groups"`
	DnsCacheAccelerationStatus       types.String `tfsdk:"dns_cache_acceleration_status"`
	DnsCacheAccelerationTtl          types.Int64  `tfsdk:"dns_cache_acceleration_ttl"`
	DnsHealthCheckAnycastControl     types.Bool   `tfsdk:"dns_health_check_anycast_control"`
	DnsHealthCheckDomainList         types.List   `tfsdk:"dns_health_check_domain_list"`
	DnsHealthCheckInterval           types.Int64  `tfsdk:"dns_health_check_interval"`
	DnsHealthCheckRecursionFlag      types.Bool   `tfsdk:"dns_health_check_recursion_flag"`
	DnsHealthCheckRetries            types.Int64  `tfsdk:"dns_health_check_retries"`
	DnsHealthCheckTimeout            types.Int64  `tfsdk:"dns_health_check_timeout"`
	DnsNotifyTransferSource          types.String `tfsdk:"dns_notify_transfer_source"`
	DnsNotifyTransferSourceAddress   types.String `tfsdk:"dns_notify_transfer_source_address"`
	DnsOverTlsService                types.Bool   `tfsdk:"dns_over_tls_service"`
	DnsQueryCaptureFileTimeLimit     types.Int64  `tfsdk:"dns_query_capture_file_time_limit"`
	DnsQuerySourceAddress            types.String `tfsdk:"dns_query_source_address"`
	DnsQuerySourceInterface          types.String `tfsdk:"dns_query_source_interface"`
	DnsViewAddressSettings           types.List   `tfsdk:"dns_view_address_settings"`
	DnssecBlacklistEnabled           types.Bool   `tfsdk:"dnssec_blacklist_enabled"`
	DnssecDns64Enabled               types.Bool   `tfsdk:"dnssec_dns64_enabled"`
	DnssecEnabled                    types.Bool   `tfsdk:"dnssec_enabled"`
	DnssecExpiredSignaturesEnabled   types.Bool   `tfsdk:"dnssec_expired_signatures_enabled"`
	DnssecNegativeTrustAnchors       types.List   `tfsdk:"dnssec_negative_trust_anchors"`
	DnssecNxdomainEnabled            types.Bool   `tfsdk:"dnssec_nxdomain_enabled"`
	DnssecRpzEnabled                 types.Bool   `tfsdk:"dnssec_rpz_enabled"`
	DnssecTrustedKeys                types.List   `tfsdk:"dnssec_trusted_keys"`
	DnssecValidationEnabled          types.Bool   `tfsdk:"dnssec_validation_enabled"`
	DnstapSetting                    types.Object `tfsdk:"dnstap_setting"`
	DohHttpsSessionDuration          types.Int64  `tfsdk:"doh_https_session_duration"`
	DohService                       types.Bool   `tfsdk:"doh_service"`
	DomainsToCaptureDnsQueries       types.List   `tfsdk:"domains_to_capture_dns_queries"`
	DtcDnsQueriesSpecificBehavior    types.String `tfsdk:"dtc_dns_queries_specific_behavior"`
	DtcEdnsPreferClientSubnet        types.Bool   `tfsdk:"dtc_edns_prefer_client_subnet"`
	DtcHealthSource                  types.String `tfsdk:"dtc_health_source"`
	DtcHealthSourceAddress           types.String `tfsdk:"dtc_health_source_address"`
	EdnsUdpSize                      types.Int64  `tfsdk:"edns_udp_size"`
	EnableBlackhole                  types.Bool   `tfsdk:"enable_blackhole"`
	EnableBlacklist                  types.Bool   `tfsdk:"enable_blacklist"`
	EnableCaptureDnsQueries          types.Bool   `tfsdk:"enable_capture_dns_queries"`
	EnableCaptureDnsResponses        types.Bool   `tfsdk:"enable_capture_dns_responses"`
	EnableDns                        types.Bool   `tfsdk:"enable_dns"`
	EnableDns64                      types.Bool   `tfsdk:"enable_dns64"`
	EnableDnsCacheAcceleration       types.Bool   `tfsdk:"enable_dns_cache_acceleration"`
	EnableDnsHealthCheck             types.Bool   `tfsdk:"enable_dns_health_check"`
	EnableDnstapQueries              types.Bool   `tfsdk:"enable_dnstap_queries"`
	EnableDnstapResponses            types.Bool   `tfsdk:"enable_dnstap_responses"`
	EnableDnstapViolationsTls        types.Bool   `tfsdk:"enable_dnstap_violations_tls"`
	EnableExcludedDomainNames        types.Bool   `tfsdk:"enable_excluded_domain_names"`
	EnableFixedRrsetOrderFqdns       types.Bool   `tfsdk:"enable_fixed_rrset_order_fqdns"`
	EnableFtc                        types.Bool   `tfsdk:"enable_ftc"`
	EnableGssTsig                    types.Bool   `tfsdk:"enable_gss_tsig"`
	EnableNotifySourcePort           types.Bool   `tfsdk:"enable_notify_source_port"`
	EnableQueryRewrite               types.Bool   `tfsdk:"enable_query_rewrite"`
	EnableQuerySourcePort            types.Bool   `tfsdk:"enable_query_source_port"`
	ExcludedDomainNames              types.List   `tfsdk:"excluded_domain_names"`
	ExtAttrs                         types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll                      types.Map    `tfsdk:"extattrs_all"`
	FileTransferSetting              types.Object `tfsdk:"file_transfer_setting"`
	FilterAaaa                       types.String `tfsdk:"filter_aaaa"`
	FilterAaaaList                   types.List   `tfsdk:"filter_aaaa_list"`
	FixedRrsetOrderFqdns             types.List   `tfsdk:"fixed_rrset_order_fqdns"`
	ForwardOnly                      types.Bool   `tfsdk:"forward_only"`
	ForwardUpdates                   types.Bool   `tfsdk:"forward_updates"`
	Forwarders                       types.List   `tfsdk:"forwarders"`
	FtcExpiredRecordTimeout          types.Int64  `tfsdk:"ftc_expired_record_timeout"`
	FtcExpiredRecordTtl              types.Int64  `tfsdk:"ftc_expired_record_ttl"`
	GlueRecordAddresses              types.List   `tfsdk:"glue_record_addresses"`
	GssTsigKeys                      types.List   `tfsdk:"gss_tsig_keys"`
	HostName                         types.String `tfsdk:"host_name"`
	Ipv4addr                         types.String `tfsdk:"ipv4addr"`
	Ipv6GlueRecordAddresses          types.List   `tfsdk:"ipv6_glue_record_addresses"`
	Ipv6addr                         types.String `tfsdk:"ipv6addr"`
	LoggingCategories                types.Object `tfsdk:"logging_categories"`
	MaxCacheTtl                      types.Int64  `tfsdk:"max_cache_ttl"`
	MaxCachedLifetime                types.Int64  `tfsdk:"max_cached_lifetime"`
	MaxNcacheTtl                     types.Int64  `tfsdk:"max_ncache_ttl"`
	MaxUdpSize                       types.Int64  `tfsdk:"max_udp_size"`
	MinimalResp                      types.Bool   `tfsdk:"minimal_resp"`
	NotifyDelay                      types.Int64  `tfsdk:"notify_delay"`
	NotifySourcePort                 types.Int64  `tfsdk:"notify_source_port"`
	NxdomainLogQuery                 types.Bool   `tfsdk:"nxdomain_log_query"`
	NxdomainRedirect                 types.Bool   `tfsdk:"nxdomain_redirect"`
	NxdomainRedirectAddresses        types.List   `tfsdk:"nxdomain_redirect_addresses"`
	NxdomainRedirectAddressesV6      types.List   `tfsdk:"nxdomain_redirect_addresses_v6"`
	NxdomainRedirectTtl              types.Int64  `tfsdk:"nxdomain_redirect_ttl"`
	NxdomainRulesets                 types.List   `tfsdk:"nxdomain_rulesets"`
	QuerySourcePort                  types.Int64  `tfsdk:"query_source_port"`
	RecordNamePolicy                 types.String `tfsdk:"record_name_policy"`
	RecursiveClientLimit             types.Int64  `tfsdk:"recursive_client_limit"`
	RecursiveQueryList               types.List   `tfsdk:"recursive_query_list"`
	RecursiveResolver                types.String `tfsdk:"recursive_resolver"`
	ResolverQueryTimeout             types.Int64  `tfsdk:"resolver_query_timeout"`
	ResponseRateLimiting             types.Object `tfsdk:"response_rate_limiting"`
	RootNameServerType               types.String `tfsdk:"root_name_server_type"`
	RpzDisableNsdnameNsip            types.Bool   `tfsdk:"rpz_disable_nsdname_nsip"`
	RpzDropIpRuleEnabled             types.Bool   `tfsdk:"rpz_drop_ip_rule_enabled"`
	RpzDropIpRuleMinPrefixLengthIpv4 types.Int64  `tfsdk:"rpz_drop_ip_rule_min_prefix_length_ipv4"`
	RpzDropIpRuleMinPrefixLengthIpv6 types.Int64  `tfsdk:"rpz_drop_ip_rule_min_prefix_length_ipv6"`
	RpzQnameWaitRecurse              types.Bool   `tfsdk:"rpz_qname_wait_recurse"`
	SerialQueryRate                  types.Int64  `tfsdk:"serial_query_rate"`
	ServerIdDirective                types.String `tfsdk:"server_id_directive"`
	ServerIdDirectiveString          types.String `tfsdk:"server_id_directive_string"`
	SkipInGridRpzQueries             types.Bool   `tfsdk:"skip_in_grid_rpz_queries"`
	Sortlist                         types.List   `tfsdk:"sortlist"`
	StoreLocally                     types.Bool   `tfsdk:"store_locally"`
	SyslogFacility                   types.String `tfsdk:"syslog_facility"`
	TcpIdleTimeout                   types.Int64  `tfsdk:"tcp_idle_timeout"`
	TlsSessionDuration               types.Int64  `tfsdk:"tls_session_duration"`
	TransferExcludedServers          types.List   `tfsdk:"transfer_excluded_servers"`
	TransferFormat                   types.String `tfsdk:"transfer_format"`
	TransfersIn                      types.Int64  `tfsdk:"transfers_in"`
	TransfersOut                     types.Int64  `tfsdk:"transfers_out"`
	TransfersPerNs                   types.Int64  `tfsdk:"transfers_per_ns"`
	UpstreamAddressFamilyPreference  types.String `tfsdk:"upstream_address_family_preference"`
	UseAddClientIpMacOptions         types.Bool   `tfsdk:"use_add_client_ip_mac_options"`
	UseAllowQuery                    types.Bool   `tfsdk:"use_allow_query"`
	UseAllowTransfer                 types.Bool   `tfsdk:"use_allow_transfer"`
	UseAttackMitigation              types.Bool   `tfsdk:"use_attack_mitigation"`
	UseAutoBlackhole                 types.Bool   `tfsdk:"use_auto_blackhole"`
	UseBindHostnameDirective         types.Bool   `tfsdk:"use_bind_hostname_directive"`
	UseBlackhole                     types.Bool   `tfsdk:"use_blackhole"`
	UseBlacklist                     types.Bool   `tfsdk:"use_blacklist"`
	UseCaptureDnsQueriesOnAllDomains types.Bool   `tfsdk:"use_capture_dns_queries_on_all_domains"`
	UseCopyClientIpMacOptions        types.Bool   `tfsdk:"use_copy_client_ip_mac_options"`
	UseCopyXferToNotify              types.Bool   `tfsdk:"use_copy_xfer_to_notify"`
	UseDisableEdns                   types.Bool   `tfsdk:"use_disable_edns"`
	UseDns64                         types.Bool   `tfsdk:"use_dns64"`
	UseDnsCacheAccelerationTtl       types.Bool   `tfsdk:"use_dns_cache_acceleration_ttl"`
	UseDnsHealthCheck                types.Bool   `tfsdk:"use_dns_health_check"`
	UseDnssec                        types.Bool   `tfsdk:"use_dnssec"`
	UseDnstapSetting                 types.Bool   `tfsdk:"use_dnstap_setting"`
	UseDtcDnsQueriesSpecificBehavior types.Bool   `tfsdk:"use_dtc_dns_queries_specific_behavior"`
	UseDtcEdnsPreferClientSubnet     types.Bool   `tfsdk:"use_dtc_edns_prefer_client_subnet"`
	UseEdnsUdpSize                   types.Bool   `tfsdk:"use_edns_udp_size"`
	UseEnableCaptureDns              types.Bool   `tfsdk:"use_enable_capture_dns"`
	UseEnableExcludedDomainNames     types.Bool   `tfsdk:"use_enable_excluded_domain_names"`
	UseEnableGssTsig                 types.Bool   `tfsdk:"use_enable_gss_tsig"`
	UseEnableQueryRewrite            types.Bool   `tfsdk:"use_enable_query_rewrite"`
	UseFilterAaaa                    types.Bool   `tfsdk:"use_filter_aaaa"`
	UseFixedRrsetOrderFqdns          types.Bool   `tfsdk:"use_fixed_rrset_order_fqdns"`
	UseForwardUpdates                types.Bool   `tfsdk:"use_forward_updates"`
	UseForwarders                    types.Bool   `tfsdk:"use_forwarders"`
	UseFtc                           types.Bool   `tfsdk:"use_ftc"`
	UseGssTsigKeys                   types.Bool   `tfsdk:"use_gss_tsig_keys"`
	UseLan2Ipv6Port                  types.Bool   `tfsdk:"use_lan2_ipv6_port"`
	UseLan2Port                      types.Bool   `tfsdk:"use_lan2_port"`
	UseLanIpv6Port                   types.Bool   `tfsdk:"use_lan_ipv6_port"`
	UseLanPort                       types.Bool   `tfsdk:"use_lan_port"`
	UseLoggingCategories             types.Bool   `tfsdk:"use_logging_categories"`
	UseMaxCacheTtl                   types.Bool   `tfsdk:"use_max_cache_ttl"`
	UseMaxCachedLifetime             types.Bool   `tfsdk:"use_max_cached_lifetime"`
	UseMaxNcacheTtl                  types.Bool   `tfsdk:"use_max_ncache_ttl"`
	UseMaxUdpSize                    types.Bool   `tfsdk:"use_max_udp_size"`
	UseMgmtIpv6Port                  types.Bool   `tfsdk:"use_mgmt_ipv6_port"`
	UseMgmtPort                      types.Bool   `tfsdk:"use_mgmt_port"`
	UseNotifyDelay                   types.Bool   `tfsdk:"use_notify_delay"`
	UseNxdomainRedirect              types.Bool   `tfsdk:"use_nxdomain_redirect"`
	UseRecordNamePolicy              types.Bool   `tfsdk:"use_record_name_policy"`
	UseRecursiveClientLimit          types.Bool   `tfsdk:"use_recursive_client_limit"`
	UseRecursiveQuerySetting         types.Bool   `tfsdk:"use_recursive_query_setting"`
	UseResolverQueryTimeout          types.Bool   `tfsdk:"use_resolver_query_timeout"`
	UseResponseRateLimiting          types.Bool   `tfsdk:"use_response_rate_limiting"`
	UseRootNameServer                types.Bool   `tfsdk:"use_root_name_server"`
	UseRootServerForAllViews         types.Bool   `tfsdk:"use_root_server_for_all_views"`
	UseRpzDisableNsdnameNsip         types.Bool   `tfsdk:"use_rpz_disable_nsdname_nsip"`
	UseRpzDropIpRule                 types.Bool   `tfsdk:"use_rpz_drop_ip_rule"`
	UseRpzQnameWaitRecurse           types.Bool   `tfsdk:"use_rpz_qname_wait_recurse"`
	UseSerialQueryRate               types.Bool   `tfsdk:"use_serial_query_rate"`
	UseServerIdDirective             types.Bool   `tfsdk:"use_server_id_directive"`
	UseSortlist                      types.Bool   `tfsdk:"use_sortlist"`
	UseSourcePorts                   types.Bool   `tfsdk:"use_source_ports"`
	UseSyslogFacility                types.Bool   `tfsdk:"use_syslog_facility"`
	UseTransfersIn                   types.Bool   `tfsdk:"use_transfers_in"`
	UseTransfersOut                  types.Bool   `tfsdk:"use_transfers_out"`
	UseTransfersPerNs                types.Bool   `tfsdk:"use_transfers_per_ns"`
	UseUpdateSetting                 types.Bool   `tfsdk:"use_update_setting"`
	UseZoneTransferFormat            types.Bool   `tfsdk:"use_zone_transfer_format"`
	Views                            types.List   `tfsdk:"views"`
}

var MemberDnsAttrTypes = map[string]attr.Type{
	"ref":                                     types.StringType,
	"add_client_ip_mac_options":               types.BoolType,
	"additional_ip_list":                      types.ListType{ElemType: types.StringType},
	"additional_ip_list_struct":               types.ListType{ElemType: types.ObjectType{AttrTypes: MemberDnsAdditionalIpListStructAttrTypes}},
	"allow_gss_tsig_zone_updates":             types.BoolType,
	"allow_query":                             types.ListType{ElemType: types.ObjectType{AttrTypes: MemberDnsAllowQueryAttrTypes}},
	"allow_recursive_query":                   types.BoolType,
	"allow_transfer":                          types.ListType{ElemType: types.ObjectType{AttrTypes: MemberDnsAllowTransferAttrTypes}},
	"allow_update":                            types.ListType{ElemType: types.ObjectType{AttrTypes: MemberDnsAllowUpdateAttrTypes}},
	"anonymize_response_logging":              types.BoolType,
	"atc_fwd_enable":                          types.BoolType,
	"attack_mitigation":                       types.ObjectType{AttrTypes: MemberDnsAttackMitigationAttrTypes},
	"auto_blackhole":                          types.ObjectType{AttrTypes: MemberDnsAutoBlackholeAttrTypes},
	"auto_create_a_and_ptr_for_lan2":          types.BoolType,
	"auto_create_aaaa_and_ipv6ptr_for_lan2":   types.BoolType,
	"auto_sort_views":                         types.BoolType,
	"bind_check_names_policy":                 types.StringType,
	"bind_hostname_directive":                 types.StringType,
	"bind_hostname_directive_fqdn":            types.StringType,
	"blackhole_list":                          types.ListType{ElemType: types.ObjectType{AttrTypes: MemberDnsBlackholeListAttrTypes}},
	"blacklist_action":                        types.StringType,
	"blacklist_log_query":                     types.BoolType,
	"blacklist_redirect_addresses":            types.ListType{ElemType: types.StringType},
	"blacklist_redirect_ttl":                  types.Int64Type,
	"blacklist_rulesets":                      types.ListType{ElemType: types.StringType},
	"capture_dns_queries_on_all_domains":      types.BoolType,
	"check_names_for_ddns_and_zone_transfer":  types.BoolType,
	"copy_client_ip_mac_options":              types.BoolType,
	"copy_xfer_to_notify":                     types.BoolType,
	"custom_root_name_servers":                types.ListType{ElemType: types.ObjectType{AttrTypes: MemberDnsCustomRootNameServersAttrTypes}},
	"disable_edns":                            types.BoolType,
	"dns64_groups":                            types.ListType{ElemType: types.StringType},
	"dns_cache_acceleration_status":           types.StringType,
	"dns_cache_acceleration_ttl":              types.Int64Type,
	"dns_health_check_anycast_control":        types.BoolType,
	"dns_health_check_domain_list":            types.ListType{ElemType: types.StringType},
	"dns_health_check_interval":               types.Int64Type,
	"dns_health_check_recursion_flag":         types.BoolType,
	"dns_health_check_retries":                types.Int64Type,
	"dns_health_check_timeout":                types.Int64Type,
	"dns_notify_transfer_source":              types.StringType,
	"dns_notify_transfer_source_address":      types.StringType,
	"dns_over_tls_service":                    types.BoolType,
	"dns_query_capture_file_time_limit":       types.Int64Type,
	"dns_query_source_address":                types.StringType,
	"dns_query_source_interface":              types.StringType,
	"dns_view_address_settings":               types.ListType{ElemType: types.ObjectType{AttrTypes: MemberDnsDnsViewAddressSettingsAttrTypes}},
	"dnssec_blacklist_enabled":                types.BoolType,
	"dnssec_dns64_enabled":                    types.BoolType,
	"dnssec_enabled":                          types.BoolType,
	"dnssec_expired_signatures_enabled":       types.BoolType,
	"dnssec_negative_trust_anchors":           types.ListType{ElemType: types.StringType},
	"dnssec_nxdomain_enabled":                 types.BoolType,
	"dnssec_rpz_enabled":                      types.BoolType,
	"dnssec_trusted_keys":                     types.ListType{ElemType: types.ObjectType{AttrTypes: MemberDnsDnssecTrustedKeysAttrTypes}},
	"dnssec_validation_enabled":               types.BoolType,
	"dnstap_setting":                          types.ObjectType{AttrTypes: MemberDnsDnstapSettingAttrTypes},
	"doh_https_session_duration":              types.Int64Type,
	"doh_service":                             types.BoolType,
	"domains_to_capture_dns_queries":          types.ListType{ElemType: types.StringType},
	"dtc_dns_queries_specific_behavior":       types.StringType,
	"dtc_edns_prefer_client_subnet":           types.BoolType,
	"dtc_health_source":                       types.StringType,
	"dtc_health_source_address":               types.StringType,
	"edns_udp_size":                           types.Int64Type,
	"enable_blackhole":                        types.BoolType,
	"enable_blacklist":                        types.BoolType,
	"enable_capture_dns_queries":              types.BoolType,
	"enable_capture_dns_responses":            types.BoolType,
	"enable_dns":                              types.BoolType,
	"enable_dns64":                            types.BoolType,
	"enable_dns_cache_acceleration":           types.BoolType,
	"enable_dns_health_check":                 types.BoolType,
	"enable_dnstap_queries":                   types.BoolType,
	"enable_dnstap_responses":                 types.BoolType,
	"enable_dnstap_violations_tls":            types.BoolType,
	"enable_excluded_domain_names":            types.BoolType,
	"enable_fixed_rrset_order_fqdns":          types.BoolType,
	"enable_ftc":                              types.BoolType,
	"enable_gss_tsig":                         types.BoolType,
	"enable_notify_source_port":               types.BoolType,
	"enable_query_rewrite":                    types.BoolType,
	"enable_query_source_port":                types.BoolType,
	"excluded_domain_names":                   types.ListType{ElemType: types.StringType},
	"extattrs":                                types.MapType{ElemType: types.StringType},
	"file_transfer_setting":                   types.ObjectType{AttrTypes: MemberDnsFileTransferSettingAttrTypes},
	"filter_aaaa":                             types.StringType,
	"filter_aaaa_list":                        types.ListType{ElemType: types.ObjectType{AttrTypes: MemberDnsFilterAaaaListAttrTypes}},
	"fixed_rrset_order_fqdns":                 types.ListType{ElemType: types.ObjectType{AttrTypes: MemberDnsFixedRrsetOrderFqdnsAttrTypes}},
	"forward_only":                            types.BoolType,
	"forward_updates":                         types.BoolType,
	"forwarders":                              types.ListType{ElemType: types.StringType},
	"ftc_expired_record_timeout":              types.Int64Type,
	"ftc_expired_record_ttl":                  types.Int64Type,
	"glue_record_addresses":                   types.ListType{ElemType: types.ObjectType{AttrTypes: MemberDnsGlueRecordAddressesAttrTypes}},
	"gss_tsig_keys":                           types.ListType{ElemType: types.StringType},
	"host_name":                               types.StringType,
	"ipv4addr":                                types.StringType,
	"ipv6_glue_record_addresses":              types.ListType{ElemType: types.ObjectType{AttrTypes: MemberDnsIpv6GlueRecordAddressesAttrTypes}},
	"ipv6addr":                                types.StringType,
	"logging_categories":                      types.ObjectType{AttrTypes: MemberDnsLoggingCategoriesAttrTypes},
	"max_cache_ttl":                           types.Int64Type,
	"max_cached_lifetime":                     types.Int64Type,
	"max_ncache_ttl":                          types.Int64Type,
	"max_udp_size":                            types.Int64Type,
	"minimal_resp":                            types.BoolType,
	"notify_delay":                            types.Int64Type,
	"notify_source_port":                      types.Int64Type,
	"nxdomain_log_query":                      types.BoolType,
	"nxdomain_redirect":                       types.BoolType,
	"nxdomain_redirect_addresses":             types.ListType{ElemType: types.StringType},
	"nxdomain_redirect_addresses_v6":          types.ListType{ElemType: types.StringType},
	"nxdomain_redirect_ttl":                   types.Int64Type,
	"nxdomain_rulesets":                       types.ListType{ElemType: types.StringType},
	"query_source_port":                       types.Int64Type,
	"record_name_policy":                      types.StringType,
	"recursive_client_limit":                  types.Int64Type,
	"recursive_query_list":                    types.ListType{ElemType: types.ObjectType{AttrTypes: MemberDnsRecursiveQueryListAttrTypes}},
	"recursive_resolver":                      types.StringType,
	"resolver_query_timeout":                  types.Int64Type,
	"response_rate_limiting":                  types.ObjectType{AttrTypes: MemberDnsResponseRateLimitingAttrTypes},
	"root_name_server_type":                   types.StringType,
	"rpz_disable_nsdname_nsip":                types.BoolType,
	"rpz_drop_ip_rule_enabled":                types.BoolType,
	"rpz_drop_ip_rule_min_prefix_length_ipv4": types.Int64Type,
	"rpz_drop_ip_rule_min_prefix_length_ipv6": types.Int64Type,
	"rpz_qname_wait_recurse":                  types.BoolType,
	"serial_query_rate":                       types.Int64Type,
	"server_id_directive":                     types.StringType,
	"server_id_directive_string":              types.StringType,
	"skip_in_grid_rpz_queries":                types.BoolType,
	"sortlist":                                types.ListType{ElemType: types.ObjectType{AttrTypes: MemberDnsSortlistAttrTypes}},
	"store_locally":                           types.BoolType,
	"syslog_facility":                         types.StringType,
	"tcp_idle_timeout":                        types.Int64Type,
	"tls_session_duration":                    types.Int64Type,
	"transfer_excluded_servers":               types.ListType{ElemType: types.StringType},
	"transfer_format":                         types.StringType,
	"transfers_in":                            types.Int64Type,
	"transfers_out":                           types.Int64Type,
	"transfers_per_ns":                        types.Int64Type,
	"upstream_address_family_preference":      types.StringType,
	"use_add_client_ip_mac_options":           types.BoolType,
	"use_allow_query":                         types.BoolType,
	"use_allow_transfer":                      types.BoolType,
	"use_attack_mitigation":                   types.BoolType,
	"use_auto_blackhole":                      types.BoolType,
	"use_bind_hostname_directive":             types.BoolType,
	"use_blackhole":                           types.BoolType,
	"use_blacklist":                           types.BoolType,
	"use_capture_dns_queries_on_all_domains":  types.BoolType,
	"use_copy_client_ip_mac_options":          types.BoolType,
	"use_copy_xfer_to_notify":                 types.BoolType,
	"use_disable_edns":                        types.BoolType,
	"use_dns64":                               types.BoolType,
	"use_dns_cache_acceleration_ttl":          types.BoolType,
	"use_dns_health_check":                    types.BoolType,
	"use_dnssec":                              types.BoolType,
	"use_dnstap_setting":                      types.BoolType,
	"use_dtc_dns_queries_specific_behavior":   types.BoolType,
	"use_dtc_edns_prefer_client_subnet":       types.BoolType,
	"use_edns_udp_size":                       types.BoolType,
	"use_enable_capture_dns":                  types.BoolType,
	"use_enable_excluded_domain_names":        types.BoolType,
	"use_enable_gss_tsig":                     types.BoolType,
	"use_enable_query_rewrite":                types.BoolType,
	"use_filter_aaaa":                         types.BoolType,
	"use_fixed_rrset_order_fqdns":             types.BoolType,
	"use_forward_updates":                     types.BoolType,
	"use_forwarders":                          types.BoolType,
	"use_ftc":                                 types.BoolType,
	"use_gss_tsig_keys":                       types.BoolType,
	"use_lan2_ipv6_port":                      types.BoolType,
	"use_lan2_port":                           types.BoolType,
	"use_lan_ipv6_port":                       types.BoolType,
	"use_lan_port":                            types.BoolType,
	"use_logging_categories":                  types.BoolType,
	"use_max_cache_ttl":                       types.BoolType,
	"use_max_cached_lifetime":                 types.BoolType,
	"use_max_ncache_ttl":                      types.BoolType,
	"use_max_udp_size":                        types.BoolType,
	"use_mgmt_ipv6_port":                      types.BoolType,
	"use_mgmt_port":                           types.BoolType,
	"use_notify_delay":                        types.BoolType,
	"use_nxdomain_redirect":                   types.BoolType,
	"use_record_name_policy":                  types.BoolType,
	"use_recursive_client_limit":              types.BoolType,
	"use_recursive_query_setting":             types.BoolType,
	"use_resolver_query_timeout":              types.BoolType,
	"use_response_rate_limiting":              types.BoolType,
	"use_root_name_server":                    types.BoolType,
	"use_root_server_for_all_views":           types.BoolType,
	"use_rpz_disable_nsdname_nsip":            types.BoolType,
	"use_rpz_drop_ip_rule":                    types.BoolType,
	"use_rpz_qname_wait_recurse":              types.BoolType,
	"use_serial_query_rate":                   types.BoolType,
	"use_server_id_directive":                 types.BoolType,
	"use_sortlist":                            types.BoolType,
	"use_source_ports":                        types.BoolType,
	"use_syslog_facility":                     types.BoolType,
	"use_transfers_in":                        types.BoolType,
	"use_transfers_out":                       types.BoolType,
	"use_transfers_per_ns":                    types.BoolType,
	"use_update_setting":                      types.BoolType,
	"use_zone_transfer_format":                types.BoolType,
	"views":                                   types.ListType{ElemType: types.StringType},
}

var MemberDnsResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"add_client_ip_mac_options": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Add custom IP, MAC and DNS View name ENDS0 options to outgoing recursive queries.",
	},
	"additional_ip_list": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of additional IP addresses on which DNS is enabled for a Grid member. Only one of \"additional_ip_list\" or \"additional_ip_list_struct\" should be set when modifying the object.",
	},
	"additional_ip_list_struct": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberDnsAdditionalIpListStructResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of additional IP addresses and IP Space Discriminator short names on which DNS is enabled for a Grid member. Only one of \"additional_ip_list\" or \"additional_ip_list_struct\" should be set when modifying the object.",
	},
	"allow_gss_tsig_zone_updates": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the GSS-TSIG zone updates is enabled for the Grid member.",
	},
	"allow_query": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberDnsAllowQueryResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "Determines if queries from specified IPv4 or IPv6 addresses and networks are enabled or not. The appliance can also use Transaction Signature (TSIG) keys to authenticate the queries. This setting overrides the Grid query settings.",
	},
	"allow_recursive_query": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the responses to recursive queries is enabled or not. This setting overrides Grid recursive query settings.",
	},
	"allow_transfer": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberDnsAllowTransferResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "Allows or rejects zone transfers from specified IPv4 or IPv6 addresses and networks or allows transfers from hosts authenticated by Transaction signature (TSIG) key. This setting overrides the Grid zone transfer settings.",
	},
	"allow_update": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberDnsAllowUpdateResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "Allows or rejects dynamic updates from specified IPv4 or IPv6 addresses, networks or from host authenticated by TSIG key. This setting overrides Grid update settings.",
	},
	"anonymize_response_logging": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "The flag that indicates whether the anonymization of captured DNS responses is enabled or disabled.",
	},
	"atc_fwd_enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enable DNS recursive query forwarding to Active Trust Cloud.",
	},
	"attack_mitigation": schema.SingleNestedAttribute{
		Attributes: MemberDnsAttackMitigationResourceSchemaAttributes,
		Optional:   true,
	},
	"auto_blackhole": schema.SingleNestedAttribute{
		Attributes: MemberDnsAutoBlackholeResourceSchemaAttributes,
		Optional:   true,
	},
	"auto_create_a_and_ptr_for_lan2": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the auto-generation of A and PTR records for the LAN2 IP address is enabled or not, if DNS service is enabled on LAN2.",
	},
	"auto_create_aaaa_and_ipv6ptr_for_lan2": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if auto-generation of AAAA and IPv6 PTR records for LAN2 IPv6 address is enabled or not.",
	},
	"auto_sort_views": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if a Grid member to automatically sort DNS views is enabled or not. The order of the DNS views determines the order in which the appliance checks the match lists.",
	},
	"bind_check_names_policy": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The BIND check names policy, which indicates the action the appliance takes when it encounters host names that do not comply with the Strict Hostname Checking policy. This method applies only if the host name restriction policy is set to 'Strict Hostname Checking'.",
	},
	"bind_hostname_directive": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The value of the hostname directive for BIND.",
	},
	"bind_hostname_directive_fqdn": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The value of the user-defined hostname directive for BIND. To enable user-defined hostname directive, you must set the bind_hostname_directive to \"USER_DEFINED\".",
	},
	"blackhole_list": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberDnsBlackholeListResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of IPv4 or IPv6 addresses and networks from which DNS queries are blocked. This setting overrides the Grid blackhole_list.",
	},
	"blacklist_action": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The action to perform when a domain name matches the pattern defined in a rule that is specified by the blacklist_ruleset method.",
	},
	"blacklist_log_query": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if blacklist redirection queries are logged or not.",
	},
	"blacklist_redirect_addresses": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The IP addresses the appliance includes in the response it sends in place of a blacklisted IP address.",
	},
	"blacklist_redirect_ttl": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The TTL value of the synthetic DNS responses that result from blacklist redirection.",
	},
	"blacklist_rulesets": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The DNS Ruleset object names assigned at the Grid level for blacklist redirection.",
	},
	"capture_dns_queries_on_all_domains": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "The flag that indicates whether the capture of DNS queries for all domains is enabled or disabled.",
	},
	"check_names_for_ddns_and_zone_transfer": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the application of BIND check-names for zone transfers and DDNS updates are enabled.",
	},
	"copy_client_ip_mac_options": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Copy custom IP, MAC and DNS View name ENDS0 options from incoming to outgoing recursive queries.",
	},
	"copy_xfer_to_notify": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Copies the allowed IPs from the zone transfer list into the also-notify statement in the named.conf file.",
	},
	"custom_root_name_servers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberDnsCustomRootNameServersResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of custom root name servers. You can either select and use Internet root name servers or specify custom root name servers by providing a host name and IP address to which the Infoblox appliance can send queries.",
	},
	"disable_edns": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "The EDNS0 support for queries that require recursive resolution on Grid members.",
	},
	"dns64_groups": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of DNS64 synthesis groups associated with this member.",
	},
	"dns_cache_acceleration_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The DNS cache acceleration status.",
	},
	"dns_cache_acceleration_ttl": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The minimum TTL value, in seconds, that a DNS record must have in order for it to be cached by the DNS Cache Acceleration service. An integer from 1 to 65000 that represents the TTL in seconds.",
	},
	"dns_health_check_anycast_control": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "The flag that indicates whether the anycast failure (BFD session down) is enabled on member failure or not.",
	},
	"dns_health_check_domain_list": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of domain names for the DNS health check.",
	},
	"dns_health_check_interval": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The time interval (in seconds) for DNS health check.",
	},
	"dns_health_check_recursion_flag": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "The flag that indicates whether the recursive DNS health check is enabled or not.",
	},
	"dns_health_check_retries": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The number of DNS health check retries.",
	},
	"dns_health_check_timeout": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The DNS health check timeout interval (in seconds).",
	},
	"dns_notify_transfer_source": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Determines which IP address is used as the source for DDNS notify and transfer operations.",
	},
	"dns_notify_transfer_source_address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The source address used if dns_notify_transfer_source type is \"IP\".",
	},
	"dns_over_tls_service": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enables DNS over TLS service.",
	},
	"dns_query_capture_file_time_limit": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The time limit (in minutes) for the DNS query capture file.",
	},
	"dns_query_source_address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The source address used if dns_query_source_interface type is \"IP\".",
	},
	"dns_query_source_interface": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Determines which IP address is used as the source for DDNS query operations.",
	},
	"dns_view_address_settings": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberDnsDnsViewAddressSettingsResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "Array of notify/query source settings for views.",
	},
	"dnssec_blacklist_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the blacklist rules for DNSSEC-enabled clients are enabled or not.",
	},
	"dnssec_dns64_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the DNS64 groups for DNSSEC-enabled clients are enabled or not.",
	},
	"dnssec_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the DNS security extension is enabled or not.",
	},
	"dnssec_expired_signatures_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines when the DNS member accepts expired signatures.",
	},
	"dnssec_negative_trust_anchors": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "A list of zones for which the server does not perform DNSSEC validation.",
	},
	"dnssec_nxdomain_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the NXDOMAIN rules for DNSSEC-enabled clients are enabled or not.",
	},
	"dnssec_rpz_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the RPZ policies for DNSSEC-enabled clients are enabled or not.",
	},
	"dnssec_trusted_keys": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberDnsDnssecTrustedKeysResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of trusted keys for the DNSSEC feature.",
	},
	"dnssec_validation_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the DNS security validation is enabled or not.",
	},
	"dnstap_setting": schema.SingleNestedAttribute{
		Attributes: MemberDnsDnstapSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"doh_https_session_duration": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "DNS over HTTPS sessions duration.",
	},
	"doh_service": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enables DNS over HTTPS service.",
	},
	"domains_to_capture_dns_queries": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of domains for DNS query capture.",
	},
	"dtc_dns_queries_specific_behavior": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Setting to control specific behavior for DTC DNS responses for incoming lbdn matched queries.",
	},
	"dtc_edns_prefer_client_subnet": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether to prefer the client address from the edns-client-subnet option for DTC or not.",
	},
	"dtc_health_source": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The health check source type.",
	},
	"dtc_health_source_address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The source address used if dtc_health_source type is \"IP\".",
	},
	"edns_udp_size": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Advertises the EDNS0 buffer size to the upstream server. The value should be between 512 and 4096 bytes. The recommended value is between 512 and 1220 bytes.",
	},
	"enable_blackhole": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the blocking of DNS queries is enabled or not. This setting overrides the Grid enable_blackhole settings.",
	},
	"enable_blacklist": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if a blacklist is enabled or not on the Grid member.",
	},
	"enable_capture_dns_queries": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "The flag that indicates whether the capture of DNS queries is enabled or disabled.",
	},
	"enable_capture_dns_responses": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "The flag that indicates whether the capture of DNS responses is enabled or disabled.",
	},
	"enable_dns": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the DNS service of a member is enabled or not.",
	},
	"enable_dns64": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the DNS64 support is enabled or not for this member.",
	},
	"enable_dns_cache_acceleration": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the DNS Cache Acceleration service is enabled or not for a member.",
	},
	"enable_dns_health_check": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "The flag that indicates whether the DNS health check is enabled or not.",
	},
	"enable_dnstap_queries": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the query messages need to be forwarded to DNSTAP or not.",
	},
	"enable_dnstap_responses": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the response messages need to be forwarded to DNSTAP or not.",
	},
	"enable_dnstap_violations_tls": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the violations messages need to be forwarded to DNSTAP or not.",
	},
	"enable_excluded_domain_names": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "The flag that indicates whether excluding domain names from captured DNS queries and responses is enabled or disabled.",
	},
	"enable_fixed_rrset_order_fqdns": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the fixed RRset order FQDN is enabled or not.",
	},
	"enable_ftc": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether Fault Tolerant Caching (FTC) is enabled.",
	},
	"enable_gss_tsig": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the appliance is enabled to receive GSS-TSIG authenticated updates from DHCP clients.",
	},
	"enable_notify_source_port": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the notify source port for a member is enabled or not.",
	},
	"enable_query_rewrite": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the DNS query rewrite is enabled or not for this member.",
	},
	"enable_query_source_port": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the query source port for a memer is enabled or not.",
	},
	"excluded_domain_names": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of domains that are excluded from DNS query and response capture.",
	},
	"extattrs": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.",
	},
	"extattrs_all": schema.MapAttribute{
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object , including default attributes.",
		ElementType:         types.StringType,
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
	},
	"file_transfer_setting": schema.SingleNestedAttribute{
		Attributes: MemberDnsFileTransferSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"filter_aaaa": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The type of AAAA filtering for this member DNS object.",
	},
	"filter_aaaa_list": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberDnsFilterAaaaListResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of IPv4 addresses and networks from which queries are received. AAAA filtering is applied to these addresses.",
	},
	"fixed_rrset_order_fqdns": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberDnsFixedRrsetOrderFqdnsResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The fixed RRset order FQDN. If this field does not contain an empty value, the appliance will automatically set the enable_fixed_rrset_order_fqdns field to 'true', unless the same request sets the enable field to 'false'.",
	},
	"forward_only": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Permits this member to send queries to forwarders only. When the value is \"true\", the member sends queries to forwarders only, and not to other internal or Internet root servers.",
	},
	"forward_updates": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Allows secondary servers to forward updates to the DNS server. This setting overrides grid update settings.",
	},
	"forwarders": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The forwarders for the member. A forwarder is essentially a name server to which other name servers first send all of their off-site queries. The forwarder builds up a cache of information, avoiding the need for the other name servers to send queries off-site. This setting overrides the Grid level setting.",
	},
	"ftc_expired_record_timeout": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The timeout interval (in seconds) after which the expired Fault Tolerant Caching (FTC)record is stale and no longer valid.",
	},
	"ftc_expired_record_ttl": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The TTL value (in seconds) of the expired Fault Tolerant Caching (FTC) record in DNS responses.",
	},
	"glue_record_addresses": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberDnsGlueRecordAddressesResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of glue record addresses.",
	},
	"gss_tsig_keys": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of GSS-TSIG keys for a member DNS object.",
	},
	"host_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The host name of the Grid member.",
	},
	"ipv4addr": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The IPv4 Address of the Grid member.",
	},
	"ipv6_glue_record_addresses": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberDnsIpv6GlueRecordAddressesResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of IPv6 glue record addresses.",
	},
	"ipv6addr": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The IPv6 Address of the Grid member.",
	},
	"logging_categories": schema.SingleNestedAttribute{
		Attributes: MemberDnsLoggingCategoriesResourceSchemaAttributes,
		Optional:   true,
	},
	"max_cache_ttl": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The maximum time (in seconds) for which the server will cache positive answers.",
	},
	"max_cached_lifetime": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The maximum time in seconds a DNS response can be stored in the hardware acceleration cache. Valid values are unsigned integer between 60 and 86400, inclusive.",
	},
	"max_ncache_ttl": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The maximum time (in seconds) for which the server will cache negative (NXDOMAIN) responses. The maximum allowed value is 604800.",
	},
	"max_udp_size": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The value is used by authoritative DNS servers to never send DNS responses larger than the configured value. The value should be between 512 and 4096 bytes. The recommended value is between 512 and 1220 bytes.",
	},
	"minimal_resp": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enables the ability to return a minimal amount of data in response to a query. This capability speeds up the DNS services provided by the appliance.",
	},
	"notify_delay": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Specifies the number of seconds of delay the notify messages are sent to secondaries.",
	},
	"notify_source_port": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The source port for notify messages. When requesting zone transfers from the primary server, some secondary DNS servers use the source port number (the primary server used to send the notify message) as the destination port number in the zone transfer request. This setting overrides Grid static source port settings. Valid values are between 1 and 63999. The default is selected by BIND.",
	},
	"nxdomain_log_query": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if NXDOMAIN redirection queries are logged or not.",
	},
	"nxdomain_redirect": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enables NXDOMAIN redirection.",
	},
	"nxdomain_redirect_addresses": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The IPv4 NXDOMAIN redirection addresses.",
	},
	"nxdomain_redirect_addresses_v6": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The IPv6 NXDOMAIN redirection addresses.",
	},
	"nxdomain_redirect_ttl": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The TTL value of synthetic DNS responses that result from NXDOMAIN redirection.",
	},
	"nxdomain_rulesets": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The names of the Ruleset objects assigned at the Grid level for NXDOMAIN redirection.",
	},
	"query_source_port": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The source port for queries. Specifying a source port number for recursive queries ensures that a firewall will allow the response. Valid values are between 1 and 63999. The default is selected by BIND.",
	},
	"record_name_policy": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The record name restriction policy.",
	},
	"recursive_client_limit": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "A limit on the number of concurrent recursive clients.",
	},
	"recursive_query_list": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberDnsRecursiveQueryListResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of IPv4 or IPv6 addresses, networks or hosts authenticated by Transaction signature (TSIG) key from which recursive queries are allowed or denied.",
	},
	"recursive_resolver": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The recursive resolver for member DNS. UNBOUND support has been deprecated from NIOS 9.0 onwards.",
	},
	"resolver_query_timeout": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The recursive query timeout for the member. The value must be 0 or between 10 and 30.",
	},
	"response_rate_limiting": schema.SingleNestedAttribute{
		Attributes: MemberDnsResponseRateLimitingResourceSchemaAttributes,
		Optional:   true,
	},
	"root_name_server_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Determines the type of root name servers.",
	},
	"rpz_disable_nsdname_nsip": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enables NSDNAME and NSIP resource records from RPZ feeds at member level.",
	},
	"rpz_drop_ip_rule_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enables the appliance to ignore RPZ-IP triggers with prefix lengths less than the specified minimum prefix length.",
	},
	"rpz_drop_ip_rule_min_prefix_length_ipv4": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The minimum prefix length for IPv4 RPZ-IP triggers. The appliance ignores RPZ-IP triggers with prefix lengths less than the specified minimum IPv4 prefix length.",
	},
	"rpz_drop_ip_rule_min_prefix_length_ipv6": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The minimum prefix length for IPv6 RPZ-IP triggers. The appliance ignores RPZ-IP triggers with prefix lengths less than the specified minimum IPv6 prefix length.",
	},
	"rpz_qname_wait_recurse": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "The flag that indicates whether recursive RPZ lookups are enabled.",
	},
	"serial_query_rate": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The number of maximum concurrent SOA queries per second for the member.",
	},
	"server_id_directive": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The value of the server-id directive for BIND and Unbound DNS.",
	},
	"server_id_directive_string": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The value of the user-defined hostname directive for BIND DNS. To enable user-defined hostname directive, you must set the bind_hostname_directive to \"USER_DEFINED\".",
	},
	"skip_in_grid_rpz_queries": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if RPZ rules are applied to queries originated from this member and received by other Grid members.",
	},
	"sortlist": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberDnsSortlistResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "A sort list determines the order of addresses in responses made to DNS queries. This setting overrides Grid sort list settings.",
	},
	"store_locally": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "The flag that indicates whether the storage of query capture reports on the appliance is enabled or disabled.",
	},
	"syslog_facility": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The syslog facility. This is the location on the syslog server to which you want to sort the DNS logging messages. This setting overrides the Grid logging facility settings.",
	},
	"tcp_idle_timeout": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "TCP Idle timeout for DNS over TLS connections.",
	},
	"tls_session_duration": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "DNS over TLS sessions duration.",
	},
	"transfer_excluded_servers": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "Excludes specified DNS servers during zone transfers.",
	},
	"transfer_format": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The BIND format for a zone transfer. This provides tracking capabilities for single or multiple transfers and their associated servers.",
	},
	"transfers_in": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The number of maximum concurrent transfers for the member.",
	},
	"transfers_out": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The number of maximum outbound concurrent zone transfers for the member.",
	},
	"transfers_per_ns": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The number of maximum concurrent transfers per member for the member.",
	},
	"upstream_address_family_preference": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Upstream address family preference when dual mode is configured.",
	},
	"use_add_client_ip_mac_options": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: add_client_ip_mac_options",
	},
	"use_allow_query": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: allow_query",
	},
	"use_allow_transfer": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: allow_transfer",
	},
	"use_attack_mitigation": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: attack_mitigation",
	},
	"use_auto_blackhole": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: auto_blackhole",
	},
	"use_bind_hostname_directive": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: bind_hostname_directive",
	},
	"use_blackhole": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: enable_blackhole",
	},
	"use_blacklist": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: blackhole_list , blacklist_action, blacklist_log_query, blacklist_redirect_addresses, blacklist_redirect_ttl, blacklist_rulesets, enable_blacklist",
	},
	"use_capture_dns_queries_on_all_domains": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: capture_dns_queries_on_all_domains",
	},
	"use_copy_client_ip_mac_options": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: copy_client_ip_mac_options",
	},
	"use_copy_xfer_to_notify": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: copy_xfer_to_notify",
	},
	"use_disable_edns": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: disable_edns",
	},
	"use_dns64": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: enable_dns64 , dns64_groups",
	},
	"use_dns_cache_acceleration_ttl": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: dns_cache_acceleration_ttl",
	},
	"use_dns_health_check": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: dns_health_check_domain_list , dns_health_check_recursion_flag, dns_health_check_anycast_control, enable_dns_health_check, dns_health_check_interval, dns_health_check_timeout, dns_health_check_retries",
	},
	"use_dnssec": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: dnssec_enabled , dnssec_expired_signatures_enabled, dnssec_validation_enabled, dnssec_trusted_keys",
	},
	"use_dnstap_setting": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: enable_dnstap_queries , enable_dnstap_responses, enable_dnstap_violations_tls, dnstap_setting",
	},
	"use_dtc_dns_queries_specific_behavior": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: dtc_dns_queries_specific_behavior",
	},
	"use_dtc_edns_prefer_client_subnet": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: dtc_edns_prefer_client_subnet",
	},
	"use_edns_udp_size": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: edns_udp_size",
	},
	"use_enable_capture_dns": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: enable_capture_dns_queries , enable_capture_dns_responses",
	},
	"use_enable_excluded_domain_names": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: enable_excluded_domain_names",
	},
	"use_enable_gss_tsig": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: enable_gss_tsig",
	},
	"use_enable_query_rewrite": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: enable_query_rewrite",
	},
	"use_filter_aaaa": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: filter_aaaa , filter_aaaa_list",
	},
	"use_fixed_rrset_order_fqdns": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: fixed_rrset_order_fqdns , enable_fixed_rrset_order_fqdns",
	},
	"use_forward_updates": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: forward_updates",
	},
	"use_forwarders": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: forwarders , forward_only",
	},
	"use_ftc": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: enable_ftc , ftc_expired_record_ttl, ftc_expired_record_timeout",
	},
	"use_gss_tsig_keys": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: gss_tsig_keys",
	},
	"use_lan2_ipv6_port": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the DNS service on the IPv6 LAN2 port is enabled or not.",
	},
	"use_lan2_port": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the DNS service on the LAN2 port is enabled or not.",
	},
	"use_lan_ipv6_port": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the DNS service on the IPv6 LAN port is enabled or not.",
	},
	"use_lan_port": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines the status of the use of DNS services on the IPv4 LAN1 port.",
	},
	"use_logging_categories": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: logging_categories",
	},
	"use_max_cache_ttl": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: max_cache_ttl",
	},
	"use_max_cached_lifetime": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: max_cached_lifetime",
	},
	"use_max_ncache_ttl": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: max_ncache_ttl",
	},
	"use_max_udp_size": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: max_udp_size",
	},
	"use_mgmt_ipv6_port": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the DNS services on the IPv6 MGMT port is enabled or not.",
	},
	"use_mgmt_port": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the DNS services on the MGMT port is enabled or not.",
	},
	"use_notify_delay": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: notify_delay",
	},
	"use_nxdomain_redirect": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: nxdomain_redirect , nxdomain_redirect_addresses, nxdomain_redirect_addresses_v6, nxdomain_redirect_ttl, nxdomain_log_query, nxdomain_rulesets",
	},
	"use_record_name_policy": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: record_name_policy",
	},
	"use_recursive_client_limit": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: recursive_client_limit",
	},
	"use_recursive_query_setting": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: allow_recursive_query , recursive_query_list",
	},
	"use_resolver_query_timeout": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: resolver_query_timeout",
	},
	"use_response_rate_limiting": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: response_rate_limiting",
	},
	"use_root_name_server": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: root_name_server_type , custom_root_name_servers, use_root_server_for_all_views",
	},
	"use_root_server_for_all_views": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if root name servers should be applied to all views or only to Default view.",
	},
	"use_rpz_disable_nsdname_nsip": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: rpz_disable_nsdname_nsip",
	},
	"use_rpz_drop_ip_rule": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: rpz_drop_ip_rule_enabled , rpz_drop_ip_rule_min_prefix_length_ipv4, rpz_drop_ip_rule_min_prefix_length_ipv6",
	},
	"use_rpz_qname_wait_recurse": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: rpz_qname_wait_recurse",
	},
	"use_serial_query_rate": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: serial_query_rate",
	},
	"use_server_id_directive": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: server_id_directive",
	},
	"use_sortlist": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: sortlist",
	},
	"use_source_ports": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: enable_notify_source_port , notify_source_port, enable_query_source_port, query_source_port",
	},
	"use_syslog_facility": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: syslog_facility",
	},
	"use_transfers_in": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: transfers_in",
	},
	"use_transfers_out": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: transfers_out",
	},
	"use_transfers_per_ns": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: transfers_per_ns",
	},
	"use_update_setting": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: allow_update , allow_gss_tsig_zone_updates",
	},
	"use_zone_transfer_format": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: transfer_excluded_servers , transfer_format",
	},
	"views": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of views associated with this member.",
	},
}

func ExpandMemberDns(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberDns {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberDnsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberDnsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberDns {
	if m == nil {
		return nil
	}
	to := &grid.MemberDns{
		Ref:                              flex.ExpandStringPointer(m.Ref),
		AddClientIpMacOptions:            flex.ExpandBoolPointer(m.AddClientIpMacOptions),
		AdditionalIpList:                 flex.ExpandFrameworkListString(ctx, m.AdditionalIpList, diags),
		AdditionalIpListStruct:           flex.ExpandFrameworkListNestedBlock(ctx, m.AdditionalIpListStruct, diags, ExpandMemberDnsAdditionalIpListStruct),
		AllowGssTsigZoneUpdates:          flex.ExpandBoolPointer(m.AllowGssTsigZoneUpdates),
		AllowQuery:                       flex.ExpandFrameworkListNestedBlock(ctx, m.AllowQuery, diags, ExpandMemberDnsAllowQuery),
		AllowRecursiveQuery:              flex.ExpandBoolPointer(m.AllowRecursiveQuery),
		AllowTransfer:                    flex.ExpandFrameworkListNestedBlock(ctx, m.AllowTransfer, diags, ExpandMemberDnsAllowTransfer),
		AllowUpdate:                      flex.ExpandFrameworkListNestedBlock(ctx, m.AllowUpdate, diags, ExpandMemberDnsAllowUpdate),
		AnonymizeResponseLogging:         flex.ExpandBoolPointer(m.AnonymizeResponseLogging),
		AtcFwdEnable:                     flex.ExpandBoolPointer(m.AtcFwdEnable),
		AttackMitigation:                 ExpandMemberDnsAttackMitigation(ctx, m.AttackMitigation, diags),
		AutoBlackhole:                    ExpandMemberDnsAutoBlackhole(ctx, m.AutoBlackhole, diags),
		AutoCreateAAndPtrForLan2:         flex.ExpandBoolPointer(m.AutoCreateAAndPtrForLan2),
		AutoCreateAaaaAndIpv6ptrForLan2:  flex.ExpandBoolPointer(m.AutoCreateAaaaAndIpv6ptrForLan2),
		AutoSortViews:                    flex.ExpandBoolPointer(m.AutoSortViews),
		BindCheckNamesPolicy:             flex.ExpandStringPointer(m.BindCheckNamesPolicy),
		BindHostnameDirective:            flex.ExpandStringPointer(m.BindHostnameDirective),
		BindHostnameDirectiveFqdn:        flex.ExpandStringPointer(m.BindHostnameDirectiveFqdn),
		BlackholeList:                    flex.ExpandFrameworkListNestedBlock(ctx, m.BlackholeList, diags, ExpandMemberDnsBlackholeList),
		BlacklistAction:                  flex.ExpandStringPointer(m.BlacklistAction),
		BlacklistLogQuery:                flex.ExpandBoolPointer(m.BlacklistLogQuery),
		BlacklistRedirectAddresses:       flex.ExpandFrameworkListString(ctx, m.BlacklistRedirectAddresses, diags),
		BlacklistRedirectTtl:             flex.ExpandInt64Pointer(m.BlacklistRedirectTtl),
		BlacklistRulesets:                flex.ExpandFrameworkListString(ctx, m.BlacklistRulesets, diags),
		CaptureDnsQueriesOnAllDomains:    flex.ExpandBoolPointer(m.CaptureDnsQueriesOnAllDomains),
		CheckNamesForDdnsAndZoneTransfer: flex.ExpandBoolPointer(m.CheckNamesForDdnsAndZoneTransfer),
		CopyClientIpMacOptions:           flex.ExpandBoolPointer(m.CopyClientIpMacOptions),
		CopyXferToNotify:                 flex.ExpandBoolPointer(m.CopyXferToNotify),
		CustomRootNameServers:            flex.ExpandFrameworkListNestedBlock(ctx, m.CustomRootNameServers, diags, ExpandMemberDnsCustomRootNameServers),
		DisableEdns:                      flex.ExpandBoolPointer(m.DisableEdns),
		Dns64Groups:                      flex.ExpandFrameworkListString(ctx, m.Dns64Groups, diags),
		DnsCacheAccelerationTtl:          flex.ExpandInt64Pointer(m.DnsCacheAccelerationTtl),
		DnsHealthCheckAnycastControl:     flex.ExpandBoolPointer(m.DnsHealthCheckAnycastControl),
		DnsHealthCheckDomainList:         flex.ExpandFrameworkListString(ctx, m.DnsHealthCheckDomainList, diags),
		DnsHealthCheckInterval:           flex.ExpandInt64Pointer(m.DnsHealthCheckInterval),
		DnsHealthCheckRecursionFlag:      flex.ExpandBoolPointer(m.DnsHealthCheckRecursionFlag),
		DnsHealthCheckRetries:            flex.ExpandInt64Pointer(m.DnsHealthCheckRetries),
		DnsHealthCheckTimeout:            flex.ExpandInt64Pointer(m.DnsHealthCheckTimeout),
		DnsNotifyTransferSource:          flex.ExpandStringPointer(m.DnsNotifyTransferSource),
		DnsNotifyTransferSourceAddress:   flex.ExpandStringPointer(m.DnsNotifyTransferSourceAddress),
		DnsOverTlsService:                flex.ExpandBoolPointer(m.DnsOverTlsService),
		DnsQueryCaptureFileTimeLimit:     flex.ExpandInt64Pointer(m.DnsQueryCaptureFileTimeLimit),
		DnsQuerySourceAddress:            flex.ExpandStringPointer(m.DnsQuerySourceAddress),
		DnsQuerySourceInterface:          flex.ExpandStringPointer(m.DnsQuerySourceInterface),
		DnsViewAddressSettings:           flex.ExpandFrameworkListNestedBlock(ctx, m.DnsViewAddressSettings, diags, ExpandMemberDnsDnsViewAddressSettings),
		DnssecBlacklistEnabled:           flex.ExpandBoolPointer(m.DnssecBlacklistEnabled),
		DnssecDns64Enabled:               flex.ExpandBoolPointer(m.DnssecDns64Enabled),
		DnssecEnabled:                    flex.ExpandBoolPointer(m.DnssecEnabled),
		DnssecExpiredSignaturesEnabled:   flex.ExpandBoolPointer(m.DnssecExpiredSignaturesEnabled),
		DnssecNegativeTrustAnchors:       flex.ExpandFrameworkListString(ctx, m.DnssecNegativeTrustAnchors, diags),
		DnssecNxdomainEnabled:            flex.ExpandBoolPointer(m.DnssecNxdomainEnabled),
		DnssecRpzEnabled:                 flex.ExpandBoolPointer(m.DnssecRpzEnabled),
		DnssecTrustedKeys:                flex.ExpandFrameworkListNestedBlock(ctx, m.DnssecTrustedKeys, diags, ExpandMemberDnsDnssecTrustedKeys),
		DnssecValidationEnabled:          flex.ExpandBoolPointer(m.DnssecValidationEnabled),
		DnstapSetting:                    ExpandMemberDnsDnstapSetting(ctx, m.DnstapSetting, diags),
		DohHttpsSessionDuration:          flex.ExpandInt64Pointer(m.DohHttpsSessionDuration),
		DohService:                       flex.ExpandBoolPointer(m.DohService),
		DomainsToCaptureDnsQueries:       flex.ExpandFrameworkListString(ctx, m.DomainsToCaptureDnsQueries, diags),
		DtcDnsQueriesSpecificBehavior:    flex.ExpandStringPointer(m.DtcDnsQueriesSpecificBehavior),
		DtcEdnsPreferClientSubnet:        flex.ExpandBoolPointer(m.DtcEdnsPreferClientSubnet),
		DtcHealthSource:                  flex.ExpandStringPointer(m.DtcHealthSource),
		DtcHealthSourceAddress:           flex.ExpandStringPointer(m.DtcHealthSourceAddress),
		EdnsUdpSize:                      flex.ExpandInt64Pointer(m.EdnsUdpSize),
		EnableBlackhole:                  flex.ExpandBoolPointer(m.EnableBlackhole),
		EnableBlacklist:                  flex.ExpandBoolPointer(m.EnableBlacklist),
		EnableCaptureDnsQueries:          flex.ExpandBoolPointer(m.EnableCaptureDnsQueries),
		EnableCaptureDnsResponses:        flex.ExpandBoolPointer(m.EnableCaptureDnsResponses),
		EnableDns:                        flex.ExpandBoolPointer(m.EnableDns),
		EnableDns64:                      flex.ExpandBoolPointer(m.EnableDns64),
		EnableDnsCacheAcceleration:       flex.ExpandBoolPointer(m.EnableDnsCacheAcceleration),
		EnableDnsHealthCheck:             flex.ExpandBoolPointer(m.EnableDnsHealthCheck),
		EnableDnstapQueries:              flex.ExpandBoolPointer(m.EnableDnstapQueries),
		EnableDnstapResponses:            flex.ExpandBoolPointer(m.EnableDnstapResponses),
		EnableDnstapViolationsTls:        flex.ExpandBoolPointer(m.EnableDnstapViolationsTls),
		EnableExcludedDomainNames:        flex.ExpandBoolPointer(m.EnableExcludedDomainNames),
		EnableFixedRrsetOrderFqdns:       flex.ExpandBoolPointer(m.EnableFixedRrsetOrderFqdns),
		EnableFtc:                        flex.ExpandBoolPointer(m.EnableFtc),
		EnableGssTsig:                    flex.ExpandBoolPointer(m.EnableGssTsig),
		EnableNotifySourcePort:           flex.ExpandBoolPointer(m.EnableNotifySourcePort),
		EnableQueryRewrite:               flex.ExpandBoolPointer(m.EnableQueryRewrite),
		EnableQuerySourcePort:            flex.ExpandBoolPointer(m.EnableQuerySourcePort),
		ExcludedDomainNames:              flex.ExpandFrameworkListString(ctx, m.ExcludedDomainNames, diags),
		ExtAttrs:                         ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		FileTransferSetting:              ExpandMemberDnsFileTransferSetting(ctx, m.FileTransferSetting, diags),
		FilterAaaa:                       flex.ExpandStringPointer(m.FilterAaaa),
		FilterAaaaList:                   flex.ExpandFrameworkListNestedBlock(ctx, m.FilterAaaaList, diags, ExpandMemberDnsFilterAaaaList),
		FixedRrsetOrderFqdns:             flex.ExpandFrameworkListNestedBlock(ctx, m.FixedRrsetOrderFqdns, diags, ExpandMemberDnsFixedRrsetOrderFqdns),
		ForwardOnly:                      flex.ExpandBoolPointer(m.ForwardOnly),
		ForwardUpdates:                   flex.ExpandBoolPointer(m.ForwardUpdates),
		Forwarders:                       flex.ExpandFrameworkListString(ctx, m.Forwarders, diags),
		FtcExpiredRecordTimeout:          flex.ExpandInt64Pointer(m.FtcExpiredRecordTimeout),
		FtcExpiredRecordTtl:              flex.ExpandInt64Pointer(m.FtcExpiredRecordTtl),
		GlueRecordAddresses:              flex.ExpandFrameworkListNestedBlock(ctx, m.GlueRecordAddresses, diags, ExpandMemberDnsGlueRecordAddresses),
		GssTsigKeys:                      flex.ExpandFrameworkListString(ctx, m.GssTsigKeys, diags),
		Ipv6GlueRecordAddresses:          flex.ExpandFrameworkListNestedBlock(ctx, m.Ipv6GlueRecordAddresses, diags, ExpandMemberDnsIpv6GlueRecordAddresses),
		LoggingCategories:                ExpandMemberDnsLoggingCategories(ctx, m.LoggingCategories, diags),
		MaxCacheTtl:                      flex.ExpandInt64Pointer(m.MaxCacheTtl),
		MaxCachedLifetime:                flex.ExpandInt64Pointer(m.MaxCachedLifetime),
		MaxNcacheTtl:                     flex.ExpandInt64Pointer(m.MaxNcacheTtl),
		MaxUdpSize:                       flex.ExpandInt64Pointer(m.MaxUdpSize),
		MinimalResp:                      flex.ExpandBoolPointer(m.MinimalResp),
		NotifyDelay:                      flex.ExpandInt64Pointer(m.NotifyDelay),
		NotifySourcePort:                 flex.ExpandInt64Pointer(m.NotifySourcePort),
		NxdomainLogQuery:                 flex.ExpandBoolPointer(m.NxdomainLogQuery),
		NxdomainRedirect:                 flex.ExpandBoolPointer(m.NxdomainRedirect),
		NxdomainRedirectAddresses:        flex.ExpandFrameworkListString(ctx, m.NxdomainRedirectAddresses, diags),
		NxdomainRedirectAddressesV6:      flex.ExpandFrameworkListString(ctx, m.NxdomainRedirectAddressesV6, diags),
		NxdomainRedirectTtl:              flex.ExpandInt64Pointer(m.NxdomainRedirectTtl),
		NxdomainRulesets:                 flex.ExpandFrameworkListString(ctx, m.NxdomainRulesets, diags),
		QuerySourcePort:                  flex.ExpandInt64Pointer(m.QuerySourcePort),
		RecordNamePolicy:                 flex.ExpandStringPointer(m.RecordNamePolicy),
		RecursiveClientLimit:             flex.ExpandInt64Pointer(m.RecursiveClientLimit),
		RecursiveQueryList:               flex.ExpandFrameworkListNestedBlock(ctx, m.RecursiveQueryList, diags, ExpandMemberDnsRecursiveQueryList),
		RecursiveResolver:                flex.ExpandStringPointer(m.RecursiveResolver),
		ResolverQueryTimeout:             flex.ExpandInt64Pointer(m.ResolverQueryTimeout),
		ResponseRateLimiting:             ExpandMemberDnsResponseRateLimiting(ctx, m.ResponseRateLimiting, diags),
		RootNameServerType:               flex.ExpandStringPointer(m.RootNameServerType),
		RpzDisableNsdnameNsip:            flex.ExpandBoolPointer(m.RpzDisableNsdnameNsip),
		RpzDropIpRuleEnabled:             flex.ExpandBoolPointer(m.RpzDropIpRuleEnabled),
		RpzDropIpRuleMinPrefixLengthIpv4: flex.ExpandInt64Pointer(m.RpzDropIpRuleMinPrefixLengthIpv4),
		RpzDropIpRuleMinPrefixLengthIpv6: flex.ExpandInt64Pointer(m.RpzDropIpRuleMinPrefixLengthIpv6),
		RpzQnameWaitRecurse:              flex.ExpandBoolPointer(m.RpzQnameWaitRecurse),
		SerialQueryRate:                  flex.ExpandInt64Pointer(m.SerialQueryRate),
		ServerIdDirective:                flex.ExpandStringPointer(m.ServerIdDirective),
		ServerIdDirectiveString:          flex.ExpandStringPointer(m.ServerIdDirectiveString),
		SkipInGridRpzQueries:             flex.ExpandBoolPointer(m.SkipInGridRpzQueries),
		Sortlist:                         flex.ExpandFrameworkListNestedBlock(ctx, m.Sortlist, diags, ExpandMemberDnsSortlist),
		StoreLocally:                     flex.ExpandBoolPointer(m.StoreLocally),
		SyslogFacility:                   flex.ExpandStringPointer(m.SyslogFacility),
		TcpIdleTimeout:                   flex.ExpandInt64Pointer(m.TcpIdleTimeout),
		TlsSessionDuration:               flex.ExpandInt64Pointer(m.TlsSessionDuration),
		TransferExcludedServers:          flex.ExpandFrameworkListString(ctx, m.TransferExcludedServers, diags),
		TransferFormat:                   flex.ExpandStringPointer(m.TransferFormat),
		TransfersIn:                      flex.ExpandInt64Pointer(m.TransfersIn),
		TransfersOut:                     flex.ExpandInt64Pointer(m.TransfersOut),
		TransfersPerNs:                   flex.ExpandInt64Pointer(m.TransfersPerNs),
		UpstreamAddressFamilyPreference:  flex.ExpandStringPointer(m.UpstreamAddressFamilyPreference),
		UseAddClientIpMacOptions:         flex.ExpandBoolPointer(m.UseAddClientIpMacOptions),
		UseAllowQuery:                    flex.ExpandBoolPointer(m.UseAllowQuery),
		UseAllowTransfer:                 flex.ExpandBoolPointer(m.UseAllowTransfer),
		UseAttackMitigation:              flex.ExpandBoolPointer(m.UseAttackMitigation),
		UseAutoBlackhole:                 flex.ExpandBoolPointer(m.UseAutoBlackhole),
		UseBindHostnameDirective:         flex.ExpandBoolPointer(m.UseBindHostnameDirective),
		UseBlackhole:                     flex.ExpandBoolPointer(m.UseBlackhole),
		UseBlacklist:                     flex.ExpandBoolPointer(m.UseBlacklist),
		UseCaptureDnsQueriesOnAllDomains: flex.ExpandBoolPointer(m.UseCaptureDnsQueriesOnAllDomains),
		UseCopyClientIpMacOptions:        flex.ExpandBoolPointer(m.UseCopyClientIpMacOptions),
		UseCopyXferToNotify:              flex.ExpandBoolPointer(m.UseCopyXferToNotify),
		UseDisableEdns:                   flex.ExpandBoolPointer(m.UseDisableEdns),
		UseDns64:                         flex.ExpandBoolPointer(m.UseDns64),
		UseDnsCacheAccelerationTtl:       flex.ExpandBoolPointer(m.UseDnsCacheAccelerationTtl),
		UseDnsHealthCheck:                flex.ExpandBoolPointer(m.UseDnsHealthCheck),
		UseDnssec:                        flex.ExpandBoolPointer(m.UseDnssec),
		UseDnstapSetting:                 flex.ExpandBoolPointer(m.UseDnstapSetting),
		UseDtcDnsQueriesSpecificBehavior: flex.ExpandBoolPointer(m.UseDtcDnsQueriesSpecificBehavior),
		UseDtcEdnsPreferClientSubnet:     flex.ExpandBoolPointer(m.UseDtcEdnsPreferClientSubnet),
		UseEdnsUdpSize:                   flex.ExpandBoolPointer(m.UseEdnsUdpSize),
		UseEnableCaptureDns:              flex.ExpandBoolPointer(m.UseEnableCaptureDns),
		UseEnableExcludedDomainNames:     flex.ExpandBoolPointer(m.UseEnableExcludedDomainNames),
		UseEnableGssTsig:                 flex.ExpandBoolPointer(m.UseEnableGssTsig),
		UseEnableQueryRewrite:            flex.ExpandBoolPointer(m.UseEnableQueryRewrite),
		UseFilterAaaa:                    flex.ExpandBoolPointer(m.UseFilterAaaa),
		UseFixedRrsetOrderFqdns:          flex.ExpandBoolPointer(m.UseFixedRrsetOrderFqdns),
		UseForwardUpdates:                flex.ExpandBoolPointer(m.UseForwardUpdates),
		UseForwarders:                    flex.ExpandBoolPointer(m.UseForwarders),
		UseFtc:                           flex.ExpandBoolPointer(m.UseFtc),
		UseGssTsigKeys:                   flex.ExpandBoolPointer(m.UseGssTsigKeys),
		UseLan2Ipv6Port:                  flex.ExpandBoolPointer(m.UseLan2Ipv6Port),
		UseLan2Port:                      flex.ExpandBoolPointer(m.UseLan2Port),
		UseLanIpv6Port:                   flex.ExpandBoolPointer(m.UseLanIpv6Port),
		UseLanPort:                       flex.ExpandBoolPointer(m.UseLanPort),
		UseLoggingCategories:             flex.ExpandBoolPointer(m.UseLoggingCategories),
		UseMaxCacheTtl:                   flex.ExpandBoolPointer(m.UseMaxCacheTtl),
		UseMaxCachedLifetime:             flex.ExpandBoolPointer(m.UseMaxCachedLifetime),
		UseMaxNcacheTtl:                  flex.ExpandBoolPointer(m.UseMaxNcacheTtl),
		UseMaxUdpSize:                    flex.ExpandBoolPointer(m.UseMaxUdpSize),
		UseMgmtIpv6Port:                  flex.ExpandBoolPointer(m.UseMgmtIpv6Port),
		UseMgmtPort:                      flex.ExpandBoolPointer(m.UseMgmtPort),
		UseNotifyDelay:                   flex.ExpandBoolPointer(m.UseNotifyDelay),
		UseNxdomainRedirect:              flex.ExpandBoolPointer(m.UseNxdomainRedirect),
		UseRecordNamePolicy:              flex.ExpandBoolPointer(m.UseRecordNamePolicy),
		UseRecursiveClientLimit:          flex.ExpandBoolPointer(m.UseRecursiveClientLimit),
		UseRecursiveQuerySetting:         flex.ExpandBoolPointer(m.UseRecursiveQuerySetting),
		UseResolverQueryTimeout:          flex.ExpandBoolPointer(m.UseResolverQueryTimeout),
		UseResponseRateLimiting:          flex.ExpandBoolPointer(m.UseResponseRateLimiting),
		UseRootNameServer:                flex.ExpandBoolPointer(m.UseRootNameServer),
		UseRootServerForAllViews:         flex.ExpandBoolPointer(m.UseRootServerForAllViews),
		UseRpzDisableNsdnameNsip:         flex.ExpandBoolPointer(m.UseRpzDisableNsdnameNsip),
		UseRpzDropIpRule:                 flex.ExpandBoolPointer(m.UseRpzDropIpRule),
		UseRpzQnameWaitRecurse:           flex.ExpandBoolPointer(m.UseRpzQnameWaitRecurse),
		UseSerialQueryRate:               flex.ExpandBoolPointer(m.UseSerialQueryRate),
		UseServerIdDirective:             flex.ExpandBoolPointer(m.UseServerIdDirective),
		UseSortlist:                      flex.ExpandBoolPointer(m.UseSortlist),
		UseSourcePorts:                   flex.ExpandBoolPointer(m.UseSourcePorts),
		UseSyslogFacility:                flex.ExpandBoolPointer(m.UseSyslogFacility),
		UseTransfersIn:                   flex.ExpandBoolPointer(m.UseTransfersIn),
		UseTransfersOut:                  flex.ExpandBoolPointer(m.UseTransfersOut),
		UseTransfersPerNs:                flex.ExpandBoolPointer(m.UseTransfersPerNs),
		UseUpdateSetting:                 flex.ExpandBoolPointer(m.UseUpdateSetting),
		UseZoneTransferFormat:            flex.ExpandBoolPointer(m.UseZoneTransferFormat),
		Views:                            flex.ExpandFrameworkListString(ctx, m.Views, diags),
	}
	return to
}

func FlattenMemberDns(ctx context.Context, from *grid.MemberDns, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberDnsAttrTypes)
	}
	m := MemberDnsModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, MemberDnsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberDnsModel) Flatten(ctx context.Context, from *grid.MemberDns, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberDnsModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.AddClientIpMacOptions = types.BoolPointerValue(from.AddClientIpMacOptions)
	m.AdditionalIpList = flex.FlattenFrameworkListString(ctx, from.AdditionalIpList, diags)
	m.AdditionalIpListStruct = flex.FlattenFrameworkListNestedBlock(ctx, from.AdditionalIpListStruct, MemberDnsAdditionalIpListStructAttrTypes, diags, FlattenMemberDnsAdditionalIpListStruct)
	m.AllowGssTsigZoneUpdates = types.BoolPointerValue(from.AllowGssTsigZoneUpdates)
	m.AllowQuery = flex.FlattenFrameworkListNestedBlock(ctx, from.AllowQuery, MemberDnsAllowQueryAttrTypes, diags, FlattenMemberDnsAllowQuery)
	m.AllowRecursiveQuery = types.BoolPointerValue(from.AllowRecursiveQuery)
	m.AllowTransfer = flex.FlattenFrameworkListNestedBlock(ctx, from.AllowTransfer, MemberDnsAllowTransferAttrTypes, diags, FlattenMemberDnsAllowTransfer)
	m.AllowUpdate = flex.FlattenFrameworkListNestedBlock(ctx, from.AllowUpdate, MemberDnsAllowUpdateAttrTypes, diags, FlattenMemberDnsAllowUpdate)
	m.AnonymizeResponseLogging = types.BoolPointerValue(from.AnonymizeResponseLogging)
	m.AtcFwdEnable = types.BoolPointerValue(from.AtcFwdEnable)
	m.AttackMitigation = FlattenMemberDnsAttackMitigation(ctx, from.AttackMitigation, diags)
	m.AutoBlackhole = FlattenMemberDnsAutoBlackhole(ctx, from.AutoBlackhole, diags)
	m.AutoCreateAAndPtrForLan2 = types.BoolPointerValue(from.AutoCreateAAndPtrForLan2)
	m.AutoCreateAaaaAndIpv6ptrForLan2 = types.BoolPointerValue(from.AutoCreateAaaaAndIpv6ptrForLan2)
	m.AutoSortViews = types.BoolPointerValue(from.AutoSortViews)
	m.BindCheckNamesPolicy = flex.FlattenStringPointer(from.BindCheckNamesPolicy)
	m.BindHostnameDirective = flex.FlattenStringPointer(from.BindHostnameDirective)
	m.BindHostnameDirectiveFqdn = flex.FlattenStringPointer(from.BindHostnameDirectiveFqdn)
	m.BlackholeList = flex.FlattenFrameworkListNestedBlock(ctx, from.BlackholeList, MemberDnsBlackholeListAttrTypes, diags, FlattenMemberDnsBlackholeList)
	m.BlacklistAction = flex.FlattenStringPointer(from.BlacklistAction)
	m.BlacklistLogQuery = types.BoolPointerValue(from.BlacklistLogQuery)
	m.BlacklistRedirectAddresses = flex.FlattenFrameworkListString(ctx, from.BlacklistRedirectAddresses, diags)
	m.BlacklistRedirectTtl = flex.FlattenInt64Pointer(from.BlacklistRedirectTtl)
	m.BlacklistRulesets = flex.FlattenFrameworkListString(ctx, from.BlacklistRulesets, diags)
	m.CaptureDnsQueriesOnAllDomains = types.BoolPointerValue(from.CaptureDnsQueriesOnAllDomains)
	m.CheckNamesForDdnsAndZoneTransfer = types.BoolPointerValue(from.CheckNamesForDdnsAndZoneTransfer)
	m.CopyClientIpMacOptions = types.BoolPointerValue(from.CopyClientIpMacOptions)
	m.CopyXferToNotify = types.BoolPointerValue(from.CopyXferToNotify)
	m.CustomRootNameServers = flex.FlattenFrameworkListNestedBlock(ctx, from.CustomRootNameServers, MemberDnsCustomRootNameServersAttrTypes, diags, FlattenMemberDnsCustomRootNameServers)
	m.DisableEdns = types.BoolPointerValue(from.DisableEdns)
	m.Dns64Groups = flex.FlattenFrameworkListString(ctx, from.Dns64Groups, diags)
	m.DnsCacheAccelerationStatus = flex.FlattenStringPointer(from.DnsCacheAccelerationStatus)
	m.DnsCacheAccelerationTtl = flex.FlattenInt64Pointer(from.DnsCacheAccelerationTtl)
	m.DnsHealthCheckAnycastControl = types.BoolPointerValue(from.DnsHealthCheckAnycastControl)
	m.DnsHealthCheckDomainList = flex.FlattenFrameworkListString(ctx, from.DnsHealthCheckDomainList, diags)
	m.DnsHealthCheckInterval = flex.FlattenInt64Pointer(from.DnsHealthCheckInterval)
	m.DnsHealthCheckRecursionFlag = types.BoolPointerValue(from.DnsHealthCheckRecursionFlag)
	m.DnsHealthCheckRetries = flex.FlattenInt64Pointer(from.DnsHealthCheckRetries)
	m.DnsHealthCheckTimeout = flex.FlattenInt64Pointer(from.DnsHealthCheckTimeout)
	m.DnsNotifyTransferSource = flex.FlattenStringPointer(from.DnsNotifyTransferSource)
	m.DnsNotifyTransferSourceAddress = flex.FlattenStringPointer(from.DnsNotifyTransferSourceAddress)
	m.DnsOverTlsService = types.BoolPointerValue(from.DnsOverTlsService)
	m.DnsQueryCaptureFileTimeLimit = flex.FlattenInt64Pointer(from.DnsQueryCaptureFileTimeLimit)
	m.DnsQuerySourceAddress = flex.FlattenStringPointer(from.DnsQuerySourceAddress)
	m.DnsQuerySourceInterface = flex.FlattenStringPointer(from.DnsQuerySourceInterface)
	m.DnsViewAddressSettings = flex.FlattenFrameworkListNestedBlock(ctx, from.DnsViewAddressSettings, MemberDnsDnsViewAddressSettingsAttrTypes, diags, FlattenMemberDnsDnsViewAddressSettings)
	m.DnssecBlacklistEnabled = types.BoolPointerValue(from.DnssecBlacklistEnabled)
	m.DnssecDns64Enabled = types.BoolPointerValue(from.DnssecDns64Enabled)
	m.DnssecEnabled = types.BoolPointerValue(from.DnssecEnabled)
	m.DnssecExpiredSignaturesEnabled = types.BoolPointerValue(from.DnssecExpiredSignaturesEnabled)
	m.DnssecNegativeTrustAnchors = flex.FlattenFrameworkListString(ctx, from.DnssecNegativeTrustAnchors, diags)
	m.DnssecNxdomainEnabled = types.BoolPointerValue(from.DnssecNxdomainEnabled)
	m.DnssecRpzEnabled = types.BoolPointerValue(from.DnssecRpzEnabled)
	m.DnssecTrustedKeys = flex.FlattenFrameworkListNestedBlock(ctx, from.DnssecTrustedKeys, MemberDnsDnssecTrustedKeysAttrTypes, diags, FlattenMemberDnsDnssecTrustedKeys)
	m.DnssecValidationEnabled = types.BoolPointerValue(from.DnssecValidationEnabled)
	m.DnstapSetting = FlattenMemberDnsDnstapSetting(ctx, from.DnstapSetting, diags)
	m.DohHttpsSessionDuration = flex.FlattenInt64Pointer(from.DohHttpsSessionDuration)
	m.DohService = types.BoolPointerValue(from.DohService)
	m.DomainsToCaptureDnsQueries = flex.FlattenFrameworkListString(ctx, from.DomainsToCaptureDnsQueries, diags)
	m.DtcDnsQueriesSpecificBehavior = flex.FlattenStringPointer(from.DtcDnsQueriesSpecificBehavior)
	m.DtcEdnsPreferClientSubnet = types.BoolPointerValue(from.DtcEdnsPreferClientSubnet)
	m.DtcHealthSource = flex.FlattenStringPointer(from.DtcHealthSource)
	m.DtcHealthSourceAddress = flex.FlattenStringPointer(from.DtcHealthSourceAddress)
	m.EdnsUdpSize = flex.FlattenInt64Pointer(from.EdnsUdpSize)
	m.EnableBlackhole = types.BoolPointerValue(from.EnableBlackhole)
	m.EnableBlacklist = types.BoolPointerValue(from.EnableBlacklist)
	m.EnableCaptureDnsQueries = types.BoolPointerValue(from.EnableCaptureDnsQueries)
	m.EnableCaptureDnsResponses = types.BoolPointerValue(from.EnableCaptureDnsResponses)
	m.EnableDns = types.BoolPointerValue(from.EnableDns)
	m.EnableDns64 = types.BoolPointerValue(from.EnableDns64)
	m.EnableDnsCacheAcceleration = types.BoolPointerValue(from.EnableDnsCacheAcceleration)
	m.EnableDnsHealthCheck = types.BoolPointerValue(from.EnableDnsHealthCheck)
	m.EnableDnstapQueries = types.BoolPointerValue(from.EnableDnstapQueries)
	m.EnableDnstapResponses = types.BoolPointerValue(from.EnableDnstapResponses)
	m.EnableDnstapViolationsTls = types.BoolPointerValue(from.EnableDnstapViolationsTls)
	m.EnableExcludedDomainNames = types.BoolPointerValue(from.EnableExcludedDomainNames)
	m.EnableFixedRrsetOrderFqdns = types.BoolPointerValue(from.EnableFixedRrsetOrderFqdns)
	m.EnableFtc = types.BoolPointerValue(from.EnableFtc)
	m.EnableGssTsig = types.BoolPointerValue(from.EnableGssTsig)
	m.EnableNotifySourcePort = types.BoolPointerValue(from.EnableNotifySourcePort)
	m.EnableQueryRewrite = types.BoolPointerValue(from.EnableQueryRewrite)
	m.EnableQuerySourcePort = types.BoolPointerValue(from.EnableQuerySourcePort)
	m.ExcludedDomainNames = flex.FlattenFrameworkListString(ctx, from.ExcludedDomainNames, diags)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.FileTransferSetting = FlattenMemberDnsFileTransferSetting(ctx, from.FileTransferSetting, diags)
	m.FilterAaaa = flex.FlattenStringPointer(from.FilterAaaa)
	m.FilterAaaaList = flex.FlattenFrameworkListNestedBlock(ctx, from.FilterAaaaList, MemberDnsFilterAaaaListAttrTypes, diags, FlattenMemberDnsFilterAaaaList)
	m.FixedRrsetOrderFqdns = flex.FlattenFrameworkListNestedBlock(ctx, from.FixedRrsetOrderFqdns, MemberDnsFixedRrsetOrderFqdnsAttrTypes, diags, FlattenMemberDnsFixedRrsetOrderFqdns)
	m.ForwardOnly = types.BoolPointerValue(from.ForwardOnly)
	m.ForwardUpdates = types.BoolPointerValue(from.ForwardUpdates)
	m.Forwarders = flex.FlattenFrameworkListString(ctx, from.Forwarders, diags)
	m.FtcExpiredRecordTimeout = flex.FlattenInt64Pointer(from.FtcExpiredRecordTimeout)
	m.FtcExpiredRecordTtl = flex.FlattenInt64Pointer(from.FtcExpiredRecordTtl)
	m.GlueRecordAddresses = flex.FlattenFrameworkListNestedBlock(ctx, from.GlueRecordAddresses, MemberDnsGlueRecordAddressesAttrTypes, diags, FlattenMemberDnsGlueRecordAddresses)
	m.GssTsigKeys = flex.FlattenFrameworkListString(ctx, from.GssTsigKeys, diags)
	m.HostName = flex.FlattenStringPointer(from.HostName)
	m.Ipv4addr = flex.FlattenStringPointer(from.Ipv4addr)
	m.Ipv6GlueRecordAddresses = flex.FlattenFrameworkListNestedBlock(ctx, from.Ipv6GlueRecordAddresses, MemberDnsIpv6GlueRecordAddressesAttrTypes, diags, FlattenMemberDnsIpv6GlueRecordAddresses)
	m.Ipv6addr = flex.FlattenStringPointer(from.Ipv6addr)
	m.LoggingCategories = FlattenMemberDnsLoggingCategories(ctx, from.LoggingCategories, diags)
	m.MaxCacheTtl = flex.FlattenInt64Pointer(from.MaxCacheTtl)
	m.MaxCachedLifetime = flex.FlattenInt64Pointer(from.MaxCachedLifetime)
	m.MaxNcacheTtl = flex.FlattenInt64Pointer(from.MaxNcacheTtl)
	m.MaxUdpSize = flex.FlattenInt64Pointer(from.MaxUdpSize)
	m.MinimalResp = types.BoolPointerValue(from.MinimalResp)
	m.NotifyDelay = flex.FlattenInt64Pointer(from.NotifyDelay)
	m.NotifySourcePort = flex.FlattenInt64Pointer(from.NotifySourcePort)
	m.NxdomainLogQuery = types.BoolPointerValue(from.NxdomainLogQuery)
	m.NxdomainRedirect = types.BoolPointerValue(from.NxdomainRedirect)
	m.NxdomainRedirectAddresses = flex.FlattenFrameworkListString(ctx, from.NxdomainRedirectAddresses, diags)
	m.NxdomainRedirectAddressesV6 = flex.FlattenFrameworkListString(ctx, from.NxdomainRedirectAddressesV6, diags)
	m.NxdomainRedirectTtl = flex.FlattenInt64Pointer(from.NxdomainRedirectTtl)
	m.NxdomainRulesets = flex.FlattenFrameworkListString(ctx, from.NxdomainRulesets, diags)
	m.QuerySourcePort = flex.FlattenInt64Pointer(from.QuerySourcePort)
	m.RecordNamePolicy = flex.FlattenStringPointer(from.RecordNamePolicy)
	m.RecursiveClientLimit = flex.FlattenInt64Pointer(from.RecursiveClientLimit)
	m.RecursiveQueryList = flex.FlattenFrameworkListNestedBlock(ctx, from.RecursiveQueryList, MemberDnsRecursiveQueryListAttrTypes, diags, FlattenMemberDnsRecursiveQueryList)
	m.RecursiveResolver = flex.FlattenStringPointer(from.RecursiveResolver)
	m.ResolverQueryTimeout = flex.FlattenInt64Pointer(from.ResolverQueryTimeout)
	m.ResponseRateLimiting = FlattenMemberDnsResponseRateLimiting(ctx, from.ResponseRateLimiting, diags)
	m.RootNameServerType = flex.FlattenStringPointer(from.RootNameServerType)
	m.RpzDisableNsdnameNsip = types.BoolPointerValue(from.RpzDisableNsdnameNsip)
	m.RpzDropIpRuleEnabled = types.BoolPointerValue(from.RpzDropIpRuleEnabled)
	m.RpzDropIpRuleMinPrefixLengthIpv4 = flex.FlattenInt64Pointer(from.RpzDropIpRuleMinPrefixLengthIpv4)
	m.RpzDropIpRuleMinPrefixLengthIpv6 = flex.FlattenInt64Pointer(from.RpzDropIpRuleMinPrefixLengthIpv6)
	m.RpzQnameWaitRecurse = types.BoolPointerValue(from.RpzQnameWaitRecurse)
	m.SerialQueryRate = flex.FlattenInt64Pointer(from.SerialQueryRate)
	m.ServerIdDirective = flex.FlattenStringPointer(from.ServerIdDirective)
	m.ServerIdDirectiveString = flex.FlattenStringPointer(from.ServerIdDirectiveString)
	m.SkipInGridRpzQueries = types.BoolPointerValue(from.SkipInGridRpzQueries)
	m.Sortlist = flex.FlattenFrameworkListNestedBlock(ctx, from.Sortlist, MemberDnsSortlistAttrTypes, diags, FlattenMemberDnsSortlist)
	m.StoreLocally = types.BoolPointerValue(from.StoreLocally)
	m.SyslogFacility = flex.FlattenStringPointer(from.SyslogFacility)
	m.TcpIdleTimeout = flex.FlattenInt64Pointer(from.TcpIdleTimeout)
	m.TlsSessionDuration = flex.FlattenInt64Pointer(from.TlsSessionDuration)
	m.TransferExcludedServers = flex.FlattenFrameworkListString(ctx, from.TransferExcludedServers, diags)
	m.TransferFormat = flex.FlattenStringPointer(from.TransferFormat)
	m.TransfersIn = flex.FlattenInt64Pointer(from.TransfersIn)
	m.TransfersOut = flex.FlattenInt64Pointer(from.TransfersOut)
	m.TransfersPerNs = flex.FlattenInt64Pointer(from.TransfersPerNs)
	m.UpstreamAddressFamilyPreference = flex.FlattenStringPointer(from.UpstreamAddressFamilyPreference)
	m.UseAddClientIpMacOptions = types.BoolPointerValue(from.UseAddClientIpMacOptions)
	m.UseAllowQuery = types.BoolPointerValue(from.UseAllowQuery)
	m.UseAllowTransfer = types.BoolPointerValue(from.UseAllowTransfer)
	m.UseAttackMitigation = types.BoolPointerValue(from.UseAttackMitigation)
	m.UseAutoBlackhole = types.BoolPointerValue(from.UseAutoBlackhole)
	m.UseBindHostnameDirective = types.BoolPointerValue(from.UseBindHostnameDirective)
	m.UseBlackhole = types.BoolPointerValue(from.UseBlackhole)
	m.UseBlacklist = types.BoolPointerValue(from.UseBlacklist)
	m.UseCaptureDnsQueriesOnAllDomains = types.BoolPointerValue(from.UseCaptureDnsQueriesOnAllDomains)
	m.UseCopyClientIpMacOptions = types.BoolPointerValue(from.UseCopyClientIpMacOptions)
	m.UseCopyXferToNotify = types.BoolPointerValue(from.UseCopyXferToNotify)
	m.UseDisableEdns = types.BoolPointerValue(from.UseDisableEdns)
	m.UseDns64 = types.BoolPointerValue(from.UseDns64)
	m.UseDnsCacheAccelerationTtl = types.BoolPointerValue(from.UseDnsCacheAccelerationTtl)
	m.UseDnsHealthCheck = types.BoolPointerValue(from.UseDnsHealthCheck)
	m.UseDnssec = types.BoolPointerValue(from.UseDnssec)
	m.UseDnstapSetting = types.BoolPointerValue(from.UseDnstapSetting)
	m.UseDtcDnsQueriesSpecificBehavior = types.BoolPointerValue(from.UseDtcDnsQueriesSpecificBehavior)
	m.UseDtcEdnsPreferClientSubnet = types.BoolPointerValue(from.UseDtcEdnsPreferClientSubnet)
	m.UseEdnsUdpSize = types.BoolPointerValue(from.UseEdnsUdpSize)
	m.UseEnableCaptureDns = types.BoolPointerValue(from.UseEnableCaptureDns)
	m.UseEnableExcludedDomainNames = types.BoolPointerValue(from.UseEnableExcludedDomainNames)
	m.UseEnableGssTsig = types.BoolPointerValue(from.UseEnableGssTsig)
	m.UseEnableQueryRewrite = types.BoolPointerValue(from.UseEnableQueryRewrite)
	m.UseFilterAaaa = types.BoolPointerValue(from.UseFilterAaaa)
	m.UseFixedRrsetOrderFqdns = types.BoolPointerValue(from.UseFixedRrsetOrderFqdns)
	m.UseForwardUpdates = types.BoolPointerValue(from.UseForwardUpdates)
	m.UseForwarders = types.BoolPointerValue(from.UseForwarders)
	m.UseFtc = types.BoolPointerValue(from.UseFtc)
	m.UseGssTsigKeys = types.BoolPointerValue(from.UseGssTsigKeys)
	m.UseLan2Ipv6Port = types.BoolPointerValue(from.UseLan2Ipv6Port)
	m.UseLan2Port = types.BoolPointerValue(from.UseLan2Port)
	m.UseLanIpv6Port = types.BoolPointerValue(from.UseLanIpv6Port)
	m.UseLanPort = types.BoolPointerValue(from.UseLanPort)
	m.UseLoggingCategories = types.BoolPointerValue(from.UseLoggingCategories)
	m.UseMaxCacheTtl = types.BoolPointerValue(from.UseMaxCacheTtl)
	m.UseMaxCachedLifetime = types.BoolPointerValue(from.UseMaxCachedLifetime)
	m.UseMaxNcacheTtl = types.BoolPointerValue(from.UseMaxNcacheTtl)
	m.UseMaxUdpSize = types.BoolPointerValue(from.UseMaxUdpSize)
	m.UseMgmtIpv6Port = types.BoolPointerValue(from.UseMgmtIpv6Port)
	m.UseMgmtPort = types.BoolPointerValue(from.UseMgmtPort)
	m.UseNotifyDelay = types.BoolPointerValue(from.UseNotifyDelay)
	m.UseNxdomainRedirect = types.BoolPointerValue(from.UseNxdomainRedirect)
	m.UseRecordNamePolicy = types.BoolPointerValue(from.UseRecordNamePolicy)
	m.UseRecursiveClientLimit = types.BoolPointerValue(from.UseRecursiveClientLimit)
	m.UseRecursiveQuerySetting = types.BoolPointerValue(from.UseRecursiveQuerySetting)
	m.UseResolverQueryTimeout = types.BoolPointerValue(from.UseResolverQueryTimeout)
	m.UseResponseRateLimiting = types.BoolPointerValue(from.UseResponseRateLimiting)
	m.UseRootNameServer = types.BoolPointerValue(from.UseRootNameServer)
	m.UseRootServerForAllViews = types.BoolPointerValue(from.UseRootServerForAllViews)
	m.UseRpzDisableNsdnameNsip = types.BoolPointerValue(from.UseRpzDisableNsdnameNsip)
	m.UseRpzDropIpRule = types.BoolPointerValue(from.UseRpzDropIpRule)
	m.UseRpzQnameWaitRecurse = types.BoolPointerValue(from.UseRpzQnameWaitRecurse)
	m.UseSerialQueryRate = types.BoolPointerValue(from.UseSerialQueryRate)
	m.UseServerIdDirective = types.BoolPointerValue(from.UseServerIdDirective)
	m.UseSortlist = types.BoolPointerValue(from.UseSortlist)
	m.UseSourcePorts = types.BoolPointerValue(from.UseSourcePorts)
	m.UseSyslogFacility = types.BoolPointerValue(from.UseSyslogFacility)
	m.UseTransfersIn = types.BoolPointerValue(from.UseTransfersIn)
	m.UseTransfersOut = types.BoolPointerValue(from.UseTransfersOut)
	m.UseTransfersPerNs = types.BoolPointerValue(from.UseTransfersPerNs)
	m.UseUpdateSetting = types.BoolPointerValue(from.UseUpdateSetting)
	m.UseZoneTransferFormat = types.BoolPointerValue(from.UseZoneTransferFormat)
	m.Views = flex.FlattenFrameworkListString(ctx, from.Views, diags)
}
