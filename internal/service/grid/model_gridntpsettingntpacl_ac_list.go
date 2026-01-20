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

type GridntpsettingntpaclAcListModel struct {
	AddressAc types.Object `tfsdk:"address_ac"`
	Service   types.String `tfsdk:"service"`
}

var GridntpsettingntpaclAcListAttrTypes = map[string]attr.Type{
	"address_ac": types.ObjectType{AttrTypes: GridntpsettingntpaclaclistAddressAcAttrTypes},
	"service":    types.StringType,
}

var GridntpsettingntpaclAcListResourceSchemaAttributes = map[string]schema.Attribute{
	"address_ac": schema.SingleNestedAttribute{
		Attributes: GridntpsettingntpaclaclistAddressAcResourceSchemaAttributes,
		Optional:   true,
	},
	"service": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The type of service with access control.",
	},
}

func ExpandGridntpsettingntpaclAcList(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridntpsettingntpaclAcList {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridntpsettingntpaclAcListModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridntpsettingntpaclAcListModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridntpsettingntpaclAcList {
	if m == nil {
		return nil
	}
	to := &grid.GridntpsettingntpaclAcList{
		AddressAc: ExpandGridntpsettingntpaclaclistAddressAc(ctx, m.AddressAc, diags),
		Service:   flex.ExpandStringPointer(m.Service),
	}
	return to
}

func FlattenGridntpsettingntpaclAcList(ctx context.Context, from *grid.GridntpsettingntpaclAcList, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridntpsettingntpaclAcListAttrTypes)
	}
	m := GridntpsettingntpaclAcListModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridntpsettingntpaclAcListAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridntpsettingntpaclAcListModel) Flatten(ctx context.Context, from *grid.GridntpsettingntpaclAcList, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridntpsettingntpaclAcListModel{}
	}
	m.AddressAc = FlattenGridntpsettingntpaclaclistAddressAc(ctx, from.AddressAc, diags)
	m.Service = flex.FlattenStringPointer(from.Service)
}
