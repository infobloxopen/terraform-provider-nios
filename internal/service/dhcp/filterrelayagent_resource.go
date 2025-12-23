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

	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForFilterrelayagent = "circuit_id_name,circuit_id_substring_length,circuit_id_substring_offset,comment,extattrs,is_circuit_id,is_circuit_id_substring,is_remote_id,is_remote_id_substring,name,remote_id_name,remote_id_substring_length,remote_id_substring_offset"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &FilterrelayagentResource{}
var _ resource.ResourceWithImportState = &FilterrelayagentResource{}

var _ resource.ResourceWithValidateConfig = &FilterrelayagentResource{}

func NewFilterrelayagentResource() resource.Resource {
	return &FilterrelayagentResource{}
}

// FilterrelayagentResource defines the resource implementation.
type FilterrelayagentResource struct {
	client *niosclient.APIClient
}

func (r *FilterrelayagentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "dhcp_filterrelayagent"
}

func (r *FilterrelayagentResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a DHCP Filter Relay Agent resource.",
		Attributes:          FilterrelayagentResourceSchemaAttributes,
	}
}

func (r *FilterrelayagentResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *FilterrelayagentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var diags diag.Diagnostics
	var data FilterrelayagentModel

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
		FilterrelayagentAPI.
		Create(ctx).
		Filterrelayagent(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForFilterrelayagent).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Filterrelayagent, got error: %s", err))
		return
	}

	res := apiRes.CreateFilterrelayagentResponseAsObject.GetResult()
	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		resp.Diagnostics.AddError("Client Error", "Error while creating Filterrelayagent due to inherited Extensible attributes")
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FilterrelayagentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var diags diag.Diagnostics
	var data FilterrelayagentModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	associateInternalId, diags := req.Private.GetKey(ctx, "associate_internal_id")
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.DHCPAPI.
		FilterrelayagentAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForFilterrelayagent).
		ReturnAsObject(1).
		Execute()

	// If the resource is not found, try searching using Extensible Attributes
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound && r.ReadByExtAttrs(ctx, &data, resp) {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Filterrelayagent, got error: %s", err))
		return
	}

	res := apiRes.GetFilterrelayagentResponseObjectAsResult.GetResult()

	apiTerraformId, ok := (*res.ExtAttrs)[terraformInternalIDEA]
	if !ok {
		apiTerraformId.Value = ""
	}

	if associateInternalId == nil {
		stateExtAttrs := ExpandExtAttrs(ctx, data.ExtAttrsAll, &diags)
		if stateExtAttrs == nil {
			resp.Diagnostics.AddError(
				"Missing Internal ID",
				"Unable to read Filterrelayagent because the internal ID (from extattrs_all) is missing or invalid.",
			)
			return
		}

		stateTerraformId := (*stateExtAttrs)[terraformInternalIDEA]
		if apiTerraformId.Value != stateTerraformId.Value {
			if r.ReadByExtAttrs(ctx, &data, resp) {
				return
			}
		}
	}

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		resp.Diagnostics.AddError("Client Error", "Error while reading Filterrelayagent due to inherited Extensible attributes")
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FilterrelayagentResource) ReadByExtAttrs(ctx context.Context, data *FilterrelayagentModel, resp *resource.ReadResponse) bool {
	var diags diag.Diagnostics

	if data.ExtAttrsAll.IsNull() {
		return false
	}

	internalIdExtAttr := *ExpandExtAttrs(ctx, data.ExtAttrsAll, &diags)
	if diags.HasError() {
		return false
	}

	internalId := internalIdExtAttr[terraformInternalIDEA].Value
	if internalId == "" {
		return false
	}

	idMap := map[string]interface{}{
		terraformInternalIDEA: internalId,
	}

	apiRes, _, err := r.client.DHCPAPI.
		FilterrelayagentAPI.
		List(ctx).
		Extattrfilter(idMap).
		ReturnAsObject(1).
		ReturnFieldsPlus(readableAttributesForFilterrelayagent).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Filterrelayagent by extattrs, got error: %s", err))
		return true
	}

	results := apiRes.ListFilterrelayagentResponseObject.GetResult()

	// If the list is empty, the resource no longer exists so remove it from state
	if len(results) == 0 {
		resp.State.RemoveResource(ctx)
		return true
	}

	res := results[0]

	// Remove inherited external attributes from extattrs
	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		return true
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)

	return true
}

func (r *FilterrelayagentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data FilterrelayagentModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	planExtAttrs := data.ExtAttrs
	diags = req.State.GetAttribute(ctx, path.Root("ref"), &data.Ref)
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
		FilterrelayagentAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Filterrelayagent(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForFilterrelayagent).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Filterrelayagent, got error: %s", err))
		return
	}

	res := apiRes.UpdateFilterrelayagentResponseAsObject.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, planExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		resp.Diagnostics.AddError("Client Error", "Error while updating Filterrelayagent due to inherited Extensible attributes")
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if associateInternalId != nil {
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", nil)...)
	}
}

func (r *FilterrelayagentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data FilterrelayagentModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.DHCPAPI.
		FilterrelayagentAPI.
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Filterrelayagent, got error: %s", err))
		return
	}
}

func (r *FilterrelayagentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("ref"), req.ID)...)
	resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", []byte("true"))...)
}

