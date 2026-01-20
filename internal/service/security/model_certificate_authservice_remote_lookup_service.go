package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type CertificateAuthserviceRemoteLookupServiceModel struct {
	Ref  types.String `tfsdk:"ref"`
	Uuid types.String `tfsdk:"uuid"`
}

var CertificateAuthserviceRemoteLookupServiceAttrTypes = map[string]attr.Type{
	"ref":  types.StringType,
	"uuid": types.StringType,
}

var CertificateAuthserviceRemoteLookupServiceResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the LDAP auth service object.",
	},
}

func ExpandCertificateAuthserviceRemoteLookupService(ctx context.Context, s types.String, diags *diag.Diagnostics) *security.CertificateAuthserviceRemoteLookupService {
	if s.IsNull() || s.IsUnknown() {
		return nil
	}

	stringPtr := flex.ExpandStringPointer(s)
	if stringPtr == nil {
		return nil
	}

	return &security.CertificateAuthserviceRemoteLookupService{
		String: stringPtr,
	}
}

func (m *CertificateAuthserviceRemoteLookupServiceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.CertificateAuthserviceRemoteLookupService {
	if m == nil {
		return nil
	}
	to := &security.CertificateAuthserviceRemoteLookupService{
		String: flex.ExpandStringPointer(m.Ref),
	}
	return to
}

func FlattenCertificateAuthserviceRemoteLookupService(ctx context.Context, from *security.CertificateAuthserviceRemoteLookupService, diags *diag.Diagnostics) types.String {
	if from == nil {
		return types.StringNull()
	}
	if from.CertificateAuthserviceRemoteLookupServiceOneOf == nil || from.CertificateAuthserviceRemoteLookupServiceOneOf.Ref == nil {
		return types.StringNull()
	}
	t := from.CertificateAuthserviceRemoteLookupServiceOneOf.Ref
	return flex.FlattenStringPointer(t)
}

func (m *CertificateAuthserviceRemoteLookupServiceModel) Flatten(ctx context.Context, from *security.CertificateAuthserviceRemoteLookupService, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = CertificateAuthserviceRemoteLookupServiceModel{}
	}
	// Check if OneOf structure exists
	if from.CertificateAuthserviceRemoteLookupServiceOneOf == nil || from.CertificateAuthserviceRemoteLookupServiceOneOf.Ref == nil {
		m.Ref = types.StringNull()
		return
	}

	// Extract the Ref from the OneOf structure
	m.Ref = flex.FlattenStringPointer(from.CertificateAuthserviceRemoteLookupServiceOneOf.Ref)
	
}
