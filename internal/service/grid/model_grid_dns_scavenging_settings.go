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

type GridDnsScavengingSettingsModel struct {
	EnableScavenging          types.Bool   `tfsdk:"enable_scavenging"`
	EnableRecurrentScavenging types.Bool   `tfsdk:"enable_recurrent_scavenging"`
	EnableAutoReclamation     types.Bool   `tfsdk:"enable_auto_reclamation"`
	EnableRrLastQueried       types.Bool   `tfsdk:"enable_rr_last_queried"`
	EnableZoneLastQueried     types.Bool   `tfsdk:"enable_zone_last_queried"`
	ReclaimAssociatedRecords  types.Bool   `tfsdk:"reclaim_associated_records"`
	ScavengingSchedule        types.Object `tfsdk:"scavenging_schedule"`
	ExpressionList            types.List   `tfsdk:"expression_list"`
	EaExpressionList          types.List   `tfsdk:"ea_expression_list"`
}

var GridDnsScavengingSettingsAttrTypes = map[string]attr.Type{
	"enable_scavenging":           types.BoolType,
	"enable_recurrent_scavenging": types.BoolType,
	"enable_auto_reclamation":     types.BoolType,
	"enable_rr_last_queried":      types.BoolType,
	"enable_zone_last_queried":    types.BoolType,
	"reclaim_associated_records":  types.BoolType,
	"scavenging_schedule":         types.ObjectType{AttrTypes: GriddnsscavengingsettingsScavengingScheduleAttrTypes},
	"expression_list":             types.ListType{ElemType: types.ObjectType{AttrTypes: GriddnsscavengingsettingsExpressionListAttrTypes}},
	"ea_expression_list":          types.ListType{ElemType: types.ObjectType{AttrTypes: GriddnsscavengingsettingsEaExpressionListAttrTypes}},
}

var GridDnsScavengingSettingsResourceSchemaAttributes = map[string]schema.Attribute{
	"enable_scavenging": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "This flag indicates if the resource record scavenging is enabled or not.",
	},
	"enable_recurrent_scavenging": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "This flag indicates if the recurrent resource record scavenging is enabled or not.",
	},
	"enable_auto_reclamation": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "This flag indicates if the automatic resource record scavenging is enabled or not.",
	},
	"enable_rr_last_queried": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "This flag indicates if the resource record last queried monitoring in affected zones is enabled or not.",
	},
	"enable_zone_last_queried": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "This flag indicates if the last queried monitoring for affected zones is enabled or not.",
	},
	"reclaim_associated_records": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "This flag indicates if the associated resource record scavenging is enabled or not.",
	},
	"scavenging_schedule": schema.SingleNestedAttribute{
		Attributes: GriddnsscavengingsettingsScavengingScheduleResourceSchemaAttributes,
		Optional:   true,
	},
	"expression_list": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: GriddnsscavengingsettingsExpressionListResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The expression list. The particular record is treated as reclaimable if expression condition evaluates to 'true' for given record if scavenging hasn't been manually disabled on a given resource record.",
	},
	"ea_expression_list": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: GriddnsscavengingsettingsEaExpressionListResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The extensible attributes expression list. The particular record is treated as reclaimable if extensible attributes expression condition evaluates to 'true' for given record if scavenging hasn't been manually disabled on a given resource record.",
	},
}

func ExpandGridDnsScavengingSettings(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridDnsScavengingSettings {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridDnsScavengingSettingsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridDnsScavengingSettingsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridDnsScavengingSettings {
	if m == nil {
		return nil
	}
	to := &grid.GridDnsScavengingSettings{
		EnableScavenging:          flex.ExpandBoolPointer(m.EnableScavenging),
		EnableRecurrentScavenging: flex.ExpandBoolPointer(m.EnableRecurrentScavenging),
		EnableAutoReclamation:     flex.ExpandBoolPointer(m.EnableAutoReclamation),
		EnableRrLastQueried:       flex.ExpandBoolPointer(m.EnableRrLastQueried),
		EnableZoneLastQueried:     flex.ExpandBoolPointer(m.EnableZoneLastQueried),
		ReclaimAssociatedRecords:  flex.ExpandBoolPointer(m.ReclaimAssociatedRecords),
		ScavengingSchedule:        ExpandGriddnsscavengingsettingsScavengingSchedule(ctx, m.ScavengingSchedule, diags),
		ExpressionList:            flex.ExpandFrameworkListNestedBlock(ctx, m.ExpressionList, diags, ExpandGriddnsscavengingsettingsExpressionList),
		EaExpressionList:          flex.ExpandFrameworkListNestedBlock(ctx, m.EaExpressionList, diags, ExpandGriddnsscavengingsettingsEaExpressionList),
	}
	return to
}

func FlattenGridDnsScavengingSettings(ctx context.Context, from *grid.GridDnsScavengingSettings, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridDnsScavengingSettingsAttrTypes)
	}
	m := GridDnsScavengingSettingsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridDnsScavengingSettingsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridDnsScavengingSettingsModel) Flatten(ctx context.Context, from *grid.GridDnsScavengingSettings, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridDnsScavengingSettingsModel{}
	}
	m.EnableScavenging = types.BoolPointerValue(from.EnableScavenging)
	m.EnableRecurrentScavenging = types.BoolPointerValue(from.EnableRecurrentScavenging)
	m.EnableAutoReclamation = types.BoolPointerValue(from.EnableAutoReclamation)
	m.EnableRrLastQueried = types.BoolPointerValue(from.EnableRrLastQueried)
	m.EnableZoneLastQueried = types.BoolPointerValue(from.EnableZoneLastQueried)
	m.ReclaimAssociatedRecords = types.BoolPointerValue(from.ReclaimAssociatedRecords)
	m.ScavengingSchedule = FlattenGriddnsscavengingsettingsScavengingSchedule(ctx, from.ScavengingSchedule, diags)
	m.ExpressionList = flex.FlattenFrameworkListNestedBlock(ctx, from.ExpressionList, GriddnsscavengingsettingsExpressionListAttrTypes, diags, FlattenGriddnsscavengingsettingsExpressionList)
	m.EaExpressionList = flex.FlattenFrameworkListNestedBlock(ctx, from.EaExpressionList, GriddnsscavengingsettingsEaExpressionListAttrTypes, diags, FlattenGriddnsscavengingsettingsEaExpressionList)
}
