package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dtc"

	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/flex"
)

type DtcServerMonitorsModel struct {
	Monitor types.String `tfsdk:"monitor"`
	Host    types.String `tfsdk:"host"`
}

var DtcServerMonitorsAttrTypes = map[string]attr.Type{
	"monitor": types.StringType,
	"host":    types.StringType,
}

var DtcServerMonitorsResourceSchemaAttributes = map[string]schema.Attribute{
	"monitor": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The monitor related to server.",
	},
	"host": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "IP address or FQDN of the server used for monitoring.",
	},
}

func ExpandDtcServerMonitors(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dtc.DtcServerMonitors {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DtcServerMonitorsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DtcServerMonitorsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dtc.DtcServerMonitors {
	if m == nil {
		return nil
	}
	to := &dtc.DtcServerMonitors{
		Monitor: flex.ExpandStringPointer(m.Monitor),
		Host:    flex.ExpandStringPointer(m.Host),
	}
	return to
}

func FlattenDtcServerMonitors(ctx context.Context, from *dtc.DtcServerMonitors, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DtcServerMonitorsAttrTypes)
	}
	m := DtcServerMonitorsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DtcServerMonitorsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DtcServerMonitorsModel) Flatten(ctx context.Context, from *dtc.DtcServerMonitors, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DtcServerMonitorsModel{}
	}
	m.Monitor = flex.FlattenStringPointer(from.Monitor)
	m.Host = flex.FlattenStringPointer(from.Host)
}
