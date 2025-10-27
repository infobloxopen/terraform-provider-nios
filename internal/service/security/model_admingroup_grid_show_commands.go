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

type AdmingroupGridShowCommandsModel struct {
	ShowTestPromoteMaster types.Bool `tfsdk:"show_test_promote_master"`
	ShowToken             types.Bool `tfsdk:"show_token"`
	EnableAll             types.Bool `tfsdk:"enable_all"`
	DisableAll            types.Bool `tfsdk:"disable_all"`
	ShowDscp              types.Bool `tfsdk:"show_dscp"`
}

var AdmingroupGridShowCommandsAttrTypes = map[string]attr.Type{
	"show_test_promote_master": types.BoolType,
	"show_token":               types.BoolType,
	"enable_all":               types.BoolType,
	"disable_all":              types.BoolType,
	"show_dscp":                types.BoolType,
}

var AdmingroupGridShowCommandsResourceSchemaAttributes = map[string]schema.Attribute{
	"show_test_promote_master": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_token": schema.BoolAttribute{
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
	"show_dscp": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
}

func ExpandAdmingroupGridShowCommands(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.AdmingroupGridShowCommands {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AdmingroupGridShowCommandsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AdmingroupGridShowCommandsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.AdmingroupGridShowCommands {
	if m == nil {
		return nil
	}
	to := &security.AdmingroupGridShowCommands{
		ShowTestPromoteMaster: flex.ExpandBoolPointer(m.ShowTestPromoteMaster),
		ShowToken:             flex.ExpandBoolPointer(m.ShowToken),
		EnableAll:             flex.ExpandBoolPointer(m.EnableAll),
		DisableAll:            flex.ExpandBoolPointer(m.DisableAll),
		ShowDscp:              flex.ExpandBoolPointer(m.ShowDscp),
	}
	return to
}

func FlattenAdmingroupGridShowCommands(ctx context.Context, from *security.AdmingroupGridShowCommands, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AdmingroupGridShowCommandsAttrTypes)
	}
	m := AdmingroupGridShowCommandsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AdmingroupGridShowCommandsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AdmingroupGridShowCommandsModel) Flatten(ctx context.Context, from *security.AdmingroupGridShowCommands, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AdmingroupGridShowCommandsModel{}
	}
	m.ShowTestPromoteMaster = types.BoolPointerValue(from.ShowTestPromoteMaster)
	m.ShowToken = types.BoolPointerValue(from.ShowToken)
	m.EnableAll = types.BoolPointerValue(from.EnableAll)
	m.DisableAll = types.BoolPointerValue(from.DisableAll)
	m.ShowDscp = types.BoolPointerValue(from.ShowDscp)
}
