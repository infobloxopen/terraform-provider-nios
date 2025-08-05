package validator

import (
    "context"
    "fmt"

    "github.com/hashicorp/terraform-plugin-framework/path"
    "github.com/hashicorp/terraform-plugin-framework/schema/validator"
    "github.com/hashicorp/terraform-plugin-framework/types"
)

// QuorumValidator validates that the quorum attribute is only set when availability is "QUORUM"
// and is required when availability is "QUORUM".
type QuorumValidator struct{}

// Description returns a human-readable description of the validator.
func (v QuorumValidator) Description(ctx context.Context) string {
    return "Validates quorum value is only set when availability is 'QUORUM' and is required when availability is 'QUORUM'"
}

// MarkdownDescription returns a markdown description of the validator.
func (v QuorumValidator) MarkdownDescription(ctx context.Context) string {
    return "Validates quorum value is only set when availability is 'QUORUM' and is required when availability is 'QUORUM'"
}

// ValidateInt64 implements the validator.Int64 interface.
func (v QuorumValidator) ValidateInt64(ctx context.Context, req validator.Int64Request, resp *validator.Int64Response) {
    // If quorum is null or unknown, we don't need to validate anything
    if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
        return
    }

    // Get the availability attribute value
    availabilityPath := path.Root("availability")
    var availability types.String
    diags := req.Config.GetAttribute(ctx, availabilityPath, &availability)
    if diags.HasError() {
        resp.Diagnostics.Append(diags...)
        return
    }

    // If availability is null or unknown, we can't validate further
    if availability.IsNull() || availability.IsUnknown() {
        return
    }

    // Quorum is set but availability is not "QUORUM"
    if availability.ValueString() != "QUORUM" {
        resp.Diagnostics.AddAttributeError(
            req.Path,
            "Invalid Attribute Combination",
            fmt.Sprintf("The quorum attribute can only be used when availability is set to 'QUORUM', but got '%s'", availability.ValueString()),
        )
        return
    }
}

// NewQuorumValidator returns a new instance of QuorumValidator.
func NewQuorumValidator() QuorumValidator {
    return QuorumValidator{}
}

// AvailabilityQuorumValidator validates that the quorum attribute is set when availability is "QUORUM".
type AvailabilityQuorumValidator struct{}

// Description returns a human-readable description of the validator.
func (v AvailabilityQuorumValidator) Description(ctx context.Context) string {
    return "Validates quorum value is set when availability is 'QUORUM'"
}

// MarkdownDescription returns a markdown description of the validator.
func (v AvailabilityQuorumValidator) MarkdownDescription(ctx context.Context) string {
    return "Validates quorum value is set when availability is 'QUORUM'"
}

// ValidateString implements the validator.String interface.
func (v AvailabilityQuorumValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
    // If availability is null, unknown, or not "QUORUM", we don't need to check for quorum
    if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() || req.ConfigValue.ValueString() != "QUORUM" {
        return
    }

    // Check if quorum is set when availability is "QUORUM"
    quorumPath := path.Root("quorum")
    var quorum types.Int64
    diags := req.Config.GetAttribute(ctx, quorumPath, &quorum)
    if diags.HasError() {
        resp.Diagnostics.Append(diags...)
        return
    }

    // If quorum is null or unknown when availability is "QUORUM", report an error
    if quorum.IsNull() || quorum.IsUnknown() {
        resp.Diagnostics.AddAttributeError(
            req.Path,
            "Missing Required Attribute",
            "When availability is set to 'QUORUM', the 'quorum' attribute must be specified",
        )
        return
    }
}

// NewAvailabilityQuorumValidator returns a new instance of AvailabilityQuorumValidator.
func NewAvailabilityQuorumValidator() AvailabilityQuorumValidator {
    return AvailabilityQuorumValidator{}
}