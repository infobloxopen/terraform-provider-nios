package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type MemberExternalSyslogBackupServersModel struct {
	AddressOrFqdn types.String `tfsdk:"address_or_fqdn"`
	DirectoryPath types.String `tfsdk:"directory_path"`
	Enable        types.Bool   `tfsdk:"enable"`
	Password      types.String `tfsdk:"password"`
	Port          types.Int64  `tfsdk:"port"`
	Protocol      types.String `tfsdk:"protocol"`
	Username      types.String `tfsdk:"username"`
}

var MemberExternalSyslogBackupServersAttrTypes = map[string]attr.Type{
	"address_or_fqdn": types.StringType,
	"directory_path":  types.StringType,
	"enable":          types.BoolType,
	"password":        types.StringType,
	"port":            types.Int64Type,
	"protocol":        types.StringType,
	"username":        types.StringType,
}

var MemberExternalSyslogBackupServersResourceSchemaAttributes = map[string]schema.Attribute{
	"address_or_fqdn": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The IPv4 or IPv6 address or FQDN of the backup syslog server.",
	},
	"directory_path": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The directory path for the replication of the rotated syslog files.",
	},
	"enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If set to True, the syslog backup server is enabled.",
	},
	"password": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The password of the backup syslog server.",
	},
	"port": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The port used to connect to the backup syslog server.",
	},
	"protocol": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The transport protocol used to connect to the backup syslog server.",
	},
	"username": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The username of the backup syslog server.",
	},
}

func ExpandMemberExternalSyslogBackupServers(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberExternalSyslogBackupServers {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberExternalSyslogBackupServersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberExternalSyslogBackupServersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberExternalSyslogBackupServers {
	if m == nil {
		return nil
	}
	to := &grid.MemberExternalSyslogBackupServers{
		AddressOrFqdn: flex.ExpandStringPointer(m.AddressOrFqdn),
		DirectoryPath: flex.ExpandStringPointer(m.DirectoryPath),
		Enable:        flex.ExpandBoolPointer(m.Enable),
		Password:      flex.ExpandStringPointer(m.Password),
		Port:          flex.ExpandInt64Pointer(m.Port),
		Protocol:      flex.ExpandStringPointer(m.Protocol),
		Username:      flex.ExpandStringPointer(m.Username),
	}
	return to
}

func FlattenMemberExternalSyslogBackupServers(ctx context.Context, from *grid.MemberExternalSyslogBackupServers, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberExternalSyslogBackupServersAttrTypes)
	}
	m := MemberExternalSyslogBackupServersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberExternalSyslogBackupServersAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberExternalSyslogBackupServersModel) Flatten(ctx context.Context, from *grid.MemberExternalSyslogBackupServers, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberExternalSyslogBackupServersModel{}
	}
	m.AddressOrFqdn = flex.FlattenStringPointer(from.AddressOrFqdn)
	m.DirectoryPath = flex.FlattenStringPointer(from.DirectoryPath)
	m.Enable = types.BoolPointerValue(from.Enable)
	m.Password = flex.FlattenStringPointer(from.Password)
	m.Port = flex.FlattenInt64Pointer(from.Port)
	m.Protocol = flex.FlattenStringPointer(from.Protocol)
	m.Username = flex.FlattenStringPointer(from.Username)
}
