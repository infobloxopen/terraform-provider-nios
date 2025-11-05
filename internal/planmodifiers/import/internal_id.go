package planmodifiers

import (
	"context"
	"maps"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const terraformInternalIDEA = "Terraform Internal ID"

func (m *associateInternalId) PlanModifyMap(ctx context.Context, req planmodifier.MapRequest, resp *planmodifier.MapResponse) {
	associateInternalId, diags := req.Private.GetKey(ctx, "associate_internal_id")
	resp.Diagnostics.Append(diags...)
	if associateInternalId == nil {
		return
	}

	planMap := make(map[string]attr.Value)

	if !req.PlanValue.IsNull() {
		maps.Copy(planMap, req.PlanValue.Elements())
	}

	planMap[terraformInternalIDEA] = types.StringUnknown()

	newPlanValue, diags := types.MapValue(types.StringType, planMap)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.PlanValue = newPlanValue
}

type associateInternalId struct{}

func (m *associateInternalId) Description(ctx context.Context) string {
	return "Adds the Terraform Internal ID to the plan during import"
}

func (m *associateInternalId) MarkdownDescription(ctx context.Context) string {
	return "Adds the Terraform Internal ID to the plan during import"
}

func AssociateInternalId() planmodifier.Map {
	return &associateInternalId{}
}
