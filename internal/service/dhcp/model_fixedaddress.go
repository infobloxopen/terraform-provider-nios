package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/hwtypes"
	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type FixedaddressModel struct {
	Ref                            types.String        `tfsdk:"ref"`
	AgentCircuitId                 types.String        `tfsdk:"agent_circuit_id"`
	AgentRemoteId                  types.String        `tfsdk:"agent_remote_id"`
	AllowTelnet                    types.Bool          `tfsdk:"allow_telnet"`
	AlwaysUpdateDns                types.Bool          `tfsdk:"always_update_dns"`
	Bootfile                       types.String        `tfsdk:"bootfile"`
	Bootserver                     types.String        `tfsdk:"bootserver"`
	CliCredentials                 types.List          `tfsdk:"cli_credentials"`
	ClientIdentifierPrependZero    types.Bool          `tfsdk:"client_identifier_prepend_zero"`
	CloudInfo                      types.Object        `tfsdk:"cloud_info"`
	Comment                        types.String        `tfsdk:"comment"`
	DdnsDomainname                 types.String        `tfsdk:"ddns_domainname"`
	DdnsHostname                   types.String        `tfsdk:"ddns_hostname"`
	DenyBootp                      types.Bool          `tfsdk:"deny_bootp"`
	DeviceDescription              types.String        `tfsdk:"device_description"`
	DeviceLocation                 types.String        `tfsdk:"device_location"`
	DeviceType                     types.String        `tfsdk:"device_type"`
	DeviceVendor                   types.String        `tfsdk:"device_vendor"`
	DhcpClientIdentifier           types.String        `tfsdk:"dhcp_client_identifier"`
	Disable                        types.Bool          `tfsdk:"disable"`
	DisableDiscovery               types.Bool          `tfsdk:"disable_discovery"`
	DiscoverNowStatus              types.String        `tfsdk:"discover_now_status"`
	DiscoveredData                 types.Object        `tfsdk:"discovered_data"`
	EnableDdns                     types.Bool          `tfsdk:"enable_ddns"`
	EnableImmediateDiscovery       types.Bool          `tfsdk:"enable_immediate_discovery"`
	EnablePxeLeaseTime             types.Bool          `tfsdk:"enable_pxe_lease_time"`
	ExtAttrs                       types.Map           `tfsdk:"extattrs"`
	ExtAttrsAll                    types.Map           `tfsdk:"extattrs_all"`
	IgnoreDhcpOptionListRequest    types.Bool          `tfsdk:"ignore_dhcp_option_list_request"`
	Ipv4addr                       iptypes.IPv4Address `tfsdk:"ipv4addr"`
	FuncCall                       types.Object        `tfsdk:"func_call"`
	IsInvalidMac                   types.Bool          `tfsdk:"is_invalid_mac"`
	LogicFilterRules               types.List          `tfsdk:"logic_filter_rules"`
	Mac                            hwtypes.MACAddress  `tfsdk:"mac"`
	MatchClient                    types.String        `tfsdk:"match_client"`
	MsAdUserData                   types.Object        `tfsdk:"ms_ad_user_data"`
	MsOptions                      types.List          `tfsdk:"ms_options"`
	MsServer                       types.Object        `tfsdk:"ms_server"`
	Name                           types.String        `tfsdk:"name"`
	Network                        types.String        `tfsdk:"network"`
	NetworkView                    types.String        `tfsdk:"network_view"`
	Nextserver                     types.String        `tfsdk:"nextserver"`
	Options                        types.List          `tfsdk:"options"`
	PxeLeaseTime                   types.Int64         `tfsdk:"pxe_lease_time"`
	ReservedInterface              types.String        `tfsdk:"reserved_interface"`
	RestartIfNeeded                types.Bool          `tfsdk:"restart_if_needed"`
	Snmp3Credential                types.Object        `tfsdk:"snmp3_credential"`
	SnmpCredential                 types.Object        `tfsdk:"snmp_credential"`
	Template                       types.String        `tfsdk:"template"`
	UseBootfile                    types.Bool          `tfsdk:"use_bootfile"`
	UseBootserver                  types.Bool          `tfsdk:"use_bootserver"`
	UseCliCredentials              types.Bool          `tfsdk:"use_cli_credentials"`
	UseDdnsDomainname              types.Bool          `tfsdk:"use_ddns_domainname"`
	UseDenyBootp                   types.Bool          `tfsdk:"use_deny_bootp"`
	UseEnableDdns                  types.Bool          `tfsdk:"use_enable_ddns"`
	UseIgnoreDhcpOptionListRequest types.Bool          `tfsdk:"use_ignore_dhcp_option_list_request"`
	UseLogicFilterRules            types.Bool          `tfsdk:"use_logic_filter_rules"`
	UseMsOptions                   types.Bool          `tfsdk:"use_ms_options"`
	UseNextserver                  types.Bool          `tfsdk:"use_nextserver"`
	UseOptions                     types.Bool          `tfsdk:"use_options"`
	UsePxeLeaseTime                types.Bool          `tfsdk:"use_pxe_lease_time"`
	UseSnmp3Credential             types.Bool          `tfsdk:"use_snmp3_credential"`
	UseSnmpCredential              types.Bool          `tfsdk:"use_snmp_credential"`
}

