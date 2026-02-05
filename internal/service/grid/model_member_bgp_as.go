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

type MemberBgpAsModel struct {
	As         types.Int64 `tfsdk:"as"`
	Keepalive  types.Int64 `tfsdk:"keepalive"`
	Holddown   types.Int64 `tfsdk:"holddown"`
	Neighbors  types.List  `tfsdk:"neighbors"`
	LinkDetect types.Bool  `tfsdk:"link_detect"`
}

var MemberBgpAsAttrTypes = map[string]attr.Type{
	"as":          types.Int64Type,
	"keepalive":   types.Int64Type,
	"holddown":    types.Int64Type,
	"neighbors":   types.ListType{ElemType: types.ObjectType{AttrTypes: MemberbgpasNeighborsAttrTypes}},
	"link_detect": types.BoolType,
}

var MemberBgpAsResourceSchemaAttributes = map[string]schema.Attribute{
	"as": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The number of this autonomous system.",
	},
	"keepalive": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The AS keepalive timer (in seconds). The valid value is from 1 to 21845.",
	},
	"holddown": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The AS holddown timer (in seconds). The valid value is from 3 to 65535.",
	},
	"neighbors": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberbgpasNeighborsResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The BGP neighbors for this AS.",
	},
	"link_detect": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if link detection on the interface is enabled or not.",
	},
}

func ExpandMemberBgpAs(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberBgpAs {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberBgpAsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberBgpAsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberBgpAs {
	if m == nil {
		return nil
	}
	to := &grid.MemberBgpAs{
		As:         flex.ExpandInt64Pointer(m.As),
		Keepalive:  flex.ExpandInt64Pointer(m.Keepalive),
		Holddown:   flex.ExpandInt64Pointer(m.Holddown),
		Neighbors:  flex.ExpandFrameworkListNestedBlock(ctx, m.Neighbors, diags, ExpandMemberbgpasNeighbors),
		LinkDetect: flex.ExpandBoolPointer(m.LinkDetect),
	}
	return to
}

func FlattenMemberBgpAs(ctx context.Context, from *grid.MemberBgpAs, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberBgpAsAttrTypes)
	}
	m := MemberBgpAsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberBgpAsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberBgpAsModel) Flatten(ctx context.Context, from *grid.MemberBgpAs, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberBgpAsModel{}
	}
	m.As = flex.FlattenInt64Pointer(from.As)
	m.Keepalive = flex.FlattenInt64Pointer(from.Keepalive)
	m.Holddown = flex.FlattenInt64Pointer(from.Holddown)
	m.Neighbors = flex.FlattenFrameworkListNestedBlock(ctx, from.Neighbors, MemberbgpasNeighborsAttrTypes, diags, FlattenMemberbgpasNeighbors)
	m.LinkDetect = types.BoolPointerValue(from.LinkDetect)
}
