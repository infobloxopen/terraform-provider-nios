package misc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/misc"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type OutboundCloudclientOutboundCloudClientEventsModel struct {
	EventType types.String `tfsdk:"event_type"`
	Enabled   types.Bool   `tfsdk:"enabled"`
}

var OutboundCloudclientOutboundCloudClientEventsAttrTypes = map[string]attr.Type{
	"event_type": types.StringType,
	"enabled":    types.BoolType,
}

var OutboundCloudclientOutboundCloudClientEventsResourceSchemaAttributes = map[string]schema.Attribute{
	"event_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The event type enum rpz and analytics.",
	},
	"enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the event type is enabled or not.",
	},
}

func ExpandOutboundCloudclientOutboundCloudClientEvents(ctx context.Context, o types.Object, diags *diag.Diagnostics) *misc.OutboundCloudclientOutboundCloudClientEvents {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m OutboundCloudclientOutboundCloudClientEventsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *OutboundCloudclientOutboundCloudClientEventsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *misc.OutboundCloudclientOutboundCloudClientEvents {
	if m == nil {
		return nil
	}
	to := &misc.OutboundCloudclientOutboundCloudClientEvents{
		EventType: flex.ExpandStringPointer(m.EventType),
		Enabled:   flex.ExpandBoolPointer(m.Enabled),
	}
	return to
}

func FlattenOutboundCloudclientOutboundCloudClientEvents(ctx context.Context, from *misc.OutboundCloudclientOutboundCloudClientEvents, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(OutboundCloudclientOutboundCloudClientEventsAttrTypes)
	}
	m := OutboundCloudclientOutboundCloudClientEventsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, OutboundCloudclientOutboundCloudClientEventsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *OutboundCloudclientOutboundCloudClientEventsModel) Flatten(ctx context.Context, from *misc.OutboundCloudclientOutboundCloudClientEvents, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = OutboundCloudclientOutboundCloudClientEventsModel{}
	}
	m.EventType = flex.FlattenStringPointer(from.EventType)
	m.Enabled = types.BoolPointerValue(from.Enabled)
}
