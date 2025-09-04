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

type AdmingroupLicensingSetCommandsModel struct {
	SetLicense               types.Bool `tfsdk:"set_license"`
	SetReportingResetLicense types.Bool `tfsdk:"set_reporting_reset_license"`
	SetTempLicense           types.Bool `tfsdk:"set_temp_license"`
	EnableAll                types.Bool `tfsdk:"enable_all"`
	DisableAll               types.Bool `tfsdk:"disable_all"`
}

var AdmingroupLicensingSetCommandsAttrTypes = map[string]attr.Type{
	"set_license":                 types.BoolType,
	"set_reporting_reset_license": types.BoolType,
	"set_temp_license":            types.BoolType,
	"enable_all":                  types.BoolType,
	"disable_all":                 types.BoolType,
}

var AdmingroupLicensingSetCommandsResourceSchemaAttributes = map[string]schema.Attribute{
	"set_license": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_reporting_reset_license": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_temp_license": schema.BoolAttribute{
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

func ExpandAdmingroupLicensingSetCommands(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.AdmingroupLicensingSetCommands {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AdmingroupLicensingSetCommandsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AdmingroupLicensingSetCommandsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.AdmingroupLicensingSetCommands {
	if m == nil {
		return nil
	}
	to := &security.AdmingroupLicensingSetCommands{
		SetLicense:               flex.ExpandBoolPointer(m.SetLicense),
		SetReportingResetLicense: flex.ExpandBoolPointer(m.SetReportingResetLicense),
		SetTempLicense:           flex.ExpandBoolPointer(m.SetTempLicense),
		EnableAll:                flex.ExpandBoolPointer(m.EnableAll),
		DisableAll:               flex.ExpandBoolPointer(m.DisableAll),
	}
	return to
}

func FlattenAdmingroupLicensingSetCommands(ctx context.Context, from *security.AdmingroupLicensingSetCommands, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AdmingroupLicensingSetCommandsAttrTypes)
	}
	m := AdmingroupLicensingSetCommandsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AdmingroupLicensingSetCommandsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AdmingroupLicensingSetCommandsModel) Flatten(ctx context.Context, from *security.AdmingroupLicensingSetCommands, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AdmingroupLicensingSetCommandsModel{}
	}
	m.SetLicense = types.BoolPointerValue(from.SetLicense)
	m.SetReportingResetLicense = types.BoolPointerValue(from.SetReportingResetLicense)
	m.SetTempLicense = types.BoolPointerValue(from.SetTempLicense)
	m.EnableAll = types.BoolPointerValue(from.EnableAll)
	m.DisableAll = types.BoolPointerValue(from.DisableAll)
}
