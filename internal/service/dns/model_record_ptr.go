package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/boolvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int32validator"
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

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dns"

	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/flex"
)

type RecordPtrModel struct {
	Ref                types.String `tfsdk:"ref"`
	AwsRte53RecordInfo types.Object `tfsdk:"aws_rte53_record_info"`
	CloudInfo          types.Object `tfsdk:"cloud_info"`
	Comment            types.String `tfsdk:"comment"`
	CreationTime       types.Int32  `tfsdk:"creation_time"`
	Creator            types.String `tfsdk:"creator"`
	DdnsPrincipal      types.String `tfsdk:"ddns_principal"`
	DdnsProtected      types.Bool   `tfsdk:"ddns_protected"`
	Disable            types.Bool   `tfsdk:"disable"`
	DiscoveredData     types.Object `tfsdk:"discovered_data"`
	DnsName            types.String `tfsdk:"dns_name"`
	DnsPtrdname        types.String `tfsdk:"dns_ptrdname"`
	ExtAttrs           types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll        types.Map    `tfsdk:"extattrs_all"`
	ForbidReclamation  types.Bool   `tfsdk:"forbid_reclamation"`
	Ipv4addr           types.String `tfsdk:"ipv4addr"`
	FuncCall           types.Object `tfsdk:"func_call"`
	Ipv6addr           types.String `tfsdk:"ipv6addr"`
	LastQueried        types.Int32  `tfsdk:"last_queried"`
	MsAdUserData       types.Object `tfsdk:"ms_ad_user_data"`
	Name               types.String `tfsdk:"name"`
	Ptrdname           types.String `tfsdk:"ptrdname"`
	Reclaimable        types.Bool   `tfsdk:"reclaimable"`
	SharedRecordGroup  types.String `tfsdk:"shared_record_group"`
	Ttl                types.Int32  `tfsdk:"ttl"`
	UseTtl             types.Bool   `tfsdk:"use_ttl"`
	View               types.String `tfsdk:"view"`
	Zone               types.String `tfsdk:"zone"`
}

var RecordPtrAttrTypes = map[string]attr.Type{
	"ref":                   types.StringType,
	"aws_rte53_record_info": types.ObjectType{AttrTypes: RecordPtrAwsRte53RecordInfoAttrTypes},
	"cloud_info":            types.ObjectType{AttrTypes: RecordPtrCloudInfoAttrTypes},
	"comment":               types.StringType,
	"creation_time":         types.Int32Type,
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
	"ipv4addr":              types.StringType,
	"func_call":             types.ObjectType{AttrTypes: FuncCallAttrTypes},
	"ipv6addr":              types.StringType,
	"last_queried":          types.Int32Type,
	"ms_ad_user_data":       types.ObjectType{AttrTypes: RecordPtrMsAdUserDataAttrTypes},
	"name":                  types.StringType,
	"ptrdname":              types.StringType,
	"reclaimable":           types.BoolType,
	"shared_record_group":   types.StringType,
	"ttl":                   types.Int32Type,
	"use_ttl":               types.BoolType,
	"view":                  types.StringType,
	"zone":                  types.StringType,
}

var RecordPtrResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"aws_rte53_record_info": schema.SingleNestedAttribute{
		Attributes: RecordPtrAwsRte53RecordInfoResourceSchemaAttributes,
		Computed:   true,
	},
	"cloud_info": schema.SingleNestedAttribute{
		Attributes: RecordPtrCloudInfoResourceSchemaAttributes,
		Computed:   true,
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Comment for the record; maximum 256 characters.",
	},
	"creation_time": schema.Int32Attribute{
		Computed:            true,
		MarkdownDescription: "The time of the record creation in Epoch seconds format.",
	},
	"creator": schema.StringAttribute{
		Optional: true,
		Computed: true,
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
		Attributes: RecordPtrDiscoveredDataResourceSchemaAttributes,
		Computed:   true,
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
		Optional:            true,
		Computed:            true,
		ElementType:         types.StringType,
		Default:             mapdefault.StaticValue(types.MapNull(types.StringType)),
		MarkdownDescription: "Extensible attributes associated with the object.",
	},
	"extattrs_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "Extensible attributes associated with the object , including default attributes.",
	},
	"forbid_reclamation": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the reclamation is allowed for the record or not.",
	},
	"ipv4addr": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.ExactlyOneOf(
				path.MatchRoot("ipv4addr"),
				path.MatchRoot("ipv6addr"),
			),
		},
		MarkdownDescription: "The IPv4 Address of the record.",
	},
	"func_call": schema.SingleNestedAttribute{
		Attributes:          FuncCallResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Function call to be executed.",
	},
	"ipv6addr": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.ExactlyOneOf(
				path.MatchRoot("ipv4addr"),
				path.MatchRoot("ipv6addr"),
			),
		},
		MarkdownDescription: "The IPv6 Address of the record.",
	},
	"last_queried": schema.Int32Attribute{
		Computed:            true,
		MarkdownDescription: "The time of the last DNS query in Epoch seconds format.",
	},
	"ms_ad_user_data": schema.SingleNestedAttribute{
		Computed:            true,
		Attributes:          RecordPtrMsAdUserDataResourceSchemaAttributes,
		MarkdownDescription: "The Microsoft Active Directory user related information.",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The name of the DNS PTR record in FQDN format.",
	},
	"ptrdname": schema.StringAttribute{
		Required:            true,
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
	"ttl": schema.Int32Attribute{
		Optional: true,
		Validators: []validator.Int32{
			int32validator.AlsoRequires(path.MatchRoot("use_ttl")),
		},
		MarkdownDescription: "Time To Live (TTL) value for the record. A 32-bit unsigned integer that represents the duration, in seconds, that the record is valid (cached). Zero indicates that the record should not be cached.",
	},
	"use_ttl": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("ttl")),
		},
		MarkdownDescription: "Flag to indicate whether the TTL value should be used for the A record.",
	},
	"view": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Name of the DNS View in which the record resides, for example \"external\".",
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
		Ref:                flex.ExpandStringPointer(m.Ref),
		AwsRte53RecordInfo: ExpandRecordPtrAwsRte53RecordInfo(ctx, m.AwsRte53RecordInfo, diags),
		CloudInfo:          ExpandRecordPtrCloudInfo(ctx, m.CloudInfo, diags),
		Comment:            flex.ExpandStringPointer(m.Comment),
		Creator:            flex.ExpandStringPointer(m.Creator),
		DdnsPrincipal:      flex.ExpandStringPointer(m.DdnsPrincipal),
		DdnsProtected:      flex.ExpandBoolPointer(m.DdnsProtected),
		Disable:            flex.ExpandBoolPointer(m.Disable),
		DiscoveredData:     ExpandRecordPtrDiscoveredData(ctx, m.DiscoveredData, diags),
		ExtAttrs:           ExpandExtAttr(ctx, m.ExtAttrs, diags),
		ForbidReclamation:  flex.ExpandBoolPointer(m.ForbidReclamation),
		Ipv4addr:           ExpandRecordPtrIpv4addr(m.Ipv4addr),
		FuncCall:           ExpandFuncCall(ctx, m.FuncCall, diags),
		Ipv6addr:           ExpandRecordPtrIpv6addr(m.Ipv6addr),
		MsAdUserData:       ExpandRecordPtrMsAdUserData(ctx, m.MsAdUserData, diags),
		Name:               flex.ExpandStringPointer(m.Name),
		Ptrdname:           flex.ExpandStringPointer(m.Ptrdname),
		Ttl:                flex.ExpandInt32Pointer(m.Ttl),
		UseTtl:             flex.ExpandBoolPointer(m.UseTtl),
	}
	if isCreate {
		to.View = flex.ExpandStringPointer(m.View)
	}
	return to
}

