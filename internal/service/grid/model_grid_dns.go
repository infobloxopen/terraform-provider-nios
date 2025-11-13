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

type GridDnsModel struct {
	Ref                                 types.String `tfsdk:"ref"`
	AddClientIpMacOptions               types.Bool   `tfsdk:"add_client_ip_mac_options"`
	AllowBulkhostDdns                   types.String `tfsdk:"allow_bulkhost_ddns"`
	AllowGssTsigZoneUpdates             types.Bool   `tfsdk:"allow_gss_tsig_zone_updates"`
	AllowQuery                          types.List   `tfsdk:"allow_query"`
	AllowRecursiveQuery                 types.Bool   `tfsdk:"allow_recursive_query"`
	AllowTransfer                       types.List   `tfsdk:"allow_transfer"`
	AllowUpdate                         types.List   `tfsdk:"allow_update"`
	AnonymizeResponseLogging            types.Bool   `tfsdk:"anonymize_response_logging"`
	AttackMitigation                    types.Object `tfsdk:"attack_mitigation"`
	AutoBlackhole                       types.Object `tfsdk:"auto_blackhole"`
	BindCheckNamesPolicy                types.String `tfsdk:"bind_check_names_policy"`
	BindHostnameDirective               types.String `tfsdk:"bind_hostname_directive"`
	BlackholeList                       types.List   `tfsdk:"blackhole_list"`
	BlacklistAction                     types.String `tfsdk:"blacklist_action"`
	BlacklistLogQuery                   types.Bool   `tfsdk:"blacklist_log_query"`
	BlacklistRedirectAddresses          types.List   `tfsdk:"blacklist_redirect_addresses"`
	BlacklistRedirectTtl                types.Int64  `tfsdk:"blacklist_redirect_ttl"`
	BlacklistRulesets                   types.List   `tfsdk:"blacklist_rulesets"`
	BulkHostNameTemplates               types.List   `tfsdk:"bulk_host_name_templates"`
	CaptureDnsQueriesOnAllDomains       types.Bool   `tfsdk:"capture_dns_queries_on_all_domains"`
	CheckNamesForDdnsAndZoneTransfer    types.Bool   `tfsdk:"check_names_for_ddns_and_zone_transfer"`
	ClientSubnetDomains                 types.List   `tfsdk:"client_subnet_domains"`
	ClientSubnetIpv4PrefixLength        types.Int64  `tfsdk:"client_subnet_ipv4_prefix_length"`
	ClientSubnetIpv6PrefixLength        types.Int64  `tfsdk:"client_subnet_ipv6_prefix_length"`
	CopyClientIpMacOptions              types.Bool   `tfsdk:"copy_client_ip_mac_options"`
	CopyXferToNotify                    types.Bool   `tfsdk:"copy_xfer_to_notify"`
	CustomRootNameServers               types.List   `tfsdk:"custom_root_name_servers"`
	DdnsForceCreationTimestampUpdate    types.Bool   `tfsdk:"ddns_force_creation_timestamp_update"`
	DdnsPrincipalGroup                  types.String `tfsdk:"ddns_principal_group"`
	DdnsPrincipalTracking               types.Bool   `tfsdk:"ddns_principal_tracking"`
	DdnsRestrictPatterns                types.Bool   `tfsdk:"ddns_restrict_patterns"`
	DdnsRestrictPatternsList            types.List   `tfsdk:"ddns_restrict_patterns_list"`
	DdnsRestrictProtected               types.Bool   `tfsdk:"ddns_restrict_protected"`
	DdnsRestrictSecure                  types.Bool   `tfsdk:"ddns_restrict_secure"`
	DdnsRestrictStatic                  types.Bool   `tfsdk:"ddns_restrict_static"`
	DefaultBulkHostNameTemplate         types.String `tfsdk:"default_bulk_host_name_template"`
	DefaultTtl                          types.Int64  `tfsdk:"default_ttl"`
	DisableEdns                         types.Bool   `tfsdk:"disable_edns"`
	Dns64Groups                         types.List   `tfsdk:"dns64_groups"`
	DnsCacheAccelerationTtl             types.Int64  `tfsdk:"dns_cache_acceleration_ttl"`
	DnsHealthCheckAnycastControl        types.Bool   `tfsdk:"dns_health_check_anycast_control"`
	DnsHealthCheckDomainList            types.List   `tfsdk:"dns_health_check_domain_list"`
	DnsHealthCheckInterval              types.Int64  `tfsdk:"dns_health_check_interval"`
	DnsHealthCheckRecursionFlag         types.Bool   `tfsdk:"dns_health_check_recursion_flag"`
	DnsHealthCheckRetries               types.Int64  `tfsdk:"dns_health_check_retries"`
	DnsHealthCheckTimeout               types.Int64  `tfsdk:"dns_health_check_timeout"`
	DnsQueryCaptureFileTimeLimit        types.Int64  `tfsdk:"dns_query_capture_file_time_limit"`
	DnssecBlacklistEnabled              types.Bool   `tfsdk:"dnssec_blacklist_enabled"`
	DnssecDns64Enabled                  types.Bool   `tfsdk:"dnssec_dns64_enabled"`
	DnssecEnabled                       types.Bool   `tfsdk:"dnssec_enabled"`
	DnssecExpiredSignaturesEnabled      types.Bool   `tfsdk:"dnssec_expired_signatures_enabled"`
	DnssecKeyParams                     types.Object `tfsdk:"dnssec_key_params"`
	DnssecNegativeTrustAnchors          types.List   `tfsdk:"dnssec_negative_trust_anchors"`
	DnssecNxdomainEnabled               types.Bool   `tfsdk:"dnssec_nxdomain_enabled"`
	DnssecRpzEnabled                    types.Bool   `tfsdk:"dnssec_rpz_enabled"`
	DnssecTrustedKeys                   types.List   `tfsdk:"dnssec_trusted_keys"`
	DnssecValidationEnabled             types.Bool   `tfsdk:"dnssec_validation_enabled"`
	DnstapSetting                       types.Object `tfsdk:"dnstap_setting"`
	DomainsToCaptureDnsQueries          types.List   `tfsdk:"domains_to_capture_dns_queries"`
	DtcDnsQueriesSpecificBehavior       types.String `tfsdk:"dtc_dns_queries_specific_behavior"`
	DtcDnssecMode                       types.String `tfsdk:"dtc_dnssec_mode"`
	DtcEdnsPreferClientSubnet           types.Bool   `tfsdk:"dtc_edns_prefer_client_subnet"`
	DtcScheduledBackup                  types.Object `tfsdk:"dtc_scheduled_backup"`
	DtcTopologyEaList                   types.List   `tfsdk:"dtc_topology_ea_list"`
	EdnsUdpSize                         types.Int64  `tfsdk:"edns_udp_size"`
	Email                               types.String `tfsdk:"email"`
	EnableBlackhole                     types.Bool   `tfsdk:"enable_blackhole"`
	EnableBlacklist                     types.Bool   `tfsdk:"enable_blacklist"`
	EnableCaptureDnsQueries             types.Bool   `tfsdk:"enable_capture_dns_queries"`
	EnableCaptureDnsResponses           types.Bool   `tfsdk:"enable_capture_dns_responses"`
	EnableClientSubnetForwarding        types.Bool   `tfsdk:"enable_client_subnet_forwarding"`
	EnableClientSubnetRecursive         types.Bool   `tfsdk:"enable_client_subnet_recursive"`
	EnableDeleteAssociatedPtr           types.Bool   `tfsdk:"enable_delete_associated_ptr"`
	EnableDns64                         types.Bool   `tfsdk:"enable_dns64"`
	EnableDnsHealthCheck                types.Bool   `tfsdk:"enable_dns_health_check"`
	EnableDnstapQueries                 types.Bool   `tfsdk:"enable_dnstap_queries"`
	EnableDnstapResponses               types.Bool   `tfsdk:"enable_dnstap_responses"`
	EnableDnstapViolationsTls           types.Bool   `tfsdk:"enable_dnstap_violations_tls"`
	EnableExcludedDomainNames           types.Bool   `tfsdk:"enable_excluded_domain_names"`
	EnableFixedRrsetOrderFqdns          types.Bool   `tfsdk:"enable_fixed_rrset_order_fqdns"`
	EnableFtc                           types.Bool   `tfsdk:"enable_ftc"`
	EnableGssTsig                       types.Bool   `tfsdk:"enable_gss_tsig"`
	EnableHostRrsetOrder                types.Bool   `tfsdk:"enable_host_rrset_order"`
	EnableHsmSigning                    types.Bool   `tfsdk:"enable_hsm_signing"`
	EnableNotifySourcePort              types.Bool   `tfsdk:"enable_notify_source_port"`
	EnableQueryRewrite                  types.Bool   `tfsdk:"enable_query_rewrite"`
	EnableQuerySourcePort               types.Bool   `tfsdk:"enable_query_source_port"`
	ExcludedDomainNames                 types.List   `tfsdk:"excluded_domain_names"`
	ExpireAfter                         types.Int64  `tfsdk:"expire_after"`
	FileTransferSetting                 types.Object `tfsdk:"file_transfer_setting"`
	FilterAaaa                          types.String `tfsdk:"filter_aaaa"`
	FilterAaaaList                      types.List   `tfsdk:"filter_aaaa_list"`
	FixedRrsetOrderFqdns                types.List   `tfsdk:"fixed_rrset_order_fqdns"`
	ForwardOnly                         types.Bool   `tfsdk:"forward_only"`
	ForwardUpdates                      types.Bool   `tfsdk:"forward_updates"`
	Forwarders                          types.List   `tfsdk:"forwarders"`
	FtcExpiredRecordTimeout             types.Int64  `tfsdk:"ftc_expired_record_timeout"`
	FtcExpiredRecordTtl                 types.Int64  `tfsdk:"ftc_expired_record_ttl"`
	GenEadbFromHosts                    types.Bool   `tfsdk:"gen_eadb_from_hosts"`
	GenEadbFromNetworkContainers        types.Bool   `tfsdk:"gen_eadb_from_network_containers"`
	GenEadbFromNetworks                 types.Bool   `tfsdk:"gen_eadb_from_networks"`
	GenEadbFromRanges                   types.Bool   `tfsdk:"gen_eadb_from_ranges"`
	GssTsigKeys                         types.List   `tfsdk:"gss_tsig_keys"`
	LastQueriedAcl                      types.List   `tfsdk:"last_queried_acl"`
	LoggingCategories                   types.Object `tfsdk:"logging_categories"`
	MaxCacheTtl                         types.Int64  `tfsdk:"max_cache_ttl"`
	MaxCachedLifetime                   types.Int64  `tfsdk:"max_cached_lifetime"`
	MaxNcacheTtl                        types.Int64  `tfsdk:"max_ncache_ttl"`
	MaxUdpSize                          types.Int64  `tfsdk:"max_udp_size"`
	MemberSecondaryNotify               types.Bool   `tfsdk:"member_secondary_notify"`
	NegativeTtl                         types.Int64  `tfsdk:"negative_ttl"`
	NotifyDelay                         types.Int64  `tfsdk:"notify_delay"`
	NotifySourcePort                    types.Int64  `tfsdk:"notify_source_port"`
	NsgroupDefault                      types.String `tfsdk:"nsgroup_default"`
	Nsgroups                            types.List   `tfsdk:"nsgroups"`
	NxdomainLogQuery                    types.Bool   `tfsdk:"nxdomain_log_query"`
	NxdomainRedirect                    types.Bool   `tfsdk:"nxdomain_redirect"`
	NxdomainRedirectAddresses           types.List   `tfsdk:"nxdomain_redirect_addresses"`
	NxdomainRedirectAddressesV6         types.List   `tfsdk:"nxdomain_redirect_addresses_v6"`
	NxdomainRedirectTtl                 types.Int64  `tfsdk:"nxdomain_redirect_ttl"`
	NxdomainRulesets                    types.List   `tfsdk:"nxdomain_rulesets"`
	PreserveHostRrsetOrderOnSecondaries types.Bool   `tfsdk:"preserve_host_rrset_order_on_secondaries"`
	ProtocolRecordNamePolicies          types.List   `tfsdk:"protocol_record_name_policies"`
	QueryRewriteDomainNames             types.List   `tfsdk:"query_rewrite_domain_names"`
	QueryRewritePrefix                  types.String `tfsdk:"query_rewrite_prefix"`
	QuerySourcePort                     types.Int64  `tfsdk:"query_source_port"`
	RecursiveQueryList                  types.List   `tfsdk:"recursive_query_list"`
	RefreshTimer                        types.Int64  `tfsdk:"refresh_timer"`
	ResolverQueryTimeout                types.Int64  `tfsdk:"resolver_query_timeout"`
	ResponseRateLimiting                types.Object `tfsdk:"response_rate_limiting"`
	RestartSetting                      types.Object `tfsdk:"restart_setting"`
	RetryTimer                          types.Int64  `tfsdk:"retry_timer"`
	RootNameServerType                  types.String `tfsdk:"root_name_server_type"`
	RpzDisableNsdnameNsip               types.Bool   `tfsdk:"rpz_disable_nsdname_nsip"`
	RpzDropIpRuleEnabled                types.Bool   `tfsdk:"rpz_drop_ip_rule_enabled"`
	RpzDropIpRuleMinPrefixLengthIpv4    types.Int64  `tfsdk:"rpz_drop_ip_rule_min_prefix_length_ipv4"`
	RpzDropIpRuleMinPrefixLengthIpv6    types.Int64  `tfsdk:"rpz_drop_ip_rule_min_prefix_length_ipv6"`
	RpzQnameWaitRecurse                 types.Bool   `tfsdk:"rpz_qname_wait_recurse"`
	ScavengingSettings                  types.Object `tfsdk:"scavenging_settings"`
	SerialQueryRate                     types.Int64  `tfsdk:"serial_query_rate"`
	ServerIdDirective                   types.String `tfsdk:"server_id_directive"`
	Sortlist                            types.List   `tfsdk:"sortlist"`
	StoreLocally                        types.Bool   `tfsdk:"store_locally"`
	SyslogFacility                      types.String `tfsdk:"syslog_facility"`
	TransferExcludedServers             types.List   `tfsdk:"transfer_excluded_servers"`
	TransferFormat                      types.String `tfsdk:"transfer_format"`
	TransfersIn                         types.Int64  `tfsdk:"transfers_in"`
	TransfersOut                        types.Int64  `tfsdk:"transfers_out"`
	TransfersPerNs                      types.Int64  `tfsdk:"transfers_per_ns"`
	ZoneDeletionDoubleConfirm           types.Bool   `tfsdk:"zone_deletion_double_confirm"`
}

