package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type RecordaaaacloudinfoDelegatedMemberModel struct {
	Ipv4addr types.String `tfsdk:"ipv4addr"`
	Ipv6addr types.String `tfsdk:"ipv6addr"`
	Name     types.String `tfsdk:"name"`
}

var RecordaaaacloudinfoDelegatedMemberAttrTypes = map[string]attr.Type{
	"ipv4addr": types.StringType,
	"ipv6addr": types.StringType,
	"name":     types.StringType,
}

var RecordaaaacloudinfoDelegatedMemberResourceSchemaAttributes = map[string]schema.Attribute{
	"ipv4addr": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The IPv4 Address of the Grid Member.",
	},
	"ipv6addr": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The IPv6 Address of the Grid Member.",
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The Grid member name",
	},
}

func (m *RecordaaaacloudinfoDelegatedMemberModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.RecordaaaacloudinfoDelegatedMember {
	if m == nil {
		return nil
	}
	to := &dns.RecordaaaacloudinfoDelegatedMember{}
	return to
}

func FlattenRecordaaaacloudinfoDelegatedMember(ctx context.Context, from *dns.RecordaaaacloudinfoDelegatedMember, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordaaaacloudinfoDelegatedMemberAttrTypes)
	}
	m := RecordaaaacloudinfoDelegatedMemberModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RecordaaaacloudinfoDelegatedMemberAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RecordaaaacloudinfoDelegatedMemberModel) Flatten(ctx context.Context, from *dns.RecordaaaacloudinfoDelegatedMember, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RecordaaaacloudinfoDelegatedMemberModel{}
	}
	m.Ipv4addr = flex.FlattenStringPointer(from.Ipv4addr)
	m.Ipv6addr = flex.FlattenStringPointer(from.Ipv6addr)
	m.Name = flex.FlattenStringPointer(from.Name)
}
