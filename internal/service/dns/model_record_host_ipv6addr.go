package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	"github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/plancontrol"
)

type RecordHostIpv6addrModel struct {
	Ref                  types.String        `tfsdk:"ref"`
	AddressType          types.String        `tfsdk:"address_type"`
	ConfigureForDhcp     types.Bool          `tfsdk:"configure_for_dhcp"`
	DiscoverNowStatus    types.String        `tfsdk:"discover_now_status"`
	DiscoveredData       types.Object        `tfsdk:"discovered_data"`
	DomainName           types.String        `tfsdk:"domain_name"`
	DomainNameServers    types.List          `tfsdk:"domain_name_servers"`
	Duid                 types.String        `tfsdk:"duid"`
	Host                 types.String        `tfsdk:"host"`
	Ipv6addr             iptypes.IPv6Address `tfsdk:"ipv6addr"`
	FuncCall             types.Object        `tfsdk:"func_call"`
	Ipv6prefix           types.String        `tfsdk:"ipv6prefix"`
	Ipv6prefixBits       types.Int64         `tfsdk:"ipv6prefix_bits"`
	LastQueried          types.Int64         `tfsdk:"last_queried"`
	LogicFilterRules     types.List          `tfsdk:"logic_filter_rules"`
	Mac                  types.String        `tfsdk:"mac"`
	MatchClient          types.String        `tfsdk:"match_client"`
	MsAdUserData         types.Object        `tfsdk:"ms_ad_user_data"`
	Network              types.String        `tfsdk:"network"`
	NetworkView          types.String        `tfsdk:"network_view"`
	Options              types.List          `tfsdk:"options"`
	PreferredLifetime    types.Int64         `tfsdk:"preferred_lifetime"`
	ReservedInterface    types.String        `tfsdk:"reserved_interface"`
	UseDomainName        types.Bool          `tfsdk:"use_domain_name"`
	UseDomainNameServers types.Bool          `tfsdk:"use_domain_name_servers"`
	UseForEaInheritance  types.Bool          `tfsdk:"use_for_ea_inheritance"`
	UseLogicFilterRules  types.Bool          `tfsdk:"use_logic_filter_rules"`
	UseOptions           types.Bool          `tfsdk:"use_options"`
	UsePreferredLifetime types.Bool          `tfsdk:"use_preferred_lifetime"`
	UseValidLifetime     types.Bool          `tfsdk:"use_valid_lifetime"`
	ValidLifetime        types.Int64         `tfsdk:"valid_lifetime"`
}

var RecordHostIpv6addrAttrTypes = map[string]attr.Type{
	"ref":                     types.StringType,
	"address_type":            types.StringType,
	"configure_for_dhcp":      types.BoolType,
	"discover_now_status":     types.StringType,
	"discovered_data":         types.ObjectType{AttrTypes: RecordHostIpv6addrDiscoveredDataAttrTypes},
	"domain_name":             types.StringType,
	"domain_name_servers":     types.ListType{ElemType: types.StringType},
	"duid":                    types.StringType,
	"host":                    types.StringType,
	"ipv6addr":                iptypes.IPv6AddressType{},
	"func_call":               types.ObjectType{AttrTypes: FuncCallAttrTypes},
	"ipv6prefix":              types.StringType,
	"ipv6prefix_bits":         types.Int64Type,
	"last_queried":            types.Int64Type,
	"logic_filter_rules":      types.ListType{ElemType: types.ObjectType{AttrTypes: RecordHostIpv6addrLogicFilterRulesAttrTypes}},
	"mac":                     types.StringType,
	"match_client":            types.StringType,
	"ms_ad_user_data":         types.ObjectType{AttrTypes: RecordHostIpv6addrMsAdUserDataAttrTypes},
	"network":                 types.StringType,
	"network_view":            types.StringType,
	"options":                 types.ListType{ElemType: types.ObjectType{AttrTypes: RecordHostIpv6addrOptionsAttrTypes}},
	"preferred_lifetime":      types.Int64Type,
	"reserved_interface":      types.StringType,
	"use_domain_name":         types.BoolType,
	"use_domain_name_servers": types.BoolType,
	"use_for_ea_inheritance":  types.BoolType,
	"use_logic_filter_rules":  types.BoolType,
	"use_options":             types.BoolType,
	"use_preferred_lifetime":  types.BoolType,
	"use_valid_lifetime":      types.BoolType,
	"valid_lifetime":          types.Int64Type,
}

