package dns

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"net/http"
	"os"
	"strings"
	"time"

	niosclient "github.com/Infoblox-CTO/infoblox-nios-go-client/client"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/utils"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var OperationTimeout = 30 * time.Second
var d *RecordADataSource
var r *RecordAResource

// TODO : Add readable attributes for the resource
var readableAttributesForRecordCname = "aws_rte53_record_info,canonical,cloud_info,comment,creation_time,creator,ddns_principal,ddns_protected,disable,dns_canonical,dns_name,extattrs,forbid_reclamation,last_queried,name,reclaimable,shared_record_group,ttl,use_ttl,view,zone"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &RecordCnameResource{}
var _ resource.ResourceWithImportState = &RecordCnameResource{}

func NewRecordCnameResource() resource.Resource {
	return &RecordCnameResource{}
}

// RecordCnameResource defines the resource implementation.
type RecordCnameResource struct {
	client *niosclient.APIClient
}

func (r *RecordCnameResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "dns_record_cname"
}

func (r *RecordCnameResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "",
		Attributes:          RecordCnameResourceSchemaAttributes,
	}
}

func (r *RecordCnameResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *RecordCnameResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var diags diag.Diagnostics
	var data RecordCnameModel

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

	filters := map[string]interface{}{
		"name": data.Name,
	}
	err := retry.RetryContext(ctx, OperationTimeout, func() *retry.RetryError {
		apiRes, _, err := r.client.DNSAPI.
			RecordCnameAPI.
			Create(ctx).
			RecordCname(*data.Expand(ctx, &resp.Diagnostics)).
			ReturnFieldsPlus(readableAttributesForRecordCname).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			if strings.Contains(err.Error(), "already exists.") {
				tflog.Debug(ctx, "Waiting for state to stabilize, will retry", map[string]interface{}{"error": err.Error()})
				fdip := os.Getenv("FDIP")
				dimbt := os.Getenv("DIMBT")
				if checkFDIP(fdip) {
					if checkRecords(ctx, diags, filters, dimbt) {
						return retry.RetryableError(err)
					}
				}
				//apiRes2, httpRes, err2 := d.client.DNSAPI.
				//	RecordAAPI.
				//	List(ctx).
				//	Extattrfilter(filters).
				//	ReturnAsObject(1).
				//	ReturnFieldsPlus(readableAttributesForRecordA).
				//	Execute()
				//if err2 == nil && len(apiRes2.ListRecordAResponseObject.GetResult()) > 0 {
				//	httpRes3, err3 := r.client.DNSAPI.
				//		RecordAAPI.
				//		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
				//		Execute()
				//	if err != nil {
				//		if httpRes3 != nil && httpRes3.StatusCode == http.StatusNotFound {
				//			return retry.RetryableError(err3)
				//		}
				//		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Found Record A with same name , unable to delete, got error: %s", err))
				//		return retry.NonRetryableError(err)
				//	}
				//}
				//if err2 != nil {
				//	if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
				//		resp.State.RemoveResource(ctx)
				//		return
				//	}
				//	resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read RecordA, got error: %s", err))
				//	return
				//}
				//return retry.RetryableError(err)
			}
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Record, got error: %s", err))
			return retry.NonRetryableError(err)
		}
		res := apiRes.CreateRecordCnameResponseAsObject.GetResult()
		res.ExtAttrs, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
		if diags.HasError() {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while create RecordCname due inherited Extensible attributes, got error: %s", err))
			return nil
		}
		data.Flatten(ctx, &res, &resp.Diagnostics)

		return nil
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create RecordCname due to error: %s", err))
		return
	}

	//apiRes, _, err := r.client.DNSAPI.
	//	RecordCnameAPI.
	//	Create(ctx).
	//	RecordCname(*data.Expand(ctx, &resp.Diagnostics)).
	//	ReturnFieldsPlus(readableAttributesForRecordCname).
	//	ReturnAsObject(1).
	//	Execute()
	//if err != nil {
	//	resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create RecordCname, got error: %s", err))
	//	return
	//}
	//
	//res := apiRes.CreateRecordCnameResponseAsObject.GetResult()
	//res.ExtAttrs, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	//if diags.HasError() {
	//	resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while create RecordCname due inherited Extensible attributes, got error: %s", err))
	//}
	//data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RecordCnameResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var diags diag.Diagnostics
	var data RecordCnameModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.DNSAPI.
		RecordCnameAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForRecordCname).
		ReturnAsObject(1).
		Execute()

	// If the resource is not found, try searching using Extensible Attributes
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound && r.ReadByExtAttrs(ctx, &data, resp) {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read RecordCname, got error: %s", err))
		return
	}

	res := apiRes.GetRecordCnameResponseObjectAsResult.GetResult()
	if res.ExtAttrs == nil {
		resp.Diagnostics.AddError(
			"Missing Extensible Attributes",
			"Unable to read RecordCname because no extensible attributes were returned from the API.",
		)
		return
	}

	res.ExtAttrs, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while reading RecordCname due inherited Extensible attributes, got error: %s", diags))
		return
	}

	apiTerraformId, ok := (*res.ExtAttrs)["Terraform Internal ID"]
	if !ok {
		resp.Diagnostics.AddError(
			"Missing Terraform internal id Attributes",
			"Unable to read RecordCname because terraform internal id does not exist.",
		)
		return
	}

	stateExtAttrs := ExpandExtAttr(ctx, data.ExtAttrsAll, &diags)
	if stateExtAttrs == nil {
		resp.Diagnostics.AddError(
			"Missing Internal ID",
			"Unable to read RecordCname because the internal ID (from extattrs_all) is missing or invalid.",
		)
		return
	}

	stateTerraformId := (*stateExtAttrs)["Terraform Internal ID"]

	if apiTerraformId.Value != stateTerraformId.Value {
		if r.ReadByExtAttrs(ctx, &data, resp) {
			return
		}
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RecordCnameResource) ReadByExtAttrs(ctx context.Context, data *RecordCnameModel, resp *resource.ReadResponse) bool {
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

	apiRes, httpRes, err := r.client.DNSAPI.
		RecordCnameAPI.
		List(ctx).
		Extattrfilter(idMap).
		ReturnAsObject(1).
		ReturnFieldsPlus(readableAttributesForRecordCname).
		Execute()

	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)
			return true
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read RecordCname by extattrs, got error: %s", err))
		return true
	}

	if len(apiRes.ListRecordCnameResponseObject.GetResult()) > 0 {
		res := apiRes.ListRecordCnameResponseObject.GetResult()[0]

		// Remove inherited external attributes and check for errors
		res.ExtAttrs, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
		if diags.HasError() {
			return true
		}

		data.Flatten(ctx, &res, &resp.Diagnostics)
		resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
	}
	return true
}

