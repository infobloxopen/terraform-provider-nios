package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	// "github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dtc"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type DtcTopologyRuleDestinationLinkModel struct {
	Ref  types.String `tfsdk:"ref"`
	Host types.String `tfsdk:"host"`
	Name types.String `tfsdk:"name"`
}

var DtcTopologyRuleDestinationLinkAttrTypes = map[string]attr.Type{
	"ref":  types.StringType,
	"host": types.StringType,
	"name": types.StringType,
}

var DtcTopologyRuleDestinationLinkResourceSchemaAttributes = map[string]schema.Attribute{
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

func ExpandDtcTopologyRuleDestinationLink(ctx context.Context, o types.String, diags *diag.Diagnostics) *dtc.DtcTopologyRuleDestinationLink {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	return &dtc.DtcTopologyRuleDestinationLink{
		String: flex.ExpandStringPointer(o),
	}
}

// func (m *DtcTopologyRuleDestinationLinkModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dtc.DtcTopologyRuleDestinationLink {
// 	if m == nil {
// 		return nil
// 	}
// 	to := &dtc.DtcTopologyRuleDestinationLink{
// 		Ref:  flex.ExpandStringPointer(m.Ref),
// 		Host: flex.ExpandStringPointer(m.Host),
// 		Name: flex.ExpandStringPointer(m.Name),
// 	}
// 	return to
// }

func FlattenDtcTopologyRuleDestinationLink(ctx context.Context, from *dtc.DtcTopologyRuleDestinationLink, diags *diag.Diagnostics) types.String {
	if from == nil {
		return types.StringNull()
	}
	return flex.FlattenStringPointer(from.DtcTopologyRuleDestinationLinkOneOf.Ref)
}

// func (m *DtcTopologyRuleDestinationLinkModel) Flatten(ctx context.Context, from *dtc.DtcTopologyRuleDestinationLink, diags *diag.Diagnostics) {
// 	if from == nil {
// 		return
// 	}
// 	if m == nil {
// 		*m = DtcTopologyRuleDestinationLinkModel{}
// 	}
// 	m.Ref = flex.FlattenStringPointer(from.Ref)
// 	m.Host = flex.FlattenStringPointer(from.Host)
// 	m.Name = flex.FlattenStringPointer(from.Name)
// }
