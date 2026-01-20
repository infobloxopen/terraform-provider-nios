package misc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/misc"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type ScheduledtaskModel struct {
	Ref                  types.String `tfsdk:"ref"`
	Uuid                 types.String `tfsdk:"uuid"`
	ApprovalStatus       types.String `tfsdk:"approval_status"`
	Approver             types.String `tfsdk:"approver"`
	ApproverComment      types.String `tfsdk:"approver_comment"`
	AutomaticRestart     types.Bool   `tfsdk:"automatic_restart"`
	ChangedObjects       types.List   `tfsdk:"changed_objects"`
	DependentTasks       types.List   `tfsdk:"dependent_tasks"`
	ExecuteNow           types.Bool   `tfsdk:"execute_now"`
	ExecutionDetails     types.List   `tfsdk:"execution_details"`
	ExecutionDetailsType types.String `tfsdk:"execution_details_type"`
	ExecutionStatus      types.String `tfsdk:"execution_status"`
	ExecutionTime        types.Int64  `tfsdk:"execution_time"`
	IsNetworkInsightTask types.Bool   `tfsdk:"is_network_insight_task"`
	Member               types.String `tfsdk:"member"`
	PredecessorTask      types.String `tfsdk:"predecessor_task"`
	ReExecuteTask        types.Bool   `tfsdk:"re_execute_task"`
	ScheduledTime        types.Int64  `tfsdk:"scheduled_time"`
	SubmitTime           types.Int64  `tfsdk:"submit_time"`
	Submitter            types.String `tfsdk:"submitter"`
	SubmitterComment     types.String `tfsdk:"submitter_comment"`
	TaskId               types.Int64  `tfsdk:"task_id"`
	TaskType             types.String `tfsdk:"task_type"`
	TicketNumber         types.String `tfsdk:"ticket_number"`
}

var ScheduledtaskAttrTypes = map[string]attr.Type{
	"ref":                     types.StringType,
	"uuid":                    types.StringType,
	"approval_status":         types.StringType,
	"approver":                types.StringType,
	"approver_comment":        types.StringType,
	"automatic_restart":       types.BoolType,
	"changed_objects":         types.ListType{ElemType: types.ObjectType{AttrTypes: ScheduledtaskChangedObjectsAttrTypes}},
	"dependent_tasks":         types.ListType{ElemType: types.StringType},
	"execute_now":             types.BoolType,
	"execution_details":       types.ListType{ElemType: types.StringType},
	"execution_details_type":  types.StringType,
	"execution_status":        types.StringType,
	"execution_time":          types.Int64Type,
	"is_network_insight_task": types.BoolType,
	"member":                  types.StringType,
	"predecessor_task":        types.StringType,
	"re_execute_task":         types.BoolType,
	"scheduled_time":          types.Int64Type,
	"submit_time":             types.Int64Type,
	"submitter":               types.StringType,
	"submitter_comment":       types.StringType,
	"task_id":                 types.Int64Type,
	"task_type":               types.StringType,
	"ticket_number":           types.StringType,
}

var ScheduledtaskResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The UUID of the object.",
	},
	"approval_status": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The approval status of the task.",
	},
	"approver": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The approver of the task.",
	},
	"approver_comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The comment specified by the approver of the task.",
	},
	"automatic_restart": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Indicates whether there will be an automatic restart when the appliance executes the task.",
	},
	"changed_objects": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ScheduledtaskChangedObjectsResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Computed:            true,
		MarkdownDescription: "A list of objects that are affected by the task.",
	},
	"dependent_tasks": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Computed:            true,
		MarkdownDescription: "If this scheduled task has dependent tasks, their references will be returned in this field.",
	},
	"execute_now": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If this field is set to True the specified task will be executed immediately.",
	},
	"execution_details": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Computed:            true,
		MarkdownDescription: "Messages generated by the execution of the scheduled task after its completion.",
	},
	"execution_details_type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The type of details generated by the execution of the scheduled task after its completion.",
	},
	"execution_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The execution status of the task.",
	},
	"execution_time": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The time when the appliance executed the task.",
	},
	"is_network_insight_task": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Indicates whether this is a Network Insight scheduled task.",
	},
	"member": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The member where this task was created.",
	},
	"predecessor_task": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "If this scheduled task has a predecessor task set, its reference will be returned in this field.",
	},
	"re_execute_task": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If set to True, if the scheduled task is a Network Insight task and it failed, a new task will be cloned from this task and re-executed.",
	},
	"scheduled_time": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The time when the task is scheduled to occur.",
	},
	"submit_time": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The time when the task was submitted.",
	},
	"submitter": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The submitter of the task.",
	},
	"submitter_comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The comment specified by the submitter of the task.",
	},
	"task_id": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The task ID.",
	},
	"task_type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The task type.",
	},
	"ticket_number": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The task ticket number.",
	},
}

