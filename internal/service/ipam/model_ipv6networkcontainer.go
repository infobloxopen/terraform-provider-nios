package ipam

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
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	internaltypes "github.com/infobloxopen/terraform-provider-nios/internal/types"
)

type Ipv6networkcontainerModel struct {
	Ref                              types.String                     `tfsdk:"ref"`
	AutoCreateReversezone            types.Bool                       `tfsdk:"auto_create_reversezone"`
	CloudInfo                        types.Object                     `tfsdk:"cloud_info"`
	Comment                          types.String                     `tfsdk:"comment"`
	DdnsDomainname                   types.String                     `tfsdk:"ddns_domainname"`
	DdnsEnableOptionFqdn             types.Bool                       `tfsdk:"ddns_enable_option_fqdn"`
	DdnsGenerateHostname             types.Bool                       `tfsdk:"ddns_generate_hostname"`
	DdnsServerAlwaysUpdates          types.Bool                       `tfsdk:"ddns_server_always_updates"`
	DdnsTtl                          types.Int64                      `tfsdk:"ddns_ttl"`
	DeleteReason                     types.String                     `tfsdk:"delete_reason"`
	DiscoverNowStatus                types.String                     `tfsdk:"discover_now_status"`
	DiscoveryBasicPollSettings       types.Object                     `tfsdk:"discovery_basic_poll_settings"`
	DiscoveryBlackoutSetting         types.Object                     `tfsdk:"discovery_blackout_setting"`
	DiscoveryEngineType              types.String                     `tfsdk:"discovery_engine_type"`
	DiscoveryMember                  types.String                     `tfsdk:"discovery_member"`
	DomainNameServers                types.List                       `tfsdk:"domain_name_servers"`
	EnableDdns                       types.Bool                       `tfsdk:"enable_ddns"`
	EnableDiscovery                  types.Bool                       `tfsdk:"enable_discovery"`
	EnableImmediateDiscovery         types.Bool                       `tfsdk:"enable_immediate_discovery"`
	EndpointSources                  types.List                       `tfsdk:"endpoint_sources"`
	ExtAttrs                         types.Map                        `tfsdk:"extattrs"`
	ExtAttrsAll                      types.Map                        `tfsdk:"extattrs_all"`
	FederatedRealms                  types.List                       `tfsdk:"federated_realms"`
	LastRirRegistrationUpdateSent    types.Int64                      `tfsdk:"last_rir_registration_update_sent"`
	LastRirRegistrationUpdateStatus  types.String                     `tfsdk:"last_rir_registration_update_status"`
	LogicFilterRules                 types.List                       `tfsdk:"logic_filter_rules"`
	MgmPrivate                       types.Bool                       `tfsdk:"mgm_private"`
	MgmPrivateOverridable            types.Bool                       `tfsdk:"mgm_private_overridable"`
	MsAdUserData                     types.Object                     `tfsdk:"ms_ad_user_data"`
	Network                          types.String                     `tfsdk:"network"`
	FuncCall                         types.Object                     `tfsdk:"func_call"`
	NetworkContainer                 types.String                     `tfsdk:"network_container"`
	NetworkView                      types.String                     `tfsdk:"network_view"`
	Options                          internaltypes.UnorderedListValue `tfsdk:"options"`
	PortControlBlackoutSetting       types.Object                     `tfsdk:"port_control_blackout_setting"`
	PreferredLifetime                types.Int64                      `tfsdk:"preferred_lifetime"`
	RemoveSubnets                    types.Bool                       `tfsdk:"remove_subnets"`
	RestartIfNeeded                  types.Bool                       `tfsdk:"restart_if_needed"`
	Rir                              types.String                     `tfsdk:"rir"`
	RirOrganization                  types.String                     `tfsdk:"rir_organization"`
	RirRegistrationAction            types.String                     `tfsdk:"rir_registration_action"`
	RirRegistrationStatus            types.String                     `tfsdk:"rir_registration_status"`
	SamePortControlDiscoveryBlackout types.Bool                       `tfsdk:"same_port_control_discovery_blackout"`
	SendRirRequest                   types.Bool                       `tfsdk:"send_rir_request"`
	SubscribeSettings                types.Object                     `tfsdk:"subscribe_settings"`
	Unmanaged                        types.Bool                       `tfsdk:"unmanaged"`
	UpdateDnsOnLeaseRenewal          types.Bool                       `tfsdk:"update_dns_on_lease_renewal"`
	UseBlackoutSetting               types.Bool                       `tfsdk:"use_blackout_setting"`
	UseDdnsDomainname                types.Bool                       `tfsdk:"use_ddns_domainname"`
	UseDdnsEnableOptionFqdn          types.Bool                       `tfsdk:"use_ddns_enable_option_fqdn"`
	UseDdnsGenerateHostname          types.Bool                       `tfsdk:"use_ddns_generate_hostname"`
	UseDdnsTtl                       types.Bool                       `tfsdk:"use_ddns_ttl"`
	UseDiscoveryBasicPollingSettings types.Bool                       `tfsdk:"use_discovery_basic_polling_settings"`
	UseDomainNameServers             types.Bool                       `tfsdk:"use_domain_name_servers"`
	UseEnableDdns                    types.Bool                       `tfsdk:"use_enable_ddns"`
	UseEnableDiscovery               types.Bool                       `tfsdk:"use_enable_discovery"`
	UseLogicFilterRules              types.Bool                       `tfsdk:"use_logic_filter_rules"`
	UseMgmPrivate                    types.Bool                       `tfsdk:"use_mgm_private"`
	UseOptions                       types.Bool                       `tfsdk:"use_options"`
	UsePreferredLifetime             types.Bool                       `tfsdk:"use_preferred_lifetime"`
	UseSubscribeSettings             types.Bool                       `tfsdk:"use_subscribe_settings"`
	UseUpdateDnsOnLeaseRenewal       types.Bool                       `tfsdk:"use_update_dns_on_lease_renewal"`
	UseValidLifetime                 types.Bool                       `tfsdk:"use_valid_lifetime"`
	UseZoneAssociations              types.Bool                       `tfsdk:"use_zone_associations"`
	Utilization                      types.Int64                      `tfsdk:"utilization"`
	ValidLifetime                    types.Int64                      `tfsdk:"valid_lifetime"`
	ZoneAssociations                 types.List                       `tfsdk:"zone_associations"`
}

