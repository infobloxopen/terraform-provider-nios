package validator

import (
	"context"
	"fmt"
	"net/netip"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// IPOrFQDNValidator validates that the provided string is either:
// - A valid IPv4 address
// - A valid IPv6 address
// - A valid FQDN
type IPOrFQDNValidator struct{}

func (v IPOrFQDNValidator) Description(ctx context.Context) string {
	return "String must be a valid IPv4 address, IPv6 address, or FQDN"
}

func (v IPOrFQDNValidator) MarkdownDescription(ctx context.Context) string {
	return "String must be a valid IPv4 address, IPv6 address, or FQDN"
}

func (v IPOrFQDNValidator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	value := request.ConfigValue.ValueString()

	// Check if the value is a valid IPv4 address
	if ip, err := netip.ParseAddr(value); err == nil && ip.Is4() {
		return
	}

	// Check if the value is a valid IPv6 address
	if ip, err := netip.ParseAddr(value); err == nil && ip.Is6() {
		return
	}

	// If it looks like an IP address but failed to parse, surface an error
	if looksLikeIPv4(value) {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Invalid IPv4 Address",
			fmt.Sprintf("The value '%s' appears to be an IPv4 address but is not valid.", value),
		)
		return
	}

	if looksLikeIPv6(value) {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Invalid IPv6 Address",
			fmt.Sprintf("The value '%s' appears to be an IPv6 address but is not valid.", value),
		)
		return
	}

	// Check if it's a valid FQDN using the domain name validator
	domainValidator := IsValidDomainName()
	fqdnResp := &validator.StringResponse{}
	domainValidator.ValidateString(ctx, request, fqdnResp)

	if !fqdnResp.Diagnostics.HasError() {
		return
	}

	// If all validations fail, add an error
	response.Diagnostics.AddAttributeError(
		request.Path,
		"Invalid Value Format",
		fmt.Sprintf("The value '%s' is not a valid IPv4 address, IPv6 address, or FQDN.", value),
	)
}

// looksLikeIPv6 returns true if the string resembles an IPv6 address
func looksLikeIPv6(s string) bool {
	return strings.Count(s, ":") >= 2
}

// IsValidIPOrFQDN returns a validator that ensures the input is a valid IPv4, IPv6, or FQDN.
func IsValidIPOrFQDN() validator.String {
	return IPOrFQDNValidator{}
}
