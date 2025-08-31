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

type RecordSrvModel struct {
	Ref                types.String `tfsdk:"ref"`
	AwsRte53RecordInfo types.Object `tfsdk:"aws_rte53_record_info"`
	CloudInfo          types.Object `tfsdk:"cloud_info"`
	Comment            types.String `tfsdk:"comment"`
	CreationTime       types.Int64  `tfsdk:"creation_time"`
	Creator            types.String `tfsdk:"creator"`
	DdnsPrincipal      types.String `tfsdk:"ddns_principal"`
	DdnsProtected      types.Bool   `tfsdk:"ddns_protected"`
	Disable            types.Bool   `tfsdk:"disable"`
	DnsName            types.String `tfsdk:"dns_name"`
	DnsTarget          types.String `tfsdk:"dns_target"`
	ExtAttrs           types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll        types.Map    `tfsdk:"extattrs_all"`
	ForbidReclamation  types.Bool   `tfsdk:"forbid_reclamation"`
	LastQueried        types.Int64  `tfsdk:"last_queried"`
	Name               types.String `tfsdk:"name"`
	Port               types.Int64  `tfsdk:"port"`
	Priority           types.Int64  `tfsdk:"priority"`
	Reclaimable        types.Bool   `tfsdk:"reclaimable"`
	SharedRecordGroup  types.String `tfsdk:"shared_record_group"`
	Target             types.String `tfsdk:"target"`
	Ttl                types.Int64  `tfsdk:"ttl"`
	UseTtl             types.Bool   `tfsdk:"use_ttl"`
	View               types.String `tfsdk:"view"`
	Weight             types.Int64  `tfsdk:"weight"`
	Zone               types.String `tfsdk:"zone"`
}

var RecordSrvAttrTypes = map[string]attr.Type{
	"ref":                   types.StringType,
	"aws_rte53_record_info": types.ObjectType{AttrTypes: RecordSrvAwsRte53RecordInfoAttrTypes},
	"cloud_info":            types.ObjectType{AttrTypes: RecordSrvCloudInfoAttrTypes},
	"comment":               types.StringType,
	"creation_time":         types.Int64Type,
	"creator":               types.StringType,
	"ddns_principal":        types.StringType,
	"ddns_protected":        types.BoolType,
	"disable":               types.BoolType,
	"dns_name":              types.StringType,
	"dns_target":            types.StringType,
	"extattrs":              types.MapType{ElemType: types.StringType},
	"extattrs_all":          types.MapType{ElemType: types.StringType},
	"forbid_reclamation":    types.BoolType,
	"last_queried":          types.Int64Type,
	"name":                  types.StringType,
	"port":                  types.Int64Type,
	"priority":              types.Int64Type,
	"reclaimable":           types.BoolType,
	"shared_record_group":   types.StringType,
	"target":                types.StringType,
	"ttl":                   types.Int64Type,
	"use_ttl":               types.BoolType,
	"view":                  types.StringType,
	"weight":                types.Int64Type,
	"zone":                  types.StringType,
}

var RecordSrvResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"aws_rte53_record_info": schema.SingleNestedAttribute{
		Attributes:          RecordSrvAwsRte53RecordInfoResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "The AWS Route53 record information associated with the record.",
	},
	"cloud_info": schema.SingleNestedAttribute{
		Attributes:          RecordSrvCloudInfoResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "The cloud information associated with the record.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
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
	"dns_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name for an SRV record in punycode format.",
	},
	"dns_target": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name for a SRV record in punycode format.",
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
		MarkdownDescription: "A name in FQDN format. This value can be in unicode format.",
	},
	"port": schema.Int64Attribute{
		Required: true,
		Validators: []validator.Int64{
			int64validator.Between(0, 65535),
		},
		MarkdownDescription: "The port of the SRV record. Valid values are from 0 to 65535 (inclusive), in 32-bit unsigned integer format.",
	},
	"priority": schema.Int64Attribute{
		Required: true,
		Validators: []validator.Int64{
			int64validator.Between(0, 65535),
		},
		MarkdownDescription: "The priority of the SRV record. Valid values are from 0 to 65535 (inclusive), in 32-bit unsigned integer format.",
	},
	"reclaimable": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines if the record is reclaimable or not.",
	},
	"shared_record_group": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the shared record group in which the record resides. This field exists only on db_objects if this record is a shared record.",
	},
	"target": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The target of the SRV record in FQDN format. This value can be in unicode format.",
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
	"weight": schema.Int64Attribute{
		Required: true,
		Validators: []validator.Int64{
			int64validator.Between(0, 65535),
		},
		MarkdownDescription: "The weight of the SRV record. Valid values are from 0 to 65535 (inclusive), in 32-bit unsigned integer format.",
	},
	"zone": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the zone in which the record resides. Example: \"zone.com\". If a view is not specified when searching by zone, the default view is used.",
	},
}

