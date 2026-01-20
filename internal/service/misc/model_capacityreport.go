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

type CapacityreportModel struct {
	Ref          types.String `tfsdk:"ref"`
	Uuid         types.String `tfsdk:"uuid"`
	HardwareType types.String `tfsdk:"hardware_type"`
	MaxCapacity  types.Int64  `tfsdk:"max_capacity"`
	Name         types.String `tfsdk:"name"`
	ObjectCounts types.List   `tfsdk:"object_counts"`
	PercentUsed  types.Int64  `tfsdk:"percent_used"`
	Role         types.String `tfsdk:"role"`
	TotalObjects types.Int64  `tfsdk:"total_objects"`
}

var CapacityreportAttrTypes = map[string]attr.Type{
	"ref":           types.StringType,
	"uuid":          types.StringType,
	"hardware_type": types.StringType,
	"max_capacity":  types.Int64Type,
	"name":          types.StringType,
	"object_counts": types.ListType{ElemType: types.ObjectType{AttrTypes: CapacityreportObjectCountsAttrTypes}},
	"percent_used":  types.Int64Type,
	"role":          types.StringType,
	"total_objects": types.Int64Type,
}

var CapacityreportResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The UUID of the object.",
	},
	"hardware_type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Hardware type of a Grid member.",
	},
	"max_capacity": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The maximum amount of capacity available for the Grid member.",
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The Grid member name.",
	},
	"object_counts": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: CapacityreportObjectCountsResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Computed:            true,
		MarkdownDescription: "A list of instance counts for object types created on the Grid member.",
	},
	"percent_used": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The percentage of the capacity in use by the Grid member.",
	},
	"role": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The Grid member role.",
	},
	"total_objects": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The total number of objects created by the Grid member.",
	},
}

func ExpandCapacityreport(ctx context.Context, o types.Object, diags *diag.Diagnostics) *misc.Capacityreport {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m CapacityreportModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *CapacityreportModel) Expand(ctx context.Context, diags *diag.Diagnostics) *misc.Capacityreport {
	if m == nil {
		return nil
	}
	to := &misc.Capacityreport{
		Ref:  flex.ExpandStringPointer(m.Ref),
		Uuid: flex.ExpandStringPointer(m.Uuid),
	}
	return to
}

func FlattenCapacityreport(ctx context.Context, from *misc.Capacityreport, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(CapacityreportAttrTypes)
	}
	m := CapacityreportModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, CapacityreportAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *CapacityreportModel) Flatten(ctx context.Context, from *misc.Capacityreport, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = CapacityreportModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.HardwareType = flex.FlattenStringPointer(from.HardwareType)
	m.MaxCapacity = flex.FlattenInt64Pointer(from.MaxCapacity)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.ObjectCounts = flex.FlattenFrameworkListNestedBlock(ctx, from.ObjectCounts, CapacityreportObjectCountsAttrTypes, diags, FlattenCapacityreportObjectCounts)
	m.PercentUsed = flex.FlattenInt64Pointer(from.PercentUsed)
	m.Role = flex.FlattenStringPointer(from.Role)
	m.TotalObjects = flex.FlattenInt64Pointer(from.TotalObjects)
}
