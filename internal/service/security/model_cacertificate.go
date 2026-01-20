package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type CacertificateModel struct {
	Ref               types.String `tfsdk:"ref"`
	Uuid              types.String `tfsdk:"uuid"`
	DistinguishedName types.String `tfsdk:"distinguished_name"`
	Issuer            types.String `tfsdk:"issuer"`
	Serial            types.String `tfsdk:"serial"`
	UsedBy            types.String `tfsdk:"used_by"`
	ValidNotAfter     types.Int64  `tfsdk:"valid_not_after"`
	ValidNotBefore    types.Int64  `tfsdk:"valid_not_before"`
}

var CacertificateAttrTypes = map[string]attr.Type{
	"ref":                types.StringType,
	"uuid":               types.StringType,
	"distinguished_name": types.StringType,
	"issuer":             types.StringType,
	"serial":             types.StringType,
	"used_by":            types.StringType,
	"valid_not_after":    types.Int64Type,
	"valid_not_before":   types.Int64Type,
}

var CacertificateResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The uuid to the object.",
	},
	"distinguished_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The certificate subject name.",
	},
	"issuer": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The certificate issuer subject name.",
	},
	"serial": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The certificate serial number in hex format.",
	},
	"used_by": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Information about the CA certificate usage.",
	},
	"valid_not_after": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The date after which the certificate becomes invalid.",
	},
	"valid_not_before": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The date before which the certificate is not valid.",
	},
}

func ExpandCacertificate(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.Cacertificate {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m CacertificateModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *CacertificateModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.Cacertificate {
	if m == nil {
		return nil
	}
	to := &security.Cacertificate{
		Ref:  flex.ExpandStringPointer(m.Ref),
		Uuid: flex.ExpandStringPointer(m.Uuid),
	}
	return to
}

func FlattenCacertificate(ctx context.Context, from *security.Cacertificate, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(CacertificateAttrTypes)
	}
	m := CacertificateModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, CacertificateAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *CacertificateModel) Flatten(ctx context.Context, from *security.Cacertificate, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = CacertificateModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.DistinguishedName = flex.FlattenStringPointer(from.DistinguishedName)
	m.Issuer = flex.FlattenStringPointer(from.Issuer)
	m.Serial = flex.FlattenStringPointer(from.Serial)
	m.UsedBy = flex.FlattenStringPointer(from.UsedBy)
	m.ValidNotAfter = flex.FlattenInt64Pointer(from.ValidNotAfter)
	m.ValidNotBefore = flex.FlattenInt64Pointer(from.ValidNotBefore)
}
