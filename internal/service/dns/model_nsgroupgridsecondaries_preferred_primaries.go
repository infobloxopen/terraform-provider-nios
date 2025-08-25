package dns

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type NsgroupgridsecondariesPreferredPrimariesModel struct {
	Address                      iptypes.IPAddress `tfsdk:"address"`
	Name                         types.String      `tfsdk:"name"`
	SharedWithMsParentDelegation types.Bool        `tfsdk:"shared_with_ms_parent_delegation"`
	Stealth                      types.Bool        `tfsdk:"stealth"`
	TsigKey                      types.String      `tfsdk:"tsig_key"`
	TsigKeyAlg                   types.String      `tfsdk:"tsig_key_alg"`
	TsigKeyName                  types.String      `tfsdk:"tsig_key_name"`
	UseTsigKeyName               types.Bool        `tfsdk:"use_tsig_key_name"`
}

var NsgroupgridsecondariesPreferredPrimariesAttrTypes = map[string]attr.Type{
	"address":                          iptypes.IPAddressType{},
	"name":                             types.StringType,
	"shared_with_ms_parent_delegation": types.BoolType,
	"stealth":                          types.BoolType,
	"tsig_key":                         types.StringType,
	"tsig_key_alg":                     types.StringType,
	"tsig_key_name":                    types.StringType,
	"use_tsig_key_name":                types.BoolType,
}

var NsgroupgridsecondariesPreferredPrimariesResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		CustomType: iptypes.IPAddressType{},
		Required:   true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Should not have leading or trailing whitespace",
			),
		},
		MarkdownDescription: "The IPv4 Address or IPv6 Address of the server.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Should not have leading or trailing whitespace",
			),
		},
		MarkdownDescription: "A resolvable domain name for the external DNS server.",
	},
	"shared_with_ms_parent_delegation": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "This flag represents whether the name server is shared with the parent Microsoft primary zone's delegation server.",
	},
	"stealth": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Set this flag to hide the NS record for the primary name server from DNS queries.",
	},
	"tsig_key": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Should not have leading or trailing whitespace",
			),
		},
		MarkdownDescription: "A generated TSIG key.",
	},
	"tsig_key_alg": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("HMAC-MD5"),
		Validators: []validator.String{
			stringvalidator.OneOf("HMAC-MD5", "HMAC-SHA256"),
		},
		MarkdownDescription: "The TSIG key algorithm.",
	},
	"tsig_key_name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Should not have leading or trailing whitespace",
			),
			stringvalidator.AlsoRequires(path.MatchRoot("use_tsig_key_name")),
		},
		MarkdownDescription: "The TSIG key name.",
	},
	"use_tsig_key_name": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: tsig_key_name",
	},
}

func ExpandNsgroupgridsecondariesPreferredPrimaries(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.NsgroupgridsecondariesPreferredPrimaries {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NsgroupgridsecondariesPreferredPrimariesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NsgroupgridsecondariesPreferredPrimariesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.NsgroupgridsecondariesPreferredPrimaries {
	if m == nil {
		return nil
	}
	to := &dns.NsgroupgridsecondariesPreferredPrimaries{
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

func FlattenNsgroupgridsecondariesPreferredPrimaries(ctx context.Context, from *dns.NsgroupgridsecondariesPreferredPrimaries, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NsgroupgridsecondariesPreferredPrimariesAttrTypes)
	}
	m := NsgroupgridsecondariesPreferredPrimariesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NsgroupgridsecondariesPreferredPrimariesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NsgroupgridsecondariesPreferredPrimariesModel) Flatten(ctx context.Context, from *dns.NsgroupgridsecondariesPreferredPrimaries, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NsgroupgridsecondariesPreferredPrimariesModel{}
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
