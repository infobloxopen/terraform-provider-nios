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
	return "Validator to check if string is in the correct IPv6 ARPA format 'x.x.x...x.ip6.arpa' with exactly 32 hex nibbles."
}

func (v ArpaIPv6Validator) MarkdownDescription(ctx context.Context) string {
	return "Validator to check if string is in the correct IPv6 ARPA format 'x.x.x...x.ip6.arpa' with exactly 32 hex nibbles."
}

func (v ArpaIPv6Validator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()
	normalized := strings.TrimSuffix(value, ".")

	if !strings.HasSuffix(normalized, ".ip6.arpa") {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid IPv6 ARPA Format",
			fmt.Sprintf("The value '%s' must end with '.ip6.arpa'", value),
		)
		return
	}

	nibbles := strings.Split(strings.TrimSuffix(normalized, ".ip6.arpa"), ".")
	if len(nibbles) != 32 {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid IPv6 ARPA Format",
			fmt.Sprintf("The value '%s' must contain exactly 32 hexadecimal nibbles before '.ip6.arpa'", value),
		)
		return
	}

	hexNibble := regexp.MustCompile(`^[0-9a-fA-F]$`)
	for _, nibble := range nibbles {
		if !hexNibble.MatchString(nibble) {
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
