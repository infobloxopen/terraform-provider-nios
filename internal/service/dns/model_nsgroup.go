package dns

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type NsgroupModel struct {
	Ref                 types.String `tfsdk:"ref"`
	Comment             types.String `tfsdk:"comment"`
	ExtAttrs            types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll         types.Map    `tfsdk:"extattrs_all"`
	ExternalPrimaries   types.List   `tfsdk:"external_primaries"`
	ExternalSecondaries types.List   `tfsdk:"external_secondaries"`
	GridPrimary         types.List   `tfsdk:"grid_primary"`
	GridSecondaries     types.List   `tfsdk:"grid_secondaries"`
	IsGridDefault       types.Bool   `tfsdk:"is_grid_default"`
	IsMultimaster       types.Bool   `tfsdk:"is_multimaster"`
	Name                types.String `tfsdk:"name"`
	UseExternalPrimary  types.Bool   `tfsdk:"use_external_primary"`
}

var NsgroupAttrTypes = map[string]attr.Type{
	"ref":                  types.StringType,
	"comment":              types.StringType,
	"extattrs":             types.MapType{ElemType: types.StringType},
	"extattrs_all":         types.MapType{ElemType: types.StringType},
	"external_primaries":   types.ListType{ElemType: types.ObjectType{AttrTypes: NsgroupExternalPrimariesAttrTypes}},
	"external_secondaries": types.ListType{ElemType: types.ObjectType{AttrTypes: NsgroupExternalSecondariesAttrTypes}},
	"grid_primary":         types.ListType{ElemType: types.ObjectType{AttrTypes: NsgroupGridPrimaryAttrTypes}},
	"grid_secondaries":     types.ListType{ElemType: types.ObjectType{AttrTypes: NsgroupGridSecondariesAttrTypes}},
	"is_grid_default":      types.BoolType,
	"is_multimaster":       types.BoolType,
	"name":                 types.StringType,
	"use_external_primary": types.BoolType,
}

var NsgroupResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Should not have leading or trailing whitespace",
			),
		},
		MarkdownDescription: "Comment for the name server group; maximum 256 characters.",
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
	"external_primaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NsgroupExternalPrimariesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_external_primary")),
		},
		MarkdownDescription: "The list of external primary servers.",
	},
	"external_secondaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NsgroupExternalSecondariesResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "The list of external secondary servers.",
	},
	"grid_primary": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NsgroupGridPrimaryResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "The grid primary servers for this group.",
	},
	"grid_secondaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NsgroupGridSecondariesResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "The list with Grid members that are secondary servers for this group.",
	},
	"is_grid_default": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if this name server group is the Grid default.",
	},
	"is_multimaster": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Determines if the \"multiple DNS primaries\" feature is enabled for the group.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Should not have leading or trailing whitespace",
			),
		},
		MarkdownDescription: "The name of this name server group.",
	},
	"use_external_primary": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This flag controls whether the group is using an external primary. Note that modification of this field requires passing values for \"grid_secondaries\" and \"external_primaries\".",
	},
}

func (m *NsgroupModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *dns.Nsgroup {
	if m == nil {
		return nil
	}
	to := &dns.Nsgroup{
		Comment:             flex.ExpandStringPointer(m.Comment),
		ExtAttrs:            ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		ExternalPrimaries:   flex.ExpandFrameworkListNestedBlock(ctx, m.ExternalPrimaries, diags, ExpandNsgroupExternalPrimaries),
		ExternalSecondaries: flex.ExpandFrameworkListNestedBlock(ctx, m.ExternalSecondaries, diags, ExpandNsgroupExternalSecondaries),
		GridPrimary:         flex.ExpandFrameworkListNestedBlock(ctx, m.GridPrimary, diags, ExpandNsgroupGridPrimary),
		GridSecondaries:     flex.ExpandFrameworkListNestedBlock(ctx, m.GridSecondaries, diags, ExpandNsgroupGridSecondaries),
		IsGridDefault:       flex.ExpandBoolPointer(m.IsGridDefault),
		IsMultimaster:       flex.ExpandBoolPointer(m.IsMultimaster),
		Name:                flex.ExpandStringPointer(m.Name),
		UseExternalPrimary:  flex.ExpandBoolPointer(m.UseExternalPrimary),
	}
	if isCreate {
		to.IsMultimaster = flex.ExpandBoolPointer(m.IsMultimaster)
	}
	return to
}

func FlattenNsgroup(ctx context.Context, from *dns.Nsgroup, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NsgroupAttrTypes)
	}
	m := NsgroupModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, NsgroupAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NsgroupModel) Flatten(ctx context.Context, from *dns.Nsgroup, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NsgroupModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.ExternalPrimaries = flex.FlattenFrameworkListNestedBlock(ctx, from.ExternalPrimaries, NsgroupExternalPrimariesAttrTypes, diags, FlattenNsgroupExternalPrimaries)
	m.ExternalSecondaries = flex.FlattenFrameworkListNestedBlock(ctx, from.ExternalSecondaries, NsgroupExternalSecondariesAttrTypes, diags, FlattenNsgroupExternalSecondaries)
	m.GridPrimary = flex.FlattenFrameworkListNestedBlock(ctx, from.GridPrimary, NsgroupGridPrimaryAttrTypes, diags, FlattenNsgroupGridPrimary)
	m.GridSecondaries = flex.FlattenFrameworkListNestedBlock(ctx, from.GridSecondaries, NsgroupGridSecondariesAttrTypes, diags, FlattenNsgroupGridSecondaries)
	m.IsGridDefault = types.BoolPointerValue(from.IsGridDefault)
	m.IsMultimaster = types.BoolPointerValue(from.IsMultimaster)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.UseExternalPrimary = types.BoolPointerValue(from.UseExternalPrimary)
}
