package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/boolvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type Ipv6rangetemplateModel struct {
	Ref                   types.String `tfsdk:"ref"`
	Uuid                  types.String `tfsdk:"uuid"`
	CloudApiCompatible    types.Bool   `tfsdk:"cloud_api_compatible"`
	Comment               types.String `tfsdk:"comment"`
	DelegatedMember       types.Object `tfsdk:"delegated_member"`
	Exclude               types.List   `tfsdk:"exclude"`
	LogicFilterRules      types.List   `tfsdk:"logic_filter_rules"`
	Member                types.Object `tfsdk:"member"`
	Name                  types.String `tfsdk:"name"`
	NumberOfAddresses     types.Int64  `tfsdk:"number_of_addresses"`
	Offset                types.Int64  `tfsdk:"offset"`
	OptionFilterRules     types.List   `tfsdk:"option_filter_rules"`
	RecycleLeases         types.Bool   `tfsdk:"recycle_leases"`
	ServerAssociationType types.String `tfsdk:"server_association_type"`
	UseLogicFilterRules   types.Bool   `tfsdk:"use_logic_filter_rules"`
	UseRecycleLeases      types.Bool   `tfsdk:"use_recycle_leases"`
}

var Ipv6rangetemplateAttrTypes = map[string]attr.Type{
	"ref":                     types.StringType,
	"uuid":                    types.StringType,
	"cloud_api_compatible":    types.BoolType,
	"comment":                 types.StringType,
	"delegated_member":        types.ObjectType{AttrTypes: Ipv6rangetemplateDelegatedMemberAttrTypes},
	"exclude":                 types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6rangetemplateExcludeAttrTypes}},
	"logic_filter_rules":      types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6rangetemplateLogicFilterRulesAttrTypes}},
	"member":                  types.ObjectType{AttrTypes: Ipv6rangetemplateMemberAttrTypes},
	"name":                    types.StringType,
	"number_of_addresses":     types.Int64Type,
	"offset":                  types.Int64Type,
	"option_filter_rules":     types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6rangetemplateOptionFilterRulesAttrTypes}},
	"recycle_leases":          types.BoolType,
	"server_association_type": types.StringType,
	"use_logic_filter_rules":  types.BoolType,
	"use_recycle_leases":      types.BoolType,
}

var Ipv6rangetemplateResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The uuid to the object.",
	},
	"cloud_api_compatible": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Determines whether the IPv6 DHCP range template can be used to create network objects in a cloud-computing deployment. The cloud_api_compatible attribute must be set to true if any extensible attributes, such as the Terraform Internal ID, require cloud access; otherwise, it must be set to false.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The IPv6 DHCP range template descriptive comment.",
	},
	"delegated_member": schema.SingleNestedAttribute{
		Attributes:          Ipv6rangetemplateDelegatedMemberResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The vConnector member that the object should be delegated to when created from the IPv6 DHCP range template. The vConnector refers to VMware vConnector.",
	},
	"exclude": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6rangetemplateExcludeResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "These are ranges of IPv6 addresses that the appliance does not use to assign to clients. You can use these excluded addresses as static IPv6 addresses. They contain the start and end addresses of the excluded range, and optionally, information about this excluded range.",
	},
	"logic_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6rangetemplateLogicFilterRulesResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
			listvalidator.AlsoRequires(path.MatchRoot("use_logic_filter_rules")),
		},
		Optional:            true,
		MarkdownDescription: "This field contains the logic filters to be applied on this IPv6 range. This list corresponds to the match rules that are written to the DHCPv6 configuration file.",
	},
	"member": schema.SingleNestedAttribute{
		Attributes:          Ipv6rangetemplateMemberResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The member that will provide service for the IPv6 DHCP range. Set `server_association_type` to `MEMBER` if you want the server specified here to serve the range. For searching by this field, use an HTTP method that contains a body (POST or PUT) with MS DHCP server structure and the request should have option _method=GET.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Name of the IPv6 DHCP range template.",
	},
	"number_of_addresses": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "The number of addresses for the IPv6 DHCP range.",
	},
	"offset": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "The start address offset for the IPv6 DHCP range.",
	},
	"option_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6rangetemplateOptionFilterRulesResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "This field contains the Option filters to be applied to this IPv6 range. The appliance uses the matching rules of these filters to select the address range from which it assigns a lease.",
	},
	"recycle_leases": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(true),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_recycle_leases")),
		},
		MarkdownDescription: "Determines whether the leases are kept in Recycle Bin until one week after expiry. If this is set to False, the leases are permanently deleted.",
	},
	"server_association_type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("NONE"),
		Validators: []validator.String{
			stringvalidator.OneOf("MEMBER", "NONE"),
		},
		MarkdownDescription: "The type of server that is going to serve the IPv6 DHCP range. Valid values are `MEMBER` and `NONE`.",
	},
	"use_logic_filter_rules": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: logic_filter_rules",
	},
	"use_recycle_leases": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: recycle_leases",
	},
}