func ExpandScheduledtask(ctx context.Context, o types.Object, diags *diag.Diagnostics) *misc.Scheduledtask {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ScheduledtaskModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ScheduledtaskModel) Expand(ctx context.Context, diags *diag.Diagnostics) *misc.Scheduledtask {
	if m == nil {
		return nil
	}
	to := &misc.Scheduledtask{
		Ref:              flex.ExpandStringPointer(m.Ref),
		Uuid:             flex.ExpandStringPointer(m.Uuid),
		ApprovalStatus:   flex.ExpandStringPointer(m.ApprovalStatus),
		ApproverComment:  flex.ExpandStringPointer(m.ApproverComment),
		AutomaticRestart: flex.ExpandBoolPointer(m.AutomaticRestart),
		ExecuteNow:       flex.ExpandBoolPointer(m.ExecuteNow),
		ReExecuteTask:    flex.ExpandBoolPointer(m.ReExecuteTask),
		ScheduledTime:    flex.ExpandInt64Pointer(m.ScheduledTime),
		SubmitterComment: flex.ExpandStringPointer(m.SubmitterComment),
	}
	return to
}

func FlattenScheduledtask(ctx context.Context, from *misc.Scheduledtask, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ScheduledtaskAttrTypes)
	}
	m := ScheduledtaskModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ScheduledtaskAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ScheduledtaskModel) Flatten(ctx context.Context, from *misc.Scheduledtask, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ScheduledtaskModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.ApprovalStatus = flex.FlattenStringPointer(from.ApprovalStatus)
	m.Approver = flex.FlattenStringPointer(from.Approver)
	m.ApproverComment = flex.FlattenStringPointer(from.ApproverComment)
	m.AutomaticRestart = types.BoolPointerValue(from.AutomaticRestart)
	m.ChangedObjects = flex.FlattenFrameworkListNestedBlock(ctx, from.ChangedObjects, ScheduledtaskChangedObjectsAttrTypes, diags, FlattenScheduledtaskChangedObjects)
	m.DependentTasks = flex.FlattenFrameworkListString(ctx, from.DependentTasks, diags)
	m.ExecuteNow = types.BoolPointerValue(from.ExecuteNow)
	m.ExecutionDetails = flex.FlattenFrameworkListString(ctx, from.ExecutionDetails, diags)
	m.ExecutionDetailsType = flex.FlattenStringPointer(from.ExecutionDetailsType)
	m.ExecutionStatus = flex.FlattenStringPointer(from.ExecutionStatus)
	m.ExecutionTime = flex.FlattenInt64Pointer(from.ExecutionTime)
	m.IsNetworkInsightTask = types.BoolPointerValue(from.IsNetworkInsightTask)
	m.Member = flex.FlattenStringPointer(from.Member)
	m.PredecessorTask = flex.FlattenStringPointer(from.PredecessorTask)
	m.ReExecuteTask = types.BoolPointerValue(from.ReExecuteTask)
	m.ScheduledTime = flex.FlattenInt64Pointer(from.ScheduledTime)
	m.SubmitTime = flex.FlattenInt64Pointer(from.SubmitTime)
	m.Submitter = flex.FlattenStringPointer(from.Submitter)
	m.SubmitterComment = flex.FlattenStringPointer(from.SubmitterComment)
	m.TaskId = flex.FlattenInt64Pointer(from.TaskId)
	m.TaskType = flex.FlattenStringPointer(from.TaskType)
	m.TicketNumber = flex.FlattenStringPointer(from.TicketNumber)
}
