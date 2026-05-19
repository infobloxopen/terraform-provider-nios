package misc

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/infoblox-nios-go-client/misc"
	"github.com/infobloxopen/terraform-provider-nios/internal/config"
	"github.com/infobloxopen/terraform-provider-nios/internal/retry"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForBfdtemplate = "authentication_key_id,authentication_type,detection_multiplier,min_rx_interval,min_tx_interval,name"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &BfdtemplateResource{}
var _ resource.ResourceWithImportState = &BfdtemplateResource{}
var _ resource.ResourceWithModifyPlan = &BfdtemplateResource{}

func NewBfdtemplateResource() resource.Resource {
	return &BfdtemplateResource{}
}

// BfdtemplateResource defines the resource implementation.
type BfdtemplateResource struct {
	client *niosclient.APIClient
}

type secretsHashState struct {
	AuthenticationKey string `json:"authentication_key_hash"`
}

func (r *BfdtemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "misc_bfdtemplate"
}

func (r *BfdtemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a BFD template.",
		Attributes:          BfdtemplateResourceSchemaAttributes,
	}
}

func (r *BfdtemplateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *BfdtemplateResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return
	}

	var stateRev types.Int64
	resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("secret_revision"), &stateRev)...)
	if resp.Diagnostics.HasError() {
		return
	}

	curRev := int64(0)
	if !stateRev.IsNull() && !stateRev.IsUnknown() {
		curRev = stateRev.ValueInt64()
	}

	var planSecret types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("authentication_key"), &planSecret)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var prev struct {
		Algo string `json:"algo"`
		Hash string `json:"hash"`
	}
	if b, diags := req.Private.GetKey(ctx, "authentication_key_hash"); diags != nil {
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else if b != nil {
		if err := json.Unmarshal(b, &prev); err != nil {
			prev.Hash = ""
		}
	}

	prevHashes := secretsHashState{}
	if prev.Hash != "" {
		_ = json.Unmarshal([]byte(prev.Hash), &prevHashes)
	}
	plannedHashes := prevHashes
	computeNewHash := !planSecret.IsNull() && !planSecret.IsUnknown()
	plannedHash := prev.Hash

	if computeNewHash {
		h := sha256.New()
		h.Write([]byte(planSecret.ValueString()))
		plannedHashes.AuthenticationKey = hex.EncodeToString(h.Sum(nil))
		if data, err := json.Marshal(plannedHashes); err == nil {
			plannedHash = string(data)
		}
	}

	if computeNewHash && plannedHash != prev.Hash {
		resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("secret_revision"), types.Int64Value(curRev+1))...)
		val := map[string]string{"algo": "sha256", "hash": plannedHash}
		b, err := json.Marshal(val)
		if err != nil {
			resp.Diagnostics.AddError("Private State Marshal Error", err.Error())
			return
		}
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, "authentication_key_hash", b)...)
		return
	}

	resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("secret_revision"), types.Int64Value(curRev))...)
}

func (r *BfdtemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data BfdtemplateModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	payload := data.Expand(ctx, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	secret_revisionValue := types.Int64Value(0)
	var authenticationKey types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("authentication_key"), &authenticationKey)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !authenticationKey.IsNull() && !authenticationKey.IsUnknown() {
		secretVal := authenticationKey.ValueString()
		payload.AuthenticationKey = &secretVal
		secret_revisionValue = types.Int64Value(1)

		secretData := secretsHashState{}
		h := sha256.New()
		h.Write([]byte(authenticationKey.ValueString()))
		secretData.AuthenticationKey = hex.EncodeToString(h.Sum(nil))
		secretDataJSON, _ := json.Marshal(secretData)
		val := map[string]string{"algo": "sha256", "hash": string(secretDataJSON)}
		hashedSecret, err := json.Marshal(val)
		if err != nil {
			resp.Diagnostics.AddError("Private State Marshal Error", err.Error())
			return
		}
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, "authentication_key_hash", hashedSecret)...)
	}

	var apiRes *misc.CreateBfdtemplateResponse

	err := retry.Do(ctx, retry.TransientErrors, func(ctx context.Context) (int, error) {
		var (
			httpRes *http.Response
			callErr error
		)
		apiRes, httpRes, callErr = r.client.MiscAPI.
			BfdtemplateAPI.
			Create(ctx).
			Bfdtemplate(*payload).
			ReturnFieldsPlus(readableAttributesForBfdtemplate).
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Bfdtemplate, got error: %s", err))
		return
	}

	res := apiRes.CreateBfdtemplateResponseAsObject.GetResult()

	data.SecretRevision = secret_revisionValue
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BfdtemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data BfdtemplateModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resourceRef := utils.ExtractResourceRef(data.Ref.ValueString())

	var (
		httpRes *http.Response
		apiRes  *misc.GetBfdtemplateResponse
	)

	err := retry.Do(ctx, nil, func(ctx context.Context) (int, error) {
		var callErr error
		apiRes, httpRes, callErr = r.client.MiscAPI.
			BfdtemplateAPI.
			Read(ctx, resourceRef).
			ReturnFieldsPlus(readableAttributesForBfdtemplate).
			ReturnAsObject(1).
			ProxySearch(config.GetProxySearch()).
			Execute()

		if httpRes != nil {
			return httpRes.StatusCode, callErr
		}
		return 0, callErr
	})

	// Handle case not found
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			// Resource no longer exists, remove from state
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Bfdtemplate, got error: %s", err))
		return
	}

	res := apiRes.GetBfdtemplateResponseObjectAsResult.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BfdtemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data BfdtemplateModel

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

	var authenticationKey types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("authentication_key"), &authenticationKey)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if !authenticationKey.IsNull() && !authenticationKey.IsUnknown() {
		secretVal := authenticationKey.ValueString()
		payload.AuthenticationKey = &secretVal
	}

	resourceRef := utils.ExtractResourceRef(data.Ref.ValueString())

	var apiRes *misc.UpdateBfdtemplateResponse

	err := retry.Do(ctx, retry.TransientErrors, func(ctx context.Context) (int, error) {
		var (
			httpRes *http.Response
			callErr error
		)
		apiRes, httpRes, callErr = r.client.MiscAPI.
			BfdtemplateAPI.
			Update(ctx, resourceRef).
			Bfdtemplate(*payload).
			ReturnFieldsPlus(readableAttributesForBfdtemplate).
			ReturnAsObject(1).
			Execute()

		if httpRes != nil {
			return httpRes.StatusCode, callErr
		}
		return 0, callErr
	})

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Bfdtemplate, got error: %s", err))
		return
	}

	res := apiRes.UpdateBfdtemplateResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BfdtemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data BfdtemplateModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resourceRef := utils.ExtractResourceRef(data.Ref.ValueString())

	err := retry.Do(ctx, retry.TransientErrors, func(ctx context.Context) (int, error) {
		httpRes, callErr := r.client.MiscAPI.
			BfdtemplateAPI.
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Bfdtemplate, got error: %s", err))
		return
	}
}

func (r *BfdtemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("ref"), req, resp)
}
