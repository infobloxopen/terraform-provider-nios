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
	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForRangetemplate = "bootfile,bootserver,cloud_api_compatible,comment,ddns_domainname,ddns_generate_hostname,delegated_member,deny_all_clients,deny_bootp,email_list,enable_ddns,enable_dhcp_thresholds,enable_email_warnings,enable_pxe_lease_time,enable_snmp_warnings,exclude,extattrs,failover_association,fingerprint_filter_rules,high_water_mark,high_water_mark_reset,ignore_dhcp_option_list_request,known_clients,lease_scavenge_time,logic_filter_rules,low_water_mark,low_water_mark_reset,mac_filter_rules,member,ms_options,ms_server,nac_filter_rules,name,nextserver,number_of_addresses,offset,option_filter_rules,options,pxe_lease_time,recycle_leases,relay_agent_filter_rules,server_association_type,unknown_clients,update_dns_on_lease_renewal,use_bootfile,use_bootserver,use_ddns_domainname,use_ddns_generate_hostname,use_deny_bootp,use_email_list,use_enable_ddns,use_enable_dhcp_thresholds,use_ignore_dhcp_option_list_request,use_known_clients,use_lease_scavenge_time,use_logic_filter_rules,use_ms_options,use_nextserver,use_options,use_pxe_lease_time,use_recycle_leases,use_unknown_clients,use_update_dns_on_lease_renewal"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &RangetemplateResource{}
var _ resource.ResourceWithImportState = &RangetemplateResource{}

func NewRangetemplateResource() resource.Resource {
	return &RangetemplateResource{}
}

// RangetemplateResource defines the resource implementation.
type RangetemplateResource struct {
	client *niosclient.APIClient
}

func (r *RangetemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "dhcp_range_template"
}

func (r *RangetemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a DHCP Range Template.",
		Attributes:          RangetemplateResourceSchemaAttributes,
	}
}

func (r *RangetemplateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *RangetemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var diags diag.Diagnostics
	var data RangetemplateModel

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

	apiRes, _, err := r.client.DHCPAPI.
		RangetemplateAPI.
		Create(ctx).
		Rangetemplate(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForRangetemplate).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Rangetemplate, got error: %s", err))
		return
	}

	res := apiRes.CreateRangetemplateResponseAsObject.GetResult()
	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while create Rangetemplate due inherited Extensible attributes, got error: %s", err))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RangetemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var diags diag.Diagnostics
	var data RangetemplateModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.DHCPAPI.
		RangetemplateAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForRangetemplate).
		ReturnAsObject(1).
		Execute()

	// If the resource is not found, try searching using Extensible Attributes
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound && r.ReadByExtAttrs(ctx, &data, resp) {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Rangetemplate, got error: %s", err))
		return
	}

	res := apiRes.GetRangetemplateResponseObjectAsResult.GetResult()

	apiTerraformId, ok := (*res.ExtAttrs)[terraformInternalIDEA]
	if !ok {
		apiTerraformId.Value = ""
	}

	stateExtAttrs := ExpandExtAttrs(ctx, data.ExtAttrsAll, &diags)
	if stateExtAttrs == nil {
		resp.Diagnostics.AddError(
			"Missing Internal ID",
			"Unable to read Rangetemplate because the internal ID (from extattrs_all) is missing or invalid.",
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while reading Rangetemplate due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RangetemplateResource) ReadByExtAttrs(ctx context.Context, data *RangetemplateModel, resp *resource.ReadResponse) bool {
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
		RangetemplateAPI.
		List(ctx).
		Extattrfilter(idMap).
		ReturnAsObject(1).
		ReturnFieldsPlus(readableAttributesForRangetemplate).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Rangetemplate by extattrs, got error: %s", err))
		return true
	}

	results := apiRes.ListRangetemplateResponseObject.GetResult()

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

func (r *RangetemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data RangetemplateModel

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

	apiRes, _, err := r.client.DHCPAPI.
		RangetemplateAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Rangetemplate(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForRangetemplate).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Rangetemplate, got error: %s", err))
		return
	}

	res := apiRes.UpdateRangetemplateResponseAsObject.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, planExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while update Rangetemplate due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RangetemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data RangetemplateModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.DHCPAPI.
		RangetemplateAPI.
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Rangetemplate, got error: %s", err))
		return
	}
}

