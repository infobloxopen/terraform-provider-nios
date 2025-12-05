package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type GmcscheduleModel struct {
	Ref                      types.String `tfsdk:"ref"`
	ActivateGmcGroupSchedule types.Bool   `tfsdk:"activate_gmc_group_schedule"`
	GmcGroups                types.List   `tfsdk:"gmc_groups"`
}

var GmcscheduleAttrTypes = map[string]attr.Type{
	"ref":                         types.StringType,
	"activate_gmc_group_schedule": types.BoolType,
	"gmc_groups":                  types.ListType{ElemType: types.StringType},
}

var GmcscheduleResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"activate_gmc_group_schedule": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the gmc schedule is active.",
	},
	"gmc_groups": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Computed:            true,
		MarkdownDescription: "Object array of gmc groups",
	},
}

func ExpandGmcschedule(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.Gmcschedule {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GmcscheduleModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GmcscheduleModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.Gmcschedule {
	if m == nil {
		return nil
	}
	to := &grid.Gmcschedule{
		Ref:                      flex.ExpandStringPointer(m.Ref),
		ActivateGmcGroupSchedule: flex.ExpandBoolPointer(m.ActivateGmcGroupSchedule),
	}
	return to
}

func FlattenGmcschedule(ctx context.Context, from *grid.Gmcschedule, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GmcscheduleAttrTypes)
	}
	m := GmcscheduleModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GmcscheduleAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GmcscheduleModel) Flatten(ctx context.Context, from *grid.Gmcschedule, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GmcscheduleModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.ActivateGmcGroupSchedule = types.BoolPointerValue(from.ActivateGmcGroupSchedule)
	m.GmcGroups = flex.FlattenFrameworkListString(ctx, from.GmcGroups, diags)
}
