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

type MsserverDnsModel struct {
	Ref                        types.String `tfsdk:"ref"`
    Uuid        types.String `tfsdk:"uuid"`
	Address                    types.String `tfsdk:"address"`
	EnableDnsReportsSync       types.Bool   `tfsdk:"enable_dns_reports_sync"`
	LoginName                  types.String `tfsdk:"login_name"`
	LoginPassword              types.String `tfsdk:"login_password"`
	SynchronizationInterval    types.Int64  `tfsdk:"synchronization_interval"`
	UseEnableDnsReportsSync    types.Bool   `tfsdk:"use_enable_dns_reports_sync"`
	UseLogin                   types.Bool   `tfsdk:"use_login"`
	UseSynchronizationInterval types.Bool   `tfsdk:"use_synchronization_interval"`
}

var MsserverDnsAttrTypes = map[string]attr.Type{
	"ref":                          types.StringType,
    "uuid":        types.StringType,
	"address":                      types.StringType,
	"enable_dns_reports_sync":      types.BoolType,
	"login_name":                   types.StringType,
	"login_password":               types.StringType,
	"synchronization_interval":     types.Int64Type,
	"use_enable_dns_reports_sync":  types.BoolType,
	"use_login":                    types.BoolType,
	"use_synchronization_interval": types.BoolType,
}

var MsserverDnsResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
    "uuid": schema.StringAttribute{
        Computed:            true,
        MarkdownDescription: "The uuid to the object.",
    },
	"address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The address or FQDN of the DNS Microsoft Server.",
	},
	"enable_dns_reports_sync": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if synchronization of DNS reporting data from the Microsoft server is enabled or not.",
	},
	"login_name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The login name of the DNS Microsoft Server.",
	},
	"login_password": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The login password of the DNS Microsoft Server.",
	},
	"synchronization_interval": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The minimum number of minutes between two synchronizations.",
	},
	"use_enable_dns_reports_sync": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: enable_dns_reports_sync",
	},
	"use_login": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: login_name , login_password",
	},
	"use_synchronization_interval": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: synchronization_interval",
	},
}

func ExpandMsserverDns(ctx context.Context, o types.Object, diags *diag.Diagnostics) *microsoftserver.MsserverDns {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MsserverDnsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MsserverDnsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *microsoftserver.MsserverDns {
	if m == nil {
		return nil
	}
	to := &microsoftserver.MsserverDns{
		Ref:                        flex.ExpandStringPointer(m.Ref),
		EnableDnsReportsSync:       flex.ExpandBoolPointer(m.EnableDnsReportsSync),
		LoginName:                  flex.ExpandStringPointer(m.LoginName),
		LoginPassword:              flex.ExpandStringPointer(m.LoginPassword),
		SynchronizationInterval:    flex.ExpandInt64Pointer(m.SynchronizationInterval),
		UseEnableDnsReportsSync:    flex.ExpandBoolPointer(m.UseEnableDnsReportsSync),
		UseLogin:                   flex.ExpandBoolPointer(m.UseLogin),
		UseSynchronizationInterval: flex.ExpandBoolPointer(m.UseSynchronizationInterval),
	}
	return to
}

func FlattenMsserverDns(ctx context.Context, from *microsoftserver.MsserverDns, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MsserverDnsAttrTypes)
	}
	m := MsserverDnsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MsserverDnsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MsserverDnsModel) Flatten(ctx context.Context, from *microsoftserver.MsserverDns, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MsserverDnsModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
    m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Address = flex.FlattenStringPointer(from.Address)
	m.EnableDnsReportsSync = types.BoolPointerValue(from.EnableDnsReportsSync)
	m.LoginName = flex.FlattenStringPointer(from.LoginName)
	m.LoginPassword = flex.FlattenStringPointer(from.LoginPassword)
	m.SynchronizationInterval = flex.FlattenInt64Pointer(from.SynchronizationInterval)
	m.UseEnableDnsReportsSync = types.BoolPointerValue(from.UseEnableDnsReportsSync)
	m.UseLogin = types.BoolPointerValue(from.UseLogin)
	m.UseSynchronizationInterval = types.BoolPointerValue(from.UseSynchronizationInterval)
}
