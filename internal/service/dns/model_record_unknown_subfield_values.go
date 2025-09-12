package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type RecordUnknownSubfieldValuesModel struct {
	FieldValue    types.String `tfsdk:"field_value"`
	FieldType     types.String `tfsdk:"field_type"`
	IncludeLength types.String `tfsdk:"include_length"`
}

var RecordUnknownSubfieldValuesAttrTypes = map[string]attr.Type{
	"field_value":    types.StringType,
	"field_type":     types.StringType,
	"include_length": types.StringType,
}

var RecordUnknownSubfieldValuesResourceSchemaAttributes = map[string]schema.Attribute{
	"field_value": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "String representation of subfield value.",
	},
	"field_type": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.OneOf("B", "S", "I", "H", "6", "4", "N", "T", "X"),
		},
		MarkdownDescription: "Type of field. \"B\": unsigned 8-bit integer, \"S\": unsigned 16-bit integer, \"I\": unsigned 32-bit integer. \"H\": BASE64, \"6\": an IPv6 address, \"4\": an IPv4 address, \"N\": a domain name, \"T\": text string, \"X\": opaque binary data",
	},
	"include_length": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.OneOf("8_BIT", "16_BIT", "NONE"),
		},
		MarkdownDescription: "The 'size of 'length' sub-sub field to be included in RDATA.",
	},
}

func ExpandRecordUnknownSubfieldValues(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.RecordUnknownSubfieldValues {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RecordUnknownSubfieldValuesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *RecordUnknownSubfieldValuesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.RecordUnknownSubfieldValues {
	if m == nil {
		return nil
	}
	to := &dns.RecordUnknownSubfieldValues{
		FieldValue:    flex.ExpandStringPointer(m.FieldValue),
		FieldType:     flex.ExpandStringPointer(m.FieldType),
		IncludeLength: flex.ExpandStringPointer(m.IncludeLength),
	}
	return to
}

func FlattenRecordUnknownSubfieldValues(ctx context.Context, from *dns.RecordUnknownSubfieldValues, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordUnknownSubfieldValuesAttrTypes)
	}
	m := RecordUnknownSubfieldValuesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RecordUnknownSubfieldValuesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RecordUnknownSubfieldValuesModel) Flatten(ctx context.Context, from *dns.RecordUnknownSubfieldValues, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RecordUnknownSubfieldValuesModel{}
	}
	m.FieldValue = flex.FlattenStringPointer(from.FieldValue)
	m.FieldType = flex.FlattenStringPointer(from.FieldType)
	m.IncludeLength = flex.FlattenStringPointer(from.IncludeLength)
}
