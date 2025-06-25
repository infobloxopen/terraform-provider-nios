package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/ipam"

	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/flex"
)

type NetworkcontainerOptionsModel struct {
	Name        types.String `tfsdk:"name"`
	Num         types.Int64  `tfsdk:"num"`
	VendorClass types.String `tfsdk:"vendor_class"`
	Value       types.String `tfsdk:"value"`
	UseOption   types.Bool   `tfsdk:"use_option"`
}

var NetworkcontainerOptionsAttrTypes = map[string]attr.Type{
	"name":         types.StringType,
	"num":          types.Int64Type,
	"vendor_class": types.StringType,
	"value":        types.StringType,
	"use_option":   types.BoolType,
}

var NetworkcontainerOptionsResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Name of the DHCP option.",
		Computed:            true,
		Default:             stringdefault.StaticString("dhcp-lease-time"),
	},
	"num": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The code of the DHCP option.",
		Computed:            true,
		Default:             int64default.StaticInt64(51),
	},
	"vendor_class": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the space this DHCP option is associated to.",
		Computed:            true,
		Default:             stringdefault.StaticString("DHCP"),
	},
	"value": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Value of the DHCP option",
		Computed:            true,
		Default:             stringdefault.StaticString("43200"),
	},
	"use_option": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Only applies to special options that are displayed separately from other options and have a use flag. These options are: * routers * router-templates * domain-name-servers * domain-name * broadcast-address * broadcast-address-offset * dhcp-lease-time * dhcp6.name-servers",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
}

func ExpandNetworkcontainerOptions(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.NetworkcontainerOptions {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkcontainerOptionsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NetworkcontainerOptionsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.NetworkcontainerOptions {
	if m == nil {
		return nil
	}
	to := &ipam.NetworkcontainerOptions{
		Name:        flex.ExpandStringPointer(m.Name),
		Num:         flex.ExpandInt64Pointer(m.Num),
		VendorClass: flex.ExpandStringPointer(m.VendorClass),
		Value:       flex.ExpandStringPointer(m.Value),
		UseOption:   flex.ExpandBoolPointer(m.UseOption),
	}
	return to
}

func FlattenNetworkcontainerOptions(ctx context.Context, from *ipam.NetworkcontainerOptions, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkcontainerOptionsAttrTypes)
	}
	m := NetworkcontainerOptionsModel{}
	m.Flatten(ctx, from, diags)
	// m.ExtAttrs = m.ExtAttrsAll
	t, d := types.ObjectValueFrom(ctx, NetworkcontainerOptionsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NetworkcontainerOptionsModel) Flatten(ctx context.Context, from *ipam.NetworkcontainerOptions, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NetworkcontainerOptionsModel{}
	}
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Num = flex.FlattenInt64Pointer(from.Num)
	m.VendorClass = flex.FlattenStringPointer(from.VendorClass)
	m.Value = flex.FlattenStringPointer(from.Value)
	m.UseOption = types.BoolPointerValue(from.UseOption)
}
