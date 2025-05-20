package dtc

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dtc"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/flex"
)

type ExtAttrModel struct {
	Value types.String `tfsdk:"value"`
}

var ExtAttrAttrTypes = map[string]attr.Type{
	"value": types.StringType,
}

var ExtAttrResourceSchemaAttributes = map[string]schema.Attribute{
	"value": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The value of the extensible attribute.",
	},
}

func ExpandExtAttr(ctx context.Context, tfMap types.Map, diags *diag.Diagnostics) *map[string]dtc.ExtAttrs {
	if tfMap.IsNull() || tfMap.IsUnknown() {
		return nil
	}

	result := make(map[string]dtc.ExtAttrs)
	for key, val := range tfMap.Elements() {
		objVal, ok := val.(basetypes.ObjectValue)
		if !ok {
			diags.Append(diag.NewErrorDiagnostic(
				"Unexpected object value in extattrs map",
				fmt.Sprintf("Expected object value in extattrs map for key %q", key),
			))
			continue
		}

		var m ExtAttrModel
		diags.Append(objVal.As(ctx, &m, basetypes.ObjectAsOptions{})...)
		if diags.HasError() {
			return nil
		}

		result[key] = *m.Expand(ctx, diags)
	}

	return &result
}

func (m *ExtAttrModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dtc.ExtAttrs {
	if m == nil {
		return nil
	}
	to := &dtc.ExtAttrs{
		Value: flex.ExpandString(m.Value),
	}
	return to
}

func FlattenExtAttr(ctx context.Context, input map[string]dtc.ExtAttrs, diags *diag.Diagnostics) types.Map {

	elementsOg := make(map[string]attr.Value)

	for k, v := range input {
		if &v == nil {
			continue
		}
		m := ExtAttrModel{}
		m.Flatten(ctx, &v, diags)
		t, d := types.ObjectValueFrom(ctx, ExtAttrAttrTypes, m)
		diags.Append(d...)
		elementsOg[k] = t
	}

	elements, d := types.MapValue(types.ObjectType{AttrTypes: ExtAttrAttrTypes}, elementsOg)
	diags.Append(d...)
	return elements
}

func (m *ExtAttrModel) Flatten(ctx context.Context, from *dtc.ExtAttrs, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ExtAttrModel{}
	}
	m.Value = flex.FlattenString(from.Value)
}

func RemoveInheritedExtAttrs(ctx context.Context, planExtAttrs types.Map, respExtAttrs map[string]dtc.ExtAttrs) (*map[string]dtc.ExtAttrs, diag.Diagnostics) {
	var diags diag.Diagnostics
	newRespMap := make(map[string]dtc.ExtAttrs, len(respExtAttrs))

	if planExtAttrs.IsNull() || planExtAttrs.IsUnknown() {
		if v, ok := respExtAttrs["Terraform Internal ID"]; ok {
			newRespMap["Terraform Internal ID"] = v
		}
		return &newRespMap, nil
	}

	planMap := *ExpandExtAttr(ctx, planExtAttrs, &diags)
	if diags.HasError() {
		return nil, diags
	}

	for k, v := range respExtAttrs {
		if k == "Terraform Internal ID" {
			newRespMap[k] = v
			continue
		}

		if respExtAttrs[k].AdditionalProperties["inheritance_source"] != nil {
			if planVal, ok := planMap[k]; ok {
				newRespMap[k] = planVal
			}
			continue
		}
		newRespMap[k] = respExtAttrs[k]
	}
	return &newRespMap, diags
}
