package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
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

type ZonerpgridprimaryPreferredPrimariesModel struct {
	Address                      iptypes.IPAddress `tfsdk:"address"`
	Name                         types.String      `tfsdk:"name"`
	SharedWithMsParentDelegation types.Bool        `tfsdk:"shared_with_ms_parent_delegation"`
	Stealth                      types.Bool        `tfsdk:"stealth"`
	TsigKey                      types.String      `tfsdk:"tsig_key"`
	TsigKeyAlg                   types.String      `tfsdk:"tsig_key_alg"`
	TsigKeyName                  types.String      `tfsdk:"tsig_key_name"`
	UseTsigKeyName               types.Bool        `tfsdk:"use_tsig_key_name"`
}

var ZonerpgridprimaryPreferredPrimariesAttrTypes = map[string]attr.Type{
	"address":                          iptypes.IPAddressType{},
	"name":                             types.StringType,
	"shared_with_ms_parent_delegation": types.BoolType,
	"stealth":                          types.BoolType,
	"tsig_key":                         types.StringType,
	"tsig_key_alg":                     types.StringType,
	"tsig_key_name":                    types.StringType,
	"use_tsig_key_name":                types.BoolType,
}

// ZonerpgridprimaryPreferredPrimariesResourceSchemaAttributes has certain fields set to Computed: true because these fields need to be set at the GRID Level and Cannot be set while Configuring Grid Primary
var ZonerpgridprimaryPreferredPrimariesResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		CustomType:          iptypes.IPAddressType{},
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

func ExpandZonerpgridprimaryPreferredPrimaries(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.ZonerpgridprimaryPreferredPrimaries {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZonerpgridprimaryPreferredPrimariesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ZonerpgridprimaryPreferredPrimariesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.ZonerpgridprimaryPreferredPrimaries {
	if m == nil {
		return nil
	}
	to := &dns.ZonerpgridprimaryPreferredPrimaries{
		Address:        flex.ExpandIPAddress(m.Address),
		Name:           flex.ExpandStringPointer(m.Name),
		Stealth:        flex.ExpandBoolPointer(m.Stealth),
		TsigKey:        flex.ExpandStringPointer(m.TsigKey),
		TsigKeyAlg:     flex.ExpandStringPointer(m.TsigKeyAlg),
		TsigKeyName:    flex.ExpandStringPointer(m.TsigKeyName),
		UseTsigKeyName: flex.ExpandBoolPointer(m.UseTsigKeyName),
	}
	return to
}

func FlattenZonerpgridprimaryPreferredPrimaries(ctx context.Context, from *dns.ZonerpgridprimaryPreferredPrimaries, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZonerpgridprimaryPreferredPrimariesAttrTypes)
	}
	m := ZonerpgridprimaryPreferredPrimariesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZonerpgridprimaryPreferredPrimariesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ZonerpgridprimaryPreferredPrimariesModel) Flatten(ctx context.Context, from *dns.ZonerpgridprimaryPreferredPrimaries, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ZonerpgridprimaryPreferredPrimariesModel{}
	}
	m.Address = flex.FlattenIPAddress(from.Address)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.SharedWithMsParentDelegation = types.BoolPointerValue(from.SharedWithMsParentDelegation)
	m.Stealth = types.BoolPointerValue(from.Stealth)
	m.TsigKey = flex.FlattenStringPointer(from.TsigKey)
	m.TsigKeyAlg = flex.FlattenStringPointer(from.TsigKeyAlg)
	m.TsigKeyName = flex.FlattenStringPointer(from.TsigKeyName)
	m.UseTsigKeyName = types.BoolPointerValue(from.UseTsigKeyName)
}
