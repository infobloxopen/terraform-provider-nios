package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dtc"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/import"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type DtcMonitorSnmpModel struct {
	Ref         types.String `tfsdk:"ref"`
	Comment     types.String `tfsdk:"comment"`
	Community   types.String `tfsdk:"community"`
	Context     types.String `tfsdk:"context"`
	EngineId    types.String `tfsdk:"engine_id"`
	ExtAttrs    types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll types.Map    `tfsdk:"extattrs_all"`
	Interval    types.Int64  `tfsdk:"interval"`
	Name        types.String `tfsdk:"name"`
	Oids        types.List   `tfsdk:"oids"`
	Port        types.Int64  `tfsdk:"port"`
	RetryDown   types.Int64  `tfsdk:"retry_down"`
	RetryUp     types.Int64  `tfsdk:"retry_up"`
	Timeout     types.Int64  `tfsdk:"timeout"`
	User        types.String `tfsdk:"user"`
	Version     types.String `tfsdk:"version"`
}

var DtcMonitorSnmpAttrTypes = map[string]attr.Type{
	"ref":          types.StringType,
	"comment":      types.StringType,
	"community":    types.StringType,
	"context":      types.StringType,
	"engine_id":    types.StringType,
	"extattrs":     types.MapType{ElemType: types.StringType},
	"extattrs_all": types.MapType{ElemType: types.StringType},
	"interval":     types.Int64Type,
	"name":         types.StringType,
	"oids":         types.ListType{ElemType: types.ObjectType{AttrTypes: DtcMonitorSnmpOidsAttrTypes}},
	"port":         types.Int64Type,
	"retry_down":   types.Int64Type,
	"retry_up":     types.Int64Type,
	"timeout":      types.Int64Type,
	"user":         types.StringType,
	"version":      types.StringType,
}

var DtcMonitorSnmpResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 256),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Comment for this DTC monitor; maximum 256 characters.",
	},
	"community": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString("public"),
		MarkdownDescription: "The SNMP community string for SNMP authentication.",
	},
	"context": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The SNMPv3 context.",
	},
	"engine_id": schema.StringAttribute{
		Optional: true,
		Computed: true,
		MarkdownDescription: "The SNMPv3 engine identifier.",
	},
	"extattrs": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Extensible attributes associated with the object.",
	},
	"extattrs_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "Extensible attributes associated with the object , including default attributes.",
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
	},
	"interval": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(5),
		MarkdownDescription: "The interval for SNMP health check.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The display name for this DTC monitor.",
	},
	"oids": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: DtcMonitorSnmpOidsResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "A list of OIDs for SNMP monitoring.",
	},
	"port": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(161),
		Validators: []validator.Int64{
			int64validator.Between(1, 65535),
		},
		MarkdownDescription: "The port value for SNMP requests.",
	},
	"retry_down": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(1),
		MarkdownDescription: "The value of how many times the server should appear as down to be treated as dead after it was alive.",
	},
	"retry_up": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(1),
		MarkdownDescription: "The value of how many times the server should appear as up to be treated as alive after it was dead.",
	},
	"timeout": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(15),
		MarkdownDescription: "The timeout for SNMP health check in seconds.",
	},
	"user": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The SNMPv3 user setting.",
	},
	"version": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("V2C"),
		Validators: []validator.String{
			stringvalidator.OneOf("V1", "V2C", "V3"),
		},
		MarkdownDescription: "The SNMP protocol version for the SNMP health check.",
	},
}

func (m *DtcMonitorSnmpModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dtc.DtcMonitorSnmp {
	if m == nil {
		return nil
	}
	to := &dtc.DtcMonitorSnmp{
		Comment:   flex.ExpandStringPointer(m.Comment),
		Community: flex.ExpandStringPointer(m.Community),
		Context:   flex.ExpandStringPointer(m.Context),
		EngineId:  flex.ExpandStringPointer(m.EngineId),
		ExtAttrs:  ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		Interval:  flex.ExpandInt64Pointer(m.Interval),
		Name:      flex.ExpandStringPointer(m.Name),
		Oids:      flex.ExpandFrameworkListNestedBlock(ctx, m.Oids, diags, ExpandDtcMonitorSnmpOids),
		Port:      flex.ExpandInt64Pointer(m.Port),
		RetryDown: flex.ExpandInt64Pointer(m.RetryDown),
		RetryUp:   flex.ExpandInt64Pointer(m.RetryUp),
		Timeout:   flex.ExpandInt64Pointer(m.Timeout),
		User:      flex.ExpandStringPointer(m.User),
		Version:   flex.ExpandStringPointer(m.Version),
	}
	return to
}

func FlattenDtcMonitorSnmp(ctx context.Context, from *dtc.DtcMonitorSnmp, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DtcMonitorSnmpAttrTypes)
	}
	m := DtcMonitorSnmpModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, DtcMonitorSnmpAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DtcMonitorSnmpModel) Flatten(ctx context.Context, from *dtc.DtcMonitorSnmp, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DtcMonitorSnmpModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Community = flex.FlattenStringPointer(from.Community)
	m.Context = flex.FlattenStringPointer(from.Context)
	m.EngineId = flex.FlattenStringPointer(from.EngineId)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.Interval = flex.FlattenInt64Pointer(from.Interval)
	m.Name = flex.FlattenStringPointer(from.Name)
	planOids := m.Oids
	m.Oids = flex.FlattenFrameworkListNestedBlock(ctx, from.Oids, DtcMonitorSnmpOidsAttrTypes, diags, FlattenDtcMonitorSnmpOids)
	if !planOids.IsUnknown() {
		reOrderedList, diags := utils.ReorderAndFilterNestedListResponse(ctx, planOids, m.Oids, "oid")
		if !diags.HasError() {
			m.Oids = reOrderedList.(basetypes.ListValue)
		}
	}
	m.Port = flex.FlattenInt64Pointer(from.Port)
	m.RetryDown = flex.FlattenInt64Pointer(from.RetryDown)
	m.RetryUp = flex.FlattenInt64Pointer(from.RetryUp)
	m.Timeout = flex.FlattenInt64Pointer(from.Timeout)
	m.User = flex.FlattenStringPointerNilAsNotEmpty(from.User)
	m.Version = flex.FlattenStringPointer(from.Version)
}
