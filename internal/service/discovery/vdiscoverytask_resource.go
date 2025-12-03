package discovery

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForVdiscoverytask = "accounts_list,allow_unsecured_connection,auto_consolidate_cloud_ea,auto_consolidate_managed_tenant,auto_consolidate_managed_vm,auto_create_dns_hostname_template,auto_create_dns_record,auto_create_dns_record_type,cdiscovery_file_token,comment,credentials_type,dns_view_private_ip,dns_view_public_ip,domain_name,driver_type,enable_filter,enabled,fqdn_or_ip,govcloud_enabled,identity_version,last_run,member,merge_data,multiple_accounts_sync_policy,name,network_filter,network_list,port,private_network_view,private_network_view_mapping_policy,protocol,public_network_view,public_network_view_mapping_policy,role_arn,scheduled_run,selected_regions,service_account_file,service_account_file_token,state,state_msg,sync_child_accounts,update_dns_view_private_ip,update_dns_view_public_ip,update_metadata,use_identity,username"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &VdiscoverytaskResource{}
var _ resource.ResourceWithImportState = &VdiscoverytaskResource{}
var _ resource.ResourceWithValidateConfig = &VdiscoverytaskResource{}

func NewVdiscoverytaskResource() resource.Resource {
	return &VdiscoverytaskResource{}
}

// VdiscoverytaskResource defines the resource implementation.
type VdiscoverytaskResource struct {
	client *niosclient.APIClient
}

func (r *VdiscoverytaskResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "discovery_vdiscovery_task"
}

