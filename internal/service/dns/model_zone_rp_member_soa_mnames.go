package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type ZoneRpMemberSoaMnamesModel struct {
	GridPrimary     types.String `tfsdk:"grid_primary"`
	MsServerPrimary types.String `tfsdk:"ms_server_primary"`
	Mname           types.String `tfsdk:"mname"`
	DnsMname        types.String `tfsdk:"dns_mname"`
}

var ZoneRpMemberSoaMnamesAttrTypes = map[string]attr.Type{
	"grid_primary":      types.StringType,
	"ms_server_primary": types.StringType,
	"mname":             types.StringType,
	"dns_mname":         types.StringType,
}

var ZoneRpMemberSoaMnamesResourceSchemaAttributes = map[string]schema.Attribute{
	"grid_primary": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("ms_server_primary")),
		},
		MarkdownDescription: "The grid primary server for the zone. Only one of \"grid_primary\" or \"ms_server_primary\" should be set when modifying or creating the object.",
	},
	"ms_server_primary": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The primary MS server for the zone. Only one of \"grid_primary\" or \"ms_server_primary\" should be set when modifying or creating the object.",
	},
	"mname": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Master's SOA MNAME. This value can be in unicode format.",
	},
	"dns_mname": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Master's SOA MNAME in punycode format.",
	},
}

func ExpandZoneRpMemberSoaMnames(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.ZoneRpMemberSoaMnames {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneRpMemberSoaMnamesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ZoneRpMemberSoaMnamesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.ZoneRpMemberSoaMnames {
	if m == nil {
		return nil
	}
	to := &dns.ZoneRpMemberSoaMnames{
		GridPrimary:     flex.ExpandStringPointer(m.GridPrimary),
		MsServerPrimary: flex.ExpandStringPointer(m.MsServerPrimary),
		Mname:           flex.ExpandStringPointer(m.Mname),
	}
	return to
}

func FlattenZoneRpMemberSoaMnames(ctx context.Context, from *dns.ZoneRpMemberSoaMnames, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneRpMemberSoaMnamesAttrTypes)
	}
	m := ZoneRpMemberSoaMnamesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZoneRpMemberSoaMnamesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ZoneRpMemberSoaMnamesModel) Flatten(ctx context.Context, from *dns.ZoneRpMemberSoaMnames, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ZoneRpMemberSoaMnamesModel{}
	}
	m.GridPrimary = flex.FlattenStringPointer(from.GridPrimary)
	m.MsServerPrimary = flex.FlattenStringPointer(from.MsServerPrimary)
	m.Mname = flex.FlattenStringPointer(from.Mname)
	m.DnsMname = flex.FlattenStringPointer(from.DnsMname)
}
