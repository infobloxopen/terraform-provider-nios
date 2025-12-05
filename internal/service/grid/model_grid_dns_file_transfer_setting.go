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

type GridDnsFileTransferSettingModel struct {
	Directory           types.String `tfsdk:"directory"`
	ServerAddressOrFqdn types.String `tfsdk:"server_address_or_fqdn"`
	Password            types.String `tfsdk:"password"`
	Type                types.String `tfsdk:"type"`
	Username            types.String `tfsdk:"username"`
	Port                types.Int64  `tfsdk:"port"`
}

var GridDnsFileTransferSettingAttrTypes = map[string]attr.Type{
	"directory":              types.StringType,
	"server_address_or_fqdn": types.StringType,
	"password":               types.StringType,
	"type":                   types.StringType,
	"username":               types.StringType,
	"port":                   types.Int64Type,
}

var GridDnsFileTransferSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"directory": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The directory to save the captured DNS queries and responses.",
	},
	"server_address_or_fqdn": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The server address or a FQDN name of the destination server for DNS capture transfer.",
	},
	"password": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The password to access the destination server directory.",
	},
	"type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The transfer protocol for the captured DNS queries and responses.",
	},
	"username": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The username to access the destination server directory.",
	},
	"port": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Transfer scp port.",
	},
}

func ExpandGridDnsFileTransferSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridDnsFileTransferSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridDnsFileTransferSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridDnsFileTransferSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridDnsFileTransferSetting {
	if m == nil {
		return nil
	}
	to := &grid.GridDnsFileTransferSetting{
		Directory:           flex.ExpandStringPointer(m.Directory),
		ServerAddressOrFqdn: flex.ExpandStringPointer(m.ServerAddressOrFqdn),
		Password:            flex.ExpandStringPointer(m.Password),
		Type:                flex.ExpandStringPointer(m.Type),
		Username:            flex.ExpandStringPointer(m.Username),
		Port:                flex.ExpandInt64Pointer(m.Port),
	}
	return to
}

func FlattenGridDnsFileTransferSetting(ctx context.Context, from *grid.GridDnsFileTransferSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridDnsFileTransferSettingAttrTypes)
	}
	m := GridDnsFileTransferSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridDnsFileTransferSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridDnsFileTransferSettingModel) Flatten(ctx context.Context, from *grid.GridDnsFileTransferSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridDnsFileTransferSettingModel{}
	}
	m.Directory = flex.FlattenStringPointer(from.Directory)
	m.ServerAddressOrFqdn = flex.FlattenStringPointer(from.ServerAddressOrFqdn)
	m.Password = flex.FlattenStringPointer(from.Password)
	m.Type = flex.FlattenStringPointer(from.Type)
	m.Username = flex.FlattenStringPointer(from.Username)
	m.Port = flex.FlattenInt64Pointer(from.Port)
}
