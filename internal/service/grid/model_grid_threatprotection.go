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

type GridThreatprotectionModel struct {
	Ref                                   types.String `tfsdk:"ref"`
	CurrentRuleset                        types.String `tfsdk:"current_ruleset"`
	DisableMultipleDnsTcpRequest          types.Bool   `tfsdk:"disable_multiple_dns_tcp_request"`
	EnableAccelRespBeforeThreatProtection types.Bool   `tfsdk:"enable_accel_resp_before_threat_protection"`
	EnableAutoDownload                    types.Bool   `tfsdk:"enable_auto_download"`
	EnableNatRules                        types.Bool   `tfsdk:"enable_nat_rules"`
	EnableScheduledDownload               types.Bool   `tfsdk:"enable_scheduled_download"`
	EventsPerSecondPerRule                types.Int64  `tfsdk:"events_per_second_per_rule"`
	GridName                              types.String `tfsdk:"grid_name"`
	LastCheckedForUpdate                  types.Int64  `tfsdk:"last_checked_for_update"`
	LastRuleUpdateTimestamp               types.Int64  `tfsdk:"last_rule_update_timestamp"`
	LastRuleUpdateVersion                 types.String `tfsdk:"last_rule_update_version"`
	NatRules                              types.List   `tfsdk:"nat_rules"`
	OutboundSettings                      types.Object `tfsdk:"outbound_settings"`
	RuleUpdatePolicy                      types.String `tfsdk:"rule_update_policy"`
	ScheduledDownload                     types.Object `tfsdk:"scheduled_download"`
}

var GridThreatprotectionAttrTypes = map[string]attr.Type{
	"ref":                              types.StringType,
	"current_ruleset":                  types.StringType,
	"disable_multiple_dns_tcp_request": types.BoolType,
	"enable_accel_resp_before_threat_protection": types.BoolType,
	"enable_auto_download":                       types.BoolType,
	"enable_nat_rules":                           types.BoolType,
	"enable_scheduled_download":                  types.BoolType,
	"events_per_second_per_rule":                 types.Int64Type,
	"grid_name":                                  types.StringType,
	"last_checked_for_update":                    types.Int64Type,
	"last_rule_update_timestamp":                 types.Int64Type,
	"last_rule_update_version":                   types.StringType,
	"nat_rules":                                  types.ListType{ElemType: types.ObjectType{AttrTypes: GridThreatprotectionNatRulesAttrTypes}},
	"outbound_settings":                          types.ObjectType{AttrTypes: GridThreatprotectionOutboundSettingsAttrTypes},
	"rule_update_policy":                         types.StringType,
	"scheduled_download":                         types.ObjectType{AttrTypes: GridThreatprotectionScheduledDownloadAttrTypes},
}

var GridThreatprotectionResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"current_ruleset": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The current Grid ruleset.",
	},
	"disable_multiple_dns_tcp_request": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if multiple BIND responses via TCP connection are disabled.",
	},
	"enable_accel_resp_before_threat_protection": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if DNS responses are sent from acceleration cache before applying Threat Protection rules. Recommended for better performance when using DNS Cache Acceleration.",
	},
	"enable_auto_download": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if auto download service is enabled.",
	},
	"enable_nat_rules": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if NAT (Network Address Translation) mapping for threat protection is enabled or not.",
	},
	"enable_scheduled_download": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if scheduled download is enabled. The default frequency is once in every 24 hours if it is disabled.",
	},
	"events_per_second_per_rule": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The number of events logged per second per rule.",
	},
	"grid_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The Grid name.",
	},
	"last_checked_for_update": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The time when the Grid last checked for updates.",
	},
	"last_rule_update_timestamp": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The last rule update timestamp.",
	},
	"last_rule_update_version": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The version of last rule update.",
	},
	"nat_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: GridThreatprotectionNatRulesResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of NAT mapping rules for threat protection.",
	},
	"outbound_settings": schema.SingleNestedAttribute{
		Attributes: GridThreatprotectionOutboundSettingsResourceSchemaAttributes,
		Optional:   true,
	},
	"rule_update_policy": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The update rule policy.",
	},
	"scheduled_download": schema.SingleNestedAttribute{
		Attributes: GridThreatprotectionScheduledDownloadResourceSchemaAttributes,
		Optional:   true,
	},
}

