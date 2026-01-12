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

type MemberDnsResolverSettingModel struct {
	Resolvers     types.List `tfsdk:"resolvers"`
	SearchDomains types.List `tfsdk:"search_domains"`
}

var MemberDnsResolverSettingAttrTypes = map[string]attr.Type{
	"resolvers":      types.ListType{ElemType: types.StringType},
	"search_domains": types.ListType{ElemType: types.StringType},
}

var MemberDnsResolverSettingResourceSchemaAttributes = map[string]schema.Attribute{
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

func ExpandMemberDnsResolverSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberDnsResolverSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberDnsResolverSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberDnsResolverSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberDnsResolverSetting {
	if m == nil {
		return nil
	}
	to := &grid.MemberDnsResolverSetting{
		Resolvers:     flex.ExpandFrameworkListString(ctx, m.Resolvers, diags),
		SearchDomains: flex.ExpandFrameworkListString(ctx, m.SearchDomains, diags),
	}
	return to
}

func FlattenMemberDnsResolverSetting(ctx context.Context, from *grid.MemberDnsResolverSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberDnsResolverSettingAttrTypes)
	}
	m := MemberDnsResolverSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberDnsResolverSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberDnsResolverSettingModel) Flatten(ctx context.Context, from *grid.MemberDnsResolverSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberDnsResolverSettingModel{}
	}
	m.Resolvers = flex.FlattenFrameworkListString(ctx, from.Resolvers, diags)
	m.SearchDomains = flex.FlattenFrameworkListString(ctx, from.SearchDomains, diags)
}
