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

type GridDnsDnstapSettingModel struct {
	DnstapReceiverAddressOrFqdn types.String `tfsdk:"dnstap_receiver_address_or_fqdn"`
	DnstapReceiverPort          types.Int64  `tfsdk:"dnstap_receiver_port"`
	DnstapIdentity              types.String `tfsdk:"dnstap_identity"`
	DnstapVersion               types.String `tfsdk:"dnstap_version"`
}

var GridDnsDnstapSettingAttrTypes = map[string]attr.Type{
	"dnstap_receiver_address_or_fqdn": types.StringType,
	"dnstap_receiver_port":            types.Int64Type,
	"dnstap_identity":                 types.StringType,
	"dnstap_version":                  types.StringType,
}

var GridDnsDnstapSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"dnstap_receiver_address_or_fqdn": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Address or FQDN of DNSTAP receiver.",
	},
	"dnstap_receiver_port": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "DNSTAP receiver port number.",
	},
	"dnstap_identity": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "DNSTAP id string.",
	},
	"dnstap_version": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "DNSTAP version.",
	},
}

func ExpandGridDnsDnstapSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridDnsDnstapSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridDnsDnstapSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridDnsDnstapSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridDnsDnstapSetting {
	if m == nil {
		return nil
	}
	to := &grid.GridDnsDnstapSetting{
		DnstapReceiverAddressOrFqdn: flex.ExpandStringPointer(m.DnstapReceiverAddressOrFqdn),
		DnstapReceiverPort:          flex.ExpandInt64Pointer(m.DnstapReceiverPort),
	}
	return to
}

func FlattenGridDnsDnstapSetting(ctx context.Context, from *grid.GridDnsDnstapSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridDnsDnstapSettingAttrTypes)
	}
	m := GridDnsDnstapSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridDnsDnstapSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridDnsDnstapSettingModel) Flatten(ctx context.Context, from *grid.GridDnsDnstapSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridDnsDnstapSettingModel{}
	}
	m.DnstapReceiverAddressOrFqdn = flex.FlattenStringPointer(from.DnstapReceiverAddressOrFqdn)
	m.DnstapReceiverPort = flex.FlattenInt64Pointer(from.DnstapReceiverPort)
	m.DnstapIdentity = flex.FlattenStringPointer(from.DnstapIdentity)
	m.DnstapVersion = flex.FlattenStringPointer(from.DnstapVersion)
}
