package security

import (
	"context"

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

	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type SnmpuserModel struct {
	Ref                    types.String `tfsdk:"ref"`
	AuthenticationPassword types.String `tfsdk:"authentication_password"`
	AuthenticationProtocol types.String `tfsdk:"authentication_protocol"`
	Comment                types.String `tfsdk:"comment"`
	Disable                types.Bool   `tfsdk:"disable"`
	ExtAttrs               types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll            types.Map    `tfsdk:"extattrs_all"`
	Name                   types.String `tfsdk:"name"`
	PrivacyPassword        types.String `tfsdk:"privacy_password"`
	PrivacyProtocol        types.String `tfsdk:"privacy_protocol"`
}

var SnmpuserAttrTypes = map[string]attr.Type{
	"ref":                     types.StringType,
	"authentication_password": types.StringType,
	"authentication_protocol": types.StringType,
	"comment":                 types.StringType,
	"disable":                 types.BoolType,
	"extattrs":                types.MapType{ElemType: types.StringType},
	"extattrs_all":            types.MapType{ElemType: types.StringType},
	"name":                    types.StringType,
	"privacy_password":        types.StringType,
	"privacy_protocol":        types.StringType,
}

var SnmpuserResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"authentication_password": schema.StringAttribute{
		Optional:  true,
		Sensitive: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("authentication_protocol")),
			stringvalidator.LengthBetween(8, 256),
		},
		MarkdownDescription: "Determines an authentication password for the user. This is a write-only attribute. Must be between 8 and 256 characters.",
	},
	"authentication_protocol": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.OneOf("MD5", "SHA", "NONE"),
		},
		MarkdownDescription: "The authentication protocol to be used for this user.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "A descriptive comment for the SNMPv3 User.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if SNMPv3 user is disabled or not.",
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
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object , including default attributes.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of the user.",
	},
	"privacy_password": schema.StringAttribute{
		Optional:  true,
		Sensitive: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("privacy_protocol")),
			stringvalidator.LengthBetween(8, 256),
		},
		MarkdownDescription: "Determines a password for the privacy protocol.",
	},
	"privacy_protocol": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.OneOf("DES", "AES", "NONE"),
		},
		MarkdownDescription: "The privacy protocol to be used for this user.",
	},
}

func (m *SnmpuserModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.Snmpuser {
	if m == nil {
		return nil
	}
	to := &security.Snmpuser{
		AuthenticationPassword: flex.ExpandStringPointer(m.AuthenticationPassword),
		AuthenticationProtocol: flex.ExpandStringPointer(m.AuthenticationProtocol),
		Comment:                flex.ExpandStringPointer(m.Comment),
		Disable:                flex.ExpandBoolPointer(m.Disable),
		ExtAttrs:               ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		Name:                   flex.ExpandStringPointer(m.Name),
		PrivacyPassword:        flex.ExpandStringPointer(m.PrivacyPassword),
		PrivacyProtocol:        flex.ExpandStringPointer(m.PrivacyProtocol),
	}
	return to
}

func FlattenSnmpuser(ctx context.Context, from *security.Snmpuser, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SnmpuserAttrTypes)
	}
	m := SnmpuserModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, SnmpuserAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *SnmpuserModel) Flatten(ctx context.Context, from *security.Snmpuser, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = SnmpuserModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.AuthenticationProtocol = flex.FlattenStringPointer(from.AuthenticationProtocol)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.PrivacyProtocol = flex.FlattenStringPointer(from.PrivacyProtocol)
}
