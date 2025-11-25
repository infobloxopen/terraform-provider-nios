package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/cidrtypes"
	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/boolvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	planmodifiers "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/immutable"
	importmod "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/import"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type Ipv6rangeModel struct {
	Ref                              types.String         `tfsdk:"ref"`
	AddressType                      types.String         `tfsdk:"address_type"`
	CloudInfo                        types.Object         `tfsdk:"cloud_info"`
	Comment                          types.String         `tfsdk:"comment"`
	Disable                          types.Bool           `tfsdk:"disable"`
	DiscoverNowStatus                types.String         `tfsdk:"discover_now_status"`
	DiscoveryBasicPollSettings       types.Object         `tfsdk:"discovery_basic_poll_settings"`
	DiscoveryBlackoutSetting         types.Object         `tfsdk:"discovery_blackout_setting"`
	DiscoveryMember                  types.String         `tfsdk:"discovery_member"`
	EnableDiscovery                  types.Bool           `tfsdk:"enable_discovery"`
	EnableImmediateDiscovery         types.Bool           `tfsdk:"enable_immediate_discovery"`
	EndAddr                          iptypes.IPv6Address  `tfsdk:"end_addr"`
	EndpointSources                  types.List           `tfsdk:"endpoint_sources"`
	Exclude                          types.List           `tfsdk:"exclude"`
	ExtAttrs                         types.Map            `tfsdk:"extattrs"`
	ExtAttrsAll                      types.Map            `tfsdk:"extattrs_all"`
	Ipv6EndPrefix                    iptypes.IPv6Address  `tfsdk:"ipv6_end_prefix"`
	Ipv6PrefixBits                   types.Int64          `tfsdk:"ipv6_prefix_bits"`
	Ipv6StartPrefix                  iptypes.IPv6Address  `tfsdk:"ipv6_start_prefix"`
	LogicFilterRules                 types.List           `tfsdk:"logic_filter_rules"`
	Member                           types.Object         `tfsdk:"member"`
	Name                             types.String         `tfsdk:"name"`
	Network                          cidrtypes.IPv6Prefix `tfsdk:"network"`
	NetworkView                      types.String         `tfsdk:"network_view"`
	OptionFilterRules                types.List           `tfsdk:"option_filter_rules"`
	PortControlBlackoutSetting       types.Object         `tfsdk:"port_control_blackout_setting"`
	RecycleLeases                    types.Bool           `tfsdk:"recycle_leases"`
	RestartIfNeeded                  types.Bool           `tfsdk:"restart_if_needed"`
	SamePortControlDiscoveryBlackout types.Bool           `tfsdk:"same_port_control_discovery_blackout"`
	ServerAssociationType            types.String         `tfsdk:"server_association_type"`
	StartAddr                        iptypes.IPv6Address  `tfsdk:"start_addr"`
	SubscribeSettings                types.Object         `tfsdk:"subscribe_settings"`
	Template                         types.String         `tfsdk:"template"`
	UseBlackoutSetting               types.Bool           `tfsdk:"use_blackout_setting"`
	UseDiscoveryBasicPollingSettings types.Bool           `tfsdk:"use_discovery_basic_polling_settings"`
	UseEnableDiscovery               types.Bool           `tfsdk:"use_enable_discovery"`
	UseLogicFilterRules              types.Bool           `tfsdk:"use_logic_filter_rules"`
	UseRecycleLeases                 types.Bool           `tfsdk:"use_recycle_leases"`
	UseSubscribeSettings             types.Bool           `tfsdk:"use_subscribe_settings"`
}

