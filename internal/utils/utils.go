package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const ReadPageSizeLimit int32 = 1000

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
