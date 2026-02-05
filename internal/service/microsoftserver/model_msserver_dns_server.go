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

type MsserverDnsServerModel struct {
	UseLogin                   types.Bool   `tfsdk:"use_login"`
	LoginName                  types.String `tfsdk:"login_name"`
	LoginPassword              types.String `tfsdk:"login_password"`
	Managed                    types.Bool   `tfsdk:"managed"`
	NextSyncControl            types.String `tfsdk:"next_sync_control"`
	Status                     types.String `tfsdk:"status"`
	StatusDetail               types.String `tfsdk:"status_detail"`
	StatusLastUpdated          types.Int64  `tfsdk:"status_last_updated"`
	LastSyncTs                 types.Int64  `tfsdk:"last_sync_ts"`
	LastSyncStatus             types.String `tfsdk:"last_sync_status"`
	LastSyncDetail             types.String `tfsdk:"last_sync_detail"`
	Forwarders                 types.String `tfsdk:"forwarders"`
	SupportsIpv6               types.Bool   `tfsdk:"supports_ipv6"`
	SupportsIpv6Reverse        types.Bool   `tfsdk:"supports_ipv6_reverse"`
	SupportsRrDname            types.Bool   `tfsdk:"supports_rr_dname"`
	SupportsDnssec             types.Bool   `tfsdk:"supports_dnssec"`
	SupportsActiveDirectory    types.Bool   `tfsdk:"supports_active_directory"`
	Address                    types.String `tfsdk:"address"`
	SupportsRrNaptr            types.Bool   `tfsdk:"supports_rr_naptr"`
	UseEnableMonitoring        types.Bool   `tfsdk:"use_enable_monitoring"`
	EnableMonitoring           types.Bool   `tfsdk:"enable_monitoring"`
	UseSynchronizationMinDelay types.Bool   `tfsdk:"use_synchronization_min_delay"`
	SynchronizationMinDelay    types.Int64  `tfsdk:"synchronization_min_delay"`
	UseEnableDnsReportsSync    types.Bool   `tfsdk:"use_enable_dns_reports_sync"`
	EnableDnsReportsSync       types.Bool   `tfsdk:"enable_dns_reports_sync"`
}

var MsserverDnsServerAttrTypes = map[string]attr.Type{
	"use_login":                     types.BoolType,
	"login_name":                    types.StringType,
	"login_password":                types.StringType,
	"managed":                       types.BoolType,
	"next_sync_control":             types.StringType,
	"status":                        types.StringType,
	"status_detail":                 types.StringType,
	"status_last_updated":           types.Int64Type,
	"last_sync_ts":                  types.Int64Type,
	"last_sync_status":              types.StringType,
	"last_sync_detail":              types.StringType,
	"forwarders":                    types.StringType,
	"supports_ipv6":                 types.BoolType,
	"supports_ipv6_reverse":         types.BoolType,
	"supports_rr_dname":             types.BoolType,
	"supports_dnssec":               types.BoolType,
	"supports_active_directory":     types.BoolType,
	"address":                       types.StringType,
	"supports_rr_naptr":             types.BoolType,
	"use_enable_monitoring":         types.BoolType,
	"enable_monitoring":             types.BoolType,
	"use_synchronization_min_delay": types.BoolType,
	"synchronization_min_delay":     types.Int64Type,
	"use_enable_dns_reports_sync":   types.BoolType,
	"enable_dns_reports_sync":       types.BoolType,
}

var MsserverDnsServerResourceSchemaAttributes = map[string]schema.Attribute{
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
	"status_detail": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Detailed status of the DNS status",
	},
	"status_last_updated": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Timestamp of the last update",
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
		MarkdownDescription: "Detailled status of the last synchronization attempt",
	},
	"forwarders": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Ordered list of IP addresses to forward queries to",
	},
	"supports_ipv6": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Flag indicating if the server supports IPv6",
	},
	"supports_ipv6_reverse": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Flag indicating if the server supports reverse IPv6 zones",
	},
	"supports_rr_dname": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Flag indicating if the server supports DNAME records",
	},
	"supports_dnssec": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Flag indicating if the server supports",
	},
	"supports_active_directory": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Flag indicating if the server supports AD integrated zones",
	},
	"address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "MS Server ip address",
	},
	"supports_rr_naptr": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Flag indicating if the server supports NAPTR records",
	},
	"use_enable_monitoring": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Override enable monitoring inherited from grid level",
	},
	"enable_monitoring": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Flag indicating if the DNS service is monitored and controlled",
	},
	"use_synchronization_min_delay": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Flag to override synchronization interval from the MS Server",
	},
	"synchronization_min_delay": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Minimum number of minutes between two synchronizations",
	},
	"use_enable_dns_reports_sync": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Override enable reports data inherited from grid level",
	},
	"enable_dns_reports_sync": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enable or Disable MS DNS data for reports from this MS Server",
	},
}

