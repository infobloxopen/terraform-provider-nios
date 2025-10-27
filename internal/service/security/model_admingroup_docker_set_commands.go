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

type AdmingroupDockerSetCommandsModel struct {
	SetDockerBridge types.Bool `tfsdk:"set_docker_bridge"`
	EnableAll       types.Bool `tfsdk:"enable_all"`
	DisableAll      types.Bool `tfsdk:"disable_all"`
}

var AdmingroupDockerSetCommandsAttrTypes = map[string]attr.Type{
	"set_docker_bridge": types.BoolType,
	"enable_all":        types.BoolType,
	"disable_all":       types.BoolType,
}

var AdmingroupDockerSetCommandsResourceSchemaAttributes = map[string]schema.Attribute{
	"set_docker_bridge": schema.BoolAttribute{
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

func ExpandAdmingroupDockerSetCommands(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.AdmingroupDockerSetCommands {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AdmingroupDockerSetCommandsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AdmingroupDockerSetCommandsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.AdmingroupDockerSetCommands {
	if m == nil {
		return nil
	}
	to := &security.AdmingroupDockerSetCommands{
		SetDockerBridge: flex.ExpandBoolPointer(m.SetDockerBridge),
		EnableAll:       flex.ExpandBoolPointer(m.EnableAll),
		DisableAll:      flex.ExpandBoolPointer(m.DisableAll),
	}
	return to
}

func FlattenAdmingroupDockerSetCommands(ctx context.Context, from *security.AdmingroupDockerSetCommands, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AdmingroupDockerSetCommandsAttrTypes)
	}
	m := AdmingroupDockerSetCommandsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AdmingroupDockerSetCommandsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AdmingroupDockerSetCommandsModel) Flatten(ctx context.Context, from *security.AdmingroupDockerSetCommands, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AdmingroupDockerSetCommandsModel{}
	}
	m.SetDockerBridge = types.BoolPointerValue(from.SetDockerBridge)
	m.EnableAll = types.BoolPointerValue(from.EnableAll)
	m.DisableAll = types.BoolPointerValue(from.DisableAll)
}
