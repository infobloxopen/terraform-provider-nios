package dtc

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/infoblox-nios-go-client/dtc"

	"github.com/infobloxopen/terraform-provider-nios/internal/config"
)

var readableAttributesForDtcTopology = "comment,extattrs,name,rules"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &DtcTopologyResource{}
var _ resource.ResourceWithImportState = &DtcTopologyResource{}

func NewDtcTopologyResource() resource.Resource {
	return &DtcTopologyResource{}
}

// DtcTopologyResource defines the resource implementation.
type DtcTopologyResource struct {
	client *niosclient.APIClient
}

func (r *DtcTopologyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "dtc_topology"
}

func (r *DtcTopologyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a DTC Topology",
		Attributes:          DtcTopologyResourceSchemaAttributes,
	}
}

func (r *DtcTopologyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *DtcTopologyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var diags diag.Diagnostics
	var data DtcTopologyModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Add internal ID exists in the Extensible Attributes if not already present
	data.ExtAttrs, diags = AddInternalIDToExtAttrs(ctx, data.ExtAttrs, diags)
	if diags.HasError() {
		return
	}

	apiRes, _, err := r.client.DTCAPI.
		DtcTopologyAPI.
		Create(ctx).
		DtcTopology(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForDtcTopology).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create DtcTopology, got error: %s", err))
		return
	}

	res := apiRes.CreateDtcTopologyResponseAsObject.GetResult()
	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while create DtcTopology due inherited Extensible attributes, got error: %s", err))
		return
	}

	r.populateTopologyRules(ctx, &res, &diags)

	if diags.HasError() {
		return
	}
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DtcTopologyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var diags diag.Diagnostics
	var data DtcTopologyModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.DTCAPI.
		DtcTopologyAPI.
		Read(ctx, data.Uuid.ValueString()).
		ReturnFieldsPlus(readableAttributesForDtcTopology).
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read DtcTopology, got error: %s", err))
		return
	}

	res := apiRes.GetDtcTopologyResponseObjectAsResult.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while reading DtcTopology due inherited Extensible attributes, got error: %s", diags))
		return
	}

	r.populateTopologyRules(ctx, &res, &diags)

	if diags.HasError() {
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}


func (r *DtcTopologyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data DtcTopologyModel

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
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}
	if associateInternalId != nil {
		data.ExtAttrs, diags = AddInternalIDToExtAttrs(ctx, data.ExtAttrs, diags)
		if diags.HasError() {
			return
		}
	}

	// Add Inherited Extensible Attributes
	data.ExtAttrs, diags = AddInheritedExtAttrs(ctx, data.ExtAttrs, data.ExtAttrsAll)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	apiRes, _, err := r.client.DTCAPI.
		DtcTopologyAPI.
		Update(ctx, data.Uuid.ValueString()).
		DtcTopology(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForDtcTopology).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update DtcTopology, got error: %s", err))
		return
	}

	res := apiRes.UpdateDtcTopologyResponseAsObject.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, planExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while update DtcTopology due inherited Extensible attributes, got error: %s", diags))
		return
	}

	r.populateTopologyRules(ctx, &res, &diags)

	if diags.HasError() {
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if associateInternalId != nil {
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", nil)...)
	}
}

func (r *DtcTopologyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DtcTopologyModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.DTCAPI.
		DtcTopologyAPI.
		Delete(ctx, data.Uuid.ValueString()).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete DtcTopology, got error: %s", err))
		return
	}
}

func (r *DtcTopologyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("uuid"), req.ID)...)
	resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", []byte("true"))...)
}

func (r *DtcTopologyResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data DtcTopologyModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Rules.IsNull() || data.Rules.IsUnknown() {
		return
	}

	rules := data.Rules.Elements()

	firstDestType := ""

	for _, rule := range rules {
		ruleObj, ok := rule.(types.Object)
		if !ok {
			resp.Diagnostics.AddError("Type Assertion Error", fmt.Sprintf("Expected types.Object, got: %T", rule))
			return
		}

		destTypeAttr, exists := ruleObj.Attributes()["dest_type"]
		if !exists {
			continue
		}

		if destValue, ok := destTypeAttr.(types.String); ok {
			destType := destValue.ValueString()

			if firstDestType == "" {
				firstDestType = destType
			} else if firstDestType != destType {
				resp.Diagnostics.AddError("The Topology resource cannot have rules with different dest_type values", fmt.Sprintf("Found different dest_type values: %s and %s.", firstDestType, destType))
				return
			}
		}
	}
}

func UpdateDtcTopologyRules(ctx context.Context, r *DtcTopologyResource, ruleUuid string, diags *diag.Diagnostics) *dtc.DtcTopologyRulesInnerOneOf1 {
	apiRes, _, err := r.client.DTCAPI.
		DtcTopologyRuleAPI.
		Read(ctx, ruleUuid).
		ReturnFieldsPlus(readableAttributesForDtcTopologyRule).
		ReturnAsObject(1).
		Execute()

	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read DTC Topology Rules %s", err))
	}
	res := apiRes.GetDtcTopologyRuleResponseObjectAsResult.GetResult()

	ruleData := &dtc.DtcTopologyRulesInnerOneOf1{}

	if destType, ok := res.GetDestTypeOk(); ok {
		ruleData.SetDestType(*destType)
	}

	if destLink, ok := res.GetDestinationLinkOk(); ok {
		ruleData.SetDestinationLink(*destLink.DtcTopologyRuleDestinationLinkOneOf.Ref)
	}

	if returnType, ok := res.GetReturnTypeOk(); ok {
		ruleData.SetReturnType(*returnType)
	}

	if topology, ok := res.GetTopologyOk(); ok {
		ruleData.SetTopology(*topology)
	}

	if valid, ok := res.GetValidOk(); ok {
		ruleData.SetValid(*valid)
	}

	if sources, ok := res.GetSourcesOk(); ok {
		convertedSources := make([]dtc.DtcTopologyRulesInnerOneOf1SourcesInner, len(sources))
		for i, source := range sources {
			innerSource := dtc.DtcTopologyRulesInnerOneOf1SourcesInner{}

			if sourceOp, ok := source.GetSourceOpOk(); ok {
				innerSource.SourceOp = sourceOp
			}
			if sourceType, ok := source.GetSourceTypeOk(); ok {
				innerSource.SourceType = sourceType
			}
			if sourceValue, ok := source.GetSourceValueOk(); ok {
				innerSource.SourceValue = sourceValue
			}

			convertedSources[i] = innerSource
		}
		ruleData.SetSources(convertedSources)
	}

	return ruleData
}

func (r *DtcTopologyResource) populateTopologyRules(ctx context.Context, res *dtc.DtcTopology, diags *diag.Diagnostics) {
	for i, rule := range res.Rules {
		ruleUuid := rule.DtcTopologyRulesInnerOneOf.Uuid
		if ruleUuid == nil {
			continue
		}
		res.Rules[i].DtcTopologyRulesInnerOneOf1 = UpdateDtcTopologyRules(ctx, r, *ruleUuid, diags)
	}
}
