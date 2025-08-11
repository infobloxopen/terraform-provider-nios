package dns

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type NsgroupDelegationModel struct {
	Ref         types.String `tfsdk:"ref"`
	Comment     types.String `tfsdk:"comment"`
	DelegateTo  types.List   `tfsdk:"delegate_to"`
	ExtAttrs    types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll types.Map    `tfsdk:"extattrs_all"`
	Name        types.String `tfsdk:"name"`
}

var NsgroupDelegationAttrTypes = map[string]attr.Type{
	"ref":          types.StringType,
	"comment":      types.StringType,
	"delegate_to":  types.ListType{ElemType: types.ObjectType{AttrTypes: NsgroupDelegationDelegateToAttrTypes}},
	"extattrs":     types.MapType{ElemType: types.StringType},
	"extattrs_all": types.MapType{ElemType: types.StringType},
	"name":         types.StringType,
}

var NsgroupDelegationResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Should not have leading or trailing whitespace",
			),
		},
		MarkdownDescription: "The comment for the delegated NS group.",
	},
	"delegate_to": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NsgroupDelegationDelegateToResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "The list of delegated servers for the delegated NS group.",
	},
	"extattrs": schema.MapAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object.",
		ElementType:         types.StringType,
		Default:             mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
	},
	"extattrs_all": schema.MapAttribute{
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object , including default attributes.",
		ElementType:         types.StringType,
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of the delegated NS group.",
	},
}

func ExpandNsgroupDelegation(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.NsgroupDelegation {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NsgroupDelegationModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NsgroupDelegationModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.NsgroupDelegation {
	if m == nil {
		return nil
	}
	to := &dns.NsgroupDelegation{
		Ref:        flex.ExpandStringPointer(m.Ref),
		Comment:    flex.ExpandStringPointer(m.Comment),
		DelegateTo: flex.ExpandFrameworkListNestedBlock(ctx, m.DelegateTo, diags, ExpandNsgroupDelegationDelegateTo),
		ExtAttrs:   ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		Name:       flex.ExpandStringPointer(m.Name),
	}
	return to
}

func FlattenNsgroupDelegation(ctx context.Context, from *dns.NsgroupDelegation, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NsgroupDelegationAttrTypes)
	}
	m := NsgroupDelegationModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, NsgroupDelegationAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NsgroupDelegationModel) Flatten(ctx context.Context, from *dns.NsgroupDelegation, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NsgroupDelegationModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.DelegateTo = flex.FlattenFrameworkListNestedBlock(ctx, from.DelegateTo, NsgroupDelegationDelegateToAttrTypes, diags, FlattenNsgroupDelegationDelegateTo)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.Name = flex.FlattenStringPointer(from.Name)
}
