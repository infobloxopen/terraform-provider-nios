package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type UpgradescheduleUpgradeGroupsModel struct {
	Name                       types.String `tfsdk:"name"`
	TimeZone                   types.String `tfsdk:"time_zone"`
	DistributionDependentGroup types.String `tfsdk:"distribution_dependent_group"`
	UpgradeDependentGroup      types.String `tfsdk:"upgrade_dependent_group"`
	DistributionTime           types.Int64  `tfsdk:"distribution_time"`
	UpgradeTime                types.Int64  `tfsdk:"upgrade_time"`
}

var UpgradescheduleUpgradeGroupsAttrTypes = map[string]attr.Type{
	"name":                         types.StringType,
	"time_zone":                    types.StringType,
	"distribution_dependent_group": types.StringType,
	"upgrade_dependent_group":      types.StringType,
	"distribution_time":            types.Int64Type,
	"upgrade_time":                 types.Int64Type,
}

var UpgradescheduleUpgradeGroupsResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The upgrade group name.",
	},
	"time_zone": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The time zone for scheduling operations.",
	},
	"distribution_dependent_group": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The distribution dependent group name.",
	},
	"upgrade_dependent_group": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The upgrade dependent group name.",
	},
	"distribution_time": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The time of the next scheduled distribution.",
	},
	"upgrade_time": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The time of the next scheduled upgrade.",
	},
}

func ExpandUpgradescheduleUpgradeGroups(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.UpgradescheduleUpgradeGroups {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m UpgradescheduleUpgradeGroupsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *UpgradescheduleUpgradeGroupsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.UpgradescheduleUpgradeGroups {
	if m == nil {
		return nil
	}
	to := &grid.UpgradescheduleUpgradeGroups{
		Name:                       flex.ExpandStringPointer(m.Name),
		DistributionDependentGroup: flex.ExpandStringPointer(m.DistributionDependentGroup),
		UpgradeDependentGroup:      flex.ExpandStringPointer(m.UpgradeDependentGroup),
		DistributionTime:           flex.ExpandInt64Pointer(m.DistributionTime),
		UpgradeTime:                flex.ExpandInt64Pointer(m.UpgradeTime),
	}
	return to
}

func FlattenUpgradescheduleUpgradeGroups(ctx context.Context, from *grid.UpgradescheduleUpgradeGroups, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(UpgradescheduleUpgradeGroupsAttrTypes)
	}
	m := UpgradescheduleUpgradeGroupsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, UpgradescheduleUpgradeGroupsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *UpgradescheduleUpgradeGroupsModel) Flatten(ctx context.Context, from *grid.UpgradescheduleUpgradeGroups, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = UpgradescheduleUpgradeGroupsModel{}
	}
	m.Name = flex.FlattenStringPointer(from.Name)
	m.TimeZone = flex.FlattenStringPointer(from.TimeZone)
	m.DistributionDependentGroup = flex.FlattenStringPointer(from.DistributionDependentGroup)
	m.UpgradeDependentGroup = flex.FlattenStringPointer(from.UpgradeDependentGroup)
	m.DistributionTime = flex.FlattenInt64Pointer(from.DistributionTime)
	m.UpgradeTime = flex.FlattenInt64Pointer(from.UpgradeTime)
}
