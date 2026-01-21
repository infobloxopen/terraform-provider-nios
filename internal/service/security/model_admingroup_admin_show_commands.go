package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type AdmingroupAdminShowCommandsModel struct {
	ShowAdminGroupAcl             types.Bool `tfsdk:"show_admin_group_acl"`
	ShowAnalyticsParameter        types.Bool `tfsdk:"show_analytics_parameter"`
	ShowArp                       types.Bool `tfsdk:"show_arp"`
	ShowBfd                       types.Bool `tfsdk:"show_bfd"`
	ShowBgp                       types.Bool `tfsdk:"show_bgp"`
	ShowCapacity                  types.Bool `tfsdk:"show_capacity"`
	ShowClusterdInfo              types.Bool `tfsdk:"show_clusterd_info"`
	ShowConfig                    types.Bool `tfsdk:"show_config"`
	ShowCpu                       types.Bool `tfsdk:"show_cpu"`
	ShowDate                      types.Bool `tfsdk:"show_date"`
	ShowDebug                     types.Bool `tfsdk:"show_debug"`
	ShowDebugAnalytics            types.Bool `tfsdk:"show_debug_analytics"`
	ShowDeleteTasksInterval       types.Bool `tfsdk:"show_delete_tasks_interval"`
	ShowDisk                      types.Bool `tfsdk:"show_disk"`
	ShowHardwareType              types.Bool `tfsdk:"show_hardware_type"`
	ShowHardwareStatus            types.Bool `tfsdk:"show_hardware_status"`
	ShowHwid                      types.Bool `tfsdk:"show_hwid"`
	ShowIbtrap                    types.Bool `tfsdk:"show_ibtrap"`
	ShowHwIdent                   types.Bool `tfsdk:"show_hw_ident"`
	ShowLog                       types.Bool `tfsdk:"show_log"`
	ShowLogfiles                  types.Bool `tfsdk:"show_logfiles"`
	ShowMemory                    types.Bool `tfsdk:"show_memory"`
	ShowNtp                       types.Bool `tfsdk:"show_ntp"`
	ShowReportingUserCapabilities types.Bool `tfsdk:"show_reporting_user_capabilities"`
	ShowRpzRecursiveOnly          types.Bool `tfsdk:"show_rpz_recursive_only"`
	ShowScheduled                 types.Bool `tfsdk:"show_scheduled"`
	ShowSnmp                      types.Bool `tfsdk:"show_snmp"`
	ShowStatus                    types.Bool `tfsdk:"show_status"`
	ShowTechSupport               types.Bool `tfsdk:"show_tech_support"`
	ShowTemperature               types.Bool `tfsdk:"show_temperature"`
	ShowThresholdtrap             types.Bool `tfsdk:"show_thresholdtrap"`
	ShowUpgradeHistory            types.Bool `tfsdk:"show_upgrade_history"`
	ShowUptime                    types.Bool `tfsdk:"show_uptime"`
	ShowVersion                   types.Bool `tfsdk:"show_version"`
	ShowAnalyticsDatabaseDumps    types.Bool `tfsdk:"show_analytics_database_dumps"`
	ShowCores                     types.Bool `tfsdk:"show_cores"`
	ShowCoresummary               types.Bool `tfsdk:"show_coresummary"`
	ShowCspThreatDb               types.Bool `tfsdk:"show_csp_threat_db"`
	ShowHsmGroup                  types.Bool `tfsdk:"show_hsm_group"`
	ShowHsmInfo                   types.Bool `tfsdk:"show_hsm_info"`
	ShowPmap                      types.Bool `tfsdk:"show_pmap"`
	ShowProcess                   types.Bool `tfsdk:"show_process"`
	ShowPstack                    types.Bool `tfsdk:"show_pstack"`
	ShowSafenetSupportInfo        types.Bool `tfsdk:"show_safenet_support_info"`
	ShowWredStats                 types.Bool `tfsdk:"show_wred_stats"`
	ShowWredStatus                types.Bool `tfsdk:"show_wred_status"`
	ShowNtpStratum                types.Bool `tfsdk:"show_ntp_stratum"`
	ShowPcDomain                  types.Bool `tfsdk:"show_pc_domain"`
	ShowReportFrequency           types.Bool `tfsdk:"show_report_frequency"`
	EnableAll                     types.Bool `tfsdk:"enable_all"`
	DisableAll                    types.Bool `tfsdk:"disable_all"`
}

