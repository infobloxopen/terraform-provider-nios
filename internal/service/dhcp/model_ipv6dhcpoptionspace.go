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

type Ipv6dhcpoptionspaceModel struct {
	Ref               types.String `tfsdk:"ref"`
    Uuid        types.String `tfsdk:"uuid"`
	Comment           types.String `tfsdk:"comment"`
	EnterpriseNumber  types.Int64  `tfsdk:"enterprise_number"`
	Name              types.String `tfsdk:"name"`
	OptionDefinitions types.List   `tfsdk:"option_definitions"`
}

var Ipv6dhcpoptionspaceAttrTypes = map[string]attr.Type{
	"ref":                types.StringType,
    "uuid":        types.StringType,
	"comment":            types.StringType,
	"enterprise_number":  types.Int64Type,
	"name":               types.StringType,
	"option_definitions": types.ListType{ElemType: types.StringType},
}

var Ipv6dhcpoptionspaceResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
    "uuid": schema.StringAttribute{
        Computed:            true,
        MarkdownDescription: "The uuid to the object.",
    },
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 256),
		},
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "A descriptive comment of a DHCP IPv6 option space object.",
	},
	"enterprise_number": schema.Int64Attribute{
		Required: true,
		Validators: []validator.Int64{
			int64validator.Between(0, 4294967295),
		},
		MarkdownDescription: "The enterprise number of a DHCP IPv6 option space object.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of a DHCP IPv6 option space object.",
	},
	"option_definitions": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The list of DHCP IPv6 option definition objects.",
	},
}

func (m *Ipv6dhcpoptionspaceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.Ipv6dhcpoptionspace {
	if m == nil {
		return nil
	}
	to := &dhcp.Ipv6dhcpoptionspace{
		Comment:          flex.ExpandStringPointer(m.Comment),
		EnterpriseNumber: flex.ExpandInt64Pointer(m.EnterpriseNumber),
		Name:             flex.ExpandStringPointer(m.Name),
	}
	return to
}

func FlattenIpv6dhcpoptionspace(ctx context.Context, from *dhcp.Ipv6dhcpoptionspace, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6dhcpoptionspaceAttrTypes)
	}
	m := Ipv6dhcpoptionspaceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6dhcpoptionspaceAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Ipv6dhcpoptionspaceModel) Flatten(ctx context.Context, from *dhcp.Ipv6dhcpoptionspace, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Ipv6dhcpoptionspaceModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
    m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.EnterpriseNumber = flex.FlattenInt64Pointer(from.EnterpriseNumber)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.OptionDefinitions = flex.FlattenFrameworkListString(ctx, from.OptionDefinitions, diags)
}
