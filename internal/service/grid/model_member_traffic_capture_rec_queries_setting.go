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

type MemberTrafficCaptureRecQueriesSettingModel struct {
	RecursiveClientsCountTriggerEnable types.Bool  `tfsdk:"recursive_clients_count_trigger_enable"`
	RecursiveClientsCountThreshold     types.Int64 `tfsdk:"recursive_clients_count_threshold"`
	RecursiveClientsCountReset         types.Int64 `tfsdk:"recursive_clients_count_reset"`
}

var MemberTrafficCaptureRecQueriesSettingAttrTypes = map[string]attr.Type{
	"recursive_clients_count_trigger_enable": types.BoolType,
	"recursive_clients_count_threshold":      types.Int64Type,
	"recursive_clients_count_reset":          types.Int64Type,
}

var MemberTrafficCaptureRecQueriesSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"recursive_clients_count_trigger_enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enable triggering automated traffic capture based on outgoing recursive queries count.",
	},
	"recursive_clients_count_threshold": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Concurrent outgoing recursive queries count below which traffic capture will be triggered.",
	},
	"recursive_clients_count_reset": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Concurrent outgoing recursive queries count below which traffic capture will be stopped.",
	},
}

func ExpandMemberTrafficCaptureRecQueriesSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberTrafficCaptureRecQueriesSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberTrafficCaptureRecQueriesSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberTrafficCaptureRecQueriesSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberTrafficCaptureRecQueriesSetting {
	if m == nil {
		return nil
	}
	to := &grid.MemberTrafficCaptureRecQueriesSetting{
		RecursiveClientsCountTriggerEnable: flex.ExpandBoolPointer(m.RecursiveClientsCountTriggerEnable),
		RecursiveClientsCountThreshold:     flex.ExpandInt64Pointer(m.RecursiveClientsCountThreshold),
		RecursiveClientsCountReset:         flex.ExpandInt64Pointer(m.RecursiveClientsCountReset),
	}
	return to
}

func FlattenMemberTrafficCaptureRecQueriesSetting(ctx context.Context, from *grid.MemberTrafficCaptureRecQueriesSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberTrafficCaptureRecQueriesSettingAttrTypes)
	}
	m := MemberTrafficCaptureRecQueriesSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberTrafficCaptureRecQueriesSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberTrafficCaptureRecQueriesSettingModel) Flatten(ctx context.Context, from *grid.MemberTrafficCaptureRecQueriesSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberTrafficCaptureRecQueriesSettingModel{}
	}
	m.RecursiveClientsCountTriggerEnable = types.BoolPointerValue(from.RecursiveClientsCountTriggerEnable)
	m.RecursiveClientsCountThreshold = flex.FlattenInt64Pointer(from.RecursiveClientsCountThreshold)
	m.RecursiveClientsCountReset = flex.FlattenInt64Pointer(from.RecursiveClientsCountReset)
}
