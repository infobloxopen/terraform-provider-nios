package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/boolvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
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

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	planmodifiers "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/immutable"
	importmod "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/import"
	internaltypes "github.com/infobloxopen/terraform-provider-nios/internal/types"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type RoaminghostModel struct {
	Ref                            types.String                             `tfsdk:"ref"`
	Uuid                           types.String                             `tfsdk:"uuid"`
	AddressType                    types.String                             `tfsdk:"address_type"`
	Bootfile                       types.String                             `tfsdk:"bootfile"`
	Bootserver                     types.String                             `tfsdk:"bootserver"`
	ClientIdentifierPrependZero    types.Bool                               `tfsdk:"client_identifier_prepend_zero"`
	Comment                        types.String                             `tfsdk:"comment"`
	DdnsDomainname                 internaltypes.CaseInsensitiveStringValue `tfsdk:"ddns_domainname"`
	DdnsHostname                   types.String                             `tfsdk:"ddns_hostname"`
	DenyBootp                      types.Bool                               `tfsdk:"deny_bootp"`
	DhcpClientIdentifier           types.String                             `tfsdk:"dhcp_client_identifier"`
	Disable                        types.Bool                               `tfsdk:"disable"`
	EnableDdns                     types.Bool                               `tfsdk:"enable_ddns"`
	EnablePxeLeaseTime             types.Bool                               `tfsdk:"enable_pxe_lease_time"`
	ExtAttrs                       types.Map                                `tfsdk:"extattrs"`
	ExtAttrsAll                    types.Map                                `tfsdk:"extattrs_all"`
	ForceRoamingHostname           types.Bool                               `tfsdk:"force_roaming_hostname"`
	IgnoreDhcpOptionListRequest    types.Bool                               `tfsdk:"ignore_dhcp_option_list_request"`
	Ipv6ClientHostname             types.String                             `tfsdk:"ipv6_client_hostname"`
	Ipv6DdnsDomainname             types.String                             `tfsdk:"ipv6_ddns_domainname"`
	Ipv6DdnsHostname               types.String                             `tfsdk:"ipv6_ddns_hostname"`
	Ipv6DomainName                 types.String                             `tfsdk:"ipv6_domain_name"`
	Ipv6DomainNameServers          types.List                               `tfsdk:"ipv6_domain_name_servers"`
	Ipv6Duid                       internaltypes.DUIDValue                  `tfsdk:"ipv6_duid"`
	Ipv6EnableDdns                 types.Bool                               `tfsdk:"ipv6_enable_ddns"`
	Ipv6ForceRoamingHostname       types.Bool                               `tfsdk:"ipv6_force_roaming_hostname"`
	Ipv6MacAddress                 internaltypes.MACAddressValue            `tfsdk:"ipv6_mac_address"`
	Ipv6MatchOption                types.String                             `tfsdk:"ipv6_match_option"`
	Ipv6Options                    types.List                               `tfsdk:"ipv6_options"`
	Ipv6Template                   types.String                             `tfsdk:"ipv6_template"`
	Mac                            internaltypes.MACAddressValue            `tfsdk:"mac"`
	MatchClient                    types.String                             `tfsdk:"match_client"`
	Name                           types.String                             `tfsdk:"name"`
	NetworkView                    types.String                             `tfsdk:"network_view"`
	Nextserver                     types.String                             `tfsdk:"nextserver"`
	Options                        types.List                               `tfsdk:"options"`
	PreferredLifetime              types.Int64                              `tfsdk:"preferred_lifetime"`
	PxeLeaseTime                   types.Int64                              `tfsdk:"pxe_lease_time"`
	Template                       types.String                             `tfsdk:"template"`
	UseBootfile                    types.Bool                               `tfsdk:"use_bootfile"`
	UseBootserver                  types.Bool                               `tfsdk:"use_bootserver"`
	UseDdnsDomainname              types.Bool                               `tfsdk:"use_ddns_domainname"`
	UseDenyBootp                   types.Bool                               `tfsdk:"use_deny_bootp"`
	UseEnableDdns                  types.Bool                               `tfsdk:"use_enable_ddns"`
	UseIgnoreDhcpOptionListRequest types.Bool                               `tfsdk:"use_ignore_dhcp_option_list_request"`
	UseIpv6DdnsDomainname          types.Bool                               `tfsdk:"use_ipv6_ddns_domainname"`
	UseIpv6DomainName              types.Bool                               `tfsdk:"use_ipv6_domain_name"`
	UseIpv6DomainNameServers       types.Bool                               `tfsdk:"use_ipv6_domain_name_servers"`
	UseIpv6EnableDdns              types.Bool                               `tfsdk:"use_ipv6_enable_ddns"`
	UseIpv6Options                 types.Bool                               `tfsdk:"use_ipv6_options"`
	UseNextserver                  types.Bool                               `tfsdk:"use_nextserver"`
	UseOptions                     types.Bool                               `tfsdk:"use_options"`
	UsePreferredLifetime           types.Bool                               `tfsdk:"use_preferred_lifetime"`
	UsePxeLeaseTime                types.Bool                               `tfsdk:"use_pxe_lease_time"`
	UseValidLifetime               types.Bool                               `tfsdk:"use_valid_lifetime"`
	ValidLifetime                  types.Int64                              `tfsdk:"valid_lifetime"`
}

