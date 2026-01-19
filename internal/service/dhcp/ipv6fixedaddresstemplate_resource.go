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

var readableAttributesForIpv6fixedaddresstemplate = "comment,domain_name,domain_name_servers,extattrs,logic_filter_rules,name,number_of_addresses,offset,options,preferred_lifetime,use_domain_name,use_domain_name_servers,use_logic_filter_rules,use_options,use_preferred_lifetime,use_valid_lifetime,valid_lifetime"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &Ipv6fixedaddresstemplateResource{}
var _ resource.ResourceWithImportState = &Ipv6fixedaddresstemplateResource{}
var _ resource.ResourceWithValidateConfig = &Ipv6fixedaddresstemplateResource{}

func NewIpv6fixedaddresstemplateResource() resource.Resource {
	return &Ipv6fixedaddresstemplateResource{}
}

// Ipv6fixedaddresstemplateResource defines the resource implementation.
type Ipv6fixedaddresstemplateResource struct {
	client *niosclient.APIClient
}

func (r *Ipv6fixedaddresstemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "dhcp_ipv6fixedaddresstemplate"
}

func (r *Ipv6fixedaddresstemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a DHCP IPv6 Fixed Address Template.",
		Attributes:          Ipv6fixedaddresstemplateResourceSchemaAttributes,
	}
}

func (r *Ipv6fixedaddresstemplateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *Ipv6fixedaddresstemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var diags diag.Diagnostics
	var data Ipv6fixedaddresstemplateModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.DHCPAPI.
		Ipv6fixedaddresstemplateAPI.
		Create(ctx).
		Ipv6fixedaddresstemplate(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForIpv6fixedaddresstemplate).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Ipv6fixedaddresstemplate, got error: %s", err))
		return
	}

	res := apiRes.CreateIpv6fixedaddresstemplateResponseAsObject.GetResult()
	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while create Ipv6fixedaddresstemplate due inherited Extensible attributes, got error: %s", err))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Ipv6fixedaddresstemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var diags diag.Diagnostics
	var data Ipv6fixedaddresstemplateModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.DHCPAPI.
		Ipv6fixedaddresstemplateAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForIpv6fixedaddresstemplate).
		ReturnAsObject(1).
		Execute()

	// If the resource is not found, try searching using Extensible Attributes
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Ipv6fixedaddresstemplate, got error: %s", err))
		return
	}

	res := apiRes.GetIpv6fixedaddresstemplateResponseObjectAsResult.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while reading Ipv6fixedaddresstemplate due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Ipv6fixedaddresstemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data Ipv6fixedaddresstemplateModel

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
		Ipv6fixedaddresstemplateAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Ipv6fixedaddresstemplate(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForIpv6fixedaddresstemplate).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Ipv6fixedaddresstemplate, got error: %s", err))
		return
	}

	res := apiRes.UpdateIpv6fixedaddresstemplateResponseAsObject.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, planExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while update Ipv6fixedaddresstemplate due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Ipv6fixedaddresstemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Ipv6fixedaddresstemplateModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.DHCPAPI.
		Ipv6fixedaddresstemplateAPI.
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Ipv6fixedaddresstemplate, got error: %s", err))
		return
	}
}

func (r *Ipv6fixedaddresstemplateResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data Ipv6fixedaddresstemplateModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var options []Ipv6fixedaddresstemplateOptionsModel
	diags := data.Options.ElementsAs(ctx, &options, false)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var dhcpLeaseTimeValue string
	var hasDhcpLeaseTime bool

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

			if option.Name.ValueString() == "dhcp-lease-time" {
				hasDhcpLeaseTime = true
				dhcpLeaseTimeValue = option.Value.ValueString()
			}

			// domain_name attribute must match the value of option 'domain-name'
			if option.Name.ValueString() == "domain-name" {
				if !data.DomainName.IsNull() && !data.DomainName.IsUnknown() &&
					!option.Value.IsNull() && !option.Value.IsUnknown() &&
					option.Value.ValueString() != data.DomainName.ValueString() {
					resp.Diagnostics.AddAttributeError(
						path.Root("domain_name"),
						"Invalid configuration for Domain Name",
						"domain_name attribute must match the 'value' attribute for DHCP Option 'domain-name'.",
					)
				}
			}
		}

		// When dhcp-lease-time option is set, valid_lifetime attribute must have the same value as option value
		if hasDhcpLeaseTime && !data.ValidLifetime.IsNull() && !data.ValidLifetime.IsUnknown() {
			if dhcpLeaseTimeValue != strconv.FormatInt(data.ValidLifetime.ValueInt64(), 10) {
				resp.Diagnostics.AddAttributeError(
					path.Root("valid_lifetime"),
					"Invalid configuration for Valid Lifetime",
					"valid_lifetime attribute must match the 'value' attribute for DHCP Option 'dhcp-lease-time'.",
				)
			}
		}
	}

	// Preferred lifetime must be less than or equal to valid lifetime
	if !data.PreferredLifetime.IsNull() && !data.PreferredLifetime.IsUnknown() {
		if !data.ValidLifetime.IsNull() && !data.ValidLifetime.IsUnknown() {
			if data.PreferredLifetime.ValueInt64() > data.ValidLifetime.ValueInt64() {
				resp.Diagnostics.AddAttributeError(
					path.Root("preferred_lifetime"),
					"Invalid configuration",
					"The 'preferred_lifetime' must be less than or equal to 'valid_lifetime'.",
				)
			}
		} else if hasDhcpLeaseTime {
			// if valid_lifetime is not set, compare with DHCP lease time
			if dhcpLeaseTimeInt, err := strconv.ParseInt(dhcpLeaseTimeValue, 10, 64); err == nil {
				if data.PreferredLifetime.ValueInt64() > dhcpLeaseTimeInt {
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

func (r *Ipv6fixedaddresstemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("ref"), req.ID)...)
}
