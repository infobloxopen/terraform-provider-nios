package notification

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"

	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForNotificationRule = "all_members,comment,disable,enable_event_deduplication,enable_event_deduplication_log,event_deduplication_fields,event_deduplication_lookback_period,event_priority,event_type,expression_list,name,notification_action,notification_target,publish_settings,scheduled_event,selected_members,template_instance,use_publish_settings"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NotificationRuleResource{}
var _ resource.ResourceWithImportState = &NotificationRuleResource{}
var _ resource.ResourceWithValidateConfig = &NotificationRuleResource{}

func NewNotificationRuleResource() resource.Resource {
	return &NotificationRuleResource{}
}

// NotificationRuleResource defines the resource implementation.
type NotificationRuleResource struct {
	client *niosclient.APIClient
}

func (r *NotificationRuleResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "notification_rule"
}

func (r *NotificationRuleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a Notification Rule.",
		Attributes:          NotificationRuleResourceSchemaAttributes,
	}
}

func (r *NotificationRuleResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*niosclient.APIClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *niosclient.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *NotificationRuleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NotificationRuleModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.NotificationAPI.
		NotificationRuleAPI.
		Create(ctx).
		NotificationRule(*data.Expand(ctx, &resp.Diagnostics, true)).
		ReturnFieldsPlus(readableAttributesForNotificationRule).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create NotificationRule, got error: %s", err))
		return
	}

	res := apiRes.CreateNotificationRuleResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NotificationRuleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NotificationRuleModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.NotificationAPI.
		NotificationRuleAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForNotificationRule).
		ReturnAsObject(1).
		Execute()

	// Handle not found case
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			// Resource no longer exists, remove from state
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read NotificationRule, got error: %s", err))
		return
	}

	res := apiRes.GetNotificationRuleResponseObjectAsResult.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NotificationRuleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data NotificationRuleModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	diags = req.State.GetAttribute(ctx, path.Root("ref"), &data.Ref)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	apiRes, _, err := r.client.NotificationAPI.
		NotificationRuleAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		NotificationRule(*data.Expand(ctx, &resp.Diagnostics, false)).
		ReturnFieldsPlus(readableAttributesForNotificationRule).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update NotificationRule, got error: %s", err))
		return
	}

	res := apiRes.UpdateNotificationRuleResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NotificationRuleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NotificationRuleModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.NotificationAPI.
		NotificationRuleAPI.
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete NotificationRule, got error: %s", err))
		return
	}
}

func (r *NotificationRuleResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data NotificationRuleModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Validate scheduled_event configuration
	if !data.ScheduledEvent.IsNull() && !data.ScheduledEvent.IsUnknown() {
		utils.ValidateScheduleConfig(
			data.ScheduledEvent,
			"",
			path.Root("scheduled_event"),
			&resp.Diagnostics,
		)
	}
}

func (r *NotificationRuleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("ref"), req, resp)
}