var RoaminghostAttrTypes = map[string]attr.Type{
	"ref":                                 types.StringType,
	"uuid":                                types.StringType,
	"address_type":                        types.StringType,
	"bootfile":                            types.StringType,
	"bootserver":                          types.StringType,
	"client_identifier_prepend_zero":      types.BoolType,
	"comment":                             types.StringType,
	"ddns_domainname":                     internaltypes.CaseInsensitiveString{},
	"ddns_hostname":                       types.StringType,
	"deny_bootp":                          types.BoolType,
	"dhcp_client_identifier":              types.StringType,
	"disable":                             types.BoolType,
	"enable_ddns":                         types.BoolType,
	"enable_pxe_lease_time":               types.BoolType,
	"extattrs":                            types.MapType{ElemType: types.StringType},
	"extattrs_all":                        types.MapType{ElemType: types.StringType},
	"force_roaming_hostname":              types.BoolType,
	"ignore_dhcp_option_list_request":     types.BoolType,
	"ipv6_client_hostname":                types.StringType,
	"ipv6_ddns_domainname":                types.StringType,
	"ipv6_ddns_hostname":                  types.StringType,
	"ipv6_domain_name":                    types.StringType,
	"ipv6_domain_name_servers":            types.ListType{ElemType: types.StringType},
	"ipv6_duid":                           internaltypes.DUIDType{},
	"ipv6_enable_ddns":                    types.BoolType,
	"ipv6_force_roaming_hostname":         types.BoolType,
	"ipv6_mac_address":                    internaltypes.MACAddressType{},
	"ipv6_match_option":                   types.StringType,
	"ipv6_options":                        types.ListType{ElemType: types.ObjectType{AttrTypes: RoaminghostIpv6OptionsAttrTypes}},
	"ipv6_template":                       types.StringType,
	"mac":                                 internaltypes.MACAddressType{},
	"match_client":                        types.StringType,
	"name":                                types.StringType,
	"network_view":                        types.StringType,
	"nextserver":                          types.StringType,
	"options":                             types.ListType{ElemType: types.ObjectType{AttrTypes: RoaminghostOptionsAttrTypes}},
	"preferred_lifetime":                  types.Int64Type,
	"pxe_lease_time":                      types.Int64Type,
	"template":                            types.StringType,
	"use_bootfile":                        types.BoolType,
	"use_bootserver":                      types.BoolType,
	"use_ddns_domainname":                 types.BoolType,
	"use_deny_bootp":                      types.BoolType,
	"use_enable_ddns":                     types.BoolType,
	"use_ignore_dhcp_option_list_request": types.BoolType,
	"use_ipv6_ddns_domainname":            types.BoolType,
	"use_ipv6_domain_name":                types.BoolType,
	"use_ipv6_domain_name_servers":        types.BoolType,
	"use_ipv6_enable_ddns":                types.BoolType,
	"use_ipv6_options":                    types.BoolType,
	"use_nextserver":                      types.BoolType,
	"use_options":                         types.BoolType,
	"use_preferred_lifetime":              types.BoolType,
	"use_pxe_lease_time":                  types.BoolType,
	"use_valid_lifetime":                  types.BoolType,
	"valid_lifetime":                      types.Int64Type,
}

var RoaminghostResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The uuid to the object.",
	},
	"address_type": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Default:  stringdefault.StaticString("IPV4"),
		Validators: []validator.String{
			stringvalidator.OneOf("BOTH", "IPV4", "IPV6"),
		},
		MarkdownDescription: "The address type for this roaming host.",
	},
	"bootfile": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_bootfile")),
		},
		MarkdownDescription: "The bootfile name for the roaming host. You can configure the DHCP server to support clients that use the boot file name option in their DHCPREQUEST messages.",
	},
	"bootserver": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_bootserver")),
		},
		MarkdownDescription: "The boot server address for the roaming host. You can specify the name and/or IP address of the boot server that the host needs to boot. The boot server IPv4 Address or name in FQDN format.",
	},
	"client_identifier_prepend_zero": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This field controls whether there is a prepend for the dhcp-client-identifier of a roaming host.",
	},
	"comment": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 256),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Comment for the roaming host; maximum 256 characters.",
	},
	"ddns_domainname": schema.StringAttribute{
		CustomType: internaltypes.CaseInsensitiveString{},
		Computed:   true,
		Optional:   true,
		Default:    stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_ddns_domainname")),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The DDNS domain name for this roaming host.",
	},
	"ddns_hostname": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The DDNS host name for this roaming host.",
	},
	"deny_bootp": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_deny_bootp")),
		},
		MarkdownDescription: "If set to true, BOOTP settings are disabled and BOOTP requests will be denied.",
	},
	"dhcp_client_identifier": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The DHCP client ID for the roaming host.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether a roaming host is disabled or not. When this is set to False, the roaming host is enabled.",
	},
	"enable_ddns": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_enable_ddns")),
		},
		MarkdownDescription: "The dynamic DNS updates flag of the roaming host object. If set to True, the DHCP server sends DDNS updates to DNS servers in the same Grid, and to external DNS servers.",
	},
	"enable_pxe_lease_time": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("pxe_lease_time")),
		},
		MarkdownDescription: "Set this to True if you want the DHCP server to use a different lease time for PXE clients.",
	},
	"extattrs": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Extensible attributes associated with the object.",
	},
	"extattrs_all": schema.MapAttribute{
		Computed: true,
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
		MarkdownDescription: "Extensible attributes associated with the object, including default and internal attributes.",
		ElementType:         types.StringType,
	},
	"force_roaming_hostname": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Set this to True to use the roaming host name as its ddns_hostname.",
	},
	"ignore_dhcp_option_list_request": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_ignore_dhcp_option_list_request")),
		},
		MarkdownDescription: "If this field is set to False, the appliance returns all the DHCP options the client is eligible to receive, rather than only the list of options the client has requested.",
	},
	"ipv6_client_hostname": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The client hostname of a DHCP roaming host object. This field specifies the host name that the DHCP client sends to the Infoblox appliance using DHCP option 12.",
	},
	"ipv6_ddns_domainname": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_ipv6_ddns_domainname")),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The IPv6 DDNS domain name for this roaming host.",
	},
	"ipv6_ddns_hostname": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The IPv6 DDNS host name for this roaming host.",
	},
	"ipv6_domain_name": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_ipv6_domain_name")),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The IPv6 domain name for this roaming host.",
	},
	"ipv6_domain_name_servers": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
		Default:     listdefault.StaticValue(types.ListNull(types.StringType)),
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
			listvalidator.AlsoRequires(path.MatchRoot("use_ipv6_domain_name_servers")),
			listvalidator.ValueStringsAre(customvalidator.IsValidIPv6Address()),
		},
		MarkdownDescription: "The IPv6 addresses of DNS recursive name servers to which the DHCP client can send name resolution requests. The DHCP server includes this information in the DNS Recursive Name Server option in Advertise, Rebind, Information-Request, and Reply messages.",
	},
	"ipv6_duid": schema.StringAttribute{
		CustomType: internaltypes.DUIDType{},
		Computed:   true,
		Optional:   true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The DUID value for this roaming host.",
	},
	"ipv6_enable_ddns": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_ipv6_enable_ddns")),
		},
		MarkdownDescription: "Set this to True to enable IPv6 DDNS.",
	},
	"ipv6_force_roaming_hostname": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Set this to True to use the roaming host name as its ddns_hostname.",
	},
	"ipv6_mac_address": schema.StringAttribute{
		CustomType: internaltypes.MACAddressType{},
		Computed:   true,
		Optional:   true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
			customvalidator.IsValidMacAddress(),
		},
		MarkdownDescription: "The MAC address for this roaming host.",
	},
	"ipv6_match_option": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Validators: []validator.String{
			stringvalidator.OneOf("DUID", "V6_MAC_ADDRESS"),
		},
		MarkdownDescription: "The identification method for an IPv6 or mixed IPv4/IPv6 roaming host. The supported values for this field are \"DUID\" or \"V6_MAC_ADDRESS\", which specify what option should be used to identify the specific DHCPv6 client.",
	},
	"ipv6_options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RoaminghostIpv6OptionsResourceSchemaAttributes,
		},
		Computed: true,
		Optional: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
			listvalidator.AlsoRequires(path.MatchRoot("use_ipv6_options")),
		},
		MarkdownDescription: "An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object.",
	},
	"ipv6_template": schema.StringAttribute{
		Computed: true,
		Optional: true,
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
		MarkdownDescription: "If set on creation, the roaming host will be created according to the values specified in the named IPv6 roaming host template.",
	},
	"mac": schema.StringAttribute{
		CustomType: internaltypes.MACAddressType{},
		Computed:   true,
		Optional:   true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
			customvalidator.IsValidMacAddress(),
		},
		MarkdownDescription: "The MAC address for this roaming host.",
	},
	"match_client": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Validators: []validator.String{
			stringvalidator.OneOf("CLIENT_ID", "MAC_ADDRESS"),
		},
		MarkdownDescription: "The match-client value for this roaming host. Valid values are: \"MAC_ADDRESS\": The fixed IP address is leased to the matching MAC address. \"CLIENT_ID\": The fixed IP address is leased to the matching DHCP client identifier.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of this roaming host.",
	},
	"network_view": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Default:  stringdefault.StaticString("default"),
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
		MarkdownDescription: "The name of the network view in which this roaming host resides.",
	},
	"nextserver": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_nextserver")),
		},
		MarkdownDescription: "The name in FQDN and/or IPv4 Address format of the next server that the host needs to boot.",
	},
	"options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RoaminghostOptionsResourceSchemaAttributes,
		},
		Computed: true,
		Optional: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
			listvalidator.AlsoRequires(path.MatchRoot("use_options")),
		},
		MarkdownDescription: "An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object.",
	},
	"preferred_lifetime": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(27000),
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_preferred_lifetime")),
		},
		MarkdownDescription: "The preferred lifetime value for this roaming host object.",
	},
	"pxe_lease_time": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_pxe_lease_time")),
		},
		MarkdownDescription: "The PXE lease time value for this roaming host object. Some hosts use PXE (Preboot Execution Environment) to boot remotely from a server. To better manage your IP resources, set a different lease time for PXE boot requests. You can configure the DHCP server to allocate an IP address with a shorter lease time to hosts that send PXE boot requests, so IP addresses are not leased longer than necessary. A 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached.",
	},
	"template": schema.StringAttribute{
		Computed: true,
		Optional: true,
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
		MarkdownDescription: "If set on creation, the roaming host will be created according to the values specified in the named template.",
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
	"use_ipv6_ddns_domainname": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: ipv6_ddns_domainname",
	},
	"use_ipv6_domain_name": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: ipv6_domain_name",
	},
	"use_ipv6_domain_name_servers": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: ipv6_domain_name_servers",
	},
	"use_ipv6_enable_ddns": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: ipv6_enable_ddns",
	},
	"use_ipv6_options": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: ipv6_options",
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
	"use_preferred_lifetime": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: preferred_lifetime",
	},
	"use_pxe_lease_time": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: pxe_lease_time",
	},
	"use_valid_lifetime": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: valid_lifetime",
	},
	"valid_lifetime": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_valid_lifetime")),
		},
		MarkdownDescription: "The valid lifetime value for this roaming host object.",
	},
}