var Ipv6networkcontainerAttrTypes = map[string]attr.Type{
	"ref":                                  types.StringType,
	"auto_create_reversezone":              types.BoolType,
	"cloud_info":                           types.ObjectType{AttrTypes: Ipv6networkcontainerCloudInfoAttrTypes},
	"comment":                              types.StringType,
	"ddns_domainname":                      types.StringType,
	"ddns_enable_option_fqdn":              types.BoolType,
	"ddns_generate_hostname":               types.BoolType,
	"ddns_server_always_updates":           types.BoolType,
	"ddns_ttl":                             types.Int64Type,
	"delete_reason":                        types.StringType,
	"discover_now_status":                  types.StringType,
	"discovery_basic_poll_settings":        types.ObjectType{AttrTypes: Ipv6networkcontainerDiscoveryBasicPollSettingsAttrTypes},
	"discovery_blackout_setting":           types.ObjectType{AttrTypes: Ipv6networkcontainerDiscoveryBlackoutSettingAttrTypes},
	"discovery_engine_type":                types.StringType,
	"discovery_member":                     types.StringType,
	"domain_name_servers":                  types.ListType{ElemType: types.StringType},
	"enable_ddns":                          types.BoolType,
	"enable_discovery":                     types.BoolType,
	"enable_immediate_discovery":           types.BoolType,
	"endpoint_sources":                     types.ListType{ElemType: types.StringType},
	"extattrs":                             types.MapType{ElemType: types.StringType},
	"extattrs_all":                         types.MapType{ElemType: types.StringType},
	"federated_realms":                     types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6networkcontainerFederatedRealmsAttrTypes}},
	"last_rir_registration_update_sent":    types.Int64Type,
	"last_rir_registration_update_status":  types.StringType,
	"logic_filter_rules":                   types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6networkcontainerLogicFilterRulesAttrTypes}},
	"mgm_private":                          types.BoolType,
	"mgm_private_overridable":              types.BoolType,
	"ms_ad_user_data":                      types.ObjectType{AttrTypes: Ipv6networkcontainerMsAdUserDataAttrTypes},
	"network":                              types.StringType,
	"func_call":                            types.ObjectType{AttrTypes: FuncCallAttrTypes},
	"network_container":                    types.StringType,
	"network_view":                         types.StringType,
	"options":                              internaltypes.UnorderedList{ListType: basetypes.ListType{ElemType: basetypes.ObjectType{AttrTypes: Ipv6networkcontainerOptionsAttrTypes}}},
	"port_control_blackout_setting":        types.ObjectType{AttrTypes: Ipv6networkcontainerPortControlBlackoutSettingAttrTypes},
	"preferred_lifetime":                   types.Int64Type,
	"remove_subnets":                       types.BoolType,
	"restart_if_needed":                    types.BoolType,
	"rir":                                  types.StringType,
	"rir_organization":                     types.StringType,
	"rir_registration_action":              types.StringType,
	"rir_registration_status":              types.StringType,
	"same_port_control_discovery_blackout": types.BoolType,
	"send_rir_request":                     types.BoolType,
	"subscribe_settings":                   types.ObjectType{AttrTypes: Ipv6networkcontainerSubscribeSettingsAttrTypes},
	"unmanaged":                            types.BoolType,
	"update_dns_on_lease_renewal":          types.BoolType,
	"use_blackout_setting":                 types.BoolType,
	"use_ddns_domainname":                  types.BoolType,
	"use_ddns_enable_option_fqdn":          types.BoolType,
	"use_ddns_generate_hostname":           types.BoolType,
	"use_ddns_ttl":                         types.BoolType,
	"use_discovery_basic_polling_settings": types.BoolType,
	"use_domain_name_servers":              types.BoolType,
	"use_enable_ddns":                      types.BoolType,
	"use_enable_discovery":                 types.BoolType,
	"use_logic_filter_rules":               types.BoolType,
	"use_mgm_private":                      types.BoolType,
	"use_options":                          types.BoolType,
	"use_preferred_lifetime":               types.BoolType,
	"use_subscribe_settings":               types.BoolType,
	"use_update_dns_on_lease_renewal":      types.BoolType,
	"use_valid_lifetime":                   types.BoolType,
	"use_zone_associations":                types.BoolType,
	"utilization":                          types.Int64Type,
	"valid_lifetime":                       types.Int64Type,
	"zone_associations":                    types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6networkcontainerZoneAssociationsAttrTypes}},
}

