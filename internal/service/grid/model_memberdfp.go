package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type MemberdfpModel struct {
	Ref             types.String `tfsdk:"ref"`
	DfpForwardFirst types.Bool   `tfsdk:"dfp_forward_first"`
	HostName        types.String `tfsdk:"host_name"`
	IsDfpOverride   types.Bool   `tfsdk:"is_dfp_override"`
}

var MemberdfpAttrTypes = map[string]attr.Type{
	"ref":               types.StringType,
	"dfp_forward_first": types.BoolType,
	"host_name":         types.StringType,
	"is_dfp_override":   types.BoolType,
}

var MemberdfpResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"dfp_forward_first": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Option to resolve DNS query if resolution over Active Trust Cloud failed.",
	},
	// 	Optional:    true,
	// 	Computed:    true,
	// 	ElementType: types.StringType,
	// 	Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
	// 	Validators: []validator.Map{
	// 		mapvalidator.SizeAtLeast(1),
	// 	},
	// 	MarkdownDescription: "Extensible attributes associated with the object.",
	// },
	// "extattrs_all": schema.MapAttribute{
	// 	Computed: true,
	// 	PlanModifiers: []planmodifier.Map{
	// 		importmod.AssociateInternalId(),
	// 	},
	// 	MarkdownDescription: "Extensible attributes associated with the object, including default and internal attributes.",
	// 	ElementType:         types.StringType,
	// },
	"host_name": schema.StringAttribute{
		Computed: true,
		Optional:            true,
		MarkdownDescription: "Host name of the parent Member",
	},
	"is_dfp_override": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "DFP override lock'.",
	},
}

func (m *MemberdfpModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.Memberdfp {
	if m == nil {
		return nil
	}
	to := &grid.Memberdfp{
		DfpForwardFirst: flex.ExpandBoolPointer(m.DfpForwardFirst),
		IsDfpOverride:   flex.ExpandBoolPointer(m.IsDfpOverride),
	}
	return to
}

func FlattenMemberdfp(ctx context.Context, from *grid.Memberdfp, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberdfpAttrTypes)
	}
	m := MemberdfpModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberdfpAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberdfpModel) Flatten(ctx context.Context, from *grid.Memberdfp, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberdfpModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.DfpForwardFirst = types.BoolPointerValue(from.DfpForwardFirst)
	m.HostName = flex.FlattenStringPointer(from.HostName)
	m.IsDfpOverride = types.BoolPointerValue(from.IsDfpOverride)
}
