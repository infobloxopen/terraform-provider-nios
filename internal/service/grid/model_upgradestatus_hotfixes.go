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

type UpgradestatusHotfixesModel struct {
	StatusText types.String `tfsdk:"status_text"`
	UniqueId   types.String `tfsdk:"unique_id"`
}

var UpgradestatusHotfixesAttrTypes = map[string]attr.Type{
	"status_text": types.StringType,
	"unique_id":   types.StringType,
}

var UpgradestatusHotfixesResourceSchemaAttributes = map[string]schema.Attribute{
	"status_text": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The status text of the hotfix.",
	},
	"unique_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Unique ID of the hotfix.",
	},
}

func ExpandUpgradestatusHotfixes(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.UpgradestatusHotfixes {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m UpgradestatusHotfixesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *UpgradestatusHotfixesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.UpgradestatusHotfixes {
	if m == nil {
		return nil
	}
	to := &grid.UpgradestatusHotfixes{
		StatusText: flex.ExpandStringPointer(m.StatusText),
		UniqueId:   flex.ExpandStringPointer(m.UniqueId),
	}
	return to
}

func FlattenUpgradestatusHotfixes(ctx context.Context, from *grid.UpgradestatusHotfixes, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(UpgradestatusHotfixesAttrTypes)
	}
	m := UpgradestatusHotfixesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, UpgradestatusHotfixesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *UpgradestatusHotfixesModel) Flatten(ctx context.Context, from *grid.UpgradestatusHotfixes, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = UpgradestatusHotfixesModel{}
	}
	m.StatusText = flex.FlattenStringPointer(from.StatusText)
	m.UniqueId = flex.FlattenStringPointer(from.UniqueId)
}
