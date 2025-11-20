package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/import"
)

type ApprovalworkflowModel struct {
	Ref                     types.String `tfsdk:"ref"`
    Uuid        types.String `tfsdk:"uuid"`
	ApprovalGroup           types.String `tfsdk:"approval_group"`
	ApprovalNotifyTo        types.String `tfsdk:"approval_notify_to"`
	ApprovedNotifyTo        types.String `tfsdk:"approved_notify_to"`
	ApproverComment         types.String `tfsdk:"approver_comment"`
	EnableApprovalNotify    types.Bool   `tfsdk:"enable_approval_notify"`
	EnableApprovedNotify    types.Bool   `tfsdk:"enable_approved_notify"`
	EnableFailedNotify      types.Bool   `tfsdk:"enable_failed_notify"`
	EnableNotifyGroup       types.Bool   `tfsdk:"enable_notify_group"`
	EnableNotifyUser        types.Bool   `tfsdk:"enable_notify_user"`
	EnableRejectedNotify    types.Bool   `tfsdk:"enable_rejected_notify"`
	EnableRescheduledNotify types.Bool   `tfsdk:"enable_rescheduled_notify"`
	EnableSucceededNotify   types.Bool   `tfsdk:"enable_succeeded_notify"`
	ExtAttrs                types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll             types.Map    `tfsdk:"extattrs_all"`
	FailedNotifyTo          types.String `tfsdk:"failed_notify_to"`
	RejectedNotifyTo        types.String `tfsdk:"rejected_notify_to"`
	RescheduledNotifyTo     types.String `tfsdk:"rescheduled_notify_to"`
	SubmitterComment        types.String `tfsdk:"submitter_comment"`
	SubmitterGroup          types.String `tfsdk:"submitter_group"`
	SucceededNotifyTo       types.String `tfsdk:"succeeded_notify_to"`
	TicketNumber            types.String `tfsdk:"ticket_number"`
}

var ApprovalworkflowAttrTypes = map[string]attr.Type{
	"ref":                       types.StringType,
    "uuid":        types.StringType,
	"approval_group":            types.StringType,
	"approval_notify_to":        types.StringType,
	"approved_notify_to":        types.StringType,
	"approver_comment":          types.StringType,
	"enable_approval_notify":    types.BoolType,
	"enable_approved_notify":    types.BoolType,
	"enable_failed_notify":      types.BoolType,
	"enable_notify_group":       types.BoolType,
	"enable_notify_user":        types.BoolType,
	"enable_rejected_notify":    types.BoolType,
	"enable_rescheduled_notify": types.BoolType,
	"enable_succeeded_notify":   types.BoolType,
	"extattrs":                  types.MapType{ElemType: types.StringType},
	"extattrs_all":              types.MapType{ElemType: types.StringType},
	"failed_notify_to":          types.StringType,
	"rejected_notify_to":        types.StringType,
	"rescheduled_notify_to":     types.StringType,
	"submitter_comment":         types.StringType,
	"submitter_group":           types.StringType,
	"succeeded_notify_to":       types.StringType,
	"ticket_number":             types.StringType,
}

var ApprovalworkflowResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
    "uuid": schema.StringAttribute{
        Computed:            true,
        MarkdownDescription: "The uuid to the object.",
    },
	"approval_group": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The approval administration group.",
	},
	"approval_notify_to": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The destination for approval task notifications.",
	},
	"approved_notify_to": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The destination for approved task notifications.",
	},
	"approver_comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The requirement for the comment when an approver approves a submitted task.",
	},
	"enable_approval_notify": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether approval task notifications are enabled.",
	},
	"enable_approved_notify": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether approved task notifications are enabled.",
	},
	"enable_failed_notify": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether failed task notifications are enabled.",
	},
	"enable_notify_group": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether e-mail notifications to admin group's e-mail address are enabled.",
	},
	"enable_notify_user": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether e-mail notifications to an admin member's e-mail address are enabled.",
	},
	"enable_rejected_notify": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether rejected task notifications are enabled.",
	},
	"enable_rescheduled_notify": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether rescheduled task notifications are enabled.",
	},
	"enable_succeeded_notify": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether succeeded task notifications are enabled.",
	},
	"extattrs": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.",
	},
	"extattrs_all": schema.MapAttribute{
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object , including default attributes.",
		ElementType:         types.StringType,
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
	},
	"failed_notify_to": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The destination for failed task notifications.",
	},
	"rejected_notify_to": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The destination for rejected task notifications.",
	},
	"rescheduled_notify_to": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The destination for rescheduled task notifications.",
	},
	"submitter_comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The requirement for the comment when a submitter submits a task for approval.",
	},
	"submitter_group": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The submitter admininstration group.",
	},
	"succeeded_notify_to": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The destination for succeeded task notifications.",
	},
	"ticket_number": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The requirement for the ticket number when a submitter submits a task for approval.",
	},
}

