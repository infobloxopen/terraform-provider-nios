package dhcp

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	planmodifiers "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/immutable"
	internaltypes "github.com/infobloxopen/terraform-provider-nios/internal/types"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/import"
)

type Ipv6fixedaddressModel struct {
	Ref                      types.String                  `tfsdk:"ref"`
	AddressType              types.String                  `tfsdk:"address_type"`
	AllowTelnet              types.Bool                    `tfsdk:"allow_telnet"`
	CliCredentials           types.List                    `tfsdk:"cli_credentials"`
	CloudInfo                types.Object                  `tfsdk:"cloud_info"`
	Comment                  types.String                  `tfsdk:"comment"`
	DeviceDescription        types.String                  `tfsdk:"device_description"`
	DeviceLocation           types.String                  `tfsdk:"device_location"`
	DeviceType               types.String                  `tfsdk:"device_type"`
	DeviceVendor             types.String                  `tfsdk:"device_vendor"`
	Disable                  types.Bool                    `tfsdk:"disable"`
	DisableDiscovery         types.Bool                    `tfsdk:"disable_discovery"`
	DiscoverNowStatus        types.String                  `tfsdk:"discover_now_status"`
	DiscoveredData           types.Object                  `tfsdk:"discovered_data"`
	DomainName               types.String                  `tfsdk:"domain_name"`
	DomainNameServers        types.List                    `tfsdk:"domain_name_servers"`
	Duid                     types.String                  `tfsdk:"duid"`
	EnableImmediateDiscovery types.Bool                    `tfsdk:"enable_immediate_discovery"`
	ExtAttrs                 types.Map                     `tfsdk:"extattrs"`
	Ipv6addr                 iptypes.IPv6Address           `tfsdk:"ipv6addr"`
	FuncCall                 types.Object                  `tfsdk:"func_call"`
	Ipv6prefix               types.String                  `tfsdk:"ipv6prefix"`
	Ipv6prefixBits           types.Int64                   `tfsdk:"ipv6prefix_bits"`
	LogicFilterRules         types.List                    `tfsdk:"logic_filter_rules"`
	MacAddress               internaltypes.MACAddressValue `tfsdk:"mac_address"`
	MatchClient              types.String                  `tfsdk:"match_client"`
	MsAdUserData             types.Object                  `tfsdk:"ms_ad_user_data"`
	Name                     types.String                  `tfsdk:"name"`
	Network                  types.String                  `tfsdk:"network"`
	NetworkView              types.String                  `tfsdk:"network_view"`
	Options                  types.List                    `tfsdk:"options"`
	PreferredLifetime        types.Int64                   `tfsdk:"preferred_lifetime"`
	ReservedInterface        types.String                  `tfsdk:"reserved_interface"`
	RestartIfNeeded          types.Bool                    `tfsdk:"restart_if_needed"`
	Snmp3Credential          types.Object                  `tfsdk:"snmp3_credential"`
	SnmpCredential           types.Object                  `tfsdk:"snmp_credential"`
	Template                 types.String                  `tfsdk:"template"`
	UseCliCredentials        types.Bool                    `tfsdk:"use_cli_credentials"`
	UseDomainName            types.Bool                    `tfsdk:"use_domain_name"`
	UseDomainNameServers     types.Bool                    `tfsdk:"use_domain_name_servers"`
	UseLogicFilterRules      types.Bool                    `tfsdk:"use_logic_filter_rules"`
	UseOptions               types.Bool                    `tfsdk:"use_options"`
	UsePreferredLifetime     types.Bool                    `tfsdk:"use_preferred_lifetime"`
	UseSnmp3Credential       types.Bool                    `tfsdk:"use_snmp3_credential"`
	UseSnmpCredential        types.Bool                    `tfsdk:"use_snmp_credential"`
	UseValidLifetime         types.Bool                    `tfsdk:"use_valid_lifetime"`
	ValidLifetime            types.Int64                   `tfsdk:"valid_lifetime"`
	ExtAttrsAll              types.Map                     `tfsdk:"extattrs_all"`
}

