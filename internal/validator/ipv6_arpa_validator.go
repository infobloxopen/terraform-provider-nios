package validator

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// ArpaIPv6Validator validates that the string matches a valid ARPA IPv6 format
type ArpaIPv6Validator struct{}

func (v ArpaIPv6Validator) Description(ctx context.Context) string {
	return "String must be a valid ARPA IPv6 address in the format 'x.x.x...x.ip6.arpa' with 1–32 hex nibbles"
}

func (v ArpaIPv6Validator) MarkdownDescription(ctx context.Context) string {
	return "String must be a valid ARPA IPv6 address in the format `x.x.x...x.ip6.arpa` with 1–32 hex nibbles"
}

func (v ArpaIPv6Validator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()

	// Must end with .ip6.arpa
	if !strings.HasSuffix(value, ".ip6.arpa") {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid IPv6 ARPA Format",
			fmt.Sprintf("The value '%s' must end with '.ip6.arpa'", value),
		)
		return
	}

	// Strip suffix and split nibbles
	nibbles := strings.Split(strings.TrimSuffix(value, ".ip6.arpa"), ".")

	// Must be between 1 and 32 nibbles
	if len(nibbles) < 1 || len(nibbles) > 32 {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid IPv6 ARPA Format",
			fmt.Sprintf("The value '%s' must contain between 1 and 32 hexadecimal nibbles before '.ip6.arpa'", value),
		)
		return
	}

	// Validate each nibble
	for _, nibble := range nibbles {
		if !regexp.MustCompile(`^[0-9a-fA-F]$`).MatchString(nibble) {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Invalid IPv6 ARPA Format",
				fmt.Sprintf("Nibble '%s' in '%s' is not a valid hexadecimal digit", nibble, value),
			)
			return
		}
	}
}

// IsValidArpaIPv6 returns a validator that ensures the input is a valid ARPA IPv6 address.
func IsValidArpaIPv6() validator.String {
	return ArpaIPv6Validator{}
}
