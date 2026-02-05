package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type GridThreatprotectionNatRulesModel struct {
	RuleType     types.String `tfsdk:"rule_type"`
	Address      types.String `tfsdk:"address"`
	Network      types.String `tfsdk:"network"`
	Cidr         types.Int64  `tfsdk:"cidr"`
	StartAddress types.String `tfsdk:"start_address"`
	EndAddress   types.String `tfsdk:"end_address"`
	NatPorts     types.List   `tfsdk:"nat_ports"`
}

var GridThreatprotectionNatRulesAttrTypes = map[string]attr.Type{
	"rule_type":     types.StringType,
	"address":       types.StringType,
	"network":       types.StringType,
	"cidr":          types.Int64Type,
	"start_address": types.StringType,
	"end_address":   types.StringType,
	"nat_ports":     types.ListType{ElemType: types.ObjectType{AttrTypes: GridthreatprotectionnatrulesNatPortsAttrTypes}},
}

var GridThreatprotectionNatRulesResourceSchemaAttributes = map[string]schema.Attribute{
	"rule_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The rule type for the threat protection NAT mapping rule.",
	},
	"address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The IP address for the threat protection NAT mapping rule.",
	},
	"network": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The network address for the threat protection NAT mapping rule.",
	},
	"cidr": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The network CIDR for the threat protection NAT mapping rule.",
	},
	"start_address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The start address for the range of the threat protection NAT mapping rule.",
	},
	"end_address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The end address for the range of the threat protection NAT mapping rule.",
	},
	"nat_ports": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: GridthreatprotectionnatrulesNatPortsResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The NAT port configuration for the threat protection NAT mapping rule.",
	},
}

func ExpandGridThreatprotectionNatRules(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridThreatprotectionNatRules {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridThreatprotectionNatRulesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridThreatprotectionNatRulesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridThreatprotectionNatRules {
	if m == nil {
		return nil
	}
	to := &grid.GridThreatprotectionNatRules{
		RuleType:     flex.ExpandStringPointer(m.RuleType),
		Address:      flex.ExpandStringPointer(m.Address),
		Network:      flex.ExpandStringPointer(m.Network),
		Cidr:         flex.ExpandInt64Pointer(m.Cidr),
		StartAddress: flex.ExpandStringPointer(m.StartAddress),
		EndAddress:   flex.ExpandStringPointer(m.EndAddress),
		NatPorts:     flex.ExpandFrameworkListNestedBlock(ctx, m.NatPorts, diags, ExpandGridthreatprotectionnatrulesNatPorts),
	}
	return to
}

func FlattenGridThreatprotectionNatRules(ctx context.Context, from *grid.GridThreatprotectionNatRules, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridThreatprotectionNatRulesAttrTypes)
	}
	m := GridThreatprotectionNatRulesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridThreatprotectionNatRulesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridThreatprotectionNatRulesModel) Flatten(ctx context.Context, from *grid.GridThreatprotectionNatRules, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridThreatprotectionNatRulesModel{}
	}
	m.RuleType = flex.FlattenStringPointer(from.RuleType)
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Network = flex.FlattenStringPointer(from.Network)
	m.Cidr = flex.FlattenInt64Pointer(from.Cidr)
	m.StartAddress = flex.FlattenStringPointer(from.StartAddress)
	m.EndAddress = flex.FlattenStringPointer(from.EndAddress)
	m.NatPorts = flex.FlattenFrameworkListNestedBlock(ctx, from.NatPorts, GridthreatprotectionnatrulesNatPortsAttrTypes, diags, FlattenGridthreatprotectionnatrulesNatPorts)
}
