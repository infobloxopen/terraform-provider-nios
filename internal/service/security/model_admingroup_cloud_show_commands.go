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

type AdmingroupCloudShowCommandsModel struct {
	ShowCloudServicesPortal types.Bool `tfsdk:"show_cloud_services_portal"`
	EnableAll               types.Bool `tfsdk:"enable_all"`
	DisableAll              types.Bool `tfsdk:"disable_all"`
}

var AdmingroupCloudShowCommandsAttrTypes = map[string]attr.Type{
	"show_cloud_services_portal": types.BoolType,
	"enable_all":                 types.BoolType,
	"disable_all":                types.BoolType,
}

var AdmingroupCloudShowCommandsResourceSchemaAttributes = map[string]schema.Attribute{
	"show_cloud_services_portal": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"enable_all": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then enable all fields",
	},
	"disable_all": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then disable all fields",
	},
}

func ExpandAdmingroupCloudShowCommands(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.AdmingroupCloudShowCommands {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AdmingroupCloudShowCommandsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AdmingroupCloudShowCommandsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.AdmingroupCloudShowCommands {
	if m == nil {
		return nil
	}
	to := &security.AdmingroupCloudShowCommands{
		ShowCloudServicesPortal: flex.ExpandBoolPointer(m.ShowCloudServicesPortal),
		EnableAll:               flex.ExpandBoolPointer(m.EnableAll),
		DisableAll:              flex.ExpandBoolPointer(m.DisableAll),
	}
	return to
}

func FlattenAdmingroupCloudShowCommands(ctx context.Context, from *security.AdmingroupCloudShowCommands, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AdmingroupCloudShowCommandsAttrTypes)
	}
	m := AdmingroupCloudShowCommandsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AdmingroupCloudShowCommandsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AdmingroupCloudShowCommandsModel) Flatten(ctx context.Context, from *security.AdmingroupCloudShowCommands, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AdmingroupCloudShowCommandsModel{}
	}
	m.ShowCloudServicesPortal = types.BoolPointerValue(from.ShowCloudServicesPortal)
	m.EnableAll = types.BoolPointerValue(from.EnableAll)
	m.DisableAll = types.BoolPointerValue(from.DisableAll)
}
