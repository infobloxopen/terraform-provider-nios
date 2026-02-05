package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type NatgroupModel struct {
	Ref     types.String `tfsdk:"ref"`
	Uuid    types.String `tfsdk:"uuid"`
	Comment types.String `tfsdk:"comment"`
	Name    types.String `tfsdk:"name"`
}

var NatgroupAttrTypes = map[string]attr.Type{
	"ref":     types.StringType,
	"uuid":    types.StringType,
	"comment": types.StringType,
	"name":    types.StringType,
}

var NatgroupResourceSchemaAttributes = map[string]schema.Attribute{
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
			customvalidator.ValidateTrimmedString(),
		},
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The NAT group descriptive comment; maximum 256 characters.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of a NAT group object.",
	},
}

func (m *NatgroupModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.Natgroup {
	if m == nil {
		return nil
	}
	to := &grid.Natgroup{
		Comment: flex.ExpandStringPointer(m.Comment),
		Name:    flex.ExpandStringPointer(m.Name),
	}
	return to
}

func FlattenNatgroup(ctx context.Context, from *grid.Natgroup, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NatgroupAttrTypes)
	}
	m := NatgroupModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NatgroupAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NatgroupModel) Flatten(ctx context.Context, from *grid.Natgroup, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NatgroupModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Name = flex.FlattenStringPointer(from.Name)
}
