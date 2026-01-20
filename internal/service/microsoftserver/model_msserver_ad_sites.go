package microsoftserver

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/microsoftserver"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type MsserverAdSitesModel struct {
	UseDefaultIpSiteLink       types.Bool   `tfsdk:"use_default_ip_site_link"`
	DefaultIpSiteLink          types.String `tfsdk:"default_ip_site_link"`
	UseLogin                   types.Bool   `tfsdk:"use_login"`
	LoginName                  types.String `tfsdk:"login_name"`
	LoginPassword              types.String `tfsdk:"login_password"`
	UseSynchronizationMinDelay types.Bool   `tfsdk:"use_synchronization_min_delay"`
	SynchronizationMinDelay    types.Int64  `tfsdk:"synchronization_min_delay"`
	UseLdapTimeout             types.Bool   `tfsdk:"use_ldap_timeout"`
	LdapTimeout                types.Int64  `tfsdk:"ldap_timeout"`
	LdapAuthPort               types.Int64  `tfsdk:"ldap_auth_port"`
	LdapEncryption             types.String `tfsdk:"ldap_encryption"`
	Managed                    types.Bool   `tfsdk:"managed"`
	ReadOnly                   types.Bool   `tfsdk:"read_only"`
	LastSyncTs                 types.Int64  `tfsdk:"last_sync_ts"`
	LastSyncStatus             types.String `tfsdk:"last_sync_status"`
	LastSyncDetail             types.String `tfsdk:"last_sync_detail"`
	SupportsIpv6               types.Bool   `tfsdk:"supports_ipv6"`
}

var MsserverAdSitesAttrTypes = map[string]attr.Type{
	"use_default_ip_site_link":      types.BoolType,
	"default_ip_site_link":          types.StringType,
	"use_login":                     types.BoolType,
	"login_name":                    types.StringType,
	"login_password":                types.StringType,
	"use_synchronization_min_delay": types.BoolType,
	"synchronization_min_delay":     types.Int64Type,
	"use_ldap_timeout":              types.BoolType,
	"ldap_timeout":                  types.Int64Type,
	"ldap_auth_port":                types.Int64Type,
	"ldap_encryption":               types.StringType,
	"managed":                       types.BoolType,
	"read_only":                     types.BoolType,
	"last_sync_ts":                  types.Int64Type,
	"last_sync_status":              types.StringType,
	"last_sync_detail":              types.StringType,
	"supports_ipv6":                 types.BoolType,
}

var MsserverAdSitesResourceSchemaAttributes = map[string]schema.Attribute{
	"use_default_ip_site_link": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Flag to override MS Server default IP site link",
	},
	"default_ip_site_link": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Default IP site link for sites created from NIOS",
	},
	"use_login": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Flag to override login name and password from the MS Server",
	},
	"login_name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Microsoft Server login name, with optional",
	},
	"login_password": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Microsoft Server login password.",
	},
	"use_synchronization_min_delay": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Flag to override synchronization interval from the MS Server",
	},
	"synchronization_min_delay": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Minimum number of minutes between two synchronizations",
	},
	"use_ldap_timeout": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Flag to override cluster LDAP timeoutMS Server",
	},
	"ldap_timeout": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Timeout in seconds for LDAP connections for this MS Server",
	},
	"ldap_auth_port": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "TCP port for LDAP connections for this",
	},
	"ldap_encryption": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Encryption for LDAP connections for this MS Server",
	},
	"managed": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Controls whether the Sites of this MS Server are to be synchronized by the assigned managing member or not",
	},
	"read_only": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enable/disable read-only synchronization of Sites for this Active Directory domain",
	},
	"last_sync_ts": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Timestamp of the last synchronization attempt",
	},
	"last_sync_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Status of the last synchronization attempt",
	},
	"last_sync_detail": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The detailed status of the last synchronization attempt.",
	},
	"supports_ipv6": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Flag indicating if the server supports IPv6",
	},
}

func ExpandMsserverAdSites(ctx context.Context, o types.Object, diags *diag.Diagnostics) *microsoftserver.MsserverAdSites {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MsserverAdSitesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MsserverAdSitesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *microsoftserver.MsserverAdSites {
	if m == nil {
		return nil
	}
	to := &microsoftserver.MsserverAdSites{
		UseDefaultIpSiteLink:       flex.ExpandBoolPointer(m.UseDefaultIpSiteLink),
		DefaultIpSiteLink:          flex.ExpandStringPointer(m.DefaultIpSiteLink),
		UseLogin:                   flex.ExpandBoolPointer(m.UseLogin),
		LoginName:                  flex.ExpandStringPointer(m.LoginName),
		LoginPassword:              flex.ExpandStringPointer(m.LoginPassword),
		UseSynchronizationMinDelay: flex.ExpandBoolPointer(m.UseSynchronizationMinDelay),
		SynchronizationMinDelay:    flex.ExpandInt64Pointer(m.SynchronizationMinDelay),
		UseLdapTimeout:             flex.ExpandBoolPointer(m.UseLdapTimeout),
		LdapTimeout:                flex.ExpandInt64Pointer(m.LdapTimeout),
		LdapAuthPort:               flex.ExpandInt64Pointer(m.LdapAuthPort),
		LdapEncryption:             flex.ExpandStringPointer(m.LdapEncryption),
		Managed:                    flex.ExpandBoolPointer(m.Managed),
		ReadOnly:                   flex.ExpandBoolPointer(m.ReadOnly),
	}
	return to
}

func FlattenMsserverAdSites(ctx context.Context, from *microsoftserver.MsserverAdSites, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MsserverAdSitesAttrTypes)
	}
	m := MsserverAdSitesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MsserverAdSitesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MsserverAdSitesModel) Flatten(ctx context.Context, from *microsoftserver.MsserverAdSites, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MsserverAdSitesModel{}
	}
	m.UseDefaultIpSiteLink = types.BoolPointerValue(from.UseDefaultIpSiteLink)
	m.DefaultIpSiteLink = flex.FlattenStringPointer(from.DefaultIpSiteLink)
	m.UseLogin = types.BoolPointerValue(from.UseLogin)
	m.LoginName = flex.FlattenStringPointer(from.LoginName)
	m.LoginPassword = flex.FlattenStringPointer(from.LoginPassword)
	m.UseSynchronizationMinDelay = types.BoolPointerValue(from.UseSynchronizationMinDelay)
	m.SynchronizationMinDelay = flex.FlattenInt64Pointer(from.SynchronizationMinDelay)
	m.UseLdapTimeout = types.BoolPointerValue(from.UseLdapTimeout)
	m.LdapTimeout = flex.FlattenInt64Pointer(from.LdapTimeout)
	m.LdapAuthPort = flex.FlattenInt64Pointer(from.LdapAuthPort)
	m.LdapEncryption = flex.FlattenStringPointer(from.LdapEncryption)
	m.Managed = types.BoolPointerValue(from.Managed)
	m.ReadOnly = types.BoolPointerValue(from.ReadOnly)
	m.LastSyncTs = flex.FlattenInt64Pointer(from.LastSyncTs)
	m.LastSyncStatus = flex.FlattenStringPointer(from.LastSyncStatus)
	m.LastSyncDetail = flex.FlattenStringPointer(from.LastSyncDetail)
	m.SupportsIpv6 = types.BoolPointerValue(from.SupportsIpv6)
}
