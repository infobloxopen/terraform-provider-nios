package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dtc"

	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/flex"
)

type DtcPoolServersModel struct {
	Server types.String `tfsdk:"server"`
	Ratio  types.Int64  `tfsdk:"ratio"`
}

var DtcPoolServersAttrTypes = map[string]attr.Type{
	"server": types.StringType,
	"ratio":  types.Int64Type,
}

var DtcPoolServersResourceSchemaAttributes = map[string]schema.Attribute{
	"server": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The server to link with.",
	},
	"ratio": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The weight of server.",
	},
}

func ExpandDtcPoolServers(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dtc.DtcPoolServers {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DtcPoolServersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DtcPoolServersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dtc.DtcPoolServers {
	if m == nil {
		return nil
	}
	to := &dtc.DtcPoolServers{
		Server: flex.ExpandStringPointer(m.Server),
		Ratio:  flex.ExpandInt64Pointer(m.Ratio),
	}
	return to
}

func FlattenDtcPoolServers(ctx context.Context, from *dtc.DtcPoolServers, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DtcPoolServersAttrTypes)
	}
	m := DtcPoolServersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DtcPoolServersAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DtcPoolServersModel) Flatten(ctx context.Context, from *dtc.DtcPoolServers, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DtcPoolServersModel{}
	}
	m.Server = flex.FlattenStringPointer(from.Server)
	m.Ratio = flex.FlattenInt64Pointer(from.Ratio)
}