var Ipv6networkcontainerResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		MarkdownDescription: "The reference to the object.",
		Computed:            true,
	},
	"auto_create_reversezone": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "This flag controls whether reverse zones are automatically created when the network is added.",
		Default:             booldefault.StaticBool(false),
		Computed:            true,
	},
	"cloud_info": schema.SingleNestedAttribute{
		Attributes:          Ipv6networkcontainerCloudInfoResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Structure containing all cloud API related information for this object.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Comment for the network; maximum 256 characters.",
		Computed:            true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Should not have leading or trailing whitespace",
			),
		},
	},
	"ddns_domainname": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The dynamic DNS domain name the appliance uses specifically for DDNS updates for this network container.",
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_ddns_domainname")),
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Should not have leading or trailing whitespace"),
		},
		Computed: true,
	},
	"ddns_enable_option_fqdn": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use this method to set or retrieve the ddns_enable_option_fqdn flag of a DHCP IPv6 Network Container object. This method controls whether the FQDN option sent by the client is to be used, or if the server can automatically generate the FQDN. This setting overrides the upper-level settings.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_ddns_enable_option_fqdn")),
		},
	},
	"ddns_generate_hostname": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If this field is set to True, the DHCP server generates a hostname and updates DNS with it when the DHCP client request does not contain a hostname.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_ddns_generate_hostname")),
		},
	},
	"ddns_server_always_updates": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "This field controls whether the DHCP server is allowed to update DNS, regardless of the DHCP client requests. Note that changes for this field take effect only if ddns_enable_option_fqdn is True.",
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("ddns_enable_option_fqdn")),
		},
	},
	"ddns_ttl": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The DNS update Time to Live (TTL) value of a DHCP network container object. The TTL is a 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached.",
		Computed:            true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_ddns_ttl")),
		},
		Default: int64default.StaticInt64(0),
	},
	"delete_reason": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reason for deleting the RIR registration request.",
	},
	"discover_now_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Discover now status for this network container.",
	},
	"discovery_basic_poll_settings": schema.SingleNestedAttribute{
		Attributes: Ipv6networkcontainerDiscoveryBasicPollSettingsResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Validators: []validator.Object{
			objectvalidator.AlsoRequires(path.MatchRoot("use_discovery_basic_polling_settings")),
		},
	},
	"discovery_blackout_setting": schema.SingleNestedAttribute{
		Attributes: Ipv6networkcontainerDiscoveryBlackoutSettingResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Validators: []validator.Object{
			objectvalidator.AlsoRequires(path.MatchRoot("use_blackout_setting")),
		},
		MarkdownDescription: "The discovery blackout setting for this network container.",
	},
	"discovery_engine_type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The network discovery engine type.",
	},
	"discovery_member": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The member that will run discovery for this network container.",
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_enable_discovery")),
		},
		Computed: true,
	},
	"domain_name_servers": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "Use this method to set or retrieve the dynamic DNS updates flag of a DHCP IPv6 Network Container object. The DHCP server can send DDNS updates to DNS servers in the same Grid and to external DNS servers. This setting overrides the member level settings.",
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_domain_name_servers")),
		},
	},
	"enable_ddns": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "The dynamic DNS updates flag of a DHCP IPv6 network container object. If set to True, the DHCP server sends DDNS updates to DNS servers in the same Grid, and to external DNS servers.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_enable_ddns")),
		},
	},
	"enable_discovery": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether a discovery is enabled or not for this network container. When this is set to False, the network container discovery is disabled.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_enable_discovery")),
		},
	},
	"enable_immediate_discovery": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the discovery for the network container should be immediately enabled.",
	},
	"endpoint_sources": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The endpoints that provides data for the DHCP IPv6 Network Container.",
	},
	"extattrs": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.",
		Computed:            true,
		Default:             mapdefault.StaticValue(types.MapNull(types.StringType)),
	},
	"extattrs_all": schema.MapAttribute{
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object , including default attributes.",
		ElementType:         types.StringType,
	},
	"federated_realms": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6networkcontainerFederatedRealmsResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "This field contains the federated realms associated to this network container.",
	},
	"last_rir_registration_update_sent": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The timestamp when the last RIR registration update was sent.",
	},
	"last_rir_registration_update_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Last RIR registration update status.",
	},
	"logic_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6networkcontainerLogicFilterRulesResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "This field contains the logic filters to be applied on the this network container. This list corresponds to the match rules that are written to the dhcpd configuration file.",
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_logic_filter_rules")),
		},
	},
	"mgm_private": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "This field controls whether this object is synchronized with the Multi-Grid Master. If this field is set to True, objects are not synchronized.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_mgm_private")),
		},
	},
	"mgm_private_overridable": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "This field is assumed to be True unless filled by any conforming objects, such as Network, IPv6 Network, Network Container, IPv6 Network Container, and Network View. This value is set to False if mgm_private is set to True in the parent object.",
	},
	"ms_ad_user_data": schema.SingleNestedAttribute{
		Attributes:          Ipv6networkcontainerMsAdUserDataResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "The Microsoft Active Directory user data associated with the network container.",
	},
	"network": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The network address in IPv6 Address/CIDR format. For regular expression searches, only the IPv6 Address portion is supported. Searches for the CIDR portion is always an exact match. For example, both network containers 16::0/28 and 26::0/24 are matched by expression '.6' and only 26::0/24 is matched by '.6/24'.",
		Computed:            true,
	},
	"func_call": schema.SingleNestedAttribute{
		Computed:            true,
		Attributes:          FuncCallResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "The function call to be executed on the object.",
	},
	"network_container": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The network container to which this network belongs, if any.",
	},
	"network_view": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the network view in which this network resides.",
		Computed:            true,
		Default:             stringdefault.StaticString("default"),
	},
	"options": schema.ListNestedAttribute{
		CustomType: internaltypes.UnorderedList{ListType: basetypes.ListType{ElemType: basetypes.ObjectType{AttrTypes: Ipv6networkcontainerOptionsAttrTypes}}},
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6networkcontainerOptionsResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object.",
		Computed:            true,
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_options")),
		},
	},
	"port_control_blackout_setting": schema.SingleNestedAttribute{
		Attributes: Ipv6networkcontainerPortControlBlackoutSettingResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Validators: []validator.Object{
			objectvalidator.AlsoRequires(path.MatchRoot("use_blackout_setting")),
		},
	},
	"preferred_lifetime": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Use this method to set or retrieve the preferred lifetime value of a DHCP IPv6 Network Container object.",
		Computed:            true,
		Default:             int64default.StaticInt64(27000),
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_preferred_lifetime")),
		},
	},
	"remove_subnets": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Remove subnets delete option. Determines whether all child objects should be removed alongside with the IPv6 network container or child objects should be assigned to another parental container. By default child objects are deleted with this network container.",
	},
	"restart_if_needed": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Restarts the member service.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"rir": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The registry (RIR) that allocated the IPv6 network container address space.",
	},
	"rir_organization": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The RIR organization associated with the IPv6 network container.",
		Computed:            true,
	},
	"rir_registration_action": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The RIR registration action.",
		Computed:            true,
	},
	"rir_registration_status": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The registration status of the IPv6 network container in RIR.",
		Computed:            true,
		Default:             stringdefault.StaticString("NOT_REGISTERED"),
		Validators: []validator.String{
			stringvalidator.OneOf(
				"NOT_REGISTERED",
				"REGISTERED",
			),
		},
	},
	"same_port_control_discovery_blackout": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If the field is set to True, the discovery blackout setting will be used for port control blackout setting.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_blackout_setting")),
		},
	},
	"send_rir_request": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether to send the RIR registration request.",
	},
	"subscribe_settings": schema.SingleNestedAttribute{
		Attributes: Ipv6networkcontainerSubscribeSettingsResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Validators: []validator.Object{
			objectvalidator.AlsoRequires(path.MatchRoot("use_subscribe_settings")),
		},
	},
	"unmanaged": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the network container is unmanaged or not.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"update_dns_on_lease_renewal": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "This field controls whether the DHCP server updates DNS when a DHCP lease is renewed.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_update_dns_on_lease_renewal")),
		},
	},
	"use_blackout_setting": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: discovery_blackout_setting , port_control_blackout_setting, same_port_control_discovery_blackout",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_ddns_domainname": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ddns_domainname",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_ddns_enable_option_fqdn": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ddns_enable_option_fqdn",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_ddns_generate_hostname": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ddns_generate_hostname",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_ddns_ttl": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ddns_ttl",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_discovery_basic_polling_settings": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: discovery_basic_poll_settings",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_domain_name_servers": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: domain_name_servers",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_enable_ddns": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: enable_ddns",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_enable_discovery": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: discovery_member , enable_discovery",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_logic_filter_rules": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: logic_filter_rules",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_mgm_private": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: mgm_private",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_options": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: options",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_preferred_lifetime": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: preferred_lifetime",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_subscribe_settings": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: subscribe_settings",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_update_dns_on_lease_renewal": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: update_dns_on_lease_renewal",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_valid_lifetime": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: valid_lifetime",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_zone_associations": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: zone_associations",
		Computed:            true,
		Default:             booldefault.StaticBool(true),
	},
	"utilization": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The network container utilization in percentage.",
	},
	"valid_lifetime": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Use this method to set or retrieve the valid lifetime value of a DHCP IPv6 Network Container object.",
		Computed:            true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_valid_lifetime")),
		},
		Default: int64default.StaticInt64(43200),
	},
	"zone_associations": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6networkcontainerZoneAssociationsResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "The list of zones associated with this network container.",
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_zone_associations")),
		},
	},
}

