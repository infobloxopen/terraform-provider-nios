package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type MemberThreatprotectionModel struct {
	Ref                                      types.String `tfsdk:"ref"`
	Comment                                  types.String `tfsdk:"comment"`
	CurrentRuleset                           types.String `tfsdk:"current_ruleset"`
	DisableMultipleDnsTcpRequest             types.Bool   `tfsdk:"disable_multiple_dns_tcp_request"`
	EnableAccelRespBeforeThreatProtection    types.Bool   `tfsdk:"enable_accel_resp_before_threat_protection"`
	EnableNatRules                           types.Bool   `tfsdk:"enable_nat_rules"`
	EnableService                            types.Bool   `tfsdk:"enable_service"`
	EventsPerSecondPerRule                   types.Int64  `tfsdk:"events_per_second_per_rule"`
	HardwareModel                            types.String `tfsdk:"hardware_model"`
	HardwareType                             types.String `tfsdk:"hardware_type"`
	HostName                                 types.String `tfsdk:"host_name"`
	Ipv4address                              types.String `tfsdk:"ipv4address"`
	Ipv6address                              types.String `tfsdk:"ipv6address"`
	NatRules                                 types.List   `tfsdk:"nat_rules"`
	OutboundSettings                         types.Object `tfsdk:"outbound_settings"`
	Profile                                  types.String `tfsdk:"profile"`
	UseCurrentRuleset                        types.Bool   `tfsdk:"use_current_ruleset"`
	UseDisableMultipleDnsTcpRequest          types.Bool   `tfsdk:"use_disable_multiple_dns_tcp_request"`
	UseEnableAccelRespBeforeThreatProtection types.Bool   `tfsdk:"use_enable_accel_resp_before_threat_protection"`
	UseEnableNatRules                        types.Bool   `tfsdk:"use_enable_nat_rules"`
	UseEventsPerSecondPerRule                types.Bool   `tfsdk:"use_events_per_second_per_rule"`
	UseOutboundSettings                      types.Bool   `tfsdk:"use_outbound_settings"`
}

var MemberThreatprotectionAttrTypes = map[string]attr.Type{
	"ref":                              types.StringType,
	"comment":                          types.StringType,
	"current_ruleset":                  types.StringType,
	"disable_multiple_dns_tcp_request": types.BoolType,
	"enable_accel_resp_before_threat_protection": types.BoolType,
	"enable_nat_rules":                           types.BoolType,
	"enable_service":                             types.BoolType,
	"events_per_second_per_rule":                 types.Int64Type,
	"hardware_model":                             types.StringType,
	"hardware_type":                              types.StringType,
	"host_name":                                  types.StringType,
	"ipv4address":                                types.StringType,
	"ipv6address":                                types.StringType,
	"nat_rules":                                  types.ListType{ElemType: types.ObjectType{AttrTypes: MemberThreatprotectionNatRulesAttrTypes}},
	"outbound_settings":                          types.ObjectType{AttrTypes: MemberThreatprotectionOutboundSettingsAttrTypes},
	"profile":                                    types.StringType,
	"use_current_ruleset":                        types.BoolType,
	"use_disable_multiple_dns_tcp_request":       types.BoolType,
	"use_enable_accel_resp_before_threat_protection": types.BoolType,
	"use_enable_nat_rules":                           types.BoolType,
	"use_events_per_second_per_rule":                 types.BoolType,
	"use_outbound_settings":                          types.BoolType,
}

var MemberThreatprotectionResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"comment": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The human readable comment for member threat protection properties.",
	},
	"current_ruleset": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The ruleset used for threat protection.",
	},
	"disable_multiple_dns_tcp_request": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if multiple BIND responses via TCP connection is enabled or not.",
	},
	"enable_accel_resp_before_threat_protection": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if DNS responses are sent from acceleration cache before applying Threat Protection rules. Recommended for better performance when using DNS Cache Acceleration.",
	},
	"enable_nat_rules": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if NAT (Network Address Translation) mapping for threat protection is enabled or not.",
	},
	"enable_service": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the Threat protection service is enabled or not.",
	},
	"events_per_second_per_rule": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The number of events logged per second per rule.",
	},
	"hardware_model": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The hardware model of the member.",
	},
	"hardware_type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The hardware type of the member.",
	},
	"host_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "A Grid member name.",
	},
	"ipv4address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The IPv4 address of member threat protection service.",
	},
	"ipv6address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The IPv6 address of member threat protection service.",
	},
	"nat_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberThreatprotectionNatRulesResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of NAT rules.",
	},
	"outbound_settings": schema.SingleNestedAttribute{
		Attributes: MemberThreatprotectionOutboundSettingsResourceSchemaAttributes,
		Optional:   true,
	},
	"profile": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The Threat Protection profile associated with the member.",
	},
	"use_current_ruleset": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: current_ruleset",
	},
	"use_disable_multiple_dns_tcp_request": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: disable_multiple_dns_tcp_request",
	},
	"use_enable_accel_resp_before_threat_protection": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: enable_accel_resp_before_threat_protection",
	},
	"use_enable_nat_rules": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: enable_nat_rules",
	},
	"use_events_per_second_per_rule": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: events_per_second_per_rule",
	},
	"use_outbound_settings": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: outbound_settings",
	},
}

