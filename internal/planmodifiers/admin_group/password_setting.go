package admin_group

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type passwordSettingPlanModifier struct{}

func (m passwordSettingPlanModifier) Description(ctx context.Context) string {
	return "Validates that `password_setting` can only be updated when `use_password_setting` is set to true."
}

func (m passwordSettingPlanModifier) MarkdownDescription(ctx context.Context) string {
	return "Validates that `password_setting` can only be updated when `use_password_setting` is set to true."
}

func (m passwordSettingPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	if req.StateValue.IsNull() || req.StateValue.IsUnknown() {
		return
	}
	if req.PlanValue.IsNull() || req.PlanValue.IsUnknown() {
		return
	}

	// Retrieve the current and planned values
	var currentPasswordSetting, plannedPasswordSetting attr.Value
	var usePasswordSetting attr.Value

	req.State.GetAttribute(ctx, path.Root("password_setting"), &currentPasswordSetting)
	req.Plan.GetAttribute(ctx, path.Root("password_setting"), &plannedPasswordSetting)
	req.Plan.GetAttribute(ctx, path.Root("use_password_setting"), &usePasswordSetting)

	// Check if `password_setting` has changed
	if !currentPasswordSetting.Equal(plannedPasswordSetting) {
		// Ensure `use_password_setting` is true
		if usePasswordSetting.IsNull() || usePasswordSetting.IsUnknown() || !usePasswordSetting.(types.Bool).ValueBool() {
			resp.Diagnostics.AddAttributeError(
				path.Root("password_setting"),
				"Invalid Configuration",
				"The `password_setting` field can only be updated when `use_password_setting` is set to true.",
			)
		}
	}

}

// PasswordSettingPlanModifier returns a plan modifier for `password_setting`.
func PasswordSettingPlanModifier() planmodifier.Object {
	return passwordSettingPlanModifier{}
}
