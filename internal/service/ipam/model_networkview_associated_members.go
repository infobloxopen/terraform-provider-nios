package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type NetworkviewAssociatedMembersModel struct {
	Member    types.String `tfsdk:"member"`
	Failovers types.List   `tfsdk:"failovers"`
}

var NetworkviewAssociatedMembersAttrTypes = map[string]attr.Type{
	"member":    types.StringType,
	"failovers": types.ListType{ElemType: types.StringType},
}

var NetworkviewAssociatedMembersResourceSchemaAttributes = map[string]schema.Attribute{
	"member": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The member object associated with a network view.",
	},
	"failovers": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The list of failover objects associated with each member.",
	},
}

func ExpandNetworkviewAssociatedMembers(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.NetworkviewAssociatedMembers {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkviewAssociatedMembersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NetworkviewAssociatedMembersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.NetworkviewAssociatedMembers {
	if m == nil {
		return nil
	}
	to := &ipam.NetworkviewAssociatedMembers{}
	return to
}

func FlattenNetworkviewAssociatedMembers(ctx context.Context, from *ipam.NetworkviewAssociatedMembers, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkviewAssociatedMembersAttrTypes)
	}
	m := NetworkviewAssociatedMembersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkviewAssociatedMembersAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NetworkviewAssociatedMembersModel) Flatten(ctx context.Context, from *ipam.NetworkviewAssociatedMembers, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NetworkviewAssociatedMembersModel{}
	}
	m.Member = flex.FlattenStringPointer(from.Member)
	m.Failovers = flex.FlattenFrameworkListString(ctx, from.Failovers, diags)
}
