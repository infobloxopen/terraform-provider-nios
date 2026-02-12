package rir

import (
	"context"
	"fmt"
	"net/http"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"

	"github.com/infobloxopen/terraform-provider-nios/internal/config"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForRirOrganization = "extattrs,id,maintainer,name,rir,sender_email"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &RirOrganizationResource{}
var _ resource.ResourceWithImportState = &RirOrganizationResource{}

func NewRirOrganizationResource() resource.Resource {
	return &RirOrganizationResource{}
}

// RirOrganizationResource defines the resource implementation.
type RirOrganizationResource struct {
	client *niosclient.APIClient
}

func (r *RirOrganizationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "rir_organization"
}

func (r *RirOrganizationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a RirOrganization resource object.",
		Attributes:          RirOrganizationResourceSchemaAttributes,
	}
}

func (r *RirOrganizationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *RirOrganizationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data RirOrganizationModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.RIRAPI.
		RirOrganizationAPI.
		Create(ctx).
		RirOrganization(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForRirOrganization).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create RirOrganization, got error: %s", err))
		return
	}

	res := apiRes.CreateRirOrganizationResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RirOrganizationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data RirOrganizationModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.RIRAPI.
		RirOrganizationAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForRirOrganization).
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read RirOrganization, got error: %s", err))
		return
	}

	res := apiRes.GetRirOrganizationResponseObjectAsResult.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RirOrganizationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data RirOrganizationModel

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

	apiRes, _, err := r.client.RIRAPI.
		RirOrganizationAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		RirOrganization(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForRirOrganization).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update RirOrganization, got error: %s", err))
		return
	}

	res := apiRes.UpdateRirOrganizationResponseAsObject.GetResult()

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RirOrganizationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data RirOrganizationModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.RIRAPI.
		RirOrganizationAPI.
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete RirOrganization, got error: %s", err))
		return
	}
}

