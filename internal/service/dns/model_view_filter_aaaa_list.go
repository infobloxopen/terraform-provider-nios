package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type ViewFilterAaaaListModel struct {
	Ref        types.String `tfsdk:"ref"`
	Uuid      types.String `tfsdk:"uuid"`
	Address    types.String `tfsdk:"address"`
	Permission types.String `tfsdk:"permission"`
}

var ViewFilterAaaaListAttrTypes = map[string]attr.Type{
	"ref":        types.StringType,
	"uuid":      types.StringType,
	"address":    types.StringType,
	"permission": types.StringType,
}

var ViewFilterAaaaListResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("address")),
			stringvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("permission")),
		},
		MarkdownDescription: "The reference to the Named ACL object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The UUID of the object.",
	},
	"address": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The address this rule applies to or \"Any\".",
	},
	"permission": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("ALLOW", "DENY"),
		},
		MarkdownDescription: "The permission to use for this address.",
	},
}

func ExpandViewFilterAaaaList(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.ViewFilterAaaaList {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ViewFilterAaaaListModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ViewFilterAaaaListModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.ViewFilterAaaaList {
	if m == nil {
		return nil
	}
	to := &dns.ViewFilterAaaaList{
		Ref:        flex.ExpandStringPointer(m.Ref),
		Address:    flex.ExpandStringPointer(m.Address),
		Permission: flex.ExpandStringPointer(m.Permission),
	}
	return to
}

func FlattenViewFilterAaaaList(ctx context.Context, from *dns.ViewFilterAaaaList, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ViewFilterAaaaListAttrTypes)
	}
	m := ViewFilterAaaaListModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ViewFilterAaaaListAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ViewFilterAaaaListModel) Flatten(ctx context.Context, from *dns.ViewFilterAaaaList, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ViewFilterAaaaListModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Permission = flex.FlattenStringPointer(from.Permission)
}
