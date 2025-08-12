package security

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-uuid"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

const terraformInternalIDEA = "Terraform Internal ID"

func ExpandExtAttrs(ctx context.Context, extattrs types.Map, diags *diag.Diagnostics) *map[string]security.ExtAttrs {
	if extattrs.IsNull() || extattrs.IsUnknown() {
		return nil
	}
	var extAttrsMap map[string]string
	diags.Append(extattrs.ElementsAs(ctx, &extAttrsMap, false)...)
	if diags.HasError() {
		return nil
	}

	result := make(map[string]security.ExtAttrs)

	for key, valStr := range extAttrsMap {
		parsedValue := utils.ParseInterfaceValue(valStr)
		result[key] = security.ExtAttrs{Value: parsedValue}
	}
	return &result
}

func FlattenExtAttrs(ctx context.Context, planExtAttrs types.Map, extattrs *map[string]security.ExtAttrs, diags *diag.Diagnostics) types.Map {
	result := make(map[string]attr.Value)
	planExtAttrsMap := planExtAttrs.Elements()
	if extattrs == nil || len(*extattrs) == 0 {
		return types.MapNull(types.StringType)
	}

	for key, extAttr := range *extattrs {
		if extAttr.Value == nil {
			continue
		}

		// Convert value to string based on its type
		switch v := extAttr.Value.(type) {
		case []interface{}:
			// Convert list to JSON string
			jsonBytes, err := json.Marshal(v)
			if err != nil {
				diags.AddError(
					"Error converting list to JSON",
					fmt.Sprintf("Could not convert list value for key %s: %s", key, err),
				)
				result[key] = types.StringValue(fmt.Sprintf("%v", v))
			} else {
				value := string(jsonBytes)
				if _, ok := planExtAttrsMap[key]; ok {
					if strings.Contains(planExtAttrsMap[key].String(), "'") {
						value = strings.ReplaceAll(value, "\"", "'")
					}
				}
				result[key] = types.StringValue(value)
			}
		default:
			// Convert primitive values to string
			result[key] = types.StringValue(fmt.Sprintf("%v", v))
		}
	}

	mapVal, mapDiags := types.MapValue(types.StringType, result)
	diags.Append(mapDiags...)
	return mapVal
}

func RemoveInheritedExtAttrs(ctx context.Context, planExtAttrs types.Map, respExtAttrs map[string]security.ExtAttrs) (*map[string]security.ExtAttrs, types.Map, diag.Diagnostics) {
	var diags diag.Diagnostics
	extAttrsRespMap := make(map[string]security.ExtAttrs, len(planExtAttrs.Elements()))
	extAttrsAllRespMap := make(map[string]security.ExtAttrs)
	var extAttrAll types.Map

	if planExtAttrs.IsNull() || planExtAttrs.IsUnknown() {
		extAttrAll = FlattenExtAttrs(ctx, planExtAttrs, &respExtAttrs, &diags)
		return nil, extAttrAll, nil
	}

	planMap := *ExpandExtAttrs(ctx, planExtAttrs, &diags)
	if diags.HasError() {
		return nil, extAttrAll, diags
	}

	for k, v := range respExtAttrs {
		if k == terraformInternalIDEA {
			extAttrsAllRespMap[k] = v
			continue
		}

		// If the EA is inherited , if the state is override , add it to the ExtAttrs.
		// If the EA is inherited and state is inherited , add it ExtAttrsAll
		if respExtAttrs[k].AdditionalProperties["inheritance_source"] != nil {
			if planVal, ok := planMap[k]; ok {
				extAttrsRespMap[k] = planVal
			} else {
				extAttrsAllRespMap[k] = respExtAttrs[k]
			}
			continue
		}
		extAttrsRespMap[k] = v
	}
	extAttrAll = FlattenExtAttrs(ctx, planExtAttrs, &extAttrsAllRespMap, &diags)
	return &extAttrsRespMap, extAttrAll, diags
}

func AddInheritedExtAttrs(ctx context.Context, planExtAttrs types.Map, stateExtAttrs types.Map) (types.Map, diag.Diagnostics) {
	var diags diag.Diagnostics
	stateExtAttrsMap := stateExtAttrs.Elements()
	if len(stateExtAttrsMap) == 0 {
		return planExtAttrs, diags
	}
	planExtAttrsMap := planExtAttrs.Elements()

	for k, v := range stateExtAttrsMap {
		// if the key is not in planExtAttrsMap , we add it
		if _, ok := planExtAttrsMap[k]; !ok {
			planExtAttrsMap[k] = v
		}
	}

	// Convert the updated map back to types.Map
	newRespMap, diags := types.MapValue(types.StringType, planExtAttrsMap)
	if diags.HasError() {
		return planExtAttrs, diags
	}

	return newRespMap, diags
}

func AddInternalIDToExtAttrs(ctx context.Context, extAttrs types.Map, diags diag.Diagnostics) (types.Map, diag.Diagnostics) {

	internalId, err := uuid.GenerateUUID()
	if err != nil {
		diags.AddError("Error generating UUID", fmt.Sprintf("Unable to generate internal ID for Extensible Attributes: %s", err))
		return extAttrs, diags
	}

	extAttrsMap := extAttrs.Elements()
	extAttrsMap[terraformInternalIDEA] = types.StringValue(internalId)

	extAttrs, diags = types.MapValue(types.StringType, extAttrsMap)
	if diags.HasError() {
		return extAttrs, diags
	}

	return extAttrs, nil
}
