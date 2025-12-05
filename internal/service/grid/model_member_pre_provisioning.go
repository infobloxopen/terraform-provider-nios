package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type MemberPreProvisioningModel struct {
	HardwareInfo types.List `tfsdk:"hardware_info"`
	Licenses     types.List `tfsdk:"licenses"`
}

var MemberPreProvisioningAttrTypes = map[string]attr.Type{
	"hardware_info": types.ListType{ElemType: types.ObjectType{AttrTypes: MemberpreprovisioningHardwareInfoAttrTypes}},
	"licenses":      types.ListType{ElemType: types.StringType},
}

var MemberPreProvisioningResourceSchemaAttributes = map[string]schema.Attribute{
	"hardware_info": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberpreprovisioningHardwareInfoResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "An array of structures that describe the hardware being pre-provisioned.",
	},
	"licenses": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "An array of license types the pre-provisioned member should have in order to join the Grid, or the licenses that must be allocated to the member when it joins the Grid using the token-based authentication.",
	},
}

func ExpandMemberPreProvisioning(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberPreProvisioning {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberPreProvisioningModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberPreProvisioningModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberPreProvisioning {
	if m == nil {
		return nil
	}
	to := &grid.MemberPreProvisioning{
		HardwareInfo: flex.ExpandFrameworkListNestedBlock(ctx, m.HardwareInfo, diags, ExpandMemberpreprovisioningHardwareInfo),
		Licenses:     flex.ExpandFrameworkListString(ctx, m.Licenses, diags),
	}
	return to
}

func FlattenMemberPreProvisioning(ctx context.Context, from *grid.MemberPreProvisioning, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberPreProvisioningAttrTypes)
	}
	m := MemberPreProvisioningModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberPreProvisioningAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberPreProvisioningModel) Flatten(ctx context.Context, from *grid.MemberPreProvisioning, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberPreProvisioningModel{}
	}
	m.HardwareInfo = flex.FlattenFrameworkListNestedBlock(ctx, from.HardwareInfo, MemberpreprovisioningHardwareInfoAttrTypes, diags, FlattenMemberpreprovisioningHardwareInfo)
	m.Licenses = flex.FlattenFrameworkListString(ctx, from.Licenses, diags)
}
