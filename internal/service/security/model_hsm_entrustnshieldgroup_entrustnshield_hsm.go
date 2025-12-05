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

type HsmEntrustnshieldgroupEntrustnshieldHsmModel struct {
	RemoteIp   types.String `tfsdk:"remote_ip"`
	RemotePort types.Int64  `tfsdk:"remote_port"`
	Status     types.String `tfsdk:"status"`
	RemoteEsn  types.String `tfsdk:"remote_esn"`
	Keyhash    types.String `tfsdk:"keyhash"`
	Disable    types.Bool   `tfsdk:"disable"`
}

var HsmEntrustnshieldgroupEntrustnshieldHsmAttrTypes = map[string]attr.Type{
	"remote_ip":   types.StringType,
	"remote_port": types.Int64Type,
	"status":      types.StringType,
	"remote_esn":  types.StringType,
	"keyhash":     types.StringType,
	"disable":     types.BoolType,
}

var HsmEntrustnshieldgroupEntrustnshieldHsmResourceSchemaAttributes = map[string]schema.Attribute{
	"remote_ip": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The IPv4 Address of the Entrust nShield HSM device.",
	},
	"remote_port": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The Entrust nShield HSM device destination port.",
	},
	"status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The Entrust nShield HSM device status.",
	},
	"remote_esn": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The Entrust nShield HSM device electronic serial number.",
	},
	"keyhash": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The Entrust nShield HSM device public key digest.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the Entrust nShield HSM device is disabled.",
	},
}

func ExpandHsmEntrustnshieldgroupEntrustnshieldHsm(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.HsmEntrustnshieldgroupEntrustnshieldHsm {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m HsmEntrustnshieldgroupEntrustnshieldHsmModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *HsmEntrustnshieldgroupEntrustnshieldHsmModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.HsmEntrustnshieldgroupEntrustnshieldHsm {
	if m == nil {
		return nil
	}
	to := &security.HsmEntrustnshieldgroupEntrustnshieldHsm{
		RemoteIp:   flex.ExpandStringPointer(m.RemoteIp),
		RemotePort: flex.ExpandInt64Pointer(m.RemotePort),
		Keyhash:    flex.ExpandStringPointer(m.Keyhash),
		Disable:    flex.ExpandBoolPointer(m.Disable),
	}
	return to
}

func FlattenHsmEntrustnshieldgroupEntrustnshieldHsm(ctx context.Context, from *security.HsmEntrustnshieldgroupEntrustnshieldHsm, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(HsmEntrustnshieldgroupEntrustnshieldHsmAttrTypes)
	}
	m := HsmEntrustnshieldgroupEntrustnshieldHsmModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, HsmEntrustnshieldgroupEntrustnshieldHsmAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *HsmEntrustnshieldgroupEntrustnshieldHsmModel) Flatten(ctx context.Context, from *security.HsmEntrustnshieldgroupEntrustnshieldHsm, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = HsmEntrustnshieldgroupEntrustnshieldHsmModel{}
	}
	m.RemoteIp = flex.FlattenStringPointer(from.RemoteIp)
	m.RemotePort = flex.FlattenInt64Pointer(from.RemotePort)
	m.Status = flex.FlattenStringPointer(from.Status)
	m.RemoteEsn = flex.FlattenStringPointer(from.RemoteEsn)
	m.Keyhash = flex.FlattenStringPointer(from.Keyhash)
	m.Disable = types.BoolPointerValue(from.Disable)
}
