package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dtc"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type DtcTopologyRuleDestinationModel struct {
	DestinationLink types.String `tfsdk:"destination_link"`
	Priority        types.Int64  `tfsdk:"priority"`
}

var DtcTopologyRuleDestinationAttrTypes = map[string]attr.Type{
	"destination_link": types.StringType,
	"priority":         types.Int64Type,
}

var DtcTopologyRuleDestinationResourceSchemaAttributes = map[string]schema.Attribute{
	"destination_link": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The reference to the destination DTC pool or DTC server.",
	},
	"priority": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Priority.",
	},
}

func ExpandDtcTopologyRuleDestination(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dtc.DtcTopologyRuleDestination {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DtcTopologyRuleDestinationModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DtcTopologyRuleDestinationModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dtc.DtcTopologyRuleDestination {
	if m == nil {
		return nil
	}
	to := &dtc.DtcTopologyRuleDestination{
		DestinationLink: ExpandDtcTopologyRuleDestinationLink(ctx, m.DestinationLink, diags),
		Priority:        flex.ExpandInt64Pointer(m.Priority),
	}
	return to
}

func FlattenDtcTopologyRuleDestination(ctx context.Context, from *dtc.DtcTopologyRuleDestination, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DtcTopologyRuleDestinationAttrTypes)
	}
	m := DtcTopologyRuleDestinationModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DtcTopologyRuleDestinationAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DtcTopologyRuleDestinationModel) Flatten(ctx context.Context, from *dtc.DtcTopologyRuleDestination, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DtcTopologyRuleDestinationModel{}
	}
	m.DestinationLink = flex.FlattenStringPointer(from.DestinationLink.DtcTopologyRuleDestinationDestinationLinkOneOf.Ref)
	m.Priority = flex.FlattenInt64Pointer(from.Priority)
}

func ExpandDtcTopologyRuleDestinationLink(ctx context.Context, o types.String, diags *diag.Diagnostics) *dtc.DtcTopologyRuleDestinationDestinationLink {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	return &dtc.DtcTopologyRuleDestinationDestinationLink{
		String: flex.ExpandStringPointer(o),
	}
}

func FlattenDtcTopologyRuleDestinationLink(ctx context.Context, from *dtc.DtcTopologyRuleDestinationDestinationLink, diags *diag.Diagnostics) types.String {
	if from == nil {
		return types.StringNull()
	}
	return flex.FlattenStringPointer(from.String)
}
