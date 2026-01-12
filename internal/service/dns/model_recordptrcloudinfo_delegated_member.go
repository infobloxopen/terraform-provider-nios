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

type RecordptrcloudinfoDelegatedMemberModel struct {
	Ipv4addr types.String `tfsdk:"ipv4addr"`
	Ipv6addr types.String `tfsdk:"ipv6addr"`
	Name     types.String `tfsdk:"name"`
}

var RecordptrcloudinfoDelegatedMemberAttrTypes = map[string]attr.Type{
	"ipv4addr": types.StringType,
	"ipv6addr": types.StringType,
	"name":     types.StringType,
}

var RecordptrcloudinfoDelegatedMemberResourceSchemaAttributes = map[string]schema.Attribute{
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

func (m *RecordptrcloudinfoDelegatedMemberModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.RecordptrcloudinfoDelegatedMember {
	if m == nil {
		return nil
	}
	to := &dns.RecordptrcloudinfoDelegatedMember{}
	return to
}

func FlattenRecordptrcloudinfoDelegatedMember(ctx context.Context, from *dns.RecordptrcloudinfoDelegatedMember, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordptrcloudinfoDelegatedMemberAttrTypes)
	}
	m := RecordptrcloudinfoDelegatedMemberModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RecordptrcloudinfoDelegatedMemberAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RecordptrcloudinfoDelegatedMemberModel) Flatten(ctx context.Context, from *dns.RecordptrcloudinfoDelegatedMember, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RecordptrcloudinfoDelegatedMemberModel{}
	}
	m.Ipv4addr = flex.FlattenStringPointer(from.Ipv4addr)
	m.Ipv6addr = flex.FlattenStringPointer(from.Ipv6addr)
	m.Name = flex.FlattenStringPointer(from.Name)
}
