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

type RangetemplateExcludeModel struct {
	Offset            types.Int64  `tfsdk:"offset"`
	NumberOfAddresses types.Int64  `tfsdk:"number_of_addresses"`
	Comment           types.String `tfsdk:"comment"`
}

var RangetemplateExcludeAttrTypes = map[string]attr.Type{
	"offset":              types.Int64Type,
	"number_of_addresses": types.Int64Type,
	"comment":             types.StringType,
}

var RangetemplateExcludeResourceSchemaAttributes = map[string]schema.Attribute{
	"offset": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The address offset of the DHCP exclusion range template.",
	},
	"number_of_addresses": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The number of addresses in the DHCP exclusion range template.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "A descriptive comment of a DHCP exclusion range template.",
	},
}

func ExpandRangetemplateExclude(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dhcp.RangetemplateExclude {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RangetemplateExcludeModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *RangetemplateExcludeModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.RangetemplateExclude {
	if m == nil {
		return nil
	}
	to := &dhcp.RangetemplateExclude{
		Offset:            flex.ExpandInt64Pointer(m.Offset),
		NumberOfAddresses: flex.ExpandInt64Pointer(m.NumberOfAddresses),
		Comment:           flex.ExpandStringPointer(m.Comment),
	}
	return to
}

func FlattenRangetemplateExclude(ctx context.Context, from *dhcp.RangetemplateExclude, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RangetemplateExcludeAttrTypes)
	}
	m := RangetemplateExcludeModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RangetemplateExcludeAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RangetemplateExcludeModel) Flatten(ctx context.Context, from *dhcp.RangetemplateExclude, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RangetemplateExcludeModel{}
	}
	m.Offset = flex.FlattenInt64Pointer(from.Offset)
	m.NumberOfAddresses = flex.FlattenInt64Pointer(from.NumberOfAddresses)
	m.Comment = flex.FlattenStringPointer(from.Comment)
}
