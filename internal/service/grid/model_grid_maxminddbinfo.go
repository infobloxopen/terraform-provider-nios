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

type GridMaxminddbinfoModel struct {
	Ref                types.String `tfsdk:"ref"`
	Uuid               types.String `tfsdk:"uuid"`
	BinaryMajorVersion types.Int64  `tfsdk:"binary_major_version"`
	BinaryMinorVersion types.Int64  `tfsdk:"binary_minor_version"`
	BuildTime          types.Int64  `tfsdk:"build_time"`
	DatabaseType       types.String `tfsdk:"database_type"`
	DeploymentTime     types.Int64  `tfsdk:"deployment_time"`
	Member             types.String `tfsdk:"member"`
	TopologyType       types.String `tfsdk:"topology_type"`
}

var GridMaxminddbinfoAttrTypes = map[string]attr.Type{
	"ref":                  types.StringType,
	"binary_major_version": types.Int64Type,
	"binary_minor_version": types.Int64Type,
	"build_time":           types.Int64Type,
	"database_type":        types.StringType,
	"deployment_time":      types.Int64Type,
	"member":               types.StringType,
	"topology_type":        types.StringType,
}

var GridMaxminddbinfoResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The UUID of the object.",
	},
	"binary_major_version": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The major version of DB binary format.",
	},
	"binary_minor_version": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The minor version of DB binary format.",
	},
	"build_time": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The time at which the DB was built.",
	},
	"database_type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The structure of data records (\"GeoLite2-Country\", GeoLite2-City\", etc.).",
	},
	"deployment_time": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The time at which the current Topology DB was deployed.",
	},
	"member": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The member for testing the connection.",
	},
	"topology_type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The topology type.",
	},
}

func ExpandGridMaxminddbinfo(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridMaxminddbinfo {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridMaxminddbinfoModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridMaxminddbinfoModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridMaxminddbinfo {
	if m == nil {
		return nil
	}
	to := &grid.GridMaxminddbinfo{
		Ref: flex.ExpandStringPointer(m.Ref),
	}
	return to
}

func FlattenGridMaxminddbinfo(ctx context.Context, from *grid.GridMaxminddbinfo, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridMaxminddbinfoAttrTypes)
	}
	m := GridMaxminddbinfoModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridMaxminddbinfoAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridMaxminddbinfoModel) Flatten(ctx context.Context, from *grid.GridMaxminddbinfo, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridMaxminddbinfoModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.BinaryMajorVersion = flex.FlattenInt64Pointer(from.BinaryMajorVersion)
	m.BinaryMinorVersion = flex.FlattenInt64Pointer(from.BinaryMinorVersion)
	m.BuildTime = flex.FlattenInt64Pointer(from.BuildTime)
	m.DatabaseType = flex.FlattenStringPointer(from.DatabaseType)
	m.DeploymentTime = flex.FlattenInt64Pointer(from.DeploymentTime)
	m.Member = flex.FlattenStringPointer(from.Member)
	m.TopologyType = flex.FlattenStringPointer(from.TopologyType)
}
