package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"
)

type ExtensibleattributedefDefaultValueModel struct {
}

var ExtensibleattributedefDefaultValueAttrTypes = map[string]attr.Type{}

var ExtensibleattributedefDefaultValueResourceSchemaAttributes = map[string]schema.Attribute{}

func ExpandExtensibleattributedefDefaultValue(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.ExtensibleattributedefDefaultValue {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ExtensibleattributedefDefaultValueModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ExtensibleattributedefDefaultValueModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.ExtensibleattributedefDefaultValue {
	if m == nil {
		return nil
	}
	to := &grid.ExtensibleattributedefDefaultValue{}
	return to
}

func FlattenExtensibleattributedefDefaultValue(ctx context.Context, from *grid.ExtensibleattributedefDefaultValue, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ExtensibleattributedefDefaultValueAttrTypes)
	}
	m := ExtensibleattributedefDefaultValueModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ExtensibleattributedefDefaultValueAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ExtensibleattributedefDefaultValueModel) Flatten(ctx context.Context, from *grid.ExtensibleattributedefDefaultValue, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ExtensibleattributedefDefaultValueModel{}
	}
}
