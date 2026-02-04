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

type RecordNaptrModel struct {
	Ref               types.String `tfsdk:"ref"`
	Uuid              types.String `tfsdk:"uuid"`
	CloudInfo         types.Object `tfsdk:"cloud_info"`
	Comment           types.String `tfsdk:"comment"`
	CreationTime      types.Int64  `tfsdk:"creation_time"`
	Creator           types.String `tfsdk:"creator"`
	DdnsPrincipal     types.String `tfsdk:"ddns_principal"`
	DdnsProtected     types.Bool   `tfsdk:"ddns_protected"`
	Disable           types.Bool   `tfsdk:"disable"`
	DnsName           types.String `tfsdk:"dns_name"`
	DnsReplacement    types.String `tfsdk:"dns_replacement"`
	ExtAttrs          types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll       types.Map    `tfsdk:"extattrs_all"`
	Flags             types.String `tfsdk:"flags"`
	ForbidReclamation types.Bool   `tfsdk:"forbid_reclamation"`
	LastQueried       types.Int64  `tfsdk:"last_queried"`
	Name              types.String `tfsdk:"name"`
	Order             types.Int64  `tfsdk:"order"`
	Preference        types.Int64  `tfsdk:"preference"`
	Reclaimable       types.Bool   `tfsdk:"reclaimable"`
	Regexp            types.String `tfsdk:"regexp"`
	Replacement       types.String `tfsdk:"replacement"`
	Services          types.String `tfsdk:"services"`
	Ttl               types.Int64  `tfsdk:"ttl"`
	UseTtl            types.Bool   `tfsdk:"use_ttl"`
	View              types.String `tfsdk:"view"`
	Zone              types.String `tfsdk:"zone"`
}

var RecordNaptrAttrTypes = map[string]attr.Type{
	"ref":                types.StringType,
	"uuid":               types.StringType,
	"cloud_info":         types.ObjectType{AttrTypes: RecordNaptrCloudInfoAttrTypes},
	"comment":            types.StringType,
	"creation_time":      types.Int64Type,
	"creator":            types.StringType,
	"ddns_principal":     types.StringType,
	"ddns_protected":     types.BoolType,
	"disable":            types.BoolType,
	"dns_name":           types.StringType,
	"dns_replacement":    types.StringType,
	"extattrs":           types.MapType{ElemType: types.StringType},
	"extattrs_all":       types.MapType{ElemType: types.StringType},
	"flags":              types.StringType,
	"forbid_reclamation": types.BoolType,
	"last_queried":       types.Int64Type,
	"name":               types.StringType,
	"order":              types.Int64Type,
	"preference":         types.Int64Type,
	"reclaimable":        types.BoolType,
	"regexp":             types.StringType,
	"replacement":        types.StringType,
	"services":           types.StringType,
	"ttl":                types.Int64Type,
	"use_ttl":            types.BoolType,
	"view":               types.StringType,
	"zone":               types.StringType,
}

var RecordNaptrResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The uuid to the object.",
	},
	"cloud_info": schema.SingleNestedAttribute{
		Attributes:          RecordNaptrCloudInfoResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "The cloud information associated with the record.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
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
			stringvalidator.OneOf("DYNAMIC", "STATIC", "SYSTEM"),
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
		MarkdownDescription: "The name of the NAPTR record in punycode format.",
	},
	"dns_replacement": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The replacement field of the NAPTR record in punycode format.",
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
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
	},
	"flags": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
			stringvalidator.OneOf("U", "S", "P", "A", ""),
		},
		MarkdownDescription: "The flags used to control the interpretation of the fields for an NAPTR record object. Supported values for the flags field are \"U\", \"S\", \"P\" and \"A\".",
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
			customvalidator.IsValidDomainName(),
		},
		MarkdownDescription: "The name of the NAPTR record in FQDN format. This value can be in unicode format.",
	},
	"order": schema.Int64Attribute{
		Required: true,
		Validators: []validator.Int64{
			int64validator.Between(0, 65535),
		},
		MarkdownDescription: "The order parameter of the NAPTR records. This parameter specifies the order in which the NAPTR rules are applied when multiple rules are present. Valid values are from 0 to 65535 (inclusive), in 32-bit unsigned integer format.",
	},
	"preference": schema.Int64Attribute{
		Required: true,
		Validators: []validator.Int64{
			int64validator.Between(0, 65535),
		},
		MarkdownDescription: "The preference of the NAPTR record. The preference field determines the order NAPTR records are processed when multiple records with the same order parameter are present. Valid values are from 0 to 65535 (inclusive), in 32-bit unsigned integer format.",
	},
	"reclaimable": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines if the record is reclaimable or not.",
	},
	"regexp": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The regular expression-based rewriting rule of the NAPTR record. This should be a POSIX compliant regular expression, including the substitution rule and flags. Refer to RFC 2915 for the field syntax details.",
	},
	"replacement": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.IsValidDomainName(),
		},
		MarkdownDescription: "The replacement field of the NAPTR record object. For nonterminal NAPTR records, this field specifies the next domain name to look up. This value can be in unicode format.",
	},
	"services": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
			stringvalidator.LengthBetween(0, 128),
		},
		MarkdownDescription: "The services field of the NAPTR record object; maximum 128 characters. The services field contains protocol and service identifiers, such as \"http+E2U\" or \"SIPS+D2T\".",
	},
	"ttl": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_ttl")),
		},
		MarkdownDescription: "The Time to Live (TTL) value for the NAPTR record. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached.",
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
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
	},
	"zone": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the zone in which the record resides. Example: \"zone.com\". If a view is not specified when searching by zone, the default view is used.",
	},
}

