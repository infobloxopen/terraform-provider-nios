package cloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/cloud"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	planmodifiers "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/immutable"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type Awsrte53taskgroupModel struct {
	Ref                        types.String `tfsdk:"ref"`
	AccountId                  types.String `tfsdk:"account_id"`
	AccountsList               types.String `tfsdk:"accounts_list"`
	AwsAccountIdsFileToken     types.String `tfsdk:"aws_account_ids_file_token"`
	AwsAccountIdsFilePath      types.String `tfsdk:"aws_account_ids_file_path"`
	Comment                    types.String `tfsdk:"comment"`
	ConsolidateZones           types.Bool   `tfsdk:"consolidate_zones"`
	ConsolidatedView           types.String `tfsdk:"consolidated_view"`
	Disabled                   types.Bool   `tfsdk:"disabled"`
	GridMember                 types.String `tfsdk:"grid_member"`
	MultipleAccountsSyncPolicy types.String `tfsdk:"multiple_accounts_sync_policy"`
	Name                       types.String `tfsdk:"name"`
	NetworkView                types.String `tfsdk:"network_view"`
	NetworkViewMappingPolicy   types.String `tfsdk:"network_view_mapping_policy"`
	RoleArn                    types.String `tfsdk:"role_arn"`
	SyncChildAccounts          types.Bool   `tfsdk:"sync_child_accounts"`
	SyncStatus                 types.String `tfsdk:"sync_status"`
	TaskList                   types.List   `tfsdk:"task_list"`
}

var Awsrte53taskgroupAttrTypes = map[string]attr.Type{
	"ref":                           types.StringType,
	"account_id":                    types.StringType,
	"accounts_list":                 types.StringType,
	"aws_account_ids_file_token":    types.StringType,
	"aws_account_ids_file_path":     types.StringType,
	"comment":                       types.StringType,
	"consolidate_zones":             types.BoolType,
	"consolidated_view":             types.StringType,
	"disabled":                      types.BoolType,
	"grid_member":                   types.StringType,
	"multiple_accounts_sync_policy": types.StringType,
	"name":                          types.StringType,
	"network_view":                  types.StringType,
	"network_view_mapping_policy":   types.StringType,
	"role_arn":                      types.StringType,
	"sync_child_accounts":           types.BoolType,
	"sync_status":                   types.StringType,
	"task_list":                     types.ListType{ElemType: types.ObjectType{AttrTypes: Awsrte53taskgroupTaskListAttrTypes}},
}

var Awsrte53taskgroupResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"account_id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The AWS Account ID associated with this task group.",
	},
	"accounts_list": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The AWS Account IDs list associated with this task group.",
	},
	"aws_account_ids_file_token": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The AWS account IDs file's token.",
	},
	"aws_account_ids_file_path": schema.StringAttribute{
		Optional:            true, // change this  based on validate config
		MarkdownDescription: "The AWS account IDs file's path.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 256),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Comment for the task group; maximum 256 characters.",
	},
	"consolidate_zones": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Indicates if all zones need to be saved into a single view.",
		PlanModifiers: []planmodifier.Bool{
			planmodifiers.ImmutableBool(),
		},
	},
	"consolidated_view": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("consolidate_zones")),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of the DNS view for consolidating zones.",
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
	},
	"disabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Indicates if the task group is enabled or disabled.",
	},
	"grid_member": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Member on which the tasks in this task group will be run.",
	},
	"multiple_accounts_sync_policy": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("NONE"),
		Validators: []validator.String{
			stringvalidator.OneOf("DISCOVER_CHILDREN", "NONE", "UPLOAD_CHILDREN"),
		},
		MarkdownDescription: "Discover all child accounts or Upload child account ids to discover..",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of this AWS Route53 sync task group.",
	},
	"network_view": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of the tenant's network view.",
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
	},
	"network_view_mapping_policy": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("AUTO_CREATE"),
		Validators: []validator.String{
			stringvalidator.OneOf("AUTO_CREATE", "DIRECT"),
		},
		MarkdownDescription: "The network view mapping policy.",
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
	},
	"role_arn": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Role ARN for syncing child accounts; maximum 128 characters.",
	},
	"sync_child_accounts": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Synchronizing child accounts is enabled or disabled.",
	},
	"sync_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Indicate the overall sync status of this task group.",
	},
	"task_list": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Awsrte53taskgroupTaskListResourceSchemaAttributes,
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "List of AWS Route53 tasks in this group.",
	},
}

