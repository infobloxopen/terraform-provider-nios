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

type ScavengingtaskModel struct {
	Ref                types.String `tfsdk:"ref"`
	Uuid               types.String `tfsdk:"uuid"`
	Action             types.String `tfsdk:"action"`
	AssociatedObject   types.String `tfsdk:"associated_object"`
	EndTime            types.Int64  `tfsdk:"end_time"`
	ProcessedRecords   types.Int64  `tfsdk:"processed_records"`
	ReclaimableRecords types.Int64  `tfsdk:"reclaimable_records"`
	ReclaimedRecords   types.Int64  `tfsdk:"reclaimed_records"`
	StartTime          types.Int64  `tfsdk:"start_time"`
	Status             types.String `tfsdk:"status"`
}

var ScavengingtaskAttrTypes = map[string]attr.Type{
	"ref":                 types.StringType,
	"uuid":                types.StringType,
	"action":              types.StringType,
	"associated_object":   types.StringType,
	"end_time":            types.Int64Type,
	"processed_records":   types.Int64Type,
	"reclaimable_records": types.Int64Type,
	"reclaimed_records":   types.Int64Type,
	"start_time":          types.Int64Type,
	"status":              types.StringType,
}

var ScavengingtaskResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The UUID of the object.",
	},
	"action": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The scavenging action.",
	},
	"associated_object": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object associated with the scavenging task.",
	},
	"end_time": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The scavenging process end time.",
	},
	"processed_records": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of processed during scavenging resource records.",
	},
	"reclaimable_records": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of resource records that are allowed to be reclaimed during the scavenging process.",
	},
	"reclaimed_records": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of reclaimed during the scavenging process resource records.",
	},
	"start_time": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The scavenging process start time.",
	},
	"status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The scavenging process status. This is a read-only attribute.",
	},
}

func ExpandScavengingtask(ctx context.Context, o types.Object, diags *diag.Diagnostics) *misc.Scavengingtask {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ScavengingtaskModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ScavengingtaskModel) Expand(ctx context.Context, diags *diag.Diagnostics) *misc.Scavengingtask {
	if m == nil {
		return nil
	}
	to := &misc.Scavengingtask{
		Ref:  flex.ExpandStringPointer(m.Ref),
		Uuid: flex.ExpandStringPointer(m.Uuid),
	}
	return to
}

func FlattenScavengingtask(ctx context.Context, from *misc.Scavengingtask, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ScavengingtaskAttrTypes)
	}
	m := ScavengingtaskModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ScavengingtaskAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ScavengingtaskModel) Flatten(ctx context.Context, from *misc.Scavengingtask, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ScavengingtaskModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Action = flex.FlattenStringPointer(from.Action)
	m.AssociatedObject = flex.FlattenStringPointer(from.AssociatedObject)
	m.EndTime = flex.FlattenInt64Pointer(from.EndTime)
	m.ProcessedRecords = flex.FlattenInt64Pointer(from.ProcessedRecords)
	m.ReclaimableRecords = flex.FlattenInt64Pointer(from.ReclaimableRecords)
	m.ReclaimedRecords = flex.FlattenInt64Pointer(from.ReclaimedRecords)
	m.StartTime = flex.FlattenInt64Pointer(from.StartTime)
	m.Status = flex.FlattenStringPointer(from.Status)
}
