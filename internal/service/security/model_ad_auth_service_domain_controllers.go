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

type AdAuthServiceDomainControllersModel struct {
	FqdnOrIp    types.String `tfsdk:"fqdn_or_ip"`
	AuthPort    types.Int64  `tfsdk:"auth_port"`
	Comment     types.String `tfsdk:"comment"`
	Disabled    types.Bool   `tfsdk:"disabled"`
	Encryption  types.String `tfsdk:"encryption"`
	MgmtPort    types.Bool   `tfsdk:"mgmt_port"`
	UseMgmtPort types.Bool   `tfsdk:"use_mgmt_port"`
}

var AdAuthServiceDomainControllersAttrTypes = map[string]attr.Type{
	"fqdn_or_ip":    types.StringType,
	"auth_port":     types.Int64Type,
	"comment":       types.StringType,
	"disabled":      types.BoolType,
	"encryption":    types.StringType,
	"mgmt_port":     types.BoolType,
	"use_mgmt_port": types.BoolType,
}

var AdAuthServiceDomainControllersResourceSchemaAttributes = map[string]schema.Attribute{
	"fqdn_or_ip": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The FQDN (Fully Qualified Domain Name) or IP address of the server.",
	},
	"auth_port": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The authentication port.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The descriptive comment for the AD authentication server.",
	},
	"disabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the AD authorization server is disabled.",
	},
	"encryption": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The type of encryption to use.",
	},
	"mgmt_port": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determine if the MGMT port is enabled for the AD authentication server.",
	},
	"use_mgmt_port": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: mgmt_port",
	},
}

func ExpandAdAuthServiceDomainControllers(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.AdAuthServiceDomainControllers {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AdAuthServiceDomainControllersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AdAuthServiceDomainControllersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.AdAuthServiceDomainControllers {
	if m == nil {
		return nil
	}
	to := &security.AdAuthServiceDomainControllers{
		FqdnOrIp:    flex.ExpandStringPointer(m.FqdnOrIp),
		AuthPort:    flex.ExpandInt64Pointer(m.AuthPort),
		Comment:     flex.ExpandStringPointer(m.Comment),
		Disabled:    flex.ExpandBoolPointer(m.Disabled),
		Encryption:  flex.ExpandStringPointer(m.Encryption),
		MgmtPort:    flex.ExpandBoolPointer(m.MgmtPort),
		UseMgmtPort: flex.ExpandBoolPointer(m.UseMgmtPort),
	}
	return to
}

func FlattenAdAuthServiceDomainControllers(ctx context.Context, from *security.AdAuthServiceDomainControllers, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AdAuthServiceDomainControllersAttrTypes)
	}
	m := AdAuthServiceDomainControllersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AdAuthServiceDomainControllersAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AdAuthServiceDomainControllersModel) Flatten(ctx context.Context, from *security.AdAuthServiceDomainControllers, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AdAuthServiceDomainControllersModel{}
	}
	m.FqdnOrIp = flex.FlattenStringPointer(from.FqdnOrIp)
	m.AuthPort = flex.FlattenInt64Pointer(from.AuthPort)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disabled = types.BoolPointerValue(from.Disabled)
	m.Encryption = flex.FlattenStringPointer(from.Encryption)
	m.MgmtPort = types.BoolPointerValue(from.MgmtPort)
	m.UseMgmtPort = types.BoolPointerValue(from.UseMgmtPort)
}