func (r *RirOrganizationResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data RirOrganizationModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if !data.Rir.IsUnknown() && !data.Rir.IsNull() && data.Rir.ValueString() == "RIPE" {
		var extattrsMap map[string]string
		resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("extattrs"), &extattrsMap)...)
		if resp.Diagnostics.HasError() {
			return
		}

		for key, value := range extattrsMap {
			if key == "RIPE Country" {
				if value == "" {
					resp.Diagnostics.AddError("Invalid RIPE Country", "RIPE Country cannot be empty.")
					continue
				}
				eaRipeCountryList := []string{
					"Afghanistan (AF)", "Åland Islands (AX)", "Albania (AL)", "Algeria (DZ)", "American Samoa (AS)", "Andorra (AD)", "Angola (AO)", "Anguilla (AI)", "Antarctica (AQ)", "Antigua and Barbuda (AG)",
					"Argentina (AR)", "Armenia (AM)", "Aruba (AW)", "Australia (AU)", "Austria (AT)", "Azerbaijan (AZ)", "Bahamas (BS)", "Bahrain (BH)", "Bangladesh (BD)", "Barbados (BB)",
					"Belarus (BY)", "Belgium (BE)", "Belize (BZ)", "Benin (BJ)", "Bermuda (BM)", "Bhutan (BT)", "Bolivia, Plurinational State of (BO)", "Bonaire, Sint Eustatius and Saba (BQ)", "Bosnia and Herzegovina (BA)", "Botswana (BW)",
					"Bouvet Island (BV)", "Brazil (BR)", "British Indian Ocean Territory (IO)", "Brunei Darussalam (BN)", "Bulgaria (BG)", "Burkina Faso (BF)", "Burundi (BI)", "Cambodia (KH)", "Cameroon (CM)", "Canada (CA)",
					"Cape Verde (CV)", "Cayman Islands (KY)", "Central African Republic (CF)", "Chad (TD)", "Chile (CL)", "China (CN)", "Christmas Island (CX)", "Cocos (Keeling) Islands (CC)", "Colombia (CO)", "Comoros (KM)",
					"Congo (CG)", "Congo, The Democratic Republic of the (CD)", "Cook Islands (CK)", "Costa Rica (CR)", "Côte d'Ivoire (CI)", "Croatia (HR)", "Cuba (CU)", "Curaçao (CW)", "Cyprus (CY)", "Czech Republic (CZ)",
					"Denmark (DK)", "Djibouti (DJ)", "Dominica (DM)", "Dominican Republic (DO)", "Ecuador (EC)", "Egypt (EG)", "El Salvador (SV)", "Equatorial Guinea (GQ)", "Eritrea (ER)", "Estonia (EE)",
					"Ethiopia (ET)", "Falkland Islands (Malvinas) (FK)", "Faroe Islands (FO)", "Fiji (FJ)", "Finland (FI)", "France (FR)", "French Guiana (GF)", "French Polynesia (PF)", "French Southern Territories (TF)", "Gabon (GA)",
					"Gambia (GM)", "Georgia (GE)", "Germany (DE)", "Ghana (GH)", "Gibraltar (GI)", "Greece (GR)", "Greenland (GL)", "Grenada (GD)", "Guadeloupe (GP)", "Guam (GU)",
					"Guatemala (GT)", "Guernsey (GG)", "Guinea (GN)", "Guinea-Bissau (GW)", "Guyana (GY)", "Haiti (HT)", "Heard Island and McDonald Islands (HM)", "Holy See (Vatican City State) (VA)", "Honduras (HN)", "Hong Kong (HK)",
					"Hungary (HU)", "Iceland (IS)", "India (IN)", "Indonesia (ID)", "Iran, Islamic Republic of (IR)", "Iraq (IQ)", "Ireland (IE)", "Isle of Man (IM)", "Israel (IL)", "Italy (IT)",
					"Jamaica (JM)", "Japan (JP)", "Jersey (JE)", "Jordan (JO)", "Kazakhstan (KZ)", "Kenya (KE)", "Kiribati (KI)", "Korea, Democratic People's Republic of (KP)", "Korea, Republic of (KR)", "Kuwait (KW)",
					"Kyrgyzstan (KG)", "Lao People's Democratic Republic (LA)", "Latvia (LV)", "Lebanon (LB)", "Lesotho (LS)", "Liberia (LR)", "Libya (LY)", "Liechtenstein (LI)", "Lithuania (LT)", "Luxembourg (LU)",
					"Macao (MO)", "Macedonia, The Former Yugoslav Republic of (MK)", "Madagascar (MG)", "Malawi (MW)", "Malaysia (MY)", "Maldives (MV)", "Mali (ML)", "Malta (MT)", "Marshall Islands (MH)", "Martinique (MQ)",
					"Mauritania (MR)", "Mauritius (MU)", "Mayotte (YT)", "Mexico (MX)", "Micronesia, Federated States of (FM)", "Moldova, Republic of (MD)", "Monaco (MC)", "Mongolia (MN)", "Montenegro (ME)", "Montserrat (MS)",
					"Morocco (MA)", "Mozambique (MZ)", "Myanmar (MM)", "Namibia (NA)", "Nauru (NR)", "Nepal (NP)", "Netherlands (NL)", "New Caledonia (NC)", "New Zealand (NZ)", "Nicaragua (NI)",
					"Niger (NE)", "Nigeria (NG)", "Niue (NU)", "Norfolk Island (NF)", "Northern Mariana Islands (MP)", "Norway (NO)", "Oman (OM)", "Pakistan (PK)", "Palau (PW)", "Palestinian Territory, Occupied (PS)",
					"Panama (PA)", "Papua New Guinea (PG)", "Paraguay (PY)", "Peru (PE)", "Philippines (PH)", "Pitcairn (PN)", "Poland (PL)", "Portugal (PT)", "Puerto Rico (PR)", "Qatar (QA)",
					"Réunion (RE)", "Romania (RO)", "Russian Federation (RU)", "Rwanda (RW)", "Saint Barthélemy (BL)", "Saint Helena, Ascension and Tristan da Cunha (SH)", "Saint Kitts and Nevis (KN)", "Saint Lucia (LC)", "Saint Martin (French part) (MF)", "Saint Pierre and Miquelon (PM)",
					"Saint Vincent and the Grenadines (VC)", "Samoa (WS)", "San Marino (SM)", "Sao Tome and Principe (ST)", "Saudi Arabia (SA)", "Senegal (SN)", "Serbia (RS)", "Seychelles (SC)", "Sierra Leone (SL)", "Singapore (SG)",
					"Sint Maarten (Dutch part) (SX)", "Slovakia (SK)", "Slovenia (SI)", "Solomon Islands (SB)", "Somalia (SO)", "South Africa (ZA)", "South Georgia and the South Sandwich Islands (GS)", "South Sudan (SS)", "Spain (ES)", "Sri Lanka (LK)",
					"Sudan (SD)", "Suriname (SR)", "Svalbard and Jan Mayen (SJ)", "Swaziland (SZ)", "Sweden (SE)", "Switzerland (CH)", "Syrian Arab Republic (SY)", "Taiwan, Province of China (TW)", "Tajikistan (TJ)", "Tanzania, United Republic of (TZ)",
					"Thailand (TH)", "Timor-Leste (TL)", "Togo (TG)", "Tokelau (TK)", "Tonga (TO)", "Trinidad and Tobago (TT)", "Tunisia (TN)", "Turkey (TR)", "Turkmenistan (TM)", "Turks and Caicos Islands (TC)",
					"Tuvalu (TV)", "Uganda (UG)", "Ukraine (UA)", "United Arab Emirates (AE)", "United Kingdom (GB)", "United States (US)", "United States Minor Outlying Islands (UM)", "Uruguay (UY)", "Uzbekistan (UZ)", "Vanuatu (VU)",
					"Venezuela, Bolivarian Republic of (VE)", "Viet Nam (VN)", "Virgin Islands, British (VG)", "Virgin Islands, U.S. (VI)", "Wallis and Futuna (WF)", "Western Sahara (EH)", "Yemen (YE)", "Zambia (ZM)", "Zimbabwe (ZW)",
				}
				found := false
				for _, country := range eaRipeCountryList {
					if value == country {
						found = true
						break
					}
				}
				if !found {
					resp.Diagnostics.AddError("Invalid RIPE Country", fmt.Sprintf("RIPE Country '%s' is not a valid option. Valid Values are %q", value, eaRipeCountryList))
				}
			} else if key == "RIPE Email" {
				if matched, _ := regexp.MatchString(`^[^@]+@[^@]+\.com$`, value); !matched {
					resp.Diagnostics.AddError("Invalid RIPE Email", fmt.Sprintf("RIPE Email '%s' is not a valid .com email address.", value))
				}
			} else if key == "RIPE Technical Contact" {
				if matched, _ := regexp.MatchString(`^[A-Za-z]{2,4}(?:[1-9][0-9]{0,5})?-[A-Za-z0-9]{1,9}$`, value); !matched {
					resp.Diagnostics.AddError("Invalid RIPE Technical Contact", fmt.Sprintf("RIPE Technical Contact '%s' is not a valid value. Valid format is 'AB123-XYZ'", value))
				}
			} else if key == "RIPE Organization Type" {
				eaRipeOrgTypeList := []string{
					"IANA", "RIR", "NIR", "LIR", "WHITEPAGES",
					"DIRECT_ASSIGNMENT", "OTHER",
				}
				if value != "" {
					found := false
					for _, orgType := range eaRipeOrgTypeList {
						if value == orgType {
							found = true
							break
						}
					}
					if !found {
						resp.Diagnostics.AddError("Invalid RIPE Organization Type", fmt.Sprintf("RIPE Organization Type '%s' is not a valid option. Valid Values are %q", value, eaRipeOrgTypeList))
					}
				}
			}
		}
	}
}

func (r *RirOrganizationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("ref"), req, resp)
}
