package dns

import (
	"context"
	"regexp"

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

type RecordMxModel struct {
	Ref                types.String `tfsdk:"ref"`
	AwsRte53RecordInfo types.Object `tfsdk:"aws_rte53_record_info"`
	CloudInfo          types.Object `tfsdk:"cloud_info"`
	Comment            types.String `tfsdk:"comment"`
	CreationTime       types.Int64  `tfsdk:"creation_time"`
	Creator            types.String `tfsdk:"creator"`
	DdnsPrincipal      types.String `tfsdk:"ddns_principal"`
	DdnsProtected      types.Bool   `tfsdk:"ddns_protected"`
	Disable            types.Bool   `tfsdk:"disable"`
	DnsMailExchanger   types.String `tfsdk:"dns_mail_exchanger"`
	DnsName            types.String `tfsdk:"dns_name"`
	ExtAttrs           types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll        types.Map    `tfsdk:"extattrs_all"`
	ForbidReclamation  types.Bool   `tfsdk:"forbid_reclamation"`
	LastQueried        types.Int64  `tfsdk:"last_queried"`
	MailExchanger      types.String `tfsdk:"mail_exchanger"`
	Name               types.String `tfsdk:"name"`
	Preference         types.Int64  `tfsdk:"preference"`
	Reclaimable        types.Bool   `tfsdk:"reclaimable"`
	SharedRecordGroup  types.String `tfsdk:"shared_record_group"`
	Ttl                types.Int64  `tfsdk:"ttl"`
	UseTtl             types.Bool   `tfsdk:"use_ttl"`
	View               types.String `tfsdk:"view"`
	Zone               types.String `tfsdk:"zone"`
}

var RecordMxAttrTypes = map[string]attr.Type{
	"ref":                   types.StringType,
	"aws_rte53_record_info": types.ObjectType{AttrTypes: RecordMxAwsRte53RecordInfoAttrTypes},
	"cloud_info":            types.ObjectType{AttrTypes: RecordMxCloudInfoAttrTypes},
	"comment":               types.StringType,
	"creation_time":         types.Int64Type,
	"creator":               types.StringType,
	"ddns_principal":        types.StringType,
	"ddns_protected":        types.BoolType,
	"disable":               types.BoolType,
	"dns_mail_exchanger":    types.StringType,
	"dns_name":              types.StringType,
	"extattrs":              types.MapType{ElemType: types.StringType},
	"extattrs_all":          types.MapType{ElemType: types.StringType},
	"forbid_reclamation":    types.BoolType,
	"last_queried":          types.Int64Type,
	"mail_exchanger":        types.StringType,
	"name":                  types.StringType,
	"preference":            types.Int64Type,
	"reclaimable":           types.BoolType,
	"shared_record_group":   types.StringType,
	"ttl":                   types.Int64Type,
	"use_ttl":               types.BoolType,
	"view":                  types.StringType,
	"zone":                  types.StringType,
}

var RecordMxResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"aws_rte53_record_info": schema.SingleNestedAttribute{
		Attributes:          RecordMxAwsRte53RecordInfoResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "The AWS Route53 record information associated with the record.",
	},
	"cloud_info": schema.SingleNestedAttribute{
		Attributes:          RecordMxCloudInfoResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "The cloud information associated with the record.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Should not have leading or trailing whitespace",
			),
		},
		Default:             stringdefault.StaticString(""),
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
		Default:             stringdefault.StaticString(""),
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
	"dns_mail_exchanger": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The Mail exchanger name in punycode format.",
	},
	"dns_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name for a MX record in punycode format.",
	},
	"extattrs": schema.MapAttribute{
		Optional:            true,
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.",
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
	"forbid_reclamation": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the reclamation is allowed for the record or not.",
	},
	"last_queried": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The time of the last DNS query in Epoch seconds format.",
	},
	"mail_exchanger": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.IsValidFQDN(),
		},
		MarkdownDescription: "Mail exchanger name in FQDN format. This value can be in unicode format.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.IsValidFQDN(),
		},
		MarkdownDescription: "Name for the MX record in FQDN format. This value can be in unicode format.",
	},
	"preference": schema.Int64Attribute{
		Required: true,
		Validators: []validator.Int64{
			int64validator.Between(0, 65535),
		},
		MarkdownDescription: "Preference value, 0 to 65535 (inclusive) in 32-bit unsigned integer format.",
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
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The Time To Live (TTL) value for record. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached.",
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_ttl")),
		},
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
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Should not have leading or trailing whitespace",
			),
		},
		Default:             stringdefault.StaticString("default"),
		MarkdownDescription: "The name of the DNS view in which the record resides. Example: \"external\".",
	},
	"zone": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the zone in which the record resides. Example: \"zone.com\". If a view is not specified when searching by zone, the default view is used.",
	},
}

func (m *RecordMxModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.RecordMx {
	if m == nil {
		return nil
	}
	to := &dns.RecordMx{
		Comment:           flex.ExpandStringPointer(m.Comment),
		Creator:           flex.ExpandStringPointer(m.Creator),
		DdnsPrincipal:     flex.ExpandStringPointer(m.DdnsPrincipal),
		DdnsProtected:     flex.ExpandBoolPointer(m.DdnsProtected),
		Disable:           flex.ExpandBoolPointer(m.Disable),
		ExtAttrs:          ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		ForbidReclamation: flex.ExpandBoolPointer(m.ForbidReclamation),
		MailExchanger:     flex.ExpandStringPointer(m.MailExchanger),
		Name:              flex.ExpandStringPointer(m.Name),
		Preference:        flex.ExpandInt64Pointer(m.Preference),
		Ttl:               flex.ExpandInt64Pointer(m.Ttl),
		UseTtl:            flex.ExpandBoolPointer(m.UseTtl),
		View:              flex.ExpandStringPointer(m.View),
	}
	return to
}

func FlattenRecordMx(ctx context.Context, from *dns.RecordMx, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordMxAttrTypes)
	}
	m := RecordMxModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RecordMxAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RecordMxModel) Flatten(ctx context.Context, from *dns.RecordMx, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RecordMxModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.AwsRte53RecordInfo = FlattenRecordMxAwsRte53RecordInfo(ctx, from.AwsRte53RecordInfo, diags)
	m.CloudInfo = FlattenRecordMxCloudInfo(ctx, from.CloudInfo, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CreationTime = flex.FlattenInt64Pointer(from.CreationTime)
	m.Creator = flex.FlattenStringPointer(from.Creator)
	m.DdnsPrincipal = flex.FlattenStringPointer(from.DdnsPrincipal)
	m.DdnsProtected = types.BoolPointerValue(from.DdnsProtected)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.DnsMailExchanger = flex.FlattenStringPointer(from.DnsMailExchanger)
	m.DnsName = flex.FlattenStringPointer(from.DnsName)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.ForbidReclamation = types.BoolPointerValue(from.ForbidReclamation)
	m.LastQueried = flex.FlattenInt64Pointer(from.LastQueried)
	m.MailExchanger = flex.FlattenStringPointer(from.MailExchanger)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Preference = flex.FlattenInt64Pointer(from.Preference)
	m.Reclaimable = types.BoolPointerValue(from.Reclaimable)
	m.SharedRecordGroup = flex.FlattenStringPointer(from.SharedRecordGroup)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
	m.UseTtl = types.BoolPointerValue(from.UseTtl)
	m.View = flex.FlattenStringPointer(from.View)
	m.Zone = flex.FlattenStringPointer(from.Zone)
}
