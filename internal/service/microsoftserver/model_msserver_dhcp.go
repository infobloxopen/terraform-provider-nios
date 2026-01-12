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

type MsserverDhcpModel struct {
	Ref                        types.String `tfsdk:"ref"`
	Address                    types.String `tfsdk:"address"`
	Comment                    types.String `tfsdk:"comment"`
	DhcpUtilization            types.Int64  `tfsdk:"dhcp_utilization"`
	DhcpUtilizationStatus      types.String `tfsdk:"dhcp_utilization_status"`
	DynamicHosts               types.Int64  `tfsdk:"dynamic_hosts"`
	LastSyncTs                 types.Int64  `tfsdk:"last_sync_ts"`
	LoginName                  types.String `tfsdk:"login_name"`
	LoginPassword              types.String `tfsdk:"login_password"`
	NetworkView                types.String `tfsdk:"network_view"`
	NextSyncControl            types.String `tfsdk:"next_sync_control"`
	ReadOnly                   types.Bool   `tfsdk:"read_only"`
	ServerName                 types.String `tfsdk:"server_name"`
	StaticHosts                types.Int64  `tfsdk:"static_hosts"`
	Status                     types.String `tfsdk:"status"`
	StatusDetail               types.String `tfsdk:"status_detail"`
	StatusLastUpdated          types.Int64  `tfsdk:"status_last_updated"`
	SupportsFailover           types.Bool   `tfsdk:"supports_failover"`
	SynchronizationInterval    types.Int64  `tfsdk:"synchronization_interval"`
	TotalHosts                 types.Int64  `tfsdk:"total_hosts"`
	UseLogin                   types.Bool   `tfsdk:"use_login"`
	UseSynchronizationInterval types.Bool   `tfsdk:"use_synchronization_interval"`
}

var MsserverDhcpAttrTypes = map[string]attr.Type{
	"ref":                          types.StringType,
	"address":                      types.StringType,
	"comment":                      types.StringType,
	"dhcp_utilization":             types.Int64Type,
	"dhcp_utilization_status":      types.StringType,
	"dynamic_hosts":                types.Int64Type,
	"last_sync_ts":                 types.Int64Type,
	"login_name":                   types.StringType,
	"login_password":               types.StringType,
	"network_view":                 types.StringType,
	"next_sync_control":            types.StringType,
	"read_only":                    types.BoolType,
	"server_name":                  types.StringType,
	"static_hosts":                 types.Int64Type,
	"status":                       types.StringType,
	"status_detail":                types.StringType,
	"status_last_updated":          types.Int64Type,
	"supports_failover":            types.BoolType,
	"synchronization_interval":     types.Int64Type,
	"total_hosts":                  types.Int64Type,
	"use_login":                    types.BoolType,
	"use_synchronization_interval": types.BoolType,
}

var MsserverDhcpResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The address or FQDN of the DHCP Microsoft Server.",
	},
	"comment": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Comment from Microsoft Server",
	},
	"dhcp_utilization": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The percentage of the total DHCP utilization of DHCP objects belonging to the DHCP Microsoft Server multiplied by 1000. This is the percentage of the total number of available IP addresses from all the DHCP objects belonging to the DHCP Microsoft Server versus the total number of all IP addresses in all of the DHCP objects on the DHCP Microsoft Server.",
	},
	"dhcp_utilization_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "A string describing the utilization level of DHCP objects that belong to the DHCP Microsoft Server.",
	},
	"dynamic_hosts": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The total number of DHCP leases issued for the DHCP objects on the DHCP Microsoft Server.",
	},
	"last_sync_ts": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Timestamp of the last synchronization attempt",
	},
	"login_name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The login name of the DHCP Microsoft Server.",
	},
	"login_password": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The login password of the DHCP Microsoft Server.",
	},
	"network_view": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Network view to update",
	},
	"next_sync_control": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Defines what control to apply on the DHCP server",
	},
	"read_only": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Whether Microsoft server is read only",
	},
	"server_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Microsoft server address",
	},
	"static_hosts": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of static DHCP addresses configured in DHCP objects that belong to the DHCP Microsoft Server.",
	},
	"status": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Status of the Microsoft DHCP Service",
	},
	"status_detail": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Detailed status of the DHCP status",
	},
	"status_last_updated": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Timestamp of the last update",
	},
	"supports_failover": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Flag indicating if the DHCP supports Failover",
	},
	"synchronization_interval": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The minimum number of minutes between two synchronizations.",
	},
	"total_hosts": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The total number of DHCP addresses configured in DHCP objects that belong to the DHCP Microsoft Server.",
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

func ExpandMsserverDhcp(ctx context.Context, o types.Object, diags *diag.Diagnostics) *microsoftserver.MsserverDhcp {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MsserverDhcpModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MsserverDhcpModel) Expand(ctx context.Context, diags *diag.Diagnostics) *microsoftserver.MsserverDhcp {
	if m == nil {
		return nil
	}
	to := &microsoftserver.MsserverDhcp{
		Ref:                        flex.ExpandStringPointer(m.Ref),
		LoginName:                  flex.ExpandStringPointer(m.LoginName),
		LoginPassword:              flex.ExpandStringPointer(m.LoginPassword),
		NextSyncControl:            flex.ExpandStringPointer(m.NextSyncControl),
		Status:                     flex.ExpandStringPointer(m.Status),
		SynchronizationInterval:    flex.ExpandInt64Pointer(m.SynchronizationInterval),
		UseLogin:                   flex.ExpandBoolPointer(m.UseLogin),
		UseSynchronizationInterval: flex.ExpandBoolPointer(m.UseSynchronizationInterval),
	}
	return to
}

func FlattenMsserverDhcp(ctx context.Context, from *microsoftserver.MsserverDhcp, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MsserverDhcpAttrTypes)
	}
	m := MsserverDhcpModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MsserverDhcpAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MsserverDhcpModel) Flatten(ctx context.Context, from *microsoftserver.MsserverDhcp, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MsserverDhcpModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.DhcpUtilization = flex.FlattenInt64Pointer(from.DhcpUtilization)
	m.DhcpUtilizationStatus = flex.FlattenStringPointer(from.DhcpUtilizationStatus)
	m.DynamicHosts = flex.FlattenInt64Pointer(from.DynamicHosts)
	m.LastSyncTs = flex.FlattenInt64Pointer(from.LastSyncTs)
	m.LoginName = flex.FlattenStringPointer(from.LoginName)
	m.LoginPassword = flex.FlattenStringPointer(from.LoginPassword)
	m.NetworkView = flex.FlattenStringPointer(from.NetworkView)
	m.NextSyncControl = flex.FlattenStringPointer(from.NextSyncControl)
	m.ReadOnly = types.BoolPointerValue(from.ReadOnly)
	m.ServerName = flex.FlattenStringPointer(from.ServerName)
	m.StaticHosts = flex.FlattenInt64Pointer(from.StaticHosts)
	m.Status = flex.FlattenStringPointer(from.Status)
	m.StatusDetail = flex.FlattenStringPointer(from.StatusDetail)
	m.StatusLastUpdated = flex.FlattenInt64Pointer(from.StatusLastUpdated)
	m.SupportsFailover = types.BoolPointerValue(from.SupportsFailover)
	m.SynchronizationInterval = flex.FlattenInt64Pointer(from.SynchronizationInterval)
	m.TotalHosts = flex.FlattenInt64Pointer(from.TotalHosts)
	m.UseLogin = types.BoolPointerValue(from.UseLogin)
	m.UseSynchronizationInterval = types.BoolPointerValue(from.UseSynchronizationInterval)
}
