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

type GridModel struct {
	Ref                             types.String `tfsdk:"ref"`
	AllowRecursiveDeletion          types.String `tfsdk:"allow_recursive_deletion"`
	AuditLogFormat                  types.String `tfsdk:"audit_log_format"`
	AuditToSyslogEnable             types.Bool   `tfsdk:"audit_to_syslog_enable"`
	AutomatedTrafficCaptureSetting  types.Object `tfsdk:"automated_traffic_capture_setting"`
	ConsentBannerSetting            types.Object `tfsdk:"consent_banner_setting"`
	CspApiConfig                    types.Object `tfsdk:"csp_api_config"`
	CspGridSetting                  types.Object `tfsdk:"csp_grid_setting"`
	DenyMgmSnapshots                types.Bool   `tfsdk:"deny_mgm_snapshots"`
	DescendantsAction               types.Object `tfsdk:"descendants_action"`
	DnsResolverSetting              types.Object `tfsdk:"dns_resolver_setting"`
	Dscp                            types.Int64  `tfsdk:"dscp"`
	EmailSetting                    types.Object `tfsdk:"email_setting"`
	EnableFederation                types.Bool   `tfsdk:"enable_federation"`
	EnableForceSyncJoinTokenToGmc   types.Bool   `tfsdk:"enable_force_sync_join_token_to_gmc"`
	EnableGuiApiForLanVip           types.Bool   `tfsdk:"enable_gui_api_for_lan_vip"`
	EnableLom                       types.Bool   `tfsdk:"enable_lom"`
	EnableMemberRedirect            types.Bool   `tfsdk:"enable_member_redirect"`
	EnableRecycleBin                types.Bool   `tfsdk:"enable_recycle_bin"`
	EnableRirSwip                   types.Bool   `tfsdk:"enable_rir_swip"`
	ExternalSyslogBackupServers     types.List   `tfsdk:"external_syslog_backup_servers"`
	ExternalSyslogServerEnable      types.Bool   `tfsdk:"external_syslog_server_enable"`
	HttpProxyServerSetting          types.Object `tfsdk:"http_proxy_server_setting"`
	InformationalBannerSetting      types.Object `tfsdk:"informational_banner_setting"`
	IsGridVisualizationVisible      types.Bool   `tfsdk:"is_grid_visualization_visible"`
	LockoutSetting                  types.Object `tfsdk:"lockout_setting"`
	LomUsers                        types.List   `tfsdk:"lom_users"`
	MgmStrictDelegateMode           types.Bool   `tfsdk:"mgm_strict_delegate_mode"`
	MsSetting                       types.Object `tfsdk:"ms_setting"`
	Name                            types.String `tfsdk:"name"`
	NatGroups                       types.List   `tfsdk:"nat_groups"`
	NtpSetting                      types.Object `tfsdk:"ntp_setting"`
	ObjectsChangesTrackingSetting   types.Object `tfsdk:"objects_changes_tracking_setting"`
	PasswordSetting                 types.Object `tfsdk:"password_setting"`
	RestartBannerSetting            types.Object `tfsdk:"restart_banner_setting"`
	RestartStatus                   types.String `tfsdk:"restart_status"`
	RpzHitRateInterval              types.Int64  `tfsdk:"rpz_hit_rate_interval"`
	RpzHitRateMaxQuery              types.Int64  `tfsdk:"rpz_hit_rate_max_query"`
	RpzHitRateMinQuery              types.Int64  `tfsdk:"rpz_hit_rate_min_query"`
	ScheduledBackup                 types.Object `tfsdk:"scheduled_backup"`
	Secret                          types.String `tfsdk:"secret"`
	SecurityBannerSetting           types.Object `tfsdk:"security_banner_setting"`
	SecuritySetting                 types.Object `tfsdk:"security_setting"`
	ServiceStatus                   types.String `tfsdk:"service_status"`
	SnmpSetting                     types.Object `tfsdk:"snmp_setting"`
	SupportBundleDownloadTimeout    types.Int64  `tfsdk:"support_bundle_download_timeout"`
	SyslogFacility                  types.String `tfsdk:"syslog_facility"`
	SyslogServers                   types.List   `tfsdk:"syslog_servers"`
	SyslogSize                      types.Int64  `tfsdk:"syslog_size"`
	ThresholdTraps                  types.List   `tfsdk:"threshold_traps"`
	TimeZone                        types.String `tfsdk:"time_zone"`
	TokenUsageDelay                 types.Int64  `tfsdk:"token_usage_delay"`
	TrafficCaptureAuthDnsSetting    types.Object `tfsdk:"traffic_capture_auth_dns_setting"`
	TrafficCaptureChrSetting        types.Object `tfsdk:"traffic_capture_chr_setting"`
	TrafficCaptureQpsSetting        types.Object `tfsdk:"traffic_capture_qps_setting"`
	TrafficCaptureRecDnsSetting     types.Object `tfsdk:"traffic_capture_rec_dns_setting"`
	TrafficCaptureRecQueriesSetting types.Object `tfsdk:"traffic_capture_rec_queries_setting"`
	TrapNotifications               types.List   `tfsdk:"trap_notifications"`
	UpdatesDownloadMemberConfig     types.List   `tfsdk:"updates_download_member_config"`
	VpnPort                         types.Int64  `tfsdk:"vpn_port"`
}

