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

type GridtrafficcapturerecdnssettingKpiMonitoredDomainsModel struct {
	DomainName types.String `tfsdk:"domain_name"`
	RecordType types.String `tfsdk:"record_type"`
}

var GridtrafficcapturerecdnssettingKpiMonitoredDomainsAttrTypes = map[string]attr.Type{
	"domain_name": types.StringType,
	"record_type": types.StringType,
}

var GridtrafficcapturerecdnssettingKpiMonitoredDomainsResourceSchemaAttributes = map[string]schema.Attribute{
	"domain_name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Domain name (FQDN to Query).",
	},
	"record_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Record type(record to query).",
	},
}

func ExpandGridtrafficcapturerecdnssettingKpiMonitoredDomains(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridtrafficcapturerecdnssettingKpiMonitoredDomains {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridtrafficcapturerecdnssettingKpiMonitoredDomainsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridtrafficcapturerecdnssettingKpiMonitoredDomainsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridtrafficcapturerecdnssettingKpiMonitoredDomains {
	if m == nil {
		return nil
	}
	to := &grid.GridtrafficcapturerecdnssettingKpiMonitoredDomains{
		DomainName: flex.ExpandStringPointer(m.DomainName),
		RecordType: flex.ExpandStringPointer(m.RecordType),
	}
	return to
}

func FlattenGridtrafficcapturerecdnssettingKpiMonitoredDomains(ctx context.Context, from *grid.GridtrafficcapturerecdnssettingKpiMonitoredDomains, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridtrafficcapturerecdnssettingKpiMonitoredDomainsAttrTypes)
	}
	m := GridtrafficcapturerecdnssettingKpiMonitoredDomainsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridtrafficcapturerecdnssettingKpiMonitoredDomainsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridtrafficcapturerecdnssettingKpiMonitoredDomainsModel) Flatten(ctx context.Context, from *grid.GridtrafficcapturerecdnssettingKpiMonitoredDomains, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridtrafficcapturerecdnssettingKpiMonitoredDomainsModel{}
	}
	m.DomainName = flex.FlattenStringPointer(from.DomainName)
	m.RecordType = flex.FlattenStringPointer(from.RecordType)
}
