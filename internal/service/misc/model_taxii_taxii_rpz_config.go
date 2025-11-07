package misc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/misc"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type TaxiiTaxiiRpzConfigModel struct {
	CollectionName types.String `tfsdk:"collection_name"`
	Zone           types.String `tfsdk:"zone"`
}

var TaxiiTaxiiRpzConfigAttrTypes = map[string]attr.Type{
	"collection_name": types.StringType,
	"zone":            types.StringType,
}

var TaxiiTaxiiRpzConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"collection_name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The STIX collection name.",
	},
	"zone": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the RPZ in which rules are created through the Taxii protocol requests.",
	},
}

func ExpandTaxiiTaxiiRpzConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *misc.TaxiiTaxiiRpzConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m TaxiiTaxiiRpzConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *TaxiiTaxiiRpzConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *misc.TaxiiTaxiiRpzConfig {
	if m == nil {
		return nil
	}
	to := &misc.TaxiiTaxiiRpzConfig{
		CollectionName: flex.ExpandStringPointer(m.CollectionName),
		Zone:           flex.ExpandStringPointer(m.Zone),
	}
	return to
}

func FlattenTaxiiTaxiiRpzConfig(ctx context.Context, from *misc.TaxiiTaxiiRpzConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(TaxiiTaxiiRpzConfigAttrTypes)
	}
	m := TaxiiTaxiiRpzConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, TaxiiTaxiiRpzConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *TaxiiTaxiiRpzConfigModel) Flatten(ctx context.Context, from *misc.TaxiiTaxiiRpzConfig, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = TaxiiTaxiiRpzConfigModel{}
	}
	m.CollectionName = flex.FlattenStringPointer(from.CollectionName)
	m.Zone = flex.FlattenStringPointer(from.Zone)
}
