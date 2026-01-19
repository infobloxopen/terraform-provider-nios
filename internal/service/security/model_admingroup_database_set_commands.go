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

type AdmingroupDatabaseSetCommandsModel struct {
	SetNamedMaxJournalSize types.Bool `tfsdk:"set_named_max_journal_size"`
	SetTxnTrace            types.Bool `tfsdk:"set_txn_trace"`
	SetDatabaseTransfer    types.Bool `tfsdk:"set_database_transfer"`
	SetDbRollover          types.Bool `tfsdk:"set_db_rollover"`
	SetDbSnapshot          types.Bool `tfsdk:"set_db_snapshot"`
	SetDbsize              types.Bool `tfsdk:"set_dbsize"`
	SetDebugTools          types.Bool `tfsdk:"set_debug_tools"`
	SetCircTxnTrace        types.Bool `tfsdk:"set_circ_txn_trace"`
	EnableAll              types.Bool `tfsdk:"enable_all"`
	DisableAll             types.Bool `tfsdk:"disable_all"`
}

var AdmingroupDatabaseSetCommandsAttrTypes = map[string]attr.Type{
	"set_named_max_journal_size": types.BoolType,
	"set_txn_trace":              types.BoolType,
	"set_database_transfer":      types.BoolType,
	"set_db_rollover":            types.BoolType,
	"set_db_snapshot":            types.BoolType,
	"set_dbsize":                 types.BoolType,
	"set_debug_tools":            types.BoolType,
	"set_circ_txn_trace":         types.BoolType,
	"enable_all":                 types.BoolType,
	"disable_all":                types.BoolType,
}

var AdmingroupDatabaseSetCommandsResourceSchemaAttributes = map[string]schema.Attribute{
	"set_named_max_journal_size": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_txn_trace": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_database_transfer": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_db_rollover": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_db_snapshot": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_dbsize": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_debug_tools": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"set_circ_txn_trace": schema.BoolAttribute{
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

func ExpandAdmingroupDatabaseSetCommands(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.AdmingroupDatabaseSetCommands {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AdmingroupDatabaseSetCommandsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AdmingroupDatabaseSetCommandsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.AdmingroupDatabaseSetCommands {
	if m == nil {
		return nil
	}
	to := &security.AdmingroupDatabaseSetCommands{
		SetNamedMaxJournalSize: flex.ExpandBoolPointer(m.SetNamedMaxJournalSize),
		SetTxnTrace:            flex.ExpandBoolPointer(m.SetTxnTrace),
		SetDatabaseTransfer:    flex.ExpandBoolPointer(m.SetDatabaseTransfer),
		SetDbRollover:          flex.ExpandBoolPointer(m.SetDbRollover),
		SetDbSnapshot:          flex.ExpandBoolPointer(m.SetDbSnapshot),
		SetDbsize:              flex.ExpandBoolPointer(m.SetDbsize),
		SetDebugTools:          flex.ExpandBoolPointer(m.SetDebugTools),
		SetCircTxnTrace:        flex.ExpandBoolPointer(m.SetCircTxnTrace),
	}
	return to
}

func FlattenAdmingroupDatabaseSetCommands(ctx context.Context, from *security.AdmingroupDatabaseSetCommands, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AdmingroupDatabaseSetCommandsAttrTypes)
	}
	m := AdmingroupDatabaseSetCommandsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AdmingroupDatabaseSetCommandsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AdmingroupDatabaseSetCommandsModel) Flatten(ctx context.Context, from *security.AdmingroupDatabaseSetCommands, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AdmingroupDatabaseSetCommandsModel{}
	}
	m.SetNamedMaxJournalSize = types.BoolPointerValue(from.SetNamedMaxJournalSize)
	m.SetTxnTrace = types.BoolPointerValue(from.SetTxnTrace)
	m.SetDatabaseTransfer = types.BoolPointerValue(from.SetDatabaseTransfer)
	m.SetDbRollover = types.BoolPointerValue(from.SetDbRollover)
	m.SetDbSnapshot = types.BoolPointerValue(from.SetDbSnapshot)
	m.SetDbsize = types.BoolPointerValue(from.SetDbsize)
	m.SetDebugTools = types.BoolPointerValue(from.SetDebugTools)
	m.SetCircTxnTrace = types.BoolPointerValue(from.SetCircTxnTrace)
	m.EnableAll = types.BoolPointerValue(from.EnableAll)
	m.DisableAll = types.BoolPointerValue(from.DisableAll)
}
