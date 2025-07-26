package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type Ipv6networkcontainerLogicFilterRulesModel struct {
	Filter types.String `tfsdk:"filter"`
	Type   types.String `tfsdk:"type"`
}

var Ipv6networkcontainerLogicFilterRulesAttrTypes = map[string]attr.Type{
	"filter": types.StringType,
	"type":   types.StringType,
}

var Ipv6networkcontainerLogicFilterRulesResourceSchemaAttributes = map[string]schema.Attribute{
	"filter": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The filter name.",
	},
	"type": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The filter type. Valid values are: * MAC * NAC * Option",
		Validators: []validator.String{
			stringvalidator.OneOf("MAC", "NAC", "Option"),
		},
	},
}

func ExpandIpv6networkcontainerLogicFilterRules(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.Ipv6networkcontainerLogicFilterRules {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6networkcontainerLogicFilterRulesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *Ipv6networkcontainerLogicFilterRulesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.Ipv6networkcontainerLogicFilterRules {
	if m == nil {
		return nil
	}
	to := &ipam.Ipv6networkcontainerLogicFilterRules{
		Filter: flex.ExpandStringPointer(m.Filter),
		Type:   flex.ExpandStringPointer(m.Type),
	}
	return to
}

func FlattenIpv6networkcontainerLogicFilterRules(ctx context.Context, from *ipam.Ipv6networkcontainerLogicFilterRules, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6networkcontainerLogicFilterRulesAttrTypes)
	}
	m := Ipv6networkcontainerLogicFilterRulesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6networkcontainerLogicFilterRulesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Ipv6networkcontainerLogicFilterRulesModel) Flatten(ctx context.Context, from *ipam.Ipv6networkcontainerLogicFilterRules, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Ipv6networkcontainerLogicFilterRulesModel{}
	}
	m.Filter = flex.FlattenStringPointer(from.Filter)
	m.Type = flex.FlattenStringPointer(from.Type)
}
