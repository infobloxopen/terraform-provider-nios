package dns

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	internaltypes "github.com/infobloxopen/terraform-provider-nios/internal/types"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type IPAssociationModel struct {
	Ref              types.String                  `tfsdk:"ref"`
	Uuid             types.String                  `tfsdk:"uuid"`
	ConfigureForDhcp types.Bool                    `tfsdk:"configure_for_dhcp"`
	Duid             internaltypes.DUIDValue       `tfsdk:"duid"`
	MacAddr          internaltypes.MACAddressValue `tfsdk:"mac"`
	MatchClient      types.String                  `tfsdk:"match_client"`
}

var IPAssociationAttrTypes = map[string]attr.Type{
	"ref":                types.StringType,
	"uuid":               types.StringType,
	"configure_for_dhcp": types.BoolType,
	"duid":               internaltypes.DUIDType{},
	"mac":                internaltypes.MACAddressType{},
	"match_client":       types.StringType,
}

var IpAssociationResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The uuid to the object.",
	},
	"configure_for_dhcp": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Set this to True to enable the DHCP configuration for the IP association.",
		Default:             booldefault.StaticBool(false),
	},
	"duid": schema.StringAttribute{
		CustomType:          internaltypes.DUIDType{},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The DUID of the IP association.",
		Default:             stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.IsValidDUID(),
		},
	},
	"mac": schema.StringAttribute{
		CustomType:          internaltypes.MACAddressType{},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The MAC address of the IP association.",
		Default:             stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.IsValidMacAddress(),
		},
	},
	"match_client": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The match_client value for this IP association. Valid values are: \"DUID\": The host IP address is leased to the matching DUID. \"MAC_ADDRESS\": The host IP address is leased to the matching MAC address.",
		Default:             stringdefault.StaticString("DUID"),
		Validators: []validator.String{
			stringvalidator.OneOf("DUID", "MAC_ADDRESS"),
		},
	},
}
