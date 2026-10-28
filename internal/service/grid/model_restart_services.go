package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type RestartServicesModel struct {
	Groups        types.List   `tfsdk:"groups"`
	Members       types.List   `tfsdk:"members"`
	Mode          types.String `tfsdk:"mode"`
	RestartOption types.String `tfsdk:"restart_option"`
	Services      types.List   `tfsdk:"services"`
	UserName      types.String `tfsdk:"user_name"`
}

var RestartServicesActionSchemaAttributes = map[string]schema.Attribute{
	"groups": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "Service Restart Group names to restart. `groups` conflicts with `members`. If both are omitted, NIOS uses the default restart group.",
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
			listvalidator.UniqueValues(),
			listvalidator.ConflictsWith(path.MatchRoot("members")),
		},
	},
	"members": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "Grid Member names to restart. `members` conflicts with `groups`. If both are omitted, NIOS uses the default restart group.",
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
			listvalidator.UniqueValues(),
		},
	},
	"mode": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Grid restart method. Valid values are `GROUPED`, `SEQUENTIAL`, and `SIMULTANEOUS`.",
		Validators: []validator.String{
			stringvalidator.OneOf("GROUPED", "SEQUENTIAL", "SIMULTANEOUS"),
		},
	},
	"restart_option": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Controls whether NIOS restarts services only when needed or unconditionally. Valid values are `RESTART_IF_NEEDED` and `FORCE_RESTART`. Defaults to `RESTART_IF_NEEDED`.",
		Validators: []validator.String{
			stringvalidator.OneOf("RESTART_IF_NEEDED", "FORCE_RESTART"),
		},
	},
	"services": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "Services to restart. Valid values are `ALL`, `DNS`, `DHCP`, `DHCPV4`, and `DHCPV6`. Defaults to `ALL`.",
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
			listvalidator.UniqueValues(),
			listvalidator.ValueStringsAre(
				stringvalidator.OneOf("ALL", "DNS", "DHCP", "DHCPV4", "DHCPV6"),
			),
		},
	},
	"user_name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Name of the user requesting the restart, recorded by NIOS.",
	},
}

func (m *RestartServicesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridFunctionCall {
	if m == nil {
		return nil
	}
	return &grid.GridFunctionCall{
		RestartServiceActionRequest: &grid.RestartServiceActionRequest{
			Groups:        flex.ExpandFrameworkListStringEmptyAsNil(ctx, m.Groups, diags),
			Members:       flex.ExpandFrameworkListStringEmptyAsNil(ctx, m.Members, diags),
			Mode:          flex.ExpandStringPointer(m.Mode),
			RestartOption: flex.ExpandStringPointer(m.RestartOption),
			Services:      flex.ExpandFrameworkListStringEmptyAsNil(ctx, m.Services, diags),
			UserName:      flex.ExpandStringPointer(m.UserName),
		},
	}
}