func (m *RecordNaptrModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *dns.RecordNaptr {
	if m == nil {
		return nil
	}
	to := &dns.RecordNaptr{
		Comment:           flex.ExpandStringPointer(m.Comment),
		Creator:           flex.ExpandStringPointer(m.Creator),
		DdnsPrincipal:     flex.ExpandStringPointer(m.DdnsPrincipal),
		DdnsProtected:     flex.ExpandBoolPointer(m.DdnsProtected),
		Disable:           flex.ExpandBoolPointer(m.Disable),
		ExtAttrs:          ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		Flags:             flex.ExpandStringPointer(m.Flags),
		ForbidReclamation: flex.ExpandBoolPointer(m.ForbidReclamation),
		Name:              flex.ExpandStringPointer(m.Name),
		Order:             flex.ExpandInt64Pointer(m.Order),
		Preference:        flex.ExpandInt64Pointer(m.Preference),
		Regexp:            flex.ExpandStringPointer(m.Regexp),
		Replacement:       flex.ExpandStringPointer(m.Replacement),
		Services:          flex.ExpandStringPointer(m.Services),
		Ttl:               flex.ExpandInt64Pointer(m.Ttl),
		UseTtl:            flex.ExpandBoolPointer(m.UseTtl),
	}
	if isCreate {
		to.View = flex.ExpandStringPointer(m.View)
	}
	return to
}

func FlattenRecordNaptr(ctx context.Context, from *dns.RecordNaptr, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordNaptrAttrTypes)
	}
	m := RecordNaptrModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, RecordNaptrAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RecordNaptrModel) Flatten(ctx context.Context, from *dns.RecordNaptr, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RecordNaptrModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.CloudInfo = FlattenRecordNaptrCloudInfo(ctx, from.CloudInfo, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CreationTime = flex.FlattenInt64Pointer(from.CreationTime)
	m.Creator = flex.FlattenStringPointer(from.Creator)
	m.DdnsPrincipal = flex.FlattenStringPointer(from.DdnsPrincipal)
	m.DdnsProtected = types.BoolPointerValue(from.DdnsProtected)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.DnsName = flex.FlattenStringPointer(from.DnsName)
	m.DnsReplacement = flex.FlattenStringPointer(from.DnsReplacement)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.Flags = flex.FlattenStringPointer(from.Flags)
	m.ForbidReclamation = types.BoolPointerValue(from.ForbidReclamation)
	m.LastQueried = flex.FlattenInt64Pointer(from.LastQueried)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Order = flex.FlattenInt64Pointer(from.Order)
	m.Preference = flex.FlattenInt64Pointer(from.Preference)
	m.Reclaimable = types.BoolPointerValue(from.Reclaimable)
	m.Regexp = flex.FlattenStringPointer(from.Regexp)
	m.Replacement = flex.FlattenStringPointer(from.Replacement)
	m.Services = flex.FlattenStringPointer(from.Services)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
	m.UseTtl = types.BoolPointerValue(from.UseTtl)
	m.View = flex.FlattenStringPointer(from.View)
	m.Zone = flex.FlattenStringPointer(from.Zone)
}
