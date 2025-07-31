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

type DtcPoolHealthModel struct {
	Availability types.String `tfsdk:"availability"`
	EnabledState types.String `tfsdk:"enabled_state"`
	Description  types.String `tfsdk:"description"`
}

var DtcPoolHealthAttrTypes = map[string]attr.Type{
	"availability":  types.StringType,
	"enabled_state": types.StringType,
	"description":   types.StringType,
}

var DtcPoolHealthResourceSchemaAttributes = map[string]schema.Attribute{
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

func ExpandDtcPoolHealth(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dtc.DtcPoolHealth {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DtcPoolHealthModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DtcPoolHealthModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dtc.DtcPoolHealth {
	if m == nil {
		return nil
	}
	to := &dtc.DtcPoolHealth{}
	return to
}

func FlattenDtcPoolHealth(ctx context.Context, from *dtc.DtcPoolHealth, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DtcPoolHealthAttrTypes)
	}
	m := DtcPoolHealthModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DtcPoolHealthAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DtcPoolHealthModel) Flatten(ctx context.Context, from *dtc.DtcPoolHealth, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DtcPoolHealthModel{}
	}
	m.Availability = flex.FlattenStringPointer(from.Availability)
	m.EnabledState = flex.FlattenStringPointer(from.EnabledState)
	m.Description = flex.FlattenStringPointer(from.Description)
}
