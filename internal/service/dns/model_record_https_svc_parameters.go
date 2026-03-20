package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type RecordHttpsSvcParametersModel struct {
	SvcKey    types.String `tfsdk:"svc_key"`
	SvcValue  types.List   `tfsdk:"svc_value"`
	Mandatory types.Bool   `tfsdk:"mandatory"`
}

var RecordHttpsSvcParametersAttrTypes = map[string]attr.Type{
	"svc_key":   types.StringType,
	"svc_value": types.ListType{ElemType: types.StringType},
	"mandatory": types.BoolType,
}

var RecordHttpsSvcParametersResourceSchemaAttributes = map[string]schema.Attribute{
	"svc_key": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Svc param key.",
	},
	"svc_value": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Required:            true,
		MarkdownDescription: "Svc param value.",
	},
	"mandatory": schema.BoolAttribute{
		Required:            true,
		MarkdownDescription: "Specifies if this is mandatory key for this RR.",
	},
}

func ExpandRecordHttpsSvcParameters(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.RecordHttpsSvcParameters {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RecordHttpsSvcParametersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *RecordHttpsSvcParametersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.RecordHttpsSvcParameters {
	if m == nil {
		return nil
	}
	to := &dns.RecordHttpsSvcParameters{
		SvcKey:    flex.ExpandStringPointer(m.SvcKey),
		SvcValue:  flex.ExpandFrameworkListString(ctx, m.SvcValue, diags),
		Mandatory: flex.ExpandBoolPointer(m.Mandatory),
	}
	return to
}

func FlattenRecordHttpsSvcParameters(ctx context.Context, from *dns.RecordHttpsSvcParameters, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordHttpsSvcParametersAttrTypes)
	}
	m := RecordHttpsSvcParametersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RecordHttpsSvcParametersAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RecordHttpsSvcParametersModel) Flatten(ctx context.Context, from *dns.RecordHttpsSvcParameters, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RecordHttpsSvcParametersModel{}
	}
	m.SvcKey = flex.FlattenStringPointer(from.SvcKey)
	m.SvcValue = flex.FlattenFrameworkListString(ctx, from.SvcValue, diags)
	m.Mandatory = types.BoolPointerValue(from.Mandatory)
}
