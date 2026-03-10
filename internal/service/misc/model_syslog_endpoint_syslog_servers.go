package misc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/misc"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type SyslogEndpointSyslogServersModel struct {
	Address             types.String `tfsdk:"address"`
	ConnectionType      types.String `tfsdk:"connection_type"`
	Port                types.Int64  `tfsdk:"port"`
	Hostname            types.String `tfsdk:"hostname"`
	Format              types.String `tfsdk:"format"`
	Facility            types.String `tfsdk:"facility"`
	Severity            types.String `tfsdk:"severity"`
	Certificate         types.Object `tfsdk:"certificate"`
	CertificateToken    types.String `tfsdk:"certificate_token"`
	CertificateFilePath types.String `tfsdk:"certificate_file_path"`
}

var SyslogEndpointSyslogServersAttrTypes = map[string]attr.Type{
	"address":               types.StringType,
	"connection_type":       types.StringType,
	"port":                  types.Int64Type,
	"hostname":              types.StringType,
	"format":                types.StringType,
	"facility":              types.StringType,
	"severity":              types.StringType,
	"certificate":           types.ObjectType{AttrTypes: SyslogEndpointSyslogServersCertificateAttrTypes},
	"certificate_token":     types.StringType,
	"certificate_file_path": types.StringType,
}

var SyslogEndpointSyslogServersResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Syslog Server IP address",
	},
	"connection_type": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Validators: []validator.String{
			stringvalidator.OneOf("stcp", "udp", "tls"),
		},
		Default:             stringdefault.StaticString("udp"),
		MarkdownDescription: "Connection type values",
	},
	"port": schema.Int64Attribute{
		Computed:            true,
		Optional:            true,
		Default:             int64default.StaticInt64(514),
		MarkdownDescription: "The port this server listens on.",
	},
	"hostname": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Validators: []validator.String{
			stringvalidator.OneOf("HOSTNAME", "FQDN", "IP_ADDRESS"),
		},
		Default:             stringdefault.StaticString("HOSTNAME"),
		MarkdownDescription: "List of hostnames",
	},
	"format": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Validators: []validator.String{
			stringvalidator.OneOf("formatted", "raw"),
		},
		Default:             stringdefault.StaticString("raw"),
		MarkdownDescription: "Format vlues for syslog endpoint server",
	},
	"facility": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Validators: []validator.String{
			stringvalidator.OneOf("local0", "local1", "local2", "local3", "local4", "local5", "local6", "local7"),
		},
		Default:             stringdefault.StaticString("local0"),
		MarkdownDescription: "Facility values for syslog endpoint server",
	},
	"severity": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Validators: []validator.String{
			stringvalidator.OneOf("alert", "critic", "debug", "emerg", "err", "info", "notice", "warning"),
		},
		Default:             stringdefault.StaticString("debug"),
		MarkdownDescription: "Severity values for syslog endpoint server.",
	},
	"certificate": schema.SingleNestedAttribute{
		Computed:            true,
		Attributes:          SyslogEndpointSyslogServersCertificateResourceSchemaAttributes,
		MarkdownDescription: "Reference for creating sysog endpoint server.",
	},
	"certificate_token": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The token returned by the uploadinit function call in object fileop.",
	},
	"certificate_file_path": schema.StringAttribute{
		Computed:            true,
		Optional:            true,
		MarkdownDescription: "The file path to the certificate.",
	},
}

