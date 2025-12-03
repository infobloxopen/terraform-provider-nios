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

type HsmEntrustnshieldgroupModel struct {
	Ref               types.String `tfsdk:"ref"`
	CardName          types.String `tfsdk:"card_name"`
	Comment           types.String `tfsdk:"comment"`
	EntrustnshieldHsm types.List   `tfsdk:"entrustnshield_hsm"`
	KeyServerIp       types.String `tfsdk:"key_server_ip"`
	KeyServerPort     types.Int64  `tfsdk:"key_server_port"`
	Name              types.String `tfsdk:"name"`
	PassPhrase        types.String `tfsdk:"pass_phrase"`
	Protection        types.String `tfsdk:"protection"`
	Status            types.String `tfsdk:"status"`
}

var HsmEntrustnshieldgroupAttrTypes = map[string]attr.Type{
	"ref":                types.StringType,
	"card_name":          types.StringType,
	"comment":            types.StringType,
	"entrustnshield_hsm": types.ListType{ElemType: types.ObjectType{AttrTypes: HsmEntrustnshieldgroupEntrustnshieldHsmAttrTypes}},
	"key_server_ip":      types.StringType,
	"key_server_port":    types.Int64Type,
	"name":               types.StringType,
	"pass_phrase":        types.StringType,
	"protection":         types.StringType,
	"status":             types.StringType,
}

var HsmEntrustnshieldgroupResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"card_name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The Entrust nShield HSM softcard name.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The Entrust nShield HSM group comment.",
	},
	"entrustnshield_hsm": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: HsmEntrustnshieldgroupEntrustnshieldHsmResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of Entrust nShield HSM devices.",
	},
	"key_server_ip": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The remote file server (RFS) IPv4 Address.",
	},
	"key_server_port": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The remote file server (RFS) port.",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The Entrust nShield HSM group name.",
	},
	"pass_phrase": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The password phrase used to unlock the Entrust nShield HSM keystore.",
	},
	"protection": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The level of protection that the HSM group uses for the DNSSEC key data.",
	},
	"status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The status of all Entrust nShield HSM devices in the group.",
	},
}

func ExpandHsmEntrustnshieldgroup(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.HsmEntrustnshieldgroup {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m HsmEntrustnshieldgroupModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *HsmEntrustnshieldgroupModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.HsmEntrustnshieldgroup {
	if m == nil {
		return nil
	}
	to := &security.HsmEntrustnshieldgroup{
		Ref:               flex.ExpandStringPointer(m.Ref),
		CardName:          flex.ExpandStringPointer(m.CardName),
		Comment:           flex.ExpandStringPointer(m.Comment),
		EntrustnshieldHsm: flex.ExpandFrameworkListNestedBlock(ctx, m.EntrustnshieldHsm, diags, ExpandHsmEntrustnshieldgroupEntrustnshieldHsm),
		KeyServerIp:       flex.ExpandStringPointer(m.KeyServerIp),
		KeyServerPort:     flex.ExpandInt64Pointer(m.KeyServerPort),
		Name:              flex.ExpandStringPointer(m.Name),
		PassPhrase:        flex.ExpandStringPointer(m.PassPhrase),
		Protection:        flex.ExpandStringPointer(m.Protection),
	}
	return to
}

func FlattenHsmEntrustnshieldgroup(ctx context.Context, from *security.HsmEntrustnshieldgroup, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(HsmEntrustnshieldgroupAttrTypes)
	}
	m := HsmEntrustnshieldgroupModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, HsmEntrustnshieldgroupAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *HsmEntrustnshieldgroupModel) Flatten(ctx context.Context, from *security.HsmEntrustnshieldgroup, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = HsmEntrustnshieldgroupModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.CardName = flex.FlattenStringPointer(from.CardName)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.EntrustnshieldHsm = flex.FlattenFrameworkListNestedBlock(ctx, from.EntrustnshieldHsm, HsmEntrustnshieldgroupEntrustnshieldHsmAttrTypes, diags, FlattenHsmEntrustnshieldgroupEntrustnshieldHsm)
	m.KeyServerIp = flex.FlattenStringPointer(from.KeyServerIp)
	m.KeyServerPort = flex.FlattenInt64Pointer(from.KeyServerPort)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.PassPhrase = flex.FlattenStringPointer(from.PassPhrase)
	m.Protection = flex.FlattenStringPointer(from.Protection)
	m.Status = flex.FlattenStringPointer(from.Status)
}
