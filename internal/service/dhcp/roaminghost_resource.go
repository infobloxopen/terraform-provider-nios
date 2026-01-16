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

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"

	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForRoaminghost = "address_type,bootfile,bootserver,client_identifier_prepend_zero,comment,ddns_domainname,ddns_hostname,deny_bootp,dhcp_client_identifier,disable,enable_ddns,enable_pxe_lease_time,extattrs,force_roaming_hostname,ignore_dhcp_option_list_request,ipv6_client_hostname,ipv6_ddns_domainname,ipv6_ddns_hostname,ipv6_domain_name,ipv6_domain_name_servers,ipv6_duid,ipv6_enable_ddns,ipv6_force_roaming_hostname,ipv6_mac_address,ipv6_match_option,ipv6_options,mac,match_client,name,network_view,nextserver,options,preferred_lifetime,pxe_lease_time,use_bootfile,use_bootserver,use_ddns_domainname,use_deny_bootp,use_enable_ddns,use_ignore_dhcp_option_list_request,use_ipv6_ddns_domainname,use_ipv6_domain_name,use_ipv6_domain_name_servers,use_ipv6_enable_ddns,use_ipv6_options,use_nextserver,use_options,use_preferred_lifetime,use_pxe_lease_time,use_valid_lifetime,valid_lifetime"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &RoaminghostResource{}
var _ resource.ResourceWithImportState = &RoaminghostResource{}
var _ resource.ResourceWithValidateConfig = &RoaminghostResource{}

func NewRoaminghostResource() resource.Resource {
	return &RoaminghostResource{}
}

// RoaminghostResource defines the resource implementation.
type RoaminghostResource struct {
	client *niosclient.APIClient
}

func (r *RoaminghostResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "dhcp_roaminghost"
}

func (r *RoaminghostResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a DHCP Roaming Host.",
		Attributes:          RoaminghostResourceSchemaAttributes,
	}
}

func (r *RoaminghostResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *RoaminghostResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var diags diag.Diagnostics
	var data RoaminghostModel

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

	apiRes, _, err := r.client.DHCPAPI.
		RoaminghostAPI.
		Create(ctx).
		Roaminghost(*data.Expand(ctx, &resp.Diagnostics, true)).
		ReturnFieldsPlus(readableAttributesForRoaminghost).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Roaminghost, got error: %s", err))
		return
	}

	res := apiRes.CreateRoaminghostResponseAsObject.GetResult()
	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		resp.Diagnostics.AddError("Client Error", "Error while creating Roaminghost due to inherited Extensible attributes")
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RoaminghostResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var diags diag.Diagnostics
	var data RoaminghostModel

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
		RoaminghostAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForRoaminghost).
		ReturnAsObject(1).
		Execute()

	// If the resource is not found, try searching using Extensible Attributes
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound && r.ReadByExtAttrs(ctx, &data, resp) {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Roaminghost, got error: %s", err))
		return
	}

	res := apiRes.GetRoaminghostResponseObjectAsResult.GetResult()

	apiTerraformId, ok := (*res.ExtAttrs)[terraformInternalIDEA]
	if !ok {
		apiTerraformId.Value = ""
	}

	if associateInternalId == nil {
		stateExtAttrs := ExpandExtAttrs(ctx, data.ExtAttrsAll, &diags)
		if stateExtAttrs == nil {
			resp.Diagnostics.AddError(
				"Missing Internal ID",
				"Unable to read Roaminghost because the internal ID (from extattrs_all) is missing or invalid.",
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
		resp.Diagnostics.AddError("Client Error", "Error while reading Roaminghost due to inherited Extensible attributes")
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RoaminghostResource) ReadByExtAttrs(ctx context.Context, data *RoaminghostModel, resp *resource.ReadResponse) bool {
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
		RoaminghostAPI.
		List(ctx).
		Extattrfilter(idMap).
		ReturnAsObject(1).
		ReturnFieldsPlus(readableAttributesForRoaminghost).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Roaminghost by extattrs, got error: %s", err))
		return true
	}

	results := apiRes.ListRoaminghostResponseObject.GetResult()

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

func (r *RoaminghostResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data RoaminghostModel

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

	apiRes, _, err := r.client.DHCPAPI.
		RoaminghostAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Roaminghost(*data.Expand(ctx, &resp.Diagnostics, false)).
		ReturnFieldsPlus(readableAttributesForRoaminghost).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Roaminghost, got error: %s", err))
		return
	}

	res := apiRes.UpdateRoaminghostResponseAsObject.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, planExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		resp.Diagnostics.AddError("Client Error", "Error while updating Roaminghost due to inherited Extensible attributes")
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if associateInternalId != nil {
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", nil)...)
	}
}

func (r *RoaminghostResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data RoaminghostModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.DHCPAPI.
		RoaminghostAPI.
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Roaminghost, got error: %s", err))
		return
	}
}

