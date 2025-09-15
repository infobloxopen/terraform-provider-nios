package dns

import (
	"context"

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

type ZoneauthdnsseckeyparamsKskAlgorithmsModel struct {
	Algorithm types.String `tfsdk:"algorithm"`
	Size      types.Int64  `tfsdk:"size"`
}

var ZoneauthdnsseckeyparamsKskAlgorithmsAttrTypes = map[string]attr.Type{
	"algorithm": types.StringType,
	"size":      types.Int64Type,
}

var ZoneauthdnsseckeyparamsKskAlgorithmsResourceSchemaAttributes = map[string]schema.Attribute{
	"algorithm": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Validators: []validator.String{
			stringvalidator.OneOf("ECDSAP256SHA256", "ECDSAP384SHA384", "RSASHA1", "RSASHA256", "RSASHA512"),
		},
		MarkdownDescription: "The signing key algorithm.",
	},
	"size": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The signing key size, in bits.",
	},
}

func ExpandZoneauthdnsseckeyparamsKskAlgorithms(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.ZoneauthdnsseckeyparamsKskAlgorithms {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneauthdnsseckeyparamsKskAlgorithmsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ZoneauthdnsseckeyparamsKskAlgorithmsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.ZoneauthdnsseckeyparamsKskAlgorithms {
	if m == nil {
		return nil
	}
	to := &dns.ZoneauthdnsseckeyparamsKskAlgorithms{
		Algorithm: flex.ExpandStringPointer(m.Algorithm),
		Size:      flex.ExpandInt64Pointer(m.Size),
	}
	return to
}

func FlattenZoneauthdnsseckeyparamsKskAlgorithms(ctx context.Context, from *dns.ZoneauthdnsseckeyparamsKskAlgorithms, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneauthdnsseckeyparamsKskAlgorithmsAttrTypes)
	}
	m := ZoneauthdnsseckeyparamsKskAlgorithmsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZoneauthdnsseckeyparamsKskAlgorithmsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ZoneauthdnsseckeyparamsKskAlgorithmsModel) Flatten(ctx context.Context, from *dns.ZoneauthdnsseckeyparamsKskAlgorithms, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ZoneauthdnsseckeyparamsKskAlgorithmsModel{}
	}
	m.Algorithm = flex.FlattenStringPointer(from.Algorithm)
	m.Size = flex.FlattenInt64Pointer(from.Size)
}
