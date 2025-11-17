package dtc

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

	"github.com/infobloxopen/infoblox-nios-go-client/dtc"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/import"
	internaltypes "github.com/infobloxopen/terraform-provider-nios/internal/types"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type DtcPoolModel struct {
	Ref                      types.String                     `tfsdk:"ref"`
	AutoConsolidatedMonitors types.Bool                       `tfsdk:"auto_consolidated_monitors"`
	Availability             types.String                     `tfsdk:"availability"`
	Comment                  types.String                     `tfsdk:"comment"`
	ConsolidatedMonitors     types.List                       `tfsdk:"consolidated_monitors"`
	Disable                  types.Bool                       `tfsdk:"disable"`
	ExtAttrs                 types.Map                        `tfsdk:"extattrs"`
	ExtAttrsAll              types.Map                        `tfsdk:"extattrs_all"`
	Health                   types.Object                     `tfsdk:"health"`
	LbAlternateMethod        types.String                     `tfsdk:"lb_alternate_method"`
	LbAlternateTopology      types.String                     `tfsdk:"lb_alternate_topology"`
	LbDynamicRatioAlternate  types.Object                     `tfsdk:"lb_dynamic_ratio_alternate"`
	LbDynamicRatioPreferred  types.Object                     `tfsdk:"lb_dynamic_ratio_preferred"`
	LbPreferredMethod        types.String                     `tfsdk:"lb_preferred_method"`
	LbPreferredTopology      types.String                     `tfsdk:"lb_preferred_topology"`
	Monitors                 internaltypes.UnorderedListValue `tfsdk:"monitors"`
	Name                     types.String                     `tfsdk:"name"`
	Quorum                   types.Int64                      `tfsdk:"quorum"`
	Servers                  types.List                       `tfsdk:"servers"`
	Ttl                      types.Int64                      `tfsdk:"ttl"`
	UseTtl                   types.Bool                       `tfsdk:"use_ttl"`
}

var DtcPoolAttrTypes = map[string]attr.Type{
	"ref":                        types.StringType,
	"auto_consolidated_monitors": types.BoolType,
	"availability":               types.StringType,
	"comment":                    types.StringType,
	"consolidated_monitors":      types.ListType{ElemType: types.ObjectType{AttrTypes: DtcPoolConsolidatedMonitorsAttrTypes}},
	"disable":                    types.BoolType,
	"extattrs":                   types.MapType{ElemType: types.StringType},
	"extattrs_all":               types.MapType{ElemType: types.StringType},
	"health":                     types.ObjectType{AttrTypes: DtcPoolHealthAttrTypes},
	"lb_alternate_method":        types.StringType,
	"lb_alternate_topology":      types.StringType,
	"lb_dynamic_ratio_alternate": types.ObjectType{AttrTypes: DtcPoolLbDynamicRatioAlternateAttrTypes},
	"lb_dynamic_ratio_preferred": types.ObjectType{AttrTypes: DtcPoolLbDynamicRatioPreferredAttrTypes},
	"lb_preferred_method":        types.StringType,
	"lb_preferred_topology":      types.StringType,
	"monitors":                   internaltypes.UnorderedListOfStringType,
	"name":                       types.StringType,
	"quorum":                     types.Int64Type,
	"servers":                    types.ListType{ElemType: types.ObjectType{AttrTypes: DtcPoolServersAttrTypes}},
	"ttl":                        types.Int64Type,
	"use_ttl":                    types.BoolType,
}

var DtcPoolResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"auto_consolidated_monitors": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Flag for enabling auto managing DTC Consolidated Monitors in DTC Pool.",
	},
	"availability": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("ALL"),
		Validators: []validator.String{
			stringvalidator.OneOf("ALL", "ANY", "QUORUM"),
		},
		MarkdownDescription: "A resource in the pool is available if ANY, at least QUORUM, or ALL monitors for the pool say that it is up.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The comment for the DTC Pool; maximum 256 characters.",
	},
	"consolidated_monitors": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: DtcPoolConsolidatedMonitorsResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "List of monitors and associated members statuses of which are shared across members and consolidated in server availability determination.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether the DTC Pool is disabled or not. When this is set to False, the fixed address is enabled.",
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
	"health": schema.SingleNestedAttribute{
		Attributes:          DtcPoolHealthResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "The health status of DTC Pool",
	},
	"lb_alternate_method": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("NONE"),
		Validators: []validator.String{
			stringvalidator.OneOf("ALL_AVAILABLE", "DYNAMIC_RATIO", "GLOBAL_AVAILABILITY", "NONE", "RATIO", "ROUND_ROBIN", "SOURCE_IP_HASH", "TOPOLOGY"),
		},
		MarkdownDescription: "The alternate load balancing method. Use this to select a method type from the pool if the preferred method does not return any results.",
	},
	"lb_alternate_topology": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The alternate topology for load balancing.",
	},
	"lb_dynamic_ratio_alternate": schema.SingleNestedAttribute{
		Attributes:          DtcPoolLbDynamicRatioAlternateResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The DTC Pool settings for dynamic ratio when its selected as alternate method.",
	},
	"lb_dynamic_ratio_preferred": schema.SingleNestedAttribute{
		Attributes:          DtcPoolLbDynamicRatioPreferredResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The DTC Pool settings for dynamic ratio when its selected as preferred method.",
	},
	"lb_preferred_method": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.OneOf("ALL_AVAILABLE", "DYNAMIC_RATIO", "GLOBAL_AVAILABILITY", "RATIO", "ROUND_ROBIN", "SOURCE_IP_HASH", "TOPOLOGY"),
		},
		MarkdownDescription: "The preferred load balancing method. Use this to select a method type from the pool.",
	},
	"lb_preferred_topology": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The preferred topology for load balancing.",
	},
	"monitors": schema.ListAttribute{
		CustomType:  internaltypes.UnorderedListOfStringType,
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The monitors related to pool.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The DTC Pool display name.",
	},
	"quorum": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "For availability mode QUORUM, at least this many monitors must report the resource as up for it to be available",
	},
	"servers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: DtcPoolServersResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The servers related to the pool.",
	},
	"ttl": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The Time To Live (TTL) value for the DTC Pool. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached.",
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_ttl")),
		},
	},
	"use_ttl": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Flag to indicate whether the TTL value should be used for the DTC Pool.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
}

