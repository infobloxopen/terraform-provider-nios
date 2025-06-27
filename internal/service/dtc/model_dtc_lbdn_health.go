package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dtc"

	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/flex"
)

type DtcLbdnHealthModel struct {
	Availability types.String `tfsdk:"availability"`
	EnabledState types.String `tfsdk:"enabled_state"`
	Description  types.String `tfsdk:"description"`
}

var DtcLbdnHealthAttrTypes = map[string]attr.Type{
	"availability":  types.StringType,
	"enabled_state": types.StringType,
	"description":   types.StringType,
}

var DtcLbdnHealthResourceSchemaAttributes = map[string]schema.Attribute{
	"availability": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The availability color status.",
	},
	"enabled_state": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The enabled state of the object.",
	},
	"description": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The textual description of the object's status.",
	},
}

func ExpandDtcLbdnHealth(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dtc.DtcLbdnHealth {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DtcLbdnHealthModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DtcLbdnHealthModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dtc.DtcLbdnHealth {
	if m == nil {
		return nil
	}
	to := &dtc.DtcLbdnHealth{
		Availability: flex.ExpandStringPointer(m.Availability),
		EnabledState: flex.ExpandStringPointer(m.EnabledState),
		Description:  flex.ExpandStringPointer(m.Description),
	}
	return to
}

func FlattenDtcLbdnHealth(ctx context.Context, from *dtc.DtcLbdnHealth, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DtcLbdnHealthAttrTypes)
	}
	m := DtcLbdnHealthModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DtcLbdnHealthAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DtcLbdnHealthModel) Flatten(ctx context.Context, from *dtc.DtcLbdnHealth, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DtcLbdnHealthModel{}
	}
	m.Availability = flex.FlattenStringPointer(from.Availability)
	m.EnabledState = flex.FlattenStringPointer(from.EnabledState)
	m.Description = flex.FlattenStringPointer(from.Description)
}
