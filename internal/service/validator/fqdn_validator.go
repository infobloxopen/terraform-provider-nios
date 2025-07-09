package validator

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = dnsRecordNameValidator{}

// dnsRecordNameValidator validates if the provided value is a valid DNS record name.
type dnsRecordNameValidator struct {
}

func (validator dnsRecordNameValidator) Description(ctx context.Context) string {
	return "value must be a valid DNS record name in FQDN format without leading/trailing whitespace or trailing dot"
}

func (validator dnsRecordNameValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

func (validator dnsRecordNameValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	// Only validate the attribute configuration value if it is known.
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()

	if strings.TrimSpace(value) != value {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeTypeDiagnostic(
			req.Path,
			"must not contain leading or trailing whitespaces",
			req.ConfigValue.ValueString(),
		))
	}
	// Check for trailing dot
	if strings.HasSuffix(value, ".") {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeTypeDiagnostic(
			req.Path,
			"must not end with a dot",
			req.ConfigValue.ValueString(),
		))
	}
}

// IsRecordNameValid returns an AttributeValidator which ensures that any configured
// attribute value:
//
//   - Contains no whitespace.
//   - Does not end with a dot
func IsRecordNameValid() validator.String {
	return dnsRecordNameValidator{}
}
