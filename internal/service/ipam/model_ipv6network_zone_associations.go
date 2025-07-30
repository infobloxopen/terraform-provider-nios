package ipam

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

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type Ipv6networkZoneAssociationsModel struct {
	Fqdn      types.String `tfsdk:"fqdn"`
	IsDefault types.Bool   `tfsdk:"is_default"`
	View      types.String `tfsdk:"view"`
}

var Ipv6networkZoneAssociationsAttrTypes = map[string]attr.Type{
	"fqdn":       types.StringType,
	"is_default": types.BoolType,
	"view":       types.StringType,
}

var Ipv6networkZoneAssociationsResourceSchemaAttributes = map[string]schema.Attribute{
	"fqdn": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The FQDN of the authoritative forward zone.",
		Validators: []validator.String{
			customvalidator.IsValidFQDN(),
		},
	},
	"is_default": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "True if this is the default zone.",
	},
	"view": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The view to which the zone belongs. If a view is not specified, the default view is used.",
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Should not have leading or trailing whitespace",
			),
		},
	},
}

func ExpandIpv6networkZoneAssociations(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.Ipv6networkZoneAssociations {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6networkZoneAssociationsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *Ipv6networkZoneAssociationsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.Ipv6networkZoneAssociations {
	if m == nil {
		return nil
	}
	to := &ipam.Ipv6networkZoneAssociations{
		Fqdn:      flex.ExpandStringPointer(m.Fqdn),
		IsDefault: flex.ExpandBoolPointer(m.IsDefault),
		View:      flex.ExpandStringPointer(m.View),
	}
	return to
}

func FlattenIpv6networkZoneAssociations(ctx context.Context, from *ipam.Ipv6networkZoneAssociations, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6networkZoneAssociationsAttrTypes)
	}
	m := Ipv6networkZoneAssociationsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6networkZoneAssociationsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Ipv6networkZoneAssociationsModel) Flatten(ctx context.Context, from *ipam.Ipv6networkZoneAssociations, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Ipv6networkZoneAssociationsModel{}
	}
	m.Fqdn = flex.FlattenStringPointer(from.Fqdn)
	m.IsDefault = types.BoolPointerValue(from.IsDefault)
	m.View = flex.FlattenStringPointer(from.View)
}
