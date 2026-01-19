package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type AdmingroupMachineControlToplevelCommandsModel struct {
	Reboot     types.Bool `tfsdk:"reboot"`
	Reset      types.Bool `tfsdk:"reset"`
	Shutdown   types.Bool `tfsdk:"shutdown"`
	Restart    types.Bool `tfsdk:"restart"`
	EnableAll  types.Bool `tfsdk:"enable_all"`
	DisableAll types.Bool `tfsdk:"disable_all"`
}

var AdmingroupMachineControlToplevelCommandsAttrTypes = map[string]attr.Type{
	"reboot":      types.BoolType,
	"reset":       types.BoolType,
	"shutdown":    types.BoolType,
	"restart":     types.BoolType,
	"enable_all":  types.BoolType,
	"disable_all": types.BoolType,
}

var AdmingroupMachineControlToplevelCommandsResourceSchemaAttributes = map[string]schema.Attribute{
	"reboot": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"reset": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"shutdown": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"restart": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"enable_all": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "If True then enable all fields",
	},
	"disable_all": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "If True then disable all fields",
	},
}

func ExpandAdmingroupMachineControlToplevelCommands(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.AdmingroupMachineControlToplevelCommands {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AdmingroupMachineControlToplevelCommandsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AdmingroupMachineControlToplevelCommandsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.AdmingroupMachineControlToplevelCommands {
	if m == nil {
		return nil
	}
	to := &security.AdmingroupMachineControlToplevelCommands{
		Reboot:   flex.ExpandBoolPointer(m.Reboot),
		Reset:    flex.ExpandBoolPointer(m.Reset),
		Shutdown: flex.ExpandBoolPointer(m.Shutdown),
		Restart:  flex.ExpandBoolPointer(m.Restart),
	}
	return to
}

func FlattenAdmingroupMachineControlToplevelCommands(ctx context.Context, from *security.AdmingroupMachineControlToplevelCommands, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AdmingroupMachineControlToplevelCommandsAttrTypes)
	}
	m := AdmingroupMachineControlToplevelCommandsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AdmingroupMachineControlToplevelCommandsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AdmingroupMachineControlToplevelCommandsModel) Flatten(ctx context.Context, from *security.AdmingroupMachineControlToplevelCommands, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AdmingroupMachineControlToplevelCommandsModel{}
	}
	m.Reboot = types.BoolPointerValue(from.Reboot)
	m.Reset = types.BoolPointerValue(from.Reset)
	m.Shutdown = types.BoolPointerValue(from.Shutdown)
	m.Restart = types.BoolPointerValue(from.Restart)
	m.EnableAll = types.BoolPointerValue(from.EnableAll)
	m.DisableAll = types.BoolPointerValue(from.DisableAll)
}
