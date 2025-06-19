package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dhcp"

	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/flex"
)

type FixedaddresstemplateModel struct {
	Ref                            types.String `tfsdk:"ref"`
	Bootfile                       types.String `tfsdk:"bootfile"`
	Bootserver                     types.String `tfsdk:"bootserver"`
	Comment                        types.String `tfsdk:"comment"`
	DdnsDomainname                 types.String `tfsdk:"ddns_domainname"`
	DdnsHostname                   types.String `tfsdk:"ddns_hostname"`
	DenyBootp                      types.Bool   `tfsdk:"deny_bootp"`
	EnableDdns                     types.Bool   `tfsdk:"enable_ddns"`
	EnablePxeLeaseTime             types.Bool   `tfsdk:"enable_pxe_lease_time"`
	ExtAttrs                       types.Map    `tfsdk:"extattrs"`
	IgnoreDhcpOptionListRequest    types.Bool   `tfsdk:"ignore_dhcp_option_list_request"`
	LogicFilterRules               types.List   `tfsdk:"logic_filter_rules"`
	Name                           types.String `tfsdk:"name"`
	Nextserver                     types.String `tfsdk:"nextserver"`
	NumberOfAddresses              types.Int64  `tfsdk:"number_of_addresses"`
	Offset                         types.Int64  `tfsdk:"offset"`
	Options                        types.List   `tfsdk:"options"`
	PxeLeaseTime                   types.Int64  `tfsdk:"pxe_lease_time"`
	UseBootfile                    types.Bool   `tfsdk:"use_bootfile"`
	UseBootserver                  types.Bool   `tfsdk:"use_bootserver"`
	UseDdnsDomainname              types.Bool   `tfsdk:"use_ddns_domainname"`
	UseDenyBootp                   types.Bool   `tfsdk:"use_deny_bootp"`
	UseEnableDdns                  types.Bool   `tfsdk:"use_enable_ddns"`
	UseIgnoreDhcpOptionListRequest types.Bool   `tfsdk:"use_ignore_dhcp_option_list_request"`
	UseLogicFilterRules            types.Bool   `tfsdk:"use_logic_filter_rules"`
	UseNextserver                  types.Bool   `tfsdk:"use_nextserver"`
	UseOptions                     types.Bool   `tfsdk:"use_options"`
	UsePxeLeaseTime                types.Bool   `tfsdk:"use_pxe_lease_time"`
}

var FixedaddresstemplateAttrTypes = map[string]attr.Type{
	"_ref":                                types.StringType,
	"bootfile":                            types.StringType,
	"bootserver":                          types.StringType,
	"comment":                             types.StringType,
	"ddns_domainname":                     types.StringType,
	"ddns_hostname":                       types.StringType,
	"deny_bootp":                          types.BoolType,
	"enable_ddns":                         types.BoolType,
	"enable_pxe_lease_time":               types.BoolType,
	"extattrs":                            types.MapType{ElemType: types.StringType},
	"ignore_dhcp_option_list_request":     types.BoolType,
	"logic_filter_rules":                  types.ListType{ElemType: types.ObjectType{AttrTypes: FixedaddresstemplateLogicFilterRulesAttrTypes}},
	"name":                                types.StringType,
	"nextserver":                          types.StringType,
	"number_of_addresses":                 types.Int64Type,
	"offset":                              types.Int64Type,
	"options":                             types.ListType{ElemType: types.ObjectType{AttrTypes: FixedaddresstemplateOptionsAttrTypes}},
	"pxe_lease_time":                      types.Int64Type,
	"use_bootfile":                        types.BoolType,
	"use_bootserver":                      types.BoolType,
	"use_ddns_domainname":                 types.BoolType,
	"use_deny_bootp":                      types.BoolType,
	"use_enable_ddns":                     types.BoolType,
	"use_ignore_dhcp_option_list_request": types.BoolType,
	"use_logic_filter_rules":              types.BoolType,
	"use_nextserver":                      types.BoolType,
	"use_options":                         types.BoolType,
	"use_pxe_lease_time":                  types.BoolType,
}

