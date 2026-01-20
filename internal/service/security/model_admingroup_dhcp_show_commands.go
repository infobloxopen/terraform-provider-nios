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

type AdmingroupDhcpShowCommandsModel struct {
	ShowDhcpGssTsig          types.Bool `tfsdk:"show_dhcp_gss_tsig"`
	ShowDhcpv6GssTsig        types.Bool `tfsdk:"show_dhcpv6_gss_tsig"`
	ShowDhcpdRecvSockBufSize types.Bool `tfsdk:"show_dhcpd_recv_sock_buf_size"`
	ShowOverloadBootp        types.Bool `tfsdk:"show_overload_bootp"`
	ShowLogTxnId             types.Bool `tfsdk:"show_log_txn_id"`
	EnableAll                types.Bool `tfsdk:"enable_all"`
	DisableAll               types.Bool `tfsdk:"disable_all"`
}

var AdmingroupDhcpShowCommandsAttrTypes = map[string]attr.Type{
	"show_dhcp_gss_tsig":            types.BoolType,
	"show_dhcpv6_gss_tsig":          types.BoolType,
	"show_dhcpd_recv_sock_buf_size": types.BoolType,
	"show_overload_bootp":           types.BoolType,
	"show_log_txn_id":               types.BoolType,
	"enable_all":                    types.BoolType,
	"disable_all":                   types.BoolType,
}

var AdmingroupDhcpShowCommandsResourceSchemaAttributes = map[string]schema.Attribute{
	"show_dhcp_gss_tsig": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_dhcpv6_gss_tsig": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_dhcpd_recv_sock_buf_size": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_overload_bootp": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_log_txn_id": schema.BoolAttribute{
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

func ExpandAdmingroupDhcpShowCommands(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.AdmingroupDhcpShowCommands {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AdmingroupDhcpShowCommandsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AdmingroupDhcpShowCommandsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.AdmingroupDhcpShowCommands {
	if m == nil {
		return nil
	}
	to := &security.AdmingroupDhcpShowCommands{
		ShowDhcpGssTsig:          flex.ExpandBoolPointer(m.ShowDhcpGssTsig),
		ShowDhcpv6GssTsig:        flex.ExpandBoolPointer(m.ShowDhcpv6GssTsig),
		ShowDhcpdRecvSockBufSize: flex.ExpandBoolPointer(m.ShowDhcpdRecvSockBufSize),
		ShowOverloadBootp:        flex.ExpandBoolPointer(m.ShowOverloadBootp),
		ShowLogTxnId:             flex.ExpandBoolPointer(m.ShowLogTxnId),
	}
	return to
}

func FlattenAdmingroupDhcpShowCommands(ctx context.Context, from *security.AdmingroupDhcpShowCommands, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AdmingroupDhcpShowCommandsAttrTypes)
	}
	m := AdmingroupDhcpShowCommandsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AdmingroupDhcpShowCommandsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AdmingroupDhcpShowCommandsModel) Flatten(ctx context.Context, from *security.AdmingroupDhcpShowCommands, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AdmingroupDhcpShowCommandsModel{}
	}
	m.ShowDhcpGssTsig = types.BoolPointerValue(from.ShowDhcpGssTsig)
	m.ShowDhcpv6GssTsig = types.BoolPointerValue(from.ShowDhcpv6GssTsig)
	m.ShowDhcpdRecvSockBufSize = types.BoolPointerValue(from.ShowDhcpdRecvSockBufSize)
	m.ShowOverloadBootp = types.BoolPointerValue(from.ShowOverloadBootp)
	m.ShowLogTxnId = types.BoolPointerValue(from.ShowLogTxnId)
	m.EnableAll = types.BoolPointerValue(from.EnableAll)
	m.DisableAll = types.BoolPointerValue(from.DisableAll)
}
