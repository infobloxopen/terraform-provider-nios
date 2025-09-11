package validator

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// exactlyOneOfValidator validates that exactly one of the specified attributes is set.
type exactlyOneOfValidator struct {
	attributes []string
}

func (v exactlyOneOfValidator) Description(ctx context.Context) string {
	return "Validates that exactly one of the specified attributes is set"
}

func (v exactlyOneOfValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v exactlyOneOfValidator) ValidateList(ctx context.Context, request validator.ListRequest, response *validator.ListResponse) {
	// Only validate the attribute configuration value if it is known.
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	count := 0

	for _, attrName := range v.attributes {
		// Get the attribute value
		var attrValue attr.Value
		diags := request.Config.GetAttribute(ctx, path.Root(attrName), &attrValue)
		response.Diagnostics.Append(diags...)

		if diags.HasError() {
			return
		}

		// Check if the attribute is specified
		if !attrValue.IsNull() && !attrValue.IsUnknown() {
			count++
		}
	}

	if count == 0 || count > 1 {
		response.Diagnostics.Append(validatordiag.InvalidAttributeCombinationDiagnostic(
			request.Path,
			fmt.Sprintf(
				"When %s is specified, exactly one of these attributes should be specified: %s",
				request.Path.String(),
				strings.Join(v.attributes, ", "),
			),
		))

	}
}

// ExactlyOneOf returns a validator which ensures that exactly one of the
// specified attributes is set.
func ExactlyOneOf(attributes ...string) validator.List {
	return exactlyOneOfValidator{
		attributes: attributes,
	}
}
