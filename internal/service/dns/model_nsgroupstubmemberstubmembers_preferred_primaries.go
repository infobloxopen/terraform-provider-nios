package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	"github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/plancontrol"
)

type NsgroupstubmemberstubmembersPreferredPrimariesModel struct {
	Address                      types.String `tfsdk:"address"`
	Name                         types.String `tfsdk:"name"`
	SharedWithMsParentDelegation types.Bool   `tfsdk:"shared_with_ms_parent_delegation"`
	Stealth                      types.Bool   `tfsdk:"stealth"`
	TsigKey                      types.String `tfsdk:"tsig_key"`
	TsigKeyAlg                   types.String `tfsdk:"tsig_key_alg"`
	TsigKeyName                  types.String `tfsdk:"tsig_key_name"`
	UseTsigKeyName               types.Bool   `tfsdk:"use_tsig_key_name"`
}

var NsgroupstubmemberstubmembersPreferredPrimariesAttrTypes = map[string]attr.Type{
	"address":                          types.StringType,
	"name":                             types.StringType,
	"shared_with_ms_parent_delegation": types.BoolType,
	"stealth":                          types.BoolType,
	"tsig_key":                         types.StringType,
	"tsig_key_alg":                     types.StringType,
	"tsig_key_name":                    types.StringType,
	"use_tsig_key_name":                types.BoolType,
}

var NsgroupstubmemberstubmembersPreferredPrimariesResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The IPv4 Address or IPv6 Address of the server.",
		PlanModifiers: []planmodifier.String{
			plancontrol.UseStateForUnknownString(),
		},
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "A resolvable domain name for the external DNS server.",
		PlanModifiers: []planmodifier.String{
			plancontrol.UseStateForUnknownString(),
		},
	},
	"shared_with_ms_parent_delegation": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "This flag represents whether the name server is shared with the parent Microsoft primary zone's delegation server.",
		PlanModifiers: []planmodifier.Bool{
			plancontrol.UseStateForUnknownBool(),
		},
	},
	"stealth": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Set this flag to hide the NS record for the primary name server from DNS queries.",
		PlanModifiers: []planmodifier.Bool{
			plancontrol.UseStateForUnknownBool(),
		},
	},
	"tsig_key": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "A generated TSIG key.",
		PlanModifiers: []planmodifier.String{
			plancontrol.UseStateForUnknownString(),
		},
	},
	"tsig_key_alg": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The TSIG key algorithm.",
		PlanModifiers: []planmodifier.String{
			plancontrol.UseStateForUnknownString(),
		},
	},
	"tsig_key_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The TSIG key name.",
		PlanModifiers: []planmodifier.String{
			plancontrol.UseStateForUnknownString(),
		},
	},
	"use_tsig_key_name": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Use flag for: tsig_key_name",
		PlanModifiers: []planmodifier.Bool{
			plancontrol.UseStateForUnknownBool(),
		},
	},
}

func ExpandNsgroupstubmemberstubmembersPreferredPrimaries(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.NsgroupstubmemberstubmembersPreferredPrimaries {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NsgroupstubmemberstubmembersPreferredPrimariesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NsgroupstubmemberstubmembersPreferredPrimariesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.NsgroupstubmemberstubmembersPreferredPrimaries {
	if m == nil {
		return nil
	}
	to := &dns.NsgroupstubmemberstubmembersPreferredPrimaries{}
	return to
}

func FlattenNsgroupstubmemberstubmembersPreferredPrimaries(ctx context.Context, from *dns.NsgroupstubmemberstubmembersPreferredPrimaries, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NsgroupstubmemberstubmembersPreferredPrimariesAttrTypes)
	}
	m := NsgroupstubmemberstubmembersPreferredPrimariesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NsgroupstubmemberstubmembersPreferredPrimariesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NsgroupstubmemberstubmembersPreferredPrimariesModel) Flatten(ctx context.Context, from *dns.NsgroupstubmemberstubmembersPreferredPrimaries, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NsgroupstubmemberstubmembersPreferredPrimariesModel{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.SharedWithMsParentDelegation = types.BoolPointerValue(from.SharedWithMsParentDelegation)
	m.Stealth = types.BoolPointerValue(from.Stealth)
	m.TsigKey = flex.FlattenStringPointer(from.TsigKey)
	m.TsigKeyAlg = flex.FlattenStringPointer(from.TsigKeyAlg)
	m.TsigKeyName = flex.FlattenStringPointer(from.TsigKeyName)
	m.UseTsigKeyName = types.BoolPointerValue(from.UseTsigKeyName)
}
