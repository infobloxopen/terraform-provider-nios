package misc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/misc"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type OutboundCloudclientModel struct {
	Ref                       types.String `tfsdk:"ref"`
	Enable                    types.Bool   `tfsdk:"enable"`
	GridMember                types.String `tfsdk:"grid_member"`
	Interval                  types.Int64  `tfsdk:"interval"`
	OutboundCloudClientEvents types.List   `tfsdk:"outbound_cloud_client_events"`
}

var OutboundCloudclientAttrTypes = map[string]attr.Type{
	"ref":                          types.StringType,
	"enable":                       types.BoolType,
	"grid_member":                  types.StringType,
	"interval":                     types.Int64Type,
	"outbound_cloud_client_events": types.ListType{ElemType: types.ObjectType{AttrTypes: OutboundCloudclientOutboundCloudClientEventsAttrTypes}},
}

var OutboundCloudclientResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the OutBound Cloud Client is enabled.",
	},
	"grid_member": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The Grid member where our outbound is running.",
	},
	"interval": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The time interval (in seconds) for requesting newly detected domains by the Infoblox Outbound Cloud Client and applying them to the list of configured RPZs.",
	},
	"outbound_cloud_client_events": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: OutboundCloudclientOutboundCloudClientEventsResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "List of event types to request",
	},
}

func ExpandOutboundCloudclient(ctx context.Context, o types.Object, diags *diag.Diagnostics) *misc.OutboundCloudclient {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m OutboundCloudclientModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *OutboundCloudclientModel) Expand(ctx context.Context, diags *diag.Diagnostics) *misc.OutboundCloudclient {
	if m == nil {
		return nil
	}
	to := &misc.OutboundCloudclient{
		Ref:                       flex.ExpandStringPointer(m.Ref),
		Enable:                    flex.ExpandBoolPointer(m.Enable),
		GridMember:                flex.ExpandStringPointer(m.GridMember),
		Interval:                  flex.ExpandInt64Pointer(m.Interval),
		OutboundCloudClientEvents: flex.ExpandFrameworkListNestedBlock(ctx, m.OutboundCloudClientEvents, diags, ExpandOutboundCloudclientOutboundCloudClientEvents),
	}
	return to
}

func FlattenOutboundCloudclient(ctx context.Context, from *misc.OutboundCloudclient, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(OutboundCloudclientAttrTypes)
	}
	m := OutboundCloudclientModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, OutboundCloudclientAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *OutboundCloudclientModel) Flatten(ctx context.Context, from *misc.OutboundCloudclient, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = OutboundCloudclientModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Enable = types.BoolPointerValue(from.Enable)
	m.GridMember = flex.FlattenStringPointer(from.GridMember)
	m.Interval = flex.FlattenInt64Pointer(from.Interval)
	m.OutboundCloudClientEvents = flex.FlattenFrameworkListNestedBlock(ctx, from.OutboundCloudClientEvents, OutboundCloudclientOutboundCloudClientEventsAttrTypes, diags, FlattenOutboundCloudclientOutboundCloudClientEvents)
}
