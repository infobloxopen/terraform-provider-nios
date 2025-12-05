package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/import"
)

type MemberdfpModel struct {
	Ref             types.String `tfsdk:"ref"`
    Uuid        types.String `tfsdk:"uuid"`
	DfpForwardFirst types.Bool   `tfsdk:"dfp_forward_first"`
	ExtAttrs        types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll     types.Map    `tfsdk:"extattrs_all"`
	HostName        types.String `tfsdk:"host_name"`
	IsDfpOverride   types.Bool   `tfsdk:"is_dfp_override"`
}

var MemberdfpAttrTypes = map[string]attr.Type{
	"ref":               types.StringType,
    "uuid":        types.StringType,
	"dfp_forward_first": types.BoolType,
	"extattrs":          types.MapType{ElemType: types.StringType},
	"extattrs_all":      types.MapType{ElemType: types.StringType},
	"host_name":         types.StringType,
	"is_dfp_override":   types.BoolType,
}

var MemberdfpResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
    "uuid": schema.StringAttribute{
        Computed:            true,
        MarkdownDescription: "The uuid to the object.",
    },
	"dfp_forward_first": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Option to resolve DNS query if resolution over Active Trust Cloud failed.",
	},
	"extattrs": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.",
	},
	"extattrs_all": schema.MapAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "Host name of the parent Member",
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
	},
	"is_dfp_override": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "DFP override lock'.",
	},
}

func ExpandMemberdfp(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.Memberdfp {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberdfpModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberdfpModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.Memberdfp {
	if m == nil {
		return nil
	}
	to := &grid.Memberdfp{
		Ref:             flex.ExpandStringPointer(m.Ref),
		DfpForwardFirst: flex.ExpandBoolPointer(m.DfpForwardFirst),
		ExtAttrs:        ExpandExtAttrs(ctx, m.ExtAttrs, diags),
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
	m.ExtAttrsAll = types.MapNull(types.StringType)
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
    m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.DfpForwardFirst = types.BoolPointerValue(from.DfpForwardFirst)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.HostName = flex.FlattenStringPointer(from.HostName)
	m.IsDfpOverride = types.BoolPointerValue(from.IsDfpOverride)
}
