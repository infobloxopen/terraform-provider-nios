package plancontrol

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/infobloxopen/terraform-provider-nios/config"
)

func UseStateForUnknownInt64() planmodifier.Int64 {
	return useStateForUnknownInt64{}
}

// useStateForUnknownInt64 implements the plan modifier.
type useStateForUnknownInt64 struct{}

// Description returns a human-readable description of the plan modifier.
func (m useStateForUnknownInt64) Description(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m useStateForUnknownInt64) MarkdownDescription(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// PlanModifyInt64 implements the plan modification logic.
func (m useStateForUnknownInt64) PlanModifyInt64(_ context.Context, req planmodifier.Int64Request, resp *planmodifier.Int64Response) {
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
