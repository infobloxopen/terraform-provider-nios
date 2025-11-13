package microsoftserver

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/microsoftserver"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/import"
)

type MsserverModel struct {
	Ref                         types.String `tfsdk:"ref"`
	AdDomain                    types.String `tfsdk:"ad_domain"`
	AdSites                     types.Object `tfsdk:"ad_sites"`
	AdUser                      types.Object `tfsdk:"ad_user"`
	Address                     types.String `tfsdk:"address"`
	Comment                     types.String `tfsdk:"comment"`
	ConnectionStatus            types.String `tfsdk:"connection_status"`
	ConnectionStatusDetail      types.String `tfsdk:"connection_status_detail"`
	DhcpServer                  types.Object `tfsdk:"dhcp_server"`
	Disabled                    types.Bool   `tfsdk:"disabled"`
	DnsServer                   types.Object `tfsdk:"dns_server"`
	DnsView                     types.String `tfsdk:"dns_view"`
	ExtAttrs                    types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll                 types.Map    `tfsdk:"extattrs_all"`
	GridMember                  types.String `tfsdk:"grid_member"`
	LastSeen                    types.Int64  `tfsdk:"last_seen"`
	LogDestination              types.String `tfsdk:"log_destination"`
	LogLevel                    types.String `tfsdk:"log_level"`
	LoginName                   types.String `tfsdk:"login_name"`
	LoginPassword               types.String `tfsdk:"login_password"`
	ManagingMember              types.String `tfsdk:"managing_member"`
	MsMaxConnection             types.Int64  `tfsdk:"ms_max_connection"`
	MsRpcTimeoutInSeconds       types.Int64  `tfsdk:"ms_rpc_timeout_in_seconds"`
	NetworkView                 types.String `tfsdk:"network_view"`
	ReadOnly                    types.Bool   `tfsdk:"read_only"`
	RootAdDomain                types.String `tfsdk:"root_ad_domain"`
	ServerName                  types.String `tfsdk:"server_name"`
	SynchronizationMinDelay     types.Int64  `tfsdk:"synchronization_min_delay"`
	SynchronizationStatus       types.String `tfsdk:"synchronization_status"`
	SynchronizationStatusDetail types.String `tfsdk:"synchronization_status_detail"`
	UseLogDestination           types.Bool   `tfsdk:"use_log_destination"`
	UseMsMaxConnection          types.Bool   `tfsdk:"use_ms_max_connection"`
	UseMsRpcTimeoutInSeconds    types.Bool   `tfsdk:"use_ms_rpc_timeout_in_seconds"`
	Version                     types.String `tfsdk:"version"`
}

var MsserverAttrTypes = map[string]attr.Type{
	"ref":                           types.StringType,
	"ad_domain":                     types.StringType,
	"ad_sites":                      types.ObjectType{AttrTypes: MsserverAdSitesAttrTypes},
	"ad_user":                       types.ObjectType{AttrTypes: MsserverAdUserAttrTypes},
	"address":                       types.StringType,
	"comment":                       types.StringType,
	"connection_status":             types.StringType,
	"connection_status_detail":      types.StringType,
	"dhcp_server":                   types.ObjectType{AttrTypes: MsserverDhcpServerAttrTypes},
	"disabled":                      types.BoolType,
	"dns_server":                    types.ObjectType{AttrTypes: MsserverDnsServerAttrTypes},
	"dns_view":                      types.StringType,
	"extattrs":                      types.MapType{ElemType: types.StringType},
	"extattrs_all":                  types.MapType{ElemType: types.StringType},
	"grid_member":                   types.StringType,
	"last_seen":                     types.Int64Type,
	"log_destination":               types.StringType,
	"log_level":                     types.StringType,
	"login_name":                    types.StringType,
	"login_password":                types.StringType,
	"managing_member":               types.StringType,
	"ms_max_connection":             types.Int64Type,
	"ms_rpc_timeout_in_seconds":     types.Int64Type,
	"network_view":                  types.StringType,
	"read_only":                     types.BoolType,
	"root_ad_domain":                types.StringType,
	"server_name":                   types.StringType,
	"synchronization_min_delay":     types.Int64Type,
	"synchronization_status":        types.StringType,
	"synchronization_status_detail": types.StringType,
	"use_log_destination":           types.BoolType,
	"use_ms_max_connection":         types.BoolType,
	"use_ms_rpc_timeout_in_seconds": types.BoolType,
	"version":                       types.StringType,
}

var MsserverResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"ad_domain": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The Active Directory domain to which this server belongs (if applicable).",
	},
	"ad_sites": schema.SingleNestedAttribute{
		Attributes: MsserverAdSitesResourceSchemaAttributes,
		Optional:   true,
	},
	"ad_user": schema.SingleNestedAttribute{
		Attributes: MsserverAdUserResourceSchemaAttributes,
		Optional:   true,
	},
	"address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The address or FQDN of the server.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "User comments for this Microsoft Server",
	},
	"connection_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Result of the last RPC connection attempt made",
	},
	"connection_status_detail": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Detail of the last connection attempt made",
	},
	"dhcp_server": schema.SingleNestedAttribute{
		Attributes: MsserverDhcpServerResourceSchemaAttributes,
		Optional:   true,
	},
	"disabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Allow/forbids usage of this Microsoft Server",
	},
	"dns_server": schema.SingleNestedAttribute{
		Attributes: MsserverDnsServerResourceSchemaAttributes,
		Optional:   true,
	},
	"dns_view": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Reference to the DNS view",
	},
	"extattrs": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.",
	},
	"extattrs_all": schema.MapAttribute{
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object , including default attributes.",
		ElementType:         types.StringType,
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
	},
	"grid_member": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "eference to the assigned grid member",
	},
	"last_seen": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Timestamp of the last message received",
	},
	"log_destination": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Directs logging of sync messages either to syslog or mslog",
	},
	"log_level": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Log level for this Microsoft Server",
	},
	"login_name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Microsoft Server login name, with optional domainname",
	},
	"login_password": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Microsoft Server login password",
	},
	"managing_member": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Hostname of grid member managing this Microsoft Server",
	},
	"ms_max_connection": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Maximum number of connections to MS server",
	},
	"ms_rpc_timeout_in_seconds": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Timeout in seconds of RPC connections for this MS Server",
	},
	"network_view": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Reference to the network view",
	},
	"read_only": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enable read-only management for this Microsoft Server",
	},
	"root_ad_domain": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The root Active Directory domain to which this server belongs (if applicable).",
	},
	"server_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Gives the server name as reported by itself",
	},
	"synchronization_min_delay": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Minimum number of minutes between two synchronizations",
	},
	"synchronization_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Synchronization status summary",
	},
	"synchronization_status_detail": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Detail status if synchronization_status is ERROR",
	},
	"use_log_destination": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Override log_destination inherited from grid level",
	},
	"use_ms_max_connection": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Override grid ms_max_connection setting",
	},
	"use_ms_rpc_timeout_in_seconds": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Flag to override cluster RPC timeout",
	},
	"version": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Version of the Microsoft Server",
	},
}

