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

type MemberEmailSettingModel struct {
	Enabled           types.Bool   `tfsdk:"enabled"`
	FromAddress       types.String `tfsdk:"from_address"`
	Address           types.String `tfsdk:"address"`
	RelayEnabled      types.Bool   `tfsdk:"relay_enabled"`
	Relay             types.String `tfsdk:"relay"`
	Password          types.String `tfsdk:"password"`
	Smtps             types.Bool   `tfsdk:"smtps"`
	PortNumber        types.Int64  `tfsdk:"port_number"`
	UseAuthentication types.Bool   `tfsdk:"use_authentication"`
}

var MemberEmailSettingAttrTypes = map[string]attr.Type{
	"enabled":            types.BoolType,
	"from_address":       types.StringType,
	"address":            types.StringType,
	"relay_enabled":      types.BoolType,
	"relay":              types.StringType,
	"password":           types.StringType,
	"smtps":              types.BoolType,
	"port_number":        types.Int64Type,
	"use_authentication": types.BoolType,
}

var MemberEmailSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if email notification is enabled or not.",
	},
	"from_address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The email address of a Grid Member for 'from' field in notification.",
	},
	"address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The notification email address of a Grid member.",
	},
	"relay_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if email relay is enabled or not.",
	},
	"relay": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The relay name or IP address.",
	},
	"password": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Password to validate from address",
	},
	"smtps": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "SMTP over TLS",
	},
	"port_number": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "SMTP port number",
	},
	"use_authentication": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enable or disable SMTP auth",
	},
}

func ExpandMemberEmailSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberEmailSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberEmailSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberEmailSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberEmailSetting {
	if m == nil {
		return nil
	}
	to := &grid.MemberEmailSetting{
		Enabled:           flex.ExpandBoolPointer(m.Enabled),
		FromAddress:       flex.ExpandStringPointer(m.FromAddress),
		Address:           flex.ExpandStringPointer(m.Address),
		RelayEnabled:      flex.ExpandBoolPointer(m.RelayEnabled),
		Relay:             flex.ExpandStringPointer(m.Relay),
		Password:          flex.ExpandStringPointer(m.Password),
		Smtps:             flex.ExpandBoolPointer(m.Smtps),
		PortNumber:        flex.ExpandInt64Pointer(m.PortNumber),
		UseAuthentication: flex.ExpandBoolPointer(m.UseAuthentication),
	}
	return to
}

func FlattenMemberEmailSetting(ctx context.Context, from *grid.MemberEmailSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberEmailSettingAttrTypes)
	}
	m := MemberEmailSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberEmailSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberEmailSettingModel) Flatten(ctx context.Context, from *grid.MemberEmailSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberEmailSettingModel{}
	}
	m.Enabled = types.BoolPointerValue(from.Enabled)
	m.FromAddress = flex.FlattenStringPointer(from.FromAddress)
	m.Address = flex.FlattenStringPointer(from.Address)
	m.RelayEnabled = types.BoolPointerValue(from.RelayEnabled)
	m.Relay = flex.FlattenStringPointer(from.Relay)
	m.Password = flex.FlattenStringPointer(from.Password)
	m.Smtps = types.BoolPointerValue(from.Smtps)
	m.PortNumber = flex.FlattenInt64Pointer(from.PortNumber)
	m.UseAuthentication = types.BoolPointerValue(from.UseAuthentication)
}
