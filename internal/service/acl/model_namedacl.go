package acl

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/acl"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type NamedaclModel struct {
	Ref                types.String `tfsdk:"ref"`
	AccessList         types.List   `tfsdk:"access_list"`
	Comment            types.String `tfsdk:"comment"`
	ExplodedAccessList types.List   `tfsdk:"exploded_access_list"`
	ExtAttrs           types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll        types.Map    `tfsdk:"extattrs_all"`
	Name               types.String `tfsdk:"name"`
}

var NamedaclAttrTypes = map[string]attr.Type{
	"ref":                  types.StringType,
	"access_list":          types.ListType{ElemType: types.ObjectType{AttrTypes: NamedaclAccessListAttrTypes}},
	"comment":              types.StringType,
	"exploded_access_list": types.ListType{ElemType: types.ObjectType{AttrTypes: NamedaclExplodedAccessListAttrTypes}},
	"extattrs":             types.MapType{ElemType: types.StringType},
	"extattrs_all":         types.MapType{ElemType: types.StringType},
	"name":                 types.StringType,
}

var NamedaclResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"access_list": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NamedaclAccessListResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The access control list of IPv4/IPv6 addresses, networks, TSIG-based anonymous access controls, and other named ACLs.",
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
		MarkdownDescription: "Comment for the named ACL; maximum 256 characters.",
	},
	"exploded_access_list": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NamedaclExplodedAccessListResourceSchemaAttributes,
		},
		Computed:            true,
		MarkdownDescription: "The exploded access list for the named ACL. This list displays all the access control entries in a named ACL and its nested named ACLs, if applicable.",
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
		MarkdownDescription: "Extensible attributes associated with the object, including default attributes.",
		ElementType:         types.StringType,
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Should not have leading or trailing whitespace",
			),
		},
		MarkdownDescription: "The name of the named ACL.",
	},
}

func (m *NamedaclModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *acl.Namedacl {
	if m == nil {
		return nil
	}
	to := &acl.Namedacl{
		Ref:        flex.ExpandStringPointer(m.Ref),
		AccessList: flex.ExpandFrameworkListNestedBlock(ctx, m.AccessList, diags, ExpandNamedaclAccessList),
		Comment:    flex.ExpandStringPointer(m.Comment),
		ExtAttrs:   ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		Name:       flex.ExpandStringPointer(m.Name),
	}
	return to
}

func FlattenNamedacl(ctx context.Context, from *acl.Namedacl, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NamedaclAttrTypes)
	}
	m := NamedaclModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, NamedaclAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NamedaclModel) Flatten(ctx context.Context, from *acl.Namedacl, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NamedaclModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.AccessList = flex.FlattenFrameworkListNestedBlock(ctx, from.AccessList, NamedaclAccessListAttrTypes, diags, FlattenNamedaclAccessList)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.ExplodedAccessList = flex.FlattenFrameworkListNestedBlock(ctx, from.ExplodedAccessList, NamedaclExplodedAccessListAttrTypes, diags, FlattenNamedaclExplodedAccessList)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.Name = flex.FlattenStringPointer(from.Name)
}
