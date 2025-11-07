package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type UserprofileModel struct {
	Ref                    types.String `tfsdk:"ref"`
	ActiveDashboardType    types.String `tfsdk:"active_dashboard_type"`
	AdminGroup             types.String `tfsdk:"admin_group"`
	DaysToExpire           types.Int64  `tfsdk:"days_to_expire"`
	Email                  types.String `tfsdk:"email"`
	GlobalSearchOnEa       types.Bool   `tfsdk:"global_search_on_ea"`
	GlobalSearchOnNiData   types.Bool   `tfsdk:"global_search_on_ni_data"`
	GridAdminGroups        types.List   `tfsdk:"grid_admin_groups"`
	LastLogin              types.Int64  `tfsdk:"last_login"`
	LbTreeNodesAtGenLevel  types.Int64  `tfsdk:"lb_tree_nodes_at_gen_level"`
	LbTreeNodesAtLastLevel types.Int64  `tfsdk:"lb_tree_nodes_at_last_level"`
	MaxCountWidgets        types.Int64  `tfsdk:"max_count_widgets"`
	Name                   types.String `tfsdk:"name"`
	OldPassword            types.String `tfsdk:"old_password"`
	Password               types.String `tfsdk:"password"`
	TableSize              types.Int64  `tfsdk:"table_size"`
	TimeZone               types.String `tfsdk:"time_zone"`
	UseTimeZone            types.Bool   `tfsdk:"use_time_zone"`
	UserType               types.String `tfsdk:"user_type"`
}

var UserprofileAttrTypes = map[string]attr.Type{
	"ref":                         types.StringType,
	"active_dashboard_type":       types.StringType,
	"admin_group":                 types.StringType,
	"days_to_expire":              types.Int64Type,
	"email":                       types.StringType,
	"global_search_on_ea":         types.BoolType,
	"global_search_on_ni_data":    types.BoolType,
	"grid_admin_groups":           types.ListType{ElemType: types.StringType},
	"last_login":                  types.Int64Type,
	"lb_tree_nodes_at_gen_level":  types.Int64Type,
	"lb_tree_nodes_at_last_level": types.Int64Type,
	"max_count_widgets":           types.Int64Type,
	"name":                        types.StringType,
	"old_password":                types.StringType,
	"password":                    types.StringType,
	"table_size":                  types.Int64Type,
	"time_zone":                   types.StringType,
	"use_time_zone":               types.BoolType,
	"user_type":                   types.StringType,
}

var UserprofileResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"active_dashboard_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Determines the active dashboard type.",
	},
	"admin_group": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The Admin Group object to which the admin belongs. An admin user can belong to only one admin group at a time.",
	},
	"days_to_expire": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of days left before the admin's password expires.",
	},
	"email": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The email address of the admin.",
	},
	"global_search_on_ea": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if extensible attribute values will be returned by global search or not.",
	},
	"global_search_on_ni_data": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if global search will search for network insight devices and interfaces or not.",
	},
	"grid_admin_groups": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Computed:            true,
		MarkdownDescription: "List of Admin Group objects that the current user is mapped to.",
	},
	"last_login": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The timestamp when the admin last logged in.",
	},
	"lb_tree_nodes_at_gen_level": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Determines how many nodes are displayed at generation levels.",
	},
	"lb_tree_nodes_at_last_level": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Determines how many nodes are displayed at the last level.",
	},
	"max_count_widgets": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The maximum count of widgets that can be added to one dashboard.",
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The admin name.",
	},
	"old_password": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The current password that will be replaced by a new password. To change a password in the database, you must provide both the current and new password values. This is a write-only attribute.",
	},
	"password": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The new password of the admin. To change a password in the database, you must provide both the current and new password values. This is a write-only attribute.",
	},
	"table_size": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The number of lines of data a table or a single list view can contain.",
	},
	"time_zone": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The time zone of the admin user.",
	},
	"use_time_zone": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: time_zone",
	},
	"user_type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The admin type.",
	},
}

func ExpandUserprofile(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.Userprofile {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m UserprofileModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *UserprofileModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.Userprofile {
	if m == nil {
		return nil
	}
	to := &security.Userprofile{
		Ref:                    flex.ExpandStringPointer(m.Ref),
		ActiveDashboardType:    flex.ExpandStringPointer(m.ActiveDashboardType),
		Email:                  flex.ExpandStringPointer(m.Email),
		GlobalSearchOnEa:       flex.ExpandBoolPointer(m.GlobalSearchOnEa),
		GlobalSearchOnNiData:   flex.ExpandBoolPointer(m.GlobalSearchOnNiData),
		LbTreeNodesAtGenLevel:  flex.ExpandInt64Pointer(m.LbTreeNodesAtGenLevel),
		LbTreeNodesAtLastLevel: flex.ExpandInt64Pointer(m.LbTreeNodesAtLastLevel),
		MaxCountWidgets:        flex.ExpandInt64Pointer(m.MaxCountWidgets),
		OldPassword:            flex.ExpandStringPointer(m.OldPassword),
		Password:               flex.ExpandStringPointer(m.Password),
		TableSize:              flex.ExpandInt64Pointer(m.TableSize),
		TimeZone:               flex.ExpandStringPointer(m.TimeZone),
		UseTimeZone:            flex.ExpandBoolPointer(m.UseTimeZone),
	}
	return to
}

func FlattenUserprofile(ctx context.Context, from *security.Userprofile, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(UserprofileAttrTypes)
	}
	m := UserprofileModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, UserprofileAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *UserprofileModel) Flatten(ctx context.Context, from *security.Userprofile, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = UserprofileModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.ActiveDashboardType = flex.FlattenStringPointer(from.ActiveDashboardType)
	m.AdminGroup = flex.FlattenStringPointer(from.AdminGroup)
	m.DaysToExpire = flex.FlattenInt64Pointer(from.DaysToExpire)
	m.Email = flex.FlattenStringPointer(from.Email)
	m.GlobalSearchOnEa = types.BoolPointerValue(from.GlobalSearchOnEa)
	m.GlobalSearchOnNiData = types.BoolPointerValue(from.GlobalSearchOnNiData)
	m.GridAdminGroups = flex.FlattenFrameworkListString(ctx, from.GridAdminGroups, diags)
	m.LastLogin = flex.FlattenInt64Pointer(from.LastLogin)
	m.LbTreeNodesAtGenLevel = flex.FlattenInt64Pointer(from.LbTreeNodesAtGenLevel)
	m.LbTreeNodesAtLastLevel = flex.FlattenInt64Pointer(from.LbTreeNodesAtLastLevel)
	m.MaxCountWidgets = flex.FlattenInt64Pointer(from.MaxCountWidgets)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.OldPassword = flex.FlattenStringPointer(from.OldPassword)
	m.Password = flex.FlattenStringPointer(from.Password)
	m.TableSize = flex.FlattenInt64Pointer(from.TableSize)
	m.TimeZone = flex.FlattenStringPointer(from.TimeZone)
	m.UseTimeZone = types.BoolPointerValue(from.UseTimeZone)
	m.UserType = flex.FlattenStringPointer(from.UserType)
}
