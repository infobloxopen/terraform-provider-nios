package validator

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = nacFilterExpressionValidator{}

// nacFilterExpressionValidator validates that the provided string is a valid filter expression
type nacFilterExpressionValidator struct{}

func (v nacFilterExpressionValidator) Description(ctx context.Context) string {
	return "A filter expression should start and end with parenthesis and contain option field and values with AND/OR."
}

func (v nacFilterExpressionValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v nacFilterExpressionValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()

	// Allow empty string
	if value == "" {
		return
	}

	pattern := `^\(.+\)$`
	matched, err := regexp.MatchString(pattern, value)
	if err != nil {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			req.Path,
			"Failed to validate expression pattern.",
			value,
		))
		return
	}

	if !matched {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			req.Path,
			"Expression must begin with '(' and end with ')' with content inside, for example:\n(Radius.ServerError=\"false\" AND Sophos.ComplianceState=\"NonCompliant\" AND Radius.ServerResponse=\"accept\" AND (Radius.ServerState=\"disabled\" OR Sophos.UserClass=\"example_user_class\" OR ()))",
			value,
		))
	}
}

// ValidateNACFilterExpression returns a validator that ensures the string is a valid filter expression
// wrapped in parentheses with content inside.
func ValidateNACFilterExpression() validator.String {
	return nacFilterExpressionValidator{}
}
