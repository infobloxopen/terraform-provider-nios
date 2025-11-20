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
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type RecordPtrModel struct {
	Ref                types.String        `tfsdk:"ref"`
    Uuid        types.String `tfsdk:"uuid"`
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
	DnsPtrdname        types.String        `tfsdk:"dns_ptrdname"`
	ExtAttrs           types.Map           `tfsdk:"extattrs"`
	ExtAttrsAll        types.Map           `tfsdk:"extattrs_all"`
	ForbidReclamation  types.Bool          `tfsdk:"forbid_reclamation"`
	Ipv4addr           iptypes.IPv4Address `tfsdk:"ipv4addr"`
	FuncCall           types.Object        `tfsdk:"func_call"`
	Ipv6addr           iptypes.IPv6Address `tfsdk:"ipv6addr"`
	LastQueried        types.Int64         `tfsdk:"last_queried"`
	MsAdUserData       types.Object        `tfsdk:"ms_ad_user_data"`
	Name               types.String        `tfsdk:"name"`
	Ptrdname           types.String        `tfsdk:"ptrdname"`
	Reclaimable        types.Bool          `tfsdk:"reclaimable"`
	SharedRecordGroup  types.String        `tfsdk:"shared_record_group"`
	Ttl                types.Int64         `tfsdk:"ttl"`
	UseTtl             types.Bool          `tfsdk:"use_ttl"`
	View               types.String        `tfsdk:"view"`
	Zone               types.String        `tfsdk:"zone"`
}

var RecordPtrAttrTypes = map[string]attr.Type{
	"ref":                   types.StringType,
    "uuid":        types.StringType,
	"aws_rte53_record_info": types.ObjectType{AttrTypes: RecordPtrAwsRte53RecordInfoAttrTypes},
	"cloud_info":            types.ObjectType{AttrTypes: RecordPtrCloudInfoAttrTypes},
	"comment":               types.StringType,
	"creation_time":         types.Int64Type,
	"creator":               types.StringType,
	"ddns_principal":        types.StringType,
	"ddns_protected":        types.BoolType,
	"disable":               types.BoolType,
	"discovered_data":       types.ObjectType{AttrTypes: RecordPtrDiscoveredDataAttrTypes},
	"dns_name":              types.StringType,
	"dns_ptrdname":          types.StringType,
	"extattrs":              types.MapType{ElemType: types.StringType},
	"extattrs_all":          types.MapType{ElemType: types.StringType},
	"forbid_reclamation":    types.BoolType,
	"ipv4addr":              iptypes.IPv4AddressType{},
	"func_call":             types.ObjectType{AttrTypes: FuncCallAttrTypes},
	"ipv6addr":              iptypes.IPv6AddressType{},
	"last_queried":          types.Int64Type,
	"ms_ad_user_data":       types.ObjectType{AttrTypes: RecordPtrMsAdUserDataAttrTypes},
	"name":                  types.StringType,
	"ptrdname":              types.StringType,
	"reclaimable":           types.BoolType,
	"shared_record_group":   types.StringType,
	"ttl":                   types.Int64Type,
	"use_ttl":               types.BoolType,
	"view":                  types.StringType,
	"zone":                  types.StringType,
}

var RecordPtrResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
    "uuid": schema.StringAttribute{
        Computed:            true,
        MarkdownDescription: "The uuid to the object.",
    },
	"aws_rte53_record_info": schema.SingleNestedAttribute{
		Attributes:          RecordPtrAwsRte53RecordInfoResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "The AWS Route53 record information associated with the record.",
	},
	"cloud_info": schema.SingleNestedAttribute{
		Attributes:          RecordPtrCloudInfoResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "The cloud information associated with the record.",
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
	},
	"creator": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("STATIC", "DYNAMIC", "SYSTEM"),
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
		Attributes:          RecordPtrDiscoveredDataResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "The discovered data for the record.",
	},
	"dns_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name for a DNS PTR record in punycode format.",
	},
	"dns_ptrdname": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The domain name of the DNS PTR record in punycode format.",
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
		ElementType:         types.StringType,
		MarkdownDescription: "Extensible attributes associated with the object , including default attributes.",
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
	"ipv4addr": schema.StringAttribute{
		CustomType: iptypes.IPv4AddressType{},
		Optional:   true,
		Computed:   true,
		Validators: []validator.String{
			stringvalidator.ExactlyOneOf(
				path.MatchRoot("ipv4addr"),
				path.MatchRoot("ipv6addr"),
				path.MatchRoot("name"),
				path.MatchRoot("func_call"),
			),
		},
		MarkdownDescription: "The IPv4 Address of the record. Either of `ipv4addr`,`ipv6addr`, `name` or `func_call` to invoke `next_available_ip` is required.",
	},
	"func_call": schema.SingleNestedAttribute{
		Attributes:          FuncCallResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Specifies the function call to execute. The `next_available_ip` function is supported for Record PTR.",
	},
	"ipv6addr": schema.StringAttribute{
		CustomType:          iptypes.IPv6AddressType{},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The IPv6 Address of the record. Either of `ipv4addr`,`ipv6addr`, `name` or `func_call` to invoke `next_available_ip` is required.",
	},
	"last_queried": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The time of the last DNS query in Epoch seconds format.",
	},
	"ms_ad_user_data": schema.SingleNestedAttribute{
		Computed:            true,
		Attributes:          RecordPtrMsAdUserDataResourceSchemaAttributes,
		MarkdownDescription: "The Microsoft Active Directory user related information.",
	},
	"name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.Any(
				customvalidator.IsValidArpaIPv4(),
				customvalidator.IsValidArpaIPv6(),
			),
		},
		MarkdownDescription: "The name of the DNS PTR record in FQDN format. Either of `ipv4addr`,`ipv6addr`, `name` or `func_call` to invoke `next_available_ip` is required.",
	},
	"ptrdname": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.IsValidFQDN(),
		},
		MarkdownDescription: "The domain name of the DNS PTR record in FQDN format.",
	},
	"reclaimable": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines if the record is reclaimable or not.",
	},
	"shared_record_group": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the shared record group in which the record resides. This field exists only on db_objects if this record is a shared record.",
	},
	"ttl": schema.Int64Attribute{
		Optional: true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_ttl")),
		},
		MarkdownDescription: "Time To Live (TTL) value for the record. A 32-bit unsigned integer that represents the duration, in seconds, that the record is valid (cached). Zero indicates that the record should not be cached.",
	},
	"use_ttl": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Flag to indicate whether the TTL value should be used for the A record.",
	},
	"view": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("default"),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Name of the DNS View in which the record resides, for example \"external\".",
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
	},
	"zone": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the zone in which the record resides. For example: \"zone.com\". If a view is not specified when searching by zone, the default view is used.",
	},
}

