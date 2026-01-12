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

type GridDnsResponseRateLimitingModel struct {
	EnableRrl          types.Bool  `tfsdk:"enable_rrl"`
	LogOnly            types.Bool  `tfsdk:"log_only"`
	ResponsesPerSecond types.Int64 `tfsdk:"responses_per_second"`
	Window             types.Int64 `tfsdk:"window"`
	Slip               types.Int64 `tfsdk:"slip"`
}

var GridDnsResponseRateLimitingAttrTypes = map[string]attr.Type{
	"enable_rrl":           types.BoolType,
	"log_only":             types.BoolType,
	"responses_per_second": types.Int64Type,
	"window":               types.Int64Type,
	"slip":                 types.Int64Type,
}

var GridDnsResponseRateLimitingResourceSchemaAttributes = map[string]schema.Attribute{
	"enable_rrl": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the response rate limiting is enabled or not.",
	},
	"log_only": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if logging for response rate limiting without dropping any requests is enabled or not.",
	},
	"responses_per_second": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The number of responses per client per second.",
	},
	"window": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The time interval in seconds over which responses are tracked.",
	},
	"slip": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The response rate limiting slip. Note that if slip is not equal to 0 every n-th rate-limited UDP request is sent a truncated response instead of being dropped.",
	},
}

func ExpandGridDnsResponseRateLimiting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridDnsResponseRateLimiting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridDnsResponseRateLimitingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridDnsResponseRateLimitingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridDnsResponseRateLimiting {
	if m == nil {
		return nil
	}
	to := &grid.GridDnsResponseRateLimiting{
		EnableRrl:          flex.ExpandBoolPointer(m.EnableRrl),
		LogOnly:            flex.ExpandBoolPointer(m.LogOnly),
		ResponsesPerSecond: flex.ExpandInt64Pointer(m.ResponsesPerSecond),
		Window:             flex.ExpandInt64Pointer(m.Window),
		Slip:               flex.ExpandInt64Pointer(m.Slip),
	}
	return to
}

func FlattenGridDnsResponseRateLimiting(ctx context.Context, from *grid.GridDnsResponseRateLimiting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridDnsResponseRateLimitingAttrTypes)
	}
	m := GridDnsResponseRateLimitingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridDnsResponseRateLimitingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridDnsResponseRateLimitingModel) Flatten(ctx context.Context, from *grid.GridDnsResponseRateLimiting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridDnsResponseRateLimitingModel{}
	}
	m.EnableRrl = types.BoolPointerValue(from.EnableRrl)
	m.LogOnly = types.BoolPointerValue(from.LogOnly)
	m.ResponsesPerSecond = flex.FlattenInt64Pointer(from.ResponsesPerSecond)
	m.Window = flex.FlattenInt64Pointer(from.Window)
	m.Slip = flex.FlattenInt64Pointer(from.Slip)
}