var GridAttrTypes = map[string]attr.Type{
	"ref":                                 types.StringType,
	"allow_recursive_deletion":            types.StringType,
	"audit_log_format":                    types.StringType,
	"audit_to_syslog_enable":              types.BoolType,
	"automated_traffic_capture_setting":   types.ObjectType{AttrTypes: GridAutomatedTrafficCaptureSettingAttrTypes},
	"consent_banner_setting":              types.ObjectType{AttrTypes: GridConsentBannerSettingAttrTypes},
	"csp_api_config":                      types.ObjectType{AttrTypes: GridCspApiConfigAttrTypes},
	"csp_grid_setting":                    types.ObjectType{AttrTypes: GridCspGridSettingAttrTypes},
	"deny_mgm_snapshots":                  types.BoolType,
	"descendants_action":                  types.ObjectType{AttrTypes: GridDescendantsActionAttrTypes},
	"dns_resolver_setting":                types.ObjectType{AttrTypes: GridDnsResolverSettingAttrTypes},
	"dscp":                                types.Int64Type,
	"email_setting":                       types.ObjectType{AttrTypes: GridEmailSettingAttrTypes},
	"enable_federation":                   types.BoolType,
	"enable_force_sync_join_token_to_gmc": types.BoolType,
	"enable_gui_api_for_lan_vip":          types.BoolType,
	"enable_lom":                          types.BoolType,
	"enable_member_redirect":              types.BoolType,
	"enable_recycle_bin":                  types.BoolType,
	"enable_rir_swip":                     types.BoolType,
	"external_syslog_backup_servers":      types.ListType{ElemType: types.ObjectType{AttrTypes: GridExternalSyslogBackupServersAttrTypes}},
	"external_syslog_server_enable":       types.BoolType,
	"http_proxy_server_setting":           types.ObjectType{AttrTypes: GridHttpProxyServerSettingAttrTypes},
	"informational_banner_setting":        types.ObjectType{AttrTypes: GridInformationalBannerSettingAttrTypes},
	"is_grid_visualization_visible":       types.BoolType,
	"lockout_setting":                     types.ObjectType{AttrTypes: GridLockoutSettingAttrTypes},
	"lom_users":                           types.ListType{ElemType: types.ObjectType{AttrTypes: GridLomUsersAttrTypes}},
	"mgm_strict_delegate_mode":            types.BoolType,
	"ms_setting":                          types.ObjectType{AttrTypes: GridMsSettingAttrTypes},
	"name":                                types.StringType,
	"nat_groups":                          types.ListType{ElemType: types.StringType},
	"ntp_setting":                         types.ObjectType{AttrTypes: GridNtpSettingAttrTypes},
	"objects_changes_tracking_setting":    types.ObjectType{AttrTypes: GridObjectsChangesTrackingSettingAttrTypes},
	"password_setting":                    types.ObjectType{AttrTypes: GridPasswordSettingAttrTypes},
	"restart_banner_setting":              types.ObjectType{AttrTypes: GridRestartBannerSettingAttrTypes},
	"restart_status":                      types.StringType,
	"rpz_hit_rate_interval":               types.Int64Type,
	"rpz_hit_rate_max_query":              types.Int64Type,
	"rpz_hit_rate_min_query":              types.Int64Type,
	"scheduled_backup":                    types.ObjectType{AttrTypes: GridScheduledBackupAttrTypes},
	"secret":                              types.StringType,
	"security_banner_setting":             types.ObjectType{AttrTypes: GridSecurityBannerSettingAttrTypes},
	"security_setting":                    types.ObjectType{AttrTypes: GridSecuritySettingAttrTypes},
	"service_status":                      types.StringType,
	"snmp_setting":                        types.ObjectType{AttrTypes: GridSnmpSettingAttrTypes},
	"support_bundle_download_timeout":     types.Int64Type,
	"syslog_facility":                     types.StringType,
	"syslog_servers":                      types.ListType{ElemType: types.ObjectType{AttrTypes: GridSyslogServersAttrTypes}},
	"syslog_size":                         types.Int64Type,
	"threshold_traps":                     types.ListType{ElemType: types.ObjectType{AttrTypes: GridThresholdTrapsAttrTypes}},
	"time_zone":                           types.StringType,
	"token_usage_delay":                   types.Int64Type,
	"traffic_capture_auth_dns_setting":    types.ObjectType{AttrTypes: GridTrafficCaptureAuthDnsSettingAttrTypes},
	"traffic_capture_chr_setting":         types.ObjectType{AttrTypes: GridTrafficCaptureChrSettingAttrTypes},
	"traffic_capture_qps_setting":         types.ObjectType{AttrTypes: GridTrafficCaptureQpsSettingAttrTypes},
	"traffic_capture_rec_dns_setting":     types.ObjectType{AttrTypes: GridTrafficCaptureRecDnsSettingAttrTypes},
	"traffic_capture_rec_queries_setting": types.ObjectType{AttrTypes: GridTrafficCaptureRecQueriesSettingAttrTypes},
	"trap_notifications":                  types.ListType{ElemType: types.ObjectType{AttrTypes: GridTrapNotificationsAttrTypes}},
	"updates_download_member_config":      types.ListType{ElemType: types.ObjectType{AttrTypes: GridUpdatesDownloadMemberConfigAttrTypes}},
	"vpn_port":                            types.Int64Type,
}

var GridResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"allow_recursive_deletion": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The property to allow recursive deletion. Determines the users who can choose to perform recursive deletion on networks or zones from the GUI only.",
	},
	"audit_log_format": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Determines the audit log format.",
	},
	"audit_to_syslog_enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If set to True, audit log messages are also copied to the syslog.",
	},
	"automated_traffic_capture_setting": schema.SingleNestedAttribute{
		Attributes: GridAutomatedTrafficCaptureSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"consent_banner_setting": schema.SingleNestedAttribute{
		Attributes: GridConsentBannerSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"csp_api_config": schema.SingleNestedAttribute{
		Attributes: GridCspApiConfigResourceSchemaAttributes,
		Optional:   true,
	},
	"csp_grid_setting": schema.SingleNestedAttribute{
		Attributes: GridCspGridSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"deny_mgm_snapshots": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If set to True, the managed Grid will not send snapshots to the Multi-Grid Master.",
	},
	"descendants_action": schema.SingleNestedAttribute{
		Attributes: GridDescendantsActionResourceSchemaAttributes,
		Optional:   true,
	},
	"dns_resolver_setting": schema.SingleNestedAttribute{
		Attributes: GridDnsResolverSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"dscp": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The DSCP value. Valid values are integers between 0 and 63 inclusive.",
	},
	"email_setting": schema.SingleNestedAttribute{
		Attributes: GridEmailSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"enable_federation": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the federation feature is enabled or not. Test Setting will be performed for any change in enable_federation.",
	},
	"enable_force_sync_join_token_to_gmc": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the force sync join token from GM to GMC is enabled or not.",
	},
	"enable_gui_api_for_lan_vip": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If set to True, GUI and API access are enabled on the LAN/VIP port and MGMT port (if configured).",
	},
	"enable_lom": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the LOM functionality is enabled or not.",
	},
	"enable_member_redirect": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines redirections is enabled or not for members.",
	},
	"enable_recycle_bin": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the Recycle Bin is enabled or not.",
	},
	"enable_rir_swip": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the RIR/SWIP support is enabled or not.",
	},
	"external_syslog_backup_servers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: GridExternalSyslogBackupServersResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of external backup syslog servers.",
	},
	"external_syslog_server_enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If set to True, external syslog servers are enabled.",
	},
	"http_proxy_server_setting": schema.SingleNestedAttribute{
		Attributes: GridHttpProxyServerSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"informational_banner_setting": schema.SingleNestedAttribute{
		Attributes: GridInformationalBannerSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"is_grid_visualization_visible": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If set to True, graphical visualization of the Grid is enabled.",
	},
	"lockout_setting": schema.SingleNestedAttribute{
		Attributes: GridLockoutSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"lom_users": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: GridLomUsersResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of LOM users.",
	},
	"mgm_strict_delegate_mode": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if strict delegate mode for the Grid managed by the Master Grid is enabled or not.",
	},
	"ms_setting": schema.SingleNestedAttribute{
		Attributes: GridMsSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The grid name.",
	},
	"nat_groups": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of all Network Address Translation (NAT) groups configured on the Grid.",
	},
	"ntp_setting": schema.SingleNestedAttribute{
		Attributes: GridNtpSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"objects_changes_tracking_setting": schema.SingleNestedAttribute{
		Attributes: GridObjectsChangesTrackingSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"password_setting": schema.SingleNestedAttribute{
		Attributes: GridPasswordSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"restart_banner_setting": schema.SingleNestedAttribute{
		Attributes: GridRestartBannerSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"restart_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The restart status for the Grid.",
	},
	"rpz_hit_rate_interval": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The time interval (in seconds) that determines how often the appliance calculates the RPZ hit rate.",
	},
	"rpz_hit_rate_max_query": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The maximum number of incoming queries between the RPZ hit rate checks.",
	},
	"rpz_hit_rate_min_query": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The minimum number of incoming queries between the RPZ hit rate checks.",
	},
	"scheduled_backup": schema.SingleNestedAttribute{
		Attributes: GridScheduledBackupResourceSchemaAttributes,
		Optional:   true,
	},
	"secret": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The shared secret of the Grid. This is a write-only attribute.",
	},
	"security_banner_setting": schema.SingleNestedAttribute{
		Attributes: GridSecurityBannerSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"security_setting": schema.SingleNestedAttribute{
		Attributes: GridSecuritySettingResourceSchemaAttributes,
		Optional:   true,
	},
	"service_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Determines overall service status of the Grid.",
	},
	"snmp_setting": schema.SingleNestedAttribute{
		Attributes: GridSnmpSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"support_bundle_download_timeout": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Support bundle download timeout in seconds.",
	},
	"syslog_facility": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "If 'audit_to_syslog_enable' is set to True, the facility that determines the processes and daemons from which the log messages are generated.",
	},
	"syslog_servers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: GridSyslogServersResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of external syslog servers.",
	},
	"syslog_size": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The maximum size for the syslog file expressed in megabytes.",
	},
	"threshold_traps": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: GridThresholdTrapsResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "Determines the list of threshold traps. The user can only change the values for each trap or remove traps.",
	},
	"time_zone": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The time zone of the Grid. The UTC string that represents the time zone, such as \"US/Eastern\".",
	},
	"token_usage_delay": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The delayed usage (in minutes) of a permission token.",
	},
	"traffic_capture_auth_dns_setting": schema.SingleNestedAttribute{
		Attributes: GridTrafficCaptureAuthDnsSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"traffic_capture_chr_setting": schema.SingleNestedAttribute{
		Attributes: GridTrafficCaptureChrSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"traffic_capture_qps_setting": schema.SingleNestedAttribute{
		Attributes: GridTrafficCaptureQpsSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"traffic_capture_rec_dns_setting": schema.SingleNestedAttribute{
		Attributes: GridTrafficCaptureRecDnsSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"traffic_capture_rec_queries_setting": schema.SingleNestedAttribute{
		Attributes: GridTrafficCaptureRecQueriesSettingResourceSchemaAttributes,
		Optional:   true,
	},
	"trap_notifications": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: GridTrapNotificationsResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "Determines configuration of the trap notifications.",
	},
	"updates_download_member_config": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: GridUpdatesDownloadMemberConfigResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of member configuration structures, which provides information and settings for configuring the member that is responsible for downloading updates.",
	},
	"vpn_port": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The VPN port.",
	},
}

