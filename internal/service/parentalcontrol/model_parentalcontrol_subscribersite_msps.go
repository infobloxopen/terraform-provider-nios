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

type ParentalcontrolSubscribersiteMspsModel struct {
	IpAddress types.String `tfsdk:"ip_address"`
}

var ParentalcontrolSubscribersiteMspsAttrTypes = map[string]attr.Type{
	"ip_address": types.StringType,
}

var ParentalcontrolSubscribersiteMspsResourceSchemaAttributes = map[string]schema.Attribute{
	"ip_address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The IP Address of MSP.",
	},
}

func ExpandParentalcontrolSubscribersiteMsps(ctx context.Context, o types.Object, diags *diag.Diagnostics) *parentalcontrol.ParentalcontrolSubscribersiteMsps {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ParentalcontrolSubscribersiteMspsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ParentalcontrolSubscribersiteMspsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *parentalcontrol.ParentalcontrolSubscribersiteMsps {
	if m == nil {
		return nil
	}
	to := &parentalcontrol.ParentalcontrolSubscribersiteMsps{
		IpAddress: flex.ExpandStringPointer(m.IpAddress),
	}
	return to
}

func FlattenParentalcontrolSubscribersiteMsps(ctx context.Context, from *parentalcontrol.ParentalcontrolSubscribersiteMsps, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ParentalcontrolSubscribersiteMspsAttrTypes)
	}
	m := ParentalcontrolSubscribersiteMspsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ParentalcontrolSubscribersiteMspsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ParentalcontrolSubscribersiteMspsModel) Flatten(ctx context.Context, from *parentalcontrol.ParentalcontrolSubscribersiteMsps, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ParentalcontrolSubscribersiteMspsModel{}
	}
	m.IpAddress = flex.FlattenStringPointer(from.IpAddress)
}
