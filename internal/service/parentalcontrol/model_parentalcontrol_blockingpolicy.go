package parentalcontrol

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/parentalcontrol"
	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type ParentalcontrolBlockingpolicyModel struct {
	Ref   types.String `tfsdk:"ref"`
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

var ParentalcontrolBlockingpolicyAttrTypes = map[string]attr.Type{
	"ref":   types.StringType,
	"name":  types.StringType,
	"value": types.StringType,
}

var hex32Regex = regexp.MustCompile(`^[0-9a-fA-F]{32}$`)

var ParentalcontrolBlockingpolicyResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of the blocking policy.",
	},
	"value": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
			stringvalidator.RegexMatches(hex32Regex, "must be exactly 32 hexadecimal digits"),
		},
		MarkdownDescription: "The 32 bit hex value of the blocking policy.",
	},
}

func (m *ParentalcontrolBlockingpolicyModel) Expand(ctx context.Context, diags *diag.Diagnostics) *parentalcontrol.ParentalcontrolBlockingpolicy {
	if m == nil {
		return nil
	}
	to := &parentalcontrol.ParentalcontrolBlockingpolicy{
		Name:  flex.ExpandStringPointer(m.Name),
		Value: flex.ExpandStringPointer(m.Value),
	}
	return to
}

func FlattenParentalcontrolBlockingpolicy(ctx context.Context, from *parentalcontrol.ParentalcontrolBlockingpolicy, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ParentalcontrolBlockingpolicyAttrTypes)
	}
	m := ParentalcontrolBlockingpolicyModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ParentalcontrolBlockingpolicyAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ParentalcontrolBlockingpolicyModel) Flatten(ctx context.Context, from *parentalcontrol.ParentalcontrolBlockingpolicy, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ParentalcontrolBlockingpolicyModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Value = flex.FlattenStringPointer(from.Value)
}