var FixedaddressAttrTypes = map[string]attr.Type{
	"ref":                                 types.StringType,
	"agent_circuit_id":                    types.StringType,
	"agent_remote_id":                     types.StringType,
	"allow_telnet":                        types.BoolType,
	"always_update_dns":                   types.BoolType,
	"bootfile":                            types.StringType,
	"bootserver":                          types.StringType,
	"cli_credentials":                     types.ListType{ElemType: types.ObjectType{AttrTypes: FixedaddressCliCredentialsAttrTypes}},
	"client_identifier_prepend_zero":      types.BoolType,
	"cloud_info":                          types.ObjectType{AttrTypes: FixedaddressCloudInfoAttrTypes},
	"comment":                             types.StringType,
	"ddns_domainname":                     types.StringType,
	"ddns_hostname":                       types.StringType,
	"deny_bootp":                          types.BoolType,
	"device_description":                  types.StringType,
	"device_location":                     types.StringType,
	"device_type":                         types.StringType,
	"device_vendor":                       types.StringType,
	"dhcp_client_identifier":              types.StringType,
	"disable":                             types.BoolType,
	"disable_discovery":                   types.BoolType,
	"discover_now_status":                 types.StringType,
	"discovered_data":                     types.ObjectType{AttrTypes: FixedaddressDiscoveredDataAttrTypes},
	"enable_ddns":                         types.BoolType,
	"enable_immediate_discovery":          types.BoolType,
	"enable_pxe_lease_time":               types.BoolType,
	"extattrs":                            types.MapType{ElemType: types.StringType},
	"extattrs_all":                        types.MapType{ElemType: types.StringType},
	"ignore_dhcp_option_list_request":     types.BoolType,
	"ipv4addr":                            iptypes.IPv4AddressType{},
	"func_call":                           types.ObjectType{AttrTypes: FuncCallAttrTypes},
	"is_invalid_mac":                      types.BoolType,
	"logic_filter_rules":                  types.ListType{ElemType: types.ObjectType{AttrTypes: FixedaddressLogicFilterRulesAttrTypes}},
	"mac":                                 hwtypes.MACAddressType{},
	"match_client":                        types.StringType,
	"ms_ad_user_data":                     types.ObjectType{AttrTypes: FixedaddressMsAdUserDataAttrTypes},
	"ms_options":                          types.ListType{ElemType: types.ObjectType{AttrTypes: FixedaddressMsOptionsAttrTypes}},
	"ms_server":                           types.ObjectType{AttrTypes: FixedaddressMsServerAttrTypes},
	"name":                                types.StringType,
	"network":                             types.StringType,
	"network_view":                        types.StringType,
	"nextserver":                          types.StringType,
	"options":                             types.ListType{ElemType: types.ObjectType{AttrTypes: FixedaddressOptionsAttrTypes}},
	"pxe_lease_time":                      types.Int64Type,
	"reserved_interface":                  types.StringType,
	"restart_if_needed":                   types.BoolType,
	"snmp3_credential":                    types.ObjectType{AttrTypes: FixedaddressSnmp3CredentialAttrTypes},
	"snmp_credential":                     types.ObjectType{AttrTypes: FixedaddressSnmpCredentialAttrTypes},
	"template":                            types.StringType,
	"use_bootfile":                        types.BoolType,
	"use_bootserver":                      types.BoolType,
	"use_cli_credentials":                 types.BoolType,
	"use_ddns_domainname":                 types.BoolType,
	"use_deny_bootp":                      types.BoolType,
	"use_enable_ddns":                     types.BoolType,
	"use_ignore_dhcp_option_list_request": types.BoolType,
	"use_logic_filter_rules":              types.BoolType,
	"use_ms_options":                      types.BoolType,
	"use_nextserver":                      types.BoolType,
	"use_options":                         types.BoolType,
	"use_pxe_lease_time":                  types.BoolType,
	"use_snmp3_credential":                types.BoolType,
	"use_snmp_credential":                 types.BoolType,
}

var FixedaddressResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"agent_circuit_id": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The agent circuit ID for the fixed address.",
	},
	"agent_remote_id": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The agent remote ID for the fixed address.",
	},
	"allow_telnet": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This field controls whether the credential is used for both the Telnet and SSH credentials. If set to False, the credential is used only for SSH.",
	},
	"always_update_dns": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This field controls whether only the DHCP server is allowed to update DNS, regardless of the DHCP client requests.",
	},
	"bootfile": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_bootfile")),
		},
		MarkdownDescription: "The bootfile name for the fixed address. You can configure the DHCP server to support clients that use the boot file name option in their DHCPREQUEST messages.",
	},
	"bootserver": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_bootserver")),
			customvalidator.IsValidFQDN(),
		},
		MarkdownDescription: "The bootserver address for the fixed address. You can specify the name and/or IP address of the boot server that the host needs to boot. The boot server IPv4 Address or name in FQDN format.",
	},
	"cli_credentials": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: FixedaddressCliCredentialsResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The CLI credentials for the fixed address.",
	},
	"client_identifier_prepend_zero": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This field controls whether there is a prepend for the dhcp-client-identifier of a fixed address.",
	},
	"cloud_info": schema.SingleNestedAttribute{
		Attributes:          FixedaddressCloudInfoResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "Structure containing all cloud API related information for this object.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Comment for the fixed address; maximum 256 characters.",
	},
	"ddns_domainname": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_ddns_domainname")),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The dynamic DNS domain name the appliance uses specifically for DDNS updates for this fixed address.",
	},
	"ddns_hostname": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The DDNS host name for this fixed address.",
	},
	"deny_bootp": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_deny_bootp")),
		},
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If set to true, BOOTP settings are disabled and BOOTP requests will be denied.",
	},
	"device_description": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The description of the device.",
	},
	"device_location": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The location of the device.",
	},
	"device_type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The type of the device.",
	},
	"device_vendor": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The vendor of the device.",
	},
	"dhcp_client_identifier": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The DHCP client ID for the fixed address.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether a fixed address is disabled or not. When this is set to False, the fixed address is enabled.",
	},
	"disable_discovery": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the discovery for this fixed address is disabled or not. False means that the discovery is enabled.",
	},
	"discover_now_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The discovery status of this fixed address.",
	},
	"discovered_data": schema.SingleNestedAttribute{
		Attributes:          FixedaddressDiscoveredDataResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "The discovered data for this fixed address.",
	},
	"enable_ddns": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_enable_ddns")),
		},
		MarkdownDescription: "The dynamic DNS updates flag of a DHCP Fixed Address object. If set to True, the DHCP server sends DDNS updates to DNS servers in the same Grid, and to external DNS servers.",
	},
	"enable_immediate_discovery": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the discovery for the fixed address should be immediately enabled.",
	},
	"enable_pxe_lease_time": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Set this to True if you want the DHCP server to use a different lease time for PXE clients.",
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
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.",
	},
	"ignore_dhcp_option_list_request": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_ignore_dhcp_option_list_request")),
		},
		MarkdownDescription: "If this field is set to False, the appliance returns all DHCP options the client is eligible to receive, rather than only the list of options the client has requested.",
	},
	"ipv4addr": schema.StringAttribute{
		CustomType:          iptypes.IPv4AddressType{},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The IPv4 Address of the record.",
	},
	"func_call": schema.SingleNestedAttribute{
		Attributes: FuncCallResourceSchemaAttributes,
		Optional:   true,
		Validators: []validator.Object{
			objectvalidator.ConflictsWith(path.MatchRoot("ipv4addr")),
		},
		MarkdownDescription: "Function call to be executed for Fixed Address",
	},
	"is_invalid_mac": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "This flag reflects whether the MAC address for this fixed address is invalid.",
	},
	"logic_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: FixedaddressLogicFilterRulesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_logic_filter_rules")),
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "This field contains the logic filters to be applied on the this fixed address. This list corresponds to the match rules that are written to the dhcpd configuration file.",
	},
	"mac": schema.StringAttribute{
		CustomType: hwtypes.MACAddressType{},
		Optional:   true,
		Computed:   true,
		Validators: []validator.String{
			stringvalidator.ExactlyOneOf(
				path.MatchRoot("mac"),
				path.MatchRoot("agent_circuit_id"),
				path.MatchRoot("agent_remote_id"),
				path.MatchRoot("dhcp_client_identifier")),
		},
		MarkdownDescription: "The MAC address value for this fixed address.",
	},
	"match_client": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("CIRCUIT_ID", "CLIENT_ID", "MAC_ADDRESS", "REMOTE_ID", "RESERVED"),
		},
		Default:             stringdefault.StaticString("MAC_ADDRESS"),
		MarkdownDescription: "The match_client value for this fixed address. Valid values are: \"MAC_ADDRESS\": The fixed IP address is leased to the matching MAC address. \"CLIENT_ID\": The fixed IP address is leased to the matching DHCP client identifier. \"RESERVED\": The fixed IP address is reserved for later use with a MAC address that only has zeros. \"CIRCUIT_ID\": The fixed IP address is leased to the DHCP client with a matching circuit ID. Note that the \"agent_circuit_id\" field must be set in this case. \"REMOTE_ID\": The fixed IP address is leased to the DHCP client with a matching remote ID. Note that the \"agent_remote_id\" field must be set in this case.",
	},
	"ms_ad_user_data": schema.SingleNestedAttribute{
		Attributes:          FixedaddressMsAdUserDataResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "The Microsoft Active Directory user related information.",
	},
	"ms_options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: FixedaddressMsOptionsResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_ms_options")),
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "This field contains the Microsoft DHCP options for this fixed address.",
	},
	"ms_server": schema.SingleNestedAttribute{
		Attributes:          FixedaddressMsServerResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "The Microsoft server associated with this fixed address.",
	},
	"name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "This field contains the name of this fixed address.",
	},
	"network": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The network to which this fixed address belongs, in IPv4 Address/CIDR format.",
	},
	"network_view": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString("default"),
		MarkdownDescription: "The name of the network view in which this fixed address resides.",
	},
	"nextserver": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_nextserver")),
		},
		MarkdownDescription: "The name in FQDN and/or IPv4 Address format of the next server that the host needs to boot.",
	},
	"options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: FixedaddressOptionsResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_options")),
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object.",
	},
	"pxe_lease_time": schema.Int64Attribute{
		Optional: true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_pxe_lease_time")),
			int64validator.Between(0, 4294967295),
		},
		MarkdownDescription: "The PXE lease time value for a DHCP Fixed Address object. Some hosts use PXE (Preboot Execution Environment) to boot remotely from a server. To better manage your IP resources, set a different lease time for PXE boot requests. You can configure the DHCP server to allocate an IP address with a shorter lease time to hosts that send PXE boot requests, so IP addresses are not leased longer than necessary. A 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached.",
	},
	"reserved_interface": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The ref to the reserved interface to which the device belongs.",
	},
	"restart_if_needed": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Restarts the member service. The restart_if_needed flag can trigger a restart on DHCP services only when it is enabled on CP member.",
	},
	"snmp3_credential": schema.SingleNestedAttribute{
		Attributes:          FixedaddressSnmp3CredentialResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "The SNMPv3 credential for this fixed address.",
	},
	"snmp_credential": schema.SingleNestedAttribute{
		Attributes:          FixedaddressSnmpCredentialResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "The SNMP credential for this fixed address. If set to true, the SNMP credential will override member-level settings.",
	},
	"template": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If set on creation, the fixed address will be created according to the values specified in the named template.",
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
	"use_cli_credentials": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If set to true, the CLI credential will override member-level settings.",
	},
	"use_ddns_domainname": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: ddns_domainname",
	},
	"use_deny_bootp": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: deny_bootp",
	},
	"use_enable_ddns": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: enable_ddns",
	},
	"use_ignore_dhcp_option_list_request": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: ignore_dhcp_option_list_request",
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
	"use_snmp3_credential": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the SNMPv3 credential should be used for the fixed address.",
	},
	"use_snmp_credential": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If set to true, the SNMP credential will override member-level settings.",
	},
}

