package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type RangetemplateMsServerModel struct {
	Ipv4addr types.String `tfsdk:"ipv4addr"`
}

var RangetemplateMsServerAttrTypes = map[string]attr.Type{
	"ipv4addr": types.StringType,
}

var RangetemplateMsServerResourceSchemaAttributes = map[string]schema.Attribute{
	"ipv4addr": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The IPv4 Address or FQDN of the Microsoft server.",
	},
}

func ExpandRangetemplateMsServer(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dhcp.RangetemplateMsServer {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RangetemplateMsServerModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *RangetemplateMsServerModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.RangetemplateMsServer {
	if m == nil {
		return nil
	}
	to := &dhcp.RangetemplateMsServer{
		Ipv4addr: flex.ExpandStringPointer(m.Ipv4addr),
	}
	return to
}

func FlattenRangetemplateMsServer(ctx context.Context, from *dhcp.RangetemplateMsServer, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RangetemplateMsServerAttrTypes)
	}
	m := RangetemplateMsServerModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RangetemplateMsServerAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RangetemplateMsServerModel) Flatten(ctx context.Context, from *dhcp.RangetemplateMsServer, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RangetemplateMsServerModel{}
	}
	m.Ipv4addr = flex.FlattenStringPointer(from.Ipv4addr)
}
