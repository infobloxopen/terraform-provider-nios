package dhcp

import (
	"context"
	"fmt"
	"net/http"
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
	// Ipv6OptionSpaceOperationTimeout is the maximum amount of time to wait for eventual consistency
	Ipv6OptionSpaceOperationTimeout = 2 * time.Minute
)

var readableAttributesForIpv6dhcpoptionspace = "comment,enterprise_number,name,option_definitions"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &Ipv6dhcpoptionspaceResource{}
var _ resource.ResourceWithImportState = &Ipv6dhcpoptionspaceResource{}

func NewIpv6dhcpoptionspaceResource() resource.Resource {
	return &Ipv6dhcpoptionspaceResource{}
}

// Ipv6dhcpoptionspaceResource defines the resource implementation.
type Ipv6dhcpoptionspaceResource struct {
	client *niosclient.APIClient
}

func (r *Ipv6dhcpoptionspaceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "dhcp_ipv6optionspace"
}

func (r *Ipv6dhcpoptionspaceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages an IPv6 DHCP Option Space.",
		Attributes:          Ipv6dhcpoptionspaceResourceSchemaAttributes,
	}
}

func (r *Ipv6dhcpoptionspaceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *Ipv6dhcpoptionspaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Ipv6dhcpoptionspaceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	payload := data.Expand(ctx, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	var apiRes *dhcp.CreateIpv6dhcpoptionspaceResponse

	err := retry.Do(ctx, retry.TransientErrors, func(ctx context.Context) (int, error) {
		var (
			httpRes *http.Response
			callErr error
		)
		apiRes, httpRes, callErr = r.client.DHCPAPI.
			Ipv6dhcpoptionspaceAPI.
			Create(ctx).
			Ipv6dhcpoptionspace(*payload).
			ReturnFieldsPlus(readableAttributesForIpv6dhcpoptionspace).
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Ipv6dhcpoptionspace, got error: %s", err))
		return
	}

	res := apiRes.CreateIpv6dhcpoptionspaceResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Ipv6dhcpoptionspaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Ipv6dhcpoptionspaceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resourceRef := utils.ExtractResourceRef(data.Ref.ValueString())

	var (
		httpRes *http.Response
		apiRes  *dhcp.GetIpv6dhcpoptionspaceResponse
	)

	err := retry.Do(ctx, nil, func(ctx context.Context) (int, error) {
		var callErr error
		apiRes, httpRes, callErr = r.client.DHCPAPI.
			Ipv6dhcpoptionspaceAPI.
			Read(ctx, resourceRef).
			ReturnFieldsPlus(readableAttributesForIpv6dhcpoptionspace).
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Ipv6dhcpoptionspace, got error: %s", err))
		return
	}

	res := apiRes.GetIpv6dhcpoptionspaceResponseObjectAsResult.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Ipv6dhcpoptionspaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data Ipv6dhcpoptionspaceModel

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

	var apiRes *dhcp.UpdateIpv6dhcpoptionspaceResponse

	err := retry.Do(ctx, retry.TransientErrors, func(ctx context.Context) (int, error) {
		var (
			httpRes *http.Response
			callErr error
		)
		apiRes, httpRes, callErr = r.client.DHCPAPI.
			Ipv6dhcpoptionspaceAPI.
			Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
			Ipv6dhcpoptionspace(*payload).
			ReturnFieldsPlus(readableAttributesForIpv6dhcpoptionspace).
			ReturnAsObject(1).
			Execute()

		if httpRes != nil {
			return httpRes.StatusCode, callErr
		}
		return 0, callErr

	})

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Ipv6dhcpoptionspace, got error: %s", err))
		return
	}

	res := apiRes.UpdateIpv6dhcpoptionspaceResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Ipv6dhcpoptionspaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Ipv6dhcpoptionspaceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resourceRef := utils.ExtractResourceRef(data.Ref.ValueString())

	err := retry.DoWithTimeout(ctx, Ipv6OptionSpaceOperationTimeout, IsOptionDefinitionReferenced, func(ctx context.Context) (int, error) {
		httpRes, callErr := r.client.DHCPAPI.
			Ipv6dhcpoptionspaceAPI.
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Ipv6dhcpoptionspace, got error: %s", err))
		return
	}
}

func (r *Ipv6dhcpoptionspaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("ref"), req, resp)
}
