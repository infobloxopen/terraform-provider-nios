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
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type RecordAliasModel struct {
	Ref                types.String `tfsdk:"ref"`
    Uuid        types.String `tfsdk:"uuid"`
	AwsRte53RecordInfo types.Object `tfsdk:"aws_rte53_record_info"`
	CloudInfo          types.Object `tfsdk:"cloud_info"`
	Comment            types.String `tfsdk:"comment"`
	Creator            types.String `tfsdk:"creator"`
	Disable            types.Bool   `tfsdk:"disable"`
	DnsName            types.String `tfsdk:"dns_name"`
	DnsTargetName      types.String `tfsdk:"dns_target_name"`
	ExtAttrs           types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll        types.Map    `tfsdk:"extattrs_all"`
	LastQueried        types.Int64  `tfsdk:"last_queried"`
	Name               types.String `tfsdk:"name"`
	TargetName         types.String `tfsdk:"target_name"`
	TargetType         types.String `tfsdk:"target_type"`
	Ttl                types.Int64  `tfsdk:"ttl"`
	UseTtl             types.Bool   `tfsdk:"use_ttl"`
	View               types.String `tfsdk:"view"`
	Zone               types.String `tfsdk:"zone"`
}

var RecordAliasAttrTypes = map[string]attr.Type{
	"ref":                   types.StringType,
    "uuid":        types.StringType,
	"aws_rte53_record_info": types.ObjectType{AttrTypes: RecordAliasAwsRte53RecordInfoAttrTypes},
	"cloud_info":            types.ObjectType{AttrTypes: RecordAliasCloudInfoAttrTypes},
	"comment":               types.StringType,
	"creator":               types.StringType,
	"disable":               types.BoolType,
	"dns_name":              types.StringType,
	"dns_target_name":       types.StringType,
	"extattrs":              types.MapType{ElemType: types.StringType},
	"extattrs_all":          types.MapType{ElemType: types.StringType},
	"last_queried":          types.Int64Type,
	"name":                  types.StringType,
	"target_name":           types.StringType,
	"target_type":           types.StringType,
	"ttl":                   types.Int64Type,
	"use_ttl":               types.BoolType,
	"view":                  types.StringType,
	"zone":                  types.StringType,
}

var RecordAliasResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
    "uuid": schema.StringAttribute{
        Computed:            true,
        MarkdownDescription: "The uuid to the object.",
    },
	"aws_rte53_record_info": schema.SingleNestedAttribute{
		Attributes:          RecordAliasAwsRte53RecordInfoResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "The AWS Route53 record information associated with the record.",
	},
	"cloud_info": schema.SingleNestedAttribute{
		Attributes:          RecordAliasCloudInfoResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "The cloud information associated with the record.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Comment for the record; maximum 256 characters.",
	},
	"creator": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("STATIC"),
		},
		Default:             stringdefault.StaticString("STATIC"),
		MarkdownDescription: "The record creator.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the record is disabled or not. False means that the record is enabled.",
	},
	"dns_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name for an Alias record in punycode format.",
	},
	"dns_target_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Target name in punycode format.",
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
		Required:            true,
		MarkdownDescription: "The Name of the Alias record.",
		Validators: []validator.String{
			customvalidator.IsValidFQDN(),
		},
	},
	"target_name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Target name in FQDN format. This value can be in unicode format.",
	},
	"target_type": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.OneOf("A", "AAAA", "MX", "NAPTR", "PTR", "SPF", "SRV", "TXT"),
		},
		MarkdownDescription: "Target type.",
	},
	"ttl": schema.Int64Attribute{
		Computed:            true,
		Optional:            true,
		MarkdownDescription: "Time-to-live value of the record, in seconds.",
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_ttl")),
		},
	},
	"use_ttl": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Flag to indicate whether the TTL value should be used for the A record.",
	},
	"view": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "View that this record is part of.",
	},
	"zone": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The zone in which the record resides.",
	},
}

func (m *RecordAliasModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.RecordAlias {
	if m == nil {
		return nil
	}
	to := &dns.RecordAlias{
		Comment:    flex.ExpandStringPointer(m.Comment),
		Creator:    flex.ExpandStringPointer(m.Creator),
		Disable:    flex.ExpandBoolPointer(m.Disable),
		ExtAttrs:   ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		Name:       flex.ExpandStringPointer(m.Name),
		TargetName: flex.ExpandStringPointer(m.TargetName),
		TargetType: flex.ExpandStringPointer(m.TargetType),
		Ttl:        flex.ExpandInt64Pointer(m.Ttl),
		UseTtl:     flex.ExpandBoolPointer(m.UseTtl),
		View:       flex.ExpandStringPointer(m.View),
	}
	return to
}

func FlattenRecordAlias(ctx context.Context, from *dns.RecordAlias, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordAliasAttrTypes)
	}
	m := RecordAliasModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, RecordAliasAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RecordAliasModel) Flatten(ctx context.Context, from *dns.RecordAlias, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RecordAliasModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
    m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.AwsRte53RecordInfo = FlattenRecordAliasAwsRte53RecordInfo(ctx, from.AwsRte53RecordInfo, diags)
	m.CloudInfo = FlattenRecordAliasCloudInfo(ctx, from.CloudInfo, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Creator = flex.FlattenStringPointer(from.Creator)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.DnsName = flex.FlattenStringPointer(from.DnsName)
	m.DnsTargetName = flex.FlattenStringPointer(from.DnsTargetName)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.LastQueried = flex.FlattenInt64Pointer(from.LastQueried)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.TargetName = flex.FlattenStringPointer(from.TargetName)
	m.TargetType = flex.FlattenStringPointer(from.TargetType)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
	m.UseTtl = types.BoolPointerValue(from.UseTtl)
	m.View = flex.FlattenStringPointer(from.View)
	m.Zone = flex.FlattenStringPointer(from.Zone)
}
