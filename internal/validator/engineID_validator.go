package validator

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = engineIDValidator{}


type engineIDValidator struct{}

func (v engineIDValidator) Description(_ context.Context) string {
    return "value must be a valid SNMP Engine ID"
}

func (v engineIDValidator) MarkdownDescription(_ context.Context) string {
    return "Value must be a valid SNMP Engine ID. " +
        "Engine ID must contain only hexadecimal digits (0-9, a-f, A-F), " +
        "must have an even number of characters, " +
        "and must be between 10 and 64 characters in length (5 to 32 octets)."
}

func (v engineIDValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
    if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
        return
    }

    engineID := req.ConfigValue.ValueString()

    // Engine ID must be between 10 and 64 hexadecimal digits (5 to 32 octets)
    if len(engineID) < 10 || len(engineID) > 64 {
        resp.Diagnostics.AddAttributeError(
            req.Path,
            "Invalid Engine ID",
            "Engine ID must contain between 10 and 64 hexadecimal digits (5 to 32 octets).",
        )
        return
    }

    // Engine ID must contain only hexadecimal digits
    if !isValidHexadecimal(engineID) {
        resp.Diagnostics.AddAttributeError(
            req.Path,
            "Invalid Engine ID Format",
            "Engine ID must contain only hexadecimal digits (0-9, a-f, A-F).",
        )
        return
    }

    // Number of characters must be even (each octet is represented by 2 hex digits)
    if len(engineID)%2 != 0 {
        resp.Diagnostics.AddAttributeError(
            req.Path,
            "Invalid Engine ID Format",
            "Engine ID must contain an even number of hexadecimal digits.",
        )
        return
    }
}

func EngineIDValidator() validator.String {
    return engineIDValidator{}
}
