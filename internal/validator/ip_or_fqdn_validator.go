package validator

import (
	"context"
	"net"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = iPOrFQDNValidator{}

// iPOrFQDNValidator validates if the provided value is a valid IP address or a string without leading/trailing whitespace, trailing dot, or uppercase characters.
type iPOrFQDNValidator struct{}

func (v iPOrFQDNValidator) Description(ctx context.Context) string {
	return "value must be a valid IP address or a string without leading/trailing whitespace, trailing dot, or uppercase characters"
}

func (v iPOrFQDNValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v iPOrFQDNValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()

	if ip := net.ParseIP(value); ip != nil {
		return // valid IP address
	}

	// If it looks like an IP address but failed to parse, surface an error.
	if looksLikeIPv4(value) {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			req.Path,
			"Invalid IP address",
			value,
		))
		return
	}

	validFQDN := fqdnValidator{}
	fqdnResp := &validator.StringResponse{}
	validFQDN.ValidateString(ctx, req, fqdnResp)

	if fqdnResp.Diagnostics.HasError() {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			req.Path,
			"Invalid FQDN. Must be a valid DNS record name in FQDN format without leading/trailing whitespace or trailing dot or uppercase characters.",
			value,
		))
	}
}

// looksLikeIPv4 returns true if the string resembles an IPv4 address (N.N.N.N with digits only)
func looksLikeIPv4(s string) bool {
	if strings.Count(s, ".") != 3 {
		return false
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if (c < '0' || c > '9') && c != '.' {
			return false
		}
	}
	// Between 1 and 3 dots to catch incomplete IPv4 like "10.0.0"
	dots := strings.Count(s, ".")
	if dots < 1 || dots > 3 {
		return false
	}
	return true
}

func IsValidIPv4OrFQDN() validator.String {
	return iPOrFQDNValidator{}
}
