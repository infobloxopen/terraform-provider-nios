package validator

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type useFlagMustBeTrueValidator struct {
	useFlagPath path.Expression
}

// UseFlagMustBeTrue returns a validator that checks if the referenced
// boolean "use flag" attribute is set to true when the current attribute is configured.
func UseFlagMustBeTrue(useFlagPath path.Expression) validator.Bool {
	return useFlagMustBeTrueValidator{useFlagPath: useFlagPath}
}

func (v useFlagMustBeTrueValidator) Description(_ context.Context) string {
	return fmt.Sprintf("If this attribute is set, %s must be set to true.", v.useFlagPath)
}

func (v useFlagMustBeTrueValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v useFlagMustBeTrueValidator) ValidateBool(ctx context.Context, req validator.BoolRequest, resp *validator.BoolResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	var useFlag types.Bool
	diags := req.Config.GetAttribute(ctx, path.Root(v.useFlagPath.String()), &useFlag)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if useFlag.IsUnknown() || useFlag.ValueBool() {
		return
	}

	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Missing required use flag",
		fmt.Sprintf(
			"Attribute %q is set but %q is not set to true. Set %q to true to use this attribute.",
			req.Path, v.useFlagPath, v.useFlagPath,
		),
	)
}
