package utils

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const (
	ReadPageSizeLimit   int32  = 1000
	NaiveDatetimeLayout string = "2006-01-02T15:04:05"
)

// Ptr is a helper routine that returns a pointer to given value.
func Ptr[T any](t T) *T {
	return &t
}

// DataSourceAttributeMap converts a map of resource schema attributes to data source schema attributes
func DataSourceAttributeMap(r map[string]resourceschema.Attribute, diags *diag.Diagnostics) map[string]datasourceschema.Attribute {
	d := map[string]datasourceschema.Attribute{}
	for k, v := range r {
		d[k] = DataSourceAttribute(k, v, diags)
	}
	return d
}

// DataSourceNestedAttributeObject converts a resource schema nested attribute object to data source schema nested attribute object
func DataSourceNestedAttributeObject(r resourceschema.NestedAttributeObject, diags *diag.Diagnostics) datasourceschema.NestedAttributeObject {
	return datasourceschema.NestedAttributeObject{
		Attributes: DataSourceAttributeMap(r.Attributes, diags),
		CustomType: r.CustomType,
		Validators: r.Validators,
	}
}

// DataSourceAttribute converts a resource schema attribute to data source schema attribute
func DataSourceAttribute(name string, val resourceschema.Attribute, diags *diag.Diagnostics) datasourceschema.Attribute {
	switch a := val.(type) {
	case resourceschema.BoolAttribute:
		return datasourceschema.BoolAttribute{
			CustomType:          a.CustomType,
			Required:            a.Required,
			Optional:            a.Optional,
			Computed:            a.Computed,
			Sensitive:           a.Sensitive,
			Description:         a.Description,
			MarkdownDescription: a.MarkdownDescription,
			DeprecationMessage:  a.DeprecationMessage,
			Validators:          a.Validators,
		}
	case resourceschema.StringAttribute:
		return datasourceschema.StringAttribute{
			CustomType:          a.CustomType,
			Required:            a.Required,
			Optional:            a.Optional,
			Computed:            a.Computed,
			Sensitive:           a.Sensitive,
			Description:         a.Description,
			MarkdownDescription: a.MarkdownDescription,
			DeprecationMessage:  a.DeprecationMessage,
			Validators:          a.Validators,
		}
	case resourceschema.Int32Attribute:
		return datasourceschema.Int32Attribute{
			CustomType:          a.CustomType,
			Required:            a.Required,
			Optional:            a.Optional,
			Computed:            a.Computed,
			Sensitive:           a.Sensitive,
			Description:         a.Description,
			MarkdownDescription: a.MarkdownDescription,
			DeprecationMessage:  a.DeprecationMessage,
			Validators:          a.Validators,
		}
	case resourceschema.Int64Attribute:
		return datasourceschema.Int64Attribute{
			CustomType:          a.CustomType,
			Required:            a.Required,
			Optional:            a.Optional,
			Computed:            a.Computed,
			Sensitive:           a.Sensitive,
			Description:         a.Description,
			MarkdownDescription: a.MarkdownDescription,
			DeprecationMessage:  a.DeprecationMessage,
			Validators:          a.Validators,
		}
	case resourceschema.Float32Attribute:
		return datasourceschema.Float32Attribute{
			CustomType:          a.CustomType,
			Required:            a.Required,
			Optional:            a.Optional,
			Computed:            a.Computed,
			Sensitive:           a.Sensitive,
			Description:         a.Description,
			MarkdownDescription: a.MarkdownDescription,
			DeprecationMessage:  a.DeprecationMessage,
			Validators:          a.Validators,
		}
	case resourceschema.Float64Attribute:
		return datasourceschema.Float64Attribute{
			CustomType:          a.CustomType,
			Required:            a.Required,
			Optional:            a.Optional,
			Computed:            a.Computed,
			Sensitive:           a.Sensitive,
			Description:         a.Description,
			MarkdownDescription: a.MarkdownDescription,
			DeprecationMessage:  a.DeprecationMessage,
			Validators:          a.Validators,
		}
	case resourceschema.NumberAttribute:
		return datasourceschema.NumberAttribute{
			CustomType:          a.CustomType,
			Required:            a.Required,
			Optional:            a.Optional,
			Computed:            a.Computed,
			Sensitive:           a.Sensitive,
			Description:         a.Description,
			MarkdownDescription: a.MarkdownDescription,
			DeprecationMessage:  a.DeprecationMessage,
			Validators:          a.Validators,
		}
	case resourceschema.ObjectAttribute:
		return datasourceschema.ObjectAttribute{
			AttributeTypes:      a.AttributeTypes,
			CustomType:          a.CustomType,
			Required:            a.Required,
			Optional:            a.Optional,
			Computed:            a.Computed,
			Sensitive:           a.Sensitive,
			Description:         a.Description,
			MarkdownDescription: a.MarkdownDescription,
			DeprecationMessage:  a.DeprecationMessage,
			Validators:          a.Validators,
		}
	case resourceschema.ListAttribute:
		return datasourceschema.ListAttribute{
			ElementType:         a.ElementType,
			CustomType:          a.CustomType,
			Required:            a.Required,
			Optional:            a.Optional,
			Computed:            a.Computed,
			Sensitive:           a.Sensitive,
			Description:         a.Description,
			MarkdownDescription: a.MarkdownDescription,
			DeprecationMessage:  a.DeprecationMessage,
			Validators:          a.Validators,
		}
	case resourceschema.ListNestedAttribute:
		return datasourceschema.ListNestedAttribute{
			NestedObject:        DataSourceNestedAttributeObject(a.NestedObject, diags),
			CustomType:          a.CustomType,
			Required:            a.Required,
			Optional:            a.Optional,
			Computed:            a.Computed,
			Sensitive:           a.Sensitive,
			Description:         a.Description,
			MarkdownDescription: a.MarkdownDescription,
			DeprecationMessage:  a.DeprecationMessage,
			Validators:          a.Validators,
		}
	case resourceschema.MapAttribute:
		return datasourceschema.MapAttribute{
			ElementType:         a.ElementType,
			CustomType:          a.CustomType,
			Required:            a.Required,
			Optional:            a.Optional,
			Computed:            a.Computed,
			Sensitive:           a.Sensitive,
			Description:         a.Description,
			MarkdownDescription: a.MarkdownDescription,
			DeprecationMessage:  a.DeprecationMessage,
			Validators:          a.Validators,
		}
	case resourceschema.MapNestedAttribute:
		return datasourceschema.MapNestedAttribute{
			NestedObject:        DataSourceNestedAttributeObject(a.NestedObject, diags),
			CustomType:          a.CustomType,
			Required:            a.Required,
			Optional:            a.Optional,
			Computed:            a.Computed,
			Sensitive:           a.Sensitive,
			Description:         a.Description,
			MarkdownDescription: a.MarkdownDescription,
			DeprecationMessage:  a.DeprecationMessage,
			Validators:          a.Validators,
		}
	case resourceschema.SetAttribute:
		return datasourceschema.SetAttribute{
			ElementType:         a.ElementType,
			CustomType:          a.CustomType,
			Required:            a.Required,
			Optional:            a.Optional,
			Computed:            a.Computed,
			Sensitive:           a.Sensitive,
			Description:         a.Description,
			MarkdownDescription: a.MarkdownDescription,
			DeprecationMessage:  a.DeprecationMessage,
			Validators:          a.Validators,
		}
	case resourceschema.SetNestedAttribute:
		return datasourceschema.SetNestedAttribute{
			NestedObject:        DataSourceNestedAttributeObject(a.NestedObject, diags),
			CustomType:          a.CustomType,
			Required:            a.Required,
			Optional:            a.Optional,
			Computed:            a.Computed,
			Sensitive:           a.Sensitive,
			Description:         a.Description,
			MarkdownDescription: a.MarkdownDescription,
			DeprecationMessage:  a.DeprecationMessage,
			Validators:          a.Validators,
		}
	case resourceschema.SingleNestedAttribute:
		return datasourceschema.SingleNestedAttribute{
			Attributes:          DataSourceAttributeMap(a.Attributes, diags),
			CustomType:          a.CustomType,
			Required:            a.Required,
			Optional:            a.Optional,
			Computed:            a.Computed,
			Sensitive:           a.Sensitive,
			Description:         a.Description,
			MarkdownDescription: a.MarkdownDescription,
			DeprecationMessage:  a.DeprecationMessage,
			Validators:          a.Validators,
		}
	}
	diags.AddError("Provider error",
		fmt.Sprintf("Failed to convert schema attribute of type '%T' for '%s'", val, name))
	return nil
}

