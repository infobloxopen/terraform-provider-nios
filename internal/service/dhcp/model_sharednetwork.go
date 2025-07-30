package dhcp

import (
	"context"
	"regexp"

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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	internaltypes "github.com/infobloxopen/terraform-provider-nios/internal/types"
)

// TODO: networks fields to accept list of IPs (current implementation accepts list of networks' references)
// TODO: ignore_id, ignore_client_identifier need to be checked as `ignore_id` instead of `ignore_client_identifier` in version WAPI 1.8 or higher.

type SharednetworkModel struct {
	Ref                            types.String                     `tfsdk:"ref"`
	Authority                      types.Bool                       `tfsdk:"authority"`
	Bootfile                       types.String                     `tfsdk:"bootfile"`
	Bootserver                     types.String                     `tfsdk:"bootserver"`
	Comment                        types.String                     `tfsdk:"comment"`
	DdnsGenerateHostname           types.Bool                       `tfsdk:"ddns_generate_hostname"`
	DdnsServerAlwaysUpdates        types.Bool                       `tfsdk:"ddns_server_always_updates"`
	DdnsTtl                        types.Int64                      `tfsdk:"ddns_ttl"`
	DdnsUpdateFixedAddresses       types.Bool                       `tfsdk:"ddns_update_fixed_addresses"`
	DdnsUseOption81                types.Bool                       `tfsdk:"ddns_use_option81"`
	DenyBootp                      types.Bool                       `tfsdk:"deny_bootp"`
	DhcpUtilization                types.Int64                      `tfsdk:"dhcp_utilization"`
	DhcpUtilizationStatus          types.String                     `tfsdk:"dhcp_utilization_status"`
	Disable                        types.Bool                       `tfsdk:"disable"`
	DynamicHosts                   types.Int64                      `tfsdk:"dynamic_hosts"`
	EnableDdns                     types.Bool                       `tfsdk:"enable_ddns"`
	EnablePxeLeaseTime             types.Bool                       `tfsdk:"enable_pxe_lease_time"`
	ExtAttrs                       types.Map                        `tfsdk:"extattrs"`
	ExtAttrsAll                    types.Map                        `tfsdk:"extattrs_all"`
	IgnoreClientIdentifier         types.Bool                       `tfsdk:"ignore_client_identifier"`
	IgnoreDhcpOptionListRequest    types.Bool                       `tfsdk:"ignore_dhcp_option_list_request"`
	IgnoreId                       types.String                     `tfsdk:"ignore_id"`
	IgnoreMacAddresses             internaltypes.UnorderedListValue `tfsdk:"ignore_mac_addresses"`
	LeaseScavengeTime              types.Int64                      `tfsdk:"lease_scavenge_time"`
	LogicFilterRules               types.List                       `tfsdk:"logic_filter_rules"`
	MsAdUserData                   types.Object                     `tfsdk:"ms_ad_user_data"`
	Name                           types.String                     `tfsdk:"name"`
	NetworkView                    types.String                     `tfsdk:"network_view"`
	Networks                       internaltypes.UnorderedListValue `tfsdk:"networks"`
	Nextserver                     types.String                     `tfsdk:"nextserver"`
	Options                        internaltypes.UnorderedListValue `tfsdk:"options"`
	PxeLeaseTime                   types.Int64                      `tfsdk:"pxe_lease_time"`
	StaticHosts                    types.Int64                      `tfsdk:"static_hosts"`
	TotalHosts                     types.Int64                      `tfsdk:"total_hosts"`
	UpdateDnsOnLeaseRenewal        types.Bool                       `tfsdk:"update_dns_on_lease_renewal"`
	UseAuthority                   types.Bool                       `tfsdk:"use_authority"`
	UseBootfile                    types.Bool                       `tfsdk:"use_bootfile"`
	UseBootserver                  types.Bool                       `tfsdk:"use_bootserver"`
	UseDdnsGenerateHostname        types.Bool                       `tfsdk:"use_ddns_generate_hostname"`
	UseDdnsTtl                     types.Bool                       `tfsdk:"use_ddns_ttl"`
	UseDdnsUpdateFixedAddresses    types.Bool                       `tfsdk:"use_ddns_update_fixed_addresses"`
	UseDdnsUseOption81             types.Bool                       `tfsdk:"use_ddns_use_option81"`
	UseDenyBootp                   types.Bool                       `tfsdk:"use_deny_bootp"`
	UseEnableDdns                  types.Bool                       `tfsdk:"use_enable_ddns"`
	UseIgnoreClientIdentifier      types.Bool                       `tfsdk:"use_ignore_client_identifier"`
	UseIgnoreDhcpOptionListRequest types.Bool                       `tfsdk:"use_ignore_dhcp_option_list_request"`
	UseIgnoreId                    types.Bool                       `tfsdk:"use_ignore_id"`
	UseLeaseScavengeTime           types.Bool                       `tfsdk:"use_lease_scavenge_time"`
	UseLogicFilterRules            types.Bool                       `tfsdk:"use_logic_filter_rules"`
	UseNextserver                  types.Bool                       `tfsdk:"use_nextserver"`
	UseOptions                     types.Bool                       `tfsdk:"use_options"`
	UsePxeLeaseTime                types.Bool                       `tfsdk:"use_pxe_lease_time"`
	UseUpdateDnsOnLeaseRenewal     types.Bool                       `tfsdk:"use_update_dns_on_lease_renewal"`
}

