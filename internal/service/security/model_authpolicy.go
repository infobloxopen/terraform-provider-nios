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

type AuthpolicyModel struct {
	Ref          types.String `tfsdk:"ref"`
	Uuid         types.String `tfsdk:"uuid"`
	AdminGroups  types.List   `tfsdk:"admin_groups"`
	AuthServices types.List   `tfsdk:"auth_services"`
	DefaultGroup types.String `tfsdk:"default_group"`
	UsageType    types.String `tfsdk:"usage_type"`
}

var AuthpolicyAttrTypes = map[string]attr.Type{
	"ref":           types.StringType,
	"uuid":          types.StringType,
	"admin_groups":  types.ListType{ElemType: types.StringType},
	"auth_services": types.ListType{ElemType: types.StringType},
	"default_group": types.StringType,
	"usage_type":    types.StringType,
}

var AuthpolicyResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The uuid to the object.",
	},
	"admin_groups": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "List of names of local administration groups that are mapped to remote administration groups.",
	},
	"auth_services": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The array that contains an ordered list of refs to :doc:`localuser:authservice object </objects/localuser.authservice>`, ldap_auth_service object ldap_auth_service, :doc:`radius:authservice object </objects/radius.authservice>`, :doc:`tacacsplus:authservice object </objects/tacacsplus.authservice>`, ad_auth_service object ad_auth_service, :doc:`certificate:authservice object </objects/certificate.authservice>`. :doc:`saml:authservice object </objects/saml.authservice>`,",
	},
	"default_group": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The default admin group that provides authentication in case no valid group is found.",
	},
	"usage_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Remote policies usage.",
	},
}

func ExpandAuthpolicy(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.Authpolicy {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AuthpolicyModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AuthpolicyModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.Authpolicy {
	if m == nil {
		return nil
	}
	to := &security.Authpolicy{
		Ref:          flex.ExpandStringPointer(m.Ref),
		Uuid:         flex.ExpandStringPointer(m.Uuid),
		AdminGroups:  flex.ExpandFrameworkListString(ctx, m.AdminGroups, diags),
		AuthServices: flex.ExpandFrameworkListString(ctx, m.AuthServices, diags),
		DefaultGroup: flex.ExpandStringPointer(m.DefaultGroup),
		UsageType:    flex.ExpandStringPointer(m.UsageType),
	}
	return to
}

func FlattenAuthpolicy(ctx context.Context, from *security.Authpolicy, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AuthpolicyAttrTypes)
	}
	m := AuthpolicyModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AuthpolicyAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AuthpolicyModel) Flatten(ctx context.Context, from *security.Authpolicy, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AuthpolicyModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.AdminGroups = flex.FlattenFrameworkListString(ctx, from.AdminGroups, diags)
	m.AuthServices = flex.FlattenFrameworkListString(ctx, from.AuthServices, diags)
	m.DefaultGroup = flex.FlattenStringPointer(from.DefaultGroup)
	m.UsageType = flex.FlattenStringPointer(from.UsageType)
}
