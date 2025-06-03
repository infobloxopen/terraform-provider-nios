package dtc

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	niosclient "github.com/Infoblox-CTO/infoblox-nios-go-client/client"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/utils"
)

// TODO : Add readable attributes for the resource
var readableAttributesForDtcPool = "extattrs,lb_preferred_method,auto_consolidated_monitors,availability,comment,consolidated_monitors,disable,health,lb_alternate_method,lb_alternate_topology,lb_dynamic_ratio_alternate,lb_dynamic_ratio_preferred,lb_preferred_topology,name,quorum,servers,ttl,use_ttl,monitors"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &DtcpoolResource{}
var _ resource.ResourceWithImportState = &DtcpoolResource{}

func NewDtcpoolResource() resource.Resource {
	return &DtcpoolResource{}
}

// DtcpoolResource defines the resource implementation.
type DtcpoolResource struct {
	client *niosclient.APIClient
}

func (r *DtcpoolResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "resource_nios_DtcPool"
}

func (r *DtcpoolResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "",
		Attributes:          DtcPoolResourceSchemaAttributes,
	}
}

func (r *DtcpoolResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *DtcpoolResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var diags diag.Diagnostics
	var data DtcPoolModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Add internal ID exists in the Extensible Attributes if not already present
	if err := r.addInternalIDToExtAttrs(ctx, &data); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to add internal ID to Extensible Attributes, got error: %s", err))
		return
	}

	apiRes, _, err := r.client.DTCAPI.
		DtcpoolAPI.
		Post(ctx).
		DtcPool(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFields2(readableAttributesForDtcPool).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Dtcpool, got error: %s", err))
		return
	}

	res := apiRes.CreateDtcPoolResponseAsObject.GetResult()
	res.Extattrs, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.Extattrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while create Dtcpool due inherited Extensible attributes, got error: %s", err))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DtcpoolResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var diags diag.Diagnostics
	var data DtcPoolModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.DTCAPI.
		DtcpoolAPI.
		ReferenceGet(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFields2(readableAttributesForDtcPool).
		ReturnAsObject(1).
		Execute()

	// If the resource is not found, try searching using Extensible Attributes
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound && r.ReadByExtAttrs(ctx, &data, resp) {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Dtcpool, got error: %s", err))
		return
	}

	res := apiRes.GetDtcPoolResponseObjectAsResult.GetResult()
	res.Extattrs, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.Extattrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while reading Dtcpool due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DtcpoolResource) ReadByExtAttrs(ctx context.Context, data *DtcPoolModel, resp *resource.ReadResponse) bool {
	var diags diag.Diagnostics

	if data.ExtAttrsAll.IsNull() {
		return false
	}

	internalIdExtAttr := *ExpandExtAttr(ctx, data.ExtAttrsAll, &diags)
	if diags.HasError() {
		return false
	}

	internalId := internalIdExtAttr["Terraform Internal ID"].Value
	if internalId == "" {
		return false
	}

	idMap := map[string]interface{}{
		"Terraform Internal ID": internalId,
	}

	apiRes, httpRes, err := r.client.DTCAPI.
		DtcpoolAPI.
		Get(ctx).
		Extattrfilter(idMap).
		ReturnAsObject(1).
		ReturnFields2(readableAttributesForDtcPool).
		Execute()

	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)
			return true
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Dtcpool by extattrs, got error: %s", err))
		return true
	}

	if len(apiRes.ListDtcPoolResponseObject.GetResult()) > 0 {
		res := apiRes.ListDtcPoolResponseObject.GetResult()[0]

		// Remove inherited external attributes and check for errors
		res.Extattrs, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.Extattrs)
		if diags.HasError() {
			return true
		}

		data.Flatten(ctx, &res, &resp.Diagnostics)
		resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
	}
	return true
}

func (r *DtcpoolResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data DtcPoolModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Add internal ID exists in the Extensible Attributes if not already present
	if err := r.addInternalIDToExtAttrs(ctx, &data); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to add internal ID to Extensible Attributes, got error: %s", err))
		return
	}

	apiRes, _, err := r.client.DTCAPI.
		DtcpoolAPI.
		ReferencePut(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		DtcPool(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFields2(readableAttributesForDtcPool).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Dtcpool, got error: %s", err))
		return
	}

	res := apiRes.UpdateDtcPoolResponseAsObject.GetResult()

	res.Extattrs, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.Extattrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while update Dtcpool due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DtcpoolResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DtcPoolModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.DTCAPI.
		DtcpoolAPI.
		ReferenceDelete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Dtcpool, got error: %s", err))
		return
	}
}

func (r *DtcpoolResource) addInternalIDToExtAttrs(ctx context.Context, data *DtcPoolModel) error {
	_, exists := data.ExtAttrsAll.Elements()["Terraform Internal ID"]
	if exists {
		return nil
	}

	internalId, err := uuid.GenerateUUID()
	if err != nil {
		return err
	}

	// Inject default tag for update
	r.client.DTCAPI.APIClient.Cfg.DefaultExtAttrs = map[string]struct{ Value string }{
		"Terraform Internal ID": {Value: internalId},
	}

	return nil
}

func (r *DtcpoolResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("ref"), req, resp)
}
