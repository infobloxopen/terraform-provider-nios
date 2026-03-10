package misc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/misc"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type SyslogEndpointSyslogServersCertificateModel struct {
	Ref            types.String `tfsdk:"ref"`
	Issuer         types.String `tfsdk:"issuer"`
	Serial         types.String `tfsdk:"serial"`
	Subject        types.String `tfsdk:"subject"`
	ValidNotAfter  types.Int64  `tfsdk:"valid_not_after"`
	ValidNotBefore types.Int64  `tfsdk:"valid_not_before"`
}

var SyslogEndpointSyslogServersCertificateAttrTypes = map[string]attr.Type{
	"ref":              types.StringType,
	"issuer":           types.StringType,
	"serial":           types.StringType,
	"subject":          types.StringType,
	"valid_not_after":  types.Int64Type,
	"valid_not_before": types.Int64Type,
}

var SyslogEndpointSyslogServersCertificateResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the Syslog endpoint server certificate.",
	},
	"issuer": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The certificate issuer subject name.",
	},
	"serial": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The certificate serial number in hex format.",
	},
	"subject": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The certificate subject name.",
	},
	"valid_not_after": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The date after which the certificate becomes invalid.",
	},
	"valid_not_before": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The date before which the certificate is not valid.",
	},
}

func ExpandSyslogEndpointSyslogServersCertificate(ctx context.Context, o types.Object, diags *diag.Diagnostics) *misc.SyslogEndpointSyslogServersCertificate {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m SyslogEndpointSyslogServersCertificateModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *SyslogEndpointSyslogServersCertificateModel) Expand(ctx context.Context, diags *diag.Diagnostics) *misc.SyslogEndpointSyslogServersCertificate {
	if m == nil {
		return nil
	}
	to := &misc.SyslogEndpointSyslogServersCertificate{}
	return to
}

func FlattenSyslogEndpointSyslogServersCertificate(ctx context.Context, from *misc.SyslogEndpointSyslogServersCertificate, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SyslogEndpointSyslogServersCertificateAttrTypes)
	}
	m := SyslogEndpointSyslogServersCertificateModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, SyslogEndpointSyslogServersCertificateAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *SyslogEndpointSyslogServersCertificateModel) Flatten(ctx context.Context, from *misc.SyslogEndpointSyslogServersCertificate, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = SyslogEndpointSyslogServersCertificateModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Issuer = flex.FlattenStringPointer(from.Issuer)
	m.Serial = flex.FlattenStringPointer(from.Serial)
	m.Subject = flex.FlattenStringPointer(from.Subject)
	m.ValidNotAfter = flex.FlattenInt64Pointer(from.ValidNotAfter)
	m.ValidNotBefore = flex.FlattenInt64Pointer(from.ValidNotBefore)
}
