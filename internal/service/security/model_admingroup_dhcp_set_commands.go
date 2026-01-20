package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	"github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/plancontrol"
)

type AdmingroupDhcpSetCommandsModel struct {
	SetDhcpdRecvSockBufSize      types.Bool `tfsdk:"set_dhcpd_recv_sock_buf_size"`
	SetLogTxnId                  types.Bool `tfsdk:"set_log_txn_id"`
	SetOverloadBootp             types.Bool `tfsdk:"set_overload_bootp"`
	SetRegenerateDhcpUpdaterKeys types.Bool `tfsdk:"set_regenerate_dhcp_updater_keys"`
	EnableAll                    types.Bool `tfsdk:"enable_all"`
	DisableAll                   types.Bool `tfsdk:"disable_all"`
}

var AdmingroupDhcpSetCommandsAttrTypes = map[string]attr.Type{
	"set_dhcpd_recv_sock_buf_size":     types.BoolType,
	"set_log_txn_id":                   types.BoolType,
	"set_overload_bootp":               types.BoolType,
	"set_regenerate_dhcp_updater_keys": types.BoolType,
	"enable_all":                       types.BoolType,
	"disable_all":                      types.BoolType,
}

var AdmingroupDhcpSetCommandsResourceSchemaAttributes = map[string]schema.Attribute{
	"set_dhcpd_recv_sock_buf_size": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_log_txn_id": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_overload_bootp": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_regenerate_dhcp_updater_keys": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"enable_all": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "If True then enable all fields",
		PlanModifiers: []planmodifier.Bool{
			plancontrol.UseStateForUnknownBool(),
		},
	},
	"disable_all": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "If True then disable all fields",
		PlanModifiers: []planmodifier.Bool{
			plancontrol.UseStateForUnknownBool(),
		},
	},
}

func ExpandAdmingroupDhcpSetCommands(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.AdmingroupDhcpSetCommands {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AdmingroupDhcpSetCommandsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AdmingroupDhcpSetCommandsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.AdmingroupDhcpSetCommands {
	if m == nil {
		return nil
	}
	to := &security.AdmingroupDhcpSetCommands{
		SetDhcpdRecvSockBufSize:      flex.ExpandBoolPointer(m.SetDhcpdRecvSockBufSize),
		SetLogTxnId:                  flex.ExpandBoolPointer(m.SetLogTxnId),
		SetOverloadBootp:             flex.ExpandBoolPointer(m.SetOverloadBootp),
		SetRegenerateDhcpUpdaterKeys: flex.ExpandBoolPointer(m.SetRegenerateDhcpUpdaterKeys),
	}
	return to
}

func FlattenAdmingroupDhcpSetCommands(ctx context.Context, from *security.AdmingroupDhcpSetCommands, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AdmingroupDhcpSetCommandsAttrTypes)
	}
	m := AdmingroupDhcpSetCommandsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AdmingroupDhcpSetCommandsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AdmingroupDhcpSetCommandsModel) Flatten(ctx context.Context, from *security.AdmingroupDhcpSetCommands, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AdmingroupDhcpSetCommandsModel{}
	}
	m.SetDhcpdRecvSockBufSize = types.BoolPointerValue(from.SetDhcpdRecvSockBufSize)
	m.SetLogTxnId = types.BoolPointerValue(from.SetLogTxnId)
	m.SetOverloadBootp = types.BoolPointerValue(from.SetOverloadBootp)
	m.SetRegenerateDhcpUpdaterKeys = types.BoolPointerValue(from.SetRegenerateDhcpUpdaterKeys)
	m.EnableAll = types.BoolPointerValue(from.EnableAll)
	m.DisableAll = types.BoolPointerValue(from.DisableAll)
}
