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

type AdmingroupDatabaseShowCommandsModel struct {
	ShowNamedMaxJournalSize    types.Bool `tfsdk:"show_named_max_journal_size"`
	ShowTxnTrace               types.Bool `tfsdk:"show_txn_trace"`
	ShowDatabaseTransferStatus types.Bool `tfsdk:"show_database_transfer_status"`
	ShowBackup                 types.Bool `tfsdk:"show_backup"`
	ShowDbPh                   types.Bool `tfsdk:"show_db_ph"`
	ShowDbsize                 types.Bool `tfsdk:"show_dbsize"`
	ShowIbdbstat               types.Bool `tfsdk:"show_ibdbstat"`
	EnableAll                  types.Bool `tfsdk:"enable_all"`
	DisableAll                 types.Bool `tfsdk:"disable_all"`
}

var AdmingroupDatabaseShowCommandsAttrTypes = map[string]attr.Type{
	"show_named_max_journal_size":   types.BoolType,
	"show_txn_trace":                types.BoolType,
	"show_database_transfer_status": types.BoolType,
	"show_backup":                   types.BoolType,
	"show_db_ph":                    types.BoolType,
	"show_dbsize":                   types.BoolType,
	"show_ibdbstat":                 types.BoolType,
	"enable_all":                    types.BoolType,
	"disable_all":                   types.BoolType,
}

var AdmingroupDatabaseShowCommandsResourceSchemaAttributes = map[string]schema.Attribute{
	"show_named_max_journal_size": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_txn_trace": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_database_transfer_status": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_backup": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_db_ph": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_dbsize": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"show_ibdbstat": schema.BoolAttribute{
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

func ExpandAdmingroupDatabaseShowCommands(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.AdmingroupDatabaseShowCommands {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AdmingroupDatabaseShowCommandsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AdmingroupDatabaseShowCommandsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.AdmingroupDatabaseShowCommands {
	if m == nil {
		return nil
	}
	to := &security.AdmingroupDatabaseShowCommands{
		ShowNamedMaxJournalSize:    flex.ExpandBoolPointer(m.ShowNamedMaxJournalSize),
		ShowTxnTrace:               flex.ExpandBoolPointer(m.ShowTxnTrace),
		ShowDatabaseTransferStatus: flex.ExpandBoolPointer(m.ShowDatabaseTransferStatus),
		ShowBackup:                 flex.ExpandBoolPointer(m.ShowBackup),
		ShowDbPh:                   flex.ExpandBoolPointer(m.ShowDbPh),
		ShowDbsize:                 flex.ExpandBoolPointer(m.ShowDbsize),
		ShowIbdbstat:               flex.ExpandBoolPointer(m.ShowIbdbstat),
		EnableAll:                  flex.ExpandBoolPointer(m.EnableAll),
		DisableAll:                 flex.ExpandBoolPointer(m.DisableAll),
	}
	return to
}

func FlattenAdmingroupDatabaseShowCommands(ctx context.Context, from *security.AdmingroupDatabaseShowCommands, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AdmingroupDatabaseShowCommandsAttrTypes)
	}
	m := AdmingroupDatabaseShowCommandsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AdmingroupDatabaseShowCommandsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AdmingroupDatabaseShowCommandsModel) Flatten(ctx context.Context, from *security.AdmingroupDatabaseShowCommands, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AdmingroupDatabaseShowCommandsModel{}
	}
	m.ShowNamedMaxJournalSize = types.BoolPointerValue(from.ShowNamedMaxJournalSize)
	m.ShowTxnTrace = types.BoolPointerValue(from.ShowTxnTrace)
	m.ShowDatabaseTransferStatus = types.BoolPointerValue(from.ShowDatabaseTransferStatus)
	m.ShowBackup = types.BoolPointerValue(from.ShowBackup)
	m.ShowDbPh = types.BoolPointerValue(from.ShowDbPh)
	m.ShowDbsize = types.BoolPointerValue(from.ShowDbsize)
	m.ShowIbdbstat = types.BoolPointerValue(from.ShowIbdbstat)
	m.EnableAll = types.BoolPointerValue(from.EnableAll)
	m.DisableAll = types.BoolPointerValue(from.DisableAll)
}
