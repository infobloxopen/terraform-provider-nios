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
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type SharedrecordSrvModel struct {
	Ref               types.String `tfsdk:"ref"`
	Comment           types.String `tfsdk:"comment"`
	Disable           types.Bool   `tfsdk:"disable"`
	DnsName           types.String `tfsdk:"dns_name"`
	DnsTarget         types.String `tfsdk:"dns_target"`
	ExtAttrs          types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll       types.Map    `tfsdk:"extattrs_all"`
	Name              types.String `tfsdk:"name"`
	Port              types.Int64  `tfsdk:"port"`
	Priority          types.Int64  `tfsdk:"priority"`
	SharedRecordGroup types.String `tfsdk:"shared_record_group"`
	Target            types.String `tfsdk:"target"`
	Ttl               types.Int64  `tfsdk:"ttl"`
	UseTtl            types.Bool   `tfsdk:"use_ttl"`
	Weight            types.Int64  `tfsdk:"weight"`
}

var SharedrecordSrvAttrTypes = map[string]attr.Type{
	"ref":                 types.StringType,
	"comment":             types.StringType,
	"disable":             types.BoolType,
	"dns_name":            types.StringType,
	"dns_target":          types.StringType,
	"extattrs":            types.MapType{ElemType: types.StringType},
	"extattrs_all":        types.MapType{ElemType: types.StringType},
	"name":                types.StringType,
	"port":                types.Int64Type,
	"priority":            types.Int64Type,
	"shared_record_group": types.StringType,
	"target":              types.StringType,
	"ttl":                 types.Int64Type,
	"use_ttl":             types.BoolType,
	"weight":              types.Int64Type,
}

var SharedrecordSrvResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
			stringvalidator.LengthAtMost(256),
		},
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "Comment for this shared record; maximum 256 characters.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if this shared record is disabled or not. False means that the record is enabled.",
	},
	"dns_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name for this shared record in punycode format.",
	},
	"dns_target": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name for a shared SRV record in punycode format.",
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
	"extattrs_all": schema.MapAttribute{
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object , including default attributes.",
		ElementType:         types.StringType,
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Name for this shared record. This value can be in unicode format.",
	},
	"port": schema.Int64Attribute{
		Required: true,
		Validators: []validator.Int64{
			int64validator.Between(0, 65535),
		},
		MarkdownDescription: "The port of the shared SRV record. Valid values are from 0 to 65535 (inclusive), in 32-bit unsigned integer format.",
	},
	"priority": schema.Int64Attribute{
		Required: true,
		Validators: []validator.Int64{
			int64validator.Between(0, 65535),
		},
		MarkdownDescription: "The priority of the shared SRV record. Valid values are from 0 to 65535 (inclusive), in 32-bit unsigned integer format.",
	},
	"shared_record_group": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
		MarkdownDescription: "The name of the shared record group in which the record resides.",
	},
	"target": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.IsValidDomainName(),
		},
		MarkdownDescription: "The target of the shared SRV record in FQDN format. This value can be in unicode format.",
	},
	"ttl": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_ttl")),
		},
		MarkdownDescription: "The Time To Live (TTL) value for this shared record. A 32-bit unsigned integer that represents the duration, in seconds, for which the shared record is valid (cached). Zero indicates that the shared record should not be cached.",
	},
	"use_ttl": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: ttl",
	},
	"weight": schema.Int64Attribute{
		Required: true,
		Validators: []validator.Int64{
			int64validator.Between(0, 65535),
		},
		MarkdownDescription: "The weight of the shared SRV record. Valid values are from 0 to 65535 (inclusive), in 32-bit unsigned integer format.",
	},
}

func (m *SharedrecordSrvModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *dns.SharedrecordSrv {
	if m == nil {
		return nil
	}
	to := &dns.SharedrecordSrv{
		Comment:  flex.ExpandStringPointer(m.Comment),
		Disable:  flex.ExpandBoolPointer(m.Disable),
		ExtAttrs: ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		Name:     flex.ExpandStringPointer(m.Name),
		Port:     flex.ExpandInt64Pointer(m.Port),
		Priority: flex.ExpandInt64Pointer(m.Priority),
		Target:   flex.ExpandStringPointer(m.Target),
		Ttl:      flex.ExpandInt64Pointer(m.Ttl),
		UseTtl:   flex.ExpandBoolPointer(m.UseTtl),
		Weight:   flex.ExpandInt64Pointer(m.Weight),
	}
	if isCreate {
		to.SharedRecordGroup = flex.ExpandStringPointer(m.SharedRecordGroup)
	}

	return to
}

func FlattenSharedrecordSrv(ctx context.Context, from *dns.SharedrecordSrv, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SharedrecordSrvAttrTypes)
	}
	m := SharedrecordSrvModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, SharedrecordSrvAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *SharedrecordSrvModel) Flatten(ctx context.Context, from *dns.SharedrecordSrv, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = SharedrecordSrvModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.DnsName = flex.FlattenStringPointer(from.DnsName)
	m.DnsTarget = flex.FlattenStringPointer(from.DnsTarget)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Port = flex.FlattenInt64Pointer(from.Port)
	m.Priority = flex.FlattenInt64Pointer(from.Priority)
	m.SharedRecordGroup = flex.FlattenStringPointer(from.SharedRecordGroup)
	m.Target = flex.FlattenStringPointer(from.Target)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
	m.UseTtl = types.BoolPointerValue(from.UseTtl)
	m.Weight = flex.FlattenInt64Pointer(from.Weight)
}
