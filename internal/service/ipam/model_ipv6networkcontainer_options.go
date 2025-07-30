package ipam

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type Ipv6networkcontainerOptionsModel struct {
	Name        types.String `tfsdk:"name"`
	Num         types.Int64  `tfsdk:"num"`
	VendorClass types.String `tfsdk:"vendor_class"`
	Value       types.String `tfsdk:"value"`
	UseOption   types.Bool   `tfsdk:"use_option"`
}

var Ipv6networkcontainerOptionsAttrTypes = map[string]attr.Type{
	"name":         types.StringType,
	"num":          types.Int64Type,
	"vendor_class": types.StringType,
	"value":        types.StringType,
	"use_option":   types.BoolType,
}

var Ipv6networkcontainerOptionsResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Name of the DHCP option.",
		Computed:            true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Should not have leading or trailing whitespace",
			),
		},
	},
	"num": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The code of the DHCP option.",
		Computed:            true,
	},
	"vendor_class": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the space this DHCP option is associated to.",
		Computed:            true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Should not have leading or trailing whitespace",
			),
		},
		Default: stringdefault.StaticString("DHCP"),
	},
	"value": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Value of the DHCP option",
		Computed:            true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Should not have leading or trailing whitespace",
			),
		},
	},
	"use_option": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Only applies to special options that are displayed separately from other options and have a use flag. These options are: * routers * router-templates * domain-name-servers * domain-name * broadcast-address * broadcast-address-offset * dhcp-lease-time * dhcp6.name-servers",
		Computed:            true,
	},
}

func ExpandIpv6networkcontainerOptions(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.Ipv6networkcontainerOptions {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6networkcontainerOptionsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *Ipv6networkcontainerOptionsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.Ipv6networkcontainerOptions {
	if m == nil {
		return nil
	}
	to := &ipam.Ipv6networkcontainerOptions{
		Name:        flex.ExpandStringPointer(m.Name),
		Num:         flex.ExpandInt64Pointer(m.Num),
		VendorClass: flex.ExpandStringPointer(m.VendorClass),
		Value:       flex.ExpandStringPointer(m.Value),
		UseOption:   flex.ExpandBoolPointer(m.UseOption),
	}
	return to
}

func FlattenIpv6networkcontainerOptions(ctx context.Context, from *ipam.Ipv6networkcontainerOptions, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6networkcontainerOptionsAttrTypes)
	}
	m := Ipv6networkcontainerOptionsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6networkcontainerOptionsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Ipv6networkcontainerOptionsModel) Flatten(ctx context.Context, from *ipam.Ipv6networkcontainerOptions, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Ipv6networkcontainerOptionsModel{}
	}
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Num = flex.FlattenInt64Pointer(from.Num)
	m.VendorClass = flex.FlattenStringPointer(from.VendorClass)
	m.Value = flex.FlattenStringPointer(from.Value)
	m.UseOption = types.BoolPointerValue(from.UseOption)
}
