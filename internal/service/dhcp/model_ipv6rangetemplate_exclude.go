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

type Ipv6rangetemplateExcludeModel struct {
	Offset            types.Int64  `tfsdk:"offset"`
	NumberOfAddresses types.Int64  `tfsdk:"number_of_addresses"`
	Comment           types.String `tfsdk:"comment"`
}

var Ipv6rangetemplateExcludeAttrTypes = map[string]attr.Type{
	"offset":              types.Int64Type,
	"number_of_addresses": types.Int64Type,
	"comment":             types.StringType,
}

var Ipv6rangetemplateExcludeResourceSchemaAttributes = map[string]schema.Attribute{
	"offset": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The address offset of the DHCP exclusion range template.",
	},
	"number_of_addresses": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The number of addresses in the DHCP exclusion range template.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "A descriptive comment of a DHCP exclusion range template.",
	},
}

func ExpandIpv6rangetemplateExclude(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dhcp.Ipv6rangetemplateExclude {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6rangetemplateExcludeModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *Ipv6rangetemplateExcludeModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.Ipv6rangetemplateExclude {
	if m == nil {
		return nil
	}
	to := &dhcp.Ipv6rangetemplateExclude{
		Offset:            flex.ExpandInt64Pointer(m.Offset),
		NumberOfAddresses: flex.ExpandInt64Pointer(m.NumberOfAddresses),
		Comment:           flex.ExpandStringPointer(m.Comment),
	}
	return to
}

func FlattenIpv6rangetemplateExclude(ctx context.Context, from *dhcp.Ipv6rangetemplateExclude, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6rangetemplateExcludeAttrTypes)
	}
	m := Ipv6rangetemplateExcludeModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, Ipv6rangetemplateExcludeAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Ipv6rangetemplateExcludeModel) Flatten(ctx context.Context, from *dhcp.Ipv6rangetemplateExclude, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Ipv6rangetemplateExcludeModel{}
	}
	m.Offset = flex.FlattenInt64Pointer(from.Offset)
	m.NumberOfAddresses = flex.FlattenInt64Pointer(from.NumberOfAddresses)
	m.Comment = flex.FlattenStringPointer(from.Comment)
}
