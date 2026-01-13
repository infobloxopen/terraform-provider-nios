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

type RecordRpzCnameClientipaddressModel struct {
	Ref         types.String `tfsdk:"ref"`
	Canonical   types.String `tfsdk:"canonical"`
	Comment     types.String `tfsdk:"comment"`
	Disable     types.Bool   `tfsdk:"disable"`
	ExtAttrs    types.Map    `tfsdk:"extattrs"`
	IsIpv4      types.Bool   `tfsdk:"is_ipv4"`
	Name        types.String `tfsdk:"name"`
	RpZone      types.String `tfsdk:"rp_zone"`
	Ttl         types.Int64  `tfsdk:"ttl"`
	UseTtl      types.Bool   `tfsdk:"use_ttl"`
	View        types.String `tfsdk:"view"`
	Zone        types.String `tfsdk:"zone"`
	ExtAttrsAll types.Map    `tfsdk:"extattrs_all"`
}

var RecordRpzCnameClientipaddressAttrTypes = map[string]attr.Type{
	"ref":          types.StringType,
	"canonical":    types.StringType,
	"comment":      types.StringType,
	"disable":      types.BoolType,
	"extattrs":     types.MapType{ElemType: types.StringType},
	"is_ipv4":      types.BoolType,
	"name":         types.StringType,
	"rp_zone":      types.StringType,
	"ttl":          types.Int64Type,
	"use_ttl":      types.BoolType,
	"view":         types.StringType,
	"zone":         types.StringType,
	"extattrs_all": types.MapType{ElemType: types.StringType},
}

var RecordRpzCnameClientipaddressResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"canonical": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The canonical name in FQDN format. This value can be in unicode format.",
		Validators: []validator.String{
			customvalidator.IsValidFQDN(),
		},
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
	"is_ipv4": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Indicates whether the record is an IPv4 record. If the return value is \"true\", it is an IPv4 record. Ohterwise, it is an IPv6 record.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.IsValidDomainName(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name for a record in FQDN format. This value cannot be in unicode format.",
	},
	"rp_zone": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
		MarkdownDescription: "The name of a response policy zone in which the record resides.",
	},
	"ttl": schema.Int64Attribute{
		Computed: true,
		Optional: true,
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
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
		Default:             stringdefault.StaticString("default"),
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

func (m *RecordRpzCnameClientipaddressModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *rpz.RecordRpzCnameClientipaddress {
	if m == nil {
		return nil
	}
	to := &rpz.RecordRpzCnameClientipaddress{
		Canonical: flex.ExpandStringPointer(m.Canonical),
		Comment:   flex.ExpandStringPointer(m.Comment),
		Disable:   flex.ExpandBoolPointer(m.Disable),
		ExtAttrs:  ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		Name:      flex.ExpandStringPointer(m.Name),
		Ttl:       flex.ExpandInt64Pointer(m.Ttl),
		UseTtl:    flex.ExpandBoolPointer(m.UseTtl),
	}

	if isCreate {
		to.RpZone = flex.ExpandStringPointer(m.RpZone)
		to.View = flex.ExpandStringPointer(m.View)
	}
	return to
}

func FlattenRecordRpzCnameClientipaddress(ctx context.Context, from *rpz.RecordRpzCnameClientipaddress, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordRpzCnameClientipaddressAttrTypes)
	}
	m := RecordRpzCnameClientipaddressModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, RecordRpzCnameClientipaddressAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RecordRpzCnameClientipaddressModel) Flatten(ctx context.Context, from *rpz.RecordRpzCnameClientipaddress, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RecordRpzCnameClientipaddressModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Canonical = flex.FlattenStringPointer(from.Canonical)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.IsIpv4 = types.BoolPointerValue(from.IsIpv4)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.RpZone = flex.FlattenStringPointer(from.RpZone)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
	m.UseTtl = types.BoolPointerValue(from.UseTtl)
	m.View = flex.FlattenStringPointer(from.View)
	m.Zone = flex.FlattenStringPointer(from.Zone)
}
