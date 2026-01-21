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

type AdmingroupAdminSetCommandsModel struct {
	SetAdminGroupAcl             types.Bool `tfsdk:"set_admin_group_acl"`
	EtBfd                        types.Bool `tfsdk:"et_bfd"`
	SetBfd                       types.Bool `tfsdk:"set_bfd"`
	SetBgp                       types.Bool `tfsdk:"set_bgp"`
	SetCleanMscache              types.Bool `tfsdk:"set_clean_mscache"`
	SetDebug                     types.Bool `tfsdk:"set_debug"`
	SetDebugAnalytics            types.Bool `tfsdk:"set_debug_analytics"`
	SetDeleteTasksInterval       types.Bool `tfsdk:"set_delete_tasks_interval"`
	SetDisableGuiOneClickSupport types.Bool `tfsdk:"set_disable_gui_one_click_support"`
	SetHardwareType              types.Bool `tfsdk:"set_hardware_type"`
	SetIbtrap                    types.Bool `tfsdk:"set_ibtrap"`
	SetHwIdent                   types.Bool `tfsdk:"set_hw_ident"`
	SetLines                     types.Bool `tfsdk:"set_lines"`
	SetMsMaxConnection           types.Bool `tfsdk:"set_ms_max_connection"`
	SetNosafemode                types.Bool `tfsdk:"set_nosafemode"`
	SetOcsp                      types.Bool `tfsdk:"set_ocsp"`
	SetPurgeRestartObjects       types.Bool `tfsdk:"set_purge_restart_objects"`
	SetReportingUserCapabilities types.Bool `tfsdk:"set_reporting_user_capabilities"`
	SetRpzRecursiveOnly          types.Bool `tfsdk:"set_rpz_recursive_only"`
	SetSafemode                  types.Bool `tfsdk:"set_safemode"`
	SetScheduled                 types.Bool `tfsdk:"set_scheduled"`
	SetSnmptrap                  types.Bool `tfsdk:"set_snmptrap"`
	SetSysname                   types.Bool `tfsdk:"set_sysname"`
	SetTerm                      types.Bool `tfsdk:"set_term"`
	SetThresholdtrap             types.Bool `tfsdk:"set_thresholdtrap"`
	SetExpertmode                types.Bool `tfsdk:"set_expertmode"`
	SetMaintenancemode           types.Bool `tfsdk:"set_maintenancemode"`
	SetTransferReportingData     types.Bool `tfsdk:"set_transfer_reporting_data"`
	SetTransferSupportbundle     types.Bool `tfsdk:"set_transfer_supportbundle"`
	SetAnalyticsDatabaseDump     types.Bool `tfsdk:"set_analytics_database_dump"`
	SetAnalyticsParameter        types.Bool `tfsdk:"set_analytics_parameter"`
	SetCollectOldLogs            types.Bool `tfsdk:"set_collect_old_logs"`
	SetCoreFilesQuota            types.Bool `tfsdk:"set_core_files_quota"`
	SetHsmGroup                  types.Bool `tfsdk:"set_hsm_group"`
	SetWred                      types.Bool `tfsdk:"set_wred"`
	SetEnableDohKeyLogging       types.Bool `tfsdk:"set_enable_doh_key_logging"`
	SetEnableDotKeyLogging       types.Bool `tfsdk:"set_enable_dot_key_logging"`
	SetHotfix                    types.Bool `tfsdk:"set_hotfix"`
	SetMgm                       types.Bool `tfsdk:"set_mgm"`
	SetNtpStratum                types.Bool `tfsdk:"set_ntp_stratum"`
	SetPcDomain                  types.Bool `tfsdk:"set_pc_domain"`
	SetReportFrequency           types.Bool `tfsdk:"set_report_frequency"`
	EnableAll                    types.Bool `tfsdk:"enable_all"`
	DisableAll                   types.Bool `tfsdk:"disable_all"`
}

