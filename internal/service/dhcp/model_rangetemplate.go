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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	internaltypes "github.com/infobloxopen/terraform-provider-nios/internal/types"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type RangetemplateModel struct {
	Ref                            types.String                     `tfsdk:"ref"`
	Bootfile                       types.String                     `tfsdk:"bootfile"`
	Bootserver                     types.String                     `tfsdk:"bootserver"`
	CloudApiCompatible             types.Bool                       `tfsdk:"cloud_api_compatible"`
	Comment                        types.String                     `tfsdk:"comment"`
	DdnsDomainname                 types.String                     `tfsdk:"ddns_domainname"`
	DdnsGenerateHostname           types.Bool                       `tfsdk:"ddns_generate_hostname"`
	DelegatedMember                types.Object                     `tfsdk:"delegated_member"`
	DenyAllClients                 types.Bool                       `tfsdk:"deny_all_clients"`
	DenyBootp                      types.Bool                       `tfsdk:"deny_bootp"`
	EmailList                      internaltypes.UnorderedListValue `tfsdk:"email_list"`
	EnableDdns                     types.Bool                       `tfsdk:"enable_ddns"`
	EnableDhcpThresholds           types.Bool                       `tfsdk:"enable_dhcp_thresholds"`
	EnableEmailWarnings            types.Bool                       `tfsdk:"enable_email_warnings"`
	EnablePxeLeaseTime             types.Bool                       `tfsdk:"enable_pxe_lease_time"`
	EnableSnmpWarnings             types.Bool                       `tfsdk:"enable_snmp_warnings"`
	Exclude                        types.List                       `tfsdk:"exclude"`
	ExtAttrs                       types.Map                        `tfsdk:"extattrs"`
	ExtAttrsAll                    types.Map                        `tfsdk:"extattrs_all"`
	FailoverAssociation            types.String                     `tfsdk:"failover_association"`
	FingerprintFilterRules         types.List                       `tfsdk:"fingerprint_filter_rules"`
	HighWaterMark                  types.Int64                      `tfsdk:"high_water_mark"`
	HighWaterMarkReset             types.Int64                      `tfsdk:"high_water_mark_reset"`
	IgnoreDhcpOptionListRequest    types.Bool                       `tfsdk:"ignore_dhcp_option_list_request"`
	KnownClients                   types.String                     `tfsdk:"known_clients"`
	LeaseScavengeTime              types.Int64                      `tfsdk:"lease_scavenge_time"`
	LogicFilterRules               types.List                       `tfsdk:"logic_filter_rules"`
	LowWaterMark                   types.Int64                      `tfsdk:"low_water_mark"`
	LowWaterMarkReset              types.Int64                      `tfsdk:"low_water_mark_reset"`
	MacFilterRules                 types.List                       `tfsdk:"mac_filter_rules"`
	Member                         types.Object                     `tfsdk:"member"`
	MsOptions                      types.List                       `tfsdk:"ms_options"`
	MsServer                       types.Object                     `tfsdk:"ms_server"`
	NacFilterRules                 types.List                       `tfsdk:"nac_filter_rules"`
	Name                           types.String                     `tfsdk:"name"`
	Nextserver                     types.String                     `tfsdk:"nextserver"`
	NumberOfAddresses              types.Int64                      `tfsdk:"number_of_addresses"`
	Offset                         types.Int64                      `tfsdk:"offset"`
	OptionFilterRules              types.List                       `tfsdk:"option_filter_rules"`
	Options                        types.List                       `tfsdk:"options"`
	PxeLeaseTime                   types.Int64                      `tfsdk:"pxe_lease_time"`
	RecycleLeases                  types.Bool                       `tfsdk:"recycle_leases"`
	RelayAgentFilterRules          types.List                       `tfsdk:"relay_agent_filter_rules"`
	ServerAssociationType          types.String                     `tfsdk:"server_association_type"`
	UnknownClients                 types.String                     `tfsdk:"unknown_clients"`
	UpdateDnsOnLeaseRenewal        types.Bool                       `tfsdk:"update_dns_on_lease_renewal"`
	UseBootfile                    types.Bool                       `tfsdk:"use_bootfile"`
	UseBootserver                  types.Bool                       `tfsdk:"use_bootserver"`
	UseDdnsDomainname              types.Bool                       `tfsdk:"use_ddns_domainname"`
	UseDdnsGenerateHostname        types.Bool                       `tfsdk:"use_ddns_generate_hostname"`
	UseDenyBootp                   types.Bool                       `tfsdk:"use_deny_bootp"`
	UseEmailList                   types.Bool                       `tfsdk:"use_email_list"`
	UseEnableDdns                  types.Bool                       `tfsdk:"use_enable_ddns"`
	UseEnableDhcpThresholds        types.Bool                       `tfsdk:"use_enable_dhcp_thresholds"`
	UseIgnoreDhcpOptionListRequest types.Bool                       `tfsdk:"use_ignore_dhcp_option_list_request"`
	UseKnownClients                types.Bool                       `tfsdk:"use_known_clients"`
	UseLeaseScavengeTime           types.Bool                       `tfsdk:"use_lease_scavenge_time"`
	UseLogicFilterRules            types.Bool                       `tfsdk:"use_logic_filter_rules"`
	UseMsOptions                   types.Bool                       `tfsdk:"use_ms_options"`
	UseNextserver                  types.Bool                       `tfsdk:"use_nextserver"`
	UseOptions                     types.Bool                       `tfsdk:"use_options"`
	UsePxeLeaseTime                types.Bool                       `tfsdk:"use_pxe_lease_time"`
	UseRecycleLeases               types.Bool                       `tfsdk:"use_recycle_leases"`
	UseUnknownClients              types.Bool                       `tfsdk:"use_unknown_clients"`
	UseUpdateDnsOnLeaseRenewal     types.Bool                       `tfsdk:"use_update_dns_on_lease_renewal"`
}

