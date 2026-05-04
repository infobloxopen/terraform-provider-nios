package dhcp

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-nios/internal/config"
	"github.com/infobloxopen/terraform-provider-nios/internal/retry"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForIpv6rangetemplate = "cloud_api_compatible,comment,delegated_member,exclude,logic_filter_rules,member,name,number_of_addresses,offset,option_filter_rules,recycle_leases,server_association_type,use_logic_filter_rules,use_recycle_leases"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &Ipv6rangetemplateResource{}
var _ resource.ResourceWithImportState = &Ipv6rangetemplateResource{}

func NewIpv6rangetemplateResource() resource.Resource {
	return &Ipv6rangetemplateResource{}
}

// Ipv6rangetemplateResource defines the resource implementation.
type Ipv6rangetemplateResource struct {
	client *niosclient.APIClient
}

func (r *Ipv6rangetemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "dhcp_ipv6_range_template"
}

func (r *Ipv6rangetemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages an IPV6 Range Template.",
		Attributes:          Ipv6rangetemplateResourceSchemaAttributes,
	}
}

func (r *Ipv6rangetemplateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *Ipv6rangetemplateResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var config Ipv6rangetemplateModel

	// Parse the configuration
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	// For Configuration object, any attributes not defined by the user appear as null, unless derived from another instance.
	// We perform IsUnknown() check to handle variables from .tfvars that are resolved
	// during the plan phase rather than validation phase, preventing false validation errors.

	var serverAssociationType string

	if !config.ServerAssociationType.IsUnknown() {
		serverAssociationType = "NONE"
		if !config.ServerAssociationType.IsNull() {
			serverAssociationType = config.ServerAssociationType.ValueString()
		}
	}

	// If server_association_type is MEMBER, member field must be set
	if serverAssociationType == "MEMBER" {
		if config.Member.IsNull() || config.Member.IsUnknown() {
			resp.Diagnostics.AddAttributeError(
				path.Root("member"),
				"Invalid Configuration",
				"The 'member' field must be set when 'server_association_type' is set to 'MEMBER'.",
			)
		}
	}

	// If server_association_type is NONE, member field cannot be set
	if serverAssociationType == "NONE" {
		if !config.Member.IsNull() && !config.Member.IsUnknown() {
			resp.Diagnostics.AddAttributeError(
				path.Root("member"),
				"Invalid Configuration",
				"The 'member' field cannot be set when 'server_association_type' is set to 'NONE' (default).",
			)
		}
	}
}

func (r *Ipv6rangetemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Ipv6rangetemplateModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	payload := data.Expand(ctx, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	var apiRes *dhcp.CreateIpv6rangetemplateResponse

	err := retry.Do(ctx, retry.TransientErrors, func(ctx context.Context) (int, error) {
		var (
			httpRes *http.Response
			callErr error
		)
		apiRes, httpRes, callErr = r.client.DHCPAPI.
			Ipv6rangetemplateAPI.
			Create(ctx).
			Ipv6rangetemplate(*payload).
			ReturnFieldsPlus(readableAttributesForIpv6rangetemplate).
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Ipv6rangetemplate, got error: %s", err))
		return
	}

	res := apiRes.CreateIpv6rangetemplateResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Ipv6rangetemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Ipv6rangetemplateModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resourceRef := utils.ExtractResourceRef(data.Ref.ValueString())

	var (
		httpRes *http.Response
		apiRes  *dhcp.GetIpv6rangetemplateResponse
	)

	err := retry.Do(ctx, nil, func(ctx context.Context) (int, error) {
		var callErr error
		apiRes, httpRes, callErr = r.client.DHCPAPI.
			Ipv6rangetemplateAPI.
			Read(ctx, resourceRef).
			ReturnFieldsPlus(readableAttributesForIpv6rangetemplate).
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Ipv6rangetemplate, got error: %s", err))
		return
	}

	res := apiRes.GetIpv6rangetemplateResponseObjectAsResult.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Ipv6rangetemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data Ipv6rangetemplateModel

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

	var apiRes *dhcp.UpdateIpv6rangetemplateResponse

	err := retry.Do(ctx, retry.TransientErrors, func(ctx context.Context) (int, error) {
		var (
			httpRes *http.Response
			callErr error
		)
		apiRes, httpRes, callErr = r.client.DHCPAPI.
			Ipv6rangetemplateAPI.
			Update(ctx, resourceRef).
			Ipv6rangetemplate(*payload).
			ReturnFieldsPlus(readableAttributesForIpv6rangetemplate).
			ReturnAsObject(1).
			Execute()

		if httpRes != nil {
			return httpRes.StatusCode, callErr
		}
		return 0, callErr
	})

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Ipv6rangetemplate, got error: %s", err))
		return
	}

	res := apiRes.UpdateIpv6rangetemplateResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Ipv6rangetemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Ipv6rangetemplateModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resourceRef := utils.ExtractResourceRef(data.Ref.ValueString())

	err := retry.Do(ctx, retry.TransientErrors, func(ctx context.Context) (int, error) {
		httpRes, callErr := r.client.DHCPAPI.
			Ipv6rangetemplateAPI.
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Ipv6rangetemplate, got error: %s", err))
		return
	}
}

func (r *Ipv6rangetemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("ref"), req, resp)
}
