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

type GridScheduledBackupModel struct {
	Status          types.String `tfsdk:"status"`
	Execute         types.String `tfsdk:"execute"`
	Operation       types.String `tfsdk:"operation"`
	BackupType      types.String `tfsdk:"backup_type"`
	KeepLocalCopy   types.Bool   `tfsdk:"keep_local_copy"`
	BackupFrequency types.String `tfsdk:"backup_frequency"`
	Weekday         types.String `tfsdk:"weekday"`
	HourOfDay       types.Int64  `tfsdk:"hour_of_day"`
	MinutesPastHour types.Int64  `tfsdk:"minutes_past_hour"`
	Username        types.String `tfsdk:"username"`
	Password        types.String `tfsdk:"password"`
	BackupServer    types.String `tfsdk:"backup_server"`
	Path            types.String `tfsdk:"path"`
	RestoreType     types.String `tfsdk:"restore_type"`
	RestoreServer   types.String `tfsdk:"restore_server"`
	RestoreUsername types.String `tfsdk:"restore_username"`
	RestorePassword types.String `tfsdk:"restore_password"`
	RestorePath     types.String `tfsdk:"restore_path"`
	NiosData        types.Bool   `tfsdk:"nios_data"`
	DiscoveryData   types.Bool   `tfsdk:"discovery_data"`
	SplunkAppData   types.Bool   `tfsdk:"splunk_app_data"`
	Enable          types.Bool   `tfsdk:"enable"`
	UseKeys         types.Bool   `tfsdk:"use_keys"`
	KeyType         types.String `tfsdk:"key_type"`
	UploadKeys      types.Bool   `tfsdk:"upload_keys"`
	DownloadKeys    types.Bool   `tfsdk:"download_keys"`
}

var GridScheduledBackupAttrTypes = map[string]attr.Type{
	"status":            types.StringType,
	"execute":           types.StringType,
	"operation":         types.StringType,
	"backup_type":       types.StringType,
	"keep_local_copy":   types.BoolType,
	"backup_frequency":  types.StringType,
	"weekday":           types.StringType,
	"hour_of_day":       types.Int64Type,
	"minutes_past_hour": types.Int64Type,
	"username":          types.StringType,
	"password":          types.StringType,
	"backup_server":     types.StringType,
	"path":              types.StringType,
	"restore_type":      types.StringType,
	"restore_server":    types.StringType,
	"restore_username":  types.StringType,
	"restore_password":  types.StringType,
	"restore_path":      types.StringType,
	"nios_data":         types.BoolType,
	"discovery_data":    types.BoolType,
	"splunk_app_data":   types.BoolType,
	"enable":            types.BoolType,
	"use_keys":          types.BoolType,
	"key_type":          types.StringType,
	"upload_keys":       types.BoolType,
	"download_keys":     types.BoolType,
}

var GridScheduledBackupResourceSchemaAttributes = map[string]schema.Attribute{
	"status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The status of the scheduled backup.",
	},
	"execute": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The state for scheduled backup or restore operation.",
	},
	"operation": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The scheduled backup operation.",
	},
	"backup_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The destination of the backup files.",
	},
	"keep_local_copy": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the local backup performed before uploading backup to remote storage.",
	},
	"backup_frequency": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The frequency of backups.",
	},
	"weekday": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The day of the week when the backup is performed.",
	},
	"hour_of_day": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The hour of the day past 12:00 AM the backup is performed.",
	},
	"minutes_past_hour": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The minute of the hour when the backup is performed.",
	},
	"username": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The user name on the backup server.",
	},
	"password": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The user password on the backup server.",
	},
	"backup_server": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The IP address of the backup server.",
	},
	"path": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The directory path to the backup file stored on the server.",
	},
	"restore_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The destination of the restore files.",
	},
	"restore_server": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The IP address of the restore server.",
	},
	"restore_username": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The user name on the restore server.",
	},
	"restore_password": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The password on the restore server.",
	},
	"restore_path": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The directory path to the restored file on the server.",
	},
	"nios_data": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the restore of the NIOS data is enabled.",
	},
	"discovery_data": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the restore the NetMRI data is enabled.",
	},
	"splunk_app_data": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the restore of the Splunk application data is enabled.",
	},
	"enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the scheduled backup is enabled.",
	},
	"use_keys": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If set, scp backup support based on keys",
	},
	"key_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "If set, scp backup support based on keys type",
	},
	"upload_keys": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If set, scp backup support to upload keys",
	},
	"download_keys": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If set, scp backup support to download keys",
	},
}