var RangetemplateAttrTypes = map[string]attr.Type{
	"ref":                                 types.StringType,
	"bootfile":                            types.StringType,
	"bootserver":                          types.StringType,
	"cloud_api_compatible":                types.BoolType,
	"comment":                             types.StringType,
	"ddns_domainname":                     types.StringType,
	"ddns_generate_hostname":              types.BoolType,
	"delegated_member":                    types.ObjectType{AttrTypes: RangetemplateDelegatedMemberAttrTypes},
	"deny_all_clients":                    types.BoolType,
	"deny_bootp":                          types.BoolType,
	"email_list":                          internaltypes.UnorderedListOfStringType,
	"enable_ddns":                         types.BoolType,
	"enable_dhcp_thresholds":              types.BoolType,
	"enable_email_warnings":               types.BoolType,
	"enable_pxe_lease_time":               types.BoolType,
	"enable_snmp_warnings":                types.BoolType,
	"exclude":                             types.ListType{ElemType: types.ObjectType{AttrTypes: RangetemplateExcludeAttrTypes}},
	"extattrs":                            types.MapType{ElemType: types.StringType},
	"extattrs_all":                        types.MapType{ElemType: types.StringType},
	"failover_association":                types.StringType,
	"fingerprint_filter_rules":            types.ListType{ElemType: types.ObjectType{AttrTypes: RangetemplateFingerprintFilterRulesAttrTypes}},
	"high_water_mark":                     types.Int64Type,
	"high_water_mark_reset":               types.Int64Type,
	"ignore_dhcp_option_list_request":     types.BoolType,
	"known_clients":                       types.StringType,
	"lease_scavenge_time":                 types.Int64Type,
	"logic_filter_rules":                  types.ListType{ElemType: types.ObjectType{AttrTypes: RangetemplateLogicFilterRulesAttrTypes}},
	"low_water_mark":                      types.Int64Type,
	"low_water_mark_reset":                types.Int64Type,
	"mac_filter_rules":                    types.ListType{ElemType: types.ObjectType{AttrTypes: RangetemplateMacFilterRulesAttrTypes}},
	"member":                              types.ObjectType{AttrTypes: RangetemplateMemberAttrTypes},
	"ms_options":                          types.ListType{ElemType: types.ObjectType{AttrTypes: RangetemplateMsOptionsAttrTypes}},
	"ms_server":                           types.ObjectType{AttrTypes: RangetemplateMsServerAttrTypes},
	"nac_filter_rules":                    types.ListType{ElemType: types.ObjectType{AttrTypes: RangetemplateNacFilterRulesAttrTypes}},
	"name":                                types.StringType,
	"nextserver":                          types.StringType,
	"number_of_addresses":                 types.Int64Type,
	"offset":                              types.Int64Type,
	"option_filter_rules":                 types.ListType{ElemType: types.ObjectType{AttrTypes: RangetemplateOptionFilterRulesAttrTypes}},
	"options":                             types.ListType{ElemType: types.ObjectType{AttrTypes: RangetemplateOptionsAttrTypes}},
	"pxe_lease_time":                      types.Int64Type,
	"recycle_leases":                      types.BoolType,
	"relay_agent_filter_rules":            types.ListType{ElemType: types.ObjectType{AttrTypes: RangetemplateRelayAgentFilterRulesAttrTypes}},
	"server_association_type":             types.StringType,
	"unknown_clients":                     types.StringType,
	"update_dns_on_lease_renewal":         types.BoolType,
	"use_bootfile":                        types.BoolType,
	"use_bootserver":                      types.BoolType,
	"use_ddns_domainname":                 types.BoolType,
	"use_ddns_generate_hostname":          types.BoolType,
	"use_deny_bootp":                      types.BoolType,
	"use_email_list":                      types.BoolType,
	"use_enable_ddns":                     types.BoolType,
	"use_enable_dhcp_thresholds":          types.BoolType,
	"use_ignore_dhcp_option_list_request": types.BoolType,
	"use_known_clients":                   types.BoolType,
	"use_lease_scavenge_time":             types.BoolType,
	"use_logic_filter_rules":              types.BoolType,
	"use_ms_options":                      types.BoolType,
	"use_nextserver":                      types.BoolType,
	"use_options":                         types.BoolType,
	"use_pxe_lease_time":                  types.BoolType,
	"use_recycle_leases":                  types.BoolType,
	"use_unknown_clients":                 types.BoolType,
	"use_update_dns_on_lease_renewal":     types.BoolType,
}

var RangetemplateResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"bootfile": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_bootfile")),
		},
		MarkdownDescription: "The bootfile name for the range. You can configure the DHCP server to support clients that use the boot file name option in their DHCPREQUEST messages.",
	},
	"bootserver": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_bootserver")),
		},
		MarkdownDescription: "The bootserver address for the range. You can specify the name and/or IP address of the boot server that the host needs to boot. The boot server IPv4 Address or name in FQDN format.",
	},
	"cloud_api_compatible": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "This flag controls whether this template can be used to create network objects in a cloud-computing deployment.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "A descriptive comment of a range template object.",
	},
	"ddns_domainname": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_ddns_domainname")),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The dynamic DNS domain name the appliance uses specifically for DDNS updates for this range.",
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
	"delegated_member": schema.SingleNestedAttribute{
		Attributes:          RangetemplateDelegatedMemberResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The vconnector member that the object should be delegated to when created from this range template.",
	},
	"deny_all_clients": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True, send NAK forcing the client to take the new address.",
	},
	"deny_bootp": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_deny_bootp")),
		},
		MarkdownDescription: "Determines if BOOTP settings are disabled and BOOTP requests will be denied.",
	},
	"email_list": schema.ListAttribute{
		CustomType:  internaltypes.UnorderedListOfStringType,
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_email_list")),
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The e-mail lists to which the appliance sends DHCP threshold alarm e-mail messages.",
	},
	"enable_ddns": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_enable_ddns")),
		},
		MarkdownDescription: "Determines if the DHCP server sends DDNS updates to DNS servers in the same Grid, and to external DNS servers.",
	},
	"enable_dhcp_thresholds": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_enable_dhcp_thresholds")),
		},
		MarkdownDescription: "Determines if DHCP thresholds are enabled for the range.",
	},
	"enable_email_warnings": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if DHCP threshold warnings are sent through email.",
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
		MarkdownDescription: "Determines if DHCP threshold warnings are sent through SNMP.",
	},
	"exclude": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RangetemplateExcludeResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "These are ranges of IP addresses that the appliance does not use to assign to clients. You can use these exclusion addresses as static IP addresses. They contain the start and end addresses of the exclusion range, and optionally, information about this exclusion range.",
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
	"failover_association": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("server_association_type")),
		},
		MarkdownDescription: "The name of the failover association: the server in this failover association will serve the IPv4 range in case the main server is out of service. {rangetemplate:rangetemplate} must be set to 'FAILOVER' or 'FAILOVER_MS' if you want the failover association specified here to serve the range.",
	},
	"fingerprint_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RangetemplateFingerprintFilterRulesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
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
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_ignore_dhcp_option_list_request")),
		},
		MarkdownDescription: "If this field is set to False, the appliance returns all DHCP options the client is eligible to receive, rather than only the list of options the client has requested.",
	},
	"known_clients": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			stringvalidator.OneOf("Allow", "Deny"),
			stringvalidator.AlsoRequires(path.MatchRoot("use_known_clients")),
		},
		MarkdownDescription: "Permission for known clients. If set to 'Deny' known clients will be denied IP addresses. Known clients include roaming hosts and clients with fixed addresses or DHCP host entries. Unknown clients include clients that are not roaming hosts and clients that do not have fixed addresses or DHCP host entries.",
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
			Attributes: RangetemplateLogicFilterRulesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_logic_filter_rules")),
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "This field contains the logic filters to be applied on this range. This list corresponds to the match rules that are written to the dhcpd configuration file.",
	},
	"low_water_mark": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(0),
		Validators: []validator.Int64{
			int64validator.Between(1, 100),
		},
		MarkdownDescription: "The percentage of DHCP range usage below which the Infoblox appliance generates a syslog message and sends a warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100.",
	},
	"low_water_mark_reset": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(10),
		Validators: []validator.Int64{
			int64validator.Between(1, 100),
		},
		MarkdownDescription: "The percentage of DHCP range usage threshold below which range usage is not expected and may warrant your attention. When the low watermark is crossed, the Infoblox appliance generates a syslog message and sends a warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100. The low watermark reset value must be higher than the low watermark value.",
	},
	"mac_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RangetemplateMacFilterRulesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "This field contains the MAC filters to be applied to this range. The appliance uses the matching rules of these filters to select the address range from which it assigns a lease.",
	},
	"member": schema.SingleNestedAttribute{
		Attributes:          RangetemplateMemberResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The member that will provide service for this range.",
	},
	"ms_options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RangetemplateMsOptionsResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_ms_options")),
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The Microsoft DHCP options for this range.",
	},
	"ms_server": schema.SingleNestedAttribute{
		Attributes:          RangetemplateMsServerResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The Microsoft server that will provide service for this range. server_association_type needs to be set to ‘MS_SERVER’ if you want the server specified here to serve the range. For searching by this field you should use a HTTP method that contains a body (POST or PUT) with MS DHCP server structure and the request should have option _method=GET.",
	},
	"nac_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RangetemplateNacFilterRulesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "This field contains the NAC filters to be applied to this range. The appliance uses the matching rules of these filters to select the address range from which it assigns a lease.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of a range template object.",
	},
	"nextserver": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_nextserver")),
		},
		MarkdownDescription: "The name in FQDN and/or IPv4 Address format of the next server that the host needs to boot.",
	},
	"number_of_addresses": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "The number of addresses for this range.",
	},
	"offset": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "The start address offset for this range.",
	},
	"option_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RangetemplateOptionFilterRulesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "This field contains the Option filters to be applied to this range. The appliance uses the matching rules of these filters to select the address range from which it assigns a lease.",
	},
	"options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RangetemplateOptionsResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Default: listdefault.StaticValue(
			types.ListValueMust(
				types.ObjectType{AttrTypes: RangetemplateOptionsAttrTypes},
				[]attr.Value{},
			),
		),
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_options")),
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "An unordered set of DHCP option dhcpoption structs that lists the DHCP options associated with the object.",
	},
	"pxe_lease_time": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_pxe_lease_time")),
		},
		MarkdownDescription: "The PXE lease time value for a range object. Some hosts use PXE (Preboot Execution Environment) to boot remotely from a server. To better manage your IP resources, set a different lease time for PXE boot requests. You can configure the DHCP server to allocate an IP address with a shorter lease time to hosts that send PXE boot requests, so IP addresses are not leased longer than necessary. A 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached.",
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
			Attributes: RangetemplateRelayAgentFilterRulesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "This field contains the Relay Agent filters to be applied to this range. The appliance uses the matching rules of these filters to select the address range from which it assigns a lease.",
	},
	"server_association_type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("NONE"),
		Validators: []validator.String{
			stringvalidator.OneOf("FAILOVER", "MEMBER", "MS_FAILOVER", "MS_SERVER", "NONE"),
		},
		MarkdownDescription: "The type of server that is going to serve the range.",
	},
	"unknown_clients": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			stringvalidator.OneOf("Allow", "Deny"),
			stringvalidator.AlsoRequires(path.MatchRoot("use_unknown_clients")),
		},
		MarkdownDescription: "Permission for unknown clients. If set to 'Deny' unknown clients will be denied IP addresses. Known clients include roaming hosts and clients with fixed addresses or DHCP host entries. Unknown clients include clients that are not roaming hosts and clients that do not have fixed addresses or DHCP host entries.",
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
	"use_ignore_dhcp_option_list_request": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: ignore_dhcp_option_list_request",
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

