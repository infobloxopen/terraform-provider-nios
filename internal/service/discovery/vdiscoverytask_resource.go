package discovery

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForVdiscoverytask = "accounts_list,allow_unsecured_connection,auto_consolidate_cloud_ea,auto_consolidate_managed_tenant,auto_consolidate_managed_vm,auto_create_dns_hostname_template,auto_create_dns_record,auto_create_dns_record_type,cdiscovery_file_token,comment,credentials_type,dns_view_private_ip,dns_view_public_ip,domain_name,driver_type,enable_filter,enabled,fqdn_or_ip,govcloud_enabled,identity_version,last_run,member,merge_data,multiple_accounts_sync_policy,name,network_filter,network_list,port,private_network_view,private_network_view_mapping_policy,protocol,public_network_view,public_network_view_mapping_policy,role_arn,scheduled_run,selected_regions,service_account_file,service_account_file_token,state,state_msg,sync_child_accounts,update_dns_view_private_ip,update_dns_view_public_ip,update_metadata,use_identity,username"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &VdiscoverytaskResource{}
var _ resource.ResourceWithImportState = &VdiscoverytaskResource{}

func NewVdiscoverytaskResource() resource.Resource {
	return &VdiscoverytaskResource{}
}

// VdiscoverytaskResource defines the resource implementation.
type VdiscoverytaskResource struct {
	client *niosclient.APIClient
}

func (r *VdiscoverytaskResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "discovery_vdiscoverytask"
}

func (r *VdiscoverytaskResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a Vdiscoverytask.",
		Attributes:          VdiscoverytaskResourceSchemaAttributes,
	}
}

func (r *VdiscoverytaskResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *VdiscoverytaskResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VdiscoverytaskModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Process GCP service account file if provided
	if data.DriverType.ValueString() == "GCP" {
		if !r.processGCPServiceAccountFile(ctx, &data, &resp.Diagnostics) {
			return
		}
	}

	// Process CDiscovery file if multiple_accounts_sync_policy is UPLOAD
	if !data.MultipleAccountsSyncPolicy.IsNull() && data.MultipleAccountsSyncPolicy.ValueString() == "UPLOAD" {
		if !r.processCDiscoveryFile(ctx, &data, &resp.Diagnostics) {
			return
		}
	}
	apiRes, _, err := r.client.DiscoveryAPI.
		VdiscoverytaskAPI.
		Create(ctx).
		Vdiscoverytask(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForVdiscoverytask).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Vdiscoverytask, got error: %s", err))
		return
	}

	res := apiRes.CreateVdiscoverytaskResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VdiscoverytaskResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VdiscoverytaskModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.DiscoveryAPI.
		VdiscoverytaskAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForVdiscoverytask).
		ReturnAsObject(1).
		Execute()

		// Handle not found case
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			// Resource no longer exists, remove from state
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Vdiscoverytask, got error: %s", err))
		return
	}

	res := apiRes.GetVdiscoverytaskResponseObjectAsResult.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VdiscoverytaskResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data VdiscoverytaskModel

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

	// Process GCP service account file if provided
	if data.DriverType.ValueString() == "GCP" {
		if !r.processGCPServiceAccountFile(ctx, &data, &resp.Diagnostics) {
			return
		}
	}

	// Process CDiscovery file if multiple_accounts_sync_policy is UPLOAD
	if !data.MultipleAccountsSyncPolicy.IsNull() && data.MultipleAccountsSyncPolicy.ValueString() == "UPLOAD" {
		if !r.processCDiscoveryFile(ctx, &data, &resp.Diagnostics) {
			return
		}
	}

	apiRes, _, err := r.client.DiscoveryAPI.
		VdiscoverytaskAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Vdiscoverytask(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForVdiscoverytask).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Vdiscoverytask, got error: %s", err))
		return
	}

	res := apiRes.UpdateVdiscoverytaskResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VdiscoverytaskResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VdiscoverytaskModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.DiscoveryAPI.
		VdiscoverytaskAPI.
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Vdiscoverytask, got error: %s", err))
		return
	}
}

func (r *VdiscoverytaskResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("ref"), req, resp)
}

// function that will process your GCP service account file and return the token
func (r *VdiscoverytaskResource) processGCPServiceAccountFile(ctx context.Context, data *VdiscoverytaskModel, diags *diag.Diagnostics) bool {
	// Check if service_account_file is provided
	if data.ServiceAccountFile.IsNull() || data.ServiceAccountFile.IsUnknown() {
		return true // No file to process, continue
	}

	// // Skip if token is already provided (don't override existing token)
	// if !data.ServiceAccountFileToken.IsNull() && !data.ServiceAccountFileToken.IsUnknown() {
	// 	return true // Token already exists, continue
	// }

	// Get connection details from client configuration
	baseUrl := r.client.SecurityAPI.Cfg.NIOSHostURL
	username := r.client.SecurityAPI.Cfg.NIOSUsername
	password := r.client.SecurityAPI.Cfg.NIOSPassword

	// Get the file path from the model
	filePath := data.ServiceAccountFile.ValueString()

	// Upload the GCP service account file and get the token
	token, err := utils.UploadPEMFileWithToken(ctx, baseUrl, filePath, username, password)
	if err != nil {
		diags.AddError(
			"Client Error",
			fmt.Sprintf("Unable to process GCP service account file %s, got error: %s", filePath, err),
		)
		return false
	}

	// Store the token in the service_account_file_token field
	data.ServiceAccountFileToken = types.StringValue(token)
	return true
}

// function that will process your CDiscovery file and return the token
func (r *VdiscoverytaskResource) processCDiscoveryFile(ctx context.Context, data *VdiscoverytaskModel, diags *diag.Diagnostics) bool {
	// Check if cdiscovery_file is provided
	if data.CdiscoveryFile.IsNull() || data.CdiscoveryFile.IsUnknown() {
		return true // No file to process, continue
	}

	// Skip if token is already provided (don't override existing token)
	// if !data.CdiscoveryFileToken.IsNull() && !data.CdiscoveryFileToken.IsUnknown() {
	// 	return true // Token already exists, continue
	// }

	// Get connection details from client configuration
	baseUrl := r.client.SecurityAPI.Cfg.NIOSHostURL
	username := r.client.SecurityAPI.Cfg.NIOSUsername
	password := r.client.SecurityAPI.Cfg.NIOSPassword

	// Get the file path from the model
	filePath := data.CdiscoveryFile.ValueString()

	// Upload the CDiscovery file and get the token
	token, err := utils.UploadPEMFileWithToken(ctx, baseUrl, filePath, username, password)
	if err != nil {
		diags.AddError(
			"Client Error",
			fmt.Sprintf("Unable to process CDiscovery file %s, got error: %s", filePath, err),
		)
		return false
	}

	// Store the token in the cdiscovery_file_token field
	data.CdiscoveryFileToken = types.StringValue(token)
	return true
}
