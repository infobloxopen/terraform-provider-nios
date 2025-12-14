package validator

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type ConditionalIPv4OrFQDNValidator struct {
	IPv4Validator validator.String
	FQDNValidator validator.String
}

func (v ConditionalIPv4OrFQDNValidator) Description(ctx context.Context) string {
	return "Validates input as either a valid IPv4 address or a valid FQDN."
}

func (v ConditionalIPv4OrFQDNValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v ConditionalIPv4OrFQDNValidator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	value := request.ConfigValue.ValueString()

	// validate if it is an IPv4 address
	if regexp.MustCompile(`^\d+\.\d+\.\d+\.\d+$`).MatchString(value) {
		ipv4Resp := &validator.StringResponse{}

		v.IPv4Validator.ValidateString(ctx, request, ipv4Resp)
		if ipv4Resp.Diagnostics.HasError() {
			response.Diagnostics.AddAttributeError(
				request.Path,
				"Invalid IPv4 Address",
				fmt.Sprintf("The value '%s' must be a valid IPv4 address or a valid FQDN.", value),
			)
		}
		return
	}

	// else validate as FQDN
	v.FQDNValidator.ValidateString(ctx, request, response)
}

// IsValidIPv4OrFQDN creates a new instance of the validator.
func IsValidIPv4OrFQDN(ipv4, fqdn validator.String) validator.String {
	return ConditionalIPv4OrFQDNValidator{
		IPv4Validator: ipv4,
		FQDNValidator: fqdn,
	}
}
