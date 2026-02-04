package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
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

type DhcpoptionspaceModel struct {
	Ref               types.String `tfsdk:"ref"`
	Uuid              types.String `tfsdk:"uuid"`
	Comment           types.String `tfsdk:"comment"`
	Name              types.String `tfsdk:"name"`
	OptionDefinitions types.List   `tfsdk:"option_definitions"`
	SpaceType         types.String `tfsdk:"space_type"`
}

var DhcpoptionspaceAttrTypes = map[string]attr.Type{
	"ref":                types.StringType,
	"uuid":               types.StringType,
	"comment":            types.StringType,
	"name":               types.StringType,
	"option_definitions": types.ListType{ElemType: types.StringType},
	"space_type":         types.StringType,
}

var DhcpoptionspaceResourceSchemaAttributes = map[string]schema.Attribute{
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
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "A descriptive comment of a DHCP option space object.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of a DHCP option space object.",
	},
	"option_definitions": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Computed:            true,
		MarkdownDescription: "The list of DHCP option definition objects.",
	},
	"space_type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The type of a DHCP option space object.",
	},
}

func (m *DhcpoptionspaceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.Dhcpoptionspace {
	if m == nil {
		return nil
	}
	to := &dhcp.Dhcpoptionspace{
		Comment: flex.ExpandStringPointer(m.Comment),
		Name:    flex.ExpandStringPointer(m.Name),
	}
	return to
}

func FlattenDhcpoptionspace(ctx context.Context, from *dhcp.Dhcpoptionspace, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DhcpoptionspaceAttrTypes)
	}
	m := DhcpoptionspaceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DhcpoptionspaceAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DhcpoptionspaceModel) Flatten(ctx context.Context, from *dhcp.Dhcpoptionspace, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DhcpoptionspaceModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.OptionDefinitions = flex.FlattenFrameworkListString(ctx, from.OptionDefinitions, diags)
	m.SpaceType = flex.FlattenStringPointer(from.SpaceType)
}
