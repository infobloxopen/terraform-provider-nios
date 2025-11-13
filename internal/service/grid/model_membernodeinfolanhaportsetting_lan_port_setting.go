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

type MembernodeinfolanhaportsettingLanPortSettingModel struct {
	AutoPortSettingEnabled types.Bool   `tfsdk:"auto_port_setting_enabled"`
	Speed                  types.String `tfsdk:"speed"`
	Duplex                 types.String `tfsdk:"duplex"`
}

var MembernodeinfolanhaportsettingLanPortSettingAttrTypes = map[string]attr.Type{
	"auto_port_setting_enabled": types.BoolType,
	"speed":                     types.StringType,
	"duplex":                    types.StringType,
}

var MembernodeinfolanhaportsettingLanPortSettingResourceSchemaAttributes = map[string]schema.Attribute{
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

func ExpandMembernodeinfolanhaportsettingLanPortSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MembernodeinfolanhaportsettingLanPortSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MembernodeinfolanhaportsettingLanPortSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MembernodeinfolanhaportsettingLanPortSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MembernodeinfolanhaportsettingLanPortSetting {
	if m == nil {
		return nil
	}
	to := &grid.MembernodeinfolanhaportsettingLanPortSetting{
		AutoPortSettingEnabled: flex.ExpandBoolPointer(m.AutoPortSettingEnabled),
		Speed:                  flex.ExpandStringPointer(m.Speed),
		Duplex:                 flex.ExpandStringPointer(m.Duplex),
	}
	return to
}

func FlattenMembernodeinfolanhaportsettingLanPortSetting(ctx context.Context, from *grid.MembernodeinfolanhaportsettingLanPortSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MembernodeinfolanhaportsettingLanPortSettingAttrTypes)
	}
	m := MembernodeinfolanhaportsettingLanPortSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MembernodeinfolanhaportsettingLanPortSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MembernodeinfolanhaportsettingLanPortSettingModel) Flatten(ctx context.Context, from *grid.MembernodeinfolanhaportsettingLanPortSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MembernodeinfolanhaportsettingLanPortSettingModel{}
	}
	m.AutoPortSettingEnabled = types.BoolPointerValue(from.AutoPortSettingEnabled)
	m.Speed = flex.FlattenStringPointer(from.Speed)
	m.Duplex = flex.FlattenStringPointer(from.Duplex)
}
