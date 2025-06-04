package dtc

import (
	"context"

	internaltypes "github.com/Infoblox-CTO/infoblox-nios-terraform/internal/types"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dtc"

	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/flex"
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
	"extattrs":                   types.MapType{ElemType: types.ObjectType{AttrTypes: ExtAttrAttrTypes}},
	"extattrs_all":               types.MapType{ElemType: types.ObjectType{AttrTypes: ExtAttrAttrTypes}},
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
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
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
		//validation for all
		MarkdownDescription: "A resource in the pool is available if ANY, at least QUORUM, or ALL monitors for the pool say that it is up.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The comment for the DTC Pool; maximum 256 characters.",
	},
	"consolidated_monitors": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: DtcPoolConsolidatedMonitorsResourceSchemaAttributes,
		},
		Optional:            true,
		Computed:            true,
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
		ElementType:         types.ObjectType{AttrTypes: ExtAttrAttrTypes},
		Default:             mapdefault.StaticValue(types.MapNull(types.ObjectType{AttrTypes: ExtAttrAttrTypes})),
		MarkdownDescription: "Extensible attributes associated with the object.",
	},
	"extattrs_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.ObjectType{AttrTypes: ExtAttrAttrTypes},
		MarkdownDescription: "Extensible attributes associated with the object , including default attributes.",
	},
	"health": schema.SingleNestedAttribute{
		Attributes: DtcPoolHealthResourceSchemaAttributes,
		Computed:   true,
	},
	"lb_alternate_method": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString("NONE"),
		MarkdownDescription: "The alternate load balancing method. Use this to select a method type from the pool if the preferred method does not return any results.",
	},
	"lb_alternate_topology": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The alternate topology for load balancing.",
	},
	"lb_dynamic_ratio_alternate": schema.SingleNestedAttribute{
		Attributes: DtcPoolLbDynamicRatioAlternateResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DtcPoolLbDynamicRatioAlternateAttrTypes,
				map[string]attr.Value{
					"method":                types.StringValue("MONITOR"),
					"monitor":               types.StringNull(),
					"monitor_metric":        types.StringNull(),
					"monitor_weighing":      types.StringValue("RATIO"),
					"invert_monitor_metric": types.BoolValue(false),
				},
			),
		),
	},
	"lb_dynamic_ratio_preferred": schema.SingleNestedAttribute{
		Attributes: DtcPoolLbDynamicRatioPreferredResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DtcPoolLbDynamicRatioAlternateAttrTypes,
				map[string]attr.Value{
					"method":                types.StringValue("MONITOR"),
					"monitor":               types.StringNull(),
					"monitor_metric":        types.StringNull(),
					"monitor_weighing":      types.StringValue("RATIO"),
					"invert_monitor_metric": types.BoolValue(false),
				},
			),
		),
	},
	"lb_preferred_method": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The preferred load balancing method. Use this to select a method type from the pool.",
	},
	"lb_preferred_topology": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The preferred topology for load balancing.",
	},
	"monitors": schema.ListAttribute{
		CustomType:          internaltypes.UnorderedListOfStringType,
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The monitors related to pool.",
	},
	"name": schema.StringAttribute{
		Optional:            true,
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
		Optional:            true,
		MarkdownDescription: "The servers related to the pool.",
	},
	"ttl": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The Time To Live (TTL) value for the DTC Pool. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached.",
	},
	"use_ttl": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: ttl",
	},
}

func ExpandDtcPool(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dtc.DtcPool {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DtcPoolModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DtcPoolModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dtc.DtcPool {
	if m == nil {
		return nil
	}
	to := &dtc.DtcPool{
		Ref:                      flex.ExpandStringPointer(m.Ref),
		AutoConsolidatedMonitors: flex.ExpandBoolPointer(m.AutoConsolidatedMonitors),
		Availability:             flex.ExpandStringPointer(m.Availability),
		Comment:                  flex.ExpandStringPointer(m.Comment),
		ConsolidatedMonitors:     flex.ExpandFrameworkListNestedBlock(ctx, m.ConsolidatedMonitors, diags, ExpandDtcPoolConsolidatedMonitors),
		Disable:                  flex.ExpandBoolPointer(m.Disable),
		Health:                   ExpandDtcPoolHealth(ctx, m.Health, diags),
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
		// Extattrs // TODO: should have been expanded, but generator didn't know how to.
		Extattrs: ExpandExtAttr(ctx, m.ExtAttrs, diags),
	}
	return to
}

func FlattenDtcPool(ctx context.Context, from *dtc.DtcPool, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DtcPoolAttrTypes)
	}
	m := DtcPoolModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrs = m.ExtAttrsAll
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
	if from.AutoConsolidatedMonitors == nil || !*from.AutoConsolidatedMonitors {
        m.ConsolidatedMonitors = flex.FlattenFrameworkListNestedBlock(ctx, from.ConsolidatedMonitors, DtcPoolConsolidatedMonitorsAttrTypes, diags, FlattenDtcPoolConsolidatedMonitors)
    } else {
        m.ConsolidatedMonitors = types.ListNull(types.ObjectType{AttrTypes: DtcPoolConsolidatedMonitorsAttrTypes})
    }
	m.Disable = types.BoolPointerValue(from.Disable)
	m.Health = FlattenDtcPoolHealth(ctx, from.Health, diags)
	m.LbAlternateMethod = flex.FlattenStringPointer(from.LbAlternateMethod)
	m.LbAlternateTopology = flex.FlattenStringPointer(from.LbAlternateTopology)
	m.LbDynamicRatioAlternate = FlattenDtcPoolLbDynamicRatioAlternate(ctx, from.LbDynamicRatioAlternate, diags)
	m.LbDynamicRatioPreferred = FlattenDtcPoolLbDynamicRatioPreferred(ctx, from.LbDynamicRatioPreferred, diags)
	m.LbPreferredMethod = flex.FlattenStringPointer(from.LbPreferredMethod)
	m.LbPreferredTopology = flex.FlattenStringPointer(from.LbPreferredTopology)
	m.Monitors = flex.FlattenFrameworkUnorderedList(ctx, types.StringType, from.Monitors, diags)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Quorum = flex.FlattenInt64Pointer(from.Quorum)
	m.Servers = flex.FlattenFrameworkListNestedBlock(ctx, from.Servers, DtcPoolServersAttrTypes, diags, FlattenDtcPoolServers)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
	m.UseTtl = types.BoolPointerValue(from.UseTtl)
	// Extattrs // TODO: should have been flattened, but generator didn't know how to.
	m.ExtAttrsAll = FlattenExtAttr(ctx, *from.Extattrs, diags)
}