var Ipv6fixedaddressAttrTypes = map[string]attr.Type{
	"ref":                        types.StringType,
	"address_type":               types.StringType,
	"allow_telnet":               types.BoolType,
	"cli_credentials":            types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6fixedaddressCliCredentialsAttrTypes}},
	"cloud_info":                 types.ObjectType{AttrTypes: Ipv6fixedaddressCloudInfoAttrTypes},
	"comment":                    types.StringType,
	"device_description":         types.StringType,
	"device_location":            types.StringType,
	"device_type":                types.StringType,
	"device_vendor":              types.StringType,
	"disable":                    types.BoolType,
	"disable_discovery":          types.BoolType,
	"discover_now_status":        types.StringType,
	"discovered_data":            types.ObjectType{AttrTypes: Ipv6fixedaddressDiscoveredDataAttrTypes},
	"domain_name":                types.StringType,
	"domain_name_servers":        types.ListType{ElemType: types.StringType},
	"duid":                       types.StringType,
	"enable_immediate_discovery": types.BoolType,
	"extattrs":                   types.MapType{ElemType: types.StringType},
	"ipv6addr":                   iptypes.IPv6AddressType{},
	"func_call":                  types.ObjectType{AttrTypes: FuncCallAttrTypes},
	"ipv6prefix":                 types.StringType,
	"ipv6prefix_bits":            types.Int64Type,
	"logic_filter_rules":         types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6fixedaddressLogicFilterRulesAttrTypes}},
	"mac_address":                internaltypes.MACAddressType{},
	"match_client":               types.StringType,
	"ms_ad_user_data":            types.ObjectType{AttrTypes: Ipv6fixedaddressMsAdUserDataAttrTypes},
	"name":                       types.StringType,
	"network":                    types.StringType,
	"network_view":               types.StringType,
	"options":                    types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6fixedaddressOptionsAttrTypes}},
	"preferred_lifetime":         types.Int64Type,
	"reserved_interface":         types.StringType,
	"restart_if_needed":          types.BoolType,
	"snmp3_credential":           types.ObjectType{AttrTypes: Ipv6fixedaddressSnmp3CredentialAttrTypes},
	"snmp_credential":            types.ObjectType{AttrTypes: Ipv6fixedaddressSnmpCredentialAttrTypes},
	"template":                   types.StringType,
	"use_cli_credentials":        types.BoolType,
	"use_domain_name":            types.BoolType,
	"use_domain_name_servers":    types.BoolType,
	"use_logic_filter_rules":     types.BoolType,
	"use_options":                types.BoolType,
	"use_preferred_lifetime":     types.BoolType,
	"use_snmp3_credential":       types.BoolType,
	"use_snmp_credential":        types.BoolType,
	"use_valid_lifetime":         types.BoolType,
	"valid_lifetime":             types.Int64Type,
	"extattrs_all":               types.MapType{ElemType: types.StringType},
}

