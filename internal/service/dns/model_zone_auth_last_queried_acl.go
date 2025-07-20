package dns

import (
	"context"
	"regexp"

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
)

type ZoneAuthLastQueriedAclModel struct {
	Address    types.String `tfsdk:"address"`
	Permission types.String `tfsdk:"permission"`
}

var ZoneAuthLastQueriedAclAttrTypes = map[string]attr.Type{
	"address":    types.StringType,
	"permission": types.StringType,
}

var ZoneAuthLastQueriedAclResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The address this rule applies to or \"Any\".",
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Address should not have leading or trailing whitespace",
			),
		},
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

func ExpandZoneAuthLastQueriedAcl(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.ZoneAuthLastQueriedAcl {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneAuthLastQueriedAclModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ZoneAuthLastQueriedAclModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.ZoneAuthLastQueriedAcl {
	if m == nil {
		return nil
	}
	to := &dns.ZoneAuthLastQueriedAcl{
		Address:    flex.ExpandStringPointer(m.Address),
		Permission: flex.ExpandStringPointer(m.Permission),
	}
	return to
}

func FlattenZoneAuthLastQueriedAcl(ctx context.Context, from *dns.ZoneAuthLastQueriedAcl, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneAuthLastQueriedAclAttrTypes)
	}
	m := ZoneAuthLastQueriedAclModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZoneAuthLastQueriedAclAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ZoneAuthLastQueriedAclModel) Flatten(ctx context.Context, from *dns.ZoneAuthLastQueriedAcl, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ZoneAuthLastQueriedAclModel{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Permission = flex.FlattenStringPointer(from.Permission)
}
