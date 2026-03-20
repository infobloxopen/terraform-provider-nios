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

type RecordSvcbModel struct {
	Ref                types.String `tfsdk:"ref"`
	Uuid               types.String `tfsdk:"uuid"`
	AwsRte53RecordInfo types.Object `tfsdk:"aws_rte53_record_info"`
	CloudInfo          types.Object `tfsdk:"cloud_info"`
	Comment            types.String `tfsdk:"comment"`
	CreationTime       types.Int64  `tfsdk:"creation_time"`
	Creator            types.String `tfsdk:"creator"`
	DdnsPrincipal      types.String `tfsdk:"ddns_principal"`
	DdnsProtected      types.Bool   `tfsdk:"ddns_protected"`
	Disable            types.Bool   `tfsdk:"disable"`
	ExtAttrs           types.Map    `tfsdk:"extattrs"`
	ForbidReclamation  types.Bool   `tfsdk:"forbid_reclamation"`
	LastQueried        types.Int64  `tfsdk:"last_queried"`
	Name               types.String `tfsdk:"name"`
	Priority           types.Int64  `tfsdk:"priority"`
	Reclaimable        types.Bool   `tfsdk:"reclaimable"`
	SvcParameters      types.List   `tfsdk:"svc_parameters"`
	TargetName         types.String `tfsdk:"target_name"`
	Ttl                types.Int64  `tfsdk:"ttl"`
	UseTtl             types.Bool   `tfsdk:"use_ttl"`
	View               types.String `tfsdk:"view"`
	Zone               types.String `tfsdk:"zone"`
	ExtAttrsAll        types.Map    `tfsdk:"extattrs_all"`
}

var RecordSvcbAttrTypes = map[string]attr.Type{
	"ref":                   types.StringType,
	"uuid":                  types.StringType,
	"aws_rte53_record_info": types.ObjectType{AttrTypes: RecordSvcbAwsRte53RecordInfoAttrTypes},
	"cloud_info":            types.ObjectType{AttrTypes: RecordSvcbCloudInfoAttrTypes},
	"comment":               types.StringType,
	"creation_time":         types.Int64Type,
	"creator":               types.StringType,
	"ddns_principal":        types.StringType,
	"ddns_protected":        types.BoolType,
	"disable":               types.BoolType,
	"extattrs":              types.MapType{ElemType: types.StringType},
	"forbid_reclamation":    types.BoolType,
	"last_queried":          types.Int64Type,
	"name":                  types.StringType,
	"priority":              types.Int64Type,
	"reclaimable":           types.BoolType,
	"svc_parameters":        types.ListType{ElemType: types.ObjectType{AttrTypes: RecordSvcbSvcParametersAttrTypes}},
	"target_name":           types.StringType,
	"ttl":                   types.Int64Type,
	"use_ttl":               types.BoolType,
	"view":                  types.StringType,
	"zone":                  types.StringType,
	"extattrs_all":          types.MapType{ElemType: types.StringType},
}

var RecordSvcbResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Universally Unique ID assigned for this object",
	},
	"aws_rte53_record_info": schema.SingleNestedAttribute{
		Attributes: RecordSvcbAwsRte53RecordInfoResourceSchemaAttributes,
		Computed:   true,
	},
	"cloud_info": schema.SingleNestedAttribute{
		Attributes: RecordSvcbCloudInfoResourceSchemaAttributes,
		Computed:   true,
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
	"creation_time": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The time of the record creation in Epoch seconds format.",
	},
	"creator": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("STATIC"),
		Validators: []validator.String{
			stringvalidator.OneOf("STATIC", "DYNAMIC", "SYSTEM"),
		},
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
		MarkdownDescription: "Determines if the record is disabled or not.False means that the record is enabled.",
	},
	"extattrs": schema.MapAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.",
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
			customvalidator.ValidateTrimmedString(),
			customvalidator.IsValidDomainName(),
		},
		MarkdownDescription: "The name for a SVCB record in FQDN format. This value can be in unicode format. Regular expression search is not supported for unicode values.",
	},
	"priority": schema.Int64Attribute{
		Required: true,
		Validators: []validator.Int64{
			int64validator.Between(0, 65535),
		},
		MarkdownDescription: "The priority of the SVCB record. Valid values are from 0 to 65535 (inclusive), in 32-bit unsigned integer format.",
	},
	"reclaimable": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines if the record is reclaimable or not.",
	},
	"svc_parameters": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RecordSvcbSvcParametersResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "Structure to represent SVC Params.",
	},
	"target_name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
			customvalidator.IsValidDomainName(),
		},
		MarkdownDescription: "Target name in FQDN format. This value can be in unicode format.",
	},
	"ttl": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_ttl")),
		},
		MarkdownDescription: "The Time To Live (TTL) value for record. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached.",
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
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
		MarkdownDescription: "The name of the DNS view in which the record resides. Example: \"external\".",
	},
	"zone": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the zone in which the record resides. Example: \"zone.com\". If a view is not specified when searching by zone, the default view is used.",
	},
	"extattrs_all": schema.MapAttribute{
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object, including default attributes.",
		ElementType:         types.StringType,
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
	},
}

