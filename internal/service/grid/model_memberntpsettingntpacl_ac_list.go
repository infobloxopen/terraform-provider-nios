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

type MemberntpsettingntpaclAcListModel struct {
	AddressAc types.Object `tfsdk:"address_ac"`
	Service   types.String `tfsdk:"service"`
}

var MemberntpsettingntpaclAcListAttrTypes = map[string]attr.Type{
	"address_ac": types.ObjectType{AttrTypes: MemberntpsettingntpaclaclistAddressAcAttrTypes},
	"service":    types.StringType,
}

var MemberntpsettingntpaclAcListResourceSchemaAttributes = map[string]schema.Attribute{
	"address_ac": schema.SingleNestedAttribute{
		Attributes: MemberntpsettingntpaclaclistAddressAcResourceSchemaAttributes,
		Optional:   true,
	},
	"service": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The type of service with access control.",
	},
}

func ExpandMemberntpsettingntpaclAcList(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberntpsettingntpaclAcList {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberntpsettingntpaclAcListModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberntpsettingntpaclAcListModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberntpsettingntpaclAcList {
	if m == nil {
		return nil
	}
	to := &grid.MemberntpsettingntpaclAcList{
		AddressAc: ExpandMemberntpsettingntpaclaclistAddressAc(ctx, m.AddressAc, diags),
		Service:   flex.ExpandStringPointer(m.Service),
	}
	return to
}

func FlattenMemberntpsettingntpaclAcList(ctx context.Context, from *grid.MemberntpsettingntpaclAcList, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberntpsettingntpaclAcListAttrTypes)
	}
	m := MemberntpsettingntpaclAcListModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberntpsettingntpaclAcListAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberntpsettingntpaclAcListModel) Flatten(ctx context.Context, from *grid.MemberntpsettingntpaclAcList, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberntpsettingntpaclAcListModel{}
	}
	m.AddressAc = FlattenMemberntpsettingntpaclaclistAddressAc(ctx, from.AddressAc, diags)
	m.Service = flex.FlattenStringPointer(from.Service)
}
