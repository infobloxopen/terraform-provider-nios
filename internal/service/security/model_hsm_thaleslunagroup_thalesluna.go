package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type HsmThaleslunagroupThaleslunaModel struct {
	Name                  types.String `tfsdk:"name"`
	PartitionSerialNumber types.String `tfsdk:"partition_serial_number"`
	Disable               types.Bool   `tfsdk:"disable"`
	PartitionId           types.String `tfsdk:"partition_id"`
	IsFipsCompliant       types.Bool   `tfsdk:"is_fips_compliant"`
	ServerCert            types.String `tfsdk:"server_cert"`
	PartitionCapacity     types.Int64  `tfsdk:"partition_capacity"`
	Status                types.String `tfsdk:"status"`
}

var HsmThaleslunagroupThaleslunaAttrTypes = map[string]attr.Type{
	"name":                    types.StringType,
	"partition_serial_number": types.StringType,
	"disable":                 types.BoolType,
	"partition_id":            types.StringType,
	"is_fips_compliant":       types.BoolType,
	"server_cert":             types.StringType,
	"partition_capacity":      types.Int64Type,
	"status":                  types.StringType,
}

var HsmThaleslunagroupThaleslunaResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The HSM Thales Luna device IPv4 Address or FQDN.",
	},
	"partition_serial_number": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The HSM Thales Luna device partition serial number (PSN).",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the HSM Thales Luna device is disabled.",
	},
	"partition_id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Partition ID that is displayed after the appliance has successfully connected to the HSM Thales Luna device.",
	},
	"is_fips_compliant": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines whether the HSM Thales Luna device is FIPS compliant.",
	},
	"server_cert": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The token returned by the uploadinit function call in object fileop for a Thales Luna HSM device certificate.",
	},
	"partition_capacity": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The HSM Thales Luna device partition capacity percentage used.",
	},
	"status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The HSM Thales Luna device status.",
	},
}

func ExpandHsmThaleslunagroupThalesluna(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.HsmThaleslunagroupThalesluna {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m HsmThaleslunagroupThaleslunaModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *HsmThaleslunagroupThaleslunaModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.HsmThaleslunagroupThalesluna {
	if m == nil {
		return nil
	}
	to := &security.HsmThaleslunagroupThalesluna{
		Name:                  flex.ExpandStringPointer(m.Name),
		PartitionSerialNumber: flex.ExpandStringPointer(m.PartitionSerialNumber),
		Disable:               flex.ExpandBoolPointer(m.Disable),
		ServerCert:            flex.ExpandStringPointer(m.ServerCert),
	}
	return to
}

func FlattenHsmThaleslunagroupThalesluna(ctx context.Context, from *security.HsmThaleslunagroupThalesluna, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(HsmThaleslunagroupThaleslunaAttrTypes)
	}
	m := HsmThaleslunagroupThaleslunaModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, HsmThaleslunagroupThaleslunaAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *HsmThaleslunagroupThaleslunaModel) Flatten(ctx context.Context, from *security.HsmThaleslunagroupThalesluna, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = HsmThaleslunagroupThaleslunaModel{}
	}
	m.Name = flex.FlattenStringPointer(from.Name)
	m.PartitionSerialNumber = flex.FlattenStringPointer(from.PartitionSerialNumber)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.PartitionId = flex.FlattenStringPointer(from.PartitionId)
	m.IsFipsCompliant = types.BoolPointerValue(from.IsFipsCompliant)
	m.ServerCert = flex.FlattenStringPointer(from.ServerCert)
	m.PartitionCapacity = flex.FlattenInt64Pointer(from.PartitionCapacity)
	m.Status = flex.FlattenStringPointer(from.Status)
}
