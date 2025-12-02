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

type Ipv6sharednetworkModel struct {
	Ref                        types.String                             `tfsdk:"ref"`
	Comment                    types.String                             `tfsdk:"comment"`
	DdnsDomainname             internaltypes.CaseInsensitiveStringValue `tfsdk:"ddns_domainname"`
	DdnsGenerateHostname       types.Bool                               `tfsdk:"ddns_generate_hostname"`
	DdnsServerAlwaysUpdates    types.Bool                               `tfsdk:"ddns_server_always_updates"`
	DdnsTtl                    types.Int64                              `tfsdk:"ddns_ttl"`
	DdnsUseOption81            types.Bool                               `tfsdk:"ddns_use_option81"`
	Disable                    types.Bool                               `tfsdk:"disable"`
	DomainName                 internaltypes.CaseInsensitiveStringValue `tfsdk:"domain_name"`
	DomainNameServers          types.List                               `tfsdk:"domain_name_servers"`
	EnableDdns                 types.Bool                               `tfsdk:"enable_ddns"`
	ExtAttrs                   types.Map                                `tfsdk:"extattrs"`
	ExtAttrsAll                types.Map                                `tfsdk:"extattrs_all"`
	LogicFilterRules           types.List                               `tfsdk:"logic_filter_rules"`
	Name                       types.String                             `tfsdk:"name"`
	NetworkView                types.String                             `tfsdk:"network_view"`
	Networks                   internaltypes.UnorderedListValue         `tfsdk:"networks"`
	Options                    types.List                               `tfsdk:"options"`
	PreferredLifetime          types.Int64                              `tfsdk:"preferred_lifetime"`
	UpdateDnsOnLeaseRenewal    types.Bool                               `tfsdk:"update_dns_on_lease_renewal"`
	UseDdnsDomainname          types.Bool                               `tfsdk:"use_ddns_domainname"`
	UseDdnsGenerateHostname    types.Bool                               `tfsdk:"use_ddns_generate_hostname"`
	UseDdnsTtl                 types.Bool                               `tfsdk:"use_ddns_ttl"`
	UseDdnsUseOption81         types.Bool                               `tfsdk:"use_ddns_use_option81"`
	UseDomainName              types.Bool                               `tfsdk:"use_domain_name"`
	UseDomainNameServers       types.Bool                               `tfsdk:"use_domain_name_servers"`
	UseEnableDdns              types.Bool                               `tfsdk:"use_enable_ddns"`
	UseLogicFilterRules        types.Bool                               `tfsdk:"use_logic_filter_rules"`
	UseOptions                 types.Bool                               `tfsdk:"use_options"`
	UsePreferredLifetime       types.Bool                               `tfsdk:"use_preferred_lifetime"`
	UseUpdateDnsOnLeaseRenewal types.Bool                               `tfsdk:"use_update_dns_on_lease_renewal"`
	UseValidLifetime           types.Bool                               `tfsdk:"use_valid_lifetime"`
	ValidLifetime              types.Int64                              `tfsdk:"valid_lifetime"`
}

var Ipv6sharednetworkAttrTypes = map[string]attr.Type{
	"ref":                             types.StringType,
	"comment":                         types.StringType,
	"ddns_domainname":                 internaltypes.CaseInsensitiveString{},
	"ddns_generate_hostname":          types.BoolType,
	"ddns_server_always_updates":      types.BoolType,
	"ddns_ttl":                        types.Int64Type,
	"ddns_use_option81":               types.BoolType,
	"disable":                         types.BoolType,
	"domain_name":                     internaltypes.CaseInsensitiveString{},
	"domain_name_servers":             types.ListType{ElemType: types.StringType},
	"enable_ddns":                     types.BoolType,
	"extattrs":                        types.MapType{ElemType: types.StringType},
	"extattrs_all":                    types.MapType{ElemType: types.StringType},
	"logic_filter_rules":              types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6sharednetworkLogicFilterRulesAttrTypes}},
	"name":                            types.StringType,
	"network_view":                    types.StringType,
	"networks":                        internaltypes.UnorderedListOfStringType,
	"options":                         types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6sharednetworkOptionsAttrTypes}},
	"preferred_lifetime":              types.Int64Type,
	"update_dns_on_lease_renewal":     types.BoolType,
	"use_ddns_domainname":             types.BoolType,
	"use_ddns_generate_hostname":      types.BoolType,
	"use_ddns_ttl":                    types.BoolType,
	"use_ddns_use_option81":           types.BoolType,
	"use_domain_name":                 types.BoolType,
	"use_domain_name_servers":         types.BoolType,
	"use_enable_ddns":                 types.BoolType,
	"use_logic_filter_rules":          types.BoolType,
	"use_options":                     types.BoolType,
	"use_preferred_lifetime":          types.BoolType,
	"use_update_dns_on_lease_renewal": types.BoolType,
	"use_valid_lifetime":              types.BoolType,
	"valid_lifetime":                  types.Int64Type,
}