var GridDnsAttrTypes = map[string]attr.Type{
	"ref":                                      types.StringType,
	"add_client_ip_mac_options":                types.BoolType,
	"allow_bulkhost_ddns":                      types.StringType,
	"allow_gss_tsig_zone_updates":              types.BoolType,
	"allow_query":                              types.ListType{ElemType: types.ObjectType{AttrTypes: GridDnsAllowQueryAttrTypes}},
	"allow_recursive_query":                    types.BoolType,
	"allow_transfer":                           types.ListType{ElemType: types.ObjectType{AttrTypes: GridDnsAllowTransferAttrTypes}},
	"allow_update":                             types.ListType{ElemType: types.ObjectType{AttrTypes: GridDnsAllowUpdateAttrTypes}},
	"anonymize_response_logging":               types.BoolType,
	"attack_mitigation":                        types.ObjectType{AttrTypes: GridDnsAttackMitigationAttrTypes},
	"auto_blackhole":                           types.ObjectType{AttrTypes: GridDnsAutoBlackholeAttrTypes},
	"bind_check_names_policy":                  types.StringType,
	"bind_hostname_directive":                  types.StringType,
	"blackhole_list":                           types.ListType{ElemType: types.ObjectType{AttrTypes: GridDnsBlackholeListAttrTypes}},
	"blacklist_action":                         types.StringType,
	"blacklist_log_query":                      types.BoolType,
	"blacklist_redirect_addresses":             types.ListType{ElemType: types.StringType},
	"blacklist_redirect_ttl":                   types.Int64Type,
	"blacklist_rulesets":                       types.ListType{ElemType: types.StringType},
	"bulk_host_name_templates":                 types.ListType{ElemType: types.StringType},
	"capture_dns_queries_on_all_domains":       types.BoolType,
	"check_names_for_ddns_and_zone_transfer":   types.BoolType,
	"client_subnet_domains":                    types.ListType{ElemType: types.ObjectType{AttrTypes: GridDnsClientSubnetDomainsAttrTypes}},
	"client_subnet_ipv4_prefix_length":         types.Int64Type,
	"client_subnet_ipv6_prefix_length":         types.Int64Type,
	"copy_client_ip_mac_options":               types.BoolType,
	"copy_xfer_to_notify":                      types.BoolType,
	"custom_root_name_servers":                 types.ListType{ElemType: types.ObjectType{AttrTypes: GridDnsCustomRootNameServersAttrTypes}},
	"ddns_force_creation_timestamp_update":     types.BoolType,
	"ddns_principal_group":                     types.StringType,
	"ddns_principal_tracking":                  types.BoolType,
	"ddns_restrict_patterns":                   types.BoolType,
	"ddns_restrict_patterns_list":              types.ListType{ElemType: types.StringType},
	"ddns_restrict_protected":                  types.BoolType,
	"ddns_restrict_secure":                     types.BoolType,
	"ddns_restrict_static":                     types.BoolType,
	"default_bulk_host_name_template":          types.StringType,
	"default_ttl":                              types.Int64Type,
	"disable_edns":                             types.BoolType,
	"dns64_groups":                             types.ListType{ElemType: types.StringType},
	"dns_cache_acceleration_ttl":               types.Int64Type,
	"dns_health_check_anycast_control":         types.BoolType,
	"dns_health_check_domain_list":             types.ListType{ElemType: types.StringType},
	"dns_health_check_interval":                types.Int64Type,
	"dns_health_check_recursion_flag":          types.BoolType,
	"dns_health_check_retries":                 types.Int64Type,
	"dns_health_check_timeout":                 types.Int64Type,
	"dns_query_capture_file_time_limit":        types.Int64Type,
	"dnssec_blacklist_enabled":                 types.BoolType,
	"dnssec_dns64_enabled":                     types.BoolType,
	"dnssec_enabled":                           types.BoolType,
	"dnssec_expired_signatures_enabled":        types.BoolType,
	"dnssec_key_params":                        types.ObjectType{AttrTypes: GridDnsDnssecKeyParamsAttrTypes},
	"dnssec_negative_trust_anchors":            types.ListType{ElemType: types.StringType},
	"dnssec_nxdomain_enabled":                  types.BoolType,
	"dnssec_rpz_enabled":                       types.BoolType,
	"dnssec_trusted_keys":                      types.ListType{ElemType: types.ObjectType{AttrTypes: GridDnsDnssecTrustedKeysAttrTypes}},
	"dnssec_validation_enabled":                types.BoolType,
	"dnstap_setting":                           types.ObjectType{AttrTypes: GridDnsDnstapSettingAttrTypes},
	"domains_to_capture_dns_queries":           types.ListType{ElemType: types.StringType},
	"dtc_dns_queries_specific_behavior":        types.StringType,
	"dtc_dnssec_mode":                          types.StringType,
	"dtc_edns_prefer_client_subnet":            types.BoolType,
	"dtc_scheduled_backup":                     types.ObjectType{AttrTypes: GridDnsDtcScheduledBackupAttrTypes},
	"dtc_topology_ea_list":                     types.ListType{ElemType: types.StringType},
	"edns_udp_size":                            types.Int64Type,
	"email":                                    types.StringType,
	"enable_blackhole":                         types.BoolType,
	"enable_blacklist":                         types.BoolType,
	"enable_capture_dns_queries":               types.BoolType,
	"enable_capture_dns_responses":             types.BoolType,
	"enable_client_subnet_forwarding":          types.BoolType,
	"enable_client_subnet_recursive":           types.BoolType,
	"enable_delete_associated_ptr":             types.BoolType,
	"enable_dns64":                             types.BoolType,
	"enable_dns_health_check":                  types.BoolType,
	"enable_dnstap_queries":                    types.BoolType,
	"enable_dnstap_responses":                  types.BoolType,
	"enable_dnstap_violations_tls":             types.BoolType,
	"enable_excluded_domain_names":             types.BoolType,
	"enable_fixed_rrset_order_fqdns":           types.BoolType,
	"enable_ftc":                               types.BoolType,
	"enable_gss_tsig":                          types.BoolType,
	"enable_host_rrset_order":                  types.BoolType,
	"enable_hsm_signing":                       types.BoolType,
	"enable_notify_source_port":                types.BoolType,
	"enable_query_rewrite":                     types.BoolType,
	"enable_query_source_port":                 types.BoolType,
	"excluded_domain_names":                    types.ListType{ElemType: types.StringType},
	"expire_after":                             types.Int64Type,
	"file_transfer_setting":                    types.ObjectType{AttrTypes: GridDnsFileTransferSettingAttrTypes},
	"filter_aaaa":                              types.StringType,
	"filter_aaaa_list":                         types.ListType{ElemType: types.ObjectType{AttrTypes: GridDnsFilterAaaaListAttrTypes}},
	"fixed_rrset_order_fqdns":                  types.ListType{ElemType: types.ObjectType{AttrTypes: GridDnsFixedRrsetOrderFqdnsAttrTypes}},
	"forward_only":                             types.BoolType,
	"forward_updates":                          types.BoolType,
	"forwarders":                               types.ListType{ElemType: types.StringType},
	"ftc_expired_record_timeout":               types.Int64Type,
	"ftc_expired_record_ttl":                   types.Int64Type,
	"gen_eadb_from_hosts":                      types.BoolType,
	"gen_eadb_from_network_containers":         types.BoolType,
	"gen_eadb_from_networks":                   types.BoolType,
	"gen_eadb_from_ranges":                     types.BoolType,
	"gss_tsig_keys":                            types.ListType{ElemType: types.StringType},
	"last_queried_acl":                         types.ListType{ElemType: types.ObjectType{AttrTypes: GridDnsLastQueriedAclAttrTypes}},
	"logging_categories":                       types.ObjectType{AttrTypes: GridDnsLoggingCategoriesAttrTypes},
	"max_cache_ttl":                            types.Int64Type,
	"max_cached_lifetime":                      types.Int64Type,
	"max_ncache_ttl":                           types.Int64Type,
	"max_udp_size":                             types.Int64Type,
	"member_secondary_notify":                  types.BoolType,
	"negative_ttl":                             types.Int64Type,
	"notify_delay":                             types.Int64Type,
	"notify_source_port":                       types.Int64Type,
	"nsgroup_default":                          types.StringType,
	"nsgroups":                                 types.ListType{ElemType: types.StringType},
	"nxdomain_log_query":                       types.BoolType,
	"nxdomain_redirect":                        types.BoolType,
	"nxdomain_redirect_addresses":              types.ListType{ElemType: types.StringType},
	"nxdomain_redirect_addresses_v6":           types.ListType{ElemType: types.StringType},
	"nxdomain_redirect_ttl":                    types.Int64Type,
	"nxdomain_rulesets":                        types.ListType{ElemType: types.StringType},
	"preserve_host_rrset_order_on_secondaries": types.BoolType,
	"protocol_record_name_policies":            types.ListType{ElemType: types.StringType},
	"query_rewrite_domain_names":               types.ListType{ElemType: types.StringType},
	"query_rewrite_prefix":                     types.StringType,
	"query_source_port":                        types.Int64Type,
	"recursive_query_list":                     types.ListType{ElemType: types.ObjectType{AttrTypes: GridDnsRecursiveQueryListAttrTypes}},
	"refresh_timer":                            types.Int64Type,
	"resolver_query_timeout":                   types.Int64Type,
	"response_rate_limiting":                   types.ObjectType{AttrTypes: GridDnsResponseRateLimitingAttrTypes},
	"restart_setting":                          types.ObjectType{AttrTypes: GridDnsRestartSettingAttrTypes},
	"retry_timer":                              types.Int64Type,
	"root_name_server_type":                    types.StringType,
	"rpz_disable_nsdname_nsip":                 types.BoolType,
	"rpz_drop_ip_rule_enabled":                 types.BoolType,
	"rpz_drop_ip_rule_min_prefix_length_ipv4":  types.Int64Type,
	"rpz_drop_ip_rule_min_prefix_length_ipv6":  types.Int64Type,
	"rpz_qname_wait_recurse":                   types.BoolType,
	"scavenging_settings":                      types.ObjectType{AttrTypes: GridDnsScavengingSettingsAttrTypes},
	"serial_query_rate":                        types.Int64Type,
	"server_id_directive":                      types.StringType,
	"sortlist":                                 types.ListType{ElemType: types.ObjectType{AttrTypes: GridDnsSortlistAttrTypes}},
	"store_locally":                            types.BoolType,
	"syslog_facility":                          types.StringType,
	"transfer_excluded_servers":                types.ListType{ElemType: types.StringType},
	"transfer_format":                          types.StringType,
	"transfers_in":                             types.Int64Type,
	"transfers_out":                            types.Int64Type,
	"transfers_per_ns":                         types.Int64Type,
	"zone_deletion_double_confirm":             types.BoolType,
}

var GridDnsResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"add_client_ip_mac_options": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Add custom IP, MAC and DNS View name ENDS0 options to outgoing recursive queries.",
	},
	"allow_bulkhost_ddns": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if DDNS bulk host is allowed or not.",
	},
	"allow_gss_tsig_zone_updates": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether GSS-TSIG zone update is enabled for all Grid members.",
	},
	"allow_query": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: GridDnsAllowQueryResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "Determines if queries from the specified IPv4 or IPv6 addresses and networks are allowed or not. The appliance can also use Transaction Signature (TSIG) keys to authenticate the queries.",
	},
	"allow_recursive_query": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the responses to recursive queries are enabled or not.",
	},
	"allow_transfer": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: GridDnsAllowTransferResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "Determines if zone transfers from specified IPv4 or IPv6 addresses and networks or transfers from hosts authenticated by Transaction signature (TSIG) key are allowed or not.",
	},
	"allow_update": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: GridDnsAllowUpdateResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "Determines if dynamic updates from specified IPv4 or IPv6 addresses, networks or from host authenticated by TSIG key are allowed or not.",
	},
	"anonymize_response_logging": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the anonymization of captured DNS responses is enabled or disabled.",
	},
	"attack_mitigation": schema.SingleNestedAttribute{
		Attributes: GridDnsAttackMitigationResourceSchemaAttributes,
		Optional:   true,
	},
	"auto_blackhole": schema.SingleNestedAttribute{
		Attributes: GridDnsAutoBlackholeResourceSchemaAttributes,
		Optional:   true,
	},
	"bind_check_names_policy": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The BIND check names policy, which indicates the action the appliance takes when it encounters host names that do not comply with the Strict Hostname Checking policy. This method applies only if the host name restriction policy is set to \"Strict Hostname Checking\".",
	},
	"bind_hostname_directive": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The value of the hostname directive for BIND.",
	},
	"blackhole_list": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: GridDnsBlackholeListResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of IPv4 or IPv6 addresses and networks from which DNS queries are blocked.",
	},
	"blacklist_action": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The action to perform when a domain name matches the pattern defined in a rule that is specified by the blacklist ruleset.",
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
		MarkdownDescription: "The TTL value (in seconds) of the synthetic DNS responses that result from blacklist redirection.",
	},
	"blacklist_rulesets": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The DNS Ruleset object names assigned at the Grid level for blacklist redirection.",
	},
	"bulk_host_name_templates": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of bulk host name templates. There are four Infoblox predefined bulk host name templates. Template Name Template Format \"Four Octets\" -$1-$2-$3-$4 \"Three Octets\" -$2-$3-$4 \"Two Octets\" -$3-$4 \"One Octet\" -$4",
	},
	"capture_dns_queries_on_all_domains": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the capture of DNS queries for all domains is enabled or disabled.",
	},
	"check_names_for_ddns_and_zone_transfer": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the application of BIND check-names for zone transfers and DDNS updates are enabled.",
	},
	"client_subnet_domains": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: GridDnsClientSubnetDomainsResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of zone domain names that are allowed or forbidden for EDNS client subnet (ECS) recursion.",
	},
	"client_subnet_ipv4_prefix_length": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Default IPv4 Source Prefix-Length used when sending queries with EDNS client subnet option.",
	},
	"client_subnet_ipv6_prefix_length": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Default IPv6 Source Prefix-Length used when sending queries with EDNS client subnet option.",
	},
	"copy_client_ip_mac_options": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Copy custom IP, MAC and DNS View name ENDS0 options from incoming to outgoing recursive queries.",
	},
	"copy_xfer_to_notify": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "The allowed IPs, from the zone transfer list, added to the also-notify statement in the named.conf file.",
	},
	"custom_root_name_servers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: GridDnsCustomRootNameServersResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of customized root nameserver(s). You can use Internet root name servers or specify host names and IP addresses of custom root name servers.",
	},
	"ddns_force_creation_timestamp_update": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Defines whether creation timestamp of RR should be updated ' when DDNS update happens even if there is no change to ' the RR.",
	},
	"ddns_principal_group": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The DDNS Principal cluster group name.",
	},
	"ddns_principal_tracking": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the DDNS principal track is enabled or disabled.",
	},
	"ddns_restrict_patterns": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if an option to restrict DDNS update request based on FQDN patterns is enabled or disabled.",
	},
	"ddns_restrict_patterns_list": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The unordered list of restriction patterns for an option of to restrict DDNS updates based on FQDN patterns.",
	},
	"ddns_restrict_protected": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if an option to restrict DDNS update request to protected resource records is enabled or disabled.",
	},
	"ddns_restrict_secure": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if DDNS update request for principal other than target resource record's principal is restricted.",
	},
	"ddns_restrict_static": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if an option to restrict DDNS update request to resource records which are marked as 'STATIC' is enabled or disabled.",
	},
	"default_bulk_host_name_template": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Default bulk host name of a Grid DNS.",
	},
	"default_ttl": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The default TTL value of a Grid DNS object. This interval tells the secondary how long the data can be cached.",
	},
	"disable_edns": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the EDNS0 support for queries that require recursive resolution on Grid members is enabled or not.",
	},
	"dns64_groups": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of DNS64 synthesis groups associated with this Grid DNS object.",
	},
	"dns_cache_acceleration_ttl": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The minimum TTL value, in seconds, that a DNS record must have in order for it to be cached by the DNS Cache Acceleration service. An integer from 1 to 65000 that represents the TTL in seconds.",
	},
	"dns_health_check_anycast_control": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the anycast failure (BFD session down) is enabled on member failure or not.",
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
		MarkdownDescription: "Determines if the recursive DNS health check is enabled or not.",
	},
	"dns_health_check_retries": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The number of DNS health check retries.",
	},
	"dns_health_check_timeout": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The DNS health check timeout interval (in seconds).",
	},
	"dns_query_capture_file_time_limit": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The time limit (in minutes) for the DNS query capture file.",
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
	"dnssec_key_params": schema.SingleNestedAttribute{
		Attributes: GridDnsDnssecKeyParamsResourceSchemaAttributes,
		Optional:   true,
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
			Attributes: GridDnsDnssecTrustedKeysResourceSchemaAttributes,
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
		Attributes: GridDnsDnstapSettingResourceSchemaAttributes,
		Optional:   true,
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
	"dtc_dnssec_mode": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "DTC DNSSEC operation mode.",
	},
	"dtc_edns_prefer_client_subnet": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether to prefer the client address from the edns-client-subnet option for DTC or not.",
	},
	"dtc_scheduled_backup": schema.SingleNestedAttribute{
		Attributes: GridDnsDtcScheduledBackupResourceSchemaAttributes,
		Optional:   true,
	},
	"dtc_topology_ea_list": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The DTC topology extensible attribute definition list. When configuring a DTC topology, users may configure classification as either \"Geographic\" or \"Extensible Attributes\". Selecting extensible attributes will replace supported Topology database labels (Continent, Country, Subdivision, City) with the names of the selection EA types and provide values extracted from DHCP Network Container, Network and Range objects with those extensible attributes.",
	},
	"edns_udp_size": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Advertises the EDNS0 buffer size to the upstream server. The value should be between 512 and 4096 bytes. The recommended value is between 512 and 1220 bytes.",
	},
	"email": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The email address of a Grid DNS object.",
	},
	"enable_blackhole": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the blocking of DNS queries is enabled or not.",
	},
	"enable_blacklist": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if a blacklist is enabled or not.",
	},
	"enable_capture_dns_queries": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the capture of DNS queries is enabled or disabled.",
	},
	"enable_capture_dns_responses": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the capture of DNS responses is enabled or disabled.",
	},
	"enable_client_subnet_forwarding": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether to enable forwarding EDNS client subnet options to upstream servers.",
	},
	"enable_client_subnet_recursive": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether to enable adding EDNS client subnet options in recursive resolution. The client_subnet_domains parameter value must not be empty to enable the enable_client_subnet_recursive parameter.",
	},
	"enable_delete_associated_ptr": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the ability to automatically remove associated PTR records while deleting A or AAAA records is enabled or not.",
	},
	"enable_dns64": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the DNS64 support is enabled or not.",
	},
	"enable_dns_health_check": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the DNS health check is enabled or not.",
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
		MarkdownDescription: "Determines if excluding domain names from captured DNS queries and responses is enabled or disabled.",
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
		MarkdownDescription: "Determines whether all appliances in the Grid are enabled to receive GSS-TSIG authenticated updates from DNS clients.",
	},
	"enable_host_rrset_order": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the host RRset order is enabled or not.",
	},
	"enable_hsm_signing": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether Hardware Security Modules (HSMs) are enabled for key generation and signing. Note, that you must configure the HSM group with at least one enabled HSM.",
	},
	"enable_notify_source_port": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the notify source port at the Grid Level is enabled or not.",
	},
	"enable_query_rewrite": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the DNS query rewrite is enabled or not.",
	},
	"enable_query_source_port": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the query source port at the Grid Level is enabled or not.",
	},
	"excluded_domain_names": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of domains that are excluded from DNS query and response capture.",
	},
	"expire_after": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The expiration time of a Grid DNS object. If the secondary DNS server fails to contact the primary server for the specified interval, the secondary server stops giving out answers about the zone because the zone data is too old to be useful.",
	},
	"file_transfer_setting": schema.SingleNestedAttribute{
		Attributes: GridDnsFileTransferSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"filter_aaaa": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The type of AAAA filtering for this member DNS object.",
	},
	"filter_aaaa_list": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: GridDnsFilterAaaaListResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of IPv4 addresses and networks from which queries are received. AAAA filtering is applied to these addresses.",
	},
	"fixed_rrset_order_fqdns": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: GridDnsFixedRrsetOrderFqdnsResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The fixed RRset order FQDN. If this field does not contain an empty value, the appliance will automatically set the enable_fixed_rrset_order_fqdns field to 'true', unless the same request sets the enable field to 'false'.",
	},
	"forward_only": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if member sends queries to forwarders only. When the value is \"true\", the member sends queries to forwarders only, and not to other internal or Internet root servers.",
	},
	"forward_updates": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if secondary servers is allowed to forward updates to the DNS server or not.",
	},
	"forwarders": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The forwarders for the member. A forwarder is essentially a name server to which other name servers first send all of their off-site queries. The forwarder builds up a cache of information, avoiding the need for the other name servers to send queries off-site.",
	},
	"ftc_expired_record_timeout": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The timeout interval (in seconds) after which the expired Fault Tolerant Caching (FTC)record is stale and no longer valid.",
	},
	"ftc_expired_record_ttl": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The TTL value (in seconds) of the expired Fault Tolerant Caching (FTC) record in DNS responses.",
	},
	"gen_eadb_from_hosts": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Flag for taking EA values from IPAM Hosts into consideration for the DTC topology EA database.",
	},
	"gen_eadb_from_network_containers": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Flag for taking EA values from IPAM Network Containers into consideration for the DTC topology EA database.",
	},
	"gen_eadb_from_networks": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Flag for taking EA values from IPAM Network into consideration for the DTC topology EA database.",
	},
	"gen_eadb_from_ranges": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Flag for taking EA values from IPAM Ranges into consideration for the DTC topology EA database.",
	},
	"gss_tsig_keys": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of GSS-TSIG keys for a Grid DNS object.",
	},
	"last_queried_acl": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: GridDnsLastQueriedAclResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "Determines last queried ACL for the specified IPv4 or IPv6 addresses and networks in scavenging settings.",
	},
	"logging_categories": schema.SingleNestedAttribute{
		Attributes: GridDnsLoggingCategoriesResourceSchemaAttributes,
		Optional:   true,
	},
	"max_cache_ttl": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The maximum time (in seconds) for which the server will cache positive answers.",
	},
	"max_cached_lifetime": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The maximum time (in seconds) a DNS response can be stored in the hardware acceleration cache. Valid values are unsigned integer between 60 and 86400, inclusive.",
	},
	"max_ncache_ttl": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The maximum time (in seconds) for which the server will cache negative (NXDOMAIN) responses. The maximum allowed value is 604800.",
	},
	"max_udp_size": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The value is used by authoritative DNS servers to never send DNS responses larger than the configured value. The value should be between 512 and 4096 bytes. The recommended value is between 512 and 1220 bytes.",
	},
	"member_secondary_notify": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if Grid members that are authoritative secondary servers are allowed to send notification messages to external name servers, if the Grid member that is primary for a zone fails or loses connectivity.",
	},
	"negative_ttl": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The negative TTL value of a Grid DNS object. This interval tells the secondary how long data can be cached for \"Does Not Respond\" responses.",
	},
	"notify_delay": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Specifies with how many seconds of delay the notify messages are sent to secondaries.",
	},
	"notify_source_port": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The source port for notify messages. When requesting zone transfers from the primary server, some secondary DNS servers use the source port number (the primary server used to send the notify message) as the destination port number in the zone transfer request. Valid values are between 1 and 63999. The default is picked by BIND.",
	},
	"nsgroup_default": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The default nameserver group.",
	},
	"nsgroups": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "A name server group is a collection of one primary DNS server and one or more secondary DNS servers.",
	},
	"nxdomain_log_query": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if NXDOMAIN redirection queries are logged or not.",
	},
	"nxdomain_redirect": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if NXDOMAIN redirection is enabled or not.",
	},
	"nxdomain_redirect_addresses": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of IPv4 NXDOMAIN redirection addresses.",
	},
	"nxdomain_redirect_addresses_v6": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of IPv6 NXDOMAIN redirection addresses.",
	},
	"nxdomain_redirect_ttl": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The TTL value (in seconds) of synthetic DNS responses that result from NXDOMAIN redirection.",
	},
	"nxdomain_rulesets": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The Ruleset object names assigned at the Grid level for NXDOMAIN redirection.",
	},
	"preserve_host_rrset_order_on_secondaries": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the host RRset order on secondaries is preserved or not.",
	},
	"protocol_record_name_policies": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of record name policies.",
	},
	"query_rewrite_domain_names": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of domain names that trigger DNS query rewrite.",
	},
	"query_rewrite_prefix": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The domain name prefix for DNS query rewrite.",
	},
	"query_source_port": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The source port for queries. Specifying a source port number for recursive queries ensures that a firewall will allow the response. Valid values are between 1 and 63999. The default is picked by BIND.",
	},
	"recursive_query_list": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: GridDnsRecursiveQueryListResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of IPv4 or IPv6 addresses, networks or hosts authenticated by Transaction signature (TSIG) key from which recursive queries are allowed or denied.",
	},
	"refresh_timer": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The refresh time. This interval tells the secondary how often to send a message to the primary for a zone to check that its data is current, and retrieve fresh data if it is not.",
	},
	"resolver_query_timeout": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The recursive query timeout for the member.",
	},
	"response_rate_limiting": schema.SingleNestedAttribute{
		Attributes: GridDnsResponseRateLimitingResourceSchemaAttributes,
		Optional:   true,
	},
	"restart_setting": schema.SingleNestedAttribute{
		Attributes: GridDnsRestartSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"retry_timer": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The retry time. This interval tells the secondary how long to wait before attempting to recontact the primary after a connection failure occurs between the two servers.",
	},
	"root_name_server_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Determines the type of root name servers.",
	},
	"rpz_disable_nsdname_nsip": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if NSDNAME and NSIP resource records from RPZ feeds are enabled or not.",
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
		MarkdownDescription: "Determines if recursive RPZ lookups are enabled.",
	},
	"scavenging_settings": schema.SingleNestedAttribute{
		Attributes: GridDnsScavengingSettingsResourceSchemaAttributes,
		Optional:   true,
	},
	"serial_query_rate": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The number of maximum concurrent SOA queries per second. Valid values are unsigned integer between 20 and 1000, inclusive.",
	},
	"server_id_directive": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The value of the server-id directive for BIND DNS.",
	},
	"sortlist": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: GridDnsSortlistResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "A sort list determines the order of addresses in responses made to DNS queries.",
	},
	"store_locally": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the storage of query capture reports on the appliance is enabled or disabled.",
	},
	"syslog_facility": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The syslog facility. This is the location on the syslog server to which you want to sort the DNS logging messages.",
	},
	"transfer_excluded_servers": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of excluded DNS servers during zone transfers.",
	},
	"transfer_format": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The BIND format for a zone transfer. This provides tracking capabilities for single or multiple transfers and their associated servers.",
	},
	"transfers_in": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The number of maximum concurrent transfers for the Grid. Valid values are unsigned integer between 10 and 10000, inclusive.",
	},
	"transfers_out": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The number of maximum outbound concurrent zone transfers. Valid values are unsigned integer between 10 and 10000, inclusive.",
	},
	"transfers_per_ns": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The number of maximum concurrent transfers per member. Valid values are unsigned integer between 2 and 10000, inclusive.",
	},
	"zone_deletion_double_confirm": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the double confirmation during zone deletion is enabled or not.",
	},
}

