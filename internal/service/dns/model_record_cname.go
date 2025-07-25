package dns

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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

type RecordCnameModel struct {
	Ref                types.String `tfsdk:"ref"`
	AwsRte53RecordInfo types.Object `tfsdk:"aws_rte53_record_info"`
	Canonical          types.String `tfsdk:"canonical"`
	CloudInfo          types.Object `tfsdk:"cloud_info"`
	Comment            types.String `tfsdk:"comment"`
	CreationTime       types.Int64  `tfsdk:"creation_time"`
	Creator            types.String `tfsdk:"creator"`
	DdnsPrincipal      types.String `tfsdk:"ddns_principal"`
	DdnsProtected      types.Bool   `tfsdk:"ddns_protected"`
	Disable            types.Bool   `tfsdk:"disable"`
	DnsCanonical       types.String `tfsdk:"dns_canonical"`
	DnsName            types.String `tfsdk:"dns_name"`
	ExtAttrs           types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll        types.Map    `tfsdk:"extattrs_all"`
	ForbidReclamation  types.Bool   `tfsdk:"forbid_reclamation"`
	LastQueried        types.Int64  `tfsdk:"last_queried"`
	Name               types.String `tfsdk:"name"`
	Reclaimable        types.Bool   `tfsdk:"reclaimable"`
	SharedRecordGroup  types.String `tfsdk:"shared_record_group"`
	Ttl                types.Int64  `tfsdk:"ttl"`
	UseTtl             types.Bool   `tfsdk:"use_ttl"`
	View               types.String `tfsdk:"view"`
	Zone               types.String `tfsdk:"zone"`
}

var RecordCnameAttrTypes = map[string]attr.Type{
	"ref":                   types.StringType,
	"aws_rte53_record_info": types.ObjectType{AttrTypes: RecordCnameAwsRte53RecordInfoAttrTypes},
	"canonical":             types.StringType,
	"cloud_info":            types.ObjectType{AttrTypes: RecordCnameCloudInfoAttrTypes},
	"comment":               types.StringType,
	"creation_time":         types.Int64Type,
	"creator":               types.StringType,
	"ddns_principal":        types.StringType,
	"ddns_protected":        types.BoolType,
	"disable":               types.BoolType,
	"dns_canonical":         types.StringType,
	"dns_name":              types.StringType,
	"extattrs":              types.MapType{ElemType: types.StringType},
	"extattrs_all":          types.MapType{ElemType: types.StringType},
	"forbid_reclamation":    types.BoolType,
	"last_queried":          types.Int64Type,
	"name":                  types.StringType,
	"reclaimable":           types.BoolType,
	"shared_record_group":   types.StringType,
	"ttl":                   types.Int64Type,
	"use_ttl":               types.BoolType,
	"view":                  types.StringType,
	"zone":                  types.StringType,
}

var RecordCnameResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"aws_rte53_record_info": schema.SingleNestedAttribute{
		Attributes:          RecordCnameAwsRte53RecordInfoResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "The AWS Route53 record information associated with the record.",
	},
	"canonical": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.IsValidFQDN(),
		},
		MarkdownDescription: "Canonical name in FQDN format. This value can be in unicode format.",
	},
	"cloud_info": schema.SingleNestedAttribute{
		Attributes:          RecordCnameCloudInfoResourceSchemaAttributes,
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
		MarkdownDescription: "The record creator.",
	},
	"ddns_principal": schema.StringAttribute{
		Computed:            true,
		Optional:            true,
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
	"dns_canonical": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Canonical name in punycode format.",
	},
	"dns_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name for the CNAME record in punycode format.",
	},
	"extattrs": schema.MapAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object.",
		ElementType:         types.StringType,
		Default:             mapdefault.StaticValue(types.MapNull(types.StringType)),
	},
	"extattrs_all": schema.MapAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object , including default attributes.",
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
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.IsValidFQDN(),
		},
		MarkdownDescription: "The name for a CNAME record in FQDN format. This value can be in unicode format. Regular expression search is not supported for unicode values.",
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
		Computed:            true,
		Optional:            true,
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
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString("default"),
		MarkdownDescription: "The name of the DNS view in which the record resides. Example: \"external\".",
	},
	"zone": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the zone in which the record resides. Example: \"zone.com\". If a view is not specified when searching by zone, the default view is used.",
	},
}

func (m *RecordCnameModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *dns.RecordCname {
	if m == nil {
		return nil
	}
	to := &dns.RecordCname{
		Canonical:         flex.ExpandStringPointer(m.Canonical),
		Comment:           flex.ExpandStringPointer(m.Comment),
		Creator:           flex.ExpandStringPointer(m.Creator),
		DdnsPrincipal:     flex.ExpandStringPointer(m.DdnsPrincipal),
		DdnsProtected:     flex.ExpandBoolPointer(m.DdnsProtected),
		Disable:           flex.ExpandBoolPointer(m.Disable),
		ExtAttrs:          ExpandExtAttr(ctx, m.ExtAttrs, diags),
		ForbidReclamation: flex.ExpandBoolPointer(m.ForbidReclamation),
		Name:              flex.ExpandStringPointer(m.Name),
		Ttl:               flex.ExpandInt64Pointer(m.Ttl),
		UseTtl:            flex.ExpandBoolPointer(m.UseTtl),
	}
	if isCreate {
		to.View = flex.ExpandStringPointer(m.View)
	}
	return to
}

func FlattenRecordCname(ctx context.Context, from *dns.RecordCname, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordCnameAttrTypes)
	}
	m := RecordCnameModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrs = m.ExtAttrsAll
	t, d := types.ObjectValueFrom(ctx, RecordCnameAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RecordCnameModel) Flatten(ctx context.Context, from *dns.RecordCname, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RecordCnameModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.AwsRte53RecordInfo = FlattenRecordCnameAwsRte53RecordInfo(ctx, from.AwsRte53RecordInfo, diags)
	m.Canonical = flex.FlattenStringPointer(from.Canonical)
	m.CloudInfo = FlattenRecordCnameCloudInfo(ctx, from.CloudInfo, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CreationTime = flex.FlattenInt64Pointer(from.CreationTime)
	m.Creator = flex.FlattenStringPointer(from.Creator)
	m.DdnsPrincipal = flex.FlattenStringPointer(from.DdnsPrincipal)
	m.DdnsProtected = types.BoolPointerValue(from.DdnsProtected)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.DnsCanonical = flex.FlattenStringPointer(from.DnsCanonical)
	m.DnsName = flex.FlattenStringPointer(from.DnsName)
	m.ExtAttrsAll = FlattenExtAttr(ctx, from.ExtAttrs, diags)
	m.ForbidReclamation = types.BoolPointerValue(from.ForbidReclamation)
	m.LastQueried = flex.FlattenInt64Pointer(from.LastQueried)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Reclaimable = types.BoolPointerValue(from.Reclaimable)
	m.SharedRecordGroup = flex.FlattenStringPointer(from.SharedRecordGroup)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
	m.UseTtl = types.BoolPointerValue(from.UseTtl)
	m.View = flex.FlattenStringPointer(from.View)
	m.Zone = flex.FlattenStringPointer(from.Zone)
}
