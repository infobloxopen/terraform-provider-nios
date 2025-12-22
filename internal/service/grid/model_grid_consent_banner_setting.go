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

type GridConsentBannerSettingModel struct {
	Enable  types.Bool   `tfsdk:"enable"`
	Message types.String `tfsdk:"message"`
}

var GridConsentBannerSettingAttrTypes = map[string]attr.Type{
	"enable":  types.BoolType,
	"message": types.StringType,
}

var GridConsentBannerSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the consent banner is enabled.",
	},
	"message": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The message included in the consent banner.",
	},
}

func ExpandGridConsentBannerSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridConsentBannerSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridConsentBannerSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridConsentBannerSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridConsentBannerSetting {
	if m == nil {
		return nil
	}
	to := &grid.GridConsentBannerSetting{
		Enable:  flex.ExpandBoolPointer(m.Enable),
		Message: flex.ExpandStringPointer(m.Message),
	}
	return to
}

func FlattenGridConsentBannerSetting(ctx context.Context, from *grid.GridConsentBannerSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridConsentBannerSettingAttrTypes)
	}
	m := GridConsentBannerSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridConsentBannerSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridConsentBannerSettingModel) Flatten(ctx context.Context, from *grid.GridConsentBannerSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridConsentBannerSettingModel{}
	}
	m.Enable = types.BoolPointerValue(from.Enable)
	m.Message = flex.FlattenStringPointer(from.Message)
}
