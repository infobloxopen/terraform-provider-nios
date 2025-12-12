package validator

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type RPZAIPv4Validator struct{}

func (v RPZAIPv4Validator) Description(ctx context.Context) string {
	return "Validator to check if string is in the RPZ A IPv4 format '<ipv4>.<rp-zone>' or '<ipv4>/<prefix>.<rp-zone>'"
}

func (v RPZAIPv4Validator) MarkdownDescription(ctx context.Context) string {
	return "Validator to check if string is in the RPZ A IPv4 format '<ipv4>.<rp-zone>' or '<ipv4>/<prefix>.<rp-zone>'"
}

func (v RPZAIPv4Validator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
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
		`(?:25[0-5]|2[0-4]\d|[01]?\d?\d)(?:\.(?:25[0-5]|2[0-4]\d|[01]?\d?\d)){3}` + // ipv4
		`(?:/(?:[0-9]|[1-2][0-9]|3[0-2]))?` + // optional /0-32
		`)\.[A-Za-z0-9.-]+$`) // rp-zone

	if !re.MatchString(value) {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Invalid RPZ IPv4 Format",
			fmt.Sprintf("The value '%s' is not a valid RPZ A IPv4. Expected format: '<ipv4>.<rp-zone>' or '<ipv4-cidr>.<rp-zone>' (example: '10.10.0.1.rpz.example.com' or '10.0.0.0/16.rpz.example.com').", value),
		)
	}
}

// IsValidRPZAIPv4 returns a validator that ensures the input is a valid RPZ A IPv4.
func IsValidRPZAIPv4() validator.String {
	return RPZAIPv4Validator{}
}
