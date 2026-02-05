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

type ThreatinsightAllowlistModel struct {
	Ref     types.String `tfsdk:"ref"`
	Comment types.String `tfsdk:"comment"`
	Disable types.Bool   `tfsdk:"disable"`
	Fqdn    types.String `tfsdk:"fqdn"`
	Type    types.String `tfsdk:"type"`
}

var ThreatinsightAllowlistAttrTypes = map[string]attr.Type{
	"ref":     types.StringType,
	"comment": types.StringType,
	"disable": types.BoolType,
	"fqdn":    types.StringType,
	"type":    types.StringType,
}

var ThreatinsightAllowlistResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The descriptive comment for the threat insight allowlist.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the threat insight allowlist is disabled.",
	},
	"fqdn": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The FQDN of the threat insight allowlist.",
	},
	"type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The type of the threat insight allowlist.",
	},
}

func ExpandThreatinsightAllowlist(ctx context.Context, o types.Object, diags *diag.Diagnostics) *threatinsight.ThreatinsightAllowlist {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ThreatinsightAllowlistModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ThreatinsightAllowlistModel) Expand(ctx context.Context, diags *diag.Diagnostics) *threatinsight.ThreatinsightAllowlist {
	if m == nil {
		return nil
	}
	to := &threatinsight.ThreatinsightAllowlist{
		Ref:     flex.ExpandStringPointer(m.Ref),
		Comment: flex.ExpandStringPointer(m.Comment),
		Disable: flex.ExpandBoolPointer(m.Disable),
		Fqdn:    flex.ExpandStringPointer(m.Fqdn),
	}
	return to
}

func FlattenThreatinsightAllowlist(ctx context.Context, from *threatinsight.ThreatinsightAllowlist, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ThreatinsightAllowlistAttrTypes)
	}
	m := ThreatinsightAllowlistModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ThreatinsightAllowlistAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ThreatinsightAllowlistModel) Flatten(ctx context.Context, from *threatinsight.ThreatinsightAllowlist, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ThreatinsightAllowlistModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.Fqdn = flex.FlattenStringPointer(from.Fqdn)
	m.Type = flex.FlattenStringPointer(from.Type)
}
