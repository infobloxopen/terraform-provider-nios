package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
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
)

type Ipv6filteroptionModel struct {
	Ref          types.String `tfsdk:"ref"`
	Uuid         types.String `tfsdk:"uuid"`
	ApplyAsClass types.Bool   `tfsdk:"apply_as_class"`
	Comment      types.String `tfsdk:"comment"`
	Expression   types.String `tfsdk:"expression"`
	ExtAttrs     types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll  types.Map    `tfsdk:"extattrs_all"`
	LeaseTime    types.Int64  `tfsdk:"lease_time"`
	Name         types.String `tfsdk:"name"`
	OptionList   types.List   `tfsdk:"option_list"`
	OptionSpace  types.String `tfsdk:"option_space"`
}

var Ipv6filteroptionAttrTypes = map[string]attr.Type{
	"ref":            types.StringType,
	"uuid":           types.StringType,
	"apply_as_class": types.BoolType,
	"comment":        types.StringType,
	"expression":     types.StringType,
	"extattrs":       types.MapType{ElemType: types.StringType},
	"extattrs_all":   types.MapType{ElemType: types.StringType},
	"lease_time":     types.Int64Type,
	"name":           types.StringType,
	"option_list":    types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6filteroptionOptionListAttrTypes}},
	"option_space":   types.StringType,
}

var Ipv6filteroptionResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The uuid to the object.",
	},
	"apply_as_class": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Determines if apply as class is enabled or not. If this flag is set to \"true\" the filter is treated as global DHCP class, e.g it is written to DHCPv6 configuration file even if it is not present in any DHCP range.",
	},
	"comment": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "The descriptive comment of a DHCP IPv6 filter option object.",
	},
	"expression": schema.StringAttribute{
		Computed:            true,
		Optional:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The conditional expression of a DHCP IPv6 filter option object.",
	},
	"extattrs": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Extensible attributes associated with the object.",
	},
	"extattrs_all": schema.MapAttribute{
		Computed: true,
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
		MarkdownDescription: "Extensible attributes associated with the object, including default and internal attributes.",
		ElementType:         types.StringType,
	},
	"lease_time": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Determines the lease time of a DHCP IPv6 filter option object.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of a DHCP IPv6 option filter object.",
	},
	"option_list": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6filteroptionOptionListResourceSchemaAttributes,
		},
		Computed: true,
		Optional: true,
		Default: listdefault.StaticValue(
			types.ListNull(
				types.ObjectType{AttrTypes: Ipv6filteroptionOptionListAttrTypes},
			),
		),
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object.",
	},
	"option_space": schema.StringAttribute{
		Computed:            true,
		Optional:            true,
		Default:             stringdefault.StaticString("DHCPv6"),
		MarkdownDescription: "The option space of a DHCP IPv6 filter option object.",
	},
}

func (m *Ipv6filteroptionModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.Ipv6filteroption {
	if m == nil {
		return nil
	}
	to := &dhcp.Ipv6filteroption{
		ApplyAsClass: flex.ExpandBoolPointer(m.ApplyAsClass),
		Comment:      flex.ExpandStringPointer(m.Comment),
		Expression:   flex.ExpandStringPointer(m.Expression),
		ExtAttrs:     ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		LeaseTime:    flex.ExpandInt64Pointer(m.LeaseTime),
		Name:         flex.ExpandStringPointer(m.Name),
		OptionList:   flex.ExpandFrameworkListNestedBlock(ctx, m.OptionList, diags, ExpandIpv6filteroptionOptionList),
		OptionSpace:  flex.ExpandStringPointer(m.OptionSpace),
	}
	return to
}

func FlattenIpv6filteroption(ctx context.Context, from *dhcp.Ipv6filteroption, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6filteroptionAttrTypes)
	}
	m := Ipv6filteroptionModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, Ipv6filteroptionAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Ipv6filteroptionModel) Flatten(ctx context.Context, from *dhcp.Ipv6filteroption, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Ipv6filteroptionModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.ApplyAsClass = types.BoolPointerValue(from.ApplyAsClass)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Expression = flex.FlattenStringPointer(from.Expression)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.LeaseTime = flex.FlattenInt64Pointer(from.LeaseTime)
	m.Name = flex.FlattenStringPointer(from.Name)
	planoptionList := m.OptionList
	m.OptionList = flex.FlattenFrameworkListNestedBlock(ctx, from.OptionList, Ipv6filteroptionOptionListAttrTypes, diags, FlattenIpv6filteroptionOptionList)
	if !planoptionList.IsUnknown() {
		reOrderedOptionList, diags := utils.ReorderAndFilterDHCPOptions(ctx, planoptionList, m.OptionList)
		if !diags.HasError() {
			m.OptionList = reOrderedOptionList.(basetypes.ListValue)
		}
	}
	m.OptionSpace = flex.FlattenStringPointer(from.OptionSpace)
}
