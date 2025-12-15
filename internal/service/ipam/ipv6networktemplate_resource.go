package ipam

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

var readableAttributesForIpv6networktemplate = "allow_any_netmask,auto_create_reversezone,cidr,cloud_api_compatible,comment,ddns_domainname,ddns_enable_option_fqdn,ddns_generate_hostname,ddns_server_always_updates,ddns_ttl,delegated_member,domain_name,domain_name_servers,enable_ddns,extattrs,fixed_address_templates,ipv6prefix,logic_filter_rules,members,name,options,preferred_lifetime,range_templates,recycle_leases,rir,rir_organization,rir_registration_action,rir_registration_status,send_rir_request,update_dns_on_lease_renewal,use_ddns_domainname,use_ddns_enable_option_fqdn,use_ddns_generate_hostname,use_ddns_ttl,use_domain_name,use_domain_name_servers,use_enable_ddns,use_logic_filter_rules,use_options,use_preferred_lifetime,use_recycle_leases,use_update_dns_on_lease_renewal,use_valid_lifetime,valid_lifetime"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &Ipv6networktemplateResource{}
var _ resource.ResourceWithImportState = &Ipv6networktemplateResource{}

func NewIpv6networktemplateResource() resource.Resource {
	return &Ipv6networktemplateResource{}
}

// Ipv6networktemplateResource defines the resource implementation.
type Ipv6networktemplateResource struct {
	client *niosclient.APIClient
}

func (r *Ipv6networktemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "ipam_ipv6networktemplate"
}

func (r *Ipv6networktemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages an IPv6 Network Template.",
		Attributes:          Ipv6networktemplateResourceSchemaAttributes,
	}
}

func (r *Ipv6networktemplateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *Ipv6networktemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var diags diag.Diagnostics
	var data Ipv6networktemplateModel

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

	apiRes, _, err := r.client.IPAMAPI.
		Ipv6networktemplateAPI.
		Create(ctx).
		Ipv6networktemplate(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForIpv6networktemplate).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Ipv6networktemplate, got error: %s", err))
		return
	}

	res := apiRes.CreateIpv6networktemplateResponseAsObject.GetResult()
	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while creating Ipv6networktemplate due to inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Ipv6networktemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var diags diag.Diagnostics
	var data Ipv6networktemplateModel

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

	apiRes, httpRes, err := r.client.IPAMAPI.
		Ipv6networktemplateAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForIpv6networktemplate).
		ReturnAsObject(1).
		Execute()

	// If the resource is not found, try searching using Extensible Attributes
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound && r.ReadByExtAttrs(ctx, &data, resp) {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Ipv6networktemplate, got error: %s", err))
		return
	}

	res := apiRes.GetIpv6networktemplateResponseObjectAsResult.GetResult()

	apiTerraformId, ok := (*res.ExtAttrs)[terraformInternalIDEA]
	if !ok {
		apiTerraformId.Value = ""
	}

	if associateInternalId == nil {
		stateExtAttrs := ExpandExtAttrs(ctx, data.ExtAttrsAll, &diags)
		if stateExtAttrs == nil {
			resp.Diagnostics.AddError(
				"Missing Internal ID",
				"Unable to read Ipv6networktemplate because the internal ID (from extattrs_all) is missing or invalid.",
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while reading Ipv6networktemplate due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Ipv6networktemplateResource) ReadByExtAttrs(ctx context.Context, data *Ipv6networktemplateModel, resp *resource.ReadResponse) bool {
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

	apiRes, _, err := r.client.IPAMAPI.
		Ipv6networktemplateAPI.
		List(ctx).
		Extattrfilter(idMap).
		ReturnAsObject(1).
		ReturnFieldsPlus(readableAttributesForIpv6networktemplate).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Ipv6networktemplate by extattrs, got error: %s", err))
		return true
	}

	results := apiRes.ListIpv6networktemplateResponseObject.GetResult()

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

func (r *Ipv6networktemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data Ipv6networktemplateModel

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

	apiRes, _, err := r.client.IPAMAPI.
		Ipv6networktemplateAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Ipv6networktemplate(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForIpv6networktemplate).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Ipv6networktemplate, got error: %s", err))
		return
	}

	res := apiRes.UpdateIpv6networktemplateResponseAsObject.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, planExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while updating Ipv6networktemplate due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if associateInternalId != nil {
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", nil)...)
	}
}

func (r *Ipv6networktemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Ipv6networktemplateModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.IPAMAPI.
		Ipv6networktemplateAPI.
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Ipv6networktemplate, got error: %s", err))
		return
	}
}

func (r *Ipv6networktemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("ref"), req.ID)...)
	resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", []byte("true"))...)
}
