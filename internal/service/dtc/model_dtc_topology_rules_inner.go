package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dtc"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type DtcTopologyRulesInnerModel struct {
	DestType        types.String `tfsdk:"dest_type"`
	DestinationLink types.String `tfsdk:"destination_link"`
	ReturnType      types.String `tfsdk:"return_type"`
	Topology        types.String `tfsdk:"topology"`
	Valid           types.Bool   `tfsdk:"valid"`
	Sources         types.List   `tfsdk:"sources"`
}

var DtcTopologyRulesInnerAttrTypes = map[string]attr.Type{
	"dest_type":        types.StringType,
	"destination_link": types.StringType,
	"return_type":      types.StringType,
	"topology":         types.StringType,
	"valid":            types.BoolType,
	"sources":          types.ListType{ElemType: types.ObjectType{AttrTypes: DtcTopologyRulesInnerOneOf1SourcesInnerAttrTypes}},
}

var DtcTopologyRulesInnerResourceSchemaAttributes = map[string]schema.Attribute{
	"dest_type": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.OneOf("POOL", "SERVER"),
		},
		MarkdownDescription: "The type of the destination for this rule.",
	},
	"destination_link": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The reference to the destination object.",
	},
	"return_type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("NOERR", "NXDOMAIN", "REGULAR"),
		},
		MarkdownDescription: "The type of the return value for this source.",
	},
	"topology": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The topology for this rule.",
	},
	"valid": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Indicates whether the rule is valid.",
	},
	"sources": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: DtcTopologyRulesInnerOneOf1SourcesInnerResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Conditions for matching sources.",
	},
}

func ExpandDtcTopologyRulesInner(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dtc.DtcTopologyRulesInner {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DtcTopologyRulesInnerModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DtcTopologyRulesInnerModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dtc.DtcTopologyRulesInner {
	if m == nil {
		return nil
	}
	to := &dtc.DtcTopologyRulesInner{
		DtcTopologyRulesInnerOneOf1: &dtc.DtcTopologyRulesInnerOneOf1{
			DestType:        flex.ExpandStringPointer(m.DestType),
			DestinationLink: flex.ExpandStringPointer(m.DestinationLink),
			ReturnType:      flex.ExpandStringPointer(m.ReturnType),
			Sources:         flex.ExpandFrameworkListNestedBlock(ctx, m.Sources, diags, ExpandDtcTopologyRulesInnerOneOf1SourcesInner),
		},
	}
	return to
}

func FlattenDtcTopologyRulesInner(ctx context.Context, from *dtc.DtcTopologyRulesInner, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DtcTopologyRulesInnerAttrTypes)
	}
	m := DtcTopologyRulesInnerModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DtcTopologyRulesInnerAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DtcTopologyRulesInnerModel) Flatten(ctx context.Context, from *dtc.DtcTopologyRulesInner, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DtcTopologyRulesInnerModel{}
	}

	m.DestType = flex.FlattenStringPointer(from.DtcTopologyRulesInnerOneOf1.DestType)
	m.DestinationLink = flex.FlattenStringPointer(from.DtcTopologyRulesInnerOneOf1.DestinationLink)
	m.ReturnType = flex.FlattenStringPointer(from.DtcTopologyRulesInnerOneOf1.ReturnType)
	m.Topology = flex.FlattenStringPointer(from.DtcTopologyRulesInnerOneOf1.Topology)
    m.Valid = types.BoolPointerValue(from.DtcTopologyRulesInnerOneOf1.Valid)
	m.Sources = flex.FlattenFrameworkListNestedBlock(ctx, from.DtcTopologyRulesInnerOneOf1.Sources, DtcTopologyRulesInnerOneOf1SourcesInnerAttrTypes, diags, FlattenDtcTopologyRulesInnerOneOf1SourcesInner)
}
