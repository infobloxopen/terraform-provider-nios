package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type CertificateAuthserviceModel struct {
	Ref                   types.String `tfsdk:"ref"`
	AutoPopulateLogin     types.String `tfsdk:"auto_populate_login"`
	CaCertificates        types.List   `tfsdk:"ca_certificates"`
	Comment               types.String `tfsdk:"comment"`
	Disabled              types.Bool   `tfsdk:"disabled"`
	EnablePasswordRequest types.Bool   `tfsdk:"enable_password_request"`
	EnableRemoteLookup    types.Bool   `tfsdk:"enable_remote_lookup"`
	MaxRetries            types.Int64  `tfsdk:"max_retries"`
	Name                  types.String `tfsdk:"name"`
	OcspCheck             types.String `tfsdk:"ocsp_check"`
	OcspResponders        types.List   `tfsdk:"ocsp_responders"`
	RecoveryInterval      types.Int64  `tfsdk:"recovery_interval"`
	RemoteLookupPassword  types.String `tfsdk:"remote_lookup_password"`
	RemoteLookupService   types.String `tfsdk:"remote_lookup_service"`
	RemoteLookupUsername  types.String `tfsdk:"remote_lookup_username"`
	ResponseTimeout       types.Int64  `tfsdk:"response_timeout"`
	TrustModel            types.String `tfsdk:"trust_model"`
	UserMatchType         types.String `tfsdk:"user_match_type"`
}

var CertificateAuthserviceAttrTypes = map[string]attr.Type{
	"ref":                     types.StringType,
	"auto_populate_login":     types.StringType,
	"ca_certificates":         types.ListType{ElemType: types.StringType},
	"comment":                 types.StringType,
	"disabled":                types.BoolType,
	"enable_password_request": types.BoolType,
	"enable_remote_lookup":    types.BoolType,
	"max_retries":             types.Int64Type,
	"name":                    types.StringType,
	"ocsp_check":              types.StringType,
	"ocsp_responders":         types.ListType{ElemType: types.ObjectType{AttrTypes: CertificateAuthserviceOcspRespondersAttrTypes}},
	"recovery_interval":       types.Int64Type,
	"remote_lookup_password":  types.StringType,
	"remote_lookup_service":   types.StringType,
	"remote_lookup_username":  types.StringType,
	"response_timeout":        types.Int64Type,
	"trust_model":             types.StringType,
	"user_match_type":         types.StringType,
}

var CertificateAuthserviceResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"auto_populate_login": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("S_DN_CN"),
		Validators: []validator.String{
			stringvalidator.OneOf("AD_SUBJECT_ISSUER", "SAN_EMAIL", "SAN_UPN", "SERIAL_NUMBER", "S_DN_CN", "S_DN_EMAIL"),
		},
		MarkdownDescription: "Specifies the value of the client certificate for automatically populating the NIOS login name.",
	},
	"ca_certificates": schema.ListAttribute{
		ElementType:         types.StringType,
		Required:            true,
		MarkdownDescription: "The list of CA certificates.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The descriptive comment for the certificate authentication service.",
	},
	"disabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if this certificate authentication service is enabled or disabled.",
	},
	"enable_password_request": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Determines if username/password authentication together with client certificate authentication is enabled or disabled.",
	},
	"enable_remote_lookup": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the lookup for user group membership information on remote services is enabled or disabled.",
	},
	"max_retries": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(0),
		MarkdownDescription: "The number of validation attempts before the appliance contacts the next responder.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of the certificate authentication service.",
	},
	"ocsp_check": schema.StringAttribute{
		Optional: true,
		Computed: true,
		//Default:  stringdefault.StaticString("MANUAL"),
		Validators: []validator.String{
			stringvalidator.OneOf("AIA_AND_MANUAL", "AIA_ONLY", "DISABLED", "MANUAL"),
		},
		MarkdownDescription: "Specifies the source of OCSP settings.",
	},
	"ocsp_responders": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: CertificateAuthserviceOcspRespondersResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "An ordered list of OCSP responders that are part of the certificate authentication service.",
	},
	"recovery_interval": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(30),
		MarkdownDescription: "The period of time the appliance waits before it attempts to contact a responder that is out of service again. The value must be between 1 and 600 seconds.",
	},
	"remote_lookup_password": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The password for the service account.",
	},
	"remote_lookup_service": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The service that will be used for remote lookup.",
	},
	"remote_lookup_username": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The username for the service account.",
	},
	"response_timeout": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(1000),
		MarkdownDescription: "The validation timeout period in milliseconds.",
	},
	"trust_model": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("DIRECT"),
		Validators: []validator.String{
			stringvalidator.OneOf("DELEGATED", "DIRECT"),
		},
		MarkdownDescription: "The OCSP trust model.",
	},
	"user_match_type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("AUTO_MATCH"),
		Validators: []validator.String{
			stringvalidator.OneOf("AUTO_MATCH", "DIRECT_MATCH"),
		},
		MarkdownDescription: "Specifies how to search for a user.",
	},
}

