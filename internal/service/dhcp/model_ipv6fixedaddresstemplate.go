package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type Ipv6fixedaddresstemplateModel struct {
	Ref                  types.String `tfsdk:"ref"`
	Comment              types.String `tfsdk:"comment"`
	DomainName           types.String `tfsdk:"domain_name"`
	DomainNameServers    types.List   `tfsdk:"domain_name_servers"`
	ExtAttrs             types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll          types.Map    `tfsdk:"extattrs_all"`
	LogicFilterRules     types.List   `tfsdk:"logic_filter_rules"`
	Name                 types.String `tfsdk:"name"`
	NumberOfAddresses    types.Int64  `tfsdk:"number_of_addresses"`
	Offset               types.Int64  `tfsdk:"offset"`
	Options              types.List   `tfsdk:"options"`
	PreferredLifetime    types.Int64  `tfsdk:"preferred_lifetime"`
	UseDomainName        types.Bool   `tfsdk:"use_domain_name"`
	UseDomainNameServers types.Bool   `tfsdk:"use_domain_name_servers"`
	UseLogicFilterRules  types.Bool   `tfsdk:"use_logic_filter_rules"`
	UseOptions           types.Bool   `tfsdk:"use_options"`
	UsePreferredLifetime types.Bool   `tfsdk:"use_preferred_lifetime"`
	UseValidLifetime     types.Bool   `tfsdk:"use_valid_lifetime"`
	ValidLifetime        types.Int64  `tfsdk:"valid_lifetime"`
}

var Ipv6fixedaddresstemplateAttrTypes = map[string]attr.Type{
	"ref":                     types.StringType,
	"comment":                 types.StringType,
	"domain_name":             types.StringType,
	"domain_name_servers":     types.ListType{ElemType: types.StringType},
	"extattrs":                types.MapType{ElemType: types.StringType},
	"extattrs_all":            types.MapType{ElemType: types.StringType},
	"logic_filter_rules":      types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6fixedaddresstemplateLogicFilterRulesAttrTypes}},
	"name":                    types.StringType,
	"number_of_addresses":     types.Int64Type,
	"offset":                  types.Int64Type,
	"options":                 types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6fixedaddresstemplateOptionsAttrTypes}},
	"preferred_lifetime":      types.Int64Type,
	"use_domain_name":         types.BoolType,
	"use_domain_name_servers": types.BoolType,
	"use_logic_filter_rules":  types.BoolType,
	"use_options":             types.BoolType,
	"use_preferred_lifetime":  types.BoolType,
	"use_valid_lifetime":      types.BoolType,
	"valid_lifetime":          types.Int64Type,
}

var Ipv6fixedaddresstemplateResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "A descriptive comment of an IPv6 fixed address template object.",
	},
	"domain_name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.IsValidDomainName(),
			stringvalidator.AlsoRequires(path.MatchRoot("use_domain_name")),
		},
		MarkdownDescription: "Domain name of the IPv6 fixed address template object.",
	},
	"domain_name_servers": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
		Default:     listdefault.StaticValue(types.ListNull(types.StringType)),
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
			listvalidator.AlsoRequires(path.MatchRoot("use_domain_name_servers")),
		},
		MarkdownDescription: "The IPv6 addresses of DNS recursive name servers to which the DHCP client can send name resolution requests. The DHCP server includes this information in the DNS Recursive Name Server option in Advertise, Rebind, Information-Request, and Reply messages.",
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
	"logic_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6fixedaddresstemplateLogicFilterRulesResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_logic_filter_rules")),
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "This field contains the logic filters to be applied to this IPv6 fixed address. This list corresponds to the match rules that are written to the DHCPv6 configuration file.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Name of an IPv6 fixed address template object.",
	},
	"number_of_addresses": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("offset")),
		},
		MarkdownDescription: "The number of IPv6 addresses for this fixed address.",
	},
	"offset": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("number_of_addresses")),
		},
		MarkdownDescription: "The start address offset for this IPv6 fixed address.",
	},
	"options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6fixedaddresstemplateOptionsResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Default: listdefault.StaticValue(
			types.ListValueMust(
				types.ObjectType{AttrTypes: Ipv6fixedaddresstemplateOptionsAttrTypes},
				[]attr.Value{},
			),
		),
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_options")),
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object.",
	},
	"preferred_lifetime": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_preferred_lifetime")),
		},
		MarkdownDescription: "The preferred lifetime value for this DHCP IPv6 fixed address template object.",
	},
	"use_domain_name": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: domain_name",
	},
	"use_domain_name_servers": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: domain_name_servers",
	},
	"use_logic_filter_rules": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: logic_filter_rules",
	},
	"use_options": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: options",
	},
	"use_preferred_lifetime": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: preferred_lifetime",
	},
	"use_valid_lifetime": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: valid_lifetime",
	},
	"valid_lifetime": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_valid_lifetime")),
		},
		MarkdownDescription: "The valid lifetime value for this DHCP IPv6 fixed address template object.",
	},
}

