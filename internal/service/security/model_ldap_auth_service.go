package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type LdapAuthServiceModel struct {
	Ref                         types.String `tfsdk:"ref"`
    Uuid        types.String `tfsdk:"uuid"`
	Comment                     types.String `tfsdk:"comment"`
	Disable                     types.Bool   `tfsdk:"disable"`
	EaMapping                   types.List   `tfsdk:"ea_mapping"`
	LdapGroupAttribute          types.String `tfsdk:"ldap_group_attribute"`
	LdapGroupAuthenticationType types.String `tfsdk:"ldap_group_authentication_type"`
	LdapUserAttribute           types.String `tfsdk:"ldap_user_attribute"`
	Mode                        types.String `tfsdk:"mode"`
	Name                        types.String `tfsdk:"name"`
	RecoveryInterval            types.Int64  `tfsdk:"recovery_interval"`
	Retries                     types.Int64  `tfsdk:"retries"`
	SearchScope                 types.String `tfsdk:"search_scope"`
	Servers                     types.List   `tfsdk:"servers"`
	Timeout                     types.Int64  `tfsdk:"timeout"`
}

var LdapAuthServiceAttrTypes = map[string]attr.Type{
	"ref":                            types.StringType,
    "uuid":        types.StringType,
	"comment":                        types.StringType,
	"disable":                        types.BoolType,
	"ea_mapping":                     types.ListType{ElemType: types.ObjectType{AttrTypes: LdapAuthServiceEaMappingAttrTypes}},
	"ldap_group_attribute":           types.StringType,
	"ldap_group_authentication_type": types.StringType,
	"ldap_user_attribute":            types.StringType,
	"mode":                           types.StringType,
	"name":                           types.StringType,
	"recovery_interval":              types.Int64Type,
	"retries":                        types.Int64Type,
	"search_scope":                   types.StringType,
	"servers":                        types.ListType{ElemType: types.ObjectType{AttrTypes: LdapAuthServiceServersAttrTypes}},
	"timeout":                        types.Int64Type,
}

var LdapAuthServiceResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
    "uuid": schema.StringAttribute{
        Computed:            true,
        MarkdownDescription: "The uuid to the object.",
    },
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The LDAP descriptive comment.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the LDAP authentication service is disabled.",
	},
	"ea_mapping": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: LdapAuthServiceEaMappingResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The mapping LDAP fields to extensible attributes.",
	},
	"ldap_group_attribute": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the LDAP attribute that defines group membership.",
	},
	"ldap_group_authentication_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The LDAP group authentication type.",
	},
	"ldap_user_attribute": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The LDAP userid attribute that is used for search.",
	},
	"mode": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The LDAP authentication mode.",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The LDAP authentication service name.",
	},
	"recovery_interval": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The period of time in seconds to wait before trying to contact a LDAP server that has been marked as 'DOWN'.",
	},
	"retries": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The maximum number of LDAP authentication attempts.",
	},
	"search_scope": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The starting point of the LDAP search.",
	},
	"servers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: LdapAuthServiceServersResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of LDAP servers used for authentication.",
	},
	"timeout": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The LDAP authentication timeout in seconds.",
	},
}

func ExpandLdapAuthService(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.LdapAuthService {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m LdapAuthServiceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *LdapAuthServiceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.LdapAuthService {
	if m == nil {
		return nil
	}
	to := &security.LdapAuthService{
		Ref:                         flex.ExpandStringPointer(m.Ref),
		Comment:                     flex.ExpandStringPointer(m.Comment),
		Disable:                     flex.ExpandBoolPointer(m.Disable),
		EaMapping:                   flex.ExpandFrameworkListNestedBlock(ctx, m.EaMapping, diags, ExpandLdapAuthServiceEaMapping),
		LdapGroupAttribute:          flex.ExpandStringPointer(m.LdapGroupAttribute),
		LdapGroupAuthenticationType: flex.ExpandStringPointer(m.LdapGroupAuthenticationType),
		LdapUserAttribute:           flex.ExpandStringPointer(m.LdapUserAttribute),
		Mode:                        flex.ExpandStringPointer(m.Mode),
		Name:                        flex.ExpandStringPointer(m.Name),
		RecoveryInterval:            flex.ExpandInt64Pointer(m.RecoveryInterval),
		Retries:                     flex.ExpandInt64Pointer(m.Retries),
		SearchScope:                 flex.ExpandStringPointer(m.SearchScope),
		Servers:                     flex.ExpandFrameworkListNestedBlock(ctx, m.Servers, diags, ExpandLdapAuthServiceServers),
		Timeout:                     flex.ExpandInt64Pointer(m.Timeout),
	}
	return to
}

func FlattenLdapAuthService(ctx context.Context, from *security.LdapAuthService, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(LdapAuthServiceAttrTypes)
	}
	m := LdapAuthServiceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, LdapAuthServiceAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *LdapAuthServiceModel) Flatten(ctx context.Context, from *security.LdapAuthService, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = LdapAuthServiceModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
    m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.EaMapping = flex.FlattenFrameworkListNestedBlock(ctx, from.EaMapping, LdapAuthServiceEaMappingAttrTypes, diags, FlattenLdapAuthServiceEaMapping)
	m.LdapGroupAttribute = flex.FlattenStringPointer(from.LdapGroupAttribute)
	m.LdapGroupAuthenticationType = flex.FlattenStringPointer(from.LdapGroupAuthenticationType)
	m.LdapUserAttribute = flex.FlattenStringPointer(from.LdapUserAttribute)
	m.Mode = flex.FlattenStringPointer(from.Mode)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.RecoveryInterval = flex.FlattenInt64Pointer(from.RecoveryInterval)
	m.Retries = flex.FlattenInt64Pointer(from.Retries)
	m.SearchScope = flex.FlattenStringPointer(from.SearchScope)
	m.Servers = flex.FlattenFrameworkListNestedBlock(ctx, from.Servers, LdapAuthServiceServersAttrTypes, diags, FlattenLdapAuthServiceServers)
	m.Timeout = flex.FlattenInt64Pointer(from.Timeout)
}
