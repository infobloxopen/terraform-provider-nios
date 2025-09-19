package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	internaltypes "github.com/infobloxopen/terraform-provider-nios/internal/types"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

// ZoneStubStubFromModel defines the model for a stub zone's primary servers.
// All attributes of this model except 'name' and 'address' are Computed, as input values for stub from are ignored.

type ZoneStubStubFromModel struct {
	Address                      iptypes.IPAddress                        `tfsdk:"address"`
	Name                         internaltypes.CaseInsensitiveStringValue `tfsdk:"name"`
	SharedWithMsParentDelegation types.Bool                               `tfsdk:"shared_with_ms_parent_delegation"`
	Stealth                      types.Bool                               `tfsdk:"stealth"`
	TsigKey                      types.String                             `tfsdk:"tsig_key"`
	TsigKeyAlg                   types.String                             `tfsdk:"tsig_key_alg"`
	TsigKeyName                  types.String                             `tfsdk:"tsig_key_name"`
	UseTsigKeyName               types.Bool                               `tfsdk:"use_tsig_key_name"`
}

var ZoneStubStubFromAttrTypes = map[string]attr.Type{
	"address":                          iptypes.IPAddressType{},
	"name":                             internaltypes.CaseInsensitiveString{},
	"shared_with_ms_parent_delegation": types.BoolType,
	"stealth":                          types.BoolType,
	"tsig_key":                         types.StringType,
	"tsig_key_alg":                     types.StringType,
	"tsig_key_name":                    types.StringType,
	"use_tsig_key_name":                types.BoolType,
}

var ZoneStubStubFromResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		CustomType:          iptypes.IPAddressType{},
		Required:            true,
		MarkdownDescription: "The IPv4 Address or IPv6 Address of the server.",
	},
	"name": schema.StringAttribute{
		CustomType: internaltypes.CaseInsensitiveString{},
		Required:   true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "A resolvable domain name for the external DNS server.",
	},
	"shared_with_ms_parent_delegation": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "This flag represents whether the name server is shared with the parent Microsoft primary zone's delegation server.",
	},
	"stealth": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Set this flag to hide the NS record for the primary name server from DNS queries.",
	},
	"tsig_key": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "A generated TSIG key.",
	},
	"tsig_key_alg": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The TSIG key algorithm.",
	},
	"tsig_key_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The TSIG key name.",
	},
	"use_tsig_key_name": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Use flag for: tsig_key_name",
	},
}

func ExpandZoneStubStubFrom(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.ZoneStubStubFrom {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneStubStubFromModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ZoneStubStubFromModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.ZoneStubStubFrom {
	if m == nil {
		return nil
	}
	to := &dns.ZoneStubStubFrom{
		Address: flex.ExpandIPAddress(m.Address),
		Name:    flex.ExpandStringPointer(m.Name.StringValue),
	}
	return to
}

func FlattenZoneStubStubFrom(ctx context.Context, from *dns.ZoneStubStubFrom, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneStubStubFromAttrTypes)
	}
	m := ZoneStubStubFromModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZoneStubStubFromAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ZoneStubStubFromModel) Flatten(ctx context.Context, from *dns.ZoneStubStubFrom, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ZoneStubStubFromModel{}
	}
	m.Address = flex.FlattenIPAddress(from.Address)
	m.Name.StringValue = flex.FlattenStringPointer(from.Name)
	m.SharedWithMsParentDelegation = types.BoolPointerValue(from.SharedWithMsParentDelegation)
	m.Stealth = types.BoolPointerValue(from.Stealth)
	m.TsigKey = flex.FlattenStringPointer(from.TsigKey)
	m.TsigKeyAlg = flex.FlattenStringPointer(from.TsigKeyAlg)
	m.TsigKeyName = flex.FlattenStringPointer(from.TsigKeyName)
	m.UseTsigKeyName = types.BoolPointerValue(from.UseTsigKeyName)
}
