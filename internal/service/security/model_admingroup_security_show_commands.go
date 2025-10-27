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

type AdmingroupSecurityShowCommandsModel struct {
	ShowFipsMode                types.Bool `tfsdk:"show_fips_mode"`
	ShowCcMode                  types.Bool `tfsdk:"show_cc_mode"`
	ShowCertificateAuthAdmins   types.Bool `tfsdk:"show_certificate_auth_admins"`
	ShowCertificateAuthServices types.Bool `tfsdk:"show_certificate_auth_services"`
	ShowCheckAuthNs             types.Bool `tfsdk:"show_check_auth_ns"`
	ShowCheckSslCertificate     types.Bool `tfsdk:"show_check_ssl_certificate"`
	ShowSecurity                types.Bool `tfsdk:"show_security"`
	ShowSessionTimeout          types.Bool `tfsdk:"show_session_timeout"`
	ShowSubscriberSecureData    types.Bool `tfsdk:"show_subscriber_secure_data"`
	ShowSupportAccess           types.Bool `tfsdk:"show_support_access"`
	ShowVpnCertDates            types.Bool `tfsdk:"show_vpn_cert_dates"`
	ShowAdp                     types.Bool `tfsdk:"show_adp"`
	ShowAdpDebug                types.Bool `tfsdk:"show_adp_debug"`
	ShowSupportTimeout          types.Bool `tfsdk:"show_support_timeout"`
	EnableAll                   types.Bool `tfsdk:"enable_all"`
	DisableAll                  types.Bool `tfsdk:"disable_all"`
}

var AdmingroupSecurityShowCommandsAttrTypes = map[string]attr.Type{
	"show_fips_mode":                 types.BoolType,
	"show_cc_mode":                   types.BoolType,
	"show_certificate_auth_admins":   types.BoolType,
	"show_certificate_auth_services": types.BoolType,
	"show_check_auth_ns":             types.BoolType,
	"show_check_ssl_certificate":     types.BoolType,
	"show_security":                  types.BoolType,
	"show_session_timeout":           types.BoolType,
	"show_subscriber_secure_data":    types.BoolType,
	"show_support_access":            types.BoolType,
	"show_vpn_cert_dates":            types.BoolType,
	"show_adp":                       types.BoolType,
	"show_adp_debug":                 types.BoolType,
	"show_support_timeout":           types.BoolType,
	"enable_all":                     types.BoolType,
	"disable_all":                    types.BoolType,
}

var AdmingroupSecurityShowCommandsResourceSchemaAttributes = map[string]schema.Attribute{
	"show_fips_mode": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_cc_mode": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_certificate_auth_admins": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_certificate_auth_services": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_check_auth_ns": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_check_ssl_certificate": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_security": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_session_timeout": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_subscriber_secure_data": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_support_access": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_vpn_cert_dates": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_adp": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_adp_debug": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_support_timeout": schema.BoolAttribute{
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

func ExpandAdmingroupSecurityShowCommands(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.AdmingroupSecurityShowCommands {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AdmingroupSecurityShowCommandsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AdmingroupSecurityShowCommandsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.AdmingroupSecurityShowCommands {
	if m == nil {
		return nil
	}
	to := &security.AdmingroupSecurityShowCommands{
		ShowFipsMode:                flex.ExpandBoolPointer(m.ShowFipsMode),
		ShowCcMode:                  flex.ExpandBoolPointer(m.ShowCcMode),
		ShowCertificateAuthAdmins:   flex.ExpandBoolPointer(m.ShowCertificateAuthAdmins),
		ShowCertificateAuthServices: flex.ExpandBoolPointer(m.ShowCertificateAuthServices),
		ShowCheckAuthNs:             flex.ExpandBoolPointer(m.ShowCheckAuthNs),
		ShowCheckSslCertificate:     flex.ExpandBoolPointer(m.ShowCheckSslCertificate),
		ShowSecurity:                flex.ExpandBoolPointer(m.ShowSecurity),
		ShowSessionTimeout:          flex.ExpandBoolPointer(m.ShowSessionTimeout),
		ShowSubscriberSecureData:    flex.ExpandBoolPointer(m.ShowSubscriberSecureData),
		ShowSupportAccess:           flex.ExpandBoolPointer(m.ShowSupportAccess),
		ShowVpnCertDates:            flex.ExpandBoolPointer(m.ShowVpnCertDates),
		ShowAdp:                     flex.ExpandBoolPointer(m.ShowAdp),
		ShowAdpDebug:                flex.ExpandBoolPointer(m.ShowAdpDebug),
		ShowSupportTimeout:          flex.ExpandBoolPointer(m.ShowSupportTimeout),
		EnableAll:                   flex.ExpandBoolPointer(m.EnableAll),
		DisableAll:                  flex.ExpandBoolPointer(m.DisableAll),
	}
	return to
}

func FlattenAdmingroupSecurityShowCommands(ctx context.Context, from *security.AdmingroupSecurityShowCommands, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AdmingroupSecurityShowCommandsAttrTypes)
	}
	m := AdmingroupSecurityShowCommandsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AdmingroupSecurityShowCommandsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AdmingroupSecurityShowCommandsModel) Flatten(ctx context.Context, from *security.AdmingroupSecurityShowCommands, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AdmingroupSecurityShowCommandsModel{}
	}
	m.ShowFipsMode = types.BoolPointerValue(from.ShowFipsMode)
	m.ShowCcMode = types.BoolPointerValue(from.ShowCcMode)
	m.ShowCertificateAuthAdmins = types.BoolPointerValue(from.ShowCertificateAuthAdmins)
	m.ShowCertificateAuthServices = types.BoolPointerValue(from.ShowCertificateAuthServices)
	m.ShowCheckAuthNs = types.BoolPointerValue(from.ShowCheckAuthNs)
	m.ShowCheckSslCertificate = types.BoolPointerValue(from.ShowCheckSslCertificate)
	m.ShowSecurity = types.BoolPointerValue(from.ShowSecurity)
	m.ShowSessionTimeout = types.BoolPointerValue(from.ShowSessionTimeout)
	m.ShowSubscriberSecureData = types.BoolPointerValue(from.ShowSubscriberSecureData)
	m.ShowSupportAccess = types.BoolPointerValue(from.ShowSupportAccess)
	m.ShowVpnCertDates = types.BoolPointerValue(from.ShowVpnCertDates)
	m.ShowAdp = types.BoolPointerValue(from.ShowAdp)
	m.ShowAdpDebug = types.BoolPointerValue(from.ShowAdpDebug)
	m.ShowSupportTimeout = types.BoolPointerValue(from.ShowSupportTimeout)
	m.EnableAll = types.BoolPointerValue(from.EnableAll)
	m.DisableAll = types.BoolPointerValue(from.DisableAll)
}
