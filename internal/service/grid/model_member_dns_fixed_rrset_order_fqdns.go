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

type MemberDnsFixedRrsetOrderFqdnsModel struct {
	Fqdn       types.String `tfsdk:"fqdn"`
	RecordType types.String `tfsdk:"record_type"`
}

var MemberDnsFixedRrsetOrderFqdnsAttrTypes = map[string]attr.Type{
	"fqdn":        types.StringType,
	"record_type": types.StringType,
}

var MemberDnsFixedRrsetOrderFqdnsResourceSchemaAttributes = map[string]schema.Attribute{
	"fqdn": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The FQDN of the fixed RRset configuration item.",
	},
	"record_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The record type for the specified FQDN in the fixed RRset configuration.",
	},
}

func ExpandMemberDnsFixedRrsetOrderFqdns(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberDnsFixedRrsetOrderFqdns {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberDnsFixedRrsetOrderFqdnsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberDnsFixedRrsetOrderFqdnsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberDnsFixedRrsetOrderFqdns {
	if m == nil {
		return nil
	}
	to := &grid.MemberDnsFixedRrsetOrderFqdns{
		Fqdn:       flex.ExpandStringPointer(m.Fqdn),
		RecordType: flex.ExpandStringPointer(m.RecordType),
	}
	return to
}

func FlattenMemberDnsFixedRrsetOrderFqdns(ctx context.Context, from *grid.MemberDnsFixedRrsetOrderFqdns, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberDnsFixedRrsetOrderFqdnsAttrTypes)
	}
	m := MemberDnsFixedRrsetOrderFqdnsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberDnsFixedRrsetOrderFqdnsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberDnsFixedRrsetOrderFqdnsModel) Flatten(ctx context.Context, from *grid.MemberDnsFixedRrsetOrderFqdns, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberDnsFixedRrsetOrderFqdnsModel{}
	}
	m.Fqdn = flex.FlattenStringPointer(from.Fqdn)
	m.RecordType = flex.FlattenStringPointer(from.RecordType)
}
