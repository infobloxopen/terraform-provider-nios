package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type RecordHostIpv4addrLogicFilterRulesModel struct {
	Filter types.String `tfsdk:"filter"`
	Type   types.String `tfsdk:"type"`
}

var RecordHostIpv4addrLogicFilterRulesAttrTypes = map[string]attr.Type{
	"filter": types.StringType,
	"type":   types.StringType,
}

var RecordHostIpv4addrLogicFilterRulesResourceSchemaAttributes = map[string]schema.Attribute{
	"filter": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The filter name.",
	},
	"type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The filter type. Valid values are: * MAC * NAC * Option",
	},
}

func ExpandRecordHostIpv4addrLogicFilterRules(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.RecordHostIpv4addrLogicFilterRules {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RecordHostIpv4addrLogicFilterRulesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *RecordHostIpv4addrLogicFilterRulesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.RecordHostIpv4addrLogicFilterRules {
	if m == nil {
		return nil
	}
	to := &dns.RecordHostIpv4addrLogicFilterRules{
		Filter: flex.ExpandStringPointer(m.Filter),
		Type:   flex.ExpandStringPointer(m.Type),
	}
	return to
}

func FlattenRecordHostIpv4addrLogicFilterRules(ctx context.Context, from *dns.RecordHostIpv4addrLogicFilterRules, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordHostIpv4addrLogicFilterRulesAttrTypes)
	}
	m := RecordHostIpv4addrLogicFilterRulesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RecordHostIpv4addrLogicFilterRulesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RecordHostIpv4addrLogicFilterRulesModel) Flatten(ctx context.Context, from *dns.RecordHostIpv4addrLogicFilterRules, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RecordHostIpv4addrLogicFilterRulesModel{}
	}
	m.Filter = flex.FlattenStringPointer(from.Filter)
	m.Type = flex.FlattenStringPointer(from.Type)
}
