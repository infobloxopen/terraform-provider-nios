package parentalcontrol

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/infoblox-nios-go-client/parentalcontrol"

	"github.com/infobloxopen/terraform-provider-nios/internal/config"
	"github.com/infobloxopen/terraform-provider-nios/internal/retry"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForParentalcontrolSubscriberrecord = "accounting_session_id,alt_ip_addr,ans0,ans1,ans2,ans3,ans4,black_list,bwflag,dynamic_category_policy,flags,ip_addr,ipsd,localid,nas_contextual,op_code,parental_control_policy,prefix,proxy_all,site,subscriber_id,subscriber_secure_policy,unknown_category_policy,white_list,wpc_category_policy"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ParentalcontrolSubscriberrecordResource{}
var _ resource.ResourceWithImportState = &ParentalcontrolSubscriberrecordResource{}

func NewParentalcontrolSubscriberrecordResource() resource.Resource {
	return &ParentalcontrolSubscriberrecordResource{}
}

// ParentalcontrolSubscriberrecordResource defines the resource implementation.
type ParentalcontrolSubscriberrecordResource struct {
	client *niosclient.APIClient
}

func (r *ParentalcontrolSubscriberrecordResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "parentalcontrol_subscriberrecord"
}

func (r *ParentalcontrolSubscriberrecordResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a Parental Control Subscriber Record.",
		Attributes:          ParentalcontrolSubscriberrecordResourceSchemaAttributes,
	}
}

func (r *ParentalcontrolSubscriberrecordResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ParentalcontrolSubscriberrecordResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {

	var data ParentalcontrolSubscriberrecordModel
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Check if white_list or blacklist are provided and bwflag is set to true, if not return error
	if !data.WhiteList.IsNull() || !data.BlackList.IsNull() {
		if data.Bwflag.IsNull() || !data.Bwflag.ValueBool() {
			resp.Diagnostics.AddError(
				"Invalid Configuration",
				"bwflag must be set to true when white_list or black_list is provided.",
			)
			return
		}
	}
}

func (r *ParentalcontrolSubscriberrecordResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ParentalcontrolSubscriberrecordModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	payload := data.Expand(ctx, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	var apiRes *parentalcontrol.CreateParentalcontrolSubscriberrecordResponse

	err := retry.Do(ctx, retry.TransientErrors, func(ctx context.Context) (int, error) {
		var (
			httpRes *http.Response
			callErr error
		)
		apiRes, httpRes, callErr = r.client.ParentalControlAPI.
			ParentalcontrolSubscriberrecordAPI.
			Create(ctx).
			ParentalcontrolSubscriberrecord(*payload).
			ReturnFieldsPlus(readableAttributesForParentalcontrolSubscriberrecord).
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create ParentalcontrolSubscriberrecord, got error: %s", err))
		return
	}

	res := apiRes.CreateParentalcontrolSubscriberrecordResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ParentalcontrolSubscriberrecordResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ParentalcontrolSubscriberrecordModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resourceRef := utils.ExtractResourceRef(data.Ref.ValueString())

	var (
		httpRes *http.Response
		apiRes  *parentalcontrol.GetParentalcontrolSubscriberrecordResponse
	)

	err := retry.Do(ctx, nil, func(ctx context.Context) (int, error) {
		var callErr error
		apiRes, httpRes, callErr = r.client.ParentalControlAPI.
			ParentalcontrolSubscriberrecordAPI.
			Read(ctx, resourceRef).
			ReturnFieldsPlus(readableAttributesForParentalcontrolSubscriberrecord).
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read ParentalcontrolSubscriberrecord, got error: %s", err))
		return
	}

	res := apiRes.GetParentalcontrolSubscriberrecordResponseObjectAsResult.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ParentalcontrolSubscriberrecordResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data ParentalcontrolSubscriberrecordModel

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

	payload := data.Expand(ctx, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	resourceRef := utils.ExtractResourceRef(data.Ref.ValueString())

	var apiRes *parentalcontrol.UpdateParentalcontrolSubscriberrecordResponse

	err := retry.Do(ctx, retry.TransientErrors, func(ctx context.Context) (int, error) {
		var (
			httpRes *http.Response
			callErr error
		)
		apiRes, httpRes, callErr = r.client.ParentalControlAPI.
			ParentalcontrolSubscriberrecordAPI.
			Update(ctx, resourceRef).
			ParentalcontrolSubscriberrecord(*payload).
			ReturnFieldsPlus(readableAttributesForParentalcontrolSubscriberrecord).
			ReturnAsObject(1).
			Execute()

		if httpRes != nil {
			return httpRes.StatusCode, callErr
		}
		return 0, callErr
	})

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update ParentalcontrolSubscriberrecord, got error: %s", err))
		return
	}

	res := apiRes.UpdateParentalcontrolSubscriberrecordResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ParentalcontrolSubscriberrecordResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ParentalcontrolSubscriberrecordModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resourceRef := utils.ExtractResourceRef(data.Ref.ValueString())

	err := retry.Do(ctx, retry.TransientErrors, func(ctx context.Context) (int, error) {
		httpRes, callErr := r.client.ParentalControlAPI.
			ParentalcontrolSubscriberrecordAPI.
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete ParentalcontrolSubscriberrecord, got error: %s", err))
		return
	}
}

func (r *ParentalcontrolSubscriberrecordResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("ref"), req, resp)
}
