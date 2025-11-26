package rpz

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
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

type RecordRpzPtrModel struct {
	Ref         types.String        `tfsdk:"ref"`
	Comment     types.String        `tfsdk:"comment"`
	Disable     types.Bool          `tfsdk:"disable"`
	ExtAttrs    types.Map           `tfsdk:"extattrs"`
	Ipv4addr    iptypes.IPv4Address `tfsdk:"ipv4addr"`
	Ipv6addr    iptypes.IPv6Address `tfsdk:"ipv6addr"`
	Name        types.String        `tfsdk:"name"`
	Ptrdname    types.String        `tfsdk:"ptrdname"`
	RpZone      types.String        `tfsdk:"rp_zone"`
	Ttl         types.Int64         `tfsdk:"ttl"`
	UseTtl      types.Bool          `tfsdk:"use_ttl"`
	View        types.String        `tfsdk:"view"`
	Zone        types.String        `tfsdk:"zone"`
	ExtAttrsAll types.Map           `tfsdk:"extattrs_all"`
}

var RecordRpzPtrAttrTypes = map[string]attr.Type{
	"ref":          types.StringType,
	"comment":      types.StringType,
	"disable":      types.BoolType,
	"extattrs":     types.MapType{ElemType: types.StringType},
	"ipv4addr":     iptypes.IPv4AddressType{},
	"ipv6addr":     iptypes.IPv6AddressType{},
	"name":         types.StringType,
	"ptrdname":     types.StringType,
	"rp_zone":      types.StringType,
	"ttl":          types.Int64Type,
	"use_ttl":      types.BoolType,
	"view":         types.StringType,
	"zone":         types.StringType,
	"extattrs_all": types.MapType{ElemType: types.StringType},
}

var RecordRpzPtrResourceSchemaAttributes = map[string]schema.Attribute{
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
	"ipv4addr": schema.StringAttribute{
		CustomType:          iptypes.IPv4AddressType{},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The IPv4 Address of the substitute rule.",
	},
	"ipv6addr": schema.StringAttribute{
		CustomType:          iptypes.IPv6AddressType{},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The IPv6 Address of the substitute rule.",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The name of the RPZ Substitute (PTR Record) Rule object in FQDN format.",
		Validators: []validator.String{
			stringvalidator.Any(
				customvalidator.IsValidArpaIPv4RPZ(),
				customvalidator.IsValidArpaIPv6RPZ(),
			),
		},
	},
	"ptrdname": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.IsValidDomainName(),
		},
		MarkdownDescription: "The domain name of the RPZ Substitute (PTR Record) Rule object in FQDN format.",
	},
	"rp_zone": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of a response policy zone in which the record resides.",
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
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
		Computed:            true,
		Optional:            true,
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

func (m *RecordRpzPtrModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *rpz.RecordRpzPtr {
	if m == nil {
		return nil
	}
	to := &rpz.RecordRpzPtr{
		Comment:  flex.ExpandStringPointer(m.Comment),
		Disable:  flex.ExpandBoolPointer(m.Disable),
		ExtAttrs: ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		Ipv4addr: flex.ExpandIPv4Address(m.Ipv4addr),
		Ipv6addr: flex.ExpandIPv6Address(m.Ipv6addr),
		Name:     flex.ExpandStringPointer(m.Name),
		Ptrdname: flex.ExpandStringPointer(m.Ptrdname),
		RpZone:   flex.ExpandStringPointer(m.RpZone),
		Ttl:      flex.ExpandInt64Pointer(m.Ttl),
		UseTtl:   flex.ExpandBoolPointer(m.UseTtl),
	}
	if isCreate {
		to.View = flex.ExpandStringPointer(m.View)
	}
	return to
}

func FlattenRecordRpzPtr(ctx context.Context, from *rpz.RecordRpzPtr, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordRpzPtrAttrTypes)
	}
	m := RecordRpzPtrModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, RecordRpzPtrAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RecordRpzPtrModel) Flatten(ctx context.Context, from *rpz.RecordRpzPtr, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RecordRpzPtrModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.Ipv4addr = flex.FlattenIPv4Address(from.Ipv4addr)
	m.Ipv6addr = flex.FlattenIPv6Address(from.Ipv6addr)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Ptrdname = flex.FlattenStringPointer(from.Ptrdname)
	m.RpZone = flex.FlattenStringPointer(from.RpZone)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
	m.UseTtl = types.BoolPointerValue(from.UseTtl)
	m.View = flex.FlattenStringPointer(from.View)
	m.Zone = flex.FlattenStringPointer(from.Zone)
}