var RecordHostIpv6addrResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"address_type": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Type of the DHCP IPv6 Host Address object.",
	},
	"configure_for_dhcp": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Set this to True to enable the DHCP configuration for this IPv6 host address.",
		PlanModifiers: []planmodifier.Bool{
			plancontrol.UseStateForUnknownBool(),
		},
	},
	"discover_now_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The discovery status of this IPv6 Host Address.",
		PlanModifiers: []planmodifier.String{
			plancontrol.UseStateForUnknownString(),
		},
	},
	"discovered_data": schema.SingleNestedAttribute{
		Attributes: RecordHostIpv6addrDiscoveredDataResourceSchemaAttributes,
		Computed:   true,
		PlanModifiers: []planmodifier.Object{
			plancontrol.UseStateForUnknownObject(),
		},
	},
	"domain_name": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Use this method to set or retrieve the domain_name value of the DHCP IPv6 Host Address object.",
	},
	"domain_name_servers": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The IPv6 addresses of DNS recursive name servers to which the DHCP client can send name resolution requests. The DHCP server includes this information in the DNS Recursive Name Server option in Advertise, Rebind, Information-Request, and Reply messages.",
	},
	"duid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "DHCPv6 Unique Identifier (DUID) of the address object.",
		PlanModifiers: []planmodifier.String{
			plancontrol.UseStateForUnknownString(),
		},
	},
	"host": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The host to which the IPv6 host address belongs, in FQDN format. It is only present when the host address object is not returned as part of a host.",
		PlanModifiers: []planmodifier.String{
			plancontrol.UseStateForUnknownString(),
		},
	},
	"ipv6addr": schema.StringAttribute{
		CustomType:          iptypes.IPv6AddressType{},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The IPv6 Address of the record.",
	},
	"func_call": schema.SingleNestedAttribute{
		Attributes: FuncCallResourceSchemaAttributes,
		Optional:   true,
		Validators: []validator.Object{
			objectvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("ipv6addr")),
		},
		MarkdownDescription: "Function call to be executed for Fixed Address",
	},
	"ipv6prefix": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The IPv6 Address prefix of the DHCP IPv6 Host Address object.",
	},
	"ipv6prefix_bits": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Prefix bits of the DHCP IPv6 Host Address object.",
	},
	"last_queried": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The time of the last DNS query in Epoch seconds format.",
		PlanModifiers: []planmodifier.Int64{
			plancontrol.UseStateForUnknownInt64(),
		},
	},
	"logic_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RecordHostIpv6addrLogicFilterRulesResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "This field contains the logic filters to be applied on the this host address. This list corresponds to the match rules that are written to the dhcpd configuration file.",
	},
	"mac": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The MAC address for this host address.",
		PlanModifiers: []planmodifier.String{
			plancontrol.UseStateForUnknownString(),
		},
	},
	"match_client": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The match_client value for this fixed address. Valid values are: \"DUID\": The host IP address is leased to the matching DUID. \"MAC_ADDRESS\": The host IP address is leased to the matching MAC address.",
		PlanModifiers: []planmodifier.String{
			plancontrol.UseStateForUnknownString(),
		},
	},
	"ms_ad_user_data": schema.SingleNestedAttribute{
		Attributes: RecordHostIpv6addrMsAdUserDataResourceSchemaAttributes,
		Optional:   true,
	},
	"network": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The network of the host address, in FQDN/CIDR format.",
		PlanModifiers: []planmodifier.String{
			plancontrol.UseStateForUnknownString(),
		},
	},
	"network_view": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the network view in which the host address resides.",
		PlanModifiers: []planmodifier.String{
			plancontrol.UseStateForUnknownString(),
		},
	},
	"options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RecordHostIpv6addrOptionsResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object.",
	},
	"preferred_lifetime": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Use this method to set or retrieve the preferred lifetime value of the DHCP IPv6 Host Address object.",
	},
	"reserved_interface": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The reference to the reserved interface to which the device belongs.",
	},
	"use_domain_name": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: domain_name",
	},
	"use_domain_name_servers": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: domain_name_servers",
	},
	"use_for_ea_inheritance": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Set this to True when using this host address for EA inheritance.",
	},
	"use_logic_filter_rules": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: logic_filter_rules",
	},
	"use_options": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: options",
	},
	"use_preferred_lifetime": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: preferred_lifetime",
	},
	"use_valid_lifetime": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: valid_lifetime",
	},
	"valid_lifetime": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Use this method to set or retrieve the valid lifetime value of the DHCP IPv6 Host Address object.",
	},
}

