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

type RecordSvcbSvcParametersModel struct {
	SvcKey    types.String `tfsdk:"svc_key"`
	SvcValue  types.List   `tfsdk:"svc_value"`
	Mandatory types.Bool   `tfsdk:"mandatory"`
}

var RecordSvcbSvcParametersAttrTypes = map[string]attr.Type{
	"svc_key":   types.StringType,
	"svc_value": types.ListType{ElemType: types.StringType},
	"mandatory": types.BoolType,
}

var RecordSvcbSvcParametersResourceSchemaAttributes = map[string]schema.Attribute{
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

func ExpandRecordSvcbSvcParameters(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.RecordSvcbSvcParameters {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RecordSvcbSvcParametersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *RecordSvcbSvcParametersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.RecordSvcbSvcParameters {
	if m == nil {
		return nil
	}
	to := &dns.RecordSvcbSvcParameters{
		SvcKey:    flex.ExpandStringPointer(m.SvcKey),
		SvcValue:  flex.ExpandFrameworkListString(ctx, m.SvcValue, diags),
		Mandatory: flex.ExpandBoolPointer(m.Mandatory),
	}
	return to
}

func FlattenRecordSvcbSvcParameters(ctx context.Context, from *dns.RecordSvcbSvcParameters, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordSvcbSvcParametersAttrTypes)
	}
	m := RecordSvcbSvcParametersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RecordSvcbSvcParametersAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RecordSvcbSvcParametersModel) Flatten(ctx context.Context, from *dns.RecordSvcbSvcParameters, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RecordSvcbSvcParametersModel{}
	}
	m.SvcKey = flex.FlattenStringPointer(from.SvcKey)
	m.SvcValue = flex.FlattenFrameworkListString(ctx, from.SvcValue, diags)
	m.Mandatory = types.BoolPointerValue(from.Mandatory)
}
