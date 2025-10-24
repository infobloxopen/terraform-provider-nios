package notification

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/infoblox-nios-go-client/notification"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForNotificationRestEndpoint = "client_certificate_subject,client_certificate_valid_from,client_certificate_valid_to,comment,extattrs,log_level,name,outbound_member_type,outbound_members,server_cert_validation,sync_disabled,template_instance,timeout,uri,username,vendor_identifier,wapi_user_name"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NotificationRestEndpointResource{}
var _ resource.ResourceWithImportState = &NotificationRestEndpointResource{}
var _ resource.ResourceWithValidateConfig = &NotificationRestEndpointResource{}

func NewNotificationRestEndpointResource() resource.Resource {
	return &NotificationRestEndpointResource{}
}

// NotificationRestEndpointResource defines the resource implementation.
type NotificationRestEndpointResource struct {
	client *niosclient.APIClient
}

func (r *NotificationRestEndpointResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "notification_rest_endpoint"
}

func (r *NotificationRestEndpointResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a Notification REST Endpoint.",
		Attributes:          NotificationRestEndpointResourceSchemaAttributes,
	}
}

func (r *NotificationRestEndpointResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *NotificationRestEndpointResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data NotificationRestEndpointModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Outbound Members Validation
	if data.OutboundMemberType.ValueString() == "MEMBER" {
		if data.OutboundMembers.IsNull() || data.OutboundMembers.IsUnknown() {
			resp.Diagnostics.AddAttributeError(
				path.Root("outbound_members"),
				"Invalid Configuration",
				"Attribute 'outbound_members' must be specified when 'outbound_member_type' is set to 'MEMBER'.",
			)
		}
	} else if !data.OutboundMembers.IsNull() && !data.OutboundMembers.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("outbound_members"),
			"Invalid Configuration",
			"Attribute 'outbound_members' cannot be specified when 'outbound_member_type' is set to 'GM'.",
		)
	}

	// URI Validation
	uri := data.Uri.ValueString()
	_, err := url.ParseRequestURI(uri)
	if err != nil {
		resp.Diagnostics.AddAttributeError(
			path.Root("uri"),
			"Invalid URI",
			"URI must contain a valid value.",
		)
	}

	// Template Instance Parameters Validation
	if !data.TemplateInstance.IsNull() && !data.TemplateInstance.IsUnknown() {
		var templateInstanceModel NotificationRestEndpointTemplateInstanceModel

		resp.Diagnostics.Append(data.TemplateInstance.As(ctx, &templateInstanceModel, basetypes.ObjectAsOptions{})...)
		if resp.Diagnostics.HasError() {
			return
		}

		var templateInstanceParametersModel []NotificationrestendpointtemplateinstanceParametersModel

		resp.Diagnostics.Append(templateInstanceModel.Parameters.ElementsAs(ctx, &templateInstanceParametersModel, false)...)
		if resp.Diagnostics.HasError() {
			return
		}
		for i, param := range templateInstanceParametersModel {
			// Skip the validation if syntax or value is Null or Unknown
			if param.Syntax.IsNull() || param.Syntax.IsUnknown() ||
				param.Value.IsNull() || param.Value.IsUnknown() {
				continue
			}

			syntax := param.Syntax.ValueString()
			value := param.Value.ValueString()

			if syntax == "INT" {
				if _, err := strconv.Atoi(value); err != nil {
					resp.Diagnostics.AddAttributeError(
						path.Root("template_instance").AtName("parameters").AtListIndex(i).AtName("value"),
						"Invalid Value for INT Syntax",
						fmt.Sprintf("The value of the parameter definition '%s' is incorrect. The value type should be %s. Got: %s", param.Name.ValueString(), syntax, value),
					)
				}
			}
			if syntax == "BOOL" {
				if value != "True" && value != "False" {
					resp.Diagnostics.AddAttributeError(
						path.Root("template_instance").AtName("parameters").AtListIndex(i).AtName("value"),
						"Invalid Value for BOOL Syntax",
						fmt.Sprintf("The value of the parameter definition '%s' is incorrect. The value type should be %s, either True/False (Case Sensitive). Got: %s", param.Name.ValueString(), syntax, value),
					)
				}
			}
		}
	}
}

