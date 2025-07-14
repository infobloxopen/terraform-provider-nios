package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dtc"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type DtcServerHealthModel struct {
	Availability types.String `tfsdk:"availability"`
	EnabledState types.String `tfsdk:"enabled_state"`
	Description  types.String `tfsdk:"description"`
}

var DtcServerHealthAttrTypes = map[string]attr.Type{
	"availability":  types.StringType,
	"enabled_state": types.StringType,
	"description":   types.StringType,
}

var DtcServerHealthResourceSchemaAttributes = map[string]schema.Attribute{
	"availability": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The availability color status.",
	},
	"enabled_state": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The enabled state of the object.",
	},
	"description": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The textual description of the object's status.",
	},
}

func ExpandDtcServerHealth(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dtc.DtcServerHealth {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DtcServerHealthModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DtcServerHealthModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dtc.DtcServerHealth {
	if m == nil {
		return nil
	}
	to := &dtc.DtcServerHealth{
		Availability: flex.ExpandStringPointer(m.Availability),
		EnabledState: flex.ExpandStringPointer(m.EnabledState),
		Description:  flex.ExpandStringPointer(m.Description),
	}
	return to
}

func FlattenDtcServerHealth(ctx context.Context, from *dtc.DtcServerHealth, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DtcServerHealthAttrTypes)
	}
	m := DtcServerHealthModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DtcServerHealthAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DtcServerHealthModel) Flatten(ctx context.Context, from *dtc.DtcServerHealth, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DtcServerHealthModel{}
	}
	m.Availability = flex.FlattenStringPointer(from.Availability)
	m.EnabledState = flex.FlattenStringPointer(from.EnabledState)
	m.Description = flex.FlattenStringPointer(from.Description)
}
