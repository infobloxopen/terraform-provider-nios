package dhcp

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForIpv6fixedaddress = "address_type,allow_telnet,cli_credentials,cloud_info,comment,device_description,device_location,device_type,device_vendor,disable,disable_discovery,discover_now_status,discovered_data,domain_name,domain_name_servers,duid,extattrs,ipv6addr,ipv6prefix,ipv6prefix_bits,logic_filter_rules,mac_address,match_client,ms_ad_user_data,name,network,network_view,options,preferred_lifetime,reserved_interface,snmp3_credential,snmp_credential,use_cli_credentials,use_domain_name,use_domain_name_servers,use_logic_filter_rules,use_options,use_preferred_lifetime,use_snmp3_credential,use_snmp_credential,use_valid_lifetime,valid_lifetime"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &Ipv6fixedaddressResource{}
var _ resource.ResourceWithImportState = &Ipv6fixedaddressResource{}
var _ resource.ResourceWithValidateConfig = &Ipv6fixedaddressResource{}

func NewIpv6fixedaddressResource() resource.Resource {
	return &Ipv6fixedaddressResource{}
}

// Ipv6fixedaddressResource defines the resource implementation.
type Ipv6fixedaddressResource struct {
	client *niosclient.APIClient
}

func (r *Ipv6fixedaddressResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "dhcp_ipv6fixedaddress"
}

func (r *Ipv6fixedaddressResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a DHCP IPV6 Fixed Address.",
		Attributes:          Ipv6fixedaddressResourceSchemaAttributes,
	}
}

func (r *Ipv6fixedaddressResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *Ipv6fixedaddressResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var diags diag.Diagnostics
	var data Ipv6fixedaddressModel

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
		Ipv6fixedaddressAPI.
		Create(ctx).
		Ipv6fixedaddress(*data.Expand(ctx, &resp.Diagnostics, true)).
		ReturnFieldsPlus(readableAttributesForIpv6fixedaddress).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Ipv6fixedaddress, got error: %s", err))
		return
	}

	res := apiRes.CreateIpv6fixedaddressResponseAsObject.GetResult()
	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while create Ipv6fixedaddress due inherited Extensible attributes, got error: %s", err))
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

func (r *Ipv6fixedaddressResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var diags diag.Diagnostics
	var data Ipv6fixedaddressModel

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

	apiRes, httpRes, err := r.client.DHCPAPI.
		Ipv6fixedaddressAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForIpv6fixedaddress).
		ReturnAsObject(1).
		Execute()

	// If the resource is not found, try searching using Extensible Attributes
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound && r.ReadByExtAttrs(ctx, &data, resp) {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Ipv6fixedaddress, got error: %s", err))
		return
	}

	res := apiRes.GetIpv6fixedaddressResponseObjectAsResult.GetResult()

	apiTerraformId, ok := (*res.ExtAttrs)[terraformInternalIDEA]
	if !ok {
		apiTerraformId.Value = ""
	}

	if associateInternalId == nil {
		stateExtAttrs := ExpandExtAttrs(ctx, data.ExtAttrsAll, &diags)
		if stateExtAttrs == nil {
			resp.Diagnostics.AddError(
				"Missing Internal ID",
				"Unable to read Ipv6fixedaddress because the internal ID (from extattrs_all) is missing or invalid.",
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while reading Ipv6fixedaddress due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Ipv6fixedaddressResource) ReadByExtAttrs(ctx context.Context, data *Ipv6fixedaddressModel, resp *resource.ReadResponse) bool {
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
		Ipv6fixedaddressAPI.
		List(ctx).
		Extattrfilter(idMap).
		ReturnAsObject(1).
		ReturnFieldsPlus(readableAttributesForIpv6fixedaddress).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Ipv6fixedaddress by extattrs, got error: %s", err))
		return true
	}

	results := apiRes.ListIpv6fixedaddressResponseObject.GetResult()

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

func (r *Ipv6fixedaddressResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data Ipv6fixedaddressModel

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

	apiRes, _, err := r.client.DHCPAPI.
		Ipv6fixedaddressAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Ipv6fixedaddress(*data.Expand(ctx, &resp.Diagnostics, false)).
		ReturnFieldsPlus(readableAttributesForIpv6fixedaddress).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Ipv6fixedaddress, got error: %s", err))
		return
	}

	res := apiRes.UpdateIpv6fixedaddressResponseAsObject.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, planExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while update Ipv6fixedaddress due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if associateInternalId != nil {
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", nil)...)
	}
}

func (r *Ipv6fixedaddressResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Ipv6fixedaddressModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.DHCPAPI.
		Ipv6fixedaddressAPI.
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Ipv6fixedaddress, got error: %s", err))
		return
	}
}

func (r *Ipv6fixedaddressResource) UpdateFuncCallAttributeName(ctx context.Context, data Ipv6fixedaddressModel, diags *diag.Diagnostics) types.Object {

	updatedFuncCallAttrs := data.FuncCall.Attributes()
	attrVal := updatedFuncCallAttrs["attribute_name"].(types.String).ValueString()
	pathVar, err := utils.FindModelFieldByTFSdkTag(data, attrVal)
	if !err {
		diags.AddError("Client Error", fmt.Sprintf("Unable to find attribute '%s' in Ipv6fixedaddress model, got error", attrVal))
		return types.ObjectNull(FuncCallAttrTypes)
	}
	updatedFuncCallAttrs["attribute_name"] = types.StringValue(pathVar)

	return types.ObjectValueMust(FuncCallAttrTypes, updatedFuncCallAttrs)
}

func (r *Ipv6fixedaddressResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("ref"), req.ID)...)
	resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", []byte("true"))...)
}