func (m *Ipv6networkcontainerModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *ipam.Ipv6networkcontainer {
	if m == nil {
		return nil
	}
	to := &ipam.Ipv6networkcontainer{
		CloudInfo:                        ExpandIpv6networkcontainerCloudInfo(ctx, m.CloudInfo, diags),
		Comment:                          flex.ExpandStringPointer(m.Comment),
		DdnsDomainname:                   flex.ExpandStringPointer(m.DdnsDomainname),
		DdnsEnableOptionFqdn:             flex.ExpandBoolPointer(m.DdnsEnableOptionFqdn),
		DdnsGenerateHostname:             flex.ExpandBoolPointer(m.DdnsGenerateHostname),
		DdnsServerAlwaysUpdates:          flex.ExpandBoolPointer(m.DdnsServerAlwaysUpdates),
		DdnsTtl:                          flex.ExpandInt64Pointer(m.DdnsTtl),
		DeleteReason:                     flex.ExpandStringPointer(m.DeleteReason),
		DiscoveryBasicPollSettings:       ExpandIpv6networkcontainerDiscoveryBasicPollSettings(ctx, m.DiscoveryBasicPollSettings, diags),
		DiscoveryBlackoutSetting:         ExpandIpv6networkcontainerDiscoveryBlackoutSetting(ctx, m.DiscoveryBlackoutSetting, diags),
		DiscoveryMember:                  flex.ExpandStringPointer(m.DiscoveryMember),
		DomainNameServers:                flex.ExpandFrameworkListString(ctx, m.DomainNameServers, diags),
		EnableDdns:                       flex.ExpandBoolPointer(m.EnableDdns),
		EnableDiscovery:                  flex.ExpandBoolPointer(m.EnableDiscovery),
		EnableImmediateDiscovery:         flex.ExpandBoolPointer(m.EnableImmediateDiscovery),
		ExtAttrs:                         ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		FederatedRealms:                  flex.ExpandFrameworkListNestedBlock(ctx, m.FederatedRealms, diags, ExpandIpv6networkcontainerFederatedRealms),
		LogicFilterRules:                 flex.ExpandFrameworkListNestedBlock(ctx, m.LogicFilterRules, diags, ExpandIpv6networkcontainerLogicFilterRules),
		MgmPrivate:                       flex.ExpandBoolPointer(m.MgmPrivate),
		MsAdUserData:                     ExpandIpv6networkcontainerMsAdUserData(ctx, m.MsAdUserData, diags),
		Network:                          ExpandIpv6NetworkcontainerNetwork(m.Network),
		FuncCall:                         ExpandFuncCall(ctx, m.FuncCall, diags),
		NetworkView:                      flex.ExpandStringPointer(m.NetworkView),
		Options:                          flex.ExpandFrameworkListNestedBlock(ctx, m.Options, diags, ExpandIpv6networkcontainerOptions),
		PortControlBlackoutSetting:       ExpandIpv6networkcontainerPortControlBlackoutSetting(ctx, m.PortControlBlackoutSetting, diags),
		PreferredLifetime:                flex.ExpandInt64Pointer(m.PreferredLifetime),
		RemoveSubnets:                    flex.ExpandBoolPointer(m.RemoveSubnets),
		RestartIfNeeded:                  flex.ExpandBoolPointer(m.RestartIfNeeded),
		RirOrganization:                  flex.ExpandStringPointer(m.RirOrganization),
		RirRegistrationAction:            flex.ExpandStringPointer(m.RirRegistrationAction),
		RirRegistrationStatus:            flex.ExpandStringPointer(m.RirRegistrationStatus),
		SamePortControlDiscoveryBlackout: flex.ExpandBoolPointer(m.SamePortControlDiscoveryBlackout),
		SendRirRequest:                   flex.ExpandBoolPointer(m.SendRirRequest),
		SubscribeSettings:                ExpandIpv6networkcontainerSubscribeSettings(ctx, m.SubscribeSettings, diags),
		Unmanaged:                        flex.ExpandBoolPointer(m.Unmanaged),
		UpdateDnsOnLeaseRenewal:          flex.ExpandBoolPointer(m.UpdateDnsOnLeaseRenewal),
		UseBlackoutSetting:               flex.ExpandBoolPointer(m.UseBlackoutSetting),
		UseDdnsDomainname:                flex.ExpandBoolPointer(m.UseDdnsDomainname),
		UseDdnsEnableOptionFqdn:          flex.ExpandBoolPointer(m.UseDdnsEnableOptionFqdn),
		UseDdnsGenerateHostname:          flex.ExpandBoolPointer(m.UseDdnsGenerateHostname),
		UseDdnsTtl:                       flex.ExpandBoolPointer(m.UseDdnsTtl),
		UseDiscoveryBasicPollingSettings: flex.ExpandBoolPointer(m.UseDiscoveryBasicPollingSettings),
		UseDomainNameServers:             flex.ExpandBoolPointer(m.UseDomainNameServers),
		UseEnableDdns:                    flex.ExpandBoolPointer(m.UseEnableDdns),
		UseEnableDiscovery:               flex.ExpandBoolPointer(m.UseEnableDiscovery),
		UseLogicFilterRules:              flex.ExpandBoolPointer(m.UseLogicFilterRules),
		UseMgmPrivate:                    flex.ExpandBoolPointer(m.UseMgmPrivate),
		UseOptions:                       flex.ExpandBoolPointer(m.UseOptions),
		UsePreferredLifetime:             flex.ExpandBoolPointer(m.UsePreferredLifetime),
		UseSubscribeSettings:             flex.ExpandBoolPointer(m.UseSubscribeSettings),
		UseUpdateDnsOnLeaseRenewal:       flex.ExpandBoolPointer(m.UseUpdateDnsOnLeaseRenewal),
		UseValidLifetime:                 flex.ExpandBoolPointer(m.UseValidLifetime),
		UseZoneAssociations:              flex.ExpandBoolPointer(m.UseZoneAssociations),
		ValidLifetime:                    flex.ExpandInt64Pointer(m.ValidLifetime),
		ZoneAssociations:                 flex.ExpandFrameworkListNestedBlock(ctx, m.ZoneAssociations, diags, ExpandIpv6networkcontainerZoneAssociations),
	}
	if isCreate {
		to.NetworkContainer = flex.ExpandStringPointer(m.NetworkContainer)
		to.NetworkView = flex.ExpandStringPointer(m.NetworkView)
		to.Network = ExpandIpv6NetworkcontainerNetwork(m.Network)
		to.AutoCreateReversezone = flex.ExpandBoolPointer(m.AutoCreateReversezone)
	}
	return to
}

func FlattenIpv6networkcontainer(ctx context.Context, from *ipam.Ipv6networkcontainer, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6networkcontainerAttrTypes)
	}
	m := Ipv6networkcontainerModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrs = m.ExtAttrsAll
	t, d := types.ObjectValueFrom(ctx, Ipv6networkcontainerAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Ipv6networkcontainerModel) Flatten(ctx context.Context, from *ipam.Ipv6networkcontainer, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Ipv6networkcontainerModel{}
	}
	// from.Options = RemoveDefaultDHCPOptions(ctx, diags, from.Options, m.Options)

	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.CloudInfo = FlattenIpv6networkcontainerCloudInfo(ctx, from.CloudInfo, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.DdnsDomainname = flex.FlattenStringPointer(from.DdnsDomainname)
	m.DdnsEnableOptionFqdn = types.BoolPointerValue(from.DdnsEnableOptionFqdn)
	m.DdnsGenerateHostname = types.BoolPointerValue(from.DdnsGenerateHostname)
	m.DdnsServerAlwaysUpdates = types.BoolPointerValue(from.DdnsServerAlwaysUpdates)
	m.DdnsTtl = flex.FlattenInt64Pointer(from.DdnsTtl)
	m.DiscoverNowStatus = flex.FlattenStringPointer(from.DiscoverNowStatus)
	m.DiscoveryBasicPollSettings = FlattenIpv6networkcontainerDiscoveryBasicPollSettings(ctx, from.DiscoveryBasicPollSettings, diags)
	m.DiscoveryBlackoutSetting = FlattenIpv6networkcontainerDiscoveryBlackoutSetting(ctx, from.DiscoveryBlackoutSetting, diags)
	m.DiscoveryEngineType = flex.FlattenStringPointer(from.DiscoveryEngineType)
	m.DiscoveryMember = flex.FlattenStringPointer(from.DiscoveryMember)
	m.DomainNameServers = flex.FlattenFrameworkListString(ctx, from.DomainNameServers, diags)
	m.EnableDdns = types.BoolPointerValue(from.EnableDdns)
	m.EnableDiscovery = types.BoolPointerValue(from.EnableDiscovery)
	m.EnableImmediateDiscovery = types.BoolPointerValue(from.EnableImmediateDiscovery)
	m.EndpointSources = flex.FlattenFrameworkListString(ctx, from.EndpointSources, diags)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.FederatedRealms = flex.FlattenFrameworkListNestedBlock(ctx, from.FederatedRealms, Ipv6networkcontainerFederatedRealmsAttrTypes, diags, FlattenIpv6networkcontainerFederatedRealms)
	m.LastRirRegistrationUpdateSent = flex.FlattenInt64Pointer(from.LastRirRegistrationUpdateSent)
	m.LastRirRegistrationUpdateStatus = flex.FlattenStringPointer(from.LastRirRegistrationUpdateStatus)
	m.LogicFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.LogicFilterRules, Ipv6networkcontainerLogicFilterRulesAttrTypes, diags, FlattenIpv6networkcontainerLogicFilterRules)
	m.MgmPrivate = types.BoolPointerValue(from.MgmPrivate)
	m.MgmPrivateOverridable = types.BoolPointerValue(from.MgmPrivateOverridable)
	m.MsAdUserData = FlattenIpv6networkcontainerMsAdUserData(ctx, from.MsAdUserData, diags)
	m.Network = FlattenIpv6NetworkcontainerNetwork(from.Network)
	if m.FuncCall.IsNull() || m.FuncCall.IsUnknown() {
		m.FuncCall = FlattenFuncCall(ctx, from.FuncCall, diags)
	}
	m.NetworkContainer = flex.FlattenStringPointer(from.NetworkContainer)
	m.NetworkView = flex.FlattenStringPointer(from.NetworkView)
	m.Options = RemoveIpv6networkcontainerDefaultDHCPOptions(ctx, diags, from.Options, m.Options)
	m.PortControlBlackoutSetting = FlattenIpv6networkcontainerPortControlBlackoutSetting(ctx, from.PortControlBlackoutSetting, diags)
	m.PreferredLifetime = flex.FlattenInt64Pointer(from.PreferredLifetime)
	m.RemoveSubnets = types.BoolPointerValue(from.RemoveSubnets)
	m.Rir = flex.FlattenStringPointer(from.Rir)
	m.RirOrganization = flex.FlattenStringPointer(from.RirOrganization)
	m.RirRegistrationAction = flex.FlattenStringPointer(from.RirRegistrationAction)
	m.RirRegistrationStatus = flex.FlattenStringPointer(from.RirRegistrationStatus)
	m.SamePortControlDiscoveryBlackout = types.BoolPointerValue(from.SamePortControlDiscoveryBlackout)
	m.SendRirRequest = types.BoolPointerValue(from.SendRirRequest)
	m.SubscribeSettings = FlattenIpv6networkcontainerSubscribeSettings(ctx, from.SubscribeSettings, diags)
	m.Unmanaged = types.BoolPointerValue(from.Unmanaged)
	m.UpdateDnsOnLeaseRenewal = types.BoolPointerValue(from.UpdateDnsOnLeaseRenewal)
	m.UseBlackoutSetting = types.BoolPointerValue(from.UseBlackoutSetting)
	m.UseDdnsDomainname = types.BoolPointerValue(from.UseDdnsDomainname)
	m.UseDdnsEnableOptionFqdn = types.BoolPointerValue(from.UseDdnsEnableOptionFqdn)
	m.UseDdnsGenerateHostname = types.BoolPointerValue(from.UseDdnsGenerateHostname)
	m.UseDdnsTtl = types.BoolPointerValue(from.UseDdnsTtl)
	m.UseDiscoveryBasicPollingSettings = types.BoolPointerValue(from.UseDiscoveryBasicPollingSettings)
	m.UseDomainNameServers = types.BoolPointerValue(from.UseDomainNameServers)
	m.UseEnableDdns = types.BoolPointerValue(from.UseEnableDdns)
	m.UseEnableDiscovery = types.BoolPointerValue(from.UseEnableDiscovery)
	m.UseLogicFilterRules = types.BoolPointerValue(from.UseLogicFilterRules)
	m.UseMgmPrivate = types.BoolPointerValue(from.UseMgmPrivate)
	m.UseOptions = types.BoolPointerValue(from.UseOptions)
	m.UsePreferredLifetime = types.BoolPointerValue(from.UsePreferredLifetime)
	m.UseSubscribeSettings = types.BoolPointerValue(from.UseSubscribeSettings)
	m.UseUpdateDnsOnLeaseRenewal = types.BoolPointerValue(from.UseUpdateDnsOnLeaseRenewal)
	m.UseValidLifetime = types.BoolPointerValue(from.UseValidLifetime)
	m.UseZoneAssociations = types.BoolPointerValue(from.UseZoneAssociations)
	m.Utilization = flex.FlattenInt64Pointer(from.Utilization)
	m.ValidLifetime = flex.FlattenInt64Pointer(from.ValidLifetime)
	m.ZoneAssociations = flex.FlattenFrameworkListNestedBlock(ctx, from.ZoneAssociations, Ipv6networkcontainerZoneAssociationsAttrTypes, diags, FlattenIpv6networkcontainerZoneAssociations)
}

