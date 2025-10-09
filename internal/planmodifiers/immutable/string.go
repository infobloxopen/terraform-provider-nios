package immutable

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

var _ planmodifier.String = immutableStringModifier{}

// immutableStringModifier validates that the provided string is not mutated after resource creation.
type immutableStringModifier struct{}

func (m immutableStringModifier) Description(ctx context.Context) string {
	return "Ensures this attribute cannot be changed after resource creation"
}

func (m immutableStringModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m immutableStringModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.StateValue.IsNull() {
		return
	}

	if req.PlanValue.IsNull() {
		return
	}

	if req.StateValue.IsUnknown() || req.PlanValue.IsUnknown() {
		return
	}

	if req.StateValue.Equal(req.PlanValue) {
		return
	}

	stateVal := req.StateValue.ValueString()
	planVal := req.PlanValue.ValueString()

	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Immutable Attribute Changed",
		fmt.Sprintf("The attribute cannot be changed after creation. "+
			"Existing value: %q, Planned value: %q. "+
			"To change this value, the resource must be destroyed and recreated.",
			stateVal,
			planVal,
		),
	)
}

// ImmutableString returns a plan modifier that ensures the given string attribute cannot be changed after creation.
func ImmutableString() planmodifier.String {
	return immutableStringModifier{}
}