var Ipv6sharednetworkResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"comment": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 256),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Comment for the IPv6 shared network, maximum 256 characters.",
	},
	"ddns_domainname": schema.StringAttribute{
		CustomType: internaltypes.CaseInsensitiveString{},
		Computed:   true,
		Optional:   true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_ddns_domainname")),
			customvalidator.IsValidFQDN(),
		},
		MarkdownDescription: "The dynamic DNS domain name the appliance uses specifically for DDNS updates for this network.",
	},
	"ddns_generate_hostname": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_ddns_generate_hostname")),
		},
		MarkdownDescription: "If this field is set to True, the DHCP server generates a hostname and updates DNS with it when the DHCP client request does not contain a hostname.",
	},
	"ddns_server_always_updates": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(true),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("ddns_use_option81")),
		},
		MarkdownDescription: "This field controls whether only the DHCP server is allowed to update DNS, regardless of the DHCP clients requests. Note that changes for this field take effect only if ddns_use_option81 is True.",
	},
	"ddns_ttl": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(0),
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_ddns_ttl")),
		},
		MarkdownDescription: "The DNS update Time to Live (TTL) value of an IPv6 shared network object. The TTL is a 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached.",
	},
	"ddns_use_option81": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_ddns_use_option81")),
		},
		MarkdownDescription: "The support for DHCP Option 81 at the IPv6 shared network level.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether an IPv6 shared network is disabled or not. When this is set to False, the IPv6 shared network is enabled.",
	},
	"domain_name": schema.StringAttribute{
		CustomType: internaltypes.CaseInsensitiveString{},
		Computed:   true,
		Optional:   true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_domain_name")),
			customvalidator.IsValidFQDN(),
		},
		MarkdownDescription: "Use this method to set or retrieve the domain_name value of a DHCP IPv6 Shared Network object.",
	},
	"domain_name_servers": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
			listvalidator.AlsoRequires(path.MatchRoot("use_domain_name_servers")),
		},
		MarkdownDescription: "Use this method to set or retrieve the dynamic DNS updates flag of a DHCP IPv6 Shared Network object. The DHCP server can send DDNS updates to DNS servers in the same Grid and to external DNS servers. This setting overrides the member level settings.",
	},
	"enable_ddns": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_enable_ddns")),
		},
		MarkdownDescription: "The dynamic DNS updates flag of an IPv6 shared network object. If set to True, the DHCP server sends DDNS updates to DNS servers in the same Grid, and to external DNS servers.",
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
		MarkdownDescription: "Extensible attributes associated with the object , including default and internal attributes.",
		ElementType:         types.StringType,
	},
	"logic_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6sharednetworkLogicFilterRulesResourceSchemaAttributes,
		},
		Computed: true,
		Optional: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
			listvalidator.AlsoRequires(path.MatchRoot("use_logic_filter_rules")),
		},
		MarkdownDescription: "This field contains the logic filters to be applied on the this IPv6 shared network. This list corresponds to the match rules that are written to the DHCPv6 configuration file.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of the IPv6 Shared Network.",
	},
	"network_view": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Default:  stringdefault.StaticString("default"),
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of the network view in which this IPv6 shared network resides.",
	},
	"networks": schema.ListAttribute{
		ElementType: types.StringType,
		CustomType:  internaltypes.UnorderedListOfStringType,
		Required:    true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "A list of IPv6 networks belonging to the shared network Each individual list item must be specified as an object containing a 'ref' parameter to a network reference, for example:: [{ \"ref\": \"ipv6network/ZG5zdHdvcmskMTAuAvMTYvMA\", }] if the reference of the wanted network is not known, it is possible to specify search parameters for the network instead in the following way:: [{ \"ref\": { 'network': 'aabb::/64', } }] note that in this case the search must match exactly one network for the assignment to be successful.",
	},
	"options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6sharednetworkOptionsResourceSchemaAttributes,
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
		MarkdownDescription: "Use this method to set or retrieve the preferred lifetime value of a DHCP IPv6 Shared Network object.",
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
	"use_ddns_ttl": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: ddns_ttl",
	},
	"use_ddns_use_option81": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: ddns_use_option81",
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
	"use_enable_ddns": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: enable_ddns",
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
	"use_update_dns_on_lease_renewal": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: update_dns_on_lease_renewal",
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
		MarkdownDescription: "Use this method to set or retrieve the valid lifetime value of a DHCP IPv6 Shared Network object.",
	},
}

