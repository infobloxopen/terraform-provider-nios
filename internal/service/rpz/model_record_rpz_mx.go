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

type RecordRpzMxModel struct {
	Ref           types.String `tfsdk:"ref"`
	Uuid          types.String `tfsdk:"uuid"`
	Comment       types.String `tfsdk:"comment"`
	Disable       types.Bool   `tfsdk:"disable"`
	ExtAttrs      types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll   types.Map    `tfsdk:"extattrs_all"`
	MailExchanger types.String `tfsdk:"mail_exchanger"`
	Name          types.String `tfsdk:"name"`
	Preference    types.Int64  `tfsdk:"preference"`
	RpZone        types.String `tfsdk:"rp_zone"`
	Ttl           types.Int64  `tfsdk:"ttl"`
	UseTtl        types.Bool   `tfsdk:"use_ttl"`
	View          types.String `tfsdk:"view"`
	Zone          types.String `tfsdk:"zone"`
}

var RecordRpzMxAttrTypes = map[string]attr.Type{
	"ref":            types.StringType,
	"uuid":           types.StringType,
	"comment":        types.StringType,
	"disable":        types.BoolType,
	"extattrs":       types.MapType{ElemType: types.StringType},
	"extattrs_all":   types.MapType{ElemType: types.StringType},
	"mail_exchanger": types.StringType,
	"name":           types.StringType,
	"preference":     types.Int64Type,
	"rp_zone":        types.StringType,
	"ttl":            types.Int64Type,
	"use_ttl":        types.BoolType,
	"view":           types.StringType,
	"zone":           types.StringType,
}

var RecordRpzMxResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The uuid to the object.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
			stringvalidator.LengthBetween(0, 256),
		},
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The comment for the record; maximum 256 characters.",
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
	"extattrs_all": schema.MapAttribute{
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object, including default attributes.",
		ElementType:         types.StringType,
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
	},
	"mail_exchanger": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.IsValidFQDN(),
		},
		MarkdownDescription: "Mail exchanger name in FQDN format. This value can be in unicode format.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.IsValidFQDN(),
		},
		MarkdownDescription: "The name for a record in FQDN format. This value cannot be in unicode format.",
	},
	"preference": schema.Int64Attribute{
		Required: true,
		Validators: []validator.Int64{
			int64validator.Between(0, 65535),
		},
		MarkdownDescription: "Preference value, 0 to 65535 (inclusive) in 32-bit unsigned integer format.",
	},
	"rp_zone": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
		MarkdownDescription: "The name of a response policy zone in which the record resides.",
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
}

func (m *RecordRpzMxModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *rpz.RecordRpzMx {
	if m == nil {
		return nil
	}
	to := &rpz.RecordRpzMx{
		Comment:       flex.ExpandStringPointer(m.Comment),
		Disable:       flex.ExpandBoolPointer(m.Disable),
		ExtAttrs:      ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		MailExchanger: flex.ExpandStringPointer(m.MailExchanger),
		Name:          flex.ExpandStringPointer(m.Name),
		Preference:    flex.ExpandInt64Pointer(m.Preference),
		Ttl:           flex.ExpandInt64Pointer(m.Ttl),
		UseTtl:        flex.ExpandBoolPointer(m.UseTtl),
	}
	if isCreate {
		to.RpZone = flex.ExpandStringPointer(m.RpZone)
		to.View = flex.ExpandStringPointer(m.View)
	}
	return to
}

func FlattenRecordRpzMx(ctx context.Context, from *rpz.RecordRpzMx, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordRpzMxAttrTypes)
	}
	m := RecordRpzMxModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, RecordRpzMxAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RecordRpzMxModel) Flatten(ctx context.Context, from *rpz.RecordRpzMx, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RecordRpzMxModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.MailExchanger = flex.FlattenStringPointer(from.MailExchanger)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Preference = flex.FlattenInt64Pointer(from.Preference)
	m.RpZone = flex.FlattenStringPointer(from.RpZone)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
	m.UseTtl = types.BoolPointerValue(from.UseTtl)
	m.View = flex.FlattenStringPointer(from.View)
	m.Zone = flex.FlattenStringPointer(from.Zone)
}
