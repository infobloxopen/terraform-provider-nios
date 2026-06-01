package security

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/config"
	"github.com/infobloxopen/terraform-provider-nios/internal/retry"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForRadiusAuthservice = "acct_retries,acct_timeout,auth_retries,auth_timeout,cache_ttl,comment,disable,enable_cache,mode,name,recovery_interval,servers"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &RadiusAuthserviceResource{}
var _ resource.ResourceWithImportState = &RadiusAuthserviceResource{}

func NewRadiusAuthserviceResource() resource.Resource {
	return &RadiusAuthserviceResource{}
}

// RadiusAuthserviceResource defines the resource implementation.
type RadiusAuthserviceResource struct {
	client *niosclient.APIClient
}

func (r *RadiusAuthserviceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "security_radius_authservice"
}

func (r *RadiusAuthserviceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a Radius Authentication Service.",
		Attributes:          RadiusAuthserviceResourceSchemaAttributes,
	}
}

func (r *RadiusAuthserviceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *RadiusAuthserviceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data RadiusAuthserviceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	payload := data.Expand(ctx, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	var apiRes *security.CreateRadiusAuthserviceResponse

	err := retry.Do(ctx, retry.TransientErrors, func(ctx context.Context) (int, error) {
		var (
			httpRes *http.Response
			callErr error
		)
		apiRes, httpRes, callErr = r.client.SecurityAPI.
			RadiusAuthserviceAPI.
			Create(ctx).
			RadiusAuthservice(*payload).
			ReturnFieldsPlus(readableAttributesForRadiusAuthservice).
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create RadiusAuthservice, got error: %s", err))
		return
	}

	res := apiRes.CreateRadiusAuthserviceResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RadiusAuthserviceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data RadiusAuthserviceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resourceRef := utils.ExtractResourceRef(data.Ref.ValueString())

	var (
		httpRes *http.Response
		apiRes  *security.GetRadiusAuthserviceResponse
	)

	err := retry.Do(ctx, nil, func(ctx context.Context) (int, error) {
		var callErr error
		apiRes, httpRes, callErr = r.client.SecurityAPI.
			RadiusAuthserviceAPI.
			Read(ctx, resourceRef).
			ReturnFieldsPlus(readableAttributesForRadiusAuthservice).
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read RadiusAuthservice, got error: %s", err))
		return
	}

	res := apiRes.GetRadiusAuthserviceResponseObjectAsResult.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RadiusAuthserviceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data RadiusAuthserviceModel

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

	var apiRes *security.UpdateRadiusAuthserviceResponse

	err := retry.Do(ctx, retry.TransientErrors, func(ctx context.Context) (int, error) {
		var (
			httpRes *http.Response
			callErr error
		)
		apiRes, httpRes, callErr = r.client.SecurityAPI.
			RadiusAuthserviceAPI.
			Update(ctx, resourceRef).
			RadiusAuthservice(*payload).
			ReturnFieldsPlus(readableAttributesForRadiusAuthservice).
			ReturnAsObject(1).
			Execute()

		if httpRes != nil {
			return httpRes.StatusCode, callErr
		}
		return 0, callErr
	})

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update RadiusAuthservice, got error: %s", err))
		return
	}

	res := apiRes.UpdateRadiusAuthserviceResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RadiusAuthserviceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data RadiusAuthserviceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resourceRef := utils.ExtractResourceRef(data.Ref.ValueString())

	err := retry.Do(ctx, retry.TransientErrors, func(ctx context.Context) (int, error) {
		httpRes, callErr := r.client.SecurityAPI.
			RadiusAuthserviceAPI.
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete RadiusAuthservice, got error: %s", err))
		return
	}
}

func (r *RadiusAuthserviceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("ref"), req, resp)
}
