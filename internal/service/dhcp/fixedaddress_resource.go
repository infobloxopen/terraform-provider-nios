package dhcp

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
	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForFixedaddress = "agent_circuit_id,agent_remote_id,allow_telnet,always_update_dns,bootfile,bootserver,cli_credentials,client_identifier_prepend_zero,cloud_info,comment,ddns_domainname,ddns_hostname,deny_bootp,device_description,device_location,device_type,device_vendor,dhcp_client_identifier,disable,disable_discovery,discover_now_status,discovered_data,enable_ddns,enable_pxe_lease_time,extattrs,ignore_dhcp_option_list_request,ipv4addr,is_invalid_mac,logic_filter_rules,mac,match_client,ms_ad_user_data,ms_options,ms_server,name,network,network_view,nextserver,options,pxe_lease_time,reserved_interface,snmp3_credential,snmp_credential,use_bootfile,use_bootserver,use_cli_credentials,use_ddns_domainname,use_deny_bootp,use_enable_ddns,use_ignore_dhcp_option_list_request,use_logic_filter_rules,use_ms_options,use_nextserver,use_options,use_pxe_lease_time,use_snmp3_credential,use_snmp_credential"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &FixedaddressResource{}
var _ resource.ResourceWithImportState = &FixedaddressResource{}

func NewFixedaddressResource() resource.Resource {
	return &FixedaddressResource{}
}

// FixedaddressResource defines the resource implementation.
type FixedaddressResource struct {
	client *niosclient.APIClient
}

func (r *FixedaddressResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "dhcp_fixed_address"
}

func (r *FixedaddressResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a Fixed Address.",
		Attributes:          FixedaddressResourceSchemaAttributes,
	}
}

func (r *FixedaddressResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *FixedaddressResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var diags diag.Diagnostics
	var data FixedaddressModel

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

	// If the function call attributes are set, update the attribute name to match tfsdk tag
	origFunCallAttrs := data.FuncCall.Attributes()
	if len(origFunCallAttrs) > 0 {
		data.FuncCall = r.UpdateFuncCallAttributeName(ctx, data, &resp.Diagnostics)
	}

	apiRes, _, err := r.client.DHCPAPI.
		FixedaddressAPI.
		Create(ctx).
		Fixedaddress(*data.Expand(ctx, &resp.Diagnostics, true)).
		ReturnFieldsPlus(readableAttributesForFixedaddress).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Fixedaddress, got error: %s", err))
		return
	}

	res := apiRes.CreateFixedaddressResponseAsObject.GetResult()
	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while create Fixedaddress due inherited Extensible attributes, got error: %s", err))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Retain the original function call attributes
	if len(origFunCallAttrs) > 0 {
		data.FuncCall = types.ObjectValueMust(FuncCallAttrTypes, origFunCallAttrs)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FixedaddressResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var diags diag.Diagnostics
	var data FixedaddressModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.DHCPAPI.
		FixedaddressAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForFixedaddress).
		ReturnAsObject(1).
		Execute()

	// If the resource is not found, try searching using Extensible Attributes
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound && r.ReadByExtAttrs(ctx, &data, resp) {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Fixedaddress, got error: %s", err))
		return
	}

	res := apiRes.GetFixedaddressResponseObjectAsResult.GetResult()

	apiTerraformId, ok := (*res.ExtAttrs)[terraformInternalIDEA]
	if !ok {
		apiTerraformId.Value = ""
	}

	stateExtAttrs := ExpandExtAttrs(ctx, data.ExtAttrsAll, &diags)
	if stateExtAttrs == nil {
		resp.Diagnostics.AddError(
			"Missing Internal ID",
			"Unable to read Fixedaddress because the internal ID (from extattrs_all) is missing or invalid.",
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while reading Fixedaddress due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FixedaddressResource) ReadByExtAttrs(ctx context.Context, data *FixedaddressModel, resp *resource.ReadResponse) bool {
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
		FixedaddressAPI.
		List(ctx).
		Extattrfilter(idMap).
		ReturnAsObject(1).
		ReturnFieldsPlus(readableAttributesForFixedaddress).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Fixedaddress by extattrs, got error: %s", err))
		return true
	}

	results := apiRes.ListFixedaddressResponseObject.GetResult()

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

func (r *FixedaddressResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data FixedaddressModel

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
		FixedaddressAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Fixedaddress(*data.Expand(ctx, &resp.Diagnostics, false)).
		ReturnFieldsPlus(readableAttributesForFixedaddress).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Fixedaddress, got error: %s", err))
		return
	}

	res := apiRes.UpdateFixedaddressResponseAsObject.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, planExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while update Fixedaddress due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FixedaddressResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data FixedaddressModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.DHCPAPI.
		FixedaddressAPI.
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Fixedaddress, got error: %s", err))
		return
	}
}

func (r *FixedaddressResource) UpdateFuncCallAttributeName(ctx context.Context, data FixedaddressModel, diags *diag.Diagnostics) types.Object {

	updatedFuncCallAttrs := data.FuncCall.Attributes()
	attrVal := updatedFuncCallAttrs["attribute_name"].(types.String).ValueString()
	pathVar, err := utils.FindModelFieldByTFSdkTag(data, attrVal)
	if !err {
		diags.AddError("Client Error", fmt.Sprintf("Unable to find attribute '%s' in Fixedaddress model, got error", attrVal))
		return types.ObjectNull(FuncCallAttrTypes)
	}
	updatedFuncCallAttrs["attribute_name"] = types.StringValue(pathVar)

	return types.ObjectValueMust(FuncCallAttrTypes, updatedFuncCallAttrs)
}

