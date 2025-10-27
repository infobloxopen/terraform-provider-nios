package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/cidrtypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/boolvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	planmodifiers "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/immutable"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type Ipv6networkModel struct {
	Ref                              types.String         `tfsdk:"ref"`
	AutoCreateReversezone            types.Bool           `tfsdk:"auto_create_reversezone"`
	CloudInfo                        types.Object         `tfsdk:"cloud_info"`
	Comment                          types.String         `tfsdk:"comment"`
	DdnsDomainname                   types.String         `tfsdk:"ddns_domainname"`
	DdnsEnableOptionFqdn             types.Bool           `tfsdk:"ddns_enable_option_fqdn"`
	DdnsGenerateHostname             types.Bool           `tfsdk:"ddns_generate_hostname"`
	DdnsServerAlwaysUpdates          types.Bool           `tfsdk:"ddns_server_always_updates"`
	DdnsTtl                          types.Int64          `tfsdk:"ddns_ttl"`
	DeleteReason                     types.String         `tfsdk:"delete_reason"`
	Disable                          types.Bool           `tfsdk:"disable"`
	DiscoverNowStatus                types.String         `tfsdk:"discover_now_status"`
	DiscoveredBgpAs                  types.String         `tfsdk:"discovered_bgp_as"`
	DiscoveredBridgeDomain           types.String         `tfsdk:"discovered_bridge_domain"`
	DiscoveredTenant                 types.String         `tfsdk:"discovered_tenant"`
	DiscoveredVlanId                 types.String         `tfsdk:"discovered_vlan_id"`
	DiscoveredVlanName               types.String         `tfsdk:"discovered_vlan_name"`
	DiscoveredVrfDescription         types.String         `tfsdk:"discovered_vrf_description"`
	DiscoveredVrfName                types.String         `tfsdk:"discovered_vrf_name"`
	DiscoveredVrfRd                  types.String         `tfsdk:"discovered_vrf_rd"`
	DiscoveryBasicPollSettings       types.Object         `tfsdk:"discovery_basic_poll_settings"`
	DiscoveryBlackoutSetting         types.Object         `tfsdk:"discovery_blackout_setting"`
	DiscoveryEngineType              types.String         `tfsdk:"discovery_engine_type"`
	DiscoveryMember                  types.String         `tfsdk:"discovery_member"`
	DomainName                       types.String         `tfsdk:"domain_name"`
	DomainNameServers                types.List           `tfsdk:"domain_name_servers"`
	EnableDdns                       types.Bool           `tfsdk:"enable_ddns"`
	EnableDiscovery                  types.Bool           `tfsdk:"enable_discovery"`
	EnableIfmapPublishing            types.Bool           `tfsdk:"enable_ifmap_publishing"`
	EnableImmediateDiscovery         types.Bool           `tfsdk:"enable_immediate_discovery"`
	EndpointSources                  types.List           `tfsdk:"endpoint_sources"`
	ExtAttrs                         types.Map            `tfsdk:"extattrs"`
	ExtAttrsAll                      types.Map            `tfsdk:"extattrs_all"`
	FederatedRealms                  types.List           `tfsdk:"federated_realms"`
	LastRirRegistrationUpdateSent    types.Int64          `tfsdk:"last_rir_registration_update_sent"`
	LastRirRegistrationUpdateStatus  types.String         `tfsdk:"last_rir_registration_update_status"`
	LogicFilterRules                 types.List           `tfsdk:"logic_filter_rules"`
	Members                          types.List           `tfsdk:"members"`
	MgmPrivate                       types.Bool           `tfsdk:"mgm_private"`
	MgmPrivateOverridable            types.Bool           `tfsdk:"mgm_private_overridable"`
	MsAdUserData                     types.Object         `tfsdk:"ms_ad_user_data"`
	Network                          cidrtypes.IPv6Prefix `tfsdk:"network"`
	FuncCall                         types.Object         `tfsdk:"func_call"`
	NetworkContainer                 types.String         `tfsdk:"network_container"`
	NetworkView                      types.String         `tfsdk:"network_view"`
	Options                          types.List           `tfsdk:"options"`
	PortControlBlackoutSetting       types.Object         `tfsdk:"port_control_blackout_setting"`
	PreferredLifetime                types.Int64          `tfsdk:"preferred_lifetime"`
	RecycleLeases                    types.Bool           `tfsdk:"recycle_leases"`
	RestartIfNeeded                  types.Bool           `tfsdk:"restart_if_needed"`
	Rir                              types.String         `tfsdk:"rir"`
	RirOrganization                  types.String         `tfsdk:"rir_organization"`
	RirRegistrationAction            types.String         `tfsdk:"rir_registration_action"`
	RirRegistrationStatus            types.String         `tfsdk:"rir_registration_status"`
	SamePortControlDiscoveryBlackout types.Bool           `tfsdk:"same_port_control_discovery_blackout"`
	SendRirRequest                   types.Bool           `tfsdk:"send_rir_request"`
	SubscribeSettings                types.Object         `tfsdk:"subscribe_settings"`
	Template                         types.String         `tfsdk:"template"`
	Unmanaged                        types.Bool           `tfsdk:"unmanaged"`
	UnmanagedCount                   types.Int64          `tfsdk:"unmanaged_count"`
	UpdateDnsOnLeaseRenewal          types.Bool           `tfsdk:"update_dns_on_lease_renewal"`
	UseBlackoutSetting               types.Bool           `tfsdk:"use_blackout_setting"`
	UseDdnsDomainname                types.Bool           `tfsdk:"use_ddns_domainname"`
	UseDdnsEnableOptionFqdn          types.Bool           `tfsdk:"use_ddns_enable_option_fqdn"`
	UseDdnsGenerateHostname          types.Bool           `tfsdk:"use_ddns_generate_hostname"`
	UseDdnsTtl                       types.Bool           `tfsdk:"use_ddns_ttl"`
	UseDiscoveryBasicPollingSettings types.Bool           `tfsdk:"use_discovery_basic_polling_settings"`
	UseDomainName                    types.Bool           `tfsdk:"use_domain_name"`
	UseDomainNameServers             types.Bool           `tfsdk:"use_domain_name_servers"`
	UseEnableDdns                    types.Bool           `tfsdk:"use_enable_ddns"`
	UseEnableDiscovery               types.Bool           `tfsdk:"use_enable_discovery"`
	UseEnableIfmapPublishing         types.Bool           `tfsdk:"use_enable_ifmap_publishing"`
	UseLogicFilterRules              types.Bool           `tfsdk:"use_logic_filter_rules"`
	UseMgmPrivate                    types.Bool           `tfsdk:"use_mgm_private"`
	UseOptions                       types.Bool           `tfsdk:"use_options"`
	UsePreferredLifetime             types.Bool           `tfsdk:"use_preferred_lifetime"`
	UseRecycleLeases                 types.Bool           `tfsdk:"use_recycle_leases"`
	UseSubscribeSettings             types.Bool           `tfsdk:"use_subscribe_settings"`
	UseUpdateDnsOnLeaseRenewal       types.Bool           `tfsdk:"use_update_dns_on_lease_renewal"`
	UseValidLifetime                 types.Bool           `tfsdk:"use_valid_lifetime"`
	UseZoneAssociations              types.Bool           `tfsdk:"use_zone_associations"`
	ValidLifetime                    types.Int64          `tfsdk:"valid_lifetime"`
	Vlans                            types.List           `tfsdk:"vlans"`
	ZoneAssociations                 types.List           `tfsdk:"zone_associations"`
}

