package rpz

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

	"github.com/infobloxopen/infoblox-nios-go-client/rpz"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	planmodifiers "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/immutable"
	importmod "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/import"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type RecordRpzNaptrModel struct {
	Ref         types.String `tfsdk:"ref"`
	Comment     types.String `tfsdk:"comment"`
	Disable     types.Bool   `tfsdk:"disable"`
	ExtAttrs    types.Map    `tfsdk:"extattrs"`
	Flags       types.String `tfsdk:"flags"`
	LastQueried types.Int64  `tfsdk:"last_queried"`
	Name        types.String `tfsdk:"name"`
	Order       types.Int64  `tfsdk:"order"`
	Preference  types.Int64  `tfsdk:"preference"`
	Regexp      types.String `tfsdk:"regexp"`
	Replacement types.String `tfsdk:"replacement"`
	RpZone      types.String `tfsdk:"rp_zone"`
	Services    types.String `tfsdk:"services"`
	Ttl         types.Int64  `tfsdk:"ttl"`
	UseTtl      types.Bool   `tfsdk:"use_ttl"`
	View        types.String `tfsdk:"view"`
	Zone        types.String `tfsdk:"zone"`
	ExtAttrsAll types.Map    `tfsdk:"extattrs_all"`
}

var RecordRpzNaptrAttrTypes = map[string]attr.Type{
	"ref":          types.StringType,
	"comment":      types.StringType,
	"disable":      types.BoolType,
	"extattrs":     types.MapType{ElemType: types.StringType},
	"flags":        types.StringType,
	"last_queried": types.Int64Type,
	"name":         types.StringType,
	"order":        types.Int64Type,
	"preference":   types.Int64Type,
	"regexp":       types.StringType,
	"replacement":  types.StringType,
	"rp_zone":      types.StringType,
	"services":     types.StringType,
	"ttl":          types.Int64Type,
	"use_ttl":      types.BoolType,
	"view":         types.StringType,
	"zone":         types.StringType,
	"extattrs_all": types.MapType{ElemType: types.StringType},
}

var RecordRpzNaptrResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The comment for the record; maximum 256 characters.",
		Default:             stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 256),
			customvalidator.ValidateTrimmedString(),
		},
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the record is disabled or not. False means that the record is enabled.",
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
	"flags": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
			stringvalidator.OneOf("U", "S", "P", "A", ""),
		},
		MarkdownDescription: "The flags used to control the interpretation of the fields for a Substitute (NAPTR Record) Rule object. Supported values for the flags field are \"U\", \"S\", \"P\" and \"A\".",
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
		MarkdownDescription: "The name for a record in FQDN format. This value cannot be in unicode format.",
	},
	"order": schema.Int64Attribute{
		Required: true,
		Validators: []validator.Int64{
			int64validator.Between(0, 65535),
		},
		MarkdownDescription: "The order parameter of the Substitute (NAPTR Record) Rule records. This parameter specifies the order in which the NAPTR rules are applied when multiple rules are present. Valid values are from 0 to 65535 (inclusive), in 32-bit unsigned integer format.",
	},
	"preference": schema.Int64Attribute{
		Required: true,
		Validators: []validator.Int64{
			int64validator.Between(0, 65535),
		},
		MarkdownDescription: "The preference of the Substitute (NAPTR Record) Rule record. The preference field determines the order NAPTR records are processed when multiple records with the same order parameter are present. Valid values are from 0 to 65535 (inclusive), in 32-bit unsigned integer format.",
	},
	"regexp": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The regular expression-based rewriting rule of the Substitute (NAPTR Record) Rule record. This should be a POSIX compliant regular expression, including the substitution rule and flags. Refer to RFC 2915 for the field syntax details.",
	},
	"replacement": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The replacement field of the Substitute (NAPTR Record) Rule object. For nonterminal NAPTR records, this field specifies the next domain name to look up. This value can be in unicode format.",
	},
	"rp_zone": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of a response policy zone in which the record resides.",
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
	},
	"services": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
			stringvalidator.LengthBetween(0, 128),
		},
		MarkdownDescription: "The services field of the Substitute (NAPTR Record) Rule object; maximum 128 characters. The services field contains protocol and service identifiers, such as \"http+E2U\" or \"SIPS+D2T\".",
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
		MarkdownDescription: "The name of the DNS View in which the record resides. Example: \"external\".",
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

func (m *RecordRpzNaptrModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *rpz.RecordRpzNaptr {
	if m == nil {
		return nil
	}
	to := &rpz.RecordRpzNaptr{
		Comment:     flex.ExpandStringPointer(m.Comment),
		Disable:     flex.ExpandBoolPointer(m.Disable),
		ExtAttrs:    ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		Flags:       flex.ExpandStringPointer(m.Flags),
		Name:        flex.ExpandStringPointer(m.Name),
		Order:       flex.ExpandInt64Pointer(m.Order),
		Preference:  flex.ExpandInt64Pointer(m.Preference),
		Regexp:      flex.ExpandStringPointer(m.Regexp),
		Replacement: flex.ExpandStringPointer(m.Replacement),
		Services:    flex.ExpandStringPointer(m.Services),
		Ttl:         flex.ExpandInt64Pointer(m.Ttl),
		UseTtl:      flex.ExpandBoolPointer(m.UseTtl),
	}
	if isCreate {
		to.RpZone = flex.ExpandStringPointer(m.RpZone)
		to.View = flex.ExpandStringPointer(m.View)
	}
	return to
}

func FlattenRecordRpzNaptr(ctx context.Context, from *rpz.RecordRpzNaptr, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordRpzNaptrAttrTypes)
	}
	m := RecordRpzNaptrModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, RecordRpzNaptrAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RecordRpzNaptrModel) Flatten(ctx context.Context, from *rpz.RecordRpzNaptr, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RecordRpzNaptrModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.Flags = flex.FlattenStringPointer(from.Flags)
	m.LastQueried = flex.FlattenInt64Pointer(from.LastQueried)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Order = flex.FlattenInt64Pointer(from.Order)
	m.Preference = flex.FlattenInt64Pointer(from.Preference)
	m.Regexp = flex.FlattenStringPointer(from.Regexp)
	m.Replacement = flex.FlattenStringPointer(from.Replacement)
	m.RpZone = flex.FlattenStringPointer(from.RpZone)
	m.Services = flex.FlattenStringPointer(from.Services)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
	m.UseTtl = types.BoolPointerValue(from.UseTtl)
	m.View = flex.FlattenStringPointer(from.View)
	m.Zone = flex.FlattenStringPointer(from.Zone)
}
