package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
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

type RecordUnknownModel struct {
	Ref                  types.String `tfsdk:"ref"`
	CloudInfo            types.Object `tfsdk:"cloud_info"`
	Comment              types.String `tfsdk:"comment"`
	Creator              types.String `tfsdk:"creator"`
	Disable              types.Bool   `tfsdk:"disable"`
	DisplayRdata         types.String `tfsdk:"display_rdata"`
	DnsName              types.String `tfsdk:"dns_name"`
	EnableHostNamePolicy types.Bool   `tfsdk:"enable_host_name_policy"`
	ExtAttrs             types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll          types.Map    `tfsdk:"extattrs_all"`
	LastQueried          types.Int64  `tfsdk:"last_queried"`
	Name                 types.String `tfsdk:"name"`
	Policy               types.String `tfsdk:"policy"`
	RecordType           types.String `tfsdk:"record_type"`
	SubfieldValues       types.List   `tfsdk:"subfield_values"`
	Ttl                  types.Int64  `tfsdk:"ttl"`
	UseTtl               types.Bool   `tfsdk:"use_ttl"`
	View                 types.String `tfsdk:"view"`
	Zone                 types.String `tfsdk:"zone"`
}

var RecordUnknownAttrTypes = map[string]attr.Type{
	"ref":                     types.StringType,
	"cloud_info":              types.ObjectType{AttrTypes: RecordUnknownCloudInfoAttrTypes},
	"comment":                 types.StringType,
	"creator":                 types.StringType,
	"disable":                 types.BoolType,
	"display_rdata":           types.StringType,
	"dns_name":                types.StringType,
	"enable_host_name_policy": types.BoolType,
	"extattrs":                types.MapType{ElemType: types.StringType},
	"extattrs_all":            types.MapType{ElemType: types.StringType},
	"last_queried":            types.Int64Type,
	"name":                    types.StringType,
	"policy":                  types.StringType,
	"record_type":             types.StringType,
	"subfield_values":         types.ListType{ElemType: types.ObjectType{AttrTypes: RecordUnknownSubfieldValuesAttrTypes}},
	"ttl":                     types.Int64Type,
	"use_ttl":                 types.BoolType,
	"view":                    types.StringType,
	"zone":                    types.StringType,
}

var RecordUnknownResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"cloud_info": schema.SingleNestedAttribute{
		Attributes:          RecordUnknownCloudInfoResourceSchemaAttributes,
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
	"display_rdata": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Standard textual representation of the RDATA.",
	},
	"dns_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the unknown record in punycode format.",
	},
	"enable_host_name_policy": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if host name policy is applicable for the record.",
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
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
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
		MarkdownDescription: "The Unknown record name in FQDN format. This value can be in unicode format.",
	},
	"policy": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The host name policy for the record.",
	},
	"record_type": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Specifies type of unknown resource record.",
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
	},
	"subfield_values": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RecordUnknownSubfieldValuesResourceSchemaAttributes,
		},
		Required: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The list of rdata subfield values of unknown resource record.",
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

func (m *RecordUnknownModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *dns.RecordUnknown {
	if m == nil {
		return nil
	}
	to := &dns.RecordUnknown{
		Comment:              flex.ExpandStringPointer(m.Comment),
		Creator:              flex.ExpandStringPointer(m.Creator),
		Disable:              flex.ExpandBoolPointer(m.Disable),
		EnableHostNamePolicy: flex.ExpandBoolPointer(m.EnableHostNamePolicy),
		ExtAttrs:             ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		Name:                 flex.ExpandStringPointer(m.Name),
		SubfieldValues:       flex.ExpandFrameworkListNestedBlock(ctx, m.SubfieldValues, diags, ExpandRecordUnknownSubfieldValues),
		Ttl:                  flex.ExpandInt64Pointer(m.Ttl),
		UseTtl:               flex.ExpandBoolPointer(m.UseTtl),
		View:                 flex.ExpandStringPointer(m.View),
	}
	if isCreate {
		to.RecordType = flex.ExpandStringPointer(m.RecordType)
	}
	return to
}

func FlattenRecordUnknown(ctx context.Context, from *dns.RecordUnknown, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordUnknownAttrTypes)
	}
	m := RecordUnknownModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, RecordUnknownAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RecordUnknownModel) Flatten(ctx context.Context, from *dns.RecordUnknown, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RecordUnknownModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.CloudInfo = FlattenRecordUnknownCloudInfo(ctx, from.CloudInfo, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Creator = flex.FlattenStringPointer(from.Creator)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.DisplayRdata = flex.FlattenStringPointer(from.DisplayRdata)
	m.DnsName = flex.FlattenStringPointer(from.DnsName)
	m.EnableHostNamePolicy = types.BoolPointerValue(from.EnableHostNamePolicy)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.LastQueried = flex.FlattenInt64Pointer(from.LastQueried)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Policy = flex.FlattenStringPointer(from.Policy)
	m.RecordType = flex.FlattenStringPointer(from.RecordType)
	m.SubfieldValues = flex.FlattenFrameworkListNestedBlock(ctx, from.SubfieldValues, RecordUnknownSubfieldValuesAttrTypes, diags, FlattenRecordUnknownSubfieldValues)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
	m.UseTtl = types.BoolPointerValue(from.UseTtl)
	m.View = flex.FlattenStringPointer(from.View)
	m.Zone = flex.FlattenStringPointer(from.Zone)
}
