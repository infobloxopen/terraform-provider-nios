package dhcp

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type RangeExcludeModel struct {
	StartAddress types.String `tfsdk:"start_address"`
	EndAddress   types.String `tfsdk:"end_address"`
	Comment      types.String `tfsdk:"comment"`
}

var RangeExcludeAttrTypes = map[string]attr.Type{
	"start_address": types.StringType,
	"end_address":   types.StringType,
	"comment":       types.StringType,
}

var RangeExcludeResourceSchemaAttributes = map[string]schema.Attribute{
	"start_address": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The IPv4 Address starting address of the exclusion range.",
	},
	"end_address": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The IPv4 Address ending address of the exclusion range.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^\S.*\S$`),
				"should not have leading or trailing whitespace",
			),
		},
		MarkdownDescription: "Comment for the exclusion range; maximum 256 characters.",
	},
}

func ExpandRangeExclude(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dhcp.RangeExclude {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RangeExcludeModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *RangeExcludeModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.RangeExclude {
	if m == nil {
		return nil
	}
	to := &dhcp.RangeExclude{
		StartAddress: flex.ExpandStringPointer(m.StartAddress),
		EndAddress:   flex.ExpandStringPointer(m.EndAddress),
		Comment:      flex.ExpandStringPointer(m.Comment),
	}
	return to
}

func FlattenRangeExclude(ctx context.Context, from *dhcp.RangeExclude, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RangeExcludeAttrTypes)
	}
	m := RangeExcludeModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RangeExcludeAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RangeExcludeModel) Flatten(ctx context.Context, from *dhcp.RangeExclude, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RangeExcludeModel{}
	}
	m.StartAddress = flex.FlattenStringPointer(from.StartAddress)
	m.EndAddress = flex.FlattenStringPointer(from.EndAddress)
	m.Comment = flex.FlattenStringPointer(from.Comment)
}
