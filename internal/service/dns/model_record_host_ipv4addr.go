package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/boolvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type RecordHostIpv4addrModel struct {
	Ref                             types.String        `tfsdk:"ref"`
	Uuid                            types.String        `tfsdk:"uuid"`
	Bootfile                        types.String        `tfsdk:"bootfile"`
	Bootserver                      types.String        `tfsdk:"bootserver"`
	ConfigureForDhcp                types.Bool          `tfsdk:"configure_for_dhcp"`
	DenyBootp                       types.Bool          `tfsdk:"deny_bootp"`
	DiscoverNowStatus               types.String        `tfsdk:"discover_now_status"`
	DiscoveredData                  types.Object        `tfsdk:"discovered_data"`
	EnablePxeLeaseTime              types.Bool          `tfsdk:"enable_pxe_lease_time"`
	Host                            types.String        `tfsdk:"host"`
	IgnoreClientRequestedOptions    types.Bool          `tfsdk:"ignore_client_requested_options"`
	Ipv4addr                        iptypes.IPv4Address `tfsdk:"ipv4addr"`
	FuncCall                        types.Object        `tfsdk:"func_call"`
	IsInvalidMac                    types.Bool          `tfsdk:"is_invalid_mac"`
	LastQueried                     types.Int64         `tfsdk:"last_queried"`
	LogicFilterRules                types.List          `tfsdk:"logic_filter_rules"`
	Mac                             types.String        `tfsdk:"mac"`
	MatchClient                     types.String        `tfsdk:"match_client"`
	MsAdUserData                    types.Object        `tfsdk:"ms_ad_user_data"`
	Network                         types.String        `tfsdk:"network"`
	NetworkView                     types.String        `tfsdk:"network_view"`
	Nextserver                      types.String        `tfsdk:"nextserver"`
	Options                         types.List          `tfsdk:"options"`
	PxeLeaseTime                    types.Int64         `tfsdk:"pxe_lease_time"`
	ReservedInterface               types.String        `tfsdk:"reserved_interface"`
	UseBootfile                     types.Bool          `tfsdk:"use_bootfile"`
	UseBootserver                   types.Bool          `tfsdk:"use_bootserver"`
	UseDenyBootp                    types.Bool          `tfsdk:"use_deny_bootp"`
	UseForEaInheritance             types.Bool          `tfsdk:"use_for_ea_inheritance"`
	UseIgnoreClientRequestedOptions types.Bool          `tfsdk:"use_ignore_client_requested_options"`
	UseLogicFilterRules             types.Bool          `tfsdk:"use_logic_filter_rules"`
	UseNextserver                   types.Bool          `tfsdk:"use_nextserver"`
	UseOptions                      types.Bool          `tfsdk:"use_options"`
	UsePxeLeaseTime                 types.Bool          `tfsdk:"use_pxe_lease_time"`
}

var RecordHostIpv4addrAttrTypes = map[string]attr.Type{
	"ref":                                 types.StringType,
	"uuid":                                types.StringType,
	"bootfile":                            types.StringType,
	"bootserver":                          types.StringType,
	"configure_for_dhcp":                  types.BoolType,
	"deny_bootp":                          types.BoolType,
	"discover_now_status":                 types.StringType,
	"discovered_data":                     types.ObjectType{AttrTypes: RecordHostIpv4addrDiscoveredDataAttrTypes},
	"enable_pxe_lease_time":               types.BoolType,
	"host":                                types.StringType,
	"ignore_client_requested_options":     types.BoolType,
	"ipv4addr":                            iptypes.IPv4AddressType{},
	"func_call":                           types.ObjectType{AttrTypes: FuncCallAttrTypes},
	"is_invalid_mac":                      types.BoolType,
	"last_queried":                        types.Int64Type,
	"logic_filter_rules":                  types.ListType{ElemType: types.ObjectType{AttrTypes: RecordHostIpv4addrLogicFilterRulesAttrTypes}},
	"mac":                                 types.StringType,
	"match_client":                        types.StringType,
	"ms_ad_user_data":                     types.ObjectType{AttrTypes: RecordHostIpv4addrMsAdUserDataAttrTypes},
	"network":                             types.StringType,
	"network_view":                        types.StringType,
	"nextserver":                          types.StringType,
	"options":                             types.ListType{ElemType: types.ObjectType{AttrTypes: RecordHostIpv4addrOptionsAttrTypes}},
	"pxe_lease_time":                      types.Int64Type,
	"reserved_interface":                  types.StringType,
	"use_bootfile":                        types.BoolType,
	"use_bootserver":                      types.BoolType,
	"use_deny_bootp":                      types.BoolType,
	"use_for_ea_inheritance":              types.BoolType,
	"use_ignore_client_requested_options": types.BoolType,
	"use_logic_filter_rules":              types.BoolType,
	"use_nextserver":                      types.BoolType,
	"use_options":                         types.BoolType,
	"use_pxe_lease_time":                  types.BoolType,
}

