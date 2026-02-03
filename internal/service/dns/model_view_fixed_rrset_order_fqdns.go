package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type ViewFixedRrsetOrderFqdnsModel struct {
	Fqdn       types.String `tfsdk:"fqdn"`
	RecordType types.String `tfsdk:"record_type"`
}

var ViewFixedRrsetOrderFqdnsAttrTypes = map[string]attr.Type{
	"fqdn":        types.StringType,
	"record_type": types.StringType,
}

var ViewFixedRrsetOrderFqdnsResourceSchemaAttributes = map[string]schema.Attribute{
	"fqdn": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.IsValidDomainName(),
		},
		MarkdownDescription: "The FQDN of the fixed RRset configuration item.",
	},
	"record_type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("A", "AAAA", "BOTH"),
		},
		Default:             stringdefault.StaticString("A"),
		MarkdownDescription: "The record type for the specified FQDN in the fixed RRset configuration.",
	},
}

func ExpandViewFixedRrsetOrderFqdns(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.ViewFixedRrsetOrderFqdns {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ViewFixedRrsetOrderFqdnsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ViewFixedRrsetOrderFqdnsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.ViewFixedRrsetOrderFqdns {
	if m == nil {
		return nil
	}
	to := &dns.ViewFixedRrsetOrderFqdns{
		Fqdn:       flex.ExpandStringPointer(m.Fqdn),
		RecordType: flex.ExpandStringPointer(m.RecordType),
	}
	return to
}

func FlattenViewFixedRrsetOrderFqdns(ctx context.Context, from *dns.ViewFixedRrsetOrderFqdns, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ViewFixedRrsetOrderFqdnsAttrTypes)
	}
	m := ViewFixedRrsetOrderFqdnsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ViewFixedRrsetOrderFqdnsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ViewFixedRrsetOrderFqdnsModel) Flatten(ctx context.Context, from *dns.ViewFixedRrsetOrderFqdns, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ViewFixedRrsetOrderFqdnsModel{}
	}
	m.Fqdn = flex.FlattenStringPointer(from.Fqdn)
	m.RecordType = flex.FlattenStringPointer(from.RecordType)
}
