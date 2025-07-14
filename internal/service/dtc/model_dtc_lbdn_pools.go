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

type DtcLbdnPoolsModel struct {
	Pool  types.String `tfsdk:"pool"`
	Ratio types.Int64  `tfsdk:"ratio"`
}

var DtcLbdnPoolsAttrTypes = map[string]attr.Type{
	"pool":  types.StringType,
	"ratio": types.Int64Type,
}

var DtcLbdnPoolsResourceSchemaAttributes = map[string]schema.Attribute{
	"pool": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The pool to link with.",
	},
	"ratio": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The weight of pool.",
	},
}

func ExpandDtcLbdnPools(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dtc.DtcLbdnPools {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DtcLbdnPoolsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DtcLbdnPoolsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dtc.DtcLbdnPools {
	if m == nil {
		return nil
	}
	to := &dtc.DtcLbdnPools{
		Pool:  flex.ExpandStringPointer(m.Pool),
		Ratio: flex.ExpandInt64Pointer(m.Ratio),
	}
	return to
}

func FlattenDtcLbdnPools(ctx context.Context, from *dtc.DtcLbdnPools, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DtcLbdnPoolsAttrTypes)
	}
	m := DtcLbdnPoolsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DtcLbdnPoolsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DtcLbdnPoolsModel) Flatten(ctx context.Context, from *dtc.DtcLbdnPools, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DtcLbdnPoolsModel{}
	}
	m.Pool = flex.FlattenStringPointer(from.Pool)
	m.Ratio = flex.FlattenInt64Pointer(from.Ratio)
}
