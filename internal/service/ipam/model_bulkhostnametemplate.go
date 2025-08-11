package ipam

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

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type BulkhostnametemplateModel struct {
	Ref            types.String `tfsdk:"ref"`
	IsGridDefault  types.Bool   `tfsdk:"is_grid_default"`
	PreDefined     types.Bool   `tfsdk:"pre_defined"`
	TemplateFormat types.String `tfsdk:"template_format"`
	TemplateName   types.String `tfsdk:"template_name"`
}

var BulkhostnametemplateAttrTypes = map[string]attr.Type{
	"ref":             types.StringType,
	"is_grid_default": types.BoolType,
	"pre_defined":     types.BoolType,
	"template_format": types.StringType,
	"template_name":   types.StringType,
}

var BulkhostnametemplateResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"is_grid_default": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "True if this template is Grid default.",
	},
	"pre_defined": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "True if this is a pre-defined template, False otherwise.",
	},
	"template_format": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`.*(\$4|#4).*`),
				"Template format must contain at least one of $4 or #4 placeholders",
			),
		},
		MarkdownDescription: "The format of bulk host name template. It should follow certain rules (please use Administration Guide as reference).",
	},
	"template_name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s]([^\s]|.*[^\s])?$`),
				"Should not have leading or trailing whitespace",
			),
		},
		MarkdownDescription: "The name of bulk host name template.",
	},
}

func ExpandBulkhostnametemplate(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.Bulkhostnametemplate {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m BulkhostnametemplateModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *BulkhostnametemplateModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.Bulkhostnametemplate {
	if m == nil {
		return nil
	}
	to := &ipam.Bulkhostnametemplate{
		Ref:            flex.ExpandStringPointer(m.Ref),
		TemplateFormat: flex.ExpandStringPointer(m.TemplateFormat),
		TemplateName:   flex.ExpandStringPointer(m.TemplateName),
	}
	return to
}

func FlattenBulkhostnametemplate(ctx context.Context, from *ipam.Bulkhostnametemplate, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(BulkhostnametemplateAttrTypes)
	}
	m := BulkhostnametemplateModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, BulkhostnametemplateAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *BulkhostnametemplateModel) Flatten(ctx context.Context, from *ipam.Bulkhostnametemplate, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = BulkhostnametemplateModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.IsGridDefault = types.BoolPointerValue(from.IsGridDefault)
	m.PreDefined = types.BoolPointerValue(from.PreDefined)
	m.TemplateFormat = flex.FlattenStringPointer(from.TemplateFormat)
	m.TemplateName = flex.FlattenStringPointer(from.TemplateName)
}
