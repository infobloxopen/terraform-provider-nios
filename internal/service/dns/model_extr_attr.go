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

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dns"
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

func FlattenExtAttr(ctx context.Context, extattrs map[string]dns.ExtAttrs, diags *diag.Diagnostics) types.Map {

	result := make(map[string]attr.Value)

	for key, extAttr := range extattrs {
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

	// JSON list with either single or double quotes
	if strings.HasPrefix(valStr, "[") && strings.HasSuffix(valStr, "]") {
		var listVal []interface{}

		// Try standard JSON first (with double quotes)
		err := json.Unmarshal([]byte(valStr), &listVal)

		// If that fails and we have single quotes, replace them with double quotes
		if err != nil && strings.Contains(valStr, "'") {
			processedStr := strings.ReplaceAll(valStr, "'", "\"")
			err = json.Unmarshal([]byte(processedStr), &listVal)
		}

		// If we successfully parsed it as a list
		if err == nil {
			return listVal
		}
	}
	// Integer values
	if intVal, err := strconv.ParseInt(valStr, 10, 64); err == nil {
		return intVal
	}
	return valStr
}

func FlattenUserExtAttrs(ctx context.Context, extattrs *map[string]dns.ExtAttrs, diags *diag.Diagnostics) types.Map {
	if extattrs == nil {
		return types.MapNull(types.StringType)
	}

	sys_extattrs := map[string]bool{
		"Terraform Internal ID": true,
	}

	userExtAttrs := make(map[string]dns.ExtAttrs)
	for key, val := range *extattrs {
		if !sys_extattrs[key] {
			userExtAttrs[key] = val
		}
	}

	if len(userExtAttrs) > 0 {
		return FlattenExtAttr(ctx, userExtAttrs, diags)
	}

	return types.MapNull(types.StringType)
}
