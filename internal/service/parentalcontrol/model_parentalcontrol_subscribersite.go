package parentalcontrol

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/parentalcontrol"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/import"
)

type ParentalcontrolSubscribersiteModel struct {
	Ref                      types.String `tfsdk:"ref"`
	Abss                     types.List   `tfsdk:"abss"`
	ApiMembers               types.List   `tfsdk:"api_members"`
	ApiPort                  types.Int64  `tfsdk:"api_port"`
	BlockSize                types.Int64  `tfsdk:"block_size"`
	BlockingIpv4Vip1         types.String `tfsdk:"blocking_ipv4_vip1"`
	BlockingIpv4Vip2         types.String `tfsdk:"blocking_ipv4_vip2"`
	BlockingIpv6Vip1         types.String `tfsdk:"blocking_ipv6_vip1"`
	BlockingIpv6Vip2         types.String `tfsdk:"blocking_ipv6_vip2"`
	Comment                  types.String `tfsdk:"comment"`
	DcaSubBwList             types.Bool   `tfsdk:"dca_sub_bw_list"`
	DcaSubQueryCount         types.Bool   `tfsdk:"dca_sub_query_count"`
	EnableGlobalAllowListRpz types.Bool   `tfsdk:"enable_global_allow_list_rpz"`
	EnableRpzFilteringBypass types.Bool   `tfsdk:"enable_rpz_filtering_bypass"`
	ExtAttrs                 types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll              types.Map    `tfsdk:"extattrs_all"`
	FirstPort                types.Int64  `tfsdk:"first_port"`
	GlobalAllowListRpz       types.Int64  `tfsdk:"global_allow_list_rpz"`
	MaximumSubscribers       types.Int64  `tfsdk:"maximum_subscribers"`
	Members                  types.List   `tfsdk:"members"`
	Msps                     types.List   `tfsdk:"msps"`
	Name                     types.String `tfsdk:"name"`
	NasGateways              types.List   `tfsdk:"nas_gateways"`
	NasPort                  types.Int64  `tfsdk:"nas_port"`
	ProxyRpzPassthru         types.Bool   `tfsdk:"proxy_rpz_passthru"`
	Spms                     types.List   `tfsdk:"spms"`
	StopAnycast              types.Bool   `tfsdk:"stop_anycast"`
	StrictNat                types.Bool   `tfsdk:"strict_nat"`
	SubscriberCollectionType types.String `tfsdk:"subscriber_collection_type"`
}

var ParentalcontrolSubscribersiteAttrTypes = map[string]attr.Type{
	"ref":                          types.StringType,
	"abss":                         types.ListType{ElemType: types.ObjectType{AttrTypes: ParentalcontrolSubscribersiteAbssAttrTypes}},
	"api_members":                  types.ListType{ElemType: types.ObjectType{AttrTypes: ParentalcontrolSubscribersiteApiMembersAttrTypes}},
	"api_port":                     types.Int64Type,
	"block_size":                   types.Int64Type,
	"blocking_ipv4_vip1":           types.StringType,
	"blocking_ipv4_vip2":           types.StringType,
	"blocking_ipv6_vip1":           types.StringType,
	"blocking_ipv6_vip2":           types.StringType,
	"comment":                      types.StringType,
	"dca_sub_bw_list":              types.BoolType,
	"dca_sub_query_count":          types.BoolType,
	"enable_global_allow_list_rpz": types.BoolType,
	"enable_rpz_filtering_bypass":  types.BoolType,
	"extattrs":                     types.MapType{ElemType: types.StringType},
	"extattrs_all":                 types.MapType{ElemType: types.StringType},
	"first_port":                   types.Int64Type,
	"global_allow_list_rpz":        types.Int64Type,
	"maximum_subscribers":          types.Int64Type,
	"members":                      types.ListType{ElemType: types.ObjectType{AttrTypes: ParentalcontrolSubscribersiteMembersAttrTypes}},
	"msps":                         types.ListType{ElemType: types.ObjectType{AttrTypes: ParentalcontrolSubscribersiteMspsAttrTypes}},
	"name":                         types.StringType,
	"nas_gateways":                 types.ListType{ElemType: types.ObjectType{AttrTypes: ParentalcontrolSubscribersiteNasGatewaysAttrTypes}},
	"nas_port":                     types.Int64Type,
	"proxy_rpz_passthru":           types.BoolType,
	"spms":                         types.ListType{ElemType: types.ObjectType{AttrTypes: ParentalcontrolSubscribersiteSpmsAttrTypes}},
	"stop_anycast":                 types.BoolType,
	"strict_nat":                   types.BoolType,
	"subscriber_collection_type":   types.StringType,
}

var ParentalcontrolSubscribersiteResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"abss": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ParentalcontrolSubscribersiteAbssResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of ABS for the site.",
	},
	"api_members": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ParentalcontrolSubscribersiteApiMembersResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of API members for the site.",
	},
	"api_port": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The port number for gRPC API server.",
	},
	"block_size": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The size of the Deterministic NAT block-size.",
	},
	"blocking_ipv4_vip1": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The IPv4 Address of the blocking server.",
	},
	"blocking_ipv4_vip2": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The IPv4 Address of the blocking server.",
	},
	"blocking_ipv6_vip1": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The IPv6 Address of the blocking server.",
	},
	"blocking_ipv6_vip2": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The IPv6 Address of the blocking server.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The human readable comment for the site.",
	},
	"dca_sub_bw_list": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enable/disable the DCA subscriber B/W list support.",
	},
	"dca_sub_query_count": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enable/disable the DCA subscriber query count.",
	},
	"enable_global_allow_list_rpz": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enable/disable global allow list RPZ setting.",
	},
	"enable_rpz_filtering_bypass": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enable/disable Subscriber Secure Policy Bypass for Allowed list.",
	},
	"extattrs": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.",
	},
	"extattrs_all": schema.MapAttribute{
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object , including default attributes.",
		ElementType:         types.StringType,
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
	},
	"first_port": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The start of the first Deterministic block.",
	},
	"global_allow_list_rpz": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Global allow list RPZ index. Valid values are between 0 and 63.",
	},
	"maximum_subscribers": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The max number of subscribers for the site. It is used to configure the cache size.",
	},
	"members": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ParentalcontrolSubscribersiteMembersResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of members for the site.",
	},
	"msps": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ParentalcontrolSubscribersiteMspsResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of MSP for the site.",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the site.",
	},
	"nas_gateways": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ParentalcontrolSubscribersiteNasGatewaysResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of accounting log servers.",
	},
	"nas_port": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The port number to reach the collector.",
	},
	"proxy_rpz_passthru": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enforce the global proxy list.",
	},
	"spms": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ParentalcontrolSubscribersiteSpmsResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of SPM for the site.",
	},
	"stop_anycast": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Stop the anycast service when the subscriber service is in the interim state.",
	},
	"strict_nat": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Restrict subscriber cache entries to NATed clients.",
	},
	"subscriber_collection_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Subscriber collection type either RADIUS or API.",
	},
}

func ExpandParentalcontrolSubscribersite(ctx context.Context, o types.Object, diags *diag.Diagnostics) *parentalcontrol.ParentalcontrolSubscribersite {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ParentalcontrolSubscribersiteModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ParentalcontrolSubscribersiteModel) Expand(ctx context.Context, diags *diag.Diagnostics) *parentalcontrol.ParentalcontrolSubscribersite {
	if m == nil {
		return nil
	}
	to := &parentalcontrol.ParentalcontrolSubscribersite{
		Ref:                      flex.ExpandStringPointer(m.Ref),
		Abss:                     flex.ExpandFrameworkListNestedBlock(ctx, m.Abss, diags, ExpandParentalcontrolSubscribersiteAbss),
		ApiMembers:               flex.ExpandFrameworkListNestedBlock(ctx, m.ApiMembers, diags, ExpandParentalcontrolSubscribersiteApiMembers),
		BlockSize:                flex.ExpandInt64Pointer(m.BlockSize),
		BlockingIpv4Vip1:         flex.ExpandStringPointer(m.BlockingIpv4Vip1),
		BlockingIpv4Vip2:         flex.ExpandStringPointer(m.BlockingIpv4Vip2),
		BlockingIpv6Vip1:         flex.ExpandStringPointer(m.BlockingIpv6Vip1),
		BlockingIpv6Vip2:         flex.ExpandStringPointer(m.BlockingIpv6Vip2),
		Comment:                  flex.ExpandStringPointer(m.Comment),
		DcaSubBwList:             flex.ExpandBoolPointer(m.DcaSubBwList),
		DcaSubQueryCount:         flex.ExpandBoolPointer(m.DcaSubQueryCount),
		EnableGlobalAllowListRpz: flex.ExpandBoolPointer(m.EnableGlobalAllowListRpz),
		EnableRpzFilteringBypass: flex.ExpandBoolPointer(m.EnableRpzFilteringBypass),
		ExtAttrs:                 ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		FirstPort:                flex.ExpandInt64Pointer(m.FirstPort),
		GlobalAllowListRpz:       flex.ExpandInt64Pointer(m.GlobalAllowListRpz),
		MaximumSubscribers:       flex.ExpandInt64Pointer(m.MaximumSubscribers),
		Members:                  flex.ExpandFrameworkListNestedBlock(ctx, m.Members, diags, ExpandParentalcontrolSubscribersiteMembers),
		Msps:                     flex.ExpandFrameworkListNestedBlock(ctx, m.Msps, diags, ExpandParentalcontrolSubscribersiteMsps),
		Name:                     flex.ExpandStringPointer(m.Name),
		NasGateways:              flex.ExpandFrameworkListNestedBlock(ctx, m.NasGateways, diags, ExpandParentalcontrolSubscribersiteNasGateways),
		NasPort:                  flex.ExpandInt64Pointer(m.NasPort),
		ProxyRpzPassthru:         flex.ExpandBoolPointer(m.ProxyRpzPassthru),
		Spms:                     flex.ExpandFrameworkListNestedBlock(ctx, m.Spms, diags, ExpandParentalcontrolSubscribersiteSpms),
		StopAnycast:              flex.ExpandBoolPointer(m.StopAnycast),
		StrictNat:                flex.ExpandBoolPointer(m.StrictNat),
		SubscriberCollectionType: flex.ExpandStringPointer(m.SubscriberCollectionType),
	}
	return to
}

