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

type MemberTrapNotificationsModel struct {
	TrapType    types.String `tfsdk:"trap_type"`
	EnableEmail types.Bool   `tfsdk:"enable_email"`
	EnableTrap  types.Bool   `tfsdk:"enable_trap"`
}

var MemberTrapNotificationsAttrTypes = map[string]attr.Type{
	"trap_type":    types.StringType,
	"enable_email": types.BoolType,
	"enable_trap":  types.BoolType,
}

var MemberTrapNotificationsResourceSchemaAttributes = map[string]schema.Attribute{
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

func ExpandMemberTrapNotifications(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberTrapNotifications {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberTrapNotificationsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberTrapNotificationsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberTrapNotifications {
	if m == nil {
		return nil
	}
	to := &grid.MemberTrapNotifications{
		TrapType:    flex.ExpandStringPointer(m.TrapType),
		EnableEmail: flex.ExpandBoolPointer(m.EnableEmail),
		EnableTrap:  flex.ExpandBoolPointer(m.EnableTrap),
	}
	return to
}

func FlattenMemberTrapNotifications(ctx context.Context, from *grid.MemberTrapNotifications, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberTrapNotificationsAttrTypes)
	}
	m := MemberTrapNotificationsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberTrapNotificationsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberTrapNotificationsModel) Flatten(ctx context.Context, from *grid.MemberTrapNotifications, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberTrapNotificationsModel{}
	}
	m.TrapType = flex.FlattenStringPointer(from.TrapType)
	m.EnableEmail = types.BoolPointerValue(from.EnableEmail)
	m.EnableTrap = types.BoolPointerValue(from.EnableTrap)
}