func ExpandGridThreatprotection(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridThreatprotection {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridThreatprotectionModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridThreatprotectionModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridThreatprotection {
	if m == nil {
		return nil
	}
	to := &grid.GridThreatprotection{
		Ref:                                   flex.ExpandStringPointer(m.Ref),
		CurrentRuleset:                        flex.ExpandStringPointer(m.CurrentRuleset),
		DisableMultipleDnsTcpRequest:          flex.ExpandBoolPointer(m.DisableMultipleDnsTcpRequest),
		EnableAccelRespBeforeThreatProtection: flex.ExpandBoolPointer(m.EnableAccelRespBeforeThreatProtection),
		EnableAutoDownload:                    flex.ExpandBoolPointer(m.EnableAutoDownload),
		EnableNatRules:                        flex.ExpandBoolPointer(m.EnableNatRules),
		EnableScheduledDownload:               flex.ExpandBoolPointer(m.EnableScheduledDownload),
		EventsPerSecondPerRule:                flex.ExpandInt64Pointer(m.EventsPerSecondPerRule),
		NatRules:                              flex.ExpandFrameworkListNestedBlock(ctx, m.NatRules, diags, ExpandGridThreatprotectionNatRules),
		OutboundSettings:                      ExpandGridThreatprotectionOutboundSettings(ctx, m.OutboundSettings, diags),
		RuleUpdatePolicy:                      flex.ExpandStringPointer(m.RuleUpdatePolicy),
		ScheduledDownload:                     ExpandGridThreatprotectionScheduledDownload(ctx, m.ScheduledDownload, diags),
	}
	return to
}

func FlattenGridThreatprotection(ctx context.Context, from *grid.GridThreatprotection, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridThreatprotectionAttrTypes)
	}
	m := GridThreatprotectionModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridThreatprotectionAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridThreatprotectionModel) Flatten(ctx context.Context, from *grid.GridThreatprotection, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridThreatprotectionModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.CurrentRuleset = flex.FlattenStringPointer(from.CurrentRuleset)
	m.DisableMultipleDnsTcpRequest = types.BoolPointerValue(from.DisableMultipleDnsTcpRequest)
	m.EnableAccelRespBeforeThreatProtection = types.BoolPointerValue(from.EnableAccelRespBeforeThreatProtection)
	m.EnableAutoDownload = types.BoolPointerValue(from.EnableAutoDownload)
	m.EnableNatRules = types.BoolPointerValue(from.EnableNatRules)
	m.EnableScheduledDownload = types.BoolPointerValue(from.EnableScheduledDownload)
	m.EventsPerSecondPerRule = flex.FlattenInt64Pointer(from.EventsPerSecondPerRule)
	m.GridName = flex.FlattenStringPointer(from.GridName)
	m.LastCheckedForUpdate = flex.FlattenInt64Pointer(from.LastCheckedForUpdate)
	m.LastRuleUpdateTimestamp = flex.FlattenInt64Pointer(from.LastRuleUpdateTimestamp)
	m.LastRuleUpdateVersion = flex.FlattenStringPointer(from.LastRuleUpdateVersion)
	m.NatRules = flex.FlattenFrameworkListNestedBlock(ctx, from.NatRules, GridThreatprotectionNatRulesAttrTypes, diags, FlattenGridThreatprotectionNatRules)
	m.OutboundSettings = FlattenGridThreatprotectionOutboundSettings(ctx, from.OutboundSettings, diags)
	m.RuleUpdatePolicy = flex.FlattenStringPointer(from.RuleUpdatePolicy)
	m.ScheduledDownload = FlattenGridThreatprotectionScheduledDownload(ctx, from.ScheduledDownload, diags)
}
