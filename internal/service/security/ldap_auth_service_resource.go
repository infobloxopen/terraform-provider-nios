package security

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/config"
	"github.com/infobloxopen/terraform-provider-nios/internal/retry"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForLdapAuthService = "comment,disable,ea_mapping,ldap_group_attribute,ldap_group_authentication_type,ldap_user_attribute,mode,name,recovery_interval,retries,search_scope,servers,timeout"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &LdapAuthServiceResource{}
var _ resource.ResourceWithImportState = &LdapAuthServiceResource{}
var _ resource.ResourceWithModifyPlan = &LdapAuthServiceResource{}

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
		MarkdownDescription: "Manages an LDAP Auth Service.",
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

type serversBindPasswordHashState struct {
	PasswordHash string `json:"servers_bind_password_hash"`
}

func (r *LdapAuthServiceResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return
	}

	var passwordVersion types.Int64
	curRev := int64(0)

	if !req.State.Raw.IsNull() && req.State.Raw.IsKnown() {
		resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("password_version"), &passwordVersion)...)
		if resp.Diagnostics.HasError() {
			return
		}
		if !passwordVersion.IsNull() && !passwordVersion.IsUnknown() {
			curRev = passwordVersion.ValueInt64()
		}
	}

	var data LdapAuthServiceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	servers, diags := extractLdapServers(ctx, data.Servers)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if len(servers) == 0 {
		return
	}

	var prev struct {
		Algo string `json:"algo"`
		Hash string `json:"hash"`
	}
	prevHashes := serversBindPasswordHashState{}

	if b, diags := req.Private.GetKey(ctx, "servers_bind_password_hash"); diags != nil {
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else if b != nil {
		if err := json.Unmarshal(b, &prev); err != nil {
			prev.Hash = ""
		}
		if prev.Hash != "" {
			if err := json.Unmarshal([]byte(prev.Hash), &prevHashes); err != nil {
				prevHashes = serversBindPasswordHashState{}
			}
		}
	}

	newHash, err := hashLdapServers(servers)
	if err != nil {
		resp.Diagnostics.AddError("Hash Error", err.Error())
		return
	}

	plannedHashes := serversBindPasswordHashState{PasswordHash: newHash}
	plannedHashJSON, err := json.Marshal(plannedHashes)
	if err != nil {
		resp.Diagnostics.AddError("Private State Marshal Error", err.Error())
		return
	}

	if plannedHashes.PasswordHash != "" && plannedHashes.PasswordHash != prevHashes.PasswordHash {
		resp.Diagnostics.Append(
			resp.Plan.SetAttribute(ctx, path.Root("password_version"), types.Int64Value(curRev+1))...,
		)

		val := map[string]string{
			"algo": "sha256",
			"hash": string(plannedHashJSON),
		}
		b, err := json.Marshal(val)
		if err != nil {
			resp.Diagnostics.AddError("Private State Marshal Error", err.Error())
			return
		}
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, "servers_bind_password_hash", b)...)
	} else {
		resp.Diagnostics.Append(
			resp.Plan.SetAttribute(ctx, path.Root("password_version"), types.Int64Value(curRev))...,
		)
	}
}

func (r *LdapAuthServiceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LdapAuthServiceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	payload := data.Expand(ctx, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	passwordVersionValue := types.Int64Value(0)
	secretData := serversBindPasswordHashState{}

	servers, diags := extractLdapServers(ctx, data.Servers)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if len(servers) > 0 && len(payload.Servers) == len(servers) {
		for i := range servers {
			if !servers[i].BindPassword.IsNull() && !servers[i].BindPassword.IsUnknown() {
				payload.Servers[i].BindPassword = servers[i].BindPassword.ValueStringPointer()
			}
		}

		serversHash, err := hashLdapServers(servers)
		if err != nil {
			resp.Diagnostics.AddError("Hash Error", err.Error())
			return
		}

		secretData.PasswordHash = serversHash
		secretDataJSON, _ := json.Marshal(secretData)
		val := map[string]string{"algo": "sha256", "hash": string(secretDataJSON)}
		hashedSecret, err := json.Marshal(val)
		if err != nil {
			resp.Diagnostics.AddError("Private State Marshal Error", err.Error())
			return
		}
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, "servers_bind_password_hash", hashedSecret)...)
		passwordVersionValue = types.Int64Value(1)
	}

	var apiRes *security.CreateLdapAuthServiceResponse

	err := retry.Do(ctx, retry.TransientErrors, func(ctx context.Context) (int, error) {
		var (
			httpRes *http.Response
			callErr error
		)
		apiRes, httpRes, callErr = r.client.SecurityAPI.
			LdapAuthServiceAPI.
			Create(ctx).
			LdapAuthService(*payload).
			ReturnFieldsPlus(readableAttributesForLdapAuthService).
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create LdapAuthService, got error: %s", err))
		return
	}

	res := apiRes.CreateLdapAuthServiceResponseAsObject.GetResult()
	data.PasswordVersion = passwordVersionValue

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

	resourceRef := utils.ExtractResourceRef(data.Ref.ValueString())

	var (
		httpRes *http.Response
		apiRes  *security.GetLdapAuthServiceResponse
	)

	err := retry.Do(ctx, nil, func(ctx context.Context) (int, error) {
		var callErr error
		apiRes, httpRes, callErr = r.client.SecurityAPI.
			LdapAuthServiceAPI.
			Read(ctx, resourceRef).
			ReturnFieldsPlus(readableAttributesForLdapAuthService).
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
	var plannedPasswordVersion types.Int64

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("password_version"), &plannedPasswordVersion)...)

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

	servers, diags := extractLdapServers(ctx, data.Servers)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if len(servers) > 0 && len(payload.Servers) == len(servers) {
		for i := range servers {
			if !servers[i].BindPassword.IsNull() && !servers[i].BindPassword.IsUnknown() {
				payload.Servers[i].BindPassword = servers[i].BindPassword.ValueStringPointer()
			}
		}
	}

	resourceRef := utils.ExtractResourceRef(data.Ref.ValueString())

	var apiRes *security.UpdateLdapAuthServiceResponse

	err := retry.Do(ctx, retry.TransientErrors, func(ctx context.Context) (int, error) {
		var (
			httpRes *http.Response
			callErr error
		)
		apiRes, httpRes, callErr = r.client.SecurityAPI.
			LdapAuthServiceAPI.
			Update(ctx, resourceRef).
			LdapAuthService(*payload).
			ReturnFieldsPlus(readableAttributesForLdapAuthService).
			ReturnAsObject(1).
			Execute()

		if httpRes != nil {
			return httpRes.StatusCode, callErr
		}
		return 0, callErr
	})

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update LdapAuthService, got error: %s", err))
		return
	}

	res := apiRes.UpdateLdapAuthServiceResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)
	data.PasswordVersion = plannedPasswordVersion

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

	resourceRef := utils.ExtractResourceRef(data.Ref.ValueString())

	err := retry.Do(ctx, retry.TransientErrors, func(ctx context.Context) (int, error) {
		httpRes, callErr := r.client.SecurityAPI.
			LdapAuthServiceAPI.
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete LdapAuthService, got error: %s", err))
		return
	}
}

func (r *LdapAuthServiceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("ref"), req, resp)
}