func (r *RangetemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var diags diag.Diagnostics
	var data RangetemplateModel
	var goClientData dhcp.Rangetemplate

	resourceRef := utils.ExtractResourceRef(req.ID)
	extattrs, diags := AddInternalIDToExtAttrs(ctx, data.ExtAttrs, diags)
	if diags.HasError() {
		return
	}
	goClientData.ExtAttrsPlus = ExpandExtAttrs(ctx, extattrs, &diags)
	data.ExtAttrsAll = extattrs

	updateRes, _, err := r.client.DHCPAPI.
		RangetemplateAPI.
		Update(ctx, resourceRef).
		Rangetemplate(goClientData).
		ReturnFieldsPlus(readableAttributesForRangetemplate).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Import Failed", fmt.Sprintf("Unable to update Rangetemplate for import, got error: %s", err))
		return
	}

	res := updateRes.UpdateRangetemplateResponseAsObject.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrsAll, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while update Rangetemplate due inherited Extensible attributes for import, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("ref"), req.ID)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("extattrs_all"), data.ExtAttrsAll)...)
}

func (r *RangetemplateResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data RangetemplateModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Check if options are defined
	if !data.Options.IsNull() && !data.Options.IsUnknown() {
		// Special DHCP option names that require use_option to be set
		specialOptions := map[string]bool{
			"routers":                  true,
			"router-templates":         true,
			"domain-name-servers":      true,
			"domain-name":              true,
			"broadcast-address":        true,
			"broadcast-address-offset": true,
			"dhcp-lease-time":          true,
			"dhcp6.name-servers":       true,
		}

		specialOptionsNum := map[int64]bool{
			3:  true,
			6:  true,
			15: true,
			28: true,
			51: true,
			23: true,
		}

		var options []RangetemplateOptionsModel
		diags := data.Options.ElementsAs(ctx, &options, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		for i, option := range options {
			isSpecialOption := false
			optionName := ""
			if option.Value.IsNull() || option.Value.IsUnknown() {
				resp.Diagnostics.AddAttributeError(
					path.Root("options").AtListIndex(i).AtName("value"),
					"Invalid configuration for DHCP Option",
					"The 'value' attribute is a required field and must be set for all DHCP Options.",
				)
			}
			if !option.Name.IsNull() && !option.Name.IsUnknown() {
				optionName = option.Name.ValueString()
				isSpecialOption = specialOptions[optionName]
			} else if !option.Num.IsNull() && !option.Num.IsUnknown() {
				optionNum := option.Num.ValueInt64()
				isSpecialOption = specialOptionsNum[optionNum]
				optionName = fmt.Sprintf("with num = %d", optionNum)
			} else {
				resp.Diagnostics.AddAttributeError(
					path.Root("options").AtListIndex(i).AtName("name"),
					"Invalid configuration for DHCP Option",
					"Either the 'name' or 'num' attribute must be set for all DHCP Options. "+
						"Missing both attributes for 'option' at index "+fmt.Sprint(i)+".",
				)
				continue
			}

			if option.Value.ValueString() == "" {
				if !isSpecialOption {
					resp.Diagnostics.AddAttributeError(
						path.Root("options").AtListIndex(i).AtName("value"),
						"Invalid configuration for DHCP Option",
						"The 'value' attribute cannot be set as empty for Custom DHCP Option '"+optionName+"'.",
					)
				} else if !option.UseOption.IsUnknown() && !option.UseOption.IsNull() && !option.UseOption.ValueBool() {
					resp.Diagnostics.AddAttributeError(
						path.Root("options").AtListIndex(i).AtName("value"),
						"Invalid configuration for DHCP Option",
						"The 'value' attribute cannot be set as empty for Special DHCP Option '"+optionName+"' when 'use_option' is set to false.",
					)
				}
			}

			if !isSpecialOption && !option.UseOption.IsNull() && !option.UseOption.IsUnknown() {
				resp.Diagnostics.AddAttributeError(
					path.Root("options").AtListIndex(i).AtName("use_option"),
					"Invalid configuration",
					fmt.Sprintf("The 'use_option' attribute should not be set for Custom DHCP Option '%s'. "+
						"It is only applicable for Special Options: routers, router-templates, domain-name-servers, "+
						"domain-name, broadcast-address, broadcast-address-offset, dhcp-lease-time, dhcp6.name-servers.",
						optionName),
				)
			}
		}
	}
}
