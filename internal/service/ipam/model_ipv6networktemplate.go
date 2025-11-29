package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
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

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/import"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type Ipv6networktemplateModel struct {
	Ref                        types.String        `tfsdk:"ref"`
	AllowAnyNetmask            types.Bool          `tfsdk:"allow_any_netmask"`
	AutoCreateReversezone      types.Bool          `tfsdk:"auto_create_reversezone"`
	Cidr                       types.Int64         `tfsdk:"cidr"`
	CloudApiCompatible         types.Bool          `tfsdk:"cloud_api_compatible"`
	Comment                    types.String        `tfsdk:"comment"`
	DdnsDomainname             types.String        `tfsdk:"ddns_domainname"`
	DdnsEnableOptionFqdn       types.Bool          `tfsdk:"ddns_enable_option_fqdn"`
	DdnsGenerateHostname       types.Bool          `tfsdk:"ddns_generate_hostname"`
	DdnsServerAlwaysUpdates    types.Bool          `tfsdk:"ddns_server_always_updates"`
	DdnsTtl                    types.Int64         `tfsdk:"ddns_ttl"`
	DelegatedMember            types.Object        `tfsdk:"delegated_member"`
	DomainName                 types.String        `tfsdk:"domain_name"`
	DomainNameServers          types.List          `tfsdk:"domain_name_servers"`
	EnableDdns                 types.Bool          `tfsdk:"enable_ddns"`
	ExtAttrs                   types.Map           `tfsdk:"extattrs"`
	ExtAttrsAll                types.Map           `tfsdk:"extattrs_all"`
	FixedAddressTemplates      types.List          `tfsdk:"fixed_address_templates"`
	Ipv6prefix                 iptypes.IPv6Address `tfsdk:"ipv6prefix"`
	LogicFilterRules           types.List          `tfsdk:"logic_filter_rules"`
	Members                    types.List          `tfsdk:"members"`
	Name                       types.String        `tfsdk:"name"`
	Options                    types.List          `tfsdk:"options"`
	PreferredLifetime          types.Int64         `tfsdk:"preferred_lifetime"`
	RangeTemplates             types.List          `tfsdk:"range_templates"`
	RecycleLeases              types.Bool          `tfsdk:"recycle_leases"`
	Rir                        iptypes.IPv6Address `tfsdk:"rir"`
	RirOrganization            types.String        `tfsdk:"rir_organization"`
	RirRegistrationAction      types.String        `tfsdk:"rir_registration_action"`
	RirRegistrationStatus      types.String        `tfsdk:"rir_registration_status"`
	SendRirRequest             types.Bool          `tfsdk:"send_rir_request"`
	UpdateDnsOnLeaseRenewal    types.Bool          `tfsdk:"update_dns_on_lease_renewal"`
	UseDdnsDomainname          types.Bool          `tfsdk:"use_ddns_domainname"`
	UseDdnsEnableOptionFqdn    types.Bool          `tfsdk:"use_ddns_enable_option_fqdn"`
	UseDdnsGenerateHostname    types.Bool          `tfsdk:"use_ddns_generate_hostname"`
	UseDdnsTtl                 types.Bool          `tfsdk:"use_ddns_ttl"`
	UseDomainName              types.Bool          `tfsdk:"use_domain_name"`
	UseDomainNameServers       types.Bool          `tfsdk:"use_domain_name_servers"`
	UseEnableDdns              types.Bool          `tfsdk:"use_enable_ddns"`
	UseLogicFilterRules        types.Bool          `tfsdk:"use_logic_filter_rules"`
	UseOptions                 types.Bool          `tfsdk:"use_options"`
	UsePreferredLifetime       types.Bool          `tfsdk:"use_preferred_lifetime"`
	UseRecycleLeases           types.Bool          `tfsdk:"use_recycle_leases"`
	UseUpdateDnsOnLeaseRenewal types.Bool          `tfsdk:"use_update_dns_on_lease_renewal"`
	UseValidLifetime           types.Bool          `tfsdk:"use_valid_lifetime"`
	ValidLifetime              types.Int64         `tfsdk:"valid_lifetime"`
}

