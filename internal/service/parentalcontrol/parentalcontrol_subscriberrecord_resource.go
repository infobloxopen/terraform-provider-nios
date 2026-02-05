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
		MarkdownDescription: "",
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

func (r *ParentalcontrolSubscriberrecordResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ParentalcontrolSubscriberrecordModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.ParentalControlAPI.
		ParentalcontrolSubscriberrecordAPI.
		Create(ctx).
		ParentalcontrolSubscriberrecord(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForParentalcontrolSubscriberrecord).
		ReturnAsObject(1).
		Execute()
	if err != nil {
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

	apiRes, httpRes, err := r.client.ParentalControlAPI.
		ParentalcontrolSubscriberrecordAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForParentalcontrolSubscriberrecord).
		ReturnAsObject(1).
		Execute()

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

	apiRes, _, err := r.client.ParentalControlAPI.
		ParentalcontrolSubscriberrecordAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ParentalcontrolSubscriberrecord(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForParentalcontrolSubscriberrecord).
		ReturnAsObject(1).
		Execute()
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

	httpRes, err := r.client.ParentalControlAPI.
		ParentalcontrolSubscriberrecordAPI.
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete ParentalcontrolSubscriberrecord, got error: %s", err))
		return
	}
}

func (r *ParentalcontrolSubscriberrecordResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("ref"), req, resp)
}
