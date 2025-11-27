package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	planmodifiers "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/immutable"
	importmod "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/import"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type VlanrangeModel struct {
	Ref            types.String `tfsdk:"ref"`
	Comment        types.String `tfsdk:"comment"`
	DeleteVlans    types.Bool   `tfsdk:"delete_vlans"`
	EndVlanId      types.Int64  `tfsdk:"end_vlan_id"`
	ExtAttrs       types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll    types.Map    `tfsdk:"extattrs_all"`
	Name           types.String `tfsdk:"name"`
	PreCreateVlan  types.Bool   `tfsdk:"pre_create_vlan"`
	StartVlanId    types.Int64  `tfsdk:"start_vlan_id"`
	VlanNamePrefix types.String `tfsdk:"vlan_name_prefix"`
	VlanView       types.String `tfsdk:"vlan_view"`
}

var VlanrangeAttrTypes = map[string]attr.Type{
	"ref":              types.StringType,
	"comment":          types.StringType,
	"delete_vlans":     types.BoolType,
	"end_vlan_id":      types.Int64Type,
	"extattrs":         types.MapType{ElemType: types.StringType},
	"extattrs_all":     types.MapType{ElemType: types.StringType},
	"name":             types.StringType,
	"pre_create_vlan":  types.BoolType,
	"start_vlan_id":    types.Int64Type,
	"vlan_name_prefix": types.StringType,
	"vlan_view":        types.StringType,
}

var VlanrangeResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"comment": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 256),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "A descriptive comment for this VLAN Range.",
	},
	// DeleteVlans can only be set during delete operation, so it is computed here.
	"delete_vlans": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Vlans delete option. Determines whether all child objects should be removed alongside with the VLAN Range or child objects should be assigned to another parental VLAN Range/View. By default child objects are re-parented.",
	},
	"end_vlan_id": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "End ID for VLAN Range.",
	},
	"extattrs": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Extensible attributes associated with the object.",
	},
	"extattrs_all": schema.MapAttribute{
		Computed: true,
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
		MarkdownDescription: "Extensible attributes associated with the object , including default and internal attributes.",
		ElementType:         types.StringType,
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Name of the VLAN Range.",
	},
	"pre_create_vlan": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.Bool{
			planmodifiers.ImmutableBool(),
		},
		MarkdownDescription: "If set on creation VLAN objects will be created once VLAN Range created.",
	},
	"start_vlan_id": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "Start ID for VLAN Range.",
	},
	"vlan_name_prefix": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
		MarkdownDescription: "If set on creation prefix string will be used for VLAN name.",
	},
	"vlan_view": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The VLAN View to which this VLAN Range belongs.",
	},
}

func (m *VlanrangeModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *ipam.Vlanrange {
	if m == nil {
		return nil
	}
	to := &ipam.Vlanrange{
		Comment:     flex.ExpandStringPointer(m.Comment),
		EndVlanId:   flex.ExpandInt64Pointer(m.EndVlanId),
		ExtAttrs:    ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		Name:        flex.ExpandStringPointer(m.Name),
		StartVlanId: flex.ExpandInt64Pointer(m.StartVlanId),
		VlanView:    ExpandVlanView(m.VlanView),
	}
	if isCreate {
		to.PreCreateVlan = flex.ExpandBoolPointer(m.PreCreateVlan)
		to.VlanNamePrefix = flex.ExpandStringPointer(m.VlanNamePrefix)
	}
	return to
}

func FlattenVlanrange(ctx context.Context, from *ipam.Vlanrange, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(VlanrangeAttrTypes)
	}
	m := VlanrangeModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, VlanrangeAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *VlanrangeModel) Flatten(ctx context.Context, from *ipam.Vlanrange, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = VlanrangeModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.DeleteVlans = types.BoolPointerValue(from.DeleteVlans)
	m.EndVlanId = flex.FlattenInt64Pointer(from.EndVlanId)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.Name = flex.FlattenStringPointer(from.Name)
	if m.PreCreateVlan.IsUnknown() || m.PreCreateVlan.IsNull() {
		m.PreCreateVlan = types.BoolPointerValue(from.PreCreateVlan)
	}
	m.StartVlanId = flex.FlattenInt64Pointer(from.StartVlanId)
	if m.VlanNamePrefix.IsUnknown() || m.VlanNamePrefix.IsNull() {
		m.VlanNamePrefix = flex.FlattenStringPointer(from.VlanNamePrefix)
	}
	m.VlanView = FlattenVlanView(from.VlanView)
}

func ExpandVlanView(str types.String) *ipam.VlanrangeVlanView {
	if str.IsNull() {
		return &ipam.VlanrangeVlanView{}
	}
	var m ipam.VlanrangeVlanView
	m.String = flex.ExpandStringPointer(str)

	return &m
}

func FlattenVlanView(from *ipam.VlanrangeVlanView) types.String {
	if from.VlanrangeVlanViewOneOf == nil {
		return types.StringNull()
	}
	m := flex.FlattenStringPointer(from.VlanrangeVlanViewOneOf.Ref)
	return m
}
