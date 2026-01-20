package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	"github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/plancontrol"
)

type RecordcnamecloudinfoDelegatedMemberModel struct {
	Ipv4addr types.String `tfsdk:"ipv4addr"`
	Ipv6addr types.String `tfsdk:"ipv6addr"`
	Name     types.String `tfsdk:"name"`
}

var RecordcnamecloudinfoDelegatedMemberAttrTypes = map[string]attr.Type{
	"ipv4addr": types.StringType,
	"ipv6addr": types.StringType,
	"name":     types.StringType,
}

var RecordcnamecloudinfoDelegatedMemberResourceSchemaAttributes = map[string]schema.Attribute{
	"ipv4addr": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The IPv4 Address of the Grid Member.",
		PlanModifiers: []planmodifier.String{
			plancontrol.UseStateForUnknownString(),
		},
	},
	"ipv6addr": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The IPv6 Address of the Grid Member.",
		PlanModifiers: []planmodifier.String{
			plancontrol.UseStateForUnknownString(),
		},
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The Grid member name",
		PlanModifiers: []planmodifier.String{
			plancontrol.UseStateForUnknownString(),
		},
	},
}

func (m *RecordcnamecloudinfoDelegatedMemberModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.RecordcnamecloudinfoDelegatedMember {
	if m == nil {
		return nil
	}
	to := &dns.RecordcnamecloudinfoDelegatedMember{}
	return to
}

func FlattenRecordcnamecloudinfoDelegatedMember(ctx context.Context, from *dns.RecordcnamecloudinfoDelegatedMember, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordcnamecloudinfoDelegatedMemberAttrTypes)
	}
	m := RecordcnamecloudinfoDelegatedMemberModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RecordcnamecloudinfoDelegatedMemberAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RecordcnamecloudinfoDelegatedMemberModel) Flatten(ctx context.Context, from *dns.RecordcnamecloudinfoDelegatedMember, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RecordcnamecloudinfoDelegatedMemberModel{}
	}
	m.Ipv4addr = flex.FlattenStringPointer(from.Ipv4addr)
	m.Ipv6addr = flex.FlattenStringPointer(from.Ipv6addr)
	m.Name = flex.FlattenStringPointer(from.Name)
}
