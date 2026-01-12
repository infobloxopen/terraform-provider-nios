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

type RestartservicestatusModel struct {
	Ref             types.String `tfsdk:"ref"`
	DhcpStatus      types.String `tfsdk:"dhcp_status"`
	DnsStatus       types.String `tfsdk:"dns_status"`
	Member          types.String `tfsdk:"member"`
	ReportingStatus types.String `tfsdk:"reporting_status"`
}

var RestartservicestatusAttrTypes = map[string]attr.Type{
	"ref":              types.StringType,
	"dhcp_status":      types.StringType,
	"dns_status":       types.StringType,
	"member":           types.StringType,
	"reporting_status": types.StringType,
}

var RestartservicestatusResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"dhcp_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The status of the DHCP service.",
	},
	"dns_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The status of the DNS service.",
	},
	"member": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of this Grid member in FQDN format.",
	},
	"reporting_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The status of the reporting service.",
	},
}

func ExpandRestartservicestatus(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.Restartservicestatus {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RestartservicestatusModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *RestartservicestatusModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.Restartservicestatus {
	if m == nil {
		return nil
	}
	to := &grid.Restartservicestatus{
		Ref: flex.ExpandStringPointer(m.Ref),
	}
	return to
}

func FlattenRestartservicestatus(ctx context.Context, from *grid.Restartservicestatus, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RestartservicestatusAttrTypes)
	}
	m := RestartservicestatusModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RestartservicestatusAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RestartservicestatusModel) Flatten(ctx context.Context, from *grid.Restartservicestatus, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RestartservicestatusModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.DhcpStatus = flex.FlattenStringPointer(from.DhcpStatus)
	m.DnsStatus = flex.FlattenStringPointer(from.DnsStatus)
	m.Member = flex.FlattenStringPointer(from.Member)
	m.ReportingStatus = flex.FlattenStringPointer(from.ReportingStatus)
}