func FlattenParentalcontrolSubscribersite(ctx context.Context, from *parentalcontrol.ParentalcontrolSubscribersite, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ParentalcontrolSubscribersiteAttrTypes)
	}
	m := ParentalcontrolSubscribersiteModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, ParentalcontrolSubscribersiteAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ParentalcontrolSubscribersiteModel) Flatten(ctx context.Context, from *parentalcontrol.ParentalcontrolSubscribersite, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ParentalcontrolSubscribersiteModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Abss = flex.FlattenFrameworkListNestedBlock(ctx, from.Abss, ParentalcontrolSubscribersiteAbssAttrTypes, diags, FlattenParentalcontrolSubscribersiteAbss)
	m.ApiMembers = flex.FlattenFrameworkListNestedBlock(ctx, from.ApiMembers, ParentalcontrolSubscribersiteApiMembersAttrTypes, diags, FlattenParentalcontrolSubscribersiteApiMembers)
	m.ApiPort = flex.FlattenInt64Pointer(from.ApiPort)
	m.BlockSize = flex.FlattenInt64Pointer(from.BlockSize)
	m.BlockingIpv4Vip1 = flex.FlattenStringPointer(from.BlockingIpv4Vip1)
	m.BlockingIpv4Vip2 = flex.FlattenStringPointer(from.BlockingIpv4Vip2)
	m.BlockingIpv6Vip1 = flex.FlattenStringPointer(from.BlockingIpv6Vip1)
	m.BlockingIpv6Vip2 = flex.FlattenStringPointer(from.BlockingIpv6Vip2)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.DcaSubBwList = types.BoolPointerValue(from.DcaSubBwList)
	m.DcaSubQueryCount = types.BoolPointerValue(from.DcaSubQueryCount)
	m.EnableGlobalAllowListRpz = types.BoolPointerValue(from.EnableGlobalAllowListRpz)
	m.EnableRpzFilteringBypass = types.BoolPointerValue(from.EnableRpzFilteringBypass)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.FirstPort = flex.FlattenInt64Pointer(from.FirstPort)
	m.GlobalAllowListRpz = flex.FlattenInt64Pointer(from.GlobalAllowListRpz)
	m.MaximumSubscribers = flex.FlattenInt64Pointer(from.MaximumSubscribers)
	m.Members = flex.FlattenFrameworkListNestedBlock(ctx, from.Members, ParentalcontrolSubscribersiteMembersAttrTypes, diags, FlattenParentalcontrolSubscribersiteMembers)
	m.Msps = flex.FlattenFrameworkListNestedBlock(ctx, from.Msps, ParentalcontrolSubscribersiteMspsAttrTypes, diags, FlattenParentalcontrolSubscribersiteMsps)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.NasGateways = flex.FlattenFrameworkListNestedBlock(ctx, from.NasGateways, ParentalcontrolSubscribersiteNasGatewaysAttrTypes, diags, FlattenParentalcontrolSubscribersiteNasGateways)
	m.NasPort = flex.FlattenInt64Pointer(from.NasPort)
	m.ProxyRpzPassthru = types.BoolPointerValue(from.ProxyRpzPassthru)
	m.Spms = flex.FlattenFrameworkListNestedBlock(ctx, from.Spms, ParentalcontrolSubscribersiteSpmsAttrTypes, diags, FlattenParentalcontrolSubscribersiteSpms)
	m.StopAnycast = types.BoolPointerValue(from.StopAnycast)
	m.StrictNat = types.BoolPointerValue(from.StrictNat)
	m.SubscriberCollectionType = flex.FlattenStringPointer(from.SubscriberCollectionType)
}
