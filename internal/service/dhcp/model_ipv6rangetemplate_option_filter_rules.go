package dhcp

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type Ipv6rangetemplateOptionFilterRulesModel struct {
	Filter     types.String `tfsdk:"filter"`
	Permission types.String `tfsdk:"permission"`
}

var Ipv6rangetemplateOptionFilterRulesAttrTypes = map[string]attr.Type{
	"filter":     types.StringType,
	"permission": types.StringType,
}

var Ipv6rangetemplateOptionFilterRulesResourceSchemaAttributes = map[string]schema.Attribute{
	"filter": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of the DHCP filter.",
	},
	"permission": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			stringvalidator.OneOf("Allow", "Deny"),
		},
		MarkdownDescription: "The permission to be applied.",
	},
}

func ExpandIpv6rangetemplateOptionFilterRules(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dhcp.Ipv6rangetemplateOptionFilterRules {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6rangetemplateOptionFilterRulesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *Ipv6rangetemplateOptionFilterRulesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.Ipv6rangetemplateOptionFilterRules {
	if m == nil {
		return nil
	}
	to := &dhcp.Ipv6rangetemplateOptionFilterRules{
		Filter:     flex.ExpandStringPointer(m.Filter),
		Permission: flex.ExpandStringPointer(m.Permission),
	}
	return to
}

func FlattenIpv6rangetemplateOptionFilterRules(ctx context.Context, from *dhcp.Ipv6rangetemplateOptionFilterRules, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6rangetemplateOptionFilterRulesAttrTypes)
	}
	m := Ipv6rangetemplateOptionFilterRulesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6rangetemplateOptionFilterRulesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Ipv6rangetemplateOptionFilterRulesModel) Flatten(ctx context.Context, from *dhcp.Ipv6rangetemplateOptionFilterRules, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Ipv6rangetemplateOptionFilterRulesModel{}
	}
	m.Filter = flex.FlattenStringPointer(from.Filter)
	m.Permission = flex.FlattenStringPointer(from.Permission)
}
