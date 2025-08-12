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

type AdmingroupUserAccessModel struct {
	Address    types.String `tfsdk:"address"`
	Permission types.String `tfsdk:"permission"`
}

var AdmingroupUserAccessAttrTypes = map[string]attr.Type{
	"address":    types.StringType,
	"permission": types.StringType,
}

var AdmingroupUserAccessResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The address this rule applies to or \"Any\".",
	},
	"permission": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The permission to use for this address.",
	},
}

func ExpandAdmingroupUserAccess(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.AdmingroupUserAccess {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AdmingroupUserAccessModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AdmingroupUserAccessModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.AdmingroupUserAccess {
	if m == nil {
		return nil
	}
	to := &security.AdmingroupUserAccess{
		Address:    flex.ExpandStringPointer(m.Address),
		Permission: flex.ExpandStringPointer(m.Permission),
	}
	return to
}

func FlattenAdmingroupUserAccess(ctx context.Context, from *security.AdmingroupUserAccess, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AdmingroupUserAccessAttrTypes)
	}
	m := AdmingroupUserAccessModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AdmingroupUserAccessAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AdmingroupUserAccessModel) Flatten(ctx context.Context, from *security.AdmingroupUserAccess, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AdmingroupUserAccessModel{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Permission = flex.FlattenStringPointer(from.Permission)
}
