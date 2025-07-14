package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type NetworkcontainerLogicFilterRulesModel struct {
	Filter types.String `tfsdk:"filter"`
	Type   types.String `tfsdk:"type"`
}

var NetworkcontainerLogicFilterRulesAttrTypes = map[string]attr.Type{
	"filter": types.StringType,
	"type":   types.StringType,
}

var NetworkcontainerLogicFilterRulesResourceSchemaAttributes = map[string]schema.Attribute{
	"filter": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The filter name.",
		Computed:            true,
	},
	"type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The filter type. Valid values are: * MAC * NAC * Option",
		Validators: []validator.String{
			stringvalidator.OneOf("MAC", "NAC", "Option"),
		},
		Computed: true,
	},
}

func ExpandNetworkcontainerLogicFilterRules(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.NetworkcontainerLogicFilterRules {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkcontainerLogicFilterRulesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NetworkcontainerLogicFilterRulesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.NetworkcontainerLogicFilterRules {
	if m == nil {
		return nil
	}
	to := &ipam.NetworkcontainerLogicFilterRules{
		Filter: flex.ExpandStringPointer(m.Filter),
		Type:   flex.ExpandStringPointer(m.Type),
	}
	return to
}

func FlattenNetworkcontainerLogicFilterRules(ctx context.Context, from *ipam.NetworkcontainerLogicFilterRules, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkcontainerLogicFilterRulesAttrTypes)
	}
	m := NetworkcontainerLogicFilterRulesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkcontainerLogicFilterRulesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NetworkcontainerLogicFilterRulesModel) Flatten(ctx context.Context, from *ipam.NetworkcontainerLogicFilterRules, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NetworkcontainerLogicFilterRulesModel{}
	}
	m.Filter = flex.FlattenStringPointer(from.Filter)
	m.Type = flex.FlattenStringPointer(from.Type)
}
