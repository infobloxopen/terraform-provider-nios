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

type RecordHostSnmp3CredentialModel struct {
	User                   types.String `tfsdk:"user"`
	AuthenticationProtocol types.String `tfsdk:"authentication_protocol"`
	AuthenticationPassword types.String `tfsdk:"authentication_password"`
	PrivacyProtocol        types.String `tfsdk:"privacy_protocol"`
	PrivacyPassword        types.String `tfsdk:"privacy_password"`
	Comment                types.String `tfsdk:"comment"`
	CredentialGroup        types.String `tfsdk:"credential_group"`
}

var RecordHostSnmp3CredentialAttrTypes = map[string]attr.Type{
	"user":                    types.StringType,
	"authentication_protocol": types.StringType,
	"authentication_password": types.StringType,
	"privacy_protocol":        types.StringType,
	"privacy_password":        types.StringType,
	"comment":                 types.StringType,
	"credential_group":        types.StringType,
}

var RecordHostSnmp3CredentialResourceSchemaAttributes = map[string]schema.Attribute{
	"user": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The SNMPv3 user name.",
	},
	"authentication_protocol": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Authentication protocol for the SNMPv3 user.",
	},
	"authentication_password": schema.StringAttribute{
		Optional:            true,
		WriteOnly:           true,
		MarkdownDescription: "Authentication password for the SNMPv3 user.",
	},
	"privacy_protocol": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Privacy protocol for the SNMPv3 user.",
	},
	"privacy_password": schema.StringAttribute{
		Optional:            true,
		WriteOnly:           true,
		MarkdownDescription: "Privacy password for the SNMPv3 user.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Comments for the SNMPv3 user.",
	},
	"credential_group": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Group for the SNMPv3 credential.",
	},
}

func ExpandRecordHostSnmp3Credential(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.RecordHostSnmp3Credential {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RecordHostSnmp3CredentialModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *RecordHostSnmp3CredentialModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.RecordHostSnmp3Credential {
	if m == nil {
		return nil
	}
	to := &dns.RecordHostSnmp3Credential{
		User:                   flex.ExpandStringPointer(m.User),
		AuthenticationProtocol: flex.ExpandStringPointer(m.AuthenticationProtocol),
		AuthenticationPassword: flex.ExpandStringPointer(m.AuthenticationPassword),
		PrivacyProtocol:        flex.ExpandStringPointer(m.PrivacyProtocol),
		PrivacyPassword:        flex.ExpandStringPointer(m.PrivacyPassword),
		Comment:                flex.ExpandStringPointer(m.Comment),
		CredentialGroup:        flex.ExpandStringPointer(m.CredentialGroup),
	}
	return to
}

func FlattenRecordHostSnmp3Credential(ctx context.Context, from *dns.RecordHostSnmp3Credential, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordHostSnmp3CredentialAttrTypes)
	}
	m := RecordHostSnmp3CredentialModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RecordHostSnmp3CredentialAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RecordHostSnmp3CredentialModel) Flatten(ctx context.Context, from *dns.RecordHostSnmp3Credential, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RecordHostSnmp3CredentialModel{}
	}
	m.User = flex.FlattenStringPointer(from.User)
	m.AuthenticationProtocol = flex.FlattenStringPointer(from.AuthenticationProtocol)
	m.AuthenticationPassword = flex.FlattenStringPointer(from.AuthenticationPassword)
	m.PrivacyProtocol = flex.FlattenStringPointer(from.PrivacyProtocol)
	m.PrivacyPassword = flex.FlattenStringPointer(from.PrivacyPassword)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CredentialGroup = flex.FlattenStringPointer(from.CredentialGroup)
}
