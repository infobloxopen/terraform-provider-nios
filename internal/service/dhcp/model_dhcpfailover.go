package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/boolvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/import"
)

type DhcpfailoverModel struct {
	Ref                        types.String `tfsdk:"ref"`
	AssociationType            types.String `tfsdk:"association_type"`
	Comment                    types.String `tfsdk:"comment"`
	ExtAttrs                   types.Map    `tfsdk:"extattrs"`
	FailoverPort               types.Int64  `tfsdk:"failover_port"`
	LoadBalanceSplit           types.Int64  `tfsdk:"load_balance_split"`
	MaxClientLeadTime          types.Int64  `tfsdk:"max_client_lead_time"`
	MaxLoadBalanceDelay        types.Int64  `tfsdk:"max_load_balance_delay"`
	MaxResponseDelay           types.Int64  `tfsdk:"max_response_delay"`
	MaxUnackedUpdates          types.Int64  `tfsdk:"max_unacked_updates"`
	MsAssociationMode          types.String `tfsdk:"ms_association_mode"`
	MsEnableAuthentication     types.Bool   `tfsdk:"ms_enable_authentication"`
	MsEnableSwitchoverInterval types.Bool   `tfsdk:"ms_enable_switchover_interval"`
	MsFailoverMode             types.String `tfsdk:"ms_failover_mode"`
	MsFailoverPartner          types.String `tfsdk:"ms_failover_partner"`
	MsHotstandbyPartnerRole    types.String `tfsdk:"ms_hotstandby_partner_role"`
	MsIsConflict               types.Bool   `tfsdk:"ms_is_conflict"`
	MsPreviousState            types.String `tfsdk:"ms_previous_state"`
	MsServer                   types.String `tfsdk:"ms_server"`
	MsSharedSecret             types.String `tfsdk:"ms_shared_secret"`
	MsState                    types.String `tfsdk:"ms_state"`
	MsSwitchoverInterval       types.Int64  `tfsdk:"ms_switchover_interval"`
	Name                       types.String `tfsdk:"name"`
	Primary                    types.String `tfsdk:"primary"`
	PrimaryServerType          types.String `tfsdk:"primary_server_type"`
	PrimaryState               types.String `tfsdk:"primary_state"`
	RecycleLeases              types.Bool   `tfsdk:"recycle_leases"`
	Secondary                  types.String `tfsdk:"secondary"`
	SecondaryServerType        types.String `tfsdk:"secondary_server_type"`
	SecondaryState             types.String `tfsdk:"secondary_state"`
	UseFailoverPort            types.Bool   `tfsdk:"use_failover_port"`
	UseMsSwitchoverInterval    types.Bool   `tfsdk:"use_ms_switchover_interval"`
	UseRecycleLeases           types.Bool   `tfsdk:"use_recycle_leases"`
	ExtAttrsAll                types.Map    `tfsdk:"extattrs_all"`
}

var DhcpfailoverAttrTypes = map[string]attr.Type{
	"ref":                           types.StringType,
	"association_type":              types.StringType,
	"comment":                       types.StringType,
	"extattrs":                      types.MapType{ElemType: types.StringType},
	"failover_port":                 types.Int64Type,
	"load_balance_split":            types.Int64Type,
	"max_client_lead_time":          types.Int64Type,
	"max_load_balance_delay":        types.Int64Type,
	"max_response_delay":            types.Int64Type,
	"max_unacked_updates":           types.Int64Type,
	"ms_association_mode":           types.StringType,
	"ms_enable_authentication":      types.BoolType,
	"ms_enable_switchover_interval": types.BoolType,
	"ms_failover_mode":              types.StringType,
	"ms_failover_partner":           types.StringType,
	"ms_hotstandby_partner_role":    types.StringType,
	"ms_is_conflict":                types.BoolType,
	"ms_previous_state":             types.StringType,
	"ms_server":                     types.StringType,
	"ms_shared_secret":              types.StringType,
	"ms_state":                      types.StringType,
	"ms_switchover_interval":        types.Int64Type,
	"name":                          types.StringType,
	"primary":                       types.StringType,
	"primary_server_type":           types.StringType,
	"primary_state":                 types.StringType,
	"recycle_leases":                types.BoolType,
	"secondary":                     types.StringType,
	"secondary_server_type":         types.StringType,
	"secondary_state":               types.StringType,
	"use_failover_port":             types.BoolType,
	"use_ms_switchover_interval":    types.BoolType,
	"use_recycle_leases":            types.BoolType,
	"extattrs_all":                  types.MapType{ElemType: types.StringType},
}

var DhcpfailoverResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"association_type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The value indicating whether the failover association is Microsoft or Grid based. This is a read-only attribute.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The descriptive comment of a DHCP MAC Filter object.",
	},
	"extattrs": schema.MapAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object.",
		ElementType:         types.StringType,
		Default:             mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
	},
	"extattrs_all": schema.MapAttribute{
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object , including default attributes.",
		ElementType:         types.StringType,
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
	},
	"failover_port": schema.Int64Attribute{
		Computed: true,
		Optional: true,
		Default:  int64default.StaticInt64(647),
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_failover_port")),
			int64validator.Between(1, 63999),
		},
		MarkdownDescription: "Determines the TCP port on which the server should listen for connections from its failover peer. Valid values are between 1 and 63999.",
	},
	"load_balance_split": schema.Int64Attribute{
		Computed: true,
		Optional: true,
		Default:  int64default.StaticInt64(128),
		Validators: []validator.Int64{
			int64validator.Between(0, 256),
		},
		MarkdownDescription: "A load balancing split value of a DHCP failover object. Specify the value of the maximum load balancing delay in a 8-bit integer format (range from 0 to 256).",
	},
	"max_client_lead_time": schema.Int64Attribute{
		Computed: true,
		Optional: true,
		Default:  int64default.StaticInt64(3600),
		Validators: []validator.Int64{
			int64validator.Between(1, 4294967295),
		},
		MarkdownDescription: "The maximum client lead time value of a DHCP failover object. Specify the value of the maximum client lead time in a 32-bit integer format (range from 0 to 4294967295) that represents the duration in seconds. Valid values are between 1 and 4294967295.",
	},
	"max_load_balance_delay": schema.Int64Attribute{
		Computed: true,
		Optional: true,
		Default:  int64default.StaticInt64(3),
		Validators: []validator.Int64{
			int64validator.Between(1, 4294967295),
		},
		MarkdownDescription: "The maximum load balancing delay value of a DHCP failover object. Specify the value of the maximum load balancing delay in a 32-bit integer format (range from 0 to 4294967295) that represents the duration in seconds. Valid values are between 1 and 4294967295.",
	},
	"max_response_delay": schema.Int64Attribute{
		Computed: true,
		Optional: true,
		Default:  int64default.StaticInt64(60),
		Validators: []validator.Int64{
			int64validator.Between(1, 4294967295),
		},
		MarkdownDescription: "The maximum response delay value of a DHCP failover object. Specify the value of the maximum response delay in a 32-bit integer format (range from 0 to 4294967295) that represents the duration in seconds. Valid values are between 1 and 4294967295.",
	},
	"max_unacked_updates": schema.Int64Attribute{
		Computed: true,
		Optional: true,
		Default:  int64default.StaticInt64(10),
		Validators: []validator.Int64{
			int64validator.Between(1, 4294967295),
		},
		MarkdownDescription: "The maximum number of unacked updates value of a DHCP failover object. Specify the value of the maximum number of unacked updates in a 32-bit integer format (range from 0 to 4294967295) that represents the number of messages. Valid values are between 1 and 4294967295.",
	},
	"ms_association_mode": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The value that indicates whether the failover association is read-write or read-only. This is a read-only attribute.",
	},
	"ms_enable_authentication": schema.BoolAttribute{
		Computed:            true,
		Optional:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Determines if the authentication for the failover association is enabled or not.",
	},
	"ms_enable_switchover_interval": schema.BoolAttribute{
		Computed: true,
		Optional: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_ms_switchover_interval")),
		},
		MarkdownDescription: "Determines if the switchover interval is enabled or not.",
	},
	"ms_failover_mode": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Default:  stringdefault.StaticString("LOADBALANCE"),
		Validators: []validator.String{
			stringvalidator.OneOf("LOADBALANCE", "HOTSTANDBY"),
		},
		MarkdownDescription: "The mode for the failover association.",
	},
	"ms_failover_partner": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Failover partner defined in the association with the Microsoft Server.",
	},
	"ms_hotstandby_partner_role": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Validators: []validator.String{
			stringvalidator.OneOf("ACTIVE", "PASSIVE"),
		},
		MarkdownDescription: "The partner role in the case of HotStandby.",
	},
	"ms_is_conflict": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines if the matching Microsoft failover association (if any) is in synchronization (False) or not (True). If there is no matching failover association the returned values is False. This is a read-only attribute.",
	},
	"ms_previous_state": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The previous failover association state. This is a read-only attribute.",
	},
	"ms_server": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The primary Microsoft Server.",
	},
	"ms_shared_secret": schema.StringAttribute{
		Computed:            true,
		Optional:            true,
		Sensitive:           true,
		MarkdownDescription: "The failover association authentication. This is a write-only attribute.",
	},
	"ms_state": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The failover association state. This is a read-only attribute.",
	},
	"ms_switchover_interval": schema.Int64Attribute{
		Computed: true,
		Optional: true,
		Default:  int64default.StaticInt64(3600),
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_ms_switchover_interval")),
		},
		MarkdownDescription: "The time (in seconds) that DHCPv4 server will wait before transitioning the server from the COMMUNICATION-INT state to PARTNER-DOWN state.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of a DHCP failover object.",
	},
	"primary": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The primary server of a DHCP failover object.",
	},
	"primary_server_type": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.OneOf("EXTERNAL", "GRID"),
		},
		MarkdownDescription: "The type of the primary server of DHCP Failover association object.",
	},
	"primary_state": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The primary server status of a DHCP failover object.",
	},
	"recycle_leases": schema.BoolAttribute{
		Computed: true,
		Optional: true,
		Default:  booldefault.StaticBool(true),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_recycle_leases")),
		},
		MarkdownDescription: "Determines if the leases are kept in recycle bin until one week after expiration or not.",
	},
	"secondary": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The secondary server of a DHCP failover object.",
	},
	"secondary_server_type": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.OneOf("EXTERNAL", "GRID"),
		},
		MarkdownDescription: "The type of the secondary server of DHCP Failover association object.",
	},
	"secondary_state": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The secondary server status of a DHCP failover object.",
	},
	"use_failover_port": schema.BoolAttribute{
		Computed:            true,
		Optional:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: failover_port",
	},
	"use_ms_switchover_interval": schema.BoolAttribute{
		Computed:            true,
		Optional:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: ms_switchover_interval",
	},
	"use_recycle_leases": schema.BoolAttribute{
		Computed:            true,
		Optional:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: recycle_leases",
	},
}

