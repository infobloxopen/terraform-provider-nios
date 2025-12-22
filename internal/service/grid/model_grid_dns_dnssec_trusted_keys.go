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

type GridDnsDnssecTrustedKeysModel struct {
	Fqdn               types.String `tfsdk:"fqdn"`
	Algorithm          types.String `tfsdk:"algorithm"`
	Key                types.String `tfsdk:"key"`
	SecureEntryPoint   types.Bool   `tfsdk:"secure_entry_point"`
	DnssecMustBeSecure types.Bool   `tfsdk:"dnssec_must_be_secure"`
}

var GridDnsDnssecTrustedKeysAttrTypes = map[string]attr.Type{
	"fqdn":                  types.StringType,
	"algorithm":             types.StringType,
	"key":                   types.StringType,
	"secure_entry_point":    types.BoolType,
	"dnssec_must_be_secure": types.BoolType,
}

var GridDnsDnssecTrustedKeysResourceSchemaAttributes = map[string]schema.Attribute{
	"fqdn": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The FQDN of the domain for which the member validates responses to recursive queries.",
	},
	"algorithm": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The DNSSEC algorithm used to generate the key.",
	},
	"key": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The DNSSEC key.",
	},
	"secure_entry_point": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "The secure entry point flag, if set it means this is a KSK configuration.",
	},
	"dnssec_must_be_secure": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Responses must be DNSSEC secure for this hierarchy/domain.",
	},
}

func ExpandGridDnsDnssecTrustedKeys(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridDnsDnssecTrustedKeys {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridDnsDnssecTrustedKeysModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridDnsDnssecTrustedKeysModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridDnsDnssecTrustedKeys {
	if m == nil {
		return nil
	}
	to := &grid.GridDnsDnssecTrustedKeys{
		Fqdn:               flex.ExpandStringPointer(m.Fqdn),
		Algorithm:          flex.ExpandStringPointer(m.Algorithm),
		Key:                flex.ExpandStringPointer(m.Key),
		SecureEntryPoint:   flex.ExpandBoolPointer(m.SecureEntryPoint),
		DnssecMustBeSecure: flex.ExpandBoolPointer(m.DnssecMustBeSecure),
	}
	return to
}

func FlattenGridDnsDnssecTrustedKeys(ctx context.Context, from *grid.GridDnsDnssecTrustedKeys, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridDnsDnssecTrustedKeysAttrTypes)
	}
	m := GridDnsDnssecTrustedKeysModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridDnsDnssecTrustedKeysAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridDnsDnssecTrustedKeysModel) Flatten(ctx context.Context, from *grid.GridDnsDnssecTrustedKeys, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridDnsDnssecTrustedKeysModel{}
	}
	m.Fqdn = flex.FlattenStringPointer(from.Fqdn)
	m.Algorithm = flex.FlattenStringPointer(from.Algorithm)
	m.Key = flex.FlattenStringPointer(from.Key)
	m.SecureEntryPoint = types.BoolPointerValue(from.SecureEntryPoint)
	m.DnssecMustBeSecure = types.BoolPointerValue(from.DnssecMustBeSecure)
}
