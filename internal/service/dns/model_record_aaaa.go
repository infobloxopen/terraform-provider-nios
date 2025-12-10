package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	planmodifiers "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/immutable"
	importmod "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/import"
	"github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/plancontrol"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type RecordAaaaModel struct {
	Ref                types.String        `tfsdk:"ref"`
	AwsRte53RecordInfo types.Object        `tfsdk:"aws_rte53_record_info"`
	CloudInfo          types.Object        `tfsdk:"cloud_info"`
	Comment            types.String        `tfsdk:"comment"`
	CreationTime       types.Int64         `tfsdk:"creation_time"`
	Creator            types.String        `tfsdk:"creator"`
	DdnsPrincipal      types.String        `tfsdk:"ddns_principal"`
	DdnsProtected      types.Bool          `tfsdk:"ddns_protected"`
	Disable            types.Bool          `tfsdk:"disable"`
	DiscoveredData     types.Object        `tfsdk:"discovered_data"`
	DnsName            types.String        `tfsdk:"dns_name"`
	ExtAttrs           types.Map           `tfsdk:"extattrs"`
	ExtAttrsAll        types.Map           `tfsdk:"extattrs_all"`
	ForbidReclamation  types.Bool          `tfsdk:"forbid_reclamation"`
	Ipv6addr           iptypes.IPv6Address `tfsdk:"ipv6addr"`
	FuncCall           types.Object        `tfsdk:"func_call"`
	LastQueried        types.Int64         `tfsdk:"last_queried"`
	MsAdUserData       types.Object        `tfsdk:"ms_ad_user_data"`
	Name               types.String        `tfsdk:"name"`
	Reclaimable        types.Bool          `tfsdk:"reclaimable"`
	SharedRecordGroup  types.String        `tfsdk:"shared_record_group"`
	Ttl                types.Int64         `tfsdk:"ttl"`
	UseTtl             types.Bool          `tfsdk:"use_ttl"`
	View               types.String        `tfsdk:"view"`
	Zone               types.String        `tfsdk:"zone"`
}

var RecordAaaaAttrTypes = map[string]attr.Type{
	"ref":                   types.StringType,
	"aws_rte53_record_info": types.ObjectType{AttrTypes: RecordAaaaAwsRte53RecordInfoAttrTypes},
	"cloud_info":            types.ObjectType{AttrTypes: RecordAaaaCloudInfoAttrTypes},
	"comment":               types.StringType,
	"creation_time":         types.Int64Type,
	"creator":               types.StringType,
	"ddns_principal":        types.StringType,
	"ddns_protected":        types.BoolType,
	"disable":               types.BoolType,
	"discovered_data":       types.ObjectType{AttrTypes: RecordAaaaDiscoveredDataAttrTypes},
	"dns_name":              types.StringType,
	"extattrs":              types.MapType{ElemType: types.StringType},
	"extattrs_all":          types.MapType{ElemType: types.StringType},
	"forbid_reclamation":    types.BoolType,
	"ipv6addr":              iptypes.IPv6AddressType{},
	"func_call":             types.ObjectType{AttrTypes: FuncCallAttrTypes},
	"last_queried":          types.Int64Type,
	"ms_ad_user_data":       types.ObjectType{AttrTypes: RecordAaaaMsAdUserDataAttrTypes},
	"name":                  types.StringType,
	"reclaimable":           types.BoolType,
	"shared_record_group":   types.StringType,
	"ttl":                   types.Int64Type,
	"use_ttl":               types.BoolType,
	"view":                  types.StringType,
	"zone":                  types.StringType,
}

var RecordAaaaResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"aws_rte53_record_info": schema.SingleNestedAttribute{
		Attributes:          RecordAaaaAwsRte53RecordInfoResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "The AWS Route53 record information associated with the record.",
		PlanModifiers: []planmodifier.Object{
			plancontrol.UseStateForUnknownObject(),
		},
	},
	"cloud_info": schema.SingleNestedAttribute{
		Attributes:          RecordAaaaCloudInfoResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "The cloud information associated with the record.",
		PlanModifiers: []planmodifier.Object{
			plancontrol.UseStateForUnknownObject(),
		},
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},

		MarkdownDescription: "Comment for the record; maximum 256 characters.",
	},
	"creation_time": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The time of the record creation in Epoch seconds format.",
		PlanModifiers: []planmodifier.Int64{
			plancontrol.UseStateForUnknownInt64(),
		},
	},
	"creator": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Validators: []validator.String{
			stringvalidator.OneOf("STATIC", "DYNAMIC"),
		},
		Default:             stringdefault.StaticString("STATIC"),
		MarkdownDescription: "The record creator. Note that changing creator from or to 'SYSTEM' value is not allowed.",
	},
	"ddns_principal": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The GSS-TSIG principal that owns this record.",
	},
	"ddns_protected": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the DDNS updates for this record are allowed or not.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the record is disabled or not. False means that the record is enabled.",
	},
	"discovered_data": schema.SingleNestedAttribute{
		Attributes:          RecordAaaaDiscoveredDataResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "The discovered data for the record.",
		PlanModifiers: []planmodifier.Object{
			plancontrol.UseStateForUnknownObject(),
		},
	},
	"dns_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name for an AAAA record in punycode format.",
		PlanModifiers: []planmodifier.String{
			plancontrol.UseStateForUnknownString(),
		},
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
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object, including default attributes.",
		ElementType:         types.StringType,
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
	},
	"forbid_reclamation": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the reclamation is allowed for the record or not.",
	},
	"ipv6addr": schema.StringAttribute{
		CustomType: iptypes.IPv6AddressType{},
		Optional:   true,
		Computed:   true,
		Validators: []validator.String{
			stringvalidator.ExactlyOneOf(
				path.MatchRoot("ipv6addr"),
				path.MatchRoot("func_call"),
			),
		},
		MarkdownDescription: "The IPv6 Address of the record. This field is `required` unless a `func_call` is specified to invoke `next_available_ip`.",
	},
	"func_call": schema.SingleNestedAttribute{
		Attributes:          FuncCallResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Specifies the function call to execute. The `next_available_ip` function is supported for Record AAAA.",
	},
	"last_queried": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The time of the last DNS query in Epoch seconds format.",
		PlanModifiers: []planmodifier.Int64{
			plancontrol.UseStateForUnknownInt64(),
		},
	},
	"ms_ad_user_data": schema.SingleNestedAttribute{
		Attributes:          RecordAaaaMsAdUserDataResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "The Microsoft Active Directory user related information.",
		PlanModifiers: []planmodifier.Object{
			plancontrol.UseStateForUnknownObject(),
		},
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Name for the AAAA record in FQDN format. This value can be in unicode format.",
		Validators: []validator.String{
			customvalidator.IsValidFQDN(),
		},
	},
	"reclaimable": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines if the record is reclaimable or not.",
		PlanModifiers: []planmodifier.Bool{
			plancontrol.UseStateForUnknownBool(),
		},
	},
	"shared_record_group": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the shared record group in which the record resides. This field exists only on db_objects if this record is a shared record.",
		PlanModifiers: []planmodifier.String{
			plancontrol.UseStateForUnknownString(),
		},
	},
	"ttl": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The Time To Live (TTL) value for the record. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached.",
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_ttl")),
		},
	},
	"use_ttl": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Flag to indicate whether the TTL value should be used for the AAAA record.",
	},
	"view": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString("default"),
		MarkdownDescription: "The name of the DNS view in which the record resides. Example: \"external\".",
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
	},
	"zone": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the zone in which the record resides. Example: \"zone.com\". If a view is not specified when searching by zone, the default view is used.",
		PlanModifiers: []planmodifier.String{
			plancontrol.UseStateForUnknownString(),
		},
	},
}

