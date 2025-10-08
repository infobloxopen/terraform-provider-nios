package planmodifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

// RequiresReplaceIfRemoved returns a plan modifier that marks the resource
// for replacement when the attribute is removed/set to null in config
// while the state has a non-empty value
func RequiresReplaceIfRemoved() planmodifier.String {
	return &requiresReplaceIfRemovedModifier{}
}

type requiresReplaceIfRemovedModifier struct{}

func (m *requiresReplaceIfRemovedModifier) Description(ctx context.Context) string {
	return "Requires replacement if the attribute is removed from config while state has a value"
}

func (m *requiresReplaceIfRemovedModifier) MarkdownDescription(ctx context.Context) string {
	return "Requires replacement if the attribute is removed from config while state has a value"
}

func (m *requiresReplaceIfRemovedModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// Skip if we're creating a new resource
	if req.StateValue.IsNull() {
		return
	}

	// Skip if we're destroying the resource
	if req.PlanValue.IsNull() && req.ConfigValue.IsNull() {
		return
	}

	// Check if attribute had a non-empty value in state
	stateHasValue := !req.StateValue.IsNull() &&
		!req.StateValue.IsUnknown() &&
		req.StateValue.ValueString() != ""

	configIsNull := req.ConfigValue.IsNull()

	if stateHasValue && configIsNull {
		resp.RequiresReplace = true
	}
}
