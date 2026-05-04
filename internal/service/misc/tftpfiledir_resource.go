package misc

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/infoblox-nios-go-client/misc"

	"github.com/infobloxopen/terraform-provider-nios/internal/config"
	"github.com/infobloxopen/terraform-provider-nios/internal/retry"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForTftpfiledir = "directory,is_synced_to_gm,last_modify,name,type,vtftp_dir_members"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &TftpfiledirResource{}
var _ resource.ResourceWithImportState = &TftpfiledirResource{}
var _ resource.ResourceWithValidateConfig = &TftpfiledirResource{}

func NewTftpfiledirResource() resource.Resource {
	return &TftpfiledirResource{}
}

// TftpfiledirResource defines the resource implementation.
type TftpfiledirResource struct {
	client *niosclient.APIClient
}

func (r *TftpfiledirResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "misc_tftpfiledir"
}

func (r *TftpfiledirResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a TFTP file/directory object.",
		Attributes:          TftpfiledirResourceSchemaAttributes,
	}
}

func (r *TftpfiledirResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *TftpfiledirResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data TftpfiledirModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var apiRes *misc.CreateTftpfiledirResponse

	err := retry.Do(ctx, retry.TransientErrors, func(ctx context.Context) (int, error) {
		var (
			httpRes *http.Response
			callErr error
		)

		apiRes, httpRes, callErr = r.client.MiscAPI.
			TftpfiledirAPI.
			Create(ctx).
			Tftpfiledir(*data.Expand(ctx, &resp.Diagnostics, true)).
			ReturnFieldsPlus(readableAttributesForTftpfiledir).
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
		if strings.Contains(err.Error(), "The operation failed: Failed system call") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Tftpfiledir. This error may occur if the TFTP directory '%s' specified in the 'directory' attribute does not exist on the Infoblox appliance. Please ensure that the directory exists and has the appropriate permissions, then try again.", data.Directory.ValueString()))
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Tftpfiledir, got error: %s", err))
		return
	}

	res := apiRes.CreateTftpfiledirResponseAsObject.GetResult()

	if !data.VtftpDirMembers.IsUnknown() && !data.VtftpDirMembers.IsNull() {
		apiRes2, _, err2 := r.client.MiscAPI.
			TftpfiledirAPI.
			Update(ctx, utils.ExtractResourceRef(*res.Ref)).
			Tftpfiledir(*data.Expand(ctx, &resp.Diagnostics, false)).
			ReturnFieldsPlus(readableAttributesForTftpfiledir).
			ReturnAsObject(1).
			Execute()
		if err2 != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Tftpfiledir, got error: %s", err2))
			return
		}
		res2 := apiRes2.UpdateTftpfiledirResponseAsObject.GetResult()
		data.Flatten(ctx, &res2, &resp.Diagnostics)
	} else {
		data.Flatten(ctx, &res, &resp.Diagnostics)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TftpfiledirResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data TftpfiledirModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resourceRef := utils.ExtractResourceRef(data.Ref.ValueString())

	var (
		httpRes *http.Response
		apiRes  *misc.GetTftpfiledirResponse
	)

	err := retry.Do(ctx, nil, func(ctx context.Context) (int, error) {
		var callErr error

		apiRes, httpRes, callErr = r.client.MiscAPI.
			TftpfiledirAPI.
			Read(ctx, resourceRef).
			ReturnFieldsPlus(readableAttributesForTftpfiledir).
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Tftpfiledir, got error: %s", err))
		return
	}

	res := apiRes.GetTftpfiledirResponseObjectAsResult.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TftpfiledirResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data TftpfiledirModel

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

	resourceRef := utils.ExtractResourceRef(data.Ref.ValueString())

	var apiRes *misc.UpdateTftpfiledirResponse

	err := retry.Do(ctx, retry.TransientErrors, func(ctx context.Context) (int, error) {
		var (
			httpRes *http.Response
			callErr error
		)
		apiRes, httpRes, callErr = r.client.MiscAPI.
			TftpfiledirAPI.
			Update(ctx, resourceRef).
			Tftpfiledir(*data.Expand(ctx, &resp.Diagnostics, false)).
			ReturnFieldsPlus(readableAttributesForTftpfiledir).
			ReturnAsObject(1).
			Execute()

		if httpRes != nil {
			return httpRes.StatusCode, callErr
		}
		return 0, callErr
	})

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Tftpfiledir, got error: %s", err))
		return
	}

	res := apiRes.UpdateTftpfiledirResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TftpfiledirResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data TftpfiledirModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resourceRef := utils.ExtractResourceRef(data.Ref.ValueString())

	err := retry.Do(ctx, retry.TransientErrors, func(ctx context.Context) (int, error) {
		httpRes, callErr := r.client.MiscAPI.
			TftpfiledirAPI.
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Tftpfiledir, got error: %s", err))
		return
	}
}

func (r *TftpfiledirResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data TftpfiledirModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.Type.ValueString() == "FILE" {
		resp.Diagnostics.AddError(
			"Invalid Configuration",
			"File Type is currently not supported for TFTP file system entities.",
		)
	}
}

func (r *TftpfiledirResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("ref"), req, resp)
}
