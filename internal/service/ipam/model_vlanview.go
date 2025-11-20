package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type VlanviewModel struct {
	Ref                   types.String `tfsdk:"ref"`
	AllowRangeOverlapping types.Bool   `tfsdk:"allow_range_overlapping"`
	Comment               types.String `tfsdk:"comment"`
	EndVlanId             types.Int64  `tfsdk:"end_vlan_id"`
	ExtAttrs              types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll           types.Map    `tfsdk:"extattrs_all"`
	Name                  types.String `tfsdk:"name"`
	PreCreateVlan         types.Bool   `tfsdk:"pre_create_vlan"`
	StartVlanId           types.Int64  `tfsdk:"start_vlan_id"`
	VlanNamePrefix        types.String `tfsdk:"vlan_name_prefix"`
}

var VlanviewAttrTypes = map[string]attr.Type{
	"ref":                     types.StringType,
	"allow_range_overlapping": types.BoolType,
	"comment":                 types.StringType,
	"end_vlan_id":             types.Int64Type,
	"extattrs":                types.MapType{ElemType: types.StringType},
	"extattrs_all":            types.MapType{ElemType: types.StringType},
	"name":                    types.StringType,
	"pre_create_vlan":         types.BoolType,
	"start_vlan_id":           types.Int64Type,
	"vlan_name_prefix":        types.StringType,
}

var VlanviewResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"allow_range_overlapping": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "When set to true VLAN Ranges under VLAN View can have overlapping ID.",
	},
	"comment": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
			stringvalidator.LengthBetween(0, 256),
		},
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "A descriptive comment for this VLAN View.",
	},
	"end_vlan_id": schema.Int64Attribute{
		Required: true,
		Validators: []validator.Int64{
			int64validator.Between(1, 4094),
		},
		MarkdownDescription: "End ID for VLAN View.",
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
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object , including default and internal attributes.",
		ElementType:         types.StringType,
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Name of the VLAN View.",
	},
	"pre_create_vlan": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If set on creation VLAN objects will be created once VLAN View created.",
	},
	"start_vlan_id": schema.Int64Attribute{
		Required: true,
		Validators: []validator.Int64{
			int64validator.Between(1, 4094),
		},
		MarkdownDescription: "Start ID for VLAN View.",
	},
	"vlan_name_prefix": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "If set on creation prefix string will be used for VLAN name.",
	},
}

func (m *VlanviewModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *ipam.Vlanview {
	if m == nil {
		return nil
	}
	to := &ipam.Vlanview{
		AllowRangeOverlapping: flex.ExpandBoolPointer(m.AllowRangeOverlapping),
		Comment:               flex.ExpandStringPointer(m.Comment),
		EndVlanId:             flex.ExpandInt64Pointer(m.EndVlanId),
		ExtAttrs:              ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		Name:                  flex.ExpandStringPointer(m.Name),
		StartVlanId:           flex.ExpandInt64Pointer(m.StartVlanId),
	}
	if isCreate {
		to.PreCreateVlan = flex.ExpandBoolPointer(m.PreCreateVlan)
		to.VlanNamePrefix = flex.ExpandStringPointer(m.VlanNamePrefix)
	}
	return to
}

func FlattenVlanview(ctx context.Context, from *ipam.Vlanview, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(VlanviewAttrTypes)
	}
	m := VlanviewModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, VlanviewAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *VlanviewModel) Flatten(ctx context.Context, from *ipam.Vlanview, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = VlanviewModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.AllowRangeOverlapping = types.BoolPointerValue(from.AllowRangeOverlapping)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.EndVlanId = flex.FlattenInt64Pointer(from.EndVlanId)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.PreCreateVlan = types.BoolPointerValue(from.PreCreateVlan)
	m.StartVlanId = flex.FlattenInt64Pointer(from.StartVlanId)
	m.VlanNamePrefix = flex.FlattenStringPointer(from.VlanNamePrefix)
}