func (m *RoaminghostModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *dhcp.Roaminghost {
	if m == nil {
		return nil
	}
	to := &dhcp.Roaminghost{
		AddressType:                    flex.ExpandStringPointer(m.AddressType),
		Bootfile:                       flex.ExpandStringPointer(m.Bootfile),
		Bootserver:                     flex.ExpandStringPointer(m.Bootserver),
		ClientIdentifierPrependZero:    flex.ExpandBoolPointer(m.ClientIdentifierPrependZero),
		Comment:                        flex.ExpandStringPointer(m.Comment),
		DdnsDomainname:                 flex.ExpandStringPointer(m.DdnsDomainname.StringValue),
		DdnsHostname:                   flex.ExpandStringPointer(m.DdnsHostname),
		DenyBootp:                      flex.ExpandBoolPointer(m.DenyBootp),
		DhcpClientIdentifier:           flex.ExpandStringPointer(m.DhcpClientIdentifier),
		Disable:                        flex.ExpandBoolPointer(m.Disable),
		EnableDdns:                     flex.ExpandBoolPointer(m.EnableDdns),
		EnablePxeLeaseTime:             flex.ExpandBoolPointer(m.EnablePxeLeaseTime),
		ExtAttrs:                       ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		ForceRoamingHostname:           flex.ExpandBoolPointer(m.ForceRoamingHostname),
		IgnoreDhcpOptionListRequest:    flex.ExpandBoolPointer(m.IgnoreDhcpOptionListRequest),
		Ipv6DdnsDomainname:             flex.ExpandStringPointer(m.Ipv6DdnsDomainname),
		Ipv6DdnsHostname:               flex.ExpandStringPointer(m.Ipv6DdnsHostname),
		Ipv6DomainName:                 flex.ExpandStringPointer(m.Ipv6DomainName),
		Ipv6DomainNameServers:          flex.ExpandFrameworkListString(ctx, m.Ipv6DomainNameServers, diags),
		Ipv6Duid:                       flex.ExpandDUID(m.Ipv6Duid),
		Ipv6EnableDdns:                 flex.ExpandBoolPointer(m.Ipv6EnableDdns),
		Ipv6ForceRoamingHostname:       flex.ExpandBoolPointer(m.Ipv6ForceRoamingHostname),
		Ipv6MacAddress:                 flex.ExpandMACAddr(m.Ipv6MacAddress),
		Ipv6MatchOption:                flex.ExpandStringPointer(m.Ipv6MatchOption),
		Ipv6Options:                    flex.ExpandFrameworkListNestedBlock(ctx, m.Ipv6Options, diags, ExpandRoaminghostIpv6Options),
		Mac:                            flex.ExpandMACAddr(m.Mac),
		MatchClient:                    flex.ExpandStringPointer(m.MatchClient),
		Name:                           flex.ExpandStringPointer(m.Name),
		Nextserver:                     flex.ExpandStringPointer(m.Nextserver),
		Options:                        flex.ExpandFrameworkListNestedBlock(ctx, m.Options, diags, ExpandRoaminghostOptions),
		PreferredLifetime:              flex.ExpandInt64Pointer(m.PreferredLifetime),
		PxeLeaseTime:                   flex.ExpandInt64Pointer(m.PxeLeaseTime),
		UseBootfile:                    flex.ExpandBoolPointer(m.UseBootfile),
		UseBootserver:                  flex.ExpandBoolPointer(m.UseBootserver),
		UseDdnsDomainname:              flex.ExpandBoolPointer(m.UseDdnsDomainname),
		UseDenyBootp:                   flex.ExpandBoolPointer(m.UseDenyBootp),
		UseEnableDdns:                  flex.ExpandBoolPointer(m.UseEnableDdns),
		UseIgnoreDhcpOptionListRequest: flex.ExpandBoolPointer(m.UseIgnoreDhcpOptionListRequest),
		UseIpv6DdnsDomainname:          flex.ExpandBoolPointer(m.UseIpv6DdnsDomainname),
		UseIpv6DomainName:              flex.ExpandBoolPointer(m.UseIpv6DomainName),
		UseIpv6DomainNameServers:       flex.ExpandBoolPointer(m.UseIpv6DomainNameServers),
		UseIpv6EnableDdns:              flex.ExpandBoolPointer(m.UseIpv6EnableDdns),
		UseIpv6Options:                 flex.ExpandBoolPointer(m.UseIpv6Options),
		UseNextserver:                  flex.ExpandBoolPointer(m.UseNextserver),
		UseOptions:                     flex.ExpandBoolPointer(m.UseOptions),
		UsePreferredLifetime:           flex.ExpandBoolPointer(m.UsePreferredLifetime),
		UsePxeLeaseTime:                flex.ExpandBoolPointer(m.UsePxeLeaseTime),
		UseValidLifetime:               flex.ExpandBoolPointer(m.UseValidLifetime),
		ValidLifetime:                  flex.ExpandInt64Pointer(m.ValidLifetime),
	}
	if isCreate {
		to.Template = flex.ExpandStringPointer(m.Template)
		to.Ipv6Template = flex.ExpandStringPointer(m.Ipv6Template)
		to.NetworkView = flex.ExpandStringPointer(m.NetworkView)
	}
	return to
}

func FlattenRoaminghost(ctx context.Context, from *dhcp.Roaminghost, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RoaminghostAttrTypes)
	}
	m := RoaminghostModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, RoaminghostAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RoaminghostModel) Flatten(ctx context.Context, from *dhcp.Roaminghost, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RoaminghostModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.AddressType = flex.FlattenStringPointer(from.AddressType)
	m.Bootfile = flex.FlattenStringPointer(from.Bootfile)
	m.Bootserver = flex.FlattenStringPointer(from.Bootserver)
	m.ClientIdentifierPrependZero = types.BoolPointerValue(from.ClientIdentifierPrependZero)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.DdnsDomainname.StringValue = flex.FlattenStringPointer(from.DdnsDomainname)
	m.DdnsHostname = flex.FlattenStringPointer(from.DdnsHostname)
	m.DenyBootp = types.BoolPointerValue(from.DenyBootp)
	m.DhcpClientIdentifier = flex.FlattenStringPointer(from.DhcpClientIdentifier)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.EnableDdns = types.BoolPointerValue(from.EnableDdns)
	m.EnablePxeLeaseTime = types.BoolPointerValue(from.EnablePxeLeaseTime)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.ForceRoamingHostname = types.BoolPointerValue(from.ForceRoamingHostname)
	m.IgnoreDhcpOptionListRequest = types.BoolPointerValue(from.IgnoreDhcpOptionListRequest)
	m.Ipv6ClientHostname = flex.FlattenStringPointer(from.Ipv6ClientHostname)
	m.Ipv6DdnsDomainname = flex.FlattenStringPointer(from.Ipv6DdnsDomainname)
	m.Ipv6DdnsHostname = flex.FlattenStringPointer(from.Ipv6DdnsHostname)
	m.Ipv6DomainName = flex.FlattenStringPointer(from.Ipv6DomainName)
	m.Ipv6DomainNameServers = flex.FlattenFrameworkListString(ctx, from.Ipv6DomainNameServers, diags)
	m.Ipv6Duid = flex.FlattenDUID(from.Ipv6Duid)
	m.Ipv6EnableDdns = types.BoolPointerValue(from.Ipv6EnableDdns)
	m.Ipv6ForceRoamingHostname = types.BoolPointerValue(from.Ipv6ForceRoamingHostname)
	m.Ipv6MacAddress = flex.FlattenMACAddr(from.Ipv6MacAddress)
	m.Ipv6MatchOption = flex.FlattenStringPointer(from.Ipv6MatchOption)
	planIpv6Options := m.Ipv6Options
	m.Ipv6Options = flex.FlattenFrameworkListNestedBlock(ctx, from.Ipv6Options, RoaminghostIpv6OptionsAttrTypes, diags, FlattenRoaminghostIpv6Options)
	if !planIpv6Options.IsUnknown() {
		reOrderedOptions, diags := utils.ReorderAndFilterDHCPOptions(ctx, planIpv6Options, m.Ipv6Options)
		if !diags.HasError() {
			m.Ipv6Options = reOrderedOptions.(basetypes.ListValue)
		}
	}
	m.Mac = flex.FlattenMACAddr(from.Mac)
	m.MatchClient = flex.FlattenStringPointer(from.MatchClient)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.NetworkView = flex.FlattenStringPointer(from.NetworkView)
	m.Nextserver = flex.FlattenStringPointer(from.Nextserver)
	planOptions := m.Options
	m.Options = flex.FlattenFrameworkListNestedBlock(ctx, from.Options, RoaminghostOptionsAttrTypes, diags, FlattenRoaminghostOptions)
	if !planOptions.IsUnknown() {
		reOrderedOptions, diags := utils.ReorderAndFilterDHCPOptions(ctx, planOptions, m.Options)
		if !diags.HasError() {
			m.Options = reOrderedOptions.(basetypes.ListValue)
		}
	}
	m.PreferredLifetime = flex.FlattenInt64Pointer(from.PreferredLifetime)
	m.PxeLeaseTime = flex.FlattenInt64Pointer(from.PxeLeaseTime)
	m.Template = flex.FlattenStringPointer(from.Template)
	m.Ipv6Template = flex.FlattenStringPointer(from.Ipv6Template)
	m.UseBootfile = types.BoolPointerValue(from.UseBootfile)
	m.UseBootserver = types.BoolPointerValue(from.UseBootserver)
	m.UseDdnsDomainname = types.BoolPointerValue(from.UseDdnsDomainname)
	m.UseDenyBootp = types.BoolPointerValue(from.UseDenyBootp)
	m.UseEnableDdns = types.BoolPointerValue(from.UseEnableDdns)
	m.UseIgnoreDhcpOptionListRequest = types.BoolPointerValue(from.UseIgnoreDhcpOptionListRequest)
	m.UseIpv6DdnsDomainname = types.BoolPointerValue(from.UseIpv6DdnsDomainname)
	m.UseIpv6DomainName = types.BoolPointerValue(from.UseIpv6DomainName)
	m.UseIpv6DomainNameServers = types.BoolPointerValue(from.UseIpv6DomainNameServers)
	m.UseIpv6EnableDdns = types.BoolPointerValue(from.UseIpv6EnableDdns)
	m.UseIpv6Options = types.BoolPointerValue(from.UseIpv6Options)
	m.UseNextserver = types.BoolPointerValue(from.UseNextserver)
	m.UseOptions = types.BoolPointerValue(from.UseOptions)
	m.UsePreferredLifetime = types.BoolPointerValue(from.UsePreferredLifetime)
	m.UsePxeLeaseTime = types.BoolPointerValue(from.UsePxeLeaseTime)
	m.UseValidLifetime = types.BoolPointerValue(from.UseValidLifetime)
	m.ValidLifetime = flex.FlattenInt64Pointer(from.ValidLifetime)
}
