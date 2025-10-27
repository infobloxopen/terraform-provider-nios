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

type AdmingroupDockerShowCommandsModel struct {
	ShowDockerBridge types.Bool `tfsdk:"show_docker_bridge"`
	EnableAll        types.Bool `tfsdk:"enable_all"`
	DisableAll       types.Bool `tfsdk:"disable_all"`
}

var AdmingroupDockerShowCommandsAttrTypes = map[string]attr.Type{
	"show_docker_bridge": types.BoolType,
	"enable_all":         types.BoolType,
	"disable_all":        types.BoolType,
}

var AdmingroupDockerShowCommandsResourceSchemaAttributes = map[string]schema.Attribute{
	"show_docker_bridge": schema.BoolAttribute{
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

func ExpandAdmingroupDockerShowCommands(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.AdmingroupDockerShowCommands {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AdmingroupDockerShowCommandsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AdmingroupDockerShowCommandsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.AdmingroupDockerShowCommands {
	if m == nil {
		return nil
	}
	to := &security.AdmingroupDockerShowCommands{
		ShowDockerBridge: flex.ExpandBoolPointer(m.ShowDockerBridge),
		EnableAll:        flex.ExpandBoolPointer(m.EnableAll),
		DisableAll:       flex.ExpandBoolPointer(m.DisableAll),
	}
	return to
}

func FlattenAdmingroupDockerShowCommands(ctx context.Context, from *security.AdmingroupDockerShowCommands, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AdmingroupDockerShowCommandsAttrTypes)
	}
	m := AdmingroupDockerShowCommandsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AdmingroupDockerShowCommandsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AdmingroupDockerShowCommandsModel) Flatten(ctx context.Context, from *security.AdmingroupDockerShowCommands, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AdmingroupDockerShowCommandsModel{}
	}
	m.ShowDockerBridge = types.BoolPointerValue(from.ShowDockerBridge)
	m.EnableAll = types.BoolPointerValue(from.EnableAll)
	m.DisableAll = types.BoolPointerValue(from.DisableAll)
}
