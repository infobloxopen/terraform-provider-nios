package parentalcontrol

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/parentalcontrol"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type ParentalcontrolSubscribersiteNasGatewaysModel struct {
	Name         types.String `tfsdk:"name"`
	IpAddress    types.String `tfsdk:"ip_address"`
	SharedSecret types.String `tfsdk:"shared_secret"`
	SendAck      types.Bool   `tfsdk:"send_ack"`
	MessageRate  types.Int64  `tfsdk:"message_rate"`
	Comment      types.String `tfsdk:"comment"`
}

var ParentalcontrolSubscribersiteNasGatewaysAttrTypes = map[string]attr.Type{
	"name":          types.StringType,
	"ip_address":    types.StringType,
	"shared_secret": types.StringType,
	"send_ack":      types.BoolType,
	"message_rate":  types.Int64Type,
	"comment":       types.StringType,
}

var ParentalcontrolSubscribersiteNasGatewaysResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of NAS gateway.",
	},
	"ip_address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The IP address of NAS gateway.",
	},
	"shared_secret": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The protocol MD5 phrase.",
	},
	"send_ack": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether an acknowledge needs to be sent.",
	},
	"message_rate": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The message rate per server.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The human readable comment for NAS gateway.",
	},
}

func ExpandParentalcontrolSubscribersiteNasGateways(ctx context.Context, o types.Object, diags *diag.Diagnostics) *parentalcontrol.ParentalcontrolSubscribersiteNasGateways {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ParentalcontrolSubscribersiteNasGatewaysModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ParentalcontrolSubscribersiteNasGatewaysModel) Expand(ctx context.Context, diags *diag.Diagnostics) *parentalcontrol.ParentalcontrolSubscribersiteNasGateways {
	if m == nil {
		return nil
	}
	to := &parentalcontrol.ParentalcontrolSubscribersiteNasGateways{
		Name:         flex.ExpandStringPointer(m.Name),
		IpAddress:    flex.ExpandStringPointer(m.IpAddress),
		SharedSecret: flex.ExpandStringPointer(m.SharedSecret),
		SendAck:      flex.ExpandBoolPointer(m.SendAck),
		Comment:      flex.ExpandStringPointer(m.Comment),
	}
	return to
}

func FlattenParentalcontrolSubscribersiteNasGateways(ctx context.Context, from *parentalcontrol.ParentalcontrolSubscribersiteNasGateways, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ParentalcontrolSubscribersiteNasGatewaysAttrTypes)
	}
	m := ParentalcontrolSubscribersiteNasGatewaysModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ParentalcontrolSubscribersiteNasGatewaysAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ParentalcontrolSubscribersiteNasGatewaysModel) Flatten(ctx context.Context, from *parentalcontrol.ParentalcontrolSubscribersiteNasGateways, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ParentalcontrolSubscribersiteNasGatewaysModel{}
	}
	m.Name = flex.FlattenStringPointer(from.Name)
	m.IpAddress = flex.FlattenStringPointer(from.IpAddress)
	m.SharedSecret = flex.FlattenStringPointer(from.SharedSecret)
	m.SendAck = types.BoolPointerValue(from.SendAck)
	m.MessageRate = flex.FlattenInt64Pointer(from.MessageRate)
	m.Comment = flex.FlattenStringPointer(from.Comment)
}
