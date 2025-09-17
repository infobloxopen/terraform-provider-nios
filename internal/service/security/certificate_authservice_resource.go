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

var readableAttributesForCertificateAuthservice = "auto_populate_login,ca_certificates,comment,disabled,enable_password_request,enable_remote_lookup,max_retries,name,ocsp_check,ocsp_responders,recovery_interval,remote_lookup_service,remote_lookup_username,response_timeout,trust_model,user_match_type"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &CertificateAuthserviceResource{}
var _ resource.ResourceWithImportState = &CertificateAuthserviceResource{}

func NewCertificateAuthserviceResource() resource.Resource {
	return &CertificateAuthserviceResource{}
}

// CertificateAuthserviceResource defines the resource implementation.
type CertificateAuthserviceResource struct {
	client *niosclient.APIClient
}

func (r *CertificateAuthserviceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "security_certificate_authservice"
}

func (r *CertificateAuthserviceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages an Certificate Authservice.",
		Attributes:          CertificateAuthserviceResourceSchemaAttributes,
	}
}

func (r *CertificateAuthserviceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *CertificateAuthserviceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	//var diags diag.Diagnostics
	var data CertificateAuthserviceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	token, err := utils.UploadPEMFile(ctx, r.client, "/Users/chaithra/go/src/github.com/infobloxopen/terraform-provider-nios/internal/utils/cert.pem")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to upload PEM file, got error: %s", err))
		return
	}
	fmt.Printf("Uploaded PEM file, received token: %s\n", token)

	apiRes, _, err := r.client.SecurityAPI.
		CertificateAuthserviceAPI.
		Create(ctx).
		CertificateAuthservice(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForCertificateAuthservice).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create CertificateAuthservice, got error: %s", err))
		return
	}

	res := apiRes.CreateCertificateAuthserviceResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CertificateAuthserviceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	//var diags diag.Diagnostics
	var data CertificateAuthserviceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.SecurityAPI.
		CertificateAuthserviceAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForCertificateAuthservice).
		ReturnAsObject(1).
		Execute()

	//remove from the state if not found 
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			// Resource no longer exists, remove from state
			resp.State.RemoveResource(ctx)
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read CertificateAuthservice, got error: %s", err))
		return
	}

	res := apiRes.GetCertificateAuthserviceResponseObjectAsResult.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}


func (r *CertificateAuthserviceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data CertificateAuthserviceModel

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
		CertificateAuthserviceAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		CertificateAuthservice(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForCertificateAuthservice).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update CertificateAuthservice, got error: %s", err))
		return
	}

	res := apiRes.UpdateCertificateAuthserviceResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CertificateAuthserviceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CertificateAuthserviceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.SecurityAPI.
		CertificateAuthserviceAPI.
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete CertificateAuthservice, got error: %s", err))
		return
	}
}

func (r *CertificateAuthserviceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("ref"), req, resp)
}