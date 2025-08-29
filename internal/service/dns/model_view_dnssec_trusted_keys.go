package dns

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type ViewDnssecTrustedKeysModel struct {
	Fqdn               types.String `tfsdk:"fqdn"`
	Algorithm          types.String `tfsdk:"algorithm"`
	Key                types.String `tfsdk:"key"`
	SecureEntryPoint   types.Bool   `tfsdk:"secure_entry_point"`
	DnssecMustBeSecure types.Bool   `tfsdk:"dnssec_must_be_secure"`
}

var ViewDnssecTrustedKeysAttrTypes = map[string]attr.Type{
	"fqdn":                  types.StringType,
	"algorithm":             types.StringType,
	"key":                   types.StringType,
	"secure_entry_point":    types.BoolType,
	"dnssec_must_be_secure": types.BoolType,
}

var ViewDnssecTrustedKeysResourceSchemaAttributes = map[string]schema.Attribute{
	"fqdn": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Should not have leading or trailing whitespace",
			),
		},
		MarkdownDescription: "The FQDN of the domain for which the member validates responses to recursive queries.",
	},
	"algorithm": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Should not have leading or trailing whitespace",
			),
		},
		MarkdownDescription: "The DNSSEC algorithm used to generate the key.",
	},
	"key": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Should not have leading or trailing whitespace",
			),
		},
		MarkdownDescription: "The DNSSEC key.",
	},
	"secure_entry_point": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The secure entry point flag, if set it means this is a KSK configuration.",
	},
	"dnssec_must_be_secure": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Responses must be DNSSEC secure for this hierarchy/domain.",
	},
}

func ExpandViewDnssecTrustedKeys(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.ViewDnssecTrustedKeys {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ViewDnssecTrustedKeysModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ViewDnssecTrustedKeysModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.ViewDnssecTrustedKeys {
	if m == nil {
		return nil
	}
	to := &dns.ViewDnssecTrustedKeys{
		Fqdn:               flex.ExpandStringPointer(m.Fqdn),
		Algorithm:          flex.ExpandStringPointer(m.Algorithm),
		Key:                flex.ExpandStringPointer(m.Key),
		SecureEntryPoint:   flex.ExpandBoolPointer(m.SecureEntryPoint),
		DnssecMustBeSecure: flex.ExpandBoolPointer(m.DnssecMustBeSecure),
	}
	return to
}

func FlattenViewDnssecTrustedKeys(ctx context.Context, from *dns.ViewDnssecTrustedKeys, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ViewDnssecTrustedKeysAttrTypes)
	}
	m := ViewDnssecTrustedKeysModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ViewDnssecTrustedKeysAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ViewDnssecTrustedKeysModel) Flatten(ctx context.Context, from *dns.ViewDnssecTrustedKeys, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ViewDnssecTrustedKeysModel{}
	}
	m.Fqdn = flex.FlattenStringPointer(from.Fqdn)
	m.Algorithm = flex.FlattenStringPointer(from.Algorithm)
	m.Key = flex.FlattenStringPointer(from.Key)
	m.SecureEntryPoint = types.BoolPointerValue(from.SecureEntryPoint)
	m.DnssecMustBeSecure = types.BoolPointerValue(from.DnssecMustBeSecure)
}
