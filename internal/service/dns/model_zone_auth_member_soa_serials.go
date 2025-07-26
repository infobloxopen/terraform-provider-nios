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

type ZoneAuthMemberSoaSerialsModel struct {
	GridPrimary     types.String `tfsdk:"grid_primary"`
	MsServerPrimary types.String `tfsdk:"ms_server_primary"`
	Serial          types.Int64  `tfsdk:"serial"`
}

var ZoneAuthMemberSoaSerialsAttrTypes = map[string]attr.Type{
	"grid_primary":      types.StringType,
	"ms_server_primary": types.StringType,
	"serial":            types.Int64Type,
}

var ZoneAuthMemberSoaSerialsResourceSchemaAttributes = map[string]schema.Attribute{
	"grid_primary": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The grid primary server for the zone. Only one of \"grid_primary\" or \"ms_server_primary\" will be set when the object is retrieved from the server.",
	},
	"ms_server_primary": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The primary MS server for the zone. Only one of \"grid_primary\" or \"ms_server_primary\" will be set when the object is retrieved from the server.",
	},
	"serial": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The SOA serial number.",
	},
}

func ExpandZoneAuthMemberSoaSerials(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.ZoneAuthMemberSoaSerials {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneAuthMemberSoaSerialsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ZoneAuthMemberSoaSerialsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.ZoneAuthMemberSoaSerials {
	if m == nil {
		return nil
	}
	to := &dns.ZoneAuthMemberSoaSerials{}
	return to
}

func FlattenZoneAuthMemberSoaSerials(ctx context.Context, from *dns.ZoneAuthMemberSoaSerials, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneAuthMemberSoaSerialsAttrTypes)
	}
	m := ZoneAuthMemberSoaSerialsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZoneAuthMemberSoaSerialsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ZoneAuthMemberSoaSerialsModel) Flatten(ctx context.Context, from *dns.ZoneAuthMemberSoaSerials, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ZoneAuthMemberSoaSerialsModel{}
	}
	m.GridPrimary = flex.FlattenStringPointer(from.GridPrimary)
	m.MsServerPrimary = flex.FlattenStringPointer(from.MsServerPrimary)
	m.Serial = flex.FlattenInt64Pointer(from.Serial)
}
