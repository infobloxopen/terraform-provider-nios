package grid

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"
	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type ExtensibleattributedefDefaultValueModel struct {
}

var ExtensibleattributedefDefaultValueAttrTypes = map[string]attr.Type{}

var ExtensibleattributedefDefaultValueResourceSchemaAttributes = map[string]schema.Attribute{}

func ExpandExtensibleattributedefDefaultValue(ctx context.Context, s types.String, diags *diag.Diagnostics) *grid.ExtensibleattributedefDefaultValue {

	if s.IsNull() || s.IsUnknown() {
		return nil
	}

	stringPtr := flex.ExpandStringPointer(s)
	if stringPtr == nil {
		return nil
	}

	return &grid.ExtensibleattributedefDefaultValue{
		// Assuming the API expects a string value for DefaultValue
		// Adjust the type conversion as necessary based on actual API requirements
		String: stringPtr,
	}
}

func (m *ExtensibleattributedefDefaultValueModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.ExtensibleattributedefDefaultValue {
	if m == nil {
		return nil
	}
	to := &grid.ExtensibleattributedefDefaultValue{}
	return to
}

func FlattenExtensibleattributedefDefaultValue(ctx context.Context, from *grid.ExtensibleattributedefDefaultValue, diags *diag.Diagnostics) types.String {
	if from == nil {
		return types.StringNull()
	}

	// Check which field is set and convert back to string
	if from.String != nil {
		return flex.FlattenStringPointer(from.String)
	} else if from.Int32 != nil {
		value := strconv.FormatInt(int64(*from.Int32), 10)
		return flex.FlattenStringPointer(&value)
	}

	return types.StringNull()
}

func (m *ExtensibleattributedefDefaultValueModel) Flatten(ctx context.Context, from *grid.ExtensibleattributedefDefaultValue, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ExtensibleattributedefDefaultValueModel{}
	}
}
