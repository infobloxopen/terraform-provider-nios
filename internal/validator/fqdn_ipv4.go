package validator

import (
	"context"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type IPOrFQDNValidator struct{}

func (v IPOrFQDNValidator) Description(ctx context.Context) string {
	return "Must be a valid IPv4 or FQDN"
}

func (v IPOrFQDNValidator) MarkdownDescription(ctx context.Context) string {
	return "Must be a valid IPv4 or FQDN"
}

func (v IPOrFQDNValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	val := strings.TrimSpace(req.ConfigValue.ValueString())

	// Check if it is an IPv4 address
	if regexp.MustCompile(`^\d+\.\d+\.\d+\.\d+$`).MatchString(val) {
		// Use IP/CIDR validator
		ipValidator := IsValidIPCIDR()
		ipValidator.ValidateString(ctx, req, resp)
		return
	}

	// Otherwise, validate as FQDN
	fqdnValidator := IsValidFQDN()
	fqdnValidator.ValidateString(ctx, req, resp)
}

func IsValidIPOrFQDN() validator.String {
	return IPOrFQDNValidator{}
}