var Ipv6fixedaddressResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"address_type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("ADDRESS", "BOTH", "PREFIX"),
		},
		Default:             stringdefault.StaticString("ADDRESS"),
		MarkdownDescription: "The address type value for this IPv6 fixed address. When the address type is \"ADDRESS\", a value for the 'ipv6addr' member is required. When the address type is \"PREFIX\", values for 'ipv6prefix' and 'ipv6prefix_bits' are required. When the address type is \"BOTH\", values for 'ipv6addr', 'ipv6prefix', and 'ipv6prefix_bits' are all required.",
	},
	"allow_telnet": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This field controls whether the credential is used for both the Telnet and SSH credentials. If set to False, the credential is used only for SSH.",
	},
	"cli_credentials": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6fixedaddressCliCredentialsResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
			listvalidator.AlsoRequires(path.MatchRoot("use_cli_credentials")),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The CLI credentials for the IPv6 fixed address.",
	},
	"cloud_info": schema.SingleNestedAttribute{
		Attributes: Ipv6fixedaddressCloudInfoResourceSchemaAttributes,
		Computed:   true,
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Comment for the fixed address; maximum 256 characters.",
	},
	"device_description": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The description of the device.",
	},
	"device_location": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The location of the device.",
	},
	"device_type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The type of the device.",
	},
	"device_vendor": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The vendor of the device.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether a fixed address is disabled or not. When this is set to False, the IPv6 fixed address is enabled.",
	},
	"disable_discovery": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the discovery for this IPv6 fixed address is disabled or not. False means that the discovery is enabled.",
	},
	"discover_now_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The discovery status of this IPv6 fixed address.",
	},
	"discovered_data": schema.SingleNestedAttribute{
		Attributes: Ipv6fixedaddressDiscoveredDataResourceSchemaAttributes,
		Computed:   true,
	},
	"domain_name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		//Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
			stringvalidator.AlsoRequires(path.MatchRoot("use_domain_name")),
		},
		MarkdownDescription: "The domain name for this IPv6 fixed address.",
	},
	"domain_name_servers": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
			listvalidator.AlsoRequires(path.MatchRoot("use_domain_name_servers")),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The IPv6 addresses of DNS recursive name servers to which the DHCP client can send name resolution requests. The DHCP server includes this information in the DNS Recursive Name Server option in Advertise, Rebind, Information-Request, and Reply messages.",
	},
	"duid": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The DUID value for this IPv6 fixed address.",
	},
	"enable_immediate_discovery": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Determines if the discovery for the IPv6 fixed address should be immediately enabled.",
	},
	"extattrs": schema.MapAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.",
	},
	"ipv6addr": schema.StringAttribute{
		CustomType: iptypes.IPv6AddressType{},
		Optional:   true,
		Computed:   true,
		Validators: []validator.String{
			stringvalidator.ExactlyOneOf(
				path.MatchRoot("ipv6addr"),
				path.MatchRoot("func_call"),
				path.MatchRoot("ipv6prefix"),
			),
		},
	},
	"func_call": schema.SingleNestedAttribute{
		Attributes:          FuncCallResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Specifies the function call to execute. The `next_available_ip` function is supported for IPV6 Fixed Address.",
	},
	"ipv6prefix": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The IPv6 Address prefix of the DHCP IPv6 fixed address.",
	},
	"ipv6prefix_bits": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Prefix bits of the DHCP IPv6 fixed address.",
	},
	"logic_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6fixedaddressLogicFilterRulesResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
			listvalidator.AlsoRequires(path.MatchRoot("use_logic_filter_rules")),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "This field contains the logic filters to be applied to this IPv6 fixed address. This list corresponds to the match rules that are written to the DHCPv6 configuration file.",
	},
	"mac_address": schema.StringAttribute{
		CustomType: internaltypes.MACAddressType{},
		Optional:   true,
		Computed:   true,
		Validators: []validator.String{
			stringvalidator.ExactlyOneOf(
				path.MatchRoot("duid")),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The MAC address for this IPv6 fixed address.",
	},
	"match_client": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("DUID"),
		Validators: []validator.String{
			stringvalidator.OneOf("DUID", "MAC_ADDRESS"),
		},
		MarkdownDescription: "The match_client value for this fixed address. Valid values are: \"DUID\": The fixed IP address is leased to the matching DUID. \"MAC_ADDRESS\": The fixed IP address is leased to the matching MAC address.",
	},
	"ms_ad_user_data": schema.SingleNestedAttribute{
		Attributes: Ipv6fixedaddressMsAdUserDataResourceSchemaAttributes,
		Computed:   true,
	},
	"name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "This field contains the name of this IPv6 fixed address.",
	},
	"network": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.IsValidIPCIDR(),
		},
		MarkdownDescription: "The network to which this IPv6 fixed address belongs, in IPv6 Address/CIDR format.",
	},
	"network_view": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString("default"),
		MarkdownDescription: "The name of the network view in which this IPv6 fixed address resides.",
	},
	"options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6fixedaddressOptionsResourceSchemaAttributes,
		},
		Computed: true,
		Optional: true,
		Default: listdefault.StaticValue(
			types.ListValueMust(
				types.ObjectType{AttrTypes: Ipv6fixedaddressOptionsAttrTypes},
				[]attr.Value{},
			),
		),
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_options")),
			listvalidator.SizeAtLeast(1),
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
		MarkdownDescription: "The preferred lifetime value for this DHCP IPv6 fixed address object.",
	},
	"reserved_interface": schema.StringAttribute{
		Optional: true,
		Computed: true,
		//Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The reference to the reserved interface to which the device belongs.",
	},
	"restart_if_needed": schema.BoolAttribute{
		Optional: true,
		//Computed:            true,
		//Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Restarts the member service. The restart_if_needed flag can trigger a restart on DHCP services only when it is enabled on CP member.",
	},
	"snmp3_credential": schema.SingleNestedAttribute{
		Attributes: Ipv6fixedaddressSnmp3CredentialResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Validators: []validator.Object{
			objectvalidator.AlsoRequires(path.MatchRoot("use_snmp3_credential")),
			objectvalidator.AlsoRequires(path.MatchRoot("use_cli_credentials")),
		},
	},
	"snmp_credential": schema.SingleNestedAttribute{
		Attributes: Ipv6fixedaddressSnmpCredentialResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Validators: []validator.Object{
			objectvalidator.AlsoRequires(path.MatchRoot("use_snmp_credential")),
		},
	},
	"template": schema.StringAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
		MarkdownDescription: "If set on creation, the IPv6 fixed address will be created according to the values specified in the named template.",
	},
	"use_cli_credentials": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If set to true, the CLI credential will override member-level settings.",
	},
	"use_domain_name": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: domain_name",
	},
	"use_domain_name_servers": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: domain_name_servers",
	},
	"use_logic_filter_rules": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: logic_filter_rules",
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
	"use_snmp3_credential": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the SNMPv3 credential should be used for the IPv6 fixed address.",
	},
	"use_snmp_credential": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If set to true, SNMP credential will override member level settings.",
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
		Default:  int64default.StaticInt64(43200),
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_valid_lifetime")),
		},
		MarkdownDescription: "The valid lifetime value for this DHCP IPv6 Fixed Address object.",
	},
	"extattrs_all": schema.MapAttribute{
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object, including default attributes.",
		ElementType:         types.StringType,
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
	},
}