var RecordHostIpv4addrResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The uuid to the object.",
	},
	"bootfile": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The name of the boot file the client must download.",
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_bootfile")),
		},
	},
	"bootserver": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_bootserver")),
			customvalidator.IsValidFQDN(),
		},
		MarkdownDescription: "The IP address or hostname of the boot file server where the boot file is stored.",
	},
	"configure_for_dhcp": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Set this to True to enable the DHCP configuration for this host address.",
	},
	"deny_bootp": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_deny_bootp")),
		},
		MarkdownDescription: "Set this to True to disable the BOOTP settings and deny BOOTP boot requests.",
	},
	"discover_now_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The discovery status of this Host Address.",
	},
	"discovered_data": schema.SingleNestedAttribute{
		Attributes: RecordHostIpv4addrDiscoveredDataResourceSchemaAttributes,
		Computed:   true,
	},
	"enable_pxe_lease_time": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Set this to True if you want the DHCP server to use a different lease time for PXE clients. You can specify the duration of time it takes a host to connect to a boot server, such as a TFTP server, and download the file it needs to boot. For example, set a longer lease time if the client downloads an OS (operating system) or configuration file, or set a shorter lease time if the client downloads only configuration changes. Enter the lease time for the preboot execution environment for hosts to boot remotely from a server.",
	},
	"host": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The host to which the host address belongs, in FQDN format. It is only present when the host address object is not returned as part of a host.",
	},
	"ignore_client_requested_options": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If this field is set to false, the appliance returns all DHCP options the client is eligible to receive, rather than only the list of options the client has requested.",
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
			objectvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("ipv4addr")),
		},
		MarkdownDescription: "Function call to be executed for Fixed Address",
	},
	"is_invalid_mac": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "This flag reflects whether the MAC address for this host address is invalid.",
	},
	"last_queried": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The time of the last DNS query in Epoch seconds format.",
	},
	"logic_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RecordHostIpv4addrLogicFilterRulesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_logic_filter_rules")),
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "This field contains the logic filters to be applied on the this host address. This list corresponds to the match rules that are written to the dhcpd configuration file.",
	},
	"mac": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The MAC address for this host address.",
	},
	"match_client": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Set this to 'MAC_ADDRESS' to assign the IP address to the selected host, provided that the MAC address of the requesting host matches the MAC address that you specify in the field. Set this to 'RESERVED' to reserve this particular IP address for future use, or if the IP address is statically configured on a system (the Infoblox server does not assign the address from a DHCP request).",
	},
	"ms_ad_user_data": schema.SingleNestedAttribute{
		Attributes: RecordHostIpv4addrMsAdUserDataResourceSchemaAttributes,
		Optional:   true,
	},
	"network": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The network of the host address, in FQDN/CIDR format.",
	},
	"network_view": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the network view in which the host address resides.",
	},
	"nextserver": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_nextserver")),
		},
		Computed:            true,
		MarkdownDescription: "The name in FQDN format and/or IPv4 Address of the next server that the host needs to boot.",
	},
	"options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RecordHostIpv4addrOptionsResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object.",
	},
	"pxe_lease_time": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The lease time for PXE clients, see *enable_pxe_lease_time* for more information.",
	},
	"reserved_interface": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The reference to the reserved interface to which the device belongs.",
	},
	"use_bootfile": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: bootfile",
	},
	"use_bootserver": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: bootserver",
	},
	"use_deny_bootp": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: deny_bootp",
	},
	"use_for_ea_inheritance": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Set this to True when using this host address for EA inheritance.",
	},
	"use_ignore_client_requested_options": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ignore_client_requested_options",
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

