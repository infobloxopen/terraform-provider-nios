package threatprotection

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/threatprotection"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type ThreatprotectionStatisticsModel struct {
	Ref       types.String `tfsdk:"ref"`
	Member    types.String `tfsdk:"member"`
	StatInfos types.List   `tfsdk:"stat_infos"`
}

var ThreatprotectionStatisticsAttrTypes = map[string]attr.Type{
	"ref":        types.StringType,
	"member":     types.StringType,
	"stat_infos": types.ListType{ElemType: types.ObjectType{AttrTypes: ThreatprotectionStatisticsStatInfosAttrTypes}},
}

var ThreatprotectionStatisticsResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"member": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The Grid member name to get threat protection statistics. If nothing is specified then event statistics is returned for the Grid.",
	},
	"stat_infos": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ThreatprotectionStatisticsStatInfosResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Computed:            true,
		MarkdownDescription: "The list of event statistical information for the Grid or particular members.",
	},
}

func ExpandThreatprotectionStatistics(ctx context.Context, o types.Object, diags *diag.Diagnostics) *threatprotection.ThreatprotectionStatistics {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ThreatprotectionStatisticsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ThreatprotectionStatisticsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *threatprotection.ThreatprotectionStatistics {
	if m == nil {
		return nil
	}
	to := &threatprotection.ThreatprotectionStatistics{
		Ref: flex.ExpandStringPointer(m.Ref),
	}
	return to
}

func FlattenThreatprotectionStatistics(ctx context.Context, from *threatprotection.ThreatprotectionStatistics, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ThreatprotectionStatisticsAttrTypes)
	}
	m := ThreatprotectionStatisticsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ThreatprotectionStatisticsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ThreatprotectionStatisticsModel) Flatten(ctx context.Context, from *threatprotection.ThreatprotectionStatistics, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ThreatprotectionStatisticsModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Member = flex.FlattenStringPointer(from.Member)
	m.StatInfos = flex.FlattenFrameworkListNestedBlock(ctx, from.StatInfos, ThreatprotectionStatisticsStatInfosAttrTypes, diags, FlattenThreatprotectionStatisticsStatInfos)
}