func (m *Ipv6fixedaddressModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *dhcp.Ipv6fixedaddress {
	if m == nil {
		return nil
	}
	to := &dhcp.Ipv6fixedaddress{
		AddressType:              flex.ExpandStringPointer(m.AddressType),
		AllowTelnet:              flex.ExpandBoolPointer(m.AllowTelnet),
		CliCredentials:           flex.ExpandFrameworkListNestedBlock(ctx, m.CliCredentials, diags, ExpandIpv6fixedaddressCliCredentials),
		CloudInfo:                ExpandIpv6fixedaddressCloudInfo(ctx, m.CloudInfo, diags),
		Comment:                  flex.ExpandStringPointer(m.Comment),
		DeviceDescription:        flex.ExpandStringPointer(m.DeviceDescription),
		DeviceLocation:           flex.ExpandStringPointer(m.DeviceLocation),
		DeviceType:               flex.ExpandStringPointer(m.DeviceType),
		DeviceVendor:             flex.ExpandStringPointer(m.DeviceVendor),
		Disable:                  flex.ExpandBoolPointer(m.Disable),
		DisableDiscovery:         flex.ExpandBoolPointer(m.DisableDiscovery),
		DiscoveredData:           ExpandIpv6fixedaddressDiscoveredData(ctx, m.DiscoveredData, diags),
		DomainName:               flex.ExpandStringPointer(m.DomainName),
		DomainNameServers:        flex.ExpandFrameworkListString(ctx, m.DomainNameServers, diags),
		Duid:                     flex.ExpandStringPointer(m.Duid),
		EnableImmediateDiscovery: flex.ExpandBoolPointer(m.EnableImmediateDiscovery),
		ExtAttrs:                 ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		Ipv6addr:                 ExpandIpv6fixedaddressIpv6addr(m.Ipv6addr),
		FuncCall:                 ExpandFuncCall(ctx, m.FuncCall, diags),
		Ipv6prefix:               flex.ExpandStringPointer(m.Ipv6prefix),
		Ipv6prefixBits:           flex.ExpandInt64Pointer(m.Ipv6prefixBits),
		LogicFilterRules:         flex.ExpandFrameworkListNestedBlock(ctx, m.LogicFilterRules, diags, ExpandIpv6fixedaddressLogicFilterRules),
		MacAddress:               flex.ExpandMACAddr(m.MacAddress),
		MatchClient:              flex.ExpandStringPointer(m.MatchClient),
		MsAdUserData:             ExpandIpv6fixedaddressMsAdUserData(ctx, m.MsAdUserData, diags),
		Name:                     flex.ExpandStringPointer(m.Name),
		Network:                  flex.ExpandStringPointer(m.Network),
		NetworkView:              flex.ExpandStringPointer(m.NetworkView),
		Options:                  flex.ExpandFrameworkListNestedBlock(ctx, m.Options, diags, ExpandIpv6fixedaddressOptions),
		PreferredLifetime:        flex.ExpandInt64Pointer(m.PreferredLifetime),
		ReservedInterface:        flex.ExpandStringPointer(m.ReservedInterface),
		RestartIfNeeded:          flex.ExpandBoolPointer(m.RestartIfNeeded),
		Snmp3Credential:          ExpandIpv6fixedaddressSnmp3Credential(ctx, m.Snmp3Credential, diags),
		SnmpCredential:           ExpandIpv6fixedaddressSnmpCredential(ctx, m.SnmpCredential, diags),
		UseCliCredentials:        flex.ExpandBoolPointer(m.UseCliCredentials),
		UseDomainName:            flex.ExpandBoolPointer(m.UseDomainName),
		UseDomainNameServers:     flex.ExpandBoolPointer(m.UseDomainNameServers),
		UseLogicFilterRules:      flex.ExpandBoolPointer(m.UseLogicFilterRules),
		UseOptions:               flex.ExpandBoolPointer(m.UseOptions),
		UsePreferredLifetime:     flex.ExpandBoolPointer(m.UsePreferredLifetime),
		UseSnmp3Credential:       flex.ExpandBoolPointer(m.UseSnmp3Credential),
		UseSnmpCredential:        flex.ExpandBoolPointer(m.UseSnmpCredential),
		UseValidLifetime:         flex.ExpandBoolPointer(m.UseValidLifetime),
		ValidLifetime:            flex.ExpandInt64Pointer(m.ValidLifetime),
	}
	if isCreate {
		to.Template = flex.ExpandStringPointer(m.Template)
	}
	return to
}

func FlattenIpv6fixedaddress(ctx context.Context, from *dhcp.Ipv6fixedaddress, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6fixedaddressAttrTypes)
	}
	m := Ipv6fixedaddressModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, Ipv6fixedaddressAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Ipv6fixedaddressModel) Flatten(ctx context.Context, from *dhcp.Ipv6fixedaddress, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Ipv6fixedaddressModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.AddressType = flex.FlattenStringPointer(from.AddressType)
	m.AllowTelnet = types.BoolPointerValue(from.AllowTelnet)
	planCredentials := m.CliCredentials
	m.CliCredentials = flex.FlattenFrameworkListNestedBlock(ctx, from.CliCredentials, Ipv6fixedaddressCliCredentialsAttrTypes, diags, FlattenIpv6fixedaddressCliCredentials)
	if !planCredentials.IsUnknown() {
		credentialVal, diags := utils.CopyFieldFromPlanToRespList(ctx, planCredentials, m.CliCredentials, "password")
		if !diags.HasError() {
			m.CliCredentials = credentialVal.(basetypes.ListValue)
			reOrderedCredentials, diags := utils.ReorderAndFilterNestedListResponse(ctx, planCredentials, m.CliCredentials, "credential_type")
			if !diags.HasError() {
				m.CliCredentials = reOrderedCredentials.(basetypes.ListValue)
			}
		}
	}
	m.CloudInfo = FlattenIpv6fixedaddressCloudInfo(ctx, from.CloudInfo, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.DeviceDescription = flex.FlattenStringPointer(from.DeviceDescription)
	m.DeviceLocation = flex.FlattenStringPointer(from.DeviceLocation)
	m.DeviceType = flex.FlattenStringPointer(from.DeviceType)
	m.DeviceVendor = flex.FlattenStringPointer(from.DeviceVendor)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.DisableDiscovery = types.BoolPointerValue(from.DisableDiscovery)
	m.DiscoverNowStatus = flex.FlattenStringPointer(from.DiscoverNowStatus)
	m.DiscoveredData = FlattenIpv6fixedaddressDiscoveredData(ctx, from.DiscoveredData, diags)
	m.DomainName = flex.FlattenStringPointer(from.DomainName)
	m.DomainNameServers = flex.FlattenFrameworkListString(ctx, from.DomainNameServers, diags)
	m.Duid = flex.FlattenStringPointer(from.Duid)
	m.EnableImmediateDiscovery = types.BoolPointerValue(from.EnableImmediateDiscovery)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.Ipv6addr = FlattenIpv6fixedaddressIpv6addr(from.Ipv6addr)
	m.FuncCall = FlattenFuncCall(ctx, from.FuncCall, diags)
	m.Ipv6prefix = flex.FlattenStringPointer(from.Ipv6prefix)
	m.Ipv6prefixBits = flex.FlattenInt64Pointer(from.Ipv6prefixBits)
	m.LogicFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.LogicFilterRules, Ipv6fixedaddressLogicFilterRulesAttrTypes, diags, FlattenIpv6fixedaddressLogicFilterRules)
	m.MacAddress = flex.FlattenMACAddr(from.MacAddress)
	m.MatchClient = flex.FlattenStringPointer(from.MatchClient)
	m.MsAdUserData = FlattenIpv6fixedaddressMsAdUserData(ctx, from.MsAdUserData, diags)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Network = flex.FlattenStringPointer(from.Network)
	m.NetworkView = flex.FlattenStringPointer(from.NetworkView)
	planOptions := m.Options
	m.Options = flex.FlattenFrameworkListNestedBlock(ctx, from.Options, Ipv6fixedaddressOptionsAttrTypes, diags, FlattenIpv6fixedaddressOptions)
	if !planOptions.IsUnknown() {
		reOrderedOptions, diags := utils.ReorderAndFilterDHCPOptions(ctx, planOptions, m.Options)
		if !diags.HasError() {
			m.Options = reOrderedOptions.(basetypes.ListValue)
		}
	}
	m.PreferredLifetime = flex.FlattenInt64Pointer(from.PreferredLifetime)
	m.ReservedInterface = flex.FlattenStringPointer(from.ReservedInterface)
	m.RestartIfNeeded = types.BoolPointerValue(from.RestartIfNeeded)
	planSnmp3Credential := m.Snmp3Credential
	m.Snmp3Credential = FlattenIpv6fixedaddressSnmp3Credential(ctx, from.Snmp3Credential, diags)
	if !planSnmp3Credential.IsUnknown() {
		snmp3CredentialVal, diags := utils.CopyFieldFromPlanToRespObject(ctx, planSnmp3Credential, m.Snmp3Credential, "privacy_password")
		if !diags.HasError() {
			m.Snmp3Credential = snmp3CredentialVal.(types.Object)
		}
		snmp3CredentialVal2, diags := utils.CopyFieldFromPlanToRespObject(ctx, planSnmp3Credential, m.Snmp3Credential, "authentication_password")
		if !diags.HasError() {
			m.Snmp3Credential = snmp3CredentialVal2.(types.Object)
		}
	}
	m.SnmpCredential = FlattenIpv6fixedaddressSnmpCredential(ctx, from.SnmpCredential, diags)
	m.Template = flex.FlattenStringPointer(from.Template)
	m.UseCliCredentials = types.BoolPointerValue(from.UseCliCredentials)
	m.UseDomainName = types.BoolPointerValue(from.UseDomainName)
	m.UseDomainNameServers = types.BoolPointerValue(from.UseDomainNameServers)
	m.UseLogicFilterRules = types.BoolPointerValue(from.UseLogicFilterRules)
	m.UseOptions = types.BoolPointerValue(from.UseOptions)
	m.UsePreferredLifetime = types.BoolPointerValue(from.UsePreferredLifetime)
	m.UseSnmp3Credential = types.BoolPointerValue(from.UseSnmp3Credential)
	m.UseSnmpCredential = types.BoolPointerValue(from.UseSnmpCredential)
	m.UseValidLifetime = types.BoolPointerValue(from.UseValidLifetime)
	m.ValidLifetime = flex.FlattenInt64Pointer(from.ValidLifetime)
}

func ExpandIpv6fixedaddressIpv6addr(ipv6addr iptypes.IPv6Address) *dhcp.Ipv6fixedaddressIpv6addr {
	if ipv6addr.IsNull() {
		return &dhcp.Ipv6fixedaddressIpv6addr{}
	}
	var m dhcp.Ipv6fixedaddressIpv6addr
	m.String = flex.ExpandIPv6Address(ipv6addr)
	return &m
}

func FlattenIpv6fixedaddressIpv6addr(from *dhcp.Ipv6fixedaddressIpv6addr) iptypes.IPv6Address {
	if from.String == nil {
		return iptypes.NewIPv6AddressNull()
	}
	m := flex.FlattenIPv6Address(from.String)
	return m
}
