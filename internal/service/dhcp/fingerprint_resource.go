package dhcp

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"

	"github.com/infobloxopen/terraform-provider-nios/internal/config"
)

var readableAttributesForFingerprint = "comment,device_class,disable,extattrs,ipv6_option_sequence,name,option_sequence,type,vendor_id"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &FingerprintResource{}
var _ resource.ResourceWithImportState = &FingerprintResource{}
var _ resource.ResourceWithValidateConfig = &FingerprintResource{}

func NewFingerprintResource() resource.Resource {
	return &FingerprintResource{}
}

// FingerprintResource defines the resource implementation.
type FingerprintResource struct {
	client *niosclient.APIClient
}

func (r *FingerprintResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "dhcp_fingerprint"
}

func (r *FingerprintResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a DHCP Fingerprint.",
		Attributes:          FingerprintResourceSchemaAttributes,
	}
}

func (r *FingerprintResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *FingerprintResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var diags diag.Diagnostics
	var data FingerprintModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Add internal ID exists in the Extensible Attributes if not already present
	data.ExtAttrs, diags = AddInternalIDToExtAttrs(ctx, data.ExtAttrs, diags)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	apiRes, _, err := r.client.DHCPAPI.
		FingerprintAPI.
		Create(ctx).
		Fingerprint(*data.Expand(ctx, &resp.Diagnostics, true)).
		ReturnFieldsPlus(readableAttributesForFingerprint).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Fingerprint, got error: %s", err))
		return
	}

	res := apiRes.CreateFingerprintResponseAsObject.GetResult()
	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		resp.Diagnostics.AddError("Client Error", "Error while creating Fingerprint due to inherited Extensible attributes")
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FingerprintResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var diags diag.Diagnostics
	var data FingerprintModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.DHCPAPI.
		FingerprintAPI.
		Read(ctx, data.Uuid.ValueString()).
		ReturnFieldsPlus(readableAttributesForFingerprint).
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Fingerprint, got error: %s", err))
		return
	}

	res := apiRes.GetFingerprintResponseObjectAsResult.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		resp.Diagnostics.AddError("Client Error", "Error while reading Fingerprint due to inherited Extensible attributes")
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}


func (r *FingerprintResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data FingerprintModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
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

	apiRes, _, err := r.client.DHCPAPI.
		FingerprintAPI.
		Update(ctx, data.Uuid.ValueString()).
		Fingerprint(*data.Expand(ctx, &resp.Diagnostics, false)).
		ReturnFieldsPlus(readableAttributesForFingerprint).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Fingerprint, got error: %s", err))
		return
	}

	res := apiRes.UpdateFingerprintResponseAsObject.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, planExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		resp.Diagnostics.AddError("Client Error", "Error while updating Fingerprint due to inherited Extensible attributes")
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if associateInternalId != nil {
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", nil)...)
	}
}

func (r *FingerprintResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data FingerprintModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.DHCPAPI.
		FingerprintAPI.
		Delete(ctx, data.Uuid.ValueString()).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Fingerprint, got error: %s", err))
		return
	}
}

func (r *FingerprintResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data FingerprintModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Validate option_sequence
	if !data.OptionSequence.IsNull() && !data.OptionSequence.IsUnknown() {
		for _, option := range data.OptionSequence.Elements() {
			optionStr := option.(types.String).ValueString()
			validateOptionSequence(optionStr, path.Root("option_sequence"), "option_sequence", &resp.Diagnostics)
		}
	}

	// Validate ipv6_option_sequence
	if !data.Ipv6OptionSequence.IsNull() && !data.Ipv6OptionSequence.IsUnknown() {
		for _, option := range data.Ipv6OptionSequence.Elements() {
			optionStr := option.(types.String).ValueString()
			validateOptionSequence(optionStr, path.Root("ipv6_option_sequence"), "ipv6_option_sequence", &resp.Diagnostics)
		}
	}
}

// validateOptionSequence validates that the option sequence contains comma separated numbers in the range of 0 to 255.
func validateOptionSequence(optionStr string, attrPath path.Path, attrName string, diags *diag.Diagnostics) {
	optionNumbers := strings.Split(optionStr, ",")
	for _, numStr := range optionNumbers {
		if numStr == "" {
			diags.AddAttributeError(
				attrPath,
				fmt.Sprintf("Invalid %s", attrName),
				fmt.Sprintf("Option sequence cannot be empty. %s is not valid.", optionStr),
			)
			continue
		}

		trimmedNumStr := strings.TrimSpace(numStr)
		if numStr != trimmedNumStr {
			diags.AddAttributeError(
				attrPath,
				fmt.Sprintf("Invalid %s", attrName),
				fmt.Sprintf("Leading or Trailing whitespace is not allowed in the option %s in %s '%s'", trimmedNumStr, attrName, optionStr),
			)
			continue
		}

		numInt, err := strconv.Atoi(numStr)
		if err != nil || numInt < 0 || numInt > 255 {
			diags.AddAttributeError(
				attrPath,
				fmt.Sprintf("Invalid %s", attrName),
				fmt.Sprintf("An option sequence must contain comma separated numbers in the range of 0 to 255. Invalid value: %s", numStr),
			)
		}
	}
}

func (r *FingerprintResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("uuid"), req.ID)...)
	resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", []byte("true"))...)
}
