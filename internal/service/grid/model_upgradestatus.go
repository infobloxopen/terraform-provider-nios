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

type UpgradestatusModel struct {
	Ref                         types.String `tfsdk:"ref"`
	Uuid                        types.String `tfsdk:"uuid"`
	AllowDistribution           types.Bool   `tfsdk:"allow_distribution"`
	AllowDistributionScheduling types.Bool   `tfsdk:"allow_distribution_scheduling"`
	AllowUpgrade                types.Bool   `tfsdk:"allow_upgrade"`
	AllowUpgradeCancel          types.Bool   `tfsdk:"allow_upgrade_cancel"`
	AllowUpgradePause           types.Bool   `tfsdk:"allow_upgrade_pause"`
	AllowUpgradeResume          types.Bool   `tfsdk:"allow_upgrade_resume"`
	AllowUpgradeScheduling      types.Bool   `tfsdk:"allow_upgrade_scheduling"`
	AllowUpgradeTest            types.Bool   `tfsdk:"allow_upgrade_test"`
	AllowUpload                 types.Bool   `tfsdk:"allow_upload"`
	AlternateVersion            types.String `tfsdk:"alternate_version"`
	Comment                     types.String `tfsdk:"comment"`
	CurrentVersion              types.String `tfsdk:"current_version"`
	CurrentVersionSummary       types.String `tfsdk:"current_version_summary"`
	DistributionScheduleActive  types.Bool   `tfsdk:"distribution_schedule_active"`
	DistributionScheduleTime    types.Int64  `tfsdk:"distribution_schedule_time"`
	DistributionState           types.String `tfsdk:"distribution_state"`
	DistributionVersion         types.String `tfsdk:"distribution_version"`
	DistributionVersionSummary  types.String `tfsdk:"distribution_version_summary"`
	ElementStatus               types.String `tfsdk:"element_status"`
	GridState                   types.String `tfsdk:"grid_state"`
	GroupState                  types.String `tfsdk:"group_state"`
	HaStatus                    types.String `tfsdk:"ha_status"`
	Hotfixes                    types.List   `tfsdk:"hotfixes"`
	Ipv4Address                 types.String `tfsdk:"ipv4_address"`
	Ipv6Address                 types.String `tfsdk:"ipv6_address"`
	Member                      types.String `tfsdk:"member"`
	Message                     types.String `tfsdk:"message"`
	PnodeRole                   types.String `tfsdk:"pnode_role"`
	Reverted                    types.Bool   `tfsdk:"reverted"`
	StatusTime                  types.Int64  `tfsdk:"status_time"`
	StatusValue                 types.String `tfsdk:"status_value"`
	StatusValueUpdateTime       types.Int64  `tfsdk:"status_value_update_time"`
	Steps                       types.List   `tfsdk:"steps"`
	StepsCompleted              types.Int64  `tfsdk:"steps_completed"`
	StepsTotal                  types.Int64  `tfsdk:"steps_total"`
	SubelementType              types.String `tfsdk:"subelement_type"`
	SubelementsCompleted        types.Int64  `tfsdk:"subelements_completed"`
	SubelementsStatus           types.List   `tfsdk:"subelements_status"`
	SubelementsTotal            types.Int64  `tfsdk:"subelements_total"`
	Type                        types.String `tfsdk:"type"`
	UpgradeGroup                types.String `tfsdk:"upgrade_group"`
	UpgradeScheduleActive       types.Bool   `tfsdk:"upgrade_schedule_active"`
	UpgradeState                types.String `tfsdk:"upgrade_state"`
	UpgradeTestStatus           types.String `tfsdk:"upgrade_test_status"`
	UploadVersion               types.String `tfsdk:"upload_version"`
	UploadVersionSummary        types.String `tfsdk:"upload_version_summary"`
}

var UpgradestatusAttrTypes = map[string]attr.Type{
	"ref":                           types.StringType,
	"uuid":                          types.StringType,
	"allow_distribution":            types.BoolType,
	"allow_distribution_scheduling": types.BoolType,
	"allow_upgrade":                 types.BoolType,
	"allow_upgrade_cancel":          types.BoolType,
	"allow_upgrade_pause":           types.BoolType,
	"allow_upgrade_resume":          types.BoolType,
	"allow_upgrade_scheduling":      types.BoolType,
	"allow_upgrade_test":            types.BoolType,
	"allow_upload":                  types.BoolType,
	"alternate_version":             types.StringType,
	"comment":                       types.StringType,
	"current_version":               types.StringType,
	"current_version_summary":       types.StringType,
	"distribution_schedule_active":  types.BoolType,
	"distribution_schedule_time":    types.Int64Type,
	"distribution_state":            types.StringType,
	"distribution_version":          types.StringType,
	"distribution_version_summary":  types.StringType,
	"element_status":                types.StringType,
	"grid_state":                    types.StringType,
	"group_state":                   types.StringType,
	"ha_status":                     types.StringType,
	"hotfixes":                      types.ListType{ElemType: types.ObjectType{AttrTypes: UpgradestatusHotfixesAttrTypes}},
	"ipv4_address":                  types.StringType,
	"ipv6_address":                  types.StringType,
	"member":                        types.StringType,
	"message":                       types.StringType,
	"pnode_role":                    types.StringType,
	"reverted":                      types.BoolType,
	"status_time":                   types.Int64Type,
	"status_value":                  types.StringType,
	"status_value_update_time":      types.Int64Type,
	"steps":                         types.ListType{ElemType: types.ObjectType{AttrTypes: UpgradestatusStepsAttrTypes}},
	"steps_completed":               types.Int64Type,
	"steps_total":                   types.Int64Type,
	"subelement_type":               types.StringType,
	"subelements_completed":         types.Int64Type,
	"subelements_status":            types.ListType{ElemType: types.StringType},
	"subelements_total":             types.Int64Type,
	"type":                          types.StringType,
	"upgrade_group":                 types.StringType,
	"upgrade_schedule_active":       types.BoolType,
	"upgrade_state":                 types.StringType,
	"upgrade_test_status":           types.StringType,
	"upload_version":                types.StringType,
	"upload_version_summary":        types.StringType,
}

var UpgradestatusResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The UUID of the object.",
	},
	"allow_distribution": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines if distribution is allowed for the Grid.",
	},
	"allow_distribution_scheduling": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines if distribution scheduling is allowed.",
	},
	"allow_upgrade": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines if upgrade is allowed for the Grid.",
	},
	"allow_upgrade_cancel": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines if the Grid is allowed to cancel an upgrade.",
	},
	"allow_upgrade_pause": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines if the Grid is allowed to pause an upgrade.",
	},
	"allow_upgrade_resume": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines if the Grid is allowed to resume an upgrade.",
	},
	"allow_upgrade_scheduling": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determine if the Grid is allowed to schedule an upgrade.",
	},
	"allow_upgrade_test": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines if the Grid is allowed to test an upgrade.",
	},
	"allow_upload": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determine if the Grid is allowed to upload a build.",
	},
	"alternate_version": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The alternative version.",
	},
	"comment": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Comment in readable format for an upgrade group a or virtual node.",
	},
	"current_version": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The current version.",
	},
	"current_version_summary": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Current version summary for the 'type' requested. This field can be requested for the Grid, a certain group that has virtual nodes as subelements, or for the overall group status.",
	},
	"distribution_schedule_active": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines if the distribution schedule is active for the Grid.",
	},
	"distribution_schedule_time": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The Grid master distribution schedule time.",
	},
	"distribution_state": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The current state of distribution process.",
	},
	"distribution_version": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The version that is distributed.",
	},
	"distribution_version_summary": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Distribution version summary for the 'type' requested. This field can be requested for the Grid, a certain group that has virtual nodes as subelements, or for the overall group status.",
	},
	"element_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The status of a certain element with regards to the type requested.",
	},
	"grid_state": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The state of the Grid.",
	},
	"group_state": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The state of a group.",
	},
	"ha_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Status of the HA pair.",
	},
	"hotfixes": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: UpgradestatusHotfixesResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Computed:            true,
		MarkdownDescription: "The list of hotfixes.",
	},
	"ipv4_address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The IPv4 Address of virtual node or physical one.",
	},
	"ipv6_address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The IPv6 Address of virtual node or physical one.",
	},
	"member": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Member that participates in the upgrade process.",
	},
	"message": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The Grid message.",
	},
	"pnode_role": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Status of the physical node in the HA pair.",
	},
	"reverted": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines if the upgrade process is reverted.",
	},
	"status_time": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The status time.",
	},
	"status_value": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Status of a certain group, virtual node or physical node.",
	},
	"status_value_update_time": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Timestamp of when the status was updated.",
	},
	"steps": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: UpgradestatusStepsResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Computed:            true,
		MarkdownDescription: "The list of upgrade process steps.",
	},
	"steps_completed": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of steps done.",
	},
	"steps_total": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Total number steps in the upgrade process.",
	},
	"subelement_type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The type of subelements to be requested. If 'type' is 'GROUP', or 'VNODE', then 'upgrade_group' or 'member' should have proper values for an operation to return data specific for the values passed. Otherwise, overall data is returned for every group or physical node.",
	},
	"subelements_completed": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Number of subelements that have accomplished an upgrade.",
	},
	"subelements_status": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The upgrade process information of subelements.",
	},
	"subelements_total": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Number of subelements number in a certain group, virtual node, or the Grid.",
	},
	"type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The type of upper level elements to be requested.",
	},
	"upgrade_group": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Upgrade group that participates in the upgrade process.",
	},
	"upgrade_schedule_active": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines if the upgrade schedule is active.",
	},
	"upgrade_state": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The upgrade state of the Grid.",
	},
	"upgrade_test_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The upgrade test status of the Grid.",
	},
	"upload_version": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The version that is uploaded.",
	},
	"upload_version_summary": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Upload version summary for the 'type' requested. This field can be requested for the Grid, a certain group that has virtual nodes as subelements, or overall group status.",
	},
}

