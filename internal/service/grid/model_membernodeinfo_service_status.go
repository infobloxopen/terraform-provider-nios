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

type MembernodeinfoServiceStatusModel struct {
	Description types.String `tfsdk:"description"`
	Status      types.String `tfsdk:"status"`
	Service     types.String `tfsdk:"service"`
}

var MembernodeinfoServiceStatusAttrTypes = map[string]attr.Type{
	"description": types.StringType,
	"status":      types.StringType,
	"service":     types.StringType,
}

var MembernodeinfoServiceStatusResourceSchemaAttributes = map[string]schema.Attribute{
	"description": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The description of the current service status.",
	},
	"status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The service status.",
	},
	"service": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The service identifier.",
	},
}

func ExpandMembernodeinfoServiceStatus(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MembernodeinfoServiceStatus {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MembernodeinfoServiceStatusModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MembernodeinfoServiceStatusModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MembernodeinfoServiceStatus {
	if m == nil {
		return nil
	}
	to := &grid.MembernodeinfoServiceStatus{}
	return to
}

func FlattenMembernodeinfoServiceStatus(ctx context.Context, from *grid.MembernodeinfoServiceStatus, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MembernodeinfoServiceStatusAttrTypes)
	}
	m := MembernodeinfoServiceStatusModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MembernodeinfoServiceStatusAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MembernodeinfoServiceStatusModel) Flatten(ctx context.Context, from *grid.MembernodeinfoServiceStatus, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MembernodeinfoServiceStatusModel{}
	}
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Status = flex.FlattenStringPointer(from.Status)
	m.Service = flex.FlattenStringPointer(from.Service)
}