var Ipv6networkAttrTypes = map[string]attr.Type{
	"ref":                                  types.StringType,
	"auto_create_reversezone":              types.BoolType,
	"cloud_info":                           types.ObjectType{AttrTypes: Ipv6networkCloudInfoAttrTypes},
	"comment":                              types.StringType,
	"ddns_domainname":                      types.StringType,
	"ddns_enable_option_fqdn":              types.BoolType,
	"ddns_generate_hostname":               types.BoolType,
	"ddns_server_always_updates":           types.BoolType,
	"ddns_ttl":                             types.Int64Type,
	"delete_reason":                        types.StringType,
	"disable":                              types.BoolType,
	"discover_now_status":                  types.StringType,
	"discovered_bgp_as":                    types.StringType,
	"discovered_bridge_domain":             types.StringType,
	"discovered_tenant":                    types.StringType,
	"discovered_vlan_id":                   types.StringType,
	"discovered_vlan_name":                 types.StringType,
	"discovered_vrf_description":           types.StringType,
	"discovered_vrf_name":                  types.StringType,
	"discovered_vrf_rd":                    types.StringType,
	"discovery_basic_poll_settings":        types.ObjectType{AttrTypes: Ipv6networkDiscoveryBasicPollSettingsAttrTypes},
	"discovery_blackout_setting":           types.ObjectType{AttrTypes: Ipv6networkDiscoveryBlackoutSettingAttrTypes},
	"discovery_engine_type":                types.StringType,
	"discovery_member":                     types.StringType,
	"domain_name":                          types.StringType,
	"domain_name_servers":                  types.ListType{ElemType: types.StringType},
	"enable_ddns":                          types.BoolType,
	"enable_discovery":                     types.BoolType,
	"enable_ifmap_publishing":              types.BoolType,
	"enable_immediate_discovery":           types.BoolType,
	"endpoint_sources":                     types.ListType{ElemType: types.StringType},
	"extattrs":                             types.MapType{ElemType: types.StringType},
	"extattrs_all":                         types.MapType{ElemType: types.StringType},
	"federated_realms":                     types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6networkFederatedRealmsAttrTypes}},
	"last_rir_registration_update_sent":    types.Int64Type,
	"last_rir_registration_update_status":  types.StringType,
	"logic_filter_rules":                   types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6networkLogicFilterRulesAttrTypes}},
	"members":                              types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6networkMembersAttrTypes}},
	"mgm_private":                          types.BoolType,
	"mgm_private_overridable":              types.BoolType,
	"ms_ad_user_data":                      types.ObjectType{AttrTypes: Ipv6networkMsAdUserDataAttrTypes},
	"network":                              cidrtypes.IPv6PrefixType{},
	"func_call":                            types.ObjectType{AttrTypes: FuncCallAttrTypes},
	"network_container":                    types.StringType,
	"network_view":                         types.StringType,
	"options":                              types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6networkOptionsAttrTypes}},
	"port_control_blackout_setting":        types.ObjectType{AttrTypes: Ipv6networkPortControlBlackoutSettingAttrTypes},
	"preferred_lifetime":                   types.Int64Type,
	"recycle_leases":                       types.BoolType,
	"restart_if_needed":                    types.BoolType,
	"rir":                                  types.StringType,
	"rir_organization":                     types.StringType,
	"rir_registration_action":              types.StringType,
	"rir_registration_status":              types.StringType,
	"same_port_control_discovery_blackout": types.BoolType,
	"send_rir_request":                     types.BoolType,
	"subscribe_settings":                   types.ObjectType{AttrTypes: Ipv6networkSubscribeSettingsAttrTypes},
	"template":                             types.StringType,
	"unmanaged":                            types.BoolType,
	"unmanaged_count":                      types.Int64Type,
	"update_dns_on_lease_renewal":          types.BoolType,
	"use_blackout_setting":                 types.BoolType,
	"use_ddns_domainname":                  types.BoolType,
	"use_ddns_enable_option_fqdn":          types.BoolType,
	"use_ddns_generate_hostname":           types.BoolType,
	"use_ddns_ttl":                         types.BoolType,
	"use_discovery_basic_polling_settings": types.BoolType,
	"use_domain_name":                      types.BoolType,
	"use_domain_name_servers":              types.BoolType,
	"use_enable_ddns":                      types.BoolType,
	"use_enable_discovery":                 types.BoolType,
	"use_enable_ifmap_publishing":          types.BoolType,
	"use_logic_filter_rules":               types.BoolType,
	"use_mgm_private":                      types.BoolType,
	"use_options":                          types.BoolType,
	"use_preferred_lifetime":               types.BoolType,
	"use_recycle_leases":                   types.BoolType,
	"use_subscribe_settings":               types.BoolType,
	"use_update_dns_on_lease_renewal":      types.BoolType,
	"use_valid_lifetime":                   types.BoolType,
	"use_zone_associations":                types.BoolType,
	"valid_lifetime":                       types.Int64Type,
	"vlans":                                types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6networkVlansAttrTypes}},
	"zone_associations":                    types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6networkZoneAssociationsAttrTypes}},
}