var Ipv6networktemplateAttrTypes = map[string]attr.Type{
	"ref":                             types.StringType,
	"allow_any_netmask":               types.BoolType,
	"auto_create_reversezone":         types.BoolType,
	"cidr":                            types.Int64Type,
	"cloud_api_compatible":            types.BoolType,
	"comment":                         types.StringType,
	"ddns_domainname":                 types.StringType,
	"ddns_enable_option_fqdn":         types.BoolType,
	"ddns_generate_hostname":          types.BoolType,
	"ddns_server_always_updates":      types.BoolType,
	"ddns_ttl":                        types.Int64Type,
	"delegated_member":                types.ObjectType{AttrTypes: Ipv6networktemplateDelegatedMemberAttrTypes},
	"domain_name":                     types.StringType,
	"domain_name_servers":             types.ListType{ElemType: types.StringType},
	"enable_ddns":                     types.BoolType,
	"extattrs":                        types.MapType{ElemType: types.StringType},
	"extattrs_all":                    types.MapType{ElemType: types.StringType},
	"fixed_address_templates":         types.ListType{ElemType: types.StringType},
	"ipv6prefix":                      iptypes.IPv6AddressType{},
	"logic_filter_rules":              types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6networktemplateLogicFilterRulesAttrTypes}},
	"members":                         types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6networktemplateMembersAttrTypes}},
	"name":                            types.StringType,
	"options":                         types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6networktemplateOptionsAttrTypes}},
	"preferred_lifetime":              types.Int64Type,
	"range_templates":                 types.ListType{ElemType: types.StringType},
	"recycle_leases":                  types.BoolType,
	"rir":                             iptypes.IPv6AddressType{},
	"rir_organization":                types.StringType,
	"rir_registration_action":         types.StringType,
	"rir_registration_status":         types.StringType,
	"send_rir_request":                types.BoolType,
	"update_dns_on_lease_renewal":     types.BoolType,
	"use_ddns_domainname":             types.BoolType,
	"use_ddns_enable_option_fqdn":     types.BoolType,
	"use_ddns_generate_hostname":      types.BoolType,
	"use_ddns_ttl":                    types.BoolType,
	"use_domain_name":                 types.BoolType,
	"use_domain_name_servers":         types.BoolType,
	"use_enable_ddns":                 types.BoolType,
	"use_logic_filter_rules":          types.BoolType,
	"use_options":                     types.BoolType,
	"use_preferred_lifetime":          types.BoolType,
	"use_recycle_leases":              types.BoolType,
	"use_update_dns_on_lease_renewal": types.BoolType,
	"use_valid_lifetime":              types.BoolType,
	"valid_lifetime":                  types.Int64Type,
}

