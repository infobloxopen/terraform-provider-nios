package validator

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type RPZAAAAIPv6Validator struct{}

func (v RPZAAAAIPv6Validator) Description(ctx context.Context) string {
	return "Validator to check if string is in the RPZ AAAA IPv6 format '<ipv6>.<rp-zone>' or '<ipv6-cidr>.<rp-zone>'"
}

func (v RPZAAAAIPv6Validator) MarkdownDescription(ctx context.Context) string {
	return "Validator to check if string is in the RPZ AAAA IPv6 format '<ipv6>.<rp-zone>' or '<ipv6-cidr>.<rp-zone>'"
}

func (v RPZAAAAIPv6Validator) ValidateString(
	ctx context.Context,
	request validator.StringRequest,
	response *validator.StringResponse,
) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	value := request.ConfigValue.ValueString()

	// Check for whitespace
	if strings.TrimSpace(value) != value {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Invalid Whitespace",
			"The value must not contain leading or trailing whitespace.",
		)
		return
	}

	re := regexp.MustCompile(`^(?:` +
		`[0-9A-Fa-f:]+` + // ipv6
		`(?:/(?:[0-9]|[1-9][0-9]|1[01][0-9]|12[0-8]))?` + // optional /0-128
		`)\.[A-Za-z0-9.-]+$`) // rp-zone

	if !re.MatchString(value) {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Invalid RPZ IPv6 Format",
			fmt.Sprintf(
				"The value '%s' is not a valid RPZ AAAA IPv6. Expected format: '<ipv6>.<rp-zone>' or '<ipv6-cidr>.<rp-zone>' (example: '2001:db8::1.rpz.example.com' or '2001:db8::/64.rpz.example.com').",
				value,
			),
		)
	}
}

// IsValidRPZAAAAIPv6 returns a validator that ensures the input is a valid RPZ AAAA IPv6.
func IsValidRPZAAAAIPv6() validator.String {
	return RPZAAAAIPv6Validator{}
}