var Ipv6networkResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"auto_create_reversezone": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "This flag controls whether reverse zones are automatically created when the network is added.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		PlanModifiers: []planmodifier.Bool{
			planmodifiers.ImmutableBool(),
		},
	},
	"cloud_info": schema.SingleNestedAttribute{
		Optional:            true,
		Attributes:          Ipv6networkCloudInfoResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "Structure containing all cloud API related information for this object.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Comment for the network; maximum 256 characters.",
		Computed:            true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
	},
	"ddns_domainname": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The dynamic DNS domain name the appliance uses specifically for DDNS updates for this network.",
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_ddns_domainname")),
			customvalidator.ValidateTrimmedString(),
		},
		Computed: true,
	},
	"ddns_enable_option_fqdn": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use this method to set or retrieve the ddns_enable_option_fqdn flag of a DHCP IPv6 Network object. This method controls whether the FQDN option sent by the client is to be used, or if the server can automatically generate the FQDN. This setting overrides the upper-level settings.",
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
		MarkdownDescription: "This field controls whether only the DHCP server is allowed to update DNS, regardless of the DHCP clients requests. Note that changes for this field take effect only if ddns_enable_option_fqdn is True.",
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("ddns_enable_option_fqdn")),
		},
	},
	"ddns_ttl": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The DNS update Time to Live (TTL) value of a DHCP network object. The TTL is a 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached.",
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
	"disable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether a network is disabled or not. When this is set to False, the network is enabled.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"discover_now_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Discover now status for this network.",
		Default:             stringdefault.StaticString("NONE"),
	},
	"discovered_bgp_as": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Number of the discovered BGP AS. When multiple BGP autonomous systems are discovered in the network, this field displays \"Multiple\".",
	},
	"discovered_bridge_domain": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Discovered bridge domain.",
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
	},
	"discovered_tenant": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Discovered tenant.",
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
	},
	"discovered_vlan_id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The identifier of the discovered VLAN. When multiple VLANs are discovered in the network, this field displays \"Multiple\".",
	},
	"discovered_vlan_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the discovered VLAN. When multiple VLANs are discovered in the network, this field displays \"Multiple\".",
	},
	"discovered_vrf_description": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Description of the discovered VRF. When multiple VRFs are discovered in the network, this field displays \"Multiple\".",
	},
	"discovered_vrf_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the discovered VRF. When multiple VRFs are discovered in the network, this field displays \"Multiple\".",
	},
	"discovered_vrf_rd": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Route distinguisher of the discovered VRF. When multiple VRFs are discovered in the network, this field displays \"Multiple\".",
	},
	"discovery_basic_poll_settings": schema.SingleNestedAttribute{
		Attributes: Ipv6networkDiscoveryBasicPollSettingsResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Validators: []validator.Object{
			objectvalidator.AlsoRequires(path.MatchRoot("use_discovery_basic_polling_settings")),
		},
		MarkdownDescription: "The discovery basic poll settings for this network",
	},
	"discovery_blackout_setting": schema.SingleNestedAttribute{
		Attributes: Ipv6networkDiscoveryBlackoutSettingResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Validators: []validator.Object{
			objectvalidator.AlsoRequires(path.MatchRoot("use_blackout_setting")),
		},
		MarkdownDescription: "The discovery blackout setting for this network.",
	},
	"discovery_engine_type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The network discovery engine type.",
		Default:             stringdefault.StaticString("NONE"),
	},
	"discovery_member": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The member that will run discovery for this network.",
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_enable_discovery")),
		},
		Computed: true,
	},
	"domain_name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Use this method to set or retrieve the domain_name value of a DHCP IPv6 Network object.",
		Computed:            true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_domain_name")),
		},
	},
	"domain_name_servers": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "Use this method to set or retrieve the dynamic DNS updates flag of a DHCP IPv6 Network object. The DHCP server can send DDNS updates to DNS servers in the same Grid and to external DNS servers. This setting overrides the member level settings.",
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_domain_name_servers")),
			listvalidator.SizeAtLeast(1),
		},
	},
	"enable_ddns": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "The dynamic DNS updates flag of a DHCP IPv6 network object. If set to True, the DHCP server sends DDNS updates to DNS servers in the same Grid, and to external DNS servers.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_enable_ddns")),
		},
	},
	"enable_discovery": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether a discovery is enabled or not for this network. When this is set to False, the network discovery is disabled.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_enable_discovery")),
		},
	},
	"enable_ifmap_publishing": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if IFMAP publishing is enabled for the network.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_enable_ifmap_publishing")),
		},
	},
	"enable_immediate_discovery": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the discovery for the network should be immediately enabled.",
	},
	"endpoint_sources": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The endpoints that provides data for the DHCP IPv6 Network object.",
	},
	"extattrs": schema.MapAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object.",
		ElementType:         types.StringType,
		Default:             mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
	},
	"extattrs_all": schema.MapAttribute{
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object , including default attributes.",
		ElementType:         types.StringType,
	},
	"federated_realms": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6networkFederatedRealmsResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "This field contains the federated realms associated to this network",
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
			Attributes: Ipv6networkLogicFilterRulesResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "This field contains the logic filters to be applied on this IPv6 network. This list corresponds to the match rules that are written to the DHCPv6 configuration file.",
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_logic_filter_rules")),
			listvalidator.SizeAtLeast(1),
		},
	},
	"members": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6networkMembersResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "A list of members servers that serve DHCP for the network. All members in the array must be of the same type. The struct type must be indicated in each element, by setting the \"_struct\" member to the struct type.",
		Computed:            true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
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
		Attributes:          Ipv6networkMsAdUserDataResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "The Microsoft Active Directory user related information.",
	},

	"network": schema.StringAttribute{
		CustomType:          cidrtypes.IPv6PrefixType{},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The IPv6 network address in CIDR notation. The network address must be unique within the network view.",
		Validators: []validator.String{
			stringvalidator.ExactlyOneOf(
				path.MatchRoot("network"),
				path.MatchRoot("func_call"),
			),
		},
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
	},
	"func_call": schema.SingleNestedAttribute{
		Attributes:          FuncCallResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "A function call to be executed on the object.",
	},
	"network_container": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The network container to which this network belongs, if any.",
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
	},
	"network_view": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the network view in which this network resides.",
		Computed:            true,
		Default:             stringdefault.StaticString("default"),
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
	},
	"options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6networkOptionsResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object. The option `dhcp-lease-time` cannot be configured for this object and instead 'valid_lifetime' attribute should be used.",
		Computed:            true,
		Default: listdefault.StaticValue(
			types.ListValueMust(
				types.ObjectType{AttrTypes: Ipv6networkOptionsAttrTypes},
				[]attr.Value{},
			),
		),
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_options")),
			listvalidator.SizeAtLeast(1),
		},
	},
	"port_control_blackout_setting": schema.SingleNestedAttribute{
		Attributes: Ipv6networkPortControlBlackoutSettingResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Validators: []validator.Object{
			objectvalidator.AlsoRequires(path.MatchRoot("use_blackout_setting")),
		},
		MarkdownDescription: "The port control blackout setting for this network.",
	},
	"preferred_lifetime": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Use this method to set or retrieve the preferred lifetime value of a DHCP IPv6 Network object.",
		Computed:            true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_preferred_lifetime")),
		},
	},
	"recycle_leases": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If the field is set to True, the leases are kept in the Recycle Bin until one week after expiration. Otherwise, the leases are permanently deleted.",
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_recycle_leases")),
		},
	},
	"restart_if_needed": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Restarts the member service.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"rir": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The registry (RIR) that allocated the IPv6 network address space.",
	},
	"rir_organization": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The RIR organization associated with the IPv6 network.",
		Computed:            true,
	},
	"rir_registration_action": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The RIR registration action.",
		Computed:            true,
		Validators: []validator.String{
			stringvalidator.OneOf("CREATE", "MODIFY", "DELETE", "NONE"),
		},
	},
	"rir_registration_status": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The registration status of the IPv6 network in RIR.",
		Computed:            true,
		Validators: []validator.String{
			stringvalidator.OneOf("REGISTERED", "NOT_REGISTERED"),
		},
		Default: stringdefault.StaticString("NOT_REGISTERED"),
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
		Attributes: Ipv6networkSubscribeSettingsResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Validators: []validator.Object{
			objectvalidator.AlsoRequires(path.MatchRoot("use_subscribe_settings")),
		},
		MarkdownDescription: "The DHCP IPv6 Network Cisco ISE subscribe settings.",
	},
	"template": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "If set on creation, the network is created according to the values specified in the selected template.",
	},
	"unmanaged": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the DHCP IPv6 Network is unmanaged or not.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"unmanaged_count": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of unmanaged IP addresses as discovered by network discovery.",
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
	"use_domain_name": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: domain_name",
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
	"use_enable_ifmap_publishing": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: enable_ifmap_publishing",
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
	"use_recycle_leases": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: recycle_leases",
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
	"valid_lifetime": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Use this method to set or retrieve the valid lifetime value of a DHCP IPv6 Network object.",
		Computed:            true,
		Default:             int64default.StaticInt64(43200),
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_valid_lifetime")),
		},
	},
	"vlans": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6networkVlansResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "List of VLANs assigned to Network.",
		Computed:            true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
	},
	"zone_associations": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6networkZoneAssociationsResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "The list of zones associated with this network.",
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_zone_associations")),
			listvalidator.SizeAtLeast(1),
		},
		Computed: true,
	},
}