func (m *CertificateAuthserviceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.CertificateAuthservice {
	if m == nil {
		return nil
	}
	to := &security.CertificateAuthservice{
		AutoPopulateLogin:     flex.ExpandStringPointer(m.AutoPopulateLogin),
		CaCertificates:        flex.ExpandFrameworkListString(ctx, m.CaCertificates, diags),
		Comment:               flex.ExpandStringPointer(m.Comment),
		Disabled:              flex.ExpandBoolPointer(m.Disabled),
		EnablePasswordRequest: flex.ExpandBoolPointer(m.EnablePasswordRequest),
		EnableRemoteLookup:    flex.ExpandBoolPointer(m.EnableRemoteLookup),
		MaxRetries:            flex.ExpandInt64Pointer(m.MaxRetries),
		Name:                  flex.ExpandStringPointer(m.Name),
		OcspCheck:             flex.ExpandStringPointer(m.OcspCheck),
		OcspResponders:        flex.ExpandFrameworkListNestedBlock(ctx, m.OcspResponders, diags, ExpandCertificateAuthserviceOcspResponders),
		RecoveryInterval:      flex.ExpandInt64Pointer(m.RecoveryInterval),
		RemoteLookupPassword:  flex.ExpandStringPointer(m.RemoteLookupPassword),
		RemoteLookupService:   ExpandCertificateAuthserviceRemoteLookupService(ctx, m.RemoteLookupService, diags),
		RemoteLookupUsername:  flex.ExpandStringPointer(m.RemoteLookupUsername),
		ResponseTimeout:       flex.ExpandInt64Pointer(m.ResponseTimeout),
		TrustModel:            flex.ExpandStringPointer(m.TrustModel),
		UserMatchType:         flex.ExpandStringPointer(m.UserMatchType),
	}
	return to
}

func FlattenCertificateAuthservice(ctx context.Context, from *security.CertificateAuthservice, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(CertificateAuthserviceAttrTypes)
	}
	m := CertificateAuthserviceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, CertificateAuthserviceAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *CertificateAuthserviceModel) Flatten(ctx context.Context, from *security.CertificateAuthservice, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = CertificateAuthserviceModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.AutoPopulateLogin = flex.FlattenStringPointer(from.AutoPopulateLogin)
	m.CaCertificates = flex.FlattenFrameworkListString(ctx, from.CaCertificates, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disabled = types.BoolPointerValue(from.Disabled)
	m.EnablePasswordRequest = types.BoolPointerValue(from.EnablePasswordRequest)
	m.EnableRemoteLookup = types.BoolPointerValue(from.EnableRemoteLookup)
	m.MaxRetries = flex.FlattenInt64Pointer(from.MaxRetries)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.OcspCheck = flex.FlattenStringPointer(from.OcspCheck)
	m.OcspResponders = flex.FlattenFrameworkListNestedBlock(ctx, from.OcspResponders, CertificateAuthserviceOcspRespondersAttrTypes, diags, FlattenCertificateAuthserviceOcspResponders)
	m.RecoveryInterval = flex.FlattenInt64Pointer(from.RecoveryInterval)
	m.RemoteLookupPassword = flex.FlattenStringPointer(from.RemoteLookupPassword)
	m.RemoteLookupService = FlattenCertificateAuthserviceRemoteLookupService(ctx, from.RemoteLookupService, diags)
	m.RemoteLookupUsername = flex.FlattenStringPointer(from.RemoteLookupUsername)
	m.ResponseTimeout = flex.FlattenInt64Pointer(from.ResponseTimeout)
	m.TrustModel = flex.FlattenStringPointer(from.TrustModel)
	m.UserMatchType = flex.FlattenStringPointer(from.UserMatchType)
}
