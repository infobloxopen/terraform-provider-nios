package threatinsight

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/threatinsight"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type ThreatinsightInsightAllowlistModel struct {
	Ref     types.String `tfsdk:"ref"`
	Version types.String `tfsdk:"version"`
}

var ThreatinsightInsightAllowlistAttrTypes = map[string]attr.Type{
	"ref":     types.StringType,
	"version": types.StringType,
}

var ThreatinsightInsightAllowlistResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"version": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Allowlist version string.",
	},
}

func ExpandThreatinsightInsightAllowlist(ctx context.Context, o types.Object, diags *diag.Diagnostics) *threatinsight.ThreatinsightInsightAllowlist {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ThreatinsightInsightAllowlistModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ThreatinsightInsightAllowlistModel) Expand(ctx context.Context, diags *diag.Diagnostics) *threatinsight.ThreatinsightInsightAllowlist {
	if m == nil {
		return nil
	}
	to := &threatinsight.ThreatinsightInsightAllowlist{
		Ref: flex.ExpandStringPointer(m.Ref),
	}
	return to
}

func FlattenThreatinsightInsightAllowlist(ctx context.Context, from *threatinsight.ThreatinsightInsightAllowlist, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ThreatinsightInsightAllowlistAttrTypes)
	}
	m := ThreatinsightInsightAllowlistModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ThreatinsightInsightAllowlistAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ThreatinsightInsightAllowlistModel) Flatten(ctx context.Context, from *threatinsight.ThreatinsightInsightAllowlist, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ThreatinsightInsightAllowlistModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Version = flex.FlattenStringPointer(from.Version)
}