var FixedaddresstemplateResourceSchemaAttributes = map[string]schema.Attribute{
	"_ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"bootfile": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The boot file name for the fixed address. You can configure the DHCP server to support clients that use the boot file name option in their DHCPREQUEST messages.",
	},
	"bootserver": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The boot server address for the fixed address. You can specify the name and/or IP address of the boot server that the host needs to boot. The boot server IPv4 Address or name in FQDN format.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "A descriptive comment of a fixed address template object.",
	},
	"ddns_domainname": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The dynamic DNS domain name the appliance uses specifically for DDNS updates for this fixed address.",
	},
	"ddns_hostname": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The DDNS host name for this fixed address.",
	},
	"deny_bootp": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if BOOTP settings are disabled and BOOTP requests will be denied.",
	},
	"enable_ddns": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the DHCP server sends DDNS updates to DNS servers in the same Grid, and to external DNS servers.",
	},
	"enable_pxe_lease_time": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Set this to True if you want the DHCP server to use a different lease time for PXE clients.",
	},
	"extattrs": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.",
	},
	"ignore_dhcp_option_list_request": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If this field is set to False, the appliance returns all DHCP options the client is eligible to receive, rather than only the list of options the client has requested.",
	},
	"logic_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: FixedaddresstemplateLogicFilterRulesResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "This field contains the logic filters to be applied on this fixed address. This list corresponds to the match rules that are written to the dhcpd configuration file.",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of a fixed address template object.",
	},
	"nextserver": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name in FQDN and/or IPv4 Address format of the next server that the host needs to boot.",
	},
	"number_of_addresses": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The number of addresses for this fixed address.",
	},
	"offset": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The start address offset for this fixed address.",
	},
	"options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: FixedaddresstemplateOptionsResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object.",
	},
	"pxe_lease_time": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The PXE lease time value for a DHCP Fixed Address object. Some hosts use PXE (Preboot Execution Environment) to boot remotely from a server. To better manage your IP resources, set a different lease time for PXE boot requests. You can configure the DHCP server to allocate an IP address with a shorter lease time to hosts that send PXE boot requests, so IP addresses are not leased longer than necessary. A 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached.",
	},
	"use_bootfile": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: bootfile",
	},
	"use_bootserver": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: bootserver",
	},
	"use_ddns_domainname": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ddns_domainname",
	},
	"use_deny_bootp": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: deny_bootp",
	},
	"use_enable_ddns": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: enable_ddns",
	},
	"use_ignore_dhcp_option_list_request": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ignore_dhcp_option_list_request",
	},
	"use_logic_filter_rules": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: logic_filter_rules",
	},
	"use_nextserver": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: nextserver",
	},
	"use_options": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: options",
	},
	"use_pxe_lease_time": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: pxe_lease_time",
	},
}

