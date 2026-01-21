package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/import"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type FilternacModel struct {
	Ref         types.String `tfsdk:"ref"`
    Uuid        types.String `tfsdk:"uuid"`
	Comment     types.String `tfsdk:"comment"`
	Expression  types.String `tfsdk:"expression"`
	ExtAttrs    types.Map    `tfsdk:"extattrs"`
	LeaseTime   types.Int64  `tfsdk:"lease_time"`
	Name        types.String `tfsdk:"name"`
	Options     types.List   `tfsdk:"options"`
	ExtAttrsAll types.Map    `tfsdk:"extattrs_all"`
}

var FilternacAttrTypes = map[string]attr.Type{
	"ref":          types.StringType,
    "uuid":        types.StringType,
	"comment":      types.StringType,
	"expression":   types.StringType,
	"extattrs":     types.MapType{ElemType: types.StringType},
	"lease_time":   types.Int64Type,
	"name":         types.StringType,
	"options":      types.ListType{ElemType: types.ObjectType{AttrTypes: FilternacOptionsAttrTypes}},
	"extattrs_all": types.MapType{ElemType: types.StringType},
}

var FilternacResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
    "uuid": schema.StringAttribute{
        Computed:            true,
        MarkdownDescription: "The uuid to the object.",
    },
	"comment": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		Validators:          []validator.String{stringvalidator.LengthBetween(0, 256)},
		MarkdownDescription: "The descriptive comment of a DHCP NAC Filter object.",
	},
	"expression": schema.StringAttribute{
		Computed:            true,
		Optional:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The conditional expression of a DHCP NAC Filter object.",
	},
	"extattrs": schema.MapAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.",
	},
	"lease_time": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The length of time the DHCP server leases an IP address to a client. The lease time applies to hosts that meet the filter criteria.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		Validators:          []validator.String{customvalidator.ValidateTrimmedString()},
		MarkdownDescription: "The name of a DHCP NAC Filter object.",
	},
	"options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: FilternacOptionsResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Default: listdefault.StaticValue(
			types.ListValueMust(
				types.ObjectType{AttrTypes: FilternacOptionsAttrTypes},
				[]attr.Value{},
			),
		),
		Computed:            true,
		Optional:            true,
		MarkdownDescription: "An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object.",
	},
	"extattrs_all": schema.MapAttribute{
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object, including default attributes.",
		ElementType:         types.StringType,
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
	},
}

func (m *FilternacModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.Filternac {
	if m == nil {
		return nil
	}
	to := &dhcp.Filternac{
		Comment:    flex.ExpandStringPointer(m.Comment),
		Expression: flex.ExpandStringPointer(m.Expression),
		ExtAttrs:   ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		LeaseTime:  flex.ExpandInt64Pointer(m.LeaseTime),
		Name:       flex.ExpandStringPointer(m.Name),
		Options:    flex.ExpandFrameworkListNestedBlock(ctx, m.Options, diags, ExpandFilternacOptions),
	}
	return to
}

func FlattenFilternac(ctx context.Context, from *dhcp.Filternac, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(FilternacAttrTypes)
	}
	m := FilternacModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, FilternacAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *FilternacModel) Flatten(ctx context.Context, from *dhcp.Filternac, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = FilternacModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
    m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Expression = flex.FlattenStringPointer(from.Expression)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.LeaseTime = flex.FlattenInt64Pointer(from.LeaseTime)
	m.Name = flex.FlattenStringPointer(from.Name)
	planOptions := m.Options
	m.Options = flex.FlattenFrameworkListNestedBlock(ctx, from.Options, FilternacOptionsAttrTypes, diags, FlattenFilternacOptions)
	if !planOptions.IsUnknown() {
		reOrderedOptions, diags := utils.ReorderAndFilterDHCPOptions(ctx, planOptions, m.Options)
		if !diags.HasError() {
			m.Options = reOrderedOptions.(basetypes.ListValue)
		}
	}
}
