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

type GridInformationalBannerSettingModel struct {
	Enable  types.Bool   `tfsdk:"enable"`
	Message types.String `tfsdk:"message"`
	Color   types.String `tfsdk:"color"`
}

var GridInformationalBannerSettingAttrTypes = map[string]attr.Type{
	"enable":  types.BoolType,
	"message": types.StringType,
	"color":   types.StringType,
}

var GridInformationalBannerSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the display of the informational level banner is enabled.",
	},
	"message": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The message included in the informational level banner.",
	},
	"color": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The color for the informational level banner.",
	},
}

func ExpandGridInformationalBannerSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridInformationalBannerSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridInformationalBannerSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridInformationalBannerSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridInformationalBannerSetting {
	if m == nil {
		return nil
	}
	to := &grid.GridInformationalBannerSetting{
		Enable:  flex.ExpandBoolPointer(m.Enable),
		Message: flex.ExpandStringPointer(m.Message),
		Color:   flex.ExpandStringPointer(m.Color),
	}
	return to
}

func FlattenGridInformationalBannerSetting(ctx context.Context, from *grid.GridInformationalBannerSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridInformationalBannerSettingAttrTypes)
	}
	m := GridInformationalBannerSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridInformationalBannerSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridInformationalBannerSettingModel) Flatten(ctx context.Context, from *grid.GridInformationalBannerSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridInformationalBannerSettingModel{}
	}
	m.Enable = types.BoolPointerValue(from.Enable)
	m.Message = flex.FlattenStringPointer(from.Message)
	m.Color = flex.FlattenStringPointer(from.Color)
}
