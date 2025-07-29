package dhcp

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

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	internaltypes "github.com/infobloxopen/terraform-provider-nios/internal/types"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type RangeModel struct {
	Ref                              types.String                     `tfsdk:"ref"`
	AlwaysUpdateDns                  types.Bool                       `tfsdk:"always_update_dns"`
	Bootfile                         types.String                     `tfsdk:"bootfile"`
	Bootserver                       types.String                     `tfsdk:"bootserver"`
	CloudInfo                        types.Object                     `tfsdk:"cloud_info"`
	Comment                          types.String                     `tfsdk:"comment"`
	DdnsDomainname                   types.String                     `tfsdk:"ddns_domainname"`
	DdnsGenerateHostname             types.Bool                       `tfsdk:"ddns_generate_hostname"`
	DenyAllClients                   types.Bool                       `tfsdk:"deny_all_clients"`
	DenyBootp                        types.Bool                       `tfsdk:"deny_bootp"`
	DhcpUtilization                  types.Int64                      `tfsdk:"dhcp_utilization"`
	DhcpUtilizationStatus            types.String                     `tfsdk:"dhcp_utilization_status"`
	Disable                          types.Bool                       `tfsdk:"disable"`
	DiscoverNowStatus                types.String                     `tfsdk:"discover_now_status"`
	DiscoveryBasicPollSettings       types.Object                     `tfsdk:"discovery_basic_poll_settings"`
	DiscoveryBlackoutSetting         types.Object                     `tfsdk:"discovery_blackout_setting"`
	DiscoveryMember                  types.String                     `tfsdk:"discovery_member"`
	DynamicHosts                     types.Int64                      `tfsdk:"dynamic_hosts"`
	EmailList                        internaltypes.UnorderedListValue `tfsdk:"email_list"`
	EnableDdns                       types.Bool                       `tfsdk:"enable_ddns"`
	EnableDhcpThresholds             types.Bool                       `tfsdk:"enable_dhcp_thresholds"`
	EnableDiscovery                  types.Bool                       `tfsdk:"enable_discovery"`
	EnableEmailWarnings              types.Bool                       `tfsdk:"enable_email_warnings"`
	EnableIfmapPublishing            types.Bool                       `tfsdk:"enable_ifmap_publishing"`
	EnableImmediateDiscovery         types.Bool                       `tfsdk:"enable_immediate_discovery"`
	EnablePxeLeaseTime               types.Bool                       `tfsdk:"enable_pxe_lease_time"`
	EnableSnmpWarnings               types.Bool                       `tfsdk:"enable_snmp_warnings"`
	EndAddr                          types.String                     `tfsdk:"end_addr"`
	EndpointSources                  types.List                       `tfsdk:"endpoint_sources"`
	Exclude                          types.List                       `tfsdk:"exclude"`
	ExtAttrs                         types.Map                        `tfsdk:"extattrs"`
	ExtAttrsAll                      types.Map                        `tfsdk:"extattrs_all"`
	FailoverAssociation              types.String                     `tfsdk:"failover_association"`
	FingerprintFilterRules           types.List                       `tfsdk:"fingerprint_filter_rules"`
	HighWaterMark                    types.Int64                      `tfsdk:"high_water_mark"`
	HighWaterMarkReset               types.Int64                      `tfsdk:"high_water_mark_reset"`
	IgnoreDhcpOptionListRequest      types.Bool                       `tfsdk:"ignore_dhcp_option_list_request"`
	IgnoreId                         types.String                     `tfsdk:"ignore_id"`
	IgnoreMacAddresses               types.List                       `tfsdk:"ignore_mac_addresses"`
	IsSplitScope                     types.Bool                       `tfsdk:"is_split_scope"`
	KnownClients                     types.String                     `tfsdk:"known_clients"`
	LeaseScavengeTime                types.Int64                      `tfsdk:"lease_scavenge_time"`
	LogicFilterRules                 types.List                       `tfsdk:"logic_filter_rules"`
	LowWaterMark                     types.Int64                      `tfsdk:"low_water_mark"`
	LowWaterMarkReset                types.Int64                      `tfsdk:"low_water_mark_reset"`
	MacFilterRules                   types.List                       `tfsdk:"mac_filter_rules"`
	Member                           types.Object                     `tfsdk:"member"`
	MsAdUserData                     types.Object                     `tfsdk:"ms_ad_user_data"`
	MsOptions                        types.List                       `tfsdk:"ms_options"`
	MsServer                         types.Object                     `tfsdk:"ms_server"`
	NacFilterRules                   types.List                       `tfsdk:"nac_filter_rules"`
	Name                             types.String                     `tfsdk:"name"`
	Network                          types.String                     `tfsdk:"network"`
	NetworkView                      types.String                     `tfsdk:"network_view"`
	Nextserver                       types.String                     `tfsdk:"nextserver"`
	OptionFilterRules                types.List                       `tfsdk:"option_filter_rules"`
	Options                          types.List                       `tfsdk:"options"`
	PortControlBlackoutSetting       types.Object                     `tfsdk:"port_control_blackout_setting"`
	PxeLeaseTime                     types.Int64                      `tfsdk:"pxe_lease_time"`
	RecycleLeases                    types.Bool                       `tfsdk:"recycle_leases"`
	RelayAgentFilterRules            types.List                       `tfsdk:"relay_agent_filter_rules"`
	RestartIfNeeded                  types.Bool                       `tfsdk:"restart_if_needed"`
	SamePortControlDiscoveryBlackout types.Bool                       `tfsdk:"same_port_control_discovery_blackout"`
	ServerAssociationType            types.String                     `tfsdk:"server_association_type"`
	SplitMember                      types.Object                     `tfsdk:"split_member"`
	SplitScopeExclusionPercent       types.Int64                      `tfsdk:"split_scope_exclusion_percent"`
	StartAddr                        types.String                     `tfsdk:"start_addr"`
	StaticHosts                      types.Int64                      `tfsdk:"static_hosts"`
	SubscribeSettings                types.Object                     `tfsdk:"subscribe_settings"`
	Template                         types.String                     `tfsdk:"template"`
	TotalHosts                       types.Int64                      `tfsdk:"total_hosts"`
	UnknownClients                   types.String                     `tfsdk:"unknown_clients"`
	UpdateDnsOnLeaseRenewal          types.Bool                       `tfsdk:"update_dns_on_lease_renewal"`
	UseBlackoutSetting               types.Bool                       `tfsdk:"use_blackout_setting"`
	UseBootfile                      types.Bool                       `tfsdk:"use_bootfile"`
	UseBootserver                    types.Bool                       `tfsdk:"use_bootserver"`
	UseDdnsDomainname                types.Bool                       `tfsdk:"use_ddns_domainname"`
	UseDdnsGenerateHostname          types.Bool                       `tfsdk:"use_ddns_generate_hostname"`
	UseDenyBootp                     types.Bool                       `tfsdk:"use_deny_bootp"`
	UseDiscoveryBasicPollingSettings types.Bool                       `tfsdk:"use_discovery_basic_polling_settings"`
	UseEmailList                     types.Bool                       `tfsdk:"use_email_list"`
	UseEnableDdns                    types.Bool                       `tfsdk:"use_enable_ddns"`
	UseEnableDhcpThresholds          types.Bool                       `tfsdk:"use_enable_dhcp_thresholds"`
	UseEnableDiscovery               types.Bool                       `tfsdk:"use_enable_discovery"`
	UseEnableIfmapPublishing         types.Bool                       `tfsdk:"use_enable_ifmap_publishing"`
	UseIgnoreDhcpOptionListRequest   types.Bool                       `tfsdk:"use_ignore_dhcp_option_list_request"`
	UseIgnoreId                      types.Bool                       `tfsdk:"use_ignore_id"`
	UseKnownClients                  types.Bool                       `tfsdk:"use_known_clients"`
	UseLeaseScavengeTime             types.Bool                       `tfsdk:"use_lease_scavenge_time"`
	UseLogicFilterRules              types.Bool                       `tfsdk:"use_logic_filter_rules"`
	UseMsOptions                     types.Bool                       `tfsdk:"use_ms_options"`
	UseNextserver                    types.Bool                       `tfsdk:"use_nextserver"`
	UseOptions                       types.Bool                       `tfsdk:"use_options"`
	UsePxeLeaseTime                  types.Bool                       `tfsdk:"use_pxe_lease_time"`
	UseRecycleLeases                 types.Bool                       `tfsdk:"use_recycle_leases"`
	UseSubscribeSettings             types.Bool                       `tfsdk:"use_subscribe_settings"`
	UseUnknownClients                types.Bool                       `tfsdk:"use_unknown_clients"`
	UseUpdateDnsOnLeaseRenewal       types.Bool                       `tfsdk:"use_update_dns_on_lease_renewal"`
}

