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

type ParentalcontrolSubscribersiteMembersModel struct {
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}

var ParentalcontrolSubscribersiteMembersAttrTypes = map[string]attr.Type{
	"name": types.StringType,
	"type": types.StringType,
}

var ParentalcontrolSubscribersiteMembersResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The Grid member name.",
	},
	"type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The type of member.",
	},
}

func ExpandParentalcontrolSubscribersiteMembers(ctx context.Context, o types.Object, diags *diag.Diagnostics) *parentalcontrol.ParentalcontrolSubscribersiteMembers {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ParentalcontrolSubscribersiteMembersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ParentalcontrolSubscribersiteMembersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *parentalcontrol.ParentalcontrolSubscribersiteMembers {
	if m == nil {
		return nil
	}
	to := &parentalcontrol.ParentalcontrolSubscribersiteMembers{
		Name: flex.ExpandStringPointer(m.Name),
	}
	return to
}

func FlattenParentalcontrolSubscribersiteMembers(ctx context.Context, from *parentalcontrol.ParentalcontrolSubscribersiteMembers, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ParentalcontrolSubscribersiteMembersAttrTypes)
	}
	m := ParentalcontrolSubscribersiteMembersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ParentalcontrolSubscribersiteMembersAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ParentalcontrolSubscribersiteMembersModel) Flatten(ctx context.Context, from *parentalcontrol.ParentalcontrolSubscribersiteMembers, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ParentalcontrolSubscribersiteMembersModel{}
	}
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Type = flex.FlattenStringPointer(from.Type)
}
