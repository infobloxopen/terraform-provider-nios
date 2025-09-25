package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type DistributionscheduleUpgradeGroupsModel struct {
	Name                       types.String `tfsdk:"name"`
	TimeZone                   types.String `tfsdk:"time_zone"`
	DistributionDependentGroup types.String `tfsdk:"distribution_dependent_group"`
	UpgradeDependentGroup      types.String `tfsdk:"upgrade_dependent_group"`
	DistributionTime           types.String `tfsdk:"distribution_time"`
	UpgradeTime                types.Int64  `tfsdk:"upgrade_time"`
}

var DistributionscheduleUpgradeGroupsAttrTypes = map[string]attr.Type{
	"name":                         types.StringType,
	"time_zone":                    types.StringType,
	"distribution_dependent_group": types.StringType,
	"upgrade_dependent_group":      types.StringType,
	"distribution_time":            types.StringType,
	"upgrade_time":                 types.Int64Type,
}

var DistributionscheduleUpgradeGroupsResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Computed:            true,
		Optional:            true,
		MarkdownDescription: "The upgrade group name. Required when specifying upgrade_groups",
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
	},
	"time_zone": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The time zone for scheduling operations.",
	},
	"distribution_dependent_group": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The distribution dependent group name.",
	},
	"upgrade_dependent_group": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The upgrade dependent group name.",
	},
	"distribution_time": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The time of the next scheduled distribution.",
		Validators: []validator.String{
			customvalidator.ValidateTimeFormat(),
		},
	},
	"upgrade_time": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The time of the next scheduled upgrade.",
	},
}

func ExpandDistributionscheduleUpgradeGroups(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.DistributionscheduleUpgradeGroups {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DistributionscheduleUpgradeGroupsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DistributionscheduleUpgradeGroupsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.DistributionscheduleUpgradeGroups {
	if m == nil {
		return nil
	}
	to := &grid.DistributionscheduleUpgradeGroups{
		Name:                       flex.ExpandStringPointer(m.Name),
		DistributionDependentGroup: flex.ExpandStringPointer(m.DistributionDependentGroup),
		UpgradeDependentGroup:      flex.ExpandStringPointer(m.UpgradeDependentGroup),
		UpgradeTime:                flex.ExpandInt64Pointer(m.UpgradeTime),
	}

	to.DistributionTime = flex.ExpandTimeToUnix(m.DistributionTime, diags)

	return to
}

func FlattenDistributionscheduleUpgradeGroups(ctx context.Context, from *grid.DistributionscheduleUpgradeGroups, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DistributionscheduleUpgradeGroupsAttrTypes)
	}
	m := DistributionscheduleUpgradeGroupsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DistributionscheduleUpgradeGroupsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DistributionscheduleUpgradeGroupsModel) Flatten(ctx context.Context, from *grid.DistributionscheduleUpgradeGroups, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DistributionscheduleUpgradeGroupsModel{}
	}

	m.Name = flex.FlattenStringPointer(from.Name)
	m.TimeZone = flex.FlattenStringPointer(from.TimeZone)
	m.DistributionDependentGroup = flex.FlattenStringPointer(from.DistributionDependentGroup)
	m.DistributionTime = flex.FlattenUnixTime(from.DistributionTime, diags)
	m.UpgradeDependentGroup = flex.FlattenStringPointer(from.UpgradeDependentGroup)
	m.DistributionTime = flex.FlattenUnixTime(from.DistributionTime, diags)
	m.UpgradeTime = flex.FlattenInt64Pointer(from.UpgradeTime)
}
