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

type GridCspApiConfigModel struct {
	Url      types.String `tfsdk:"url"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}

var GridCspApiConfigAttrTypes = map[string]attr.Type{
	"url":      types.StringType,
	"username": types.StringType,
	"password": types.StringType,
}

var GridCspApiConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"url": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The url for the CspApiConfig.",
	},
	"username": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The user name for the CspApiConfig.",
	},
	"password": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The password for the CspApiConfig.",
	},
}

func ExpandGridCspApiConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridCspApiConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridCspApiConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridCspApiConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridCspApiConfig {
	if m == nil {
		return nil
	}
	to := &grid.GridCspApiConfig{
		Username: flex.ExpandStringPointer(m.Username),
		Password: flex.ExpandStringPointer(m.Password),
	}
	return to
}

func FlattenGridCspApiConfig(ctx context.Context, from *grid.GridCspApiConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridCspApiConfigAttrTypes)
	}
	m := GridCspApiConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridCspApiConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridCspApiConfigModel) Flatten(ctx context.Context, from *grid.GridCspApiConfig, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridCspApiConfigModel{}
	}
	m.Url = flex.FlattenStringPointer(from.Url)
	m.Username = flex.FlattenStringPointer(from.Username)
	m.Password = flex.FlattenStringPointer(from.Password)
}
