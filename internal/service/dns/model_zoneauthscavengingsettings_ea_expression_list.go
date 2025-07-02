package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dns"

	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/flex"
)

type ZoneauthscavengingsettingsEaExpressionListModel struct {
	Op      types.String `tfsdk:"op"`
	Op1     types.String `tfsdk:"op1"`
	Op1Type types.String `tfsdk:"op1_type"`
	Op2     types.String `tfsdk:"op2"`
	Op2Type types.String `tfsdk:"op2_type"`
}

var ZoneauthscavengingsettingsEaExpressionListAttrTypes = map[string]attr.Type{
	"op":       types.StringType,
	"op1":      types.StringType,
	"op1_type": types.StringType,
	"op2":      types.StringType,
	"op2_type": types.StringType,
}

var ZoneauthscavengingsettingsEaExpressionListResourceSchemaAttributes = map[string]schema.Attribute{
	"op": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The operation name.",
	},
	"op1": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the Extensible Attribute Definition object which is used as the first operand value.",
	},
	"op1_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The first operand type.",
	},
	"op2": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The second operand value.",
	},
	"op2_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The second operand type.",
	},
}

func ExpandZoneauthscavengingsettingsEaExpressionList(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.ZoneauthscavengingsettingsEaExpressionList {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneauthscavengingsettingsEaExpressionListModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ZoneauthscavengingsettingsEaExpressionListModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.ZoneauthscavengingsettingsEaExpressionList {
	if m == nil {
		return nil
	}
	to := &dns.ZoneauthscavengingsettingsEaExpressionList{
		Op:      flex.ExpandStringPointer(m.Op),
		Op1:     flex.ExpandStringPointer(m.Op1),
		Op1Type: flex.ExpandStringPointer(m.Op1Type),
		Op2:     flex.ExpandStringPointer(m.Op2),
		Op2Type: flex.ExpandStringPointer(m.Op2Type),
	}
	return to
}

func FlattenZoneauthscavengingsettingsEaExpressionList(ctx context.Context, from *dns.ZoneauthscavengingsettingsEaExpressionList, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneauthscavengingsettingsEaExpressionListAttrTypes)
	}
	m := ZoneauthscavengingsettingsEaExpressionListModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZoneauthscavengingsettingsEaExpressionListAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ZoneauthscavengingsettingsEaExpressionListModel) Flatten(ctx context.Context, from *dns.ZoneauthscavengingsettingsEaExpressionList, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ZoneauthscavengingsettingsEaExpressionListModel{}
	}
	m.Op = flex.FlattenStringPointer(from.Op)
	m.Op1 = flex.FlattenStringPointer(from.Op1)
	m.Op1Type = flex.FlattenStringPointer(from.Op1Type)
	m.Op2 = flex.FlattenStringPointer(from.Op2)
	m.Op2Type = flex.FlattenStringPointer(from.Op2Type)
}
