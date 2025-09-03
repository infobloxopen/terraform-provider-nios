package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type RecordTlsaModel struct {
	Ref              types.String `tfsdk:"ref"`
	CertificateData  types.String `tfsdk:"certificate_data"`
	CertificateUsage types.Int64  `tfsdk:"certificate_usage"`
	CloudInfo        types.Object `tfsdk:"cloud_info"`
	Comment          types.String `tfsdk:"comment"`
	Creator          types.String `tfsdk:"creator"`
	Disable          types.Bool   `tfsdk:"disable"`
	DnsName          types.String `tfsdk:"dns_name"`
	ExtAttrs         types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll      types.Map    `tfsdk:"extattrs_all"`
	LastQueried      types.Int64  `tfsdk:"last_queried"`
	MatchedType      types.Int64  `tfsdk:"matched_type"`
	Name             types.String `tfsdk:"name"`
	Selector         types.Int64  `tfsdk:"selector"`
	Ttl              types.Int64  `tfsdk:"ttl"`
	UseTtl           types.Bool   `tfsdk:"use_ttl"`
	View             types.String `tfsdk:"view"`
	Zone             types.String `tfsdk:"zone"`
}

var RecordTlsaAttrTypes = map[string]attr.Type{
	"ref":               types.StringType,
	"certificate_data":  types.StringType,
	"certificate_usage": types.Int64Type,
	"cloud_info":        types.ObjectType{AttrTypes: RecordTlsaCloudInfoAttrTypes},
	"comment":           types.StringType,
	"creator":           types.StringType,
	"disable":           types.BoolType,
	"dns_name":          types.StringType,
	"extattrs":          types.MapType{ElemType: types.StringType},
	"extattrs_all":      types.MapType{ElemType: types.StringType},
	"last_queried":      types.Int64Type,
	"matched_type":      types.Int64Type,
	"name":              types.StringType,
	"selector":          types.Int64Type,
	"ttl":               types.Int64Type,
	"use_ttl":           types.BoolType,
	"view":              types.StringType,
	"zone":              types.StringType,
}

var RecordTlsaResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"certificate_data": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Hex dump of either raw data for matching type 0, or the hash of the raw data for matching types 1 and 2.",
	},
	"certificate_usage": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "Specifies the provided association that will be used to match the certificate presented in the TLS handshake. Based on RFC-6698.",
	},
	"cloud_info": schema.SingleNestedAttribute{
		Attributes:          RecordTlsaCloudInfoResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "The cloud information associated with the record.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 256),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Comment for the record; maximum 256 characters.",
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
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the record is disabled or not. False means that the record is enabled.",
	},
	"dns_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the TLSA record in punycode format.",
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
	"last_queried": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The time of the last DNS query in Epoch seconds format.",
	},
	"matched_type": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "Specifies how the certificate association is presented. Based on RFC-6698.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.IsValidFQDN(),
		},
		MarkdownDescription: "The TLSA record name in FQDN format. This value can be in unicode format.",
	},
	"selector": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "Specifies which part of the TLS certificate presented by the server will be matched against the association data. Based on RFC-6698.",
	},
	"ttl": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_ttl")),
		},
		MarkdownDescription: "The Time to Live (TTL) value for the record. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached.",
	},
	"use_ttl": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: ttl",
	},
	"view": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		Default:             stringdefault.StaticString("default"),
		MarkdownDescription: "The name of the DNS view in which the record resides. Example: \"external\".",
	},
	"zone": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the zone in which the record resides. Example: \"zone.com\". If a view is not specified when searching by zone, the default view is used.",
	},
}

func (m *RecordTlsaModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.RecordTlsa {
	if m == nil {
		return nil
	}
	to := &dns.RecordTlsa{
		CertificateData:  flex.ExpandStringPointer(m.CertificateData),
		CertificateUsage: flex.ExpandInt64Pointer(m.CertificateUsage),
		Comment:          flex.ExpandStringPointer(m.Comment),
		Creator:          flex.ExpandStringPointer(m.Creator),
		Disable:          flex.ExpandBoolPointer(m.Disable),
		ExtAttrs:         ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		MatchedType:      flex.ExpandInt64Pointer(m.MatchedType),
		Name:             flex.ExpandStringPointer(m.Name),
		Selector:         flex.ExpandInt64Pointer(m.Selector),
		Ttl:              flex.ExpandInt64Pointer(m.Ttl),
		UseTtl:           flex.ExpandBoolPointer(m.UseTtl),
		View:             flex.ExpandStringPointer(m.View),
	}
	return to
}

func FlattenRecordTlsa(ctx context.Context, from *dns.RecordTlsa, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordTlsaAttrTypes)
	}
	m := RecordTlsaModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, RecordTlsaAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RecordTlsaModel) Flatten(ctx context.Context, from *dns.RecordTlsa, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RecordTlsaModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.CertificateData = flex.FlattenStringPointer(from.CertificateData)
	m.CertificateUsage = flex.FlattenInt64Pointer(from.CertificateUsage)
	m.CloudInfo = FlattenRecordTlsaCloudInfo(ctx, from.CloudInfo, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Creator = flex.FlattenStringPointer(from.Creator)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.DnsName = flex.FlattenStringPointer(from.DnsName)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.LastQueried = flex.FlattenInt64Pointer(from.LastQueried)
	m.MatchedType = flex.FlattenInt64Pointer(from.MatchedType)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Selector = flex.FlattenInt64Pointer(from.Selector)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
	m.UseTtl = types.BoolPointerValue(from.UseTtl)
	m.View = flex.FlattenStringPointer(from.View)
	m.Zone = flex.FlattenStringPointer(from.Zone)
}
