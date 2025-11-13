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

type GridObjectsChangesTrackingSettingModel struct {
	Enable           types.Bool   `tfsdk:"enable"`
	EnableCompletion types.Int64  `tfsdk:"enable_completion"`
	State            types.String `tfsdk:"state"`
	MaxTimeToTrack   types.Int64  `tfsdk:"max_time_to_track"`
	MaxObjsToTrack   types.Int64  `tfsdk:"max_objs_to_track"`
}

var GridObjectsChangesTrackingSettingAttrTypes = map[string]attr.Type{
	"enable":            types.BoolType,
	"enable_completion": types.Int64Type,
	"state":             types.StringType,
	"max_time_to_track": types.Int64Type,
	"max_objs_to_track": types.Int64Type,
}

var GridObjectsChangesTrackingSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the objects changes tracking feature is enabled or not.",
	},
	"enable_completion": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Determines the percentage of completion for objects changes tracking.",
	},
	"state": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Determines the objects changes tracking enable state.",
	},
	"max_time_to_track": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Maximum time period in seconds to track the deleted objects changes. You can enter a value from 7200 - 604800 seconds.",
	},
	"max_objs_to_track": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Maximum number of deleted objects retained for tracking. You can enter a value from 2000 - 20000.",
	},
}

func ExpandGridObjectsChangesTrackingSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridObjectsChangesTrackingSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridObjectsChangesTrackingSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridObjectsChangesTrackingSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridObjectsChangesTrackingSetting {
	if m == nil {
		return nil
	}
	to := &grid.GridObjectsChangesTrackingSetting{
		Enable:         flex.ExpandBoolPointer(m.Enable),
		MaxTimeToTrack: flex.ExpandInt64Pointer(m.MaxTimeToTrack),
		MaxObjsToTrack: flex.ExpandInt64Pointer(m.MaxObjsToTrack),
	}
	return to
}

func FlattenGridObjectsChangesTrackingSetting(ctx context.Context, from *grid.GridObjectsChangesTrackingSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridObjectsChangesTrackingSettingAttrTypes)
	}
	m := GridObjectsChangesTrackingSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridObjectsChangesTrackingSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridObjectsChangesTrackingSettingModel) Flatten(ctx context.Context, from *grid.GridObjectsChangesTrackingSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridObjectsChangesTrackingSettingModel{}
	}
	m.Enable = types.BoolPointerValue(from.Enable)
	m.EnableCompletion = flex.FlattenInt64Pointer(from.EnableCompletion)
	m.State = flex.FlattenStringPointer(from.State)
	m.MaxTimeToTrack = flex.FlattenInt64Pointer(from.MaxTimeToTrack)
	m.MaxObjsToTrack = flex.FlattenInt64Pointer(from.MaxObjsToTrack)
}