var RangeAttrTypes = map[string]attr.Type{
	"ref":                                  types.StringType,
	"always_update_dns":                    types.BoolType,
	"bootfile":                             types.StringType,
	"bootserver":                           types.StringType,
	"cloud_info":                           types.ObjectType{AttrTypes: RangeCloudInfoAttrTypes},
	"comment":                              types.StringType,
	"ddns_domainname":                      types.StringType,
	"ddns_generate_hostname":               types.BoolType,
	"deny_all_clients":                     types.BoolType,
	"deny_bootp":                           types.BoolType,
	"dhcp_utilization":                     types.Int64Type,
	"dhcp_utilization_status":              types.StringType,
	"disable":                              types.BoolType,
	"discover_now_status":                  types.StringType,
	"discovery_basic_poll_settings":        types.ObjectType{AttrTypes: RangeDiscoveryBasicPollSettingsAttrTypes},
	"discovery_blackout_setting":           types.ObjectType{AttrTypes: RangeDiscoveryBlackoutSettingAttrTypes},
	"discovery_member":                     types.StringType,
	"dynamic_hosts":                        types.Int64Type,
	"email_list":                           internaltypes.UnorderedListOfStringType,
	"enable_ddns":                          types.BoolType,
	"enable_dhcp_thresholds":               types.BoolType,
	"enable_discovery":                     types.BoolType,
	"enable_email_warnings":                types.BoolType,
	"enable_ifmap_publishing":              types.BoolType,
	"enable_immediate_discovery":           types.BoolType,
	"enable_pxe_lease_time":                types.BoolType,
	"enable_snmp_warnings":                 types.BoolType,
	"end_addr":                             types.StringType,
	"endpoint_sources":                     types.ListType{ElemType: types.StringType},
	"exclude":                              types.ListType{ElemType: types.ObjectType{AttrTypes: RangeExcludeAttrTypes}},
	"extattrs":                             types.MapType{ElemType: types.StringType},
	"extattrs_all":                         types.MapType{ElemType: types.StringType},
	"failover_association":                 types.StringType,
	"fingerprint_filter_rules":             types.ListType{ElemType: types.ObjectType{AttrTypes: RangeFingerprintFilterRulesAttrTypes}},
	"high_water_mark":                      types.Int64Type,
	"high_water_mark_reset":                types.Int64Type,
	"ignore_dhcp_option_list_request":      types.BoolType,
	"ignore_id":                            types.StringType,
	"ignore_mac_addresses":                 types.ListType{ElemType: types.StringType},
	"is_split_scope":                       types.BoolType,
	"known_clients":                        types.StringType,
	"lease_scavenge_time":                  types.Int64Type,
	"logic_filter_rules":                   types.ListType{ElemType: types.ObjectType{AttrTypes: RangeLogicFilterRulesAttrTypes}},
	"low_water_mark":                       types.Int64Type,
	"low_water_mark_reset":                 types.Int64Type,
	"mac_filter_rules":                     types.ListType{ElemType: types.ObjectType{AttrTypes: RangeMacFilterRulesAttrTypes}},
	"member":                               types.ObjectType{AttrTypes: RangeMemberAttrTypes},
	"ms_ad_user_data":                      types.ObjectType{AttrTypes: RangeMsAdUserDataAttrTypes},
	"ms_options":                           types.ListType{ElemType: types.ObjectType{AttrTypes: RangeMsOptionsAttrTypes}},
	"ms_server":                            types.ObjectType{AttrTypes: RangeMsServerAttrTypes},
	"nac_filter_rules":                     types.ListType{ElemType: types.ObjectType{AttrTypes: RangeNacFilterRulesAttrTypes}},
	"name":                                 types.StringType,
	"network":                              types.StringType,
	"network_view":                         types.StringType,
	"nextserver":                           types.StringType,
	"option_filter_rules":                  types.ListType{ElemType: types.ObjectType{AttrTypes: RangeOptionFilterRulesAttrTypes}},
	"options":                              types.ListType{ElemType: types.ObjectType{AttrTypes: RangeOptionsAttrTypes}},
	"port_control_blackout_setting":        types.ObjectType{AttrTypes: RangePortControlBlackoutSettingAttrTypes},
	"pxe_lease_time":                       types.Int64Type,
	"recycle_leases":                       types.BoolType,
	"relay_agent_filter_rules":             types.ListType{ElemType: types.ObjectType{AttrTypes: RangeRelayAgentFilterRulesAttrTypes}},
	"restart_if_needed":                    types.BoolType,
	"same_port_control_discovery_blackout": types.BoolType,
	"server_association_type":              types.StringType,
	"split_member":                         types.ObjectType{AttrTypes: RangeSplitMemberAttrTypes},
	"split_scope_exclusion_percent":        types.Int64Type,
	"start_addr":                           types.StringType,
	"static_hosts":                         types.Int64Type,
	"subscribe_settings":                   types.ObjectType{AttrTypes: RangeSubscribeSettingsAttrTypes},
	"template":                             types.StringType,
	"total_hosts":                          types.Int64Type,
	"unknown_clients":                      types.StringType,
	"update_dns_on_lease_renewal":          types.BoolType,
	"use_blackout_setting":                 types.BoolType,
	"use_bootfile":                         types.BoolType,
	"use_bootserver":                       types.BoolType,
	"use_ddns_domainname":                  types.BoolType,
	"use_ddns_generate_hostname":           types.BoolType,
	"use_deny_bootp":                       types.BoolType,
	"use_discovery_basic_polling_settings": types.BoolType,
	"use_email_list":                       types.BoolType,
	"use_enable_ddns":                      types.BoolType,
	"use_enable_dhcp_thresholds":           types.BoolType,
	"use_enable_discovery":                 types.BoolType,
	"use_enable_ifmap_publishing":          types.BoolType,
	"use_ignore_dhcp_option_list_request":  types.BoolType,
	"use_ignore_id":                        types.BoolType,
	"use_known_clients":                    types.BoolType,
	"use_lease_scavenge_time":              types.BoolType,
	"use_logic_filter_rules":               types.BoolType,
	"use_ms_options":                       types.BoolType,
	"use_nextserver":                       types.BoolType,
	"use_options":                          types.BoolType,
	"use_pxe_lease_time":                   types.BoolType,
	"use_recycle_leases":                   types.BoolType,
	"use_subscribe_settings":               types.BoolType,
	"use_unknown_clients":                  types.BoolType,
	"use_update_dns_on_lease_renewal":      types.BoolType,
}

var RangeResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"always_update_dns": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This field controls whether only the DHCP server is allowed to update DNS, regardless of the DHCP clients requests.",
	},
	"bootfile": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The bootfile name for the range. You can configure the DHCP server to support clients that use the boot file name option in their DHCPREQUEST messages.",
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_bootfile")),
		},
		Default: stringdefault.StaticString(""),
	},
	"bootserver": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The bootserver address for the range. You can specify the name and/or IP address of the boot server that the host needs to boot. The boot server IPv4 Address or name in FQDN format.",
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_bootserver")),
		},
		Default: stringdefault.StaticString(""),
	},
	"cloud_info": schema.SingleNestedAttribute{
		Attributes:          RangeCloudInfoResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "A CloudInfo struct that contains information about the cloud provider and region for the range.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^\S.*\S$`),
				"should not have leading or trailing whitespace",
			),
		},
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "Comment for the range; maximum 256 characters.",
	},
	"ddns_domainname": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_ddns_domainname")),
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^\S.*\S$`),
				"should not have leading or trailing whitespace",
			),
		},
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The dynamic DNS domain name the appliance uses specifically for DDNS updates for this range.",
	},
	"ddns_generate_hostname": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If this field is set to True, the DHCP server generates a hostname and updates DNS with it when the DHCP client request does not contain a hostname.",
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_ddns_generate_hostname")),
		},
	},
	"deny_all_clients": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True, send NAK forcing the client to take the new address.",
	},
	"deny_bootp": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If set to true, BOOTP settings are disabled and BOOTP requests will be denied.",
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_deny_bootp")),
		},
	},
	"dhcp_utilization": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The percentage of the total DHCP utilization of the range multiplied by 1000. This is the percentage of the total number of available IP addresses belonging to the range versus the total number of all IP addresses in the range.",
	},
	"dhcp_utilization_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "A string describing the utilization level of the range.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether a range is disabled or not. When this is set to False, the range is enabled.",
	},
	"discover_now_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Discover now status for this range.",
	},
	"discovery_basic_poll_settings": schema.SingleNestedAttribute{
		Attributes: RangeDiscoveryBasicPollSettingsResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Validators: []validator.Object{
			objectvalidator.AlsoRequires(path.MatchRoot("use_discovery_basic_polling_settings")),
		},
		MarkdownDescription: "The basic polling settings for the discovery of this range.",
	},
	"discovery_blackout_setting": schema.SingleNestedAttribute{
		Attributes: RangeDiscoveryBlackoutSettingResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Validators: []validator.Object{
			objectvalidator.AlsoRequires(path.MatchRoot("use_blackout_setting")),
		},
		MarkdownDescription: "The blackout settings for the discovery of this range. If this is set to False, the blackout settings are not used.",
	},
	"discovery_member": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The member that will run discovery for this range.",
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_enable_discovery")),
		},
		Default: stringdefault.StaticString(""),
	},
	"dynamic_hosts": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The total number of DHCP leases issued for the range.",
	},
	"email_list": schema.ListAttribute{
		CustomType:          internaltypes.UnorderedListOfStringType,
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The e-mail lists to which the appliance sends DHCP threshold alarm e-mail messages.",
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
			listvalidator.AlsoRequires(path.MatchRoot("use_email_list")),
		},
	},
	"enable_ddns": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The dynamic DNS updates flag of a DHCP range object. If set to True, the DHCP server sends DDNS updates to DNS servers in the same Grid, and to external DNS servers.",
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_enable_ddns")),
		},
	},
	"enable_dhcp_thresholds": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if DHCP thresholds are enabled for the range.",
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_enable_dhcp_thresholds")),
		},
	},
	"enable_discovery": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether a discovery is enabled or not for this range. When this is set to False, the discovery for this range is disabled.",
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_enable_discovery")),
		},
	},
	"enable_email_warnings": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if DHCP threshold warnings are sent through email.",
	},
	"enable_ifmap_publishing": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if IFMAP publishing is enabled for the range.",
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_enable_ifmap_publishing")),
		},
	},
	"enable_immediate_discovery": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the discovery for the range should be immediately enabled.",
	},
	"enable_pxe_lease_time": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Set this to True if you want the DHCP server to use a different lease time for PXE clients.",
	},
	"enable_snmp_warnings": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if DHCP threshold warnings are send through SNMP.",
	},
	"end_addr": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The IPv4 Address end address of the range.",
	},
	"endpoint_sources": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The endpoints that provides data for the DHCP Range object.",
	},
	"exclude": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RangeExcludeResourceSchemaAttributes,
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "These are ranges of IP addresses that the appliance does not use to assign to clients. You can use these exclusion addresses as static IP addresses. They contain the start and end addresses of the exclusion range, and optionally, information about this exclusion range.",
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
	"failover_association": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The name of the failover association: the server in this failover association will serve the IPv4 range in case the main server is out of service. {range:range} must be set to 'FAILOVER' or 'FAILOVER_MS' if you want the failover association specified here to serve the range.",
	},
	"fingerprint_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RangeFingerprintFilterRulesResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "This field contains the fingerprint filters for this DHCP range. The appliance uses matching rules in these filters to select the address range from which it assigns a lease.",
	},
	"high_water_mark": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(95),
		Validators: []validator.Int64{
			int64validator.Between(1, 100),
		},
		MarkdownDescription: "The percentage of DHCP range usage threshold above which range usage is not expected and may warrant your attention. When the high watermark is reached, the Infoblox appliance generates a syslog message and sends a warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100.",
	},
	"high_water_mark_reset": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(85),
		Validators: []validator.Int64{
			int64validator.Between(1, 100),
		},
		MarkdownDescription: "The percentage of DHCP range usage below which the corresponding SNMP trap is reset. A number that specifies the percentage of allocated addresses. The range is from 1 to 100. The high watermark reset value must be lower than the high watermark value.",
	},
	"ignore_dhcp_option_list_request": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If this field is set to False, the appliance returns all DHCP options the client is eligible to receive, rather than only the list of options the client has requested.",
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_ignore_dhcp_option_list_request")),
		},
	},
	"ignore_id": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString("NONE"),
		MarkdownDescription: "Indicates whether the appliance will ignore DHCP client IDs or MAC addresses. Valid values are \"NONE\", \"CLIENT\", or \"MACADDR\". The default is \"NONE\".",
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_ignore_id")),
			stringvalidator.OneOf("CLIENT", "MACADDR", "NONE"),
		},
	},
	"ignore_mac_addresses": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "A list of MAC addresses the appliance will ignore.",
	},
	"is_split_scope": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "This field will be 'true' if this particular range is part of a split scope.",
	},
	"known_clients": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_known_clients")),
		},
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "Permission for known clients. This can be 'Allow' or 'Deny'. If set to 'Deny' known clients will be denied IP addresses. Known clients include roaming hosts and clients with fixed addresses or DHCP host entries. Unknown clients include clients that are not roaming hosts and clients that do not have fixed addresses or DHCP host entries.",
	},
	"lease_scavenge_time": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_lease_scavenge_time")),
			int64validator.Any(
				int64validator.OneOf(-1),
				int64validator.Between(86400, 2147472000),
			),
		},
		Default:             int64default.StaticInt64(-1),
		MarkdownDescription: "An integer that specifies the period of time (in seconds) that frees and backs up leases remained in the database before they are automatically deleted. To disable lease scavenging, set the parameter to -1. The minimum positive value must be greater than 86400 seconds (1 day).",
	},
	"logic_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RangeLogicFilterRulesResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "This field contains the logic filters to be applied to this range. This list corresponds to the match rules that are written to the dhcpd configuration file.",
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_logic_filter_rules")),
		},
	},
	"low_water_mark": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.Any(
				int64validator.Between(0, 100),
			),
		},
		Default:             int64default.StaticInt64(0),
		MarkdownDescription: "The percentage of DHCP range usage below which the Infoblox appliance generates a syslog message and sends a warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100.",
	},
	"low_water_mark_reset": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(10),
		Validators: []validator.Int64{
			int64validator.Any(
				int64validator.Between(1, 100),
			),
		},
		MarkdownDescription: "The percentage of DHCP range usage threshold below which range usage is not expected and may warrant your attention. When the low watermark is crossed, the Infoblox appliance generates a syslog message and sends a warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100. The low watermark reset value must be higher than the low watermark value.",
	},
	"mac_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RangeMacFilterRulesResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "This field contains the MAC filters to be applied to this range. The appliance uses the matching rules of these filters to select the address range from which it assigns a lease.",
	},
	"member": schema.SingleNestedAttribute{
		Attributes:          RangeMemberResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "This field contains the member that will run the DHCP service for this range. If this is not set, the range will be served by the member that is currently serving the network.",
	},
	"ms_ad_user_data": schema.SingleNestedAttribute{
		Attributes:          RangeMsAdUserDataResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "This field contains the Microsoft AD user data for this range. This data is used to create a user in the Microsoft AD when a lease is assigned to a host in this range.",
	},
	"ms_options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RangeMsOptionsResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_ms_options")),
		},
		MarkdownDescription: "This field contains the Microsoft DHCP options for this range.",
	},
	"ms_server": schema.SingleNestedAttribute{
		Attributes:          RangeMsServerResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "This field contains the Microsoft server that will serve this range. This is used for Microsoft failover.",
	},
	"nac_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RangeNacFilterRulesResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "This field contains the NAC filters to be applied to this range. The appliance uses the matching rules of these filters to select the address range from which it assigns a lease.",
	},
	"name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^\S.*\S$`),
				"should not have leading or trailing whitespace",
			),
		},
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "This field contains the name of the Microsoft scope.",
	},
	"network": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The network to which this range belongs, in IPv4 Address/CIDR format.",
	},
	"network_view": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString("default"),
		MarkdownDescription: "The name of the network view in which this range resides.",
	},
	"nextserver": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.IsValidFQDN(),
			stringvalidator.AlsoRequires(path.MatchRoot("use_nextserver")),
		},
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The name in FQDN and/or IPv4 Address of the next server that the host needs to boot.",
	},
	"option_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RangeOptionFilterRulesResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "This field contains the Option filters to be applied to this range. The appliance uses the matching rules of these filters to select the address range from which it assigns a lease.",
	},
	"options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RangeOptionsResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_options")),
		},
		MarkdownDescription: "An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object.",
	},
	"port_control_blackout_setting": schema.SingleNestedAttribute{
		Attributes: RangePortControlBlackoutSettingResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Validators: []validator.Object{
			objectvalidator.AlsoRequires(path.MatchRoot("use_blackout_setting")),
		},
		MarkdownDescription: "The port control blackout settings for the range. This field is used to configure the port control blackout settings for the DHCP range. It includes information about the blackout settings, such as the start and end times of the blackout period.",
	},
	"pxe_lease_time": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The PXE lease time value of a DHCP Range object. Some hosts use PXE (Preboot Execution Environment) to boot remotely from a server. To better manage your IP resources, set a different lease time for PXE boot requests. You can configure the DHCP server to allocate an IP address with a shorter lease time to hosts that send PXE boot requests, so IP addresses are not leased longer than necessary. A 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached.",
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_pxe_lease_time")),
		},
	},
	"recycle_leases": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(true),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_recycle_leases")),
		},
		MarkdownDescription: "If the field is set to True, the leases are kept in the Recycle Bin until one week after expiration. Otherwise, the leases are permanently deleted.",
	},
	"relay_agent_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RangeRelayAgentFilterRulesResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "This field contains the Relay Agent filters to be applied to this range. The appliance uses the matching rules of these filters to select the address range from which it assigns a lease.",
	},
	"restart_if_needed": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Restarts the member service.",
	},
	"same_port_control_discovery_blackout": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_blackout_setting")),
		},
		MarkdownDescription: "If the field is set to True, the discovery blackout setting will be used for port control blackout setting.",
	},
	"server_association_type": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString("NONE"),
		MarkdownDescription: "The type of server that is going to serve the range.",
		Validators: []validator.String{
			stringvalidator.OneOf("MEMBER", "FAILOVER", "MS_FAILOVER", "MS_SERVER", "NONE"),
		},
	},
	"split_member": schema.SingleNestedAttribute{
		Attributes:          RangeSplitMemberResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "This field contains the split member that will run the DHCP service for this range. If this is not set, the range will be served by the member that is currently serving the network.",
	},
	"split_scope_exclusion_percent": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "This field controls the percentage used when creating a split scope. Valid values are numbers between 1 and 99. If the value is 40, it means that the top 40% of the exclusion will be created on the DHCP range assigned to {next_available_ip:next_available_ip} and the lower 60% of the range will be assigned to DHCP range assigned to {next_available_ip:next_available_ip}",
	},
	"start_addr": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The IPv4 Address starting address of the range.",
	},
	"static_hosts": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of static DHCP addresses configured in the range.",
	},
	"subscribe_settings": schema.SingleNestedAttribute{
		Attributes: RangeSubscribeSettingsResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Validators: []validator.Object{
			objectvalidator.AlsoRequires(path.MatchRoot("use_subscribe_settings")),
		},
		MarkdownDescription: "The subscribe settings for the range. This field is used to configure the subscription settings for the DHCP range. It includes information about the subscription, such as the subscriber's email address and whether the subscription is enabled.",
	},
	"template": schema.StringAttribute{
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "If set on creation, the range will be created according to the values specified in the named template.",
	},
	"total_hosts": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The total number of DHCP addresses configured in the range.",
	},
	"unknown_clients": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_unknown_clients")),
		},
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "Permission for unknown clients. This can be 'Allow' or 'Deny'. If set to 'Deny', unknown clients will be denied IP addresses. Known clients include roaming hosts and clients with fixed addresses or DHCP host entries. Unknown clients include clients that are not roaming hosts and clients that do not have fixed addresses or DHCP host entries.",
	},
	"update_dns_on_lease_renewal": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_update_dns_on_lease_renewal")),
		},
		MarkdownDescription: "This field controls whether the DHCP server updates DNS when a DHCP lease is renewed.",
	},
	"use_blackout_setting": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: discovery_blackout_setting , port_control_blackout_setting, same_port_control_discovery_blackout",
	},
	"use_bootfile": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: bootfile",
	},
	"use_bootserver": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: bootserver",
	},
	"use_ddns_domainname": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: ddns_domainname",
	},
	"use_ddns_generate_hostname": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: ddns_generate_hostname",
	},
	"use_deny_bootp": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: deny_bootp",
	},
	"use_discovery_basic_polling_settings": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: discovery_basic_poll_settings",
	},
	"use_email_list": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: email_list",
	},
	"use_enable_ddns": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: enable_ddns",
	},
	"use_enable_dhcp_thresholds": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: enable_dhcp_thresholds",
	},
	"use_enable_discovery": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: discovery_member , enable_discovery",
	},
	"use_enable_ifmap_publishing": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: enable_ifmap_publishing",
	},
	"use_ignore_dhcp_option_list_request": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: ignore_dhcp_option_list_request",
	},
	"use_ignore_id": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: ignore_id",
	},
	"use_known_clients": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: known_clients",
	},
	"use_lease_scavenge_time": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: lease_scavenge_time",
	},
	"use_logic_filter_rules": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: logic_filter_rules",
	},
	"use_ms_options": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: ms_options",
	},
	"use_nextserver": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: nextserver",
	},
	"use_options": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: options",
	},
	"use_pxe_lease_time": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: pxe_lease_time",
	},
	"use_recycle_leases": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: recycle_leases",
	},
	"use_subscribe_settings": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: subscribe_settings",
	},
	"use_unknown_clients": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: unknown_clients",
	},
	"use_update_dns_on_lease_renewal": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: update_dns_on_lease_renewal",
	},
}

