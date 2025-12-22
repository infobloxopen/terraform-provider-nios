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

type GridSecurityBannerSettingModel struct {
	Color   types.String `tfsdk:"color"`
	Level   types.String `tfsdk:"level"`
	Message types.String `tfsdk:"message"`
	Enable  types.Bool   `tfsdk:"enable"`
}

var GridSecurityBannerSettingAttrTypes = map[string]attr.Type{
	"color":   types.StringType,
	"level":   types.StringType,
	"message": types.StringType,
	"enable":  types.BoolType,
}

var GridSecurityBannerSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"color": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The security level color.",
	},
	"level": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The security level.",
	},
	"message": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The classification message to be displayed.",
	},
	"enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If set to True, the security banner will be displayed on the header and footer of the Grid Manager screen, including the Login screen.",
	},
}

func ExpandGridSecurityBannerSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridSecurityBannerSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridSecurityBannerSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridSecurityBannerSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridSecurityBannerSetting {
	if m == nil {
		return nil
	}
	to := &grid.GridSecurityBannerSetting{
		Color:   flex.ExpandStringPointer(m.Color),
		Level:   flex.ExpandStringPointer(m.Level),
		Message: flex.ExpandStringPointer(m.Message),
		Enable:  flex.ExpandBoolPointer(m.Enable),
	}
	return to
}

func FlattenGridSecurityBannerSetting(ctx context.Context, from *grid.GridSecurityBannerSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridSecurityBannerSettingAttrTypes)
	}
	m := GridSecurityBannerSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridSecurityBannerSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridSecurityBannerSettingModel) Flatten(ctx context.Context, from *grid.GridSecurityBannerSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridSecurityBannerSettingModel{}
	}
	m.Color = flex.FlattenStringPointer(from.Color)
	m.Level = flex.FlattenStringPointer(from.Level)
	m.Message = flex.FlattenStringPointer(from.Message)
	m.Enable = types.BoolPointerValue(from.Enable)
}
