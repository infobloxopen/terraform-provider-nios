package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type HsmThaleslunagroupModel struct {
	Ref        types.String `tfsdk:"ref"`
    Uuid        types.String `tfsdk:"uuid"`
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
    "uuid":        types.StringType,
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
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
    "uuid": schema.StringAttribute{
        Computed:            true,
        MarkdownDescription: "The uuid to the object.",
    },
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The HSM Thales Luna group comment.",
	},
	"group_sn": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The HSM Thales Luna group serial number.",
	},
	"hsm_version": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The HSM Thales Luna version.",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The HSM Thales Luna group name.",
	},
	"pass_phrase": schema.StringAttribute{
		Optional:            true,
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
		Optional:            true,
		MarkdownDescription: "The list of HSM Thales Luna devices.",
	},
}

func ExpandHsmThaleslunagroup(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.HsmThaleslunagroup {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m HsmThaleslunagroupModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *HsmThaleslunagroupModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.HsmThaleslunagroup {
	if m == nil {
		return nil
	}
	to := &security.HsmThaleslunagroup{
		Ref:        flex.ExpandStringPointer(m.Ref),
		Comment:    flex.ExpandStringPointer(m.Comment),
		HsmVersion: flex.ExpandStringPointer(m.HsmVersion),
		Name:       flex.ExpandStringPointer(m.Name),
		PassPhrase: flex.ExpandStringPointer(m.PassPhrase),
		Thalesluna: flex.ExpandFrameworkListNestedBlock(ctx, m.Thalesluna, diags, ExpandHsmThaleslunagroupThalesluna),
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
    m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.GroupSn = flex.FlattenStringPointer(from.GroupSn)
	m.HsmVersion = flex.FlattenStringPointer(from.HsmVersion)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.PassPhrase = flex.FlattenStringPointer(from.PassPhrase)
	m.Status = flex.FlattenStringPointer(from.Status)
	m.Thalesluna = flex.FlattenFrameworkListNestedBlock(ctx, from.Thalesluna, HsmThaleslunagroupThaleslunaAttrTypes, diags, FlattenHsmThaleslunagroupThalesluna)
}
