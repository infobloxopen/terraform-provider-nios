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

type MemberMemberServiceCommunicationModel struct {
	Service types.String `tfsdk:"service"`
	Type    types.String `tfsdk:"type"`
	Option  types.String `tfsdk:"option"`
}

var MemberMemberServiceCommunicationAttrTypes = map[string]attr.Type{
	"service": types.StringType,
	"type":    types.StringType,
	"option":  types.StringType,
}

var MemberMemberServiceCommunicationResourceSchemaAttributes = map[string]schema.Attribute{
	"service": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The service for a Grid member.",
	},
	"type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Communication type.",
	},
	"option": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The option for communication type.",
	},
}

func ExpandMemberMemberServiceCommunication(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberMemberServiceCommunication {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberMemberServiceCommunicationModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberMemberServiceCommunicationModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberMemberServiceCommunication {
	if m == nil {
		return nil
	}
	to := &grid.MemberMemberServiceCommunication{
		Service: flex.ExpandStringPointer(m.Service),
		Type:    flex.ExpandStringPointer(m.Type),
	}
	return to
}

func FlattenMemberMemberServiceCommunication(ctx context.Context, from *grid.MemberMemberServiceCommunication, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberMemberServiceCommunicationAttrTypes)
	}
	m := MemberMemberServiceCommunicationModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberMemberServiceCommunicationAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberMemberServiceCommunicationModel) Flatten(ctx context.Context, from *grid.MemberMemberServiceCommunication, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberMemberServiceCommunicationModel{}
	}
	m.Service = flex.FlattenStringPointer(from.Service)
	m.Type = flex.FlattenStringPointer(from.Type)
	m.Option = flex.FlattenStringPointer(from.Option)
}
