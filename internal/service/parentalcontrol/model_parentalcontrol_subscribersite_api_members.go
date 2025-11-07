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

type ParentalcontrolSubscribersiteApiMembersModel struct {
	Name types.String `tfsdk:"name"`
}

var ParentalcontrolSubscribersiteApiMembersAttrTypes = map[string]attr.Type{
	"name": types.StringType,
}

var ParentalcontrolSubscribersiteApiMembersResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The Grid member name.",
	},
}

func ExpandParentalcontrolSubscribersiteApiMembers(ctx context.Context, o types.Object, diags *diag.Diagnostics) *parentalcontrol.ParentalcontrolSubscribersiteApiMembers {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ParentalcontrolSubscribersiteApiMembersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ParentalcontrolSubscribersiteApiMembersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *parentalcontrol.ParentalcontrolSubscribersiteApiMembers {
	if m == nil {
		return nil
	}
	to := &parentalcontrol.ParentalcontrolSubscribersiteApiMembers{
		Name: flex.ExpandStringPointer(m.Name),
	}
	return to
}

func FlattenParentalcontrolSubscribersiteApiMembers(ctx context.Context, from *parentalcontrol.ParentalcontrolSubscribersiteApiMembers, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ParentalcontrolSubscribersiteApiMembersAttrTypes)
	}
	m := ParentalcontrolSubscribersiteApiMembersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ParentalcontrolSubscribersiteApiMembersAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ParentalcontrolSubscribersiteApiMembersModel) Flatten(ctx context.Context, from *parentalcontrol.ParentalcontrolSubscribersiteApiMembers, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ParentalcontrolSubscribersiteApiMembersModel{}
	}
	m.Name = flex.FlattenStringPointer(from.Name)
}
