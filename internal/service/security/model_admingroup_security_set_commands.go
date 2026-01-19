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

type AdmingroupSecuritySetCommandsModel struct {
	SetAdp                          types.Bool `tfsdk:"set_adp"`
	SetApacheHttpsCert              types.Bool `tfsdk:"set_apache_https_cert"`
	SetCcMode                       types.Bool `tfsdk:"set_cc_mode"`
	SetCertificateAuthAdmins        types.Bool `tfsdk:"set_certificate_auth_admins"`
	SetCertificateAuthServices      types.Bool `tfsdk:"set_certificate_auth_services"`
	SetCheckAuthNs                  types.Bool `tfsdk:"set_check_auth_ns"`
	SetCheckSslCertificate          types.Bool `tfsdk:"set_check_ssl_certificate"`
	SetDisableHttpsCertRegeneration types.Bool `tfsdk:"set_disable_https_cert_regeneration"`
	SetFipsMode                     types.Bool `tfsdk:"set_fips_mode"`
	SetReportingCert                types.Bool `tfsdk:"set_reporting_cert"`
	SetSecurity                     types.Bool `tfsdk:"set_security"`
	SetSessionTimeout               types.Bool `tfsdk:"set_session_timeout"`
	SetSubscriberSecureData         types.Bool `tfsdk:"set_subscriber_secure_data"`
	SetSupportAccess                types.Bool `tfsdk:"set_support_access"`
	SetSupportInstall               types.Bool `tfsdk:"set_support_install"`
	SetAdpDebug                     types.Bool `tfsdk:"set_adp_debug"`
	SetSupportTimeout               types.Bool `tfsdk:"set_support_timeout"`
	SetUpdateRabbitmqPassword       types.Bool `tfsdk:"set_update_rabbitmq_password"`
	EnableAll                       types.Bool `tfsdk:"enable_all"`
	DisableAll                      types.Bool `tfsdk:"disable_all"`
}

var AdmingroupSecuritySetCommandsAttrTypes = map[string]attr.Type{
	"set_adp":                             types.BoolType,
	"set_apache_https_cert":               types.BoolType,
	"set_cc_mode":                         types.BoolType,
	"set_certificate_auth_admins":         types.BoolType,
	"set_certificate_auth_services":       types.BoolType,
	"set_check_auth_ns":                   types.BoolType,
	"set_check_ssl_certificate":           types.BoolType,
	"set_disable_https_cert_regeneration": types.BoolType,
	"set_fips_mode":                       types.BoolType,
	"set_reporting_cert":                  types.BoolType,
	"set_security":                        types.BoolType,
	"set_session_timeout":                 types.BoolType,
	"set_subscriber_secure_data":          types.BoolType,
	"set_support_access":                  types.BoolType,
	"set_support_install":                 types.BoolType,
	"set_adp_debug":                       types.BoolType,
	"set_support_timeout":                 types.BoolType,
	"set_update_rabbitmq_password":        types.BoolType,
	"enable_all":                          types.BoolType,
	"disable_all":                         types.BoolType,
}

