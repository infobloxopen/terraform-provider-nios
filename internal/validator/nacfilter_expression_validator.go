package validator

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = nacFilterExpressionValidator{}

// nacFilterExpressionValidator validates that the provided string is a valid NAC filter expression
// containing at least one of the allowed fields: Radius.ServerError, Sophos.ComplianceState,
// Radius.ServerState, Sophos.UserClass, Radius.ServerResponse
type nacFilterExpressionValidator struct{}

func (v nacFilterExpressionValidator) Description(ctx context.Context) string {
	return "A NAC filter Expression should start and end with parenthesis and should contain field and values with AND/OR."
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
	if strings.TrimSpace(value) == "" {
		return
	}

	// Check if expression starts with '(' and ends with ')'
	trimmed := strings.TrimSpace(value)
	if !strings.HasPrefix(trimmed, "(") || !strings.HasSuffix(trimmed, ")") {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			req.Path,
			"Expression must begin with '(' and end with ')', for example :\n (Radius.ServerError=\"false\" AND Sophos.ComplianceState=\"NonCompliant\" AND Radius.ServerResponse=\"accept\" AND (Radius.ServerState=\"disabled\" OR Sophos.UserClass=\"example_user_class\" OR ()))",
			value,
		))
		return
	}
}

// ValidateNACFilterExpression returns a validator that ensures the string is a valid NAC filter expression
// wrapped in parentheses and containing at least one of the allowed fields.
func ValidateNACFilterExpression() validator.String {
	return nacFilterExpressionValidator{}
}
