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

type GriddnsdnsseckeyparamsZskAlgorithmsModel struct {
	Algorithm types.String `tfsdk:"algorithm"`
	Size      types.Int64  `tfsdk:"size"`
}

var GriddnsdnsseckeyparamsZskAlgorithmsAttrTypes = map[string]attr.Type{
	"algorithm": types.StringType,
	"size":      types.Int64Type,
}

var GriddnsdnsseckeyparamsZskAlgorithmsResourceSchemaAttributes = map[string]schema.Attribute{
	"algorithm": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The signing key algorithm.",
	},
	"size": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The signing key size, in bits.",
	},
}

func ExpandGriddnsdnsseckeyparamsZskAlgorithms(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GriddnsdnsseckeyparamsZskAlgorithms {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GriddnsdnsseckeyparamsZskAlgorithmsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GriddnsdnsseckeyparamsZskAlgorithmsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GriddnsdnsseckeyparamsZskAlgorithms {
	if m == nil {
		return nil
	}
	to := &grid.GriddnsdnsseckeyparamsZskAlgorithms{
		Algorithm: flex.ExpandStringPointer(m.Algorithm),
		Size:      flex.ExpandInt64Pointer(m.Size),
	}
	return to
}

func FlattenGriddnsdnsseckeyparamsZskAlgorithms(ctx context.Context, from *grid.GriddnsdnsseckeyparamsZskAlgorithms, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GriddnsdnsseckeyparamsZskAlgorithmsAttrTypes)
	}
	m := GriddnsdnsseckeyparamsZskAlgorithmsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GriddnsdnsseckeyparamsZskAlgorithmsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GriddnsdnsseckeyparamsZskAlgorithmsModel) Flatten(ctx context.Context, from *grid.GriddnsdnsseckeyparamsZskAlgorithms, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GriddnsdnsseckeyparamsZskAlgorithmsModel{}
	}
	m.Algorithm = flex.FlattenStringPointer(from.Algorithm)
	m.Size = flex.FlattenInt64Pointer(from.Size)
}