var Ipv6rangeAttrTypes = map[string]attr.Type{
	"ref":                                  types.StringType,
	"address_type":                         types.StringType,
	"cloud_info":                           types.ObjectType{AttrTypes: Ipv6rangeCloudInfoAttrTypes},
	"comment":                              types.StringType,
	"disable":                              types.BoolType,
	"discover_now_status":                  types.StringType,
	"discovery_basic_poll_settings":        types.ObjectType{AttrTypes: Ipv6rangeDiscoveryBasicPollSettingsAttrTypes},
	"discovery_blackout_setting":           types.ObjectType{AttrTypes: Ipv6rangeDiscoveryBlackoutSettingAttrTypes},
	"discovery_member":                     types.StringType,
	"enable_discovery":                     types.BoolType,
	"enable_immediate_discovery":           types.BoolType,
	"end_addr":                             iptypes.IPv6AddressType{},
	"endpoint_sources":                     types.ListType{ElemType: types.StringType},
	"exclude":                              types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6rangeExcludeAttrTypes}},
	"extattrs":                             types.MapType{ElemType: types.StringType},
	"extattrs_all":                         types.MapType{ElemType: types.StringType},
	"ipv6_end_prefix":                      iptypes.IPv6AddressType{},
	"ipv6_prefix_bits":                     types.Int64Type,
	"ipv6_start_prefix":                    iptypes.IPv6AddressType{},
	"logic_filter_rules":                   types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6rangeLogicFilterRulesAttrTypes}},
	"member":                               types.ObjectType{AttrTypes: Ipv6rangeMemberAttrTypes},
	"name":                                 types.StringType,
	"network":                              cidrtypes.IPv4PrefixType{},
	"network_view":                         types.StringType,
	"option_filter_rules":                  types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6rangeOptionFilterRulesAttrTypes}},
	"port_control_blackout_setting":        types.ObjectType{AttrTypes: Ipv6rangePortControlBlackoutSettingAttrTypes},
	"recycle_leases":                       types.BoolType,
	"restart_if_needed":                    types.BoolType,
	"same_port_control_discovery_blackout": types.BoolType,
	"server_association_type":              types.StringType,
	"start_addr":                           iptypes.IPv6AddressType{},
	"subscribe_settings":                   types.ObjectType{AttrTypes: Ipv6rangeSubscribeSettingsAttrTypes},
	"template":                             types.StringType,
	"use_blackout_setting":                 types.BoolType,
	"use_discovery_basic_polling_settings": types.BoolType,
	"use_enable_discovery":                 types.BoolType,
	"use_logic_filter_rules":               types.BoolType,
	"use_recycle_leases":                   types.BoolType,
	"use_subscribe_settings":               types.BoolType,
}

