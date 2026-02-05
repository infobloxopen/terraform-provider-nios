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

type MembertrafficcapturerecdnssettingKpiMonitoredDomainsModel struct {
	DomainName types.String `tfsdk:"domain_name"`
	RecordType types.String `tfsdk:"record_type"`
}

var MembertrafficcapturerecdnssettingKpiMonitoredDomainsAttrTypes = map[string]attr.Type{
	"domain_name": types.StringType,
	"record_type": types.StringType,
}

var MembertrafficcapturerecdnssettingKpiMonitoredDomainsResourceSchemaAttributes = map[string]schema.Attribute{
	"domain_name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Domain name (FQDN to Query).",
	},
	"record_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Record type(record to query).",
	},
}

func ExpandMembertrafficcapturerecdnssettingKpiMonitoredDomains(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MembertrafficcapturerecdnssettingKpiMonitoredDomains {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MembertrafficcapturerecdnssettingKpiMonitoredDomainsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MembertrafficcapturerecdnssettingKpiMonitoredDomainsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MembertrafficcapturerecdnssettingKpiMonitoredDomains {
	if m == nil {
		return nil
	}
	to := &grid.MembertrafficcapturerecdnssettingKpiMonitoredDomains{
		DomainName: flex.ExpandStringPointer(m.DomainName),
		RecordType: flex.ExpandStringPointer(m.RecordType),
	}
	return to
}

func FlattenMembertrafficcapturerecdnssettingKpiMonitoredDomains(ctx context.Context, from *grid.MembertrafficcapturerecdnssettingKpiMonitoredDomains, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MembertrafficcapturerecdnssettingKpiMonitoredDomainsAttrTypes)
	}
	m := MembertrafficcapturerecdnssettingKpiMonitoredDomainsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MembertrafficcapturerecdnssettingKpiMonitoredDomainsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MembertrafficcapturerecdnssettingKpiMonitoredDomainsModel) Flatten(ctx context.Context, from *grid.MembertrafficcapturerecdnssettingKpiMonitoredDomains, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MembertrafficcapturerecdnssettingKpiMonitoredDomainsModel{}
	}
	m.DomainName = flex.FlattenStringPointer(from.DomainName)
	m.RecordType = flex.FlattenStringPointer(from.RecordType)
}