func ExpandMemberThreatprotection(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberThreatprotection {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberThreatprotectionModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberThreatprotectionModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberThreatprotection {
	if m == nil {
		return nil
	}
	to := &grid.MemberThreatprotection{
		Ref:                                      flex.ExpandStringPointer(m.Ref),
		CurrentRuleset:                           flex.ExpandStringPointer(m.CurrentRuleset),
		DisableMultipleDnsTcpRequest:             flex.ExpandBoolPointer(m.DisableMultipleDnsTcpRequest),
		EnableAccelRespBeforeThreatProtection:    flex.ExpandBoolPointer(m.EnableAccelRespBeforeThreatProtection),
		EnableNatRules:                           flex.ExpandBoolPointer(m.EnableNatRules),
		EnableService:                            flex.ExpandBoolPointer(m.EnableService),
		EventsPerSecondPerRule:                   flex.ExpandInt64Pointer(m.EventsPerSecondPerRule),
		NatRules:                                 flex.ExpandFrameworkListNestedBlock(ctx, m.NatRules, diags, ExpandMemberThreatprotectionNatRules),
		OutboundSettings:                         ExpandMemberThreatprotectionOutboundSettings(ctx, m.OutboundSettings, diags),
		Profile:                                  flex.ExpandStringPointer(m.Profile),
		UseCurrentRuleset:                        flex.ExpandBoolPointer(m.UseCurrentRuleset),
		UseDisableMultipleDnsTcpRequest:          flex.ExpandBoolPointer(m.UseDisableMultipleDnsTcpRequest),
		UseEnableAccelRespBeforeThreatProtection: flex.ExpandBoolPointer(m.UseEnableAccelRespBeforeThreatProtection),
		UseEnableNatRules:                        flex.ExpandBoolPointer(m.UseEnableNatRules),
		UseEventsPerSecondPerRule:                flex.ExpandBoolPointer(m.UseEventsPerSecondPerRule),
		UseOutboundSettings:                      flex.ExpandBoolPointer(m.UseOutboundSettings),
	}
	return to
}

func FlattenMemberThreatprotection(ctx context.Context, from *grid.MemberThreatprotection, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberThreatprotectionAttrTypes)
	}
	m := MemberThreatprotectionModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberThreatprotectionAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberThreatprotectionModel) Flatten(ctx context.Context, from *grid.MemberThreatprotection, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberThreatprotectionModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CurrentRuleset = flex.FlattenStringPointer(from.CurrentRuleset)
	m.DisableMultipleDnsTcpRequest = types.BoolPointerValue(from.DisableMultipleDnsTcpRequest)
	m.EnableAccelRespBeforeThreatProtection = types.BoolPointerValue(from.EnableAccelRespBeforeThreatProtection)
	m.EnableNatRules = types.BoolPointerValue(from.EnableNatRules)
	m.EnableService = types.BoolPointerValue(from.EnableService)
	m.EventsPerSecondPerRule = flex.FlattenInt64Pointer(from.EventsPerSecondPerRule)
	m.HardwareModel = flex.FlattenStringPointer(from.HardwareModel)
	m.HardwareType = flex.FlattenStringPointer(from.HardwareType)
	m.HostName = flex.FlattenStringPointer(from.HostName)
	m.Ipv4address = flex.FlattenStringPointer(from.Ipv4address)
	m.Ipv6address = flex.FlattenStringPointer(from.Ipv6address)
	m.NatRules = flex.FlattenFrameworkListNestedBlock(ctx, from.NatRules, MemberThreatprotectionNatRulesAttrTypes, diags, FlattenMemberThreatprotectionNatRules)
	m.OutboundSettings = FlattenMemberThreatprotectionOutboundSettings(ctx, from.OutboundSettings, diags)
	m.Profile = flex.FlattenStringPointer(from.Profile)
	m.UseCurrentRuleset = types.BoolPointerValue(from.UseCurrentRuleset)
	m.UseDisableMultipleDnsTcpRequest = types.BoolPointerValue(from.UseDisableMultipleDnsTcpRequest)
	m.UseEnableAccelRespBeforeThreatProtection = types.BoolPointerValue(from.UseEnableAccelRespBeforeThreatProtection)
	m.UseEnableNatRules = types.BoolPointerValue(from.UseEnableNatRules)
	m.UseEventsPerSecondPerRule = types.BoolPointerValue(from.UseEventsPerSecondPerRule)
	m.UseOutboundSettings = types.BoolPointerValue(from.UseOutboundSettings)
}
