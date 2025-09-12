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

type AdmingroupSamlSettingModel struct {
	AutoCreateUser         types.Bool `tfsdk:"auto_create_user"`
	PersistAutoCreatedUser types.Bool `tfsdk:"persist_auto_created_user"`
}

var AdmingroupSamlSettingAttrTypes = map[string]attr.Type{
	"auto_create_user":          types.BoolType,
	"persist_auto_created_user": types.BoolType,
}

var AdmingroupSamlSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"auto_create_user": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Whether or not auto create user.",
	},
	"persist_auto_created_user": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Whether or not persist auto created user after logout.",
	},
}

func ExpandAdmingroupSamlSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.AdmingroupSamlSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AdmingroupSamlSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AdmingroupSamlSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.AdmingroupSamlSetting {
	if m == nil {
		return nil
	}
	to := &security.AdmingroupSamlSetting{
		AutoCreateUser:         flex.ExpandBoolPointer(m.AutoCreateUser),
		PersistAutoCreatedUser: flex.ExpandBoolPointer(m.PersistAutoCreatedUser),
	}
	return to
}

func FlattenAdmingroupSamlSetting(ctx context.Context, from *security.AdmingroupSamlSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AdmingroupSamlSettingAttrTypes)
	}
	m := AdmingroupSamlSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AdmingroupSamlSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AdmingroupSamlSettingModel) Flatten(ctx context.Context, from *security.AdmingroupSamlSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AdmingroupSamlSettingModel{}
	}
	m.AutoCreateUser = types.BoolPointerValue(from.AutoCreateUser)
	m.PersistAutoCreatedUser = types.BoolPointerValue(from.PersistAutoCreatedUser)
}
