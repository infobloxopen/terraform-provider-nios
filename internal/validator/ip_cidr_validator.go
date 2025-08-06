package validator

import (
	"context"
	"fmt"
	"net"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// IPCIDRValidator validates that the provided string is either:
// - A valid FQDN
// - A valid IPv4 address or IPV4 CIDR
// - A valid IPv6 address or IPV6 CIDR
// But NOT an IPv6 reverse mapping address
type IPCIDRValidator struct{}

// Description returns a plain text description of the validator's behavior.
func (v IPCIDRValidator) Description(ctx context.Context) string {
	return "String must be a valid FQDN, IPv4, IPv4 CIDR, IPv6, or IPv6 CIDR (IPv6 reverse mapping not allowed)"
}

// MarkdownDescription returns a markdown formatted description of the validator's behavior.
func (v IPCIDRValidator) MarkdownDescription(ctx context.Context) string {
	return "String must be a valid FQDN, IPv4, IPv4 CIDR, IPv6, or IPv6 CIDR (IPv6 reverse mapping not allowed)"
}

// Validate performs the validation.
func (v IPCIDRValidator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	value := request.ConfigValue.ValueString()

	// Check if the value is an IPv6 reverse mapping format
	if isIPv6ReverseMapping(value) {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Invalid FQDN Format",
			"The provided value appears to be in IPv6 reverse mapping format, which is not allowed.",
		)
		return
	}

	// Check if the value is a valid IPv4 or IPv6 address
	if ip := net.ParseIP(value); ip != nil {
		// Valid IP address (either IPv4 or IPv6)
		return
	}

	// Check if the value is a valid CIDR notation (IPv4 or IPv6)
	if _, _, err := net.ParseCIDR(value); err == nil {
		// Valid CIDR notation (either IPv4 or IPv6)
		return
	}
	// If we get here, it's not a valid format
	response.Diagnostics.AddAttributeError(
		request.Path,
		"Invalid Value Format",
		fmt.Sprintf("The value '%s' is not a valid IPv4, IPv4 CIDR, IPv6, or IPv6 CIDR.", value),
	)
}

// isIPv6ReverseMapping checks if the string appears to be an IPv6 reverse mapping format
// IPv6 reverse mappings typically look like: x.y.z.w.ip6.arpa
func isIPv6ReverseMapping(s string) bool {
	return strings.HasSuffix(s, ".ip6.arpa") || strings.HasSuffix(s, ".ip6.arpa.")
}

// IsValidIPCIDR returns a validator that ensures the input is a valid FQDN, IPv4, IPv6 address, or CIDR notation.
func IsValidIPCIDR() validator.String {
	return IPCIDRValidator{}
}