func ExpandSyslogEndpointSyslogServers(ctx context.Context, o types.Object, diags *diag.Diagnostics) *misc.SyslogEndpointSyslogServers {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m SyslogEndpointSyslogServersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *SyslogEndpointSyslogServersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *misc.SyslogEndpointSyslogServers {
	if m == nil {
		return nil
	}
	to := &misc.SyslogEndpointSyslogServers{
		Address:          flex.ExpandStringPointer(m.Address),
		ConnectionType:   flex.ExpandStringPointer(m.ConnectionType),
		Port:             flex.ExpandInt64Pointer(m.Port),
		Hostname:         flex.ExpandStringPointer(m.Hostname),
		Format:           flex.ExpandStringPointer(m.Format),
		Facility:         flex.ExpandStringPointer(m.Facility),
		Severity:         flex.ExpandStringPointer(m.Severity),
		CertificateToken: flex.ExpandStringPointer(m.CertificateToken),
	}
	return to
}

func FlattenSyslogEndpointSyslogServers(ctx context.Context, from *misc.SyslogEndpointSyslogServers, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SyslogEndpointSyslogServersAttrTypes)
	}
	m := SyslogEndpointSyslogServersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, SyslogEndpointSyslogServersAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *SyslogEndpointSyslogServersModel) Flatten(ctx context.Context, from *misc.SyslogEndpointSyslogServers, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = SyslogEndpointSyslogServersModel{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.ConnectionType = flex.FlattenStringPointer(from.ConnectionType)
	m.Port = flex.FlattenInt64Pointer(from.Port)
	m.Hostname = flex.FlattenStringPointer(from.Hostname)
	m.Format = flex.FlattenStringPointer(from.Format)
	m.Facility = flex.FlattenStringPointer(from.Facility)
	m.Severity = flex.FlattenStringPointer(from.Severity)
	m.Certificate = flattenCertificate(ctx, from.Certificate, diags)
	m.CertificateToken = flex.FlattenStringPointer(from.CertificateToken)
	m.CertificateFilePath = types.StringNull() // This field is write-only and should not be set from API data
}

// flattenCertificate handles both string and object types returned by the API.
// When the API returns an object (for stcp connections), it extracts the certificate fields.
func flattenCertificate(ctx context.Context, cert interface{}, diags *diag.Diagnostics) types.Object {
	if cert == nil {
		return types.ObjectNull(SyslogEndpointSyslogServersCertificateAttrTypes)
	}

	switch v := cert.(type) {
	case string:
		// If certificate is a string, put it in the ref field
		m := SyslogEndpointSyslogServersCertificateModel{
			Ref: types.StringValue(v),
		}
		obj, d := types.ObjectValueFrom(ctx, SyslogEndpointSyslogServersCertificateAttrTypes, m)
		diags.Append(d...)
		return obj
	case map[string]interface{}:
		// API returns certificate as an object with _ref, issuer, serial, etc.
		m := SyslogEndpointSyslogServersCertificateModel{}
		if ref, ok := v["_ref"].(string); ok {
			m.Ref = types.StringValue(ref)
		}
		if issuer, ok := v["issuer"].(string); ok {
			m.Issuer = types.StringValue(issuer)
		}
		if serial, ok := v["serial"].(string); ok {
			m.Serial = types.StringValue(serial)
		}
		if subject, ok := v["subject"].(string); ok {
			m.Subject = types.StringValue(subject)
		}
		// Handle valid_not_after - can be float64 from JSON unmarshaling
		if validNotAfter, ok := v["valid_not_after"].(float64); ok {
			m.ValidNotAfter = types.Int64Value(int64(validNotAfter))
		} else if validNotAfter, ok := v["valid_not_after"].(int64); ok {
			m.ValidNotAfter = types.Int64Value(validNotAfter)
		}
		// Handle valid_not_before - can be float64 from JSON unmarshaling
		if validNotBefore, ok := v["valid_not_before"].(float64); ok {
			m.ValidNotBefore = types.Int64Value(int64(validNotBefore))
		} else if validNotBefore, ok := v["valid_not_before"].(int64); ok {
			m.ValidNotBefore = types.Int64Value(validNotBefore)
		}
		obj, d := types.ObjectValueFrom(ctx, SyslogEndpointSyslogServersCertificateAttrTypes, m)
		diags.Append(d...)
		return obj
	default:
		return types.ObjectNull(SyslogEndpointSyslogServersCertificateAttrTypes)
	}
}
