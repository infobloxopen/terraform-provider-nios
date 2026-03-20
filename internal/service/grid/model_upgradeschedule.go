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
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type UpgradescheduleModel struct {
	Ref           types.String `tfsdk:"ref"`
	Uuid          types.String `tfsdk:"uuid"`
	UUID          types.String `tfsdk:"uuid"`
	Active        types.Bool   `tfsdk:"active"`
	StartTime     types.String `tfsdk:"start_time"`
	TimeZone      types.String `tfsdk:"time_zone"`
	UpgradeGroups types.List   `tfsdk:"upgrade_groups"`
}

var UpgradescheduleAttrTypes = map[string]attr.Type{
	"ref":            types.StringType,
	"uuid":           types.StringType,
	"active":         types.BoolType,
	"start_time":     types.StringType,
	"time_zone":      types.StringType,
	"upgrade_groups": types.ListType{ElemType: types.ObjectType{AttrTypes: UpgradescheduleUpgradeGroupsAttrTypes}},
}

var UpgradescheduleResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Universally Unique ID assigned for this object.",
	},
	"active": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Determines whether the upgrade schedule is active.",
	},
	"start_time": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The start time of the upgrade.",
		Validators: []validator.String{
			customvalidator.ValidateTimeFormat(),
		},
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
		Computed:            true,
		MarkdownDescription: "The upgrade groups scheduling settings.",
	},
}

func (m *UpgradescheduleModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.Upgradeschedule {
	var groups []grid.UpgradescheduleUpgradeGroups

	if m == nil {
		return nil
	}

	allGroups := flex.ExpandFrameworkListNestedBlock(ctx, m.UpgradeGroups, diags, ExpandUpgradescheduleUpgradeGroups)

	for _, group := range allGroups {
		// Convert empty optional fields to nil
		if group.UpgradeDependentGroup != nil && *group.UpgradeDependentGroup == "" {
			group.UpgradeDependentGroup = nil
		}
		if group.DistributionDependentGroup != nil && *group.DistributionDependentGroup == "" {
			group.DistributionDependentGroup = nil
		}

		// UpgradeTime cannot be nil, set to 0 if not provided
		if group.UpgradeTime == nil {
			val := int64(0)
			group.UpgradeTime = &val
		}

		groups = append(groups, group)
	}

	to := &grid.Upgradeschedule{
		Active:        flex.ExpandBoolPointer(m.Active),
		UpgradeGroups: groups,
	}

	to.StartTime = flex.ExpandTimeToUnix(m.StartTime, diags)

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
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Active = types.BoolPointerValue(from.Active)
	m.StartTime = flex.FlattenUnixTime(from.StartTime, diags)
	m.TimeZone = flex.FlattenStringPointer(from.TimeZone)
	planUpgradeGroups := m.UpgradeGroups
	m.UpgradeGroups = flex.FlattenFrameworkListNestedBlock(ctx, from.UpgradeGroups, UpgradescheduleUpgradeGroupsAttrTypes, diags, FlattenUpgradescheduleUpgradeGroups)
	if !planUpgradeGroups.IsUnknown() {
		reOrderedList, diags := utils.ReorderAndFilterNestedListResponse(ctx, planUpgradeGroups, m.UpgradeGroups, "name")
		if !diags.HasError() {
			m.UpgradeGroups = reOrderedList.(basetypes.ListValue)
		}
	}
}
