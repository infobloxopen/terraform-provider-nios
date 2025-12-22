package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type MemberDnsCustomRootNameServersModel struct {
	Address                      types.String `tfsdk:"address"`
	Name                         types.String `tfsdk:"name"`
	SharedWithMsParentDelegation types.Bool   `tfsdk:"shared_with_ms_parent_delegation"`
	Stealth                      types.Bool   `tfsdk:"stealth"`
	TsigKey                      types.String `tfsdk:"tsig_key"`
	TsigKeyAlg                   types.String `tfsdk:"tsig_key_alg"`
	TsigKeyName                  types.String `tfsdk:"tsig_key_name"`
	UseTsigKeyName               types.Bool   `tfsdk:"use_tsig_key_name"`
}

var MemberDnsCustomRootNameServersAttrTypes = map[string]attr.Type{
	"address":                          types.StringType,
	"name":                             types.StringType,
	"shared_with_ms_parent_delegation": types.BoolType,
	"stealth":                          types.BoolType,
	"tsig_key":                         types.StringType,
	"tsig_key_alg":                     types.StringType,
	"tsig_key_name":                    types.StringType,
	"use_tsig_key_name":                types.BoolType,
}

var MemberDnsCustomRootNameServersResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The IPv4 Address or IPv6 Address of the server.",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "A resolvable domain name for the external DNS server.",
	},
	"shared_with_ms_parent_delegation": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "This flag represents whether the name server is shared with the parent Microsoft primary zone's delegation server.",
	},
	"stealth": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Set this flag to hide the NS record for the primary name server from DNS queries.",
	},
	"tsig_key": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "A generated TSIG key.",
	},
	"tsig_key_alg": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The TSIG key algorithm.",
	},
	"tsig_key_name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The TSIG key name.",
	},
	"use_tsig_key_name": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: tsig_key_name",
	},
}

func ExpandMemberDnsCustomRootNameServers(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberDnsCustomRootNameServers {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberDnsCustomRootNameServersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberDnsCustomRootNameServersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberDnsCustomRootNameServers {
	if m == nil {
		return nil
	}
	to := &grid.MemberDnsCustomRootNameServers{
		Address:        flex.ExpandStringPointer(m.Address),
		Name:           flex.ExpandStringPointer(m.Name),
		Stealth:        flex.ExpandBoolPointer(m.Stealth),
		TsigKey:        flex.ExpandStringPointer(m.TsigKey),
		TsigKeyAlg:     flex.ExpandStringPointer(m.TsigKeyAlg),
		TsigKeyName:    flex.ExpandStringPointer(m.TsigKeyName),
		UseTsigKeyName: flex.ExpandBoolPointer(m.UseTsigKeyName),
	}
	return to
}

func FlattenMemberDnsCustomRootNameServers(ctx context.Context, from *grid.MemberDnsCustomRootNameServers, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberDnsCustomRootNameServersAttrTypes)
	}
	m := MemberDnsCustomRootNameServersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberDnsCustomRootNameServersAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberDnsCustomRootNameServersModel) Flatten(ctx context.Context, from *grid.MemberDnsCustomRootNameServers, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberDnsCustomRootNameServersModel{}
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
