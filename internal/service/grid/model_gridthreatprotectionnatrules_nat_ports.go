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

type GridthreatprotectionnatrulesNatPortsModel struct {
	StartPort types.Int64 `tfsdk:"start_port"`
	EndPort   types.Int64 `tfsdk:"end_port"`
	BlockSize types.Int64 `tfsdk:"block_size"`
}

var GridthreatprotectionnatrulesNatPortsAttrTypes = map[string]attr.Type{
	"start_port": types.Int64Type,
	"end_port":   types.Int64Type,
	"block_size": types.Int64Type,
}

var GridthreatprotectionnatrulesNatPortsResourceSchemaAttributes = map[string]schema.Attribute{
	"start_port": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The start port value for the NAT port configuration object.",
	},
	"end_port": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The end port value for the NAT port configuration object.",
	},
	"block_size": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The block size for the NAT Port configuration object.",
	},
}

func ExpandGridthreatprotectionnatrulesNatPorts(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridthreatprotectionnatrulesNatPorts {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridthreatprotectionnatrulesNatPortsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridthreatprotectionnatrulesNatPortsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridthreatprotectionnatrulesNatPorts {
	if m == nil {
		return nil
	}
	to := &grid.GridthreatprotectionnatrulesNatPorts{
		StartPort: flex.ExpandInt64Pointer(m.StartPort),
		EndPort:   flex.ExpandInt64Pointer(m.EndPort),
		BlockSize: flex.ExpandInt64Pointer(m.BlockSize),
	}
	return to
}

func FlattenGridthreatprotectionnatrulesNatPorts(ctx context.Context, from *grid.GridthreatprotectionnatrulesNatPorts, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridthreatprotectionnatrulesNatPortsAttrTypes)
	}
	m := GridthreatprotectionnatrulesNatPortsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridthreatprotectionnatrulesNatPortsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridthreatprotectionnatrulesNatPortsModel) Flatten(ctx context.Context, from *grid.GridthreatprotectionnatrulesNatPorts, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridthreatprotectionnatrulesNatPortsModel{}
	}
	m.StartPort = flex.FlattenInt64Pointer(from.StartPort)
	m.EndPort = flex.FlattenInt64Pointer(from.EndPort)
	m.BlockSize = flex.FlattenInt64Pointer(from.BlockSize)
}