func ExpandApprovalworkflow(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.Approvalworkflow {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ApprovalworkflowModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ApprovalworkflowModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.Approvalworkflow {
	if m == nil {
		return nil
	}
	to := &security.Approvalworkflow{
		Ref:                     flex.ExpandStringPointer(m.Ref),
		ApprovalGroup:           flex.ExpandStringPointer(m.ApprovalGroup),
		ApprovalNotifyTo:        flex.ExpandStringPointer(m.ApprovalNotifyTo),
		ApprovedNotifyTo:        flex.ExpandStringPointer(m.ApprovedNotifyTo),
		ApproverComment:         flex.ExpandStringPointer(m.ApproverComment),
		EnableApprovalNotify:    flex.ExpandBoolPointer(m.EnableApprovalNotify),
		EnableApprovedNotify:    flex.ExpandBoolPointer(m.EnableApprovedNotify),
		EnableFailedNotify:      flex.ExpandBoolPointer(m.EnableFailedNotify),
		EnableNotifyGroup:       flex.ExpandBoolPointer(m.EnableNotifyGroup),
		EnableNotifyUser:        flex.ExpandBoolPointer(m.EnableNotifyUser),
		EnableRejectedNotify:    flex.ExpandBoolPointer(m.EnableRejectedNotify),
		EnableRescheduledNotify: flex.ExpandBoolPointer(m.EnableRescheduledNotify),
		EnableSucceededNotify:   flex.ExpandBoolPointer(m.EnableSucceededNotify),
		ExtAttrs:                ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		FailedNotifyTo:          flex.ExpandStringPointer(m.FailedNotifyTo),
		RejectedNotifyTo:        flex.ExpandStringPointer(m.RejectedNotifyTo),
		RescheduledNotifyTo:     flex.ExpandStringPointer(m.RescheduledNotifyTo),
		SubmitterComment:        flex.ExpandStringPointer(m.SubmitterComment),
		SubmitterGroup:          flex.ExpandStringPointer(m.SubmitterGroup),
		SucceededNotifyTo:       flex.ExpandStringPointer(m.SucceededNotifyTo),
		TicketNumber:            flex.ExpandStringPointer(m.TicketNumber),
	}
	return to
}

func FlattenApprovalworkflow(ctx context.Context, from *security.Approvalworkflow, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ApprovalworkflowAttrTypes)
	}
	m := ApprovalworkflowModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, ApprovalworkflowAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ApprovalworkflowModel) Flatten(ctx context.Context, from *security.Approvalworkflow, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ApprovalworkflowModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
    m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.ApprovalGroup = flex.FlattenStringPointer(from.ApprovalGroup)
	m.ApprovalNotifyTo = flex.FlattenStringPointer(from.ApprovalNotifyTo)
	m.ApprovedNotifyTo = flex.FlattenStringPointer(from.ApprovedNotifyTo)
	m.ApproverComment = flex.FlattenStringPointer(from.ApproverComment)
	m.EnableApprovalNotify = types.BoolPointerValue(from.EnableApprovalNotify)
	m.EnableApprovedNotify = types.BoolPointerValue(from.EnableApprovedNotify)
	m.EnableFailedNotify = types.BoolPointerValue(from.EnableFailedNotify)
	m.EnableNotifyGroup = types.BoolPointerValue(from.EnableNotifyGroup)
	m.EnableNotifyUser = types.BoolPointerValue(from.EnableNotifyUser)
	m.EnableRejectedNotify = types.BoolPointerValue(from.EnableRejectedNotify)
	m.EnableRescheduledNotify = types.BoolPointerValue(from.EnableRescheduledNotify)
	m.EnableSucceededNotify = types.BoolPointerValue(from.EnableSucceededNotify)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.FailedNotifyTo = flex.FlattenStringPointer(from.FailedNotifyTo)
	m.RejectedNotifyTo = flex.FlattenStringPointer(from.RejectedNotifyTo)
	m.RescheduledNotifyTo = flex.FlattenStringPointer(from.RescheduledNotifyTo)
	m.SubmitterComment = flex.FlattenStringPointer(from.SubmitterComment)
	m.SubmitterGroup = flex.FlattenStringPointer(from.SubmitterGroup)
	m.SucceededNotifyTo = flex.FlattenStringPointer(from.SucceededNotifyTo)
	m.TicketNumber = flex.FlattenStringPointer(from.TicketNumber)
}