func (m *RangeModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *dhcp.Range {
	if m == nil {
		return nil
	}
	to := &dhcp.Range{
		AlwaysUpdateDns:                  flex.ExpandBoolPointer(m.AlwaysUpdateDns),
		Bootfile:                         flex.ExpandStringPointer(m.Bootfile),
		Bootserver:                       flex.ExpandStringPointer(m.Bootserver),
		CloudInfo:                        ExpandRangeCloudInfo(ctx, m.CloudInfo, diags),
		Comment:                          flex.ExpandStringPointer(m.Comment),
		DdnsDomainname:                   flex.ExpandStringPointer(m.DdnsDomainname),
		DdnsGenerateHostname:             flex.ExpandBoolPointer(m.DdnsGenerateHostname),
		DenyAllClients:                   flex.ExpandBoolPointer(m.DenyAllClients),
		DenyBootp:                        flex.ExpandBoolPointer(m.DenyBootp),
		Disable:                          flex.ExpandBoolPointer(m.Disable),
		DiscoveryBasicPollSettings:       ExpandRangeDiscoveryBasicPollSettings(ctx, m.DiscoveryBasicPollSettings, diags),
		DiscoveryBlackoutSetting:         ExpandRangeDiscoveryBlackoutSetting(ctx, m.DiscoveryBlackoutSetting, diags),
		DiscoveryMember:                  flex.ExpandStringPointer(m.DiscoveryMember),
		EmailList:                        flex.ExpandFrameworkListString(ctx, m.EmailList, diags),
		EnableDdns:                       flex.ExpandBoolPointer(m.EnableDdns),
		EnableDhcpThresholds:             flex.ExpandBoolPointer(m.EnableDhcpThresholds),
		EnableDiscovery:                  flex.ExpandBoolPointer(m.EnableDiscovery),
		EnableEmailWarnings:              flex.ExpandBoolPointer(m.EnableEmailWarnings),
		EnableIfmapPublishing:            flex.ExpandBoolPointer(m.EnableIfmapPublishing),
		EnableImmediateDiscovery:         flex.ExpandBoolPointer(m.EnableImmediateDiscovery),
		EnablePxeLeaseTime:               flex.ExpandBoolPointer(m.EnablePxeLeaseTime),
		EnableSnmpWarnings:               flex.ExpandBoolPointer(m.EnableSnmpWarnings),
		EndAddr:                          flex.ExpandStringPointer(m.EndAddr),
		Exclude:                          flex.ExpandFrameworkListNestedBlock(ctx, m.Exclude, diags, ExpandRangeExclude),
		ExtAttrs:                         ExpandExtAttr(ctx, m.ExtAttrs, diags),
		FailoverAssociation:              flex.ExpandStringPointer(m.FailoverAssociation),
		FingerprintFilterRules:           flex.ExpandFrameworkListNestedBlock(ctx, m.FingerprintFilterRules, diags, ExpandRangeFingerprintFilterRules),
		HighWaterMark:                    flex.ExpandInt64Pointer(m.HighWaterMark),
		HighWaterMarkReset:               flex.ExpandInt64Pointer(m.HighWaterMarkReset),
		IgnoreDhcpOptionListRequest:      flex.ExpandBoolPointer(m.IgnoreDhcpOptionListRequest),
		IgnoreId:                         flex.ExpandStringPointer(m.IgnoreId),
		IgnoreMacAddresses:               flex.ExpandFrameworkListString(ctx, m.IgnoreMacAddresses, diags),
		KnownClients:                     flex.ExpandStringPointer(m.KnownClients),
		LeaseScavengeTime:                flex.ExpandInt64Pointer(m.LeaseScavengeTime),
		LogicFilterRules:                 flex.ExpandFrameworkListNestedBlock(ctx, m.LogicFilterRules, diags, ExpandRangeLogicFilterRules),
		LowWaterMark:                     flex.ExpandInt64Pointer(m.LowWaterMark),
		LowWaterMarkReset:                flex.ExpandInt64Pointer(m.LowWaterMarkReset),
		MacFilterRules:                   flex.ExpandFrameworkListNestedBlock(ctx, m.MacFilterRules, diags, ExpandRangeMacFilterRules),
		Member:                           ExpandRangeMember(ctx, m.Member, diags),
		MsAdUserData:                     ExpandRangeMsAdUserData(ctx, m.MsAdUserData, diags),
		MsOptions:                        flex.ExpandFrameworkListNestedBlock(ctx, m.MsOptions, diags, ExpandRangeMsOptions),
		MsServer:                         ExpandRangeMsServer(ctx, m.MsServer, diags),
		NacFilterRules:                   flex.ExpandFrameworkListNestedBlock(ctx, m.NacFilterRules, diags, ExpandRangeNacFilterRules),
		Name:                             flex.ExpandStringPointer(m.Name),
		Network:                          flex.ExpandStringPointer(m.Network),
		NetworkView:                      flex.ExpandStringPointer(m.NetworkView),
		Nextserver:                       flex.ExpandStringPointer(m.Nextserver),
		OptionFilterRules:                flex.ExpandFrameworkListNestedBlock(ctx, m.OptionFilterRules, diags, ExpandRangeOptionFilterRules),
		Options:                          flex.ExpandFrameworkListNestedBlock(ctx, m.Options, diags, ExpandRangeOptions),
		PortControlBlackoutSetting:       ExpandRangePortControlBlackoutSetting(ctx, m.PortControlBlackoutSetting, diags),
		PxeLeaseTime:                     flex.ExpandInt64Pointer(m.PxeLeaseTime),
		RecycleLeases:                    flex.ExpandBoolPointer(m.RecycleLeases),
		RelayAgentFilterRules:            flex.ExpandFrameworkListNestedBlock(ctx, m.RelayAgentFilterRules, diags, ExpandRangeRelayAgentFilterRules),
		RestartIfNeeded:                  flex.ExpandBoolPointer(m.RestartIfNeeded),
		SamePortControlDiscoveryBlackout: flex.ExpandBoolPointer(m.SamePortControlDiscoveryBlackout),
		ServerAssociationType:            flex.ExpandStringPointer(m.ServerAssociationType),
		StartAddr:                        flex.ExpandStringPointer(m.StartAddr),
		SubscribeSettings:                ExpandRangeSubscribeSettings(ctx, m.SubscribeSettings, diags),
		UnknownClients:                   flex.ExpandStringPointer(m.UnknownClients),
		UpdateDnsOnLeaseRenewal:          flex.ExpandBoolPointer(m.UpdateDnsOnLeaseRenewal),
		UseBlackoutSetting:               flex.ExpandBoolPointer(m.UseBlackoutSetting),
		UseBootfile:                      flex.ExpandBoolPointer(m.UseBootfile),
		UseBootserver:                    flex.ExpandBoolPointer(m.UseBootserver),
		UseDdnsDomainname:                flex.ExpandBoolPointer(m.UseDdnsDomainname),
		UseDdnsGenerateHostname:          flex.ExpandBoolPointer(m.UseDdnsGenerateHostname),
		UseDenyBootp:                     flex.ExpandBoolPointer(m.UseDenyBootp),
		UseDiscoveryBasicPollingSettings: flex.ExpandBoolPointer(m.UseDiscoveryBasicPollingSettings),
		UseEmailList:                     flex.ExpandBoolPointer(m.UseEmailList),
		UseEnableDdns:                    flex.ExpandBoolPointer(m.UseEnableDdns),
		UseEnableDhcpThresholds:          flex.ExpandBoolPointer(m.UseEnableDhcpThresholds),
		UseEnableDiscovery:               flex.ExpandBoolPointer(m.UseEnableDiscovery),
		UseEnableIfmapPublishing:         flex.ExpandBoolPointer(m.UseEnableIfmapPublishing),
		UseIgnoreDhcpOptionListRequest:   flex.ExpandBoolPointer(m.UseIgnoreDhcpOptionListRequest),
		UseIgnoreId:                      flex.ExpandBoolPointer(m.UseIgnoreId),
		UseKnownClients:                  flex.ExpandBoolPointer(m.UseKnownClients),
		UseLeaseScavengeTime:             flex.ExpandBoolPointer(m.UseLeaseScavengeTime),
		UseLogicFilterRules:              flex.ExpandBoolPointer(m.UseLogicFilterRules),
		UseMsOptions:                     flex.ExpandBoolPointer(m.UseMsOptions),
		UseNextserver:                    flex.ExpandBoolPointer(m.UseNextserver),
		UseOptions:                       flex.ExpandBoolPointer(m.UseOptions),
		UsePxeLeaseTime:                  flex.ExpandBoolPointer(m.UsePxeLeaseTime),
		UseRecycleLeases:                 flex.ExpandBoolPointer(m.UseRecycleLeases),
		UseSubscribeSettings:             flex.ExpandBoolPointer(m.UseSubscribeSettings),
		UseUnknownClients:                flex.ExpandBoolPointer(m.UseUnknownClients),
		UseUpdateDnsOnLeaseRenewal:       flex.ExpandBoolPointer(m.UseUpdateDnsOnLeaseRenewal),
	}
	if isCreate {
		to.SplitMember = ExpandRangeSplitMember(ctx, m.SplitMember, diags)
		to.SplitScopeExclusionPercent = flex.ExpandInt64Pointer(m.SplitScopeExclusionPercent)
		to.Template = flex.ExpandStringPointer(m.Template)
	}
	return to
}

func FlattenRange(ctx context.Context, from *dhcp.Range, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RangeAttrTypes)
	}
	m := RangeModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrs = m.ExtAttrsAll
	t, d := types.ObjectValueFrom(ctx, RangeAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RangeModel) Flatten(ctx context.Context, from *dhcp.Range, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RangeModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.AlwaysUpdateDns = types.BoolPointerValue(from.AlwaysUpdateDns)
	m.Bootfile = flex.FlattenStringPointer(from.Bootfile)
	m.Bootserver = flex.FlattenStringPointer(from.Bootserver)
	m.CloudInfo = FlattenRangeCloudInfo(ctx, from.CloudInfo, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.DdnsDomainname = flex.FlattenStringPointer(from.DdnsDomainname)
	m.DdnsGenerateHostname = types.BoolPointerValue(from.DdnsGenerateHostname)
	m.DenyAllClients = types.BoolPointerValue(from.DenyAllClients)
	m.DenyBootp = types.BoolPointerValue(from.DenyBootp)
	m.DhcpUtilization = flex.FlattenInt64Pointer(from.DhcpUtilization)
	m.DhcpUtilizationStatus = flex.FlattenStringPointer(from.DhcpUtilizationStatus)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.DiscoverNowStatus = flex.FlattenStringPointer(from.DiscoverNowStatus)
	m.DiscoveryBasicPollSettings = FlattenRangeDiscoveryBasicPollSettings(ctx, from.DiscoveryBasicPollSettings, diags)
	m.DiscoveryBlackoutSetting = FlattenRangeDiscoveryBlackoutSetting(ctx, from.DiscoveryBlackoutSetting, diags)
	m.DiscoveryMember = flex.FlattenStringPointer(from.DiscoveryMember)
	m.DynamicHosts = flex.FlattenInt64Pointer(from.DynamicHosts)
	m.EmailList = flex.FlattenFrameworkUnorderedList(ctx, types.StringType, from.EmailList, diags)
	m.EnableDdns = types.BoolPointerValue(from.EnableDdns)
	m.EnableDhcpThresholds = types.BoolPointerValue(from.EnableDhcpThresholds)
	m.EnableDiscovery = types.BoolPointerValue(from.EnableDiscovery)
	m.EnableEmailWarnings = types.BoolPointerValue(from.EnableEmailWarnings)
	m.EnableIfmapPublishing = types.BoolPointerValue(from.EnableIfmapPublishing)
	m.EnablePxeLeaseTime = types.BoolPointerValue(from.EnablePxeLeaseTime)
	m.EnableSnmpWarnings = types.BoolPointerValue(from.EnableSnmpWarnings)
	m.EndAddr = flex.FlattenStringPointer(from.EndAddr)
	m.EndpointSources = flex.FlattenFrameworkListString(ctx, from.EndpointSources, diags)
	m.Exclude = flex.FlattenFrameworkListNestedBlock(ctx, from.Exclude, RangeExcludeAttrTypes, diags, FlattenRangeExclude)
	m.ExtAttrsAll = FlattenExtAttr(ctx, from.ExtAttrs, diags)
	m.FailoverAssociation = flex.FlattenStringPointer(from.FailoverAssociation)
	m.FingerprintFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.FingerprintFilterRules, RangeFingerprintFilterRulesAttrTypes, diags, FlattenRangeFingerprintFilterRules)
	m.HighWaterMark = flex.FlattenInt64Pointer(from.HighWaterMark)
	m.HighWaterMarkReset = flex.FlattenInt64Pointer(from.HighWaterMarkReset)
	m.IgnoreDhcpOptionListRequest = types.BoolPointerValue(from.IgnoreDhcpOptionListRequest)
	m.IgnoreId = flex.FlattenStringPointer(from.IgnoreId)
	m.IgnoreMacAddresses = flex.FlattenFrameworkListString(ctx, from.IgnoreMacAddresses, diags)
	m.IsSplitScope = types.BoolPointerValue(from.IsSplitScope)
	m.KnownClients = flex.FlattenStringPointer(from.KnownClients)
	m.LeaseScavengeTime = flex.FlattenInt64Pointer(from.LeaseScavengeTime)
	m.LogicFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.LogicFilterRules, RangeLogicFilterRulesAttrTypes, diags, FlattenRangeLogicFilterRules)
	m.LowWaterMark = flex.FlattenInt64Pointer(from.LowWaterMark)
	m.LowWaterMarkReset = flex.FlattenInt64Pointer(from.LowWaterMarkReset)
	m.MacFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.MacFilterRules, RangeMacFilterRulesAttrTypes, diags, FlattenRangeMacFilterRules)
	m.Member = FlattenRangeMember(ctx, from.Member, diags)
	m.MsAdUserData = FlattenRangeMsAdUserData(ctx, from.MsAdUserData, diags)
	m.MsOptions = flex.FlattenFrameworkListNestedBlock(ctx, from.MsOptions, RangeMsOptionsAttrTypes, diags, FlattenRangeMsOptions)
	m.MsServer = FlattenRangeMsServer(ctx, from.MsServer, diags)
	m.NacFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.NacFilterRules, RangeNacFilterRulesAttrTypes, diags, FlattenRangeNacFilterRules)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Network = flex.FlattenStringPointer(from.Network)
	m.NetworkView = flex.FlattenStringPointer(from.NetworkView)
	m.Nextserver = flex.FlattenStringPointer(from.Nextserver)
	m.OptionFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.OptionFilterRules, RangeOptionFilterRulesAttrTypes, diags, FlattenRangeOptionFilterRules)
	m.Options = flex.FlattenFrameworkListNestedBlock(ctx, from.Options, RangeOptionsAttrTypes, diags, FlattenRangeOptions)
	m.PortControlBlackoutSetting = FlattenRangePortControlBlackoutSetting(ctx, from.PortControlBlackoutSetting, diags)
	m.PxeLeaseTime = flex.FlattenInt64Pointer(from.PxeLeaseTime)
	m.RecycleLeases = types.BoolPointerValue(from.RecycleLeases)
	m.RelayAgentFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.RelayAgentFilterRules, RangeRelayAgentFilterRulesAttrTypes, diags, FlattenRangeRelayAgentFilterRules)
	m.SamePortControlDiscoveryBlackout = types.BoolPointerValue(from.SamePortControlDiscoveryBlackout)
	m.ServerAssociationType = flex.FlattenStringPointer(from.ServerAssociationType)
	m.StartAddr = flex.FlattenStringPointer(from.StartAddr)
	m.StaticHosts = flex.FlattenInt64Pointer(from.StaticHosts)
	m.SubscribeSettings = FlattenRangeSubscribeSettings(ctx, from.SubscribeSettings, diags)
	m.Template = flex.FlattenStringPointer(from.Template)
	m.TotalHosts = flex.FlattenInt64Pointer(from.TotalHosts)
	m.UnknownClients = flex.FlattenStringPointer(from.UnknownClients)
	m.UpdateDnsOnLeaseRenewal = types.BoolPointerValue(from.UpdateDnsOnLeaseRenewal)
	m.UseBlackoutSetting = types.BoolPointerValue(from.UseBlackoutSetting)
	m.UseBootfile = types.BoolPointerValue(from.UseBootfile)
	m.UseBootserver = types.BoolPointerValue(from.UseBootserver)
	m.UseDdnsDomainname = types.BoolPointerValue(from.UseDdnsDomainname)
	m.UseDdnsGenerateHostname = types.BoolPointerValue(from.UseDdnsGenerateHostname)
	m.UseDenyBootp = types.BoolPointerValue(from.UseDenyBootp)
	m.UseDiscoveryBasicPollingSettings = types.BoolPointerValue(from.UseDiscoveryBasicPollingSettings)
	m.UseEmailList = types.BoolPointerValue(from.UseEmailList)
	m.UseEnableDdns = types.BoolPointerValue(from.UseEnableDdns)
	m.UseEnableDhcpThresholds = types.BoolPointerValue(from.UseEnableDhcpThresholds)
	m.UseEnableDiscovery = types.BoolPointerValue(from.UseEnableDiscovery)
	m.UseEnableIfmapPublishing = types.BoolPointerValue(from.UseEnableIfmapPublishing)
	m.UseIgnoreDhcpOptionListRequest = types.BoolPointerValue(from.UseIgnoreDhcpOptionListRequest)
	m.UseIgnoreId = types.BoolPointerValue(from.UseIgnoreId)
	m.UseKnownClients = types.BoolPointerValue(from.UseKnownClients)
	m.UseLeaseScavengeTime = types.BoolPointerValue(from.UseLeaseScavengeTime)
	m.UseLogicFilterRules = types.BoolPointerValue(from.UseLogicFilterRules)
	m.UseMsOptions = types.BoolPointerValue(from.UseMsOptions)
	m.UseNextserver = types.BoolPointerValue(from.UseNextserver)
	m.UseOptions = types.BoolPointerValue(from.UseOptions)
	m.UsePxeLeaseTime = types.BoolPointerValue(from.UsePxeLeaseTime)
	m.UseRecycleLeases = types.BoolPointerValue(from.UseRecycleLeases)
	m.UseSubscribeSettings = types.BoolPointerValue(from.UseSubscribeSettings)
	m.UseUnknownClients = types.BoolPointerValue(from.UseUnknownClients)
	m.UseUpdateDnsOnLeaseRenewal = types.BoolPointerValue(from.UseUpdateDnsOnLeaseRenewal)
}
