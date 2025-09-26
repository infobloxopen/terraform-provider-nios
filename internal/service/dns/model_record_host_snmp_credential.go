package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type RecordHostSnmpCredentialModel struct {
	CommunityString types.String `tfsdk:"community_string"`
	Comment         types.String `tfsdk:"comment"`
	CredentialGroup types.String `tfsdk:"credential_group"`
}

var RecordHostSnmpCredentialAttrTypes = map[string]attr.Type{
	"community_string": types.StringType,
	"comment":          types.StringType,
	"credential_group": types.StringType,
}

var RecordHostSnmpCredentialResourceSchemaAttributes = map[string]schema.Attribute{
	"community_string": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The public community string.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Comments for the SNMPv1 and SNMPv2 users.",
	},
	"credential_group": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Group for the SNMPv1 and SNMPv2 credential.",
	},
}

func ExpandRecordHostSnmpCredential(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.RecordHostSnmpCredential {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RecordHostSnmpCredentialModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *RecordHostSnmpCredentialModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.RecordHostSnmpCredential {
	if m == nil {
		return nil
	}
	to := &dns.RecordHostSnmpCredential{
		CommunityString: flex.ExpandStringPointer(m.CommunityString),
		Comment:         flex.ExpandStringPointer(m.Comment),
		CredentialGroup: flex.ExpandStringPointer(m.CredentialGroup),
	}
	return to
}

func FlattenRecordHostSnmpCredential(ctx context.Context, from *dns.RecordHostSnmpCredential, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordHostSnmpCredentialAttrTypes)
	}
	m := RecordHostSnmpCredentialModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RecordHostSnmpCredentialAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RecordHostSnmpCredentialModel) Flatten(ctx context.Context, from *dns.RecordHostSnmpCredential, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RecordHostSnmpCredentialModel{}
	}
	m.CommunityString = flex.FlattenStringPointer(from.CommunityString)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CredentialGroup = flex.FlattenStringPointer(from.CredentialGroup)
}
