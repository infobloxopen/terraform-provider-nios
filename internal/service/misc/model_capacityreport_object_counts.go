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

type CapacityreportObjectCountsModel struct {
	TypeName types.String `tfsdk:"type_name"`
	Count    types.Int64  `tfsdk:"count"`
}

var CapacityreportObjectCountsAttrTypes = map[string]attr.Type{
	"type_name": types.StringType,
	"count":     types.Int64Type,
}

var CapacityreportObjectCountsResourceSchemaAttributes = map[string]schema.Attribute{
	"type_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Object type name.",
	},
	"count": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Number of object type instances created.",
	},
}

func ExpandCapacityreportObjectCounts(ctx context.Context, o types.Object, diags *diag.Diagnostics) *misc.CapacityreportObjectCounts {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m CapacityreportObjectCountsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *CapacityreportObjectCountsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *misc.CapacityreportObjectCounts {
	if m == nil {
		return nil
	}
	to := &misc.CapacityreportObjectCounts{}
	return to
}

func FlattenCapacityreportObjectCounts(ctx context.Context, from *misc.CapacityreportObjectCounts, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(CapacityreportObjectCountsAttrTypes)
	}
	m := CapacityreportObjectCountsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, CapacityreportObjectCountsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *CapacityreportObjectCountsModel) Flatten(ctx context.Context, from *misc.CapacityreportObjectCounts, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = CapacityreportObjectCountsModel{}
	}
	m.TypeName = flex.FlattenStringPointer(from.TypeName)
	m.Count = flex.FlattenInt64Pointer(from.Count)
}
