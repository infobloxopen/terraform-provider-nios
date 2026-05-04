package dhcp

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-nios/internal/retry"

	"github.com/infobloxopen/terraform-provider-nios/internal/config"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

const (
	// OptionSpaceOperationTimeout is the maximum amount of time to wait for eventual consistency
	OptionSpaceOperationTimeout = 2 * time.Minute
)

var readableAttributesForDhcpoptionspace = "comment,name,option_definitions,space_type"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &DhcpoptionspaceResource{}
var _ resource.ResourceWithImportState = &DhcpoptionspaceResource{}

func NewDhcpoptionspaceResource() resource.Resource {
	return &DhcpoptionspaceResource{}
}

// DhcpoptionspaceResource defines the resource implementation.
type DhcpoptionspaceResource struct {
	client *niosclient.APIClient
}

func (r *DhcpoptionspaceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "dhcp_optionspace"
}

func (r *DhcpoptionspaceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a DHCP Option Space.",
		Attributes:          DhcpoptionspaceResourceSchemaAttributes,
	}
}

func (r *DhcpoptionspaceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *DhcpoptionspaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DhcpoptionspaceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var apiRes *dhcp.CreateDhcpoptionspaceResponse

	err := retry.Do(ctx, retry.TransientErrors, func(ctx context.Context) (int, error) {
		var (
			httpRes *http.Response
			callErr error
		)

		apiRes, httpRes, callErr = r.client.DHCPAPI.
			DhcpoptionspaceAPI.
			Create(ctx).
			Dhcpoptionspace(*data.Expand(ctx, &resp.Diagnostics)).
			ReturnFieldsPlus(readableAttributesForDhcpoptionspace).
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Dhcpoptionspace, got error: %s", err))
		return
	}

	res := apiRes.CreateDhcpoptionspaceResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DhcpoptionspaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DhcpoptionspaceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resourceRef := utils.ExtractResourceRef(data.Ref.ValueString())

	var (
		httpRes *http.Response
		apiRes  *dhcp.GetDhcpoptionspaceResponse
	)

	err := retry.Do(ctx, nil, func(ctx context.Context) (int, error) {
		var callErr error
		apiRes, httpRes, callErr = r.client.DHCPAPI.
			DhcpoptionspaceAPI.
			Read(ctx, resourceRef).
			ReturnFieldsPlus(readableAttributesForDhcpoptionspace).
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Dhcpoptionspace, got error: %s", err))
		return
	}

	res := apiRes.GetDhcpoptionspaceResponseObjectAsResult.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DhcpoptionspaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data DhcpoptionspaceModel

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

	var apiRes *dhcp.UpdateDhcpoptionspaceResponse

	err := retry.Do(ctx, retry.TransientErrors, func(ctx context.Context) (int, error) {
		var (
			httpRes *http.Response
			callErr error
		)

		apiRes, httpRes, callErr = r.client.DHCPAPI.
			DhcpoptionspaceAPI.
			Update(ctx, resourceRef).
			Dhcpoptionspace(*data.Expand(ctx, &resp.Diagnostics)).
			ReturnFieldsPlus(readableAttributesForDhcpoptionspace).
			ReturnAsObject(1).
			Execute()

		if httpRes != nil {
			return httpRes.StatusCode, callErr
		}
		return 0, callErr
	})

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Dhcpoptionspace, got error: %s", err))
		return
	}

	res := apiRes.UpdateDhcpoptionspaceResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DhcpoptionspaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DhcpoptionspaceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resourceRef := utils.ExtractResourceRef(data.Ref.ValueString())

	err := retry.DoWithTimeout(ctx, OptionSpaceOperationTimeout, IsOptionDefinitionReferenced, func(ctx context.Context) (int, error) {
		httpRes, callErr := r.client.DHCPAPI.
			DhcpoptionspaceAPI.
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Dhcpoptionspace, got error: %s", err))
		return
	}
}

func (r *DhcpoptionspaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("ref"), req, resp)
}

// IsOptionDefinitionReferenced checks if the error indicates that the option definition is referenced.
func IsOptionDefinitionReferenced(err error) bool {
	if err == nil {
		return false
	}
	errStr := strings.ToLower(err.Error())
	return strings.Contains(errStr, "cannot be deleted as it has an option referenced") || retry.TransientErrors(err)
}