func (r RoaminghostResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data RoaminghostModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var options []RoaminghostOptionsModel
	diags := data.Options.ElementsAs(ctx, &options, false)
	resp.Diagnostics.Append(diags...)
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

	addressType := "IPV4"
	if !data.AddressType.IsNull() && !data.AddressType.IsUnknown() {
		addressType = data.AddressType.ValueString()
	}

	switch addressType {
	case "IPV4":
		// When address_type is IPV4, match_client is required
		if data.MatchClient.IsNull() || data.MatchClient.IsUnknown() || data.MatchClient.ValueString() == "" {
			resp.Diagnostics.AddAttributeError(
				path.Root("match_client"),
				"Invalid configuration for Match Client",
				"When 'address_type' is set to 'IPV4' (default), the 'match_client' attribute is required.",
			)
		}
	case "IPV6":
		// When address_type is IPV6, ipv6_match_option is required
		if data.Ipv6MatchOption.IsNull() || data.Ipv6MatchOption.IsUnknown() || data.Ipv6MatchOption.ValueString() == "" {
			resp.Diagnostics.AddAttributeError(
				path.Root("ipv6_match_option"),
				"Invalid configuration for IPv6 Match Option",
				"When 'address_type' is set to 'IPV6', the 'ipv6_match_option' attribute is required.",
			)
		}
	case "BOTH":
		// When address_type is BOTH, both match_client and ipv6_match_option are required
		if data.MatchClient.IsNull() || data.MatchClient.IsUnknown() || data.MatchClient.ValueString() == "" ||
			data.Ipv6MatchOption.IsNull() || data.Ipv6MatchOption.IsUnknown() || data.Ipv6MatchOption.ValueString() == "" {
			resp.Diagnostics.AddAttributeError(
				path.Root("match_client"),
				"Invalid configuration for Match Client",
				"When 'address_type' is set to 'BOTH', both 'match_client' and 'ipv6_match_option' attributes are required.",
			)
		}
	}

	var matchClient string
	if !data.MatchClient.IsNull() && !data.MatchClient.IsUnknown() {
		matchClient = data.MatchClient.ValueString()
	}

	switch matchClient {
	case "MAC_ADDRESS":
		// When match_client is MAC_ADDRESS, mac is required and dhcp_identifier should not be set
		if data.Mac.IsNull() || data.Mac.IsUnknown() || data.Mac.ValueString() == "" {
			resp.Diagnostics.AddAttributeError(
				path.Root("mac"),
				"Invalid configuration for MAC Address",
				"When 'match_client' is set to 'MAC_ADDRESS', the 'mac' attribute is required.",
			)
		}
		if !data.DhcpClientIdentifier.IsNull() && !data.DhcpClientIdentifier.IsUnknown() && data.DhcpClientIdentifier.ValueString() != "" {
			resp.Diagnostics.AddAttributeError(
				path.Root("dhcp_client_identifier"),
				"Invalid configuration for DHCP Client Identifier",
				"When 'match_client' is set to 'MAC_ADDRESS', the 'dhcp_client_identifier' attribute should not be set.",
			)
		}
	case "CLIENT_ID":
		// When match_client is CLIENT_ID, dhcp_identifier is required and mac should not be set
		if data.DhcpClientIdentifier.IsNull() || data.DhcpClientIdentifier.IsUnknown() || data.DhcpClientIdentifier.ValueString() == "" {
			resp.Diagnostics.AddAttributeError(
				path.Root("dhcp_client_identifier"),
				"Invalid configuration for DHCP Client Identifier",
				"When 'match_client' is set to 'CLIENT_ID', the 'dhcp_client_identifier' attribute is required.",
			)
		}
		if !data.Mac.IsNull() && !data.Mac.IsUnknown() && data.Mac.ValueString() != "" {
			resp.Diagnostics.AddAttributeError(
				path.Root("mac"),
				"Invalid configuration for MAC Address",
				"When 'match_client' is set to 'CLIENT_ID', the 'mac' attribute should not be set.",
			)
		}
	}

	var ipv6MatchOption string
	if !data.Ipv6MatchOption.IsNull() && !data.Ipv6MatchOption.IsUnknown() {
		ipv6MatchOption = data.Ipv6MatchOption.ValueString()
	}

	switch ipv6MatchOption {
	case "V6_MAC_ADDRESS":
		// When ipv6_match_option is V6_MAC_ADDRESS, ipv6_mac_address is required and ipv6_duid should not be set
		if data.Ipv6MacAddress.IsNull() || data.Ipv6MacAddress.IsUnknown() || data.Ipv6MacAddress.ValueString() == "" {
			resp.Diagnostics.AddAttributeError(
				path.Root("ipv6_mac_address"),
				"Invalid configuration for IPv6 MAC Address",
				"When 'ipv6_match_option' is set to 'V6_MAC_ADDRESS', the 'ipv6_mac_address' attribute is required.",
			)
		}
		if !data.Ipv6Duid.IsNull() && !data.Ipv6Duid.IsUnknown() && data.Ipv6Duid.ValueString() != "" {
			resp.Diagnostics.AddAttributeError(
				path.Root("ipv6_duid"),
				"Invalid configuration for IPv6 DUID",
				"When 'ipv6_match_option' is set to 'V6_MAC_ADDRESS', the 'ipv6_duid' attribute should not be set.",
			)
		}
	case "DUID":
		// When ipv6_match_option is DUID, ipv6_duid is required and ipv6_mac_address should not be set
		if data.Ipv6Duid.IsNull() || data.Ipv6Duid.IsUnknown() || data.Ipv6Duid.ValueString() == "" {
			resp.Diagnostics.AddAttributeError(
				path.Root("ipv6_duid"),
				"Invalid configuration for IPv6 DUID",
				"When 'ipv6_match_option' is set to 'DUID', the 'ipv6_duid' attribute is required.",
			)
		}
		if !data.Ipv6MacAddress.IsNull() && !data.Ipv6MacAddress.IsUnknown() && data.Ipv6MacAddress.ValueString() != "" {
			resp.Diagnostics.AddAttributeError(
				path.Root("ipv6_mac_address"),
				"Invalid configuration for IPv6 MAC Address",
				"When 'ipv6_match_option' is set to 'DUID', the 'ipv6_mac_address' attribute should not be set.",
			)
		}
	}

	// Preferred lifetime must be less than or equal to valid lifetime
	if !data.PreferredLifetime.IsNull() && !data.PreferredLifetime.IsUnknown() &&
		!data.ValidLifetime.IsNull() && !data.ValidLifetime.IsUnknown() {
		if data.PreferredLifetime.ValueInt64() > data.ValidLifetime.ValueInt64() {
			resp.Diagnostics.AddAttributeError(
				path.Root("preferred_lifetime"),
				"Invalid configuration",
				"The 'preferred_lifetime' must be less than or equal to 'valid_lifetime'.",
			)
		}
	} else if !data.Options.IsNull() && !data.Options.IsUnknown() {
		for _, option := range options {
			if !option.Name.IsNull() && !option.Name.IsUnknown() && option.Name.ValueString() == "dhcp-lease-time" {
				if !option.Value.IsNull() && !option.Value.IsUnknown() &&
					strconv.FormatInt(data.PreferredLifetime.ValueInt64(), 10) > option.Value.ValueString() {
					resp.Diagnostics.AddAttributeError(
						path.Root("preferred_lifetime"),
						"Invalid configuration",
						"The 'preferred_lifetime' must be less than or equal to 'dhcp-lease-time' (valid_lifetime) option value.",
					)
				}
			}
		}
	}
}

func (r *RoaminghostResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("ref"), req.ID)...)
	resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", []byte("true"))...)
}
