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

type GridFiledistributionModel struct {
	Ref                types.String `tfsdk:"ref"`
	Uuid               types.String `tfsdk:"uuid"`
	AllowUploads       types.Bool   `tfsdk:"allow_uploads"`
	BackupStorage      types.Bool   `tfsdk:"backup_storage"`
	CurrentUsage       types.Int64  `tfsdk:"current_usage"`
	EnableAnonymousFtp types.Bool   `tfsdk:"enable_anonymous_ftp"`
	GlobalStatus       types.String `tfsdk:"global_status"`
	Name               types.String `tfsdk:"name"`
	StorageLimit       types.Int64  `tfsdk:"storage_limit"`
}

var GridFiledistributionAttrTypes = map[string]attr.Type{
	"ref":                  types.StringType,
	"uuid":                 types.StringType,
	"allow_uploads":        types.BoolType,
	"backup_storage":       types.BoolType,
	"current_usage":        types.Int64Type,
	"enable_anonymous_ftp": types.BoolType,
	"global_status":        types.StringType,
	"name":                 types.StringType,
	"storage_limit":        types.Int64Type,
}

var GridFiledistributionResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The UUID of the object.",
	},
	"allow_uploads": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the uploads to Grid members are allowed.",
	},
	"backup_storage": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether to include distributed files in the backup.",
	},
	"current_usage": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The value is the percentage of the allocated TFTP storage space that is used, expressed in tenth of a percent. Valid values are from 0 to 1000.",
	},
	"enable_anonymous_ftp": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the FTP anonymous login is enabled.",
	},
	"global_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The Grid file distribution global status.",
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The Grid name.",
	},
	"storage_limit": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Maximum storage in megabytes allowed on the Grid. The maximum storage space allowed for all file distribution services on a Grid is equal to the storage space allowed to the Grid member with the smallest amount of space allowed.",
	},
}

func ExpandGridFiledistribution(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridFiledistribution {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridFiledistributionModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridFiledistributionModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridFiledistribution {
	if m == nil {
		return nil
	}
	to := &grid.GridFiledistribution{
		Ref:                flex.ExpandStringPointer(m.Ref),
		Uuid:               flex.ExpandStringPointer(m.Uuid),
		AllowUploads:       flex.ExpandBoolPointer(m.AllowUploads),
		BackupStorage:      flex.ExpandBoolPointer(m.BackupStorage),
		EnableAnonymousFtp: flex.ExpandBoolPointer(m.EnableAnonymousFtp),
		StorageLimit:       flex.ExpandInt64Pointer(m.StorageLimit),
	}
	return to
}

func FlattenGridFiledistribution(ctx context.Context, from *grid.GridFiledistribution, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridFiledistributionAttrTypes)
	}
	m := GridFiledistributionModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridFiledistributionAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridFiledistributionModel) Flatten(ctx context.Context, from *grid.GridFiledistribution, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridFiledistributionModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.AllowUploads = types.BoolPointerValue(from.AllowUploads)
	m.BackupStorage = types.BoolPointerValue(from.BackupStorage)
	m.CurrentUsage = flex.FlattenInt64Pointer(from.CurrentUsage)
	m.EnableAnonymousFtp = types.BoolPointerValue(from.EnableAnonymousFtp)
	m.GlobalStatus = flex.FlattenStringPointer(from.GlobalStatus)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.StorageLimit = flex.FlattenInt64Pointer(from.StorageLimit)
}