func ExpandUpgradestatus(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.Upgradestatus {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m UpgradestatusModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *UpgradestatusModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.Upgradestatus {
	if m == nil {
		return nil
	}
	to := &grid.Upgradestatus{
		Ref:               flex.ExpandStringPointer(m.Ref),
		SubelementsStatus: flex.ExpandFrameworkListString(ctx, m.SubelementsStatus, diags),
	}
	return to
}

func FlattenUpgradestatus(ctx context.Context, from *grid.Upgradestatus, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(UpgradestatusAttrTypes)
	}
	m := UpgradestatusModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, UpgradestatusAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *UpgradestatusModel) Flatten(ctx context.Context, from *grid.Upgradestatus, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = UpgradestatusModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.AllowDistribution = types.BoolPointerValue(from.AllowDistribution)
	m.AllowDistributionScheduling = types.BoolPointerValue(from.AllowDistributionScheduling)
	m.AllowUpgrade = types.BoolPointerValue(from.AllowUpgrade)
	m.AllowUpgradeCancel = types.BoolPointerValue(from.AllowUpgradeCancel)
	m.AllowUpgradePause = types.BoolPointerValue(from.AllowUpgradePause)
	m.AllowUpgradeResume = types.BoolPointerValue(from.AllowUpgradeResume)
	m.AllowUpgradeScheduling = types.BoolPointerValue(from.AllowUpgradeScheduling)
	m.AllowUpgradeTest = types.BoolPointerValue(from.AllowUpgradeTest)
	m.AllowUpload = types.BoolPointerValue(from.AllowUpload)
	m.AlternateVersion = flex.FlattenStringPointer(from.AlternateVersion)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CurrentVersion = flex.FlattenStringPointer(from.CurrentVersion)
	m.CurrentVersionSummary = flex.FlattenStringPointer(from.CurrentVersionSummary)
	m.DistributionScheduleActive = types.BoolPointerValue(from.DistributionScheduleActive)
	m.DistributionScheduleTime = flex.FlattenInt64Pointer(from.DistributionScheduleTime)
	m.DistributionState = flex.FlattenStringPointer(from.DistributionState)
	m.DistributionVersion = flex.FlattenStringPointer(from.DistributionVersion)
	m.DistributionVersionSummary = flex.FlattenStringPointer(from.DistributionVersionSummary)
	m.ElementStatus = flex.FlattenStringPointer(from.ElementStatus)
	m.GridState = flex.FlattenStringPointer(from.GridState)
	m.GroupState = flex.FlattenStringPointer(from.GroupState)
	m.HaStatus = flex.FlattenStringPointer(from.HaStatus)
	m.Hotfixes = flex.FlattenFrameworkListNestedBlock(ctx, from.Hotfixes, UpgradestatusHotfixesAttrTypes, diags, FlattenUpgradestatusHotfixes)
	m.Ipv4Address = flex.FlattenStringPointer(from.Ipv4Address)
	m.Ipv6Address = flex.FlattenStringPointer(from.Ipv6Address)
	m.Member = flex.FlattenStringPointer(from.Member)
	m.Message = flex.FlattenStringPointer(from.Message)
	m.PnodeRole = flex.FlattenStringPointer(from.PnodeRole)
	m.Reverted = types.BoolPointerValue(from.Reverted)
	m.StatusTime = flex.FlattenInt64Pointer(from.StatusTime)
	m.StatusValue = flex.FlattenStringPointer(from.StatusValue)
	m.StatusValueUpdateTime = flex.FlattenInt64Pointer(from.StatusValueUpdateTime)
	m.Steps = flex.FlattenFrameworkListNestedBlock(ctx, from.Steps, UpgradestatusStepsAttrTypes, diags, FlattenUpgradestatusSteps)
	m.StepsCompleted = flex.FlattenInt64Pointer(from.StepsCompleted)
	m.StepsTotal = flex.FlattenInt64Pointer(from.StepsTotal)
	m.SubelementType = flex.FlattenStringPointer(from.SubelementType)
	m.SubelementsCompleted = flex.FlattenInt64Pointer(from.SubelementsCompleted)
	m.SubelementsStatus = flex.FlattenFrameworkListString(ctx, from.SubelementsStatus, diags)
	m.SubelementsTotal = flex.FlattenInt64Pointer(from.SubelementsTotal)
	m.Type = flex.FlattenStringPointer(from.Type)
	m.UpgradeGroup = flex.FlattenStringPointer(from.UpgradeGroup)
	m.UpgradeScheduleActive = types.BoolPointerValue(from.UpgradeScheduleActive)
	m.UpgradeState = flex.FlattenStringPointer(from.UpgradeState)
	m.UpgradeTestStatus = flex.FlattenStringPointer(from.UpgradeTestStatus)
	m.UploadVersion = flex.FlattenStringPointer(from.UploadVersion)
	m.UploadVersionSummary = flex.FlattenStringPointer(from.UploadVersionSummary)
}