func (m *Ipv6rangetemplateModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.Ipv6rangetemplate {
	if m == nil {
		return nil
	}
	to := &dhcp.Ipv6rangetemplate{
		Ref:                   flex.ExpandStringPointer(m.Ref),
		CloudApiCompatible:    flex.ExpandBoolPointer(m.CloudApiCompatible),
		Comment:               flex.ExpandStringPointer(m.Comment),
		DelegatedMember:       ExpandIpv6rangetemplateDelegatedMember(ctx, m.DelegatedMember, diags),
		Exclude:               flex.ExpandFrameworkListNestedBlock(ctx, m.Exclude, diags, ExpandIpv6rangetemplateExclude),
		LogicFilterRules:      flex.ExpandFrameworkListNestedBlock(ctx, m.LogicFilterRules, diags, ExpandIpv6rangetemplateLogicFilterRules),
		Member:                ExpandIpv6rangetemplateMember(ctx, m.Member, diags),
		Name:                  flex.ExpandStringPointer(m.Name),
		NumberOfAddresses:     flex.ExpandInt64Pointer(m.NumberOfAddresses),
		Offset:                flex.ExpandInt64Pointer(m.Offset),
		OptionFilterRules:     flex.ExpandFrameworkListNestedBlock(ctx, m.OptionFilterRules, diags, ExpandIpv6rangetemplateOptionFilterRules),
		RecycleLeases:         flex.ExpandBoolPointer(m.RecycleLeases),
		ServerAssociationType: flex.ExpandStringPointer(m.ServerAssociationType),
		UseLogicFilterRules:   flex.ExpandBoolPointer(m.UseLogicFilterRules),
		UseRecycleLeases:      flex.ExpandBoolPointer(m.UseRecycleLeases),
	}
	return to
}

func FlattenIpv6rangetemplate(ctx context.Context, from *dhcp.Ipv6rangetemplate, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6rangetemplateAttrTypes)
	}
	m := Ipv6rangetemplateModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6rangetemplateAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Ipv6rangetemplateModel) Flatten(ctx context.Context, from *dhcp.Ipv6rangetemplate, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Ipv6rangetemplateModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.CloudApiCompatible = types.BoolPointerValue(from.CloudApiCompatible)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.DelegatedMember = FlattenIpv6rangetemplateDelegatedMember(ctx, from.DelegatedMember, diags)
	m.Exclude = flex.FlattenFrameworkListNestedBlock(ctx, from.Exclude, Ipv6rangetemplateExcludeAttrTypes, diags, FlattenIpv6rangetemplateExclude)
	m.LogicFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.LogicFilterRules, Ipv6rangetemplateLogicFilterRulesAttrTypes, diags, FlattenIpv6rangetemplateLogicFilterRules)
	m.Member = FlattenIpv6rangetemplateMember(ctx, from.Member, diags)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.NumberOfAddresses = flex.FlattenInt64Pointer(from.NumberOfAddresses)
	m.Offset = flex.FlattenInt64Pointer(from.Offset)
	m.OptionFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.OptionFilterRules, Ipv6rangetemplateOptionFilterRulesAttrTypes, diags, FlattenIpv6rangetemplateOptionFilterRules)
	m.RecycleLeases = types.BoolPointerValue(from.RecycleLeases)
	m.ServerAssociationType = flex.FlattenStringPointer(from.ServerAssociationType)
	m.UseLogicFilterRules = types.BoolPointerValue(from.UseLogicFilterRules)
	m.UseRecycleLeases = types.BoolPointerValue(from.UseRecycleLeases)
}
