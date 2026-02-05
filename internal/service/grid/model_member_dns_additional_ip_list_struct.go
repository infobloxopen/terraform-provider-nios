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

type MemberDnsAdditionalIpListStructModel struct {
	IpAddress types.String `tfsdk:"ip_address"`
}

var MemberDnsAdditionalIpListStructAttrTypes = map[string]attr.Type{
	"ip_address": types.StringType,
}

var MemberDnsAdditionalIpListStructResourceSchemaAttributes = map[string]schema.Attribute{
	"ip_address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The additional IP address of the member.",
	},
}

func ExpandMemberDnsAdditionalIpListStruct(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberDnsAdditionalIpListStruct {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberDnsAdditionalIpListStructModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberDnsAdditionalIpListStructModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberDnsAdditionalIpListStruct {
	if m == nil {
		return nil
	}
	to := &grid.MemberDnsAdditionalIpListStruct{
		IpAddress: flex.ExpandStringPointer(m.IpAddress),
	}
	return to
}

func FlattenMemberDnsAdditionalIpListStruct(ctx context.Context, from *grid.MemberDnsAdditionalIpListStruct, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberDnsAdditionalIpListStructAttrTypes)
	}
	m := MemberDnsAdditionalIpListStructModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberDnsAdditionalIpListStructAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberDnsAdditionalIpListStructModel) Flatten(ctx context.Context, from *grid.MemberDnsAdditionalIpListStruct, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberDnsAdditionalIpListStructModel{}
	}
	m.IpAddress = flex.FlattenStringPointer(from.IpAddress)
}
