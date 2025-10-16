package grid

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"

	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForExtensibleattributedef = "allowed_object_types,comment,default_value,flags,list_values,max,min,name,namespace,type"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ExtensibleattributedefResource{}
var _ resource.ResourceWithImportState = &ExtensibleattributedefResource{}

// Constants limits for validation
const (
	MinIntegerValue  = -2147483648
	MaxIntegerValue  = 2147483647
	MinStringLength  = 0
	MaxStringLength  = 256
	MinEnumLength    = 1
	MaxEnumLength    = 256
	MinURLLength     = 1
	MaxURLLength     = 256
	MinEmailLength   = 1
	MaxEmailLength   = 256
	MinDateTimeValue = 0
	MaxDateTimeValue = 2147483647 // Unix max time Jan 1, 2038
)

func NewExtensibleattributedefResource() resource.Resource {
	return &ExtensibleattributedefResource{}
}

// ExtensibleattributedefResource defines the resource implementation.
type ExtensibleattributedefResource struct {
	client *niosclient.APIClient
}

func (r *ExtensibleattributedefResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "grid_extensibleattributedef"
}

func (r *ExtensibleattributedefResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages an Extensible Attribute definition.",
		Attributes:          ExtensibleattributedefResourceSchemaAttributes,
	}
}

func (r *ExtensibleattributedefResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*niosclient.APIClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *niosclient.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *ExtensibleattributedefResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ExtensibleattributedefModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.GridAPI.
		ExtensibleattributedefAPI.
		Create(ctx).
		Extensibleattributedef(*data.Expand(ctx, &resp.Diagnostics, true)).
		ReturnFieldsPlus(readableAttributesForExtensibleattributedef).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Extensibleattributedef, got error: %s", err))
		return
	}

	res := apiRes.CreateExtensibleattributedefResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ExtensibleattributedefResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ExtensibleattributedefModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.GridAPI.
		ExtensibleattributedefAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForExtensibleattributedef).
		ReturnAsObject(1).
		Execute()

	// Handle not found case
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			// Resource no longer exists, remove from state
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Extensibleattributedef, got error: %s", err))
		return
	}

	res := apiRes.GetExtensibleattributedefResponseObjectAsResult.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ExtensibleattributedefResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data ExtensibleattributedefModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	diags = req.State.GetAttribute(ctx, path.Root("ref"), &data.Ref)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	apiRes, _, err := r.client.GridAPI.
		ExtensibleattributedefAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Extensibleattributedef(*data.Expand(ctx, &resp.Diagnostics, false)).
		ReturnFieldsPlus(readableAttributesForExtensibleattributedef).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Extensibleattributedef, got error: %s", err))
		return
	}

	res := apiRes.UpdateExtensibleattributedefResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ExtensibleattributedefResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ExtensibleattributedefModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.GridAPI.
		ExtensibleattributedefAPI.
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Extensibleattributedef, got error: %s", err))
		return
	}
}

func (r *ExtensibleattributedefResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("ref"), req, resp)
}

func (r *ExtensibleattributedefResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data ExtensibleattributedefModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Skip validation if type is unknown
	if data.Type.IsUnknown() {
		return
	}

	typeValue := data.Type.ValueString()

	// Validate type-specific field constraints
	r.validateTypeSpecificFields(ctx, data, resp)

	// Skip default_value validation if it's unknown or null
	if data.DefaultValue.IsUnknown() || data.DefaultValue.IsNull() {
		return
	}

	defaultValue := data.DefaultValue.ValueString()

	// Validate default_value based on type
	switch typeValue {
	case "INTEGER":
		r.validateIntegerDefaultValue(defaultValue, &data, resp)
	case "EMAIL":
		r.validateEmailDefaultValue(defaultValue, resp)
	case "URL":
		r.validateURLDefaultValue(defaultValue, resp)
	case "DATE":
		r.validateDateDefaultValue(defaultValue, resp)
	case "STRING":
		r.validateStringDefaultValue(defaultValue, &data, resp)
	case "ENUM":
		r.validateEnumDefaultValue(ctx, data, resp)
	}
}

