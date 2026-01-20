package plancontrol

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/infobloxopen/terraform-provider-nios/config"
)

func UseStateForUnknownBool() planmodifier.Bool {
	return useStateForUnknownBool{}
}

// useStateForUnknownBool implements the plan modifier.
type useStateForUnknownBool struct{}

// Description returns a human-readable description of the plan modifier.
func (m useStateForUnknownBool) Description(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m useStateForUnknownBool) MarkdownDescription(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// PlanModifyBool implements the plan modification logic.
func (m useStateForUnknownBool) PlanModifyBool(_ context.Context, req planmodifier.BoolRequest, resp *planmodifier.BoolResponse) {
	// Check if suppress computed plan is enabled
	if !config.GetSuppressComputedPlan() {
		return
	}

	// Do nothing if there is no state (resource is being created).
	if req.State.Raw.IsNull() {
		return
	}

	// Do nothing if there is a known planned value.
	if !req.PlanValue.IsUnknown() {
		return
	}

	// Do nothing if there is an unknown configuration value, otherwise interpolation gets messed up.
	if req.ConfigValue.IsUnknown() {
		return
	}

	resp.PlanValue = req.StateValue
}
