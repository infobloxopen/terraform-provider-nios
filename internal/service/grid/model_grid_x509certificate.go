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

type GridX509certificateModel struct {
	Ref            types.String `tfsdk:"ref"`
	Issuer         types.String `tfsdk:"issuer"`
	Serial         types.String `tfsdk:"serial"`
	Subject        types.String `tfsdk:"subject"`
	ValidNotAfter  types.Int64  `tfsdk:"valid_not_after"`
	ValidNotBefore types.Int64  `tfsdk:"valid_not_before"`
}

var GridX509certificateAttrTypes = map[string]attr.Type{
	"ref":              types.StringType,
	"issuer":           types.StringType,
	"serial":           types.StringType,
	"subject":          types.StringType,
	"valid_not_after":  types.Int64Type,
	"valid_not_before": types.Int64Type,
}

var GridX509certificateResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"issuer": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Certificate issuer.",
	},
	"serial": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "X509Certificate serial number.",
	},
	"subject": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "A Distinguished Name that is made of multiple relative distinguished names (RDNs).",
	},
	"valid_not_after": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Certificate expiry date.",
	},
	"valid_not_before": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Certificate validity start date.",
	},
}

func ExpandGridX509certificate(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridX509certificate {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridX509certificateModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridX509certificateModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridX509certificate {
	if m == nil {
		return nil
	}
	to := &grid.GridX509certificate{
		Ref: flex.ExpandStringPointer(m.Ref),
	}
	return to
}

func FlattenGridX509certificate(ctx context.Context, from *grid.GridX509certificate, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridX509certificateAttrTypes)
	}
	m := GridX509certificateModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridX509certificateAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridX509certificateModel) Flatten(ctx context.Context, from *grid.GridX509certificate, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridX509certificateModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Issuer = flex.FlattenStringPointer(from.Issuer)
	m.Serial = flex.FlattenStringPointer(from.Serial)
	m.Subject = flex.FlattenStringPointer(from.Subject)
	m.ValidNotAfter = flex.FlattenInt64Pointer(from.ValidNotAfter)
	m.ValidNotBefore = flex.FlattenInt64Pointer(from.ValidNotBefore)
}
