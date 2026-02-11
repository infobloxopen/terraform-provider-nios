package ipam

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

	"github.com/infobloxopen/terraform-provider-nios/internal/config"
)

var readableAttributesForSuperhost = "comment,dhcp_associated_objects,disabled,dns_associated_objects,extattrs,name"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SuperhostResource{}
var _ resource.ResourceWithImportState = &SuperhostResource{}
var _ resource.ResourceWithValidateConfig = &SuperhostResource{}

func NewSuperhostResource() resource.Resource {
	return &SuperhostResource{}
}

// SuperhostResource defines the resource implementation.
type SuperhostResource struct {
	client *niosclient.APIClient
}

func (r *SuperhostResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "ipam_superhost"
}

func (r *SuperhostResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a Super Host",
		Attributes:          SuperhostResourceSchemaAttributes,
	}
}

func (r *SuperhostResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *SuperhostResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var diags diag.Diagnostics
	var data SuperhostModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Validate that no host records are associated with DHCP Associated Objects when record ref is not provided in plan
	if !data.DhcpAssociatedObjects.IsUnknown() && !data.DhcpAssociatedObjects.IsNull() {
		var dhcpAssociatedObjects []string
		resp.Diagnostics.Append(data.DhcpAssociatedObjects.ElementsAs(ctx, &dhcpAssociatedObjects, false)...)
		if resp.Diagnostics.HasError() {
			return
		}

		for _, obj := range dhcpAssociatedObjects {
			if strings.HasPrefix(obj, "record:host") {
				resp.Diagnostics.AddError(
					"Invalid DHCP Associated Object",
					"Host record can only be associated with DNS Associated Objects, not DHCP Associated Objects.",
				)
				return
			}
		}
	}

	// Add internal ID exists in the Extensible Attributes if not already present
	data.ExtAttrs, diags = AddInternalIDToExtAttrs(ctx, data.ExtAttrs, diags)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	apiRes, _, err := r.client.IPAMAPI.
		SuperhostAPI.
		Create(ctx).
		Superhost(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForSuperhost).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Superhost, got error: %s", err))
		return
	}

	res := apiRes.CreateSuperhostResponseAsObject.GetResult()
	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		resp.Diagnostics.AddError("Client Error", "Error while creating Superhost due to inherited Extensible attributes")
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SuperhostResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var diags diag.Diagnostics
	var data SuperhostModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.IPAMAPI.
		SuperhostAPI.
		Read(ctx, data.Uuid.ValueString()).
		ReturnFieldsPlus(readableAttributesForSuperhost).
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Superhost, got error: %s", err))
		return
	}

	res := apiRes.GetSuperhostResponseObjectAsResult.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		resp.Diagnostics.AddError("Client Error", "Error while reading Superhost due to inherited Extensible attributes")
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}


func (r *SuperhostResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data SuperhostModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Validate that no host records are associated with DHCP Associated Objects when record ref is not provided in plan
	if !data.DhcpAssociatedObjects.IsUnknown() && !data.DhcpAssociatedObjects.IsNull() {
		var dhcpAssociatedObjects []string
		resp.Diagnostics.Append(data.DhcpAssociatedObjects.ElementsAs(ctx, &dhcpAssociatedObjects, false)...)
		if resp.Diagnostics.HasError() {
			return
		}

		for _, obj := range dhcpAssociatedObjects {
			if strings.HasPrefix(obj, "record:host") {
				resp.Diagnostics.AddError(
					"Invalid DHCP Associated Object",
					"Host record can only be associated with DNS Associated Objects, not DHCP Associated Objects.",
				)
				return
			}
		}
	}

	planExtAttrs := data.ExtAttrs
	diags = req.State.GetAttribute(ctx, path.Root("uuid"), &data.Uuid)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	diags = req.State.GetAttribute(ctx, path.Root("extattrs_all"), &data.ExtAttrsAll)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	associateInternalId, diags := req.Private.GetKey(ctx, "associate_internal_id")
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	if associateInternalId != nil {
		data.ExtAttrs, diags = AddInternalIDToExtAttrs(ctx, data.ExtAttrs, diags)
		if diags.HasError() {
			resp.Diagnostics.Append(diags...)
			return
		}
	}

	// Add Inherited Extensible Attributes
	data.ExtAttrs, diags = AddInheritedExtAttrs(ctx, data.ExtAttrs, data.ExtAttrsAll)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	apiRes, _, err := r.client.IPAMAPI.
		SuperhostAPI.
		Update(ctx, data.Uuid.ValueString()).
		Superhost(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForSuperhost).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Superhost, got error: %s", err))
		return
	}

	res := apiRes.UpdateSuperhostResponseAsObject.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, planExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		resp.Diagnostics.AddError("Client Error", "Error while updating Superhost due to inherited Extensible attributes")
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if associateInternalId != nil {
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", nil)...)
	}
}

func (r *SuperhostResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SuperhostModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.IPAMAPI.
		SuperhostAPI.
		Delete(ctx, data.Uuid.ValueString()).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Superhost, got error: %s", err))
		return
	}
}

func (r *SuperhostResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data SuperhostModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.DhcpAssociatedObjects.IsUnknown() || data.DhcpAssociatedObjects.IsNull() {
		return
	}

	var dhcpAssociatedObjects []string

	resp.Diagnostics.Append(data.DhcpAssociatedObjects.ElementsAs(ctx, &dhcpAssociatedObjects, true)...)
	if resp.Diagnostics.HasError() || dhcpAssociatedObjects == nil {
		return
	}

	// Validate that no host records are associated with DHCP Associated Objects when record ref is provided in plan
	for _, obj := range dhcpAssociatedObjects {
		if obj == "" {
			continue
		}
		if strings.HasPrefix(obj, "record:host") {
			resp.Diagnostics.AddError(
				"Invalid DHCP Associated Object",
				"Host record can only be associated with DNS Associated Objects, not DHCP Associated Objects.",
			)
		}
	}
}

func (r *SuperhostResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("uuid"), req.ID)...)
	resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", []byte("true"))...)
}
