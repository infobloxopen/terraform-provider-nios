package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	internaltypes "github.com/infobloxopen/terraform-provider-nios/internal/types"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type GridServicerestartGroupRecurringScheduleModel struct {
	Services internaltypes.UnorderedListValue `tfsdk:"services"`
	Mode     types.String                     `tfsdk:"mode"`
	Schedule types.Object                     `tfsdk:"schedule"`
	Force    types.Bool                       `tfsdk:"force"`
}

var GridServicerestartGroupRecurringScheduleAttrTypes = map[string]attr.Type{
	"services": internaltypes.UnorderedListOfStringType,
	"mode":     types.StringType,
	"schedule": types.ObjectType{AttrTypes: GridservicerestartgrouprecurringscheduleScheduleAttrTypes},
	"force":    types.BoolType,
}

var GridServicerestartGroupRecurringScheduleResourceSchemaAttributes = map[string]schema.Attribute{
	"services": schema.ListAttribute{
		CustomType:  internaltypes.UnorderedListOfStringType,
		ElementType: types.StringType,
		Validators: []validator.List{
			customvalidator.StringsInSlice([]string{"ALL", "DHCP", "DHCPV4", "DHCPV6", "DNS"}),
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The list of applicable services for the restart.",
	},
	"mode": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("GROUPED", "SEQUENTIAL", "SIMULTANEOUS"),
		},
		MarkdownDescription: "The restart method for a Grid restart.",
	},
	"schedule": schema.SingleNestedAttribute{
		Attributes: GridservicerestartgrouprecurringscheduleScheduleResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
	"force": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the Restart Group should have a force restart.",
	},
}

func ExpandGridServicerestartGroupRecurringSchedule(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridServicerestartGroupRecurringSchedule {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridServicerestartGroupRecurringScheduleModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridServicerestartGroupRecurringScheduleModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridServicerestartGroupRecurringSchedule {
	if m == nil {
		return nil
	}
	to := &grid.GridServicerestartGroupRecurringSchedule{
		Services: flex.ExpandFrameworkListString(ctx, m.Services, diags),
		Mode:     flex.ExpandStringPointer(m.Mode),
		Schedule: ExpandGridservicerestartgrouprecurringscheduleSchedule(ctx, m.Schedule, diags),
		Force:    flex.ExpandBoolPointer(m.Force),
	}
	return to
}

func FlattenGridServicerestartGroupRecurringSchedule(ctx context.Context, from *grid.GridServicerestartGroupRecurringSchedule, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridServicerestartGroupRecurringScheduleAttrTypes)
	}
	m := GridServicerestartGroupRecurringScheduleModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridServicerestartGroupRecurringScheduleAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridServicerestartGroupRecurringScheduleModel) Flatten(ctx context.Context, from *grid.GridServicerestartGroupRecurringSchedule, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridServicerestartGroupRecurringScheduleModel{}
	}
	m.Services = flex.FlattenFrameworkUnorderedList(ctx, types.StringType, from.Services, diags)
	m.Mode = flex.FlattenStringPointer(from.Mode)
	m.Schedule = FlattenGridservicerestartgrouprecurringscheduleSchedule(ctx, from.Schedule, diags)
	m.Force = types.BoolPointerValue(from.Force)
}
