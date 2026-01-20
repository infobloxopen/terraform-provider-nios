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

type GridCloudapiTenantModel struct {
	Ref          types.String `tfsdk:"ref"`
	Uuid         types.String `tfsdk:"uuid"`
	CloudInfo    types.Object `tfsdk:"cloud_info"`
	Comment      types.String `tfsdk:"comment"`
	CreatedTs    types.Int64  `tfsdk:"created_ts"`
	Id           types.String `tfsdk:"id"`
	LastEventTs  types.Int64  `tfsdk:"last_event_ts"`
	Name         types.String `tfsdk:"name"`
	NetworkCount types.Int64  `tfsdk:"network_count"`
	VmCount      types.Int64  `tfsdk:"vm_count"`
}

var GridCloudapiTenantAttrTypes = map[string]attr.Type{
	"ref":           types.StringType,
	"uuid":          types.StringType,
	"cloud_info":    types.ObjectType{AttrTypes: GridCloudapiTenantCloudInfoAttrTypes},
	"comment":       types.StringType,
	"created_ts":    types.Int64Type,
	"id":            types.StringType,
	"last_event_ts": types.Int64Type,
	"name":          types.StringType,
	"network_count": types.Int64Type,
	"vm_count":      types.Int64Type,
}

var GridCloudapiTenantResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The universally unique identifier (UUID) for the tenant.",
	},
	"cloud_info": schema.SingleNestedAttribute{
		Attributes: GridCloudapiTenantCloudInfoResourceSchemaAttributes,
		Optional:   true,
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Comment for the Grid Cloud API Tenant object; maximum 256 characters.",
	},
	"created_ts": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The timestamp when the tenant was first created in the system.",
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Unique ID associated with the tenant. This is set only when the tenant is first created.",
	},
	"last_event_ts": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The timestamp when the last event associated with the tenant happened.",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Name of the tenant.",
	},
	"network_count": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Number of Networks associated with the tenant.",
	},
	"vm_count": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Number of VMs associated with the tenant.",
	},
}

func ExpandGridCloudapiTenant(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridCloudapiTenant {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridCloudapiTenantModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridCloudapiTenantModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridCloudapiTenant {
	if m == nil {
		return nil
	}
	to := &grid.GridCloudapiTenant{
		Ref:       flex.ExpandStringPointer(m.Ref),
		Uuid:      flex.ExpandStringPointer(m.Uuid),
		CloudInfo: ExpandGridCloudapiTenantCloudInfo(ctx, m.CloudInfo, diags),
		Comment:   flex.ExpandStringPointer(m.Comment),
		Name:      flex.ExpandStringPointer(m.Name),
	}
	return to
}

func FlattenGridCloudapiTenant(ctx context.Context, from *grid.GridCloudapiTenant, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridCloudapiTenantAttrTypes)
	}
	m := GridCloudapiTenantModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridCloudapiTenantAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridCloudapiTenantModel) Flatten(ctx context.Context, from *grid.GridCloudapiTenant, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridCloudapiTenantModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.CloudInfo = FlattenGridCloudapiTenantCloudInfo(ctx, from.CloudInfo, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CreatedTs = flex.FlattenInt64Pointer(from.CreatedTs)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.LastEventTs = flex.FlattenInt64Pointer(from.LastEventTs)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.NetworkCount = flex.FlattenInt64Pointer(from.NetworkCount)
	m.VmCount = flex.FlattenInt64Pointer(from.VmCount)
}
