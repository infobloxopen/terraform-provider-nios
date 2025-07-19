package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
)

type FixedAddressStructModel struct {
	MsServer types.Object `tfsdk:"ms_server"`
}

var FixedAddressStructAttrTypes = map[string]attr.Type{
	"ms_server": types.ObjectType{AttrTypes: FixedAddressStructMsServerAttrTypes},
}

var FixedAddressStructResourceSchemaAttributes = map[string]schema.Attribute{
	"ms_server": schema.SingleNestedAttribute{
		Attributes:          FixedAddressStructMsServerResourceSchemaAttributes,
		Required:            true,
		MarkdownDescription: "Microsoft server information for the fixed address.",
	},
}

func ExpandFixedAddressStruct(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dhcp.FixedAddressStruct {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m FixedAddressStructModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *FixedAddressStructModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.FixedAddressStruct {
	if m == nil {
		return nil
	}
	to := &dhcp.FixedAddressStruct{
		MsServer: *ExpandFixedAddressStructMsServer(ctx, m.MsServer, diags),
	}
	return to
}

func (m *FixedAddressStructModel) Flatten(ctx context.Context, from *dhcp.FixedAddressStruct, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = FixedAddressStructModel{}
	}
	m.MsServer = FlattenFixedAddressStructMsServer(ctx, &from.MsServer, diags)
}