var SharednetworkAttrTypes = map[string]attr.Type{
	"ref":                                 types.StringType,
	"authority":                           types.BoolType,
	"bootfile":                            types.StringType,
	"bootserver":                          types.StringType,
	"comment":                             types.StringType,
	"ddns_generate_hostname":              types.BoolType,
	"ddns_server_always_updates":          types.BoolType,
	"ddns_ttl":                            types.Int64Type,
	"ddns_update_fixed_addresses":         types.BoolType,
	"ddns_use_option81":                   types.BoolType,
	"deny_bootp":                          types.BoolType,
	"dhcp_utilization":                    types.Int64Type,
	"dhcp_utilization_status":             types.StringType,
	"disable":                             types.BoolType,
	"dynamic_hosts":                       types.Int64Type,
	"enable_ddns":                         types.BoolType,
	"enable_pxe_lease_time":               types.BoolType,
	"extattrs":                            types.MapType{ElemType: types.StringType},
	"extattrs_all":                        types.MapType{ElemType: types.StringType},
	"ignore_client_identifier":            types.BoolType,
	"ignore_dhcp_option_list_request":     types.BoolType,
	"ignore_id":                           types.StringType,
	"ignore_mac_addresses":                internaltypes.UnorderedListOfStringType,
	"lease_scavenge_time":                 types.Int64Type,
	"logic_filter_rules":                  types.ListType{ElemType: types.ObjectType{AttrTypes: SharednetworkLogicFilterRulesAttrTypes}},
	"ms_ad_user_data":                     types.ObjectType{AttrTypes: SharednetworkMsAdUserDataAttrTypes},
	"name":                                types.StringType,
	"network_view":                        types.StringType,
	"networks":                            internaltypes.UnorderedList{ListType: basetypes.ListType{ElemType: basetypes.ObjectType{AttrTypes: SharednetworkNetworksAttrTypes}}},
	"nextserver":                          types.StringType,
	"options":                             internaltypes.UnorderedList{ListType: basetypes.ListType{ElemType: basetypes.ObjectType{AttrTypes: SharednetworkOptionsAttrTypes}}},
	"pxe_lease_time":                      types.Int64Type,
	"static_hosts":                        types.Int64Type,
	"total_hosts":                         types.Int64Type,
	"update_dns_on_lease_renewal":         types.BoolType,
	"use_authority":                       types.BoolType,
	"use_bootfile":                        types.BoolType,
	"use_bootserver":                      types.BoolType,
	"use_ddns_generate_hostname":          types.BoolType,
	"use_ddns_ttl":                        types.BoolType,
	"use_ddns_update_fixed_addresses":     types.BoolType,
	"use_ddns_use_option81":               types.BoolType,
	"use_deny_bootp":                      types.BoolType,
	"use_enable_ddns":                     types.BoolType,
	"use_ignore_client_identifier":        types.BoolType,
	"use_ignore_dhcp_option_list_request": types.BoolType,
	"use_ignore_id":                       types.BoolType,
	"use_lease_scavenge_time":             types.BoolType,
	"use_logic_filter_rules":              types.BoolType,
	"use_nextserver":                      types.BoolType,
	"use_options":                         types.BoolType,
	"use_pxe_lease_time":                  types.BoolType,
	"use_update_dns_on_lease_renewal":     types.BoolType,
}

var SharednetworkResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"authority": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_authority")),
		},
		MarkdownDescription: "Authority for the shared network.",
	},
	"bootfile": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_bootfile")),
		},
		MarkdownDescription: "The bootfile name for the shared network. You can configure the DHCP server to support clients that use the boot file name option in their DHCPREQUEST messages.",
	},
	"bootserver": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_bootserver")),
		},
		MarkdownDescription: "The bootserver address for the shared network. You can specify the name and/or IP address of the boot server that the host needs to boot. The boot server IPv4 Address or name in FQDN format.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Comment should not have leading or trailing whitespace",
			),
		},
		MarkdownDescription: "Comment for the shared network, maximum 256 characters.",
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
		MarkdownDescription: "The DNS update Time to Live (TTL) value of a shared network object. The TTL is a 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached.",
	},
	"ddns_update_fixed_addresses": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_ddns_update_fixed_addresses")),
		},
		MarkdownDescription: "By default, the DHCP server does not update DNS when it allocates a fixed address to a client. You can configure the DHCP server to update the A and PTR records of a client with a fixed address. When this feature is enabled and the DHCP server adds A and PTR records for a fixed address, the DHCP server never discards the records.",
	},
	"ddns_use_option81": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_ddns_use_option81")),
		},
		MarkdownDescription: "The support for DHCP Option 81 at the shared network level.",
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
	"dhcp_utilization": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The percentage of the total DHCP utilization of the networks belonging to the shared network multiplied by 1000. This is the percentage of the total number of available IP addresses from all the networks belonging to the shared network versus the total number of all IP addresses in all of the networks in the shared network.",
	},
	"dhcp_utilization_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "A string describing the utilization level of the shared network.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether a shared network is disabled or not. When this is set to False, the shared network is enabled.",
	},
	"dynamic_hosts": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The total number of DHCP leases issued for the shared network.",
	},
	"enable_ddns": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_enable_ddns")),
		},
		MarkdownDescription: "The dynamic DNS updates flag of a shared network object. If set to True, the DHCP server sends DDNS updates to DNS servers in the same Grid, and to external DNS servers.",
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
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object , including default attributes.",
		ElementType:         types.StringType,
	},
	"ignore_client_identifier": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_ignore_client_identifier")),
		},
		MarkdownDescription: "If set to true, the client identifier will be ignored.",
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
	"ignore_id": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("CLIENT", "MACADDR", "NONE"),
			stringvalidator.AlsoRequires(path.MatchRoot("use_ignore_id")),
		},
		MarkdownDescription: "Indicates whether the appliance will ignore DHCP client IDs or MAC addresses. Valid values are \"NONE\", \"CLIENT\", or \"MACADDR\". The default is \"NONE\".",
	},
	"ignore_mac_addresses": schema.ListAttribute{
		CustomType:          internaltypes.UnorderedListOfStringType,
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "A list of MAC addresses the appliance will ignore.",
	},
	"lease_scavenge_time": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(-1),
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_lease_scavenge_time")),
		},
		MarkdownDescription: "An integer that specifies the period of time (in seconds) that frees and backs up leases remained in the database before they are automatically deleted. To disable lease scavenging, set the parameter to -1. The minimum positive value must be greater than 86400 seconds (1 day).",
	},
	"logic_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: SharednetworkLogicFilterRulesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_logic_filter_rules")),
		},
		MarkdownDescription: "This field contains the logic filters to be applied on the this shared network. This list corresponds to the match rules that are written to the dhcpd configuration file.",
	},
	"ms_ad_user_data": schema.SingleNestedAttribute{
		Attributes:          SharednetworkMsAdUserDataResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "The Microsoft Active Directory user related information.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"prefix should not have leading or trailing whitespace",
			),
		},
		MarkdownDescription: "The name of the IPv6 Shared Network.",
	},
	"network_view": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("default"),
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Comment should not have leading or trailing whitespace",
			),
		},
		MarkdownDescription: "The name of the network view in which this shared network resides.",
	},
	"networks": schema.ListNestedAttribute{
		CustomType: internaltypes.UnorderedList{ListType: basetypes.ListType{ElemType: basetypes.ObjectType{AttrTypes: SharednetworkNetworksAttrTypes}}},
		NestedObject: schema.NestedAttributeObject{
			Attributes: SharednetworkNetworksResourceSchemaAttributes,
		},
		Required:            true,
		MarkdownDescription: "A list of networks belonging to the shared network. Each individual list item must be specified as an object containing a 'ref' parameter to a network reference, for example:: [{ \"ref\": \"network/ZG5zLm5ldHdvcmskMTAuMwLvMTYvMA\" }] if the reference of the wanted network is not known, it is possible to specify search parameters for the network instead in the following way:: [{ \"ref\": { 'network': '10.0.0.0/8' } }] note that in this case the search must match exactly one network for the assignment to be successful.",
	},
	"nextserver": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_nextserver")),
		},
		MarkdownDescription: "The name in FQDN and/or IPv4 Address of the next server that the host needs to boot.",
	},
	"options": schema.ListNestedAttribute{
		CustomType: internaltypes.UnorderedList{ListType: basetypes.ListType{ElemType: basetypes.ObjectType{AttrTypes: SharednetworkOptionsAttrTypes}}},
		NestedObject: schema.NestedAttributeObject{
			Attributes: SharednetworkOptionsResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_options")),
		},
		MarkdownDescription: "An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object.",
	},
	"pxe_lease_time": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_pxe_lease_time")),
		},
		MarkdownDescription: "The PXE lease time value of a shared network object. Some hosts use PXE (Preboot Execution Environment) to boot remotely from a server. To better manage your IP resources, set a different lease time for PXE boot requests. You can configure the DHCP server to allocate an IP address with a shorter lease time to hosts that send PXE boot requests, so IP addresses are not leased longer than necessary. A 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached.",
	},
	"static_hosts": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of static DHCP addresses configured in the shared network.",
	},
	"total_hosts": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The total number of DHCP addresses configured in the shared network.",
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
	"use_authority": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: authority",
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
	"use_ddns_update_fixed_addresses": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: ddns_update_fixed_addresses",
	},
	"use_ddns_use_option81": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: ddns_use_option81",
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
	"use_ignore_client_identifier": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Use flag for: ignore_client_identifier",
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
		MarkdownDescription: "Use flag for: ignore_id",
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
	"use_update_dns_on_lease_renewal": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: update_dns_on_lease_renewal",
	},
}

