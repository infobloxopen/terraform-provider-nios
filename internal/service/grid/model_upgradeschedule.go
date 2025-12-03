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

type UpgradescheduleModel struct {
	Ref           types.String `tfsdk:"ref"`
	Active        types.Bool   `tfsdk:"active"`
	StartTime     types.Int64  `tfsdk:"start_time"`
	TimeZone      types.String `tfsdk:"time_zone"`
	UpgradeGroups types.List   `tfsdk:"upgrade_groups"`
}

var UpgradescheduleAttrTypes = map[string]attr.Type{
	"ref":            types.StringType,
	"active":         types.BoolType,
	"start_time":     types.Int64Type,
	"time_zone":      types.StringType,
	"upgrade_groups": types.ListType{ElemType: types.ObjectType{AttrTypes: UpgradescheduleUpgradeGroupsAttrTypes}},
}

var UpgradescheduleResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"active": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the upgrade schedule is active.",
	},
	"start_time": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The start time of the upgrade.",
	},
	"time_zone": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The time zone for upgrade start time.",
	},
	"upgrade_groups": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: UpgradescheduleUpgradeGroupsResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The upgrade groups scheduling settings.",
	},
}

func ExpandUpgradeschedule(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.Upgradeschedule {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m UpgradescheduleModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *UpgradescheduleModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.Upgradeschedule {
	if m == nil {
		return nil
	}
	to := &grid.Upgradeschedule{
		Ref:           flex.ExpandStringPointer(m.Ref),
		Active:        flex.ExpandBoolPointer(m.Active),
		StartTime:     flex.ExpandInt64Pointer(m.StartTime),
		UpgradeGroups: flex.ExpandFrameworkListNestedBlock(ctx, m.UpgradeGroups, diags, ExpandUpgradescheduleUpgradeGroups),
	}
	return to
}

func FlattenUpgradeschedule(ctx context.Context, from *grid.Upgradeschedule, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(UpgradescheduleAttrTypes)
	}
	m := UpgradescheduleModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, UpgradescheduleAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *UpgradescheduleModel) Flatten(ctx context.Context, from *grid.Upgradeschedule, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = UpgradescheduleModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Active = types.BoolPointerValue(from.Active)
	m.StartTime = flex.FlattenInt64Pointer(from.StartTime)
	m.TimeZone = flex.FlattenStringPointer(from.TimeZone)
	m.UpgradeGroups = flex.FlattenFrameworkListNestedBlock(ctx, from.UpgradeGroups, UpgradescheduleUpgradeGroupsAttrTypes, diags, FlattenUpgradescheduleUpgradeGroups)
}
