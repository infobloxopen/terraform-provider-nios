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

type LdapAuthServiceServersModel struct {
	Address            types.String `tfsdk:"address"`
	AuthenticationType types.String `tfsdk:"authentication_type"`
	BaseDn             types.String `tfsdk:"base_dn"`
	BindPassword       types.String `tfsdk:"bind_password"`
	BindUserDn         types.String `tfsdk:"bind_user_dn"`
	Comment            types.String `tfsdk:"comment"`
	Disable            types.Bool   `tfsdk:"disable"`
	Encryption         types.String `tfsdk:"encryption"`
	Port               types.Int64  `tfsdk:"port"`
	UseMgmtPort        types.Bool   `tfsdk:"use_mgmt_port"`
	Version            types.String `tfsdk:"version"`
}

var LdapAuthServiceServersAttrTypes = map[string]attr.Type{
	"address":             types.StringType,
	"authentication_type": types.StringType,
	"base_dn":             types.StringType,
	"bind_password":       types.StringType,
	"bind_user_dn":        types.StringType,
	"comment":             types.StringType,
	"disable":             types.BoolType,
	"encryption":          types.StringType,
	"port":                types.Int64Type,
	"use_mgmt_port":       types.BoolType,
	"version":             types.StringType,
}

var LdapAuthServiceServersResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The IP address or FQDN of the LDAP server.",
	},
	"authentication_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The authentication type for the LDAP server.",
	},
	"base_dn": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The base DN for the LDAP server.",
	},
	"bind_password": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The user password for authentication.",
	},
	"bind_user_dn": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The user DN for authentication.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The LDAP descriptive comment.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the LDAP server is disabled.",
	},
	"encryption": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The LDAP server encryption type.",
	},
	"port": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The LDAP server port.",
	},
	"use_mgmt_port": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the connection via the MGMT interface is allowed.",
	},
	"version": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The LDAP server version.",
	},
}

func ExpandLdapAuthServiceServers(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.LdapAuthServiceServers {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m LdapAuthServiceServersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *LdapAuthServiceServersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.LdapAuthServiceServers {
	if m == nil {
		return nil
	}
	to := &security.LdapAuthServiceServers{
		Address:            flex.ExpandStringPointer(m.Address),
		AuthenticationType: flex.ExpandStringPointer(m.AuthenticationType),
		BaseDn:             flex.ExpandStringPointer(m.BaseDn),
		BindPassword:       flex.ExpandStringPointer(m.BindPassword),
		BindUserDn:         flex.ExpandStringPointer(m.BindUserDn),
		Comment:            flex.ExpandStringPointer(m.Comment),
		Disable:            flex.ExpandBoolPointer(m.Disable),
		Encryption:         flex.ExpandStringPointer(m.Encryption),
		Port:               flex.ExpandInt64Pointer(m.Port),
		UseMgmtPort:        flex.ExpandBoolPointer(m.UseMgmtPort),
		Version:            flex.ExpandStringPointer(m.Version),
	}
	return to
}

func FlattenLdapAuthServiceServers(ctx context.Context, from *security.LdapAuthServiceServers, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(LdapAuthServiceServersAttrTypes)
	}
	m := LdapAuthServiceServersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, LdapAuthServiceServersAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *LdapAuthServiceServersModel) Flatten(ctx context.Context, from *security.LdapAuthServiceServers, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = LdapAuthServiceServersModel{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.AuthenticationType = flex.FlattenStringPointer(from.AuthenticationType)
	m.BaseDn = flex.FlattenStringPointer(from.BaseDn)
	m.BindPassword = flex.FlattenStringPointer(from.BindPassword)
	m.BindUserDn = flex.FlattenStringPointer(from.BindUserDn)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.Encryption = flex.FlattenStringPointer(from.Encryption)
	m.Port = flex.FlattenInt64Pointer(from.Port)
	m.UseMgmtPort = types.BoolPointerValue(from.UseMgmtPort)
	m.Version = flex.FlattenStringPointer(from.Version)
}