func (m *RecordAaaaModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *dns.RecordAaaa {
	if m == nil {
		return nil
	}
	to := &dns.RecordAaaa{
		Comment:           flex.ExpandStringPointer(m.Comment),
		Creator:           flex.ExpandStringPointer(m.Creator),
		DdnsPrincipal:     flex.ExpandStringPointer(m.DdnsPrincipal),
		DdnsProtected:     flex.ExpandBoolPointer(m.DdnsProtected),
		Disable:           flex.ExpandBoolPointer(m.Disable),
		ExtAttrs:          ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		ForbidReclamation: flex.ExpandBoolPointer(m.ForbidReclamation),
		Ipv6addr:          ExpandRecordAaaaIpv6addr(m.Ipv6addr),
		FuncCall:          ExpandFuncCall(ctx, m.FuncCall, diags),
		Name:              flex.ExpandStringPointer(m.Name),
		Ttl:               flex.ExpandInt64Pointer(m.Ttl),
		UseTtl:            flex.ExpandBoolPointer(m.UseTtl),
	}
	if isCreate {
		to.View = flex.ExpandStringPointer(m.View)
	}
	return to
}

func FlattenRecordAaaa(ctx context.Context, from *dns.RecordAaaa, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordAaaaAttrTypes)
	}
	m := RecordAaaaModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, RecordAaaaAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RecordAaaaModel) Flatten(ctx context.Context, from *dns.RecordAaaa, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RecordAaaaModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.AwsRte53RecordInfo = FlattenRecordAaaaAwsRte53RecordInfo(ctx, from.AwsRte53RecordInfo, diags)
	m.CloudInfo = FlattenRecordAaaaCloudInfo(ctx, from.CloudInfo, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CreationTime = flex.FlattenInt64Pointer(from.CreationTime)
	m.Creator = flex.FlattenStringPointer(from.Creator)
	m.DdnsPrincipal = flex.FlattenStringPointer(from.DdnsPrincipal)
	m.DdnsProtected = types.BoolPointerValue(from.DdnsProtected)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.DiscoveredData = FlattenRecordAaaaDiscoveredData(ctx, from.DiscoveredData, diags)
	m.DnsName = flex.FlattenStringPointer(from.DnsName)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.ForbidReclamation = types.BoolPointerValue(from.ForbidReclamation)
	m.Ipv6addr = FlattenRecordAaaaIpv6addr(from.Ipv6addr)
	m.LastQueried = flex.FlattenInt64Pointer(from.LastQueried)
	m.MsAdUserData = FlattenRecordAaaaMsAdUserData(ctx, from.MsAdUserData, diags)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Reclaimable = types.BoolPointerValue(from.Reclaimable)
	m.SharedRecordGroup = flex.FlattenStringPointer(from.SharedRecordGroup)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
	m.UseTtl = types.BoolPointerValue(from.UseTtl)
	m.View = flex.FlattenStringPointer(from.View)
	m.Zone = flex.FlattenStringPointer(from.Zone)

	if m.FuncCall.IsNull() || m.FuncCall.IsUnknown() {
		m.FuncCall = FlattenFuncCall(ctx, from.FuncCall, diags)
	}
}

func ExpandRecordAaaaIpv6addr(str iptypes.IPv6Address) *dns.RecordAaaaIpv6addr {
	if str.IsNull() {
		return &dns.RecordAaaaIpv6addr{}
	}
	var m dns.RecordAaaaIpv6addr
	m.String = flex.ExpandIPv6Address(str)

	return &m
}

func FlattenRecordAaaaIpv6addr(from *dns.RecordAaaaIpv6addr) iptypes.IPv6Address {
	if from.String == nil {
		return iptypes.NewIPv6AddressNull()
	}
	m := flex.FlattenIPv6Address(from.String)
	return m
}
