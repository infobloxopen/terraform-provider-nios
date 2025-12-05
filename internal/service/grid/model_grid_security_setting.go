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

type GridSecuritySettingModel struct {
	AuditLogRollingEnable             types.Bool   `tfsdk:"audit_log_rolling_enable"`
	AdminAccessItems                  types.List   `tfsdk:"admin_access_items"`
	HttpRedirectEnable                types.Bool   `tfsdk:"http_redirect_enable"`
	LoginBannerEnable                 types.Bool   `tfsdk:"login_banner_enable"`
	LoginBannerText                   types.String `tfsdk:"login_banner_text"`
	RemoteConsoleAccessEnable         types.Bool   `tfsdk:"remote_console_access_enable"`
	SecurityAccessEnable              types.Bool   `tfsdk:"security_access_enable"`
	SecurityAccessRemoteConsoleEnable types.Bool   `tfsdk:"security_access_remote_console_enable"`
	SessionTimeout                    types.Int64  `tfsdk:"session_timeout"`
	SshPermEnable                     types.Bool   `tfsdk:"ssh_perm_enable"`
	SupportAccessEnable               types.Bool   `tfsdk:"support_access_enable"`
	SupportAccessInfo                 types.String `tfsdk:"support_access_info"`
	DisableConcurrentLogin            types.Bool   `tfsdk:"disable_concurrent_login"`
	InactivityLockoutSetting          types.Object `tfsdk:"inactivity_lockout_setting"`
}

var GridSecuritySettingAttrTypes = map[string]attr.Type{
	"audit_log_rolling_enable":              types.BoolType,
	"admin_access_items":                    types.ListType{ElemType: types.ObjectType{AttrTypes: GridsecuritysettingAdminAccessItemsAttrTypes}},
	"http_redirect_enable":                  types.BoolType,
	"login_banner_enable":                   types.BoolType,
	"login_banner_text":                     types.StringType,
	"remote_console_access_enable":          types.BoolType,
	"security_access_enable":                types.BoolType,
	"security_access_remote_console_enable": types.BoolType,
	"session_timeout":                       types.Int64Type,
	"ssh_perm_enable":                       types.BoolType,
	"support_access_enable":                 types.BoolType,
	"support_access_info":                   types.StringType,
	"disable_concurrent_login":              types.BoolType,
	"inactivity_lockout_setting":            types.ObjectType{AttrTypes: GridsecuritysettingInactivityLockoutSettingAttrTypes},
}

var GridSecuritySettingResourceSchemaAttributes = map[string]schema.Attribute{
	"audit_log_rolling_enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If set to True, rolling of audit logs is enabled.",
	},
	"admin_access_items": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: GridsecuritysettingAdminAccessItemsResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "A list of access control settings used for security access.",
	},
	"http_redirect_enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If set to True, HTTP connections are redirected to HTTPS.",
	},
	"login_banner_enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If set to True, the login banner is enabled.",
	},
	"login_banner_text": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The login banner text.",
	},
	"remote_console_access_enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If set to True, superuser admins can access the Infoblox CLI from a remote location using an SSH (Secure Shell) v2 client.",
	},
	"security_access_enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If set to True, HTTP access restrictions are enabled.",
	},
	"security_access_remote_console_enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If set to True, remote console access restrictions will be enabled.",
	},
	"session_timeout": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The session timeout interval in seconds.",
	},
	"ssh_perm_enable": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "If set to False, SSH access is permanently disabled.",
	},
	"support_access_enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If set to True, support access for the Grid has been enabled.",
	},
	"support_access_info": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Information string to be used for support access requests.",
	},
	"disable_concurrent_login": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Whether concurrent login allowed gridlevel",
	},
	"inactivity_lockout_setting": schema.SingleNestedAttribute{
		Attributes: GridsecuritysettingInactivityLockoutSettingResourceSchemaAttributes,
		Optional:   true,
	},
}

func ExpandGridSecuritySetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridSecuritySetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridSecuritySettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridSecuritySettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridSecuritySetting {
	if m == nil {
		return nil
	}
	to := &grid.GridSecuritySetting{
		AuditLogRollingEnable:             flex.ExpandBoolPointer(m.AuditLogRollingEnable),
		AdminAccessItems:                  flex.ExpandFrameworkListNestedBlock(ctx, m.AdminAccessItems, diags, ExpandGridsecuritysettingAdminAccessItems),
		HttpRedirectEnable:                flex.ExpandBoolPointer(m.HttpRedirectEnable),
		LoginBannerEnable:                 flex.ExpandBoolPointer(m.LoginBannerEnable),
		LoginBannerText:                   flex.ExpandStringPointer(m.LoginBannerText),
		RemoteConsoleAccessEnable:         flex.ExpandBoolPointer(m.RemoteConsoleAccessEnable),
		SecurityAccessEnable:              flex.ExpandBoolPointer(m.SecurityAccessEnable),
		SecurityAccessRemoteConsoleEnable: flex.ExpandBoolPointer(m.SecurityAccessRemoteConsoleEnable),
		SessionTimeout:                    flex.ExpandInt64Pointer(m.SessionTimeout),
		SupportAccessEnable:               flex.ExpandBoolPointer(m.SupportAccessEnable),
		SupportAccessInfo:                 flex.ExpandStringPointer(m.SupportAccessInfo),
		DisableConcurrentLogin:            flex.ExpandBoolPointer(m.DisableConcurrentLogin),
		InactivityLockoutSetting:          ExpandGridsecuritysettingInactivityLockoutSetting(ctx, m.InactivityLockoutSetting, diags),
	}
	return to
}

func FlattenGridSecuritySetting(ctx context.Context, from *grid.GridSecuritySetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridSecuritySettingAttrTypes)
	}
	m := GridSecuritySettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridSecuritySettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridSecuritySettingModel) Flatten(ctx context.Context, from *grid.GridSecuritySetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridSecuritySettingModel{}
	}
	m.AuditLogRollingEnable = types.BoolPointerValue(from.AuditLogRollingEnable)
	m.AdminAccessItems = flex.FlattenFrameworkListNestedBlock(ctx, from.AdminAccessItems, GridsecuritysettingAdminAccessItemsAttrTypes, diags, FlattenGridsecuritysettingAdminAccessItems)
	m.HttpRedirectEnable = types.BoolPointerValue(from.HttpRedirectEnable)
	m.LoginBannerEnable = types.BoolPointerValue(from.LoginBannerEnable)
	m.LoginBannerText = flex.FlattenStringPointer(from.LoginBannerText)
	m.RemoteConsoleAccessEnable = types.BoolPointerValue(from.RemoteConsoleAccessEnable)
	m.SecurityAccessEnable = types.BoolPointerValue(from.SecurityAccessEnable)
	m.SecurityAccessRemoteConsoleEnable = types.BoolPointerValue(from.SecurityAccessRemoteConsoleEnable)
	m.SessionTimeout = flex.FlattenInt64Pointer(from.SessionTimeout)
	m.SshPermEnable = types.BoolPointerValue(from.SshPermEnable)
	m.SupportAccessEnable = types.BoolPointerValue(from.SupportAccessEnable)
	m.SupportAccessInfo = flex.FlattenStringPointer(from.SupportAccessInfo)
	m.DisableConcurrentLogin = types.BoolPointerValue(from.DisableConcurrentLogin)
	m.InactivityLockoutSetting = FlattenGridsecuritysettingInactivityLockoutSetting(ctx, from.InactivityLockoutSetting, diags)
}
