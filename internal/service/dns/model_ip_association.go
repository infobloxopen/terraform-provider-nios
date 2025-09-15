package dns

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type IPAssociationModel struct {
	Ref              types.String `tfsdk:"ref"`
	ConfigureForDhcp types.Bool   `tfsdk:"configure_for_dhcp"`
	Duid             types.String `tfsdk:"duid"`
	InternalID       types.String `tfsdk:"internal_id"`
	MacAddr          types.String `tfsdk:"mac"`
}

var IPAssociationAttrTypes = map[string]attr.Type{
	"ref":                types.StringType,
	"configure_for_dhcp": types.BoolType,
	"duid":               types.StringType,
	"internal_id":        types.StringType,
	"mac":                types.StringType,
}

var IpAssociationResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"configure_for_dhcp": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Set this to True to enable the DHCP configuration for the IP association.",
		Default:             booldefault.StaticBool(true),
	},
	"duid": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The DUID of the IP association.",
		Default:             stringdefault.StaticString(""),
	},
	"internal_id": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Internal ID of the IP association.",
	},
	"mac": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The MAC address of the IP association.",
		Default:             stringdefault.StaticString(""),
	},
}
