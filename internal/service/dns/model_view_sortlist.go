package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type ViewSortlistModel struct {
	Address   iptypes.IPAddress `tfsdk:"address"`
	MatchList types.List        `tfsdk:"match_list"`
}

var ViewSortlistAttrTypes = map[string]attr.Type{
	"address":    iptypes.IPAddressType{},
	"match_list": types.ListType{ElemType: types.StringType},
}

var ViewSortlistResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		CustomType:          iptypes.IPAddressType{},
		Required:            true,
		MarkdownDescription: "The source address of a sortlist object.",
	},
	"match_list": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The match list of a sortlist.",
	},
}

func ExpandViewSortlist(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.ViewSortlist {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ViewSortlistModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ViewSortlistModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.ViewSortlist {
	if m == nil {
		return nil
	}
	to := &dns.ViewSortlist{
		Address:   flex.ExpandIPAddress(m.Address),
		MatchList: flex.ExpandFrameworkListString(ctx, m.MatchList, diags),
	}
	return to
}

func FlattenViewSortlist(ctx context.Context, from *dns.ViewSortlist, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ViewSortlistAttrTypes)
	}
	m := ViewSortlistModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ViewSortlistAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ViewSortlistModel) Flatten(ctx context.Context, from *dns.ViewSortlist, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ViewSortlistModel{}
	}
	m.Address = flex.FlattenIPAddress(from.Address)
	m.MatchList = flex.FlattenFrameworkListString(ctx, from.MatchList, diags)
}