func ExpandRecordHostIpv6addr(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.RecordHostIpv6addr {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RecordHostIpv6addrModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *RecordHostIpv6addrModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.RecordHostIpv6addr {
	if m == nil {
		return nil
	}
	to := &dns.RecordHostIpv6addr{
		Ref:                  flex.ExpandStringPointer(m.Ref),
		AddressType:          flex.ExpandStringPointer(m.AddressType),
		ConfigureForDhcp:     flex.ExpandBoolPointer(m.ConfigureForDhcp),
		DiscoveredData:       ExpandRecordHostIpv6addrDiscoveredData(ctx, m.DiscoveredData, diags),
		DomainName:           flex.ExpandStringPointer(m.DomainName),
		DomainNameServers:    flex.ExpandFrameworkListString(ctx, m.DomainNameServers, diags),
		Duid:                 flex.ExpandStringPointer(m.Duid),
		Ipv6addr:             ExpandRecordHostIpv6addrIpv6addr(m.Ipv6addr),
		FuncCall:             ExpandFuncCall(ctx, m.FuncCall, diags),
		Ipv6prefix:           flex.ExpandStringPointer(m.Ipv6prefix),
		Ipv6prefixBits:       flex.ExpandInt64Pointer(m.Ipv6prefixBits),
		LogicFilterRules:     flex.ExpandFrameworkListNestedBlock(ctx, m.LogicFilterRules, diags, ExpandRecordHostIpv6addrLogicFilterRules),
		Mac:                  flex.ExpandStringPointer(m.Mac),
		MatchClient:          flex.ExpandStringPointer(m.MatchClient),
		MsAdUserData:         ExpandRecordHostIpv6addrMsAdUserData(ctx, m.MsAdUserData, diags),
		Options:              flex.ExpandFrameworkListNestedBlock(ctx, m.Options, diags, ExpandRecordHostIpv6addrOptions),
		PreferredLifetime:    flex.ExpandInt64Pointer(m.PreferredLifetime),
		ReservedInterface:    flex.ExpandStringPointerEmptyAsNil(m.ReservedInterface),
		UseDomainName:        flex.ExpandBoolPointer(m.UseDomainName),
		UseDomainNameServers: flex.ExpandBoolPointer(m.UseDomainNameServers),
		UseForEaInheritance:  flex.ExpandBoolPointer(m.UseForEaInheritance),
		UseLogicFilterRules:  flex.ExpandBoolPointer(m.UseLogicFilterRules),
		UseOptions:           flex.ExpandBoolPointer(m.UseOptions),
		UsePreferredLifetime: flex.ExpandBoolPointer(m.UsePreferredLifetime),
		UseValidLifetime:     flex.ExpandBoolPointer(m.UseValidLifetime),
		ValidLifetime:        flex.ExpandInt64Pointer(m.ValidLifetime),
	}
	return to
}

func FlattenRecordHostIpv6addr(ctx context.Context, from *dns.RecordHostIpv6addr, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordHostIpv6addrAttrTypes)
	}
	m := RecordHostIpv6addrModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RecordHostIpv6addrAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RecordHostIpv6addrModel) Flatten(ctx context.Context, from *dns.RecordHostIpv6addr, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RecordHostIpv6addrModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.AddressType = flex.FlattenStringPointer(from.AddressType)
	m.ConfigureForDhcp = types.BoolPointerValue(from.ConfigureForDhcp)
	m.DiscoverNowStatus = flex.FlattenStringPointer(from.DiscoverNowStatus)
	m.DiscoveredData = FlattenRecordHostIpv6addrDiscoveredData(ctx, from.DiscoveredData, diags)
	m.DomainName = flex.FlattenStringPointer(from.DomainName)
	m.DomainNameServers = flex.FlattenFrameworkListString(ctx, from.DomainNameServers, diags)
	m.Duid = flex.FlattenStringPointer(from.Duid)
	m.Host = flex.FlattenStringPointer(from.Host)
	m.Ipv6addr = FlattenRecordHostIpv6addrIpv6addr(from.Ipv6addr)
	m.FuncCall = FlattenFuncCall(ctx, from.FuncCall, diags)
	m.Ipv6prefix = flex.FlattenStringPointer(from.Ipv6prefix)
	m.Ipv6prefixBits = flex.FlattenInt64Pointer(from.Ipv6prefixBits)
	m.LastQueried = flex.FlattenInt64Pointer(from.LastQueried)
	m.LogicFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.LogicFilterRules, RecordHostIpv6addrLogicFilterRulesAttrTypes, diags, FlattenRecordHostIpv6addrLogicFilterRules)
	m.Mac = flex.FlattenStringPointer(from.Mac)
	m.MatchClient = flex.FlattenStringPointer(from.MatchClient)
	m.MsAdUserData = FlattenRecordHostIpv6addrMsAdUserData(ctx, from.MsAdUserData, diags)
	m.Network = flex.FlattenStringPointer(from.Network)
	m.NetworkView = flex.FlattenStringPointer(from.NetworkView)
	m.Options = flex.FlattenFrameworkListNestedBlock(ctx, from.Options, RecordHostIpv6addrOptionsAttrTypes, diags, FlattenRecordHostIpv6addrOptions)
	m.PreferredLifetime = flex.FlattenInt64Pointer(from.PreferredLifetime)
	m.ReservedInterface = flex.FlattenStringPointerNilAsNotEmpty(from.ReservedInterface)
	m.UseDomainName = types.BoolPointerValue(from.UseDomainName)
	m.UseDomainNameServers = types.BoolPointerValue(from.UseDomainNameServers)
	m.UseForEaInheritance = types.BoolPointerValue(from.UseForEaInheritance)
	m.UseLogicFilterRules = types.BoolPointerValue(from.UseLogicFilterRules)
	m.UseOptions = types.BoolPointerValue(from.UseOptions)
	m.UsePreferredLifetime = types.BoolPointerValue(from.UsePreferredLifetime)
	m.UseValidLifetime = types.BoolPointerValue(from.UseValidLifetime)
	m.ValidLifetime = flex.FlattenInt64Pointer(from.ValidLifetime)
}

func ExpandRecordHostIpv6addrIpv6addr(ipv6addr iptypes.IPv6Address) *dns.RecordHostIpv6addrIpv6addr {
	if ipv6addr.IsNull() {
		return &dns.RecordHostIpv6addrIpv6addr{}
	}
	var m dns.RecordHostIpv6addrIpv6addr
	m.String = flex.ExpandIPv6Address(ipv6addr)

	return &m
}

func FlattenRecordHostIpv6addrIpv6addr(from *dns.RecordHostIpv6addrIpv6addr) iptypes.IPv6Address {
	if from.String == nil {
		return iptypes.NewIPv6AddressNull()
	}
	m := flex.FlattenIPv6Address(from.String)
	return m
}
