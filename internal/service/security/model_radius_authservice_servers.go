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

type RadiusAuthserviceServersModel struct {
	AcctPort      types.Int64  `tfsdk:"acct_port"`
	AuthPort      types.Int64  `tfsdk:"auth_port"`
	AuthType      types.String `tfsdk:"auth_type"`
	Comment       types.String `tfsdk:"comment"`
	Disable       types.Bool   `tfsdk:"disable"`
	Address       types.String `tfsdk:"address"`
	SharedSecret  types.String `tfsdk:"shared_secret"`
	UseAccounting types.Bool   `tfsdk:"use_accounting"`
	UseMgmtPort   types.Bool   `tfsdk:"use_mgmt_port"`
}

var RadiusAuthserviceServersAttrTypes = map[string]attr.Type{
	"acct_port":      types.Int64Type,
	"auth_port":      types.Int64Type,
	"auth_type":      types.StringType,
	"comment":        types.StringType,
	"disable":        types.BoolType,
	"address":        types.StringType,
	"shared_secret":  types.StringType,
	"use_accounting": types.BoolType,
	"use_mgmt_port":  types.BoolType,
}

var RadiusAuthserviceServersResourceSchemaAttributes = map[string]schema.Attribute{
	"acct_port": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The accounting port.",
	},
	"auth_port": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The authorization port.",
	},
	"auth_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The authentication protocol.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The RADIUS descriptive comment.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the RADIUS server is disabled.",
	},
	"address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The FQDN or the IP address of the RADIUS server that is used for authentication.",
	},
	"shared_secret": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The shared secret that the NIOS appliance and the RADIUS server use to encrypt and decrypt their messages.",
	},
	"use_accounting": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether RADIUS accounting is enabled.",
	},
	"use_mgmt_port": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether connection via the management interface is allowed.",
	},
}

func ExpandRadiusAuthserviceServers(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.RadiusAuthserviceServers {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RadiusAuthserviceServersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *RadiusAuthserviceServersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.RadiusAuthserviceServers {
	if m == nil {
		return nil
	}
	to := &security.RadiusAuthserviceServers{
		AcctPort:      flex.ExpandInt64Pointer(m.AcctPort),
		AuthPort:      flex.ExpandInt64Pointer(m.AuthPort),
		AuthType:      flex.ExpandStringPointer(m.AuthType),
		Comment:       flex.ExpandStringPointer(m.Comment),
		Disable:       flex.ExpandBoolPointer(m.Disable),
		Address:       flex.ExpandStringPointer(m.Address),
		SharedSecret:  flex.ExpandStringPointer(m.SharedSecret),
		UseAccounting: flex.ExpandBoolPointer(m.UseAccounting),
		UseMgmtPort:   flex.ExpandBoolPointer(m.UseMgmtPort),
	}
	return to
}

func FlattenRadiusAuthserviceServers(ctx context.Context, from *security.RadiusAuthserviceServers, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RadiusAuthserviceServersAttrTypes)
	}
	m := RadiusAuthserviceServersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RadiusAuthserviceServersAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RadiusAuthserviceServersModel) Flatten(ctx context.Context, from *security.RadiusAuthserviceServers, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RadiusAuthserviceServersModel{}
	}
	m.AcctPort = flex.FlattenInt64Pointer(from.AcctPort)
	m.AuthPort = flex.FlattenInt64Pointer(from.AuthPort)
	m.AuthType = flex.FlattenStringPointer(from.AuthType)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.Address = flex.FlattenStringPointer(from.Address)
	m.SharedSecret = flex.FlattenStringPointer(from.SharedSecret)
	m.UseAccounting = types.BoolPointerValue(from.UseAccounting)
	m.UseMgmtPort = types.BoolPointerValue(from.UseMgmtPort)
}
