package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type NetworkcontainerIpamTrapSettingsModel struct {
	EnableEmailWarnings types.Bool `tfsdk:"enable_email_warnings"`
	EnableSnmpWarnings  types.Bool `tfsdk:"enable_snmp_warnings"`
}

var NetworkcontainerIpamTrapSettingsAttrTypes = map[string]attr.Type{
	"enable_email_warnings": types.BoolType,
	"enable_snmp_warnings":  types.BoolType,
}

var NetworkcontainerIpamTrapSettingsResourceSchemaAttributes = map[string]schema.Attribute{
	"enable_email_warnings": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether sending warnings by email is enabled or not.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"enable_snmp_warnings": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether sending warnings by SNMP is enabled or not.",
		Computed:            true,
		Default:             booldefault.StaticBool(true),
	},
}

func ExpandNetworkcontainerIpamTrapSettings(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.NetworkcontainerIpamTrapSettings {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkcontainerIpamTrapSettingsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NetworkcontainerIpamTrapSettingsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.NetworkcontainerIpamTrapSettings {
	if m == nil {
		return nil
	}
	to := &ipam.NetworkcontainerIpamTrapSettings{
		EnableEmailWarnings: flex.ExpandBoolPointer(m.EnableEmailWarnings),
		EnableSnmpWarnings:  flex.ExpandBoolPointer(m.EnableSnmpWarnings),
	}
	return to
}

func FlattenNetworkcontainerIpamTrapSettings(ctx context.Context, from *ipam.NetworkcontainerIpamTrapSettings, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkcontainerIpamTrapSettingsAttrTypes)
	}
	m := NetworkcontainerIpamTrapSettingsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkcontainerIpamTrapSettingsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NetworkcontainerIpamTrapSettingsModel) Flatten(ctx context.Context, from *ipam.NetworkcontainerIpamTrapSettings, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NetworkcontainerIpamTrapSettingsModel{}
	}
	m.EnableEmailWarnings = types.BoolPointerValue(from.EnableEmailWarnings)
	m.EnableSnmpWarnings = types.BoolPointerValue(from.EnableSnmpWarnings)
}
