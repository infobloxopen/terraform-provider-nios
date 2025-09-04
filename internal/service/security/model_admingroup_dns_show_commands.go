package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type AdmingroupDnsShowCommandsModel struct {
	ShowLogGuestLookups              types.Bool `tfsdk:"show_log_guest_lookups"`
	ShowDnsGssTsig                   types.Bool `tfsdk:"show_dns_gss_tsig"`
	ShowDns                          types.Bool `tfsdk:"show_dns"`
	ShowDnstapStats                  types.Bool `tfsdk:"show_dnstap_stats"`
	ShowDnstapStatus                 types.Bool `tfsdk:"show_dnstap_status"`
	ShowDnsOverTlsConfig             types.Bool `tfsdk:"show_dns_over_tls_config"`
	ShowDnsOverTlsStatus             types.Bool `tfsdk:"show_dns_over_tls_status"`
	ShowDnsOverTlsStats              types.Bool `tfsdk:"show_dns_over_tls_stats"`
	ShowDohConfig                    types.Bool `tfsdk:"show_doh_config"`
	ShowDohStatus                    types.Bool `tfsdk:"show_doh_status"`
	ShowDohStats                     types.Bool `tfsdk:"show_doh_stats"`
	ShowExtraDnsNameValidations      types.Bool `tfsdk:"show_extra_dns_name_validations"`
	ShowMsStickyIp                   types.Bool `tfsdk:"show_ms_sticky_ip"`
	ShowDnsRrl                       types.Bool `tfsdk:"show_dns_rrl"`
	ShowEnableMatchRecursiveOnly     types.Bool `tfsdk:"show_enable_match_recursive_only"`
	ShowMaxRecursionDepth            types.Bool `tfsdk:"show_max_recursion_depth"`
	ShowMaxRecursionQueries          types.Bool `tfsdk:"show_max_recursion_queries"`
	ShowMonitor                      types.Bool `tfsdk:"show_monitor"`
	ShowQueryCapture                 types.Bool `tfsdk:"show_query_capture"`
	ShowDtcEa                        types.Bool `tfsdk:"show_dtc_ea"`
	ShowDtcGeoip                     types.Bool `tfsdk:"show_dtc_geoip"`
	ShowRestartAnycastWithDnsRestart types.Bool `tfsdk:"show_restart_anycast_with_dns_restart"`
	ShowRpzAddSoa                    types.Bool `tfsdk:"show_rpz_add_soa"`
	ShowDnsAccel                     types.Bool `tfsdk:"show_dns_accel"`
	ShowDnsAccelDebug                types.Bool `tfsdk:"show_dns_accel_debug"`
	ShowAllowQueryDomain             types.Bool `tfsdk:"show_allow_query_domain"`
	ShowAllowQueryDomainViews        types.Bool `tfsdk:"show_allow_query_domain_views"`
	EnableAll                        types.Bool `tfsdk:"enable_all"`
	DisableAll                       types.Bool `tfsdk:"disable_all"`
}

var AdmingroupDnsShowCommandsAttrTypes = map[string]attr.Type{
	"show_log_guest_lookups":                types.BoolType,
	"show_dns_gss_tsig":                     types.BoolType,
	"show_dns":                              types.BoolType,
	"show_dnstap_stats":                     types.BoolType,
	"show_dnstap_status":                    types.BoolType,
	"show_dns_over_tls_config":              types.BoolType,
	"show_dns_over_tls_status":              types.BoolType,
	"show_dns_over_tls_stats":               types.BoolType,
	"show_doh_config":                       types.BoolType,
	"show_doh_status":                       types.BoolType,
	"show_doh_stats":                        types.BoolType,
	"show_extra_dns_name_validations":       types.BoolType,
	"show_ms_sticky_ip":                     types.BoolType,
	"show_dns_rrl":                          types.BoolType,
	"show_enable_match_recursive_only":      types.BoolType,
	"show_max_recursion_depth":              types.BoolType,
	"show_max_recursion_queries":            types.BoolType,
	"show_monitor":                          types.BoolType,
	"show_query_capture":                    types.BoolType,
	"show_dtc_ea":                           types.BoolType,
	"show_dtc_geoip":                        types.BoolType,
	"show_restart_anycast_with_dns_restart": types.BoolType,
	"show_rpz_add_soa":                      types.BoolType,
	"show_dns_accel":                        types.BoolType,
	"show_dns_accel_debug":                  types.BoolType,
	"show_allow_query_domain":               types.BoolType,
	"show_allow_query_domain_views":         types.BoolType,
	"enable_all":                            types.BoolType,
	"disable_all":                           types.BoolType,
}

