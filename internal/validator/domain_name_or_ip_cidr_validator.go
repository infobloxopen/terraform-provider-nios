package validator

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// DomainNameOrIPCIDRValidator validates that the provided string is either:
// - A valid IPv4 address or IPv4 CIDR
// - A valid IPv6 address or IPv6 CIDR
// - A valid FQDN

type DomainNameOrIPCIDRValidator struct{}

// Description returns a plain text description of the validator's behavior.
func (v DomainNameOrIPCIDRValidator) Description(ctx context.Context) string {
	return "String must be a valid FQDN, IPv4, IPv4 CIDR, IPv6, or IPv6 CIDR"
}

// MarkdownDescription returns a markdown formatted description of the validator's behavior.
func (v DomainNameOrIPCIDRValidator) MarkdownDescription(ctx context.Context) string {
	return "String must be a valid FQDN, IPv4, IPv4 CIDR, IPv6, or IPv6 CIDR"
}

// ValidateString performs the validation.
func (v DomainNameOrIPCIDRValidator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	// First, check if it's a valid IP/CIDR using the existing validator
	ipCidrValidator := IsValidIPCIDR()
	ipCidrResponse := &validator.StringResponse{}
	ipCidrValidator.ValidateString(ctx, request, ipCidrResponse)

	// If IP/CIDR validation passes (no diagnostics) return
	if !ipCidrResponse.Diagnostics.HasError() {
		return
	}

	// If IP/CIDR validation fails, check if it's a valid FQDN using the existing validator
	domainValidator := IsValidDomainName()
	domainResponse := &validator.StringResponse{}
	domainValidator.ValidateString(ctx, request, domainResponse)

	// If domain validation also fails, show combined error
	if domainResponse.Diagnostics.HasError() {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Invalid Value Format",
			fmt.Sprintf("Not a valid IP/CIDR or FQDN: %s", domainResponse.Diagnostics.Errors()[0].Detail()),
		)
		return
	}

	// Valid FQDN
}

// IsValidDomainNameOrIPCIDR returns a validator that ensures the input is a valid IPv4, IPv6 address, CIDR notation, or FQDN.
func IsValidDomainNameOrIPCIDR() validator.String {
	return DomainNameOrIPCIDRValidator{}
}
