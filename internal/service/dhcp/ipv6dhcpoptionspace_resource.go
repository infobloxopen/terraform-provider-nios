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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"

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
var _ resource.ResourceWithModifyPlan = &Ipv6dhcpoptionspaceResource{}

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

func (r *Ipv6dhcpoptionspaceResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// If this is a read/destroy operation, no need to modify plan
	if req.Plan.Raw.IsNull() {
		return
	}

	// If this is a create operation, no need to modify plan
	if req.State.Raw.IsNull() {
		return
	}

	var plan Ipv6dhcpoptionspaceModel
	var state Ipv6dhcpoptionspaceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if plan.OptionDefinitions.IsUnknown() {
		// Fetch the latest state from the api
		apiRes, _, err := r.client.DHCPAPI.
			Ipv6dhcpoptionspaceAPI.
			Read(ctx, utils.ExtractResourceRef(state.Ref.ValueString())).
			ReturnFieldsPlus(readableAttributesForIpv6dhcpoptionspace).
			ReturnAsObject(1).
			Execute()

		if err == nil {
			res := apiRes.GetIpv6dhcpoptionspaceResponseObjectAsResult.GetResult()

			var apiData Ipv6dhcpoptionspaceModel
			apiData.Flatten(ctx, &res, &resp.Diagnostics)

			if !resp.Diagnostics.HasError() {
				// Use the api's option_definitions value
				plan.OptionDefinitions = apiData.OptionDefinitions
				resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
			}
		} else {
			if !state.OptionDefinitions.IsNull() {
				plan.OptionDefinitions = state.OptionDefinitions
				resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
			}
		}
	}
}

func (r *Ipv6dhcpoptionspaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Ipv6dhcpoptionspaceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.DHCPAPI.
		Ipv6dhcpoptionspaceAPI.
		Create(ctx).
		Ipv6dhcpoptionspace(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForIpv6dhcpoptionspace).
		ReturnAsObject(1).
		Execute()
	if err != nil {
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

	apiRes, httpRes, err := r.client.DHCPAPI.
		Ipv6dhcpoptionspaceAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForIpv6dhcpoptionspace).
		ReturnAsObject(1).
		Execute()

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

	apiRes, _, err := r.client.DHCPAPI.
		Ipv6dhcpoptionspaceAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Ipv6dhcpoptionspace(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForIpv6dhcpoptionspace).
		ReturnAsObject(1).
		Execute()
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

	err := retry.RetryContext(ctx, Ipv6OptionSpaceOperationTimeout, func() *retry.RetryError {
		httpRes, err := r.client.DHCPAPI.
			Ipv6dhcpoptionspaceAPI.
			Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
			Execute()
		if err != nil {
			if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
				return nil
			}
			if strings.Contains(err.Error(), "cannot be deleted as it has an option referenced") {
				return retry.RetryableError(err)
			}
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Ipv6dhcpoptionspace, got error: %s", err))
			return retry.NonRetryableError(err)
		}
		return nil
	})
	if err != nil {
		return
	}
}

func (r *Ipv6dhcpoptionspaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("ref"), req, resp)
}
