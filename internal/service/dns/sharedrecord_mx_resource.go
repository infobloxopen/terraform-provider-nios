package dns

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForSharedrecordMx = "comment,disable,dns_mail_exchanger,dns_name,extattrs,mail_exchanger,name,preference,shared_record_group,ttl,use_ttl"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SharedrecordMxResource{}
var _ resource.ResourceWithImportState = &SharedrecordMxResource{}

func NewSharedrecordMxResource() resource.Resource {
	return &SharedrecordMxResource{}
}

// SharedrecordMxResource defines the resource implementation.
type SharedrecordMxResource struct {
	client *niosclient.APIClient
}

func (r *SharedrecordMxResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "dns_sharedrecord_mx"
}

func (r *SharedrecordMxResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a DNS MX shared record.",
		Attributes:          SharedrecordMxResourceSchemaAttributes,
	}
}

func (r *SharedrecordMxResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *SharedrecordMxResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var diags diag.Diagnostics
	var data SharedrecordMxModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Add internal ID exists in the Extensible Attributes if not already present
	// data.ExtAttrs, diags = AddInternalIDToExtAttrs(ctx, data.ExtAttrs, diags)
	// if diags.HasError() {
	// 	return
	// }

	apiRes, _, err := r.client.DNSAPI.
		SharedrecordMxAPI.
		Create(ctx).
		SharedrecordMx(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForSharedrecordMx).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create SharedrecordMx, got error: %s", err))
		return
	}

	res := apiRes.CreateSharedrecordMxResponseAsObject.GetResult()
	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while create SharedrecordMx due inherited Extensible attributes, got error: %s", err))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SharedrecordMxResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var diags diag.Diagnostics
	var data SharedrecordMxModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.DNSAPI.
		SharedrecordMxAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForSharedrecordMx).
		ReturnAsObject(1).
		Execute()

	// If the resource is not found, try searching using Extensible Attributes
	// if err != nil {
	// 	if httpRes != nil && httpRes.StatusCode == http.StatusNotFound { // && r.ReadByExtAttrs(ctx, &data, resp) {

	// 		return
	// 	}
	// 	resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read SharedrecordMx, got error: %s", err))
	// 	return
	// }
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			// Resource no longer exists, remove from state
			resp.State.RemoveResource(ctx)
			return
		}
	}
	res := apiRes.GetSharedrecordMxResponseObjectAsResult.GetResult()

	// apiTerraformId, ok := (*res.ExtAttrs)[terraformInternalIDEA]
	// if !ok {
	// 	apiTerraformId.Value = ""
	// }

	// stateExtAttrs := ExpandExtAttrs(ctx, data.ExtAttrsAll, &diags)
	// if stateExtAttrs == nil {
	// 	resp.Diagnostics.AddError(
	// 		"Missing Internal ID",
	// 		"Unable to read SharedrecordMx because the internal ID (from extattrs_all) is missing or invalid.",
	// 	)
	// 	return
	// }

	// stateTerraformId := (*stateExtAttrs)[terraformInternalIDEA]
	// if apiTerraformId.Value != stateTerraformId.Value {
	// 	if r.ReadByExtAttrs(ctx, &data, resp) {
	// 		return
	// 	}
	// }

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while reading SharedrecordMx due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// func (r *SharedrecordMxResource) ReadByExtAttrs(ctx context.Context, data *SharedrecordMxModel, resp *resource.ReadResponse) bool {
// 	var diags diag.Diagnostics

// 	if data.ExtAttrsAll.IsNull() {
// 		return false
// 	}

// 	internalIdExtAttr := *ExpandExtAttrs(ctx, data.ExtAttrsAll, &diags)
// 	if diags.HasError() {
// 		return false
// 	}

// 	internalId := internalIdExtAttr[terraformInternalIDEA].Value
// 	if internalId == "" {
// 		return false
// 	}

// 	idMap := map[string]interface{}{
// 		terraformInternalIDEA: internalId,
// 	}

// 	apiRes, _, err := r.client.DNSAPI.
// 		SharedrecordMxAPI.
// 		List(ctx).
// 		Extattrfilter(idMap).
// 		ReturnAsObject(1).
// 		ReturnFieldsPlus(readableAttributesForSharedrecordMx).
// 		Execute()
// 	if err != nil {
// 		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read SharedrecordMx by extattrs, got error: %s", err))
// 		return true
// 	}

// 	results := apiRes.ListSharedrecordMxResponseObject.GetResult()

// 	// If the list is empty, the resource no longer exists so remove it from state
// 	if len(results) == 0 {
// 		resp.State.RemoveResource(ctx)
// 		return true
// 	}

// 	res := results[0]

// 	// Remove inherited external attributes from extattrs
// 	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
// 	if diags.HasError() {
// 		return true
// 	}

// 	data.Flatten(ctx, &res, &resp.Diagnostics)
// 	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)

// 	return true
// }

func (r *SharedrecordMxResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data SharedrecordMxModel

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

	// Add Inherited Extensible Attributes
	data.ExtAttrs, diags = AddInheritedExtAttrs(ctx, data.ExtAttrs, data.ExtAttrsAll)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	apiRes, _, err := r.client.DNSAPI.
		SharedrecordMxAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		SharedrecordMx(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForSharedrecordMx).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update SharedrecordMx, got error: %s", err))
		return
	}

	res := apiRes.UpdateSharedrecordMxResponseAsObject.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, planExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while update SharedrecordMx due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SharedrecordMxResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SharedrecordMxModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.DNSAPI.
		SharedrecordMxAPI.
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete SharedrecordMx, got error: %s", err))
		return
	}
}

func (r *SharedrecordMxResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var diags diag.Diagnostics
	var data SharedrecordMxModel
	var goClientData dns.SharedrecordMx

	resourceRef := utils.ExtractResourceRef(req.ID)
	extattrs, diags := AddInternalIDToExtAttrs(ctx, data.ExtAttrs, diags)
	if diags.HasError() {
		return
	}
	goClientData.ExtAttrsPlus = ExpandExtAttrs(ctx, extattrs, &diags)
	data.ExtAttrsAll = extattrs

	updateRes, _, err := r.client.DNSAPI.
		SharedrecordMxAPI.
		Update(ctx, resourceRef).
		SharedrecordMx(goClientData).
		ReturnFieldsPlus(readableAttributesForSharedrecordMx).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Import Failed", fmt.Sprintf("Unable to update SharedrecordMx for import, got error: %s", err))
		return
	}

	res := updateRes.UpdateSharedrecordMxResponseAsObject.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrsAll, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while update SharedrecordMx due inherited Extensible attributes for import, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("ref"), req.ID)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("extattrs_all"), data.ExtAttrsAll)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("extattrs"), data.ExtAttrs)...)
}
