package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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

type FilteroptionModel struct {
	Ref          types.String `tfsdk:"ref"`
	ApplyAsClass types.Bool   `tfsdk:"apply_as_class"`
	Bootfile     types.String `tfsdk:"bootfile"`
	Bootserver   types.String `tfsdk:"bootserver"`
	Comment      types.String `tfsdk:"comment"`
	Expression   types.String `tfsdk:"expression"`
	ExtAttrs     types.Map    `tfsdk:"extattrs"`
	LeaseTime    types.Int64  `tfsdk:"lease_time"`
	Name         types.String `tfsdk:"name"`
	NextServer   types.String `tfsdk:"next_server"`
	OptionList   types.List   `tfsdk:"option_list"`
	OptionSpace  types.String `tfsdk:"option_space"`
	PxeLeaseTime types.Int64  `tfsdk:"pxe_lease_time"`
	ExtAttrsAll  types.Map    `tfsdk:"extattrs_all"`
}

var FilteroptionAttrTypes = map[string]attr.Type{
	"ref":            types.StringType,
	"apply_as_class": types.BoolType,
	"bootfile":       types.StringType,
	"bootserver":     types.StringType,
	"comment":        types.StringType,
	"expression":     types.StringType,
	"extattrs":       types.MapType{ElemType: types.StringType},
	"lease_time":     types.Int64Type,
	"name":           types.StringType,
	"next_server":    types.StringType,
	"option_list":    types.ListType{ElemType: types.ObjectType{AttrTypes: FilteroptionOptionListAttrTypes}},
	"option_space":   types.StringType,
	"pxe_lease_time": types.Int64Type,
	"extattrs_all":   types.MapType{ElemType: types.StringType},
}

var FilteroptionResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"apply_as_class": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Determines if apply as class is enabled or not. If this flag is set to \"true\" the filter is treated as global DHCP class, e.g it is written to dhcpd config file even if it is not present in any DHCP range.",
	},
	"bootfile": schema.StringAttribute{
		Computed:            true,
		Optional:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "A name of boot file of a DHCP filter option object.",
	},
	"bootserver": schema.StringAttribute{
		Computed:            true,
		Optional:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "Determines the boot server of a DHCP filter option object. You can specify the name and/or IP address of the boot server that host needs to boot.",
	},
	"comment": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "The descriptive comment of a DHCP filter option object.",
	},
	"expression": schema.StringAttribute{
		Computed:            true,
		Optional:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The conditional expression of a DHCP filter option object.",
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
		MarkdownDescription: "Determines the lease time of a DHCP filter option object.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of a DHCP option filter object.",
	},
	"next_server": schema.StringAttribute{
		Computed:            true,
		Optional:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "Determines the next server of a DHCP filter option object. You can specify the name and/or IP address of the next server that the host needs to boot.",
	},
	"option_list": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: FilteroptionOptionListResourceSchemaAttributes,
		},
		Computed: true,
		Optional: true,
		Default: listdefault.StaticValue(
			types.ListValueMust(
				types.ObjectType{AttrTypes: FilteroptionOptionListAttrTypes},
				[]attr.Value{},
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
		Default:             stringdefault.StaticString("DHCP"),
		MarkdownDescription: "The option space of a DHCP filter option object.",
	},
	"pxe_lease_time": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.Between(0, 2147483647),
		},
		MarkdownDescription: "Determines the PXE (Preboot Execution Environment) lease time of a DHCP filter option object. To specify the duration of time it takes a host to connect to a boot server, such as a TFTP server, and download the file it needs to boot.",
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

func (m *FilteroptionModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.Filteroption {
	if m == nil {
		return nil
	}
	to := &dhcp.Filteroption{
		ApplyAsClass: flex.ExpandBoolPointer(m.ApplyAsClass),
		Bootfile:     flex.ExpandStringPointer(m.Bootfile),
		Bootserver:   flex.ExpandStringPointer(m.Bootserver),
		Comment:      flex.ExpandStringPointer(m.Comment),
		Expression:   flex.ExpandStringPointer(m.Expression),
		ExtAttrs:     ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		LeaseTime:    flex.ExpandInt64Pointer(m.LeaseTime),
		Name:         flex.ExpandStringPointer(m.Name),
		NextServer:   flex.ExpandStringPointer(m.NextServer),
		OptionList:   flex.ExpandFrameworkListNestedBlock(ctx, m.OptionList, diags, ExpandFilteroptionOptionList),
		OptionSpace:  flex.ExpandStringPointer(m.OptionSpace),
		PxeLeaseTime: flex.ExpandInt64Pointer(m.PxeLeaseTime),
	}
	return to
}

func FlattenFilteroption(ctx context.Context, from *dhcp.Filteroption, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(FilteroptionAttrTypes)
	}
	m := FilteroptionModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, FilteroptionAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *FilteroptionModel) Flatten(ctx context.Context, from *dhcp.Filteroption, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = FilteroptionModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.ApplyAsClass = types.BoolPointerValue(from.ApplyAsClass)
	m.Bootfile = flex.FlattenStringPointer(from.Bootfile)
	m.Bootserver = flex.FlattenStringPointer(from.Bootserver)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Expression = flex.FlattenStringPointer(from.Expression)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.LeaseTime = flex.FlattenInt64Pointer(from.LeaseTime)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.NextServer = flex.FlattenStringPointer(from.NextServer)
	planOptions := m.OptionList
	m.OptionList = flex.FlattenFrameworkListNestedBlock(ctx, from.OptionList, FilteroptionOptionListAttrTypes, diags, FlattenFilteroptionOptionList)
	if !planOptions.IsUnknown() {
		reOrderedOptions, diags := utils.ReorderAndFilterDHCPOptions(ctx, planOptions, m.OptionList)
		if !diags.HasError() {
			m.OptionList = reOrderedOptions.(basetypes.ListValue)
		}
	}
	m.OptionSpace = flex.FlattenStringPointer(from.OptionSpace)
	m.PxeLeaseTime = flex.FlattenInt64Pointer(from.PxeLeaseTime)
}
