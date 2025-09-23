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

type ExtensibleattributedefListValuesModel struct {
	Value types.String `tfsdk:"value"`
}

var ExtensibleattributedefListValuesAttrTypes = map[string]attr.Type{
	"value": types.StringType,
}

var ExtensibleattributedefListValuesResourceSchemaAttributes = map[string]schema.Attribute{
	"value": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Enum value",
	},
}

func ExpandExtensibleattributedefListValues(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.ExtensibleattributedefListValues {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ExtensibleattributedefListValuesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ExtensibleattributedefListValuesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.ExtensibleattributedefListValues {
	if m == nil {
		return nil
	}
	to := &grid.ExtensibleattributedefListValues{
		Value: flex.ExpandStringPointer(m.Value),
	}
	return to
}

func FlattenExtensibleattributedefListValues(ctx context.Context, from *grid.ExtensibleattributedefListValues, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ExtensibleattributedefListValuesAttrTypes)
	}
	m := ExtensibleattributedefListValuesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ExtensibleattributedefListValuesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ExtensibleattributedefListValuesModel) Flatten(ctx context.Context, from *grid.ExtensibleattributedefListValues, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ExtensibleattributedefListValuesModel{}
	}
	m.Value = flex.FlattenStringPointer(from.Value)
}
