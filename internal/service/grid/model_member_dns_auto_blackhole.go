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

type MemberDnsAutoBlackholeModel struct {
	EnableFetchesPerServer types.Bool  `tfsdk:"enable_fetches_per_server"`
	EnableFetchesPerZone   types.Bool  `tfsdk:"enable_fetches_per_zone"`
	EnableHolddown         types.Bool  `tfsdk:"enable_holddown"`
	FetchesPerServer       types.Int64 `tfsdk:"fetches_per_server"`
	FetchesPerZone         types.Int64 `tfsdk:"fetches_per_zone"`
	FpsFreq                types.Int64 `tfsdk:"fps_freq"`
	Holddown               types.Int64 `tfsdk:"holddown"`
	HolddownThreshold      types.Int64 `tfsdk:"holddown_threshold"`
	HolddownTimeout        types.Int64 `tfsdk:"holddown_timeout"`
}

var MemberDnsAutoBlackholeAttrTypes = map[string]attr.Type{
	"enable_fetches_per_server": types.BoolType,
	"enable_fetches_per_zone":   types.BoolType,
	"enable_holddown":           types.BoolType,
	"fetches_per_server":        types.Int64Type,
	"fetches_per_zone":          types.Int64Type,
	"fps_freq":                  types.Int64Type,
	"holddown":                  types.Int64Type,
	"holddown_threshold":        types.Int64Type,
	"holddown_timeout":          types.Int64Type,
}

var MemberDnsAutoBlackholeResourceSchemaAttributes = map[string]schema.Attribute{
	"enable_fetches_per_server": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enables or disables the configuration of the maximum number of concurrent recursive queries the appliance sends to each upstream DNS server.",
	},
	"enable_fetches_per_zone": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enables or disables the configuration of the maximum number of concurrent recursive queries the appliance sends to each DNS zone.",
	},
	"enable_holddown": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enables or disables the holddown configuration when the appliance stops sending queries to non-responsive servers.",
	},
	"fetches_per_server": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The maximum number of concurrent recursive queries the appliance sends to a single upstream name server before blocking additional queries to that server.",
	},
	"fetches_per_zone": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The maximum number of concurrent recursive queries that a server sends for its domains.",
	},
	"fps_freq": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Determines how often (in number of recursive responses) the appliance recalculates the average timeout ratio for each DNS server.",
	},
	"holddown": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The holddown duration for non-responsive servers.",
	},
	"holddown_threshold": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The number of consecutive timeouts before holding down a non-responsive server.",
	},
	"holddown_timeout": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The minimum time (in seconds) that needs to be passed before a timeout occurs. Note that only these timeouts are counted towards the number of consecutive timeouts.",
	},
}

func ExpandMemberDnsAutoBlackhole(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberDnsAutoBlackhole {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberDnsAutoBlackholeModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberDnsAutoBlackholeModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberDnsAutoBlackhole {
	if m == nil {
		return nil
	}
	to := &grid.MemberDnsAutoBlackhole{
		EnableFetchesPerServer: flex.ExpandBoolPointer(m.EnableFetchesPerServer),
		EnableFetchesPerZone:   flex.ExpandBoolPointer(m.EnableFetchesPerZone),
		EnableHolddown:         flex.ExpandBoolPointer(m.EnableHolddown),
		FetchesPerServer:       flex.ExpandInt64Pointer(m.FetchesPerServer),
		FetchesPerZone:         flex.ExpandInt64Pointer(m.FetchesPerZone),
		FpsFreq:                flex.ExpandInt64Pointer(m.FpsFreq),
		Holddown:               flex.ExpandInt64Pointer(m.Holddown),
		HolddownThreshold:      flex.ExpandInt64Pointer(m.HolddownThreshold),
		HolddownTimeout:        flex.ExpandInt64Pointer(m.HolddownTimeout),
	}
	return to
}

func FlattenMemberDnsAutoBlackhole(ctx context.Context, from *grid.MemberDnsAutoBlackhole, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberDnsAutoBlackholeAttrTypes)
	}
	m := MemberDnsAutoBlackholeModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberDnsAutoBlackholeAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberDnsAutoBlackholeModel) Flatten(ctx context.Context, from *grid.MemberDnsAutoBlackhole, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberDnsAutoBlackholeModel{}
	}
	m.EnableFetchesPerServer = types.BoolPointerValue(from.EnableFetchesPerServer)
	m.EnableFetchesPerZone = types.BoolPointerValue(from.EnableFetchesPerZone)
	m.EnableHolddown = types.BoolPointerValue(from.EnableHolddown)
	m.FetchesPerServer = flex.FlattenInt64Pointer(from.FetchesPerServer)
	m.FetchesPerZone = flex.FlattenInt64Pointer(from.FetchesPerZone)
	m.FpsFreq = flex.FlattenInt64Pointer(from.FpsFreq)
	m.Holddown = flex.FlattenInt64Pointer(from.Holddown)
	m.HolddownThreshold = flex.FlattenInt64Pointer(from.HolddownThreshold)
	m.HolddownTimeout = flex.FlattenInt64Pointer(from.HolddownTimeout)
}
