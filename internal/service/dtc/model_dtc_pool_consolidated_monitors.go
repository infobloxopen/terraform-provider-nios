package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dtc"
	internaltypes "github.com/Infoblox-CTO/infoblox-nios-terraform/internal/types"

	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/flex"
)

type DtcPoolConsolidatedMonitorsModel struct {
	Members                 internaltypes.UnorderedListValue `tfsdk:"members"`
	Monitor                 types.String                     `tfsdk:"monitor"`
	Availability            types.String                     `tfsdk:"availability"`
	FullHealthCommunication types.Bool                       `tfsdk:"full_health_communication"`
}

var DtcPoolConsolidatedMonitorsAttrTypes = map[string]attr.Type{
	"members":                   internaltypes.UnorderedListOfStringType,
	"monitor":                   types.StringType,
	"availability":              types.StringType,
	"full_health_communication": types.BoolType,
}

var DtcPoolConsolidatedMonitorsResourceSchemaAttributes = map[string]schema.Attribute{
	"members": schema.ListAttribute{
		CustomType:          internaltypes.UnorderedListOfStringType,
		ElementType:         types.StringType,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Members whose monitor statuses are shared across other members in a pool.",
	},
	"monitor": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Monitor whose statuses are shared across other members in a pool.",
	},
	"availability": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Servers assigned to a pool with monitor defined are healthy if ANY or ALL members report healthy status.",
	},
	"full_health_communication": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Flag for switching health performing and sharing behavior to perform health checks on each DTC grid member that serves related LBDN(s) and send them across all DTC grid members from both selected and non-selected lists.",
	},
}

func ExpandDtcPoolConsolidatedMonitors(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dtc.DtcPoolConsolidatedMonitors {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DtcPoolConsolidatedMonitorsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DtcPoolConsolidatedMonitorsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dtc.DtcPoolConsolidatedMonitors {
	if m == nil {
		return nil
	}
	to := &dtc.DtcPoolConsolidatedMonitors{
		Members:                 flex.ExpandFrameworkListString(ctx, m.Members, diags),
		Monitor:                 flex.ExpandStringPointer(m.Monitor),
		Availability:            flex.ExpandStringPointer(m.Availability),
		FullHealthCommunication: flex.ExpandBoolPointer(m.FullHealthCommunication),
	}
	return to
}

func FlattenDtcPoolConsolidatedMonitors(ctx context.Context, from *dtc.DtcPoolConsolidatedMonitors, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DtcPoolConsolidatedMonitorsAttrTypes)
	}
	m := DtcPoolConsolidatedMonitorsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DtcPoolConsolidatedMonitorsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DtcPoolConsolidatedMonitorsModel) Flatten(ctx context.Context, from *dtc.DtcPoolConsolidatedMonitors, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DtcPoolConsolidatedMonitorsModel{}
	}
	m.Members = flex.FlattenFrameworkUnorderedList(ctx, types.StringType, from.Members, diags)
	m.Monitor = flex.FlattenStringPointer(from.Monitor)
	m.Availability = flex.FlattenStringPointer(from.Availability)
	m.FullHealthCommunication = types.BoolPointerValue(from.FullHealthCommunication)
}
