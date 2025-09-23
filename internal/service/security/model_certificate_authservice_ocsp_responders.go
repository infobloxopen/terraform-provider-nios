package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type CertificateAuthserviceOcspRespondersModel struct {
	FqdnOrIp            types.String `tfsdk:"fqdn_or_ip"`
	Port                types.Int64  `tfsdk:"port"`
	Comment             types.String `tfsdk:"comment"`
	Disabled            types.Bool   `tfsdk:"disabled"`
	Certificate         types.String `tfsdk:"certificate"`
	CertificateToken    types.String `tfsdk:"certificate_token"`
	CertificateFilePath types.String `tfsdk:"certificate_file_path"`
}

var CertificateAuthserviceOcspRespondersAttrTypes = map[string]attr.Type{
	"fqdn_or_ip":            types.StringType,
	"port":                  types.Int64Type,
	"comment":               types.StringType,
	"disabled":              types.BoolType,
	"certificate":           types.StringType,
	"certificate_token":     types.StringType,
	"certificate_file_path": types.StringType,
}

var CertificateAuthserviceOcspRespondersResourceSchemaAttributes = map[string]schema.Attribute{
	"fqdn_or_ip": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The FQDN (Fully Qualified Domain Name) or IP address of the server.",
	},
	"port": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(80),
		MarkdownDescription: "The port used for connecting.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The descriptive comment for the OCSP authentication responder.",
	},
	"disabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if this OCSP authentication responder is disabled.",
	},
	"certificate": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the OCSP responder certificate.",
	},
	"certificate_token": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The token returned by the uploadinit function call in object fileop.",
	},
	"certificate_file_path": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The file path to the certificate.",
	},
}

func ExpandCertificateAuthserviceOcspResponders(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.CertificateAuthserviceOcspResponders {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m CertificateAuthserviceOcspRespondersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *CertificateAuthserviceOcspRespondersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.CertificateAuthserviceOcspResponders {
	if m == nil {
		return nil
	}
	to := &security.CertificateAuthserviceOcspResponders{
		FqdnOrIp:         flex.ExpandStringPointer(m.FqdnOrIp),
		Port:             flex.ExpandInt64Pointer(m.Port),
		Comment:          flex.ExpandStringPointer(m.Comment),
		Disabled:         flex.ExpandBoolPointer(m.Disabled),
		CertificateToken: flex.ExpandStringPointer(m.CertificateToken),
	}
	return to
}

func FlattenCertificateAuthserviceOcspResponders(ctx context.Context, from *security.CertificateAuthserviceOcspResponders, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(CertificateAuthserviceOcspRespondersAttrTypes)
	}
	m := CertificateAuthserviceOcspRespondersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, CertificateAuthserviceOcspRespondersAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *CertificateAuthserviceOcspRespondersModel) Flatten(ctx context.Context, from *security.CertificateAuthserviceOcspResponders, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = CertificateAuthserviceOcspRespondersModel{}
	}
	m.FqdnOrIp = flex.FlattenStringPointer(from.FqdnOrIp)
	m.Port = flex.FlattenInt64Pointer(from.Port)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disabled = types.BoolPointerValue(from.Disabled)
	//m.Certificate = flex.FlattenStringPointer(from.Certificate)
	m.CertificateToken = flex.FlattenStringPointer(from.CertificateToken)
}