func (m *RecordSvcbModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *dns.RecordSvcb {
	if m == nil {
		return nil
	}
	to := &dns.RecordSvcb{
		AwsRte53RecordInfo: ExpandRecordSvcbAwsRte53RecordInfo(ctx, m.AwsRte53RecordInfo, diags),
		CloudInfo:          ExpandRecordSvcbCloudInfo(ctx, m.CloudInfo, diags),
		Comment:            flex.ExpandStringPointer(m.Comment),
		Creator:            flex.ExpandStringPointer(m.Creator),
		DdnsPrincipal:      flex.ExpandStringPointer(m.DdnsPrincipal),
		DdnsProtected:      flex.ExpandBoolPointer(m.DdnsProtected),
		Disable:            flex.ExpandBoolPointer(m.Disable),
		ExtAttrs:           ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		ForbidReclamation:  flex.ExpandBoolPointer(m.ForbidReclamation),
		Name:               flex.ExpandStringPointer(m.Name),
		Priority:           flex.ExpandInt64Pointer(m.Priority),
		SvcParameters:      flex.ExpandFrameworkListNestedBlock(ctx, m.SvcParameters, diags, ExpandRecordSvcbSvcParameters),
		TargetName:         flex.ExpandStringPointer(m.TargetName),
		Ttl:                flex.ExpandInt64Pointer(m.Ttl),
		UseTtl:             flex.ExpandBoolPointer(m.UseTtl),
	}
	if isCreate {
		to.View = flex.ExpandStringPointer(m.View)
	}
	return to
}

func FlattenRecordSvcb(ctx context.Context, from *dns.RecordSvcb, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordSvcbAttrTypes)
	}
	m := RecordSvcbModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, RecordSvcbAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RecordSvcbModel) Flatten(ctx context.Context, from *dns.RecordSvcb, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RecordSvcbModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.AwsRte53RecordInfo = FlattenRecordSvcbAwsRte53RecordInfo(ctx, from.AwsRte53RecordInfo, diags)
	m.CloudInfo = FlattenRecordSvcbCloudInfo(ctx, from.CloudInfo, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CreationTime = flex.FlattenInt64Pointer(from.CreationTime)
	m.Creator = flex.FlattenStringPointer(from.Creator)
	m.DdnsPrincipal = flex.FlattenStringPointer(from.DdnsPrincipal)
	m.DdnsProtected = types.BoolPointerValue(from.DdnsProtected)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.ForbidReclamation = types.BoolPointerValue(from.ForbidReclamation)
	m.LastQueried = flex.FlattenInt64Pointer(from.LastQueried)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Priority = flex.FlattenInt64Pointer(from.Priority)
	m.Reclaimable = types.BoolPointerValue(from.Reclaimable)
	m.SvcParameters = flex.FlattenFrameworkListNestedBlock(ctx, from.SvcParameters, RecordSvcbSvcParametersAttrTypes, diags, FlattenRecordSvcbSvcParameters)
	m.TargetName = flex.FlattenStringPointer(from.TargetName)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
	m.UseTtl = types.BoolPointerValue(from.UseTtl)
	m.View = flex.FlattenStringPointer(from.View)
	m.Zone = flex.FlattenStringPointer(from.Zone)
}