func (r *NotificationRestEndpointResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var diags diag.Diagnostics
	var data NotificationRestEndpointModel

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

	if !r.processClientCertificate(ctx, &data, &resp.Diagnostics) {
		return
	}

	apiRes, _, err := r.client.NotificationAPI.
		NotificationRestEndpointAPI.
		Create(ctx).
		NotificationRestEndpoint(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForNotificationRestEndpoint).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create NotificationRestEndpoint, got error: %s", err))
		return
	}

	res := apiRes.CreateNotificationRestEndpointResponseAsObject.GetResult()
	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while create NotificationRestEndpoint due inherited Extensible attributes, got error: %s", err))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NotificationRestEndpointResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var diags diag.Diagnostics
	var data NotificationRestEndpointModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.NotificationAPI.
		NotificationRestEndpointAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForNotificationRestEndpoint).
		ReturnAsObject(1).
		Execute()

	// If the resource is not found, try searching using Extensible Attributes
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound && r.ReadByExtAttrs(ctx, &data, resp) {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read NotificationRestEndpoint, got error: %s", err))
		return
	}

	res := apiRes.GetNotificationRestEndpointResponseObjectAsResult.GetResult()

	apiTerraformId, ok := (*res.ExtAttrs)[terraformInternalIDEA]
	if !ok {
		apiTerraformId.Value = ""
	}

	stateExtAttrs := ExpandExtAttrs(ctx, data.ExtAttrsAll, &diags)
	if stateExtAttrs == nil {
		resp.Diagnostics.AddError(
			"Missing Internal ID",
			"Unable to read NotificationRestEndpoint because the internal ID (from extattrs_all) is missing or invalid.",
		)
		return
	}

	stateTerraformId := (*stateExtAttrs)[terraformInternalIDEA]
	if apiTerraformId.Value != stateTerraformId.Value {
		if r.ReadByExtAttrs(ctx, &data, resp) {
			return
		}
	}

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while reading NotificationRestEndpoint due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NotificationRestEndpointResource) ReadByExtAttrs(ctx context.Context, data *NotificationRestEndpointModel, resp *resource.ReadResponse) bool {
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

	apiRes, _, err := r.client.NotificationAPI.
		NotificationRestEndpointAPI.
		List(ctx).
		Extattrfilter(idMap).
		ReturnAsObject(1).
		ReturnFieldsPlus(readableAttributesForNotificationRestEndpoint).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read NotificationRestEndpoint by extattrs, got error: %s", err))
		return true
	}

	results := apiRes.ListNotificationRestEndpointResponseObject.GetResult()

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

func (r *NotificationRestEndpointResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data NotificationRestEndpointModel

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

	// Add Inherited Extensible Attributes
	data.ExtAttrs, diags = AddInheritedExtAttrs(ctx, data.ExtAttrs, data.ExtAttrsAll)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	apiRes, _, err := r.client.NotificationAPI.
		NotificationRestEndpointAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		NotificationRestEndpoint(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForNotificationRestEndpoint).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update NotificationRestEndpoint, got error: %s", err))
		return
	}

	res := apiRes.UpdateNotificationRestEndpointResponseAsObject.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, planExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while update NotificationRestEndpoint due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NotificationRestEndpointResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NotificationRestEndpointModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.NotificationAPI.
		NotificationRestEndpointAPI.
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete NotificationRestEndpoint, got error: %s", err))
		return
	}
}

func (r *NotificationRestEndpointResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var diags diag.Diagnostics
	var data NotificationRestEndpointModel
	var goClientData notification.NotificationRestEndpoint

	resourceRef := utils.ExtractResourceRef(req.ID)
	extattrs, diags := AddInternalIDToExtAttrs(ctx, data.ExtAttrs, diags)
	if diags.HasError() {
		return
	}
	goClientData.ExtAttrsPlus = ExpandExtAttrs(ctx, extattrs, &diags)
	data.ExtAttrsAll = extattrs

	updateRes, _, err := r.client.NotificationAPI.
		NotificationRestEndpointAPI.
		Update(ctx, resourceRef).
		NotificationRestEndpoint(goClientData).
		ReturnFieldsPlus(readableAttributesForNotificationRestEndpoint).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Import Failed", fmt.Sprintf("Unable to update NotificationRestEndpoint for import, got error: %s", err))
		return
	}

	res := updateRes.UpdateNotificationRestEndpointResponseAsObject.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrsAll, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while update NotificationRestEndpoint due inherited Extensible attributes for import, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NotificationRestEndpointResource) processClientCertificate(
	ctx context.Context,
	data *NotificationRestEndpointModel,
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