func (r *RecordCnameResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data RecordCnameModel

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

	diags = req.State.GetAttribute(ctx, path.Root("extattrs_all"), &data.ExtAttrsAll)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	// Add internal ID exists in the Extensible Attributes if not already present
	if err := r.addInternalIDToExtAttrs(ctx, &data); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to add internal ID to Extensible Attributes, got error: %s", err))
		return
	}

	apiRes, _, err := r.client.DNSAPI.
		RecordCnameAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		RecordCname(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForRecordCname).
		ReturnAsObject(1).
		Execute()

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update RecordCname, got error: %s", err))
		return
	}

	res := apiRes.UpdateRecordCnameResponseAsObject.GetResult()

	res.ExtAttrs, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while update RecordCname due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RecordCnameResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data RecordCnameModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.DNSAPI.
		RecordCnameAPI.
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete RecordCname, got error: %s", err))
		return
	}
}

func (r *RecordCnameResource) addInternalIDToExtAttrs(ctx context.Context, data *RecordCnameModel) error {
	var internalId string

	if !data.ExtAttrsAll.IsNull() {
		elements := data.ExtAttrsAll.Elements()
		if id, ok := elements["Terraform Internal ID"]; ok {
			if idStr, ok := id.(types.String); ok {
				internalId = idStr.ValueString()
			}
		}
	}

	if internalId == "" {
		var err error
		internalId, err = uuid.GenerateUUID()
		if err != nil {
			return err
		}
	}

	r.client.DNSAPI.APIClient.Cfg.DefaultExtAttrs = map[string]struct{ Value string }{
		"Terraform Internal ID": {Value: internalId},
	}

	return nil
}