func FlattenRecordPtr(ctx context.Context, from *dns.RecordPtr, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordPtrAttrTypes)
	}
	m := RecordPtrModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrs = m.ExtAttrsAll
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
	m.AwsRte53RecordInfo = FlattenRecordPtrAwsRte53RecordInfo(ctx, from.AwsRte53RecordInfo, diags)
	m.CloudInfo = FlattenRecordPtrCloudInfo(ctx, from.CloudInfo, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CreationTime = flex.FlattenInt32Pointer(from.CreationTime)
	m.Creator = flex.FlattenStringPointer(from.Creator)
	m.DdnsPrincipal = flex.FlattenStringPointer(from.DdnsPrincipal)
	m.DdnsProtected = types.BoolPointerValue(from.DdnsProtected)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.DiscoveredData = FlattenRecordPtrDiscoveredData(ctx, from.DiscoveredData, diags)
	m.DnsName = flex.FlattenStringPointer(from.DnsName)
	m.DnsPtrdname = flex.FlattenStringPointer(from.DnsPtrdname)
	m.ExtAttrsAll = FlattenExtAttr(ctx, from.ExtAttrs, diags)
	m.ForbidReclamation = types.BoolPointerValue(from.ForbidReclamation)
	m.Ipv4addr = FlattenRecordPtrIpv4addr(from.Ipv4addr)
	m.Ipv6addr = FlattenRecordPtrIpv6addr(from.Ipv6addr)
	m.LastQueried = flex.FlattenInt32Pointer(from.LastQueried)
	m.MsAdUserData = FlattenRecordPtrMsAdUserData(ctx, from.MsAdUserData, diags)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Ptrdname = flex.FlattenStringPointer(from.Ptrdname)
	m.Reclaimable = types.BoolPointerValue(from.Reclaimable)
	m.SharedRecordGroup = flex.FlattenStringPointer(from.SharedRecordGroup)
	m.Ttl = flex.FlattenInt32Pointer(from.Ttl)
	m.UseTtl = types.BoolPointerValue(from.UseTtl)
	m.View = flex.FlattenStringPointer(from.View)
	m.Zone = flex.FlattenStringPointer(from.Zone)
	if m.FuncCall.IsNull() || m.FuncCall.IsUnknown() {
		m.FuncCall = FlattenFuncCall(ctx, from.FuncCall, diags)
	}
}

func ExpandRecordPtrIpv4addr(str types.String) *dns.RecordPtrIpv4addr {
	if str.IsNull() || str.IsUnknown() {
		return nil
	}
	return &dns.RecordPtrIpv4addr{
		String: flex.ExpandStringPointer(str),
	}
}

func ExpandRecordPtrIpv6addr(str types.String) *dns.RecordPtrIpv6addr {
	if str.IsNull() || str.IsUnknown() {
		return nil
	}
	return &dns.RecordPtrIpv6addr{
		String: flex.ExpandStringPointer(str),
	}
}

// func ExpandRecordPtrIpv4addr(str types.String) *dns.RecordPtrIpv4addr {
// 	if str.IsNull() {
// 		return &dns.RecordPtrIpv4addr{}
// 	}
// 	var m dns.RecordPtrIpv4addr
// 	m.String = flex.ExpandStringPointer(str)

// 	return &m
// }

// func ExpandRecordPtrIpv6addr(str types.String) *dns.RecordPtrIpv6addr {
// 	if str.IsNull() {
// 		return &dns.RecordPtrIpv6addr{}
// 	}
// 	var m dns.RecordPtrIpv6addr
// 	m.String = flex.ExpandStringPointer(str)

// 	return &m
// }

func FlattenRecordPtrIpv4addr(from *dns.RecordPtrIpv4addr) types.String {
	if from.String == nil {
		return types.StringNull()
	}
	m := types.String{}
	m = flex.FlattenStringPointer(from.String)
	return m
}

func FlattenRecordPtrIpv6addr(from *dns.RecordPtrIpv6addr) types.String {
	if from.String == nil {
		return types.StringNull()
	}
	m := types.String{}
	m = flex.FlattenStringPointer(from.String)
	return m
}
