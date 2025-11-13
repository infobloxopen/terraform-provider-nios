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

type ParentalcontrolSubscribersiteAbssModel struct {
	IpAddress      types.String `tfsdk:"ip_address"`
	BlockingPolicy types.String `tfsdk:"blocking_policy"`
}

var ParentalcontrolSubscribersiteAbssAttrTypes = map[string]attr.Type{
	"ip_address":      types.StringType,
	"blocking_policy": types.StringType,
}

var ParentalcontrolSubscribersiteAbssResourceSchemaAttributes = map[string]schema.Attribute{
	"ip_address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The IP address of addional blocking server.",
	},
	"blocking_policy": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The blocking policy for the addional blocking server.",
	},
}

func ExpandParentalcontrolSubscribersiteAbss(ctx context.Context, o types.Object, diags *diag.Diagnostics) *parentalcontrol.ParentalcontrolSubscribersiteAbss {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ParentalcontrolSubscribersiteAbssModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ParentalcontrolSubscribersiteAbssModel) Expand(ctx context.Context, diags *diag.Diagnostics) *parentalcontrol.ParentalcontrolSubscribersiteAbss {
	if m == nil {
		return nil
	}
	to := &parentalcontrol.ParentalcontrolSubscribersiteAbss{
		IpAddress:      flex.ExpandStringPointer(m.IpAddress),
		BlockingPolicy: flex.ExpandStringPointer(m.BlockingPolicy),
	}
	return to
}

func FlattenParentalcontrolSubscribersiteAbss(ctx context.Context, from *parentalcontrol.ParentalcontrolSubscribersiteAbss, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ParentalcontrolSubscribersiteAbssAttrTypes)
	}
	m := ParentalcontrolSubscribersiteAbssModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ParentalcontrolSubscribersiteAbssAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ParentalcontrolSubscribersiteAbssModel) Flatten(ctx context.Context, from *parentalcontrol.ParentalcontrolSubscribersiteAbss, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ParentalcontrolSubscribersiteAbssModel{}
	}
	m.IpAddress = flex.FlattenStringPointer(from.IpAddress)
	m.BlockingPolicy = flex.FlattenStringPointer(from.BlockingPolicy)
}
