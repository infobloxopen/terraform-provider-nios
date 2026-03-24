package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type MemberNatSettingModel struct {
	Enabled           types.Bool          `tfsdk:"enabled"`
	ExternalVirtualIp iptypes.IPv4Address `tfsdk:"external_virtual_ip"`
	Group             types.String        `tfsdk:"group"`
}

var MemberNatSettingAttrTypes = map[string]attr.Type{
	"enabled":             types.BoolType,
	"external_virtual_ip": iptypes.IPv4AddressType{},
	"group":               types.StringType,
}

var MemberNatSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Determines if NAT should be enabled.",
	},
	"external_virtual_ip": schema.StringAttribute{
		CustomType:          iptypes.IPv4AddressType{},
		Computed:            true,
		Optional:            true,
		MarkdownDescription: "External IP address for NAT.",
	},
	"group": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The NAT group.",
	},
}

func ExpandMemberNatSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberNatSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberNatSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberNatSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberNatSetting {
	if m == nil {
		return nil
	}
	to := &grid.MemberNatSetting{
		Enabled:           flex.ExpandBoolPointer(m.Enabled),
		ExternalVirtualIp: flex.ExpandIPv4Address(m.ExternalVirtualIp),
		Group:             flex.ExpandStringPointerEmptyAsNil(m.Group),
	}
	return to
}

func FlattenMemberNatSetting(ctx context.Context, from *grid.MemberNatSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberNatSettingAttrTypes)
	}
	m := MemberNatSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberNatSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberNatSettingModel) Flatten(ctx context.Context, from *grid.MemberNatSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberNatSettingModel{}
	}
	m.Enabled = types.BoolPointerValue(from.Enabled)
	m.ExternalVirtualIp = flex.FlattenIPv4Address(from.ExternalVirtualIp)
	m.Group = flex.FlattenStringPointer(from.Group)
}
