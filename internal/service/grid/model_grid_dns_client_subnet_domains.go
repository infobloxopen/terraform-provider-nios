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

type GridDnsClientSubnetDomainsModel struct {
	Domain     types.String `tfsdk:"domain"`
	Permission types.String `tfsdk:"permission"`
}

var GridDnsClientSubnetDomainsAttrTypes = map[string]attr.Type{
	"domain":     types.StringType,
	"permission": types.StringType,
}

var GridDnsClientSubnetDomainsResourceSchemaAttributes = map[string]schema.Attribute{
	"domain": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The FQDN that represents the ECS zone domain name.",
	},
	"permission": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The ECS domain name permission.",
	},
}

func ExpandGridDnsClientSubnetDomains(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridDnsClientSubnetDomains {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridDnsClientSubnetDomainsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridDnsClientSubnetDomainsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridDnsClientSubnetDomains {
	if m == nil {
		return nil
	}
	to := &grid.GridDnsClientSubnetDomains{
		Domain:     flex.ExpandStringPointer(m.Domain),
		Permission: flex.ExpandStringPointer(m.Permission),
	}
	return to
}

func FlattenGridDnsClientSubnetDomains(ctx context.Context, from *grid.GridDnsClientSubnetDomains, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridDnsClientSubnetDomainsAttrTypes)
	}
	m := GridDnsClientSubnetDomainsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridDnsClientSubnetDomainsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridDnsClientSubnetDomainsModel) Flatten(ctx context.Context, from *grid.GridDnsClientSubnetDomains, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridDnsClientSubnetDomainsModel{}
	}
	m.Domain = flex.FlattenStringPointer(from.Domain)
	m.Permission = flex.FlattenStringPointer(from.Permission)
}
