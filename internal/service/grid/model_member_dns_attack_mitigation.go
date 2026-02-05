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

type MemberDnsAttackMitigationModel struct {
	DetectChr               types.Object `tfsdk:"detect_chr"`
	DetectChrGrace          types.Int64  `tfsdk:"detect_chr_grace"`
	DetectNxdomainResponses types.Object `tfsdk:"detect_nxdomain_responses"`
	DetectUdpDrop           types.Object `tfsdk:"detect_udp_drop"`
	Interval                types.Int64  `tfsdk:"interval"`
}

var MemberDnsAttackMitigationAttrTypes = map[string]attr.Type{
	"detect_chr":                types.ObjectType{AttrTypes: MemberdnsattackmitigationDetectChrAttrTypes},
	"detect_chr_grace":          types.Int64Type,
	"detect_nxdomain_responses": types.ObjectType{AttrTypes: MemberdnsattackmitigationDetectNxdomainResponsesAttrTypes},
	"detect_udp_drop":           types.ObjectType{AttrTypes: MemberdnsattackmitigationDetectUdpDropAttrTypes},
	"interval":                  types.Int64Type,
}

var MemberDnsAttackMitigationResourceSchemaAttributes = map[string]schema.Attribute{
	"detect_chr": schema.SingleNestedAttribute{
		Attributes: MemberdnsattackmitigationDetectChrResourceSchemaAttributes,
		Optional:   true,
	},
	"detect_chr_grace": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The cache utilization (in percentage) when Cache Hit Ratio (CHR) starts.",
	},
	"detect_nxdomain_responses": schema.SingleNestedAttribute{
		Attributes: MemberdnsattackmitigationDetectNxdomainResponsesResourceSchemaAttributes,
		Optional:   true,
	},
	"detect_udp_drop": schema.SingleNestedAttribute{
		Attributes: MemberdnsattackmitigationDetectUdpDropResourceSchemaAttributes,
		Optional:   true,
	},
	"interval": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The minimum time interval (in seconds) between changes in attack status.",
	},
}

func ExpandMemberDnsAttackMitigation(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberDnsAttackMitigation {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberDnsAttackMitigationModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberDnsAttackMitigationModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberDnsAttackMitigation {
	if m == nil {
		return nil
	}
	to := &grid.MemberDnsAttackMitigation{
		DetectChr:               ExpandMemberdnsattackmitigationDetectChr(ctx, m.DetectChr, diags),
		DetectChrGrace:          flex.ExpandInt64Pointer(m.DetectChrGrace),
		DetectNxdomainResponses: ExpandMemberdnsattackmitigationDetectNxdomainResponses(ctx, m.DetectNxdomainResponses, diags),
		DetectUdpDrop:           ExpandMemberdnsattackmitigationDetectUdpDrop(ctx, m.DetectUdpDrop, diags),
		Interval:                flex.ExpandInt64Pointer(m.Interval),
	}
	return to
}



func FlattenMemberDnsAttackMitigation(ctx context.Context, from *grid.MemberDnsAttackMitigation, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberDnsAttackMitigationAttrTypes)
	}
	m := MemberDnsAttackMitigationModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberDnsAttackMitigationAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberDnsAttackMitigationModel) Flatten(ctx context.Context, from *grid.MemberDnsAttackMitigation, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberDnsAttackMitigationModel{}
	}
	m.DetectChr = FlattenMemberdnsattackmitigationDetectChr(ctx, from.DetectChr, diags)
	m.DetectChrGrace = flex.FlattenInt64Pointer(from.DetectChrGrace)
	m.DetectNxdomainResponses = FlattenMemberdnsattackmitigationDetectNxdomainResponses(ctx, from.DetectNxdomainResponses, diags)
	m.DetectUdpDrop = FlattenMemberdnsattackmitigationDetectUdpDrop(ctx, from.DetectUdpDrop, diags)
	m.Interval = flex.FlattenInt64Pointer(from.Interval)
}
