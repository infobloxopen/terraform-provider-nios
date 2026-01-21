package dns

import (
	"context"

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
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/import"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type SharedrecordgroupModel struct {
	Ref                 types.String `tfsdk:"ref"`
	Uuid                types.String `tfsdk:"uuid"`
	Comment             types.String `tfsdk:"comment"`
	ExtAttrs            types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll         types.Map    `tfsdk:"extattrs_all"`
	Name                types.String `tfsdk:"name"`
	RecordNamePolicy    types.String `tfsdk:"record_name_policy"`
	UseRecordNamePolicy types.Bool   `tfsdk:"use_record_name_policy"`
	ZoneAssociations    types.List   `tfsdk:"zone_associations"`
}

var SharedrecordgroupAttrTypes = map[string]attr.Type{
	"ref":                    types.StringType,
	"uuid":                   types.StringType,
	"comment":                types.StringType,
	"extattrs":               types.MapType{ElemType: types.StringType},
	"extattrs_all":           types.MapType{ElemType: types.StringType},
	"name":                   types.StringType,
	"record_name_policy":     types.StringType,
	"use_record_name_policy": types.BoolType,
	"zone_associations":      types.ListType{ElemType: types.ObjectType{AttrTypes: SharedrecordgroupZoneAssociationsAttrTypes}},
}

var SharedrecordgroupResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The UUID of the object.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "The descriptive comment of this shared record group.",
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
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of this shared record group.",
	},
	"record_name_policy": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("Strict Hostname Checking", "Allow Underscore", "Allow Any"),
			stringvalidator.AlsoRequires(path.MatchRoot("use_record_name_policy")),
		},
		MarkdownDescription: "The record name policy of this shared record group.",
	},
	"use_record_name_policy": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: record_name_policy",
	},
	"zone_associations": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: SharedrecordgroupZoneAssociationsResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of zones associated with this shared record group. Starting from NIOS-9.0.6, this field has been updated to a structure that includes FQDN and DNS view details.",
	},
}

func (m *SharedrecordgroupModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.Sharedrecordgroup {
	if m == nil {
		return nil
	}
	to := &dns.Sharedrecordgroup{
		Comment:             flex.ExpandStringPointer(m.Comment),
		ExtAttrs:            ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		Name:                flex.ExpandStringPointer(m.Name),
		RecordNamePolicy:    flex.ExpandStringPointer(m.RecordNamePolicy),
		UseRecordNamePolicy: flex.ExpandBoolPointer(m.UseRecordNamePolicy),
		ZoneAssociations:    flex.ExpandFrameworkListNestedBlock(ctx, m.ZoneAssociations, diags, ExpandSharedrecordgroupZoneAssociations),
	}
	return to
}

func FlattenSharedrecordgroup(ctx context.Context, from *dns.Sharedrecordgroup, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SharedrecordgroupAttrTypes)
	}
	m := SharedrecordgroupModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, SharedrecordgroupAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *SharedrecordgroupModel) Flatten(ctx context.Context, from *dns.Sharedrecordgroup, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = SharedrecordgroupModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.RecordNamePolicy = flex.FlattenStringPointer(from.RecordNamePolicy)
	m.UseRecordNamePolicy = types.BoolPointerValue(from.UseRecordNamePolicy)
	planZoneAssociations := m.ZoneAssociations
	m.ZoneAssociations = flex.FlattenFrameworkListNestedBlock(ctx, from.ZoneAssociations, SharedrecordgroupZoneAssociationsAttrTypes, diags, FlattenSharedrecordgroupZoneAssociations)
	if !planZoneAssociations.IsUnknown() {
		reOrderedZoneAssociations, diags := utils.ReorderAndFilterNestedListResponse(ctx, planZoneAssociations, m.ZoneAssociations, "fqdn")
		if !diags.HasError() {
			m.ZoneAssociations = reOrderedZoneAssociations.(basetypes.ListValue)
		}
	}
}
