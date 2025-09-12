package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

type useStateForUnknownBool struct{}

func UseStateForUnknownBool() planmodifier.Bool {
	return useStateForUnknownBool{}
}

func (m useStateForUnknownBool) Description(ctx context.Context) string {
	return "If the configuration is unset and the planned value is unknown, use the prior state value instead."
}

func (m useStateForUnknownBool) MarkdownDescription(ctx context.Context) string {
	return "If the configuration is unset and the planned value is unknown, use the prior state value instead."
}

func (m useStateForUnknownBool) PlanModifyBool(ctx context.Context, req planmodifier.BoolRequest, resp *planmodifier.BoolResponse) {
	if req.PlanValue.IsUnknown() && req.ConfigValue.IsNull() && !req.StateValue.IsNull() {
		resp.PlanValue = req.StateValue
	}
}

type useStateForUnknownString struct{}

func UseStateForUnknownString() planmodifier.String {
	return useStateForUnknownString{}
}

func (m useStateForUnknownString) Description(ctx context.Context) string {
	return "If the configuration is unset and the planned value is unknown, use the prior state value instead."
}

func (m useStateForUnknownString) MarkdownDescription(ctx context.Context) string {
	return "If the configuration is unset and the planned value is unknown, use the prior state value instead."
}

func (m useStateForUnknownString) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.PlanValue.IsUnknown() && req.ConfigValue.IsNull() && !req.StateValue.IsNull() {
		resp.PlanValue = req.StateValue
	}
}

type useStateForUnknownList struct{}

func UseStateForUnknownList() planmodifier.List {
	return useStateForUnknownList{}
}

func (m useStateForUnknownList) Description(ctx context.Context) string {
	return "If the configuration is unset and the planned value is unknown, use the prior state value instead."
}

func (m useStateForUnknownList) MarkdownDescription(ctx context.Context) string {
	return "If the configuration is unset and the planned value is unknown, use the prior state value instead."
}

func (m useStateForUnknownList) PlanModifyList(ctx context.Context, req planmodifier.ListRequest, resp *planmodifier.ListResponse) {
	if req.PlanValue.IsUnknown() && req.ConfigValue.IsNull() && !req.StateValue.IsNull() {
		resp.PlanValue = req.StateValue
	}
}
