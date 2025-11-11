package planmodifiers

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (m *markUnknownIfPrivateSet) PlanModifyMap(ctx context.Context, req planmodifier.MapRequest, resp *planmodifier.MapResponse) {
	associateInternalId, diags := req.Private.GetKey(ctx, "associate_internal_id")
	resp.Diagnostics.Append(diags...)
	if associateInternalId == nil {
		return
	}

	resp.PlanValue = types.MapUnknown(types.StringType)
}

type markUnknownIfPrivateSet struct{}

func (m *markUnknownIfPrivateSet) Description(ctx context.Context) string {
	return "Marks the entire map attribute as unknown during plan when private key is set."
}

func (m *markUnknownIfPrivateSet) MarkdownDescription(ctx context.Context) string {
	return "Marks the entire map attribute as unknown during plan when private key is set."
}

func MarkUnknownIfPrivateSet() planmodifier.Map {
	return &markUnknownIfPrivateSet{}
}
