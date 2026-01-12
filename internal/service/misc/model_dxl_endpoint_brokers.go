package misc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/misc"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type DxlEndpointBrokersModel struct {
	HostName types.String `tfsdk:"host_name"`
	Address  types.String `tfsdk:"address"`
	Port     types.Int64  `tfsdk:"port"`
	UniqueId types.String `tfsdk:"unique_id"`
}

var DxlEndpointBrokersAttrTypes = map[string]attr.Type{
	"host_name": types.StringType,
	"address":   types.StringType,
	"port":      types.Int64Type,
	"unique_id": types.StringType,
}

var DxlEndpointBrokersResourceSchemaAttributes = map[string]schema.Attribute{
	"host_name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The FQDN for the DXL endpoint broker.",
	},
	"address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The IPv4 Address or IPv6 Address for the DXL endpoint broker.",
	},
	"port": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The communication port for the DXL endpoint broker.",
	},
	"unique_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The unique identifier for the DXL endpoint.",
	},
}

func ExpandDxlEndpointBrokers(ctx context.Context, o types.Object, diags *diag.Diagnostics) *misc.DxlEndpointBrokers {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DxlEndpointBrokersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DxlEndpointBrokersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *misc.DxlEndpointBrokers {
	if m == nil {
		return nil
	}
	to := &misc.DxlEndpointBrokers{
		HostName: flex.ExpandStringPointer(m.HostName),
		Address:  flex.ExpandStringPointer(m.Address),
		Port:     flex.ExpandInt64Pointer(m.Port),
		UniqueId: flex.ExpandStringPointer(m.UniqueId),
	}
	return to
}

func FlattenDxlEndpointBrokers(ctx context.Context, from *misc.DxlEndpointBrokers, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DxlEndpointBrokersAttrTypes)
	}
	m := DxlEndpointBrokersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DxlEndpointBrokersAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DxlEndpointBrokersModel) Flatten(ctx context.Context, from *misc.DxlEndpointBrokers, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DxlEndpointBrokersModel{}
	}
	m.HostName = flex.FlattenStringPointer(from.HostName)
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Port = flex.FlattenInt64Pointer(from.Port)
	m.UniqueId = flex.FlattenStringPointer(from.UniqueId)
}
