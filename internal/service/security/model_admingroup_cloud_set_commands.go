package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type AdmingroupCloudSetCommandsModel struct {
	SetCloudServicesPortalForceRefresh types.Bool `tfsdk:"set_cloud_services_portal_forcerefresh"`
	SetCloudServicesPortal             types.Bool `tfsdk:"set_cloud_services_portal"`
	EnableAll                          types.Bool `tfsdk:"enable_all"`
	DisableAll                         types.Bool `tfsdk:"disable_all"`
}

var AdmingroupCloudSetCommandsAttrTypes = map[string]attr.Type{
	"set_cloud_services_portal_forcerefresh": types.BoolType,
	"set_cloud_services_portal":              types.BoolType,
	"enable_all":                             types.BoolType,
	"disable_all":                            types.BoolType,
}

var AdmingroupCloudSetCommandsResourceSchemaAttributes = map[string]schema.Attribute{
	"set_cloud_services_portal_forcerefresh": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_cloud_services_portal": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"enable_all": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then enable all fields",
	},
	"disable_all": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then disable all fields",
	},
}

func ExpandAdmingroupCloudSetCommands(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.AdmingroupCloudSetCommands {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AdmingroupCloudSetCommandsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AdmingroupCloudSetCommandsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.AdmingroupCloudSetCommands {
	if m == nil {
		return nil
	}
	to := &security.AdmingroupCloudSetCommands{
		SetCloudServicesPortalForceRefresh: flex.ExpandBoolPointer(m.SetCloudServicesPortalForceRefresh),
		SetCloudServicesPortal:             flex.ExpandBoolPointer(m.SetCloudServicesPortal),
		EnableAll:                          flex.ExpandBoolPointer(m.EnableAll),
		DisableAll:                         flex.ExpandBoolPointer(m.DisableAll),
	}
	return to
}

func FlattenAdmingroupCloudSetCommands(ctx context.Context, from *security.AdmingroupCloudSetCommands, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AdmingroupCloudSetCommandsAttrTypes)
	}
	m := AdmingroupCloudSetCommandsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AdmingroupCloudSetCommandsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AdmingroupCloudSetCommandsModel) Flatten(ctx context.Context, from *security.AdmingroupCloudSetCommands, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AdmingroupCloudSetCommandsModel{}
	}
	m.SetCloudServicesPortalForceRefresh = types.BoolPointerValue(from.SetCloudServicesPortalForceRefresh)
	m.SetCloudServicesPortal = types.BoolPointerValue(from.SetCloudServicesPortal)
	m.EnableAll = types.BoolPointerValue(from.EnableAll)
	m.DisableAll = types.BoolPointerValue(from.DisableAll)
}