func ReadWithPages[T any](read func(pageID string, maxResults int32) ([]T, string, error)) ([]T, error) {
	var allResults []T
	var pageID = ""

	for {
		results, nextPageID, err := read(pageID, ReadPageSizeLimit)
		if err != nil {
			return nil, err
		}
		allResults = append(allResults, results...)
		if nextPageID == "" {
			break
		}
		pageID = nextPageID
	}

	return allResults, nil
}

// ToComputedAttributeMap converts a map of resource schema attributes to schema attributes with all fields set to "computed".
func ToComputedAttributeMap(r map[string]resourceschema.Attribute) map[string]resourceschema.Attribute {
	d := map[string]resourceschema.Attribute{}
	for k, v := range r {
		d[k] = ToComputedAttribute(k, v)
	}
	return d
}

// ToComputedNestedAttributeObject converts a resource schema nested attribute object to nested attribute object with all fields set to "computed".
func ToComputedNestedAttributeObject(r resourceschema.NestedAttributeObject) resourceschema.NestedAttributeObject {
	return resourceschema.NestedAttributeObject{
		Attributes: ToComputedAttributeMap(r.Attributes),
		CustomType: r.CustomType,
		Validators: r.Validators,
	}
}

// ToComputedAttribute converts a resource schema attribute having all attributes set to "computed".
func ToComputedAttribute(name string, val resourceschema.Attribute) resourceschema.Attribute {
	switch a := val.(type) {
	case resourceschema.StringAttribute:
		a.Required = false
		a.Optional = false
		a.Computed = true
		return a
	case resourceschema.BoolAttribute:
		a.Required = false
		a.Optional = false
		a.Computed = true
		return a
	case resourceschema.Int32Attribute:
		a.Required = false
		a.Optional = false
		a.Computed = true
		return a
	case resourceschema.Int64Attribute:
		a.Required = false
		a.Optional = false
		a.Computed = true
		return a
	case resourceschema.Float32Attribute:
		a.Required = false
		a.Optional = false
		a.Computed = true
		return a
	case resourceschema.Float64Attribute:
		a.Required = false
		a.Optional = false
		a.Computed = true
		return a
	case resourceschema.NumberAttribute:
		a.Required = false
		a.Optional = false
		a.Computed = true
		return a
	case resourceschema.ObjectAttribute:
		a.Required = false
		a.Optional = false
		a.Computed = true
		return a
	case resourceschema.ListAttribute:
		a.Required = false
		a.Optional = false
		a.Computed = true
		return a
	case resourceschema.ListNestedAttribute:
		a.NestedObject = ToComputedNestedAttributeObject(a.NestedObject)
		a.Required = false
		a.Optional = false
		a.Computed = true
		return a
	case resourceschema.MapAttribute:
		a.Required = false
		a.Optional = false
		a.Computed = true
		return a
	case resourceschema.MapNestedAttribute:
		a.NestedObject = ToComputedNestedAttributeObject(a.NestedObject)
		a.Required = false
		a.Optional = false
		a.Computed = true
		return a
	case resourceschema.SetAttribute:
		a.Required = false
		a.Optional = false
		a.Computed = true
		return a
	case resourceschema.SetNestedAttribute:
		a.NestedObject = ToComputedNestedAttributeObject(a.NestedObject)
		a.Required = false
		a.Optional = false
		a.Computed = true
		return a
	case resourceschema.SingleNestedAttribute:
		a.Attributes = ToComputedAttributeMap(a.Attributes)
		a.Required = false
		a.Optional = false
		a.Computed = true
		return a
	}

	tflog.Error(context.Background(), fmt.Sprintf("Failed to convert schema attribute of type '%T' for '%s'", val, name))
	return nil
}