func (r *FilterrelayagentResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data FilterrelayagentModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get the values or default to empty string
	isCircuitId := "ANY"
	if !data.IsCircuitId.IsNull() && !data.IsCircuitId.IsUnknown() {
		isCircuitId = data.IsCircuitId.ValueString()
	}

	isRemoteId := "ANY"
	if !data.IsRemoteId.IsNull() && !data.IsRemoteId.IsUnknown() {
		isRemoteId = data.IsRemoteId.ValueString()
	}

	// Validate that at least one must be set to a non-empty, non-ANY value
	// Both cannot be unset/empty at the same time (which API treats as "ANY")
	if isCircuitId == "ANY" && isRemoteId == "ANY" {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"At least one of is_circuit_id or is_remote_id must be set to 'MATCHES_VALUE' or 'NOT_SET'. Both cannot be set to 'ANY' at the same time.",
		)
		return
	}

	// Validate circuit_id_name is required when is_circuit_id == "MATCHES_VALUE"
	if isCircuitId == "MATCHES_VALUE" {
		if data.CircuitIdName.IsUnknown() || data.CircuitIdName.IsNull() || data.CircuitIdName.ValueString() == "" {
			resp.Diagnostics.AddAttributeError(
				path.Root("circuit_id_name"),
				"Missing Required Attribute",
				"Attribute circuit_id_name is required when is_circuit_id is set to 'MATCHES_VALUE'.",
			)
		}
	}

	// Validate remote_id_name is required when is_remote_id == "MATCHES_VALUE"
	if isRemoteId == "MATCHES_VALUE" {
		if data.RemoteIdName.IsUnknown() || data.RemoteIdName.IsNull() || data.RemoteIdName.ValueString() == "" {
			resp.Diagnostics.AddAttributeError(
				path.Root("remote_id_name"),
				"Missing Required Attribute",
				"Attribute remote_id_name is required when is_remote_id is set to 'MATCHES_VALUE'.",
			)
		}
	}

	// Validate circuit_id_substring_length, circuit_id_substring_offset is required when is_circuit_id_substring == true
	if !data.IsCircuitIdSubstring.IsUnknown() && !data.IsCircuitIdSubstring.IsNull() && data.IsCircuitIdSubstring.ValueBool() {
		if (data.CircuitIdSubstringLength.IsUnknown() || data.CircuitIdSubstringLength.IsNull()) ||
			(data.CircuitIdSubstringOffset.IsUnknown() || data.CircuitIdSubstringOffset.IsNull()) {
			resp.Diagnostics.AddAttributeError(
				path.Root("circuit_id_substring_length"),
				"Missing Required Attribute",
				"Attribute circuit_id_substring_length and circuit_id_substring_offset are required when is_circuit_id_substring is set to true.",
			)
		}
		// Validate the circuit_id_substring_length is equal to the length of circuit_id_name
		circuitIdLength := len(data.CircuitIdName.ValueString())
		if !data.CircuitIdName.IsNull() && data.CircuitIdSubstringLength.ValueInt64() != int64(circuitIdLength) {
			resp.Diagnostics.AddAttributeError(
				path.Root("circuit_id_substring_length"),
				"Invalid Attribute Value",
				"Attribute circuit_id_substring_length must be equal to the length of circuit_id_name when is_circuit_id_substring is set to true.",
			)
		}
	}

	// Validate remote_id_substring_length is required when is_remote_id_substring == true
	if !data.IsRemoteIdSubstring.IsUnknown() && !data.IsRemoteIdSubstring.IsNull() && data.IsRemoteIdSubstring.ValueBool() {
		if (data.RemoteIdSubstringLength.IsUnknown() || data.RemoteIdSubstringLength.IsNull()) ||
			(data.RemoteIdSubstringOffset.IsUnknown() || data.RemoteIdSubstringOffset.IsNull()) {
			resp.Diagnostics.AddAttributeError(
				path.Root("remote_id_substring_length"),
				"Missing Required Attribute",
				"Attribute remote_id_substring_length and remote_id_substring_offset are required when is_remote_id_substring is set to true.",
			)
		}

		// Validate the remote_id_substring_length is equal to the length of remote_id_name
		remoteIdLength := len(data.RemoteIdName.ValueString())
		if !data.RemoteIdName.IsNull() && data.RemoteIdSubstringLength.ValueInt64() != int64(remoteIdLength) {
			resp.Diagnostics.AddAttributeError(
				path.Root("remote_id_substring_length"),
				"Invalid Attribute Value",
				"Attribute remote_id_substring_length must be equal to the length of remote_id_name when is_remote_id_substring is set to true.",
			)
		}
	}

	// Validate is_circuit_id_substring is required when is_circuit_id == "MATCHES_VALUE"
	if isCircuitId == "MATCHES_VALUE" {
		if data.IsCircuitIdSubstring.IsUnknown() || data.IsCircuitIdSubstring.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("is_circuit_id_substring"),
				"Missing Required Attribute",
				"Attribute is_circuit_id_substring is required when is_circuit_id is set to 'MATCHES_VALUE'.",
			)
		}
	}

	// Validate is_remote_id_substring is required when is_remote_id == "MATCHES_VALUE"
	if isRemoteId == "MATCHES_VALUE" {
		if data.IsRemoteIdSubstring.IsUnknown() || data.IsRemoteIdSubstring.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("is_remote_id_substring"),
				"Missing Required Attribute",
				"Attribute is_remote_id_substring is required when is_remote_id is set to 'MATCHES_VALUE'.",
			)
		}
	}
}
