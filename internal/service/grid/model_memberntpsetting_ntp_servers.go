package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type MemberntpsettingNtpServersModel struct {
	Address              types.String `tfsdk:"address"`
	EnableAuthentication types.Bool   `tfsdk:"enable_authentication"`
	NtpKeyNumber         types.Int64  `tfsdk:"ntp_key_number"`
	Preferred            types.Bool   `tfsdk:"preferred"`
	Burst                types.Bool   `tfsdk:"burst"`
	Iburst               types.Bool   `tfsdk:"iburst"`
}

var MemberntpsettingNtpServersAttrTypes = map[string]attr.Type{
	"address":               types.StringType,
	"enable_authentication": types.BoolType,
	"ntp_key_number":        types.Int64Type,
	"preferred":             types.BoolType,
	"burst":                 types.BoolType,
	"iburst":                types.BoolType,
}

var MemberntpsettingNtpServersResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The NTP server IP address or FQDN.",
	},
	"enable_authentication": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the NTP authentication is enabled.",
	},
	"ntp_key_number": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The NTP authentication key number.",
	},
	"preferred": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the NTP server is a preferred one or not.",
	},
	"burst": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the BURST operation mode is enabled. In BURST operating mode, when the external server is reachable and a valid source of synchronization is available, NTP sends a burst of 8 packets with a 2 second interval between packets.",
	},
	"iburst": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the IBURST operation mode is enabled. In IBURST operating mode, when the external server is unreachable, NTP server sends a burst of 8 packets with a 2 second interval between packets.",
	},
}

func ExpandMemberntpsettingNtpServers(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberntpsettingNtpServers {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberntpsettingNtpServersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberntpsettingNtpServersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberntpsettingNtpServers {
	if m == nil {
		return nil
	}
	to := &grid.MemberntpsettingNtpServers{
		Address:              flex.ExpandStringPointer(m.Address),
		EnableAuthentication: flex.ExpandBoolPointer(m.EnableAuthentication),
		NtpKeyNumber:         flex.ExpandInt64Pointer(m.NtpKeyNumber),
		Preferred:            flex.ExpandBoolPointer(m.Preferred),
		Burst:                flex.ExpandBoolPointer(m.Burst),
		Iburst:               flex.ExpandBoolPointer(m.Iburst),
	}
	return to
}

func FlattenMemberntpsettingNtpServers(ctx context.Context, from *grid.MemberntpsettingNtpServers, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberntpsettingNtpServersAttrTypes)
	}
	m := MemberntpsettingNtpServersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberntpsettingNtpServersAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberntpsettingNtpServersModel) Flatten(ctx context.Context, from *grid.MemberntpsettingNtpServers, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberntpsettingNtpServersModel{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.EnableAuthentication = types.BoolPointerValue(from.EnableAuthentication)
	m.NtpKeyNumber = flex.FlattenInt64Pointer(from.NtpKeyNumber)
	m.Preferred = types.BoolPointerValue(from.Preferred)
	m.Burst = types.BoolPointerValue(from.Burst)
	m.Iburst = types.BoolPointerValue(from.Iburst)
}
