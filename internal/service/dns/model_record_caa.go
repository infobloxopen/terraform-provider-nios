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
	importmod "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/import"
	"github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/plancontrol"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type RecordCaaModel struct {
	Ref               types.String `tfsdk:"ref"`
	CaFlag            types.Int64  `tfsdk:"ca_flag"`
	CaTag             types.String `tfsdk:"ca_tag"`
	CaValue           types.String `tfsdk:"ca_value"`
	CloudInfo         types.Object `tfsdk:"cloud_info"`
	Comment           types.String `tfsdk:"comment"`
	CreationTime      types.Int64  `tfsdk:"creation_time"`
	Creator           types.String `tfsdk:"creator"`
	DdnsPrincipal     types.String `tfsdk:"ddns_principal"`
	DdnsProtected     types.Bool   `tfsdk:"ddns_protected"`
	Disable           types.Bool   `tfsdk:"disable"`
	DnsName           types.String `tfsdk:"dns_name"`
	ExtAttrs          types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll       types.Map    `tfsdk:"extattrs_all"`
	ForbidReclamation types.Bool   `tfsdk:"forbid_reclamation"`
	LastQueried       types.Int64  `tfsdk:"last_queried"`
	Name              types.String `tfsdk:"name"`
	Reclaimable       types.Bool   `tfsdk:"reclaimable"`
	Ttl               types.Int64  `tfsdk:"ttl"`
	UseTtl            types.Bool   `tfsdk:"use_ttl"`
	View              types.String `tfsdk:"view"`
	Zone              types.String `tfsdk:"zone"`
}

var RecordCaaAttrTypes = map[string]attr.Type{
	"ref":                types.StringType,
	"ca_flag":            types.Int64Type,
	"ca_tag":             types.StringType,
	"ca_value":           types.StringType,
	"cloud_info":         types.ObjectType{AttrTypes: RecordCaaCloudInfoAttrTypes},
	"comment":            types.StringType,
	"creation_time":      types.Int64Type,
	"creator":            types.StringType,
	"ddns_principal":     types.StringType,
	"ddns_protected":     types.BoolType,
	"disable":            types.BoolType,
	"dns_name":           types.StringType,
	"extattrs":           types.MapType{ElemType: types.StringType},
	"extattrs_all":       types.MapType{ElemType: types.StringType},
	"forbid_reclamation": types.BoolType,
	"last_queried":       types.Int64Type,
	"name":               types.StringType,
	"reclaimable":        types.BoolType,
	"ttl":                types.Int64Type,
	"use_ttl":            types.BoolType,
	"view":               types.StringType,
	"zone":               types.StringType,
}

var RecordCaaResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"ca_flag": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "Flag of CAA record.",
	},
	"ca_tag": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Tag of CAA record.",
		Validators: []validator.String{
			stringvalidator.LengthBetween(1, 255),
			customvalidator.ValidateTrimmedString(),
		},
	},
	"ca_value": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Value of CAA record",
	},
	"cloud_info": schema.SingleNestedAttribute{
		Attributes:          RecordCaaCloudInfoResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "The cloud information associated with the record.",
		PlanModifiers: []planmodifier.Object{
			plancontrol.UseStateForUnknownObject(),
		},
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
			stringvalidator.LengthBetween(0, 256),
		},
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "Comment for the record; maximum 256 characters.",
	},
	"creation_time": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The creation time of the record.",
		PlanModifiers: []planmodifier.Int64{
			plancontrol.UseStateForUnknownInt64(),
		},
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
	"dns_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the CAA record in punycode format.",
		PlanModifiers: []planmodifier.String{
			plancontrol.UseStateForUnknownString(),
		},
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
	"forbid_reclamation": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the reclamation is allowed for the record or not.",
	},
	"last_queried": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The time of the last DNS query in Epoch seconds format.",
		PlanModifiers: []planmodifier.Int64{
			plancontrol.UseStateForUnknownInt64(),
		},
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.IsValidFQDN(),
		},
		MarkdownDescription: "The CAA record name in FQDN format. This value can be in unicode format.",
	},
	"reclaimable": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines if the record is reclaimable or not.",
		PlanModifiers: []planmodifier.Bool{
			plancontrol.UseStateForUnknownBool(),
		},
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
		Default:  stringdefault.StaticString("default"),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of the DNS view in which the record resides. Example: \"external\".",
	},
	"zone": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the zone in which the record resides. Example: \"zone.com\". If a view is not specified when searching by zone, the default view is used.",
		PlanModifiers: []planmodifier.String{
			plancontrol.UseStateForUnknownString(),
		},
	},
}

func (m *RecordCaaModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.RecordCaa {
	if m == nil {
		return nil
	}
	to := &dns.RecordCaa{
		CaFlag:            flex.ExpandInt64Pointer(m.CaFlag),
		CaTag:             flex.ExpandStringPointer(m.CaTag),
		CaValue:           flex.ExpandStringPointer(m.CaValue),
		Comment:           flex.ExpandStringPointer(m.Comment),
		Creator:           flex.ExpandStringPointer(m.Creator),
		DdnsPrincipal:     flex.ExpandStringPointer(m.DdnsPrincipal),
		DdnsProtected:     flex.ExpandBoolPointer(m.DdnsProtected),
		Disable:           flex.ExpandBoolPointer(m.Disable),
		ExtAttrs:          ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		ForbidReclamation: flex.ExpandBoolPointer(m.ForbidReclamation),
		Name:              flex.ExpandStringPointer(m.Name),
		Ttl:               flex.ExpandInt64Pointer(m.Ttl),
		UseTtl:            flex.ExpandBoolPointer(m.UseTtl),
		View:              flex.ExpandStringPointer(m.View),
	}
	return to
}

func FlattenRecordCaa(ctx context.Context, from *dns.RecordCaa, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordCaaAttrTypes)
	}
	m := RecordCaaModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, RecordCaaAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RecordCaaModel) Flatten(ctx context.Context, from *dns.RecordCaa, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RecordCaaModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.CaFlag = flex.FlattenInt64Pointer(from.CaFlag)
	m.CaTag = flex.FlattenStringPointer(from.CaTag)
	m.CaValue = flex.FlattenStringPointer(from.CaValue)
	m.CloudInfo = FlattenRecordCaaCloudInfo(ctx, from.CloudInfo, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CreationTime = flex.FlattenInt64Pointer(from.CreationTime)
	m.Creator = flex.FlattenStringPointer(from.Creator)
	m.DdnsPrincipal = flex.FlattenStringPointer(from.DdnsPrincipal)
	m.DdnsProtected = types.BoolPointerValue(from.DdnsProtected)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.DnsName = flex.FlattenStringPointer(from.DnsName)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.ForbidReclamation = types.BoolPointerValue(from.ForbidReclamation)
	m.LastQueried = flex.FlattenInt64Pointer(from.LastQueried)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Reclaimable = types.BoolPointerValue(from.Reclaimable)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
	m.UseTtl = types.BoolPointerValue(from.UseTtl)
	m.View = flex.FlattenStringPointer(from.View)
	m.Zone = flex.FlattenStringPointer(from.Zone)
}
