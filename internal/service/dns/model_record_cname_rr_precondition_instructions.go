package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type RecordCnameRrPreconditionInstructionsModel struct {
	Condition types.String `tfsdk:"condition"`
	Name      types.String `tfsdk:"name"`
	Type      types.String `tfsdk:"type"`
	Rdata     types.String `tfsdk:"rdata"`
	Action    types.String `tfsdk:"action"`
}

var RecordCnameRrPreconditionInstructionsAttrTypes = map[string]attr.Type{
	"condition": types.StringType,
	"name":      types.StringType,
	"type":      types.StringType,
	"rdata":     types.StringType,
	"action":    types.StringType,
}

var RecordCnameRrPreconditionInstructionsResourceSchemaAttributes = map[string]schema.Attribute{
	"condition": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.OneOf("exist", "notexist"),
		},
		MarkdownDescription: "Condition type: exist or notexist.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Record name",
	},
	"type": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.OneOf("A", "AAAA"),
		},
		MarkdownDescription: "Record type",
	},
	"rdata": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "Record data (optional)",
	},
	"action": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.OneOf("none", "delete"),
		},
		MarkdownDescription: "Action to perform if condition is met: none or delete.",
	},
}

func ExpandRecordCnameRrPreconditionInstructions(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.RecordCnameRrPreconditionInstructions {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RecordCnameRrPreconditionInstructionsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *RecordCnameRrPreconditionInstructionsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.RecordCnameRrPreconditionInstructions {
	if m == nil {
		return nil
	}
	to := &dns.RecordCnameRrPreconditionInstructions{
		Condition: flex.ExpandStringPointer(m.Condition),
		Name:      flex.ExpandStringPointer(m.Name),
		Type:      flex.ExpandStringPointer(m.Type),
		Rdata:     flex.ExpandStringPointer(m.Rdata),
		Action:    flex.ExpandStringPointer(m.Action),
	}
	return to
}

func FlattenRecordCnameRrPreconditionInstructions(ctx context.Context, from *dns.RecordCnameRrPreconditionInstructions, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordCnameRrPreconditionInstructionsAttrTypes)
	}
	m := RecordCnameRrPreconditionInstructionsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RecordCnameRrPreconditionInstructionsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RecordCnameRrPreconditionInstructionsModel) Flatten(ctx context.Context, from *dns.RecordCnameRrPreconditionInstructions, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RecordCnameRrPreconditionInstructionsModel{}
	}
	m.Condition = flex.FlattenStringPointer(from.Condition)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Type = flex.FlattenStringPointer(from.Type)
	m.Rdata = flex.FlattenStringPointer(from.Rdata)
	m.Action = flex.FlattenStringPointer(from.Action)
}
