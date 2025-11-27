package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/dtc"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type DtcTopologyRuleModel struct {
	Ref             types.String `tfsdk:"ref"`
	DestType        types.String `tfsdk:"dest_type"`
	DestinationLink types.String `tfsdk:"destination_link"`
	ReturnType      types.String `tfsdk:"return_type"`
	Sources         types.List   `tfsdk:"sources"`
	Topology        types.String `tfsdk:"topology"`
	Valid           types.Bool   `tfsdk:"valid"`
}

var DtcTopologyRuleAttrTypes = map[string]attr.Type{
	"ref":              types.StringType,
	"dest_type":        types.StringType,
	"destination_link": types.StringType,
	"return_type":      types.StringType,
	"sources":          types.ListType{ElemType: types.ObjectType{AttrTypes: DtcTopologyRuleSourcesAttrTypes}},
	"topology":         types.StringType,
	"valid":            types.BoolType,
}

var DtcTopologyRuleResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"dest_type": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The type of the destination for this DTC Topology rule.",
	},
	"destination_link": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The link to the destination for this DTC Topology rule.",
	},
	"return_type": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString("REGULAR"),
		MarkdownDescription: "Type of the DNS response for rule.",
	},
	"sources": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: DtcTopologyRuleSourcesResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The conditions for matching sources. Should be empty to set rule as default destination.",
	},
	"topology": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The DTC Topology the rule belongs to.",
	},
	"valid": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "True if the label in the rule exists in the current Topology DB. Always true for SUBNET rules. Rules with non-existent labels may be configured but will never match.",
	},
}

func (m *DtcTopologyRuleModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dtc.DtcTopologyRule {
	if m == nil {
		return nil
	}
	to := &dtc.DtcTopologyRule{
		DestType:        flex.ExpandStringPointer(m.DestType),
		DestinationLink: ExpandDtcTopologyRuleDestinationLink(ctx, m.DestinationLink, diags),
		ReturnType:      flex.ExpandStringPointer(m.ReturnType),
		Sources:         flex.ExpandFrameworkListNestedBlock(ctx, m.Sources, diags, ExpandDtcTopologyRuleSources),
	}
	return to
}

func FlattenDtcTopologyRule(ctx context.Context, from *dtc.DtcTopologyRule, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DtcTopologyRuleAttrTypes)
	}
	m := DtcTopologyRuleModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DtcTopologyRuleAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DtcTopologyRuleModel) Flatten(ctx context.Context, from *dtc.DtcTopologyRule, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DtcTopologyRuleModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.DestType = flex.FlattenStringPointer(from.DestType)
	m.DestinationLink = FlattenDtcTopologyRuleDestinationLink(ctx, from.DestinationLink, diags)
	m.ReturnType = flex.FlattenStringPointer(from.ReturnType)
	m.Sources = flex.FlattenFrameworkListNestedBlock(ctx, from.Sources, DtcTopologyRuleSourcesAttrTypes, diags, FlattenDtcTopologyRuleSources)
	m.Topology = flex.FlattenStringPointer(from.Topology)
	m.Valid = types.BoolPointerValue(from.Valid)
}