func ExpandFixedaddresstemplate(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dhcp.Fixedaddresstemplate {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m FixedaddresstemplateModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *FixedaddresstemplateModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.Fixedaddresstemplate {
	if m == nil {
		return nil
	}
	to := &dhcp.Fixedaddresstemplate{
		Ref:                            flex.ExpandStringPointer(m.Ref),
		Bootfile:                       flex.ExpandStringPointer(m.Bootfile),
		Bootserver:                     flex.ExpandStringPointer(m.Bootserver),
		Comment:                        flex.ExpandStringPointer(m.Comment),
		DdnsDomainname:                 flex.ExpandStringPointer(m.DdnsDomainname),
		DdnsHostname:                   flex.ExpandStringPointer(m.DdnsHostname),
		DenyBootp:                      flex.ExpandBoolPointer(m.DenyBootp),
		EnableDdns:                     flex.ExpandBoolPointer(m.EnableDdns),
		EnablePxeLeaseTime:             flex.ExpandBoolPointer(m.EnablePxeLeaseTime),
		ExtAttrs:                       flex.ExpandFrameworkMapString(ctx, m.ExtAttrs, diags),
		IgnoreDhcpOptionListRequest:    flex.ExpandBoolPointer(m.IgnoreDhcpOptionListRequest),
		LogicFilterRules:               flex.ExpandFrameworkListNestedBlock(ctx, m.LogicFilterRules, diags, ExpandFixedaddresstemplateLogicFilterRules),
		Name:                           flex.ExpandStringPointer(m.Name),
		Nextserver:                     flex.ExpandStringPointer(m.Nextserver),
		NumberOfAddresses:              flex.ExpandInt64Pointer(m.NumberOfAddresses),
		Offset:                         flex.ExpandInt64Pointer(m.Offset),
		Options:                        flex.ExpandFrameworkListNestedBlock(ctx, m.Options, diags, ExpandFixedaddresstemplateOptions),
		PxeLeaseTime:                   flex.ExpandInt64Pointer(m.PxeLeaseTime),
		UseBootfile:                    flex.ExpandBoolPointer(m.UseBootfile),
		UseBootserver:                  flex.ExpandBoolPointer(m.UseBootserver),
		UseDdnsDomainname:              flex.ExpandBoolPointer(m.UseDdnsDomainname),
		UseDenyBootp:                   flex.ExpandBoolPointer(m.UseDenyBootp),
		UseEnableDdns:                  flex.ExpandBoolPointer(m.UseEnableDdns),
		UseIgnoreDhcpOptionListRequest: flex.ExpandBoolPointer(m.UseIgnoreDhcpOptionListRequest),
		UseLogicFilterRules:            flex.ExpandBoolPointer(m.UseLogicFilterRules),
		UseNextserver:                  flex.ExpandBoolPointer(m.UseNextserver),
		UseOptions:                     flex.ExpandBoolPointer(m.UseOptions),
		UsePxeLeaseTime:                flex.ExpandBoolPointer(m.UsePxeLeaseTime),
	}
	return to
}

func FlattenFixedaddresstemplate(ctx context.Context, from *dhcp.Fixedaddresstemplate, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(FixedaddresstemplateAttrTypes)
	}
	m := FixedaddresstemplateModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrs = m.ExtAttrsAll
	t, d := types.ObjectValueFrom(ctx, FixedaddresstemplateAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *FixedaddresstemplateModel) Flatten(ctx context.Context, from *dhcp.Fixedaddresstemplate, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = FixedaddresstemplateModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Bootfile = flex.FlattenStringPointer(from.Bootfile)
	m.Bootserver = flex.FlattenStringPointer(from.Bootserver)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.DdnsDomainname = flex.FlattenStringPointer(from.DdnsDomainname)
	m.DdnsHostname = flex.FlattenStringPointer(from.DdnsHostname)
	m.DenyBootp = types.BoolPointerValue(from.DenyBootp)
	m.EnableDdns = types.BoolPointerValue(from.EnableDdns)
	m.EnablePxeLeaseTime = types.BoolPointerValue(from.EnablePxeLeaseTime)
	m.ExtAttrs = flex.FlattenFrameworkMapString(ctx, from.ExtAttrs, diags)
	m.IgnoreDhcpOptionListRequest = types.BoolPointerValue(from.IgnoreDhcpOptionListRequest)
	m.LogicFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.LogicFilterRules, FixedaddresstemplateLogicFilterRulesAttrTypes, diags, FlattenFixedaddresstemplateLogicFilterRules)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Nextserver = flex.FlattenStringPointer(from.Nextserver)
	m.NumberOfAddresses = flex.FlattenInt64Pointer(from.NumberOfAddresses)
	m.Offset = flex.FlattenInt64Pointer(from.Offset)
	m.Options = flex.FlattenFrameworkListNestedBlock(ctx, from.Options, FixedaddresstemplateOptionsAttrTypes, diags, FlattenFixedaddresstemplateOptions)
	m.PxeLeaseTime = flex.FlattenInt64Pointer(from.PxeLeaseTime)
	m.UseBootfile = types.BoolPointerValue(from.UseBootfile)
	m.UseBootserver = types.BoolPointerValue(from.UseBootserver)
	m.UseDdnsDomainname = types.BoolPointerValue(from.UseDdnsDomainname)
	m.UseDenyBootp = types.BoolPointerValue(from.UseDenyBootp)
	m.UseEnableDdns = types.BoolPointerValue(from.UseEnableDdns)
	m.UseIgnoreDhcpOptionListRequest = types.BoolPointerValue(from.UseIgnoreDhcpOptionListRequest)
	m.UseLogicFilterRules = types.BoolPointerValue(from.UseLogicFilterRules)
	m.UseNextserver = types.BoolPointerValue(from.UseNextserver)
	m.UseOptions = types.BoolPointerValue(from.UseOptions)
	m.UsePxeLeaseTime = types.BoolPointerValue(from.UsePxeLeaseTime)
}
