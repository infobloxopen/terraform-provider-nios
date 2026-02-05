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

type MemberLomNetworkConfigModel struct {
	Address      types.String `tfsdk:"address"`
	Gateway      types.String `tfsdk:"gateway"`
	SubnetMask   types.String `tfsdk:"subnet_mask"`
	IsLomCapable types.Bool   `tfsdk:"is_lom_capable"`
}

var MemberLomNetworkConfigAttrTypes = map[string]attr.Type{
	"address":        types.StringType,
	"gateway":        types.StringType,
	"subnet_mask":    types.StringType,
	"is_lom_capable": types.BoolType,
}

var MemberLomNetworkConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The IPv4 Address of the Grid member.",
	},
	"gateway": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The default gateway for the Grid member.",
	},
	"subnet_mask": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The subnet mask for the Grid member.",
	},
	"is_lom_capable": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines if the physical node supports LOM or not.",
	},
}

func ExpandMemberLomNetworkConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberLomNetworkConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberLomNetworkConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberLomNetworkConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberLomNetworkConfig {
	if m == nil {
		return nil
	}
	to := &grid.MemberLomNetworkConfig{
		Address:    flex.ExpandStringPointer(m.Address),
		Gateway:    flex.ExpandStringPointer(m.Gateway),
		SubnetMask: flex.ExpandStringPointer(m.SubnetMask),
	}
	return to
}

func FlattenMemberLomNetworkConfig(ctx context.Context, from *grid.MemberLomNetworkConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberLomNetworkConfigAttrTypes)
	}
	m := MemberLomNetworkConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberLomNetworkConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberLomNetworkConfigModel) Flatten(ctx context.Context, from *grid.MemberLomNetworkConfig, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberLomNetworkConfigModel{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Gateway = flex.FlattenStringPointer(from.Gateway)
	m.SubnetMask = flex.FlattenStringPointer(from.SubnetMask)
	m.IsLomCapable = types.BoolPointerValue(from.IsLomCapable)
}
