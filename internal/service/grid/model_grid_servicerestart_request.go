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

type GridServicerestartRequestModel struct {
	Ref             types.String `tfsdk:"ref"`
	Uuid            types.String `tfsdk:"uuid"`
	Error           types.String `tfsdk:"error"`
	Forced          types.Bool   `tfsdk:"forced"`
	Group           types.String `tfsdk:"group"`
	LastUpdatedTime types.Int64  `tfsdk:"last_updated_time"`
	Member          types.String `tfsdk:"member"`
	Needed          types.String `tfsdk:"needed"`
	Order           types.Int64  `tfsdk:"order"`
	Result          types.String `tfsdk:"result"`
	Service         types.String `tfsdk:"service"`
	State           types.String `tfsdk:"state"`
}

var GridServicerestartRequestAttrTypes = map[string]attr.Type{
	"ref":               types.StringType,
	"uuid":              types.StringType,
	"error":             types.StringType,
	"forced":            types.BoolType,
	"group":             types.StringType,
	"last_updated_time": types.Int64Type,
	"member":            types.StringType,
	"needed":            types.StringType,
	"order":             types.Int64Type,
	"result":            types.StringType,
	"service":           types.StringType,
	"state":             types.StringType,
}

var GridServicerestartRequestResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"error": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The error message if restart has failed.",
	},
	"forced": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Indicates if this is a force restart.",
	},
	"group": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the Restart Group associated with the request.",
	},
	"last_updated_time": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The timestamp when the status of the request has changed.",
	},
	"member": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The member to restart.",
	},
	"needed": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Indicates if restart is needed.",
	},
	"order": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The order to restart.",
	},
	"result": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The result of the restart operation.",
	},
	"service": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The service to restart.",
	},
	"state": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The state of the request.",
	},
}

func ExpandGridServicerestartRequest(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridServicerestartRequest {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridServicerestartRequestModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridServicerestartRequestModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridServicerestartRequest {
	if m == nil {
		return nil
	}
	to := &grid.GridServicerestartRequest{
		Ref: flex.ExpandStringPointer(m.Ref),
	}
	return to
}

func FlattenGridServicerestartRequest(ctx context.Context, from *grid.GridServicerestartRequest, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridServicerestartRequestAttrTypes)
	}
	m := GridServicerestartRequestModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridServicerestartRequestAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridServicerestartRequestModel) Flatten(ctx context.Context, from *grid.GridServicerestartRequest, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridServicerestartRequestModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Error = flex.FlattenStringPointer(from.Error)
	m.Forced = types.BoolPointerValue(from.Forced)
	m.Group = flex.FlattenStringPointer(from.Group)
	m.LastUpdatedTime = flex.FlattenInt64Pointer(from.LastUpdatedTime)
	m.Member = flex.FlattenStringPointer(from.Member)
	m.Needed = flex.FlattenStringPointer(from.Needed)
	m.Order = flex.FlattenInt64Pointer(from.Order)
	m.Result = flex.FlattenStringPointer(from.Result)
	m.Service = flex.FlattenStringPointer(from.Service)
	m.State = flex.FlattenStringPointer(from.State)
}
