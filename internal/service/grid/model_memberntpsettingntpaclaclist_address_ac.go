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

type MemberntpsettingntpaclaclistAddressAcModel struct {
	Address    types.String `tfsdk:"address"`
	Permission types.String `tfsdk:"permission"`
}

var MemberntpsettingntpaclaclistAddressAcAttrTypes = map[string]attr.Type{
	"address":    types.StringType,
	"permission": types.StringType,
}

var MemberntpsettingntpaclaclistAddressAcResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The address this rule applies to or \"Any\".",
	},
	"permission": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The permission to use for this address.",
	},
}

func ExpandMemberntpsettingntpaclaclistAddressAc(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberntpsettingntpaclaclistAddressAc {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberntpsettingntpaclaclistAddressAcModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberntpsettingntpaclaclistAddressAcModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberntpsettingntpaclaclistAddressAc {
	if m == nil {
		return nil
	}
	to := &grid.MemberntpsettingntpaclaclistAddressAc{
		Address:    flex.ExpandStringPointer(m.Address),
		Permission: flex.ExpandStringPointer(m.Permission),
	}
	return to
}

func FlattenMemberntpsettingntpaclaclistAddressAc(ctx context.Context, from *grid.MemberntpsettingntpaclaclistAddressAc, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberntpsettingntpaclaclistAddressAcAttrTypes)
	}
	m := MemberntpsettingntpaclaclistAddressAcModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberntpsettingntpaclaclistAddressAcAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberntpsettingntpaclaclistAddressAcModel) Flatten(ctx context.Context, from *grid.MemberntpsettingntpaclaclistAddressAc, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberntpsettingntpaclaclistAddressAcModel{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Permission = flex.FlattenStringPointer(from.Permission)
}