func ExpandGridScheduledBackup(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridScheduledBackup {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridScheduledBackupModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridScheduledBackupModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridScheduledBackup {
	if m == nil {
		return nil
	}
	to := &grid.GridScheduledBackup{
		Execute:         flex.ExpandStringPointer(m.Execute),
		Operation:       flex.ExpandStringPointer(m.Operation),
		BackupType:      flex.ExpandStringPointer(m.BackupType),
		KeepLocalCopy:   flex.ExpandBoolPointer(m.KeepLocalCopy),
		BackupFrequency: flex.ExpandStringPointer(m.BackupFrequency),
		Weekday:         flex.ExpandStringPointer(m.Weekday),
		HourOfDay:       flex.ExpandInt64Pointer(m.HourOfDay),
		MinutesPastHour: flex.ExpandInt64Pointer(m.MinutesPastHour),
		Username:        flex.ExpandStringPointer(m.Username),
		Password:        flex.ExpandStringPointer(m.Password),
		BackupServer:    flex.ExpandStringPointer(m.BackupServer),
		Path:            flex.ExpandStringPointer(m.Path),
		RestoreType:     flex.ExpandStringPointer(m.RestoreType),
		RestoreServer:   flex.ExpandStringPointer(m.RestoreServer),
		RestoreUsername: flex.ExpandStringPointer(m.RestoreUsername),
		RestorePassword: flex.ExpandStringPointer(m.RestorePassword),
		RestorePath:     flex.ExpandStringPointer(m.RestorePath),
		NiosData:        flex.ExpandBoolPointer(m.NiosData),
		DiscoveryData:   flex.ExpandBoolPointer(m.DiscoveryData),
		SplunkAppData:   flex.ExpandBoolPointer(m.SplunkAppData),
		Enable:          flex.ExpandBoolPointer(m.Enable),
		UseKeys:         flex.ExpandBoolPointer(m.UseKeys),
		KeyType:         flex.ExpandStringPointer(m.KeyType),
		UploadKeys:      flex.ExpandBoolPointer(m.UploadKeys),
		DownloadKeys:    flex.ExpandBoolPointer(m.DownloadKeys),
	}
	return to
}

func FlattenGridScheduledBackup(ctx context.Context, from *grid.GridScheduledBackup, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridScheduledBackupAttrTypes)
	}
	m := GridScheduledBackupModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridScheduledBackupAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridScheduledBackupModel) Flatten(ctx context.Context, from *grid.GridScheduledBackup, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridScheduledBackupModel{}
	}
	m.Status = flex.FlattenStringPointer(from.Status)
	m.Execute = flex.FlattenStringPointer(from.Execute)
	m.Operation = flex.FlattenStringPointer(from.Operation)
	m.BackupType = flex.FlattenStringPointer(from.BackupType)
	m.KeepLocalCopy = types.BoolPointerValue(from.KeepLocalCopy)
	m.BackupFrequency = flex.FlattenStringPointer(from.BackupFrequency)
	m.Weekday = flex.FlattenStringPointer(from.Weekday)
	m.HourOfDay = flex.FlattenInt64Pointer(from.HourOfDay)
	m.MinutesPastHour = flex.FlattenInt64Pointer(from.MinutesPastHour)
	m.Username = flex.FlattenStringPointer(from.Username)
	m.Password = flex.FlattenStringPointer(from.Password)
	m.BackupServer = flex.FlattenStringPointer(from.BackupServer)
	m.Path = flex.FlattenStringPointer(from.Path)
	m.RestoreType = flex.FlattenStringPointer(from.RestoreType)
	m.RestoreServer = flex.FlattenStringPointer(from.RestoreServer)
	m.RestoreUsername = flex.FlattenStringPointer(from.RestoreUsername)
	m.RestorePassword = flex.FlattenStringPointer(from.RestorePassword)
	m.RestorePath = flex.FlattenStringPointer(from.RestorePath)
	m.NiosData = types.BoolPointerValue(from.NiosData)
	m.DiscoveryData = types.BoolPointerValue(from.DiscoveryData)
	m.SplunkAppData = types.BoolPointerValue(from.SplunkAppData)
	m.Enable = types.BoolPointerValue(from.Enable)
	m.UseKeys = types.BoolPointerValue(from.UseKeys)
	m.KeyType = flex.FlattenStringPointer(from.KeyType)
	m.UploadKeys = types.BoolPointerValue(from.UploadKeys)
	m.DownloadKeys = types.BoolPointerValue(from.DownloadKeys)
}