func ExpandMsserverDnsServer(ctx context.Context, o types.Object, diags *diag.Diagnostics) *microsoftserver.MsserverDnsServer {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MsserverDnsServerModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MsserverDnsServerModel) Expand(ctx context.Context, diags *diag.Diagnostics) *microsoftserver.MsserverDnsServer {
	if m == nil {
		return nil
	}
	to := &microsoftserver.MsserverDnsServer{
		UseLogin:                   flex.ExpandBoolPointer(m.UseLogin),
		LoginName:                  flex.ExpandStringPointer(m.LoginName),
		LoginPassword:              flex.ExpandStringPointer(m.LoginPassword),
		Managed:                    flex.ExpandBoolPointer(m.Managed),
		NextSyncControl:            flex.ExpandStringPointer(m.NextSyncControl),
		UseEnableMonitoring:        flex.ExpandBoolPointer(m.UseEnableMonitoring),
		EnableMonitoring:           flex.ExpandBoolPointer(m.EnableMonitoring),
		UseSynchronizationMinDelay: flex.ExpandBoolPointer(m.UseSynchronizationMinDelay),
		SynchronizationMinDelay:    flex.ExpandInt64Pointer(m.SynchronizationMinDelay),
		UseEnableDnsReportsSync:    flex.ExpandBoolPointer(m.UseEnableDnsReportsSync),
		EnableDnsReportsSync:       flex.ExpandBoolPointer(m.EnableDnsReportsSync),
	}
	return to
}

func FlattenMsserverDnsServer(ctx context.Context, from *microsoftserver.MsserverDnsServer, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MsserverDnsServerAttrTypes)
	}
	m := MsserverDnsServerModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MsserverDnsServerAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MsserverDnsServerModel) Flatten(ctx context.Context, from *microsoftserver.MsserverDnsServer, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MsserverDnsServerModel{}
	}
	m.UseLogin = types.BoolPointerValue(from.UseLogin)
	m.LoginName = flex.FlattenStringPointer(from.LoginName)
	m.LoginPassword = flex.FlattenStringPointer(from.LoginPassword)
	m.Managed = types.BoolPointerValue(from.Managed)
	m.NextSyncControl = flex.FlattenStringPointer(from.NextSyncControl)
	m.Status = flex.FlattenStringPointer(from.Status)
	m.StatusDetail = flex.FlattenStringPointer(from.StatusDetail)
	m.StatusLastUpdated = flex.FlattenInt64Pointer(from.StatusLastUpdated)
	m.LastSyncTs = flex.FlattenInt64Pointer(from.LastSyncTs)
	m.LastSyncStatus = flex.FlattenStringPointer(from.LastSyncStatus)
	m.LastSyncDetail = flex.FlattenStringPointer(from.LastSyncDetail)
	m.Forwarders = flex.FlattenStringPointer(from.Forwarders)
	m.SupportsIpv6 = types.BoolPointerValue(from.SupportsIpv6)
	m.SupportsIpv6Reverse = types.BoolPointerValue(from.SupportsIpv6Reverse)
	m.SupportsRrDname = types.BoolPointerValue(from.SupportsRrDname)
	m.SupportsDnssec = types.BoolPointerValue(from.SupportsDnssec)
	m.SupportsActiveDirectory = types.BoolPointerValue(from.SupportsActiveDirectory)
	m.Address = flex.FlattenStringPointer(from.Address)
	m.SupportsRrNaptr = types.BoolPointerValue(from.SupportsRrNaptr)
	m.UseEnableMonitoring = types.BoolPointerValue(from.UseEnableMonitoring)
	m.EnableMonitoring = types.BoolPointerValue(from.EnableMonitoring)
	m.UseSynchronizationMinDelay = types.BoolPointerValue(from.UseSynchronizationMinDelay)
	m.SynchronizationMinDelay = flex.FlattenInt64Pointer(from.SynchronizationMinDelay)
	m.UseEnableDnsReportsSync = types.BoolPointerValue(from.UseEnableDnsReportsSync)
	m.EnableDnsReportsSync = types.BoolPointerValue(from.EnableDnsReportsSync)
}