func (m *SharednetworkModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *dhcp.Sharednetwork {
	if m == nil {
		return nil
	}
	to := &dhcp.Sharednetwork{
		Authority:                      flex.ExpandBoolPointer(m.Authority),
		Bootfile:                       flex.ExpandStringPointer(m.Bootfile),
		Bootserver:                     flex.ExpandStringPointer(m.Bootserver),
		Comment:                        flex.ExpandStringPointer(m.Comment),
		DdnsGenerateHostname:           flex.ExpandBoolPointer(m.DdnsGenerateHostname),
		DdnsServerAlwaysUpdates:        flex.ExpandBoolPointer(m.DdnsServerAlwaysUpdates),
		DdnsTtl:                        flex.ExpandInt64Pointer(m.DdnsTtl),
		DdnsUpdateFixedAddresses:       flex.ExpandBoolPointer(m.DdnsUpdateFixedAddresses),
		DdnsUseOption81:                flex.ExpandBoolPointer(m.DdnsUseOption81),
		DenyBootp:                      flex.ExpandBoolPointer(m.DenyBootp),
		Disable:                        flex.ExpandBoolPointer(m.Disable),
		EnableDdns:                     flex.ExpandBoolPointer(m.EnableDdns),
		EnablePxeLeaseTime:             flex.ExpandBoolPointer(m.EnablePxeLeaseTime),
		ExtAttrs:                       ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		IgnoreClientIdentifier:         flex.ExpandBoolPointer(m.IgnoreClientIdentifier),
		IgnoreDhcpOptionListRequest:    flex.ExpandBoolPointer(m.IgnoreDhcpOptionListRequest),
		IgnoreId:                       flex.ExpandStringPointer(m.IgnoreId),
		IgnoreMacAddresses:             flex.ExpandFrameworkListString(ctx, m.IgnoreMacAddresses, diags),
		LeaseScavengeTime:              flex.ExpandInt64Pointer(m.LeaseScavengeTime),
		LogicFilterRules:               flex.ExpandFrameworkListNestedBlock(ctx, m.LogicFilterRules, diags, ExpandSharednetworkLogicFilterRules),
		MsAdUserData:                   ExpandSharednetworkMsAdUserData(ctx, m.MsAdUserData, diags),
		Name:                           flex.ExpandStringPointer(m.Name),
		Networks:                       flex.ExpandFrameworkListNestedBlock(ctx, m.Networks, diags, ExpandSharednetworkNetworks),
		Nextserver:                     flex.ExpandStringPointer(m.Nextserver),
		Options:                        flex.ExpandFrameworkListNestedBlock(ctx, m.Options, diags, ExpandSharednetworkOptions),
		PxeLeaseTime:                   flex.ExpandInt64Pointer(m.PxeLeaseTime),
		UpdateDnsOnLeaseRenewal:        flex.ExpandBoolPointer(m.UpdateDnsOnLeaseRenewal),
		UseAuthority:                   flex.ExpandBoolPointer(m.UseAuthority),
		UseBootfile:                    flex.ExpandBoolPointer(m.UseBootfile),
		UseBootserver:                  flex.ExpandBoolPointer(m.UseBootserver),
		UseDdnsGenerateHostname:        flex.ExpandBoolPointer(m.UseDdnsGenerateHostname),
		UseDdnsTtl:                     flex.ExpandBoolPointer(m.UseDdnsTtl),
		UseDdnsUpdateFixedAddresses:    flex.ExpandBoolPointer(m.UseDdnsUpdateFixedAddresses),
		UseDdnsUseOption81:             flex.ExpandBoolPointer(m.UseDdnsUseOption81),
		UseDenyBootp:                   flex.ExpandBoolPointer(m.UseDenyBootp),
		UseEnableDdns:                  flex.ExpandBoolPointer(m.UseEnableDdns),
		UseIgnoreClientIdentifier:      flex.ExpandBoolPointer(m.UseIgnoreClientIdentifier),
		UseIgnoreDhcpOptionListRequest: flex.ExpandBoolPointer(m.UseIgnoreDhcpOptionListRequest),
		UseIgnoreId:                    flex.ExpandBoolPointer(m.UseIgnoreId),
		UseLeaseScavengeTime:           flex.ExpandBoolPointer(m.UseLeaseScavengeTime),
		UseLogicFilterRules:            flex.ExpandBoolPointer(m.UseLogicFilterRules),
		UseNextserver:                  flex.ExpandBoolPointer(m.UseNextserver),
		UseOptions:                     flex.ExpandBoolPointer(m.UseOptions),
		UsePxeLeaseTime:                flex.ExpandBoolPointer(m.UsePxeLeaseTime),
		UseUpdateDnsOnLeaseRenewal:     flex.ExpandBoolPointer(m.UseUpdateDnsOnLeaseRenewal),
	}
	if isCreate {
		to.NetworkView = flex.ExpandStringPointer(m.NetworkView)
	}
	return to
}

func FlattenSharednetwork(ctx context.Context, from *dhcp.Sharednetwork, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SharednetworkAttrTypes)
	}
	m := SharednetworkModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, SharednetworkAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *SharednetworkModel) Flatten(ctx context.Context, from *dhcp.Sharednetwork, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = SharednetworkModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Authority = types.BoolPointerValue(from.Authority)
	m.Bootfile = flex.FlattenStringPointer(from.Bootfile)
	m.Bootserver = flex.FlattenStringPointer(from.Bootserver)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.DdnsGenerateHostname = types.BoolPointerValue(from.DdnsGenerateHostname)
	m.DdnsServerAlwaysUpdates = types.BoolPointerValue(from.DdnsServerAlwaysUpdates)
	m.DdnsTtl = flex.FlattenInt64Pointer(from.DdnsTtl)
	m.DdnsUpdateFixedAddresses = types.BoolPointerValue(from.DdnsUpdateFixedAddresses)
	m.DdnsUseOption81 = types.BoolPointerValue(from.DdnsUseOption81)
	m.DenyBootp = types.BoolPointerValue(from.DenyBootp)
	m.DhcpUtilization = flex.FlattenInt64Pointer(from.DhcpUtilization)
	m.DhcpUtilizationStatus = flex.FlattenStringPointer(from.DhcpUtilizationStatus)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.DynamicHosts = flex.FlattenInt64Pointer(from.DynamicHosts)
	m.EnableDdns = types.BoolPointerValue(from.EnableDdns)
	m.EnablePxeLeaseTime = types.BoolPointerValue(from.EnablePxeLeaseTime)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.IgnoreClientIdentifier = types.BoolPointerValue(from.IgnoreClientIdentifier)
	m.IgnoreDhcpOptionListRequest = types.BoolPointerValue(from.IgnoreDhcpOptionListRequest)
	m.IgnoreId = flex.FlattenStringPointer(from.IgnoreId)
	m.IgnoreMacAddresses = flex.FlattenFrameworkUnorderedList(ctx, types.StringType, from.IgnoreMacAddresses, diags)
	m.LeaseScavengeTime = flex.FlattenInt64Pointer(from.LeaseScavengeTime)
	m.LogicFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.LogicFilterRules, SharednetworkLogicFilterRulesAttrTypes, diags, FlattenSharednetworkLogicFilterRules)
	m.MsAdUserData = FlattenSharednetworkMsAdUserData(ctx, from.MsAdUserData, diags)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.NetworkView = flex.FlattenStringPointer(from.NetworkView)
	m.Networks = flex.FlattenFrameworkUnorderedListNestedBlock(ctx, from.Networks, SharednetworkNetworksAttrTypes, diags, FlattenSharednetworkNetworks)
	m.Nextserver = flex.FlattenStringPointer(from.Nextserver)
	m.Options = flex.FilterDHCPOptions(
		ctx,
		diags,
		from.Options,
		m.Options,
		RangetemplateOptionsAttrTypes,
		func(ctx context.Context, opt *dhcp.SharednetworkOptions, d *diag.Diagnostics) types.Object {
			return FlattenSharednetworkOptions(ctx, opt, d)
		},
		func(ctx context.Context, obj types.Object, d *diag.Diagnostics) *dhcp.SharednetworkOptions {
			return ExpandSharednetworkOptions(ctx, obj, d)
		},
	)
	m.PxeLeaseTime = flex.FlattenInt64Pointer(from.PxeLeaseTime)
	m.StaticHosts = flex.FlattenInt64Pointer(from.StaticHosts)
	m.TotalHosts = flex.FlattenInt64Pointer(from.TotalHosts)
	m.UpdateDnsOnLeaseRenewal = types.BoolPointerValue(from.UpdateDnsOnLeaseRenewal)
	m.UseAuthority = types.BoolPointerValue(from.UseAuthority)
	m.UseBootfile = types.BoolPointerValue(from.UseBootfile)
	m.UseBootserver = types.BoolPointerValue(from.UseBootserver)
	m.UseDdnsGenerateHostname = types.BoolPointerValue(from.UseDdnsGenerateHostname)
	m.UseDdnsTtl = types.BoolPointerValue(from.UseDdnsTtl)
	m.UseDdnsUpdateFixedAddresses = types.BoolPointerValue(from.UseDdnsUpdateFixedAddresses)
	m.UseDdnsUseOption81 = types.BoolPointerValue(from.UseDdnsUseOption81)
	m.UseDenyBootp = types.BoolPointerValue(from.UseDenyBootp)
	m.UseEnableDdns = types.BoolPointerValue(from.UseEnableDdns)
	m.UseIgnoreClientIdentifier = types.BoolPointerValue(from.UseIgnoreClientIdentifier)
	m.UseIgnoreDhcpOptionListRequest = types.BoolPointerValue(from.UseIgnoreDhcpOptionListRequest)
	m.UseIgnoreId = types.BoolPointerValue(from.UseIgnoreId)
	m.UseLeaseScavengeTime = types.BoolPointerValue(from.UseLeaseScavengeTime)
	m.UseLogicFilterRules = types.BoolPointerValue(from.UseLogicFilterRules)
	m.UseNextserver = types.BoolPointerValue(from.UseNextserver)
	m.UseOptions = types.BoolPointerValue(from.UseOptions)
	m.UsePxeLeaseTime = types.BoolPointerValue(from.UsePxeLeaseTime)
	m.UseUpdateDnsOnLeaseRenewal = types.BoolPointerValue(from.UseUpdateDnsOnLeaseRenewal)
}

