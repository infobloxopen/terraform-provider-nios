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

type ParentalcontrolSubscribersiteSpmsModel struct {
	IpAddress types.String `tfsdk:"ip_address"`
}

var ParentalcontrolSubscribersiteSpmsAttrTypes = map[string]attr.Type{
	"ip_address": types.StringType,
}

var ParentalcontrolSubscribersiteSpmsResourceSchemaAttributes = map[string]schema.Attribute{
	"ip_address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The IP Address of SPM.",
	},
}

func ExpandParentalcontrolSubscribersiteSpms(ctx context.Context, o types.Object, diags *diag.Diagnostics) *parentalcontrol.ParentalcontrolSubscribersiteSpms {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ParentalcontrolSubscribersiteSpmsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ParentalcontrolSubscribersiteSpmsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *parentalcontrol.ParentalcontrolSubscribersiteSpms {
	if m == nil {
		return nil
	}
	to := &parentalcontrol.ParentalcontrolSubscribersiteSpms{
		IpAddress: flex.ExpandStringPointer(m.IpAddress),
	}
	return to
}

func FlattenParentalcontrolSubscribersiteSpms(ctx context.Context, from *parentalcontrol.ParentalcontrolSubscribersiteSpms, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ParentalcontrolSubscribersiteSpmsAttrTypes)
	}
	m := ParentalcontrolSubscribersiteSpmsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ParentalcontrolSubscribersiteSpmsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ParentalcontrolSubscribersiteSpmsModel) Flatten(ctx context.Context, from *parentalcontrol.ParentalcontrolSubscribersiteSpms, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ParentalcontrolSubscribersiteSpmsModel{}
	}
	m.IpAddress = flex.FlattenStringPointer(from.IpAddress)
}
