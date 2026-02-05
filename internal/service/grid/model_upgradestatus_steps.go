package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type UpgradestatusStepsModel struct {
	StatusValue types.String `tfsdk:"status_value"`
	StatusText  types.String `tfsdk:"status_text"`
}

var UpgradestatusStepsAttrTypes = map[string]attr.Type{
	"status_value": types.StringType,
	"status_text":  types.StringType,
}

var UpgradestatusStepsResourceSchemaAttributes = map[string]schema.Attribute{
	"status_value": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The status value of a step.",
	},
	"status_text": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The status text that describes a step.",
	},
}

func ExpandUpgradestatusSteps(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.UpgradestatusSteps {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m UpgradestatusStepsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *UpgradestatusStepsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.UpgradestatusSteps {
	if m == nil {
		return nil
	}
	to := &grid.UpgradestatusSteps{
		StatusValue: flex.ExpandStringPointer(m.StatusValue),
		StatusText:  flex.ExpandStringPointer(m.StatusText),
	}
	return to
}

func FlattenUpgradestatusSteps(ctx context.Context, from *grid.UpgradestatusSteps, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(UpgradestatusStepsAttrTypes)
	}
	m := UpgradestatusStepsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, UpgradestatusStepsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *UpgradestatusStepsModel) Flatten(ctx context.Context, from *grid.UpgradestatusSteps, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = UpgradestatusStepsModel{}
	}
	m.StatusValue = flex.FlattenStringPointer(from.StatusValue)
	m.StatusText = flex.FlattenStringPointer(from.StatusText)
}
