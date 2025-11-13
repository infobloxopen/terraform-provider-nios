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

type GridCspGridSettingModel struct {
	CspJoinToken   types.String `tfsdk:"csp_join_token"`
	CspDnsResolver types.String `tfsdk:"csp_dns_resolver"`
	CspHttpsProxy  types.String `tfsdk:"csp_https_proxy"`
}

var GridCspGridSettingAttrTypes = map[string]attr.Type{
	"csp_join_token":   types.StringType,
	"csp_dns_resolver": types.StringType,
	"csp_https_proxy":  types.StringType,
}

var GridCspGridSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"csp_join_token": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Join token required to connect to a cluster",
	},
	"csp_dns_resolver": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "IP address of DNS resolver in DFP",
	},
	"csp_https_proxy": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "HTTP Proxy IP address of CSP Portal",
	},
}

func ExpandGridCspGridSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridCspGridSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridCspGridSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridCspGridSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridCspGridSetting {
	if m == nil {
		return nil
	}
	to := &grid.GridCspGridSetting{
		CspJoinToken:   flex.ExpandStringPointer(m.CspJoinToken),
		CspDnsResolver: flex.ExpandStringPointer(m.CspDnsResolver),
		CspHttpsProxy:  flex.ExpandStringPointer(m.CspHttpsProxy),
	}
	return to
}

func FlattenGridCspGridSetting(ctx context.Context, from *grid.GridCspGridSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridCspGridSettingAttrTypes)
	}
	m := GridCspGridSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridCspGridSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridCspGridSettingModel) Flatten(ctx context.Context, from *grid.GridCspGridSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridCspGridSettingModel{}
	}
	m.CspJoinToken = flex.FlattenStringPointer(from.CspJoinToken)
	m.CspDnsResolver = flex.FlattenStringPointer(from.CspDnsResolver)
	m.CspHttpsProxy = flex.FlattenStringPointer(from.CspHttpsProxy)
}
