package dns

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/boolvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type ViewModel struct {
	Ref                                 types.String `tfsdk:"ref"`
	BlacklistAction                     types.String `tfsdk:"blacklist_action"`
	BlacklistLogQuery                   types.Bool   `tfsdk:"blacklist_log_query"`
	BlacklistRedirectAddresses          types.List   `tfsdk:"blacklist_redirect_addresses"`
	BlacklistRedirectTtl                types.Int64  `tfsdk:"blacklist_redirect_ttl"`
	BlacklistRulesets                   types.List   `tfsdk:"blacklist_rulesets"`
	CloudInfo                           types.Object `tfsdk:"cloud_info"`
	Comment                             types.String `tfsdk:"comment"`
	CustomRootNameServers               types.List   `tfsdk:"custom_root_name_servers"`
	DdnsForceCreationTimestampUpdate    types.Bool   `tfsdk:"ddns_force_creation_timestamp_update"`
	DdnsPrincipalGroup                  types.String `tfsdk:"ddns_principal_group"`
	DdnsPrincipalTracking               types.Bool   `tfsdk:"ddns_principal_tracking"`
	DdnsRestrictPatterns                types.Bool   `tfsdk:"ddns_restrict_patterns"`
	DdnsRestrictPatternsList            types.List   `tfsdk:"ddns_restrict_patterns_list"`
	DdnsRestrictProtected               types.Bool   `tfsdk:"ddns_restrict_protected"`
	DdnsRestrictSecure                  types.Bool   `tfsdk:"ddns_restrict_secure"`
	DdnsRestrictStatic                  types.Bool   `tfsdk:"ddns_restrict_static"`
	Disable                             types.Bool   `tfsdk:"disable"`
	Dns64Enabled                        types.Bool   `tfsdk:"dns64_enabled"`
	Dns64Groups                         types.List   `tfsdk:"dns64_groups"`
	DnssecEnabled                       types.Bool   `tfsdk:"dnssec_enabled"`
	DnssecExpiredSignaturesEnabled      types.Bool   `tfsdk:"dnssec_expired_signatures_enabled"`
	DnssecNegativeTrustAnchors          types.List   `tfsdk:"dnssec_negative_trust_anchors"`
	DnssecTrustedKeys                   types.List   `tfsdk:"dnssec_trusted_keys"`
	DnssecValidationEnabled             types.Bool   `tfsdk:"dnssec_validation_enabled"`
	EdnsUdpSize                         types.Int64  `tfsdk:"edns_udp_size"`
	EnableBlacklist                     types.Bool   `tfsdk:"enable_blacklist"`
	EnableFixedRrsetOrderFqdns          types.Bool   `tfsdk:"enable_fixed_rrset_order_fqdns"`
	EnableMatchRecursiveOnly            types.Bool   `tfsdk:"enable_match_recursive_only"`
	ExtAttrs                            types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll                         types.Map    `tfsdk:"extattrs_all"`
	FilterAaaa                          types.String `tfsdk:"filter_aaaa"`
	FilterAaaaList                      types.List   `tfsdk:"filter_aaaa_list"`
	FixedRrsetOrderFqdns                types.List   `tfsdk:"fixed_rrset_order_fqdns"`
	ForwardOnly                         types.Bool   `tfsdk:"forward_only"`
	Forwarders                          types.List   `tfsdk:"forwarders"`
	IsDefault                           types.Bool   `tfsdk:"is_default"`
	LastQueriedAcl                      types.List   `tfsdk:"last_queried_acl"`
	MatchClients                        types.List   `tfsdk:"match_clients"`
	MatchDestinations                   types.List   `tfsdk:"match_destinations"`
	MaxCacheTtl                         types.Int64  `tfsdk:"max_cache_ttl"`
	MaxNcacheTtl                        types.Int64  `tfsdk:"max_ncache_ttl"`
	MaxUdpSize                          types.Int64  `tfsdk:"max_udp_size"`
	Name                                types.String `tfsdk:"name"`
	NetworkView                         types.String `tfsdk:"network_view"`
	NotifyDelay                         types.Int64  `tfsdk:"notify_delay"`
	NxdomainLogQuery                    types.Bool   `tfsdk:"nxdomain_log_query"`
	NxdomainRedirect                    types.Bool   `tfsdk:"nxdomain_redirect"`
	NxdomainRedirectAddresses           types.List   `tfsdk:"nxdomain_redirect_addresses"`
	NxdomainRedirectAddressesV6         types.List   `tfsdk:"nxdomain_redirect_addresses_v6"`
	NxdomainRedirectTtl                 types.Int64  `tfsdk:"nxdomain_redirect_ttl"`
	NxdomainRulesets                    types.List   `tfsdk:"nxdomain_rulesets"`
	Recursion                           types.Bool   `tfsdk:"recursion"`
	ResponseRateLimiting                types.Object `tfsdk:"response_rate_limiting"`
	RootNameServerType                  types.String `tfsdk:"root_name_server_type"`
	RpzDropIpRuleEnabled                types.Bool   `tfsdk:"rpz_drop_ip_rule_enabled"`
	RpzDropIpRuleMinPrefixLengthIpv4    types.Int64  `tfsdk:"rpz_drop_ip_rule_min_prefix_length_ipv4"`
	RpzDropIpRuleMinPrefixLengthIpv6    types.Int64  `tfsdk:"rpz_drop_ip_rule_min_prefix_length_ipv6"`
	RpzQnameWaitRecurse                 types.Bool   `tfsdk:"rpz_qname_wait_recurse"`
	ScavengingSettings                  types.Object `tfsdk:"scavenging_settings"`
	Sortlist                            types.List   `tfsdk:"sortlist"`
	UseBlacklist                        types.Bool   `tfsdk:"use_blacklist"`
	UseDdnsForceCreationTimestampUpdate types.Bool   `tfsdk:"use_ddns_force_creation_timestamp_update"`
	UseDdnsPatternsRestriction          types.Bool   `tfsdk:"use_ddns_patterns_restriction"`
	UseDdnsPrincipalSecurity            types.Bool   `tfsdk:"use_ddns_principal_security"`
	UseDdnsRestrictProtected            types.Bool   `tfsdk:"use_ddns_restrict_protected"`
	UseDdnsRestrictStatic               types.Bool   `tfsdk:"use_ddns_restrict_static"`
	UseDns64                            types.Bool   `tfsdk:"use_dns64"`
	UseDnssec                           types.Bool   `tfsdk:"use_dnssec"`
	UseEdnsUdpSize                      types.Bool   `tfsdk:"use_edns_udp_size"`
	UseFilterAaaa                       types.Bool   `tfsdk:"use_filter_aaaa"`
	UseFixedRrsetOrderFqdns             types.Bool   `tfsdk:"use_fixed_rrset_order_fqdns"`
	UseForwarders                       types.Bool   `tfsdk:"use_forwarders"`
	UseMaxCacheTtl                      types.Bool   `tfsdk:"use_max_cache_ttl"`
	UseMaxNcacheTtl                     types.Bool   `tfsdk:"use_max_ncache_ttl"`
	UseMaxUdpSize                       types.Bool   `tfsdk:"use_max_udp_size"`
	UseNxdomainRedirect                 types.Bool   `tfsdk:"use_nxdomain_redirect"`
	UseRecursion                        types.Bool   `tfsdk:"use_recursion"`
	UseResponseRateLimiting             types.Bool   `tfsdk:"use_response_rate_limiting"`
	UseRootNameServer                   types.Bool   `tfsdk:"use_root_name_server"`
	UseRpzDropIpRule                    types.Bool   `tfsdk:"use_rpz_drop_ip_rule"`
	UseRpzQnameWaitRecurse              types.Bool   `tfsdk:"use_rpz_qname_wait_recurse"`
	UseScavengingSettings               types.Bool   `tfsdk:"use_scavenging_settings"`
	UseSortlist                         types.Bool   `tfsdk:"use_sortlist"`
}

