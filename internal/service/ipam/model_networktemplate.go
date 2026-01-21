package ipam

import (
	"context"

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
	importmod "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/import"
	internaltypes "github.com/infobloxopen/terraform-provider-nios/internal/types"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type NetworktemplateModel struct {
	Ref                            types.String                             `tfsdk:"ref"`
	Uuid                           types.String                             `tfsdk:"uuid"`
	AllowAnyNetmask                types.Bool                               `tfsdk:"allow_any_netmask"`
	Authority                      types.Bool                               `tfsdk:"authority"`
	AutoCreateReversezone          types.Bool                               `tfsdk:"auto_create_reversezone"`
	Bootfile                       types.String                             `tfsdk:"bootfile"`
	Bootserver                     types.String                             `tfsdk:"bootserver"`
	CloudApiCompatible             types.Bool                               `tfsdk:"cloud_api_compatible"`
	Comment                        types.String                             `tfsdk:"comment"`
	DdnsDomainname                 internaltypes.CaseInsensitiveStringValue `tfsdk:"ddns_domainname"`
	DdnsGenerateHostname           types.Bool                               `tfsdk:"ddns_generate_hostname"`
	DdnsServerAlwaysUpdates        types.Bool                               `tfsdk:"ddns_server_always_updates"`
	DdnsTtl                        types.Int64                              `tfsdk:"ddns_ttl"`
	DdnsUpdateFixedAddresses       types.Bool                               `tfsdk:"ddns_update_fixed_addresses"`
	DdnsUseOption81                types.Bool                               `tfsdk:"ddns_use_option81"`
	DelegatedMember                types.Object                             `tfsdk:"delegated_member"`
	DenyBootp                      types.Bool                               `tfsdk:"deny_bootp"`
	EmailList                      internaltypes.UnorderedListValue         `tfsdk:"email_list"`
	EnableDdns                     types.Bool                               `tfsdk:"enable_ddns"`
	EnableDhcpThresholds           types.Bool                               `tfsdk:"enable_dhcp_thresholds"`
	EnableEmailWarnings            types.Bool                               `tfsdk:"enable_email_warnings"`
	EnablePxeLeaseTime             types.Bool                               `tfsdk:"enable_pxe_lease_time"`
	EnableSnmpWarnings             types.Bool                               `tfsdk:"enable_snmp_warnings"`
	ExtAttrs                       types.Map                                `tfsdk:"extattrs"`
	ExtAttrsAll                    types.Map                                `tfsdk:"extattrs_all"`
	FixedAddressTemplates          types.List                               `tfsdk:"fixed_address_templates"`
	HighWaterMark                  types.Int64                              `tfsdk:"high_water_mark"`
	HighWaterMarkReset             types.Int64                              `tfsdk:"high_water_mark_reset"`
	IgnoreDhcpOptionListRequest    types.Bool                               `tfsdk:"ignore_dhcp_option_list_request"`
	IpamEmailAddresses             types.List                               `tfsdk:"ipam_email_addresses"`
	IpamThresholdSettings          types.Object                             `tfsdk:"ipam_threshold_settings"`
	IpamTrapSettings               types.Object                             `tfsdk:"ipam_trap_settings"`
	LeaseScavengeTime              types.Int64                              `tfsdk:"lease_scavenge_time"`
	LogicFilterRules               types.List                               `tfsdk:"logic_filter_rules"`
	LowWaterMark                   types.Int64                              `tfsdk:"low_water_mark"`
	LowWaterMarkReset              types.Int64                              `tfsdk:"low_water_mark_reset"`
	Members                        types.List                               `tfsdk:"members"`
	Name                           types.String                             `tfsdk:"name"`
	Netmask                        types.Int64                              `tfsdk:"netmask"`
	Nextserver                     types.String                             `tfsdk:"nextserver"`
	Options                        types.List                               `tfsdk:"options"`
	PxeLeaseTime                   types.Int64                              `tfsdk:"pxe_lease_time"`
	RangeTemplates                 types.List                               `tfsdk:"range_templates"`
	RecycleLeases                  types.Bool                               `tfsdk:"recycle_leases"`
	Rir                            types.String                             `tfsdk:"rir"`
	RirOrganization                types.String                             `tfsdk:"rir_organization"`
	RirRegistrationAction          types.String                             `tfsdk:"rir_registration_action"`
	RirRegistrationStatus          types.String                             `tfsdk:"rir_registration_status"`
	SendRirRequest                 types.Bool                               `tfsdk:"send_rir_request"`
	UpdateDnsOnLeaseRenewal        types.Bool                               `tfsdk:"update_dns_on_lease_renewal"`
	UseAuthority                   types.Bool                               `tfsdk:"use_authority"`
	UseBootfile                    types.Bool                               `tfsdk:"use_bootfile"`
	UseBootserver                  types.Bool                               `tfsdk:"use_bootserver"`
	UseDdnsDomainname              types.Bool                               `tfsdk:"use_ddns_domainname"`
	UseDdnsGenerateHostname        types.Bool                               `tfsdk:"use_ddns_generate_hostname"`
	UseDdnsTtl                     types.Bool                               `tfsdk:"use_ddns_ttl"`
	UseDdnsUpdateFixedAddresses    types.Bool                               `tfsdk:"use_ddns_update_fixed_addresses"`
	UseDdnsUseOption81             types.Bool                               `tfsdk:"use_ddns_use_option81"`
	UseDenyBootp                   types.Bool                               `tfsdk:"use_deny_bootp"`
	UseEmailList                   types.Bool                               `tfsdk:"use_email_list"`
	UseEnableDdns                  types.Bool                               `tfsdk:"use_enable_ddns"`
	UseEnableDhcpThresholds        types.Bool                               `tfsdk:"use_enable_dhcp_thresholds"`
	UseIgnoreDhcpOptionListRequest types.Bool                               `tfsdk:"use_ignore_dhcp_option_list_request"`
	UseIpamEmailAddresses          types.Bool                               `tfsdk:"use_ipam_email_addresses"`
	UseIpamThresholdSettings       types.Bool                               `tfsdk:"use_ipam_threshold_settings"`
	UseIpamTrapSettings            types.Bool                               `tfsdk:"use_ipam_trap_settings"`
	UseLeaseScavengeTime           types.Bool                               `tfsdk:"use_lease_scavenge_time"`
	UseLogicFilterRules            types.Bool                               `tfsdk:"use_logic_filter_rules"`
	UseNextserver                  types.Bool                               `tfsdk:"use_nextserver"`
	UseOptions                     types.Bool                               `tfsdk:"use_options"`
	UsePxeLeaseTime                types.Bool                               `tfsdk:"use_pxe_lease_time"`
	UseRecycleLeases               types.Bool                               `tfsdk:"use_recycle_leases"`
	UseUpdateDnsOnLeaseRenewal     types.Bool                               `tfsdk:"use_update_dns_on_lease_renewal"`
}

var NetworktemplateAttrTypes = map[string]attr.Type{
	"ref":                                 types.StringType,
	"uuid":                                types.StringType,
	"allow_any_netmask":                   types.BoolType,
	"authority":                           types.BoolType,
	"auto_create_reversezone":             types.BoolType,
	"bootfile":                            types.StringType,
	"bootserver":                          types.StringType,
	"cloud_api_compatible":                types.BoolType,
	"comment":                             types.StringType,
	"ddns_domainname":                     internaltypes.CaseInsensitiveString{},
	"ddns_generate_hostname":              types.BoolType,
	"ddns_server_always_updates":          types.BoolType,
	"ddns_ttl":                            types.Int64Type,
	"ddns_update_fixed_addresses":         types.BoolType,
	"ddns_use_option81":                   types.BoolType,
	"delegated_member":                    types.ObjectType{AttrTypes: NetworktemplateDelegatedMemberAttrTypes},
	"deny_bootp":                          types.BoolType,
	"email_list":                          internaltypes.UnorderedListOfStringType,
	"enable_ddns":                         types.BoolType,
	"enable_dhcp_thresholds":              types.BoolType,
	"enable_email_warnings":               types.BoolType,
	"enable_pxe_lease_time":               types.BoolType,
	"enable_snmp_warnings":                types.BoolType,
	"extattrs":                            types.MapType{ElemType: types.StringType},
	"extattrs_all":                        types.MapType{ElemType: types.StringType},
	"fixed_address_templates":             types.ListType{ElemType: types.StringType},
	"high_water_mark":                     types.Int64Type,
	"high_water_mark_reset":               types.Int64Type,
	"ignore_dhcp_option_list_request":     types.BoolType,
	"ipam_email_addresses":                types.ListType{ElemType: types.StringType},
	"ipam_threshold_settings":             types.ObjectType{AttrTypes: NetworktemplateIpamThresholdSettingsAttrTypes},
	"ipam_trap_settings":                  types.ObjectType{AttrTypes: NetworktemplateIpamTrapSettingsAttrTypes},
	"lease_scavenge_time":                 types.Int64Type,
	"logic_filter_rules":                  types.ListType{ElemType: types.ObjectType{AttrTypes: NetworktemplateLogicFilterRulesAttrTypes}},
	"low_water_mark":                      types.Int64Type,
	"low_water_mark_reset":                types.Int64Type,
	"members":                             types.ListType{ElemType: types.ObjectType{AttrTypes: NetworktemplateMembersAttrTypes}},
	"name":                                types.StringType,
	"netmask":                             types.Int64Type,
	"nextserver":                          types.StringType,
	"options":                             types.ListType{ElemType: types.ObjectType{AttrTypes: NetworktemplateOptionsAttrTypes}},
	"pxe_lease_time":                      types.Int64Type,
	"range_templates":                     types.ListType{ElemType: types.StringType},
	"recycle_leases":                      types.BoolType,
	"rir":                                 types.StringType,
	"rir_organization":                    types.StringType,
	"rir_registration_action":             types.StringType,
	"rir_registration_status":             types.StringType,
	"send_rir_request":                    types.BoolType,
	"update_dns_on_lease_renewal":         types.BoolType,
	"use_authority":                       types.BoolType,
	"use_bootfile":                        types.BoolType,
	"use_bootserver":                      types.BoolType,
	"use_ddns_domainname":                 types.BoolType,
	"use_ddns_generate_hostname":          types.BoolType,
	"use_ddns_ttl":                        types.BoolType,
	"use_ddns_update_fixed_addresses":     types.BoolType,
	"use_ddns_use_option81":               types.BoolType,
	"use_deny_bootp":                      types.BoolType,
	"use_email_list":                      types.BoolType,
	"use_enable_ddns":                     types.BoolType,
	"use_enable_dhcp_thresholds":          types.BoolType,
	"use_ignore_dhcp_option_list_request": types.BoolType,
	"use_ipam_email_addresses":            types.BoolType,
	"use_ipam_threshold_settings":         types.BoolType,
	"use_ipam_trap_settings":              types.BoolType,
	"use_lease_scavenge_time":             types.BoolType,
	"use_logic_filter_rules":              types.BoolType,
	"use_nextserver":                      types.BoolType,
	"use_options":                         types.BoolType,
	"use_pxe_lease_time":                  types.BoolType,
	"use_recycle_leases":                  types.BoolType,
	"use_update_dns_on_lease_renewal":     types.BoolType,
}

var NetworktemplateResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The uuid to the object.",
	},
	"allow_any_netmask": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This flag controls whether the template allows any netmask. You must specify a netmask when creating a network using this template. If you set this parameter to false, you must specify the \"netmask\" field for the network template object.",
	},
	"authority": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_authority")),
		},
		MarkdownDescription: "Authority for the DHCP network.",
	},
	"auto_create_reversezone": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This flag controls whether reverse zones are automatically created when the network is added.",
	},
	"bootfile": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_bootfile")),
		},
		MarkdownDescription: "The boot server IPv4 Address or name in FQDN format for the network. You can specify the name and/or IP address of the boot server that the host needs to boot.",
	},
	"bootserver": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_bootserver")),
			customvalidator.IsValidFQDN(),
		},
		MarkdownDescription: "The bootserver address for the network. You can specify the name and/or IP address of the boot server that the host needs to boot. The boot server IPv4 Address or name in FQDN format.",
	},
	"cloud_api_compatible": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
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
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "This field controls whether the DHCP server is allowed to update DNS, regardless of the DHCP client requests. Note that changes for this field take effect only if ddns_use_option81 is True.",
	},
	"ddns_ttl": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(0),
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_ddns_ttl")),
			int64validator.Between(0, 4294967295),
		},
		MarkdownDescription: "The DNS update Time to Live (TTL) value of a DHCP network object. The TTL is a 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached.",
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
		MarkdownDescription: "The support for DHCP Option 81 at the network level.",
	},
	"delegated_member": schema.SingleNestedAttribute{
		Attributes:          NetworktemplateDelegatedMemberResourceSchemaAttributes,
		Computed:            true,
		Optional:            true,
		MarkdownDescription: "Reference the Cloud Platform Appliance to which authority of the object should be delegated when the object is created using the template.",
	},
	"deny_bootp": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_deny_bootp")),
		},
		MarkdownDescription: "If set to True, BOOTP settings are disabled and BOOTP requests will be denied.",
	},
	"email_list": schema.ListAttribute{
		CustomType:  internaltypes.UnorderedListOfStringType,
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
			listvalidator.AlsoRequires(path.MatchRoot("use_email_list")),
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
		MarkdownDescription: "The dynamic DNS updates flag of a DHCP network object. If set to True, the DHCP server sends DDNS updates to DNS servers in the same Grid, and to external DNS servers.",
	},
	"enable_dhcp_thresholds": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_enable_dhcp_thresholds")),
		},
		MarkdownDescription: "Determines if DHCP thresholds are enabled for the network.",
	},
	"enable_email_warnings": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if DHCP threshold warnings are sent through email.",
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
	"enable_snmp_warnings": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if DHCP threshold warnings are send through SNMP.",
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
	"fixed_address_templates": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The list of fixed address templates assigned to this network template object. When you create a network based on a network template object that contains fixed address templates, the fixed addresses are created based on the associated fixed address templates.",
	},
	"high_water_mark": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(95),
		Validators: []validator.Int64{
			int64validator.Between(1, 100),
		},
		MarkdownDescription: "The percentage of DHCP network usage threshold above which network usage is not expected and may warrant your attention. When the high watermark is reached, the Infoblox appliance generates a syslog message and sends a warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100.",
	},
	"high_water_mark_reset": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(85),
		Validators: []validator.Int64{
			int64validator.Between(1, 100),
		},
		MarkdownDescription: "The percentage of DHCP network usage below which the corresponding SNMP trap is reset. A number that specifies the percentage of allocated addresses. The range is from 1 to 100. The high watermark reset value must be lower than the high watermark value.",
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
	"ipam_email_addresses": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
			listvalidator.AlsoRequires(path.MatchRoot("use_ipam_email_addresses")),
		},
		MarkdownDescription: "The e-mail lists to which the appliance sends IPAM threshold alarm e-mail messages.",
	},
	"ipam_threshold_settings": schema.SingleNestedAttribute{
		Attributes: NetworktemplateIpamThresholdSettingsResourceSchemaAttributes,
		Computed:   true,
		Optional:   true,
		Validators: []validator.Object{
			objectvalidator.AlsoRequires(path.MatchRoot("use_ipam_threshold_settings")),
		},
		MarkdownDescription: "The IPAM Threshold settings for this network template.",
	},
	"ipam_trap_settings": schema.SingleNestedAttribute{
		Attributes: NetworktemplateIpamTrapSettingsResourceSchemaAttributes,
		Computed:   true,
		Optional:   true,
		Validators: []validator.Object{
			objectvalidator.AlsoRequires(path.MatchRoot("use_ipam_trap_settings")),
		},
		MarkdownDescription: "The IPAM Trap settings for this network template.",
	},
	"lease_scavenge_time": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(-1),
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_lease_scavenge_time")),
			int64validator.Between(86400, 2147472000),
		},
		MarkdownDescription: "An integer that specifies the period of time (in seconds) that frees and backs up leases remained in the database before they are automatically deleted. To disable lease scavenging, set the parameter to -1. The minimum positive value must be greater than 86400 seconds (1 day).",
	},
	"logic_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NetworktemplateLogicFilterRulesResourceSchemaAttributes,
		},
		Computed: true,
		Optional: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
			listvalidator.AlsoRequires(path.MatchRoot("use_logic_filter_rules")),
		},
		MarkdownDescription: "This field contains the logic filters to be applied on the this network template. This list corresponds to the match rules that are written to the dhcpd configuration file.",
	},
	"low_water_mark": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(0),
		Validators: []validator.Int64{
			int64validator.Between(1, 100),
		},
		MarkdownDescription: "The percentage of DHCP network usage below which the Infoblox appliance generates a syslog message and sends a warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100.",
	},
	"low_water_mark_reset": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(10),
		Validators: []validator.Int64{
			int64validator.Between(1, 100),
		},
		MarkdownDescription: "The percentage of DHCP network usage threshold below which network usage is not expected and may warrant your attention. When the low watermark is crossed, the Infoblox appliance generates a syslog message and sends a warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100. The low watermark reset value must be higher than the low watermark value.",
	},
	"members": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NetworktemplateMembersResourceSchemaAttributes,
		},
		Computed: true,
		Optional: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "A list of members or Microsoft (r) servers that serve DHCP for this network. All members in the array must be of the same type. The struct type must be indicated in each element, by setting the \"_struct\" member to the struct type.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of this network template.",
	},
	"netmask": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.Between(0, 32),
		},
		MarkdownDescription: "The netmask of the network in CIDR format.",
	},
	"nextserver": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_nextserver")),
		},
		MarkdownDescription: "The name in FQDN and/or IPv4 Address of the next server that the host needs to boot.",
	},
	"options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NetworktemplateOptionsResourceSchemaAttributes,
		},
		Computed: true,
		Optional: true,
		Default: listdefault.StaticValue(
			types.ListValueMust(
				types.ObjectType{AttrTypes: NetworktemplateOptionsAttrTypes},
				[]attr.Value{},
			),
		),
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
			listvalidator.AlsoRequires(path.MatchRoot("use_options")),
		},
		MarkdownDescription: "An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object.",
	},
	"pxe_lease_time": schema.Int64Attribute{
		Optional: true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_pxe_lease_time")),
			int64validator.Between(0, 399999999),
		},
		MarkdownDescription: "The PXE lease time value of a DHCP Network object. Some hosts use PXE (Preboot Execution Environment) to boot remotely from a server. To better manage your IP resources, set a different lease time for PXE boot requests. You can configure the DHCP server to allocate an IP address with a shorter lease time to hosts that send PXE boot requests, so IP addresses are not leased longer than necessary. A 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached.",
	},
	"range_templates": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The list of IP address range templates assigned to this network template object. When you create a network based on a network template object that contains range templates, the IP address ranges are created based on the associated IP address range templates.",
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
		Computed:            true,
		MarkdownDescription: "THe registry (RIR) that allocated the network address space.",
	},
	"rir_organization": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The RIR organization assoicated with the network.",
	},
	"rir_registration_action": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Default:  stringdefault.StaticString("NONE"),
		Validators: []validator.String{
			stringvalidator.OneOf("CREATE", "NONE"),
		},
		MarkdownDescription: "The RIR registration action.",
	},
	"rir_registration_status": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Default:  stringdefault.StaticString("NOT_REGISTERED"),
		Validators: []validator.String{
			stringvalidator.OneOf("NOT_REGISTERED", "REGISTERED"),
		},
		MarkdownDescription: "The registration status of the network in RIR.",
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
	"use_ipam_email_addresses": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: ipam_email_addresses",
	},
	"use_ipam_threshold_settings": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: ipam_threshold_settings",
	},
	"use_ipam_trap_settings": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: ipam_trap_settings",
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
}