func (r *VdiscoverytaskResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a vDiscovery Task.",
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

func (r *VdiscoverytaskResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data VdiscoverytaskModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	driverType := data.DriverType.ValueString()

	// Validate cdiscovery_file requirement for UPLOAD policy
	if !data.MultipleAccountsSyncPolicy.IsNull() && !data.MultipleAccountsSyncPolicy.IsUnknown() {
		if data.MultipleAccountsSyncPolicy.ValueString() == "UPLOAD" {
			if data.CdiscoveryFile.IsNull() || data.CdiscoveryFile.IsUnknown() || data.CdiscoveryFile.ValueString() == "" {
				resp.Diagnostics.AddError(
					"Missing CDDiscovery File",
					"'cdiscovery_file' is required when 'multiple_accounts_sync_policy' is set to 'UPLOAD'.",
				)
			}
		}
	}

	// Validate dns_view_private_ip requires update_dns_view_private_ip = true
	if !data.DnsViewPrivateIp.IsNull() && !data.DnsViewPrivateIp.IsUnknown() && data.DnsViewPrivateIp.ValueString() != "" {
		if data.UpdateDnsViewPrivateIp.IsNull() || data.UpdateDnsViewPrivateIp.IsUnknown() || !data.UpdateDnsViewPrivateIp.ValueBool() {
			resp.Diagnostics.AddError(
				"Invalid DNS View Configuration",
				"'update_dns_view_private_ip' must be set to true to use 'dns_view_private_ip'.",
			)
		}
	}

	// Validate dns_view_public_ip requires update_dns_view_public_ip = true
	if !data.DnsViewPublicIp.IsNull() && !data.DnsViewPublicIp.IsUnknown() && data.DnsViewPublicIp.ValueString() != "" {
		if data.UpdateDnsViewPublicIp.IsNull() || data.UpdateDnsViewPublicIp.IsUnknown() || !data.UpdateDnsViewPublicIp.ValueBool() {
			resp.Diagnostics.AddError(
				"Invalid DNS View Configuration",
				"'update_dns_view_public_ip' must be set to true to use 'dns_view_public_ip'.",
			)
		}
	}

	// Validate domain_name requirement for OPENSTACK with KEYSTONE_V3
	if driverType == "OPENSTACK" {
		if !data.IdentityVersion.IsNull() && !data.IdentityVersion.IsUnknown() {
			if data.IdentityVersion.ValueString() == "KEYSTONE_V3" {
				if data.DomainName.IsNull() || data.DomainName.IsUnknown() || data.DomainName.ValueString() == "" {
					resp.Diagnostics.AddError(
						"Missing Domain Name",
						"'domain_name' is required when 'identity_version' is set to 'KEYSTONE_V3'.",
					)
				}
			}
		}

		// Validate identity_version requirement for OPENSTACK
		if data.IdentityVersion.IsNull() || data.IdentityVersion.IsUnknown() {
			resp.Diagnostics.AddError(
				"Missing Identity Version",
				"'identity_version' is required when 'driver_type' is 'OPENSTACK'.",
			)
		}

		// Validate use_identity requirement for OPENSTACK
		if data.UseIdentity.IsNull() || data.UseIdentity.IsUnknown() {
			resp.Diagnostics.AddError(
				"Missing Use Identity",
				"'use_identity' is required when 'driver_type' is 'OPENSTACK'.",
			)
		}
	}

	// Validate credentials_type DIRECT requirements
	if !data.CredentialsType.IsNull() && !data.CredentialsType.IsUnknown() {
		if data.CredentialsType.ValueString() == "DIRECT" {
			// Password required for DIRECT credentials
			if data.Password.IsNull() || data.Password.IsUnknown() || data.Password.ValueString() == "" {
				resp.Diagnostics.AddError(
					"Missing Password",
					"'password' is required when 'credentials_type' is set to 'DIRECT'.",
				)
			}

			// Username required for DIRECT credentials
			if data.Username.IsNull() || data.Username.IsUnknown() || data.Username.ValueString() == "" {
				resp.Diagnostics.AddError(
					"Missing Username",
					"'username' is required when 'credentials_type' is set to 'DIRECT'.",
				)
			}
		}
	}

	// Validate selected_regions requirement for AWS
	if driverType == "AWS" {
		if data.SelectedRegions.IsNull() || data.SelectedRegions.IsUnknown() || data.SelectedRegions.ValueString() == "" {
			resp.Diagnostics.AddError(
				"Missing Selected Regions",
				"'selected_regions' is required when 'driver_type' is 'AWS'.",
			)
		}
	}

	// Validate service_account_file configuration
	serviceAccountFileProvided := !data.ServiceAccountFile.IsNull() && !data.ServiceAccountFile.IsUnknown() && data.ServiceAccountFile.ValueString() != ""

	if driverType == "GCP" {
		if !serviceAccountFileProvided {
			resp.Diagnostics.AddError(
				"Missing Service Account File",
				"'service_account_file' is required when 'driver_type' is 'GCP'.",
			)
		}
	} else if serviceAccountFileProvided {
		resp.Diagnostics.AddError(
			"Invalid Service Account File Configuration",
			fmt.Sprintf("'service_account_file' is only supported for GCP driver type, but got '%s'.", driverType),
		)
	}

	// Validate cdiscovery_file is only for AWS and GCP
	if !data.CdiscoveryFile.IsNull() && !data.CdiscoveryFile.IsUnknown() && data.CdiscoveryFile.ValueString() != "" {
		if driverType != "AWS" && driverType != "GCP" {
			resp.Diagnostics.AddError(
				"Invalid Cdiscovery File Configuration",
				fmt.Sprintf("'cdiscovery_file' is only supported for AWS and GCP driver types, but got '%s'.", driverType),
			)
		}
	}

	// Validate schedule configuration
	if !data.ScheduledRun.IsNull() && !data.ScheduledRun.IsUnknown() {
		var scheduledRun VdiscoverytaskScheduledRunModel
		diags := data.ScheduledRun.As(ctx, &scheduledRun, basetypes.ObjectAsOptions{})
		if diags.HasError() {
			resp.Diagnostics.Append(diags...)
			return
		}

		// Validate recurring_time conflicts
		if !scheduledRun.RecurringTime.IsNull() && !scheduledRun.RecurringTime.IsUnknown() {
			if !scheduledRun.HourOfDay.IsNull() || !scheduledRun.Year.IsNull() || !scheduledRun.Month.IsNull() || !scheduledRun.DayOfMonth.IsNull() {
				resp.Diagnostics.AddAttributeError(
					path.Root("scheduled_run").AtName("recurring_time"),
					"Invalid Configuration for Schedule",
					"Cannot set recurring_time if any of hour_of_day, year, month, day_of_month is set",
				)
			}
		}

		// Validate repeat field logic
		if !scheduledRun.Repeat.IsNull() && !scheduledRun.Repeat.IsUnknown() {
			repeatValue := scheduledRun.Repeat.ValueString()

			switch repeatValue {
			case "ONCE":
				// For ONCE: cannot set weekdays, frequency, every
				if !scheduledRun.Weekdays.IsNull() || !scheduledRun.Frequency.IsNull() || !scheduledRun.Every.IsNull() {
					resp.Diagnostics.AddAttributeError(
						path.Root("scheduled_run").AtName("repeat"),
						"Invalid Configuration for Repeat",
						"Cannot set frequency, weekdays and every if repeat is set to ONCE",
					)
				}
				// For ONCE: must set month, day_of_month, hour_of_day, minutes_past_hour
				if scheduledRun.Month.IsNull() || scheduledRun.DayOfMonth.IsNull() || scheduledRun.HourOfDay.IsNull() || scheduledRun.MinutesPastHour.IsNull() {
					resp.Diagnostics.AddAttributeError(
						path.Root("scheduled_run").AtName("repeat"),
						"Invalid Configuration for Schedule",
						"If repeat is set to ONCE, then month, day_of_month, hour_of_day and minutes_past_hour must be set",
					)
				}
			case "RECUR":
				// For RECUR: cannot set month, day_of_month, year
				if !scheduledRun.Month.IsNull() || !scheduledRun.DayOfMonth.IsNull() || !scheduledRun.Year.IsNull() {
					resp.Diagnostics.AddAttributeError(
						path.Root("scheduled_run").AtName("repeat"),
						"Invalid Configuration for Repeat",
						"Cannot set month, day_of_month and year if repeat is set to RECUR",
					)
				}

				// For RECUR: must set frequency, hour_of_day, minutes_past_hour
				if scheduledRun.Frequency.IsNull() || scheduledRun.HourOfDay.IsNull() || scheduledRun.MinutesPastHour.IsNull() {
					resp.Diagnostics.AddAttributeError(
						path.Root("scheduled_run").AtName("repeat"),
						"Invalid Configuration for Schedule",
						"If repeat is set to RECUR, then frequency, hour_of_day and minutes_past_hour must be set",
					)
				}

				// Handle weekdays validation based on frequency for RECUR only
				if !scheduledRun.Frequency.IsNull() && !scheduledRun.Frequency.IsUnknown() {
					frequencyValue := scheduledRun.Frequency.ValueString()

					if frequencyValue == "WEEKLY" {
						// WEEKLY requires weekdays
						if scheduledRun.Weekdays.IsNull() || scheduledRun.Weekdays.IsUnknown() {
							resp.Diagnostics.AddAttributeError(
								path.Root("scheduled_run").AtName("weekdays"),
								"Invalid Configuration for Weekdays",
								"Weekdays must be set if Frequency is set to WEEKLY",
							)
						}
					} else {
						// Non-WEEKLY cannot have weekdays
						if !scheduledRun.Weekdays.IsNull() && !scheduledRun.Weekdays.IsUnknown() {
							resp.Diagnostics.AddAttributeError(
								path.Root("scheduled_run").AtName("weekdays"),
								"Invalid Configuration for Weekdays",
								"Weekdays can only be set if Frequency is set to WEEKLY",
							)
						}
					}
				}
			}
		}
	}

	if !data.UseIdentity.IsNull() && !data.UseIdentity.IsUnknown() && data.UseIdentity.ValueBool() {
		// When use_identity is true, enforce standard ports
		if !data.Protocol.IsNull() && !data.Protocol.IsUnknown() && !data.Port.IsNull() && !data.Port.IsUnknown() {
			protocol := data.Protocol.ValueString()
			port := data.Port.ValueInt64()

			if protocol == "HTTPS" && port != 443 {
				resp.Diagnostics.AddAttributeError(
					path.Root("port"),
					"Invalid Port Configuration",
					fmt.Sprintf("When use_identity is true and protocol is HTTPS, port must be 443. Got: %d", port),
				)
			}

			if protocol == "HTTP" && port != 80 {
				resp.Diagnostics.AddAttributeError(
					path.Root("port"),
					"Invalid Port Configuration",
					fmt.Sprintf("When use_identity is true and protocol is HTTP, port must be 80. Got: %d", port),
				)
			}
		}
	}

	// Validate allow_unsecured_connection requirements
	if !data.AllowUnsecuredConnection.IsNull() && !data.AllowUnsecuredConnection.IsUnknown() && data.AllowUnsecuredConnection.ValueBool() {
		// When allow_unsecured_connection is true, protocol must be HTTPS
		if !data.Protocol.IsNull() && !data.Protocol.IsUnknown() {
			if data.Protocol.ValueString() != "HTTPS" {
				resp.Diagnostics.AddAttributeError(
					path.Root("protocol"),
					"Invalid Protocol Configuration",
					fmt.Sprintf("When allow_unsecured_connection is true, protocol must be HTTPS. Got: %s", data.Protocol.ValueString()),
				)
			}
		}

		// When allow_unsecured_connection is true, driver_type must be VMware or OpenStack
		if !data.DriverType.IsNull() && !data.DriverType.IsUnknown() {
			driverType := data.DriverType.ValueString()
			if driverType != "VMWARE" && driverType != "OPENSTACK" {
				resp.Diagnostics.AddAttributeError(
					path.Root("driver_type"),
					"Invalid Driver Type Configuration",
					fmt.Sprintf("When allow_unsecured_connection is true, driver_type must be either VMware or OpenStack. Got: %s", driverType),
				)
			}
		}
	}
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

	// Process Cdiscovery file if multiple_accounts_sync_policy is UPLOAD
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

	// Process Cdiscovery file if multiple_accounts_sync_policy is UPLOAD
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

	// Get connection details from client configuration
	baseUrl := r.client.SecurityAPI.Cfg.NIOSHostURL
	username := r.client.SecurityAPI.Cfg.NIOSUsername
	password := r.client.SecurityAPI.Cfg.NIOSPassword

	// Get the file path from the model
	filePath := data.ServiceAccountFile.ValueString()

	// Upload the GCP service account file and get the token
	token, err := utils.UploadFileWithToken(ctx, baseUrl, filePath, username, password)
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

	// Get connection details from client configuration
	baseUrl := r.client.SecurityAPI.Cfg.NIOSHostURL
	username := r.client.SecurityAPI.Cfg.NIOSUsername
	password := r.client.SecurityAPI.Cfg.NIOSPassword

	// Get the file path from the model
	filePath := data.CdiscoveryFile.ValueString()

	// Upload the CDiscovery file and get the token
	token, err := utils.UploadFileWithToken(ctx, baseUrl, filePath, username, password)
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
