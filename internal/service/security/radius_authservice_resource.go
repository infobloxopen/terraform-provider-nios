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

type radiusAuthserviceSecretsHashState struct {
	ServersHash string `json:"servers_secret_hash"`
}

func (r *RadiusAuthserviceResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return
	}

	var stateSecretVersion types.Int64
	curRev := int64(0)

	if !req.State.Raw.IsNull() && req.State.Raw.IsKnown() {
		resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("secret_version"), &stateSecretVersion)...)
		if resp.Diagnostics.HasError() {
			return
		}
		if !stateSecretVersion.IsNull() && !stateSecretVersion.IsUnknown() {
			curRev = stateSecretVersion.ValueInt64()
		}
	}

	var data RadiusAuthserviceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	servers, diags := extractRadiusServers(ctx, data.Servers)
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

	prevHashes := radiusAuthserviceSecretsHashState{}

	if b, diags := req.Private.GetKey(ctx, "servers_secret_hash"); diags != nil {
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
				prevHashes = radiusAuthserviceSecretsHashState{}
			}
		}
	}

	newHash, err := hashRadiusServers(servers)
	if err != nil {
		resp.Diagnostics.AddError("Hash Error", err.Error())
		return
	}

	plannedHashes := radiusAuthserviceSecretsHashState{ServersHash: newHash}
	plannedHashJSON, err := json.Marshal(plannedHashes)
	if err != nil {
		resp.Diagnostics.AddError("Private State Marshal Error", err.Error())
		return
	}

	if plannedHashes.ServersHash != "" && plannedHashes.ServersHash != prevHashes.ServersHash {
		resp.Diagnostics.Append(
			resp.Plan.SetAttribute(ctx, path.Root("secret_version"), types.Int64Value(curRev+1))...,
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
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, "servers_secret_hash", b)...)
	} else {
		resp.Diagnostics.Append(
			resp.Plan.SetAttribute(ctx, path.Root("secret_version"), types.Int64Value(curRev))...,
		)
	}
}

func (r *RadiusAuthserviceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data RadiusAuthserviceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	payload := data.Expand(ctx, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	secretVersion := types.Int64Value(0)
	secretData := radiusAuthserviceSecretsHashState{}

	servers, diags := extractRadiusServers(ctx, data.Servers)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if len(servers) > 0 && len(payload.Servers) == len(servers) {
		for i := range servers {
			if !servers[i].SharedSecret.IsNull() && !servers[i].SharedSecret.IsUnknown() {
				payload.Servers[i].SharedSecret = servers[i].SharedSecret.ValueStringPointer()
			}
		}

		serversHash, err := hashRadiusServers(servers)
		if err != nil {
			resp.Diagnostics.AddError("Hash Error", err.Error())
			return
		}

		secretData.ServersHash = serversHash
		secretDataJSON, _ := json.Marshal(secretData)
		val := map[string]string{"algo": "sha256", "hash": string(secretDataJSON)}
		hashedSecret, err := json.Marshal(val)
		if err != nil {
			resp.Diagnostics.AddError("Private State Marshal Error", err.Error())
			return
		}
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, "servers_secret_hash", hashedSecret)...)
		secretVersion = types.Int64Value(1)
	}

	apiRes, _, err := r.client.SecurityAPI.
		RadiusAuthserviceAPI.
		Create(ctx).
		RadiusAuthservice(*payload).
		ReturnFieldsPlus(readableAttributesForRadiusAuthservice).
		ReturnAsObject(1).
		Execute()

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create RadiusAuthservice, got error: %s", err))
		return
	}

	res := apiRes.CreateRadiusAuthserviceResponseAsObject.GetResult()

	data.SecretVersion = secretVersion
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

	apiRes, httpRes, err := r.client.SecurityAPI.
		RadiusAuthserviceAPI.
		Read(ctx, utils.ResolveIdentifier(data.Uuid, data.Ref)).
		ReturnFieldsPlus(readableAttributesForRadiusAuthservice).
		ReturnAsObject(1).
		ProxySearch(config.GetProxySearch()).
		Execute()

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
	var plannedSecretVersion types.Int64

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("secret_version"), &plannedSecretVersion)...)
	if resp.Diagnostics.HasError() {
		return
	}

	payload := data.Expand(ctx, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = req.State.GetAttribute(ctx, path.Root("ref"), &data.Ref)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	diags = req.State.GetAttribute(ctx, path.Root("uuid"), &data.Uuid)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	servers, diags := extractRadiusServers(ctx, data.Servers)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if len(servers) > 0 && len(payload.Servers) == len(servers) {
		for i := range servers {
			if !servers[i].SharedSecret.IsNull() && !servers[i].SharedSecret.IsUnknown() {
				payload.Servers[i].SharedSecret = servers[i].SharedSecret.ValueStringPointer()
			}
		}
	}

	resourceRef := utils.ExtractResourceRef(data.Ref.ValueString())

	var apiRes *security.UpdateRadiusAuthserviceResponse

	apiRes, _, err := r.client.SecurityAPI.
		RadiusAuthserviceAPI.
		Update(ctx, resourceRef).
		RadiusAuthservice(*payload).
		ReturnFieldsPlus(readableAttributesForRadiusAuthservice).
		ReturnAsObject(1).
		Execute()

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update RadiusAuthservice, got error: %s", err))
		return
	}

	res := apiRes.UpdateRadiusAuthserviceResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)
	data.SecretVersion = plannedSecretVersion

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

	httpRes, err := r.client.SecurityAPI.
		RadiusAuthserviceAPI.
		Delete(ctx, utils.ResolveIdentifier(data.Uuid, data.Ref)).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete RadiusAuthservice, got error: %s", err))
		return
	}
}

func (r *RadiusAuthserviceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("uuid"), req, resp)
}
