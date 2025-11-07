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

type GridDnsAttackMitigationModel struct {
	DetectChr               types.Object `tfsdk:"detect_chr"`
	DetectChrGrace          types.Int64  `tfsdk:"detect_chr_grace"`
	DetectNxdomainResponses types.Object `tfsdk:"detect_nxdomain_responses"`
	DetectUdpDrop           types.Object `tfsdk:"detect_udp_drop"`
	Interval                types.Int64  `tfsdk:"interval"`
}

var GridDnsAttackMitigationAttrTypes = map[string]attr.Type{
	"detect_chr":                types.ObjectType{AttrTypes: GriddnsattackmitigationDetectChrAttrTypes},
	"detect_chr_grace":          types.Int64Type,
	"detect_nxdomain_responses": types.ObjectType{AttrTypes: GriddnsattackmitigationDetectNxdomainResponsesAttrTypes},
	"detect_udp_drop":           types.ObjectType{AttrTypes: GriddnsattackmitigationDetectUdpDropAttrTypes},
	"interval":                  types.Int64Type,
}

var GridDnsAttackMitigationResourceSchemaAttributes = map[string]schema.Attribute{
	"detect_chr": schema.SingleNestedAttribute{
		Attributes: GriddnsattackmitigationDetectChrResourceSchemaAttributes,
		Optional:   true,
	},
	"detect_chr_grace": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The cache utilization (in percentage) when Cache Hit Ratio (CHR) starts.",
	},
	"detect_nxdomain_responses": schema.SingleNestedAttribute{
		Attributes: GriddnsattackmitigationDetectNxdomainResponsesResourceSchemaAttributes,
		Optional:   true,
	},
	"detect_udp_drop": schema.SingleNestedAttribute{
		Attributes: GriddnsattackmitigationDetectUdpDropResourceSchemaAttributes,
		Optional:   true,
	},
	"interval": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The minimum time interval (in seconds) between changes in attack status.",
	},
}

func ExpandGridDnsAttackMitigation(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridDnsAttackMitigation {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridDnsAttackMitigationModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridDnsAttackMitigationModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridDnsAttackMitigation {
	if m == nil {
		return nil
	}
	to := &grid.GridDnsAttackMitigation{
		DetectChr:               ExpandGriddnsattackmitigationDetectChr(ctx, m.DetectChr, diags),
		DetectChrGrace:          flex.ExpandInt64Pointer(m.DetectChrGrace),
		DetectNxdomainResponses: ExpandGriddnsattackmitigationDetectNxdomainResponses(ctx, m.DetectNxdomainResponses, diags),
		DetectUdpDrop:           ExpandGriddnsattackmitigationDetectUdpDrop(ctx, m.DetectUdpDrop, diags),
		Interval:                flex.ExpandInt64Pointer(m.Interval),
	}
	return to
}

func FlattenGridDnsAttackMitigation(ctx context.Context, from *grid.GridDnsAttackMitigation, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridDnsAttackMitigationAttrTypes)
	}
	m := GridDnsAttackMitigationModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridDnsAttackMitigationAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridDnsAttackMitigationModel) Flatten(ctx context.Context, from *grid.GridDnsAttackMitigation, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridDnsAttackMitigationModel{}
	}
	m.DetectChr = FlattenGriddnsattackmitigationDetectChr(ctx, from.DetectChr, diags)
	m.DetectChrGrace = flex.FlattenInt64Pointer(from.DetectChrGrace)
	m.DetectNxdomainResponses = FlattenGriddnsattackmitigationDetectNxdomainResponses(ctx, from.DetectNxdomainResponses, diags)
	m.DetectUdpDrop = FlattenGriddnsattackmitigationDetectUdpDrop(ctx, from.DetectUdpDrop, diags)
	m.Interval = flex.FlattenInt64Pointer(from.Interval)
}
