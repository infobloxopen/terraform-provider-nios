package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type GridSnmpSettingModel struct {
	EngineId               types.List   `tfsdk:"engine_id"`
	QueriesCommunityString types.String `tfsdk:"queries_community_string"`
	QueriesEnable          types.Bool   `tfsdk:"queries_enable"`
	Snmpv3QueriesEnable    types.Bool   `tfsdk:"snmpv3_queries_enable"`
	Snmpv3QueriesUsers     types.List   `tfsdk:"snmpv3_queries_users"`
	Snmpv3TrapsEnable      types.Bool   `tfsdk:"snmpv3_traps_enable"`
	Syscontact             types.List   `tfsdk:"syscontact"`
	Sysdescr               types.List   `tfsdk:"sysdescr"`
	Syslocation            types.List   `tfsdk:"syslocation"`
	Sysname                types.List   `tfsdk:"sysname"`
	TrapReceivers          types.List   `tfsdk:"trap_receivers"`
	TrapsCommunityString   types.String `tfsdk:"traps_community_string"`
	TrapsEnable            types.Bool   `tfsdk:"traps_enable"`
}

var GridSnmpSettingAttrTypes = map[string]attr.Type{
	"engine_id":                types.ListType{ElemType: types.StringType},
	"queries_community_string": types.StringType,
	"queries_enable":           types.BoolType,
	"snmpv3_queries_enable":    types.BoolType,
	"snmpv3_queries_users":     types.ListType{ElemType: types.ObjectType{AttrTypes: GridsnmpsettingSnmpv3QueriesUsersAttrTypes}},
	"snmpv3_traps_enable":      types.BoolType,
	"syscontact":               types.ListType{ElemType: types.StringType},
	"sysdescr":                 types.ListType{ElemType: types.StringType},
	"syslocation":              types.ListType{ElemType: types.StringType},
	"sysname":                  types.ListType{ElemType: types.StringType},
	"trap_receivers":           types.ListType{ElemType: types.ObjectType{AttrTypes: GridsnmpsettingTrapReceiversAttrTypes}},
	"traps_community_string":   types.StringType,
	"traps_enable":             types.BoolType,
}

var GridSnmpSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"engine_id": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Computed:            true,
		MarkdownDescription: "The engine ID of the appliance that manages the SNMP agent.",
	},
	"queries_community_string": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The community string for SNMP queries.",
	},
	"queries_enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If set to True, SNMP queries are enabled.",
	},
	"snmpv3_queries_enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If set to True, SNMPv3 queries are enabled.",
	},
	"snmpv3_queries_users": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: GridsnmpsettingSnmpv3QueriesUsersResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "A list of SNMPv3 queries users.",
	},
	"snmpv3_traps_enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If set to True, SNMPv3 traps are enabled.",
	},
	"syscontact": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The name of the contact person for the appliance. Second value is applicable only for HA pair. Otherwise second value is ignored.",
	},
	"sysdescr": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "Useful information about the appliance. Second value is applicable only for HA pair. Otherwise second value is ignored.",
	},
	"syslocation": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The physical location of the appliance. Second value is applicable only for HA pair. Otherwise second value is ignored.",
	},
	"sysname": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The FQDN (Fully Qualified Domain Name) of the appliance. Second value is applicable only for HA pair. Otherwise second value is ignored.",
	},
	"trap_receivers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: GridsnmpsettingTrapReceiversResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "A list of trap receivers.",
	},
	"traps_community_string": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "A string the NIOS appliance sends to the management system together with its traps. Note that this community string must match exactly what you enter in the management system.",
	},
	"traps_enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If set to True, SNMP traps are enabled.",
	},
}

func ExpandGridSnmpSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridSnmpSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridSnmpSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridSnmpSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridSnmpSetting {
	if m == nil {
		return nil
	}
	to := &grid.GridSnmpSetting{
		QueriesCommunityString: flex.ExpandStringPointer(m.QueriesCommunityString),
		QueriesEnable:          flex.ExpandBoolPointer(m.QueriesEnable),
		Snmpv3QueriesEnable:    flex.ExpandBoolPointer(m.Snmpv3QueriesEnable),
		Snmpv3QueriesUsers:     flex.ExpandFrameworkListNestedBlock(ctx, m.Snmpv3QueriesUsers, diags, ExpandGridsnmpsettingSnmpv3QueriesUsers),
		Snmpv3TrapsEnable:      flex.ExpandBoolPointer(m.Snmpv3TrapsEnable),
		Syscontact:             flex.ExpandFrameworkListString(ctx, m.Syscontact, diags),
		Sysdescr:               flex.ExpandFrameworkListString(ctx, m.Sysdescr, diags),
		Syslocation:            flex.ExpandFrameworkListString(ctx, m.Syslocation, diags),
		Sysname:                flex.ExpandFrameworkListString(ctx, m.Sysname, diags),
		TrapReceivers:          flex.ExpandFrameworkListNestedBlock(ctx, m.TrapReceivers, diags, ExpandGridsnmpsettingTrapReceivers),
		TrapsCommunityString:   flex.ExpandStringPointer(m.TrapsCommunityString),
		TrapsEnable:            flex.ExpandBoolPointer(m.TrapsEnable),
	}
	return to
}

func FlattenGridSnmpSetting(ctx context.Context, from *grid.GridSnmpSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridSnmpSettingAttrTypes)
	}
	m := GridSnmpSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridSnmpSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridSnmpSettingModel) Flatten(ctx context.Context, from *grid.GridSnmpSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridSnmpSettingModel{}
	}
	m.EngineId = flex.FlattenFrameworkListString(ctx, from.EngineId, diags)
	m.QueriesCommunityString = flex.FlattenStringPointer(from.QueriesCommunityString)
	m.QueriesEnable = types.BoolPointerValue(from.QueriesEnable)
	m.Snmpv3QueriesEnable = types.BoolPointerValue(from.Snmpv3QueriesEnable)
	m.Snmpv3QueriesUsers = flex.FlattenFrameworkListNestedBlock(ctx, from.Snmpv3QueriesUsers, GridsnmpsettingSnmpv3QueriesUsersAttrTypes, diags, FlattenGridsnmpsettingSnmpv3QueriesUsers)
	m.Snmpv3TrapsEnable = types.BoolPointerValue(from.Snmpv3TrapsEnable)
	m.Syscontact = flex.FlattenFrameworkListString(ctx, from.Syscontact, diags)
	m.Sysdescr = flex.FlattenFrameworkListString(ctx, from.Sysdescr, diags)
	m.Syslocation = flex.FlattenFrameworkListString(ctx, from.Syslocation, diags)
	m.Sysname = flex.FlattenFrameworkListString(ctx, from.Sysname, diags)
	m.TrapReceivers = flex.FlattenFrameworkListNestedBlock(ctx, from.TrapReceivers, GridsnmpsettingTrapReceiversAttrTypes, diags, FlattenGridsnmpsettingTrapReceivers)
	m.TrapsCommunityString = flex.FlattenStringPointer(from.TrapsCommunityString)
	m.TrapsEnable = types.BoolPointerValue(from.TrapsEnable)
}