var Ipv6networktemplateResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"allow_any_netmask": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This flag controls whether the template allows any netmask. You must specify a netmask when creating a network using this template. If you set this parameter to False, you must specify the \"cidr\" field for the network template object.",
	},
	"auto_create_reversezone": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This flag controls whether reverse zones are automatically created when the network is added.",
	},
	"cidr": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The CIDR of the network in CIDR format.",
	},
	"cloud_api_compatible": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This flag controls whether this template can be used to create network objects in a cloud-computing deployment.",
	},
	"comment": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 256),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Comment for the network; maximum 256 characters.",
	},
	"ddns_domainname": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_ddns_domainname")),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The dynamic DNS domain name the appliance uses specifically for DDNS updates for this network.",
	},
	"ddns_enable_option_fqdn": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_ddns_enable_option_fqdn")),
		},
		MarkdownDescription: "Use this method to set or retrieve the ddns_enable_option_fqdn flag of a DHCP IPv6 Network object. This method controls whether the FQDN option sent by the client is to be used, or if the server can automatically generate the FQDN. This setting overrides the upper-level settings.",
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
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "This field controls whether the DHCP server is allowed to update DNS, regardless of the DHCP client requests. Note that changes for this field take effect only if ddns_enable_option_fqdn is True.",
	},
	"ddns_ttl": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(0),
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_ddns_ttl")),
		},
		MarkdownDescription: "The DNS update Time to Live (TTL) value of a DHCP network object. The TTL is a 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached.",
	},
	"delegated_member": schema.SingleNestedAttribute{
		Attributes:          Ipv6networktemplateDelegatedMemberResourceSchemaAttributes,
		Computed:            true,
		Optional:            true,
		MarkdownDescription: "Reference the Cloud Platform Appliance to which authority of the object should be delegated when the object is created using the template.",
	},
	"domain_name": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_domain_name")),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Use this method to set or retrieve the domain_name value of a DHCP IPv6 Network object.",
	},
	"domain_name_servers": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
			listvalidator.AlsoRequires(path.MatchRoot("use_domain_name_servers")),
		},
		MarkdownDescription: "Use this method to set or retrieve the dynamic DNS updates flag of a DHCP IPv6 Network object. The DHCP server can send DDNS updates to DNS servers in the same Grid and to external DNS servers. This setting overrides the member level settings.",
	},
	"enable_ddns": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_enable_ddns")),
		},
		MarkdownDescription: "The dynamic DNS updates flag of a DHCP IPv6 network object. If set to True, the DHCP server sends DDNS updates to DNS servers in the same Grid, and to external DNS servers.",
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
	"fixed_address_templates": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The list of IPv6 fixed address templates assigned to this IPv6 network template object. When you create an IPv6 network based on an IPv6 network template object that contains IPv6 fixed address templates, the IPv6 fixed addresses are created based on the associated IPv6 fixed address templates.",
	},
	"ipv6prefix": schema.StringAttribute{
		CustomType:          iptypes.IPv6AddressType{},
		Computed:            true,
		Optional:            true,
		MarkdownDescription: "The IPv6 Address prefix of the DHCP IPv6 network.",
	},
	"logic_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6networktemplateLogicFilterRulesResourceSchemaAttributes,
		},
		Computed: true,
		Optional: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
			listvalidator.AlsoRequires(path.MatchRoot("use_logic_filter_rules")),
		},
		MarkdownDescription: "This field contains the logic filters to be applied on this IPv6 network template. This list corresponds to the match rules that are written to the DHCPv6 configuration file.",
	},
	"members": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6networktemplateMembersResourceSchemaAttributes,
		},
		Computed: true,
		Optional: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "A list of members that serve DHCP for the network. All members in the array must be of the same type. The struct type must be indicated in each element, by setting the \"_struct\" member to the struct type.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of this IPv6 network template.",
	},
	"options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6networktemplateOptionsResourceSchemaAttributes,
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
		MarkdownDescription: "Use this method to set or retrieve the preferred lifetime value of a DHCP IPv6 Network object.",
	},
	"range_templates": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The list of IPv6 address range templates assigned to this IPv6 network template object. When you create an IPv6 network based on an IPv6 network template object that contains IPv6 range templates, the IPv6 address ranges are created based on the associated IPv6 address range templates.",
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
	"rir": schema.StringAttribute{
		CustomType:          iptypes.IPv6AddressType{},
		Computed:            true,
		MarkdownDescription: "The registry (RIR) that allocated the IPv6 network address space.",
	},
	"rir_organization": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The RIR organization associated with the IPv6 network.",
	},
	"rir_registration_action": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Default:  stringdefault.StaticString("NONE"),
		Validators: []validator.String{
			stringvalidator.OneOf("CREATE", "NONE"),
		},
		MarkdownDescription: "The action for the RIR registration.",
	},
	"rir_registration_status": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Default:  stringdefault.StaticString("NOT_REGISTERED"),
		Validators: []validator.String{
			stringvalidator.OneOf("NOT_REGISTERED", "REGISTERED"),
		},
		MarkdownDescription: "The registration status of the IPv6 network in RIR.",
	},
	"send_rir_request": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether to send the RIR registration request.",
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
	"use_ddns_enable_option_fqdn": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: ddns_enable_option_fqdn",
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
	"use_recycle_leases": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: recycle_leases",
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
		MarkdownDescription: "Use this method to set or retrieve the valid lifetime value of a DHCP IPv6 Network object.",
	},
}

