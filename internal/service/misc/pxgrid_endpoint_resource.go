package misc

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"

	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForPxgridEndpoint = "address,client_certificate_subject,client_certificate_valid_from,client_certificate_valid_to,comment,disable,extattrs,log_level,name,network_view,outbound_member_type,outbound_members,publish_settings,subscribe_settings,template_instance,timeout,vendor_identifier,wapi_user_name"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &PxgridEndpointResource{}
var _ resource.ResourceWithImportState = &PxgridEndpointResource{}
var _ resource.ResourceWithValidateConfig = &PxgridEndpointResource{}

func NewPxgridEndpointResource() resource.Resource {
	return &PxgridEndpointResource{}
}

// PxgridEndpointResource defines the resource implementation.
type PxgridEndpointResource struct {
	client *niosclient.APIClient
}

func (r *PxgridEndpointResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "misc_pxgrid_endpoint"
}

func (r *PxgridEndpointResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a Pxgrid Endpoint.",
		Attributes:          PxgridEndpointResourceSchemaAttributes,
	}
}

func (r *PxgridEndpointResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *PxgridEndpointResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var diags diag.Diagnostics
	var data PxgridEndpointModel

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

	if !r.processClientCertificate(ctx, &data, &resp.Diagnostics) {
		return
	}

	apiRes, _, err := r.client.MiscAPI.
		PxgridEndpointAPI.
		Create(ctx).
		PxgridEndpoint(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForPxgridEndpoint).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create PxgridEndpoint, got error: %s", err))
		return
	}

	res := apiRes.CreatePxgridEndpointResponseAsObject.GetResult()
	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		resp.Diagnostics.AddError("Client Error", "Error while creating PxgridEndpoint due to inherited Extensible attributes")
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PxgridEndpointResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var diags diag.Diagnostics
	var data PxgridEndpointModel

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

	apiRes, httpRes, err := r.client.MiscAPI.
		PxgridEndpointAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForPxgridEndpoint).
		ReturnAsObject(1).
		Execute()

	// If the resource is not found, try searching using Extensible Attributes
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound && r.ReadByExtAttrs(ctx, &data, resp) {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read PxgridEndpoint, got error: %s", err))
		return
	}

	res := apiRes.GetPxgridEndpointResponseObjectAsResult.GetResult()

	apiTerraformId, ok := (*res.ExtAttrs)[terraformInternalIDEA]
	if !ok {
		apiTerraformId.Value = ""
	}

	if associateInternalId == nil {
		stateExtAttrs := ExpandExtAttrs(ctx, data.ExtAttrsAll, &diags)
		if stateExtAttrs == nil {
			resp.Diagnostics.AddError(
				"Missing Internal ID",
				"Unable to read PxgridEndpoint because the internal ID (from extattrs_all) is missing or invalid.",
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
		resp.Diagnostics.AddError("Client Error", "Error while reading PxgridEndpoint due to inherited Extensible attributes")
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PxgridEndpointResource) ReadByExtAttrs(ctx context.Context, data *PxgridEndpointModel, resp *resource.ReadResponse) bool {
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

	apiRes, _, err := r.client.MiscAPI.
		PxgridEndpointAPI.
		List(ctx).
		Extattrfilter(idMap).
		ReturnAsObject(1).
		ReturnFieldsPlus(readableAttributesForPxgridEndpoint).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read PxgridEndpoint by extattrs, got error: %s", err))
		return true
	}

	results := apiRes.ListPxgridEndpointResponseObject.GetResult()

	// If the list is empty, the resource no longer exists so remove it from state
	if len(results) == 0 {
		resp.State.RemoveResource(ctx)
		return true
	}

	res := results[0]

	// Remove inherited external attributes from extattrs
	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return true
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)

	return true
}

func (r *PxgridEndpointResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data PxgridEndpointModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if !r.processClientCertificate(ctx, &data, &resp.Diagnostics) {
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

	apiRes, _, err := r.client.MiscAPI.
		PxgridEndpointAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		PxgridEndpoint(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForPxgridEndpoint).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update PxgridEndpoint, got error: %s", err))
		return
	}

	res := apiRes.UpdatePxgridEndpointResponseAsObject.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, planExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		resp.Diagnostics.AddError("Client Error", "Error while updating PxgridEndpoint due to inherited Extensible attributes")
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if associateInternalId != nil {
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", nil)...)
	}
}

func (r *PxgridEndpointResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data PxgridEndpointModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.MiscAPI.
		PxgridEndpointAPI.
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete PxgridEndpoint, got error: %s", err))
		return
	}
}

func (r *PxgridEndpointResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data PxgridEndpointModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Validate that PublishSettings contains IPADDRESS in enabled_attributes
	if !data.PublishSettings.IsNull() && !data.PublishSettings.IsUnknown() {
		var publishSettings PxgridEndpointPublishSettingsModel
		resp.Diagnostics.Append(data.PublishSettings.As(ctx, &publishSettings, basetypes.ObjectAsOptions{})...)

		if resp.Diagnostics.HasError() {
			return
		}

		if !publishSettings.EnabledAttributes.IsNull() && !publishSettings.EnabledAttributes.IsUnknown() {
			var enabledAttrs []string
			resp.Diagnostics.Append(publishSettings.EnabledAttributes.ElementsAs(ctx, &enabledAttrs, false)...)

			if resp.Diagnostics.HasError() {
				return
			}

			hasIPAddress := slices.Contains(enabledAttrs, "IPADDRESS")

			if !hasIPAddress {
				resp.Diagnostics.AddAttributeError(
					path.Root("publish_settings").AtName("enabled_attributes"),
					"Invalid Configuration",
					"IP Address is a required publish data type.",
				)
			}
		}
	}

	// Outbound Members Validation
	hasOutboundMembers := !data.OutboundMembers.IsNull() && !data.OutboundMembers.IsUnknown()
	hasOutboundMemberType := !data.OutboundMemberType.IsNull() && !data.OutboundMemberType.IsUnknown()

	if hasOutboundMemberType {
		outboundMemberType := data.OutboundMemberType.ValueString()
		switch outboundMemberType {
		case "GM":
			if hasOutboundMembers {
				resp.Diagnostics.AddError(
					"Invalid Configuration",
					"'outbound_member_type' cannot be set to 'GM' when 'outbound_members' is specified.",
				)
			}
		case "MEMBER":
			if !hasOutboundMembers {
				resp.Diagnostics.AddError(
					"Invalid Configuration",
					"'outbound_member_type' cannot be set to 'MEMBER' when 'outbound_members' is not specified.",
				)
			}
		}
	}
}

func (r *PxgridEndpointResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("ref"), req.ID)...)
	resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", []byte("true"))...)
}

func (r *PxgridEndpointResource) processClientCertificate(
	ctx context.Context,
	data *PxgridEndpointModel,
	diag *diag.Diagnostics,
) bool {

	if data.ClientCertificateFile.IsNull() || data.ClientCertificateFile.IsUnknown() {
		return true
	}

	baseUrl := r.client.SecurityAPI.Cfg.NIOSHostURL
	username := r.client.SecurityAPI.Cfg.NIOSUsername
	password := r.client.SecurityAPI.Cfg.NIOSPassword

	filePath := data.ClientCertificateFile.ValueString()
	token, err := utils.UploadFileWithToken(ctx, baseUrl, filePath, username, password)
	if err != nil {
		diag.AddError(
			"Client Error",
			fmt.Sprintf("Unable to process certificate file %s, got error: %s", filePath, err),
		)
		return false
	}
	data.ClientCertificateToken = types.StringValue(token)
	return true
}
