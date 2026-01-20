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

type GridTrapNotificationsModel struct {
	TrapType    types.String `tfsdk:"trap_type"`
	EnableEmail types.Bool   `tfsdk:"enable_email"`
	EnableTrap  types.Bool   `tfsdk:"enable_trap"`
}

var GridTrapNotificationsAttrTypes = map[string]attr.Type{
	"trap_type":    types.StringType,
	"enable_email": types.BoolType,
	"enable_trap":  types.BoolType,
}

var GridTrapNotificationsResourceSchemaAttributes = map[string]schema.Attribute{
	"trap_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Determines the type of a given trap.",
	},
	"enable_email": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the email notifications for the given trap are enabled or not.",
	},
	"enable_trap": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the trap is enabled or not.",
	},
}

func ExpandGridTrapNotifications(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridTrapNotifications {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridTrapNotificationsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridTrapNotificationsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridTrapNotifications {
	if m == nil {
		return nil
	}
	to := &grid.GridTrapNotifications{
		TrapType:    flex.ExpandStringPointer(m.TrapType),
		EnableEmail: flex.ExpandBoolPointer(m.EnableEmail),
		EnableTrap:  flex.ExpandBoolPointer(m.EnableTrap),
	}
	return to
}

func FlattenGridTrapNotifications(ctx context.Context, from *grid.GridTrapNotifications, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridTrapNotificationsAttrTypes)
	}
	m := GridTrapNotificationsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridTrapNotificationsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridTrapNotificationsModel) Flatten(ctx context.Context, from *grid.GridTrapNotifications, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridTrapNotificationsModel{}
	}
	m.TrapType = flex.FlattenStringPointer(from.TrapType)
	m.EnableEmail = types.BoolPointerValue(from.EnableEmail)
	m.EnableTrap = types.BoolPointerValue(from.EnableTrap)
}