func (r *RecordCnameResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("ref"), req, resp)
}

func checkFDIP(fdip string) bool {

	if fdip == "" {
		return false
	}
	fdipLower := strings.ToLower(fdip)
	containsA := strings.Contains(fdipLower, "a")
	containsAAAA := strings.Contains(fdipLower, "aaaa")
	containsCNAME := strings.Contains(fdipLower, "cname")
	containsTrue := strings.Contains(fdipLower, "true")
	if containsA || containsAAAA || containsCNAME || containsTrue {
		return true
	}
	return false
}

func checkRecords(ctx context.Context, diags diag.Diagnostics, filters map[string]interface{}, dimbt string) bool {
	records := map[string]string{}
	mbt := map[string]bool{}
	apiRes, _, err := d.client.DNSAPI.
		RecordAAPI.
		List(ctx).
		Extattrfilter(filters).
		ReturnAsObject(1).
		ReturnFieldsPlus(readableAttributesForRecordA).
		Execute()
	if err == nil && len(apiRes.ListRecordAResponseObject.GetResult()) > 0 {
		records["A"] = *apiRes.ListRecordAResponseObject.GetResult()[0].Ref
		mbt["A"] = false
		if apiRes.ListRecordAResponseObject.GetResult()[0].ExtAttrs != nil {
			for key, _ := range *apiRes.ListRecordAResponseObject.GetResult()[0].ExtAttrs {
				if key == "Terraform Internal ID" {
					mbt["A"] = true
				}
			}
		}
	}

	apiRes, _, err = d.client.DNSAPI.
		RecordAAPI.
		List(ctx).
		Extattrfilter(filters).
		ReturnAsObject(1).
		ReturnFieldsPlus(readableAttributesForRecordA).
		Execute()
	if err == nil && len(apiRes.ListRecordAResponseObject.GetResult()) > 0 {
		records["AAAA"] = *apiRes.ListRecordAResponseObject.GetResult()[0].Ref
		mbt["AAAA"] = false
		if apiRes.ListRecordAResponseObject.GetResult()[0].ExtAttrs != nil {
			for key, _ := range *apiRes.ListRecordAResponseObject.GetResult()[0].ExtAttrs {
				if key == "Terraform Internal ID" {
					mbt["AAAA"] = true
				}
			}
		}
	}

	if len(records) < 0 {
		return true
	} else {
		for key, val := range records {
			switch key {
			case "A":
				if mbt["A"] || dimbt == "true" {
					_, _ = r.client.DNSAPI.
						RecordAAPI.
						Delete(ctx, utils.ExtractResourceRef(val)).
						Execute()
					return true
				} else {
					diags.AddError("Client Error", fmt.Sprintf("Found Record A with same name, unable to delete as either the record is not managed by terraform or  DIMBT is unset"))
					//resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Found Record A with same name, unable to delete as either the record is not managed by terraform or  DIMBT is unset"))
					return false
				}
			case "AAAA":
				if mbt["AAAA"] || dimbt == "true" {
					_, _ = r.client.DNSAPI.
						RecordAAPI.
						Delete(ctx, utils.ExtractResourceRef(val)).
						Execute()
					return true
				} else {

					diags.AddError("Client Error", fmt.Sprintf("Found Record AAAA with same name, unable to delete as either the record is not managed by terraform or  DIMBT is unset"))
					return false
				}
			default:
				diags.AddError("Client Error", fmt.Sprintf("Invalid record type %s found, unable to delete", key))
				return false
			}
		}
	}
	return true
}
