package cloud

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/infoblox-nios-go-client/cloud"

	"github.com/infobloxopen/terraform-provider-nios/internal/config"
	"github.com/infobloxopen/terraform-provider-nios/internal/retry"
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

	// Process AWS account IDs file if multiple_accounts_sync_policy is UPLOAD_CHILDREN
	if !data.MultipleAccountsSyncPolicy.IsNull() && data.MultipleAccountsSyncPolicy.ValueString() == "UPLOAD_CHILDREN" {
		if !r.processAwsAccountIdsFile(ctx, &data, &resp.Diagnostics) {
			return
		}
	}

	var apiRes *cloud.CreateAwsrte53taskgroupResponse

	err := retry.Do(ctx, retry.TransientErrors, func(ctx context.Context) (int, error) {
		var (
			httpRes *http.Response
			callErr error
		)
		apiRes, httpRes, callErr = r.client.CloudAPI.
			Awsrte53taskgroupAPI.
			Create(ctx).
			Awsrte53taskgroup(*data.Expand(ctx, &resp.Diagnostics, true)).
			ReturnFieldsPlus(readableAttributesForAwsrte53taskgroup).
			ReturnAsObject(1).
			Execute()

		if httpRes != nil {
			return httpRes.StatusCode, callErr
		}
		return 0, callErr
	})

	if err != nil {
		if retry.IsAlreadyExistsErr(err) {
			// Resource already exists, import required
			resp.Diagnostics.AddError(
				"Resource Already Exists",
				fmt.Sprintf("Resource already exists, error: %s.\nPlease import the existing resource into terraform state.", err.Error()),
			)
			return
		}
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

	resourceRef := utils.ExtractResourceRef(data.Ref.ValueString())

	var (
		httpRes *http.Response
		apiRes  *cloud.GetAwsrte53taskgroupResponse
	)

	err := retry.Do(ctx, nil, func(ctx context.Context) (int, error) {
		var callErr error
		apiRes, httpRes, callErr = r.client.CloudAPI.
			Awsrte53taskgroupAPI.
			Read(ctx, resourceRef).
			ReturnFieldsPlus(readableAttributesForAwsrte53taskgroup).
			ReturnAsObject(1).
			ProxySearch(config.GetProxySearch()).
			Execute()

		if httpRes != nil {
			return httpRes.StatusCode, callErr
		}
		return 0, callErr
	})

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

	// Process AWS account IDs file if multiple_accounts_sync_policy is UPLOAD_CHILDREN
	if !data.MultipleAccountsSyncPolicy.IsNull() && data.MultipleAccountsSyncPolicy.ValueString() == "UPLOAD_CHILDREN" {
		if !r.processAwsAccountIdsFile(ctx, &data, &resp.Diagnostics) {
			return
		}
	}

	diags = req.State.GetAttribute(ctx, path.Root("ref"), &data.Ref)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	resourceRef := utils.ExtractResourceRef(data.Ref.ValueString())

	var apiRes *cloud.UpdateAwsrte53taskgroupResponse

	err := retry.Do(ctx, retry.TransientErrors, func(ctx context.Context) (int, error) {
		var (
			httpRes *http.Response
			callErr error
		)
		apiRes, httpRes, callErr = r.client.CloudAPI.
			Awsrte53taskgroupAPI.
			Update(ctx, resourceRef).
			Awsrte53taskgroup(*data.Expand(ctx, &resp.Diagnostics, false)).
			ReturnFieldsPlus(readableAttributesForAwsrte53taskgroup).
			ReturnAsObject(1).
			Execute()

		if httpRes != nil {
			return httpRes.StatusCode, callErr
		}
		return 0, callErr
	})

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

	resourceRef := utils.ExtractResourceRef(data.Ref.ValueString())

	err := retry.Do(ctx, retry.TransientErrors, func(ctx context.Context) (int, error) {
		httpRes, callErr := r.client.CloudAPI.
			Awsrte53taskgroupAPI.
			Delete(ctx, resourceRef).
			Execute()

		if httpRes != nil {
			if httpRes.StatusCode == http.StatusNotFound {
				return 0, nil
			}
			return httpRes.StatusCode, callErr
		}
		return 0, callErr
	})

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Awsrte53taskgroup, got error: %s", err))
		return
	}
}

