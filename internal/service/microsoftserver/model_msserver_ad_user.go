package microsoftserver

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/microsoftserver"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type MsserverAdUserModel struct {
	LoginName                  types.String `tfsdk:"login_name"`
	LoginPassword              types.String `tfsdk:"login_password"`
	EnableUserSync             types.Bool   `tfsdk:"enable_user_sync"`
	SynchronizationInterval    types.Int64  `tfsdk:"synchronization_interval"`
	LastSyncTime               types.Int64  `tfsdk:"last_sync_time"`
	LastSyncStatus             types.String `tfsdk:"last_sync_status"`
	LastSyncDetail             types.String `tfsdk:"last_sync_detail"`
	LastSuccessSyncTime        types.Int64  `tfsdk:"last_success_sync_time"`
	UseLogin                   types.Bool   `tfsdk:"use_login"`
	UseEnableAdUserSync        types.Bool   `tfsdk:"use_enable_ad_user_sync"`
	UseSynchronizationMinDelay types.Bool   `tfsdk:"use_synchronization_min_delay"`
	UseEnableUserSync          types.Bool   `tfsdk:"use_enable_user_sync"`
	UseSynchronizationInterval types.Bool   `tfsdk:"use_synchronization_interval"`
}

var MsserverAdUserAttrTypes = map[string]attr.Type{
	"login_name":                    types.StringType,
	"login_password":                types.StringType,
	"enable_user_sync":              types.BoolType,
	"synchronization_interval":      types.Int64Type,
	"last_sync_time":                types.Int64Type,
	"last_sync_status":              types.StringType,
	"last_sync_detail":              types.StringType,
	"last_success_sync_time":        types.Int64Type,
	"use_login":                     types.BoolType,
	"use_enable_ad_user_sync":       types.BoolType,
	"use_synchronization_min_delay": types.BoolType,
	"use_enable_user_sync":          types.BoolType,
	"use_synchronization_interval":  types.BoolType,
}

var MsserverAdUserResourceSchemaAttributes = map[string]schema.Attribute{
	"login_name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The login name of the Microsoft Server.",
	},
	"login_password": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The login password of the DHCP Microsoft Server.",
	},
	"enable_user_sync": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the Active Directory user synchronization is enabled or not.",
	},
	"synchronization_interval": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The minimum number of minutes between two synchronizations.",
	},
	"last_sync_time": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Timestamp of the last synchronization attempt.",
	},
	"last_sync_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The status of the last synchronization attempt.",
	},
	"last_sync_detail": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The detailed status of the last synchronization attempt.",
	},
	"last_success_sync_time": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Timestamp of the last successful synchronization attempt.",
	},
	"use_login": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Flag to override login name and password from MS server",
	},
	"use_enable_ad_user_sync": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Flag to override AD User sync from grid level",
	},
	"use_synchronization_min_delay": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Flag to override synchronization interval from the MS Server",
	},
	"use_enable_user_sync": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: enable_user_sync",
	},
	"use_synchronization_interval": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: synchronization_interval",
	},
}

func ExpandMsserverAdUser(ctx context.Context, o types.Object, diags *diag.Diagnostics) *microsoftserver.MsserverAdUser {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MsserverAdUserModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MsserverAdUserModel) Expand(ctx context.Context, diags *diag.Diagnostics) *microsoftserver.MsserverAdUser {
	if m == nil {
		return nil
	}
	to := &microsoftserver.MsserverAdUser{
		LoginName:                  flex.ExpandStringPointer(m.LoginName),
		LoginPassword:              flex.ExpandStringPointer(m.LoginPassword),
		EnableUserSync:             flex.ExpandBoolPointer(m.EnableUserSync),
		SynchronizationInterval:    flex.ExpandInt64Pointer(m.SynchronizationInterval),
		UseLogin:                   flex.ExpandBoolPointer(m.UseLogin),
		UseEnableAdUserSync:        flex.ExpandBoolPointer(m.UseEnableAdUserSync),
		UseSynchronizationMinDelay: flex.ExpandBoolPointer(m.UseSynchronizationMinDelay),
		UseEnableUserSync:          flex.ExpandBoolPointer(m.UseEnableUserSync),
		UseSynchronizationInterval: flex.ExpandBoolPointer(m.UseSynchronizationInterval),
	}
	return to
}

func FlattenMsserverAdUser(ctx context.Context, from *microsoftserver.MsserverAdUser, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MsserverAdUserAttrTypes)
	}
	m := MsserverAdUserModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MsserverAdUserAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MsserverAdUserModel) Flatten(ctx context.Context, from *microsoftserver.MsserverAdUser, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MsserverAdUserModel{}
	}
	m.LoginName = flex.FlattenStringPointer(from.LoginName)
	m.LoginPassword = flex.FlattenStringPointer(from.LoginPassword)
	m.EnableUserSync = types.BoolPointerValue(from.EnableUserSync)
	m.SynchronizationInterval = flex.FlattenInt64Pointer(from.SynchronizationInterval)
	m.LastSyncTime = flex.FlattenInt64Pointer(from.LastSyncTime)
	m.LastSyncStatus = flex.FlattenStringPointer(from.LastSyncStatus)
	m.LastSyncDetail = flex.FlattenStringPointer(from.LastSyncDetail)
	m.LastSuccessSyncTime = flex.FlattenInt64Pointer(from.LastSuccessSyncTime)
	m.UseLogin = types.BoolPointerValue(from.UseLogin)
	m.UseEnableAdUserSync = types.BoolPointerValue(from.UseEnableAdUserSync)
	m.UseSynchronizationMinDelay = types.BoolPointerValue(from.UseSynchronizationMinDelay)
	m.UseEnableUserSync = types.BoolPointerValue(from.UseEnableUserSync)
	m.UseSynchronizationInterval = types.BoolPointerValue(from.UseSynchronizationInterval)
}