func (r *Ipv6fixedaddressResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data Ipv6fixedaddressModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Validate based on address_type only if func_call is not set
	if data.FuncCall.IsNull() || data.FuncCall.IsUnknown() {
		addressType := data.AddressType.ValueString()
		if addressType == "" || data.AddressType.IsNull() || data.AddressType.IsUnknown() {
			addressType = "ADDRESS"
		}

		switch addressType {
		case "ADDRESS":
			if data.Ipv6addr.IsNull() || data.Ipv6addr.IsUnknown() {
				resp.Diagnostics.AddAttributeError(
					path.Root("ipv6addr"),
					"Missing Required Attribute",
					"When address_type is set to 'ADDRESS', the 'ipv6addr' attribute must be specified.",
				)
			}
		case "PREFIX":
			if data.Ipv6prefix.IsNull() || data.Ipv6prefix.IsUnknown() {
				resp.Diagnostics.AddAttributeError(
					path.Root("ipv6prefix"),
					"Missing Required Attribute",
					"When address_type is set to 'PREFIX', the 'ipv6prefix' attribute must be specified.",
				)
			}
			if data.Ipv6prefixBits.IsNull() || data.Ipv6prefixBits.IsUnknown() {
				resp.Diagnostics.AddAttributeError(
					path.Root("ipv6prefix_bits"),
					"Missing Required Attribute",
					"When address_type is set to 'PREFIX', the 'ipv6prefix_bits' attribute must be specified.",
				)
			}
		case "BOTH":
			if data.Ipv6addr.IsNull() || data.Ipv6addr.IsUnknown() {
				resp.Diagnostics.AddAttributeError(
					path.Root("ipv6addr"),
					"Missing Required Attribute",
					"When address_type is set to 'BOTH', the 'ipv6addr' attribute must be specified.",
				)
			}
			if data.Ipv6prefix.IsNull() || data.Ipv6prefix.IsUnknown() {
				resp.Diagnostics.AddAttributeError(
					path.Root("ipv6prefix"),
					"Missing Required Attribute",
					"When address_type is set to 'BOTH', the 'ipv6prefix' attribute must be specified.",
				)
			}
			if data.Ipv6prefixBits.IsNull() || data.Ipv6prefixBits.IsUnknown() {
				resp.Diagnostics.AddAttributeError(
					path.Root("ipv6prefix_bits"),
					"Missing Required Attribute",
					"When address_type is set to 'BOTH', the 'ipv6prefix_bits' attribute must be specified.",
				)
			}
		}
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

		var options []Ipv6fixedaddressOptionsModel
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

	// Check if allow_telnet is true, then cli_credentials must contain at least one element with credentials_type set to "TELNET"
	if !data.AllowTelnet.IsUnknown() && !data.AllowTelnet.IsNull() && data.AllowTelnet.ValueBool() {
		isTelnet := false
		isSSH := false
		if !data.CliCredentials.IsNull() && !data.CliCredentials.IsUnknown() {
			// Iterate through cli_credentials to check if an element has credentials_type set to "TELNET"
			var cliCredentials []Ipv6fixedaddressCliCredentialsModel
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
				if credentialsType == "SSH" {
					isSSH = true
				}
				if credentialsType == "TELNET" {
					isTelnet = true
				}
			}
		}
		if !isSSH {
			resp.Diagnostics.AddAttributeError(
				path.Root("allow_telnet"),
				"Invalid configuration",
				"The 'cli_credentials' must contain credentials with 'credentials_type' set to 'SSH'.",
			)
		}
		if !isTelnet {
			resp.Diagnostics.AddAttributeError(
				path.Root("allow_telnet"),
				"Invalid configuration",
				"The 'allow_telnet' attribute must be set to false when 'cli_credentials' does not contain any credentials with 'credentials_type' set to 'TELNET'.",
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

	var options []Ipv6fixedaddressOptionsModel
	diags := data.Options.ElementsAs(ctx, &options, false)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// domain_name attribute must match the value of option 'domain-name'
	if !data.DomainName.IsNull() && !data.DomainName.IsUnknown() && !data.Options.IsNull() && !data.Options.IsUnknown() {
		for i, option := range options {
			if !option.Name.IsNull() && !option.Name.IsUnknown() && option.Name.ValueString() == "domain-name" {
				if !option.Value.IsNull() && !option.Value.IsUnknown() &&
					option.Value.ValueString() != data.DomainName.ValueString() {
					resp.Diagnostics.AddAttributeError(
						path.Root("options").AtListIndex(i).AtName("value"),
						"Invalid configuration for Domain Name",
						"domain_name attribute must match the 'value' attribute for DHCP Option 'domain-name'.",
					)
				}
			}
		}
	}

	// When dhcp-lease-time option is set, valid_lifetime attribute must have the same value as option value
	if !data.ValidLifetime.IsNull() && !data.ValidLifetime.IsUnknown() && !data.Options.IsNull() && !data.Options.IsUnknown() {
		for i, option := range options {
			if !option.Name.IsNull() && !option.Name.IsUnknown() && option.Name.ValueString() == "dhcp-lease-time" {
				if !option.Value.IsNull() && !option.Value.IsUnknown() &&
					option.Value.ValueString() != strconv.FormatInt(data.ValidLifetime.ValueInt64(), 10) {
					resp.Diagnostics.AddAttributeError(
						path.Root("options").AtListIndex(i).AtName("value"),
						"Invalid configuration for Valid Lifetime",
						"valid_lifetime attribute must match the 'value' attribute for DHCP Option 'dhcp-lease-time'.",
					)
				}
			}
		}
	}
}
