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

type ThreatinsightModulesetModel struct {
	Ref     types.String `tfsdk:"ref"`
	Uuid    types.String `tfsdk:"uuid"`
	Version types.String `tfsdk:"version"`
}

var ThreatinsightModulesetAttrTypes = map[string]attr.Type{
	"ref":     types.StringType,
	"uuid":    types.StringType,
	"version": types.StringType,
}

var ThreatinsightModulesetResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The uuid to the object.",
	},
	"version": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The version number of the threat insight module set.",
	},
}

func ExpandThreatinsightModuleset(ctx context.Context, o types.Object, diags *diag.Diagnostics) *threatinsight.ThreatinsightModuleset {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ThreatinsightModulesetModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ThreatinsightModulesetModel) Expand(ctx context.Context, diags *diag.Diagnostics) *threatinsight.ThreatinsightModuleset {
	if m == nil {
		return nil
	}
	to := &threatinsight.ThreatinsightModuleset{
		Ref:  flex.ExpandStringPointer(m.Ref),
		Uuid: flex.ExpandStringPointer(m.Uuid),
	}
	return to
}

func FlattenThreatinsightModuleset(ctx context.Context, from *threatinsight.ThreatinsightModuleset, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ThreatinsightModulesetAttrTypes)
	}
	m := ThreatinsightModulesetModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ThreatinsightModulesetAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ThreatinsightModulesetModel) Flatten(ctx context.Context, from *threatinsight.ThreatinsightModuleset, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ThreatinsightModulesetModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Version = flex.FlattenStringPointer(from.Version)
}
