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

type MembernodeinfoLan2PhysicalSettingModel struct {
	AutoPortSettingEnabled types.Bool   `tfsdk:"auto_port_setting_enabled"`
	Speed                  types.String `tfsdk:"speed"`
	Duplex                 types.String `tfsdk:"duplex"`
}

var MembernodeinfoLan2PhysicalSettingAttrTypes = map[string]attr.Type{
	"auto_port_setting_enabled": types.BoolType,
	"speed":                     types.StringType,
	"duplex":                    types.StringType,
}

var MembernodeinfoLan2PhysicalSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"auto_port_setting_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enable or disalbe the auto port setting.",
	},
	"speed": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The port speed; if speed is 1000, duplex is FULL.",
	},
	"duplex": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The port duplex; if speed is 1000, duplex must be FULL.",
	},
}

func ExpandMembernodeinfoLan2PhysicalSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MembernodeinfoLan2PhysicalSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MembernodeinfoLan2PhysicalSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MembernodeinfoLan2PhysicalSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MembernodeinfoLan2PhysicalSetting {
	if m == nil {
		return nil
	}
	to := &grid.MembernodeinfoLan2PhysicalSetting{
		AutoPortSettingEnabled: flex.ExpandBoolPointer(m.AutoPortSettingEnabled),
		Speed:                  flex.ExpandStringPointer(m.Speed),
		Duplex:                 flex.ExpandStringPointer(m.Duplex),
	}
	return to
}

func FlattenMembernodeinfoLan2PhysicalSetting(ctx context.Context, from *grid.MembernodeinfoLan2PhysicalSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MembernodeinfoLan2PhysicalSettingAttrTypes)
	}
	m := MembernodeinfoLan2PhysicalSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MembernodeinfoLan2PhysicalSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MembernodeinfoLan2PhysicalSettingModel) Flatten(ctx context.Context, from *grid.MembernodeinfoLan2PhysicalSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MembernodeinfoLan2PhysicalSettingModel{}
	}
	m.AutoPortSettingEnabled = types.BoolPointerValue(from.AutoPortSettingEnabled)
	m.Speed = flex.FlattenStringPointer(from.Speed)
	m.Duplex = flex.FlattenStringPointer(from.Duplex)
}
