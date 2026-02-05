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

type TacacsplusAuthserviceServersModel struct {
	Address       types.String `tfsdk:"address"`
	Port          types.Int64  `tfsdk:"port"`
	SharedSecret  types.String `tfsdk:"shared_secret"`
	AuthType      types.String `tfsdk:"auth_type"`
	Comment       types.String `tfsdk:"comment"`
	Disable       types.Bool   `tfsdk:"disable"`
	UseMgmtPort   types.Bool   `tfsdk:"use_mgmt_port"`
	UseAccounting types.Bool   `tfsdk:"use_accounting"`
}

var TacacsplusAuthserviceServersAttrTypes = map[string]attr.Type{
	"address":        types.StringType,
	"port":           types.Int64Type,
	"shared_secret":  types.StringType,
	"auth_type":      types.StringType,
	"comment":        types.StringType,
	"disable":        types.BoolType,
	"use_mgmt_port":  types.BoolType,
	"use_accounting": types.BoolType,
}

var TacacsplusAuthserviceServersResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The valid IP address or FQDN of the TACACS+ server.",
	},
	"port": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The TACACS+ server port.",
	},
	"shared_secret": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The secret key with which to connect to the TACACS+ server.",
	},
	"auth_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The authentication protocol.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The TACACS+ descriptive comment.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the TACACS+ server is disabled.",
	},
	"use_mgmt_port": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the TACACS+ server is connected via the management interface.",
	},
	"use_accounting": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the TACACS+ accounting server is used.",
	},
}

func ExpandTacacsplusAuthserviceServers(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.TacacsplusAuthserviceServers {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m TacacsplusAuthserviceServersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *TacacsplusAuthserviceServersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.TacacsplusAuthserviceServers {
	if m == nil {
		return nil
	}
	to := &security.TacacsplusAuthserviceServers{
		Address:       flex.ExpandStringPointer(m.Address),
		Port:          flex.ExpandInt64Pointer(m.Port),
		SharedSecret:  flex.ExpandStringPointer(m.SharedSecret),
		AuthType:      flex.ExpandStringPointer(m.AuthType),
		Comment:       flex.ExpandStringPointer(m.Comment),
		Disable:       flex.ExpandBoolPointer(m.Disable),
		UseMgmtPort:   flex.ExpandBoolPointer(m.UseMgmtPort),
		UseAccounting: flex.ExpandBoolPointer(m.UseAccounting),
	}
	return to
}

func FlattenTacacsplusAuthserviceServers(ctx context.Context, from *security.TacacsplusAuthserviceServers, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(TacacsplusAuthserviceServersAttrTypes)
	}
	m := TacacsplusAuthserviceServersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, TacacsplusAuthserviceServersAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *TacacsplusAuthserviceServersModel) Flatten(ctx context.Context, from *security.TacacsplusAuthserviceServers, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = TacacsplusAuthserviceServersModel{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Port = flex.FlattenInt64Pointer(from.Port)
	m.SharedSecret = flex.FlattenStringPointer(from.SharedSecret)
	m.AuthType = flex.FlattenStringPointer(from.AuthType)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.UseMgmtPort = types.BoolPointerValue(from.UseMgmtPort)
	m.UseAccounting = types.BoolPointerValue(from.UseAccounting)
}