var Ipv6rangeResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"address_type": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Default:  stringdefault.StaticString("ADDRESS"),
		Validators: []validator.String{
			stringvalidator.OneOf("ADDRESS", "PREFIX", "BOTH"),
		},
		MarkdownDescription: "Type of a DHCP IPv6 Range object. Valid values are \"ADDRESS\", \"PREFIX\", or \"BOTH\". When the address type is \"ADDRESS\", values for the 'start_addr' and 'end_addr' members are required. When the address type is \"PREFIX\", values for 'ipv6_start_prefix', 'ipv6_end_prefix', and 'ipv6_prefix_bits' are required. When the address type is \"BOTH\", values for 'start_addr', 'end_addr', 'ipv6_start_prefix', 'ipv6_end_prefix', and 'ipv6_prefix_bits' are all required.",
	},
	"cloud_info": schema.SingleNestedAttribute{
		Attributes:          Ipv6rangeCloudInfoResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Structure containing all cloud API related information for this object.",
	},
	"comment": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 256),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Comment for the range; maximum 256 characters.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether a range is disabled or not. When this is set to False, the range is enabled.",
	},
	"discover_now_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Discover now status for this range.",
	},
	"discovery_basic_poll_settings": schema.SingleNestedAttribute{
		Attributes: Ipv6rangeDiscoveryBasicPollSettingsResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Validators: []validator.Object{
			objectvalidator.AlsoRequires(path.MatchRoot("use_discovery_basic_polling_settings")),
		},
		MarkdownDescription: "The discovery basic poll settings for this range.",
	},
	"discovery_blackout_setting": schema.SingleNestedAttribute{
		Attributes: Ipv6rangeDiscoveryBlackoutSettingResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Validators: []validator.Object{
			objectvalidator.AlsoRequires(path.MatchRoot("use_blackout_setting")),
		},
		MarkdownDescription: "The discovery blackout setting for this range.",
	},
	"discovery_member": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_enable_discovery")),
		},
		MarkdownDescription: "The member that will run discovery for this range.",
	},
	"enable_discovery": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_enable_discovery")),
		},
		MarkdownDescription: "Determines whether a discovery is enabled or not for this range. When this is set to False, the discovery for this range is disabled.",
	},
	"enable_immediate_discovery": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the discovery for the range should be immediately enabled.",
	},
	"end_addr": schema.StringAttribute{
		CustomType:          iptypes.IPv6AddressType{},
		Computed:            true,
		Optional:            true,
		MarkdownDescription: "The IPv6 Address end address of the DHCP IPv6 range.",
	},
	"endpoint_sources": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The endpoints that provides data for the DHCP IPv6 Range object.",
	},
	"exclude": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6rangeExcludeResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "These are ranges of IP addresses that the appliance does not use to assign to clients. You can use these exclusion addresses as static IP addresses. They contain the start and end addresses of the exclusion range, and optionally,information about this exclusion range.",
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
	"ipv6_end_prefix": schema.StringAttribute{
		CustomType:          iptypes.IPv6AddressType{},
		Computed:            true,
		Optional:            true,
		MarkdownDescription: "The IPv6 Address end prefix of the DHCP IPv6 range.",
	},
	"ipv6_prefix_bits": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Prefix bits of the DHCP IPv6 range.",
	},
	"ipv6_start_prefix": schema.StringAttribute{
		CustomType:          iptypes.IPv6AddressType{},
		Computed:            true,
		Optional:            true,
		MarkdownDescription: "The IPv6 Address starting prefix of the DHCP IPv6 range.",
	},
	"logic_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6rangeLogicFilterRulesResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
			listvalidator.AlsoRequires(path.MatchRoot("use_logic_filter_rules")),
		},
		MarkdownDescription: "This field contains the logic filters to be applied to this IPv6 range. This list corresponds to the match rules that are written to the DHCPv6 configuration file.",
	},
	"member": schema.SingleNestedAttribute{
		Attributes:          Ipv6rangeMemberResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The member that will provide service for this range. server_association_typeneeds to be set to ‘MEMBER’ if you want the server specified here to serve the range. For searching by this field you should use a HTTP method that contains a body (POST or PUT) with :ref:Dhcp Member structure<struct:dhcpmember>and the request should have option _method=GET.",
	},
	"name": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "This field contains the name of the Microsoft scope.",
	},
	"network": schema.StringAttribute{
		CustomType:          cidrtypes.IPv6PrefixType{},
		Required:            true,
		MarkdownDescription: "The network this range belongs to, in IPv6 Address/CIDR format.",
	},
	"network_view": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Default:  stringdefault.StaticString("default"),
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
		MarkdownDescription: "The name of the network view in which this range resides.",
	},
	"option_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6rangeOptionFilterRulesResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "This field contains the Option filters to be applied to this IPv6 range. The appliance uses the matching rules of these filters to select the address range from which it assigns a lease.",
	},
	"port_control_blackout_setting": schema.SingleNestedAttribute{
		Attributes:          Ipv6rangePortControlBlackoutSettingResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The port control blackout setting for this range.",
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
	"restart_if_needed": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Restarts the member service.",
	},
	"same_port_control_discovery_blackout": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_blackout_setting")),
		},
		MarkdownDescription: "If the field is set to True, the discovery blackout setting will be used for port control blackout setting.",
	},
	"server_association_type": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Default:  stringdefault.StaticString("NONE"),
		Validators: []validator.String{
			stringvalidator.OneOf("MEMBER", "NONE"),
		},
		MarkdownDescription: "The type of server that is going to serve the range. Valid values are: * MEMBER * NONE",
	},
	"start_addr": schema.StringAttribute{
		CustomType:          iptypes.IPv6AddressType{},
		Computed:            true,
		Optional:            true,
		MarkdownDescription: "The IPv6 Address starting address of the DHCP IPv6 range.",
	},
	"subscribe_settings": schema.SingleNestedAttribute{
		Attributes:          Ipv6rangeSubscribeSettingsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The DHCP IPv6 Range Cisco ISE subscribe settings.",
	},
	"template": schema.StringAttribute{
		Computed: true,
		Optional: true,
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
		MarkdownDescription: "If set on creation, the range will be created according to the values specified in the named template.",
	},
	"use_blackout_setting": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: discovery_blackout_setting , port_control_blackout_setting, same_port_control_discovery_blackout",
	},
	"use_discovery_basic_polling_settings": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: discovery_basic_poll_settings",
	},
	"use_enable_discovery": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: discovery_member , enable_discovery",
	},
	"use_logic_filter_rules": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: logic_filter_rules",
	},
	"use_recycle_leases": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: recycle_leases",
	},
	"use_subscribe_settings": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: subscribe_settings",
	},
}