func (r *ExtensibleattributedefResource) validateTypeSpecificFields(ctx context.Context, data ExtensibleattributedefModel, resp *resource.ValidateConfigResponse) {
	typeValue := data.Type.ValueString()

	// Check if list_values is provided for non-ENUM types
	if typeValue != "ENUM" && !data.ListValues.IsNull() && !data.ListValues.IsUnknown() {
		var listValues []ExtensibleattributedefListValuesModel
		diags := data.ListValues.ElementsAs(ctx, &listValues, false)
		if !diags.HasError() && len(listValues) > 0 {
			resp.Diagnostics.AddError(
				"Invalid list_values for Type",
				fmt.Sprintf("list_values can only be specified when type is 'ENUM', but type is '%s'. Remove the list_values block.", typeValue),
			)
		}
	}

	// Check if min/max is provided for non-INTEGER and non-STRING types
	if typeValue != "INTEGER" && typeValue != "STRING" {
		if !data.Min.IsNull() || !data.Max.IsNull() {
			resp.Diagnostics.AddError(
				"Invalid min/max for Type",
				fmt.Sprintf("min and max can only be specified when type is 'INTEGER' or 'STRING', but type is '%s'. Remove the min and max fields.", typeValue),
			)
		}
	}
}

func (r *ExtensibleattributedefResource) validateIntegerDefaultValue(value string, data *ExtensibleattributedefModel, resp *resource.ValidateConfigResponse) {
	// Parse integer value
	intVal, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid Integer Default Value",
			fmt.Sprintf("The default_value '%s' is not a valid integer. Please provide a valid integer value.", value),
		)
		return
	}

	// Check if it's within int32 range
	if intVal < MinIntegerValue || intVal > MaxIntegerValue {
		resp.Diagnostics.AddError(
			"Integer Default Value Out of Range",
			fmt.Sprintf("The default_value '%d' is out of range. Integer values must be between %d and %d.", intVal, MinIntegerValue, MaxIntegerValue),
		)
		return
	}

	// Check against min/max constraints if provided
	if !data.Min.IsNull() && !data.Max.IsNull() {
		min := data.Min.ValueInt64()
		max := data.Max.ValueInt64()

		if intVal < min || intVal > max {
			resp.Diagnostics.AddError(
				"Default Value Outside Min/Max Range",
				fmt.Sprintf("The default_value '%d' is outside the specified range. Must be between %d and %d.", intVal, min, max),
			)
		}
	}
}

func (r *ExtensibleattributedefResource) validateEmailDefaultValue(value string, resp *resource.ValidateConfigResponse) {
	// Check length first
	if len(value) < MinEmailLength || len(value) > MaxEmailLength {
		resp.Diagnostics.AddError(
			"Email Default Value Length Invalid",
			fmt.Sprintf("The default_value for email type must be between %d and %d characters. Current length: %d", MinEmailLength, MaxEmailLength, len(value)),
		)
		return
	}

	// Use the email regex patterns from Python code
	emailRegex1 := regexp.MustCompile(`^[a-zA-Z0-9!#$%&'*+/=?^_` + "`" + `{|}~-]+(?:\.[a-zA-Z0-9!#$%&'*+/=?^_` + "`" + `{|}~-]+)*@(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?\.)+[a-zA-Z0-9](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`)
	emailRegex2 := regexp.MustCompile(`^[a-zA-Z0-9+_\-.]+@[0-9a-zA-Z][.\-0-9a-zA-Z]*\.[a-zA-Z]+$`)

	if !emailRegex1.MatchString(value) && !emailRegex2.MatchString(value) {
		resp.Diagnostics.AddError(
			"Invalid Email Default Value",
			fmt.Sprintf("The default_value '%s' is not a valid email format. Must be like: user@example.com", value),
		)
	}
}

func (r *ExtensibleattributedefResource) validateURLDefaultValue(value string, resp *resource.ValidateConfigResponse) {
	// Check length first
	if len(value) < MinURLLength || len(value) > MaxURLLength {
		resp.Diagnostics.AddError(
			"URL Default Value Length Invalid",
			fmt.Sprintf("The default_value for URL type must be between %d and %d characters. Current length: %d", MinURLLength, MaxURLLength, len(value)),
		)
		return
	}

	// Validate URL format - must have scheme and host
	parsedURL, err := url.ParseRequestURI(value)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		resp.Diagnostics.AddError(
			"Invalid URL Default Value",
			fmt.Sprintf("The default_value '%s' is not a valid URL. Must include scheme and host (e.g., https://example.com).", value),
		)
	}
}