func (m *Ipv6sharednetworkModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *dhcp.Ipv6sharednetwork {
	if m == nil {
		return nil
	}
	to := &dhcp.Ipv6sharednetwork{
		Comment:                    flex.ExpandStringPointer(m.Comment),
		DdnsDomainname:             flex.ExpandStringPointer(m.DdnsDomainname.StringValue),
		DdnsGenerateHostname:       flex.ExpandBoolPointer(m.DdnsGenerateHostname),
		DdnsServerAlwaysUpdates:    flex.ExpandBoolPointer(m.DdnsServerAlwaysUpdates),
		DdnsTtl:                    flex.ExpandInt64Pointer(m.DdnsTtl),
		DdnsUseOption81:            flex.ExpandBoolPointer(m.DdnsUseOption81),
		Disable:                    flex.ExpandBoolPointer(m.Disable),
		DomainName:                 flex.ExpandStringPointer(m.DomainName.StringValue),
		DomainNameServers:          flex.ExpandFrameworkListString(ctx, m.DomainNameServers, diags),
		EnableDdns:                 flex.ExpandBoolPointer(m.EnableDdns),
		ExtAttrs:                   ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		LogicFilterRules:           flex.ExpandFrameworkListNestedBlock(ctx, m.LogicFilterRules, diags, ExpandIpv6sharednetworkLogicFilterRules),
		Name:                       flex.ExpandStringPointer(m.Name),
		Networks:                   ExpandNetworks(ctx, m.Networks, diags),
		Options:                    flex.ExpandFrameworkListNestedBlock(ctx, m.Options, diags, ExpandIpv6sharednetworkOptions),
		PreferredLifetime:          flex.ExpandInt64Pointer(m.PreferredLifetime),
		UpdateDnsOnLeaseRenewal:    flex.ExpandBoolPointer(m.UpdateDnsOnLeaseRenewal),
		UseDdnsDomainname:          flex.ExpandBoolPointer(m.UseDdnsDomainname),
		UseDdnsGenerateHostname:    flex.ExpandBoolPointer(m.UseDdnsGenerateHostname),
		UseDdnsTtl:                 flex.ExpandBoolPointer(m.UseDdnsTtl),
		UseDdnsUseOption81:         flex.ExpandBoolPointer(m.UseDdnsUseOption81),
		UseDomainName:              flex.ExpandBoolPointer(m.UseDomainName),
		UseDomainNameServers:       flex.ExpandBoolPointer(m.UseDomainNameServers),
		UseEnableDdns:              flex.ExpandBoolPointer(m.UseEnableDdns),
		UseLogicFilterRules:        flex.ExpandBoolPointer(m.UseLogicFilterRules),
		UseOptions:                 flex.ExpandBoolPointer(m.UseOptions),
		UsePreferredLifetime:       flex.ExpandBoolPointer(m.UsePreferredLifetime),
		UseUpdateDnsOnLeaseRenewal: flex.ExpandBoolPointer(m.UseUpdateDnsOnLeaseRenewal),
		UseValidLifetime:           flex.ExpandBoolPointer(m.UseValidLifetime),
		ValidLifetime:              flex.ExpandInt64Pointer(m.ValidLifetime),
	}
	if isCreate {
		to.NetworkView = flex.ExpandStringPointer(m.NetworkView)
	}
	return to
}

func FlattenIpv6sharednetwork(ctx context.Context, from *dhcp.Ipv6sharednetwork, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6sharednetworkAttrTypes)
	}
	m := Ipv6sharednetworkModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, Ipv6sharednetworkAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Ipv6sharednetworkModel) Flatten(ctx context.Context, from *dhcp.Ipv6sharednetwork, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Ipv6sharednetworkModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.DdnsDomainname.StringValue = flex.FlattenStringPointer(from.DdnsDomainname)
	m.DdnsGenerateHostname = types.BoolPointerValue(from.DdnsGenerateHostname)
	m.DdnsServerAlwaysUpdates = types.BoolPointerValue(from.DdnsServerAlwaysUpdates)
	m.DdnsTtl = flex.FlattenInt64Pointer(from.DdnsTtl)
	m.DdnsUseOption81 = types.BoolPointerValue(from.DdnsUseOption81)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.DomainName.StringValue = flex.FlattenStringPointer(from.DomainName)
	m.DomainNameServers = flex.FlattenFrameworkListString(ctx, from.DomainNameServers, diags)
	m.EnableDdns = types.BoolPointerValue(from.EnableDdns)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.LogicFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.LogicFilterRules, Ipv6sharednetworkLogicFilterRulesAttrTypes, diags, FlattenIpv6sharednetworkLogicFilterRules)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.NetworkView = flex.FlattenStringPointer(from.NetworkView)
	m.Networks = FlattenNetworks(ctx, from.Networks, diags)
	planOptions := m.Options
	m.Options = flex.FlattenFrameworkListNestedBlock(ctx, from.Options, Ipv6sharednetworkOptionsAttrTypes, diags, FlattenIpv6sharednetworkOptions)
	if !planOptions.IsUnknown() {
		reOrderedOptions, diags := utils.ReorderAndFilterDHCPOptions(ctx, planOptions, m.Options)
		if !diags.HasError() {
			m.Options = reOrderedOptions.(basetypes.ListValue)
		}
	}
	m.PreferredLifetime = flex.FlattenInt64Pointer(from.PreferredLifetime)
	m.UpdateDnsOnLeaseRenewal = types.BoolPointerValue(from.UpdateDnsOnLeaseRenewal)
	m.UseDdnsDomainname = types.BoolPointerValue(from.UseDdnsDomainname)
	m.UseDdnsGenerateHostname = types.BoolPointerValue(from.UseDdnsGenerateHostname)
	m.UseDdnsTtl = types.BoolPointerValue(from.UseDdnsTtl)
	m.UseDdnsUseOption81 = types.BoolPointerValue(from.UseDdnsUseOption81)
	m.UseDomainName = types.BoolPointerValue(from.UseDomainName)
	m.UseDomainNameServers = types.BoolPointerValue(from.UseDomainNameServers)
	m.UseEnableDdns = types.BoolPointerValue(from.UseEnableDdns)
	m.UseLogicFilterRules = types.BoolPointerValue(from.UseLogicFilterRules)
	m.UseOptions = types.BoolPointerValue(from.UseOptions)
	m.UsePreferredLifetime = types.BoolPointerValue(from.UsePreferredLifetime)
	m.UseUpdateDnsOnLeaseRenewal = types.BoolPointerValue(from.UseUpdateDnsOnLeaseRenewal)
	m.UseValidLifetime = types.BoolPointerValue(from.UseValidLifetime)
	m.ValidLifetime = flex.FlattenInt64Pointer(from.ValidLifetime)
}

func ExpandNetworks(ctx context.Context, networks internaltypes.UnorderedListValue, diags *diag.Diagnostics) []dhcp.Ipv6sharednetworkNetworks {
	if networks.IsNull() || networks.IsUnknown() {
		return nil
	}

	var networkRefs []string
	diags.Append(networks.ElementsAs(ctx, &networkRefs, false)...)
	if diags.HasError() {
		return nil
	}

	result := make([]dhcp.Ipv6sharednetworkNetworks, len(networkRefs))
	for i, ref := range networkRefs {
		result[i] = dhcp.Ipv6sharednetworkNetworks{
			Ref: &ref,
		}
	}
	return result
}

func FlattenNetworks(ctx context.Context, networks []dhcp.Ipv6sharednetworkNetworks, diags *diag.Diagnostics) internaltypes.UnorderedListValue {
	if networks == nil {
		return internaltypes.NewUnorderedListValueNull(types.StringType)
	}

	networkRefs := make([]string, len(networks))
	for i, network := range networks {
		if network.Ref != nil {
			networkRefs[i] = *network.Ref
		}
	}

	listValue, d := internaltypes.NewUnorderedListValueFrom(ctx, types.StringType, networkRefs)
	diags.Append(d...)
	return listValue
}