func (m *Ipv6networktemplateModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.Ipv6networktemplate {
	if m == nil {
		return nil
	}
	to := &ipam.Ipv6networktemplate{
		AllowAnyNetmask:            flex.ExpandBoolPointer(m.AllowAnyNetmask),
		AutoCreateReversezone:      flex.ExpandBoolPointer(m.AutoCreateReversezone),
		Cidr:                       flex.ExpandInt64Pointer(m.Cidr),
		CloudApiCompatible:         flex.ExpandBoolPointer(m.CloudApiCompatible),
		Comment:                    flex.ExpandStringPointer(m.Comment),
		DdnsDomainname:             flex.ExpandStringPointer(m.DdnsDomainname),
		DdnsEnableOptionFqdn:       flex.ExpandBoolPointer(m.DdnsEnableOptionFqdn),
		DdnsGenerateHostname:       flex.ExpandBoolPointer(m.DdnsGenerateHostname),
		DdnsServerAlwaysUpdates:    flex.ExpandBoolPointer(m.DdnsServerAlwaysUpdates),
		DdnsTtl:                    flex.ExpandInt64Pointer(m.DdnsTtl),
		DelegatedMember:            ExpandIpv6networktemplateDelegatedMember(ctx, m.DelegatedMember, diags),
		DomainName:                 flex.ExpandStringPointer(m.DomainName),
		DomainNameServers:          flex.ExpandFrameworkListString(ctx, m.DomainNameServers, diags),
		EnableDdns:                 flex.ExpandBoolPointer(m.EnableDdns),
		ExtAttrs:                   ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		FixedAddressTemplates:      flex.ExpandFrameworkListString(ctx, m.FixedAddressTemplates, diags),
		Ipv6prefix:                 flex.ExpandIPv6Address(m.Ipv6prefix),
		LogicFilterRules:           flex.ExpandFrameworkListNestedBlock(ctx, m.LogicFilterRules, diags, ExpandIpv6networktemplateLogicFilterRules),
		Members:                    flex.ExpandFrameworkListNestedBlock(ctx, m.Members, diags, ExpandIpv6networktemplateMembers),
		Name:                       flex.ExpandStringPointer(m.Name),
		Options:                    flex.ExpandFrameworkListNestedBlock(ctx, m.Options, diags, ExpandIpv6networktemplateOptions),
		PreferredLifetime:          flex.ExpandInt64Pointer(m.PreferredLifetime),
		RangeTemplates:             flex.ExpandFrameworkListString(ctx, m.RangeTemplates, diags),
		RecycleLeases:              flex.ExpandBoolPointer(m.RecycleLeases),
		RirOrganization:            flex.ExpandStringPointer(m.RirOrganization),
		RirRegistrationAction:      flex.ExpandStringPointer(m.RirRegistrationAction),
		RirRegistrationStatus:      flex.ExpandStringPointer(m.RirRegistrationStatus),
		SendRirRequest:             flex.ExpandBoolPointer(m.SendRirRequest),
		UpdateDnsOnLeaseRenewal:    flex.ExpandBoolPointer(m.UpdateDnsOnLeaseRenewal),
		UseDdnsDomainname:          flex.ExpandBoolPointer(m.UseDdnsDomainname),
		UseDdnsEnableOptionFqdn:    flex.ExpandBoolPointer(m.UseDdnsEnableOptionFqdn),
		UseDdnsGenerateHostname:    flex.ExpandBoolPointer(m.UseDdnsGenerateHostname),
		UseDdnsTtl:                 flex.ExpandBoolPointer(m.UseDdnsTtl),
		UseDomainName:              flex.ExpandBoolPointer(m.UseDomainName),
		UseDomainNameServers:       flex.ExpandBoolPointer(m.UseDomainNameServers),
		UseEnableDdns:              flex.ExpandBoolPointer(m.UseEnableDdns),
		UseLogicFilterRules:        flex.ExpandBoolPointer(m.UseLogicFilterRules),
		UseOptions:                 flex.ExpandBoolPointer(m.UseOptions),
		UsePreferredLifetime:       flex.ExpandBoolPointer(m.UsePreferredLifetime),
		UseRecycleLeases:           flex.ExpandBoolPointer(m.UseRecycleLeases),
		UseUpdateDnsOnLeaseRenewal: flex.ExpandBoolPointer(m.UseUpdateDnsOnLeaseRenewal),
		UseValidLifetime:           flex.ExpandBoolPointer(m.UseValidLifetime),
		ValidLifetime:              flex.ExpandInt64Pointer(m.ValidLifetime),
	}
	return to
}

func FlattenIpv6networktemplate(ctx context.Context, from *ipam.Ipv6networktemplate, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6networktemplateAttrTypes)
	}
	m := Ipv6networktemplateModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, Ipv6networktemplateAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Ipv6networktemplateModel) Flatten(ctx context.Context, from *ipam.Ipv6networktemplate, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Ipv6networktemplateModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.AllowAnyNetmask = types.BoolPointerValue(from.AllowAnyNetmask)
	m.AutoCreateReversezone = types.BoolPointerValue(from.AutoCreateReversezone)
	m.Cidr = flex.FlattenInt64Pointer(from.Cidr)
	m.CloudApiCompatible = types.BoolPointerValue(from.CloudApiCompatible)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.DdnsDomainname = flex.FlattenStringPointer(from.DdnsDomainname)
	m.DdnsEnableOptionFqdn = types.BoolPointerValue(from.DdnsEnableOptionFqdn)
	m.DdnsGenerateHostname = types.BoolPointerValue(from.DdnsGenerateHostname)
	m.DdnsServerAlwaysUpdates = types.BoolPointerValue(from.DdnsServerAlwaysUpdates)
	m.DdnsTtl = flex.FlattenInt64Pointer(from.DdnsTtl)
	m.DelegatedMember = FlattenIpv6networktemplateDelegatedMember(ctx, from.DelegatedMember, diags)
	m.DomainName = flex.FlattenStringPointer(from.DomainName)
	m.DomainNameServers = flex.FlattenFrameworkListString(ctx, from.DomainNameServers, diags)
	m.EnableDdns = types.BoolPointerValue(from.EnableDdns)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.FixedAddressTemplates = flex.FlattenFrameworkListString(ctx, from.FixedAddressTemplates, diags)
	m.Ipv6prefix = flex.FlattenIPv6Address(from.Ipv6prefix)
	m.LogicFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.LogicFilterRules, Ipv6networktemplateLogicFilterRulesAttrTypes, diags, FlattenIpv6networktemplateLogicFilterRules)
	m.Members = flex.FlattenFrameworkListNestedBlock(ctx, from.Members, Ipv6networktemplateMembersAttrTypes, diags, FlattenIpv6networktemplateMembers)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Options = flex.FlattenFrameworkListNestedBlock(ctx, from.Options, Ipv6networktemplateOptionsAttrTypes, diags, FlattenIpv6networktemplateOptions)
	m.PreferredLifetime = flex.FlattenInt64Pointer(from.PreferredLifetime)
	m.RangeTemplates = flex.FlattenFrameworkListString(ctx, from.RangeTemplates, diags)
	m.RecycleLeases = types.BoolPointerValue(from.RecycleLeases)
	m.Rir = flex.FlattenIPv6Address(from.Rir)
	m.RirOrganization = flex.FlattenStringPointer(from.RirOrganization)
	m.RirRegistrationAction = flex.FlattenStringPointer(from.RirRegistrationAction)
	m.RirRegistrationStatus = flex.FlattenStringPointer(from.RirRegistrationStatus)
	m.SendRirRequest = types.BoolPointerValue(from.SendRirRequest)
	m.UpdateDnsOnLeaseRenewal = types.BoolPointerValue(from.UpdateDnsOnLeaseRenewal)
	m.UseDdnsDomainname = types.BoolPointerValue(from.UseDdnsDomainname)
	m.UseDdnsEnableOptionFqdn = types.BoolPointerValue(from.UseDdnsEnableOptionFqdn)
	m.UseDdnsGenerateHostname = types.BoolPointerValue(from.UseDdnsGenerateHostname)
	m.UseDdnsTtl = types.BoolPointerValue(from.UseDdnsTtl)
	m.UseDomainName = types.BoolPointerValue(from.UseDomainName)
	m.UseDomainNameServers = types.BoolPointerValue(from.UseDomainNameServers)
	m.UseEnableDdns = types.BoolPointerValue(from.UseEnableDdns)
	m.UseLogicFilterRules = types.BoolPointerValue(from.UseLogicFilterRules)
	m.UseOptions = types.BoolPointerValue(from.UseOptions)
	m.UsePreferredLifetime = types.BoolPointerValue(from.UsePreferredLifetime)
	m.UseRecycleLeases = types.BoolPointerValue(from.UseRecycleLeases)
	m.UseUpdateDnsOnLeaseRenewal = types.BoolPointerValue(from.UseUpdateDnsOnLeaseRenewal)
	m.UseValidLifetime = types.BoolPointerValue(from.UseValidLifetime)
	m.ValidLifetime = flex.FlattenInt64Pointer(from.ValidLifetime)
}
