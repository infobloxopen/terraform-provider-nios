package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type MemberpreprovisioningHardwareInfoModel struct {
	Hwtype types.String `tfsdk:"hwtype"`
}

var MemberpreprovisioningHardwareInfoAttrTypes = map[string]attr.Type{
	"hwtype": types.StringType,
}

var MemberpreprovisioningHardwareInfoResourceSchemaAttributes = map[string]schema.Attribute{
	"hwtype": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.OneOf("CP-V1405", "CP-V2205", "CP-V805", "IB-1415", "IB-1425", "IB-1516", "IB-1526", "IB-2215", "IB-2225", "IB-2326", "IB-4015", "IB-4025", "IB-4126", "IB-815", "IB-825", "IB-926", "IB-FLEX", "IB-V1415", "IB-V1425", "IB-V1516", "IB-V1526", "IB-V2215", "IB-V2225", "IB-V2326", "IB-V4015", "IB-V4025", "IB-V4126", "IB-V815", "IB-V825", "IB-V926"),
		},
		MarkdownDescription: "Hardware type.",
	},
}

func ExpandMemberpreprovisioningHardwareInfo(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberpreprovisioningHardwareInfo {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberpreprovisioningHardwareInfoModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberpreprovisioningHardwareInfoModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberpreprovisioningHardwareInfo {
	if m == nil {
		return nil
	}
	to := &grid.MemberpreprovisioningHardwareInfo{
		Hwtype: flex.ExpandStringPointer(m.Hwtype),
	}
	return to
}

func FlattenMemberpreprovisioningHardwareInfo(ctx context.Context, from *grid.MemberpreprovisioningHardwareInfo, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberpreprovisioningHardwareInfoAttrTypes)
	}
	m := MemberpreprovisioningHardwareInfoModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberpreprovisioningHardwareInfoAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberpreprovisioningHardwareInfoModel) Flatten(ctx context.Context, from *grid.MemberpreprovisioningHardwareInfo, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberpreprovisioningHardwareInfoModel{}
	}
	m.Hwtype = flex.FlattenStringPointer(from.Hwtype)
}
