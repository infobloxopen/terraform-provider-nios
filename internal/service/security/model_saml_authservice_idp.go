package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type SamlAuthserviceIdpModel struct {
	IdpType          types.String `tfsdk:"idp_type"`
	Comment          types.String `tfsdk:"comment"`
	MetadataUrl      types.String `tfsdk:"metadata_url"`
	MetadataToken    types.String `tfsdk:"metadata_token"`
	MetadataFilePath types.String `tfsdk:"metadata_file_path"`
	Groupname        types.String `tfsdk:"groupname"`
	SsoRedirectUrl   types.String `tfsdk:"sso_redirect_url"`
}

var SamlAuthserviceIdpAttrTypes = map[string]attr.Type{
	"idp_type":           types.StringType,
	"comment":            types.StringType,
	"metadata_url":       types.StringType,
	"metadata_token":     types.StringType,
	"metadata_file_path": types.StringType,
	"groupname":          types.StringType,
	"sso_redirect_url":   types.StringType,
}

var SamlAuthserviceIdpResourceSchemaAttributes = map[string]schema.Attribute{
	"idp_type": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.OneOf(
				"AZURE_SSO",
				"OKTA",
				"OTHER",
				"PING_IDENTITY",
				"SHIBBOLETH_SSO",
			),
		},
		MarkdownDescription: "SAML Identity Provider type.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The SAML Identity Provider descriptive comment.",
	},
	"metadata_url": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Identity Provider Metadata URL.",
	},
	"metadata_token": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The token returned by the uploadinit function call in object fileop.",
	},
	"metadata_file_path": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The file path to the IdP SAML metadata XML file.",
	},
	"groupname": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The SAML groupname optional user group attribute.",
	},
	"sso_redirect_url": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.IsValidIPv4OrFQDN(),
		},
		MarkdownDescription: "host name or IP address of the GM",
	},
}

func ExpandSamlAuthserviceIdp(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.SamlAuthserviceIdp {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m SamlAuthserviceIdpModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *SamlAuthserviceIdpModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.SamlAuthserviceIdp {
	if m == nil {
		return nil
	}
	to := &security.SamlAuthserviceIdp{
		IdpType:        flex.ExpandStringPointer(m.IdpType),
		Comment:        flex.ExpandStringPointer(m.Comment),
		MetadataUrl:    flex.ExpandStringPointer(m.MetadataUrl),
		MetadataToken:  flex.ExpandStringPointer(m.MetadataToken),
		Groupname:      flex.ExpandStringPointer(m.Groupname),
		SsoRedirectUrl: flex.ExpandStringPointer(m.SsoRedirectUrl),
	}
	return to
}

func FlattenSamlAuthserviceIdp(ctx context.Context, from *security.SamlAuthserviceIdp, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SamlAuthserviceIdpAttrTypes)
	}
	m := SamlAuthserviceIdpModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, SamlAuthserviceIdpAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *SamlAuthserviceIdpModel) Flatten(ctx context.Context, from *security.SamlAuthserviceIdp, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = SamlAuthserviceIdpModel{}
	}
	m.IdpType = flex.FlattenStringPointer(from.IdpType)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.MetadataUrl = flex.FlattenStringPointer(from.MetadataUrl)
	m.MetadataToken = flex.FlattenStringPointer(from.MetadataToken)
	m.Groupname = flex.FlattenStringPointer(from.Groupname)
	m.SsoRedirectUrl = flex.FlattenStringPointer(from.SsoRedirectUrl)
}