func (m *Ipv6fixedaddresstemplateModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.Ipv6fixedaddresstemplate {
	if m == nil {
		return nil
	}
	to := &dhcp.Ipv6fixedaddresstemplate{
		Comment:              flex.ExpandStringPointer(m.Comment),
		DomainName:           flex.ExpandStringPointer(m.DomainName),
		DomainNameServers:    flex.ExpandFrameworkListString(ctx, m.DomainNameServers, diags),
		ExtAttrs:             ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		LogicFilterRules:     flex.ExpandFrameworkListNestedBlock(ctx, m.LogicFilterRules, diags, ExpandIpv6fixedaddresstemplateLogicFilterRules),
		Name:                 flex.ExpandStringPointer(m.Name),
		NumberOfAddresses:    flex.ExpandInt64Pointer(m.NumberOfAddresses),
		Offset:               flex.ExpandInt64Pointer(m.Offset),
		Options:              flex.ExpandFrameworkListNestedBlock(ctx, m.Options, diags, ExpandIpv6fixedaddresstemplateOptions),
		PreferredLifetime:    flex.ExpandInt64Pointer(m.PreferredLifetime),
		UseDomainName:        flex.ExpandBoolPointer(m.UseDomainName),
		UseDomainNameServers: flex.ExpandBoolPointer(m.UseDomainNameServers),
		UseLogicFilterRules:  flex.ExpandBoolPointer(m.UseLogicFilterRules),
		UseOptions:           flex.ExpandBoolPointer(m.UseOptions),
		UsePreferredLifetime: flex.ExpandBoolPointer(m.UsePreferredLifetime),
		UseValidLifetime:     flex.ExpandBoolPointer(m.UseValidLifetime),
		ValidLifetime:        flex.ExpandInt64Pointer(m.ValidLifetime),
	}
	return to
}

func FlattenIpv6fixedaddresstemplate(ctx context.Context, from *dhcp.Ipv6fixedaddresstemplate, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6fixedaddresstemplateAttrTypes)
	}
	m := Ipv6fixedaddresstemplateModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, Ipv6fixedaddresstemplateAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Ipv6fixedaddresstemplateModel) Flatten(ctx context.Context, from *dhcp.Ipv6fixedaddresstemplate, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Ipv6fixedaddresstemplateModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.DomainName = flex.FlattenStringPointer(from.DomainName)
	m.DomainNameServers = flex.FlattenFrameworkListString(ctx, from.DomainNameServers, diags)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.LogicFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.LogicFilterRules, Ipv6fixedaddresstemplateLogicFilterRulesAttrTypes, diags, FlattenIpv6fixedaddresstemplateLogicFilterRules)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.NumberOfAddresses = flex.FlattenInt64Pointer(from.NumberOfAddresses)
	m.Offset = flex.FlattenInt64Pointer(from.Offset)
	planOptions := m.Options
	m.Options = flex.FlattenFrameworkListNestedBlock(ctx, from.Options, Ipv6fixedaddresstemplateOptionsAttrTypes, diags, FlattenIpv6fixedaddresstemplateOptions)
	if !planOptions.IsUnknown() {
		reOrderedOptions, diags := utils.ReorderAndFilterDHCPOptions(ctx, planOptions, m.Options)
		if !diags.HasError() {
			m.Options = reOrderedOptions.(basetypes.ListValue)
		}
	}
	m.PreferredLifetime = flex.FlattenInt64Pointer(from.PreferredLifetime)
	m.UseDomainName = types.BoolPointerValue(from.UseDomainName)
	m.UseDomainNameServers = types.BoolPointerValue(from.UseDomainNameServers)
	m.UseLogicFilterRules = types.BoolPointerValue(from.UseLogicFilterRules)
	m.UseOptions = types.BoolPointerValue(from.UseOptions)
	m.UsePreferredLifetime = types.BoolPointerValue(from.UsePreferredLifetime)
	m.UseValidLifetime = types.BoolPointerValue(from.UseValidLifetime)
	m.ValidLifetime = flex.FlattenInt64Pointer(from.ValidLifetime)
}
