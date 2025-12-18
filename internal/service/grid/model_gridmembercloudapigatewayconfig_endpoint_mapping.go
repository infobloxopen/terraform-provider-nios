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

type GridmembercloudapigatewayconfigEndpointMappingModel struct {
	GatewayFqdn  types.String `tfsdk:"gateway_fqdn"`
	EndpointFqdn types.String `tfsdk:"endpoint_fqdn"`
}

var GridmembercloudapigatewayconfigEndpointMappingAttrTypes = map[string]attr.Type{
	"gateway_fqdn":  types.StringType,
	"endpoint_fqdn": types.StringType,
}

var GridmembercloudapigatewayconfigEndpointMappingResourceSchemaAttributes = map[string]schema.Attribute{
	"gateway_fqdn": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Gateway FQDN.",
	},
	"endpoint_fqdn": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Endpoint FQDN.",
	},
}

func ExpandGridmembercloudapigatewayconfigEndpointMapping(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridmembercloudapigatewayconfigEndpointMapping {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridmembercloudapigatewayconfigEndpointMappingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridmembercloudapigatewayconfigEndpointMappingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridmembercloudapigatewayconfigEndpointMapping {
	if m == nil {
		return nil
	}
	to := &grid.GridmembercloudapigatewayconfigEndpointMapping{
		GatewayFqdn:  flex.ExpandStringPointer(m.GatewayFqdn),
		EndpointFqdn: flex.ExpandStringPointer(m.EndpointFqdn),
	}
	return to
}

func FlattenGridmembercloudapigatewayconfigEndpointMapping(ctx context.Context, from *grid.GridmembercloudapigatewayconfigEndpointMapping, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridmembercloudapigatewayconfigEndpointMappingAttrTypes)
	}
	m := GridmembercloudapigatewayconfigEndpointMappingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridmembercloudapigatewayconfigEndpointMappingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridmembercloudapigatewayconfigEndpointMappingModel) Flatten(ctx context.Context, from *grid.GridmembercloudapigatewayconfigEndpointMapping, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridmembercloudapigatewayconfigEndpointMappingModel{}
	}
	m.GatewayFqdn = flex.FlattenStringPointer(from.GatewayFqdn)
	m.EndpointFqdn = flex.FlattenStringPointer(from.EndpointFqdn)
}