func ExpandGridDns(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridDns {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridDnsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridDnsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridDns {
	if m == nil {
		return nil
	}
	to := &grid.GridDns{
		Ref:                                 flex.ExpandStringPointer(m.Ref),
		AddClientIpMacOptions:               flex.ExpandBoolPointer(m.AddClientIpMacOptions),
		AllowBulkhostDdns:                   flex.ExpandStringPointer(m.AllowBulkhostDdns),
		AllowGssTsigZoneUpdates:             flex.ExpandBoolPointer(m.AllowGssTsigZoneUpdates),
		AllowQuery:                          flex.ExpandFrameworkListNestedBlock(ctx, m.AllowQuery, diags, ExpandGridDnsAllowQuery),
		AllowRecursiveQuery:                 flex.ExpandBoolPointer(m.AllowRecursiveQuery),
		AllowTransfer:                       flex.ExpandFrameworkListNestedBlock(ctx, m.AllowTransfer, diags, ExpandGridDnsAllowTransfer),
		AllowUpdate:                         flex.ExpandFrameworkListNestedBlock(ctx, m.AllowUpdate, diags, ExpandGridDnsAllowUpdate),
		AnonymizeResponseLogging:            flex.ExpandBoolPointer(m.AnonymizeResponseLogging),
		AttackMitigation:                    ExpandGridDnsAttackMitigation(ctx, m.AttackMitigation, diags),
		AutoBlackhole:                       ExpandGridDnsAutoBlackhole(ctx, m.AutoBlackhole, diags),
		BindCheckNamesPolicy:                flex.ExpandStringPointer(m.BindCheckNamesPolicy),
		BindHostnameDirective:               flex.ExpandStringPointer(m.BindHostnameDirective),
		BlackholeList:                       flex.ExpandFrameworkListNestedBlock(ctx, m.BlackholeList, diags, ExpandGridDnsBlackholeList),
		BlacklistAction:                     flex.ExpandStringPointer(m.BlacklistAction),
		BlacklistLogQuery:                   flex.ExpandBoolPointer(m.BlacklistLogQuery),
		BlacklistRedirectAddresses:          flex.ExpandFrameworkListString(ctx, m.BlacklistRedirectAddresses, diags),
		BlacklistRedirectTtl:                flex.ExpandInt64Pointer(m.BlacklistRedirectTtl),
		BlacklistRulesets:                   flex.ExpandFrameworkListString(ctx, m.BlacklistRulesets, diags),
		BulkHostNameTemplates:               flex.ExpandFrameworkListString(ctx, m.BulkHostNameTemplates, diags),
		CaptureDnsQueriesOnAllDomains:       flex.ExpandBoolPointer(m.CaptureDnsQueriesOnAllDomains),
		CheckNamesForDdnsAndZoneTransfer:    flex.ExpandBoolPointer(m.CheckNamesForDdnsAndZoneTransfer),
		ClientSubnetDomains:                 flex.ExpandFrameworkListNestedBlock(ctx, m.ClientSubnetDomains, diags, ExpandGridDnsClientSubnetDomains),
		ClientSubnetIpv4PrefixLength:        flex.ExpandInt64Pointer(m.ClientSubnetIpv4PrefixLength),
		ClientSubnetIpv6PrefixLength:        flex.ExpandInt64Pointer(m.ClientSubnetIpv6PrefixLength),
		CopyClientIpMacOptions:              flex.ExpandBoolPointer(m.CopyClientIpMacOptions),
		CopyXferToNotify:                    flex.ExpandBoolPointer(m.CopyXferToNotify),
		CustomRootNameServers:               flex.ExpandFrameworkListNestedBlock(ctx, m.CustomRootNameServers, diags, ExpandGridDnsCustomRootNameServers),
		DdnsForceCreationTimestampUpdate:    flex.ExpandBoolPointer(m.DdnsForceCreationTimestampUpdate),
		DdnsPrincipalGroup:                  flex.ExpandStringPointer(m.DdnsPrincipalGroup),
		DdnsPrincipalTracking:               flex.ExpandBoolPointer(m.DdnsPrincipalTracking),
		DdnsRestrictPatterns:                flex.ExpandBoolPointer(m.DdnsRestrictPatterns),
		DdnsRestrictPatternsList:            flex.ExpandFrameworkListString(ctx, m.DdnsRestrictPatternsList, diags),
		DdnsRestrictProtected:               flex.ExpandBoolPointer(m.DdnsRestrictProtected),
		DdnsRestrictSecure:                  flex.ExpandBoolPointer(m.DdnsRestrictSecure),
		DdnsRestrictStatic:                  flex.ExpandBoolPointer(m.DdnsRestrictStatic),
		DefaultBulkHostNameTemplate:         flex.ExpandStringPointer(m.DefaultBulkHostNameTemplate),
		DefaultTtl:                          flex.ExpandInt64Pointer(m.DefaultTtl),
		DisableEdns:                         flex.ExpandBoolPointer(m.DisableEdns),
		Dns64Groups:                         flex.ExpandFrameworkListString(ctx, m.Dns64Groups, diags),
		DnsCacheAccelerationTtl:             flex.ExpandInt64Pointer(m.DnsCacheAccelerationTtl),
		DnsHealthCheckAnycastControl:        flex.ExpandBoolPointer(m.DnsHealthCheckAnycastControl),
		DnsHealthCheckDomainList:            flex.ExpandFrameworkListString(ctx, m.DnsHealthCheckDomainList, diags),
		DnsHealthCheckInterval:              flex.ExpandInt64Pointer(m.DnsHealthCheckInterval),
		DnsHealthCheckRecursionFlag:         flex.ExpandBoolPointer(m.DnsHealthCheckRecursionFlag),
		DnsHealthCheckRetries:               flex.ExpandInt64Pointer(m.DnsHealthCheckRetries),
		DnsHealthCheckTimeout:               flex.ExpandInt64Pointer(m.DnsHealthCheckTimeout),
		DnsQueryCaptureFileTimeLimit:        flex.ExpandInt64Pointer(m.DnsQueryCaptureFileTimeLimit),
		DnssecBlacklistEnabled:              flex.ExpandBoolPointer(m.DnssecBlacklistEnabled),
		DnssecDns64Enabled:                  flex.ExpandBoolPointer(m.DnssecDns64Enabled),
		DnssecEnabled:                       flex.ExpandBoolPointer(m.DnssecEnabled),
		DnssecExpiredSignaturesEnabled:      flex.ExpandBoolPointer(m.DnssecExpiredSignaturesEnabled),
		DnssecKeyParams:                     ExpandGridDnsDnssecKeyParams(ctx, m.DnssecKeyParams, diags),
		DnssecNegativeTrustAnchors:          flex.ExpandFrameworkListString(ctx, m.DnssecNegativeTrustAnchors, diags),
		DnssecNxdomainEnabled:               flex.ExpandBoolPointer(m.DnssecNxdomainEnabled),
		DnssecRpzEnabled:                    flex.ExpandBoolPointer(m.DnssecRpzEnabled),
		DnssecTrustedKeys:                   flex.ExpandFrameworkListNestedBlock(ctx, m.DnssecTrustedKeys, diags, ExpandGridDnsDnssecTrustedKeys),
		DnssecValidationEnabled:             flex.ExpandBoolPointer(m.DnssecValidationEnabled),
		DnstapSetting:                       ExpandGridDnsDnstapSetting(ctx, m.DnstapSetting, diags),
		DomainsToCaptureDnsQueries:          flex.ExpandFrameworkListString(ctx, m.DomainsToCaptureDnsQueries, diags),
		DtcDnsQueriesSpecificBehavior:       flex.ExpandStringPointer(m.DtcDnsQueriesSpecificBehavior),
		DtcDnssecMode:                       flex.ExpandStringPointer(m.DtcDnssecMode),
		DtcEdnsPreferClientSubnet:           flex.ExpandBoolPointer(m.DtcEdnsPreferClientSubnet),
		DtcScheduledBackup:                  ExpandGridDnsDtcScheduledBackup(ctx, m.DtcScheduledBackup, diags),
		DtcTopologyEaList:                   flex.ExpandFrameworkListString(ctx, m.DtcTopologyEaList, diags),
		EdnsUdpSize:                         flex.ExpandInt64Pointer(m.EdnsUdpSize),
		Email:                               flex.ExpandStringPointer(m.Email),
		EnableBlackhole:                     flex.ExpandBoolPointer(m.EnableBlackhole),
		EnableBlacklist:                     flex.ExpandBoolPointer(m.EnableBlacklist),
		EnableCaptureDnsQueries:             flex.ExpandBoolPointer(m.EnableCaptureDnsQueries),
		EnableCaptureDnsResponses:           flex.ExpandBoolPointer(m.EnableCaptureDnsResponses),
		EnableClientSubnetForwarding:        flex.ExpandBoolPointer(m.EnableClientSubnetForwarding),
		EnableClientSubnetRecursive:         flex.ExpandBoolPointer(m.EnableClientSubnetRecursive),
		EnableDeleteAssociatedPtr:           flex.ExpandBoolPointer(m.EnableDeleteAssociatedPtr),
		EnableDns64:                         flex.ExpandBoolPointer(m.EnableDns64),
		EnableDnsHealthCheck:                flex.ExpandBoolPointer(m.EnableDnsHealthCheck),
		EnableDnstapQueries:                 flex.ExpandBoolPointer(m.EnableDnstapQueries),
		EnableDnstapResponses:               flex.ExpandBoolPointer(m.EnableDnstapResponses),
		EnableDnstapViolationsTls:           flex.ExpandBoolPointer(m.EnableDnstapViolationsTls),
		EnableExcludedDomainNames:           flex.ExpandBoolPointer(m.EnableExcludedDomainNames),
		EnableFixedRrsetOrderFqdns:          flex.ExpandBoolPointer(m.EnableFixedRrsetOrderFqdns),
		EnableFtc:                           flex.ExpandBoolPointer(m.EnableFtc),
		EnableGssTsig:                       flex.ExpandBoolPointer(m.EnableGssTsig),
		EnableHostRrsetOrder:                flex.ExpandBoolPointer(m.EnableHostRrsetOrder),
		EnableHsmSigning:                    flex.ExpandBoolPointer(m.EnableHsmSigning),
		EnableNotifySourcePort:              flex.ExpandBoolPointer(m.EnableNotifySourcePort),
		EnableQueryRewrite:                  flex.ExpandBoolPointer(m.EnableQueryRewrite),
		EnableQuerySourcePort:               flex.ExpandBoolPointer(m.EnableQuerySourcePort),
		ExcludedDomainNames:                 flex.ExpandFrameworkListString(ctx, m.ExcludedDomainNames, diags),
		ExpireAfter:                         flex.ExpandInt64Pointer(m.ExpireAfter),
		FileTransferSetting:                 ExpandGridDnsFileTransferSetting(ctx, m.FileTransferSetting, diags),
		FilterAaaa:                          flex.ExpandStringPointer(m.FilterAaaa),
		FilterAaaaList:                      flex.ExpandFrameworkListNestedBlock(ctx, m.FilterAaaaList, diags, ExpandGridDnsFilterAaaaList),
		FixedRrsetOrderFqdns:                flex.ExpandFrameworkListNestedBlock(ctx, m.FixedRrsetOrderFqdns, diags, ExpandGridDnsFixedRrsetOrderFqdns),
		ForwardOnly:                         flex.ExpandBoolPointer(m.ForwardOnly),
		ForwardUpdates:                      flex.ExpandBoolPointer(m.ForwardUpdates),
		Forwarders:                          flex.ExpandFrameworkListString(ctx, m.Forwarders, diags),
		FtcExpiredRecordTimeout:             flex.ExpandInt64Pointer(m.FtcExpiredRecordTimeout),
		FtcExpiredRecordTtl:                 flex.ExpandInt64Pointer(m.FtcExpiredRecordTtl),
		GenEadbFromHosts:                    flex.ExpandBoolPointer(m.GenEadbFromHosts),
		GenEadbFromNetworkContainers:        flex.ExpandBoolPointer(m.GenEadbFromNetworkContainers),
		GenEadbFromNetworks:                 flex.ExpandBoolPointer(m.GenEadbFromNetworks),
		GenEadbFromRanges:                   flex.ExpandBoolPointer(m.GenEadbFromRanges),
		GssTsigKeys:                         flex.ExpandFrameworkListString(ctx, m.GssTsigKeys, diags),
		LastQueriedAcl:                      flex.ExpandFrameworkListNestedBlock(ctx, m.LastQueriedAcl, diags, ExpandGridDnsLastQueriedAcl),
		LoggingCategories:                   ExpandGridDnsLoggingCategories(ctx, m.LoggingCategories, diags),
		MaxCacheTtl:                         flex.ExpandInt64Pointer(m.MaxCacheTtl),
		MaxCachedLifetime:                   flex.ExpandInt64Pointer(m.MaxCachedLifetime),
		MaxNcacheTtl:                        flex.ExpandInt64Pointer(m.MaxNcacheTtl),
		MaxUdpSize:                          flex.ExpandInt64Pointer(m.MaxUdpSize),
		MemberSecondaryNotify:               flex.ExpandBoolPointer(m.MemberSecondaryNotify),
		NegativeTtl:                         flex.ExpandInt64Pointer(m.NegativeTtl),
		NotifyDelay:                         flex.ExpandInt64Pointer(m.NotifyDelay),
		NotifySourcePort:                    flex.ExpandInt64Pointer(m.NotifySourcePort),
		NsgroupDefault:                      flex.ExpandStringPointer(m.NsgroupDefault),
		Nsgroups:                            flex.ExpandFrameworkListString(ctx, m.Nsgroups, diags),
		NxdomainLogQuery:                    flex.ExpandBoolPointer(m.NxdomainLogQuery),
		NxdomainRedirect:                    flex.ExpandBoolPointer(m.NxdomainRedirect),
		NxdomainRedirectAddresses:           flex.ExpandFrameworkListString(ctx, m.NxdomainRedirectAddresses, diags),
		NxdomainRedirectAddressesV6:         flex.ExpandFrameworkListString(ctx, m.NxdomainRedirectAddressesV6, diags),
		NxdomainRedirectTtl:                 flex.ExpandInt64Pointer(m.NxdomainRedirectTtl),
		NxdomainRulesets:                    flex.ExpandFrameworkListString(ctx, m.NxdomainRulesets, diags),
		PreserveHostRrsetOrderOnSecondaries: flex.ExpandBoolPointer(m.PreserveHostRrsetOrderOnSecondaries),
		ProtocolRecordNamePolicies:          flex.ExpandFrameworkListString(ctx, m.ProtocolRecordNamePolicies, diags),
		QueryRewriteDomainNames:             flex.ExpandFrameworkListString(ctx, m.QueryRewriteDomainNames, diags),
		QueryRewritePrefix:                  flex.ExpandStringPointer(m.QueryRewritePrefix),
		QuerySourcePort:                     flex.ExpandInt64Pointer(m.QuerySourcePort),
		RecursiveQueryList:                  flex.ExpandFrameworkListNestedBlock(ctx, m.RecursiveQueryList, diags, ExpandGridDnsRecursiveQueryList),
		RefreshTimer:                        flex.ExpandInt64Pointer(m.RefreshTimer),
		ResolverQueryTimeout:                flex.ExpandInt64Pointer(m.ResolverQueryTimeout),
		ResponseRateLimiting:                ExpandGridDnsResponseRateLimiting(ctx, m.ResponseRateLimiting, diags),
		RestartSetting:                      ExpandGridDnsRestartSetting(ctx, m.RestartSetting, diags),
		RetryTimer:                          flex.ExpandInt64Pointer(m.RetryTimer),
		RootNameServerType:                  flex.ExpandStringPointer(m.RootNameServerType),
		RpzDisableNsdnameNsip:               flex.ExpandBoolPointer(m.RpzDisableNsdnameNsip),
		RpzDropIpRuleEnabled:                flex.ExpandBoolPointer(m.RpzDropIpRuleEnabled),
		RpzDropIpRuleMinPrefixLengthIpv4:    flex.ExpandInt64Pointer(m.RpzDropIpRuleMinPrefixLengthIpv4),
		RpzDropIpRuleMinPrefixLengthIpv6:    flex.ExpandInt64Pointer(m.RpzDropIpRuleMinPrefixLengthIpv6),
		RpzQnameWaitRecurse:                 flex.ExpandBoolPointer(m.RpzQnameWaitRecurse),
		ScavengingSettings:                  ExpandGridDnsScavengingSettings(ctx, m.ScavengingSettings, diags),
		SerialQueryRate:                     flex.ExpandInt64Pointer(m.SerialQueryRate),
		ServerIdDirective:                   flex.ExpandStringPointer(m.ServerIdDirective),
		Sortlist:                            flex.ExpandFrameworkListNestedBlock(ctx, m.Sortlist, diags, ExpandGridDnsSortlist),
		StoreLocally:                        flex.ExpandBoolPointer(m.StoreLocally),
		SyslogFacility:                      flex.ExpandStringPointer(m.SyslogFacility),
		TransferExcludedServers:             flex.ExpandFrameworkListString(ctx, m.TransferExcludedServers, diags),
		TransferFormat:                      flex.ExpandStringPointer(m.TransferFormat),
		TransfersIn:                         flex.ExpandInt64Pointer(m.TransfersIn),
		TransfersOut:                        flex.ExpandInt64Pointer(m.TransfersOut),
		TransfersPerNs:                      flex.ExpandInt64Pointer(m.TransfersPerNs),
		ZoneDeletionDoubleConfirm:           flex.ExpandBoolPointer(m.ZoneDeletionDoubleConfirm),
	}
	return to
}

func FlattenGridDns(ctx context.Context, from *grid.GridDns, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridDnsAttrTypes)
	}
	m := GridDnsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridDnsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridDnsModel) Flatten(ctx context.Context, from *grid.GridDns, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridDnsModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.AddClientIpMacOptions = types.BoolPointerValue(from.AddClientIpMacOptions)
	m.AllowBulkhostDdns = flex.FlattenStringPointer(from.AllowBulkhostDdns)
	m.AllowGssTsigZoneUpdates = types.BoolPointerValue(from.AllowGssTsigZoneUpdates)
	m.AllowQuery = flex.FlattenFrameworkListNestedBlock(ctx, from.AllowQuery, GridDnsAllowQueryAttrTypes, diags, FlattenGridDnsAllowQuery)
	m.AllowRecursiveQuery = types.BoolPointerValue(from.AllowRecursiveQuery)
	m.AllowTransfer = flex.FlattenFrameworkListNestedBlock(ctx, from.AllowTransfer, GridDnsAllowTransferAttrTypes, diags, FlattenGridDnsAllowTransfer)
	m.AllowUpdate = flex.FlattenFrameworkListNestedBlock(ctx, from.AllowUpdate, GridDnsAllowUpdateAttrTypes, diags, FlattenGridDnsAllowUpdate)
	m.AnonymizeResponseLogging = types.BoolPointerValue(from.AnonymizeResponseLogging)
	m.AttackMitigation = FlattenGridDnsAttackMitigation(ctx, from.AttackMitigation, diags)
	m.AutoBlackhole = FlattenGridDnsAutoBlackhole(ctx, from.AutoBlackhole, diags)
	m.BindCheckNamesPolicy = flex.FlattenStringPointer(from.BindCheckNamesPolicy)
	m.BindHostnameDirective = flex.FlattenStringPointer(from.BindHostnameDirective)
	m.BlackholeList = flex.FlattenFrameworkListNestedBlock(ctx, from.BlackholeList, GridDnsBlackholeListAttrTypes, diags, FlattenGridDnsBlackholeList)
	m.BlacklistAction = flex.FlattenStringPointer(from.BlacklistAction)
	m.BlacklistLogQuery = types.BoolPointerValue(from.BlacklistLogQuery)
	m.BlacklistRedirectAddresses = flex.FlattenFrameworkListString(ctx, from.BlacklistRedirectAddresses, diags)
	m.BlacklistRedirectTtl = flex.FlattenInt64Pointer(from.BlacklistRedirectTtl)
	m.BlacklistRulesets = flex.FlattenFrameworkListString(ctx, from.BlacklistRulesets, diags)
	m.BulkHostNameTemplates = flex.FlattenFrameworkListString(ctx, from.BulkHostNameTemplates, diags)
	m.CaptureDnsQueriesOnAllDomains = types.BoolPointerValue(from.CaptureDnsQueriesOnAllDomains)
	m.CheckNamesForDdnsAndZoneTransfer = types.BoolPointerValue(from.CheckNamesForDdnsAndZoneTransfer)
	m.ClientSubnetDomains = flex.FlattenFrameworkListNestedBlock(ctx, from.ClientSubnetDomains, GridDnsClientSubnetDomainsAttrTypes, diags, FlattenGridDnsClientSubnetDomains)
	m.ClientSubnetIpv4PrefixLength = flex.FlattenInt64Pointer(from.ClientSubnetIpv4PrefixLength)
	m.ClientSubnetIpv6PrefixLength = flex.FlattenInt64Pointer(from.ClientSubnetIpv6PrefixLength)
	m.CopyClientIpMacOptions = types.BoolPointerValue(from.CopyClientIpMacOptions)
	m.CopyXferToNotify = types.BoolPointerValue(from.CopyXferToNotify)
	m.CustomRootNameServers = flex.FlattenFrameworkListNestedBlock(ctx, from.CustomRootNameServers, GridDnsCustomRootNameServersAttrTypes, diags, FlattenGridDnsCustomRootNameServers)
	m.DdnsForceCreationTimestampUpdate = types.BoolPointerValue(from.DdnsForceCreationTimestampUpdate)
	m.DdnsPrincipalGroup = flex.FlattenStringPointer(from.DdnsPrincipalGroup)
	m.DdnsPrincipalTracking = types.BoolPointerValue(from.DdnsPrincipalTracking)
	m.DdnsRestrictPatterns = types.BoolPointerValue(from.DdnsRestrictPatterns)
	m.DdnsRestrictPatternsList = flex.FlattenFrameworkListString(ctx, from.DdnsRestrictPatternsList, diags)
	m.DdnsRestrictProtected = types.BoolPointerValue(from.DdnsRestrictProtected)
	m.DdnsRestrictSecure = types.BoolPointerValue(from.DdnsRestrictSecure)
	m.DdnsRestrictStatic = types.BoolPointerValue(from.DdnsRestrictStatic)
	m.DefaultBulkHostNameTemplate = flex.FlattenStringPointer(from.DefaultBulkHostNameTemplate)
	m.DefaultTtl = flex.FlattenInt64Pointer(from.DefaultTtl)
	m.DisableEdns = types.BoolPointerValue(from.DisableEdns)
	m.Dns64Groups = flex.FlattenFrameworkListString(ctx, from.Dns64Groups, diags)
	m.DnsCacheAccelerationTtl = flex.FlattenInt64Pointer(from.DnsCacheAccelerationTtl)
	m.DnsHealthCheckAnycastControl = types.BoolPointerValue(from.DnsHealthCheckAnycastControl)
	m.DnsHealthCheckDomainList = flex.FlattenFrameworkListString(ctx, from.DnsHealthCheckDomainList, diags)
	m.DnsHealthCheckInterval = flex.FlattenInt64Pointer(from.DnsHealthCheckInterval)
	m.DnsHealthCheckRecursionFlag = types.BoolPointerValue(from.DnsHealthCheckRecursionFlag)
	m.DnsHealthCheckRetries = flex.FlattenInt64Pointer(from.DnsHealthCheckRetries)
	m.DnsHealthCheckTimeout = flex.FlattenInt64Pointer(from.DnsHealthCheckTimeout)
	m.DnsQueryCaptureFileTimeLimit = flex.FlattenInt64Pointer(from.DnsQueryCaptureFileTimeLimit)
	m.DnssecBlacklistEnabled = types.BoolPointerValue(from.DnssecBlacklistEnabled)
	m.DnssecDns64Enabled = types.BoolPointerValue(from.DnssecDns64Enabled)
	m.DnssecEnabled = types.BoolPointerValue(from.DnssecEnabled)
	m.DnssecExpiredSignaturesEnabled = types.BoolPointerValue(from.DnssecExpiredSignaturesEnabled)
	m.DnssecKeyParams = FlattenGridDnsDnssecKeyParams(ctx, from.DnssecKeyParams, diags)
	m.DnssecNegativeTrustAnchors = flex.FlattenFrameworkListString(ctx, from.DnssecNegativeTrustAnchors, diags)
	m.DnssecNxdomainEnabled = types.BoolPointerValue(from.DnssecNxdomainEnabled)
	m.DnssecRpzEnabled = types.BoolPointerValue(from.DnssecRpzEnabled)
	m.DnssecTrustedKeys = flex.FlattenFrameworkListNestedBlock(ctx, from.DnssecTrustedKeys, GridDnsDnssecTrustedKeysAttrTypes, diags, FlattenGridDnsDnssecTrustedKeys)
	m.DnssecValidationEnabled = types.BoolPointerValue(from.DnssecValidationEnabled)
	m.DnstapSetting = FlattenGridDnsDnstapSetting(ctx, from.DnstapSetting, diags)
	m.DomainsToCaptureDnsQueries = flex.FlattenFrameworkListString(ctx, from.DomainsToCaptureDnsQueries, diags)
	m.DtcDnsQueriesSpecificBehavior = flex.FlattenStringPointer(from.DtcDnsQueriesSpecificBehavior)
	m.DtcDnssecMode = flex.FlattenStringPointer(from.DtcDnssecMode)
	m.DtcEdnsPreferClientSubnet = types.BoolPointerValue(from.DtcEdnsPreferClientSubnet)
	m.DtcScheduledBackup = FlattenGridDnsDtcScheduledBackup(ctx, from.DtcScheduledBackup, diags)
	m.DtcTopologyEaList = flex.FlattenFrameworkListString(ctx, from.DtcTopologyEaList, diags)
	m.EdnsUdpSize = flex.FlattenInt64Pointer(from.EdnsUdpSize)
	m.Email = flex.FlattenStringPointer(from.Email)
	m.EnableBlackhole = types.BoolPointerValue(from.EnableBlackhole)
	m.EnableBlacklist = types.BoolPointerValue(from.EnableBlacklist)
	m.EnableCaptureDnsQueries = types.BoolPointerValue(from.EnableCaptureDnsQueries)
	m.EnableCaptureDnsResponses = types.BoolPointerValue(from.EnableCaptureDnsResponses)
	m.EnableClientSubnetForwarding = types.BoolPointerValue(from.EnableClientSubnetForwarding)
	m.EnableClientSubnetRecursive = types.BoolPointerValue(from.EnableClientSubnetRecursive)
	m.EnableDeleteAssociatedPtr = types.BoolPointerValue(from.EnableDeleteAssociatedPtr)
	m.EnableDns64 = types.BoolPointerValue(from.EnableDns64)
	m.EnableDnsHealthCheck = types.BoolPointerValue(from.EnableDnsHealthCheck)
	m.EnableDnstapQueries = types.BoolPointerValue(from.EnableDnstapQueries)
	m.EnableDnstapResponses = types.BoolPointerValue(from.EnableDnstapResponses)
	m.EnableDnstapViolationsTls = types.BoolPointerValue(from.EnableDnstapViolationsTls)
	m.EnableExcludedDomainNames = types.BoolPointerValue(from.EnableExcludedDomainNames)
	m.EnableFixedRrsetOrderFqdns = types.BoolPointerValue(from.EnableFixedRrsetOrderFqdns)
	m.EnableFtc = types.BoolPointerValue(from.EnableFtc)
	m.EnableGssTsig = types.BoolPointerValue(from.EnableGssTsig)
	m.EnableHostRrsetOrder = types.BoolPointerValue(from.EnableHostRrsetOrder)
	m.EnableHsmSigning = types.BoolPointerValue(from.EnableHsmSigning)
	m.EnableNotifySourcePort = types.BoolPointerValue(from.EnableNotifySourcePort)
	m.EnableQueryRewrite = types.BoolPointerValue(from.EnableQueryRewrite)
	m.EnableQuerySourcePort = types.BoolPointerValue(from.EnableQuerySourcePort)
	m.ExcludedDomainNames = flex.FlattenFrameworkListString(ctx, from.ExcludedDomainNames, diags)
	m.ExpireAfter = flex.FlattenInt64Pointer(from.ExpireAfter)
	m.FileTransferSetting = FlattenGridDnsFileTransferSetting(ctx, from.FileTransferSetting, diags)
	m.FilterAaaa = flex.FlattenStringPointer(from.FilterAaaa)
	m.FilterAaaaList = flex.FlattenFrameworkListNestedBlock(ctx, from.FilterAaaaList, GridDnsFilterAaaaListAttrTypes, diags, FlattenGridDnsFilterAaaaList)
	m.FixedRrsetOrderFqdns = flex.FlattenFrameworkListNestedBlock(ctx, from.FixedRrsetOrderFqdns, GridDnsFixedRrsetOrderFqdnsAttrTypes, diags, FlattenGridDnsFixedRrsetOrderFqdns)
	m.ForwardOnly = types.BoolPointerValue(from.ForwardOnly)
	m.ForwardUpdates = types.BoolPointerValue(from.ForwardUpdates)
	m.Forwarders = flex.FlattenFrameworkListString(ctx, from.Forwarders, diags)
	m.FtcExpiredRecordTimeout = flex.FlattenInt64Pointer(from.FtcExpiredRecordTimeout)
	m.FtcExpiredRecordTtl = flex.FlattenInt64Pointer(from.FtcExpiredRecordTtl)
	m.GenEadbFromHosts = types.BoolPointerValue(from.GenEadbFromHosts)
	m.GenEadbFromNetworkContainers = types.BoolPointerValue(from.GenEadbFromNetworkContainers)
	m.GenEadbFromNetworks = types.BoolPointerValue(from.GenEadbFromNetworks)
	m.GenEadbFromRanges = types.BoolPointerValue(from.GenEadbFromRanges)
	m.GssTsigKeys = flex.FlattenFrameworkListString(ctx, from.GssTsigKeys, diags)
	m.LastQueriedAcl = flex.FlattenFrameworkListNestedBlock(ctx, from.LastQueriedAcl, GridDnsLastQueriedAclAttrTypes, diags, FlattenGridDnsLastQueriedAcl)
	m.LoggingCategories = FlattenGridDnsLoggingCategories(ctx, from.LoggingCategories, diags)
	m.MaxCacheTtl = flex.FlattenInt64Pointer(from.MaxCacheTtl)
	m.MaxCachedLifetime = flex.FlattenInt64Pointer(from.MaxCachedLifetime)
	m.MaxNcacheTtl = flex.FlattenInt64Pointer(from.MaxNcacheTtl)
	m.MaxUdpSize = flex.FlattenInt64Pointer(from.MaxUdpSize)
	m.MemberSecondaryNotify = types.BoolPointerValue(from.MemberSecondaryNotify)
	m.NegativeTtl = flex.FlattenInt64Pointer(from.NegativeTtl)
	m.NotifyDelay = flex.FlattenInt64Pointer(from.NotifyDelay)
	m.NotifySourcePort = flex.FlattenInt64Pointer(from.NotifySourcePort)
	m.NsgroupDefault = flex.FlattenStringPointer(from.NsgroupDefault)
	m.Nsgroups = flex.FlattenFrameworkListString(ctx, from.Nsgroups, diags)
	m.NxdomainLogQuery = types.BoolPointerValue(from.NxdomainLogQuery)
	m.NxdomainRedirect = types.BoolPointerValue(from.NxdomainRedirect)
	m.NxdomainRedirectAddresses = flex.FlattenFrameworkListString(ctx, from.NxdomainRedirectAddresses, diags)
	m.NxdomainRedirectAddressesV6 = flex.FlattenFrameworkListString(ctx, from.NxdomainRedirectAddressesV6, diags)
	m.NxdomainRedirectTtl = flex.FlattenInt64Pointer(from.NxdomainRedirectTtl)
	m.NxdomainRulesets = flex.FlattenFrameworkListString(ctx, from.NxdomainRulesets, diags)
	m.PreserveHostRrsetOrderOnSecondaries = types.BoolPointerValue(from.PreserveHostRrsetOrderOnSecondaries)
	m.ProtocolRecordNamePolicies = flex.FlattenFrameworkListString(ctx, from.ProtocolRecordNamePolicies, diags)
	m.QueryRewriteDomainNames = flex.FlattenFrameworkListString(ctx, from.QueryRewriteDomainNames, diags)
	m.QueryRewritePrefix = flex.FlattenStringPointer(from.QueryRewritePrefix)
	m.QuerySourcePort = flex.FlattenInt64Pointer(from.QuerySourcePort)
	m.RecursiveQueryList = flex.FlattenFrameworkListNestedBlock(ctx, from.RecursiveQueryList, GridDnsRecursiveQueryListAttrTypes, diags, FlattenGridDnsRecursiveQueryList)
	m.RefreshTimer = flex.FlattenInt64Pointer(from.RefreshTimer)
	m.ResolverQueryTimeout = flex.FlattenInt64Pointer(from.ResolverQueryTimeout)
	m.ResponseRateLimiting = FlattenGridDnsResponseRateLimiting(ctx, from.ResponseRateLimiting, diags)
	m.RestartSetting = FlattenGridDnsRestartSetting(ctx, from.RestartSetting, diags)
	m.RetryTimer = flex.FlattenInt64Pointer(from.RetryTimer)
	m.RootNameServerType = flex.FlattenStringPointer(from.RootNameServerType)
	m.RpzDisableNsdnameNsip = types.BoolPointerValue(from.RpzDisableNsdnameNsip)
	m.RpzDropIpRuleEnabled = types.BoolPointerValue(from.RpzDropIpRuleEnabled)
	m.RpzDropIpRuleMinPrefixLengthIpv4 = flex.FlattenInt64Pointer(from.RpzDropIpRuleMinPrefixLengthIpv4)
	m.RpzDropIpRuleMinPrefixLengthIpv6 = flex.FlattenInt64Pointer(from.RpzDropIpRuleMinPrefixLengthIpv6)
	m.RpzQnameWaitRecurse = types.BoolPointerValue(from.RpzQnameWaitRecurse)
	m.ScavengingSettings = FlattenGridDnsScavengingSettings(ctx, from.ScavengingSettings, diags)
	m.SerialQueryRate = flex.FlattenInt64Pointer(from.SerialQueryRate)
	m.ServerIdDirective = flex.FlattenStringPointer(from.ServerIdDirective)
	m.Sortlist = flex.FlattenFrameworkListNestedBlock(ctx, from.Sortlist, GridDnsSortlistAttrTypes, diags, FlattenGridDnsSortlist)
	m.StoreLocally = types.BoolPointerValue(from.StoreLocally)
	m.SyslogFacility = flex.FlattenStringPointer(from.SyslogFacility)
	m.TransferExcludedServers = flex.FlattenFrameworkListString(ctx, from.TransferExcludedServers, diags)
	m.TransferFormat = flex.FlattenStringPointer(from.TransferFormat)
	m.TransfersIn = flex.FlattenInt64Pointer(from.TransfersIn)
	m.TransfersOut = flex.FlattenInt64Pointer(from.TransfersOut)
	m.TransfersPerNs = flex.FlattenInt64Pointer(from.TransfersPerNs)
	m.ZoneDeletionDoubleConfirm = types.BoolPointerValue(from.ZoneDeletionDoubleConfirm)
}
