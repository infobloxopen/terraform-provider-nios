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

type GridsnmpsettingTrapReceiversModel struct {
	Address types.String `tfsdk:"address"`
	User    types.String `tfsdk:"user"`
	Comment types.String `tfsdk:"comment"`
}

var GridsnmpsettingTrapReceiversAttrTypes = map[string]attr.Type{
	"address": types.StringType,
	"user":    types.StringType,
	"comment": types.StringType,
}

var GridsnmpsettingTrapReceiversResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The address of the trap receiver.",
	},
	"user": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The SNMPv3 user for this trap receiver.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "A descriptive comment for this trap receiver.",
	},
}

func ExpandGridsnmpsettingTrapReceivers(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridsnmpsettingTrapReceivers {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridsnmpsettingTrapReceiversModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridsnmpsettingTrapReceiversModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridsnmpsettingTrapReceivers {
	if m == nil {
		return nil
	}
	to := &grid.GridsnmpsettingTrapReceivers{
		Address: flex.ExpandStringPointer(m.Address),
		User:    flex.ExpandStringPointer(m.User),
		Comment: flex.ExpandStringPointer(m.Comment),
	}
	return to
}

func FlattenGridsnmpsettingTrapReceivers(ctx context.Context, from *grid.GridsnmpsettingTrapReceivers, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridsnmpsettingTrapReceiversAttrTypes)
	}
	m := GridsnmpsettingTrapReceiversModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridsnmpsettingTrapReceiversAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridsnmpsettingTrapReceiversModel) Flatten(ctx context.Context, from *grid.GridsnmpsettingTrapReceivers, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridsnmpsettingTrapReceiversModel{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.User = flex.FlattenStringPointer(from.User)
	m.Comment = flex.FlattenStringPointer(from.Comment)
}
