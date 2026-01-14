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

type GridServicerestartStatusModel struct {
	Ref            types.String `tfsdk:"ref"`
	Uuid           types.String `tfsdk:"uuid"`
	Failures       types.Int64  `tfsdk:"failures"`
	Finished       types.Int64  `tfsdk:"finished"`
	Grouped        types.String `tfsdk:"grouped"`
	NeededRestart  types.Int64  `tfsdk:"needed_restart"`
	NoRestart      types.Int64  `tfsdk:"no_restart"`
	Parent         types.String `tfsdk:"parent"`
	Pending        types.Int64  `tfsdk:"pending"`
	PendingRestart types.Int64  `tfsdk:"pending_restart"`
	Processing     types.Int64  `tfsdk:"processing"`
	Restarting     types.Int64  `tfsdk:"restarting"`
	Success        types.Int64  `tfsdk:"success"`
	Timeouts       types.Int64  `tfsdk:"timeouts"`
}

var GridServicerestartStatusAttrTypes = map[string]attr.Type{
	"ref":             types.StringType,
	"uuid":            types.StringType,
	"failures":        types.Int64Type,
	"finished":        types.Int64Type,
	"grouped":         types.StringType,
	"needed_restart":  types.Int64Type,
	"no_restart":      types.Int64Type,
	"parent":          types.StringType,
	"pending":         types.Int64Type,
	"pending_restart": types.Int64Type,
	"processing":      types.Int64Type,
	"restarting":      types.Int64Type,
	"success":         types.Int64Type,
	"timeouts":        types.Int64Type,
}

var GridServicerestartStatusResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The uuid to the object.",
	},
	"failures": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of failed requests.",
	},
	"finished": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of finished requests.",
	},
	"grouped": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The type of grouping.",
	},
	"needed_restart": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of created yet unprocessed requests for restart.",
	},
	"no_restart": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of requests that did not require a restart.",
	},
	"parent": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "A reference to the grid or grid:servicerestart:group object associated with the request.",
	},
	"pending": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of requests that are pending a restart.",
	},
	"pending_restart": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of forced or needed requests pending for restart.",
	},
	"processing": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of not forced and not needed requests pending for restart.",
	},
	"restarting": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of service restarts for corresponding members.",
	},
	"success": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of requests associated with successful restarts.",
	},
	"timeouts": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of timeout requests.",
	},
}

func ExpandGridServicerestartStatus(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridServicerestartStatus {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridServicerestartStatusModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridServicerestartStatusModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridServicerestartStatus {
	if m == nil {
		return nil
	}
	to := &grid.GridServicerestartStatus{}
	return to
}

func FlattenGridServicerestartStatus(ctx context.Context, from *grid.GridServicerestartStatus, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridServicerestartStatusAttrTypes)
	}
	m := GridServicerestartStatusModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridServicerestartStatusAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridServicerestartStatusModel) Flatten(ctx context.Context, from *grid.GridServicerestartStatus, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridServicerestartStatusModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Failures = flex.FlattenInt64Pointer(from.Failures)
	m.Finished = flex.FlattenInt64Pointer(from.Finished)
	m.Grouped = flex.FlattenStringPointer(from.Grouped)
	m.NeededRestart = flex.FlattenInt64Pointer(from.NeededRestart)
	m.NoRestart = flex.FlattenInt64Pointer(from.NoRestart)
	m.Parent = flex.FlattenStringPointer(from.Parent)
	m.Pending = flex.FlattenInt64Pointer(from.Pending)
	m.PendingRestart = flex.FlattenInt64Pointer(from.PendingRestart)
	m.Processing = flex.FlattenInt64Pointer(from.Processing)
	m.Restarting = flex.FlattenInt64Pointer(from.Restarting)
	m.Success = flex.FlattenInt64Pointer(from.Success)
	m.Timeouts = flex.FlattenInt64Pointer(from.Timeouts)
}