func (m *RecordSrvModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *dns.RecordSrv {
	if m == nil {
		return nil
	}
	to := &dns.RecordSrv{
		Comment:           flex.ExpandStringPointer(m.Comment),
		Creator:           flex.ExpandStringPointer(m.Creator),
		DdnsPrincipal:     flex.ExpandStringPointer(m.DdnsPrincipal),
		DdnsProtected:     flex.ExpandBoolPointer(m.DdnsProtected),
		Disable:           flex.ExpandBoolPointer(m.Disable),
		ExtAttrs:          ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		ForbidReclamation: flex.ExpandBoolPointer(m.ForbidReclamation),
		Name:              flex.ExpandStringPointer(m.Name),
		Port:              flex.ExpandInt64Pointer(m.Port),
		Priority:          flex.ExpandInt64Pointer(m.Priority),
		Target:            flex.ExpandStringPointer(m.Target),
		Ttl:               flex.ExpandInt64Pointer(m.Ttl),
		UseTtl:            flex.ExpandBoolPointer(m.UseTtl),
		Weight:            flex.ExpandInt64Pointer(m.Weight),
	}
	if isCreate {
		to.View = flex.ExpandStringPointer(m.View)
	}
	return to
}

func FlattenRecordSrv(ctx context.Context, from *dns.RecordSrv, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordSrvAttrTypes)
	}
	m := RecordSrvModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, RecordSrvAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RecordSrvModel) Flatten(ctx context.Context, from *dns.RecordSrv, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RecordSrvModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.AwsRte53RecordInfo = FlattenRecordSrvAwsRte53RecordInfo(ctx, from.AwsRte53RecordInfo, diags)
	m.CloudInfo = FlattenRecordSrvCloudInfo(ctx, from.CloudInfo, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CreationTime = flex.FlattenInt64Pointer(from.CreationTime)
	m.Creator = flex.FlattenStringPointer(from.Creator)
	m.DdnsPrincipal = flex.FlattenStringPointer(from.DdnsPrincipal)
	m.DdnsProtected = types.BoolPointerValue(from.DdnsProtected)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.DnsName = flex.FlattenStringPointer(from.DnsName)
	m.DnsTarget = flex.FlattenStringPointer(from.DnsTarget)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.ForbidReclamation = types.BoolPointerValue(from.ForbidReclamation)
	m.LastQueried = flex.FlattenInt64Pointer(from.LastQueried)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Port = flex.FlattenInt64Pointer(from.Port)
	m.Priority = flex.FlattenInt64Pointer(from.Priority)
	m.Reclaimable = types.BoolPointerValue(from.Reclaimable)
	m.SharedRecordGroup = flex.FlattenStringPointer(from.SharedRecordGroup)
	m.Target = flex.FlattenStringPointer(from.Target)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
	m.UseTtl = types.BoolPointerValue(from.UseTtl)
	m.View = flex.FlattenStringPointer(from.View)
	m.Weight = flex.FlattenInt64Pointer(from.Weight)
	m.Zone = flex.FlattenStringPointer(from.Zone)
}