func (m *Awsrte53taskgroupModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *cloud.Awsrte53taskgroup {
	if m == nil {
		return nil
	}
	to := &cloud.Awsrte53taskgroup{
		AwsAccountIdsFileToken:     flex.ExpandStringPointer(m.AwsAccountIdsFileToken),
		Comment:                    flex.ExpandStringPointer(m.Comment),
		Disabled:                   flex.ExpandBoolPointer(m.Disabled),
		GridMember:                 flex.ExpandStringPointer(m.GridMember),
		MultipleAccountsSyncPolicy: flex.ExpandStringPointer(m.MultipleAccountsSyncPolicy),
		Name:                       flex.ExpandStringPointer(m.Name),
		RoleArn:                    flex.ExpandStringPointer(m.RoleArn),
		SyncChildAccounts:          flex.ExpandBoolPointer(m.SyncChildAccounts),
		TaskList:                   flex.ExpandFrameworkListNestedBlock(ctx, m.TaskList, diags, ExpandAwsrte53taskgroupTaskList),
	}
	if isCreate {
		to.ConsolidateZones = flex.ExpandBoolPointer(m.ConsolidateZones)
		to.NetworkView = flex.ExpandStringPointer(m.NetworkView)
		to.ConsolidatedView = flex.ExpandStringPointer(m.ConsolidatedView)
		to.NetworkViewMappingPolicy = flex.ExpandStringPointer(m.NetworkViewMappingPolicy)
	}
	return to
}

func FlattenAwsrte53taskgroup(ctx context.Context, from *cloud.Awsrte53taskgroup, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Awsrte53taskgroupAttrTypes)
	}
	m := Awsrte53taskgroupModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Awsrte53taskgroupAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Awsrte53taskgroupModel) Flatten(ctx context.Context, from *cloud.Awsrte53taskgroup, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Awsrte53taskgroupModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.AccountId = flex.FlattenStringPointer(from.AccountId)
	m.AccountsList = flex.FlattenStringPointer(from.AccountsList)
	m.AwsAccountIdsFileToken = flex.FlattenStringPointer(from.AwsAccountIdsFileToken)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.ConsolidateZones = types.BoolPointerValue(from.ConsolidateZones)
	m.ConsolidatedView = flex.FlattenStringPointer(from.ConsolidatedView)
	m.Disabled = types.BoolPointerValue(from.Disabled)
	m.GridMember = flex.FlattenStringPointer(from.GridMember)
	m.MultipleAccountsSyncPolicy = flex.FlattenStringPointer(from.MultipleAccountsSyncPolicy)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.NetworkView = flex.FlattenStringPointer(from.NetworkView)
	m.NetworkViewMappingPolicy = flex.FlattenStringPointer(from.NetworkViewMappingPolicy)
	m.RoleArn = flex.FlattenStringPointer(from.RoleArn)
	m.SyncChildAccounts = types.BoolPointerValue(from.SyncChildAccounts)
	m.SyncStatus = flex.FlattenStringPointer(from.SyncStatus)
	planList := m.TaskList
	m.TaskList = flex.FlattenFrameworkListNestedBlock(ctx, from.TaskList, Awsrte53taskgroupTaskListAttrTypes, diags, FlattenAwsrte53taskgroupTaskList)
	if !planList.IsUnknown() {
		reOrderedList, diags := utils.ReorderAndFilterNestedListResponse(ctx, planList, m.TaskList, "name")
		if !diags.HasError() {
			m.TaskList = reOrderedList.(basetypes.ListValue)
		}
	}
}
