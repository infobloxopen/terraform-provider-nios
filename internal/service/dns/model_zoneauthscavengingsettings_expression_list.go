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

type ZoneauthscavengingsettingsExpressionListModel struct {
	Op      types.String `tfsdk:"op"`
	Op1     types.String `tfsdk:"op1"`
	Op1Type types.String `tfsdk:"op1_type"`
	Op2     types.String `tfsdk:"op2"`
	Op2Type types.String `tfsdk:"op2_type"`
}

var ZoneauthscavengingsettingsExpressionListAttrTypes = map[string]attr.Type{
	"op":       types.StringType,
	"op1":      types.StringType,
	"op1_type": types.StringType,
	"op2":      types.StringType,
	"op2_type": types.StringType,
}

var ZoneauthscavengingsettingsExpressionListResourceSchemaAttributes = map[string]schema.Attribute{
	"op": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.OneOf(
				"AND", "ENDLIST", "EQ", "EXISTS", "GE", "GT", "LE", "LT", "MATCH_CIDR", "MATCH_IP", "MATCH_RANGE", "NOT_EQ", "NOT_EXISTS", "OR",
			),
		},
		MarkdownDescription: "The operation name.",
	},
	"op1": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The first operand value.",
	},
	"op1_type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("FIELD", "LIST", "STRING"),
		},
		MarkdownDescription: "The first operand type.",
	},
	"op2": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The second operand value.",
	},
	"op2_type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("FIELD", "LIST", "STRING"),
		},
		MarkdownDescription: "The second operand type.",
	},
}

func ExpandZoneauthscavengingsettingsExpressionList(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.ZoneauthscavengingsettingsExpressionList {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneauthscavengingsettingsExpressionListModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ZoneauthscavengingsettingsExpressionListModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.ZoneauthscavengingsettingsExpressionList {
	if m == nil {
		return nil
	}
	to := &dns.ZoneauthscavengingsettingsExpressionList{
		Op:      flex.ExpandStringPointer(m.Op),
		Op1:     flex.ExpandStringPointer(m.Op1),
		Op1Type: flex.ExpandStringPointer(m.Op1Type),
		Op2:     flex.ExpandStringPointer(m.Op2),
		Op2Type: flex.ExpandStringPointer(m.Op2Type),
	}
	return to
}

func FlattenZoneauthscavengingsettingsExpressionList(ctx context.Context, from *dns.ZoneauthscavengingsettingsExpressionList, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneauthscavengingsettingsExpressionListAttrTypes)
	}
	m := ZoneauthscavengingsettingsExpressionListModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZoneauthscavengingsettingsExpressionListAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ZoneauthscavengingsettingsExpressionListModel) Flatten(ctx context.Context, from *dns.ZoneauthscavengingsettingsExpressionList, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ZoneauthscavengingsettingsExpressionListModel{}
	}
	m.Op = flex.FlattenStringPointer(from.Op)
	m.Op1 = flex.FlattenStringPointer(from.Op1)
	m.Op1Type = flex.FlattenStringPointer(from.Op1Type)
	m.Op2 = flex.FlattenStringPointer(from.Op2)
	m.Op2Type = flex.FlattenStringPointer(from.Op2Type)
}