var AdmingroupSecuritySetCommandsResourceSchemaAttributes = map[string]schema.Attribute{
	"set_adp": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_apache_https_cert": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_cc_mode": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_certificate_auth_admins": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_certificate_auth_services": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_check_auth_ns": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_check_ssl_certificate": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_disable_https_cert_regeneration": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_fips_mode": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_reporting_cert": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_security": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_session_timeout": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_subscriber_secure_data": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_support_access": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_support_install": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_adp_debug": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_support_timeout": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_update_rabbitmq_password": schema.BoolAttribute{
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

func ExpandAdmingroupSecuritySetCommands(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.AdmingroupSecuritySetCommands {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AdmingroupSecuritySetCommandsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AdmingroupSecuritySetCommandsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.AdmingroupSecuritySetCommands {
	if m == nil {
		return nil
	}
	to := &security.AdmingroupSecuritySetCommands{
		SetAdp:                          flex.ExpandBoolPointer(m.SetAdp),
		SetApacheHttpsCert:              flex.ExpandBoolPointer(m.SetApacheHttpsCert),
		SetCcMode:                       flex.ExpandBoolPointer(m.SetCcMode),
		SetCertificateAuthAdmins:        flex.ExpandBoolPointer(m.SetCertificateAuthAdmins),
		SetCertificateAuthServices:      flex.ExpandBoolPointer(m.SetCertificateAuthServices),
		SetCheckAuthNs:                  flex.ExpandBoolPointer(m.SetCheckAuthNs),
		SetCheckSslCertificate:          flex.ExpandBoolPointer(m.SetCheckSslCertificate),
		SetDisableHttpsCertRegeneration: flex.ExpandBoolPointer(m.SetDisableHttpsCertRegeneration),
		SetFipsMode:                     flex.ExpandBoolPointer(m.SetFipsMode),
		SetReportingCert:                flex.ExpandBoolPointer(m.SetReportingCert),
		SetSecurity:                     flex.ExpandBoolPointer(m.SetSecurity),
		SetSessionTimeout:               flex.ExpandBoolPointer(m.SetSessionTimeout),
		SetSubscriberSecureData:         flex.ExpandBoolPointer(m.SetSubscriberSecureData),
		SetSupportAccess:                flex.ExpandBoolPointer(m.SetSupportAccess),
		SetSupportInstall:               flex.ExpandBoolPointer(m.SetSupportInstall),
		SetAdpDebug:                     flex.ExpandBoolPointer(m.SetAdpDebug),
		SetSupportTimeout:               flex.ExpandBoolPointer(m.SetSupportTimeout),
		SetUpdateRabbitmqPassword:       flex.ExpandBoolPointer(m.SetUpdateRabbitmqPassword),
	}
	return to
}

func FlattenAdmingroupSecuritySetCommands(ctx context.Context, from *security.AdmingroupSecuritySetCommands, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AdmingroupSecuritySetCommandsAttrTypes)
	}
	m := AdmingroupSecuritySetCommandsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AdmingroupSecuritySetCommandsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AdmingroupSecuritySetCommandsModel) Flatten(ctx context.Context, from *security.AdmingroupSecuritySetCommands, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AdmingroupSecuritySetCommandsModel{}
	}
	m.SetAdp = types.BoolPointerValue(from.SetAdp)
	m.SetApacheHttpsCert = types.BoolPointerValue(from.SetApacheHttpsCert)
	m.SetCcMode = types.BoolPointerValue(from.SetCcMode)
	m.SetCertificateAuthAdmins = types.BoolPointerValue(from.SetCertificateAuthAdmins)
	m.SetCertificateAuthServices = types.BoolPointerValue(from.SetCertificateAuthServices)
	m.SetCheckAuthNs = types.BoolPointerValue(from.SetCheckAuthNs)
	m.SetCheckSslCertificate = types.BoolPointerValue(from.SetCheckSslCertificate)
	m.SetDisableHttpsCertRegeneration = types.BoolPointerValue(from.SetDisableHttpsCertRegeneration)
	m.SetFipsMode = types.BoolPointerValue(from.SetFipsMode)
	m.SetReportingCert = types.BoolPointerValue(from.SetReportingCert)
	m.SetSecurity = types.BoolPointerValue(from.SetSecurity)
	m.SetSessionTimeout = types.BoolPointerValue(from.SetSessionTimeout)
	m.SetSubscriberSecureData = types.BoolPointerValue(from.SetSubscriberSecureData)
	m.SetSupportAccess = types.BoolPointerValue(from.SetSupportAccess)
	m.SetSupportInstall = types.BoolPointerValue(from.SetSupportInstall)
	m.SetAdpDebug = types.BoolPointerValue(from.SetAdpDebug)
	m.SetSupportTimeout = types.BoolPointerValue(from.SetSupportTimeout)
	m.SetUpdateRabbitmqPassword = types.BoolPointerValue(from.SetUpdateRabbitmqPassword)
	m.EnableAll = types.BoolPointerValue(from.EnableAll)
	m.DisableAll = types.BoolPointerValue(from.DisableAll)
}
