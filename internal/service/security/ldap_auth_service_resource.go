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

	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForLdapAuthService = "comment,disable,ea_mapping,ldap_group_attribute,ldap_group_authentication_type,ldap_user_attribute,mode,name,recovery_interval,retries,search_scope,servers,timeout"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &LdapAuthServiceResource{}
var _ resource.ResourceWithImportState = &LdapAuthServiceResource{}

func NewLdapAuthServiceResource() resource.Resource {
	return &LdapAuthServiceResource{}
}

// LdapAuthServiceResource defines the resource implementation.
type LdapAuthServiceResource struct {
	client *niosclient.APIClient
}

func (r *LdapAuthServiceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "security_ldap_auth_service"
}

func (r *LdapAuthServiceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "",
		Attributes:          LdapAuthServiceResourceSchemaAttributes,
	}
}

func (r *LdapAuthServiceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *LdapAuthServiceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LdapAuthServiceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.SecurityAPI.
		LdapAuthServiceAPI.
		Create(ctx).
		LdapAuthService(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForLdapAuthService).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create LdapAuthService, got error: %s", err))
		return
	}

	res := apiRes.CreateLdapAuthServiceResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LdapAuthServiceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LdapAuthServiceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.SecurityAPI.
		LdapAuthServiceAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForLdapAuthService).
		ReturnAsObject(1).
		Execute()

		// Handle not found case
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			// Resource no longer exists, remove from state
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read LdapAuthService, got error: %s", err))
		return
	}

	res := apiRes.GetLdapAuthServiceResponseObjectAsResult.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LdapAuthServiceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data LdapAuthServiceModel

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

	apiRes, _, err := r.client.SecurityAPI.
		LdapAuthServiceAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		LdapAuthService(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForLdapAuthService).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update LdapAuthService, got error: %s", err))
		return
	}

	res := apiRes.UpdateLdapAuthServiceResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LdapAuthServiceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LdapAuthServiceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.SecurityAPI.
		LdapAuthServiceAPI.
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete LdapAuthService, got error: %s", err))
		return
	}
}

func (r *LdapAuthServiceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("ref"), req, resp)
}