func RemoveDefaultDHCPOptions(ctx context.Context, diags *diag.Diagnostics, options []dhcp.SharednetworkOptions, planOptions internaltypes.UnorderedListValue) internaltypes.UnorderedListValue {
	defaultOptionName := "dhcp-lease-time"

	// If no options, return empty list
	if len(options) == 0 {
		return internaltypes.NewUnorderedListValueNull(types.ObjectType{AttrTypes: SharednetworkOptionsAttrTypes})
	}

	// If plan options is null or unknown, return original options
	if planOptions.IsNull() || planOptions.IsUnknown() {
		return flex.FlattenFrameworkUnorderedListNestedBlock(ctx, options, SharednetworkOptionsAttrTypes, diags, FlattenSharednetworkOptions)
	}

	// Convert plan options to a map for easy lookup
	baseList, err := planOptions.ToListValue(ctx)
	if err != nil {
		return flex.FlattenFrameworkUnorderedListNestedBlock(ctx, options, SharednetworkOptionsAttrTypes, diags, FlattenSharednetworkOptions)
	}

	planOptionsArr := flex.ExpandFrameworkListNestedBlock(ctx, baseList, diags, ExpandSharednetworkOptions)
	planOptionsMap := make(map[string]dhcp.SharednetworkOptions)
	var planOrder []string
	for _, opt := range planOptionsArr {
		if opt.Name != nil {
			planOptionsMap[*opt.Name] = opt
			planOrder = append(planOrder, *opt.Name)
		}
	}

	// Convert current options to a map
	currentOptionsMap := make(map[string]dhcp.SharednetworkOptions)
	for _, opt := range options {
		if opt.Name != nil {
			currentOptionsMap[*opt.Name] = opt
		}
	}

	// Build result maintaining plan order
	var result []dhcp.SharednetworkOptions
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

	return flex.FlattenFrameworkUnorderedListNestedBlock(ctx, result, SharednetworkOptionsAttrTypes, diags, FlattenSharednetworkOptions)
}