var ViewAttrTypes = map[string]attr.Type{
	"ref":                                      types.StringType,
	"blacklist_action":                         types.StringType,
	"blacklist_log_query":                      types.BoolType,
	"blacklist_redirect_addresses":             types.ListType{ElemType: types.StringType},
	"blacklist_redirect_ttl":                   types.Int64Type,
	"blacklist_rulesets":                       types.ListType{ElemType: types.StringType},
	"cloud_info":                               types.ObjectType{AttrTypes: ViewCloudInfoAttrTypes},
	"comment":                                  types.StringType,
	"custom_root_name_servers":                 types.ListType{ElemType: types.ObjectType{AttrTypes: ViewCustomRootNameServersAttrTypes}},
	"ddns_force_creation_timestamp_update":     types.BoolType,
	"ddns_principal_group":                     types.StringType,
	"ddns_principal_tracking":                  types.BoolType,
	"ddns_restrict_patterns":                   types.BoolType,
	"ddns_restrict_patterns_list":              types.ListType{ElemType: types.StringType},
	"ddns_restrict_protected":                  types.BoolType,
	"ddns_restrict_secure":                     types.BoolType,
	"ddns_restrict_static":                     types.BoolType,
	"disable":                                  types.BoolType,
	"dns64_enabled":                            types.BoolType,
	"dns64_groups":                             types.ListType{ElemType: types.StringType},
	"dnssec_enabled":                           types.BoolType,
	"dnssec_expired_signatures_enabled":        types.BoolType,
	"dnssec_negative_trust_anchors":            types.ListType{ElemType: types.StringType},
	"dnssec_trusted_keys":                      types.ListType{ElemType: types.ObjectType{AttrTypes: ViewDnssecTrustedKeysAttrTypes}},
	"dnssec_validation_enabled":                types.BoolType,
	"edns_udp_size":                            types.Int64Type,
	"enable_blacklist":                         types.BoolType,
	"enable_fixed_rrset_order_fqdns":           types.BoolType,
	"enable_match_recursive_only":              types.BoolType,
	"extattrs":                                 types.MapType{ElemType: types.StringType},
	"extattrs_all":                             types.MapType{ElemType: types.StringType},
	"filter_aaaa":                              types.StringType,
	"filter_aaaa_list":                         types.ListType{ElemType: types.ObjectType{AttrTypes: ViewFilterAaaaListAttrTypes}},
	"fixed_rrset_order_fqdns":                  types.ListType{ElemType: types.ObjectType{AttrTypes: ViewFixedRrsetOrderFqdnsAttrTypes}},
	"forward_only":                             types.BoolType,
	"forwarders":                               types.ListType{ElemType: types.StringType},
	"is_default":                               types.BoolType,
	"last_queried_acl":                         types.ListType{ElemType: types.ObjectType{AttrTypes: ViewLastQueriedAclAttrTypes}},
	"match_clients":                            types.ListType{ElemType: types.ObjectType{AttrTypes: ViewMatchClientsAttrTypes}},
	"match_destinations":                       types.ListType{ElemType: types.ObjectType{AttrTypes: ViewMatchDestinationsAttrTypes}},
	"max_cache_ttl":                            types.Int64Type,
	"max_ncache_ttl":                           types.Int64Type,
	"max_udp_size":                             types.Int64Type,
	"name":                                     types.StringType,
	"network_view":                             types.StringType,
	"notify_delay":                             types.Int64Type,
	"nxdomain_log_query":                       types.BoolType,
	"nxdomain_redirect":                        types.BoolType,
	"nxdomain_redirect_addresses":              types.ListType{ElemType: types.StringType},
	"nxdomain_redirect_addresses_v6":           types.ListType{ElemType: types.StringType},
	"nxdomain_redirect_ttl":                    types.Int64Type,
	"nxdomain_rulesets":                        types.ListType{ElemType: types.StringType},
	"recursion":                                types.BoolType,
	"response_rate_limiting":                   types.ObjectType{AttrTypes: ViewResponseRateLimitingAttrTypes},
	"root_name_server_type":                    types.StringType,
	"rpz_drop_ip_rule_enabled":                 types.BoolType,
	"rpz_drop_ip_rule_min_prefix_length_ipv4":  types.Int64Type,
	"rpz_drop_ip_rule_min_prefix_length_ipv6":  types.Int64Type,
	"rpz_qname_wait_recurse":                   types.BoolType,
	"scavenging_settings":                      types.ObjectType{AttrTypes: ViewScavengingSettingsAttrTypes},
	"sortlist":                                 types.ListType{ElemType: types.ObjectType{AttrTypes: ViewSortlistAttrTypes}},
	"use_blacklist":                            types.BoolType,
	"use_ddns_force_creation_timestamp_update": types.BoolType,
	"use_ddns_patterns_restriction":            types.BoolType,
	"use_ddns_principal_security":              types.BoolType,
	"use_ddns_restrict_protected":              types.BoolType,
	"use_ddns_restrict_static":                 types.BoolType,
	"use_dns64":                                types.BoolType,
	"use_dnssec":                               types.BoolType,
	"use_edns_udp_size":                        types.BoolType,
	"use_filter_aaaa":                          types.BoolType,
	"use_fixed_rrset_order_fqdns":              types.BoolType,
	"use_forwarders":                           types.BoolType,
	"use_max_cache_ttl":                        types.BoolType,
	"use_max_ncache_ttl":                       types.BoolType,
	"use_max_udp_size":                         types.BoolType,
	"use_nxdomain_redirect":                    types.BoolType,
	"use_recursion":                            types.BoolType,
	"use_response_rate_limiting":               types.BoolType,
	"use_root_name_server":                     types.BoolType,
	"use_rpz_drop_ip_rule":                     types.BoolType,
	"use_rpz_qname_wait_recurse":               types.BoolType,
	"use_scavenging_settings":                  types.BoolType,
	"use_sortlist":                             types.BoolType,
}

var ViewResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"blacklist_action": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("REDIRECT"),
		Validators: []validator.String{
			stringvalidator.OneOf("REDIRECT", "REFUSE"),
			stringvalidator.AlsoRequires(path.MatchRoot("use_blacklist")),
		},
		MarkdownDescription: "The action to perform when a domain name matches the pattern defined in a rule that is specified by the blacklist_ruleset method. Valid values are \"REDIRECT\" or \"REFUSE\". The default value is \"REFUSE\".",
	},
	"blacklist_log_query": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_blacklist")),
		},
		MarkdownDescription: "The flag that indicates whether blacklist redirection queries are logged. Specify \"true\" to enable logging, or \"false\" to disable it. The default value is \"false\".",
	},
	"blacklist_redirect_addresses": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_blacklist")),
		},
		MarkdownDescription: "The array of IP addresses the appliance includes in the response it sends in place of a blacklisted IP address.",
	},
	"blacklist_redirect_ttl": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(60),
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_blacklist")),
		},
		MarkdownDescription: "The Time To Live (TTL) value of the synthetic DNS responses resulted from blacklist redirection. The TTL value is a 32-bit unsigned integer that represents the TTL in seconds.",
	},
	"blacklist_rulesets": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The name of the Ruleset object assigned at the Grid level for blacklist redirection.",
	},
	"cloud_info": schema.SingleNestedAttribute{
		Attributes:          ViewCloudInfoResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "Structure containing all cloud API related information for this object.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Name should not have leading or trailing whitespace",
			),
		},
		MarkdownDescription: "Comment for the DNS view; maximum 64 characters.",
	},
	"custom_root_name_servers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ViewCustomRootNameServersResourceSchemaAttributes,
		},
		Optional:            true,
		Computed:            true,
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_root_name_server")),
		},
		MarkdownDescription: "The list of customized root name servers. You can either select and use Internet root name servers or specify custom root name servers by providing a host name and IP address to which the Infoblox appliance can send queries. Include the specified parameter to set the attribute value. Omit the parameter to retrieve the attribute value.",
	},
	"ddns_force_creation_timestamp_update": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_ddns_force_creation_timestamp_update")),
		},
		MarkdownDescription: "Defines whether creation timestamp of RR should be updated ' when DDNS update happens even if there is no change to ' the RR.",
	},
	"ddns_principal_group": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_ddns_principal_security")),
		},
		MarkdownDescription: "The DDNS Principal cluster group name.",
	},
	"ddns_principal_tracking": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_ddns_principal_security")),
		},
		MarkdownDescription: "The flag that indicates whether the DDNS principal track is enabled or disabled.",
	},
	"ddns_restrict_patterns": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_ddns_patterns_restriction")),
		},
		MarkdownDescription: "The flag that indicates whether an option to restrict DDNS update request based on FQDN patterns is enabled or disabled.",
	},
	"ddns_restrict_patterns_list": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_ddns_patterns_restriction")),
		},
		MarkdownDescription: "The unordered list of restriction patterns for an option of to restrict DDNS updates based on FQDN patterns.",
	},
	"ddns_restrict_protected": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_ddns_restrict_protected")),
		},
		MarkdownDescription: "The flag that indicates whether an option to restrict DDNS update request to protected resource records is enabled or disabled.",
	},
	"ddns_restrict_secure": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_ddns_principal_security")),
		},
		MarkdownDescription: "The flag that indicates whether DDNS update request for principal other than target resource record's principal is restricted.",
	},
	"ddns_restrict_static": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_ddns_restrict_static")),
		},
		MarkdownDescription: "The flag that indicates whether an option to restrict DDNS update request to resource records which are marked as 'STATIC' is enabled or disabled.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the DNS view is disabled or not. When this is set to False, the DNS view is enabled.",
	},
	"dns64_enabled": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_dns64")),
		},
		MarkdownDescription: "Determines if the DNS64 s enabled or not.",
	},
	"dns64_groups": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_dns64")),
		},
		MarkdownDescription: "The list of DNS64 synthesis groups associated with this DNS view.",
	},
	"dnssec_enabled": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_dnssec")),
		},
		MarkdownDescription: "Determines if the DNS security extension is enabled or not.",
	},
	"dnssec_expired_signatures_enabled": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_dnssec")),
		},
		MarkdownDescription: "Determines if the DNS security extension accepts expired signatures or not.",
	},
	"dnssec_negative_trust_anchors": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "A list of zones for which the server does not perform DNSSEC validation.",
	},
	"dnssec_trusted_keys": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ViewDnssecTrustedKeysResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_dnssec")),
		},
		MarkdownDescription: "The list of trusted keys for the DNS security extension.",
	},
	"dnssec_validation_enabled": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(true),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_dnssec")),
		},
		MarkdownDescription: "Determines if the DNS security validation is enabled or not.",
	},
	"edns_udp_size": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(1220),
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_edns_udp_size")),
		},
		MarkdownDescription: "Advertises the EDNS0 buffer size to the upstream server. The value should be between 512 and 4096 bytes. The recommended value is between 512 and 1220 bytes.",
	},
	"enable_blacklist": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_blacklist")),
		},
		MarkdownDescription: "Determines if the blacklist in a DNS view is enabled or not.",
	},
	"enable_fixed_rrset_order_fqdns": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_fixed_rrset_order_fqdns")),
		},
		MarkdownDescription: "Determines if the fixed RRset order FQDN is enabled or not.",
	},
	"enable_match_recursive_only": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the 'match-recursive-only' option in a DNS view is enabled or not.",
	},
	"extattrs": schema.MapAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object.",
		ElementType:         types.StringType,
		Default:             mapdefault.StaticValue(types.MapNull(types.StringType)),
	},
	"extattrs_all": schema.MapAttribute{
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object , including default attributes.",
		ElementType:         types.StringType,
	},
	"filter_aaaa": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("NO"),
		Validators: []validator.String{
			stringvalidator.OneOf("BREAK_DNSSEC", "NO", "YES"),
			stringvalidator.AlsoRequires(path.MatchRoot("use_filter_aaaa")),
		},
		MarkdownDescription: "The type of AAAA filtering for this DNS view object.",
	},
	"filter_aaaa_list": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ViewFilterAaaaListResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_filter_aaaa")),
		},
		MarkdownDescription: "Applies AAAA filtering to a named ACL, or to a list of IPv4/IPv6 addresses and networks from which queries are received. This field does not allow TSIG keys.",
	},
	"fixed_rrset_order_fqdns": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ViewFixedRrsetOrderFqdnsResourceSchemaAttributes,
		},
		Optional:            true,
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_fixed_rrset_order_fqdns")),
		},
		MarkdownDescription: "The fixed RRset order FQDN. If this field does not contain an empty value, the appliance will automatically set the enable_fixed_rrset_order_fqdns field to 'true', unless the same request sets the enable field to 'false'.",
	},
	"forward_only": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_forwarders")),
		},
		MarkdownDescription: "Determines if this DNS view sends queries to forwarders only or not. When the value is True, queries are sent to forwarders only, and not to other internal or Internet root servers.",
	},
	"forwarders": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_forwarders")),
		},
		MarkdownDescription: "The list of forwarders for the DNS view. A forwarder is a name server to which other name servers first send their off-site queries. The forwarder builds up a cache of information, avoiding the need for other name servers to send queries off-site.",
	},
	"is_default": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "The NIOS appliance provides one default DNS view. You can rename the default view and change its settings, but you cannot delete it. There must always be at least one DNS view in the appliance.",
	},
	"last_queried_acl": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ViewLastQueriedAclResourceSchemaAttributes,
		},
		Optional:            true,
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_scavenging_settings")),
		},
		MarkdownDescription: "Determines last queried ACL for the specified IPv4 or IPv6 addresses and networks in scavenging settings.",
	},
	"match_clients": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ViewMatchClientsResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "A list of forwarders for the match clients. This list specifies a named ACL, or a list of IPv4/IPv6 addresses, networks, TSIG keys of clients that are allowed or denied access to the DNS view.",
	},
	"match_destinations": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ViewMatchDestinationsResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "A list of forwarders for the match destinations. This list specifies a name ACL, or a list of IPv4/IPv6 addresses, networks, TSIG keys of clients that are allowed or denied access to the DNS view.",
	},
	"max_cache_ttl": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(604800),
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_max_cache_ttl")),
		},
		MarkdownDescription: "The maximum number of seconds to cache ordinary (positive) answers.",
	},
	"max_ncache_ttl": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(10800),
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_max_ncache_ttl")),
		},
		MarkdownDescription: "The maximum number of seconds to cache negative (NXDOMAIN) answers.",
	},
	"max_udp_size": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(1220),
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_max_udp_size")),
		},
		MarkdownDescription: "The value is used by authoritative DNS servers to never send DNS responses larger than the configured value. The value should be between 512 and 4096 bytes. The recommended value is between 512 and 1220 bytes.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Name should not have leading or trailing whitespace",
			),
		},
		MarkdownDescription: "Name of the DNS view.",
	},
	"network_view": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString("default"),
		MarkdownDescription: "The name of the network view object associated with this DNS view.",
	},
	"notify_delay": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(5),
		MarkdownDescription: "The number of seconds of delay the notify messages are sent to secondaries.",
	},
	"nxdomain_log_query": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_nxdomain_redirect")),
		},
		MarkdownDescription: "The flag that indicates whether NXDOMAIN redirection queries are logged. Specify \"true\" to enable logging, or \"false\" to disable it. The default value is \"false\".",
	},
	"nxdomain_redirect": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_nxdomain_redirect")),
		},
		MarkdownDescription: "Determines if NXDOMAIN redirection in a DNS view is enabled or not.",
	},
	"nxdomain_redirect_addresses": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_nxdomain_redirect")),
		},
		MarkdownDescription: "The array with IPv4 addresses the appliance includes in the response it sends in place of an NXDOMAIN response.",
	},
	"nxdomain_redirect_addresses_v6": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_nxdomain_redirect")),
		},
		MarkdownDescription: "The array with IPv6 addresses the appliance includes in the response it sends in place of an NXDOMAIN response.",
	},
	"nxdomain_redirect_ttl": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(60),
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_nxdomain_redirect")),
		},
		MarkdownDescription: "The Time To Live (TTL) value of the synthetic DNS responses resulted from NXDOMAIN redirection. The TTL value is a 32-bit unsigned integer that represents the TTL in seconds.",
	},
	"nxdomain_rulesets": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_nxdomain_redirect")),
		},
		MarkdownDescription: "The names of the Ruleset objects assigned at the grid level for NXDOMAIN redirection.",
	},
	"recursion": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_recursion")),
		},
		MarkdownDescription: "Determines if recursion is enabled or not.",
	},
	"response_rate_limiting": schema.SingleNestedAttribute{
		Attributes: ViewResponseRateLimitingResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Validators: []validator.Object{
			objectvalidator.AlsoRequires(path.MatchRoot("use_response_rate_limiting")),
		},
		MarkdownDescription: "The response rate limiting settings for the DNS view. This feature is used to limit the number of responses sent to a client in a given time period.",
	},
	"root_name_server_type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("INTERNET"),
		Validators: []validator.String{
			stringvalidator.OneOf("CUSTOM", "INTERNET"),
			stringvalidator.AlsoRequires(path.MatchRoot("use_root_name_server")),
		},
		MarkdownDescription: "Determines the type of root name servers.",
	},
	"rpz_drop_ip_rule_enabled": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_rpz_drop_ip_rule")),
		},
		MarkdownDescription: "Enables the appliance to ignore RPZ-IP triggers with prefix lengths less than the specified minimum prefix length.",
	},
	"rpz_drop_ip_rule_min_prefix_length_ipv4": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(29),
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_rpz_drop_ip_rule")),
		},
		MarkdownDescription: "The minimum prefix length for IPv4 RPZ-IP triggers. The appliance ignores RPZ-IP triggers with prefix lengths less than the specified minimum IPv4 prefix length.",
	},
	"rpz_drop_ip_rule_min_prefix_length_ipv6": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(112),
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_rpz_drop_ip_rule")),
		},
		MarkdownDescription: "The minimum prefix length for IPv6 RPZ-IP triggers. The appliance ignores RPZ-IP triggers with prefix lengths less than the specified minimum IPv6 prefix length.",
	},
	"rpz_qname_wait_recurse": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_rpz_qname_wait_recurse")),
		},
		MarkdownDescription: "The flag that indicates whether recursive RPZ lookups are enabled.",
	},
	"scavenging_settings": schema.SingleNestedAttribute{
		Attributes: ViewScavengingSettingsResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Validators: []validator.Object{
			objectvalidator.AlsoRequires(path.MatchRoot("use_scavenging_settings")),
		},
		MarkdownDescription: "Scavenging settings for the DNS view",
	},
	"sortlist": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ViewSortlistResourceSchemaAttributes,
		},
		Optional:            true,
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_sortlist")),
		},
		MarkdownDescription: "A sort list that determines the order of IP addresses in responses sent to DNS queries.",
	},
	"use_blacklist": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: blacklist_action , blacklist_log_query, blacklist_redirect_addresses, blacklist_redirect_ttl, blacklist_rulesets, enable_blacklist",
	},
	"use_ddns_force_creation_timestamp_update": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: ddns_force_creation_timestamp_update",
	},
	"use_ddns_patterns_restriction": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: ddns_restrict_patterns_list , ddns_restrict_patterns",
	},
	"use_ddns_principal_security": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: ddns_restrict_secure , ddns_principal_tracking, ddns_principal_group",
	},
	"use_ddns_restrict_protected": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: ddns_restrict_protected",
	},
	"use_ddns_restrict_static": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: ddns_restrict_static",
	},
	"use_dns64": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: dns64_enabled , dns64_groups",
	},
	"use_dnssec": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: dnssec_enabled , dnssec_expired_signatures_enabled, dnssec_validation_enabled, dnssec_trusted_keys",
	},
	"use_edns_udp_size": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: edns_udp_size",
	},
	"use_filter_aaaa": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: filter_aaaa , filter_aaaa_list",
	},
	"use_fixed_rrset_order_fqdns": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: fixed_rrset_order_fqdns , enable_fixed_rrset_order_fqdns",
	},
	"use_forwarders": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: forwarders , forward_only",
	},
	"use_max_cache_ttl": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: max_cache_ttl",
	},
	"use_max_ncache_ttl": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: max_ncache_ttl",
	},
	"use_max_udp_size": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: max_udp_size",
	},
	"use_nxdomain_redirect": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: nxdomain_redirect , nxdomain_redirect_addresses, nxdomain_redirect_addresses_v6, nxdomain_redirect_ttl, nxdomain_log_query, nxdomain_rulesets",
	},
	"use_recursion": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: recursion",
	},
	"use_response_rate_limiting": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: response_rate_limiting",
	},
	"use_root_name_server": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: custom_root_name_servers , root_name_server_type",
	},
	"use_rpz_drop_ip_rule": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: rpz_drop_ip_rule_enabled , rpz_drop_ip_rule_min_prefix_length_ipv4, rpz_drop_ip_rule_min_prefix_length_ipv6",
	},
	"use_rpz_qname_wait_recurse": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: rpz_qname_wait_recurse",
	},
	"use_scavenging_settings": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: scavenging_settings , last_queried_acl",
	},
	"use_sortlist": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: sortlist",
	},
}


