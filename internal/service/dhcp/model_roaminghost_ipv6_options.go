package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type RoaminghostIpv6OptionsModel struct {
	Name        types.String `tfsdk:"name"`
	Num         types.Int64  `tfsdk:"num"`
	VendorClass types.String `tfsdk:"vendor_class"`
	Value       types.String `tfsdk:"value"`
	UseOption   types.Bool   `tfsdk:"use_option"`
}

var RoaminghostIpv6OptionsAttrTypes = map[string]attr.Type{
	"name":         types.StringType,
	"num":          types.Int64Type,
	"vendor_class": types.StringType,
	"value":        types.StringType,
	"use_option":   types.BoolType,
}

var RoaminghostIpv6OptionsResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Name of the DHCP option.",
	},
	"num": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The code of the DHCP option.",
	},
	"vendor_class": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Default:  stringdefault.StaticString("DHCP"),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of the space this DHCP option is associated to.",
	},
	"value": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Value of the DHCP option",
	},
	"use_option": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Only applies to special options that are displayed separately from other options and have a use flag. These options are: * routers * router-templates * domain-name-servers * domain-name * broadcast-address * broadcast-address-offset * dhcp-lease-time * dhcp6.name-servers",
	},
}

func ExpandRoaminghostIpv6Options(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dhcp.RoaminghostIpv6Options {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RoaminghostIpv6OptionsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *RoaminghostIpv6OptionsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.RoaminghostIpv6Options {
	if m == nil {
		return nil
	}
	to := &dhcp.RoaminghostIpv6Options{
		Name:        flex.ExpandStringPointer(m.Name),
		Num:         flex.ExpandInt64Pointer(m.Num),
		VendorClass: flex.ExpandStringPointer(m.VendorClass),
		Value:       flex.ExpandStringPointer(m.Value),
		UseOption:   flex.ExpandBoolPointer(m.UseOption),
	}
	return to
}

func FlattenRoaminghostIpv6Options(ctx context.Context, from *dhcp.RoaminghostIpv6Options, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RoaminghostIpv6OptionsAttrTypes)
	}
	m := RoaminghostIpv6OptionsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RoaminghostIpv6OptionsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RoaminghostIpv6OptionsModel) Flatten(ctx context.Context, from *dhcp.RoaminghostIpv6Options, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RoaminghostIpv6OptionsModel{}
	}
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Num = flex.FlattenInt64Pointer(from.Num)
	m.VendorClass = flex.FlattenStringPointer(from.VendorClass)
	m.Value = flex.FlattenStringPointer(from.Value)
	m.UseOption = types.BoolPointerValue(from.UseOption)
}