func (m *Ipv6rangeModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *dhcp.Ipv6range {
	if m == nil {
		return nil
	}
	to := &dhcp.Ipv6range{
		AddressType:                      flex.ExpandStringPointer(m.AddressType),
		CloudInfo:                        ExpandIpv6rangeCloudInfo(ctx, m.CloudInfo, diags),
		Comment:                          flex.ExpandStringPointer(m.Comment),
		Disable:                          flex.ExpandBoolPointer(m.Disable),
		DiscoveryBasicPollSettings:       ExpandIpv6rangeDiscoveryBasicPollSettings(ctx, m.DiscoveryBasicPollSettings, diags),
		DiscoveryBlackoutSetting:         ExpandIpv6rangeDiscoveryBlackoutSetting(ctx, m.DiscoveryBlackoutSetting, diags),
		DiscoveryMember:                  flex.ExpandStringPointer(m.DiscoveryMember),
		EnableDiscovery:                  flex.ExpandBoolPointer(m.EnableDiscovery),
		EnableImmediateDiscovery:         flex.ExpandBoolPointer(m.EnableImmediateDiscovery),
		EndAddr:                          flex.ExpandIPv6Address(m.EndAddr),
		Exclude:                          flex.ExpandFrameworkListNestedBlock(ctx, m.Exclude, diags, ExpandIpv6rangeExclude),
		ExtAttrs:                         ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		Ipv6EndPrefix:                    flex.ExpandIPv6Address(m.Ipv6EndPrefix),
		Ipv6PrefixBits:                   flex.ExpandInt64Pointer(m.Ipv6PrefixBits),
		Ipv6StartPrefix:                  flex.ExpandIPv6Address(m.Ipv6StartPrefix),
		LogicFilterRules:                 flex.ExpandFrameworkListNestedBlock(ctx, m.LogicFilterRules, diags, ExpandIpv6rangeLogicFilterRules),
		Member:                           ExpandIpv6rangeMember(ctx, m.Member, diags),
		Name:                             flex.ExpandStringPointer(m.Name),
		Network:                          flex.ExpandIPv6CIDR(m.Network),
		OptionFilterRules:                flex.ExpandFrameworkListNestedBlock(ctx, m.OptionFilterRules, diags, ExpandIpv6rangeOptionFilterRules),
		PortControlBlackoutSetting:       ExpandIpv6rangePortControlBlackoutSetting(ctx, m.PortControlBlackoutSetting, diags),
		RecycleLeases:                    flex.ExpandBoolPointer(m.RecycleLeases),
		RestartIfNeeded:                  flex.ExpandBoolPointer(m.RestartIfNeeded),
		SamePortControlDiscoveryBlackout: flex.ExpandBoolPointer(m.SamePortControlDiscoveryBlackout),
		ServerAssociationType:            flex.ExpandStringPointer(m.ServerAssociationType),
		StartAddr:                        flex.ExpandIPv6Address(m.StartAddr),
		SubscribeSettings:                ExpandIpv6rangeSubscribeSettings(ctx, m.SubscribeSettings, diags),
		UseBlackoutSetting:               flex.ExpandBoolPointer(m.UseBlackoutSetting),
		UseDiscoveryBasicPollingSettings: flex.ExpandBoolPointer(m.UseDiscoveryBasicPollingSettings),
		UseEnableDiscovery:               flex.ExpandBoolPointer(m.UseEnableDiscovery),
		UseLogicFilterRules:              flex.ExpandBoolPointer(m.UseLogicFilterRules),
		UseRecycleLeases:                 flex.ExpandBoolPointer(m.UseRecycleLeases),
		UseSubscribeSettings:             flex.ExpandBoolPointer(m.UseSubscribeSettings),
		NetworkView:                      flex.ExpandStringPointer(m.NetworkView),
	}
	if isCreate {
		to.Template = flex.ExpandStringPointer(m.Template)
	}
	return to
}

func FlattenIpv6range(ctx context.Context, from *dhcp.Ipv6range, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6rangeAttrTypes)
	}
	m := Ipv6rangeModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, Ipv6rangeAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Ipv6rangeModel) Flatten(ctx context.Context, from *dhcp.Ipv6range, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Ipv6rangeModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.AddressType = flex.FlattenStringPointer(from.AddressType)
	m.CloudInfo = FlattenIpv6rangeCloudInfo(ctx, from.CloudInfo, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.DiscoverNowStatus = flex.FlattenStringPointer(from.DiscoverNowStatus)
	m.DiscoveryBasicPollSettings = FlattenIpv6rangeDiscoveryBasicPollSettings(ctx, from.DiscoveryBasicPollSettings, diags)
	m.DiscoveryBlackoutSetting = FlattenIpv6rangeDiscoveryBlackoutSetting(ctx, from.DiscoveryBlackoutSetting, diags)
	m.DiscoveryMember = flex.FlattenStringPointer(from.DiscoveryMember)
	m.EnableDiscovery = types.BoolPointerValue(from.EnableDiscovery)
	m.EndAddr = flex.FlattenIPv6Address(from.EndAddr)
	m.EndpointSources = flex.FlattenFrameworkListString(ctx, from.EndpointSources, diags)
	m.Exclude = flex.FlattenFrameworkListNestedBlock(ctx, from.Exclude, Ipv6rangeExcludeAttrTypes, diags, FlattenIpv6rangeExclude)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.Ipv6EndPrefix = flex.FlattenIPv6Address(from.Ipv6EndPrefix)
	m.Ipv6PrefixBits = flex.FlattenInt64Pointer(from.Ipv6PrefixBits)
	m.Ipv6StartPrefix = flex.FlattenIPv6Address(from.Ipv6StartPrefix)
	m.LogicFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.LogicFilterRules, Ipv6rangeLogicFilterRulesAttrTypes, diags, FlattenIpv6rangeLogicFilterRules)
	m.Member = FlattenIpv6rangeMember(ctx, from.Member, diags)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Network = flex.FlattenIPv6CIDR(from.Network)
	m.NetworkView = flex.FlattenStringPointer(from.NetworkView)
	m.OptionFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.OptionFilterRules, Ipv6rangeOptionFilterRulesAttrTypes, diags, FlattenIpv6rangeOptionFilterRules)
	m.PortControlBlackoutSetting = FlattenIpv6rangePortControlBlackoutSetting(ctx, from.PortControlBlackoutSetting, diags)
	m.RecycleLeases = types.BoolPointerValue(from.RecycleLeases)
	m.SamePortControlDiscoveryBlackout = types.BoolPointerValue(from.SamePortControlDiscoveryBlackout)
	m.ServerAssociationType = flex.FlattenStringPointer(from.ServerAssociationType)
	m.StartAddr = flex.FlattenIPv6Address(from.StartAddr)
	m.SubscribeSettings = FlattenIpv6rangeSubscribeSettings(ctx, from.SubscribeSettings, diags)
	m.Template = flex.FlattenStringPointer(from.Template)
	m.UseBlackoutSetting = types.BoolPointerValue(from.UseBlackoutSetting)
	m.UseDiscoveryBasicPollingSettings = types.BoolPointerValue(from.UseDiscoveryBasicPollingSettings)
	m.UseEnableDiscovery = types.BoolPointerValue(from.UseEnableDiscovery)
	m.UseLogicFilterRules = types.BoolPointerValue(from.UseLogicFilterRules)
	m.UseRecycleLeases = types.BoolPointerValue(from.UseRecycleLeases)
	m.UseSubscribeSettings = types.BoolPointerValue(from.UseSubscribeSettings)
}