var AdmingroupDnsShowCommandsResourceSchemaAttributes = map[string]schema.Attribute{
	"show_log_guest_lookups": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_dns_gss_tsig": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_dns": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_dnstap_stats": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_dnstap_status": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_dns_over_tls_config": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_dns_over_tls_status": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_dns_over_tls_stats": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_doh_config": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_doh_status": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_doh_stats": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_extra_dns_name_validations": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_ms_sticky_ip": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_dns_rrl": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_enable_match_recursive_only": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_max_recursion_depth": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_max_recursion_queries": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_monitor": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_query_capture": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_dtc_ea": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_dtc_geoip": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_restart_anycast_with_dns_restart": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_rpz_add_soa": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_dns_accel": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_dns_accel_debug": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_allow_query_domain": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_allow_query_domain_views": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"enable_all": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then enable all fields",
	},
	"disable_all": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then disable all fields",
	},
}

func ExpandAdmingroupDnsShowCommands(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.AdmingroupDnsShowCommands {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AdmingroupDnsShowCommandsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AdmingroupDnsShowCommandsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.AdmingroupDnsShowCommands {
	if m == nil {
		return nil
	}
	to := &security.AdmingroupDnsShowCommands{
		ShowLogGuestLookups:              flex.ExpandBoolPointer(m.ShowLogGuestLookups),
		ShowDnsGssTsig:                   flex.ExpandBoolPointer(m.ShowDnsGssTsig),
		ShowDns:                          flex.ExpandBoolPointer(m.ShowDns),
		ShowDnstapStats:                  flex.ExpandBoolPointer(m.ShowDnstapStats),
		ShowDnstapStatus:                 flex.ExpandBoolPointer(m.ShowDnstapStatus),
		ShowDnsOverTlsConfig:             flex.ExpandBoolPointer(m.ShowDnsOverTlsConfig),
		ShowDnsOverTlsStatus:             flex.ExpandBoolPointer(m.ShowDnsOverTlsStatus),
		ShowDnsOverTlsStats:              flex.ExpandBoolPointer(m.ShowDnsOverTlsStats),
		ShowDohConfig:                    flex.ExpandBoolPointer(m.ShowDohConfig),
		ShowDohStatus:                    flex.ExpandBoolPointer(m.ShowDohStatus),
		ShowDohStats:                     flex.ExpandBoolPointer(m.ShowDohStats),
		ShowExtraDnsNameValidations:      flex.ExpandBoolPointer(m.ShowExtraDnsNameValidations),
		ShowMsStickyIp:                   flex.ExpandBoolPointer(m.ShowMsStickyIp),
		ShowDnsRrl:                       flex.ExpandBoolPointer(m.ShowDnsRrl),
		ShowEnableMatchRecursiveOnly:     flex.ExpandBoolPointer(m.ShowEnableMatchRecursiveOnly),
		ShowMaxRecursionDepth:            flex.ExpandBoolPointer(m.ShowMaxRecursionDepth),
		ShowMaxRecursionQueries:          flex.ExpandBoolPointer(m.ShowMaxRecursionQueries),
		ShowMonitor:                      flex.ExpandBoolPointer(m.ShowMonitor),
		ShowQueryCapture:                 flex.ExpandBoolPointer(m.ShowQueryCapture),
		ShowDtcEa:                        flex.ExpandBoolPointer(m.ShowDtcEa),
		ShowDtcGeoip:                     flex.ExpandBoolPointer(m.ShowDtcGeoip),
		ShowRestartAnycastWithDnsRestart: flex.ExpandBoolPointer(m.ShowRestartAnycastWithDnsRestart),
		ShowRpzAddSoa:                    flex.ExpandBoolPointer(m.ShowRpzAddSoa),
		ShowDnsAccel:                     flex.ExpandBoolPointer(m.ShowDnsAccel),
		ShowDnsAccelDebug:                flex.ExpandBoolPointer(m.ShowDnsAccelDebug),
		ShowAllowQueryDomain:             flex.ExpandBoolPointer(m.ShowAllowQueryDomain),
		ShowAllowQueryDomainViews:        flex.ExpandBoolPointer(m.ShowAllowQueryDomainViews),
		EnableAll:                        flex.ExpandBoolPointer(m.EnableAll),
		DisableAll:                       flex.ExpandBoolPointer(m.DisableAll),
	}
	return to
}

func FlattenAdmingroupDnsShowCommands(ctx context.Context, from *security.AdmingroupDnsShowCommands, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AdmingroupDnsShowCommandsAttrTypes)
	}
	m := AdmingroupDnsShowCommandsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AdmingroupDnsShowCommandsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AdmingroupDnsShowCommandsModel) Flatten(ctx context.Context, from *security.AdmingroupDnsShowCommands, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AdmingroupDnsShowCommandsModel{}
	}
	m.ShowLogGuestLookups = types.BoolPointerValue(from.ShowLogGuestLookups)
	m.ShowDnsGssTsig = types.BoolPointerValue(from.ShowDnsGssTsig)
	m.ShowDns = types.BoolPointerValue(from.ShowDns)
	m.ShowDnstapStats = types.BoolPointerValue(from.ShowDnstapStats)
	m.ShowDnstapStatus = types.BoolPointerValue(from.ShowDnstapStatus)
	m.ShowDnsOverTlsConfig = types.BoolPointerValue(from.ShowDnsOverTlsConfig)
	m.ShowDnsOverTlsStatus = types.BoolPointerValue(from.ShowDnsOverTlsStatus)
	m.ShowDnsOverTlsStats = types.BoolPointerValue(from.ShowDnsOverTlsStats)
	m.ShowDohConfig = types.BoolPointerValue(from.ShowDohConfig)
	m.ShowDohStatus = types.BoolPointerValue(from.ShowDohStatus)
	m.ShowDohStats = types.BoolPointerValue(from.ShowDohStats)
	m.ShowExtraDnsNameValidations = types.BoolPointerValue(from.ShowExtraDnsNameValidations)
	m.ShowMsStickyIp = types.BoolPointerValue(from.ShowMsStickyIp)
	m.ShowDnsRrl = types.BoolPointerValue(from.ShowDnsRrl)
	m.ShowEnableMatchRecursiveOnly = types.BoolPointerValue(from.ShowEnableMatchRecursiveOnly)
	m.ShowMaxRecursionDepth = types.BoolPointerValue(from.ShowMaxRecursionDepth)
	m.ShowMaxRecursionQueries = types.BoolPointerValue(from.ShowMaxRecursionQueries)
	m.ShowMonitor = types.BoolPointerValue(from.ShowMonitor)
	m.ShowQueryCapture = types.BoolPointerValue(from.ShowQueryCapture)
	m.ShowDtcEa = types.BoolPointerValue(from.ShowDtcEa)
	m.ShowDtcGeoip = types.BoolPointerValue(from.ShowDtcGeoip)
	m.ShowRestartAnycastWithDnsRestart = types.BoolPointerValue(from.ShowRestartAnycastWithDnsRestart)
	m.ShowRpzAddSoa = types.BoolPointerValue(from.ShowRpzAddSoa)
	m.ShowDnsAccel = types.BoolPointerValue(from.ShowDnsAccel)
	m.ShowDnsAccelDebug = types.BoolPointerValue(from.ShowDnsAccelDebug)
	m.ShowAllowQueryDomain = types.BoolPointerValue(from.ShowAllowQueryDomain)
	m.ShowAllowQueryDomainViews = types.BoolPointerValue(from.ShowAllowQueryDomainViews)
	m.EnableAll = types.BoolPointerValue(from.EnableAll)
	m.DisableAll = types.BoolPointerValue(from.DisableAll)
}
