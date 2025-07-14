package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type ZoneAuthMemberSoaMnamesModel struct {
	GridPrimary     types.String `tfsdk:"grid_primary"`
	MsServerPrimary types.String `tfsdk:"ms_server_primary"`
	Mname           types.String `tfsdk:"mname"`
	DnsMname        types.String `tfsdk:"dns_mname"`
}

var ZoneAuthMemberSoaMnamesAttrTypes = map[string]attr.Type{
	"grid_primary":      types.StringType,
	"ms_server_primary": types.StringType,
	"mname":             types.StringType,
	"dns_mname":         types.StringType,
}

var ZoneAuthMemberSoaMnamesResourceSchemaAttributes = map[string]schema.Attribute{
	"grid_primary": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The grid primary server for the zone. Only one of \"grid_primary\" or \"ms_server_primary\" should be set when modifying or creating the object.",
	},
	"ms_server_primary": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The primary MS server for the zone. Only one of \"grid_primary\" or \"ms_server_primary\" should be set when modifying or creating the object.",
	},
	"mname": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Master's SOA MNAME. This value can be in unicode format.",
	},
	"dns_mname": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Master's SOA MNAME in punycode format.",
	},
}

func ExpandZoneAuthMemberSoaMnames(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.ZoneAuthMemberSoaMnames {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneAuthMemberSoaMnamesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ZoneAuthMemberSoaMnamesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.ZoneAuthMemberSoaMnames {
	if m == nil {
		return nil
	}
	to := &dns.ZoneAuthMemberSoaMnames{
		GridPrimary:     flex.ExpandStringPointer(m.GridPrimary),
		MsServerPrimary: flex.ExpandStringPointer(m.MsServerPrimary),
		Mname:           flex.ExpandStringPointer(m.Mname),
	}
	return to
}

func FlattenZoneAuthMemberSoaMnames(ctx context.Context, from *dns.ZoneAuthMemberSoaMnames, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneAuthMemberSoaMnamesAttrTypes)
	}
	m := ZoneAuthMemberSoaMnamesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZoneAuthMemberSoaMnamesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ZoneAuthMemberSoaMnamesModel) Flatten(ctx context.Context, from *dns.ZoneAuthMemberSoaMnames, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ZoneAuthMemberSoaMnamesModel{}
	}
	m.GridPrimary = flex.FlattenStringPointer(from.GridPrimary)
	m.MsServerPrimary = flex.FlattenStringPointer(from.MsServerPrimary)
	m.Mname = flex.FlattenStringPointer(from.Mname)
	m.DnsMname = flex.FlattenStringPointer(from.DnsMname)
}