func (r *ExtensibleattributedefResource) validateDateDefaultValue(value string, resp *resource.ValidateConfigResponse) {
	// Date should be Unix timestamp (seconds since epoch)
	timestamp, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid Date Default Value",
			fmt.Sprintf("The default_value '%s' must be a Unix timestamp (number of seconds since January 1st, 1970 UTC).", value),
		)
		return
	}

	// Check if it's a reasonable date (not negative, not beyond Unix time limit)
	if timestamp < MinDateTimeValue || timestamp > MaxDateTimeValue {
		resp.Diagnostics.AddError(
			"Date Default Value Out of Range",
			fmt.Sprintf("The default_value timestamp '%d' is out of valid range. Must be between %d and %d (January 1st, 1970 to January 1st, 2038).", timestamp, MinDateTimeValue, MaxDateTimeValue),
		)
		return
	}

	// Optional: Warn if date is far in the future
	futureLimit := time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
	if timestamp > futureLimit {
		resp.Diagnostics.AddWarning(
			"Future Date Default Value",
			fmt.Sprintf("The default_value timestamp '%d' represents a date far in the future (%s). Please verify this is correct.", timestamp, time.Unix(timestamp, 0).Format("2006-01-02")),
		)
	}
}

func (r *ExtensibleattributedefResource) validateStringDefaultValue(value string, data *ExtensibleattributedefModel, resp *resource.ValidateConfigResponse) {
	// Check base string length constraints
	if len(value) < MinStringLength || len(value) > MaxStringLength {
		resp.Diagnostics.AddError(
			"String Default Value Length Invalid",
			fmt.Sprintf("The default_value for string type must be between %d and %d characters. Current length: %d", MinStringLength, MaxStringLength, len(value)),
		)
		return
	}

	// Check against min/max length constraints if provided (for STRING type, min/max refer to length)
	if !data.Min.IsNull() && !data.Max.IsNull() {
		min := data.Min.ValueInt64()
		max := data.Max.ValueInt64()
		valueLen := int64(len(value))

		if valueLen < min || valueLen > max {
			resp.Diagnostics.AddError(
				"String Default Value Length Outside Range",
				fmt.Sprintf("The default_value length '%d' is outside the specified range. Must be between %d and %d characters.", valueLen, min, max),
			)
		}
	}

	// Empty string validation
	if len(value) == 0 {
		resp.Diagnostics.AddError(
			"Missing Default Value",
			"The default_value for string type cannot be empty. Please provide a valid string value.",
		)
	}
}

func (r *ExtensibleattributedefResource) validateEnumDefaultValue(ctx context.Context, data ExtensibleattributedefModel, resp *resource.ValidateConfigResponse) {
	defaultValue := data.DefaultValue.ValueString()

	// Check enum value length
	if len(defaultValue) < MinEnumLength || len(defaultValue) > MaxEnumLength {
		resp.Diagnostics.AddError(
			"Enum Default Value Length Invalid",
			fmt.Sprintf("The default_value for enum type must be between %d and %d characters. Current length: %d", MinEnumLength, MaxEnumLength, len(defaultValue)),
		)
		return
	}

	// For ENUM type, the default_value should match one of the list_values
	if data.ListValues.IsNull() || data.ListValues.IsUnknown() {
		resp.Diagnostics.AddError(
			"Missing List Values for ENUM",
			"When type is 'ENUM', list_values must be provided to define the possible enum values.",
		)
		return
	}

	var listValues []ExtensibleattributedefListValuesModel
	diags := data.ListValues.ElementsAs(ctx, &listValues, false)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	if len(listValues) == 0 {
		resp.Diagnostics.AddError(
			"Empty List Values for ENUM",
			"When type is 'ENUM', at least one list_value must be provided.",
		)
		return
	}

	// Validate each enum value in list_values and check if default_value matches
	validValues := make([]string, 0, len(listValues))
	foundMatch := false

	for i, listValue := range listValues {
		if listValue.Value.IsNull() {
			resp.Diagnostics.AddError(
				"Invalid Enum List Value",
				fmt.Sprintf("list_values[%d].value cannot be null.", i),
			)
			continue
		}

		enumValue := listValue.Value.ValueString()

		// Validate enum value length
		if len(enumValue) < MinEnumLength || len(enumValue) > MaxEnumLength {
			resp.Diagnostics.AddError(
				"Enum List Value Length Invalid",
				fmt.Sprintf("list_values[%d].value '%s' length must be between %d and %d characters.", i, enumValue, MinEnumLength, MaxEnumLength),
			)
			continue
		}

		validValues = append(validValues, enumValue)

		// Check if this matches the default value
		if enumValue == defaultValue {
			foundMatch = true
		}
	}

	if !foundMatch && len(validValues) > 0 {
		resp.Diagnostics.AddError(
			"Invalid ENUM Default Value",
			fmt.Sprintf("The default_value '%s' is not one of the allowed enum values. Valid values are: %v", defaultValue, validValues),
		)
	}
}
