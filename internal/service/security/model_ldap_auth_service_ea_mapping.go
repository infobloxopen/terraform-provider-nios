package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type LdapAuthServiceEaMappingModel struct {
	Name     types.String `tfsdk:"name"`
	MappedEa types.String `tfsdk:"mapped_ea"`
}

var LdapAuthServiceEaMappingAttrTypes = map[string]attr.Type{
	"name":      types.StringType,
	"mapped_ea": types.StringType,
}

var LdapAuthServiceEaMappingResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The LDAP attribute name.",
	},
	"mapped_ea": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the extensible attribute definition object to which the LDAP attribute is mapped.",
	},
}

func ExpandLdapAuthServiceEaMapping(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.LdapAuthServiceEaMapping {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m LdapAuthServiceEaMappingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *LdapAuthServiceEaMappingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.LdapAuthServiceEaMapping {
	if m == nil {
		return nil
	}
	to := &security.LdapAuthServiceEaMapping{
		Name:     flex.ExpandStringPointer(m.Name),
		MappedEa: flex.ExpandStringPointer(m.MappedEa),
	}
	return to
}

func FlattenLdapAuthServiceEaMapping(ctx context.Context, from *security.LdapAuthServiceEaMapping, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(LdapAuthServiceEaMappingAttrTypes)
	}
	m := LdapAuthServiceEaMappingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, LdapAuthServiceEaMappingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *LdapAuthServiceEaMappingModel) Flatten(ctx context.Context, from *security.LdapAuthServiceEaMapping, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = LdapAuthServiceEaMappingModel{}
	}
	m.Name = flex.FlattenStringPointer(from.Name)
	m.MappedEa = flex.FlattenStringPointer(from.MappedEa)
}
