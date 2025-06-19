package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dhcp"

	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/flex"
)

type FixedaddressMsOptionsModel struct {
	Num         types.Int64  `tfsdk:"num"`
	Value       types.String `tfsdk:"value"`
	Name        types.String `tfsdk:"name"`
	VendorClass types.String `tfsdk:"vendor_class"`
	UserClass   types.String `tfsdk:"user_class"`
	Type        types.String `tfsdk:"type"`
}

var FixedaddressMsOptionsAttrTypes = map[string]attr.Type{
	"num":          types.Int64Type,
	"value":        types.StringType,
	"name":         types.StringType,
	"vendor_class": types.StringType,
	"user_class":   types.StringType,
	"type":         types.StringType,
}

var FixedaddressMsOptionsResourceSchemaAttributes = map[string]schema.Attribute{
	"num": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The code of the DHCP option.",
	},
	"value": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Value of the DHCP option.",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the DHCP option.",
	},
	"vendor_class": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the vendor class with which this DHCP option is associated.",
	},
	"user_class": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the user class with which this DHCP option is associated.",
	},
	"type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The DHCP option type. Valid values are: * \"16-bit signed integer\" * \"16-bit unsigned integer\" * \"32-bit signed integer\" * \"32-bit unsigned integer\" * \"64-bit unsigned integer\" * \"8-bit signed integer\" * \"8-bit unsigned integer (1,2,4,8)\" * \"8-bit unsigned integer\" * \"array of 16-bit integer\" * \"array of 16-bit unsigned integer\" * \"array of 32-bit integer\" * \"array of 32-bit unsigned integer\" * \"array of 64-bit unsigned integer\" * \"array of 8-bit integer\" * \"array of 8-bit unsigned integer\" * \"array of ip-address pair\" * \"array of ip-address\" * \"array of string\" * \"binary\" * \"boolean array of ip-address\" * \"boolean\" * \"boolean-text\" * \"domain-list\" * \"domain-name\" * \"encapsulated\" * \"ip-address\" * \"string\" * \"text\"",
	},
}

func ExpandFixedaddressMsOptions(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dhcp.FixedaddressMsOptions {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m FixedaddressMsOptionsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *FixedaddressMsOptionsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.FixedaddressMsOptions {
	if m == nil {
		return nil
	}
	to := &dhcp.FixedaddressMsOptions{
		Num:         flex.ExpandInt64Pointer(m.Num),
		Value:       flex.ExpandStringPointer(m.Value),
		Name:        flex.ExpandStringPointer(m.Name),
		VendorClass: flex.ExpandStringPointer(m.VendorClass),
		UserClass:   flex.ExpandStringPointer(m.UserClass),
	}
	return to
}

func FlattenFixedaddressMsOptions(ctx context.Context, from *dhcp.FixedaddressMsOptions, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(FixedaddressMsOptionsAttrTypes)
	}
	m := FixedaddressMsOptionsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, FixedaddressMsOptionsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *FixedaddressMsOptionsModel) Flatten(ctx context.Context, from *dhcp.FixedaddressMsOptions, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = FixedaddressMsOptionsModel{}
	}
	m.Num = flex.FlattenInt64Pointer(from.Num)
	m.Value = flex.FlattenStringPointer(from.Value)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.VendorClass = flex.FlattenStringPointer(from.VendorClass)
	m.UserClass = flex.FlattenStringPointer(from.UserClass)
	m.Type = flex.FlattenStringPointer(from.Type)
}
