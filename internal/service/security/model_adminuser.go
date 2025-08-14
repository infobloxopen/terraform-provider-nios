package security

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type AdminuserModel struct {
	Ref                             types.String `tfsdk:"ref"`
	AdminGroups                     types.List   `tfsdk:"admin_groups"`
	AuthMethod                      types.String `tfsdk:"auth_method"`
	AuthType                        types.String `tfsdk:"auth_type"`
	CaCertificateIssuer             types.String `tfsdk:"ca_certificate_issuer"`
	ClientCertificateSerialNumber   types.String `tfsdk:"client_certificate_serial_number"`
	Comment                         types.String `tfsdk:"comment"`
	Disable                         types.Bool   `tfsdk:"disable"`
	Email                           types.String `tfsdk:"email"`
	EnableCertificateAuthentication types.Bool   `tfsdk:"enable_certificate_authentication"`
	ExtAttrs                        types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll                     types.Map    `tfsdk:"extattrs_all"`
	Name                            types.String `tfsdk:"name"`
	Password                        types.String `tfsdk:"password"`
	SshKeys                         types.List   `tfsdk:"ssh_keys"`
	Status                          types.String `tfsdk:"status"`
	TimeZone                        types.String `tfsdk:"time_zone"`
	UseSshKeys                      types.Bool   `tfsdk:"use_ssh_keys"`
	UseTimeZone                     types.Bool   `tfsdk:"use_time_zone"`
}

var AdminuserAttrTypes = map[string]attr.Type{
	"ref":                               types.StringType,
	"admin_groups":                      types.ListType{ElemType: types.StringType},
	"auth_method":                       types.StringType,
	"auth_type":                         types.StringType,
	"ca_certificate_issuer":             types.StringType,
	"client_certificate_serial_number":  types.StringType,
	"comment":                           types.StringType,
	"disable":                           types.BoolType,
	"email":                             types.StringType,
	"enable_certificate_authentication": types.BoolType,
	"extattrs":                          types.MapType{ElemType: types.StringType},
	"extattrs_all":                      types.MapType{ElemType: types.StringType},
	"name":                              types.StringType,
	"password":                          types.StringType,
	"ssh_keys":                          types.ListType{ElemType: types.ObjectType{AttrTypes: AdminuserSshKeysAttrTypes}},
	"status":                            types.StringType,
	"time_zone":                         types.StringType,
	"use_ssh_keys":                      types.BoolType,
	"use_time_zone":                     types.BoolType,
}

var AdminuserResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"admin_groups": schema.ListAttribute{
		ElementType:         types.StringType,
		Required:            true,
		Validators:          []validator.List{listvalidator.SizeAtMost(1)},
		MarkdownDescription: "The names of the Admin Groups to which this Admin User belongs. Currently, this is limited to only one Admin Group.",
	},
	"auth_method": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("KEYPAIR", "KEYPAIR_PASSWORD"),
		},
		Default:             stringdefault.StaticString("KEYPAIR"),
		MarkdownDescription: "Determines the way of authentication",
	},
	"auth_type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("LOCAL", "REMOTE", "SAML", "SAML_LOCAL"),
		},
		Default:             stringdefault.StaticString("LOCAL"),
		MarkdownDescription: "The authentication type for the admin user.",
	},
	"ca_certificate_issuer": schema.StringAttribute{
		Optional: true,
		Computed: true,
		//Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The CA certificate that is used for user lookup during authentication.",
	},
	"client_certificate_serial_number": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The serial number of the client certificate.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Name should not have leading or trailing whitespace",
			),
		},
		MarkdownDescription: "Comment for the admin user; maximum 256 characters.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether the admin user is disabled or not. When this is set to False, the admin user is enabled.",
	},
	"email": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The e-mail address for the admin user.",
	},
	"enable_certificate_authentication": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether the user is allowed to log in only with the certificate. Regular username/password authentication will be disabled for this user.",
	},
	"extattrs": schema.MapAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object.",
		ElementType:         types.StringType,
		Default:             mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
	},
	"extattrs_all": schema.MapAttribute{
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object , including default attributes.",
		ElementType:         types.StringType,
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Name should not have leading or trailing whitespace",
			),
		},
		MarkdownDescription: "The name of the admin user.",
	},
	"password": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Name should not have leading or trailing whitespace",
			),
		},
		MarkdownDescription: "The password for the administrator to use when logging in.",
	},
	"ssh_keys": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: AdminuserSshKeysResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "List of ssh keys for a particular user.",
	},
	"status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Status of the user account.",
	},
	"time_zone": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_time_zone")),
		},
		MarkdownDescription: "The time zone for this admin user.",
	},
	"use_ssh_keys": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Enable/disable the ssh keypair authentication.",
	},
	"use_time_zone": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: time_zone",
	},
}

