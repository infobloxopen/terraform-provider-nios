package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type DhcpoptiondefinitionModel struct {
	Ref   types.String `tfsdk:"ref"`
	Uuid  types.String `tfsdk:"uuid"`
	Code  types.Int64  `tfsdk:"code"`
	Name  types.String `tfsdk:"name"`
	Space types.String `tfsdk:"space"`
	Type  types.String `tfsdk:"type"`
}

var DhcpoptiondefinitionAttrTypes = map[string]attr.Type{
	"ref":   types.StringType,
	"uuid":  types.StringType,
	"code":  types.Int64Type,
	"name":  types.StringType,
	"space": types.StringType,
	"type":  types.StringType,
}

var DhcpoptiondefinitionResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The uuid to the object.",
	},
	"code": schema.Int64Attribute{
		Required: true,
		Validators: []validator.Int64{
			int64validator.Between(0, 254),
		},
		MarkdownDescription: "The code of a DHCP option definition object. An option code number is used to identify the DHCP option.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of a DHCP option definition object.",
	},
	"space": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString("DHCP"),
		MarkdownDescription: "The space of a DHCP option definition object.",
	},
	"type": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.OneOf(
				"16-bit signed integer", "16-bit unsigned integer", "32-bit signed integer", "32-bit unsigned integer",
				"64-bit unsigned integer", "8-bit signed integer", "8-bit unsigned integer",
				"array of 16-bit integer", "array of 16-bit unsigned integer", "array of 32-bit integer",
				"array of 32-bit unsigned integer", "array of 64-bit unsigned integer",
				"array of 8-bit integer", "array of 8-bit unsigned integer",
				"array of ip-address", "array of ip-address pair", "array of string",
				"binary", "boolean", "boolean array of ip-address", "boolean-text", "domain-list",
				"domain-name", "encapsulated", "ip-address", "string", "text",
			),
		},
		MarkdownDescription: "The data type of the Grid DHCP option.",
	},
}

func (m *DhcpoptiondefinitionModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.Dhcpoptiondefinition {
	if m == nil {
		return nil
	}
	to := &dhcp.Dhcpoptiondefinition{
		Code:  flex.ExpandInt64Pointer(m.Code),
		Name:  flex.ExpandStringPointer(m.Name),
		Space: flex.ExpandStringPointer(m.Space),
		Type:  flex.ExpandStringPointer(m.Type),
	}
	return to
}

func FlattenDhcpoptiondefinition(ctx context.Context, from *dhcp.Dhcpoptiondefinition, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DhcpoptiondefinitionAttrTypes)
	}
	m := DhcpoptiondefinitionModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DhcpoptiondefinitionAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DhcpoptiondefinitionModel) Flatten(ctx context.Context, from *dhcp.Dhcpoptiondefinition, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DhcpoptiondefinitionModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Code = flex.FlattenInt64Pointer(from.Code)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Space = flex.FlattenStringPointer(from.Space)
	m.Type = flex.FlattenStringPointer(from.Type)
}
