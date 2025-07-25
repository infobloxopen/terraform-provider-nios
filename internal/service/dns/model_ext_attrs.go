package dns

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/go-uuid"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
)

func ExpandExtAttr(ctx context.Context, extattrs types.Map, diags *diag.Diagnostics) *map[string]dns.ExtAttrs {
	if extattrs.IsNull() || extattrs.IsUnknown() {
		return nil
	}
	var extAttrsMap map[string]string
	diags.Append(extattrs.ElementsAs(ctx, &extAttrsMap, false)...)
	if diags.HasError() {
		return nil
	}

	result := make(map[string]dns.ExtAttrs)

	for key, valStr := range extAttrsMap {
		parsedValue := parseExtAttrValue(valStr)
		result[key] = dns.ExtAttrs{Value: parsedValue}
	}
	return &result
}

func FlattenExtAttrs(ctx context.Context, planExtAttrs types.Map, extattrs *map[string]dns.ExtAttrs, diags *diag.Diagnostics) types.Map {
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

func RemoveInheritedExtAttrs(ctx context.Context, planExtAttrs types.Map, respExtAttrs map[string]dns.ExtAttrs) (*map[string]dns.ExtAttrs, types.Map, diag.Diagnostics) {
	var diags diag.Diagnostics
	extAttrsRespMap := make(map[string]dns.ExtAttrs, len(planExtAttrs.Elements()))
	extAttrsAllRespMap := make(map[string]dns.ExtAttrs)
	var extAttrAll types.Map

	if planExtAttrs.IsNull() || planExtAttrs.IsUnknown() {
		if v, ok := respExtAttrs["Terraform Internal ID"]; ok {
			extAttrsAllRespMap["Terraform Internal ID"] = v
		}
		extAttrAll = FlattenExtAttrs(ctx, planExtAttrs, &extAttrsAllRespMap, &diags)
		return nil, extAttrAll, nil
	}

	planMap := *ExpandExtAttr(ctx, planExtAttrs, &diags)
	if diags.HasError() {
		return nil, extAttrAll, diags
	}

	for k, v := range respExtAttrs {
		if k == "Terraform Internal ID" {
			extAttrsAllRespMap[k] = v
			continue
		}

		// If the EA is inherited , if the state is override , add it to the ExtAttrs.
		// If the EA is inherited and state is inherited , add it ExtAttrsAll
		if respExtAttrs[k].AdditionalProperties["inheritance_source"] != nil {
			if planVal, ok := planMap[k]; ok {
				extAttrsRespMap[k] = planVal
			}
			extAttrsAllRespMap[k] = respExtAttrs[k]
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
	if stateExtAttrsMap == nil || len(stateExtAttrsMap) == 0 {
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

func addInternalIDToExtAttrs(ctx context.Context, extAttrs types.Map, diags diag.Diagnostics) (types.Map, diag.Diagnostics) {

	internalId, err := uuid.GenerateUUID()
	if err != nil {
		diags.AddError("Error generating UUID", fmt.Sprintf("Unable to generate internal ID for Extensible Attributes: %s", err))
		return extAttrs, diags
	}

	extAttrsMap := extAttrs.Elements()
	extAttrsMap["Terraform Internal ID"] = types.StringValue(internalId)

	extAttrs, diags = types.MapValue(types.StringType, extAttrsMap)
	if diags.HasError() {
		return extAttrs, diags
	}

	return extAttrs, nil
}

func parseExtAttrValue(valStr string) interface{} {
	// Check if the value appears to be a JSON array (enclosed in square brackets)
	if strings.HasPrefix(valStr, "[") && strings.HasSuffix(valStr, "]") {
		var listVal []interface{}

		// Parse as standard JSON with double quotes
		err := json.Unmarshal([]byte(valStr), &listVal)

		// If that fails and we have single quotes, replace them with double quotes
		if err != nil && strings.Contains(valStr, "'") {
			processedStr := strings.ReplaceAll(valStr, "'", "\"")
			err = json.Unmarshal([]byte(processedStr), &listVal)
		}

		// If either parsing attempt succeeded, return the list value
		if err == nil {
			return listVal
		}
	}

	// Try to parse the value as an integer
	if intVal, err := strconv.ParseInt(valStr, 10, 64); err == nil {
		return intVal
	}
	return valStr
}
