package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dtc"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type DtcTopologyRuleDestinationLinkOneOfModel struct {
	Ref  types.String `tfsdk:"ref"`
	Host types.String `tfsdk:"host"`
	Name types.String `tfsdk:"name"`
}

var DtcTopologyRuleDestinationLinkOneOfAttrTypes = map[string]attr.Type{
	"ref":  types.StringType,
	"host": types.StringType,
	"name": types.StringType,
}

var DtcTopologyRuleDestinationLinkOneOfResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the LDAP auth service object.",
	},
	"host": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The host of server.",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the LDAP auth service object.",
	},
}

func ExpandDtcTopologyRuleDestinationLinkOneOf(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dtc.DtcTopologyRuleDestinationLinkOneOf {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DtcTopologyRuleDestinationLinkOneOfModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DtcTopologyRuleDestinationLinkOneOfModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dtc.DtcTopologyRuleDestinationLinkOneOf {
	if m == nil {
		return nil
	}
	to := &dtc.DtcTopologyRuleDestinationLinkOneOf{
		Ref:  flex.ExpandStringPointer(m.Ref),
		Host: flex.ExpandStringPointer(m.Host),
		Name: flex.ExpandStringPointer(m.Name),
	}
	return to
}

func FlattenDtcTopologyRuleDestinationLinkOneOf(ctx context.Context, from *dtc.DtcTopologyRuleDestinationLinkOneOf, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DtcTopologyRuleDestinationLinkOneOfAttrTypes)
	}
	m := DtcTopologyRuleDestinationLinkOneOfModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DtcTopologyRuleDestinationLinkOneOfAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DtcTopologyRuleDestinationLinkOneOfModel) Flatten(ctx context.Context, from *dtc.DtcTopologyRuleDestinationLinkOneOf, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DtcTopologyRuleDestinationLinkOneOfModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Host = flex.FlattenStringPointer(from.Host)
	m.Name = flex.FlattenStringPointer(from.Name)
}
