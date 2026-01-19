package dtc

import (
	"context"
	//"fmt"

	// "net/http"
	// "crypto/tls"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	// "github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	// "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/infoblox-nios-go-client/dtc"
	// "github.com/infobloxopen/infoblox-nios-go-client/option"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	//"github.com/infobloxopen/terraform-provider-nios/internal/utils"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type DtcTopologyModel struct {
	Ref         types.String `tfsdk:"ref"`
	Comment     types.String `tfsdk:"comment"`
	ExtAttrs    types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll types.Map    `tfsdk:"extattrs_all"`
	Name        types.String `tfsdk:"name"`
	Rules       types.List   `tfsdk:"rules"`
}

var DtcTopologyAttrTypes = map[string]attr.Type{
	"ref":          types.StringType,
	"comment":      types.StringType,
	"extattrs":     types.MapType{ElemType: types.StringType},
	"extattrs_all": types.MapType{ElemType: types.StringType},
	"name":         types.StringType,
	"rules":        types.ListType{ElemType: types.ObjectType{AttrTypes: DtcTopologyRulesInnerAttrTypes}},
}

var DtcTopologyResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 256),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The comment for the DTC TOPOLOGY monitor object; maximum 256 characters.",
	},
	"extattrs": schema.MapAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object.",
		ElementType:         types.StringType,
		Default:             mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
	},
	"extattrs_all": schema.MapAttribute{
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object , including default attributes.",
		ElementType:         types.StringType,
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "Display name of the DTC Topology.",
	},
	"rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: DtcTopologyRulesInnerResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Topology rules.",
	},
}

func (m *DtcTopologyModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dtc.DtcTopology {
	if m == nil {
		return nil
	}
	to := &dtc.DtcTopology{
		Ref:      flex.ExpandStringPointer(m.Ref),
		Comment:  flex.ExpandStringPointer(m.Comment),
		ExtAttrs: ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		Name:     flex.ExpandStringPointer(m.Name),
		Rules:    flex.ExpandFrameworkListNestedBlock(ctx, m.Rules, diags, ExpandDtcTopologyRulesInner),
	}
	return to
}

func FlattenDtcTopology(ctx context.Context, from *dtc.DtcTopology, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DtcTopologyAttrTypes)
	}
	m := DtcTopologyModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, DtcTopologyAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DtcTopologyModel) Flatten(ctx context.Context, from *dtc.DtcTopology, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DtcTopologyModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Rules = flex.FlattenFrameworkListNestedBlock(ctx, from.Rules, DtcTopologyRulesInnerAttrTypes, diags, FlattenDtcTopologyRulesInner)
}
// func (r *DtcTopologyResource) getOperationOnRef(ctx context.Context, ruleRef string , diags *diag.Diagnostics) DtcTopologyRulesInnerModel{

// 	apiRes, _, err := r.client.DTCAPI.
// 	DtcTopologyRuleAPI.
// 		Read(ctx, utils.ExtractResourceRef(ruleRef)).
// 		ReturnFieldsPlus(readableAttributesForDtcTopologyRule).
// 		ReturnAsObject(1).
// 		Execute()

// 		if err != nil {
// 			diags.AddError("Client Error", fmt.Sprintf("Unable to read DTC Topology Rules %s", err))
// 		}
// 		res := apiRes.GetDtcTopologyRuleResponseObjectAsResult.GetResult()
		
// 		sources := make([]dtc.DtcTopologyRulesInnerOneOf1SourcesInner, 0, len(res.Sources))
// 		for _, source := range res.Sources {
// 			srcValues := dtc.DtcTopologyRulesInnerOneOf1SourcesInner{
// 				SourceOp:    source.SourceOp,
// 				SourceType:  source.SourceType,
// 				SourceValue: source.SourceValue,
// 			}
// 			sources = append(sources, srcValues)
// 		}
// 		rulesInner := &dtc.DtcTopologyRulesInner{
// 			DtcTopologyRulesInnerOneOf1 : &dtc.DtcTopologyRulesInnerOneOf1{
// 				DestType:        res.DestType,
// 				//DestinationLink: res.DestinationLink,
// 				ReturnType:      res.ReturnType,
// 				Topology:        res.Topology,
// 				Valid:           res.Valid,
// 				Sources:         sources,
// 			},
// 		}
// 		var rule DtcTopologyRulesInnerModel
// 		rule.Flatten(ctx,rulesInner,diags)

// 		return rule
// }
// func FlattenDtcTopologyRules(ctx context.Context , from []dtc.DtcTopologyRulesInner, diags *diag.Diagnostics , r *DtcTopologyResource) types.List {
// 	to := make([]DtcTopologyRulesInnerModel, 0, len(from))
// 	for _, rule := range from {
// 		ruleRef := rule.DtcTopologyRulesInnerOneOf.Ref

// 		ruleModel := r.getOperationOnRef(ctx , *ruleRef , diags)
// 		to = append(to, ruleModel)
// 	}

// 	t, d := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DtcTopologyRulesInnerAttrTypes}, to)
// 	diags.Append(d...)
// 	return t
// }
