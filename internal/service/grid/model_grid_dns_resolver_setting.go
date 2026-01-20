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

type GridDnsResolverSettingModel struct {
	Resolvers     types.List `tfsdk:"resolvers"`
	SearchDomains types.List `tfsdk:"search_domains"`
}

var GridDnsResolverSettingAttrTypes = map[string]attr.Type{
	"resolvers":      types.ListType{ElemType: types.StringType},
	"search_domains": types.ListType{ElemType: types.StringType},
}

var GridDnsResolverSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"resolvers": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The resolvers of a Grid member. The Grid member sends queries to the first name server address in the list. The second name server address is used if first one does not response.",
	},
	"search_domains": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The Search Domain Group, which is a group of domain names that the Infoblox device can add to partial queries that do not specify a domain name. Note that you can set this parameter only when prefer_resolver or alternate_resolver is set.",
	},
}

func ExpandGridDnsResolverSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridDnsResolverSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridDnsResolverSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridDnsResolverSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridDnsResolverSetting {
	if m == nil {
		return nil
	}
	to := &grid.GridDnsResolverSetting{
		Resolvers:     flex.ExpandFrameworkListString(ctx, m.Resolvers, diags),
		SearchDomains: flex.ExpandFrameworkListString(ctx, m.SearchDomains, diags),
	}
	return to
}

func FlattenGridDnsResolverSetting(ctx context.Context, from *grid.GridDnsResolverSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridDnsResolverSettingAttrTypes)
	}
	m := GridDnsResolverSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridDnsResolverSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridDnsResolverSettingModel) Flatten(ctx context.Context, from *grid.GridDnsResolverSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridDnsResolverSettingModel{}
	}
	m.Resolvers = flex.FlattenFrameworkListString(ctx, from.Resolvers, diags)
	m.SearchDomains = flex.FlattenFrameworkListString(ctx, from.SearchDomains, diags)
}