var AdmingroupAdminSetCommandsAttrTypes = map[string]attr.Type{
	"set_admin_group_acl":               types.BoolType,
	"et_bfd":                            types.BoolType,
	"set_bfd":                           types.BoolType,
	"set_bgp":                           types.BoolType,
	"set_clean_mscache":                 types.BoolType,
	"set_debug":                         types.BoolType,
	"set_debug_analytics":               types.BoolType,
	"set_delete_tasks_interval":         types.BoolType,
	"set_disable_gui_one_click_support": types.BoolType,
	"set_hardware_type":                 types.BoolType,
	"set_ibtrap":                        types.BoolType,
	"set_hw_ident":                      types.BoolType,
	"set_lines":                         types.BoolType,
	"set_ms_max_connection":             types.BoolType,
	"set_nosafemode":                    types.BoolType,
	"set_ocsp":                          types.BoolType,
	"set_purge_restart_objects":         types.BoolType,
	"set_reporting_user_capabilities":   types.BoolType,
	"set_rpz_recursive_only":            types.BoolType,
	"set_safemode":                      types.BoolType,
	"set_scheduled":                     types.BoolType,
	"set_snmptrap":                      types.BoolType,
	"set_sysname":                       types.BoolType,
	"set_term":                          types.BoolType,
	"set_thresholdtrap":                 types.BoolType,
	"set_expertmode":                    types.BoolType,
	"set_maintenancemode":               types.BoolType,
	"set_transfer_reporting_data":       types.BoolType,
	"set_transfer_supportbundle":        types.BoolType,
	"set_analytics_database_dump":       types.BoolType,
	"set_analytics_parameter":           types.BoolType,
	"set_collect_old_logs":              types.BoolType,
	"set_core_files_quota":              types.BoolType,
	"set_hsm_group":                     types.BoolType,
	"set_wred":                          types.BoolType,
	"set_enable_doh_key_logging":        types.BoolType,
	"set_enable_dot_key_logging":        types.BoolType,
	"set_hotfix":                        types.BoolType,
	"set_mgm":                           types.BoolType,
	"set_ntp_stratum":                   types.BoolType,
	"set_pc_domain":                     types.BoolType,
	"set_report_frequency":              types.BoolType,
	"enable_all":                        types.BoolType,
	"disable_all":                       types.BoolType,
}

