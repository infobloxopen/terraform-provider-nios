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

type GriddnsdnsseckeyparamsKskAlgorithmsModel struct {
	Algorithm types.String `tfsdk:"algorithm"`
	Size      types.Int64  `tfsdk:"size"`
}

var GriddnsdnsseckeyparamsKskAlgorithmsAttrTypes = map[string]attr.Type{
	"algorithm": types.StringType,
	"size":      types.Int64Type,
}

var GriddnsdnsseckeyparamsKskAlgorithmsResourceSchemaAttributes = map[string]schema.Attribute{
	"algorithm": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The signing key algorithm.",
	},
	"size": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The signing key size, in bits.",
	},
}

func ExpandGriddnsdnsseckeyparamsKskAlgorithms(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GriddnsdnsseckeyparamsKskAlgorithms {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GriddnsdnsseckeyparamsKskAlgorithmsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GriddnsdnsseckeyparamsKskAlgorithmsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GriddnsdnsseckeyparamsKskAlgorithms {
	if m == nil {
		return nil
	}
	to := &grid.GriddnsdnsseckeyparamsKskAlgorithms{
		Algorithm: flex.ExpandStringPointer(m.Algorithm),
		Size:      flex.ExpandInt64Pointer(m.Size),
	}
	return to
}

func FlattenGriddnsdnsseckeyparamsKskAlgorithms(ctx context.Context, from *grid.GriddnsdnsseckeyparamsKskAlgorithms, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GriddnsdnsseckeyparamsKskAlgorithmsAttrTypes)
	}
	m := GriddnsdnsseckeyparamsKskAlgorithmsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GriddnsdnsseckeyparamsKskAlgorithmsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GriddnsdnsseckeyparamsKskAlgorithmsModel) Flatten(ctx context.Context, from *grid.GriddnsdnsseckeyparamsKskAlgorithms, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GriddnsdnsseckeyparamsKskAlgorithmsModel{}
	}
	m.Algorithm = flex.FlattenStringPointer(from.Algorithm)
	m.Size = flex.FlattenInt64Pointer(from.Size)
}
