package validator

import (
    "context"
    "fmt"

    "github.com/hashicorp/terraform-plugin-framework/path"
    "github.com/hashicorp/terraform-plugin-framework/schema/validator"
    "github.com/hashicorp/terraform-plugin-framework/types"
)

var _ validator.String = equalsFieldValidator{}

type equalsFieldValidator struct {
    expression path.Expression
}

func EqualsField(expression path.Expression) validator.String {
    return equalsFieldValidator{expression: expression}
}

func (v equalsFieldValidator) Description(_ context.Context) string {
    return fmt.Sprintf("Value must not be equal to the value of %q", v.expression)
}

func (v equalsFieldValidator) MarkdownDescription(ctx context.Context) string {
    return v.Description(ctx)
}

func (v equalsFieldValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
    if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
        return
    }

    var otherVal types.String
    diags := req.Config.GetAttribute(ctx, path.Root("name"), &otherVal)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    if otherVal.IsNull() || otherVal.IsUnknown() {
        return
    }

    if req.ConfigValue.ValueString() == otherVal.ValueString() {
        resp.Diagnostics.AddAttributeError(
            req.Path,
            "Invalid Attribute Value",
            fmt.Sprintf("%q cannot have the same value as %q. Both attributes resolve to %q.", req.Path, v.expression, req.ConfigValue.ValueString()),
        )
    }
}