var AdmingroupAdminShowCommandsAttrTypes = map[string]attr.Type{
	"show_admin_group_acl":             types.BoolType,
	"show_analytics_parameter":         types.BoolType,
	"show_arp":                         types.BoolType,
	"show_bfd":                         types.BoolType,
	"show_bgp":                         types.BoolType,
	"show_capacity":                    types.BoolType,
	"show_clusterd_info":               types.BoolType,
	"show_config":                      types.BoolType,
	"show_cpu":                         types.BoolType,
	"show_date":                        types.BoolType,
	"show_debug":                       types.BoolType,
	"show_debug_analytics":             types.BoolType,
	"show_delete_tasks_interval":       types.BoolType,
	"show_disk":                        types.BoolType,
	"show_hardware_type":               types.BoolType,
	"show_hardware_status":             types.BoolType,
	"show_hwid":                        types.BoolType,
	"show_ibtrap":                      types.BoolType,
	"show_hw_ident":                    types.BoolType,
	"show_log":                         types.BoolType,
	"show_logfiles":                    types.BoolType,
	"show_memory":                      types.BoolType,
	"show_ntp":                         types.BoolType,
	"show_reporting_user_capabilities": types.BoolType,
	"show_rpz_recursive_only":          types.BoolType,
	"show_scheduled":                   types.BoolType,
	"show_snmp":                        types.BoolType,
	"show_status":                      types.BoolType,
	"show_tech_support":                types.BoolType,
	"show_temperature":                 types.BoolType,
	"show_thresholdtrap":               types.BoolType,
	"show_upgrade_history":             types.BoolType,
	"show_uptime":                      types.BoolType,
	"show_version":                     types.BoolType,
	"show_analytics_database_dumps":    types.BoolType,
	"show_cores":                       types.BoolType,
	"show_coresummary":                 types.BoolType,
	"show_csp_threat_db":               types.BoolType,
	"show_hsm_group":                   types.BoolType,
	"show_hsm_info":                    types.BoolType,
	"show_pmap":                        types.BoolType,
	"show_process":                     types.BoolType,
	"show_pstack":                      types.BoolType,
	"show_safenet_support_info":        types.BoolType,
	"show_wred_stats":                  types.BoolType,
	"show_wred_status":                 types.BoolType,
	"show_ntp_stratum":                 types.BoolType,
	"show_pc_domain":                   types.BoolType,
	"show_report_frequency":            types.BoolType,
	"enable_all":                       types.BoolType,
	"disable_all":                      types.BoolType,
}

var AdmingroupAdminShowCommandsResourceSchemaAttributes = map[string]schema.Attribute{
	"show_admin_group_acl": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_analytics_parameter": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_arp": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_bfd": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_bgp": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_capacity": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_clusterd_info": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_config": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_cpu": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_date": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_debug": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_debug_analytics": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_delete_tasks_interval": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_disk": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_hardware_type": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_hardware_status": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_hwid": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_ibtrap": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_hw_ident": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_log": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_logfiles": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_memory": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_ntp": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_reporting_user_capabilities": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_rpz_recursive_only": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_scheduled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_snmp": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_status": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_tech_support": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_temperature": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_thresholdtrap": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_upgrade_history": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_uptime": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_version": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_analytics_database_dumps": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_cores": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_coresummary": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_csp_threat_db": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_hsm_group": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_hsm_info": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_pmap": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_process": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_pstack": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_safenet_support_info": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_wred_stats": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_wred_status": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_ntp_stratum": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_pc_domain": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_report_frequency": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"enable_all": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "If True then enable all fields",
	},
	"disable_all": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "If True then disable all fields",
	},
}

