package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type FixedaddressLogicFilterRulesModel struct {
	Filter types.String `tfsdk:"filter"`
	Type   types.String `tfsdk:"type"`
}

var FixedaddressLogicFilterRulesAttrTypes = map[string]attr.Type{
	"filter": types.StringType,
	"type":   types.StringType,
}

var FixedaddressLogicFilterRulesResourceSchemaAttributes = map[string]schema.Attribute{
	"filter": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The filter name.",
	},
	"type": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.OneOf("MAC", "NAC", "Option"),
		},
		MarkdownDescription: "The filter type. Valid values are: * MAC * NAC * Option",
	},
}

func ExpandFixedaddressLogicFilterRules(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dhcp.FixedaddressLogicFilterRules {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m FixedaddressLogicFilterRulesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *FixedaddressLogicFilterRulesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.FixedaddressLogicFilterRules {
	if m == nil {
		return nil
	}
	to := &dhcp.FixedaddressLogicFilterRules{
		Filter: flex.ExpandStringPointer(m.Filter),
		Type:   flex.ExpandStringPointer(m.Type),
	}
	return to
}

func FlattenFixedaddressLogicFilterRules(ctx context.Context, from *dhcp.FixedaddressLogicFilterRules, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(FixedaddressLogicFilterRulesAttrTypes)
	}
	m := FixedaddressLogicFilterRulesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, FixedaddressLogicFilterRulesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *FixedaddressLogicFilterRulesModel) Flatten(ctx context.Context, from *dhcp.FixedaddressLogicFilterRules, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = FixedaddressLogicFilterRulesModel{}
	}
	m.Filter = flex.FlattenStringPointer(from.Filter)
	m.Type = flex.FlattenStringPointer(from.Type)
}