func (r *Awsrte53taskgroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("ref"), req, resp)
}

// function that will process your AWS account IDs file and return the token
func (r *Awsrte53taskgroupResource) processAwsAccountIdsFile(ctx context.Context, data *Awsrte53taskgroupModel, diags *diag.Diagnostics) bool {
	// Check if aws_account_ids_file_path is provided
	if data.AwsAccountIdsFilePath.IsNull() || data.AwsAccountIdsFilePath.IsUnknown() {
		return true // No file to process, continue
	}

	// Get connection details from client configuration
	baseUrl := r.client.CloudAPI.Cfg.NIOSHostURL
	username := r.client.CloudAPI.Cfg.NIOSUsername
	password := r.client.CloudAPI.Cfg.NIOSPassword

	// Get the file path from the model
	filePath := data.AwsAccountIdsFilePath.ValueString()

	// Upload the AWS account IDs file and get the token
	token, err := utils.UploadFileWithToken(ctx, baseUrl, filePath, username, password)
	if err != nil {
		diags.AddError(
			"Client Error",
			fmt.Sprintf("Unable to process AWS account IDs file %s, got error: %s", filePath, err),
		)
		return false
	}

	// Store the token in the aws_account_ids_file_token field
	data.AwsAccountIdsFileToken = types.StringValue(token)
	return true
}

func (r *Awsrte53taskgroupResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data Awsrte53taskgroupModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// For Configuration object, any attributes not defined by the user appear as null, unless derived from another instance.
	// We perform IsUnknown() check to handle variables from .tfvars that are resolved
	// during the plan phase rather than validation phase, preventing false validation errors.

	// Validation 1: Handle filter validation
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

	// Validation 2: aws_account_ids_file_path can only be used with UPLOAD_CHILDREN policy
	if !data.AwsAccountIdsFilePath.IsUnknown() && !data.MultipleAccountsSyncPolicy.IsUnknown() {
		if !data.AwsAccountIdsFilePath.IsNull() && data.AwsAccountIdsFilePath.ValueString() != "" {
			policyValue := "NONE"
			if !data.MultipleAccountsSyncPolicy.IsNull() {
				policyValue = data.MultipleAccountsSyncPolicy.ValueString()
			}
			if policyValue != "UPLOAD_CHILDREN" {
				resp.Diagnostics.AddError(
					"Invalid Configuration",
					"'aws_account_ids_file_path' can only be used when 'multiple_accounts_sync_policy' is set to 'UPLOAD_CHILDREN'. "+
						"Current policy is '"+policyValue+"'. "+
						"Either remove 'aws_account_ids_file_path' or set 'multiple_accounts_sync_policy' to 'UPLOAD_CHILDREN'.",
				)
			}
		}
	}

	// Validation 3: UPLOAD_CHILDREN policy requires aws_account_ids_file_path
	if !data.MultipleAccountsSyncPolicy.IsUnknown() {
		if !data.MultipleAccountsSyncPolicy.IsNull() && data.MultipleAccountsSyncPolicy.ValueString() == "UPLOAD_CHILDREN" {
			if !data.AwsAccountIdsFilePath.IsUnknown() {
				if data.AwsAccountIdsFilePath.IsNull() || data.AwsAccountIdsFilePath.ValueString() == "" {
					resp.Diagnostics.AddError(
						"Invalid Configuration",
						"When 'multiple_accounts_sync_policy' is 'UPLOAD_CHILDREN', 'aws_account_ids_file_path' must be provided. "+
							"Please specify the path to a file containing AWS account IDs.",
					)
				}
			}
		}
	}

	// Validation 4: sync_child_accounts and role_arn relationship
	// Only validate if sync_child_accounts is known (not unknown)
	if !data.SyncChildAccounts.IsUnknown() {
		if !data.SyncChildAccounts.IsNull() && data.SyncChildAccounts.ValueBool() {
			if !data.RoleArn.IsUnknown() {
				if data.RoleArn.IsNull() || data.RoleArn.ValueString() == "" {
					resp.Diagnostics.AddError(
						"Invalid Configuration",
						"When 'sync_child_accounts' is enabled, 'role_arn' must be provided and cannot be empty. "+
							"Please provide a valid AWS IAM role ARN for accessing child accounts.",
					)
				}
			}
		}
	}
}