func (m *RangetemplateModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.Rangetemplate {
	if m == nil {
		return nil
	}
	to := &dhcp.Rangetemplate{
		Bootfile:                       flex.ExpandStringPointer(m.Bootfile),
		Bootserver:                     flex.ExpandStringPointer(m.Bootserver),
		CloudApiCompatible:             flex.ExpandBoolPointer(m.CloudApiCompatible),
		Comment:                        flex.ExpandStringPointer(m.Comment),
		DdnsDomainname:                 flex.ExpandStringPointer(m.DdnsDomainname),
		DdnsGenerateHostname:           flex.ExpandBoolPointer(m.DdnsGenerateHostname),
		DelegatedMember:                ExpandRangetemplateDelegatedMember(ctx, m.DelegatedMember, diags),
		DenyAllClients:                 flex.ExpandBoolPointer(m.DenyAllClients),
		DenyBootp:                      flex.ExpandBoolPointer(m.DenyBootp),
		EmailList:                      flex.ExpandFrameworkListString(ctx, m.EmailList, diags),
		EnableDdns:                     flex.ExpandBoolPointer(m.EnableDdns),
		EnableDhcpThresholds:           flex.ExpandBoolPointer(m.EnableDhcpThresholds),
		EnableEmailWarnings:            flex.ExpandBoolPointer(m.EnableEmailWarnings),
		EnablePxeLeaseTime:             flex.ExpandBoolPointer(m.EnablePxeLeaseTime),
		EnableSnmpWarnings:             flex.ExpandBoolPointer(m.EnableSnmpWarnings),
		Exclude:                        flex.ExpandFrameworkListNestedBlock(ctx, m.Exclude, diags, ExpandRangetemplateExclude),
		ExtAttrs:                       ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		FailoverAssociation:            flex.ExpandStringPointer(m.FailoverAssociation),
		FingerprintFilterRules:         flex.ExpandFrameworkListNestedBlock(ctx, m.FingerprintFilterRules, diags, ExpandRangetemplateFingerprintFilterRules),
		HighWaterMark:                  flex.ExpandInt64Pointer(m.HighWaterMark),
		HighWaterMarkReset:             flex.ExpandInt64Pointer(m.HighWaterMarkReset),
		IgnoreDhcpOptionListRequest:    flex.ExpandBoolPointer(m.IgnoreDhcpOptionListRequest),
		KnownClients:                   flex.ExpandStringPointer(m.KnownClients),
		LeaseScavengeTime:              flex.ExpandInt64Pointer(m.LeaseScavengeTime),
		LogicFilterRules:               flex.ExpandFrameworkListNestedBlock(ctx, m.LogicFilterRules, diags, ExpandRangetemplateLogicFilterRules),
		LowWaterMark:                   flex.ExpandInt64Pointer(m.LowWaterMark),
		LowWaterMarkReset:              flex.ExpandInt64Pointer(m.LowWaterMarkReset),
		MacFilterRules:                 flex.ExpandFrameworkListNestedBlock(ctx, m.MacFilterRules, diags, ExpandRangetemplateMacFilterRules),
		Member:                         ExpandRangetemplateMember(ctx, m.Member, diags),
		MsOptions:                      flex.ExpandFrameworkListNestedBlock(ctx, m.MsOptions, diags, ExpandRangetemplateMsOptions),
		MsServer:                       ExpandRangetemplateMsServer(ctx, m.MsServer, diags),
		NacFilterRules:                 flex.ExpandFrameworkListNestedBlock(ctx, m.NacFilterRules, diags, ExpandRangetemplateNacFilterRules),
		Name:                           flex.ExpandStringPointer(m.Name),
		Nextserver:                     flex.ExpandStringPointer(m.Nextserver),
		NumberOfAddresses:              flex.ExpandInt64Pointer(m.NumberOfAddresses),
		Offset:                         flex.ExpandInt64Pointer(m.Offset),
		OptionFilterRules:              flex.ExpandFrameworkListNestedBlock(ctx, m.OptionFilterRules, diags, ExpandRangetemplateOptionFilterRules),
		Options:                        flex.ExpandFrameworkListNestedBlock(ctx, m.Options, diags, ExpandRangetemplateOptions),
		PxeLeaseTime:                   flex.ExpandInt64Pointer(m.PxeLeaseTime),
		RecycleLeases:                  flex.ExpandBoolPointer(m.RecycleLeases),
		RelayAgentFilterRules:          flex.ExpandFrameworkListNestedBlock(ctx, m.RelayAgentFilterRules, diags, ExpandRangetemplateRelayAgentFilterRules),
		ServerAssociationType:          flex.ExpandStringPointer(m.ServerAssociationType),
		UnknownClients:                 flex.ExpandStringPointer(m.UnknownClients),
		UpdateDnsOnLeaseRenewal:        flex.ExpandBoolPointer(m.UpdateDnsOnLeaseRenewal),
		UseBootfile:                    flex.ExpandBoolPointer(m.UseBootfile),
		UseBootserver:                  flex.ExpandBoolPointer(m.UseBootserver),
		UseDdnsDomainname:              flex.ExpandBoolPointer(m.UseDdnsDomainname),
		UseDdnsGenerateHostname:        flex.ExpandBoolPointer(m.UseDdnsGenerateHostname),
		UseDenyBootp:                   flex.ExpandBoolPointer(m.UseDenyBootp),
		UseEmailList:                   flex.ExpandBoolPointer(m.UseEmailList),
		UseEnableDdns:                  flex.ExpandBoolPointer(m.UseEnableDdns),
		UseEnableDhcpThresholds:        flex.ExpandBoolPointer(m.UseEnableDhcpThresholds),
		UseIgnoreDhcpOptionListRequest: flex.ExpandBoolPointer(m.UseIgnoreDhcpOptionListRequest),
		UseKnownClients:                flex.ExpandBoolPointer(m.UseKnownClients),
		UseLeaseScavengeTime:           flex.ExpandBoolPointer(m.UseLeaseScavengeTime),
		UseLogicFilterRules:            flex.ExpandBoolPointer(m.UseLogicFilterRules),
		UseMsOptions:                   flex.ExpandBoolPointer(m.UseMsOptions),
		UseNextserver:                  flex.ExpandBoolPointer(m.UseNextserver),
		UseOptions:                     flex.ExpandBoolPointer(m.UseOptions),
		UsePxeLeaseTime:                flex.ExpandBoolPointer(m.UsePxeLeaseTime),
		UseRecycleLeases:               flex.ExpandBoolPointer(m.UseRecycleLeases),
		UseUnknownClients:              flex.ExpandBoolPointer(m.UseUnknownClients),
		UseUpdateDnsOnLeaseRenewal:     flex.ExpandBoolPointer(m.UseUpdateDnsOnLeaseRenewal),
	}
	return to
}

func FlattenRangetemplate(ctx context.Context, from *dhcp.Rangetemplate, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RangetemplateAttrTypes)
	}
	m := RangetemplateModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, RangetemplateAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RangetemplateModel) Flatten(ctx context.Context, from *dhcp.Rangetemplate, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RangetemplateModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Bootfile = flex.FlattenStringPointer(from.Bootfile)
	m.Bootserver = flex.FlattenStringPointer(from.Bootserver)
	m.CloudApiCompatible = types.BoolPointerValue(from.CloudApiCompatible)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.DdnsDomainname = flex.FlattenStringPointer(from.DdnsDomainname)
	m.DdnsGenerateHostname = types.BoolPointerValue(from.DdnsGenerateHostname)
	m.DelegatedMember = FlattenRangetemplateDelegatedMember(ctx, from.DelegatedMember, diags)
	m.DenyAllClients = types.BoolPointerValue(from.DenyAllClients)
	m.DenyBootp = types.BoolPointerValue(from.DenyBootp)
	m.EmailList = flex.FlattenFrameworkUnorderedList(ctx, types.StringType, from.EmailList, diags)
	m.EnableDdns = types.BoolPointerValue(from.EnableDdns)
	m.EnableDhcpThresholds = types.BoolPointerValue(from.EnableDhcpThresholds)
	m.EnableEmailWarnings = types.BoolPointerValue(from.EnableEmailWarnings)
	m.EnablePxeLeaseTime = types.BoolPointerValue(from.EnablePxeLeaseTime)
	m.EnableSnmpWarnings = types.BoolPointerValue(from.EnableSnmpWarnings)
	m.Exclude = flex.FlattenFrameworkListNestedBlock(ctx, from.Exclude, RangetemplateExcludeAttrTypes, diags, FlattenRangetemplateExclude)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.FailoverAssociation = flex.FlattenStringPointer(from.FailoverAssociation)
	m.FingerprintFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.FingerprintFilterRules, RangetemplateFingerprintFilterRulesAttrTypes, diags, FlattenRangetemplateFingerprintFilterRules)
	m.HighWaterMark = flex.FlattenInt64Pointer(from.HighWaterMark)
	m.HighWaterMarkReset = flex.FlattenInt64Pointer(from.HighWaterMarkReset)
	m.IgnoreDhcpOptionListRequest = types.BoolPointerValue(from.IgnoreDhcpOptionListRequest)
	m.KnownClients = flex.FlattenStringPointerNilAsNotEmpty(from.KnownClients)
	m.LeaseScavengeTime = flex.FlattenInt64Pointer(from.LeaseScavengeTime)
	m.LogicFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.LogicFilterRules, RangetemplateLogicFilterRulesAttrTypes, diags, FlattenRangetemplateLogicFilterRules)
	m.LowWaterMark = flex.FlattenInt64Pointer(from.LowWaterMark)
	m.LowWaterMarkReset = flex.FlattenInt64Pointer(from.LowWaterMarkReset)
	m.MacFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.MacFilterRules, RangetemplateMacFilterRulesAttrTypes, diags, FlattenRangetemplateMacFilterRules)
	m.Member = FlattenRangetemplateMember(ctx, from.Member, diags)
	m.MsOptions = flex.FlattenFrameworkListNestedBlock(ctx, from.MsOptions, RangetemplateMsOptionsAttrTypes, diags, FlattenRangetemplateMsOptions)
	m.MsServer = FlattenRangetemplateMsServer(ctx, from.MsServer, diags)
	m.NacFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.NacFilterRules, RangetemplateNacFilterRulesAttrTypes, diags, FlattenRangetemplateNacFilterRules)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Nextserver = flex.FlattenStringPointer(from.Nextserver)
	m.NumberOfAddresses = flex.FlattenInt64Pointer(from.NumberOfAddresses)
	m.Offset = flex.FlattenInt64Pointer(from.Offset)
	m.OptionFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.OptionFilterRules, RangetemplateOptionFilterRulesAttrTypes, diags, FlattenRangetemplateOptionFilterRules)
	planOptions := m.Options
	m.Options = flex.FlattenFrameworkListNestedBlock(ctx, from.Options, RangetemplateOptionsAttrTypes, diags, FlattenRangetemplateOptions)
	if !planOptions.IsUnknown() {
		reOrderedOptions, diags := utils.ReorderAndFilterDHCPOptions(ctx, planOptions, m.Options)
		if !diags.HasError() {
			m.Options = reOrderedOptions.(basetypes.ListValue)
		}
	}
	m.PxeLeaseTime = flex.FlattenInt64Pointer(from.PxeLeaseTime)
	m.RecycleLeases = types.BoolPointerValue(from.RecycleLeases)
	m.RelayAgentFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.RelayAgentFilterRules, RangetemplateRelayAgentFilterRulesAttrTypes, diags, FlattenRangetemplateRelayAgentFilterRules)
	m.ServerAssociationType = flex.FlattenStringPointer(from.ServerAssociationType)
	m.UnknownClients = flex.FlattenStringPointerNilAsNotEmpty(from.UnknownClients)
	m.UpdateDnsOnLeaseRenewal = types.BoolPointerValue(from.UpdateDnsOnLeaseRenewal)
	m.UseBootfile = types.BoolPointerValue(from.UseBootfile)
	m.UseBootserver = types.BoolPointerValue(from.UseBootserver)
	m.UseDdnsDomainname = types.BoolPointerValue(from.UseDdnsDomainname)
	m.UseDdnsGenerateHostname = types.BoolPointerValue(from.UseDdnsGenerateHostname)
	m.UseDenyBootp = types.BoolPointerValue(from.UseDenyBootp)
	m.UseEmailList = types.BoolPointerValue(from.UseEmailList)
	m.UseEnableDdns = types.BoolPointerValue(from.UseEnableDdns)
	m.UseEnableDhcpThresholds = types.BoolPointerValue(from.UseEnableDhcpThresholds)
	m.UseIgnoreDhcpOptionListRequest = types.BoolPointerValue(from.UseIgnoreDhcpOptionListRequest)
	m.UseKnownClients = types.BoolPointerValue(from.UseKnownClients)
	m.UseLeaseScavengeTime = types.BoolPointerValue(from.UseLeaseScavengeTime)
	m.UseLogicFilterRules = types.BoolPointerValue(from.UseLogicFilterRules)
	m.UseMsOptions = types.BoolPointerValue(from.UseMsOptions)
	m.UseNextserver = types.BoolPointerValue(from.UseNextserver)
	m.UseOptions = types.BoolPointerValue(from.UseOptions)
	m.UsePxeLeaseTime = types.BoolPointerValue(from.UsePxeLeaseTime)
	m.UseRecycleLeases = types.BoolPointerValue(from.UseRecycleLeases)
	m.UseUnknownClients = types.BoolPointerValue(from.UseUnknownClients)
	m.UseUpdateDnsOnLeaseRenewal = types.BoolPointerValue(from.UseUpdateDnsOnLeaseRenewal)
}
