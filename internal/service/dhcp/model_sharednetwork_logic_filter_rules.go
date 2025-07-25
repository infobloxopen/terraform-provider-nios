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

type SharednetworkLogicFilterRulesModel struct {
	Filter types.String `tfsdk:"filter"`
	Type   types.String `tfsdk:"type"`
}

var SharednetworkLogicFilterRulesAttrTypes = map[string]attr.Type{
	"filter": types.StringType,
	"type":   types.StringType,
}

var SharednetworkLogicFilterRulesResourceSchemaAttributes = map[string]schema.Attribute{
	"filter": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The filter name.",
	},
	"type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("MAC", "NAC", "Option"),
		},
		MarkdownDescription: "The filter type. Valid values are: * MAC * NAC * Option",
	},
}

func ExpandSharednetworkLogicFilterRules(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dhcp.SharednetworkLogicFilterRules {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m SharednetworkLogicFilterRulesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *SharednetworkLogicFilterRulesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.SharednetworkLogicFilterRules {
	if m == nil {
		return nil
	}
	to := &dhcp.SharednetworkLogicFilterRules{
		Filter: flex.ExpandStringPointer(m.Filter),
		Type:   flex.ExpandStringPointer(m.Type),
	}
	return to
}

func FlattenSharednetworkLogicFilterRules(ctx context.Context, from *dhcp.SharednetworkLogicFilterRules, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SharednetworkLogicFilterRulesAttrTypes)
	}
	m := SharednetworkLogicFilterRulesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, SharednetworkLogicFilterRulesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *SharednetworkLogicFilterRulesModel) Flatten(ctx context.Context, from *dhcp.SharednetworkLogicFilterRules, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = SharednetworkLogicFilterRulesModel{}
	}
	m.Filter = flex.FlattenStringPointer(from.Filter)
	m.Type = flex.FlattenStringPointer(from.Type)
}
