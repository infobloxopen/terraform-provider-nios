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

type SharedrecordgroupZoneAssociationsModel struct {
	Fqdn types.String `tfsdk:"fqdn"`
	View types.String `tfsdk:"view"`
}

var SharedrecordgroupZoneAssociationsAttrTypes = map[string]attr.Type{
	"fqdn": types.StringType,
	"view": types.StringType,
}

var SharedrecordgroupZoneAssociationsResourceSchemaAttributes = map[string]schema.Attribute{
	"fqdn": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.Any(
				customvalidator.IsValidFQDN(),
				customvalidator.IsValidIPCIDR(),
			),
			customvalidator.IsNotArpa(),
		},
		MarkdownDescription: "The FQDN of the authoritative forward zone.",
	},
	"view": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString("default"),
		MarkdownDescription: "The view to which the zone belongs. If a view is not specified, the default view is used.",
	},
}

func ExpandSharedrecordgroupZoneAssociations(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.SharedrecordgroupZoneAssociations {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m SharedrecordgroupZoneAssociationsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *SharedrecordgroupZoneAssociationsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.SharedrecordgroupZoneAssociations {
	if m == nil {
		return nil
	}
	to := &dns.SharedrecordgroupZoneAssociations{
		Fqdn: flex.ExpandStringPointer(m.Fqdn),
		View: flex.ExpandStringPointer(m.View),
	}
	return to
}

func FlattenSharedrecordgroupZoneAssociations(ctx context.Context, from *dns.SharedrecordgroupZoneAssociations, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SharedrecordgroupZoneAssociationsAttrTypes)
	}
	m := SharedrecordgroupZoneAssociationsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, SharedrecordgroupZoneAssociationsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *SharedrecordgroupZoneAssociationsModel) Flatten(ctx context.Context, from *dns.SharedrecordgroupZoneAssociations, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = SharedrecordgroupZoneAssociationsModel{}
	}
	m.Fqdn = flex.FlattenStringPointer(from.Fqdn)
	m.View = flex.FlattenStringPointer(from.View)
}