func (m *FixedaddressModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *dhcp.Fixedaddress {
	if m == nil {
		return nil
	}
	to := &dhcp.Fixedaddress{
		AgentCircuitId:                 flex.ExpandStringPointer(m.AgentCircuitId),
		AgentRemoteId:                  flex.ExpandStringPointer(m.AgentRemoteId),
		AllowTelnet:                    flex.ExpandBoolPointer(m.AllowTelnet),
		AlwaysUpdateDns:                flex.ExpandBoolPointer(m.AlwaysUpdateDns),
		Bootfile:                       flex.ExpandStringPointer(m.Bootfile),
		Bootserver:                     flex.ExpandStringPointer(m.Bootserver),
		CliCredentials:                 flex.ExpandFrameworkListNestedBlock(ctx, m.CliCredentials, diags, ExpandFixedaddressCliCredentials),
		ClientIdentifierPrependZero:    flex.ExpandBoolPointer(m.ClientIdentifierPrependZero),
		Comment:                        flex.ExpandStringPointer(m.Comment),
		DdnsDomainname:                 flex.ExpandStringPointer(m.DdnsDomainname),
		DdnsHostname:                   flex.ExpandStringPointer(m.DdnsHostname),
		DenyBootp:                      flex.ExpandBoolPointer(m.DenyBootp),
		DeviceDescription:              flex.ExpandStringPointer(m.DeviceDescription),
		DeviceLocation:                 flex.ExpandStringPointer(m.DeviceLocation),
		DeviceType:                     flex.ExpandStringPointer(m.DeviceType),
		DeviceVendor:                   flex.ExpandStringPointer(m.DeviceVendor),
		DhcpClientIdentifier:           flex.ExpandStringPointer(m.DhcpClientIdentifier),
		Disable:                        flex.ExpandBoolPointer(m.Disable),
		DisableDiscovery:               flex.ExpandBoolPointer(m.DisableDiscovery),
		EnableDdns:                     flex.ExpandBoolPointer(m.EnableDdns),
		EnableImmediateDiscovery:       flex.ExpandBoolPointer(m.EnableImmediateDiscovery),
		EnablePxeLeaseTime:             flex.ExpandBoolPointer(m.EnablePxeLeaseTime),
		ExtAttrs:                       ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		IgnoreDhcpOptionListRequest:    flex.ExpandBoolPointer(m.IgnoreDhcpOptionListRequest),
		Ipv4addr:                       ExpandFixedAddressIpv4addr(m.Ipv4addr),
		FuncCall:                       ExpandFuncCall(ctx, m.FuncCall, diags),
		LogicFilterRules:               flex.ExpandFrameworkListNestedBlock(ctx, m.LogicFilterRules, diags, ExpandFixedaddressLogicFilterRules),
		Mac:                            flex.ExpandMACAddress(m.Mac),
		MatchClient:                    flex.ExpandStringPointer(m.MatchClient),
		MsAdUserData:                   ExpandFixedaddressMsAdUserData(ctx, m.MsAdUserData, diags),
		MsOptions:                      flex.ExpandFrameworkListNestedBlock(ctx, m.MsOptions, diags, ExpandFixedaddressMsOptions),
		MsServer:                       ExpandFixedaddressMsServer(ctx, m.MsServer, diags),
		Name:                           flex.ExpandStringPointer(m.Name),
		Network:                        flex.ExpandStringPointer(m.Network),
		NetworkView:                    flex.ExpandStringPointer(m.NetworkView),
		Nextserver:                     flex.ExpandStringPointer(m.Nextserver),
		Options:                        flex.ExpandFrameworkListNestedBlock(ctx, m.Options, diags, ExpandFixedaddressOptions),
		PxeLeaseTime:                   flex.ExpandInt64Pointer(m.PxeLeaseTime),
		ReservedInterface:              flex.ExpandStringPointer(m.ReservedInterface),
		RestartIfNeeded:                flex.ExpandBoolPointer(m.RestartIfNeeded),
		Snmp3Credential:                ExpandFixedaddressSnmp3Credential(ctx, m.Snmp3Credential, diags),
		SnmpCredential:                 ExpandFixedaddressSnmpCredential(ctx, m.SnmpCredential, diags),
		Template:                       flex.ExpandStringPointer(m.Template),
		UseBootfile:                    flex.ExpandBoolPointer(m.UseBootfile),
		UseBootserver:                  flex.ExpandBoolPointer(m.UseBootserver),
		UseCliCredentials:              flex.ExpandBoolPointer(m.UseCliCredentials),
		UseDdnsDomainname:              flex.ExpandBoolPointer(m.UseDdnsDomainname),
		UseDenyBootp:                   flex.ExpandBoolPointer(m.UseDenyBootp),
		UseEnableDdns:                  flex.ExpandBoolPointer(m.UseEnableDdns),
		UseIgnoreDhcpOptionListRequest: flex.ExpandBoolPointer(m.UseIgnoreDhcpOptionListRequest),
		UseLogicFilterRules:            flex.ExpandBoolPointer(m.UseLogicFilterRules),
		UseMsOptions:                   flex.ExpandBoolPointer(m.UseMsOptions),
		UseNextserver:                  flex.ExpandBoolPointer(m.UseNextserver),
		UseOptions:                     flex.ExpandBoolPointer(m.UseOptions),
		UsePxeLeaseTime:                flex.ExpandBoolPointer(m.UsePxeLeaseTime),
		UseSnmp3Credential:             flex.ExpandBoolPointer(m.UseSnmp3Credential),
		UseSnmpCredential:              flex.ExpandBoolPointer(m.UseSnmpCredential),
	}
	if isCreate {
		to.Template = flex.ExpandStringPointer(m.Template)
	}
	return to
}

func FlattenFixedaddress(ctx context.Context, from *dhcp.Fixedaddress, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(FixedaddressAttrTypes)
	}
	m := FixedaddressModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, FixedaddressAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *FixedaddressModel) Flatten(ctx context.Context, from *dhcp.Fixedaddress, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = FixedaddressModel{}
	}

	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.AgentCircuitId = flex.FlattenStringPointer(from.AgentCircuitId)
	m.AgentRemoteId = flex.FlattenStringPointer(from.AgentRemoteId)
	m.AllowTelnet = types.BoolPointerValue(from.AllowTelnet)
	m.AlwaysUpdateDns = types.BoolPointerValue(from.AlwaysUpdateDns)
	m.Bootfile = flex.FlattenStringPointer(from.Bootfile)
	m.Bootserver = flex.FlattenStringPointer(from.Bootserver)
	m.CliCredentials = flex.FlattenFrameworkListNestedBlock(ctx, from.CliCredentials, FixedaddressCliCredentialsAttrTypes, diags, FlattenFixedaddressCliCredentials)
	m.ClientIdentifierPrependZero = types.BoolPointerValue(from.ClientIdentifierPrependZero)
	m.CloudInfo = FlattenFixedaddressCloudInfo(ctx, from.CloudInfo, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.DdnsDomainname = flex.FlattenStringPointer(from.DdnsDomainname)
	m.DdnsHostname = flex.FlattenStringPointer(from.DdnsHostname)
	m.DenyBootp = types.BoolPointerValue(from.DenyBootp)
	m.DeviceDescription = flex.FlattenStringPointer(from.DeviceDescription)
	m.DeviceLocation = flex.FlattenStringPointer(from.DeviceLocation)
	m.DeviceType = flex.FlattenStringPointer(from.DeviceType)
	m.DeviceVendor = flex.FlattenStringPointer(from.DeviceVendor)
	m.DhcpClientIdentifier = flex.FlattenStringPointer(from.DhcpClientIdentifier)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.DisableDiscovery = types.BoolPointerValue(from.DisableDiscovery)
	m.DiscoverNowStatus = flex.FlattenStringPointer(from.DiscoverNowStatus)
	m.DiscoveredData = FlattenFixedaddressDiscoveredData(ctx, from.DiscoveredData, diags)
	m.EnableDdns = types.BoolPointerValue(from.EnableDdns)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.EnablePxeLeaseTime = types.BoolPointerValue(from.EnablePxeLeaseTime)
	m.IgnoreDhcpOptionListRequest = types.BoolPointerValue(from.IgnoreDhcpOptionListRequest)
	m.Ipv4addr = FlattenFixedAddressIpv4addr(from.Ipv4addr)
	m.IsInvalidMac = types.BoolPointerValue(from.IsInvalidMac)
	m.LogicFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.LogicFilterRules, FixedaddressLogicFilterRulesAttrTypes, diags, FlattenFixedaddressLogicFilterRules)
	m.Mac = flex.FlattenMACAddress(from.Mac)
	m.MatchClient = flex.FlattenStringPointer(from.MatchClient)
	m.MsAdUserData = FlattenFixedaddressMsAdUserData(ctx, from.MsAdUserData, diags)
	m.MsOptions = flex.FlattenFrameworkListNestedBlock(ctx, from.MsOptions, FixedaddressMsOptionsAttrTypes, diags, FlattenFixedaddressMsOptions)
	m.MsServer = FlattenFixedaddressMsServer(ctx, from.MsServer, diags)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Network = flex.FlattenStringPointer(from.Network)
	m.NetworkView = flex.FlattenStringPointer(from.NetworkView)
	m.Nextserver = flex.FlattenStringPointer(from.Nextserver)
	planOptions := m.Options
	m.Options = flex.FlattenFrameworkListNestedBlock(ctx, from.Options, FixedaddressOptionsAttrTypes, diags, FlattenFixedaddressOptions)
	if !planOptions.IsUnknown() {
		reOrderedOptions, diags := utils.ReorderAndFilterDHCPOptions(ctx, planOptions, m.Options)
		if !diags.HasError() {
			m.Options = reOrderedOptions.(basetypes.ListValue)
		}
	}
	m.PxeLeaseTime = flex.FlattenInt64Pointer(from.PxeLeaseTime)
	m.ReservedInterface = flex.FlattenStringPointer(from.ReservedInterface)
	m.Snmp3Credential = FlattenFixedaddressSnmp3Credential(ctx, from.Snmp3Credential, diags)
	m.SnmpCredential = FlattenFixedaddressSnmpCredential(ctx, from.SnmpCredential, diags)
	m.Template = flex.FlattenStringPointer(from.Template)
	m.UseBootfile = types.BoolPointerValue(from.UseBootfile)
	m.UseBootserver = types.BoolPointerValue(from.UseBootserver)
	m.UseCliCredentials = types.BoolPointerValue(from.UseCliCredentials)
	m.UseDdnsDomainname = types.BoolPointerValue(from.UseDdnsDomainname)
	m.UseDenyBootp = types.BoolPointerValue(from.UseDenyBootp)
	m.UseEnableDdns = types.BoolPointerValue(from.UseEnableDdns)
	m.UseIgnoreDhcpOptionListRequest = types.BoolPointerValue(from.UseIgnoreDhcpOptionListRequest)
	m.UseLogicFilterRules = types.BoolPointerValue(from.UseLogicFilterRules)
	m.UseMsOptions = types.BoolPointerValue(from.UseMsOptions)
	m.UseNextserver = types.BoolPointerValue(from.UseNextserver)
	m.UseOptions = types.BoolPointerValue(from.UseOptions)
	m.UsePxeLeaseTime = types.BoolPointerValue(from.UsePxeLeaseTime)
	m.UseSnmp3Credential = types.BoolPointerValue(from.UseSnmp3Credential)
	m.UseSnmpCredential = types.BoolPointerValue(from.UseSnmpCredential)

	if m.FuncCall.IsNull() || m.FuncCall.IsUnknown() {
		m.FuncCall = FlattenFuncCall(ctx, from.FuncCall, diags)
	}
}

func ExpandFixedAddressIpv4addr(ipv4addr iptypes.IPv4Address) *dhcp.FixedaddressIpv4addr {
	if ipv4addr.IsNull() {
		return &dhcp.FixedaddressIpv4addr{}
	}
	var m dhcp.FixedaddressIpv4addr
	m.String = flex.ExpandIPv4Address(ipv4addr)
	return &m
}

func FlattenFixedAddressIpv4addr(from *dhcp.FixedaddressIpv4addr) iptypes.IPv4Address {
	if from.String == nil {
		return iptypes.NewIPv4AddressNull()
	}
	m := flex.FlattenIPv4Address(from.String)
	return m
}