func (m *Ipv6networkModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *ipam.Ipv6network {
	if m == nil {
		return nil
	}
	to := &ipam.Ipv6network{
		CloudInfo:                        ExpandIpv6networkCloudInfo(ctx, m.CloudInfo, diags),
		Comment:                          flex.ExpandStringPointer(m.Comment),
		DdnsDomainname:                   flex.ExpandStringPointer(m.DdnsDomainname),
		DdnsEnableOptionFqdn:             flex.ExpandBoolPointer(m.DdnsEnableOptionFqdn),
		DdnsGenerateHostname:             flex.ExpandBoolPointer(m.DdnsGenerateHostname),
		DdnsServerAlwaysUpdates:          flex.ExpandBoolPointer(m.DdnsServerAlwaysUpdates),
		DdnsTtl:                          flex.ExpandInt64Pointer(m.DdnsTtl),
		DeleteReason:                     flex.ExpandStringPointer(m.DeleteReason),
		Disable:                          flex.ExpandBoolPointer(m.Disable),
		DiscoveredBridgeDomain:           flex.ExpandStringPointer(m.DiscoveredBridgeDomain),
		DiscoveredTenant:                 flex.ExpandStringPointer(m.DiscoveredTenant),
		DiscoveryBasicPollSettings:       ExpandIpv6networkDiscoveryBasicPollSettings(ctx, m.DiscoveryBasicPollSettings, diags),
		DiscoveryBlackoutSetting:         ExpandIpv6networkDiscoveryBlackoutSetting(ctx, m.DiscoveryBlackoutSetting, diags),
		DiscoveryMember:                  flex.ExpandStringPointer(m.DiscoveryMember),
		DomainName:                       flex.ExpandStringPointer(m.DomainName),
		DomainNameServers:                flex.ExpandFrameworkListString(ctx, m.DomainNameServers, diags),
		EnableDdns:                       flex.ExpandBoolPointer(m.EnableDdns),
		EnableDiscovery:                  flex.ExpandBoolPointer(m.EnableDiscovery),
		EnableIfmapPublishing:            flex.ExpandBoolPointer(m.EnableIfmapPublishing),
		EnableImmediateDiscovery:         flex.ExpandBoolPointer(m.EnableImmediateDiscovery),
		ExtAttrs:                         ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		FederatedRealms:                  flex.ExpandFrameworkListNestedBlock(ctx, m.FederatedRealms, diags, ExpandIpv6networkFederatedRealms),
		LogicFilterRules:                 flex.ExpandFrameworkListNestedBlock(ctx, m.LogicFilterRules, diags, ExpandIpv6networkLogicFilterRules),
		Members:                          flex.ExpandFrameworkListNestedBlock(ctx, m.Members, diags, ExpandIpv6networkMembers),
		MgmPrivate:                       flex.ExpandBoolPointer(m.MgmPrivate),
		MsAdUserData:                     ExpandIpv6networkMsAdUserData(ctx, m.MsAdUserData, diags),
		Network:                          ExpandIpv6NetworkNetwork(m.Network),
		FuncCall:                         ExpandFuncCall(ctx, m.FuncCall, diags),
		NetworkView:                      flex.ExpandStringPointer(m.NetworkView),
		Options:                          flex.ExpandFrameworkListNestedBlock(ctx, m.Options, diags, ExpandIpv6networkOptions),
		PortControlBlackoutSetting:       ExpandIpv6networkPortControlBlackoutSetting(ctx, m.PortControlBlackoutSetting, diags),
		PreferredLifetime:                flex.ExpandInt64Pointer(m.PreferredLifetime),
		RecycleLeases:                    flex.ExpandBoolPointer(m.RecycleLeases),
		RestartIfNeeded:                  flex.ExpandBoolPointer(m.RestartIfNeeded),
		RirOrganization:                  flex.ExpandStringPointer(m.RirOrganization),
		RirRegistrationAction:            flex.ExpandStringPointer(m.RirRegistrationAction),
		RirRegistrationStatus:            flex.ExpandStringPointer(m.RirRegistrationStatus),
		SamePortControlDiscoveryBlackout: flex.ExpandBoolPointer(m.SamePortControlDiscoveryBlackout),
		SendRirRequest:                   flex.ExpandBoolPointer(m.SendRirRequest),
		SubscribeSettings:                ExpandIpv6networkSubscribeSettings(ctx, m.SubscribeSettings, diags),
		Template:                         flex.ExpandStringPointer(m.Template),
		Unmanaged:                        flex.ExpandBoolPointer(m.Unmanaged),
		UpdateDnsOnLeaseRenewal:          flex.ExpandBoolPointer(m.UpdateDnsOnLeaseRenewal),
		UseBlackoutSetting:               flex.ExpandBoolPointer(m.UseBlackoutSetting),
		UseDdnsDomainname:                flex.ExpandBoolPointer(m.UseDdnsDomainname),
		UseDdnsEnableOptionFqdn:          flex.ExpandBoolPointer(m.UseDdnsEnableOptionFqdn),
		UseDdnsGenerateHostname:          flex.ExpandBoolPointer(m.UseDdnsGenerateHostname),
		UseDdnsTtl:                       flex.ExpandBoolPointer(m.UseDdnsTtl),
		UseDiscoveryBasicPollingSettings: flex.ExpandBoolPointer(m.UseDiscoveryBasicPollingSettings),
		UseDomainName:                    flex.ExpandBoolPointer(m.UseDomainName),
		UseDomainNameServers:             flex.ExpandBoolPointer(m.UseDomainNameServers),
		UseEnableDdns:                    flex.ExpandBoolPointer(m.UseEnableDdns),
		UseEnableDiscovery:               flex.ExpandBoolPointer(m.UseEnableDiscovery),
		UseEnableIfmapPublishing:         flex.ExpandBoolPointer(m.UseEnableIfmapPublishing),
		UseLogicFilterRules:              flex.ExpandBoolPointer(m.UseLogicFilterRules),
		UseMgmPrivate:                    flex.ExpandBoolPointer(m.UseMgmPrivate),
		UseOptions:                       flex.ExpandBoolPointer(m.UseOptions),
		UsePreferredLifetime:             flex.ExpandBoolPointer(m.UsePreferredLifetime),
		UseRecycleLeases:                 flex.ExpandBoolPointer(m.UseRecycleLeases),
		UseSubscribeSettings:             flex.ExpandBoolPointer(m.UseSubscribeSettings),
		UseUpdateDnsOnLeaseRenewal:       flex.ExpandBoolPointer(m.UseUpdateDnsOnLeaseRenewal),
		UseValidLifetime:                 flex.ExpandBoolPointer(m.UseValidLifetime),
		UseZoneAssociations:              flex.ExpandBoolPointer(m.UseZoneAssociations),
		ValidLifetime:                    flex.ExpandInt64Pointer(m.ValidLifetime),
		Vlans:                            flex.ExpandFrameworkListNestedBlock(ctx, m.Vlans, diags, ExpandIpv6networkVlans),
		ZoneAssociations:                 flex.ExpandFrameworkListNestedBlock(ctx, m.ZoneAssociations, diags, ExpandIpv6networkZoneAssociations),
	}
	if isCreate {
		to.NetworkContainer = flex.ExpandStringPointer(m.NetworkContainer)
		to.NetworkView = flex.ExpandStringPointer(m.NetworkView)
		to.Network = ExpandIpv6NetworkNetwork(m.Network)
		to.AutoCreateReversezone = flex.ExpandBoolPointer(m.AutoCreateReversezone)
	}
	return to
}

func FlattenIpv6network(ctx context.Context, from *ipam.Ipv6network, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6networkAttrTypes)
	}
	m := Ipv6networkModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, Ipv6networkAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Ipv6networkModel) Flatten(ctx context.Context, from *ipam.Ipv6network, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Ipv6networkModel{}
	}

	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.CloudInfo = FlattenIpv6networkCloudInfo(ctx, from.CloudInfo, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.DdnsDomainname = flex.FlattenStringPointer(from.DdnsDomainname)
	m.DdnsEnableOptionFqdn = types.BoolPointerValue(from.DdnsEnableOptionFqdn)
	m.DdnsGenerateHostname = types.BoolPointerValue(from.DdnsGenerateHostname)
	m.DdnsServerAlwaysUpdates = types.BoolPointerValue(from.DdnsServerAlwaysUpdates)
	m.DdnsTtl = flex.FlattenInt64Pointer(from.DdnsTtl)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.DiscoverNowStatus = flex.FlattenStringPointer(from.DiscoverNowStatus)
	m.DiscoveredBgpAs = flex.FlattenStringPointer(from.DiscoveredBgpAs)
	m.DiscoveredBridgeDomain = flex.FlattenStringPointer(from.DiscoveredBridgeDomain)
	m.DiscoveredTenant = flex.FlattenStringPointer(from.DiscoveredTenant)
	m.DiscoveredVlanId = flex.FlattenStringPointer(from.DiscoveredVlanId)
	m.DiscoveredVlanName = flex.FlattenStringPointer(from.DiscoveredVlanName)
	m.DiscoveredVrfDescription = flex.FlattenStringPointer(from.DiscoveredVrfDescription)
	m.DiscoveredVrfName = flex.FlattenStringPointer(from.DiscoveredVrfName)
	m.DiscoveredVrfRd = flex.FlattenStringPointer(from.DiscoveredVrfRd)
	m.DiscoveryBasicPollSettings = FlattenIpv6networkDiscoveryBasicPollSettings(ctx, from.DiscoveryBasicPollSettings, diags)
	m.DiscoveryBlackoutSetting = FlattenIpv6networkDiscoveryBlackoutSetting(ctx, from.DiscoveryBlackoutSetting, diags)
	m.DiscoveryEngineType = flex.FlattenStringPointer(from.DiscoveryEngineType)
	m.DiscoveryMember = flex.FlattenStringPointer(from.DiscoveryMember)
	m.DomainName = flex.FlattenStringPointer(from.DomainName)
	m.DomainNameServers = flex.FlattenFrameworkListString(ctx, from.DomainNameServers, diags)
	m.EnableDdns = types.BoolPointerValue(from.EnableDdns)
	m.EnableDiscovery = types.BoolPointerValue(from.EnableDiscovery)
	m.EnableIfmapPublishing = types.BoolPointerValue(from.EnableIfmapPublishing)
	m.EndpointSources = flex.FlattenFrameworkListString(ctx, from.EndpointSources, diags)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.FederatedRealms = flex.FlattenFrameworkListNestedBlock(ctx, from.FederatedRealms, Ipv6networkFederatedRealmsAttrTypes, diags, FlattenIpv6networkFederatedRealms)
	m.LastRirRegistrationUpdateSent = flex.FlattenInt64Pointer(from.LastRirRegistrationUpdateSent)
	m.LastRirRegistrationUpdateStatus = flex.FlattenStringPointer(from.LastRirRegistrationUpdateStatus)
	m.LogicFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.LogicFilterRules, Ipv6networkLogicFilterRulesAttrTypes, diags, FlattenIpv6networkLogicFilterRules)
	m.Members = flex.FlattenFrameworkListNestedBlock(ctx, from.Members, Ipv6networkMembersAttrTypes, diags, FlattenIpv6networkMembers)
	m.MgmPrivate = types.BoolPointerValue(from.MgmPrivate)
	m.MgmPrivateOverridable = types.BoolPointerValue(from.MgmPrivateOverridable)
	m.MsAdUserData = FlattenIpv6networkMsAdUserData(ctx, from.MsAdUserData, diags)
	m.Network = FlattenIpv6networkNetwork(from.Network)
	if m.FuncCall.IsNull() || m.FuncCall.IsUnknown() {
		m.FuncCall = FlattenFuncCall(ctx, from.FuncCall, diags)
	}
	m.NetworkContainer = flex.FlattenStringPointer(from.NetworkContainer)
	m.NetworkView = flex.FlattenStringPointer(from.NetworkView)
	planOptions := m.Options
	m.Options = flex.FlattenFrameworkListNestedBlock(ctx, from.Options, Ipv6networkOptionsAttrTypes, diags, FlattenIpv6networkOptions)
	if !planOptions.IsUnknown() {
		reOrderedOptions, diags := utils.ReorderAndFilterDHCPOptions(ctx, planOptions, m.Options)
		if !diags.HasError() {
			m.Options = reOrderedOptions.(basetypes.ListValue)
		}
	}
	m.PortControlBlackoutSetting = FlattenIpv6networkPortControlBlackoutSetting(ctx, from.PortControlBlackoutSetting, diags)
	m.PreferredLifetime = flex.FlattenInt64Pointer(from.PreferredLifetime)
	m.RecycleLeases = types.BoolPointerValue(from.RecycleLeases)
	m.Rir = flex.FlattenStringPointer(from.Rir)
	m.RirOrganization = flex.FlattenStringPointer(from.RirOrganization)
	m.RirRegistrationAction = flex.FlattenStringPointer(from.RirRegistrationAction)
	m.RirRegistrationStatus = flex.FlattenStringPointer(from.RirRegistrationStatus)
	m.SamePortControlDiscoveryBlackout = types.BoolPointerValue(from.SamePortControlDiscoveryBlackout)
	m.SubscribeSettings = FlattenIpv6networkSubscribeSettings(ctx, from.SubscribeSettings, diags)
	m.Unmanaged = types.BoolPointerValue(from.Unmanaged)
	m.UnmanagedCount = flex.FlattenInt64Pointer(from.UnmanagedCount)
	m.UpdateDnsOnLeaseRenewal = types.BoolPointerValue(from.UpdateDnsOnLeaseRenewal)
	m.UseBlackoutSetting = types.BoolPointerValue(from.UseBlackoutSetting)
	m.UseDdnsDomainname = types.BoolPointerValue(from.UseDdnsDomainname)
	m.UseDdnsEnableOptionFqdn = types.BoolPointerValue(from.UseDdnsEnableOptionFqdn)
	m.UseDdnsGenerateHostname = types.BoolPointerValue(from.UseDdnsGenerateHostname)
	m.UseDdnsTtl = types.BoolPointerValue(from.UseDdnsTtl)
	m.UseDiscoveryBasicPollingSettings = types.BoolPointerValue(from.UseDiscoveryBasicPollingSettings)
	m.UseDomainName = types.BoolPointerValue(from.UseDomainName)
	m.UseDomainNameServers = types.BoolPointerValue(from.UseDomainNameServers)
	m.UseEnableDdns = types.BoolPointerValue(from.UseEnableDdns)
	m.UseEnableDiscovery = types.BoolPointerValue(from.UseEnableDiscovery)
	m.UseEnableIfmapPublishing = types.BoolPointerValue(from.UseEnableIfmapPublishing)
	m.UseLogicFilterRules = types.BoolPointerValue(from.UseLogicFilterRules)
	m.UseMgmPrivate = types.BoolPointerValue(from.UseMgmPrivate)
	m.UseOptions = types.BoolPointerValue(from.UseOptions)
	m.UsePreferredLifetime = types.BoolPointerValue(from.UsePreferredLifetime)
	m.UseRecycleLeases = types.BoolPointerValue(from.UseRecycleLeases)
	m.UseSubscribeSettings = types.BoolPointerValue(from.UseSubscribeSettings)
	m.UseUpdateDnsOnLeaseRenewal = types.BoolPointerValue(from.UseUpdateDnsOnLeaseRenewal)
	m.UseValidLifetime = types.BoolPointerValue(from.UseValidLifetime)
	m.UseZoneAssociations = types.BoolPointerValue(from.UseZoneAssociations)
	m.ValidLifetime = flex.FlattenInt64Pointer(from.ValidLifetime)
	m.Vlans = flex.FlattenFrameworkListNestedBlock(ctx, from.Vlans, Ipv6networkVlansAttrTypes, diags, FlattenIpv6networkVlans)
	m.ZoneAssociations = flex.FlattenFrameworkListNestedBlock(ctx, from.ZoneAssociations, Ipv6networkZoneAssociationsAttrTypes, diags, FlattenIpv6networkZoneAssociations)
}

func ExpandIpv6NetworkNetwork(str cidrtypes.IPv6Prefix) *ipam.Ipv6networkNetwork {
	if str.IsNull() {
		return &ipam.Ipv6networkNetwork{}
	}
	var m ipam.Ipv6networkNetwork
	m.String = flex.ExpandIPv6CIDR(str)

	return &m
}

func FlattenIpv6networkNetwork(from *ipam.Ipv6networkNetwork) cidrtypes.IPv6Prefix {
	if from.String == nil {
		return cidrtypes.NewIPv6PrefixNull()
	}
	m := flex.FlattenIPv6CIDR(from.String)
	return m
}
