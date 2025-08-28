package dns

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

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type ViewLastQueriedAclModel struct {
	Address    types.String `tfsdk:"address"`
	Permission types.String `tfsdk:"permission"`
}

var ViewLastQueriedAclAttrTypes = map[string]attr.Type{
	"address":    types.StringType,
	"permission": types.StringType,
}

var ViewLastQueriedAclResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateNoLeadingOrTrailingWhitespace(),
		},
		MarkdownDescription: "The address this rule applies to or \"Any\".",
	},
	"permission": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("ALLOW", "DENY"),
		},
		Default:             stringdefault.StaticString("ALLOW"),
		MarkdownDescription: "The permission to use for this address.",
	},
}

func ExpandViewLastQueriedAcl(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.ViewLastQueriedAcl {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ViewLastQueriedAclModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ViewLastQueriedAclModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.ViewLastQueriedAcl {
	if m == nil {
		return nil
	}
	to := &dns.ViewLastQueriedAcl{
		Address:    flex.ExpandStringPointer(m.Address),
		Permission: flex.ExpandStringPointer(m.Permission),
	}
	return to
}

func FlattenViewLastQueriedAcl(ctx context.Context, from *dns.ViewLastQueriedAcl, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ViewLastQueriedAclAttrTypes)
	}
	m := ViewLastQueriedAclModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ViewLastQueriedAclAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ViewLastQueriedAclModel) Flatten(ctx context.Context, from *dns.ViewLastQueriedAcl, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ViewLastQueriedAclModel{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Permission = flex.FlattenStringPointer(from.Permission)
}
