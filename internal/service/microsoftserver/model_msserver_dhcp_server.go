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

type MsserverDhcpServerModel struct {
	UseLogin                   types.Bool   `tfsdk:"use_login"`
	LoginName                  types.String `tfsdk:"login_name"`
	LoginPassword              types.String `tfsdk:"login_password"`
	Managed                    types.Bool   `tfsdk:"managed"`
	NextSyncControl            types.String `tfsdk:"next_sync_control"`
	Status                     types.String `tfsdk:"status"`
	StatusLastUpdated          types.Int64  `tfsdk:"status_last_updated"`
	UseEnableMonitoring        types.Bool   `tfsdk:"use_enable_monitoring"`
	EnableMonitoring           types.Bool   `tfsdk:"enable_monitoring"`
	UseEnableInvalidMac        types.Bool   `tfsdk:"use_enable_invalid_mac"`
	EnableInvalidMac           types.Bool   `tfsdk:"enable_invalid_mac"`
	SupportsFailover           types.Bool   `tfsdk:"supports_failover"`
	UseSynchronizationMinDelay types.Bool   `tfsdk:"use_synchronization_min_delay"`
	SynchronizationMinDelay    types.Int64  `tfsdk:"synchronization_min_delay"`
}

var MsserverDhcpServerAttrTypes = map[string]attr.Type{
	"use_login":                     types.BoolType,
	"login_name":                    types.StringType,
	"login_password":                types.StringType,
	"managed":                       types.BoolType,
	"next_sync_control":             types.StringType,
	"status":                        types.StringType,
	"status_last_updated":           types.Int64Type,
	"use_enable_monitoring":         types.BoolType,
	"enable_monitoring":             types.BoolType,
	"use_enable_invalid_mac":        types.BoolType,
	"enable_invalid_mac":            types.BoolType,
	"supports_failover":             types.BoolType,
	"use_synchronization_min_delay": types.BoolType,
	"synchronization_min_delay":     types.Int64Type,
}

var MsserverDhcpServerResourceSchemaAttributes = map[string]schema.Attribute{
	"use_login": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Flag to override login name and password from the MS Server",
	},
	"login_name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Microsoft Server login name",
	},
	"login_password": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Microsoft Server login password",
	},
	"managed": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "flag indicating if the DNS service is managed",
	},
	"next_sync_control": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Defines what control to apply on the DNS server",
	},
	"status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Status of the Microsoft DNS Service",
	},
	"status_last_updated": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Timestamp of the last update",
	},
	"use_enable_monitoring": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Override enable monitoring inherited from grid level",
	},
	"enable_monitoring": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Flag indicating if the DNS service is monitored and controlled",
	},
	"use_enable_invalid_mac": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Override setting for Enable Invalid Mac Address",
	},
	"enable_invalid_mac": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enable Invalid Mac Address",
	},
	"supports_failover": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Flag indicating if the DHCP supports Failover",
	},
	"use_synchronization_min_delay": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Flag to override synchronization interval from the MS Server",
	},
	"synchronization_min_delay": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Minimum number of minutes between two synchronizations",
	},
}

func ExpandMsserverDhcpServer(ctx context.Context, o types.Object, diags *diag.Diagnostics) *microsoftserver.MsserverDhcpServer {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MsserverDhcpServerModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MsserverDhcpServerModel) Expand(ctx context.Context, diags *diag.Diagnostics) *microsoftserver.MsserverDhcpServer {
	if m == nil {
		return nil
	}
	to := &microsoftserver.MsserverDhcpServer{
		UseLogin:                   flex.ExpandBoolPointer(m.UseLogin),
		LoginName:                  flex.ExpandStringPointer(m.LoginName),
		LoginPassword:              flex.ExpandStringPointer(m.LoginPassword),
		Managed:                    flex.ExpandBoolPointer(m.Managed),
		NextSyncControl:            flex.ExpandStringPointer(m.NextSyncControl),
		UseEnableMonitoring:        flex.ExpandBoolPointer(m.UseEnableMonitoring),
		EnableMonitoring:           flex.ExpandBoolPointer(m.EnableMonitoring),
		UseEnableInvalidMac:        flex.ExpandBoolPointer(m.UseEnableInvalidMac),
		EnableInvalidMac:           flex.ExpandBoolPointer(m.EnableInvalidMac),
		UseSynchronizationMinDelay: flex.ExpandBoolPointer(m.UseSynchronizationMinDelay),
		SynchronizationMinDelay:    flex.ExpandInt64Pointer(m.SynchronizationMinDelay),
	}
	return to
}

func FlattenMsserverDhcpServer(ctx context.Context, from *microsoftserver.MsserverDhcpServer, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MsserverDhcpServerAttrTypes)
	}
	m := MsserverDhcpServerModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MsserverDhcpServerAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MsserverDhcpServerModel) Flatten(ctx context.Context, from *microsoftserver.MsserverDhcpServer, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MsserverDhcpServerModel{}
	}
	m.UseLogin = types.BoolPointerValue(from.UseLogin)
	m.LoginName = flex.FlattenStringPointer(from.LoginName)
	m.LoginPassword = flex.FlattenStringPointer(from.LoginPassword)
	m.Managed = types.BoolPointerValue(from.Managed)
	m.NextSyncControl = flex.FlattenStringPointer(from.NextSyncControl)
	m.Status = flex.FlattenStringPointer(from.Status)
	m.StatusLastUpdated = flex.FlattenInt64Pointer(from.StatusLastUpdated)
	m.UseEnableMonitoring = types.BoolPointerValue(from.UseEnableMonitoring)
	m.EnableMonitoring = types.BoolPointerValue(from.EnableMonitoring)
	m.UseEnableInvalidMac = types.BoolPointerValue(from.UseEnableInvalidMac)
	m.EnableInvalidMac = types.BoolPointerValue(from.EnableInvalidMac)
	m.SupportsFailover = types.BoolPointerValue(from.SupportsFailover)
	m.UseSynchronizationMinDelay = types.BoolPointerValue(from.UseSynchronizationMinDelay)
	m.SynchronizationMinDelay = flex.FlattenInt64Pointer(from.SynchronizationMinDelay)
}
