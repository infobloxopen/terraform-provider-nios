package dhcp

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"

	"github.com/infobloxopen/terraform-provider-nios/internal/config"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForIpv6dhcpoptiondefinition = "code,name,space,type"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &Ipv6dhcpoptiondefinitionResource{}
var _ resource.ResourceWithImportState = &Ipv6dhcpoptiondefinitionResource{}

func NewIpv6dhcpoptiondefinitionResource() resource.Resource {
	return &Ipv6dhcpoptiondefinitionResource{}
}

// Ipv6dhcpoptiondefinitionResource defines the resource implementation.
type Ipv6dhcpoptiondefinitionResource struct {
	client *niosclient.APIClient
}

func (r *Ipv6dhcpoptiondefinitionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "dhcp_ipv6optiondefinition"
}

func (r *Ipv6dhcpoptiondefinitionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages an IPv6 DHCP Option Definition.",
		Attributes:          Ipv6dhcpoptiondefinitionResourceSchemaAttributes,
	}
}

func (r *Ipv6dhcpoptiondefinitionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *Ipv6dhcpoptiondefinitionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Ipv6dhcpoptiondefinitionModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.DHCPAPI.
		Ipv6dhcpoptiondefinitionAPI.
		Create(ctx).
		Ipv6dhcpoptiondefinition(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForIpv6dhcpoptiondefinition).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Ipv6dhcpoptiondefinition, got error: %s", err))
		return
	}

	res := apiRes.CreateIpv6dhcpoptiondefinitionResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Ipv6dhcpoptiondefinitionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Ipv6dhcpoptiondefinitionModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.DHCPAPI.
		Ipv6dhcpoptiondefinitionAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForIpv6dhcpoptiondefinition).
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Ipv6dhcpoptiondefinition, got error: %s", err))
		return
	}

	res := apiRes.GetIpv6dhcpoptiondefinitionResponseObjectAsResult.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Ipv6dhcpoptiondefinitionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data Ipv6dhcpoptiondefinitionModel
	var stateData Ipv6dhcpoptiondefinitionModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

	if resp.Diagnostics.HasError() {
		return
	}

	diags = req.State.GetAttribute(ctx, path.Root("ref"), &data.Ref)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	// Update ref if space has changed
	if !data.Space.Equal(stateData.Space) {
		r.updateRefIfSpaceChanged(ctx, resp, &data, &stateData)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	apiRes, _, err := r.client.DHCPAPI.
		Ipv6dhcpoptiondefinitionAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Ipv6dhcpoptiondefinition(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForIpv6dhcpoptiondefinition).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Ipv6dhcpoptiondefinition, got error: %s", err))
		return
	}

	res := apiRes.UpdateIpv6dhcpoptiondefinitionResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Ipv6dhcpoptiondefinitionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Ipv6dhcpoptiondefinitionModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.DHCPAPI.
		Ipv6dhcpoptiondefinitionAPI.
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Ipv6dhcpoptiondefinition, got error: %s", err))
		return
	}
}

func (r *Ipv6dhcpoptiondefinitionResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {

	var data Ipv6dhcpoptiondefinitionModel
	// Read Terraform prior state data into the model.
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.Space.IsUnknown() {
		return
	}

	space := "DHCPv6"

	if !data.Space.IsNull() {
		space = data.Space.ValueString()
	}

	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		name := data.Name.ValueString()
		// If space defaults to "DHCPv6", then name should start with "dhcp6."
		if space == "DHCPv6" {
			if !strings.HasPrefix(name, "dhcp6.") {
				resp.Diagnostics.AddError(
					"Invalid Name for DHCPv6 Option Definition",
					"The name of a DHCP IPv6 option definition object in the default space (DHCPv6) must start with 'dhcp6.'.",
				)
			}
		} else {
			// If space is custom, then name should not start with "dhcp6."
			if strings.HasPrefix(name, "dhcp6.") {
				resp.Diagnostics.AddError(
					"Invalid Name for Custom DHCPv6 Option Definition",
					"The name of a DHCP IPv6 option definition object in a custom space must not start with 'dhcp6.'.",
				)
			}
		}
	}
}

func (r *Ipv6dhcpoptiondefinitionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("ref"), req, resp)
}

// updateRefIfSpaceChanged updates the ref if the option space name changes by
// finding the option definition with the new space name and updating the data model accordingly.
func (r *Ipv6dhcpoptiondefinitionResource) updateRefIfSpaceChanged(ctx context.Context, resp *resource.UpdateResponse, data *Ipv6dhcpoptiondefinitionModel, stateData *Ipv6dhcpoptiondefinitionModel) {
	if resp.Diagnostics.HasError() {
		return
	}

	// Search for the option definition with the new space
	listApiRes, _, err := r.client.DHCPAPI.
		Ipv6dhcpoptiondefinitionAPI.
		List(ctx).
		Filters(map[string]interface{}{
			"name":  stateData.Name.ValueString(),
			"space": data.Space.ValueString(),
			"code":  stateData.Code.ValueInt64(),
			"type":  stateData.Type.ValueString(),
		}).
		ReturnFieldsPlus(readableAttributesForIpv6dhcpoptiondefinition).
		ReturnAsObject(1).
		Execute()

	if err != nil {
		resp.State.RemoveResource(ctx)
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Ipv6dhcpoptiondefinition list, got error: %s", err))
		return
	}

	results := listApiRes.ListIpv6dhcpoptiondefinitionResponseObject.GetResult()

	if len(results) == 0 {
		return
	}

	data.Ref = types.StringValue(*results[0].Ref)

}
