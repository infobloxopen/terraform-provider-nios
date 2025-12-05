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

type GridUpdatesDownloadMemberConfigModel struct {
	Member    types.String `tfsdk:"member"`
	Interface types.String `tfsdk:"interface"`
	IsOnline  types.Bool   `tfsdk:"is_online"`
}

var GridUpdatesDownloadMemberConfigAttrTypes = map[string]attr.Type{
	"member":    types.StringType,
	"interface": types.StringType,
	"is_online": types.BoolType,
}

var GridUpdatesDownloadMemberConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"member": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the updates download member.",
	},
	"interface": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The source interface for updates download requests.",
	},
	"is_online": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines if the updates download member is online or not.",
	},
}

func ExpandGridUpdatesDownloadMemberConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridUpdatesDownloadMemberConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridUpdatesDownloadMemberConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridUpdatesDownloadMemberConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridUpdatesDownloadMemberConfig {
	if m == nil {
		return nil
	}
	to := &grid.GridUpdatesDownloadMemberConfig{
		Member:    flex.ExpandStringPointer(m.Member),
		Interface: flex.ExpandStringPointer(m.Interface),
	}
	return to
}

func FlattenGridUpdatesDownloadMemberConfig(ctx context.Context, from *grid.GridUpdatesDownloadMemberConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridUpdatesDownloadMemberConfigAttrTypes)
	}
	m := GridUpdatesDownloadMemberConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridUpdatesDownloadMemberConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridUpdatesDownloadMemberConfigModel) Flatten(ctx context.Context, from *grid.GridUpdatesDownloadMemberConfig, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridUpdatesDownloadMemberConfigModel{}
	}
	m.Member = flex.FlattenStringPointer(from.Member)
	m.Interface = flex.FlattenStringPointer(from.Interface)
	m.IsOnline = types.BoolPointerValue(from.IsOnline)
}
