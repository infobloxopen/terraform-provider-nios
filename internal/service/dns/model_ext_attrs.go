package dns

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

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

func FlattenExtAttrsAll(ctx context.Context, extattrs *map[string]dns.ExtAttrs, diags *diag.Diagnostics) types.Map {
	result := make(map[string]attr.Value)

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
				result[key] = types.StringValue(string(jsonBytes))
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

// FlattenExtAttrs flattens the updated external attributes based on the plan and existing attributes.
func FlattenExtAttrs(ctx context.Context, planExtAttrs types.Map, extattrs *map[string]dns.ExtAttrs, diags *diag.Diagnostics) types.Map {
	if extattrs == nil || len(*extattrs) == 0 {
		return types.MapNull(types.StringType)
	}
	planExtAttrsMapPtr := ExpandExtAttr(ctx, planExtAttrs, diags)
	if planExtAttrsMapPtr == nil || len(*planExtAttrsMapPtr) == 0 {
		return types.MapNull(types.StringType)
	}

	userExtAttrs := make(map[string]dns.ExtAttrs, len(*extattrs))
	for key, extAttrVal := range *extattrs {
		if key == "Terraform Internal ID" {
			continue // Skip the internal ID as it is not part of user-defined attributes
		}
		userExtAttrs[key] = extAttrVal
	}

	if len(userExtAttrs) == 0 {
		return types.MapNull(types.StringType)
	}
	return FlattenExtAttrsAll(ctx, &userExtAttrs, diags)
}

func RemoveInheritedExtAttrs(ctx context.Context, planExtAttrs types.Map, respExtAttrs map[string]dns.ExtAttrs) (*map[string]dns.ExtAttrs, diag.Diagnostics) {
	var diags diag.Diagnostics
	newRespMap := make(map[string]dns.ExtAttrs, len(respExtAttrs))

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