func ExpandGrid(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.Grid {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.Grid {
	if m == nil {
		return nil
	}
	to := &grid.Grid{
		Ref:                             flex.ExpandStringPointer(m.Ref),
		AllowRecursiveDeletion:          flex.ExpandStringPointer(m.AllowRecursiveDeletion),
		AuditLogFormat:                  flex.ExpandStringPointer(m.AuditLogFormat),
		AuditToSyslogEnable:             flex.ExpandBoolPointer(m.AuditToSyslogEnable),
		AutomatedTrafficCaptureSetting:  ExpandGridAutomatedTrafficCaptureSetting(ctx, m.AutomatedTrafficCaptureSetting, diags),
		ConsentBannerSetting:            ExpandGridConsentBannerSetting(ctx, m.ConsentBannerSetting, diags),
		CspApiConfig:                    ExpandGridCspApiConfig(ctx, m.CspApiConfig, diags),
		CspGridSetting:                  ExpandGridCspGridSetting(ctx, m.CspGridSetting, diags),
		DenyMgmSnapshots:                flex.ExpandBoolPointer(m.DenyMgmSnapshots),
		DescendantsAction:               ExpandGridDescendantsAction(ctx, m.DescendantsAction, diags),
		DnsResolverSetting:              ExpandGridDnsResolverSetting(ctx, m.DnsResolverSetting, diags),
		Dscp:                            flex.ExpandInt64Pointer(m.Dscp),
		EmailSetting:                    ExpandGridEmailSetting(ctx, m.EmailSetting, diags),
		EnableFederation:                flex.ExpandBoolPointer(m.EnableFederation),
		EnableForceSyncJoinTokenToGmc:   flex.ExpandBoolPointer(m.EnableForceSyncJoinTokenToGmc),
		EnableGuiApiForLanVip:           flex.ExpandBoolPointer(m.EnableGuiApiForLanVip),
		EnableLom:                       flex.ExpandBoolPointer(m.EnableLom),
		EnableMemberRedirect:            flex.ExpandBoolPointer(m.EnableMemberRedirect),
		EnableRecycleBin:                flex.ExpandBoolPointer(m.EnableRecycleBin),
		EnableRirSwip:                   flex.ExpandBoolPointer(m.EnableRirSwip),
		ExternalSyslogBackupServers:     flex.ExpandFrameworkListNestedBlock(ctx, m.ExternalSyslogBackupServers, diags, ExpandGridExternalSyslogBackupServers),
		ExternalSyslogServerEnable:      flex.ExpandBoolPointer(m.ExternalSyslogServerEnable),
		HttpProxyServerSetting:          ExpandGridHttpProxyServerSetting(ctx, m.HttpProxyServerSetting, diags),
		InformationalBannerSetting:      ExpandGridInformationalBannerSetting(ctx, m.InformationalBannerSetting, diags),
		IsGridVisualizationVisible:      flex.ExpandBoolPointer(m.IsGridVisualizationVisible),
		LockoutSetting:                  ExpandGridLockoutSetting(ctx, m.LockoutSetting, diags),
		LomUsers:                        flex.ExpandFrameworkListNestedBlock(ctx, m.LomUsers, diags, ExpandGridLomUsers),
		MgmStrictDelegateMode:           flex.ExpandBoolPointer(m.MgmStrictDelegateMode),
		MsSetting:                       ExpandGridMsSetting(ctx, m.MsSetting, diags),
		Name:                            flex.ExpandStringPointer(m.Name),
		NatGroups:                       flex.ExpandFrameworkListString(ctx, m.NatGroups, diags),
		NtpSetting:                      ExpandGridNtpSetting(ctx, m.NtpSetting, diags),
		ObjectsChangesTrackingSetting:   ExpandGridObjectsChangesTrackingSetting(ctx, m.ObjectsChangesTrackingSetting, diags),
		PasswordSetting:                 ExpandGridPasswordSetting(ctx, m.PasswordSetting, diags),
		RestartBannerSetting:            ExpandGridRestartBannerSetting(ctx, m.RestartBannerSetting, diags),
		RpzHitRateInterval:              flex.ExpandInt64Pointer(m.RpzHitRateInterval),
		RpzHitRateMaxQuery:              flex.ExpandInt64Pointer(m.RpzHitRateMaxQuery),
		RpzHitRateMinQuery:              flex.ExpandInt64Pointer(m.RpzHitRateMinQuery),
		ScheduledBackup:                 ExpandGridScheduledBackup(ctx, m.ScheduledBackup, diags),
		Secret:                          flex.ExpandStringPointer(m.Secret),
		SecurityBannerSetting:           ExpandGridSecurityBannerSetting(ctx, m.SecurityBannerSetting, diags),
		SecuritySetting:                 ExpandGridSecuritySetting(ctx, m.SecuritySetting, diags),
		SnmpSetting:                     ExpandGridSnmpSetting(ctx, m.SnmpSetting, diags),
		SupportBundleDownloadTimeout:    flex.ExpandInt64Pointer(m.SupportBundleDownloadTimeout),
		SyslogFacility:                  flex.ExpandStringPointer(m.SyslogFacility),
		SyslogServers:                   flex.ExpandFrameworkListNestedBlock(ctx, m.SyslogServers, diags, ExpandGridSyslogServers),
		SyslogSize:                      flex.ExpandInt64Pointer(m.SyslogSize),
		ThresholdTraps:                  flex.ExpandFrameworkListNestedBlock(ctx, m.ThresholdTraps, diags, ExpandGridThresholdTraps),
		TimeZone:                        flex.ExpandStringPointer(m.TimeZone),
		TokenUsageDelay:                 flex.ExpandInt64Pointer(m.TokenUsageDelay),
		TrafficCaptureAuthDnsSetting:    ExpandGridTrafficCaptureAuthDnsSetting(ctx, m.TrafficCaptureAuthDnsSetting, diags),
		TrafficCaptureChrSetting:        ExpandGridTrafficCaptureChrSetting(ctx, m.TrafficCaptureChrSetting, diags),
		TrafficCaptureQpsSetting:        ExpandGridTrafficCaptureQpsSetting(ctx, m.TrafficCaptureQpsSetting, diags),
		TrafficCaptureRecDnsSetting:     ExpandGridTrafficCaptureRecDnsSetting(ctx, m.TrafficCaptureRecDnsSetting, diags),
		TrafficCaptureRecQueriesSetting: ExpandGridTrafficCaptureRecQueriesSetting(ctx, m.TrafficCaptureRecQueriesSetting, diags),
		TrapNotifications:               flex.ExpandFrameworkListNestedBlock(ctx, m.TrapNotifications, diags, ExpandGridTrapNotifications),
		UpdatesDownloadMemberConfig:     flex.ExpandFrameworkListNestedBlock(ctx, m.UpdatesDownloadMemberConfig, diags, ExpandGridUpdatesDownloadMemberConfig),
		VpnPort:                         flex.ExpandInt64Pointer(m.VpnPort),
	}
	return to
}

func FlattenGrid(ctx context.Context, from *grid.Grid, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridAttrTypes)
	}
	m := GridModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridModel) Flatten(ctx context.Context, from *grid.Grid, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.AllowRecursiveDeletion = flex.FlattenStringPointer(from.AllowRecursiveDeletion)
	m.AuditLogFormat = flex.FlattenStringPointer(from.AuditLogFormat)
	m.AuditToSyslogEnable = types.BoolPointerValue(from.AuditToSyslogEnable)
	m.AutomatedTrafficCaptureSetting = FlattenGridAutomatedTrafficCaptureSetting(ctx, from.AutomatedTrafficCaptureSetting, diags)
	m.ConsentBannerSetting = FlattenGridConsentBannerSetting(ctx, from.ConsentBannerSetting, diags)
	m.CspApiConfig = FlattenGridCspApiConfig(ctx, from.CspApiConfig, diags)
	m.CspGridSetting = FlattenGridCspGridSetting(ctx, from.CspGridSetting, diags)
	m.DenyMgmSnapshots = types.BoolPointerValue(from.DenyMgmSnapshots)
	m.DescendantsAction = FlattenGridDescendantsAction(ctx, from.DescendantsAction, diags)
	m.DnsResolverSetting = FlattenGridDnsResolverSetting(ctx, from.DnsResolverSetting, diags)
	m.Dscp = flex.FlattenInt64Pointer(from.Dscp)
	m.EmailSetting = FlattenGridEmailSetting(ctx, from.EmailSetting, diags)
	m.EnableFederation = types.BoolPointerValue(from.EnableFederation)
	m.EnableForceSyncJoinTokenToGmc = types.BoolPointerValue(from.EnableForceSyncJoinTokenToGmc)
	m.EnableGuiApiForLanVip = types.BoolPointerValue(from.EnableGuiApiForLanVip)
	m.EnableLom = types.BoolPointerValue(from.EnableLom)
	m.EnableMemberRedirect = types.BoolPointerValue(from.EnableMemberRedirect)
	m.EnableRecycleBin = types.BoolPointerValue(from.EnableRecycleBin)
	m.EnableRirSwip = types.BoolPointerValue(from.EnableRirSwip)
	m.ExternalSyslogBackupServers = flex.FlattenFrameworkListNestedBlock(ctx, from.ExternalSyslogBackupServers, GridExternalSyslogBackupServersAttrTypes, diags, FlattenGridExternalSyslogBackupServers)
	m.ExternalSyslogServerEnable = types.BoolPointerValue(from.ExternalSyslogServerEnable)
	m.HttpProxyServerSetting = FlattenGridHttpProxyServerSetting(ctx, from.HttpProxyServerSetting, diags)
	m.InformationalBannerSetting = FlattenGridInformationalBannerSetting(ctx, from.InformationalBannerSetting, diags)
	m.IsGridVisualizationVisible = types.BoolPointerValue(from.IsGridVisualizationVisible)
	m.LockoutSetting = FlattenGridLockoutSetting(ctx, from.LockoutSetting, diags)
	m.LomUsers = flex.FlattenFrameworkListNestedBlock(ctx, from.LomUsers, GridLomUsersAttrTypes, diags, FlattenGridLomUsers)
	m.MgmStrictDelegateMode = types.BoolPointerValue(from.MgmStrictDelegateMode)
	m.MsSetting = FlattenGridMsSetting(ctx, from.MsSetting, diags)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.NatGroups = flex.FlattenFrameworkListString(ctx, from.NatGroups, diags)
	m.NtpSetting = FlattenGridNtpSetting(ctx, from.NtpSetting, diags)
	m.ObjectsChangesTrackingSetting = FlattenGridObjectsChangesTrackingSetting(ctx, from.ObjectsChangesTrackingSetting, diags)
	m.PasswordSetting = FlattenGridPasswordSetting(ctx, from.PasswordSetting, diags)
	m.RestartBannerSetting = FlattenGridRestartBannerSetting(ctx, from.RestartBannerSetting, diags)
	m.RestartStatus = flex.FlattenStringPointer(from.RestartStatus)
	m.RpzHitRateInterval = flex.FlattenInt64Pointer(from.RpzHitRateInterval)
	m.RpzHitRateMaxQuery = flex.FlattenInt64Pointer(from.RpzHitRateMaxQuery)
	m.RpzHitRateMinQuery = flex.FlattenInt64Pointer(from.RpzHitRateMinQuery)
	m.ScheduledBackup = FlattenGridScheduledBackup(ctx, from.ScheduledBackup, diags)
	m.Secret = flex.FlattenStringPointer(from.Secret)
	m.SecurityBannerSetting = FlattenGridSecurityBannerSetting(ctx, from.SecurityBannerSetting, diags)
	m.SecuritySetting = FlattenGridSecuritySetting(ctx, from.SecuritySetting, diags)
	m.ServiceStatus = flex.FlattenStringPointer(from.ServiceStatus)
	m.SnmpSetting = FlattenGridSnmpSetting(ctx, from.SnmpSetting, diags)
	m.SupportBundleDownloadTimeout = flex.FlattenInt64Pointer(from.SupportBundleDownloadTimeout)
	m.SyslogFacility = flex.FlattenStringPointer(from.SyslogFacility)
	m.SyslogServers = flex.FlattenFrameworkListNestedBlock(ctx, from.SyslogServers, GridSyslogServersAttrTypes, diags, FlattenGridSyslogServers)
	m.SyslogSize = flex.FlattenInt64Pointer(from.SyslogSize)
	m.ThresholdTraps = flex.FlattenFrameworkListNestedBlock(ctx, from.ThresholdTraps, GridThresholdTrapsAttrTypes, diags, FlattenGridThresholdTraps)
	m.TimeZone = flex.FlattenStringPointer(from.TimeZone)
	m.TokenUsageDelay = flex.FlattenInt64Pointer(from.TokenUsageDelay)
	m.TrafficCaptureAuthDnsSetting = FlattenGridTrafficCaptureAuthDnsSetting(ctx, from.TrafficCaptureAuthDnsSetting, diags)
	m.TrafficCaptureChrSetting = FlattenGridTrafficCaptureChrSetting(ctx, from.TrafficCaptureChrSetting, diags)
	m.TrafficCaptureQpsSetting = FlattenGridTrafficCaptureQpsSetting(ctx, from.TrafficCaptureQpsSetting, diags)
	m.TrafficCaptureRecDnsSetting = FlattenGridTrafficCaptureRecDnsSetting(ctx, from.TrafficCaptureRecDnsSetting, diags)
	m.TrafficCaptureRecQueriesSetting = FlattenGridTrafficCaptureRecQueriesSetting(ctx, from.TrafficCaptureRecQueriesSetting, diags)
	m.TrapNotifications = flex.FlattenFrameworkListNestedBlock(ctx, from.TrapNotifications, GridTrapNotificationsAttrTypes, diags, FlattenGridTrapNotifications)
	m.UpdatesDownloadMemberConfig = flex.FlattenFrameworkListNestedBlock(ctx, from.UpdatesDownloadMemberConfig, GridUpdatesDownloadMemberConfigAttrTypes, diags, FlattenGridUpdatesDownloadMemberConfig)
	m.VpnPort = flex.FlattenInt64Pointer(from.VpnPort)
}
