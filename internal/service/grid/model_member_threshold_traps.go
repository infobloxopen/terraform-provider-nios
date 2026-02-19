package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type MemberThresholdTrapsModel struct {
	TrapType    types.String `tfsdk:"trap_type"`
	TrapReset   types.Int64  `tfsdk:"trap_reset"`
	TrapTrigger types.Int64  `tfsdk:"trap_trigger"`
}

var MemberThresholdTrapsAttrTypes = map[string]attr.Type{
	"trap_type":    types.StringType,
	"trap_reset":   types.Int64Type,
	"trap_trigger": types.Int64Type,
}

var MemberThresholdTrapsResourceSchemaAttributes = map[string]schema.Attribute{
	"trap_type": schema.StringAttribute{
		Computed: true,
		Optional: true, //https://infoblox.atlassian.net/browse/NIOS-109174
		Validators: []validator.String{
			stringvalidator.OneOf("CpuUsage", "DBObjects", "Disk", "ExtStorage", "FDUsage", "FastpathDroppedTraffic", "Fastpathmbuffdepletion", "IPAMUtilization", "Memory", "NetworkCapacity", "RPZHitRate", "RecursiveClients", "Reporting", "ReportingVolume", "Rootfs", "SwapUsage", "TcpUdpFloodAlertRate", "TcpUdpFloodDropRate", "ThreatProtectionDroppedTraffic", "ThreatProtectionTotalTraffic", "Tmpfs"),
		},
		MarkdownDescription: "Determines the type of a given trap.",
	},
	"trap_reset": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		// Default:  int64default.StaticInt64(COMPLEX),
		// TODO: Default exists , but generator cannot determine the value.
		// Default: "",
		MarkdownDescription: "Determines the threshold value to reset the trap.",
	},
	"trap_trigger": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		// Default:  int64default.StaticInt64(COMPLEX),
		// TODO: Default exists , but generator cannot determine the value.
		// Default: "",
		MarkdownDescription: "Determines the threshold value to trigger the trap.",
	},
}

func ExpandMemberThresholdTraps(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberThresholdTraps {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberThresholdTrapsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberThresholdTrapsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberThresholdTraps {
	if m == nil {
		return nil
	}
	to := &grid.MemberThresholdTraps{
		TrapType:    flex.ExpandStringPointer(m.TrapType),
		TrapReset:   flex.ExpandInt64Pointer(m.TrapReset),
		TrapTrigger: flex.ExpandInt64Pointer(m.TrapTrigger),
	}
	return to
}

func FlattenMemberThresholdTraps(ctx context.Context, from *grid.MemberThresholdTraps, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberThresholdTrapsAttrTypes)
	}
	m := MemberThresholdTrapsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberThresholdTrapsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberThresholdTrapsModel) Flatten(ctx context.Context, from *grid.MemberThresholdTraps, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberThresholdTrapsModel{}
	}
	m.TrapType = flex.FlattenStringPointer(from.TrapType)
	m.TrapReset = flex.FlattenInt64Pointer(from.TrapReset)
	m.TrapTrigger = flex.FlattenInt64Pointer(from.TrapTrigger)
}
