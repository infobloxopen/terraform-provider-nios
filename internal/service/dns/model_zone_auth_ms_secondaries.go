package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	"github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/plancontrol"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type ZoneAuthMsSecondariesModel struct {
	Address                      iptypes.IPAddress `tfsdk:"address"`
	IsMaster                     types.Bool        `tfsdk:"is_master"`
	NsIp                         types.String      `tfsdk:"ns_ip"`
	NsName                       types.String      `tfsdk:"ns_name"`
	Stealth                      types.Bool        `tfsdk:"stealth"`
	SharedWithMsParentDelegation types.Bool        `tfsdk:"shared_with_ms_parent_delegation"`
}

var ZoneAuthMsSecondariesAttrTypes = map[string]attr.Type{
	"address":                          iptypes.IPAddressType{},
	"is_master":                        types.BoolType,
	"ns_ip":                            types.StringType,
	"ns_name":                          types.StringType,
	"stealth":                          types.BoolType,
	"shared_with_ms_parent_delegation": types.BoolType,
}

var ZoneAuthMsSecondariesResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		CustomType:          iptypes.IPAddressType{},
		Required:            true,
		MarkdownDescription: "The address of the server.",
	},
	"is_master": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "This flag indicates if this server is a synchronization master.",
	},
	"ns_ip": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "This address is used when generating the NS record in the zone, which can be different in case of multihomed hosts.",
	},
	"ns_name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "This name is used when generating the NS record in the zone, which can be different in case of multihomed hosts.",
	},
	"stealth": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Set this flag to hide the NS record for the primary name server from DNS queries.",
	},
	"shared_with_ms_parent_delegation": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "This flag represents whether the name server is shared with the parent Microsoft primary zone's delegation server.",
		PlanModifiers: []planmodifier.Bool{
			plancontrol.UseStateForUnknownBool(),
		},
	},
}

func ExpandZoneAuthMsSecondaries(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.ZoneAuthMsSecondaries {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneAuthMsSecondariesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ZoneAuthMsSecondariesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.ZoneAuthMsSecondaries {
	if m == nil {
		return nil
	}
	to := &dns.ZoneAuthMsSecondaries{
		Address:  flex.ExpandIPAddress(m.Address),
		IsMaster: flex.ExpandBoolPointer(m.IsMaster),
		NsIp:     flex.ExpandStringPointer(m.NsIp),
		NsName:   flex.ExpandStringPointer(m.NsName),
		Stealth:  flex.ExpandBoolPointer(m.Stealth),
	}
	return to
}

func FlattenZoneAuthMsSecondaries(ctx context.Context, from *dns.ZoneAuthMsSecondaries, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneAuthMsSecondariesAttrTypes)
	}
	m := ZoneAuthMsSecondariesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZoneAuthMsSecondariesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ZoneAuthMsSecondariesModel) Flatten(ctx context.Context, from *dns.ZoneAuthMsSecondaries, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ZoneAuthMsSecondariesModel{}
	}
	m.Address = flex.FlattenIPAddress(from.Address)
	m.IsMaster = types.BoolPointerValue(from.IsMaster)
	m.NsIp = flex.FlattenStringPointer(from.NsIp)
	m.NsName = flex.FlattenStringPointer(from.NsName)
	m.Stealth = types.BoolPointerValue(from.Stealth)
	m.SharedWithMsParentDelegation = types.BoolPointerValue(from.SharedWithMsParentDelegation)
}
