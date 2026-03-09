package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	planmodifiers "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/immutable"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type HsmThaleslunagroupModel struct {
	Ref        types.String `tfsdk:"ref"`
	Comment    types.String `tfsdk:"comment"`
	GroupSn    types.String `tfsdk:"group_sn"`
	HsmVersion types.String `tfsdk:"hsm_version"`
	Name       types.String `tfsdk:"name"`
	PassPhrase types.String `tfsdk:"pass_phrase"`
	Status     types.String `tfsdk:"status"`
	Thalesluna types.List   `tfsdk:"thalesluna"`
}

var HsmThaleslunagroupAttrTypes = map[string]attr.Type{
	"ref":         types.StringType,
	"comment":     types.StringType,
	"group_sn":    types.StringType,
	"hsm_version": types.StringType,
	"name":        types.StringType,
	"pass_phrase": types.StringType,
	"status":      types.StringType,
	"thalesluna":  types.ListType{ElemType: types.ObjectType{AttrTypes: HsmThaleslunagroupThaleslunaAttrTypes}},
}

var HsmThaleslunagroupResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.LengthAtMost(256),
		},
		MarkdownDescription: "The HSM Thales Luna group comment.",
	},
	"group_sn": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The HSM Thales Luna group serial number.",
	},
	"hsm_version": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.OneOf("Luna_4", "Luna_5", "Luna_6", "Luna_7_CPL"),
		},
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
		MarkdownDescription: "The HSM Thales Luna version.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
			stringvalidator.LengthBetween(1, 32),
		},
		MarkdownDescription: "The HSM Thales Luna group name.",
	},
	"pass_phrase": schema.StringAttribute{
		Required:            true,
		Sensitive:           true,
		MarkdownDescription: "The pass phrase used to unlock the HSM Thales Luna keystore.",
	},
	"status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The status of all HSM Thales Luna devices in the group.",
	},
	"thalesluna": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: HsmThaleslunagroupThaleslunaResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Required:            true,
		MarkdownDescription: "The list of HSM Thales Luna devices.",
	},
}

func (m *HsmThaleslunagroupModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *security.HsmThaleslunagroup {
	if m == nil {
		return nil
	}
	to := &security.HsmThaleslunagroup{
		Comment:    flex.ExpandStringPointer(m.Comment),
		Name:       flex.ExpandStringPointer(m.Name),
		PassPhrase: flex.ExpandStringPointer(m.PassPhrase),
		Thalesluna: flex.ExpandFrameworkListNestedBlock(ctx, m.Thalesluna, diags, ExpandHsmThaleslunagroupThalesluna),
	}
	if isCreate {
		to.HsmVersion = flex.ExpandStringPointer(m.HsmVersion)
	}
	return to
}

func FlattenHsmThaleslunagroup(ctx context.Context, from *security.HsmThaleslunagroup, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(HsmThaleslunagroupAttrTypes)
	}
	m := HsmThaleslunagroupModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, HsmThaleslunagroupAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *HsmThaleslunagroupModel) Flatten(ctx context.Context, from *security.HsmThaleslunagroup, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = HsmThaleslunagroupModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.GroupSn = flex.FlattenStringPointer(from.GroupSn)
	m.HsmVersion = flex.FlattenStringPointer(from.HsmVersion)
	m.Name = flex.FlattenStringPointer(from.Name)
	//m.PassPhrase = flex.FlattenStringPointer(from.PassPhrase)
	m.Status = flex.FlattenStringPointer(from.Status)
	//m.Thalesluna = flex.FlattenFrameworkListNestedBlock(ctx, from.Thalesluna, HsmThaleslunagroupThaleslunaAttrTypes, diags, FlattenHsmThaleslunagroupThalesluna)

	// Flatten thalesluna and preserve ServerCertFilePath
	flattenedThalesluna := flex.FlattenFrameworkListNestedBlock(ctx, from.Thalesluna, HsmThaleslunagroupThaleslunaAttrTypes, diags, FlattenHsmThaleslunagroupThalesluna)
	m.Thalesluna = preserveThaleslunaServerCertFilePath(ctx, m.Thalesluna, flattenedThalesluna, diags)

}

func preserveThaleslunaServerCertFilePath(ctx context.Context,
	originalList types.List,
	flattenedList types.List,
	diags *diag.Diagnostics) types.List {

	if originalList.IsNull() || originalList.IsUnknown() ||
		flattenedList.IsNull() || flattenedList.IsUnknown() {
		return flattenedList
	}

	var originalModels []HsmThaleslunagroupThaleslunaModel
	diags.Append(originalList.ElementsAs(ctx, &originalModels, false)...)

	var flattenedModels []HsmThaleslunagroupThaleslunaModel
	diags.Append(flattenedList.ElementsAs(ctx, &flattenedModels, false)...)

	// Loop over ALL devices and restore ServerCertFilePath for each
	for i := range flattenedModels {
		if i < len(originalModels) {
			flattenedModels[i].ServerCertFilePath = originalModels[i].ServerCertFilePath
		}
	}

	updatedList, d := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: HsmThaleslunagroupThaleslunaAttrTypes}, flattenedModels)
	diags.Append(d...)
	return updatedList
}