func (r *FixedaddressResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var diags diag.Diagnostics
	var data FixedaddressModel
	var goClientData dhcp.Fixedaddress

	resourceRef := utils.ExtractResourceRef(req.ID)
	extattrs, diags := AddInternalIDToExtAttrs(ctx, data.ExtAttrs, diags)
	if diags.HasError() {
		return
	}
	goClientData.ExtAttrsPlus = ExpandExtAttrs(ctx, extattrs, &diags)
	data.ExtAttrsAll = extattrs

	updateRes, _, err := r.client.DHCPAPI.
		FixedaddressAPI.
		Update(ctx, resourceRef).
		Fixedaddress(goClientData).
		ReturnFieldsPlus(readableAttributesForFixedaddress).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Import Failed", fmt.Sprintf("Unable to update Fixedaddress for import, got error: %s", err))
		return
	}

	res := updateRes.UpdateFixedaddressResponseAsObject.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrsAll, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while update Fixedaddress due inherited Extensible attributes for import, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FixedaddressResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data FixedaddressModel
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
			3: true,
			//"router-templates":         true,
			6:  true,
			15: true,
			28: true,
			//"broadcast-address-offset": true,
			51: true,
			//"dhcp6.name-servers":       true,
		}

		var options []FixedaddressOptionsModel
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

	// Check if allow_telnet is true, then cli_credentials must contain at least one element with credentials_type set to "TELNET"
	if !data.AllowTelnet.IsUnknown() && !data.AllowTelnet.IsNull() && data.AllowTelnet.ValueBool() {
		isTelnet := false
		if !data.CliCredentials.IsNull() && !data.CliCredentials.IsUnknown() {
			// Iterate through cli_credentials to check if an element has credentials_type set to "TELNET"
			var cliCredentials []FixedaddressCliCredentialsModel
			diags := data.CliCredentials.ElementsAs(ctx, &cliCredentials, false)
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}

			for _, credentials := range cliCredentials {
				if credentials.CredentialType.IsUnknown() || credentials.CredentialType.IsNull() {
					continue
				}
				credentialsType := credentials.CredentialType.ValueString()
				if credentialsType == "TELNET" {
					isTelnet = true
					break
				}
			}
		}
		if !isTelnet {
			resp.Diagnostics.AddAttributeError(
				path.Root("allow_telnet"),
				"Invalid configuration",
				"The 'allow_telnet' attribute must be set to false when 'cli_credentials' is not set or does not contain any credentials with 'credentials_type' set to 'TELNET'.",
			)
		}
	}

	// Check if credentials are defined, then the corresponding use_cli_credentials attribute must be set to true
	if !data.CliCredentials.IsNull() && !data.CliCredentials.IsUnknown() {
		if !data.UseCliCredentials.IsUnknown() && !data.UseCliCredentials.IsNull() && !data.UseCliCredentials.ValueBool() {
			resp.Diagnostics.AddAttributeError(
				path.Root("cli_credentials"),
				"Invalid configuration",
				"The 'cli_credentials' attribute is set, but 'use_cli_credentials' is false. "+
					"Please set 'use_cli_credentials' to true to use CLI credentials.",
			)
		}
	}
	// Check if SNMP , then the corresponding use_snmp_credential attribute must be set to true
	if !data.SnmpCredential.IsUnknown() && !data.SnmpCredential.IsNull() {
		if !data.UseSnmpCredential.IsUnknown() && !data.UseSnmpCredential.IsNull() && !data.UseSnmpCredential.ValueBool() {
			resp.Diagnostics.AddAttributeError(
				path.Root("snmp_credential"),
				"Invalid configuration",
				"The 'snmp_credential' attribute is set, but 'use_snmp_credential' is false. "+
					"Please set 'use_snmp_credential' to true to use SNMP Credentials.",
			)
		}
	}

	// Check if SNMP3 credentials are set , then the corresponding use_snmp3_credential and use_cli_credentials attribute must be set to true
	if !data.Snmp3Credential.IsUnknown() && !data.Snmp3Credential.IsNull() {
		if !data.UseSnmp3Credential.IsUnknown() && !data.UseSnmp3Credential.IsNull() && !data.UseSnmp3Credential.ValueBool() {
			resp.Diagnostics.AddAttributeError(
				path.Root("snmp3_credential"),
				"Invalid configuration",
				"The 'snmp3_credential' attribute is set, but 'use_snmp3_credential' is false. "+
					"Please set 'use_snmp3_credential' to true to use SNMP3 Credentials.",
			)
		}
		if !data.UseCliCredentials.IsUnknown() && !data.UseCliCredentials.IsNull() && !data.UseCliCredentials.ValueBool() {
			resp.Diagnostics.AddAttributeError(
				path.Root("snmp3_credential"),
				"Invalid configuration",
				"The 'snmp3_credential' attribute is set, but 'use_cli_credentials' is false. "+
					"Please set 'use_cli_credentials' to true to use SNMP3 Credentials.",
			)
		}
	}
}
