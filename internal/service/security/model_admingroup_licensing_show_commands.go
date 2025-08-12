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

type AdmingroupLicensingShowCommandsModel struct {
	ShowLicense              types.Bool `tfsdk:"show_license"`
	ShowLicensePoolContainer types.Bool `tfsdk:"show_license_pool_container"`
	ShowLicenseUid           types.Bool `tfsdk:"show_license_uid"`
	EnableAll                types.Bool `tfsdk:"enable_all"`
	DisableAll               types.Bool `tfsdk:"disable_all"`
}

var AdmingroupLicensingShowCommandsAttrTypes = map[string]attr.Type{
	"show_license":                types.BoolType,
	"show_license_pool_container": types.BoolType,
	"show_license_uid":            types.BoolType,
	"enable_all":                  types.BoolType,
	"disable_all":                 types.BoolType,
}

var AdmingroupLicensingShowCommandsResourceSchemaAttributes = map[string]schema.Attribute{
	"show_license": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_license_pool_container": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_license_uid": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"enable_all": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If True then enable all fields",
	},
	"disable_all": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If True then disable all fields",
	},
}

func ExpandAdmingroupLicensingShowCommands(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.AdmingroupLicensingShowCommands {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AdmingroupLicensingShowCommandsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AdmingroupLicensingShowCommandsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.AdmingroupLicensingShowCommands {
	if m == nil {
		return nil
	}
	to := &security.AdmingroupLicensingShowCommands{
		ShowLicense:              flex.ExpandBoolPointer(m.ShowLicense),
		ShowLicensePoolContainer: flex.ExpandBoolPointer(m.ShowLicensePoolContainer),
		ShowLicenseUid:           flex.ExpandBoolPointer(m.ShowLicenseUid),
		EnableAll:                flex.ExpandBoolPointer(m.EnableAll),
		DisableAll:               flex.ExpandBoolPointer(m.DisableAll),
	}
	return to
}

func FlattenAdmingroupLicensingShowCommands(ctx context.Context, from *security.AdmingroupLicensingShowCommands, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AdmingroupLicensingShowCommandsAttrTypes)
	}
	m := AdmingroupLicensingShowCommandsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AdmingroupLicensingShowCommandsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AdmingroupLicensingShowCommandsModel) Flatten(ctx context.Context, from *security.AdmingroupLicensingShowCommands, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AdmingroupLicensingShowCommandsModel{}
	}
	m.ShowLicense = types.BoolPointerValue(from.ShowLicense)
	m.ShowLicensePoolContainer = types.BoolPointerValue(from.ShowLicensePoolContainer)
	m.ShowLicenseUid = types.BoolPointerValue(from.ShowLicenseUid)
	m.EnableAll = types.BoolPointerValue(from.EnableAll)
	m.DisableAll = types.BoolPointerValue(from.DisableAll)
}
