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

type MemberMgmtPortSettingModel struct {
	Enabled               types.Bool `tfsdk:"enabled"`
	VpnEnabled            types.Bool `tfsdk:"vpn_enabled"`
	SecurityAccessEnabled types.Bool `tfsdk:"security_access_enabled"`
}

var MemberMgmtPortSettingAttrTypes = map[string]attr.Type{
	"enabled":                 types.BoolType,
	"vpn_enabled":             types.BoolType,
	"security_access_enabled": types.BoolType,
}

var MemberMgmtPortSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if MGMT port settings should be enabled.",
	},
	"vpn_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if VPN on the MGMT port is enabled or not.",
	},
	"security_access_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if security access on the MGMT port is enabled or not.",
	},
}

func ExpandMemberMgmtPortSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberMgmtPortSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberMgmtPortSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberMgmtPortSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberMgmtPortSetting {
	if m == nil {
		return nil
	}
	to := &grid.MemberMgmtPortSetting{
		Enabled:               flex.ExpandBoolPointer(m.Enabled),
		VpnEnabled:            flex.ExpandBoolPointer(m.VpnEnabled),
		SecurityAccessEnabled: flex.ExpandBoolPointer(m.SecurityAccessEnabled),
	}
	return to
}

func FlattenMemberMgmtPortSetting(ctx context.Context, from *grid.MemberMgmtPortSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberMgmtPortSettingAttrTypes)
	}
	m := MemberMgmtPortSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberMgmtPortSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberMgmtPortSettingModel) Flatten(ctx context.Context, from *grid.MemberMgmtPortSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberMgmtPortSettingModel{}
	}
	m.Enabled = types.BoolPointerValue(from.Enabled)
	m.VpnEnabled = types.BoolPointerValue(from.VpnEnabled)
	m.SecurityAccessEnabled = types.BoolPointerValue(from.SecurityAccessEnabled)
}
