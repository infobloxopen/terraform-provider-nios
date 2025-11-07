package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type GridMemberCloudapiGatewayConfigModel struct {
	EnableProxyService types.Bool  `tfsdk:"enable_proxy_service"`
	Port               types.Int64 `tfsdk:"port"`
	EndpointMapping    types.List  `tfsdk:"endpoint_mapping"`
}

var GridMemberCloudapiGatewayConfigAttrTypes = map[string]attr.Type{
	"enable_proxy_service": types.BoolType,
	"port":                 types.Int64Type,
	"endpoint_mapping":     types.ListType{ElemType: types.ObjectType{AttrTypes: GridmembercloudapigatewayconfigEndpointMappingAttrTypes}},
}

var GridMemberCloudapiGatewayConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"enable_proxy_service": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enable Gateway Service.",
	},
	"port": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Gateway port",
	},
	"endpoint_mapping": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: GridmembercloudapigatewayconfigEndpointMappingResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "List of Gateway FQDN to AWS Endpoint Mapping.",
	},
}

func ExpandGridMemberCloudapiGatewayConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridMemberCloudapiGatewayConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridMemberCloudapiGatewayConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridMemberCloudapiGatewayConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridMemberCloudapiGatewayConfig {
	if m == nil {
		return nil
	}
	to := &grid.GridMemberCloudapiGatewayConfig{
		EnableProxyService: flex.ExpandBoolPointer(m.EnableProxyService),
		Port:               flex.ExpandInt64Pointer(m.Port),
		EndpointMapping:    flex.ExpandFrameworkListNestedBlock(ctx, m.EndpointMapping, diags, ExpandGridmembercloudapigatewayconfigEndpointMapping),
	}
	return to
}

func FlattenGridMemberCloudapiGatewayConfig(ctx context.Context, from *grid.GridMemberCloudapiGatewayConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridMemberCloudapiGatewayConfigAttrTypes)
	}
	m := GridMemberCloudapiGatewayConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridMemberCloudapiGatewayConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridMemberCloudapiGatewayConfigModel) Flatten(ctx context.Context, from *grid.GridMemberCloudapiGatewayConfig, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridMemberCloudapiGatewayConfigModel{}
	}
	m.EnableProxyService = types.BoolPointerValue(from.EnableProxyService)
	m.Port = flex.FlattenInt64Pointer(from.Port)
	m.EndpointMapping = flex.FlattenFrameworkListNestedBlock(ctx, from.EndpointMapping, GridmembercloudapigatewayconfigEndpointMappingAttrTypes, diags, FlattenGridmembercloudapigatewayconfigEndpointMapping)
}