func ExpandAdmingroupAdminShowCommands(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.AdmingroupAdminShowCommands {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AdmingroupAdminShowCommandsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AdmingroupAdminShowCommandsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.AdmingroupAdminShowCommands {
	if m == nil {
		return nil
	}
	to := &security.AdmingroupAdminShowCommands{
		ShowAdminGroupAcl:             flex.ExpandBoolPointer(m.ShowAdminGroupAcl),
		ShowAnalyticsParameter:        flex.ExpandBoolPointer(m.ShowAnalyticsParameter),
		ShowArp:                       flex.ExpandBoolPointer(m.ShowArp),
		ShowBfd:                       flex.ExpandBoolPointer(m.ShowBfd),
		ShowBgp:                       flex.ExpandBoolPointer(m.ShowBgp),
		ShowCapacity:                  flex.ExpandBoolPointer(m.ShowCapacity),
		ShowClusterdInfo:              flex.ExpandBoolPointer(m.ShowClusterdInfo),
		ShowConfig:                    flex.ExpandBoolPointer(m.ShowConfig),
		ShowCpu:                       flex.ExpandBoolPointer(m.ShowCpu),
		ShowDate:                      flex.ExpandBoolPointer(m.ShowDate),
		ShowDebug:                     flex.ExpandBoolPointer(m.ShowDebug),
		ShowDebugAnalytics:            flex.ExpandBoolPointer(m.ShowDebugAnalytics),
		ShowDeleteTasksInterval:       flex.ExpandBoolPointer(m.ShowDeleteTasksInterval),
		ShowDisk:                      flex.ExpandBoolPointer(m.ShowDisk),
		ShowHardwareType:              flex.ExpandBoolPointer(m.ShowHardwareType),
		ShowHardwareStatus:            flex.ExpandBoolPointer(m.ShowHardwareStatus),
		ShowHwid:                      flex.ExpandBoolPointer(m.ShowHwid),
		ShowIbtrap:                    flex.ExpandBoolPointer(m.ShowIbtrap),
		ShowHwIdent:                   flex.ExpandBoolPointer(m.ShowHwIdent),
		ShowLog:                       flex.ExpandBoolPointer(m.ShowLog),
		ShowLogfiles:                  flex.ExpandBoolPointer(m.ShowLogfiles),
		ShowMemory:                    flex.ExpandBoolPointer(m.ShowMemory),
		ShowNtp:                       flex.ExpandBoolPointer(m.ShowNtp),
		ShowReportingUserCapabilities: flex.ExpandBoolPointer(m.ShowReportingUserCapabilities),
		ShowRpzRecursiveOnly:          flex.ExpandBoolPointer(m.ShowRpzRecursiveOnly),
		ShowScheduled:                 flex.ExpandBoolPointer(m.ShowScheduled),
		ShowSnmp:                      flex.ExpandBoolPointer(m.ShowSnmp),
		ShowStatus:                    flex.ExpandBoolPointer(m.ShowStatus),
		ShowTechSupport:               flex.ExpandBoolPointer(m.ShowTechSupport),
		ShowTemperature:               flex.ExpandBoolPointer(m.ShowTemperature),
		ShowThresholdtrap:             flex.ExpandBoolPointer(m.ShowThresholdtrap),
		ShowUpgradeHistory:            flex.ExpandBoolPointer(m.ShowUpgradeHistory),
		ShowUptime:                    flex.ExpandBoolPointer(m.ShowUptime),
		ShowVersion:                   flex.ExpandBoolPointer(m.ShowVersion),
		//ShowAnalyticsDatabaseDumps:    flex.ExpandBoolPointer(m.ShowAnalyticsDatabaseDumps),
		ShowCores:                     flex.ExpandBoolPointer(m.ShowCores),
		ShowCoresummary:               flex.ExpandBoolPointer(m.ShowCoresummary),
		ShowCspThreatDb:               flex.ExpandBoolPointer(m.ShowCspThreatDb),
		ShowHsmGroup:                  flex.ExpandBoolPointer(m.ShowHsmGroup),
		ShowHsmInfo:                   flex.ExpandBoolPointer(m.ShowHsmInfo),
		ShowPmap:                      flex.ExpandBoolPointer(m.ShowPmap),
		ShowProcess:                   flex.ExpandBoolPointer(m.ShowProcess),
		ShowPstack:                    flex.ExpandBoolPointer(m.ShowPstack),
		ShowSafenetSupportInfo:        flex.ExpandBoolPointer(m.ShowSafenetSupportInfo),
		ShowWredStats:                 flex.ExpandBoolPointer(m.ShowWredStats),
		ShowWredStatus:                flex.ExpandBoolPointer(m.ShowWredStatus),
		ShowNtpStratum:                flex.ExpandBoolPointer(m.ShowNtpStratum),
		ShowPcDomain:                  flex.ExpandBoolPointer(m.ShowPcDomain),
		ShowReportFrequency:           flex.ExpandBoolPointer(m.ShowReportFrequency),
	}
	return to
}

func FlattenAdmingroupAdminShowCommands(ctx context.Context, from *security.AdmingroupAdminShowCommands, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AdmingroupAdminShowCommandsAttrTypes)
	}
	m := AdmingroupAdminShowCommandsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AdmingroupAdminShowCommandsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AdmingroupAdminShowCommandsModel) Flatten(ctx context.Context, from *security.AdmingroupAdminShowCommands, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AdmingroupAdminShowCommandsModel{}
	}
	m.ShowAdminGroupAcl = types.BoolPointerValue(from.ShowAdminGroupAcl)
	m.ShowAnalyticsParameter = types.BoolPointerValue(from.ShowAnalyticsParameter)
	m.ShowArp = types.BoolPointerValue(from.ShowArp)
	m.ShowBfd = types.BoolPointerValue(from.ShowBfd)
	m.ShowBgp = types.BoolPointerValue(from.ShowBgp)
	m.ShowCapacity = types.BoolPointerValue(from.ShowCapacity)
	m.ShowClusterdInfo = types.BoolPointerValue(from.ShowClusterdInfo)
	m.ShowConfig = types.BoolPointerValue(from.ShowConfig)
	m.ShowCpu = types.BoolPointerValue(from.ShowCpu)
	m.ShowDate = types.BoolPointerValue(from.ShowDate)
	m.ShowDebug = types.BoolPointerValue(from.ShowDebug)
	m.ShowDebugAnalytics = types.BoolPointerValue(from.ShowDebugAnalytics)
	m.ShowDeleteTasksInterval = types.BoolPointerValue(from.ShowDeleteTasksInterval)
	m.ShowDisk = types.BoolPointerValue(from.ShowDisk)
	m.ShowHardwareType = types.BoolPointerValue(from.ShowHardwareType)
	m.ShowHardwareStatus = types.BoolPointerValue(from.ShowHardwareStatus)
	m.ShowHwid = types.BoolPointerValue(from.ShowHwid)
	m.ShowIbtrap = types.BoolPointerValue(from.ShowIbtrap)
	m.ShowHwIdent = types.BoolPointerValue(from.ShowHwIdent)
	m.ShowLog = types.BoolPointerValue(from.ShowLog)
	m.ShowLogfiles = types.BoolPointerValue(from.ShowLogfiles)
	m.ShowMemory = types.BoolPointerValue(from.ShowMemory)
	m.ShowNtp = types.BoolPointerValue(from.ShowNtp)
	m.ShowReportingUserCapabilities = types.BoolPointerValue(from.ShowReportingUserCapabilities)
	m.ShowRpzRecursiveOnly = types.BoolPointerValue(from.ShowRpzRecursiveOnly)
	m.ShowScheduled = types.BoolPointerValue(from.ShowScheduled)
	m.ShowSnmp = types.BoolPointerValue(from.ShowSnmp)
	m.ShowStatus = types.BoolPointerValue(from.ShowStatus)
	m.ShowTechSupport = types.BoolPointerValue(from.ShowTechSupport)
	m.ShowTemperature = types.BoolPointerValue(from.ShowTemperature)
	m.ShowThresholdtrap = types.BoolPointerValue(from.ShowThresholdtrap)
	m.ShowUpgradeHistory = types.BoolPointerValue(from.ShowUpgradeHistory)
	m.ShowUptime = types.BoolPointerValue(from.ShowUptime)
	m.ShowVersion = types.BoolPointerValue(from.ShowVersion)
	//m.ShowAnalyticsDatabaseDumps = types.BoolPointerValue(from.ShowAnalyticsDatabaseDumps)
	m.ShowCores = types.BoolPointerValue(from.ShowCores)
	m.ShowCoresummary = types.BoolPointerValue(from.ShowCoresummary)
	m.ShowCspThreatDb = types.BoolPointerValue(from.ShowCspThreatDb)
	m.ShowHsmGroup = types.BoolPointerValue(from.ShowHsmGroup)
	m.ShowHsmInfo = types.BoolPointerValue(from.ShowHsmInfo)
	m.ShowPmap = types.BoolPointerValue(from.ShowPmap)
	m.ShowProcess = types.BoolPointerValue(from.ShowProcess)
	m.ShowPstack = types.BoolPointerValue(from.ShowPstack)
	m.ShowSafenetSupportInfo = types.BoolPointerValue(from.ShowSafenetSupportInfo)
	m.ShowWredStats = types.BoolPointerValue(from.ShowWredStats)
	m.ShowWredStatus = types.BoolPointerValue(from.ShowWredStatus)
	m.ShowNtpStratum = types.BoolPointerValue(from.ShowNtpStratum)
	m.ShowPcDomain = types.BoolPointerValue(from.ShowPcDomain)
	m.ShowReportFrequency = types.BoolPointerValue(from.ShowReportFrequency)
	m.EnableAll = types.BoolPointerValue(from.EnableAll)
	m.DisableAll = types.BoolPointerValue(from.DisableAll)
}
