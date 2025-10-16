package cloud

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

var readableAttributesForAwsrte53taskgroup = "account_id,comment,consolidate_zones,consolidated_view,disabled,grid_member,name,network_view,network_view_mapping_policy,role_arn,sync_child_accounts,sync_status,task_list,multiple_accounts_sync_policy"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &Awsrte53taskgroupResource{}
var _ resource.ResourceWithImportState = &Awsrte53taskgroupResource{}

func NewAwsrte53taskgroupResource() resource.Resource {
	return &Awsrte53taskgroupResource{}
}

// Awsrte53taskgroupResource defines the resource implementation.
type Awsrte53taskgroupResource struct {
	client *niosclient.APIClient
}

func (r *Awsrte53taskgroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "cloud_aws_route53_task_group"
}

func (r *Awsrte53taskgroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages an AWS Route 53 Task Group.",
		Attributes:          Awsrte53taskgroupResourceSchemaAttributes,
	}
}

func (r *Awsrte53taskgroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *Awsrte53taskgroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Awsrte53taskgroupModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.CloudAPI.
		Awsrte53taskgroupAPI.
		Create(ctx).
		Awsrte53taskgroup(*data.Expand(ctx, &resp.Diagnostics, true)).
		ReturnFieldsPlus(readableAttributesForAwsrte53taskgroup).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Awsrte53taskgroup, got error: %s", err))
		return
	}

	res := apiRes.CreateAwsrte53taskgroupResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Awsrte53taskgroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Awsrte53taskgroupModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.CloudAPI.
		Awsrte53taskgroupAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForAwsrte53taskgroup).
		ReturnAsObject(1).
		Execute()

	// Handle not found case
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			// Resource no longer exists, remove from state
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Awsrte53taskgroup, got error: %s", err))
		return
	}

	res := apiRes.GetAwsrte53taskgroupResponseObjectAsResult.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Awsrte53taskgroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data Awsrte53taskgroupModel

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

	apiRes, _, err := r.client.CloudAPI.
		Awsrte53taskgroupAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Awsrte53taskgroup(*data.Expand(ctx, &resp.Diagnostics, false)).
		ReturnFieldsPlus(readableAttributesForAwsrte53taskgroup).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Awsrte53taskgroup, got error: %s", err))
		return
	}

	res := apiRes.UpdateAwsrte53taskgroupResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Awsrte53taskgroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Awsrte53taskgroupModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.CloudAPI.
		Awsrte53taskgroupAPI.
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Awsrte53taskgroup, got error: %s", err))
		return
	}
}

func (r *Awsrte53taskgroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("ref"), req, resp)
}

func (r *Awsrte53taskgroupResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data Awsrte53taskgroupModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Handle filter validation - warn users about empty string
	if !data.TaskList.IsNull() && !data.TaskList.IsUnknown() {
		var taskList []Awsrte53taskgroupTaskListModel
		diags := data.TaskList.ElementsAs(ctx, &taskList, false)
		if !diags.HasError() {
			for i, task := range taskList {
				// Check if filter is known and not null before checking if it's empty
				if !task.Filter.IsUnknown() && !task.Filter.IsNull() && task.Filter.ValueString() == "" {
					resp.Diagnostics.AddError(
						"Invalid Filter Configuration",
						fmt.Sprintf("task_list[%d].filter cannot be empty string. Use '*' for wildcard or omit the filter attribute.", i),
					)
				}
			}
		}
	}

	syncChildAccounts := data.SyncChildAccounts
	roleArn := data.RoleArn

	// Skip validation if values are unknown (during planning)
	if syncChildAccounts.IsUnknown() {
		return
	}

	// If sync_child_accounts is true, role_arn must be provided and non-empty
	if !syncChildAccounts.IsNull() && syncChildAccounts.ValueBool() {
		if roleArn.IsUnknown() || roleArn.IsNull() || roleArn.ValueString() == "" {
			resp.Diagnostics.AddError(
				"Invalid Configuration",
				"When 'sync_child_accounts' is enabled, 'role_arn' must be provided and cannot be empty. "+
					"Please provide a valid AWS IAM role ARN for accessing child accounts.",
			)
		}
	}
}
