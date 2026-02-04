package rpz

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"

	"github.com/infobloxopen/terraform-provider-nios/internal/config"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForRecordRpzCnameClientipaddressdn = "canonical,comment,disable,extattrs,is_ipv4,name,rp_zone,ttl,use_ttl,view,zone"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &RecordRpzCnameClientipaddressdnResource{}
var _ resource.ResourceWithImportState = &RecordRpzCnameClientipaddressdnResource{}

func NewRecordRpzCnameClientipaddressdnResource() resource.Resource {
	return &RecordRpzCnameClientipaddressdnResource{}
}

// RecordRpzCnameClientipaddressdnResource defines the resource implementation.
type RecordRpzCnameClientipaddressdnResource struct {
	client *niosclient.APIClient
}

func (r *RecordRpzCnameClientipaddressdnResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "rpz_record_cname_clientipaddressdn"
}

func (r *RecordRpzCnameClientipaddressdnResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages an RPZ CNAME Client IP Address DN record.",
		Attributes:          RecordRpzCnameClientipaddressdnResourceSchemaAttributes,
	}
}

func (r *RecordRpzCnameClientipaddressdnResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *RecordRpzCnameClientipaddressdnResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var diags diag.Diagnostics
	var data RecordRpzCnameClientipaddressdnModel

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

	apiRes, _, err := r.client.RPZAPI.
		RecordRpzCnameClientipaddressdnAPI.
		Create(ctx).
		RecordRpzCnameClientipaddressdn(*data.Expand(ctx, &resp.Diagnostics, true)).
		ReturnFieldsPlus(readableAttributesForRecordRpzCnameClientipaddressdn).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create RecordRpzCnameClientipaddressdn, got error: %s", err))
		return
	}

	res := apiRes.CreateRecordRpzCnameClientipaddressdnResponseAsObject.GetResult()
	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		resp.Diagnostics.AddError("Client Error", "Error while creating RecordRpzCnameClientipaddressdn due to inherited Extensible attributes")
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RecordRpzCnameClientipaddressdnResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var diags diag.Diagnostics
	var data RecordRpzCnameClientipaddressdnModel

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

	apiRes, httpRes, err := r.client.RPZAPI.
		RecordRpzCnameClientipaddressdnAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForRecordRpzCnameClientipaddressdn).
		ReturnAsObject(1).
		ProxySearch(config.GetProxySearch()).
		Execute()

	// If the resource is not found, try searching using Extensible Attributes
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound && r.ReadByExtAttrs(ctx, &data, resp) {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read RecordRpzCnameClientipaddressdn, got error: %s", err))
		return
	}

	res := apiRes.GetRecordRpzCnameClientipaddressdnResponseObjectAsResult.GetResult()

	apiTerraformId, ok := (*res.ExtAttrs)[terraformInternalIDEA]
	if !ok {
		apiTerraformId.Value = ""
	}

	if associateInternalId == nil {
		stateExtAttrs := ExpandExtAttrs(ctx, data.ExtAttrsAll, &diags)
		if stateExtAttrs == nil {
			resp.Diagnostics.AddError(
				"Missing Internal ID",
				"Unable to read RecordRpzCnameClientipaddressdn because the internal ID (from extattrs_all) is missing or invalid.",
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
		resp.Diagnostics.AddError("Client Error", "Error while reading RecordRpzCnameClientipaddressdn due to inherited Extensible attributes")
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RecordRpzCnameClientipaddressdnResource) ReadByExtAttrs(ctx context.Context, data *RecordRpzCnameClientipaddressdnModel, resp *resource.ReadResponse) bool {
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

	apiRes, _, err := r.client.RPZAPI.
		RecordRpzCnameClientipaddressdnAPI.
		List(ctx).
		Extattrfilter(idMap).
		ReturnAsObject(1).
		ReturnFieldsPlus(readableAttributesForRecordRpzCnameClientipaddressdn).
		ProxySearch(config.GetProxySearch()).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read RecordRpzCnameClientipaddressdn by extattrs, got error: %s", err))
		return true
	}

	results := apiRes.ListRecordRpzCnameClientipaddressdnResponseObject.GetResult()

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

func (r *RecordRpzCnameClientipaddressdnResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data RecordRpzCnameClientipaddressdnModel

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

	apiRes, _, err := r.client.RPZAPI.
		RecordRpzCnameClientipaddressdnAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		RecordRpzCnameClientipaddressdn(*data.Expand(ctx, &resp.Diagnostics, false)).
		ReturnFieldsPlus(readableAttributesForRecordRpzCnameClientipaddressdn).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update RecordRpzCnameClientipaddressdn, got error: %s", err))
		return
	}

	res := apiRes.UpdateRecordRpzCnameClientipaddressdnResponseAsObject.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, planExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		resp.Diagnostics.AddError("Client Error", "Error while updating RecordRpzCnameClientipaddressdn due to inherited Extensible attributes")
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if associateInternalId != nil {
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", nil)...)
	}
}

func (r *RecordRpzCnameClientipaddressdnResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data RecordRpzCnameClientipaddressdnModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.RPZAPI.
		RecordRpzCnameClientipaddressdnAPI.
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete RecordRpzCnameClientipaddressdn, got error: %s", err))
		return
	}
}

func (r *RecordRpzCnameClientipaddressdnResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("ref"), req.ID)...)
	resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", []byte("true"))...)
}