func (m *NetworktemplateModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.Networktemplate {
	if m == nil {
		return nil
	}
	to := &ipam.Networktemplate{
		AllowAnyNetmask:                flex.ExpandBoolPointer(m.AllowAnyNetmask),
		Authority:                      flex.ExpandBoolPointer(m.Authority),
		AutoCreateReversezone:          flex.ExpandBoolPointer(m.AutoCreateReversezone),
		Bootfile:                       flex.ExpandStringPointer(m.Bootfile),
		Bootserver:                     flex.ExpandStringPointer(m.Bootserver),
		CloudApiCompatible:             flex.ExpandBoolPointer(m.CloudApiCompatible),
		Comment:                        flex.ExpandStringPointer(m.Comment),
		DdnsDomainname:                 flex.ExpandStringPointer(m.DdnsDomainname.StringValue),
		DdnsGenerateHostname:           flex.ExpandBoolPointer(m.DdnsGenerateHostname),
		DdnsServerAlwaysUpdates:        flex.ExpandBoolPointer(m.DdnsServerAlwaysUpdates),
		DdnsTtl:                        flex.ExpandInt64Pointer(m.DdnsTtl),
		DdnsUpdateFixedAddresses:       flex.ExpandBoolPointer(m.DdnsUpdateFixedAddresses),
		DdnsUseOption81:                flex.ExpandBoolPointer(m.DdnsUseOption81),
		DelegatedMember:                ExpandNetworktemplateDelegatedMember(ctx, m.DelegatedMember, diags),
		DenyBootp:                      flex.ExpandBoolPointer(m.DenyBootp),
		EmailList:                      flex.ExpandFrameworkListString(ctx, m.EmailList, diags),
		EnableDdns:                     flex.ExpandBoolPointer(m.EnableDdns),
		EnableDhcpThresholds:           flex.ExpandBoolPointer(m.EnableDhcpThresholds),
		EnableEmailWarnings:            flex.ExpandBoolPointer(m.EnableEmailWarnings),
		EnablePxeLeaseTime:             flex.ExpandBoolPointer(m.EnablePxeLeaseTime),
		EnableSnmpWarnings:             flex.ExpandBoolPointer(m.EnableSnmpWarnings),
		ExtAttrs:                       ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		FixedAddressTemplates:          flex.ExpandFrameworkListString(ctx, m.FixedAddressTemplates, diags),
		HighWaterMark:                  flex.ExpandInt64Pointer(m.HighWaterMark),
		HighWaterMarkReset:             flex.ExpandInt64Pointer(m.HighWaterMarkReset),
		IgnoreDhcpOptionListRequest:    flex.ExpandBoolPointer(m.IgnoreDhcpOptionListRequest),
		IpamEmailAddresses:             flex.ExpandFrameworkListString(ctx, m.IpamEmailAddresses, diags),
		IpamThresholdSettings:          ExpandNetworktemplateIpamThresholdSettings(ctx, m.IpamThresholdSettings, diags),
		IpamTrapSettings:               ExpandNetworktemplateIpamTrapSettings(ctx, m.IpamTrapSettings, diags),
		LeaseScavengeTime:              flex.ExpandInt64Pointer(m.LeaseScavengeTime),
		LogicFilterRules:               flex.ExpandFrameworkListNestedBlock(ctx, m.LogicFilterRules, diags, ExpandNetworktemplateLogicFilterRules),
		LowWaterMark:                   flex.ExpandInt64Pointer(m.LowWaterMark),
		LowWaterMarkReset:              flex.ExpandInt64Pointer(m.LowWaterMarkReset),
		Members:                        flex.ExpandFrameworkListNestedBlock(ctx, m.Members, diags, ExpandNetworktemplateMembers),
		Name:                           flex.ExpandStringPointer(m.Name),
		Netmask:                        flex.ExpandInt64Pointer(m.Netmask),
		Nextserver:                     flex.ExpandStringPointer(m.Nextserver),
		Options:                        flex.ExpandFrameworkListNestedBlock(ctx, m.Options, diags, ExpandNetworktemplateOptions),
		PxeLeaseTime:                   flex.ExpandInt64Pointer(m.PxeLeaseTime),
		RangeTemplates:                 flex.ExpandFrameworkListString(ctx, m.RangeTemplates, diags),
		RecycleLeases:                  flex.ExpandBoolPointer(m.RecycleLeases),
		RirOrganization:                flex.ExpandStringPointer(m.RirOrganization),
		RirRegistrationAction:          flex.ExpandStringPointer(m.RirRegistrationAction),
		RirRegistrationStatus:          flex.ExpandStringPointer(m.RirRegistrationStatus),
		SendRirRequest:                 flex.ExpandBoolPointer(m.SendRirRequest),
		UpdateDnsOnLeaseRenewal:        flex.ExpandBoolPointer(m.UpdateDnsOnLeaseRenewal),
		UseAuthority:                   flex.ExpandBoolPointer(m.UseAuthority),
		UseBootfile:                    flex.ExpandBoolPointer(m.UseBootfile),
		UseBootserver:                  flex.ExpandBoolPointer(m.UseBootserver),
		UseDdnsDomainname:              flex.ExpandBoolPointer(m.UseDdnsDomainname),
		UseDdnsGenerateHostname:        flex.ExpandBoolPointer(m.UseDdnsGenerateHostname),
		UseDdnsTtl:                     flex.ExpandBoolPointer(m.UseDdnsTtl),
		UseDdnsUpdateFixedAddresses:    flex.ExpandBoolPointer(m.UseDdnsUpdateFixedAddresses),
		UseDdnsUseOption81:             flex.ExpandBoolPointer(m.UseDdnsUseOption81),
		UseDenyBootp:                   flex.ExpandBoolPointer(m.UseDenyBootp),
		UseEmailList:                   flex.ExpandBoolPointer(m.UseEmailList),
		UseEnableDdns:                  flex.ExpandBoolPointer(m.UseEnableDdns),
		UseEnableDhcpThresholds:        flex.ExpandBoolPointer(m.UseEnableDhcpThresholds),
		UseIgnoreDhcpOptionListRequest: flex.ExpandBoolPointer(m.UseIgnoreDhcpOptionListRequest),
		UseIpamEmailAddresses:          flex.ExpandBoolPointer(m.UseIpamEmailAddresses),
		UseIpamThresholdSettings:       flex.ExpandBoolPointer(m.UseIpamThresholdSettings),
		UseIpamTrapSettings:            flex.ExpandBoolPointer(m.UseIpamTrapSettings),
		UseLeaseScavengeTime:           flex.ExpandBoolPointer(m.UseLeaseScavengeTime),
		UseLogicFilterRules:            flex.ExpandBoolPointer(m.UseLogicFilterRules),
		UseNextserver:                  flex.ExpandBoolPointer(m.UseNextserver),
		UseOptions:                     flex.ExpandBoolPointer(m.UseOptions),
		UsePxeLeaseTime:                flex.ExpandBoolPointer(m.UsePxeLeaseTime),
		UseRecycleLeases:               flex.ExpandBoolPointer(m.UseRecycleLeases),
		UseUpdateDnsOnLeaseRenewal:     flex.ExpandBoolPointer(m.UseUpdateDnsOnLeaseRenewal),
	}
	return to
}

func FlattenNetworktemplate(ctx context.Context, from *ipam.Networktemplate, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworktemplateAttrTypes)
	}
	m := NetworktemplateModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, NetworktemplateAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NetworktemplateModel) Flatten(ctx context.Context, from *ipam.Networktemplate, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NetworktemplateModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.AllowAnyNetmask = types.BoolPointerValue(from.AllowAnyNetmask)
	m.Authority = types.BoolPointerValue(from.Authority)
	m.AutoCreateReversezone = types.BoolPointerValue(from.AutoCreateReversezone)
	m.Bootfile = flex.FlattenStringPointer(from.Bootfile)
	m.Bootserver = flex.FlattenStringPointer(from.Bootserver)
	m.CloudApiCompatible = types.BoolPointerValue(from.CloudApiCompatible)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.DdnsDomainname.StringValue = flex.FlattenStringPointer(from.DdnsDomainname)
	m.DdnsGenerateHostname = types.BoolPointerValue(from.DdnsGenerateHostname)
	m.DdnsServerAlwaysUpdates = types.BoolPointerValue(from.DdnsServerAlwaysUpdates)
	m.DdnsTtl = flex.FlattenInt64Pointer(from.DdnsTtl)
	m.DdnsUpdateFixedAddresses = types.BoolPointerValue(from.DdnsUpdateFixedAddresses)
	m.DdnsUseOption81 = types.BoolPointerValue(from.DdnsUseOption81)
	m.DelegatedMember = FlattenNetworktemplateDelegatedMember(ctx, from.DelegatedMember, diags)
	m.DenyBootp = types.BoolPointerValue(from.DenyBootp)
	m.EmailList = flex.FlattenFrameworkUnorderedList(ctx, types.StringType, from.EmailList, diags)
	m.EnableDdns = types.BoolPointerValue(from.EnableDdns)
	m.EnableDhcpThresholds = types.BoolPointerValue(from.EnableDhcpThresholds)
	m.EnableEmailWarnings = types.BoolPointerValue(from.EnableEmailWarnings)
	m.EnablePxeLeaseTime = types.BoolPointerValue(from.EnablePxeLeaseTime)
	m.EnableSnmpWarnings = types.BoolPointerValue(from.EnableSnmpWarnings)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.FixedAddressTemplates = flex.FlattenFrameworkListString(ctx, from.FixedAddressTemplates, diags)
	m.HighWaterMark = flex.FlattenInt64Pointer(from.HighWaterMark)
	m.HighWaterMarkReset = flex.FlattenInt64Pointer(from.HighWaterMarkReset)
	m.IgnoreDhcpOptionListRequest = types.BoolPointerValue(from.IgnoreDhcpOptionListRequest)
	m.IpamEmailAddresses = flex.FlattenFrameworkListString(ctx, from.IpamEmailAddresses, diags)
	m.IpamThresholdSettings = FlattenNetworktemplateIpamThresholdSettings(ctx, from.IpamThresholdSettings, diags)
	m.IpamTrapSettings = FlattenNetworktemplateIpamTrapSettings(ctx, from.IpamTrapSettings, diags)
	m.LeaseScavengeTime = flex.FlattenInt64Pointer(from.LeaseScavengeTime)
	m.LogicFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.LogicFilterRules, NetworktemplateLogicFilterRulesAttrTypes, diags, FlattenNetworktemplateLogicFilterRules)
	m.LowWaterMark = flex.FlattenInt64Pointer(from.LowWaterMark)
	m.LowWaterMarkReset = flex.FlattenInt64Pointer(from.LowWaterMarkReset)
	m.Members = flex.FlattenFrameworkListNestedBlock(ctx, from.Members, NetworktemplateMembersAttrTypes, diags, FlattenNetworktemplateMembers)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Netmask = flex.FlattenInt64Pointer(from.Netmask)
	m.Nextserver = flex.FlattenStringPointer(from.Nextserver)
	planOptions := m.Options
	m.Options = flex.FlattenFrameworkListNestedBlock(ctx, from.Options, NetworktemplateOptionsAttrTypes, diags, FlattenNetworktemplateOptions)
	if !planOptions.IsUnknown() {
		reOrderedOptions, diags := utils.ReorderAndFilterDHCPOptions(ctx, planOptions, m.Options)
		if !diags.HasError() {
			m.Options = reOrderedOptions.(basetypes.ListValue)
		}
	}
	m.PxeLeaseTime = flex.FlattenInt64Pointer(from.PxeLeaseTime)
	m.RangeTemplates = flex.FlattenFrameworkListString(ctx, from.RangeTemplates, diags)
	m.RecycleLeases = types.BoolPointerValue(from.RecycleLeases)
	m.Rir = flex.FlattenStringPointer(from.Rir)
	m.RirOrganization = flex.FlattenStringPointer(from.RirOrganization)
	m.RirRegistrationAction = flex.FlattenStringPointer(from.RirRegistrationAction)
	m.RirRegistrationStatus = flex.FlattenStringPointer(from.RirRegistrationStatus)
	m.SendRirRequest = types.BoolPointerValue(from.SendRirRequest)
	m.UpdateDnsOnLeaseRenewal = types.BoolPointerValue(from.UpdateDnsOnLeaseRenewal)
	m.UseAuthority = types.BoolPointerValue(from.UseAuthority)
	m.UseBootfile = types.BoolPointerValue(from.UseBootfile)
	m.UseBootserver = types.BoolPointerValue(from.UseBootserver)
	m.UseDdnsDomainname = types.BoolPointerValue(from.UseDdnsDomainname)
	m.UseDdnsGenerateHostname = types.BoolPointerValue(from.UseDdnsGenerateHostname)
	m.UseDdnsTtl = types.BoolPointerValue(from.UseDdnsTtl)
	m.UseDdnsUpdateFixedAddresses = types.BoolPointerValue(from.UseDdnsUpdateFixedAddresses)
	m.UseDdnsUseOption81 = types.BoolPointerValue(from.UseDdnsUseOption81)
	m.UseDenyBootp = types.BoolPointerValue(from.UseDenyBootp)
	m.UseEmailList = types.BoolPointerValue(from.UseEmailList)
	m.UseEnableDdns = types.BoolPointerValue(from.UseEnableDdns)
	m.UseEnableDhcpThresholds = types.BoolPointerValue(from.UseEnableDhcpThresholds)
	m.UseIgnoreDhcpOptionListRequest = types.BoolPointerValue(from.UseIgnoreDhcpOptionListRequest)
	m.UseIpamEmailAddresses = types.BoolPointerValue(from.UseIpamEmailAddresses)
	m.UseIpamThresholdSettings = types.BoolPointerValue(from.UseIpamThresholdSettings)
	m.UseIpamTrapSettings = types.BoolPointerValue(from.UseIpamTrapSettings)
	m.UseLeaseScavengeTime = types.BoolPointerValue(from.UseLeaseScavengeTime)
	m.UseLogicFilterRules = types.BoolPointerValue(from.UseLogicFilterRules)
	m.UseNextserver = types.BoolPointerValue(from.UseNextserver)
	m.UseOptions = types.BoolPointerValue(from.UseOptions)
	m.UsePxeLeaseTime = types.BoolPointerValue(from.UsePxeLeaseTime)
	m.UseRecycleLeases = types.BoolPointerValue(from.UseRecycleLeases)
	m.UseUpdateDnsOnLeaseRenewal = types.BoolPointerValue(from.UseUpdateDnsOnLeaseRenewal)
}
