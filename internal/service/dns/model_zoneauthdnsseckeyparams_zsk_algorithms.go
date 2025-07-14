package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type ZoneauthdnsseckeyparamsZskAlgorithmsModel struct {
	Algorithm types.String `tfsdk:"algorithm"`
	Size      types.Int64  `tfsdk:"size"`
}

var ZoneauthdnsseckeyparamsZskAlgorithmsAttrTypes = map[string]attr.Type{
	"algorithm": types.StringType,
	"size":      types.Int64Type,
}

var ZoneauthdnsseckeyparamsZskAlgorithmsResourceSchemaAttributes = map[string]schema.Attribute{
	"algorithm": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Default:  stringdefault.StaticString("RSASHA256"),
		Validators: []validator.String{
			stringvalidator.OneOf("ECDSAP256SHA256", "ECDSAP384SHA384", "RSASHA1", "RSASHA256", "RSASHA512"),
		},
		MarkdownDescription: "The signing key algorithm.",
	},
	"size": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(1024),
		MarkdownDescription: "The signing key size, in bits.",
	},
}

func ExpandZoneauthdnsseckeyparamsZskAlgorithms(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.ZoneauthdnsseckeyparamsZskAlgorithms {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneauthdnsseckeyparamsZskAlgorithmsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ZoneauthdnsseckeyparamsZskAlgorithmsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.ZoneauthdnsseckeyparamsZskAlgorithms {
	if m == nil {
		return nil
	}
	to := &dns.ZoneauthdnsseckeyparamsZskAlgorithms{
		Algorithm: flex.ExpandStringPointer(m.Algorithm),
		Size:      flex.ExpandInt64Pointer(m.Size),
	}
	return to
}

func FlattenZoneauthdnsseckeyparamsZskAlgorithms(ctx context.Context, from *dns.ZoneauthdnsseckeyparamsZskAlgorithms, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneauthdnsseckeyparamsZskAlgorithmsAttrTypes)
	}
	m := ZoneauthdnsseckeyparamsZskAlgorithmsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZoneauthdnsseckeyparamsZskAlgorithmsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ZoneauthdnsseckeyparamsZskAlgorithmsModel) Flatten(ctx context.Context, from *dns.ZoneauthdnsseckeyparamsZskAlgorithms, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ZoneauthdnsseckeyparamsZskAlgorithmsModel{}
	}
	m.Algorithm = flex.FlattenStringPointer(from.Algorithm)
	m.Size = flex.FlattenInt64Pointer(from.Size)
}