func ExpandDhcpfailover(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dhcp.Dhcpfailover {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DhcpfailoverModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DhcpfailoverModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.Dhcpfailover {
	if m == nil {
		return nil
	}
	to := &dhcp.Dhcpfailover{
		Comment:                    flex.ExpandStringPointer(m.Comment),
		ExtAttrs:                   ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		FailoverPort:               flex.ExpandInt64Pointer(m.FailoverPort),
		LoadBalanceSplit:           flex.ExpandInt64Pointer(m.LoadBalanceSplit),
		MaxClientLeadTime:          flex.ExpandInt64Pointer(m.MaxClientLeadTime),
		MaxLoadBalanceDelay:        flex.ExpandInt64Pointer(m.MaxLoadBalanceDelay),
		MaxResponseDelay:           flex.ExpandInt64Pointer(m.MaxResponseDelay),
		MaxUnackedUpdates:          flex.ExpandInt64Pointer(m.MaxUnackedUpdates),
		MsEnableAuthentication:     flex.ExpandBoolPointer(m.MsEnableAuthentication),
		MsEnableSwitchoverInterval: flex.ExpandBoolPointer(m.MsEnableSwitchoverInterval),
		MsFailoverMode:             flex.ExpandStringPointer(m.MsFailoverMode),
		MsFailoverPartner:          flex.ExpandStringPointer(m.MsFailoverPartner),
		MsHotstandbyPartnerRole:    flex.ExpandStringPointer(m.MsHotstandbyPartnerRole),
		MsServer:                   flex.ExpandStringPointer(m.MsServer),
		MsSharedSecret:             flex.ExpandStringPointer(m.MsSharedSecret),
		MsSwitchoverInterval:       flex.ExpandInt64Pointer(m.MsSwitchoverInterval),
		Name:                       flex.ExpandStringPointer(m.Name),
		Primary:                    flex.ExpandStringPointer(m.Primary),
		PrimaryServerType:          flex.ExpandStringPointer(m.PrimaryServerType),
		RecycleLeases:              flex.ExpandBoolPointer(m.RecycleLeases),
		Secondary:                  flex.ExpandStringPointer(m.Secondary),
		SecondaryServerType:        flex.ExpandStringPointer(m.SecondaryServerType),
		UseFailoverPort:            flex.ExpandBoolPointer(m.UseFailoverPort),
		UseMsSwitchoverInterval:    flex.ExpandBoolPointer(m.UseMsSwitchoverInterval),
		UseRecycleLeases:           flex.ExpandBoolPointer(m.UseRecycleLeases),
	}
	return to
}

func FlattenDhcpfailover(ctx context.Context, from *dhcp.Dhcpfailover, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DhcpfailoverAttrTypes)
	}
	m := DhcpfailoverModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, DhcpfailoverAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DhcpfailoverModel) Flatten(ctx context.Context, from *dhcp.Dhcpfailover, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DhcpfailoverModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.AssociationType = flex.FlattenStringPointer(from.AssociationType)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.FailoverPort = flex.FlattenInt64Pointer(from.FailoverPort)
	m.LoadBalanceSplit = flex.FlattenInt64Pointer(from.LoadBalanceSplit)
	m.MaxClientLeadTime = flex.FlattenInt64Pointer(from.MaxClientLeadTime)
	m.MaxLoadBalanceDelay = flex.FlattenInt64Pointer(from.MaxLoadBalanceDelay)
	m.MaxResponseDelay = flex.FlattenInt64Pointer(from.MaxResponseDelay)
	m.MaxUnackedUpdates = flex.FlattenInt64Pointer(from.MaxUnackedUpdates)
	m.MsAssociationMode = flex.FlattenStringPointer(from.MsAssociationMode)
	m.MsEnableAuthentication = types.BoolPointerValue(from.MsEnableAuthentication)
	m.MsEnableSwitchoverInterval = types.BoolPointerValue(from.MsEnableSwitchoverInterval)
	m.MsFailoverMode = flex.FlattenStringPointer(from.MsFailoverMode)
	m.MsFailoverPartner = flex.FlattenStringPointer(from.MsFailoverPartner)
	m.MsHotstandbyPartnerRole = flex.FlattenStringPointer(from.MsHotstandbyPartnerRole)
	m.MsIsConflict = types.BoolPointerValue(from.MsIsConflict)
	m.MsPreviousState = flex.FlattenStringPointer(from.MsPreviousState)
	m.MsServer = flex.FlattenStringPointer(from.MsServer)
	m.MsSharedSecret = flex.FlattenStringPointer(from.MsSharedSecret)
	m.MsState = flex.FlattenStringPointer(from.MsState)
	m.MsSwitchoverInterval = flex.FlattenInt64Pointer(from.MsSwitchoverInterval)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Primary = flex.FlattenStringPointer(from.Primary)
	m.PrimaryServerType = flex.FlattenStringPointer(from.PrimaryServerType)
	m.PrimaryState = flex.FlattenStringPointer(from.PrimaryState)
	m.RecycleLeases = types.BoolPointerValue(from.RecycleLeases)
	m.Secondary = flex.FlattenStringPointer(from.Secondary)
	m.SecondaryServerType = flex.FlattenStringPointer(from.SecondaryServerType)
	m.SecondaryState = flex.FlattenStringPointer(from.SecondaryState)
	m.UseFailoverPort = types.BoolPointerValue(from.UseFailoverPort)
	m.UseMsSwitchoverInterval = types.BoolPointerValue(from.UseMsSwitchoverInterval)
	m.UseRecycleLeases = types.BoolPointerValue(from.UseRecycleLeases)
}
