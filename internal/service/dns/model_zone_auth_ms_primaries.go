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

type ZoneAuthMsPrimariesModel struct {
	Address                      types.String `tfsdk:"address"`
	IsMaster                     types.Bool   `tfsdk:"is_master"`
	NsIp                         types.String `tfsdk:"ns_ip"`
	NsName                       types.String `tfsdk:"ns_name"`
	Stealth                      types.Bool   `tfsdk:"stealth"`
	SharedWithMsParentDelegation types.Bool   `tfsdk:"shared_with_ms_parent_delegation"`
}

var ZoneAuthMsPrimariesAttrTypes = map[string]attr.Type{
	"address":                          types.StringType,
	"is_master":                        types.BoolType,
	"ns_ip":                            types.StringType,
	"ns_name":                          types.StringType,
	"stealth":                          types.BoolType,
	"shared_with_ms_parent_delegation": types.BoolType,
}

var ZoneAuthMsPrimariesResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The address of the server.",
	},
	"is_master": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "This flag indicates if this server is a synchronization master.",
	},
	"ns_ip": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "This address is used when generating the NS record in the zone, which can be different in case of multihomed hosts.",
	},
	"ns_name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "This name is used when generating the NS record in the zone, which can be different in case of multihomed hosts.",
	},
	"stealth": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Set this flag to hide the NS record for the primary name server from DNS queries.",
	},
	"shared_with_ms_parent_delegation": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "This flag represents whether the name server is shared with the parent Microsoft primary zone's delegation server.",
	},
}

func ExpandZoneAuthMsPrimaries(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.ZoneAuthMsPrimaries {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneAuthMsPrimariesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ZoneAuthMsPrimariesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.ZoneAuthMsPrimaries {
	if m == nil {
		return nil
	}
	to := &dns.ZoneAuthMsPrimaries{
		Address:  flex.ExpandStringPointer(m.Address),
		IsMaster: flex.ExpandBoolPointer(m.IsMaster),
		NsIp:     flex.ExpandStringPointer(m.NsIp),
		NsName:   flex.ExpandStringPointer(m.NsName),
		Stealth:  flex.ExpandBoolPointer(m.Stealth),
	}
	return to
}

func FlattenZoneAuthMsPrimaries(ctx context.Context, from *dns.ZoneAuthMsPrimaries, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneAuthMsPrimariesAttrTypes)
	}
	m := ZoneAuthMsPrimariesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZoneAuthMsPrimariesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ZoneAuthMsPrimariesModel) Flatten(ctx context.Context, from *dns.ZoneAuthMsPrimaries, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ZoneAuthMsPrimariesModel{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.IsMaster = types.BoolPointerValue(from.IsMaster)
	m.NsIp = flex.FlattenStringPointer(from.NsIp)
	m.NsName = flex.FlattenStringPointer(from.NsName)
	m.Stealth = types.BoolPointerValue(from.Stealth)
	m.SharedWithMsParentDelegation = types.BoolPointerValue(from.SharedWithMsParentDelegation)
}
