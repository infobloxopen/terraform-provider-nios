package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type GridHttpProxyServerSettingModel struct {
	Address                   types.String `tfsdk:"address"`
	Port                      types.Int64  `tfsdk:"port"`
	EnableProxy               types.Bool   `tfsdk:"enable_proxy"`
	EnableContentInspection   types.Bool   `tfsdk:"enable_content_inspection"`
	VerifyCname               types.Bool   `tfsdk:"verify_cname"`
	Comment                   types.String `tfsdk:"comment"`
	Username                  types.String `tfsdk:"username"`
	Password                  types.String `tfsdk:"password"`
	Certificate               types.String `tfsdk:"certificate"`
	EnableUsernameAndPassword types.Bool   `tfsdk:"enable_username_and_password"`
}

var GridHttpProxyServerSettingAttrTypes = map[string]attr.Type{
	"address":                      types.StringType,
	"port":                         types.Int64Type,
	"enable_proxy":                 types.BoolType,
	"enable_content_inspection":    types.BoolType,
	"verify_cname":                 types.BoolType,
	"comment":                      types.StringType,
	"username":                     types.StringType,
	"password":                     types.StringType,
	"certificate":                  types.StringType,
	"enable_username_and_password": types.BoolType,
}

var GridHttpProxyServerSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The address of the HTTP proxy server.",
	},
	"port": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The port on which the HTTP proxy server listens.",
	},
	"enable_proxy": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the HTTP proxy server is enabled or not.",
	},
	"enable_content_inspection": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if HTTPS content inspection by the HTTP proxy server is enabled or not.",
	},
	"verify_cname": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the CNAME record query verification is enabled or not.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The descriptive comment for the HTTP proxy server configuration.",
	},
	"username": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The user name for the HTTP proxy server.",
	},
	"password": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The password for the HTTP proxy server.",
	},
	"certificate": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The token returned by the uploadinit function call in object fileop for the CA certificate file used in the content inspection by an HTTP proxy server.",
	},
	"enable_username_and_password": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if username and password for HTTP Proxy Server connectivity is used or not.",
	},
}

func ExpandGridHttpProxyServerSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridHttpProxyServerSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridHttpProxyServerSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridHttpProxyServerSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridHttpProxyServerSetting {
	if m == nil {
		return nil
	}
	to := &grid.GridHttpProxyServerSetting{
		Address:                   flex.ExpandStringPointer(m.Address),
		Port:                      flex.ExpandInt64Pointer(m.Port),
		EnableProxy:               flex.ExpandBoolPointer(m.EnableProxy),
		EnableContentInspection:   flex.ExpandBoolPointer(m.EnableContentInspection),
		VerifyCname:               flex.ExpandBoolPointer(m.VerifyCname),
		Comment:                   flex.ExpandStringPointer(m.Comment),
		Username:                  flex.ExpandStringPointer(m.Username),
		Password:                  flex.ExpandStringPointer(m.Password),
		Certificate:               flex.ExpandStringPointer(m.Certificate),
		EnableUsernameAndPassword: flex.ExpandBoolPointer(m.EnableUsernameAndPassword),
	}
	return to
}

func FlattenGridHttpProxyServerSetting(ctx context.Context, from *grid.GridHttpProxyServerSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridHttpProxyServerSettingAttrTypes)
	}
	m := GridHttpProxyServerSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridHttpProxyServerSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridHttpProxyServerSettingModel) Flatten(ctx context.Context, from *grid.GridHttpProxyServerSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridHttpProxyServerSettingModel{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Port = flex.FlattenInt64Pointer(from.Port)
	m.EnableProxy = types.BoolPointerValue(from.EnableProxy)
	m.EnableContentInspection = types.BoolPointerValue(from.EnableContentInspection)
	m.VerifyCname = types.BoolPointerValue(from.VerifyCname)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Username = flex.FlattenStringPointer(from.Username)
	m.Password = flex.FlattenStringPointer(from.Password)
	m.Certificate = flex.FlattenStringPointer(from.Certificate)
	m.EnableUsernameAndPassword = types.BoolPointerValue(from.EnableUsernameAndPassword)
}
