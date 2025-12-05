package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type GridCloudapiCloudstatisticsModel struct {
	Ref                     types.String `tfsdk:"ref"`
	AllocatedAvailableRatio types.Int64  `tfsdk:"allocated_available_ratio"`
	AllocatedIpCount        types.Int64  `tfsdk:"allocated_ip_count"`
	AvailableIpCount        types.String `tfsdk:"available_ip_count"`
	FixedIpCount            types.Int64  `tfsdk:"fixed_ip_count"`
	FloatingIpCount         types.Int64  `tfsdk:"floating_ip_count"`
	TenantCount             types.Int64  `tfsdk:"tenant_count"`
	TenantIpCount           types.Int64  `tfsdk:"tenant_ip_count"`
	TenantVmCount           types.Int64  `tfsdk:"tenant_vm_count"`
}

var GridCloudapiCloudstatisticsAttrTypes = map[string]attr.Type{
	"ref":                       types.StringType,
	"allocated_available_ratio": types.Int64Type,
	"allocated_ip_count":        types.Int64Type,
	"available_ip_count":        types.StringType,
	"fixed_ip_count":            types.Int64Type,
	"floating_ip_count":         types.Int64Type,
	"tenant_count":              types.Int64Type,
	"tenant_ip_count":           types.Int64Type,
	"tenant_vm_count":           types.Int64Type,
}

var GridCloudapiCloudstatisticsResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"allocated_available_ratio": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Ratio of allocated vs. available IPs",
	},
	"allocated_ip_count": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Total number of IPs allocated by tenants.",
	},
	"available_ip_count": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The total number of IP addresses available to tenants. Only IP addresses in networks that are within a delegation scope are counted.",
	},
	"fixed_ip_count": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The total number of fixed IP addresses currently in use by all tenants in the system.",
	},
	"floating_ip_count": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The total number of floating IP addresses currently in use by all tenants in the system.",
	},
	"tenant_count": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Total number of tenant currently in the system.",
	},
	"tenant_ip_count": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The total number of IP addresses currently in use by all tenants in the system.",
	},
	"tenant_vm_count": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The total number of VMs currently in use by all tenants in the system.",
	},
}

func ExpandGridCloudapiCloudstatistics(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridCloudapiCloudstatistics {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridCloudapiCloudstatisticsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridCloudapiCloudstatisticsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridCloudapiCloudstatistics {
	if m == nil {
		return nil
	}
	to := &grid.GridCloudapiCloudstatistics{
		Ref: flex.ExpandStringPointer(m.Ref),
	}
	return to
}

func FlattenGridCloudapiCloudstatistics(ctx context.Context, from *grid.GridCloudapiCloudstatistics, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridCloudapiCloudstatisticsAttrTypes)
	}
	m := GridCloudapiCloudstatisticsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridCloudapiCloudstatisticsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridCloudapiCloudstatisticsModel) Flatten(ctx context.Context, from *grid.GridCloudapiCloudstatistics, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridCloudapiCloudstatisticsModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.AllocatedAvailableRatio = flex.FlattenInt64Pointer(from.AllocatedAvailableRatio)
	m.AllocatedIpCount = flex.FlattenInt64Pointer(from.AllocatedIpCount)
	m.AvailableIpCount = flex.FlattenStringPointer(from.AvailableIpCount)
	m.FixedIpCount = flex.FlattenInt64Pointer(from.FixedIpCount)
	m.FloatingIpCount = flex.FlattenInt64Pointer(from.FloatingIpCount)
	m.TenantCount = flex.FlattenInt64Pointer(from.TenantCount)
	m.TenantIpCount = flex.FlattenInt64Pointer(from.TenantIpCount)
	m.TenantVmCount = flex.FlattenInt64Pointer(from.TenantVmCount)
}