func ExtractResourceRef(ref string) string {
	v := strings.SplitN(strings.Trim(ref, "/"), "/", 2)
	return v[1]
}

func FindModelFieldByTFSdkTag(model any, tagName string) (string, bool) {
	modelType := reflect.TypeOf(model)
	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
	}

	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)
		tag := field.Tag.Get("tfsdk")
		if tag == tagName {
			return field.Name, true
		}

		// Handle comma-separated options, like `tfsdk:"name,computed"`
		if parts := strings.Split(tag, ","); len(parts) > 0 && parts[0] == tagName {
			return field.Name, true
		}
	}

	return "", false
}

func ParseInterfaceValue(valStr string) interface{} {
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

// ConvertSliceOfMapsToHCL serializes a slice of []map[string]any into an HCL format.
func ConvertSliceOfMapsToHCL(data []map[string]any) string {
	var blocks []string

	for _, item := range data {
		var keyValues []string

		for key, value := range item {
			var formattedValue string

			switch v := value.(type) {
			case []map[string]any:
				nestedHCL := ConvertSliceOfMapsToHCL(v)
				formattedValue = nestedHCL
			case map[string]any:
				formattedValue = ConvertMapToHCL(v)
			case string:
				formattedValue = fmt.Sprintf("%q", v)
			case int, int64, float64:
				formattedValue = fmt.Sprintf("%v", v)
			case bool:
				formattedValue = fmt.Sprintf("%t", v)
			default:
				formattedValue = fmt.Sprintf("%q", fmt.Sprintf("%v", v))
			}

			keyValues = append(keyValues, fmt.Sprintf("        %s = %s", key, formattedValue))
		}

		block := fmt.Sprintf("      {\n%s\n      }", strings.Join(keyValues, "\n"))
		blocks = append(blocks, block)
	}

	result := fmt.Sprintf(`[
%s
    ]`, strings.Join(blocks, ",\n"))

	return result
}

// ConvertStringSliceToHCL converts a slice of strings to an HCL format.
func ConvertStringSliceToHCL(input []string) string {
	var quotedStrings []string
	for _, s := range input {
		quotedStrings = append(quotedStrings, fmt.Sprintf("%q", s))
	}
	return fmt.Sprintf("[%s]", strings.Join(quotedStrings, ", "))
}

// ConvertMapToHCL serializes a map[string]any into HCL format.
func ConvertMapToHCL(data map[string]any) string {
	var keyValues []string

	for key, value := range data {
		var formattedValue string

		switch v := value.(type) {
		case []map[string]any:
			// Handle slice of maps
			formattedValue = ConvertSliceOfMapsToHCL(v)
		case map[string]any:
			// Handle nested map
			formattedValue = ConvertMapToHCL(v)
		case []string:
			// Handle string slice
			formattedValue = ConvertStringSliceToHCL(v)
		case string:
			formattedValue = fmt.Sprintf("%q", v)
		case int, int64, float64:
			formattedValue = fmt.Sprintf("%v", v)
		case bool:
			formattedValue = fmt.Sprintf("%t", v)
		default:
			formattedValue = fmt.Sprintf("%q", fmt.Sprintf("%v", v))
		}

		keyValues = append(keyValues, fmt.Sprintf("  %s = %s", key, formattedValue))
	}

	return fmt.Sprintf("{\n%s\n}", strings.Join(keyValues, "\n"))
}

// ToUnixWithTimezone converts a naive datetime string (without offset) into a Unix timestamp
func ToUnixWithTimezone(datetimeStr string) (int64, error) {
	tUTC, err := time.ParseInLocation(NaiveDatetimeLayout, datetimeStr, time.UTC)
	if err != nil {
		return 0, fmt.Errorf("invalid datetime %q: %w", datetimeStr, err)
	}

	return tUTC.Unix(), nil
}

// FromUnixWithTimezone converts a Unix timestamp into a naive datetime string (without offset)
func FromUnixWithTimezone(ts int64) (string, error) {
	loc := time.UTC
	t := time.Unix(ts, 0).In(loc)
	return t.Format(NaiveDatetimeLayout), nil
}

// ReorderAndFilterNestedListResponse reorders and filters the state list to match the order of the plan list based on a primary key field.
func ReorderAndFilterNestedListResponse(
	ctx context.Context,
	planValue attr.Value,
	stateValue attr.Value,
	primaryKey string,
) (attr.Value, *diag.Diagnostics) {

	var diags diag.Diagnostics

	if planValue.IsNull() || planValue.IsUnknown() {
		return stateValue, &diags
	}
	if stateValue.IsNull() || stateValue.IsUnknown() {
		return planValue, &diags
	}

	planList, ok := planValue.(basetypes.ListValue)
	if !ok {
		diags.AddError("Type Error", "planValue must be a ListValue")
		return stateValue, &diags
	}
	stateList, ok := stateValue.(basetypes.ListValue)
	if !ok {
		diags.AddError("Type Error", "stateValue must be a ListValue")
		return planValue, &diags
	}

	// Convert state list into a lookup by primary key
	stateMap := make(map[string]attr.Value)
	for _, elem := range stateList.Elements() {
		obj := elem.(basetypes.ObjectValue)
		keyAttr, ok := obj.Attributes()[primaryKey]
		if !ok {
			diags.AddError("Missing Primary Key", fmt.Sprintf("State object missing primary key: %s", primaryKey))
			continue
		}
		if keyAttr.IsNull() || keyAttr.IsUnknown() {
			continue
		}
		key := keyAttr.(basetypes.StringValue).ValueString()
		stateMap[key] = elem
	}

	// Rebuild state list in the same order as plan
	var reordered []attr.Value
	for _, elem := range planList.Elements() {
		obj := elem.(basetypes.ObjectValue)
		keyAttr := obj.Attributes()[primaryKey]
		if keyAttr.IsNull() || keyAttr.IsUnknown() {
			continue
		}
		key := keyAttr.(basetypes.StringValue).ValueString()

		// Use existing state object if found, else use planned object
		if stateObj, exists := stateMap[key]; exists {
			reordered = append(reordered, stateObj)
		} else {
			reordered = append(reordered, elem)
		}
	}

	// Build new ListValue
	newList, d := basetypes.NewListValue(planList.ElementType(ctx), reordered)
	diags.Append(d...)

	return newList, &diags
}

// CopyFieldFromPlanToRespList copies a specific field from each object in the plan list to the corresponding object in the response list.
// The lists must be of the same length and contain objects in the same order of the plan values.
func CopyFieldFromPlanToRespList(ctx context.Context, planValue, respValue attr.Value, fieldName string) (attr.Value, *diag.Diagnostics) {
	var diags diag.Diagnostics

	// Check if both values are null or unknown
	if planValue.IsNull() || planValue.IsUnknown() {
		return respValue, &diags
	}

	if respValue.IsNull() || respValue.IsUnknown() {
		return respValue, &diags
	}

	planList, ok := planValue.(types.List)
	if !ok {
		diags.AddError(
			"Invalid Plan Value Type",
			fmt.Sprintf("Expected types.List, got %T", planValue),
		)
		return respValue, &diags
	}

	respList, ok := respValue.(types.List)
	if !ok {
		diags.AddError(
			"Invalid Response Value Type",
			fmt.Sprintf("Expected types.List, got %T", respValue),
		)
		return respValue, &diags
	}

	// Get the elements from both lists
	planElements := planList.Elements()
	respElements := respList.Elements()

	if len(planElements) != len(respElements) {
		diags.AddError(
			"List Length Mismatch",
			fmt.Sprintf("Plan list has %d elements, response list has %d elements",
				len(planElements), len(respElements)),
		)
		return respValue, &diags
	}

	modifiedElements := make([]attr.Value, len(respElements))

	for i, respElement := range respElements {
		planElement := planElements[i]

		// Convert elements to objects
		planObj, ok := planElement.(types.Object)
		if !ok {
			diags.AddError(
				"Invalid Plan Element Type",
				fmt.Sprintf("Expected types.Object at index %d, got %T", i, planElement),
			)
			return respValue, &diags
		}

		respObj, ok := respElement.(types.Object)
		if !ok {
			diags.AddError(
				"Invalid Response Element Type",
				fmt.Sprintf("Expected types.Object at index %d, got %T", i, respElement),
			)
			return respValue, &diags
		}

		planAttrs := planObj.Attributes()
		respAttrs := respObj.Attributes()

		// Check if the field exists in the plan object
		planFieldValue, exists := planAttrs[fieldName]
		if !exists {
			diags.AddError(
				"Field Not Found in Plan",
				fmt.Sprintf("Field '%s' not found in plan object at index %d", fieldName, i),
			)
			return respValue, &diags
		}

		// Check if the field exists in the response object
		if _, exists := respAttrs[fieldName]; !exists {
			diags.AddError(
				"Field Not Found in Response",
				fmt.Sprintf("Field '%s' not found in response object at index %d", fieldName, i),
			)
			return respValue, &diags
		}

		newAttrs := make(map[string]attr.Value)
		for k, v := range respAttrs {
			newAttrs[k] = v
		}
		newAttrs[fieldName] = planFieldValue

		// Create a new object with the modified attributes
		newObj, objDiags := types.ObjectValue(respObj.AttributeTypes(ctx), newAttrs)
		diags.Append(objDiags...)
		if objDiags.HasError() {
			return respValue, &diags
		}

		modifiedElements[i] = newObj
	}

	// Create a new list with the modified elements
	newList, listDiags := types.ListValue(respList.ElementType(ctx), modifiedElements)
	diags.Append(listDiags...)
	if listDiags.HasError() {
		return respValue, &diags
	}

	return newList, &diags
}

// CopyFieldFromPlanToRespObject copies a specific field from the plan object to the response object.
func CopyFieldFromPlanToRespObject(ctx context.Context, planValue, respValue attr.Value, fieldName string) (attr.Value, *diag.Diagnostics) {
	var diags diag.Diagnostics

	if planValue.IsNull() || planValue.IsUnknown() {
		return respValue, &diags
	}

	if respValue.IsNull() || respValue.IsUnknown() {
		return respValue, &diags
	}

	planObject, ok := planValue.(types.Object)
	if !ok {
		diags.AddError(
			"Invalid Plan Value Type",
			fmt.Sprintf("Expected types.Object, got %T", planValue),
		)
		return respValue, &diags
	}

	respObject, ok := respValue.(types.Object)
	if !ok {
		diags.AddError(
			"Invalid Response Value Type",
			fmt.Sprintf("Expected types.Object, got %T", respValue),
		)
		return respValue, &diags
	}

	planAttrs := planObject.Attributes()
	respAttrs := respObject.Attributes()

	planFieldValue, exists := planAttrs[fieldName]
	if !exists {
		diags.AddError(
			"Field Not Found in Plan",
			fmt.Sprintf("Field '%s' not found in plan object", fieldName),
		)
		return respValue, &diags
	}

	if _, exists := respAttrs[fieldName]; !exists {
		diags.AddError(
			"Field Not Found in Response",
			fmt.Sprintf("Field '%s' not found in response object", fieldName),
		)
		return respValue, &diags
	}

	respAttrs[fieldName] = planFieldValue

	newObj, objDiags := types.ObjectValue(respObject.AttributeTypes(ctx), respAttrs)
	diags.Append(objDiags...)
	if objDiags.HasError() {
		return respValue, &diags
	}

	return newObj, &diags
}

func ReorderAndFilterDHCPOptions(
	ctx context.Context,
	planValue attr.Value,
	stateValue attr.Value,
) (attr.Value, *diag.Diagnostics) {

	var diags diag.Diagnostics
	primaryKey := "name"
	secondaryKey := "num"

	// Handle null/unknown gracefully
	if planValue.IsNull() || planValue.IsUnknown() {
		return stateValue, &diags
	}
	if stateValue.IsNull() || stateValue.IsUnknown() {
		return planValue, &diags
	}

	planList, ok := planValue.(basetypes.ListValue)
	if !ok {
		diags.AddError("Type Error", "planValue must be a basetypes.ListValue")
		return stateValue, &diags
	}
	stateList, ok := stateValue.(basetypes.ListValue)
	if !ok {
		diags.AddError("Type Error", "stateValue must be a basetypes.ListValue")
		return planValue, &diags
	}

	// Build lookup maps from state: name->element and num->element
	nameToState := make(map[string]attr.Value)
	numToState := make(map[int64]attr.Value)

	for _, elem := range stateList.Elements() {
		obj, ok := elem.(basetypes.ObjectValue)
		if !ok {
			continue
		}
		attrs := obj.Attributes()

		// name -> basetypes.StringValue (per your note state has both keys)
		if nameAttr, has := attrs[primaryKey]; has && nameAttr != nil && !nameAttr.IsNull() && !nameAttr.IsUnknown() {
			if strVal, ok := nameAttr.(basetypes.StringValue); ok {
				nameToState[strVal.ValueString()] = elem
			}
		}

		// num -> basetypes.Int64Value (per your note state has both keys)
		if numAttr, has := attrs[secondaryKey]; has && numAttr != nil && !numAttr.IsNull() && !numAttr.IsUnknown() {
			if intVal, ok := numAttr.(basetypes.Int64Value); ok {
				numToState[intVal.ValueInt64()] = elem
			}
		}
	}

	// Rebuild ordered slice based on plan order
	var reordered []attr.Value
	for _, planElem := range planList.Elements() {
		planObj, ok := planElem.(basetypes.ObjectValue)
		if !ok {
			// if plan contains something else, append it as fallback
			reordered = append(reordered, planElem)
			continue
		}
		planAttributes := planObj.Attributes()

		var matchedState attr.Value

		// Try primaryKey (name) first if present and valid
		if primaryKeyAttribute, has := planAttributes[primaryKey]; has && primaryKeyAttribute != nil && !primaryKeyAttribute.IsNull() && !primaryKeyAttribute.IsUnknown() {
			if primaryKeyAttributeValue, ok := primaryKeyAttribute.(basetypes.StringValue); ok {
				if s, exists := nameToState[primaryKeyAttributeValue.ValueString()]; exists {
					matchedState = s
				}
			}
		}

		// If not matched by name, try secondaryKey (num)
		if matchedState == nil {
			if secondaryKeyAttribute, has := planAttributes[secondaryKey]; has && secondaryKeyAttribute != nil && !secondaryKeyAttribute.IsNull() && !secondaryKeyAttribute.IsUnknown() {
				if secondaryKeyAttributeValue, ok := secondaryKeyAttribute.(basetypes.Int64Value); ok {
					if s, exists := numToState[secondaryKeyAttributeValue.ValueInt64()]; exists {
						matchedState = s
					}
				}
			}
		}

		// If matchedState found, use it; else fall back to plan element itself
		if matchedState != nil {
			reordered = append(reordered, matchedState)
		} else {
			reordered = append(reordered, planElem)
		}
	}

	// Create new ListValue with same element type as plan list
	newList, d := basetypes.NewListValue(planList.ElementType(ctx), reordered)
	diags.Append(d...)

	return newList, &diags
}

// UploadInitResponse struct represents the response from the NIOS uploadinit API call
type UploadInitResponse struct {
	Token string `json:"token"`
	URL   string `json:"url"`
}

// GenerateUploadToken generates an upload token and URL by calling the NIOS uploadinit API
func GenerateUploadToken(ctx context.Context, baseURL, username, password string) (*UploadInitResponse, error) {
	var uploadInitResponse UploadInitResponse
	// Extract configuration values using reflection for flexibility
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// Generate upload token and URL by calling uploadinit
	uploadInitURL := fmt.Sprintf("%s/wapi/v2.13.6/fileop?_function=uploadinit", baseURL)
	req, err := http.NewRequestWithContext(ctx, "POST", uploadInitURL, bytes.NewReader([]byte("{}")))
	if err != nil {
		return &uploadInitResponse, fmt.Errorf("error creating uploadinit request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, password)

	tflog.Debug(ctx, fmt.Sprintf("Making uploadinit request to: %s", uploadInitURL))
	resp, err := httpClient.Do(req)
	if err != nil {
		return &uploadInitResponse, fmt.Errorf("error making uploadinit request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return &uploadInitResponse, fmt.Errorf("uploadinit request failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	if err := json.NewDecoder(resp.Body).Decode(&uploadInitResponse); err != nil {
		return &uploadInitResponse, fmt.Errorf("error decoding uploadinit response: %w", err)
	}

	tflog.Debug(ctx, fmt.Sprintf("Generated upload token: %s with URL: %s", uploadInitResponse.Token, uploadInitResponse.URL))
	return &uploadInitResponse, nil
}

// UploadPEMFile uploads a PEM file to the Infoblox NIOS server using the provided upload URL.
func UploadPEMFile(ctx context.Context, uploadURL, pemFilePath, username, password string) error {
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	file, err := os.Open(pemFilePath)
	if err != nil {
		return fmt.Errorf("error opening PEM file: %w", err)
	}
	defer file.Close()
	// Create a buffer for the multipart form
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Create the form file field
	part, err := writer.CreateFormFile("file", filepath.Base(pemFilePath))
	if err != nil {
		return fmt.Errorf("error creating form file: %w", err)
	}
	// Copy the file content to the form field
	if _, err = io.Copy(part, file); err != nil {
		return fmt.Errorf("error copying file content: %w", err)
	}
	//Close the multipart writer to finalize the form
	if err = writer.Close(); err != nil {
		return fmt.Errorf("error finalizing multipart form: %w", err)
	}

	// Create a new request to upload the file
	uploadReq, err := http.NewRequestWithContext(ctx, "POST", uploadURL, &requestBody)
	if err != nil {
		return fmt.Errorf("error creating file upload request: %w", err)
	}

	uploadReq.Header.Set("Content-Type", writer.FormDataContentType())
	uploadReq.SetBasicAuth(username, password)

	tflog.Debug(ctx, fmt.Sprintf("Uploading PEM file to: %s", uploadURL))
	uploadResp, err := httpClient.Do(uploadReq)
	if err != nil {
		return fmt.Errorf("error uploading PEM file: %w", err)
	}
	defer uploadResp.Body.Close()
	if uploadResp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(uploadResp.Body)
		return fmt.Errorf("file upload failed with status %d: %s", uploadResp.StatusCode, string(bodyBytes))
	}

	tflog.Info(ctx, fmt.Sprintf("PEM file %s uploaded successfully", pemFilePath))
	return nil
}

// UploadPEMFileWithToken handles the complete process of generating an upload token and uploading a PEM file
// It returns the token from the successful upload or an error if any step fails
func UploadPEMFileWithToken(ctx context.Context, baseUrl, filePath, username, password string) (string, error) {
	// Generate the upload token
	uploadInitResponse, err := GenerateUploadToken(ctx, baseUrl, username, password)
	if err != nil {
		return "", fmt.Errorf("unable to generate upload token: %w", err)
	}

	// Upload the file using the token URL
	if err = UploadPEMFile(ctx, uploadInitResponse.URL, filePath, username, password); err != nil {
		return "", fmt.Errorf("unable to upload file: %w", err)
	}

	return uploadInitResponse.Token, nil
}
