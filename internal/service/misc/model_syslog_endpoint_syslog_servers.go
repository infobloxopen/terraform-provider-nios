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
	Certificate         types.String `tfsdk:"certificate"`
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
	"certificate":           types.StringType,
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
	"certificate": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Reference for creating sysog endpoint server.",
	},
	"certificate_token": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The token returned by the uploadinit function call in object fileop.",
	},
	"certificate_file_path": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The file path of the certificate to be uploaded. This is required when the certificate is uploaded using fileop.",
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
	m.Certificate = flex.FlattenStringPointer(from.Certificate)
	m.CertificateToken = flex.FlattenStringPointer(from.CertificateToken)
}