func ExpandIpv6NetworkcontainerNetwork(str types.String) *ipam.Ipv6networkcontainerNetwork {
	if str.IsNull() {
		return &ipam.Ipv6networkcontainerNetwork{}
	}
	var m ipam.Ipv6networkcontainerNetwork
	m.String = flex.ExpandStringPointer(str)
	return &m
}

func FlattenIpv6NetworkcontainerNetwork(from *ipam.Ipv6networkcontainerNetwork) types.String {
	if from.String == nil {
		return types.StringNull()
	}
	m := flex.FlattenStringPointer(from.String)
	return m
}

func RemoveIpv6networkcontainerDefaultDHCPOptions(ctx context.Context, diags *diag.Diagnostics, options []ipam.Ipv6networkcontainerOptions, planOptions internaltypes.UnorderedListValue) internaltypes.UnorderedListValue {
	defaultOptionName := "dhcp-lease-time"

	// If no options, return empty list
	if len(options) == 0 {
		return internaltypes.NewUnorderedListValueNull(types.ObjectType{AttrTypes: Ipv6networkcontainerOptionsAttrTypes})
	}

	// If plan options is null or unknown, return original options
	if planOptions.IsNull() || planOptions.IsUnknown() {
		return flex.FlattenFrameworkUnorderedListNestedBlock(ctx, options, Ipv6networkcontainerOptionsAttrTypes, diags, FlattenIpv6networkcontainerOptions)
	}

	// Convert plan options to a map for easy lookup
	baseList, err := planOptions.ToListValue(ctx)
	if err != nil {
		return flex.FlattenFrameworkUnorderedListNestedBlock(ctx, options, Ipv6networkcontainerOptionsAttrTypes, diags, FlattenIpv6networkcontainerOptions)
	}

	planOptionsArr := flex.ExpandFrameworkListNestedBlock(ctx, baseList, diags, ExpandIpv6networkcontainerOptions)
	planOptionsMap := make(map[string]ipam.Ipv6networkcontainerOptions)
	var planOrder []string
	for _, opt := range planOptionsArr {
		if opt.Name != nil {
			planOptionsMap[*opt.Name] = opt
			planOrder = append(planOrder, *opt.Name)
		}
	}

	// Convert current options to a map
	currentOptionsMap := make(map[string]ipam.Ipv6networkcontainerOptions)
	for _, opt := range options {
		if opt.Name != nil {
			currentOptionsMap[*opt.Name] = opt
		}
	}

	// Build result maintaining plan order
	var result []ipam.Ipv6networkcontainerOptions
	for _, name := range planOrder {
		if name == defaultOptionName {
			// For lease-time option, check if values match
			planOpt, planExists := planOptionsMap[name]
			currentOpt, currentExists := currentOptionsMap[name]

			if planExists && currentExists &&
				planOpt.Value != nil && currentOpt.Value != nil &&
				*planOpt.Value == *currentOpt.Value {
				result = append(result, currentOpt)
			}
		} else {
			// For non-lease-time options, use current value if exists
			if opt, exists := currentOptionsMap[name]; exists {
				result = append(result, opt)
			}
		}
	}

	// Add any remaining options that weren't in the plan but should be kept
	for _, opt := range options {
		if opt.Name == nil {
			continue
		}
		_, inPlan := planOptionsMap[*opt.Name]
		if !inPlan && *opt.Name != defaultOptionName {
			result = append(result, opt)
		}
	}

	return flex.FlattenFrameworkUnorderedListNestedBlock(ctx, result, Ipv6networkcontainerOptionsAttrTypes, diags, FlattenIpv6networkcontainerOptions)
}
