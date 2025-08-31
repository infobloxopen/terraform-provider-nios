package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type DistributionscheduleModel struct {
	Ref           types.String `tfsdk:"ref"`
	Active        types.Bool   `tfsdk:"active"`
	StartTime     types.Int64  `tfsdk:"start_time"`
	TimeZone      types.String `tfsdk:"time_zone"`
	UpgradeGroups types.List   `tfsdk:"upgrade_groups"`
}

var DistributionscheduleAttrTypes = map[string]attr.Type{
	"ref":            types.StringType,
	"active":         types.BoolType,
	"start_time":     types.Int64Type,
	"time_zone":      types.StringType,
	"upgrade_groups": types.ListType{ElemType: types.ObjectType{AttrTypes: DistributionscheduleUpgradeGroupsAttrTypes}},
}

var DistributionscheduleResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
		PlanModifiers: []planmodifier.String{
			UseStateForUnknownString(),
		},
	},
	"active": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Determines whether the distribution schedule is active.",
		PlanModifiers: []planmodifier.Bool{
			UseStateForUnknownBool(),
		},
	},
	"start_time": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The start time of the distribution.",
		PlanModifiers: []planmodifier.Int64{
			UseStateForUnknownInt64(),
		},
	},
	"time_zone": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Time zone of the distribution start time.",
		PlanModifiers: []planmodifier.String{
			UseStateForUnknownString(),
		},
	},
	"upgrade_groups": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: DistributionscheduleUpgradeGroupsResourceSchemaAttributes,
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The upgrade groups scheduling settings.",
		PlanModifiers: []planmodifier.List{
			UseStateForUnknownList(),
		},
	},
}

func UseStateForUnknownInt64() planmodifier.Int64 {
	return useStateForUnknownInt64{}
}

type useStateForUnknownInt64 struct{}

func (m useStateForUnknownInt64) Description(ctx context.Context) string {
	return "Use state value when unknown"
}

func (m useStateForUnknownInt64) MarkdownDescription(ctx context.Context) string {
	return "Use state value when unknown"
}

func (m useStateForUnknownInt64) PlanModifyInt64(ctx context.Context, req planmodifier.Int64Request, resp *planmodifier.Int64Response) {
	if req.PlanValue.IsUnknown() && req.ConfigValue.IsNull() && !req.StateValue.IsNull() {
		resp.PlanValue = req.StateValue
	}
}

func UseStateForUnknownBool() planmodifier.Bool {
	return useStateForUnknownBool{}
}

type useStateForUnknownBool struct{}

func (m useStateForUnknownBool) Description(ctx context.Context) string {
	return "Use state value when unknown"
}

func (m useStateForUnknownBool) MarkdownDescription(ctx context.Context) string {
	return "Use state value when unknown"
}

func (m useStateForUnknownBool) PlanModifyBool(ctx context.Context, req planmodifier.BoolRequest, resp *planmodifier.BoolResponse) {
	if req.PlanValue.IsUnknown() && req.ConfigValue.IsNull() && !req.StateValue.IsNull() {
		resp.PlanValue = req.StateValue
	}
}

func UseStateForUnknownString() planmodifier.String {
	return useStateForUnknownString{}
}

type useStateForUnknownString struct{}

func (m useStateForUnknownString) Description(ctx context.Context) string {
	return "Use state value when unknown"
}

func (m useStateForUnknownString) MarkdownDescription(ctx context.Context) string {
	return "Use state value when unknown"
}

func (m useStateForUnknownString) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.PlanValue.IsUnknown() && req.ConfigValue.IsNull() && !req.StateValue.IsNull() {
		resp.PlanValue = req.StateValue
	}
}

func UseStateForUnknownList() planmodifier.List {
	return useStateForUnknownList{}
}

type useStateForUnknownList struct{}

func (m useStateForUnknownList) Description(ctx context.Context) string {
	return "Use state value when unknown"
}

func (m useStateForUnknownList) MarkdownDescription(ctx context.Context) string {
	return "Use state value when unknown"
}

func (m useStateForUnknownList) PlanModifyList(ctx context.Context, req planmodifier.ListRequest, resp *planmodifier.ListResponse) {
	if req.PlanValue.IsUnknown() && req.ConfigValue.IsNull() && !req.StateValue.IsNull() {
		resp.PlanValue = req.StateValue
	}
}

func (m *DistributionscheduleModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.Distributionschedule {
	if m == nil {
		return nil
	}
	allGroups := flex.ExpandFrameworkListNestedBlock(ctx, m.UpgradeGroups, diags, ExpandDistributionscheduleUpgradeGroups)
	var groups []grid.DistributionscheduleUpgradeGroups

	for _, g := range allGroups {
		// Skip invalid groups entirely
		if g.Name == nil || *g.Name == "" {
			diags.AddError("Invalid Upgrade Group", "Upgrade group 'name' is required and cannot be empty")
			continue
		}
		if g.DistributionTime == nil || *g.DistributionTime == 0 {
			diags.AddError("Invalid Upgrade Group", "Upgrade group 'distribution_time' is required and cannot be 0")
			continue
		}

		// Convert empty optional fields to nil
		if g.UpgradeDependentGroup != nil && *g.UpgradeDependentGroup == "" {
			g.UpgradeDependentGroup = nil
		}
		if g.DistributionDependentGroup != nil && *g.DistributionDependentGroup == "" {
			g.DistributionDependentGroup = nil
		}

		// UpgradeTime: leave 0 as is, can be omitted
		if g.UpgradeTime == nil {
			val := int64(0)
			g.UpgradeTime = &val
		}

		groups = append(groups, g)
	}

	to := &grid.Distributionschedule{
		Active:        flex.ExpandBoolPointer(m.Active),
		StartTime:     flex.ExpandInt64Pointer(m.StartTime),
		UpgradeGroups: groups,
	}
	return to
}

func FlattenDistributionschedule(ctx context.Context, from *grid.Distributionschedule, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DistributionscheduleAttrTypes)
	}
	m := DistributionscheduleModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DistributionscheduleAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DistributionscheduleModel) Flatten(ctx context.Context, from *grid.Distributionschedule, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DistributionscheduleModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Active = types.BoolPointerValue(from.Active)
	m.StartTime = flex.FlattenInt64Pointer(from.StartTime)
	m.TimeZone = flex.FlattenStringPointer(from.TimeZone)
	m.UpgradeGroups = flex.FlattenFrameworkListNestedBlock(ctx, from.UpgradeGroups, DistributionscheduleUpgradeGroupsAttrTypes, diags, FlattenDistributionscheduleUpgradeGroups)
}
