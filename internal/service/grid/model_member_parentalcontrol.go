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

type MemberParentalcontrolModel struct {
	Ref           types.String `tfsdk:"ref"`
	EnableService types.Bool   `tfsdk:"enable_service"`
	Name          types.String `tfsdk:"name"`
}

var MemberParentalcontrolAttrTypes = map[string]attr.Type{
	"ref":            types.StringType,
	"enable_service": types.BoolType,
	"name":           types.StringType,
}

var MemberParentalcontrolResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"enable_service": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the parental control service is enabled.",
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The parental control member hostname.",
	},
}

func ExpandMemberParentalcontrol(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberParentalcontrol {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberParentalcontrolModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberParentalcontrolModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberParentalcontrol {
	if m == nil {
		return nil
	}
	to := &grid.MemberParentalcontrol{
		Ref:           flex.ExpandStringPointer(m.Ref),
		EnableService: flex.ExpandBoolPointer(m.EnableService),
	}
	return to
}

func FlattenMemberParentalcontrol(ctx context.Context, from *grid.MemberParentalcontrol, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberParentalcontrolAttrTypes)
	}
	m := MemberParentalcontrolModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberParentalcontrolAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberParentalcontrolModel) Flatten(ctx context.Context, from *grid.MemberParentalcontrol, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberParentalcontrolModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.EnableService = types.BoolPointerValue(from.EnableService)
	m.Name = flex.FlattenStringPointer(from.Name)
}