func ExpandMsserver(ctx context.Context, o types.Object, diags *diag.Diagnostics) *microsoftserver.Msserver {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MsserverModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MsserverModel) Expand(ctx context.Context, diags *diag.Diagnostics) *microsoftserver.Msserver {
	if m == nil {
		return nil
	}
	to := &microsoftserver.Msserver{
		Ref:                      flex.ExpandStringPointer(m.Ref),
		AdSites:                  ExpandMsserverAdSites(ctx, m.AdSites, diags),
		AdUser:                   ExpandMsserverAdUser(ctx, m.AdUser, diags),
		Address:                  flex.ExpandStringPointer(m.Address),
		Comment:                  flex.ExpandStringPointer(m.Comment),
		DhcpServer:               ExpandMsserverDhcpServer(ctx, m.DhcpServer, diags),
		Disabled:                 flex.ExpandBoolPointer(m.Disabled),
		DnsServer:                ExpandMsserverDnsServer(ctx, m.DnsServer, diags),
		DnsView:                  flex.ExpandStringPointer(m.DnsView),
		ExtAttrs:                 ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		GridMember:               flex.ExpandStringPointer(m.GridMember),
		LogDestination:           flex.ExpandStringPointer(m.LogDestination),
		LogLevel:                 flex.ExpandStringPointer(m.LogLevel),
		LoginName:                flex.ExpandStringPointer(m.LoginName),
		LoginPassword:            flex.ExpandStringPointer(m.LoginPassword),
		MsMaxConnection:          flex.ExpandInt64Pointer(m.MsMaxConnection),
		MsRpcTimeoutInSeconds:    flex.ExpandInt64Pointer(m.MsRpcTimeoutInSeconds),
		NetworkView:              flex.ExpandStringPointer(m.NetworkView),
		ReadOnly:                 flex.ExpandBoolPointer(m.ReadOnly),
		SynchronizationMinDelay:  flex.ExpandInt64Pointer(m.SynchronizationMinDelay),
		UseLogDestination:        flex.ExpandBoolPointer(m.UseLogDestination),
		UseMsMaxConnection:       flex.ExpandBoolPointer(m.UseMsMaxConnection),
		UseMsRpcTimeoutInSeconds: flex.ExpandBoolPointer(m.UseMsRpcTimeoutInSeconds),
	}
	return to
}

func FlattenMsserver(ctx context.Context, from *microsoftserver.Msserver, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MsserverAttrTypes)
	}
	m := MsserverModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, MsserverAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MsserverModel) Flatten(ctx context.Context, from *microsoftserver.Msserver, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MsserverModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.AdDomain = flex.FlattenStringPointer(from.AdDomain)
	m.AdSites = FlattenMsserverAdSites(ctx, from.AdSites, diags)
	m.AdUser = FlattenMsserverAdUser(ctx, from.AdUser, diags)
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.ConnectionStatus = flex.FlattenStringPointer(from.ConnectionStatus)
	m.ConnectionStatusDetail = flex.FlattenStringPointer(from.ConnectionStatusDetail)
	m.DhcpServer = FlattenMsserverDhcpServer(ctx, from.DhcpServer, diags)
	m.Disabled = types.BoolPointerValue(from.Disabled)
	m.DnsServer = FlattenMsserverDnsServer(ctx, from.DnsServer, diags)
	m.DnsView = flex.FlattenStringPointer(from.DnsView)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.GridMember = flex.FlattenStringPointer(from.GridMember)
	m.LastSeen = flex.FlattenInt64Pointer(from.LastSeen)
	m.LogDestination = flex.FlattenStringPointer(from.LogDestination)
	m.LogLevel = flex.FlattenStringPointer(from.LogLevel)
	m.LoginName = flex.FlattenStringPointer(from.LoginName)
	m.LoginPassword = flex.FlattenStringPointer(from.LoginPassword)
	m.ManagingMember = flex.FlattenStringPointer(from.ManagingMember)
	m.MsMaxConnection = flex.FlattenInt64Pointer(from.MsMaxConnection)
	m.MsRpcTimeoutInSeconds = flex.FlattenInt64Pointer(from.MsRpcTimeoutInSeconds)
	m.NetworkView = flex.FlattenStringPointer(from.NetworkView)
	m.ReadOnly = types.BoolPointerValue(from.ReadOnly)
	m.RootAdDomain = flex.FlattenStringPointer(from.RootAdDomain)
	m.ServerName = flex.FlattenStringPointer(from.ServerName)
	m.SynchronizationMinDelay = flex.FlattenInt64Pointer(from.SynchronizationMinDelay)
	m.SynchronizationStatus = flex.FlattenStringPointer(from.SynchronizationStatus)
	m.SynchronizationStatusDetail = flex.FlattenStringPointer(from.SynchronizationStatusDetail)
	m.UseLogDestination = types.BoolPointerValue(from.UseLogDestination)
	m.UseMsMaxConnection = types.BoolPointerValue(from.UseMsMaxConnection)
	m.UseMsRpcTimeoutInSeconds = types.BoolPointerValue(from.UseMsRpcTimeoutInSeconds)
	m.Version = flex.FlattenStringPointer(from.Version)
}
