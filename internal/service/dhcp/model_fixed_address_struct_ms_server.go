package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type FixedAddressStructMsServerModel struct {
	Struct   types.String `tfsdk:"struct"`
	Ipv4addr types.String `tfsdk:"ipv4addr"`
}

var FixedAddressStructMsServerAttrTypes = map[string]attr.Type{
	"struct":   types.StringType,
	"ipv4addr": types.StringType,
}

var FixedAddressStructMsServerResourceSchemaAttributes = map[string]schema.Attribute{
	"struct": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Struct Field for MS Server",
	},
	"ipv4addr": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The IPv4 address to perform the search with",
	},
}

func ExpandFixedAddressStructMsServer(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dhcp.FixedAddressStructMsServer {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m FixedAddressStructMsServerModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *FixedAddressStructMsServerModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.FixedAddressStructMsServer {
	if m == nil {
		return nil
	}
	to := &dhcp.FixedAddressStructMsServer{
		Struct:   flex.ExpandStringPointer(m.Struct),
		Ipv4addr: flex.ExpandStringPointer(m.Ipv4addr),
	}
	return to
}

func FlattenFixedAddressStructMsServer(ctx context.Context, from *dhcp.FixedAddressStructMsServer, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(FixedAddressStructMsServerAttrTypes)
	}
	m := FixedAddressStructMsServerModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, FixedAddressStructMsServerAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *FixedAddressStructMsServerModel) Flatten(ctx context.Context, from *dhcp.FixedAddressStructMsServer, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = FixedAddressStructMsServerModel{}
	}
	m.Struct = flex.FlattenStringPointer(from.Struct)
	m.Ipv4addr = flex.FlattenStringPointer(from.Ipv4addr)
}