func ExpandAdminuser(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.Adminuser {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AdminuserModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AdminuserModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.Adminuser {
	if m == nil {
		return nil
	}
	to := &security.Adminuser{
		AdminGroups:                     flex.ExpandFrameworkListString(ctx, m.AdminGroups, diags),
		AuthMethod:                      flex.ExpandStringPointer(m.AuthMethod),
		AuthType:                        flex.ExpandStringPointer(m.AuthType),
		CaCertificateIssuer:             flex.ExpandStringPointer(m.CaCertificateIssuer),
		ClientCertificateSerialNumber:   flex.ExpandStringPointer(m.ClientCertificateSerialNumber),
		Comment:                         flex.ExpandStringPointer(m.Comment),
		Disable:                         flex.ExpandBoolPointer(m.Disable),
		Email:                           flex.ExpandStringPointer(m.Email),
		EnableCertificateAuthentication: flex.ExpandBoolPointer(m.EnableCertificateAuthentication),
		ExtAttrs:                        ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		Name:                            flex.ExpandStringPointer(m.Name),
		Password:                        flex.ExpandStringPointer(m.Password),
		SshKeys:                         flex.ExpandFrameworkListNestedBlockNilAsEmpty(ctx, m.SshKeys, diags, ExpandAdminuserSshKeys),
		TimeZone:                        flex.ExpandStringPointer(m.TimeZone),
		UseSshKeys:                      flex.ExpandBoolPointer(m.UseSshKeys),
		UseTimeZone:                     flex.ExpandBoolPointer(m.UseTimeZone),
	}
	return to
}

func FlattenAdminuser(ctx context.Context, from *security.Adminuser, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AdminuserAttrTypes)
	}
	m := AdminuserModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, AdminuserAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AdminuserModel) Flatten(ctx context.Context, from *security.Adminuser, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AdminuserModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.AdminGroups = flex.FlattenFrameworkListString(ctx, from.AdminGroups, diags)
	m.AuthMethod = flex.FlattenStringPointer(from.AuthMethod)
	m.AuthType = flex.FlattenStringPointer(from.AuthType)
	m.CaCertificateIssuer = flex.FlattenStringPointer(from.CaCertificateIssuer)
	m.ClientCertificateSerialNumber = flex.FlattenStringPointer(from.ClientCertificateSerialNumber)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.Email = flex.FlattenStringPointer(from.Email)
	m.EnableCertificateAuthentication = types.BoolPointerValue(from.EnableCertificateAuthentication)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.Name = flex.FlattenStringPointer(from.Name)
	//m.Password = flex.FlattenStringPointer(from.Password)
	m.SshKeys = flex.FlattenFrameworkListNestedBlock(ctx, from.SshKeys, AdminuserSshKeysAttrTypes, diags, FlattenAdminuserSshKeys)
	m.Status = flex.FlattenStringPointer(from.Status)
	m.TimeZone = flex.FlattenStringPointer(from.TimeZone)
	m.UseSshKeys = types.BoolPointerValue(from.UseSshKeys)
	m.UseTimeZone = types.BoolPointerValue(from.UseTimeZone)
}