func (m *DtcPoolModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dtc.DtcPool {
	if m == nil {
		return nil
	}
	to := &dtc.DtcPool{
		AutoConsolidatedMonitors: flex.ExpandBoolPointer(m.AutoConsolidatedMonitors),
		Availability:             flex.ExpandStringPointer(m.Availability),
		Comment:                  flex.ExpandStringPointer(m.Comment),
		ConsolidatedMonitors:     flex.ExpandFrameworkListNestedBlock(ctx, m.ConsolidatedMonitors, diags, ExpandDtcPoolConsolidatedMonitors),
		Disable:                  flex.ExpandBoolPointer(m.Disable),
		ExtAttrs:                 ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		LbAlternateMethod:        flex.ExpandStringPointer(m.LbAlternateMethod),
		LbAlternateTopology:      flex.ExpandStringPointer(m.LbAlternateTopology),
		LbDynamicRatioAlternate:  ExpandDtcPoolLbDynamicRatioAlternate(ctx, m.LbDynamicRatioAlternate, diags),
		LbDynamicRatioPreferred:  ExpandDtcPoolLbDynamicRatioPreferred(ctx, m.LbDynamicRatioPreferred, diags),
		LbPreferredMethod:        flex.ExpandStringPointer(m.LbPreferredMethod),
		LbPreferredTopology:      flex.ExpandStringPointer(m.LbPreferredTopology),
		Monitors:                 flex.ExpandFrameworkListString(ctx, m.Monitors, diags),
		Name:                     flex.ExpandStringPointer(m.Name),
		Quorum:                   flex.ExpandInt64Pointer(m.Quorum),
		Servers:                  flex.ExpandFrameworkListNestedBlock(ctx, m.Servers, diags, ExpandDtcPoolServers),
		Ttl:                      flex.ExpandInt64Pointer(m.Ttl),
		UseTtl:                   flex.ExpandBoolPointer(m.UseTtl),
	}
	return to
}

func FlattenDtcPool(ctx context.Context, from *dtc.DtcPool, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DtcPoolAttrTypes)
	}
	m := DtcPoolModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, DtcPoolAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DtcPoolModel) Flatten(ctx context.Context, from *dtc.DtcPool, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DtcPoolModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.AutoConsolidatedMonitors = types.BoolPointerValue(from.AutoConsolidatedMonitors)
	m.Availability = flex.FlattenStringPointer(from.Availability)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.ConsolidatedMonitors = flex.FlattenFrameworkListNestedBlock(ctx, from.ConsolidatedMonitors, DtcPoolConsolidatedMonitorsAttrTypes, diags, FlattenDtcPoolConsolidatedMonitors)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.Health = FlattenDtcPoolHealth(ctx, from.Health, diags)
	m.LbAlternateMethod = flex.FlattenStringPointer(from.LbAlternateMethod)
	m.LbAlternateTopology = flex.FlattenStringPointer(from.LbAlternateTopology)
	m.LbDynamicRatioAlternate = FlattenDtcPoolLbDynamicRatioAlternate(ctx, from.LbDynamicRatioAlternate, diags)
	m.LbDynamicRatioPreferred = FlattenDtcPoolLbDynamicRatioPreferred(ctx, from.LbDynamicRatioPreferred, diags)
	m.LbPreferredMethod = flex.FlattenStringPointer(from.LbPreferredMethod)
	m.LbPreferredTopology = flex.FlattenStringPointer(from.LbPreferredTopology)
	m.Monitors = flex.FlattenFrameworkUnorderedList(ctx, types.StringType, from.Monitors, diags)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Quorum = FlattenQuorumBasedOnAvailability(from.Availability, from.Quorum)
	m.Servers = flex.FlattenFrameworkListNestedBlock(ctx, from.Servers, DtcPoolServersAttrTypes, diags, FlattenDtcPoolServers)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
	m.UseTtl = types.BoolPointerValue(from.UseTtl)
}

func FlattenQuorumBasedOnAvailability(availability *string, quorum *int64) types.Int64 {
	if availability == nil || *availability != "QUORUM" {
		return types.Int64Null()
	}
	return flex.FlattenInt64Pointer(quorum)
}
