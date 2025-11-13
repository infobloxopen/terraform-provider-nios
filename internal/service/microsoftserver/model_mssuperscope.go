package microsoftserver

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/microsoftserver"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/import"
)

type MssuperscopeModel struct {
	Ref                   types.String `tfsdk:"ref"`
	Comment               types.String `tfsdk:"comment"`
	DhcpUtilization       types.Int64  `tfsdk:"dhcp_utilization"`
	DhcpUtilizationStatus types.String `tfsdk:"dhcp_utilization_status"`
	Disable               types.Bool   `tfsdk:"disable"`
	DynamicHosts          types.Int64  `tfsdk:"dynamic_hosts"`
	ExtAttrs              types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll           types.Map    `tfsdk:"extattrs_all"`
	HighWaterMark         types.Int64  `tfsdk:"high_water_mark"`
	HighWaterMarkReset    types.Int64  `tfsdk:"high_water_mark_reset"`
	LowWaterMark          types.Int64  `tfsdk:"low_water_mark"`
	LowWaterMarkReset     types.Int64  `tfsdk:"low_water_mark_reset"`
	Name                  types.String `tfsdk:"name"`
	NetworkView           types.String `tfsdk:"network_view"`
	Ranges                types.List   `tfsdk:"ranges"`
	StaticHosts           types.Int64  `tfsdk:"static_hosts"`
	TotalHosts            types.Int64  `tfsdk:"total_hosts"`
}

var MssuperscopeAttrTypes = map[string]attr.Type{
	"ref":                     types.StringType,
	"comment":                 types.StringType,
	"dhcp_utilization":        types.Int64Type,
	"dhcp_utilization_status": types.StringType,
	"disable":                 types.BoolType,
	"dynamic_hosts":           types.Int64Type,
	"extattrs":                types.MapType{ElemType: types.StringType},
	"extattrs_all":            types.MapType{ElemType: types.StringType},
	"high_water_mark":         types.Int64Type,
	"high_water_mark_reset":   types.Int64Type,
	"low_water_mark":          types.Int64Type,
	"low_water_mark_reset":    types.Int64Type,
	"name":                    types.StringType,
	"network_view":            types.StringType,
	"ranges":                  types.ListType{ElemType: types.StringType},
	"static_hosts":            types.Int64Type,
	"total_hosts":             types.Int64Type,
}

var MssuperscopeResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The superscope descriptive comment.",
	},
	"dhcp_utilization": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The percentage of the total DHCP usage of the ranges in the superscope.",
	},
	"dhcp_utilization_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Utilization level of the DHCP range objects that belong to the superscope.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the superscope is disabled.",
	},
	"dynamic_hosts": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The total number of DHCP leases issued for the DHCP range objects that belong to the superscope.",
	},
	"extattrs": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.",
	},
	"extattrs_all": schema.MapAttribute{
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object , including default attributes.",
		ElementType:         types.StringType,
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
	},
	"high_water_mark": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The percentage value for DHCP range usage after which an alarm will be active.",
	},
	"high_water_mark_reset": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The percentage value for DHCP range usage after which an alarm will be reset.",
	},
	"low_water_mark": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The percentage value for DHCP range usage below which an alarm will be active.",
	},
	"low_water_mark_reset": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The percentage value for DHCP range usage below which an alarm will be reset.",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the Microsoft DHCP superscope.",
	},
	"network_view": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the network view in which the superscope resides.",
	},
	"ranges": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of DHCP ranges that are associated with the superscope.",
	},
	"static_hosts": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of static DHCP addresses configured in DHCP range objects that belong to the superscope.",
	},
	"total_hosts": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The total number of DHCP addresses configured in DHCP range objects that belong to the superscope.",
	},
}

func ExpandMssuperscope(ctx context.Context, o types.Object, diags *diag.Diagnostics) *microsoftserver.Mssuperscope {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MssuperscopeModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MssuperscopeModel) Expand(ctx context.Context, diags *diag.Diagnostics) *microsoftserver.Mssuperscope {
	if m == nil {
		return nil
	}
	to := &microsoftserver.Mssuperscope{
		Ref:         flex.ExpandStringPointer(m.Ref),
		Comment:     flex.ExpandStringPointer(m.Comment),
		Disable:     flex.ExpandBoolPointer(m.Disable),
		ExtAttrs:    ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		Name:        flex.ExpandStringPointer(m.Name),
		NetworkView: flex.ExpandStringPointer(m.NetworkView),
		Ranges:      flex.ExpandFrameworkListString(ctx, m.Ranges, diags),
	}
	return to
}

func FlattenMssuperscope(ctx context.Context, from *microsoftserver.Mssuperscope, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MssuperscopeAttrTypes)
	}
	m := MssuperscopeModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, MssuperscopeAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MssuperscopeModel) Flatten(ctx context.Context, from *microsoftserver.Mssuperscope, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MssuperscopeModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.DhcpUtilization = flex.FlattenInt64Pointer(from.DhcpUtilization)
	m.DhcpUtilizationStatus = flex.FlattenStringPointer(from.DhcpUtilizationStatus)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.DynamicHosts = flex.FlattenInt64Pointer(from.DynamicHosts)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.HighWaterMark = flex.FlattenInt64Pointer(from.HighWaterMark)
	m.HighWaterMarkReset = flex.FlattenInt64Pointer(from.HighWaterMarkReset)
	m.LowWaterMark = flex.FlattenInt64Pointer(from.LowWaterMark)
	m.LowWaterMarkReset = flex.FlattenInt64Pointer(from.LowWaterMarkReset)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.NetworkView = flex.FlattenStringPointer(from.NetworkView)
	m.Ranges = flex.FlattenFrameworkListString(ctx, from.Ranges, diags)
	m.StaticHosts = flex.FlattenInt64Pointer(from.StaticHosts)
	m.TotalHosts = flex.FlattenInt64Pointer(from.TotalHosts)
}
