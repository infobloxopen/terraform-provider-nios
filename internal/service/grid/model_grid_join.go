package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	planmodifiers "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/immutable"
)

type GridJoinModel struct {
	MemberUsername types.String `tfsdk:"member_username"`
	MemberPassword types.String `tfsdk:"member_password"`
	MemberIP       types.String `tfsdk:"member_ip"`
	GridName       types.String `tfsdk:"grid_name"`
	Master         types.String `tfsdk:"master"`
	SharedSecret   types.String `tfsdk:"shared_secret"`
}

var GridJoinAttrTypes = map[string]attr.Type{
	"member_username": types.StringType,
	"member_password": types.StringType,
	"member_ip":       types.StringType,
	"grid_name":       types.StringType,
	"master":          types.StringType,
	"shared_secret":   types.StringType,
}

var GridJoinResourceSchemaAttributes = map[string]schema.Attribute{
	"member_username": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
		MarkdownDescription: "The username of the grid member.",
	},
	"member_password": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
		MarkdownDescription: "The password of the grid member.",
	},
	"member_ip": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
		MarkdownDescription: "The IP address of the grid member.",
	},
	"grid_name": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
		MarkdownDescription: "The name of the Grid.",
	},
	"master": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
		MarkdownDescription: "The virtual IP address of the grid master.",
	},
	"shared_secret": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
		MarkdownDescription: "The shared secret string of the grid.",
	},
}

func ExpandGridJoin(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridJoin {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridJoinModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridJoinModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridJoin {
	if m == nil {
		return nil
	}
	to := &grid.GridJoin{
		GridName:     flex.ExpandStringPointer(m.GridName),
		Master:       flex.ExpandStringPointer(m.Master),
		SharedSecret: flex.ExpandStringPointer(m.SharedSecret),
	}
	return to
}