var AdmingroupAdminSetCommandsResourceSchemaAttributes = map[string]schema.Attribute{
	"set_admin_group_acl": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"et_bfd": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_bfd": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_bgp": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_clean_mscache": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_debug": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_debug_analytics": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_delete_tasks_interval": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_disable_gui_one_click_support": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_hardware_type": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_ibtrap": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_hw_ident": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_lines": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_ms_max_connection": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_nosafemode": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_ocsp": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_purge_restart_objects": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_reporting_user_capabilities": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_rpz_recursive_only": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_safemode": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_scheduled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_snmptrap": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_sysname": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_term": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_thresholdtrap": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_expertmode": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_maintenancemode": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_transfer_reporting_data": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_transfer_supportbundle": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_analytics_database_dump": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_analytics_parameter": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_collect_old_logs": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_core_files_quota": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_hsm_group": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_wred": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_enable_doh_key_logging": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_enable_dot_key_logging": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_hotfix": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_mgm": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_ntp_stratum": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_pc_domain": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_report_frequency": schema.BoolAttribute{
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

func ExpandAdmingroupAdminSetCommands(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.AdmingroupAdminSetCommands {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AdmingroupAdminSetCommandsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AdmingroupAdminSetCommandsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.AdmingroupAdminSetCommands {
	if m == nil {
		return nil
	}
	to := &security.AdmingroupAdminSetCommands{
		SetAdminGroupAcl:             flex.ExpandBoolPointer(m.SetAdminGroupAcl),
		EtBfd:                        flex.ExpandBoolPointer(m.EtBfd),
		SetBfd:                       flex.ExpandBoolPointer(m.SetBfd),
		SetBgp:                       flex.ExpandBoolPointer(m.SetBgp),
		SetCleanMscache:              flex.ExpandBoolPointer(m.SetCleanMscache),
		SetDebug:                     flex.ExpandBoolPointer(m.SetDebug),
		SetDebugAnalytics:            flex.ExpandBoolPointer(m.SetDebugAnalytics),
		SetDeleteTasksInterval:       flex.ExpandBoolPointer(m.SetDeleteTasksInterval),
		SetDisableGuiOneClickSupport: flex.ExpandBoolPointer(m.SetDisableGuiOneClickSupport),
		SetHardwareType:              flex.ExpandBoolPointer(m.SetHardwareType),
		SetIbtrap:                    flex.ExpandBoolPointer(m.SetIbtrap),
		SetHwIdent:                   flex.ExpandBoolPointer(m.SetHwIdent),
		SetLines:                     flex.ExpandBoolPointer(m.SetLines),
		SetMsMaxConnection:           flex.ExpandBoolPointer(m.SetMsMaxConnection),
		SetNosafemode:                flex.ExpandBoolPointer(m.SetNosafemode),
		SetOcsp:                      flex.ExpandBoolPointer(m.SetOcsp),
		SetPurgeRestartObjects:       flex.ExpandBoolPointer(m.SetPurgeRestartObjects),
		SetReportingUserCapabilities: flex.ExpandBoolPointer(m.SetReportingUserCapabilities),
		SetRpzRecursiveOnly:          flex.ExpandBoolPointer(m.SetRpzRecursiveOnly),
		SetSafemode:                  flex.ExpandBoolPointer(m.SetSafemode),
		SetScheduled:                 flex.ExpandBoolPointer(m.SetScheduled),
		SetSnmptrap:                  flex.ExpandBoolPointer(m.SetSnmptrap),
		SetSysname:                   flex.ExpandBoolPointer(m.SetSysname),
		SetTerm:                      flex.ExpandBoolPointer(m.SetTerm),
		SetThresholdtrap:             flex.ExpandBoolPointer(m.SetThresholdtrap),
		SetExpertmode:                flex.ExpandBoolPointer(m.SetExpertmode),
		SetMaintenancemode:           flex.ExpandBoolPointer(m.SetMaintenancemode),
		SetTransferReportingData:     flex.ExpandBoolPointer(m.SetTransferReportingData),
		SetTransferSupportbundle:     flex.ExpandBoolPointer(m.SetTransferSupportbundle),
		//SetAnalyticsDatabaseDump:     flex.ExpandBoolPointer(m.SetAnalyticsDatabaseDump),
		SetAnalyticsParameter:        flex.ExpandBoolPointer(m.SetAnalyticsParameter),
		SetCollectOldLogs:            flex.ExpandBoolPointer(m.SetCollectOldLogs),
		SetCoreFilesQuota:            flex.ExpandBoolPointer(m.SetCoreFilesQuota),
		SetHsmGroup:                  flex.ExpandBoolPointer(m.SetHsmGroup),
		SetWred:                      flex.ExpandBoolPointer(m.SetWred),
		SetEnableDohKeyLogging:       flex.ExpandBoolPointer(m.SetEnableDohKeyLogging),
		SetEnableDotKeyLogging:       flex.ExpandBoolPointer(m.SetEnableDotKeyLogging),
		SetHotfix:                    flex.ExpandBoolPointer(m.SetHotfix),
		SetMgm:                       flex.ExpandBoolPointer(m.SetMgm),
		SetNtpStratum:                flex.ExpandBoolPointer(m.SetNtpStratum),
		SetPcDomain:                  flex.ExpandBoolPointer(m.SetPcDomain),
		SetReportFrequency:           flex.ExpandBoolPointer(m.SetReportFrequency),
	}
	return to
}

func FlattenAdmingroupAdminSetCommands(ctx context.Context, from *security.AdmingroupAdminSetCommands, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AdmingroupAdminSetCommandsAttrTypes)
	}
	m := AdmingroupAdminSetCommandsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AdmingroupAdminSetCommandsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AdmingroupAdminSetCommandsModel) Flatten(ctx context.Context, from *security.AdmingroupAdminSetCommands, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AdmingroupAdminSetCommandsModel{}
	}
	m.SetAdminGroupAcl = types.BoolPointerValue(from.SetAdminGroupAcl)
	m.EtBfd = types.BoolPointerValue(from.EtBfd)
	m.SetBfd = types.BoolPointerValue(from.SetBfd)
	m.SetBgp = types.BoolPointerValue(from.SetBgp)
	m.SetCleanMscache = types.BoolPointerValue(from.SetCleanMscache)
	m.SetDebug = types.BoolPointerValue(from.SetDebug)
	m.SetDebugAnalytics = types.BoolPointerValue(from.SetDebugAnalytics)
	m.SetDeleteTasksInterval = types.BoolPointerValue(from.SetDeleteTasksInterval)
	m.SetDisableGuiOneClickSupport = types.BoolPointerValue(from.SetDisableGuiOneClickSupport)
	m.SetHardwareType = types.BoolPointerValue(from.SetHardwareType)
	m.SetIbtrap = types.BoolPointerValue(from.SetIbtrap)
	m.SetHwIdent = types.BoolPointerValue(from.SetHwIdent)
	m.SetLines = types.BoolPointerValue(from.SetLines)
	m.SetMsMaxConnection = types.BoolPointerValue(from.SetMsMaxConnection)
	m.SetNosafemode = types.BoolPointerValue(from.SetNosafemode)
	m.SetOcsp = types.BoolPointerValue(from.SetOcsp)
	m.SetPurgeRestartObjects = types.BoolPointerValue(from.SetPurgeRestartObjects)
	m.SetReportingUserCapabilities = types.BoolPointerValue(from.SetReportingUserCapabilities)
	m.SetRpzRecursiveOnly = types.BoolPointerValue(from.SetRpzRecursiveOnly)
	m.SetSafemode = types.BoolPointerValue(from.SetSafemode)
	m.SetScheduled = types.BoolPointerValue(from.SetScheduled)
	m.SetSnmptrap = types.BoolPointerValue(from.SetSnmptrap)
	m.SetSysname = types.BoolPointerValue(from.SetSysname)
	m.SetTerm = types.BoolPointerValue(from.SetTerm)
	m.SetThresholdtrap = types.BoolPointerValue(from.SetThresholdtrap)
	m.SetExpertmode = types.BoolPointerValue(from.SetExpertmode)
	m.SetMaintenancemode = types.BoolPointerValue(from.SetMaintenancemode)
	m.SetTransferReportingData = types.BoolPointerValue(from.SetTransferReportingData)
	m.SetTransferSupportbundle = types.BoolPointerValue(from.SetTransferSupportbundle)
	//m.SetAnalyticsDatabaseDump = types.BoolPointerValue(from.SetAnalyticsDatabaseDump)
	m.SetAnalyticsParameter = types.BoolPointerValue(from.SetAnalyticsParameter)
	m.SetCollectOldLogs = types.BoolPointerValue(from.SetCollectOldLogs)
	m.SetCoreFilesQuota = types.BoolPointerValue(from.SetCoreFilesQuota)
	m.SetHsmGroup = types.BoolPointerValue(from.SetHsmGroup)
	m.SetWred = types.BoolPointerValue(from.SetWred)
	m.SetEnableDohKeyLogging = types.BoolPointerValue(from.SetEnableDohKeyLogging)
	m.SetEnableDotKeyLogging = types.BoolPointerValue(from.SetEnableDotKeyLogging)
	m.SetHotfix = types.BoolPointerValue(from.SetHotfix)
	m.SetMgm = types.BoolPointerValue(from.SetMgm)
	m.SetNtpStratum = types.BoolPointerValue(from.SetNtpStratum)
	m.SetPcDomain = types.BoolPointerValue(from.SetPcDomain)
	m.SetReportFrequency = types.BoolPointerValue(from.SetReportFrequency)
	m.EnableAll = types.BoolPointerValue(from.EnableAll)
	m.DisableAll = types.BoolPointerValue(from.DisableAll)
}
