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

var readableAttributesForParentalcontrolAvp = "comment,domain_types,is_restricted,name,type,user_defined,value_type,vendor_id,vendor_type"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ParentalcontrolAvpResource{}
var _ resource.ResourceWithImportState = &ParentalcontrolAvpResource{}
var _ resource.ResourceWithValidateConfig = &ParentalcontrolAvpResource{}

func NewParentalcontrolAvpResource() resource.Resource {
	return &ParentalcontrolAvpResource{}
}

// ParentalcontrolAvpResource defines the resource implementation.
type ParentalcontrolAvpResource struct {
	client *niosclient.APIClient
}

func (r *ParentalcontrolAvpResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "parentalcontrol_avp"
}

func (r *ParentalcontrolAvpResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a parental control AVP.",
		Attributes:          ParentalcontrolAvpResourceSchemaAttributes,
	}
}

func (r *ParentalcontrolAvpResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ParentalcontrolAvpResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data ParentalcontrolAvpModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Check if 'is_restricted' is true when domain_types' is specified
	if !data.DomainTypes.IsNull() && !data.DomainTypes.IsUnknown() {
		if !data.IsRestricted.IsNull() && !data.IsRestricted.IsUnknown() && !data.IsRestricted.ValueBool() {
			resp.Diagnostics.AddError(
				"Invalid Configuration",
				"'is_restricted' must be true when 'domain_types' is specified.",
			)
		}
	}

	// Check if 'vendor_id' and 'vendor_type' are specified when 'type'=26
	if !data.Type.IsNull() && !data.Type.IsUnknown() && data.Type.ValueInt64() == 26 {
		if !data.VendorId.IsUnknown() && !data.VendorType.IsUnknown() {
			if data.VendorId.IsNull() || data.VendorType.IsNull() {
				resp.Diagnostics.AddError(
					"Invalid Configuration",
					"'vendor_id' and 'vendor_type' must be specified when 'type' is 26.",
				)
			}
		}
	}
}

func (r *ParentalcontrolAvpResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ParentalcontrolAvpModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	payload := data.Expand(ctx, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	var apiRes *parentalcontrol.CreateParentalcontrolAvpResponse

	err := retry.Do(ctx, retry.TransientErrors, func(ctx context.Context) (int, error) {
		var (
			httpRes *http.Response
			callErr error
		)
		apiRes, httpRes, callErr = r.client.ParentalControlAPI.
			ParentalcontrolAvpAPI.
			Create(ctx).
			ParentalcontrolAvp(*payload).
			ReturnFieldsPlus(readableAttributesForParentalcontrolAvp).
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create ParentalcontrolAvp, got error: %s", err))
		return
	}

	res := apiRes.CreateParentalcontrolAvpResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ParentalcontrolAvpResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ParentalcontrolAvpModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resourceRef := utils.ExtractResourceRef(data.Ref.ValueString())

	var (
		httpRes *http.Response
		apiRes  *parentalcontrol.GetParentalcontrolAvpResponse
	)

	err := retry.Do(ctx, nil, func(ctx context.Context) (int, error) {
		var callErr error
		apiRes, httpRes, callErr = r.client.ParentalControlAPI.
			ParentalcontrolAvpAPI.
			Read(ctx, resourceRef).
			ReturnFieldsPlus(readableAttributesForParentalcontrolAvp).
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read ParentalcontrolAvp, got error: %s", err))
		return
	}

	res := apiRes.GetParentalcontrolAvpResponseObjectAsResult.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ParentalcontrolAvpResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data ParentalcontrolAvpModel

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

	var apiRes *parentalcontrol.UpdateParentalcontrolAvpResponse

	err := retry.Do(ctx, retry.TransientErrors, func(ctx context.Context) (int, error) {
		var (
			httpRes *http.Response
			callErr error
		)
		apiRes, httpRes, callErr = r.client.ParentalControlAPI.
			ParentalcontrolAvpAPI.
			Update(ctx, resourceRef).
			ParentalcontrolAvp(*payload).
			ReturnFieldsPlus(readableAttributesForParentalcontrolAvp).
			ReturnAsObject(1).
			Execute()

		if httpRes != nil {
			return httpRes.StatusCode, callErr
		}
		return 0, callErr
	})

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update ParentalcontrolAvp, got error: %s", err))
		return
	}

	res := apiRes.UpdateParentalcontrolAvpResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ParentalcontrolAvpResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ParentalcontrolAvpModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resourceRef := utils.ExtractResourceRef(data.Ref.ValueString())

	err := retry.Do(ctx, retry.TransientErrors, func(ctx context.Context) (int, error) {
		httpRes, callErr := r.client.ParentalControlAPI.
			ParentalcontrolAvpAPI.
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete ParentalcontrolAvp, got error: %s", err))
		return
	}
}

func (r *ParentalcontrolAvpResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("ref"), req, resp)
}