func (m *RecordPtrModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *dns.RecordPtr {
	if m == nil {
		return nil
	}
	to := &dns.RecordPtr{
		Comment:           flex.ExpandStringPointer(m.Comment),
		Creator:           flex.ExpandStringPointer(m.Creator),
		DdnsPrincipal:     flex.ExpandStringPointer(m.DdnsPrincipal),
		DdnsProtected:     flex.ExpandBoolPointer(m.DdnsProtected),
		Disable:           flex.ExpandBoolPointer(m.Disable),
		ExtAttrs:          ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		ForbidReclamation: flex.ExpandBoolPointer(m.ForbidReclamation),
		FuncCall:          ExpandFuncCall(ctx, m.FuncCall, diags),
		Name:              flex.ExpandStringPointer(m.Name),
		Ptrdname:          flex.ExpandStringPointer(m.Ptrdname),
		Ttl:               flex.ExpandInt64Pointer(m.Ttl),
		UseTtl:            flex.ExpandBoolPointer(m.UseTtl),
	}
	if isCreate {
		to.View = flex.ExpandStringPointer(m.View)
	}
	if !m.Ipv4addr.IsUnknown() {
		if m.Ipv4addr.ValueString() != "" {
			to.Ipv4addr = ExpandRecordPtrIpv4addr(m.Ipv4addr)
		}
	} else if !m.Ipv6addr.IsUnknown() {
		if m.Ipv6addr.ValueString() != "" {
			to.Ipv6addr = ExpandRecordPtrIpv6addr(m.Ipv6addr)
		}
	}

	if !m.FuncCall.IsUnknown() && !m.FuncCall.IsNull() {
		if m.FuncCall.Attributes()["attribute_name"].(types.String).ValueString() == "Ipv4addr" {
			to.Ipv4addr = ExpandRecordPtrIpv4addr(m.Ipv4addr)
		} else if m.FuncCall.Attributes()["attribute_name"].(types.String).ValueString() == "Ipv6addr" {
			to.Ipv6addr = ExpandRecordPtrIpv6addr(m.Ipv6addr)
		}
	}
	return to
}

func FlattenRecordPtr(ctx context.Context, from *dns.RecordPtr, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordPtrAttrTypes)
	}
	m := RecordPtrModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, RecordPtrAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RecordPtrModel) Flatten(ctx context.Context, from *dns.RecordPtr, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RecordPtrModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
    m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.AwsRte53RecordInfo = FlattenRecordPtrAwsRte53RecordInfo(ctx, from.AwsRte53RecordInfo, diags)
	m.CloudInfo = FlattenRecordPtrCloudInfo(ctx, from.CloudInfo, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CreationTime = flex.FlattenInt64Pointer(from.CreationTime)
	m.Creator = flex.FlattenStringPointer(from.Creator)
	m.DdnsPrincipal = flex.FlattenStringPointer(from.DdnsPrincipal)
	m.DdnsProtected = types.BoolPointerValue(from.DdnsProtected)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.DiscoveredData = FlattenRecordPtrDiscoveredData(ctx, from.DiscoveredData, diags)
	m.DnsName = flex.FlattenStringPointer(from.DnsName)
	m.DnsPtrdname = flex.FlattenStringPointer(from.DnsPtrdname)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.ForbidReclamation = types.BoolPointerValue(from.ForbidReclamation)
	m.Ipv4addr = FlattenRecordPtrIpv4addr(from.Ipv4addr)
	m.Ipv6addr = FlattenRecordPtrIpv6addr(from.Ipv6addr)
	m.LastQueried = flex.FlattenInt64Pointer(from.LastQueried)
	m.MsAdUserData = FlattenRecordPtrMsAdUserData(ctx, from.MsAdUserData, diags)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Ptrdname = flex.FlattenStringPointer(from.Ptrdname)
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

func ExpandRecordPtrIpv4addr(str iptypes.IPv4Address) *dns.RecordPtrIpv4addr {
	if str.IsNull() {
		return &dns.RecordPtrIpv4addr{}
	}
	var m dns.RecordPtrIpv4addr
	m.String = flex.ExpandIPv4Address(str)

	return &m
}

func ExpandRecordPtrIpv6addr(str iptypes.IPv6Address) *dns.RecordPtrIpv6addr {
	if str.IsNull() {
		return &dns.RecordPtrIpv6addr{}
	}
	var m dns.RecordPtrIpv6addr
	m.String = flex.ExpandIPv6Address(str)

	return &m
}

func FlattenRecordPtrIpv4addr(from *dns.RecordPtrIpv4addr) iptypes.IPv4Address {
	if from.String == nil {
		return iptypes.NewIPv4AddressNull()
	}
	m := flex.FlattenIPv4Address(from.String)
	return m
}

func FlattenRecordPtrIpv6addr(from *dns.RecordPtrIpv6addr) iptypes.IPv6Address {
	if from.String == nil {
		return iptypes.NewIPv6AddressNull()
	}
	m := flex.FlattenIPv6Address(from.String)
	return m
}