func (m *ViewModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.View {
	if m == nil {
		return nil
	}
	to := &dns.View{
		BlacklistAction:                     flex.ExpandStringPointer(m.BlacklistAction),
		BlacklistLogQuery:                   flex.ExpandBoolPointer(m.BlacklistLogQuery),
		BlacklistRedirectAddresses:          flex.ExpandFrameworkListString(ctx, m.BlacklistRedirectAddresses, diags),
		BlacklistRedirectTtl:                flex.ExpandInt64Pointer(m.BlacklistRedirectTtl),
		BlacklistRulesets:                   flex.ExpandFrameworkListString(ctx, m.BlacklistRulesets, diags),
		Comment:                             flex.ExpandStringPointer(m.Comment),
		CustomRootNameServers:               flex.ExpandFrameworkListNestedBlock(ctx, m.CustomRootNameServers, diags, ExpandViewCustomRootNameServers),
		DdnsForceCreationTimestampUpdate:    flex.ExpandBoolPointer(m.DdnsForceCreationTimestampUpdate),
		DdnsPrincipalGroup:                  flex.ExpandStringPointer(m.DdnsPrincipalGroup),
		DdnsPrincipalTracking:               flex.ExpandBoolPointer(m.DdnsPrincipalTracking),
		DdnsRestrictPatterns:                flex.ExpandBoolPointer(m.DdnsRestrictPatterns),
		DdnsRestrictPatternsList:            flex.ExpandFrameworkListString(ctx, m.DdnsRestrictPatternsList, diags),
		DdnsRestrictProtected:               flex.ExpandBoolPointer(m.DdnsRestrictProtected),
		DdnsRestrictSecure:                  flex.ExpandBoolPointer(m.DdnsRestrictSecure),
		DdnsRestrictStatic:                  flex.ExpandBoolPointer(m.DdnsRestrictStatic),
		Disable:                             flex.ExpandBoolPointer(m.Disable),
		Dns64Enabled:                        flex.ExpandBoolPointer(m.Dns64Enabled),
		Dns64Groups:                         flex.ExpandFrameworkListString(ctx, m.Dns64Groups, diags),
		DnssecEnabled:                       flex.ExpandBoolPointer(m.DnssecEnabled),
		DnssecExpiredSignaturesEnabled:      flex.ExpandBoolPointer(m.DnssecExpiredSignaturesEnabled),
		DnssecNegativeTrustAnchors:          flex.ExpandFrameworkListString(ctx, m.DnssecNegativeTrustAnchors, diags),
		DnssecTrustedKeys:                   flex.ExpandFrameworkListNestedBlock(ctx, m.DnssecTrustedKeys, diags, ExpandViewDnssecTrustedKeys),
		DnssecValidationEnabled:             flex.ExpandBoolPointer(m.DnssecValidationEnabled),
		EdnsUdpSize:                         flex.ExpandInt64Pointer(m.EdnsUdpSize),
		EnableBlacklist:                     flex.ExpandBoolPointer(m.EnableBlacklist),
		EnableFixedRrsetOrderFqdns:          flex.ExpandBoolPointer(m.EnableFixedRrsetOrderFqdns),
		EnableMatchRecursiveOnly:            flex.ExpandBoolPointer(m.EnableMatchRecursiveOnly),
		ExtAttrs:                            ExpandExtAttr(ctx, m.ExtAttrs, diags),
		FilterAaaa:                          flex.ExpandStringPointer(m.FilterAaaa),
		FilterAaaaList:                      flex.ExpandFrameworkListNestedBlock(ctx, m.FilterAaaaList, diags, ExpandViewFilterAaaaList),
		FixedRrsetOrderFqdns:                flex.ExpandFrameworkListNestedBlock(ctx, m.FixedRrsetOrderFqdns, diags, ExpandViewFixedRrsetOrderFqdns),
		ForwardOnly:                         flex.ExpandBoolPointer(m.ForwardOnly),
		Forwarders:                          flex.ExpandFrameworkListString(ctx, m.Forwarders, diags),
		LastQueriedAcl:                      flex.ExpandFrameworkListNestedBlock(ctx, m.LastQueriedAcl, diags, ExpandViewLastQueriedAcl),
		MatchClients:                        flex.ExpandFrameworkListNestedBlock(ctx, m.MatchClients, diags, ExpandViewMatchClients),
		MatchDestinations:                   flex.ExpandFrameworkListNestedBlock(ctx, m.MatchDestinations, diags, ExpandViewMatchDestinations),
		MaxCacheTtl:                         flex.ExpandInt64Pointer(m.MaxCacheTtl),
		MaxNcacheTtl:                        flex.ExpandInt64Pointer(m.MaxNcacheTtl),
		MaxUdpSize:                          flex.ExpandInt64Pointer(m.MaxUdpSize),
		Name:                                flex.ExpandStringPointer(m.Name),
		NetworkView:                         flex.ExpandStringPointer(m.NetworkView),
		NotifyDelay:                         flex.ExpandInt64Pointer(m.NotifyDelay),
		NxdomainLogQuery:                    flex.ExpandBoolPointer(m.NxdomainLogQuery),
		NxdomainRedirect:                    flex.ExpandBoolPointer(m.NxdomainRedirect),
		NxdomainRedirectAddresses:           flex.ExpandFrameworkListString(ctx, m.NxdomainRedirectAddresses, diags),
		NxdomainRedirectAddressesV6:         flex.ExpandFrameworkListString(ctx, m.NxdomainRedirectAddressesV6, diags),
		NxdomainRedirectTtl:                 flex.ExpandInt64Pointer(m.NxdomainRedirectTtl),
		NxdomainRulesets:                    flex.ExpandFrameworkListString(ctx, m.NxdomainRulesets, diags),
		Recursion:                           flex.ExpandBoolPointer(m.Recursion),
		ResponseRateLimiting:                ExpandViewResponseRateLimiting(ctx, m.ResponseRateLimiting, diags),
		RootNameServerType:                  flex.ExpandStringPointer(m.RootNameServerType),
		RpzDropIpRuleEnabled:                flex.ExpandBoolPointer(m.RpzDropIpRuleEnabled),
		RpzDropIpRuleMinPrefixLengthIpv4:    flex.ExpandInt64Pointer(m.RpzDropIpRuleMinPrefixLengthIpv4),
		RpzDropIpRuleMinPrefixLengthIpv6:    flex.ExpandInt64Pointer(m.RpzDropIpRuleMinPrefixLengthIpv6),
		RpzQnameWaitRecurse:                 flex.ExpandBoolPointer(m.RpzQnameWaitRecurse),
		ScavengingSettings:                  ExpandViewScavengingSettings(ctx, m.ScavengingSettings, diags),
		Sortlist:                            flex.ExpandFrameworkListNestedBlock(ctx, m.Sortlist, diags, ExpandViewSortlist),
		UseBlacklist:                        flex.ExpandBoolPointer(m.UseBlacklist),
		UseDdnsForceCreationTimestampUpdate: flex.ExpandBoolPointer(m.UseDdnsForceCreationTimestampUpdate),
		UseDdnsPatternsRestriction:          flex.ExpandBoolPointer(m.UseDdnsPatternsRestriction),
		UseDdnsPrincipalSecurity:            flex.ExpandBoolPointer(m.UseDdnsPrincipalSecurity),
		UseDdnsRestrictProtected:            flex.ExpandBoolPointer(m.UseDdnsRestrictProtected),
		UseDdnsRestrictStatic:               flex.ExpandBoolPointer(m.UseDdnsRestrictStatic),
		UseDns64:                            flex.ExpandBoolPointer(m.UseDns64),
		UseDnssec:                           flex.ExpandBoolPointer(m.UseDnssec),
		UseEdnsUdpSize:                      flex.ExpandBoolPointer(m.UseEdnsUdpSize),
		UseFilterAaaa:                       flex.ExpandBoolPointer(m.UseFilterAaaa),
		UseFixedRrsetOrderFqdns:             flex.ExpandBoolPointer(m.UseFixedRrsetOrderFqdns),
		UseForwarders:                       flex.ExpandBoolPointer(m.UseForwarders),
		UseMaxCacheTtl:                      flex.ExpandBoolPointer(m.UseMaxCacheTtl),
		UseMaxNcacheTtl:                     flex.ExpandBoolPointer(m.UseMaxNcacheTtl),
		UseMaxUdpSize:                       flex.ExpandBoolPointer(m.UseMaxUdpSize),
		UseNxdomainRedirect:                 flex.ExpandBoolPointer(m.UseNxdomainRedirect),
		UseRecursion:                        flex.ExpandBoolPointer(m.UseRecursion),
		UseResponseRateLimiting:             flex.ExpandBoolPointer(m.UseResponseRateLimiting),
		UseRootNameServer:                   flex.ExpandBoolPointer(m.UseRootNameServer),
		UseRpzDropIpRule:                    flex.ExpandBoolPointer(m.UseRpzDropIpRule),
		UseRpzQnameWaitRecurse:              flex.ExpandBoolPointer(m.UseRpzQnameWaitRecurse),
		UseScavengingSettings:               flex.ExpandBoolPointer(m.UseScavengingSettings),
		UseSortlist:                         flex.ExpandBoolPointer(m.UseSortlist),
	}
	return to
}

func FlattenView(ctx context.Context, from *dns.View, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ViewAttrTypes)
	}
	m := ViewModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrs = m.ExtAttrsAll
	t, d := types.ObjectValueFrom(ctx, ViewAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ViewModel) Flatten(ctx context.Context, from *dns.View, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ViewModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.BlacklistAction = flex.FlattenStringPointer(from.BlacklistAction)
	m.BlacklistLogQuery = types.BoolPointerValue(from.BlacklistLogQuery)
	m.BlacklistRedirectAddresses = flex.FlattenFrameworkListString(ctx, from.BlacklistRedirectAddresses, diags)
	m.BlacklistRedirectTtl = flex.FlattenInt64Pointer(from.BlacklistRedirectTtl)
	m.BlacklistRulesets = flex.FlattenFrameworkListString(ctx, from.BlacklistRulesets, diags)
	m.CloudInfo = FlattenViewCloudInfo(ctx, from.CloudInfo, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CustomRootNameServers = flex.FlattenFrameworkListNestedBlock(ctx, from.CustomRootNameServers, ViewCustomRootNameServersAttrTypes, diags, FlattenViewCustomRootNameServers)
	m.DdnsForceCreationTimestampUpdate = types.BoolPointerValue(from.DdnsForceCreationTimestampUpdate)
	m.DdnsPrincipalGroup = flex.FlattenStringPointer(from.DdnsPrincipalGroup)
	m.DdnsPrincipalTracking = types.BoolPointerValue(from.DdnsPrincipalTracking)
	m.DdnsRestrictPatterns = types.BoolPointerValue(from.DdnsRestrictPatterns)
	m.DdnsRestrictPatternsList = flex.FlattenFrameworkListString(ctx, from.DdnsRestrictPatternsList, diags)
	m.DdnsRestrictProtected = types.BoolPointerValue(from.DdnsRestrictProtected)
	m.DdnsRestrictSecure = types.BoolPointerValue(from.DdnsRestrictSecure)
	m.DdnsRestrictStatic = types.BoolPointerValue(from.DdnsRestrictStatic)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.Dns64Enabled = types.BoolPointerValue(from.Dns64Enabled)
	m.Dns64Groups = flex.FlattenFrameworkListString(ctx, from.Dns64Groups, diags)
	m.DnssecEnabled = types.BoolPointerValue(from.DnssecEnabled)
	m.DnssecExpiredSignaturesEnabled = types.BoolPointerValue(from.DnssecExpiredSignaturesEnabled)
	m.DnssecNegativeTrustAnchors = flex.FlattenFrameworkListString(ctx, from.DnssecNegativeTrustAnchors, diags)
	m.DnssecTrustedKeys = flex.FlattenFrameworkListNestedBlock(ctx, from.DnssecTrustedKeys, ViewDnssecTrustedKeysAttrTypes, diags, FlattenViewDnssecTrustedKeys)
	m.DnssecValidationEnabled = types.BoolPointerValue(from.DnssecValidationEnabled)
	m.EdnsUdpSize = flex.FlattenInt64Pointer(from.EdnsUdpSize)
	m.EnableBlacklist = types.BoolPointerValue(from.EnableBlacklist)
	m.EnableFixedRrsetOrderFqdns = types.BoolPointerValue(from.EnableFixedRrsetOrderFqdns)
	m.EnableMatchRecursiveOnly = types.BoolPointerValue(from.EnableMatchRecursiveOnly)
	m.ExtAttrsAll = FlattenExtAttr(ctx, from.ExtAttrs, diags)
	m.FilterAaaa = flex.FlattenStringPointer(from.FilterAaaa)
	m.FilterAaaaList = flex.FlattenFrameworkListNestedBlock(ctx, from.FilterAaaaList, ViewFilterAaaaListAttrTypes, diags, FlattenViewFilterAaaaList)
	m.FixedRrsetOrderFqdns = flex.FlattenFrameworkListNestedBlock(ctx, from.FixedRrsetOrderFqdns, ViewFixedRrsetOrderFqdnsAttrTypes, diags, FlattenViewFixedRrsetOrderFqdns)
	m.ForwardOnly = types.BoolPointerValue(from.ForwardOnly)
	m.Forwarders = flex.FlattenFrameworkListString(ctx, from.Forwarders, diags)
	m.IsDefault = types.BoolPointerValue(from.IsDefault)
	m.LastQueriedAcl = flex.FlattenFrameworkListNestedBlock(ctx, from.LastQueriedAcl, ViewLastQueriedAclAttrTypes, diags, FlattenViewLastQueriedAcl)
	m.MatchClients = flex.FlattenFrameworkListNestedBlock(ctx, from.MatchClients, ViewMatchClientsAttrTypes, diags, FlattenViewMatchClients)
	m.MatchDestinations = flex.FlattenFrameworkListNestedBlock(ctx, from.MatchDestinations, ViewMatchDestinationsAttrTypes, diags, FlattenViewMatchDestinations)
	m.MaxCacheTtl = flex.FlattenInt64Pointer(from.MaxCacheTtl)
	m.MaxNcacheTtl = flex.FlattenInt64Pointer(from.MaxNcacheTtl)
	m.MaxUdpSize = flex.FlattenInt64Pointer(from.MaxUdpSize)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.NetworkView = flex.FlattenStringPointer(from.NetworkView)
	m.NotifyDelay = flex.FlattenInt64Pointer(from.NotifyDelay)
	m.NxdomainLogQuery = types.BoolPointerValue(from.NxdomainLogQuery)
	m.NxdomainRedirect = types.BoolPointerValue(from.NxdomainRedirect)
	m.NxdomainRedirectAddresses = flex.FlattenFrameworkListString(ctx, from.NxdomainRedirectAddresses, diags)
	m.NxdomainRedirectAddressesV6 = flex.FlattenFrameworkListString(ctx, from.NxdomainRedirectAddressesV6, diags)
	m.NxdomainRedirectTtl = flex.FlattenInt64Pointer(from.NxdomainRedirectTtl)
	m.NxdomainRulesets = flex.FlattenFrameworkListString(ctx, from.NxdomainRulesets, diags)
	m.Recursion = types.BoolPointerValue(from.Recursion)
	m.ResponseRateLimiting = FlattenViewResponseRateLimiting(ctx, from.ResponseRateLimiting, diags)
	m.RootNameServerType = flex.FlattenStringPointer(from.RootNameServerType)
	m.RpzDropIpRuleEnabled = types.BoolPointerValue(from.RpzDropIpRuleEnabled)
	m.RpzDropIpRuleMinPrefixLengthIpv4 = flex.FlattenInt64Pointer(from.RpzDropIpRuleMinPrefixLengthIpv4)
	m.RpzDropIpRuleMinPrefixLengthIpv6 = flex.FlattenInt64Pointer(from.RpzDropIpRuleMinPrefixLengthIpv6)
	m.RpzQnameWaitRecurse = types.BoolPointerValue(from.RpzQnameWaitRecurse)
	m.ScavengingSettings = FlattenViewScavengingSettings(ctx, from.ScavengingSettings, diags)
	m.Sortlist = flex.FlattenFrameworkListNestedBlock(ctx, from.Sortlist, ViewSortlistAttrTypes, diags, FlattenViewSortlist)
	m.UseBlacklist = types.BoolPointerValue(from.UseBlacklist)
	m.UseDdnsForceCreationTimestampUpdate = types.BoolPointerValue(from.UseDdnsForceCreationTimestampUpdate)
	m.UseDdnsPatternsRestriction = types.BoolPointerValue(from.UseDdnsPatternsRestriction)
	m.UseDdnsPrincipalSecurity = types.BoolPointerValue(from.UseDdnsPrincipalSecurity)
	m.UseDdnsRestrictProtected = types.BoolPointerValue(from.UseDdnsRestrictProtected)
	m.UseDdnsRestrictStatic = types.BoolPointerValue(from.UseDdnsRestrictStatic)
	m.UseDns64 = types.BoolPointerValue(from.UseDns64)
	m.UseDnssec = types.BoolPointerValue(from.UseDnssec)
	m.UseEdnsUdpSize = types.BoolPointerValue(from.UseEdnsUdpSize)
	m.UseFilterAaaa = types.BoolPointerValue(from.UseFilterAaaa)
	m.UseFixedRrsetOrderFqdns = types.BoolPointerValue(from.UseFixedRrsetOrderFqdns)
	m.UseForwarders = types.BoolPointerValue(from.UseForwarders)
	m.UseMaxCacheTtl = types.BoolPointerValue(from.UseMaxCacheTtl)
	m.UseMaxNcacheTtl = types.BoolPointerValue(from.UseMaxNcacheTtl)
	m.UseMaxUdpSize = types.BoolPointerValue(from.UseMaxUdpSize)
	m.UseNxdomainRedirect = types.BoolPointerValue(from.UseNxdomainRedirect)
	m.UseRecursion = types.BoolPointerValue(from.UseRecursion)
	m.UseResponseRateLimiting = types.BoolPointerValue(from.UseResponseRateLimiting)
	m.UseRootNameServer = types.BoolPointerValue(from.UseRootNameServer)
	m.UseRpzDropIpRule = types.BoolPointerValue(from.UseRpzDropIpRule)
	m.UseRpzQnameWaitRecurse = types.BoolPointerValue(from.UseRpzQnameWaitRecurse)
	m.UseScavengingSettings = types.BoolPointerValue(from.UseScavengingSettings)
	m.UseSortlist = types.BoolPointerValue(from.UseSortlist)
}
