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

type SharedrecordCnameModel struct {
	Ref               types.String `tfsdk:"ref"`
	Canonical         types.String `tfsdk:"canonical"`
	Comment           types.String `tfsdk:"comment"`
	Disable           types.Bool   `tfsdk:"disable"`
	DnsCanonical      types.String `tfsdk:"dns_canonical"`
	DnsName           types.String `tfsdk:"dns_name"`
	ExtAttrs          types.Map    `tfsdk:"extattrs"`
	Name              types.String `tfsdk:"name"`
	SharedRecordGroup types.String `tfsdk:"shared_record_group"`
	Ttl               types.Int64  `tfsdk:"ttl"`
	UseTtl            types.Bool   `tfsdk:"use_ttl"`
	ExtAttrsAll       types.Map    `tfsdk:"extattrs_all"`
}

var SharedrecordCnameAttrTypes = map[string]attr.Type{
	"ref":                 types.StringType,
	"canonical":           types.StringType,
	"comment":             types.StringType,
	"disable":             types.BoolType,
	"dns_canonical":       types.StringType,
	"dns_name":            types.StringType,
	"extattrs":            types.MapType{ElemType: types.StringType},
	"name":                types.StringType,
	"shared_record_group": types.StringType,
	"ttl":                 types.Int64Type,
	"use_ttl":             types.BoolType,
	"extattrs_all":        types.MapType{ElemType: types.StringType},
}

var SharedrecordCnameResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"canonical": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.IsValidDomainName(),
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Canonical name in FQDN format. This value can be in unicode format.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "Comment for this shared record; maximum 256 characters.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if this shared record is disabled or not. False means that the record is enabled.",
	},
	"dns_canonical": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Canonical name in punycode format.",
	},
	"dns_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name for this shared record in punycode format.",
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
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Name for this shared record. This value can be in unicode format.",
	},
	"shared_record_group": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
		MarkdownDescription: "The name of the shared record group in which the record resides.",
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
	"extattrs_all": schema.MapAttribute{
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object, including default attributes.",
		ElementType:         types.StringType,
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
	},
}

func (m *SharedrecordCnameModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *dns.SharedrecordCname {
	if m == nil {
		return nil
	}
	to := &dns.SharedrecordCname{
		Canonical: flex.ExpandStringPointer(m.Canonical),
		Comment:   flex.ExpandStringPointer(m.Comment),
		Disable:   flex.ExpandBoolPointer(m.Disable),
		ExtAttrs:  ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		Name:      flex.ExpandStringPointer(m.Name),
		Ttl:       flex.ExpandInt64Pointer(m.Ttl),
		UseTtl:    flex.ExpandBoolPointer(m.UseTtl),
	}
	if isCreate {
		to.SharedRecordGroup = flex.ExpandStringPointer(m.SharedRecordGroup)
	}
	return to
}

func FlattenSharedrecordCname(ctx context.Context, from *dns.SharedrecordCname, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SharedrecordCnameAttrTypes)
	}
	m := SharedrecordCnameModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, SharedrecordCnameAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *SharedrecordCnameModel) Flatten(ctx context.Context, from *dns.SharedrecordCname, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = SharedrecordCnameModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Canonical = flex.FlattenStringPointer(from.Canonical)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.DnsCanonical = flex.FlattenStringPointer(from.DnsCanonical)
	m.DnsName = flex.FlattenStringPointer(from.DnsName)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.SharedRecordGroup = flex.FlattenStringPointer(from.SharedRecordGroup)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
	m.UseTtl = types.BoolPointerValue(from.UseTtl)
}