func ExpandRecordHostIpv4addr(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.RecordHostIpv4addr {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RecordHostIpv4addrModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *RecordHostIpv4addrModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.RecordHostIpv4addr {
	if m == nil {
		return nil
	}
	to := &dns.RecordHostIpv4addr{
		Ref:                             flex.ExpandStringPointer(m.Ref),
		Bootfile:                        flex.ExpandStringPointer(m.Bootfile),
		Bootserver:                      flex.ExpandStringPointer(m.Bootserver),
		ConfigureForDhcp:                flex.ExpandBoolPointer(m.ConfigureForDhcp),
		DenyBootp:                       flex.ExpandBoolPointer(m.DenyBootp),
		DiscoveredData:                  ExpandRecordHostIpv4addrDiscoveredData(ctx, m.DiscoveredData, diags),
		EnablePxeLeaseTime:              flex.ExpandBoolPointer(m.EnablePxeLeaseTime),
		IgnoreClientRequestedOptions:    flex.ExpandBoolPointer(m.IgnoreClientRequestedOptions),
		Ipv4addr:                        ExpandRecordHostIpv4addrIpv4addr(m.Ipv4addr),
		FuncCall:                        ExpandFuncCall(ctx, m.FuncCall, diags),
		LogicFilterRules:                flex.ExpandFrameworkListNestedBlock(ctx, m.LogicFilterRules, diags, ExpandRecordHostIpv4addrLogicFilterRules),
		Mac:                             flex.ExpandStringPointer(m.Mac),
		MatchClient:                     flex.ExpandStringPointer(m.MatchClient),
		MsAdUserData:                    ExpandRecordHostIpv4addrMsAdUserData(ctx, m.MsAdUserData, diags),
		Nextserver:                      flex.ExpandStringPointer(m.Nextserver),
		Options:                         flex.ExpandFrameworkListNestedBlock(ctx, m.Options, diags, ExpandRecordHostIpv4addrOptions),
		PxeLeaseTime:                    flex.ExpandInt64Pointer(m.PxeLeaseTime),
		ReservedInterface:               flex.ExpandStringPointerEmptyAsNil(m.ReservedInterface),
		UseBootfile:                     flex.ExpandBoolPointer(m.UseBootfile),
		UseBootserver:                   flex.ExpandBoolPointer(m.UseBootserver),
		UseDenyBootp:                    flex.ExpandBoolPointer(m.UseDenyBootp),
		UseForEaInheritance:             flex.ExpandBoolPointer(m.UseForEaInheritance),
		UseIgnoreClientRequestedOptions: flex.ExpandBoolPointer(m.UseIgnoreClientRequestedOptions),
		UseLogicFilterRules:             flex.ExpandBoolPointer(m.UseLogicFilterRules),
		UseNextserver:                   flex.ExpandBoolPointer(m.UseNextserver),
		UseOptions:                      flex.ExpandBoolPointer(m.UseOptions),
		UsePxeLeaseTime:                 flex.ExpandBoolPointer(m.UsePxeLeaseTime),
	}
	return to
}

func FlattenRecordHostIpv4addr(ctx context.Context, from *dns.RecordHostIpv4addr, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordHostIpv4addrAttrTypes)
	}
	m := RecordHostIpv4addrModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RecordHostIpv4addrAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RecordHostIpv4addrModel) Flatten(ctx context.Context, from *dns.RecordHostIpv4addr, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RecordHostIpv4addrModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Bootfile = flex.FlattenStringPointer(from.Bootfile)
	m.Bootserver = flex.FlattenStringPointer(from.Bootserver)
	m.ConfigureForDhcp = types.BoolPointerValue(from.ConfigureForDhcp)
	m.DenyBootp = types.BoolPointerValue(from.DenyBootp)
	m.DiscoverNowStatus = flex.FlattenStringPointer(from.DiscoverNowStatus)
	m.DiscoveredData = FlattenRecordHostIpv4addrDiscoveredData(ctx, from.DiscoveredData, diags)
	m.EnablePxeLeaseTime = types.BoolPointerValue(from.EnablePxeLeaseTime)
	m.Host = flex.FlattenStringPointer(from.Host)
	m.IgnoreClientRequestedOptions = types.BoolPointerValue(from.IgnoreClientRequestedOptions)
	m.Ipv4addr = FlattenRecordHostIpv4addrIpv4addr(from.Ipv4addr)
	m.IsInvalidMac = types.BoolPointerValue(from.IsInvalidMac)
	m.LastQueried = flex.FlattenInt64Pointer(from.LastQueried)
	m.LogicFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.LogicFilterRules, RecordHostIpv4addrLogicFilterRulesAttrTypes, diags, FlattenRecordHostIpv4addrLogicFilterRules)
	m.Mac = flex.FlattenStringPointer(from.Mac)
	m.MatchClient = flex.FlattenStringPointer(from.MatchClient)
	m.MsAdUserData = FlattenRecordHostIpv4addrMsAdUserData(ctx, from.MsAdUserData, diags)
	m.Network = flex.FlattenStringPointer(from.Network)
	m.NetworkView = flex.FlattenStringPointer(from.NetworkView)
	m.Nextserver = flex.FlattenStringPointer(from.Nextserver)
	m.Options = flex.FlattenFrameworkListNestedBlock(ctx, from.Options, RecordHostIpv4addrOptionsAttrTypes, diags, FlattenRecordHostIpv4addrOptions)
	m.PxeLeaseTime = flex.FlattenInt64Pointer(from.PxeLeaseTime)
	m.ReservedInterface = flex.FlattenStringPointerNilAsNotEmpty(from.ReservedInterface)
	m.UseBootfile = types.BoolPointerValue(from.UseBootfile)
	m.UseBootserver = types.BoolPointerValue(from.UseBootserver)
	m.UseDenyBootp = types.BoolPointerValue(from.UseDenyBootp)
	m.UseForEaInheritance = types.BoolPointerValue(from.UseForEaInheritance)
	m.UseIgnoreClientRequestedOptions = types.BoolPointerValue(from.UseIgnoreClientRequestedOptions)
	m.UseLogicFilterRules = types.BoolPointerValue(from.UseLogicFilterRules)
	m.UseNextserver = types.BoolPointerValue(from.UseNextserver)
	m.UseOptions = types.BoolPointerValue(from.UseOptions)
	m.UsePxeLeaseTime = types.BoolPointerValue(from.UsePxeLeaseTime)

	if m.FuncCall.IsNull() || m.FuncCall.IsUnknown() {
		m.FuncCall = FlattenFuncCall(ctx, from.FuncCall, diags)
	}
}

func ExpandRecordHostIpv4addrIpv4addr(ipv4addr iptypes.IPv4Address) *dns.RecordHostIpv4addrIpv4addr {
	if ipv4addr.IsNull() {
		return &dns.RecordHostIpv4addrIpv4addr{}
	}
	var m dns.RecordHostIpv4addrIpv4addr
	m.String = flex.ExpandIPv4Address(ipv4addr)

	return &m
}

func FlattenRecordHostIpv4addrIpv4addr(from *dns.RecordHostIpv4addrIpv4addr) iptypes.IPv4Address {
	if from.String == nil {
		return iptypes.NewIPv4AddressNull()
	}
	m := flex.FlattenIPv4Address(from.String)
	return m
}
